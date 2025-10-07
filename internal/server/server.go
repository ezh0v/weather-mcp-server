package server

import (
	"context"
	"embed"
	"html/template"
	"log"
	"os/signal"
	"syscall"

	"github.com/mark3labs/mcp-go/server"

	"github.com/ezh0v/weather-mcp-server/internal/server/services/core"
	"github.com/ezh0v/weather-mcp-server/internal/server/tools"
	"github.com/ezh0v/weather-mcp-server/pkg/weatherapi"
)

//go:embed view
var templates embed.FS

func Run(cfg *Config) error {
	tmpl, err := template.ParseFS(templates, "view/*.html")
	if err != nil {
		return err
	}

	wApi := weatherapi.New(cfg.WeatherAPIKey, cfg.WeatherAPITimeout)

	svc := core.New(tmpl, wApi)

	s := server.NewMCPServer(
		"Weather Server",
		"1.0.0",
		server.WithLogging(),
	)

	toolFuncs := []tools.ToolFunc{
		tools.CurrentWeather,
	}

	for _, tool := range toolFuncs {
		s.AddTool(tool(svc))
	}

	if cfg.ListenAddr != "" {
		return serveSSE(s, cfg.ListenAddr)
	}

	return server.ServeStdio(s)
}

func serveSSE(s *server.MCPServer, addr string) error {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	srv := server.NewSSEServer(s)

	go func() {
		if err := srv.Start(addr); err != nil {
			log.Fatal(err)
		}
	}()

	<-ctx.Done()

	return srv.Shutdown(context.TODO())
}
