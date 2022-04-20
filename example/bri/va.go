package main

import (
	"fmt"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/pandudpn/go-bank-payment/bri"
)

func main() {
	clientId := os.Getenv("BRI_CLIENT_ID")
	clientSecret := os.Getenv("BRI_CLIENT_SECRET")
	institutionCode := os.Getenv("BRI_INSTITUTION_CODE")

	opts := bri.NewOption()
	opts.SetConsumerKey(clientId)
	opts.SetConsumerSecret(clientSecret)
	opts.SetDevelopment()

	token, briErr := bri.CreateAccessToken(opts)
	if briErr != nil {
		panic(briErr)
	}

	opts.SetAccessToken(token.AccessToken)

	param := &bri.VAParam{
		Amount:          10000,
		BriVANo:         "77777",
		CustCode:        "8212227418",
		InstitutionCode: institutionCode,
		Nama:            "Pandu dwi Putra Nugroho",
		Keterangan:      "Contoh",
		ExpiredDate:     time.Now().Add(6 * time.Hour).Format("2006-01-02 15:04:05"),
	}

	va, briErr := bri.CreateVA(param, opts)
	if briErr != nil {
		panic(briErr)
	}

	fmt.Println("success create va", va)
}
