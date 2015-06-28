package api

type feedRSS struct {
	Title       string `xml:"title" json:"title"`
	Description string `xml:"description" json:"description"`
	Link        string `xml:"link" json:"link"`
}

type feedAtom struct {
	Title       string `xml:"title"`
	Description string `xml:"subtitle" json:"description"`
	Link        string `xml:"link" json:"link"`
}
