package controllers

import (
	"github.com/astaxie/beego"

	"chatroom/library/badwords"
)

type ApiController struct {
	beego.Controller
}

//刷新敏感词
func (this *ApiController) UpdateBadword() {
	result := badwords.UpdateBadword()
	this.Data["json"] = result
    this.ServeJSON()
}