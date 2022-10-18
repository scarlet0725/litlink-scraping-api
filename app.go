package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/scarlet0725/litlink-scraping-api/models"
)

func main() {

	port, flg := os.LookupEnv("PORT")
	if !flg {
		port = "8080"
	}

	http.HandleFunc("/scraiping/litlink", scrapingRequestHandler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
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

func scrapingRequestHandler(w http.ResponseWriter, r *http.Request) {
	siteUrl := r.URL.Query().Get("target_url")

	u, err := url.Parse(siteUrl)

	if u.Host != "lit.link" || err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("{\"ok\": false, \"error\": \"invalid_url\"}"))
		return
	}

	scrapingResult, err := scrapingLitlink(siteUrl)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("{\"ok\": false, \"error\": \"invalid_url\"}"))
		return
	}

	serializedResult, err := serializeLitlinkProps(&scrapingResult)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("{\"ok\": false, \"error\": \"invalid_url\"}"))
		return
	}

	str, err := json.Marshal(serializedResult)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("{\"ok\": false, \"error\": \"invalid_url\"}"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string(str))

}
