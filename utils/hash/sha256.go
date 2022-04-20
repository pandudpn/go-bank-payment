package hash

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"strings"
)

// SHA256 for hashing data with algorithm SHA-256
func SHA256(value interface{}) (string, error) {
	var data []byte

	switch value.(type) {
	case []byte:
		data = value.([]byte)
	default:
		b, err := json.Marshal(value)
		if err != nil {
			return "", err
		}
		data = b
	}

	sha := sha256.New()
	sha.Write(data)

	s := hex.EncodeToString(sha.Sum(nil))
	return strings.ToLower(s), nil
}

// HmacSHA256 for hashing signature using algorithm SHA-256 with method HMAC
func HmacSHA256(value interface{}, secret string) (string, error) {
	var data []byte

	switch value.(type) {
	case []byte:
		data = value.([]byte)
	case string:
		data = []byte(value.(string))
	default:
		b, err := json.Marshal(value)
		if err != nil {
			return "", err
		}
		data = b
	}

	h := hmac.New(sha256.New, []byte(secret))
	h.Write(data)

	return base64.StdEncoding.EncodeToString(h.Sum(nil)), nil
}
