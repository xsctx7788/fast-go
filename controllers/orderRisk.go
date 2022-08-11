package controllers

import (
	"encoding/json"
	"fast-go/models"
	_ "fast-go/models"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/mcuadros/go-defaults"
)

type OrderRiskController struct {
	beego.Controller
}

func (c *OrderRiskController) List() {
	// 获取post数据
	var Params models.ListInfo
	defaults.SetDefaults(&Params)
	body := c.Ctx.Input.RequestBody //这是request的body 的json二进制数据
	json.Unmarshal(body, &Params)   //解析二进制json，把结果放进MovieInfo_obj中

	// 调用model方法
	//unix := 1660121761
	//str := models.UnixToDate(unix)

	str := models.FindRiskList(Params)
	c.Data["json"] = map[string]interface{}{"result": str, "msg": "获取成功", "code": "200"} // 设置返回值
	c.ServeJSON()
}
