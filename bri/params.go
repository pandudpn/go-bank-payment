package bri

import (
	"encoding/json"
	"fmt"

	bankpayment "github.com/pandudpn/go-bank-payment"
	"github.com/pandudpn/go-bank-payment/utils/hash"
)

type request struct {
	secretKey   string
	url         string
	path        string
	accessToken string
	httpMethod  string
	timestamp   string
	api         bankpayment.APIRequest
	body        interface{}
}

// signature to create a signature request
func (rq *request) signature() (string, error) {
	var b []byte
	switch rq.body.(type) {
	case []byte:
		b = rq.body.([]byte)
	default:
		byt, err := json.Marshal(rq.body)
		if err != nil {
			return "", err
		}

		b = byt
	}

	val := fmt.Sprintf("path=%s&verb=%s&token=%s&timestamp=%s&body=%s", rq.path, rq.httpMethod, fmt.Sprintf("Bearer %s", rq.accessToken), rq.timestamp, string(b))
	val, err := hash.HmacSHA256(val, rq.secretKey)
	if err != nil {
		return "", err
	}

	return val, nil
}

// VAParam contains data request parameter's to Create BRI Virtual Account
type VAParam struct {
	InstitutionCode string  `json:"institutionCode"`
	BriVANo         string  `json:"brivaNo"`
	CustCode        string  `json:"custCode"`
	Nama            string  `json:"nama"`
	Amount          float64 `json:"amount"`
	Keterangan      string  `json:"keterangan"`
	ExpiredDate     string  `json:"expiredDate"`
}
