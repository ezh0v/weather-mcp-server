package tools

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	"github.com/ezh0v/weather-mcp-server/internal/server/handlers"
	"github.com/ezh0v/weather-mcp-server/internal/server/services"
)

func CurrentWeather(svc services.Services) (mcp.Tool, server.ToolHandlerFunc) {
	tool := mcp.NewTool("current_weather",
		mcp.WithDescription(`
			The service retrieves the current weather for a specified city and returns only HTML and CSS code designed 
			for displaying weather information directly on a canvas. No additional comments should be included after the response. 
			It processes the city's name, gathers weather data such as temperature, humidity, wind speed, and general weather conditions, 
			and then generates a structured HTML layout along with appropriate CSS styles. This HTML must be rendered visually on the canvas. 
			After receiving the response, the text inside the HTML should be translated into the language of the request.
		`),
		mcp.WithString("city",
			mcp.Required(),
			mcp.Description(`
				The name of the city. This field is required and must be provided in English. 
				Only one city is allowed, and it must be the last one provided by the user.
			`),
		),
	)

	handler := handlers.CurrentWeather(svc)

	return tool, handler
}
