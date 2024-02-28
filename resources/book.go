package resources

type Book struct {
	Id            int
	Title         string
	Subjects      []string
	Authors       []Person
	Translators   []Person
	Bookshelves   []string
	Languages     []string
	Copyright     bool
	MediaType     string
	Formats       []Format
	DownloadCount int
}
