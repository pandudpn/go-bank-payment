package bri_test

import (
	"testing"

	"github.com/pandudpn/go-bank-payment"
	"github.com/pandudpn/go-bank-payment/bri"
	"github.com/stretchr/testify/assert"
)

func TestCreateAccessToken(t *testing.T) {
	testCases := []struct {
		name           string
		env            bankpayment.Environment
		expectedResult *bankpayment.BRIAccessToken
		expectedError  *bankpayment.BRIError
	}{
		{
			name:           "Success",
			env:            bankpayment.Development,
			expectedResult: &bankpayment.BRIAccessToken{},
			expectedError:  nil,
		},
		{
			name:           "error",
			env:            bankpayment.Testing,
			expectedResult: nil,
			expectedError:  &bankpayment.BRIError{Message: "Post \"/oauth/client_credential/accesstoken?grant_type=client_credentials\": unsupported protocol scheme \"\"", Code: "92"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			opts := bri.NewOption()
			opts.SetConsumerSecret("secret")
			opts.SetConsumerKey("key")

			if tc.env == bankpayment.Development {
				opts.SetDevelopment()
			} else if tc.env == bankpayment.Testing {
				opts.SetTesting()
			}

			res, err := bri.CreateAccessToken(opts)

			assert.Equal(t, tc.expectedResult, res)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}
