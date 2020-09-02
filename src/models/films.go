package models

// Film holds all the information related to a film.
// It's used in the output JSON
type Film struct {
	Title        string   `json:"title,omitempty"`
	EpisodeID    int      `json:"episode_id,omitempty"`
	OpeningCrawl string   `json:"opening_crawl,omitempty"`
	Director     string   `json:"director,omitempty"`
	Producer     string   `json:"producer,omitempty"`
	ReleaseDate  string   `json:"release_date,omitempty"`
	Characters   []string `json:"characters,omitempty"`
	Planets      []string `json:"planets,omitempty"`
	Starships    []string `json:"starships,omitempty"`
	Vehicles     []string `json:"vehicles,omitempty"`
	Species      []string `json:"species,omitempty"`
	URL          string   `json:"url,omitempty"`
}

// FilmsPage is the final feed with all the films and related metadata
// (e.g. pagination links)
type FilmsPage struct {
	Count    int     `json:"count"`
	Previous *string `json:"previous"` // it's a pointer because it can be nil
	Next     *string `json:"next"`     // it's a pointer because it can be nil
	Results  []Film  `json:"results"`
}
