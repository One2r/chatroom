package auth

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
	jwt "github.com/dgrijalva/jwt-go"

	"chatroom/models"
)

//CheckWSToken 验证websocket 连接token
func CheckWSToken(tokenStr string) (*models.User, error) {
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

//CreateAdminToken 创建后台管理员token
func CreateAdminToken(username string) (string, error) {
	claims := &jwt.StandardClaims{
		NotBefore: int64(time.Now().Unix()),
		ExpiresAt: int64(time.Now().Unix() + 1000),
		Issuer:    beego.AppConfig.String("appname"),
		Subject:   username,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(beego.AppConfig.String("jwt_secret")))
	return tokenStr, err
}

//CheckAdminToken 验证后台管理员token
func CheckAdminToken(tokenStr string) (username string, errorVar error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(beego.AppConfig.String("jwt_secret")), nil
	})
	if err != nil {
		errorVar = fmt.Errorf("无效token")
		return
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		errorVar = fmt.Errorf("获取用户数据失败")
		return
	}
	username = claims["sub"].(string)
	return
}
