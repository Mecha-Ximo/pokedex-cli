package pokemon_api

type ResourceList struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
}

type LocationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocationAreasDTO struct {
	ResourceList
	Results []LocationArea `json:"results"`
}


