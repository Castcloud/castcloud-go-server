package api

type feed struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Link        string `json:"link"`
}

type feedAtom struct {
	Title       string `xml:"title"`
	Description string `xml:"subtitle" json:"description"`
	Link        string `xml:"link" json:"link"`
}
