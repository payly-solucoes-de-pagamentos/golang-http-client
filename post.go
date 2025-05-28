package httpclient

import "context"

func Post[TBody, TContent any](ctx context.Context, url string, body TBody, options ...ConfigureRequestOptions) HttpResponse[TContent] {
	request := createRequest(ctx, options)

	serializeBody(request, body)

	response, err := request.Post(url)

	model, serializationError := deserializeContent[TContent](response)

	return NewHttpResponse(model, response.StatusCode(), err, serializationError)
}
