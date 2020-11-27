package main

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

var jwtKey = []byte("heyJwt")

type Claims struct {
	Software string
	OS string
	jwt.StandardClaims
}

type ComputerInfo struct {
	Software string `json:"software"`
	OS string `json:"os"`
}

func main() {
	r := gin.Default()

	r.POST("/token", getToken)
	r.POST("/checktoken", CheckToken)

	r.Run(":8080")
}

func getToken(c *gin.Context)  {
	info := ComputerInfo{}
	err := c.BindJSON(&info)
	c.JSON(200, info)
	if err != nil {
		log.Println("数据绑定失败：", err)
		c.JSON(500, gin.H{
			"msg": "服务器内部错误",
		})
		return
	}
	token, err := info.GenerateToken()
	if err != nil {
		log.Println("token生成失败", err)
		c.JSON(500, gin.H{
			"msg": "服务器内部错误",
		})
		return
	}
	c.JSON(200, gin.H{
		"token": token,
	})
}

func CheckToken(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")

	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
		return
	}
	claims, ok := ParseToken(tokenString)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
		return
	}
	c.JSON(200, gin.H{
		"os": claims.OS,
		"software": claims.Software,
	})
}

//生成token
func (info *ComputerInfo)GenerateToken() (string, error) {
	expireTime := time.Now().Add(7 * 24 * time.Hour)
	claims := Claims{
		Software: info.Software,
		OS: info.OS,
		StandardClaims : jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "lomogo",
		},
	}
	//SigningMethodHS256 注意这里的加密算法，不同的加密算法对应的签名key类型不一致
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	tokenStr , err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

//验证token
func ParseToken(tokenStr string) (*Claims, bool) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		log.Println("解析token失败", err)
		return nil, false
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, true
	} else {
		return nil, false
	}
}


