package auth

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
	jwt "github.com/dgrijalva/jwt-go"

	"chatroom/models"
)

//CheckToken 验证token
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
	user.ID = int(claims["sub"].(map[string]interface{})["ID"].(float64))
	user.Type = claims["sub"].(map[string]interface{})["Type"].(string)
	user.Username = claims["sub"].(map[string]interface{})["Username"].(string)
	return user, errorVar
}

//CreateToken 创建后台管理员token
func CreateToken(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"nbf": int64(time.Now().Unix()),
		"exp": int64(time.Now().Unix() + 1000),
		"iss": beego.AppConfig.String("appname"),
		"sub": user,
	})
	tokenStr, err := token.SignedString([]byte(beego.AppConfig.String("jwt_secret")))
	return tokenStr, err
}
