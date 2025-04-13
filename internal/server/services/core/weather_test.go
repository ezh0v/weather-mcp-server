package core

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/TuanKiri/weather-mcp-server/internal/server/services/mock"
	viewModels "github.com/TuanKiri/weather-mcp-server/internal/server/view/models"
	"github.com/TuanKiri/weather-mcp-server/pkg/weatherapi/models"
)

func TestCurrentWeather(t *testing.T) {
	testCases := map[string]struct {
		city            string
		errString       string
		wait            string
		setupWeatherAPI func(weatherAPI *mock.MockWeatherAPIProvider)
		setupRenderer   func(renderer *mock.MockTemplateRenderer)
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
			wait: "{Location:London, United Kingdom " +
				"Icon:https://cdn.weatherapi.com/weather/64x64/day/113.png " +
				"Condition:Sunny " +
				"Temperature:18 " +
				"Humidity:45 " +
				"WindSpeed:4}",
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
							WindKph:  4,
							Humidity: 45,
							Condition: models.Condition{
								Text: "Sunny",
								Icon: "//cdn.weatherapi.com/weather/64x64/day/113.png",
							},
						},
					}, nil)
			},
			setupRenderer: func(renderer *mock.MockTemplateRenderer) {
				renderer.EXPECT().
					ExecuteTemplate(
						gomock.AssignableToTypeOf(&bytes.Buffer{}),
						"weather.html",
						gomock.AssignableToTypeOf(viewModels.CurrentWeather{}),
					).
					Do(func(wr io.Writer, _ string, data any) error {
						value, _ := data.(viewModels.CurrentWeather)
						wr.Write(fmt.Appendf([]byte{}, "%+v", value))
						return nil
					})
			},
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	renderer := mock.NewMockTemplateRenderer(ctrl)
	weatherAPI := mock.NewMockWeatherAPIProvider(ctrl)

	svc := New(renderer, weatherAPI)

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			if tc.setupWeatherAPI != nil {
				tc.setupWeatherAPI(weatherAPI)
			}

			if tc.setupRenderer != nil {
				tc.setupRenderer(renderer)
			}

			data, err := svc.Weather().Current(context.Background(), tc.city)
			if err != nil {
				assert.EqualError(t, err, tc.errString)
			}

			assert.Equal(t, tc.wait, data)
		})
	}
}
