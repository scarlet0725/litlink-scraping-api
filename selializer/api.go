package selializer

import (
	"time"

	"github.com/scarlet0725/prism-api/model"
)

type ResponseSerializer interface {
	BuildResponse(interface{}) (model.APIResponse, error)
	SelializeRyzmData(model.RyzmAPIResponse) ([]model.Event, error)
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

func (a *apiResponse) SelializeRyzmData(input model.RyzmAPIResponse) ([]model.Event, error) {
	var result []model.Event

	for _, v := range input.Data {
		jst, _ := time.LoadLocation("Asia/Tokyo")
		date, _ := time.ParseInLocation("2006-01-02", v.EventDate, jst)
		result = append(result, model.Event{
			UUID:      v.ID,
			Name:      v.Title,
			Artist:    v.Artist,
			Date:      date,
			Venue:     v.Venue,
			TicketURL: v.ReservationSetting.Platforms[0].URL,
		})
	}

	return result, nil
}
