package controllers

import (
	"github.com/astaxie/beego"

	"chatroom/library/filters/sensitive"
)

type ApiController struct {
	beego.Controller
}

//刷新敏感词
func (this *ApiController) UpdateSensitiveWords() {
	result := sensitive.UpdateSensitiveWords()
	this.Data["json"] = result
    this.ServeJSON()
}