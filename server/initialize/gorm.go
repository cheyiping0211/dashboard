package initialize

import (
    "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"os"
	"fmt"
	"ping/model"
)

// SqliteTables 注册数据库表专用
func SqliteTables(db *gorm.DB) {
	 db.AutoMigrate(
		&model.Modules{},
		&model.Chart{},
	)
	fmt.Println("数据库自动迁移成功")
}

// GormSqlite 初始化Sqlite数据库
func GormSqlite() *gorm.DB {
	pwd,_ := os.Getwd()
	if db, err := gorm.Open("sqlite3", pwd + "/initialize/ping.db"); err != nil {
        fmt.Println("数据库连接失败")
		os.Exit(0)
		return nil
	} else {
		return db
	}
}