package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Repo *gorm.DB
)

func InitMySQLDB() {
	dsn := "root:Wzzst310@163.com@tcp(wjjzst.com:3306)/dict_jp?charset=utf8&parseTime=True&loc=Local"
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         255,   // string 类型字段的默认长度
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	var err error
	if Repo, err = gorm.Open(mysql.New(mysqlConfig)); err != nil {
		panic("failed to connect database")
	} else {
		repoConf, _ := Repo.DB()
		repoConf.SetMaxIdleConns(2)
		repoConf.SetMaxOpenConns(10)
	}
}
