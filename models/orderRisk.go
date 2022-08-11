package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

type OrderRisk struct {
	ID               int
	OrderId          string
	ShowOrderId      string
	RiskRule         string
	RiskRuleId       string
	LastRiskId       int
	LastRiskStatus   int
	Status           int
	DetailCount      int
	DetailPay        string
	on_order_count   int
	on_order_pay     string
	hav_order_count  int
	hav_order_pay    string
	back_order_count int
	back_order_pay   string
	editor           string
	is_blacklist     int
	CreateTime       time.Time
	UpdateTime       time.Time
	PayTime          time.Time
}

type ListInfo struct {
	PageNow      int    `json:"pageNow" default:"1"`   // 注意首字母大写
	PageSize     int    `json:"pageSize" default:"20"` // 并且写明与json字段的映射关系，否则Unmarshal函数不好用
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

//DSN
const DSN = "root:123456@tcp(localhost)/skr-order?charset-utf8mb4&parseTime=True&loc=Local"

func main() {
	//
	//var risk OrderRisk
	//db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{NamingStrategy: schema.NamingStrategy{
	//	SingularTable: true, // 使用单数表名
	//},
	//	SkipDefaultTransaction: true,  // 跳过默认事务
	//	DryRun:                 false, // 生成 SQL 但不执行，可以用于准备或测试生成的 SQL
	//})
	//
	//if err != nil {
	//	fmt.Printf("failed to connect database")
	//}
	//db.Debug().Model(&risk).Where(OrderRisk{ID: 54}).Update("OrderId", "G0000000001") // 仅更新非零值字段

	// 根据 User 的字段创建 `deleted_users` 表
	// db.Table("deleted_users").AutoMigrate(&User{})

	// 从另一张表查询数据
	//var risks []User
	//db.Debug().Model(&risk).Where(OrderRisk{ID: 54}).Find(&risks)
	//db.Table("deleted_users").Find(&deletedUsers)

	// db.Table("deleted_users").Where("name = ?", "jinzhu").Delete(&User{})
	// DELETE FROM deleted_users WHERE name = 'jinzhu';
}

type RiskListInfo struct {
	id            int
	show_order_id string
}

func FindRiskList(data ListInfo) []RiskListInfo {
	// 组装查询条件
	//var risk OrderRisk
	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true, // 使用单数表名
	},
		SkipDefaultTransaction: true,  // 跳过默认事务
		DryRun:                 false, // 生成 SQL 但不执行，可以用于准备或测试生成的 SQL
	})

	if err != nil {
		fmt.Printf("failed to connect database")
	}
	var risks []RiskListInfo

	//field := "a.id,a.show_order_id as showOrderId,a.order_id as orderId,a.last_risk_id as lastRiskId,a.last_risk_status as lastStatus,a.status,a.is_blacklist as isBlackList,a.detail_count as detailCount,a.detail_pay as detailPay,a.on_order_count as onOrderCount,a.on_order_pay as onOrderPay,a.hav_order_count as havOrderCount,a.hav_order_pay as havOrderPay,a.back_order_pay as backOrderPay,a.back_order_count as backOrderCount,a.editor,a.create_time as createTime,a.update_time as updateTime,a.risk_rule as riskRule,a.pay_time as payTime,b.payment,b.amount,b.user_id as userId,c.first_name as firstName,c.last_name as lastName,c.mobile,c.country_name as country,CONCAT_WS(\",\",c.address_1,c.address_2,c.city,c.state_name,c.zip_code,c.country_name) as address"

	field := "id,show_order_id"
	//$db_b->table('order_risk')->alias('a')
	//->join(['order' =>'b'],'a.show_order_id = b.show_order_id','LEFT')
	//->join(['order_address_ext' =>'c'],'a.show_order_id = c.show_order_id','LEFT')
	//->field($field)
	//->where($where)
	//->where($where2)
	//->order('a.create_time desc')
	//->limit(($pageNow - 1) * $pageSize, $pageSize)
	//->select();
	db.Debug().Table("order_risk").Select(field).Where(OrderRisk{ID: 54}).Find(&risks)

	//a := 12312313
	//t := time.Unix(int64(a), 0)

	//return t.Format("2006-01-02 15:04:05")
	return risks
}
