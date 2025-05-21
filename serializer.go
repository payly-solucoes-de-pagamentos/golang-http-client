package httpclient

import (
	"encoding/json"

	"github.com/go-resty/resty/v2"
)

func deserializeContent[TContent any](response *resty.Response) (*TContent, error) {
	if !shouldSerialize(response) {
		return nil, nil
	}

	model := new(TContent)

	if err := json.Unmarshal(response.Body(), &model); err != nil {
		deserializationError := DeserializationError{Message: err.Error()}
		return model, deserializationError
	}

	return model, nil
}

func serializeBody[TBody any](request *resty.Request, body TBody) {
	request.SetBody(body)
}

func shouldSerialize(response *resty.Response) bool {
	body := response.Body()
	return body != nil && len(body) > 0
}
