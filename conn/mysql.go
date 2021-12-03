package conn

import (
	"github.com/maybgit/glog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDB() *gorm.DB {
	dsn := "root:123456@(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local" //默认就一个数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		glog.Error(err)
	}
	db = db.Debug()
	return db
}
