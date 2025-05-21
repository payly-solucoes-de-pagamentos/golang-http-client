package httpclient

import (
	"time"

	"github.com/go-resty/resty/v2"
	logging "github.com/payly-solucoes-de-pagamentos/golang-logging"
)

type attemptMetadata struct {
	Attempt    int    `json:"attempt"`
	StatusCode string `json:"statusCode"`
	Method     string `json:"method"`
	URL        string `json:"url"`
}

func newAttemptMetada(attempt int, statusCode string, method string, url string) attemptMetadata {
	return attemptMetadata{
		Attempt:    attempt,
		StatusCode: statusCode,
		Method:     method,
		URL:        url,
	}
}

func configureRetry(client *resty.Client, options *HttpRequestOptions) {
	if !options.enableRetry {
		return
	}

	options.guardForMaximumAttempts().guardForWait().guardForMaxWait()

	logger := logging.NewLogger()

	client.SetRetryCount(options.maximumAttempts)

	setWaitTime(options, client)

	setMaxWaitTime(options, client)

	retryHook(client, logger)

	client.AddRetryCondition(retryCondition)
}

func setWaitTime(options *HttpRequestOptions, client *resty.Client) {
	wait := time.Duration(options.wait)

	client.SetRetryWaitTime(wait * time.Second)
}

func setMaxWaitTime(options *HttpRequestOptions, client *resty.Client) {
	maxWait := time.Duration(options.maxWait)

	client.SetRetryMaxWaitTime(maxWait * time.Second)
}

func retryHook(client *resty.Client, logger *logging.Logger) {
	attempts := 0

	client.RetryHooks = append(client.RetryHooks, func(r *resty.Response, err error) {
		attempts++

		metadata := newAttemptMetada(attempts, r.Status(), r.Request.Method, r.Request.URL)

		logger.Standard.Info().Interface("request", metadata).Msgf("[HTTP Client] Retry Attempt: %d", attempts)
	})
}

func retryCondition(response *resty.Response, err error) bool {
	return !response.IsSuccess()
}
