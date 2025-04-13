package core

import (
	"bytes"
	"context"

	"github.com/TuanKiri/weather-mcp-server/internal/server/view/models"
)

type WeatherService struct {
	*CoreServices
}

func (ws *WeatherService) Current(ctx context.Context, city string) (string, error) {
	data, err := ws.weatherAPI.Current(ctx, city)
	if err != nil {
		return "", err
	}

	var current models.CurrentWeather
	current.FromWeatherAPI(data)

	var buf bytes.Buffer

	if err := ws.renderer.ExecuteTemplate(&buf, "weather.html", current); err != nil {
		return "", err
	}

	return buf.String(), nil
}
