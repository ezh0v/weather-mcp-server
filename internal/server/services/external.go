package services

import (
	"context"

	"github.com/ezh0v/weather-mcp-server/pkg/weatherapi/models"
)

//go:generate mockgen --source external.go --destination mock/external_mock.go --package mock

type WeatherAPIProvider interface {
	Current(ctx context.Context, city string) (*models.CurrentResponse, error)
}
