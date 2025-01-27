package pokemon_api

import (
	"encoding/json"
	"io"
	"net/http"
)

func RequestLocationAreas(url string) ([]LocationArea, string, string) {
	res, err := http.Get(url)

	if err != nil {
		return []LocationArea{}, "", ""
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return []LocationArea{}, "", ""
	}

	locationAreas := LocationAreasDTO{}

	err = json.Unmarshal(data, &locationAreas)
	if err != nil {
		return []LocationArea{}, "", ""
	}

	return locationAreas.Results, locationAreas.Next, locationAreas.Previous
}