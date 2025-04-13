package handlers

import (
	"context"
	"errors"
	"testing"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	mocks "github.com/TuanKiri/weather-mcp-server/internal/server/services/mock"
)

func TestCurrentWeather(t *testing.T) {
	testCases := map[string]struct {
		fillWeatherService func(mocksWeather *mocks.MockWeatherService)
		arguments          map[string]any
		errString          string
		wait               string
	}{
		"empty_city": {
			wait: "city must be a string",
		},
		"city_not_found": {
			fillWeatherService: func(mocksWeather *mocks.MockWeatherService) {
				mocksWeather.EXPECT().
					Current(context.Background(), "Tokyo").
					Return("", errors.New("weather API not available. Code: 400"))
			},
			arguments: map[string]any{
				"city": "Tokyo",
			},
			errString: "weather API not available. Code: 400",
		},
		"successful_request": {
			fillWeatherService: func(mocksWeather *mocks.MockWeatherService) {
				mocksWeather.EXPECT().
					Current(context.Background(), "London").
					Return("<h1>London weather data</h1>", nil)
			},
			arguments: map[string]any{
				"city": "London",
			},
			wait: "<h1>London weather data</h1>",
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocksWeather := mocks.NewMockWeatherService(ctrl)

	svc := mocks.NewMockServices(ctrl)
	svc.EXPECT().Weather().Return(mocksWeather).AnyTimes()

	handler := CurrentWeather(svc)

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			if tc.fillWeatherService != nil {
				tc.fillWeatherService(mocksWeather)
			}

			var request mcp.CallToolRequest
			request.Params.Arguments = tc.arguments

			result, err := handler(context.Background(), request)
			if err != nil {
				assert.EqualError(t, err, tc.errString)
				return
			}

			require.Len(t, result.Content, 1)
			content, ok := result.Content[0].(mcp.TextContent)
			require.True(t, ok)

			assert.Equal(t, tc.wait, content.Text)
		})
	}
}
