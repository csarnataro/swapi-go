package people

// PersonFields represent the fields available for a film in the original json file
type PersonFields struct {
	Name      string `json:"name"`
	URL       string `json:"url"`
	Starships []int  `json:"starships"`
	Vehicles  []int  `json:"vehicles"`
}

// PersonEntry represents a single record of a film in thes original json file
// It has a primary key and a list of fields. See `FilmFields` struct defined above
type PersonEntry struct {
	Fields PersonFields `json:"fields"`
	Pk     int          `json:"pk"`
}
