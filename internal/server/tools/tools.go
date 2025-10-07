package tools

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	"github.com/ezh0v/weather-mcp-server/internal/server/services"
)

type ToolFunc func(svc services.Services) (mcp.Tool, server.ToolHandlerFunc)
