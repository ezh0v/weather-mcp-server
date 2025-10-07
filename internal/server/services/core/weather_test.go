package core

import (
	"context"
	"errors"
	"html/template"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/ezh0v/weather-mcp-server/internal/server/services/mock"
	"github.com/ezh0v/weather-mcp-server/pkg/weatherapi/models"
)

func TestCurrentWeather(t *testing.T) {
	testCases := map[string]struct {
		city            string
		errString       string
		wait            string
		setupWeatherAPI func(weatherAPI *mock.MockWeatherAPIProvider)
	}{
		"city_not_found": {
			city:      "Tokyo",
			errString: "weather API not available. Code: 400",
			setupWeatherAPI: func(weatherAPI *mock.MockWeatherAPIProvider) {
				weatherAPI.EXPECT().
					Current(context.Background(), "Tokyo").
					Return(nil, errors.New("weather API not available. Code: 400"))
			},
		},
		"successful_result": {
			city: "London",
			wait: "London, United Kingdom Sunny 18 45 4 " +
				"https://cdn.weatherapi.com/weather/64x64/day/113.png",
			setupWeatherAPI: func(weatherAPI *mock.MockWeatherAPIProvider) {
				weatherAPI.EXPECT().
					Current(context.Background(), "London").
					Return(&models.CurrentResponse{
						Location: models.Location{
							Name:    "London",
							Country: "United Kingdom",
						},
						Current: models.Current{
							TempC:    18.4,
							WindKph:  4.2,
							Humidity: 45,
							Condition: models.Condition{
								Text: "Sunny",
								Icon: "//cdn.weatherapi.com/weather/64x64/day/113.png",
							},
						},
					}, nil)
			},
		},
	}

	renderer, err := template.New("weather.html").Parse(
		"{{ .Location }} {{ .Condition }} {{ .Temperature }} " +
			"{{ .Humidity }} {{ .WindSpeed }} {{ .Icon }}")
	require.NoError(t, err)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	weatherAPI := mock.NewMockWeatherAPIProvider(ctrl)

	svc := New(renderer, weatherAPI)

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			if tc.setupWeatherAPI != nil {
				tc.setupWeatherAPI(weatherAPI)
			}

			data, err := svc.Weather().Current(context.Background(), tc.city)
			if err != nil {
				assert.EqualError(t, err, tc.errString)
			}

			assert.Equal(t, tc.wait, data)
		})
	}
}
