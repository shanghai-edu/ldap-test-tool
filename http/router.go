package http

import (
	"github.com/astaxie/beego"
	"github.com/shanghai-edu/ldap-test-tool/http/controllers"
)

func ConfigRouters() {
	beego.Router("/api/v1/ldap/", &controllers.MainController{})
	beego.Router("/api/v1/ldap/health", &controllers.HealthController{})
	beego.Router("/api/v1/ldap/search/filter/:filter", &controllers.SearchFilterController{})
	beego.Router("/api/v1/ldap/search/user/:username", &controllers.SearchUserController{})
	beego.Router("/api/v1/ldap/search/multi", &controllers.SearchMultiController{})
	beego.Router("/api/v1/ldap/auth/single", &controllers.AuthSingleController{})
	beego.Router("/api/v1/ldap/auth/multi", &controllers.AuthMultiController{})
}
