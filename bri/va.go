package bri

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/pandudpn/go-bank-payment"
)

const (
	pathCreateBriVa = "/v1/briva"
)

func CreateVA(param *VAParam, opts *Option) (*bankpayment.BRIVA, *bankpayment.BRIError) {
	var rq = new(request)
	rq.secretKey = opts.secret
	rq.accessToken = opts.accessToken
	rq.httpMethod = http.MethodPost
	rq.body = param
	rq.path = pathCreateBriVa
	rq.url = apiProduction
	rq.timestamp = time.Now().UTC().Format(bankpayment.ISO8601)
	rq.api = opts.api

	if opts.env == bankpayment.Development {
		rq.url = apiDevelopment
	}

	return rq.createVa(context.Background())
}

func (rq *request) createVa(ctx context.Context) (*bankpayment.BRIVA, *bankpayment.BRIError) {
	var (
		r          *briBaseModelResponse
		resBriVa   bankpayment.BRIVA
		httpHeader = make(http.Header)
	)

	signature, err := rq.signature()
	if err != nil {
		return nil, Error(CodeSignature, err.Error())
	}

	httpHeader.Set("BRI-Timestamp", rq.timestamp)
	httpHeader.Set("BRI-Signature", signature)
	httpHeader.Set("Authorization", fmt.Sprintf("Bearer %s", rq.accessToken))

	err = rq.api.Call(ctx, rq.httpMethod, rq.url+rq.path, httpHeader, rq.body, &r)
	if err != nil {
		return nil, Error(CodeHTTPRequest, err.Error())
	}

	briErr := r.parseError()
	if briErr != nil {
		return nil, briErr
	}

	briErr = r.unmarshal(&resBriVa)
	if briErr != nil {
		return nil, briErr
	}

	return &resBriVa, nil
}
