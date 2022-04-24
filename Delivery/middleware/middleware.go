package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"strings"
	"time"
)

var ApplicationName = "Kasir"
var JwtSigningMethod = jwt.SigningMethodHS256
var JwtSignatureKey = []byte("@af3nfqwdnqn")

type MyClaim struct {
	jwt.StandardClaims
	Username  string     `json:"username"`
	CreatedAt *time.Time `json:"createdAt"`
}

type authHeader struct {
	Authorization string `header:"authorization"`
}

func AuthTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.Contains(c.Request.URL.Path, "/cashiers") || c.Request.URL.Path == "/cashiers/:cashierId/login" {
			c.Next()
		} else {
			h := authHeader{}
			if err := c.ShouldBindHeader(&h); err != nil {
				c.JSON(401, gin.H{
					"message": "Unauthorized",
				})
				c.Abort()
				return
			}
			tokenString := strings.Replace(h.Authorization, "Bearer ", "", -1)
			fmt.Println(tokenString)
			if tokenString == "" {
				c.JSON(401, gin.H{
					"message": "Unauthorized",
				})
				c.Abort()
				return
			}

			//err := usecase.Authorize(tokenString)
			//if err != nil {
			//	c.JSON(401, gin.H{
			//		"message": "Unauthorized",
			//	})
			//	c.Abort()
			//	return
			//}
			//c.Next()

			token, err := ParseToken(tokenString)
			if err != nil {
				c.JSON(401, gin.H{
					"message": "Unauthorized",
				})
				c.Abort()
				return
			}
			fmt.Println(token)
			if token["iss"] == ApplicationName {
				c.Next()
			} else {
				c.JSON(401, gin.H{
					"message": "Unauthorized",
				})
				c.Abort()
				return
			}
		}
	}
}

func ParseToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("signing method invalid")
		} else if method != JwtSigningMethod {
			return nil, fmt.Errorf("signing method invalid")
		}

		return JwtSignatureKey, nil
	})

	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, err
	}
	return claim, nil
}

func GenerateToken(userName string, createdAt *time.Time) (string, error) {
	claims := MyClaim{
		StandardClaims: jwt.StandardClaims{
			Issuer: ApplicationName,
		},
		Username:  userName,
		CreatedAt: createdAt,
	}
	token := jwt.NewWithClaims(JwtSigningMethod, claims)
	return token.SignedString(JwtSignatureKey)
}
