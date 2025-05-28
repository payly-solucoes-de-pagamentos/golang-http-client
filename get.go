package httpclient

import "context"

func Get[TContent any](ctx context.Context, url string, options ...ConfigureRequestOptions) HttpResponse[TContent] {
	request := createRequest(ctx, options)

	response, requestError := request.Get(url)

	model, serializationError := deserializeContent[TContent](response)

	return NewHttpResponse(model, response.StatusCode(), requestError, serializationError)
}
