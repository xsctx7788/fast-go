package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type OrderRisk struct {
	Id            int
	OrderId       string
	Show_order_id string
	Status        int
}

func init() {

}

func main() {
	var risk OrderRisk
	dsn := "root:123456@tcp(localhost)/skr-order?charset-utf8mb4&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	},
	})

	if err != nil {
		fmt.Printf("failed to connect database")
	}
	db.Debug().Model(&risk).Where(OrderRisk{Id: 54}).Update("OrderId", "G0000000001") // 仅更新非零值字段
}
