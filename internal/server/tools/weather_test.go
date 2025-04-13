package tools

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	mocks "github.com/TuanKiri/weather-mcp-server/internal/server/services/mock"
)

func TestCurrentWeather(t *testing.T) {
	ctrl := gomock.NewController(t)

	svc := mocks.NewMockServices(ctrl)

	tool, handler := CurrentWeather(svc)

	assert.Equal(t, "current_weather", tool.Name)
	assert.NotEmpty(t, tool.Description)
	assert.Contains(t, tool.InputSchema.Properties, "city")
	assert.ElementsMatch(t, tool.InputSchema.Required, []string{"city"})

	assert.NotNil(t, handler)
}
