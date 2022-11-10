package selializer

import (
	"github.com/scarlet0725/prism-api/model"
)

type ResponseSerializer interface {
	BuildResponse(interface{}) (model.APIResponse, error)
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
