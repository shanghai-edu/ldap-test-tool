package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/shanghai-edu/ldap-test-tool/g"
	"github.com/shanghai-edu/ldap-test-tool/models"
)

type AuthMultiController struct {
	beego.Controller
}

type MsgResult struct {
	Msg     string `json:"msg"`
	Success bool   `json:"success"`
}

type AuthResult struct {
	Success bool                     `json:"success"`
	Result  models.Multi_Auth_Result `json:"result"`
}

func (this *AuthMultiController) Post() {
	var userlist []models.User
	var msgResult MsgResult
	var authResult AuthResult
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &userlist)
	if err != nil {
		msgResult.Msg = err.Error()
		this.Data["json"] = msgResult
		this.ServeJSON()
		return
	}

	result, err := models.Multi_Auth(g.Config().Ldap, userlist)

	if err != nil {
		msgResult.Msg = err.Error()
		this.Data["json"] = msgResult
		this.ServeJSON()
		return
	}
	authResult.Success = true
	authResult.Result = result
	this.Data["json"] = authResult
	this.ServeJSON()
}
