package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["myapi/controllers:LoginController"] = append(beego.GlobalControllerRouter["myapi/controllers:LoginController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:username/:passwd`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["myapi/controllers:LoginController"] = append(beego.GlobalControllerRouter["myapi/controllers:LoginController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

}
