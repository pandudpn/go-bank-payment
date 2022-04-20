package bri_test

import (
	"net/http"
	"testing"

	"github.com/pandudpn/go-bank-payment"
	"github.com/pandudpn/go-bank-payment/bri"
)

func TestNewOption(t *testing.T) {
	t.Run("testing options", func(t *testing.T) {
		c := &bankpayment.APIImplementation{HTTPClient: &http.Client{}}
		opts := bri.NewOption()
		opts.SetConsumerSecret("abc")
		opts.SetAccessToken("abc")
		opts.SetDevelopment()
		opts.SetProduction()
		opts.SetTesting()
		opts.SetConsumerKey("abc")
		opts.GetConsumerSecret()
		opts.GetConsumerKey()
		opts.GetAccessToken()
		opts.GetEnvironment()
		opts.SetAPIRequest(c)
		opts.SetHTTPClient(&http.Client{})
	})
}
