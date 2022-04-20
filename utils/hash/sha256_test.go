package hash_test

import (
	"errors"
	"testing"

	"github.com/pandudpn/go-bank-payment/utils/hash"
	"github.com/stretchr/testify/assert"
)

func TestSHA256(t *testing.T) {
	testCases := []struct {
		name           string
		value          interface{}
		expectedResult string
		expectedError  error
	}{
		{
			name:           "Success sha256 with value string",
			value:          "test",
			expectedResult: "4d967a30111bf29f0eba01c448b375c1629b2fed01cdfcc3aed91f1b57d5dd5e",
			expectedError:  errors.New(""),
		},
		{
			name:           "Success sha256 with value integer",
			value:          123,
			expectedResult: "a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3",
			expectedError:  errors.New(""),
		},
		{
			name:           "Success sha256 with value array of byte",
			value:          []byte("test"),
			expectedResult: "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
			expectedError:  errors.New(""),
		},
		{
			name:           "error marshal",
			value:          make(chan int, 1),
			expectedResult: "",
			expectedError:  errors.New("json: unsupported type: chan int"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := hash.SHA256(tc.value)

			assert.Equal(t, tc.expectedResult, res)
			assert.Error(t, tc.expectedError, err)
		})
	}
}

func TestHmacSHA256(t *testing.T) {
	testCases := []struct {
		name           string
		value          interface{}
		secret         string
		expectedResult string
		expectedError  error
	}{
		{
			name:           "Success hmac with value string",
			value:          "test",
			secret:         "abc",
			expectedResult: "1kzPDUsUSRU9eCFfn/m5DsNzDeH9KzV+FTAmyaP62pY=",
			expectedError:  errors.New(""),
		},
		{
			name:           "Success hmac with value integer",
			value:          123,
			secret:         "abc",
			expectedResult: "a6pSztU5etJqsDWidxigdvu3hVtmtxhYhnJU3n7nN2Y=",
			expectedError:  errors.New(""),
		},
		{
			name:           "Success hmac with value array of byte",
			value:          []byte("test"),
			secret:         "abc",
			expectedResult: "1kzPDUsUSRU9eCFfn/m5DsNzDeH9KzV+FTAmyaP62pY=",
			expectedError:  errors.New(""),
		},
		{
			name:           "error marshal",
			value:          make(chan int, 1),
			secret:         "abc",
			expectedResult: "",
			expectedError:  errors.New("json: unsupported type: chan int"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := hash.HmacSHA256(tc.value, tc.secret)

			assert.Equal(t, tc.expectedResult, res)
			assert.Error(t, tc.expectedError, err)
		})
	}
}
