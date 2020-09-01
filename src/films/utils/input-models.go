package films

// FilmFields represent the fields available for a film in the original json file
type FilmFields struct {
	Title        string `json:"title"`
	Director     string `json:"director"`
	Characters   []int  `json:"characters"`
	URL          string `json:"url"`
	EpisodeID    int    `json:"episode_id"`
	OpeningCrawl string `json:"opening_crawl"`
	Planets      []int  `json:"planets"`
	Producer     string `json:"producer"`
	ReleaseDate  string `json:"release_date"`
	Species      []int  `json:"species"`
	Starships    []int  `json:"starships"`
	Vehicles     []int  `json:"vehicles"`
}

// FilmEntry represents a single record of a film in thes original json file
// It has a primary key and a list of fields. See `FilmFields` struct defined above
type FilmEntry struct {
	Fields FilmFields `json:"fields"`
	Pk     int        `json:"pk"`
}
