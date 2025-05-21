package httpclient

import (
	"crypto/tls"
	"fmt"

	"github.com/go-resty/resty/v2"
)

func setAuth(client *resty.Client, options *HttpRequestOptions) {
	if containsAuthorizationHeader(options.headers) {
		return
	}

	if options.basicAuth != nil {
		setBasicAuth(client, options)
		return
	}

	if options.basicBearer != "" {
		setBasicAuthFromBasicBearer(client, options)
		return
	}

	if options.bearerToken != "" {
		setBearerToken(client, options)
	}
}

func setBasicAuth(client *resty.Client, options *HttpRequestOptions) {
	if options.basicAuth == nil {
		return
	}

	user, password := options.basicAuth.user, options.basicAuth.password

	client.SetBasicAuth(user, password)
}

func setBasicAuthFromBasicBearer(client *resty.Client, options *HttpRequestOptions) {
	if options.basicBearer == "" {
		return
	}
	authHeaderValue := fmt.Sprintf("%s %s", BasicAuthHeaderScheme, options.basicBearer)
	client.SetHeader(AuthorizationHeader, authHeaderValue)
}

func setBearerToken(client *resty.Client, options *HttpRequestOptions) {
	if options.bearerToken == "" {
		return
	}

	client.SetAuthToken(options.bearerToken)
}

func setTLSConfig(client *resty.Client, options *HttpRequestOptions) {
	if !options.tlsInsecure {
		return
	}

	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
}
