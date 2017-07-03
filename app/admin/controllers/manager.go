package controllers

import (
	"github.com/astaxie/beego"

	"chatroom/library/admin/manager"
	"chatroom/models"
)

// ManagerController handles manager requests.
type ManagerController struct {
	beego.Controller
}

// SetRoomSilence 设置某个房间的全员禁言状态
func (this *ManagerController) SetRoomSilence() {
	room, err := this.GetInt("room")
	if room <= 0 && err != nil {
		this.Data["json"] = BizException("房间号错误", 600)
		this.ServeJSON()
		this.StopRun()
	}
	status := this.GetString("status")
	if status == "" || (status != "true" && status != "false") {
		this.Data["json"] = BizException("参数错误", 600)
		this.ServeJSON()
		this.StopRun()
	}
	if manager.SetRoomSilence(room, status) {
		this.Data["json"] = AjaxSuccReturn{Data: models.Roomconf[room].Silence}
	} else {
		this.Data["json"] = BizException("设置失败，该房间没有在线人数", 600)
	}
	this.ServeJSON()
	this.StopRun()
}

//SpeakNotAllowed 禁言某个房间的某个人
func (this *ManagerController) SpeakNotAllowed() {
	room, err := this.GetInt("room")
	if room <= 0 && err != nil {
		this.Data["json"] = BizException("房间号错误", 600)
		this.ServeJSON()
		this.StopRun()
	}

	uid, err := this.GetInt("uid")
	if uid <= 0 && err != nil {
		this.Data["json"] = BizException("用户ID错误", 600)
		this.ServeJSON()
		this.StopRun()
	}

	status := this.GetString("status")
	if status == "" || (status != "true" && status != "false") {
		this.Data["json"] = BizException("参数错误", 600)
		this.ServeJSON()
		this.StopRun()
	}

	if manager.SpeakNotAllowed(room, uid, status) {
		this.Data["json"] = AjaxSuccReturn{Data: models.Roomconf[room].SpeakNotAllowed[uid]}
	} else {
		this.Data["json"] = BizException("设置失败，该房间没有在线人数", 600)
	}
	this.ServeJSON()
	this.StopRun()
}
