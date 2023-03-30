package models

import (
	"github.com/chenyuIT/framework/database/orm"
)

type User struct {
	orm.Model
	Name   string
	Avatar string
	orm.SoftDeletes
}
