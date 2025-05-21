package httpclient

import "github.com/go-resty/resty/v2"

const AuthorizationHeader string = "Authorization"
const AcceptHeader string = "Accept"

const BearerAuthHeaderScheme string = "Bearer"
const BasicAuthHeaderScheme string = "Basic"

const ApplicationJsonHeaderValue string = "application/json"

func isEmptyHeaders(headers map[string]string) bool {
	return len(headers) <= 0
}

func setHeaders(headers map[string]string, request *resty.Request) {
	if isEmptyHeaders(headers) {
		return
	}
	request.SetHeaders(headers)
}

func containsAuthorizationHeader(headers map[string]string) bool {
	if isEmptyHeaders(headers) {
		return false
	}

	for headerKey := range headers {
		if headerKey == AuthorizationHeader {
			return true
		}
	}
	return false
}
