package gotendex

import (
	"fmt"
	"strings"
)

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

func convertIds(ids []int) ([]string) {
	convertedIds := make([]string, len(ids))

	for index, value := range ids {
		convertedIds[index] = fmt.Sprintf("%v", value)
	}

	return convertedIds
}

func (query Query) ToURI() (string) {
	queryParams := make([]string, 0)

	if query.AuthorYearStart != 0 {
		queryParams = append(
			queryParams,
			fmt.Sprintf(
				"author_year_start=%d",
				query.AuthorYearStart,
			),
		)
	}

	if query.AuthorYearEnd != 0 {
		queryParams = append(
			queryParams,
			fmt.Sprintf(
				"author_year_end=%d",
				query.AuthorYearEnd,
			),
		)
	}

	if query.Copyright {
		queryParams = append(
			queryParams,
			"copyright=true",
		)
	}

	if len(query.Ids) > 0 {
		convertedIds := convertIds(query.Ids)

		queryParams = append(
			queryParams,
			fmt.Sprintf(
				"ids=%s",
				strings.Join(convertedIds, ","),
			),
		)
	}

	if len(query.Languages) > 0 {
		queryParams = append(
			queryParams,
			fmt.Sprintf(
				"languages=%s", 
				strings.Join(query.Languages, ","),
			),
		)
	}

	if query.MimeType != "" {
		queryParams = append(
			queryParams,
			fmt.Sprintf(
				"mime_type=%s", 
				query.MimeType,
			),
		)
	}

	if query.Search != "" {
		queryParams = append(
			queryParams,
			fmt.Sprintf("search=%s", query.Search),
		)
	}

	if len(query.Sort) > 0 {
		queryParams = append(
			queryParams,
			fmt.Sprintf(
				"sort=%s", 
				strings.Join(query.Sort, ","),
			),
		)
	}

	if query.Topic != "" {
		queryParams = append(
			queryParams,
			fmt.Sprintf("topic=%s", query.Topic),
		)
	}

	return strings.Join(queryParams, "&")
}
