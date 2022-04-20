package bankpayment

import "reflect"

type BRIError struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

type BRIAccessToken struct {
	OrganizationName string `json:"organization_name"`
	TokenType        string `json:"token_type"`
	AccessToken      string `json:"access_token"`
	ExpiresIn        string `json:"expires_in"`
	ApplicationName  string `json:"application_name"`
}

type BRIVA struct {
	InstitutionCode string `json:"institutionCode"`
	BrivaNo         string `json:"brivaNo"`
	CustCode        string `json:"custCode"`
	Nama            string `json:"nama"`
	Amount          string `json:"amount"`
	Keterangan      string `json:"keterangan"`
	ExpiredDate     string `json:"expiredDate"`
}

func (be *BRIError) IsNil() bool {
	if be == nil || reflect.ValueOf(be).IsNil() {
		return true
	}

	return reflect.ValueOf(be.Code).IsZero() && reflect.ValueOf(be.Message).IsZero()
}
