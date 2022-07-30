package routers

import (
	"fast-go/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {

	// 固定路由
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user", &controllers.UserController{})

	// 正则路由
	//beego.Router("/user/?:id", &controllers.UserController{})

	// 自定义路由
	beego.Router("/order_risk/list", &controllers.OrderRiskController{}, "post:List")
}
