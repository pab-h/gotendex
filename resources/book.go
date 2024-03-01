package resources

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
