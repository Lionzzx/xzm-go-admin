package initialize

import (
	"ginEssential/global"
	"ginEssential/model"
)

// 注册数据库表专用
func DBTables() {
	db := global.GVA_DB
	db.AutoMigrate(model.SysUser{},
	)
}
