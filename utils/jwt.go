package utils

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/wagfog/hmdp_go/config/setting"
	"github.com/wagfog/hmdp_go/dto/result"
)

var jwtSecret = []byte(setting.JWTSecret)

type Claims struct {
	Phone string `json:"username"`
	jwt.StandardClaims
}

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int

		code = http.StatusOK
		token := c.Query("token")
		if token == "" {
			code = http.StatusBadRequest
		} else {
			claims, err := ParseToken(token)
			if err != nil {
				code = http.StatusBadRequest
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = http.StatusBadRequest
			}
		}

		if code != http.StatusOK {
			c.JSON(http.StatusUnauthorized, result.Fail("token error"))
			c.Abort()
			return
		}
		c.Next()
	}
}

func GenerateToken(phone string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour) //过期时间

	claims := Claims{
		phone,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "hmdp_go",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	//用于解析鉴权的声明，方法内部主要是具体的解码和校验的过程，最终返回*Token
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
