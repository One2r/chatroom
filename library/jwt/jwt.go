package jwt

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
)

//验证token
func CheckToken(tokenStr string) (*jwt.Token, error) {
	var errorVar error

	if len(tokenStr) == 0 {
		errorVar = fmt.Errorf("请输入token")
	}
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(beego.AppConfig.String("jwt_secret")), nil
	})
	if err != nil {
		errorVar = fmt.Errorf("无效token")
	}
	return token, errorVar
}
