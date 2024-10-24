package xcoder

const (
	modelEnvVarName   = "XCODER_MODEL"
	baseURLEnvVarName = "XCODER_BASE_URL"
)

type options struct {
	model   string
	baseURL string
}

type Option func(options *options)

func WithModel(model string) Option {
	return func(options *options) {
		options.model = model
	}
}

func WithBaseURL(baseURL string) Option {
	return func(options *options) {
		options.baseURL = baseURL
	}
}
