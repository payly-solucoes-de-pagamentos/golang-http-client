package httpclient

import "context"

func Patch[TBody, TContent any](ctx context.Context, url string, body TBody, options ...ConfigureRequestOptions) HttpResponse[TContent] {
	request := createRequest(ctx, options)

	serializeBody(request, body)

	response, err := request.Patch(url)

	model, serializationError := deserializeContent[TContent](response)

	return NewHttpResponse(model, response.StatusCode(), err, serializationError)
}
