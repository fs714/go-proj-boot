package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/fs714/go-proj-boot/pkg/utils/jwt_util"
	"github.com/fs714/go-proj-boot/pkg/utils/log"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := tokenFromCookie(c, "todoCookieName")
		if err != nil {
			log.Errorf("%+v", err)

			c.JSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  err.Error(),
				"data": "",
			})

			c.Abort()
			return
		}

		j := jwt_util.NewJwtUtil()
		claims, err := j.ParseToken(token)
		if err != nil {
			log.Errorf("%+v", err)

			c.JSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  err.Error(),
				"data": "",
			})

			c.Abort()
			return
		}

		bufferTime := 9999
		ExpiresTime := 9999
		if claims.ExpiresAt.Unix()-time.Now().Unix() < int64(bufferTime) {
			claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Duration(ExpiresTime) * time.Second))
			newToken, _ := j.CreateTokenByOldToken(token, *claims)
			tokenToCookie(c, "todoCookieName", newToken)
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

func tokenFromCookie(c *gin.Context, cookieName string) (string, error) {
	cookie, err := c.Cookie(cookieName)
	if err != nil {
		return "", errors.Wrapf(err, "failed to get cookie: %s", cookieName)
	}

	if cookie == "" {
		return "", EmptyTokenFromCookie
	}

	return cookie, nil
}

func tokenToCookie(c *gin.Context, cookieName string, token string) {
	c.SetCookie(
		cookieName,
		token,
		9999,
		"/+todo_path",
		"todo_CookieDomain",
		false,
		true,
	)
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
