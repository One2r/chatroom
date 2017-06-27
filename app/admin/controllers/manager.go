package controllers

import (
	"github.com/astaxie/beego"
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
	if status == "" {
		this.Data["json"] = BizException("参数错误", 600)
		this.ServeJSON()
		this.StopRun()
	}
	beego.Info(status)
}
