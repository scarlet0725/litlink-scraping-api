package model

type ScrapingResult struct {
	Data    []byte
	Request *ScrapingRequest
}

type ScrapingRequest struct {
	URL    string
	Host   string
	Type   string
	Option FetchOptions
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

type FetchOptions struct {
	IsUseCache bool
	HTTPHeader map[string]string
	HTTPParams map[string]string
}

type BindingRequest struct {
	Group string `uri:"group" binding:"required"`
}
