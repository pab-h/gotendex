# Gotendex

Wrapper for JSON web API for Project Gutenberg ebook metadata, [see here](http://gutendex.com/)

# How install

`go get https://github.com/pab-h/gotendex`

# Books

The gutendex.com/books endpoint is binded by the method
 ```go 
func (api Api) Books() *src.Response
```

The book struct is 
```go
type Book struct {
	Id            int               `json:"id"`
	Title         string            `json:"title"`
	Subjects      []string          `json:"subjects"`
	Authors       []Person          `json:"authors"`
	Translators   []Person          `json:"translators"`
	Bookshelves   []string          `json:"bookshelves"`
	Languages     []string          `json:"languages"`
	Copyright     bool              `json:"copyright"`
	MediaType     string            `json:"media_type"`
	Formats       map[string]string `json:"formats"`
	DownloadCount int               `json:"download_count"`
}
```

## Query Books
You can make a query using the method 
```go 
func (api Api) QueryBooks(query *Query) *src.Response
``` 

The Query struct is 
```go
type Query struct {
	AuthorYearStart int
	AuthorYearEnd   int
	Copyright       bool
	Ids             []int
	Languages       []string
	MimeType        string
	Search          string
	Sort            []string
	Topic           string
}
```

## Pagination
You can navigate the pagination using the following methods:
```go
func (api Api) Next(response *src.Response) (*src.Response, bool)
```
```go
func (api Api) Previous(response *src.Response) (*src.Response, bool)
```
The return `bool` indicates whether you were able to continue

## Example
```go
package main

import (
	"fmt"

	"github.com/pab-h/gotendex"
)

func main() {
	api := gotendex.NewApi()

	responseBooks := api.Books()

	fmt.Println("cout:", responseBooks.Count)

	for _, book := range responseBooks.Results {
		fmt.Println("Id:", book.Id)
		fmt.Println("Title:", book.Title)
	}

}
```

## Example with query

```go 
package main

import (
	"fmt"

	"github.com/pab-h/gotendex"
)

func main() {
	api := gotendex.NewApi()

	query := gotendex.Query{
		Languages: []string{"fr", "fi"},
	}

	responseBooks := api.QueryBooks(&query)

	fmt.Println("cout:", responseBooks.Count)

	for _, book := range responseBooks.Results {
		fmt.Println("Id:", book.Id)
		fmt.Println("Title:", book.Title)
	}

}
```
## Example with pagination

```go
package main

import (
	"fmt"

	"github.com/pab-h/gotendex"
)

func main() {
	api := gotendex.NewApi()

	query := gotendex.Query{
		Languages: []string{"fr", "fi"},
	}

	responseBooks := api.QueryBooks(&query)
	responseBooks, canGo := api.Next(responseBooks)

	if canGo {
		fmt.Println(" You can not go")
	}

	fmt.Println("cout:", responseBooks.Count)

	for _, book := range responseBooks.Results {
		fmt.Println("Id:", book.Id)
		fmt.Println("Title:", book.Title)
	}

}
```
