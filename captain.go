package bulwarkadmin

import "net/http"

// Captain is the main client struct to use with bulwarkadmin
type Captain struct {
	Account *Account
	baseUrl string
}

// NewCaptain initializer
func NewCaptain(baseUrl string, client *http.Client) *Captain {
	return &Captain{
		Account: NewAccountClient(baseUrl, client),
		baseUrl: baseUrl,
	}
}
