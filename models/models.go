package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	//_ "github.com/go-sql-driver/mysql"
)

func RegisterDB() {
	// register model
	orm.RegisterModel(new(User))

	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", "postgres://lwezcldi:RqhjUwdBM_xvT3JivKir2ZVZZue90WnZ@stampy.db.elephantsql.com:5432/lwezcldi")

}
