package parser

import (
	"encoding/json"

	"github.com/scarlet0725/prism-api/model"
)

type JsonParser interface {
	Ryzm([]byte) (model.RyzmAPIResponse, error)
}

type jsonParser struct {
}

func NewJsonParser() JsonParser {
	return &jsonParser{}
}

func (j *jsonParser) Ryzm(b []byte) (model.RyzmAPIResponse, error) {
	var data model.RyzmAPIResponse
	err := json.Unmarshal(b, &data)

	return data, err

}
