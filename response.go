package httpclient

import "reflect"

type httpResponseError struct {
	key   string
	value string
}

type HttpResponse[TContent any] struct {
	StatusCode int               `json:"statusCode,omitempty"`
	Content    *TContent         `json:"content,omitempty"`
	Errors     map[string]string `json:"errors,omitempty"`
}

func buildErrorMap(errors []error) map[string]string {

	errs := make(map[string]string)

	for _, err := range errors {
		if err == nil {
			continue
		}
		responseError := newHttpResponseError(err)
		errs[responseError.key] = responseError.value
	}

	return errs
}

func newHttpResponseError(err error) httpResponseError {
	errorName := reflect.TypeOf(err).Name()
	if errorName == "" {
		errorName = "Error"
	}
	return httpResponseError{key: errorName, value: err.Error()}
}

func NewHttpResponse[TContent any](
	content *TContent,
	statusCode int,
	errors ...error) HttpResponse[TContent] {

	errs := buildErrorMap(errors)

	return HttpResponse[TContent]{
		StatusCode: statusCode,
		Content:    content,
		Errors:     errs,
	}
}
