package utils

import (
	"errors"
	"strconv"

	"github.com/csarnataro/swapi-go/src/constants"
)

// ErrNotFound is the error generated when the entity is not found
var ErrNotFound = errors.New("Not found")

// GetPage gets the right page from whole the results
func GetPage(entries []Result, pageNumber uint64) ([]Result, error) {
	count := uint64(len(entries))
	min := (pageNumber - 1) * constants.PageSize
	if min > count {
		return make([]Result, 0), ErrNotFound
	}
	max := pageNumber * constants.PageSize

	if max > count {
		max = count // in order to avoid slice overflow
	}
	page := entries[min:max]
	if len(page) == 0 {
		return make([]Result, 0), ErrNotFound
	}
	return entries[min:max], nil
}

// GetURL builds the URL for a specific result
func GetURL(serverName string, coll []int, entityType string) []string {
	var result []string
	for _, ID := range coll {
		result = append(result, serverName+"/api/"+entityType+"/"+strconv.Itoa(ID))
	}
	return result
}

// BuildResult builds the result page
func BuildResult(entries []InputEntry, serverName string, pageNumber uint64) (ResultPage, error) {
	numOfResults := len(entries)
	result := ResultPage{}
	result.Count = numOfResults

	var results []Result
	for _, entry := range entries {
		thisEntry := entry.Build(serverName)
		results = append(results, thisEntry)
	}

	resultPage, err := GetPage(results, pageNumber)

	if pageNumber == 1 {
		result.Previous = nil
	} else {
		previous := strconv.Itoa(int(pageNumber - 1))
		previousURL := serverName + "/api/films?page=" + previous
		result.Previous = &previousURL
	}

	if len(resultPage) < constants.PageSize {
		result.Next = nil
	} else {
		nextPageIndex := strconv.Itoa(int(pageNumber) + 1)
		result.Next = &nextPageIndex
	}

	if err != nil {
		return ResultPage{}, err
	}
	result.Results = resultPage
	return result, nil
}

// func buildResult(entries []FilmEntry, serverName string, pageNumber uint64) (models.FilmsPage, error) {
// 	numOfResults := len(entries)
// 	result := models.FilmsPage{}
// 	result.Count = numOfResults

// 	var films []models.Film
// 	for _, entry := range entries {
// 		thisFilm := buildFilm(entry, serverName)
// 		films = append(films, thisFilm)
// 	}

// 	resultPage, err := getPage(films, pageNumber)

// 	if pageNumber == 1 {
// 		result.Previous = nil
// 	} else {
// 		previous := strconv.Itoa(int(pageNumber - 1))
// 		previousURL := serverName + "/api/films?page=" + previous
// 		result.Previous = &previousURL
// 	}

// 	if len(resultPage) < constants.PageSize {
// 		result.Next = nil
// 	} else {
// 		nextPageIndex := strconv.Itoa(int(pageNumber) + 1)
// 		result.Next = &nextPageIndex
// 	}

// 	if err != nil {
// 		return models.FilmsPage{}, err
// 	}
// 	result.Results = resultPage
// 	return result, nil
// }

// func buildFilm(entry FilmEntry, serverName string) models.Film {

// 	var thisFilm = models.Film{}
// 	thisFilm.URL = serverName + "/api/films/" + strconv.Itoa(entry.Pk)
// 	thisFilm.Title = entry.Fields.Title
// 	thisFilm.Director = entry.Fields.Director
// 	thisFilm.EpisodeID = entry.Fields.EpisodeID
// 	thisFilm.OpeningCrawl = entry.Fields.OpeningCrawl
// 	thisFilm.ReleaseDate = entry.Fields.ReleaseDate

// 	thisFilm.Characters = getURLs(serverName, entry.Fields.Characters, "people")
// 	thisFilm.Planets = getURLs(serverName, entry.Fields.Planets, "planets")
// 	thisFilm.Starships = getURLs(serverName, entry.Fields.Starships, "starships")
// 	thisFilm.Vehicles = getURLs(serverName, entry.Fields.Vehicles, "vehicles")
// 	thisFilm.Species = getURLs(serverName, entry.Fields.Species, "species")

// 	return thisFilm
// }
