package api

type feed struct {
	Title string `xml:"title"`
}

type feedRSS struct {
	Title string `xml:"title"`
}

type feedAtom struct {
	Title string `xml:"title"`
}
