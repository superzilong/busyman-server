package jwt

import (
	"errors"
	"gg/settings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// MyClaims is a jwt claim.
type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenAccessToken returns a pair of access token string and error infomation.
func GenAccessToken(username string) (strAcessToken string, err error) {
	c := MyClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(settings.Conf.JWTConfig.AccessTimeout) * time.Minute).Unix(),
			Issuer:    settings.Conf.JWTConfig.Issuer,
		},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	strAcessToken, err = accessToken.SignedString([]byte(settings.Conf.JWTConfig.MySecret))
	if err != nil {
		return "", err
	}
	return
}

// GenRefreshToken returns a pair of refresh token string and error infomation.
func GenRefreshToken() (strRefreshToken string, err error) {
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Duration(settings.Conf.JWTConfig.RefreshTimeout) * time.Hour).Unix(),
		Issuer:    settings.Conf.JWTConfig.Issuer,
	})
	strRefreshToken, err = refreshToken.SignedString([]byte(settings.Conf.JWTConfig.MySecret))
	if err != nil {
		return "", err
	}
	return
}

func keyFunc(token *jwt.Token) (i interface{}, err error) {
	return []byte(settings.Conf.JWTConfig.MySecret), nil
}

// ParseAccessToken return a pair of MyClaims and error.
func ParseAccessToken(tokenString string) (*MyClaims, error) {
	var c MyClaims
	token, err := jwt.ParseWithClaims(tokenString, &c, keyFunc)

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

// ParseRefreshToken return parse error message.
func ParseRefreshToken(strRefreshToken string) error {
	_, err := jwt.Parse(strRefreshToken, keyFunc)
	return err
}

// IsTokenExpiredErr judges if the input error is jwt.ValidationErrorExpired
func IsTokenExpiredErr(err error) bool {
	v, casterr := err.(*jwt.ValidationError)
	if casterr && v.Errors == jwt.ValidationErrorExpired {
		return true
	}
	return false
}
