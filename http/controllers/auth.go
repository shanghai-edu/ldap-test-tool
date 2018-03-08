package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/shanghai-edu/ldap-test-tool/g"
	"github.com/shanghai-edu/ldap-test-tool/models"
)

type AuthController struct {
	beego.Controller
}

type MsgResult struct {
	Msg     string `json:"msg"`
	Success bool   `json:"success"`
}

func (this *AuthController) Post() {
	var msgResult MsgResult
	username := this.Ctx.Input.Param(":username")
	password := this.GetString("password")
	if password == "" {
		msgResult.Msg = ("Missing password")
		this.Data["json"] = msgResult
		this.ServeJSON()
		return
	}
	success, err := models.Single_Auth(g.Config().Ldap, username, password)

	if err != nil {
		msgResult.Msg = err.Error()

	} else {
		msgResult.Msg = fmt.Sprintf("user %s Auth Successed", username)
	}
	msgResult.Success = success
	this.Data["json"] = msgResult
	this.ServeJSON()
}
