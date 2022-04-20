package bankpayment

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
)

// APIRequest is abstract of HTTP Client that will make API Call
type APIRequest interface {
	Call(ctx context.Context, method, url string, header http.Header, body, result interface{}) error
}

// APIImplementation is the default implementation of HTTP Client Request
type APIImplementation struct {
	HTTPClient *http.Client
}

func (a *APIImplementation) Call(ctx context.Context, method, url string, header http.Header, body, result interface{}) error {
	var (
		err     error
		reqBody []byte
	)

	if body != nil || (reflect.ValueOf(body).Kind() != reflect.Ptr && !reflect.ValueOf(body).IsNil()) {
		switch body.(type) {
		case []byte:
			reqBody = body.([]byte)
		default:
			reqBody, err = json.Marshal(body)
			if err != nil {
				return err
			}
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}

	if header != nil {
		req.Header = header
	}

	req.Header.Set("User-Agent", fmt.Sprintf("go-bankpayment/%s", version))
	if req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "application/json")
	}

	return a.doRequest(req, result)
}

func (a *APIImplementation) doRequest(req *http.Request, result interface{}) error {
	res, err := a.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(resBody, &result)
	if err != nil {
		return err
	}

	return nil
}
