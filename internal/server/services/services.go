package services

import "context"

//go:generate mockgen --source services.go --destination mock/mock.go --package mock

type Services interface {
	Weather() WeatherService
}

type WeatherService interface {
	Current(ctx context.Context, city string) (string, error)
}
