package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/shanghai-edu/ldap-test-tool/g"
	"github.com/shanghai-edu/ldap-test-tool/models"
)

type SearchMultiController struct {
	beego.Controller
}

type SearchMultiResult struct {
	Success bool                            `json:"success"`
	Result  models.Multi_Search_User_Result `json:"result"`
}

func (this *SearchMultiController) Post() {
	var msgResult MsgResult
	var searchMultiResult SearchMultiResult
	var userlist []string
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &userlist)
	if err != nil {
		msgResult.Msg = err.Error()
		this.Data["json"] = msgResult
		this.ServeJSON()
		return
	}

	result, err := models.Multi_Search_User(g.Config().Ldap, userlist)
	if err != nil {
		msgResult.Msg = err.Error()
		this.Data["json"] = msgResult
		this.ServeJSON()
		return
	}
	searchMultiResult.Success = true
	searchMultiResult.Result = result
	this.Data["json"] = searchMultiResult
	this.ServeJSON()
}
