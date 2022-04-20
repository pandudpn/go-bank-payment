package bri

import (
	"net/http"

	"github.com/pandudpn/go-bank-payment"
)

// Option is parameter needs for API Call
type Option struct {
	key         string
	secret      string
	accessToken string
	env         bankpayment.Environment
	api         bankpayment.APIRequest
}

// NewOption creates new Option parameter's for API Call
// first init will create environment production
func NewOption() *Option {
	httpClient := &http.Client{}
	return &Option{
		env: bankpayment.Production,
		api: &bankpayment.APIImplementation{
			HTTPClient: httpClient,
		},
	}
}

// SetConsumerKey will set key into parameter consumerKey
func (o *Option) SetConsumerKey(key string) {
	o.key = key
}

// GetConsumerKey will return a value of consumerKey
func (o *Option) GetConsumerKey() string {
	return o.key
}

// SetConsumerSecret will set secret-key into parameter consumerSecret
func (o *Option) SetConsumerSecret(secret string) {
	o.secret = secret
}

// GetConsumerSecret will return a value of consumerSecret
func (o *Option) GetConsumerSecret() string {
	return o.secret
}

// SetDevelopment will set environment access to development
func (o *Option) SetDevelopment() {
	o.env = bankpayment.Development
}

// SetProduction will set environment access to production
func (o *Option) SetProduction() {
	o.env = bankpayment.Production
}

// SetTesting will set environment access for unit_test
func (o *Option) SetTesting() {
	o.env = bankpayment.Testing
}

// GetEnvironment will return an active environment
func (o *Option) GetEnvironment() bankpayment.Environment {
	return o.env
}

// SetAccessToken will set accessToken into parameter
// here's will be used when we try to call BRI API
// if the token don't exist, create a new accessToken
// first using method CreateAccessToken()
func (o *Option) SetAccessToken(token string) {
	o.accessToken = token
}

// GetAccessToken will return a value of access-token authorization
func (o *Option) GetAccessToken() string {
	return o.accessToken
}

// SetHTTPClient will set http client into parameter API Call
func (o *Option) SetHTTPClient(httpClient *http.Client) {
	o.api = &bankpayment.APIImplementation{
		HTTPClient: httpClient,
	}
}

// SetAPIRequest will set standard API Request
func (o *Option) SetAPIRequest(api bankpayment.APIRequest) {
	o.api = api
}
