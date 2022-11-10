package model

type ScrapingResult struct {
	Data    []byte
	Request *ScrapingRequest
}

type ScrapingRequest struct {
	URL  string
	Host string
}

type CacheData struct {
	Key   string
	Value []byte
}

type LitlinkParseResult struct {
	Data LitlinkData
}

type LivepocketParseResult struct {
	Data []LivepocketApplicationData
}
