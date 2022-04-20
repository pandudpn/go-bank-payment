package bri

import (
	"context"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/pandudpn/go-bank-payment"
)

const (
	pathCreateAccessToken = "/oauth/client_credential/accesstoken?grant_type=client_credentials"
)

func CreateAccessToken(opts *Option) (*bankpayment.BRIAccessToken, *bankpayment.BRIError) {
	var (
		r     = new(request)
		param = make(url.Values)
	)
	param.Set("client_id", opts.key)
	param.Set("client_secret", opts.secret)

	r.secretKey = opts.secret
	r.accessToken = opts.accessToken
	r.httpMethod = http.MethodPost
	r.body = []byte(param.Encode())
	r.path = pathCreateAccessToken
	r.url = apiProduction
	r.timestamp = time.Now().UTC().Format(bankpayment.ISO8601)
	r.api = opts.api

	if opts.env == bankpayment.Development {
		r.url = apiDevelopment
	} else if opts.env == bankpayment.Testing {
		r.url = apiTesting
	}

	return r.createAccessToken(context.Background())
}

func (rq *request) createAccessToken(ctx context.Context) (*bankpayment.BRIAccessToken, *bankpayment.BRIError) {
	var (
		resBriToken *bankpayment.BRIAccessToken
		httpHeader  = make(http.Header)
	)

	httpHeader.Set("Content-Type", "application/x-www-form-urlencoded")

	err := rq.api.Call(ctx, rq.httpMethod, rq.url+rq.path, httpHeader, rq.body, &resBriToken)
	if err != nil {
		log.Println("error", err)
		return nil, Error(CodeHTTPRequest, err.Error())
	}

	return resBriToken, nil
}
