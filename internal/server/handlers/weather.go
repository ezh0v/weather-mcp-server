package handlers

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	"github.com/ezh0v/weather-mcp-server/internal/server/services"
)

func CurrentWeather(svc services.Services) server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		city, ok := request.Params.Arguments["city"].(string)
		if !ok {
			return mcp.NewToolResultError("city must be a string"), nil
		}

		data, err := svc.Weather().Current(ctx, city)
		if err != nil {
			return nil, err
		}

		return mcp.NewToolResultText(data), nil
	}
}
