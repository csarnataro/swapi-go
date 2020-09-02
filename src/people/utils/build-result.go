package people

import (
	"errors"
	"strconv"

	"github.com/csarnataro/swapi-go/src/constants"
	"github.com/csarnataro/swapi-go/src/models"
)

var errNotFound = errors.New("Not found")

func getPage(people []models.Person, pageNumber uint64) ([]models.Person, error) {
	count := uint64(len(people))
	min := (pageNumber - 1) * constants.PageSize
	if min > count {
		return make([]models.Person, 0), errNotFound
	}
	max := pageNumber * constants.PageSize

	if max > count {
		max = count // in order to avoid slice overflow
	}
	page := people[min:max]
	if len(page) == 0 {
		return make([]models.Person, 0), errNotFound
	}
	return people[min:max], nil
}

func getURLs(serverName string, coll []int, entityType string) []string {
	var result []string
	for _, ID := range coll {
		result = append(result, serverName+"/api/"+entityType+"/"+strconv.Itoa(ID))
	}
	return result
}

func buildResult(entries []PersonEntry, serverName string, pageNumber uint64) (models.PeoplePage, error) {
	numOfResults := len(entries)
	result := models.PeoplePage{}
	result.Count = numOfResults

	var people []models.Person
	for _, entry := range entries {
		thisPerson := buildPerson(entry, serverName)
		people = append(people, thisPerson)
	}

	resultPage, err := getPage(people, pageNumber)

	if pageNumber == 1 {
		result.Previous = nil
	} else {
		previous := strconv.Itoa(int(pageNumber - 1))
		previousURL := serverName + "/api/people?page=" + previous
		result.Previous = &previousURL
	}

	if numOfResults < 5 {
		result.Next = nil
	} else {
		nextPageIndex := strconv.Itoa(int(pageNumber) + 1)
		result.Next = &nextPageIndex
	}

	if err != nil {
		return models.PeoplePage{}, err
	}
	result.Results = resultPage
	return result, nil
}

func buildPerson(entry PersonEntry, serverName string) models.Person {

	var thisPerson = models.Person{}
	thisPerson.URL = serverName + "/api/people/" + strconv.Itoa(entry.Pk)
	thisPerson.Name = entry.Fields.Name

	thisPerson.Starships = getURLs(serverName, entry.Fields.Starships, "starships")
	thisPerson.Vehicles = getURLs(serverName, entry.Fields.Vehicles, "vehicles")

	return thisPerson
}
