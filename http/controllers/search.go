package controllers

import (
	"github.com/astaxie/beego"
	"github.com/shanghai-edu/ldap-test-tool/g"
	"github.com/shanghai-edu/ldap-test-tool/models"
)

type SearchController struct {
	beego.Controller
}

type SearchResult struct {
	Results []models.LDAP_RESULT `json:"results"`
	Success bool                 `json:"success"`
}

func (this *SearchController) Get() {

	searchFilter := this.GetString("filter")
	results, err := models.Single_Search(g.Config().Ldap, searchFilter)
	if err != nil {
		var failedResult MsgResult
		failedResult.Msg = err.Error()
		this.Data["json"] = failedResult
	} else {
		var successResult SearchResult
		successResult.Success = true
		successResult.Results = results
		this.Data["json"] = successResult
	}
	this.ServeJSON()
}
