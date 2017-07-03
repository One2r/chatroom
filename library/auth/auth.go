package auth

import (
	"fmt"

	"github.com/astaxie/beego"
	jwt "github.com/dgrijalva/jwt-go"

	"chatroom/models"
)

//验证token
func CheckToken(tokenStr string) (*models.User, error) {
	var errorVar error
	user := &models.User{}

	if len(tokenStr) == 0 {
		errorVar = fmt.Errorf("请输入token")
	}
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(beego.AppConfig.String("jwt_secret")), nil
	})
	if err != nil {
		errorVar = fmt.Errorf("无效token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		errorVar = fmt.Errorf("获取用户数据失败")
	}
	user.UserID = int(claims["sub"].(map[string]interface{})["uid"].(float64))
	return user, errorVar
}
