package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/scarlet0725/litlink-scraping-api/models"
)

func main() {

	siteUrl := os.Getenv("SITE_URL")

	scrapingResult, err := scrapingLitlink(siteUrl)

	if err != nil {
		panic(err)
	}

	serializedResult, err := serializeLitlinkProps(&scrapingResult)

	if err != nil {
		panic(err)
	}

	str, err := json.Marshal(serializedResult)

	if err != nil {
		panic(err)
	}
	fmt.Println(string(str))
}

func scrapingLitlink(url string) (models.LitlinkProps, error) {
	res, err := http.Get(url)

	if err != nil {
		return models.LitlinkProps{}, err
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return models.LitlinkProps{}, err
	}
	selection := doc.Find("#__NEXT_DATA__")
	b := []byte(selection.Text())

	var data models.LitlinkProps

	json.Unmarshal(b, &data)

	if err != nil {
		return models.LitlinkProps{}, err
	}

	return data, nil
}

func serializeLitlinkProps(props *models.LitlinkProps) (models.ApiResponse, error) {
	b := []byte(props.Props.PageProps.ProfileString)

	var profile models.LitlinkProfile
	err := json.Unmarshal(b, &profile)
	if err != nil {
		return models.ApiResponse{}, err
	}

	var profileDetails []models.LitlinkProfileDetail

	for _, v := range profile.ProfileLink.Details {
		profileDetails = append(profileDetails, models.LitlinkProfileDetail{
			Title:       v.ButtonLink.Title,
			URL:         v.ButtonLink.URL,
			Description: v.ButtonLink.Description,
		})
	}

	result := models.ApiResponse{
		Ok:           true,
		Name:         profile.Name,
		ProfileLinks: profileDetails,
	}

	return result, nil
}
