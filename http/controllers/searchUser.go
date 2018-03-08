package controllers

import (
	"github.com/astaxie/beego"
	"github.com/shanghai-edu/ldap-test-tool/g"
	"github.com/shanghai-edu/ldap-test-tool/models"
)

type SearchUserController struct {
	beego.Controller
}

type SearchUserResult struct {
	User    models.LDAP_RESULT `json:"user"`
	Success bool               `json:"success"`
}

func (this *SearchUserController) Get() {

	username := this.Ctx.Input.Param(":username")
	user, err := models.Single_Search_User(g.Config().Ldap, username)
	if err != nil {
		var failedResult MsgResult
		failedResult.Msg = err.Error()
		this.Data["json"] = failedResult
	} else {
		var successResult SearchUserResult
		successResult.Success = true
		successResult.User = user
		this.Data["json"] = successResult
	}
	this.ServeJSON()
}
