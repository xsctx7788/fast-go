package controllers

import (
	"encoding/json"
	_ "fast-go/models"
	"fmt"
	"github.com/astaxie/beego/orm"
	beego "github.com/beego/beego/v2/server/web"
	"strconv"
)

type WarehouseOutController struct {
	beego.Controller
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

func (l *WarehouseOutController) Api() {
	//step2：配置基础的东西
	o := orm.NewOrm()
	qs := o.QueryTable("Out")

	var maps []orm.Params
	//step2：expr逻辑

	//2.1筛选出包含xxx的语句  contains大小写敏感
	//iexact大小写不敏感
	//qs.Filter("bookname__contains", "围城").Values(&maps)

	//2.2查询指定的值
	//qs.Filter("bookname", "老人与海").Values(&maps)

	//qs.Filter("Id__gt", 18).Values(&maps) // WHERE profile.age > 18
	//qs.Filter("Id__gte", 18).Values(&maps) // WHERE profile.age >= 18
	//qs.Filter("Id__in", 0, 20).Values(&maps) // WHERE profile.age IN (18, 20) in好像用不了
	//qs.Filter("name__endswith", "slene") // WHERE name LIKE BINARY '%slene'
	//qs.Filter("name__startswith", "slene") // WHERE name LIKE BINARY 'slene%'
	qs.Filter("Id__gt", 15).Exclude("Id__gt", 18).Values(&maps) //// WHERE profile.age IN (15, 20) AND NOT profile_id < 18
	//step3：第一次拼接
	for _, m := range maps {
		//int的拼装用strconv.Atoi
		id, _ := strconv.Atoi(Strval(m["Id"]))
		//string的拼装用Strval
		name := Strval(m["Barcode"])
		//subMapA := map[string]string{"A_Key_1": "A_SubValue_1", "A_key_2": "A_SubValue_2"}
		//info["MapA"] = subMapA
		info = append(
			info,
			map[string]string{"id": Strval(id), "name": name},
		)

	}

	//step4：关联模型测试 假设传来的值是这个
	otest := orm.NewOrm()
	var mtest []orm.Params
	otest.Raw("SELECT * FROM books where id=2").Values(&mtest)

	res, err := otest.Raw("delete FROM books where id=125").Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		fmt.Println("mysql row affected nums: ", num)
	}
	fmt.Println()

	for _, m := range mtest {
		//int的拼装用strconv.Atoi
		id, _ := strconv.Atoi(Strval(m["id"]))
		//string的拼装用Strval
		name := Strval(m["author"])

		info1 = append(
			info1,
			map[string]string{"id": Strval(id), "name": name},
		)

	}
	//step5：最终输出
	l.Data["json"] = map[string]interface{}{"code": "200", "msg": "接收成功", "result": info, "假装是关联模型": info1}
	//这里必须要加上不然会越append越多
	info = []map[string]string{}
	info1 = []map[string]string{}
	l.ServeJSON()

}
