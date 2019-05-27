package therundown

// Config - holds your authorization
type Config struct {
	// XRapidAPIKey - your rapid api key
	XRapidAPIKey string

	// XRapidAPIHost - version
	XRapidAPIHost string

	// BaseURL - url
	BaseURL string

	// Sport - sport id
	Sport string
}

// NewConfig -
func NewConfig(xRapidAPIKey string) *Config {
	return &Config{
		XRapidAPIKey:  xRapidAPIKey,
		XRapidAPIHost: XRapidAPIHostV1,
		BaseURL:       TheRundownBaseURL,
		Sport:         SportMLB,
	}
}
