package models


type BreedData struct {
	Breed   string `json:"breed"`
	Country string `json:"country"`
	Origin  string `json:"origin"`
	Coat    string `json:"coat"`
	Pattern string `json:"pattern"`
}
type Response struct {
	Breeds   []BreedData `json:"data"`
	Pages    int         `json:"page"`
	Total    int         `json:"total"`
	LastPage int         `json:"last_page"`
	Links    []Links     `json:"links"`
}

type Links struct {
	Url    string `json:"url"`
	Label  string `json:"label"`
	Active bool `json:"active"`
}

type CatBreedByCountry map[string][]CatBreedResponseByCountry


type CatBreedResponseByCountry struct {
	Breed   string `json:"breed"`
	Origin  string `json:"origin"`
	Coat    string `json:"coat"`
	Pattern string `json:"pattern"`
}




type Payload struct {
	Str string `json:"str"`
}


