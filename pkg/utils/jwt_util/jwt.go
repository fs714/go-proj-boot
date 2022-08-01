package jwt_util

import (
	"time"

	"github.com/fs714/go-proj-boot/pkg/utils/config"
	"github.com/golang-jwt/jwt/v4"
	"github.com/oklog/ulid/v2"
	"github.com/pkg/errors"
)

var TokenInvalid = errors.New("invalid token")

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type JwtUtil struct {
	SigningKey     []byte
	ExpiresTime    int
	BufferTime     int
	CookieName     string
	CookiePath     string
	CookieDomain   string
	SecurityCookie bool
}

func NewJwtUtil() *JwtUtil {
	return &JwtUtil{
		SigningKey:     []byte(config.Config.Jwt.Secret),
		ExpiresTime:    config.Config.Jwt.ExpiresTime,
		BufferTime:     config.Config.Jwt.BufferTime,
		CookieName:     config.Config.Jwt.CookieName,
		CookiePath:     config.Config.Jwt.CookiePath,
		CookieDomain:   config.Config.Jwt.CookieDomain,
		SecurityCookie: config.Config.Jwt.SecurityCookie,
	}
}

func (j *JwtUtil) CreateClaims(username string) Claims {
	claims := Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "Config.Issuer",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(j.ExpiresTime) * time.Second)),
			NotBefore: jwt.NewNumericDate(time.Now().Add(600 * time.Second)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        ulid.Make().String(),
		},
	}

	return claims
}

func (j *JwtUtil) CreateToken(claims Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := token.SignedString(j.SigningKey)
	if err != nil {
		return "", errors.Wrap(err, "sign string by jwt token failed")
	}

	return signedString, nil
}

func (j *JwtUtil) CreateTokenByOldToken(oldToken string, claims Claims) (string, error) {
	return j.CreateToken(claims)
}

func (j *JwtUtil) ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.Wrap(ve, "token is malformed")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.Wrap(ve, "token is expired")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.Wrap(ve, "token is not active yet")
			} else {
				return nil, errors.Wrap(ve, "token is invalid")
			}
		} else {
			return nil, errors.Wrap(ve, "token is not a ValidationError")
		}
	}

	if token != nil {
		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			return claims, nil
		}

		return nil, TokenInvalid
	} else {
		return nil, TokenInvalid
	}
}
