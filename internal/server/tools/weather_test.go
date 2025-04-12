package tools

import (
	"context"
	"testing"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/TuanKiri/weather-mcp-server/internal/server/services/mock"
)

func TestCurrentWeather(t *testing.T) {
	testCases := map[string]struct {
		arguments map[string]any
		errString string
		wait      string
	}{
		"empty_city": {
			wait: "city must be a string",
		},
		"city_not_found": {
			arguments: map[string]any{
				"city": "Tokyo",
			},
			errString: "weather API not available. Code: 400",
		},
		"successful_request": {
			arguments: map[string]any{
				"city": "London",
			},
		},
	}

	svc := mock.New()

	tool, handler := CurrentWeather(svc)

	assert.Equal(t, "current_weather", tool.Name)
	assert.NotEmpty(t, tool.Description)
	assert.Contains(t, tool.InputSchema.Properties, "city")
	assert.ElementsMatch(t, tool.InputSchema.Required, []string{"city"})

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			request := mcp.CallToolRequest{
				Params: struct {
					Name      string         `json:"name"`
					Arguments map[string]any `json:"arguments,omitempty"`
					Meta      *struct {
						ProgressToken mcp.ProgressToken `json:"progressToken,omitempty"`
					} `json:"_meta,omitempty"`
				}{
					Arguments: tc.arguments,
				},
			}

			result, err := handler(context.Background(), request)
			if err != nil {
				assert.EqualError(t, err, tc.errString)
				return
			}

			require.Len(t, result.Content, 1)
			content, ok := result.Content[0].(mcp.TextContent)
			require.True(t, ok)

			assert.Equal(t, tc.wait, content.Text)
		})
	}
}
