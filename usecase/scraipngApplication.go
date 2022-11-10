package usecase

import (
	"bytes"

	"github.com/scarlet0725/prism-api/controller"
	"github.com/scarlet0725/prism-api/model"
	"github.com/scarlet0725/prism-api/parser"
	"github.com/scarlet0725/prism-api/selializer"
)

type ScrapingApplication interface {
	Execute(*model.ScrapingRequest) (model.APIResponse, error)
}

type scrapingApplication struct {
	f controller.FetchController
	s selializer.ResponseSerializer
	p parser.DocParser
}

func NewScrapingApplication(f controller.FetchController, s selializer.ResponseSerializer, p parser.DocParser) ScrapingApplication {
	return &scrapingApplication{
		f: f,
		s: s,
		p: p,
	}
}

func (s *scrapingApplication) Execute(r *model.ScrapingRequest) (model.APIResponse, error) {
	result, err := s.f.Fetch(r)
	if err != nil {
		return model.APIResponse{}, err
	}

	b := bytes.NewReader(result.Data)

	var res interface{}

	switch result.Request.Host {
	case "t.livepocket.jp":
		res, err = s.p.Livepocket(b)
	case "lit.link":
		res, err = s.p.Litlink(b)
	default:
		return model.APIResponse{}, &model.AppError{
			Code: 400,
			Msg:  "unsupported_site",
		}
	}

	if err != nil {
		return model.APIResponse{}, err
	}

	return s.s.BuildResponse(res)

}
