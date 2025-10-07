package weatherapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/ezh0v/weather-mcp-server/pkg/weatherapi/models"
)

const baseURL = "http://api.weatherapi.com"

type WeatherAPI struct {
	key     string
	baseURL string
	client  *http.Client
}

func New(key string, timeout time.Duration) *WeatherAPI {
	return &WeatherAPI{
		key:     key,
		baseURL: baseURL,
		client: &http.Client{
			Timeout: timeout,
		},
	}
}

func (w *WeatherAPI) Current(ctx context.Context, city string) (*models.CurrentResponse, error) {
	query := url.Values{
		"key": {w.key},
		"q":   {city},
	}

	request, err := http.NewRequestWithContext(ctx,
		http.MethodGet,
		w.baseURL+"/v1/current.json?"+query.Encode(),
		nil,
	)
	if err != nil {
		return nil, err
	}

	response, err := w.client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("weather API not available. Code: %d", response.StatusCode)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var data models.CurrentResponse

	if err = json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return &data, nil
}
