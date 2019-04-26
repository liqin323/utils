package access_token

import (
	"time"

	jwt "utils/jwt"
)

type accessToken struct {
}

const (
	ACCESS_TOKEN_EXPIRED  = time.Hour * 24 * 365 * 10  // 10 years for testing
	REFRESH_TOKEN_EXPIRED = time.Hour * 24 * 365 * 100 // 100 years means never
)

var AccessToken *accessToken = new(accessToken)

const token_key = "yhi7TQIDAQABAoIBACa3uMjjVI6rz4zy723Bvgr"

func (acsTk *accessToken) Signed(tkID, tkType string, expired int64) (acAcsTk string, err error) {

	// Create the token
	jwt_tk := jwt.New(jwt.GetSigningMethod("HS256"))

	// Set some claims
	jwt_tk.Claims["id"] = tkID
	jwt_tk.Claims["type"] = tkType
	jwt_tk.Claims["exp"] = expired

	// Sign and get the complete encoded token as a string
	_snx_tk, err := jwt_tk.SignedString([]byte(token_key))
	if err != nil {
		return
	}

	acAcsTk = _snx_tk
	return
}

func (acsTk *accessToken) Verify(acAcsTk string) (err error) {

	_, err = jwt.Parse(acAcsTk, func(token *jwt.Token) (interface{}, error) {
		return []byte(token_key), nil
	})

	return
}

func (acsTk *accessToken) IsExpired(acAcsTk string) (is_exp bool, err error) {

	jwt_tk, err := jwt.Parse(acAcsTk, func(token *jwt.Token) (interface{}, error) {
		return []byte(token_key), nil
	})

	if err != nil {
		is_exp = true
		return
	}

	expired := int64(jwt_tk.Claims["exp"].(float64))

	now := time.Now().Unix()
	if now > expired {
		is_exp = true
	}

	return
}

func (acsTk *accessToken) GetID(acAcsTk string) (tkID string, err error) {

	jwt_tk, err := jwt.Parse(acAcsTk, func(token *jwt.Token) (interface{}, error) {
		return []byte(token_key), nil
	})

	if err != nil {
		return
	}

	tkID = jwt_tk.Claims["id"].(string)
	return
}

func (acsTk *accessToken) GetType(acAcsTk string) (tkType string, err error) {

	jwt_tk, err := jwt.Parse(acAcsTk, func(token *jwt.Token) (interface{}, error) {
		return []byte(token_key), nil
	})

	if err != nil {
		return
	}

	tkType = jwt_tk.Claims["type"].(string)
	return
}

func (acsTk *accessToken) GetExpired(acAcsTk string) (expired int64, err error) {

	jwt_tk, err := jwt.Parse(acAcsTk, func(token *jwt.Token) (interface{}, error) {
		return []byte(token_key), nil
	})

	if err != nil {
		return
	}

	expired = int64(jwt_tk.Claims["exp"].(float64))
	return
}
