package mid

import (
	"mytechblog/utils"
	msg "mytechblog/utils/errormsg"
	"strings"
	"time"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var JwtKey = []byte(utils.JwtKey)

type Claims struct {
	Username string `json:"username"`
	// Password string `json:"password"`
	jwt.StandardClaims
}
// 生成token
func SetToken(username string/*, password string*/)(string, int){
	expireTime := time.Now().Add(time.Hour)
	setClaims := Claims{
		username, 
		// password,
		jwt.StandardClaims {
			ExpiresAt: expireTime.Unix(),
			Issuer: "kytolly",
		},
	}
	req := jwt.NewWithClaims(jwt.SigningMethodHS256, setClaims)
	token, err := req.SignedString(JwtKey)
	if err != nil{
		return "  ", msg.ERROR
	}
	return token, msg.SUCCESS
}

// 验证token
func CheckToken(token string)(*Claims, int){
	setToken,_ :=jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token)(interface{}, error){
		return JwtKey, nil
	})
	if key, ok := setToken.Claims.(*Claims); ok&&setToken.Valid{
		return key, msg.SUCCESS
	}else{
		return nil, msg.ERROR
	}
}
// jwt中间件
func JwtToken()gin.HandlerFunc {
	return func(c *gin.Context){
		tokenHeader := c.Request.Header.Get("Authorization")
		status := msg.SUCCESS
		if tokenHeader == ""{
			status = msg.ERROR_TOKEN_NOT_EXIST
			c.JSON(http.StatusOK, gin.H{
				"status": status,
				"message": msg.GetErrorMessage(status),
			})
			c.Abort()
			return 
		}
		checkToken := strings.SplitN(tokenHeader, " ", 2)
		if len(checkToken)!=2 && checkToken[0]!= "Bearer "{
			status = msg.ERROR_TOKEN_TYPE_WRONG
			c.JSON(http.StatusOK, gin.H{
				"status": status,
				"message": msg.GetErrorMessage(status),
			})
			c.Abort()
			return 
		}
		key, ok := CheckToken(checkToken[1])
		if ok == msg.ERROR{
			status = msg.ERROR_TOKEN_WRONG
			c.JSON(http.StatusOK, gin.H{
				"status": status,
				"message": msg.GetErrorMessage(status),
			})
			c.Abort()
			return 
		}
		if time.Now().Unix() > key.ExpiresAt {
			status = msg.ERROR_TOKEN_INVALID
			c.JSON(http.StatusOK, gin.H{
				"status": status,
				"message": msg.GetErrorMessage(status),
			})
			c.Abort()
			return 
		}
		c.Set("username", key.Username)
		c.Next()
	}
}