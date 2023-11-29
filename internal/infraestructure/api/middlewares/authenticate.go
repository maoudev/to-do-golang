package middlewares

import (
	"log/slog"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/maoudev/todo/internal/config"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		strToken := getToken(c)

		jwtToken, err := parseToken(strToken)
		if err != nil {
			slog.Error("Error qqui: " + err.Error())
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		if !isTokenValid(jwtToken) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		setCLaims(c, jwtToken)
	}
}

func isTokenValid(token *jwt.Token) bool {
	return token.Valid
}

func parseToken(token string) (*jwt.Token, error) {
	jwtToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.SECRET_KEY), nil
	})

	return jwtToken, err
}

func getToken(c *gin.Context) string {
	header := c.Request.Header.Get("Authorization")

	strToken := strings.TrimPrefix(header, "Bearer ")

	return strToken
}

func setCLaims(c *gin.Context, jwtToken *jwt.Token) {
	claims := jwtToken.Claims.(jwt.MapClaims)

	c.Set("userID", claims["ID"])
}
