package main

import (
	"errors"
	"fmt"
)

func commandMapf(cfg *config) error {
	locationsRes, err := cfg.pokeapiClient.GetLocationAreas(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsRes.Next
	cfg.prevLocationsURL = locationsRes.Previous

	for _, location := range locationsRes.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapb(cfg *config) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page already")
	}

	locationRes, err := cfg.pokeapiClient.GetLocationAreas(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationRes.Next
	cfg.prevLocationsURL = locationRes.Previous

	for _, location := range locationRes.Results {
		fmt.Println(location.Name)
	}

	return nil
}
