package httpclient

import (
	"context"

	"github.com/go-resty/resty/v2"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

func createRequest(ctx context.Context, opts []ConfigureRequestOptions) *resty.Request {
	client := resty.New()

	options := configureOptions(opts)

	setTLSConfig(client, options)

	configureRetry(client, options)

	request := client.R()

	setHeaders(options.headers, request)

	setAuth(client, options)

	propagator := otel.GetTextMapPropagator()
	propagator.Inject(ctx, propagation.HeaderCarrier(request.Header))

	return request
}
