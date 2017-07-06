package controllers

import (
	"github.com/astaxie/beego"

	"chatroom/library/filters/replace"
	"chatroom/library/filters/sensitive"

	"chatroom/library/admin"
	"chatroom/library/auth"
	"chatroom/models"
)

// AdminController handles admin requests.
type AdminController struct {
	beego.Controller
}

func (this *AdminController) Prepare() {
	this.Data["Appname"] = beego.AppConfig.String("appname")
	this.Data["Appver"] = beego.AppConfig.String("appver")
}

// Signin 登录
func (this *AdminController) Signin() {
	username := this.GetString("username")
	password := this.GetString("password")
	this.Data["showMsg"] = false

	if username != "" || password != "" {
		if username == beego.AppConfig.String("admin_username") && password == beego.AppConfig.String("admin_password") {
			token, err := auth.CreateToken(models.User{ID: -1, Type: models.USER_TYPE_ADMIN, Username: username})
			if err != nil {
				this.Abort("500")
			}
			this.Ctx.SetCookie("token", token)
			this.Ctx.Redirect(301, "/admin/dashboard.html")
		} else {
			this.Data["showMsg"] = true
		}
	}
	this.TplName = "admin/signin.tpl"
}

//Signout 登出
func (this *AdminController) Signout() {
	this.Ctx.SetCookie("token", "")
	this.Ctx.Redirect(301, "/admin/signin.html")
}

//Dashboard ...
func (this *AdminController) Dashboard() {
	this.Data["Statis"] = admin.GetStatis()
	this.TplName = "admin/dashboard.tpl"
}

//Service ...
func (this *AdminController) Service() {
	this.Data["Statis"] = admin.GetStatis()
	this.TplName = "admin/service.tpl"
}

//UpdateSensitiveWords 刷新敏感词
func (this *AdminController) UpdateSensitiveWords() {
	result := sensitive.UpdateSensitiveWords()
	this.Data["json"] = AjaxSuccReturn{Data: result}
	this.ServeJSON()
}

//UpdateReplaceWords 刷新替换词
func (this *AdminController) UpdateReplaceWords() {
	result := replace.UpdateReplaceWords()
	this.Data["json"] = AjaxSuccReturn{Data: result}
	this.ServeJSON()
}
