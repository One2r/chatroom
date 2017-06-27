package controllers

import (
	"github.com/astaxie/beego"

	"chatroom/library/filters/replace"
	"chatroom/library/filters/sensitive"
)

// AdminController handles admin requests.
type AdminController struct {
	beego.Controller
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
