package utils

// InputEntry is an 'input' in the sense of an entity coming from data json
type InputEntry interface {
	Build(serverName string) Result
}

// Result is a model for an entry
type Result interface{}

// ResultPage is the final feed with all the entries and related metadata
// (e.g. pagination links)
type ResultPage struct {
	Count    int      `json:"count"`
	Previous *string  `json:"previous"` // it's a pointer because it can be nil
	Next     *string  `json:"next"`     // it's a pointer because it can be nil
	Results  []Result `json:"results"`
}
