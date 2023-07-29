package utils

import (
	"fmt"
	"github.com/advengulo/go-clean-arch-test/domains"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"os"
	"time"
)

var (
	applicationName = os.Getenv("APPLICATION_NAME")
	jwtKeySecret    = []byte(os.Getenv("JWT_KEY_SECRET"))
)

type Claim struct {
	Username  string    `json:"username"`
	ExpiredAt time.Time `json:"expired_at,omitempty"`
	jwt.RegisteredClaims
}

func CreateToken(payload domains.UserPayload) (domains.Token, error) {
	expirationTime := time.Now().Add(12 * time.Hour)

	claim := &Claim{
		Username: payload.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    applicationName,
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token, err := claim.GetToken()
	if err != nil {
		return domains.Token{}, err
	}

	return domains.Token{
		Token:     token,
		ExpiredAt: expirationTime,
	}, nil
}

func (t *Claim) GetToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, t)

	tokenString, err := token.SignedString(jwtKeySecret)
	if err != nil {
		return "", fmt.Errorf("error create token")
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claim{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKeySecret, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, fmt.Errorf("invalid token")
		}
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return token, nil
}

func GetDataToken(tokenString string) (*Claim, error) {
	token, err := ValidateToken(tokenString)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claim)
	if !ok {
		return nil, fmt.Errorf("invalid token")
	}

	claims.ExpiredAt = claims.ExpiresAt.Time

	return claims, nil
}

func GetHeaderToken(c echo.Context) string {
	return c.Request().Header.Get("Authorization")
}

func IsValidToken(tokenString string) bool {
	_, err := ValidateToken(tokenString)
	if err != nil {
		return false
	}

	return true
}
