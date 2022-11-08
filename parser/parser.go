package parser

import (
	"encoding/json"
	"io"

	"github.com/PuerkitoBio/goquery"

	"github.com/scarlet0725/prism-api/model"
)

type DocParser interface {
	Execute(io.Reader) (model.APIResponse, error)
	Litlink(io.Reader) (model.APIResponse, error)
	Livepocket(io.Reader) (model.APIResponse, error)
	Kolokol(io.Reader) (model.APIResponse, error)
}

type docParser struct {
}

func CreateSerializer() DocParser {
	return &docParser{}
}

func (s *docParser) Litlink(r io.Reader) (model.APIResponse, error) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return model.APIResponse{}, err
	}

	selection := doc.Find("#__NEXT_DATA__")
	b := []byte(selection.Text())

	var data model.LitlinkProps

	json.Unmarshal(b, &data)

	if err != nil {
		return model.APIResponse{}, err
	}

	b = []byte(data.Props.PageProps.ProfileString)

	var profile model.LitlinkProfile
	err = json.Unmarshal(b, &profile)
	if err != nil {
		return model.APIResponse{}, err
	}

	var profileDetails []model.LitlinkProfileDetail

	for _, v := range profile.ProfileLink.Details {

		//TODO: この処理をjsonと構造体のMarshallとUnmarshallでやりたい(たぶんできる)
		if v.ButtonLink.URL == "" {
			continue
		}

		profileDetails = append(profileDetails, model.LitlinkProfileDetail{
			Title:       v.ButtonLink.Title,
			URL:         v.ButtonLink.URL,
			Description: v.ButtonLink.Description,
		})
	}

	result := model.APIResponse{
		Ok:         true,
		Livepocket: []model.LivepocketApplicationData{},
		Litlink:    model.LitlinkData{Name: profile.Name, ProfileLinks: profileDetails},
	}

	return result, nil
}

func (s *docParser) Livepocket(r io.Reader) (model.APIResponse, error) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return model.APIResponse{}, err
	}
	selection, _ := doc.Find("#event_ticket_groups").Attr("value")
	b := []byte(selection)

	var data []model.LivepocketApplicationData

	json.Unmarshal(b, &data)

	if err != nil {
		return model.APIResponse{}, err
	}

	result := model.APIResponse{
		Ok:         true,
		Livepocket: data,
	}

	return result, nil
}

func (s *docParser) Kolokol(r io.Reader) (model.APIResponse, error) {
	//Todo
	return model.APIResponse{}, nil
}

func (s *docParser) Execute(r io.Reader) (model.APIResponse, error) {
	//Todo
	return model.APIResponse{}, nil
}
