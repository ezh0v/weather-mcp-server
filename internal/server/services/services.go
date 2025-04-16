package services

import (
	"context"

	"github.com/TuanKiri/weather-mcp-server/pkg/weatherapi/models"
)

type WeatherAPIProvider interface {
	Current(ctx context.Context, city string) (*models.CurrentResponse, error)
}

type Services interface {
	Weather() WeatherService
}

type WeatherService interface {
	Current(ctx context.Context, city string) (string, error)
}
