package httpclient

import "context"

func Delete(ctx context.Context, url string, options ...ConfigureRequestOptions) HttpResponse[struct{}] {
	request := createRequest(ctx, options)

	response, err := request.Delete(url)

	model, serializationError := deserializeContent[struct{}](response)

	return NewHttpResponse(model, response.StatusCode(), err, serializationError)
}
