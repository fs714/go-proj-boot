package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/fs714/go-proj-boot/pkg/utils/jwt_util"
	"github.com/fs714/go-proj-boot/pkg/utils/log"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/securecookie"
	"github.com/pkg/errors"
)

var (
	hashKey  = []byte("a719b51e882d2992")
	blockKey = []byte("b69e3a9ac3a3f65f")
	sc       = securecookie.New(hashKey, blockKey)
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		j := jwt_util.NewJwtUtil()

		token, err := tokenFromCookie(c, j.CookieName, j.SecurityCookie)
		if err != nil {
			errMsg := "failed to get token from cookie"
			log.Errorw(errMsg, "err", err)

			c.JSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  errMsg,
				"data": "",
			})

			c.Abort()
			return
		}

		claims, err := j.ParseToken(token)
		if err != nil {
			errMsg := "failed to parse token"
			log.Errorw(errMsg, "err", err)

			c.JSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  errMsg,
				"data": "",
			})

			c.Abort()
			return
		}

		if claims.ExpiresAt.Unix()-time.Now().Unix() < int64(j.BufferTime) {
			claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Duration(j.ExpiresTime) * time.Second))
			newToken, _ := j.CreateTokenByOldToken(token, *claims)
			err = tokenToCookie(c, j, newToken)
			if err != nil {
				errMsg := "failed to set token to cookie"
				log.Errorw(errMsg, "err", err)

				c.JSON(http.StatusUnauthorized, gin.H{
					"code": http.StatusUnauthorized,
					"msg":  errMsg,
					"data": "",
				})

				c.Abort()
				return
			}
		}

		c.Set("claims", claims)
		c.Next()
	}
}

var (
	EmptyTokenFromCookie = errors.New("empty token from cookie")
	EmptyTokenFromHeader = errors.New("empty token from header")
	EmptyTokenFromQuery  = errors.New("empty token from query")
	InvalidTokenHeader   = errors.New("invalid token header")
)

func tokenFromCookie(c *gin.Context, cookieName string, secureCookie bool) (string, error) {
	cookie, err := c.Cookie(cookieName)
	if err != nil {
		return "", errors.Wrapf(err, "failed to get cookie: %s", cookieName)
	}

	if cookie == "" {
		return "", EmptyTokenFromCookie
	}

	if secureCookie {
		var decodedCookie string
		err = sc.Decode(cookieName, cookie, &decodedCookie)
		if err != nil {
			return "", errors.Wrapf(err, "failed to decode token for cookie: %s", cookieName)
		}

		cookie = decodedCookie
	}

	return cookie, nil
}

func tokenToCookie(c *gin.Context, j *jwt_util.JwtUtil, token string) error {
	if j.SecurityCookie {
		encodedToken, err := sc.Encode(j.CookieDomain, token)
		if err != nil {
			return errors.Wrapf(err, "failed to encode token for cookie: %s", j.CookieName)
		}

		c.SetCookie(
			j.CookieName,
			encodedToken,
			j.ExpiresTime,
			j.CookiePath,
			j.CookieDomain,
			false,
			true,
		)

		return nil
	} else {
		c.SetCookie(
			j.CookieName,
			token,
			j.ExpiresTime,
			j.CookiePath,
			j.CookieDomain,
			false,
			true,
		)

		return nil
	}
}

// Authorization:Bearer xxxxxxxxx, headerName is Authorization, tokenHeaderName is Bearer
func tokenFromHeader(c *gin.Context, headerName string, tokenHeaderName string) (string, error) {
	authHeader := c.Request.Header.Get(headerName)

	if authHeader == "" {
		return "", EmptyTokenFromHeader
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == tokenHeaderName) {
		return "", InvalidTokenHeader
	}

	return parts[1], nil
}

func tokenFromQuery(c *gin.Context, queryName string) (string, error) {
	token := c.Query(queryName)

	if token == "" {
		return "", EmptyTokenFromQuery
	}

	return token, nil
}
