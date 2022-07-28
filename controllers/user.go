package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
)

type UserController struct {
	beego.Controller
}

type userInfo struct {
	Id     int
	Email  string
	Mobile string
}

// GET 请求
// 输出JSON
func (c *UserController) Get() {
	fmt.Println("exec user controller get")
	uid, _ := c.GetInt("id")
	userInfo := userInfo{
		Id:     uid,
		Email:  "123@123.com",
		Mobile: "123456789",
	}
	c.Data["json"] = &userInfo
	c.ServeJSON()
}

// POST 请求
// 输出XML
func (c *UserController) Post() {
	userInfo := new(userInfo)
	c.ParseForm(userInfo)
	c.Data["xml"] = &userInfo
	c.ServeXML()
}
