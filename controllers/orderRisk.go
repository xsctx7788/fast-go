package controllers

import (
	"encoding/json"
	_ "fast-go/models"
	beego "github.com/beego/beego/v2/server/web"
)

type OrderRiskController struct {
	beego.Controller
}

type MovieInfo struct {
	Id   string `json:"id"`   // 注意首字母大写
	Name string `json:"name"` // 并且写明与json字段的映射关系，否则Unmarshal函数不好用
}

func (c *OrderRiskController) List() {

	var MovieInfo_obj MovieInfo
	body := c.Ctx.Input.RequestBody      //这是request的body 的json二进制数据
	json.Unmarshal(body, &MovieInfo_obj) //解析二进制json，把结果放进MovieInfo_obj中

	c.Data["json"] = map[string]interface{}{"result": MovieInfo_obj, "msg": "获取成功", "code": "200"} // 设置返回值
	c.ServeJSON()
}
