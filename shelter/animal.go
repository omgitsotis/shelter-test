package shelter

type Image struct {
	URL string `json:"url"`
}

type Animal struct {
	Name  string `json:"name"`
	Age   string `json:"age"`
	Image string `json:"image"`
	Type  string `json:"type"`
}
