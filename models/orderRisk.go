package models

import (
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
