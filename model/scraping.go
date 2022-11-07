package model

type ScrapingResult struct {
	Data []byte
}

type ScrapingRequest struct {
	URL  string
	Type string
}
