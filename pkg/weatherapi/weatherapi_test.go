package weatherapi

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ezh0v/weather-mcp-server/pkg/weatherapi/models"
)

func TestCurrentWeather(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		city      string
		errString string
		wait      *models.CurrentResponse
	}{
		"successful_request": {
			city: "London",
			wait: &models.CurrentResponse{
				Location: models.Location{
					Name:    "London",
					Country: "United Kingdom",
				},
				Current: models.Current{
					TempC:    18.4,
					WindKph:  4,
					Humidity: 45,
					Condition: models.Condition{
						Text: "Sunny",
						Icon: "//cdn.weatherapi.com/weather/64x64/day/113.png",
					},
				},
			},
		},
		"bad_request": {
			errString: "weather API not available. Code: 400",
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query().Get("q")
		if q == "" {
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		path := filepath.Join("mock", "current.json")

		data, err := os.ReadFile(path)
		if err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}))
	defer server.Close()

	weatherAPI := &WeatherAPI{
		key:     "test-key",
		baseURL: server.URL,
		client:  server.Client(),
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			result, err := weatherAPI.Current(context.Background(), tc.city)
			if err != nil {
				assert.EqualError(t, err, tc.errString)
			}

			assert.Equal(t, tc.wait, result)
		})
	}
}
