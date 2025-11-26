package pokeapi

import (
	"encoding/json"
	"fmt"
)

func (c *Client) GetLocationAreas(pageURL *string) (LocationAreasPage, error) {
	url := baseUrl + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	body, err := c.doRequest(url)
	if err != nil {
		return LocationAreasPage{}, err
	}

	var data LocationAreasPage
	if err := json.Unmarshal(body, &data); err != nil {
		return LocationAreasPage{}, fmt.Errorf("unmarshalling raised error: %w", err)
	}

	return data, nil
}
