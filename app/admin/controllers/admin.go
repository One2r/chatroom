package controllers

import (
	"github.com/astaxie/beego"

	"chatroom/library/filters/replace"
	"chatroom/library/filters/sensitive"

	"chatroom/library/admin"
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
			this.SetSession("isLogin", username)
			this.Ctx.Redirect(301, "/admin/dashboard")
		} else {
			this.Data["showMsg"] = true
		}
	}
	this.TplName = "admin/signin.tpl"
}

//Signout 登出
func (this *AdminController) Signout() {
	this.DelSession("isLogin")
	this.Ctx.Redirect(301, "/admin/signin")
}

//Dashboard ...
func (this *AdminController) Dashboard() {
	this.Data["Statis"] = admin.GetStatis()
	this.Data["isLogin"] = this.GetSession("isLogin").(string)
	this.TplName = "admin/dashboard.tpl"
}

//UpdateSensitiveWords 刷新敏感词
func (this *AdminController) UpdateSensitiveWords() {
	result := sensitive.UpdateSensitiveWords()
	this.Data["json"] = result
	this.ServeJSON()
}

//UpdateReplaceWords 刷新替换词
func (this *AdminController) UpdateReplaceWords() {
	result := replace.UpdateReplaceWords()
	this.Data["json"] = result
	this.ServeJSON()
}
