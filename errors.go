package httpclient

type HttpClientOptionsError struct {
	Message string
}

func (err HttpClientOptionsError) Error() string {
	return err.Message
}

type DeserializationError struct {
	Message string
}

func (err DeserializationError) Error() string {
	return err.Message
}
