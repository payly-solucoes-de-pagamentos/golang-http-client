package httpclient

import (
	"context"
	"io"
)

type FileMetadata struct {
	Content  io.Reader
	Param    string
	FileName string
}

func UploadFile[TResponse any](ctx context.Context, url string, fileMetadata FileMetadata, options ...ConfigureRequestOptions) HttpResponse[TResponse] {
	request := createRequest(ctx, options)

	request.SetFileReader(fileMetadata.Param, fileMetadata.FileName, fileMetadata.Content)

	response, requestError := request.Post(url)

	model, serializationError := deserializeContent[TResponse](response)

	return NewHttpResponse(model, response.StatusCode(), requestError, serializationError)
}
