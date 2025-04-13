package services

import (
	"context"
	"io"

	"github.com/TuanKiri/weather-mcp-server/pkg/weatherapi/models"
)

type TemplateRenderer interface {
	ExecuteTemplate(wr io.Writer, name string, data any) error
}

type WeatherAPIProvider interface {
	Current(ctx context.Context, city string) (*models.CurrentResponse, error)
}

type Services interface {
	Weather() WeatherService
}

type WeatherService interface {
	Current(ctx context.Context, city string) (string, error)
}
