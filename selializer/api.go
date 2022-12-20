package selializer

import (
	"time"

	"github.com/scarlet0725/prism-api/model"
)

type ResponseSerializer interface {
	BuildResponse(interface{}) (model.APIResponse, error)
	SelializeRyzmData(model.RyzmAPIResponse) ([]*model.Event, error)
}

type apiResponse struct{}

func NewResponseSerializer() ResponseSerializer {
	return &apiResponse{}
}

func (a *apiResponse) BuildResponse(i interface{}) (model.APIResponse, error) {
	var result model.APIResponse
	var err error

	switch i.(type) {
	case model.LitlinkParseResult:
		v, _ := i.(model.LitlinkParseResult)
		result = model.APIResponse{
			Ok:      true,
			Litlink: v.Data,
		}

		return result, nil
	case model.LivepocketParseResult:
		v, _ := i.(model.LivepocketParseResult)
		result = model.APIResponse{
			Ok:         true,
			Livepocket: v.Data,
		}
	default:
		err = &model.AppError{
			Code: 500,
			Msg:  "internal_server_error",
		}

	}

	return result, err
}

func (a *apiResponse) SelializeRyzmData(input model.RyzmAPIResponse) ([]*model.Event, error) {
	result := []*model.Event{}

	for _, v := range input.Data {
		jst, _ := time.LoadLocation("Asia/Tokyo")
		date, _ := time.ParseInLocation("2006-01-02", v.EventDate, jst)
		ryzmEvent := &model.RyzmEvent{
			UUID: v.ID,
		}
		event := &model.Event{
			Name:              v.Title,
			Date:              &date,
			RelatedRyzmEvents: []*model.RyzmEvent{ryzmEvent},
			UnStructuredInformation: &model.UnStructuredEventInformation{
				RyzmUUID:   v.ID,
				ArtistName: v.Artist,
				VenueName:  v.Venue,
				Price:      v.Price,
			},
		}

		if len(v.ReservationSetting.Platforms) > 0 {
			event.TicketURL = v.ReservationSetting.Platforms[0].URL
		}

		result = append(result, event)
	}

	return result, nil
}
