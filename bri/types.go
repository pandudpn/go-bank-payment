package bri

import (
	"encoding/json"

	"github.com/pandudpn/go-bank-payment"
)

// briBaseModelResponse is standard response for Virtual Account BRI
// this is can handle for error or success response
type briBaseModelResponse struct {
	Status              bool        `json:"status"`
	ErrDesc             string      `json:"errDesc,omitempty"`
	ResponseDescription string      `json:"responseDescription,omitempty"`
	ResponseCode        string      `json:"responseCode"`
	Data                interface{} `json:"data"`
}

func Error(code, message string) *bankpayment.BRIError {
	return &bankpayment.BRIError{
		Message: message,
		Code:    code,
	}
}

// handle an error
func (b *briBaseModelResponse) parseError() *bankpayment.BRIError {
	var (
		r          *bankpayment.BRIError
		code, desc string
	)

	if !b.Status {
		code = b.ResponseCode
		desc = b.ErrDesc

		if code == "" {
			code = CodeInternalServerError
		}

		if desc == "" {
			desc = ErrInternal
		}

		r = &bankpayment.BRIError{
			Code:    code,
			Message: desc,
		}
	}

	return r
}

// unmarshal used for dynamic parsing response from object key `data`
func (b *briBaseModelResponse) unmarshal(i interface{}) *bankpayment.BRIError {
	dByte, err := json.Marshal(b.Data)
	if err != nil {
		return Error(CodeUnmarshal, err.Error())
	}

	err = json.Unmarshal(dByte, i)
	if err != nil {
		return Error(CodeUnmarshal, err.Error())
	}

	return nil
}
