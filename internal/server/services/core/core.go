package core

import (
	"html/template"

	"github.com/ezh0v/weather-mcp-server/internal/server/services"
)

type CoreServices struct {
	renderer   *template.Template
	weatherAPI services.WeatherAPIProvider

	weatherService *WeatherService
}

func New(renderer *template.Template, weatherAPI services.WeatherAPIProvider) *CoreServices {
	return &CoreServices{
		renderer:   renderer,
		weatherAPI: weatherAPI,
	}
}

func (cs *CoreServices) Weather() services.WeatherService {
	if cs.weatherService == nil {
		cs.weatherService = &WeatherService{CoreServices: cs}
	}

	return cs.weatherService
}
