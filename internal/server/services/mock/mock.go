package mock

import "github.com/TuanKiri/weather-mcp-server/internal/server/services"

type MockServices struct {
	weatherService *WeatherService
}

func New() *MockServices {
	return &MockServices{}
}

func (ms *MockServices) Weather() services.WeatherService {
	if ms.weatherService == nil {
		ms.weatherService = &WeatherService{MockServices: ms}
	}

	return ms.weatherService
}
