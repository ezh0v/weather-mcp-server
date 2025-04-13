package models

import (
	"fmt"

	"github.com/TuanKiri/weather-mcp-server/pkg/weatherapi/models"
)

type CurrentWeather struct {
	Location    string
	Icon        string
	Condition   string
	Temperature string
	Humidity    string
	WindSpeed   string
}

func (w *CurrentWeather) FromWeatherAPI(data *models.CurrentResponse) {
	w.Location = fmt.Sprintf("%s, %s", data.Location.Name, data.Location.Country)
	w.Icon = "https:" + data.Current.Condition.Icon
	w.Condition = data.Current.Condition.Text
	w.Temperature = fmt.Sprintf("%.0f", data.Current.TempC)
	w.Humidity = fmt.Sprintf("%d", data.Current.Humidity)
	w.WindSpeed = fmt.Sprintf("%.0f", data.Current.WindKph)
}
