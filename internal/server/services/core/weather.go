package core

import (
	"bytes"
	"context"
	"fmt"
)

type WeatherService struct {
	*CoreServices
}

func (ws *WeatherService) Current(ctx context.Context, city string) (string, error) {
	data, err := ws.weatherAPI.Current(ctx, city)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer

	if err := ws.renderer.ExecuteTemplate(&buf, "weather.html", map[string]string{
		"Location":    fmt.Sprintf("%s, %s", data.Location.Name, data.Location.Country),
		"Icon":        "https:" + data.Current.Condition.Icon,
		"Condition":   data.Current.Condition.Text,
		"Temperature": fmt.Sprintf("%.0f", data.Current.TempC),
		"Humidity":    fmt.Sprintf("%d", data.Current.Humidity),
		"WindSpeed":   fmt.Sprintf("%.0f", data.Current.WindKph),
	}); err != nil {
		return "", err
	}

	return buf.String(), nil
}
