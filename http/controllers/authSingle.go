package controllers

import (
	"fmt"

	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/shanghai-edu/ldap-test-tool/g"
	"github.com/shanghai-edu/ldap-test-tool/models"
)

type AuthSingleController struct {
	beego.Controller
}

func (this *AuthSingleController) Post() {
	var user models.User
	var msgResult MsgResult
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &user)
	if err != nil {
		msgResult.Msg = err.Error()
		this.Data["json"] = msgResult
		this.ServeJSON()
		return
	}
	if user.Username == "" || user.Password == "" {
		msgResult.Msg = "Missing username or password"
		this.Data["json"] = msgResult
		this.ServeJSON()
		return
	}

	success, err := models.Single_Auth(g.Config().Ldap, user.Username, user.Password)

	if err != nil {
		msgResult.Msg = err.Error()

	} else {
		msgResult.Msg = fmt.Sprintf("user %s Auth Successed", user.Username)
	}
	msgResult.Success = success
	this.Data["json"] = msgResult
	this.ServeJSON()
}
