package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
)

func (c *Client) GetPokemonInformation(name string) (Pokemon, error) {
	if name == "" {
		return Pokemon{}, errors.New("no name given")
	}
	url := baseUrl + "/pokemon/" + name

	body, err := c.doRequest(url)
	if err != nil {
		return Pokemon{}, err
	}

	var data Pokemon
	if err := json.Unmarshal(body, &data); err != nil {
		return Pokemon{}, fmt.Errorf("unmarshalling raised error: %w", err)
	}

	return data, nil
}
