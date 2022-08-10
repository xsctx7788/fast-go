package controllers

import (
	"encoding/json"
	"fast-go/models"
	_ "fast-go/models"
	beego "github.com/beego/beego/v2/server/web"
	"strconv"
)

type OrderRiskController struct {
	beego.Controller
}

type ListInfo struct {
	PageNow      int    `json:"pageNow"`  // 注意首字母大写
	PageSize     int    `json:"pageSize"` // 并且写明与json字段的映射关系，否则Unmarshal函数不好用
	ShowOrderId  string `json:"showOrderId"`
	Payment      string `json:"payment"`
	StartPayTime string `json:"startPayTime"`
	EndPayTime   string `json:"endPayTime"`
	RiskRule     string `json:"riskRule"`
	Status       int    `json:"status"`
	Country      string `json:"country"`
	IsBlackList  int    `json:"isBlackList"`
	LastStatus   int    `json:"lastStatus"`
	UserId       int    `json:"userId"`
	UserName     string `json:"userName"`
	Email        string `json:"email"`
	Editor       string `json:"editor"`
}

func Strval(value interface{}) string {
	var key string
	if value == nil {
		return key
	}

	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}

	return key
}

//主模型初始化
var info []map[string]string

//关联模型初始化
var info1 []map[string]string

func (c *OrderRiskController) List() {
	// 获取post数据
	var Params ListInfo
	body := c.Ctx.Input.RequestBody //这是request的body 的json二进制数据
	json.Unmarshal(body, &Params)   //解析二进制json，把结果放进MovieInfo_obj中

	// 调用model方法
	unix := 1660121761
	str := models.UnixToDate(unix)

	Params.EndPayTime = str
	c.Data["json"] = map[string]interface{}{"result": Params, "msg": "获取成功", "code": "200"} // 设置返回值
	c.ServeJSON()
}
