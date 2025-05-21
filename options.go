package httpclient

type HttpRequestOptions struct {
	headers         map[string]string
	enableRetry     bool
	maximumAttempts int
	wait            int
	maxWait         int
	basicAuth       *basicAuthOptions
	bearerToken     string
	basicBearer     string
	tlsInsecure     bool
}

type basicAuthOptions struct {
	user     string
	password string
}

type ConfigureRequestOptions func(options *HttpRequestOptions)

const defaultEnableRetry bool = false
const defaultMaximumAttempts int = 3
const defaultWait int = 5
const defaultMaxWait int = 20
const defaultBearerToken string = ""
const defaultBasicBearer string = ""
const defaultTLSInsecure bool = false

func newDefaultHeaders() map[string]string {
	headers := make(map[string]string)

	headers[AcceptHeader] = ApplicationJsonHeaderValue

	return headers
}

func newDefaultOptions() HttpRequestOptions {
	defaultHeaders := newDefaultHeaders()

	return HttpRequestOptions{
		headers:         defaultHeaders,
		enableRetry:     defaultEnableRetry,
		maximumAttempts: defaultMaximumAttempts,
		wait:            defaultWait,
		maxWait:         defaultMaxWait,
		basicAuth:       nil,
		bearerToken:     defaultBearerToken,
		basicBearer:     defaultBasicBearer,
		tlsInsecure:     defaultTLSInsecure,
	}
}

func configureOptions(opts []ConfigureRequestOptions) *HttpRequestOptions {
	options := new(HttpRequestOptions)

	if opts != nil {
		for _, opt := range opts {
			opt(options)
		}
	} else {
		*options = newDefaultOptions()
	}

	return options
}

func (options *HttpRequestOptions) guardForMaximumAttempts() *HttpRequestOptions {
	if options.maximumAttempts > 0 {
		return options
	}
	options.maximumAttempts = defaultMaximumAttempts
	return options
}

func (options *HttpRequestOptions) guardForWait() *HttpRequestOptions {
	if options.wait > 0 {
		return options
	}
	options.wait = defaultWait
	return options
}

func (options *HttpRequestOptions) guardForMaxWait() *HttpRequestOptions {
	if options.maxWait > 0 {
		return options
	}
	options.maxWait = defaultMaxWait
	return options
}

func (options *HttpRequestOptions) SetHeaders(headers map[string]string) {
	options.headers = headers
}

func (options *HttpRequestOptions) WithRetryEnabled() *HttpRequestOptions {
	options.enableRetry = true
	return options
}

func (options *HttpRequestOptions) WithRetryDisabled() *HttpRequestOptions {
	options.enableRetry = false
	return options
}

func (options *HttpRequestOptions) WithMaximumAttempts(maximumAttempts int) *HttpRequestOptions {
	options.maximumAttempts = maximumAttempts
	return options
}

func (options *HttpRequestOptions) WithWaitingTime(wait int) *HttpRequestOptions {
	options.wait = wait
	return options
}

func (options *HttpRequestOptions) WithMaxWaitingTime(maxWait int) *HttpRequestOptions {
	options.maxWait = maxWait
	return options
}

func (options *HttpRequestOptions) SetBasicAuth(user, password string) *HttpRequestOptions {
	options.basicAuth = &basicAuthOptions{user: user, password: password}
	return options
}

func (options *HttpRequestOptions) SetBasicAuthFromBasicBearer(basicBearer string) *HttpRequestOptions {
	options.basicBearer = basicBearer
	return options
}

func (options *HttpRequestOptions) SetBearerToken(token string) *HttpRequestOptions {
	options.bearerToken = token
	return options
}

func (options *HttpRequestOptions) AllowInsecureTLS() *HttpRequestOptions {
	options.tlsInsecure = true
	return options
}
