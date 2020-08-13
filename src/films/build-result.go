package films

import (
	"errors"
	"strconv"

	"github.com/csarnataro/swapi-go/src/constants"
	"github.com/csarnataro/swapi-go/src/models"
)

var errNotFound = errors.New("Not found")

func getPage(films []models.Film, pageNumber uint64) ([]models.Film, error) {
	count := uint64(len(films))
	min := (pageNumber - 1) * constants.PageSize
	if min > count {
		return make([]models.Film, 0), errNotFound
	}
	max := pageNumber * constants.PageSize

	if max > count {
		max = count // in order to avoid slice overflow
	}
	page := films[min:max]
	if len(page) == 0 {
		return make([]models.Film, 0), errNotFound
	}
	return films[min:max], nil
}

func getURLs(coll []int, entityType string) []string {
	var result []string
	for _, ID := range coll {
		result = append(result, constants.BaseURL+"/api/"+entityType+"/"+strconv.Itoa(ID))
	}
	return result
}

func buildResult(entries []FilmEntry, pageNumber uint64) (models.FilmPage, error) {
	var result = models.FilmPage{}
	result.Count = len(entries)

	var films []models.Film
	for _, entry := range entries {
		var thisFilm = models.Film{}
		thisFilm.URL = constants.BaseURL + "/api/films/" + strconv.Itoa(entry.Pk)
		thisFilm.Title = entry.Fields.Title
		thisFilm.Director = entry.Fields.Director
		thisFilm.EpisodeID = entry.Fields.EpisodeID
		thisFilm.OpeningCrawl = entry.Fields.OpeningCrawl
		thisFilm.ReleaseDate = entry.Fields.ReleaseDate

		thisFilm.Characters = getURLs(entry.Fields.Characters, "people")
		thisFilm.Planets = getURLs(entry.Fields.Planets, "planets")
		thisFilm.Starships = getURLs(entry.Fields.Starships, "starships")
		thisFilm.Vehicles = getURLs(entry.Fields.Vehicles, "vehicles")
		thisFilm.Species = getURLs(entry.Fields.Species, "species")
		films = append(films, thisFilm)
	}

	resultPage, err := getPage(films, pageNumber)

	if pageNumber == 1 {
		result.Previous = nil
	} else {
		previous := strconv.Itoa(int(pageNumber - 1))
		previousURL := constants.BaseURL + "/api/films?page=" + previous
		result.Previous = &previousURL
	}

	next := strconv.Itoa(int(pageNumber))
	result.Next = &next

	if err != nil {
		return models.FilmPage{}, err
	}
	result.Results = resultPage
	return result, nil
}
