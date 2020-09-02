package models

// Person holds all the information related to a character.
// It's used in the output JSON
type Person struct {
	Name      string   `json:"name,omitempty"`
	Starships []string `json:"starships,omitempty"`
	Vehicles  []string `json:"vehicles,omitempty"`
	URL       string   `json:"url,omitempty"`
}

// PeoplePage is the final feed with all the films and related metadata
// (e.g. pagination links)
type PeoplePage struct {
	Count    int      `json:"count"`
	Previous *string  `json:"previous"` // it's a pointer because it can be nil
	Next     *string  `json:"next"`     // it's a pointer because it can be nil
	Results  []Person `json:"results"`
}
