package models

import "github.com/beego/beego/v2/client/orm"

type Out struct {
	Id      int
	OrderId string
	Barcode string
}

func init() {
	orm.RegisterModel(new(Out))
}
