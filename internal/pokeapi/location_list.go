package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocationAreas(pageURL *string) (LocationAreasPage, error) {
	url := baseUrl + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreasPage{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasPage{}, fmt.Errorf("error making request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(io.LimitReader(res.Body, 1024))
		return LocationAreasPage{}, fmt.Errorf("request unsuccessfull with status code: %d and\nbody:\n%s", res.StatusCode, body)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreasPage{}, fmt.Errorf("reading of body failed with err: %w", err)
	}

	var data LocationAreasPage
	if err := json.Unmarshal(body, &data); err != nil {
		return LocationAreasPage{}, fmt.Errorf("unmarshalling raised error: %w", err)
	}

	return data, nil
}
