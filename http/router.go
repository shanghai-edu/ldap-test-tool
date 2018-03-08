package http

import (
	"github.com/astaxie/beego"
	"github.com/shanghai-edu/ldap-test-tool/http/controllers"
)

func ConfigRouters() {
	beego.Router("/api/v1/ldap/", &controllers.MainController{})
	beego.Router("/api/v1/ldap/health", &controllers.HealthController{})
	beego.Router("/api/v1/ldap/search", &controllers.SearchController{})
	beego.Router("/api/v1/ldap/search/:username", &controllers.SearchUserController{})
	beego.Router("/api/v1/ldap/auth/:username", &controllers.AuthController{})
}
