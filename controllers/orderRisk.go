package controllers

import (
	_ "fast-go/models"
	beego "github.com/beego/beego/v2/server/web"
)

type OrderRiskController struct {
	beego.Controller
}

func (c *OrderRiskController) List() {

}
