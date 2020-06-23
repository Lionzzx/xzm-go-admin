package main

import (
	"ginEssential/core"
	"ginEssential/global"
	"ginEssential/initialize"
)


func main() {
	initialize.Mysql()
	initialize.Redis()
	initialize.DBTables()
	defer global.GVA_DB.Close()
	core.RunWindowsServer()
}
