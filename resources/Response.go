package resources

type Response struct {
	Count    int
	Results []Book
	Previous string
	Next     string
}
