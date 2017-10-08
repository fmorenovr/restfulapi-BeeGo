package routers

import (
	"myapi/controllers"
	"github.com/astaxie/beego"
)

func init() {
  // namespaces
  var namespaces []string = []string {"MyAPI","another"}
  // controllers
  var ctrllers []string = []string{"login","other"}

  restfulRouter := beego.NewNamespace("/"+namespaces[0],
	  beego.NSNamespace("/"+ctrllers[0],
		  beego.NSInclude(
			  &controllers.LoginController{},
		  ),
	  ),
  )
  beego.AddNamespace(restfulRouter)
}
