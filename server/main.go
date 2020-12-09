package main

import (
	"ping/core"
	"ping/global"
	"ping/initialize"
)

func main() {
	global.GVA_DB = initialize.GormSqlite()   
	initialize.SqliteTables(global.GVA_DB) 
	core.RunServer()
}
