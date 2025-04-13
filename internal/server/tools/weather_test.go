package tools

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCurrentWeather(t *testing.T) {
	tool, handler := CurrentWeather(nil)

	assert.Equal(t, "current_weather", tool.Name)
	assert.NotEmpty(t, tool.Description)
	assert.Contains(t, tool.InputSchema.Properties, "city")
	assert.ElementsMatch(t, tool.InputSchema.Required, []string{"city"})

	assert.NotNil(t, handler)
}
