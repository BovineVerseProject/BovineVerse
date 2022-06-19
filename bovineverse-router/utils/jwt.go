package utils

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired     error = fmt.Errorf("Token is expired")
	TokenNotValidYet error = fmt.Errorf("Token not active yet")
	TokenMalformed   error = fmt.Errorf("That's not even a token")
	TokenInvalid     error = fmt.Errorf("Couldn't handle this token:")
	secretKey              = "NONSENSE123456789abcdefghijklmn!"
)

type CustomClaims struct {
	ID      int    `json:"userId"`
	Name    string `json:"name"`
	Address string `json:"address"`
	jwt.StandardClaims
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(GetSignKey()),
	}
}

func GetSignKey() string {
	return secretKey
}

func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}

func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}

func GenerateToken(id int, address string) (token string, err error) {
	j := NewJWT()
	claims := CustomClaims{
		ID:      id,
		Address: address,
	}

	now := time.Now()
	claims.StandardClaims = jwt.StandardClaims{
		NotBefore: now.Unix(),
	}
	claims.StandardClaims.ExpiresAt = now.Unix() + 900
	token, err = j.CreateToken(claims)
	return
}
