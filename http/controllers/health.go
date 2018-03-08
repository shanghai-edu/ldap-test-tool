package controllers

import (
	"github.com/astaxie/beego"
	"github.com/shanghai-edu/ldap-test-tool/g"
	"github.com/shanghai-edu/ldap-test-tool/models"
)

type HealthController struct {
	beego.Controller
}

func (this *HealthController) Get() {
	var msgResult MsgResult

	success, err := models.Health_Check(g.Config().Ldap)

	if err != nil {
		msgResult.Msg = err.Error()

	} else {
		msgResult.Msg = "ok"
	}
	msgResult.Success = success
	this.Data["json"] = msgResult
	this.ServeJSON()
}
