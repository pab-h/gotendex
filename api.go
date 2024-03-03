package gotendex

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/pab-h/gotendex/resources"
)

const APIURL string = "http://gutendex.com/books/"

type Api struct {
	baseUrl string
}

func NewApi() Api {
	return Api{
		baseUrl: APIURL,
	}
}

func (api Api) Book(id int) resources.Book {
	url := fmt.Sprintf("%v%v", api.baseUrl, id)

	response, error := http.Get(url)

	if error != nil {
		log.Fatal(error)
	}

	defer response.Body.Close()

	book := resources.Book{}

	decoder := json.NewDecoder(response.Body)

	if error := decoder.Decode(&book); error != nil {
		log.Fatal(error)
	}

	return book
}

func (api Api) Books() resources.Response {
	response, error := http.Get(api.baseUrl)

	if error != nil {
		log.Fatal(error)
	}

	defer response.Body.Close()

	apiResponse := resources.Response{}

	decoder := json.NewDecoder(response.Body)

	if error := decoder.Decode(&apiResponse); error != nil {
		log.Fatal(error)
	}

	return apiResponse
}

func (api Api) QueryBooks(query Query) resources.Response {
	url := fmt.Sprintf("%v%v", api.baseUrl, query.ToURI())

	response, error := http.Get(url)

	if error != nil {
		log.Fatal(error)
	}

	defer response.Body.Close()

	apiResponse := resources.Response{}

	decoder := json.NewDecoder(response.Body)

	if error := decoder.Decode(&apiResponse); error != nil {
		log.Fatal(error)
	}

	return apiResponse
}

func (api Api) Next(response resources.Response) resources.Response {
	return resources.Response{}
}

func (api Api) Previous(response resources.Response) resources.Response {
	return resources.Response{}
}
