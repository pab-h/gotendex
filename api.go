package gotendex

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	src "github.com/pab-h/gotendex/resources"
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

func (api Api) Book(id int) (*src.Book, error) {
	url := fmt.Sprintf("%v%v", api.baseUrl, id)

	response, error := http.Get(url)

	if error != nil {
		return &src.Book{}, error
	}

	defer response.Body.Close()

	book := src.Book{}

	decoder := json.NewDecoder(response.Body)

	if response.StatusCode == http.StatusNotFound {
		return &src.Book{}, errors.New("not found")
	}

	if error := decoder.Decode(&book); error != nil {
		return &src.Book{}, error
	}

	return &book, nil
}

func (api Api) Books() *src.Response {
	response, error := http.Get(api.baseUrl)

	if error != nil {
		log.Fatal(error)
	}

	defer response.Body.Close()

	apiResponse := src.Response{}

	decoder := json.NewDecoder(response.Body)

	if error := decoder.Decode(&apiResponse); error != nil {
		log.Fatal(error)
	}

	return &apiResponse
}

func (api Api) QueryBooks(query *Query) *src.Response {
	url := fmt.Sprintf("%v?%v", api.baseUrl, query.ToURI())

	response, error := http.Get(url)

	if error != nil {
		log.Fatal(error)
	}

	defer response.Body.Close()

	apiResponse := src.Response{}

	decoder := json.NewDecoder(response.Body)

	if error := decoder.Decode(&apiResponse); error != nil {
		log.Fatal(error)
	}

	return &apiResponse
}

func (api Api) Next(response *src.Response) (*src.Response, bool) {
	if len(response.Next) == 0 {
		return &src.Response{}, false
	}

	nextResponse, error := http.Get(response.Next)

	if error != nil {
		log.Fatal(error)
	}

	defer nextResponse.Body.Close()

	apiResponse := src.Response{}

	decoder := json.NewDecoder(nextResponse.Body)

	if error := decoder.Decode(&apiResponse); error != nil {
		log.Fatal(error)
	}

	return &apiResponse, true
}

func (api Api) Previous(response *src.Response) (*src.Response, bool){
	if len(response.Previous) == 0 {
		return &src.Response{}, false
	}

	previousResponse, error := http.Get(response.Previous)

	if error != nil {
		log.Fatal(error)
	}

	defer previousResponse.Body.Close()

	apiResponse := src.Response{}

	decoder := json.NewDecoder(previousResponse.Body)

	if error := decoder.Decode(&apiResponse); error != nil {
		log.Fatal(error)
	}

	return &apiResponse, true
}
