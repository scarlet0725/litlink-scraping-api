package main

import (
	"encoding/json"
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

	http.HandleFunc("/scraiping", scrapingRequestHandler)
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

		//この処理をjsonと構造体のMarshallとUnmarshallでやりたい(たぶんできる)
		if v.ButtonLink.URL == "" {
			continue
		}

		profileDetails = append(profileDetails, models.LitlinkProfileDetail{
			Title:       v.ButtonLink.Title,
			URL:         v.ButtonLink.URL,
			Description: v.ButtonLink.Description,
		})
	}

	result := models.ApiResponse{
		Ok:             true,
		LivepocketData: &[]models.LivepocketApplicationData{},
		LitlinkData:    &models.LitlinkData{Name: profile.Name, ProfileLinks: &profileDetails},
	}

	return result, nil

}

func scraipingLivepocket(url string) ([]models.LivepocketApplicationData, error) {
	res, err := http.Get(url)

	if err != nil {
		return []models.LivepocketApplicationData{}, err
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return []models.LivepocketApplicationData{}, err
	}
	selection, _ := doc.Find("#event_ticket_groups").Attr("value")
	b := []byte(selection)

	var data []models.LivepocketApplicationData

	json.Unmarshal(b, &data)

	if err != nil {
		return []models.LivepocketApplicationData{}, err
	}

	return data, nil
}

func serializeLivepocket(data *[]models.LivepocketApplicationData) (models.ApiResponse, error) {
	result := models.ApiResponse{
		Ok:             true,
		LivepocketData: data,
	}

	return result, nil
}

func scrapingRequestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	siteUrl := r.URL.Query().Get("target_url")

	u, err := url.Parse(siteUrl)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("{\"ok\": false, \"error\": \"invalid_url\"}"))
	}

	switch u.Host {
	case "t.livepocket.jp":
		scrapingResult, err := scraipingLivepocket(siteUrl)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("{\"ok\": false, \"error\": \"scraping_error\"}"))
		}

		serializedResult, err := serializeLivepocket(&scrapingResult)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("{\"ok\": false, \"error\": \"scraping_error\"}"))
		}
		b, err := json.Marshal(&serializedResult)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("{\"ok\": false, \"error\": \"scraping_error\"}"))
		}
		respondToClient(w, &b)
		return
	case "lit.link":
		scrapingResult, err := scrapingLitlink(siteUrl)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("{\"ok\": false, \"error\": \"scraping_error\"}"))
		}
		result, err := serializeLitlinkProps(&scrapingResult)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("{\"ok\": false, \"error\": \"scraping_error\"}"))
		}
		b, err := json.Marshal(result)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("{\"ok\": false, \"error\": \"scraping_error\"}"))
		}
		respondToClient(w, &b)
		return
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("{\"ok\": false, \"error\": \"unsupported_site\"}"))
		return
	}

}

func respondToClient(w http.ResponseWriter, b *[]byte) {
	w.WriteHeader(http.StatusOK)
	w.Write(*b)
}
