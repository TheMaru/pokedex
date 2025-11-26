package pokeapi

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/TheMaru/pokedexcli/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	cache      *pokecache.Cache
}

func (c *Client) doRequest(fullURL string) ([]byte, error) {
	if data, exists := c.cache.Get(fullURL); exists {
		return data, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(io.LimitReader(res.Body, 1024))
		return nil, fmt.Errorf("request unsuccessfull with status code: %d and\nbody:\n%s", res.StatusCode, body)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("reading of body failed with err: %w", err)
	}

	c.cache.Add(fullURL, body)

	return body, nil
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		httpClient: http.Client{Timeout: timeout},
		cache:      pokecache.NewCache(cacheInterval),
	}
}
