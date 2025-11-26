package pokeapi

import (
	"encoding/json"
	"errors"
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

func (c *Client) GetLocationAreaData(name string) (LocationArea, error) {
	if name == "" {
		return LocationArea{}, errors.New("no location given")
	}
	url := baseUrl + "/location-area/" + name

	body, err := c.doRequest(url)
	if err != nil {
		return LocationArea{}, err
	}

	var data LocationArea
	if err := json.Unmarshal(body, &data); err != nil {
		return LocationArea{}, fmt.Errorf("unmarshalling raised error: %w", err)
	}

	return data, nil

}
