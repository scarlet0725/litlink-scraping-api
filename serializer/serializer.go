package serializer

import (
	"encoding/json"
	"io"

	"github.com/PuerkitoBio/goquery"

	"github.com/scarlet0725/litlink-scraping-api/model"
)

type Serializer interface {
	Execute(io.Reader) (model.ApiResponse, error)
	Litlink(io.Reader) (model.ApiResponse, error)
	Livepocket(io.Reader) (model.ApiResponse, error)
	Kolokol(io.Reader) (model.ApiResponse, error)
}

type serializer struct {
}

func CreateSerializer() *serializer {
	return &serializer{}
}

func (s *serializer) Litlink(r io.Reader) (model.ApiResponse, error) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return model.ApiResponse{}, err
	}

	selection := doc.Find("#__NEXT_DATA__")
	b := []byte(selection.Text())

	var data model.LitlinkProps

	json.Unmarshal(b, &data)

	if err != nil {
		return model.ApiResponse{}, err
	}

	b = []byte(data.Props.PageProps.ProfileString)

	var profile model.LitlinkProfile
	err = json.Unmarshal(b, &profile)
	if err != nil {
		return model.ApiResponse{}, err
	}

	var profileDetails []model.LitlinkProfileDetail

	for _, v := range profile.ProfileLink.Details {

		//この処理をjsonと構造体のMarshallとUnmarshallでやりたい(たぶんできる)
		if v.ButtonLink.URL == "" {
			continue
		}

		profileDetails = append(profileDetails, model.LitlinkProfileDetail{
			Title:       v.ButtonLink.Title,
			URL:         v.ButtonLink.URL,
			Description: v.ButtonLink.Description,
		})
	}

	result := model.ApiResponse{
		Ok:         true,
		Livepocket: &[]model.LivepocketApplicationData{},
		Litlink:    &model.LitlinkData{Name: profile.Name, ProfileLinks: &profileDetails},
	}

	return result, nil

}

func (s *serializer) Livepocket(r io.Reader) (model.ApiResponse, error) {
	return model.ApiResponse{}, nil
}

func (s *serializer) Kolokol(r io.Reader) (model.ApiResponse, error) {
	//Todo
	return model.ApiResponse{}, nil
}

func (s *serializer) Execute(b io.Reader) (model.ApiResponse, error) {
	//Todo
	return model.ApiResponse{}, nil
}
