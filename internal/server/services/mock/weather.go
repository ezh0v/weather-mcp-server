package mock

import (
	"context"
)

type WeatherService struct {
	*MockServices
}

func (ws *WeatherService) Current(ctx context.Context, city string) (string, error) {
	return "", nil
}
