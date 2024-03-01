package resources

type Response struct {
	Count    int    `json:"count"`
	Results  []Book `json:"results"`
	Previous string `json:"previous"`
	Next     string `json:"next"`
}
