package model

import (
	"fmt"
	"log"
	"time"

	"github.com/DQFSN/blog/config"
	"github.com/DQFSN/blog/server/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // mysql 驱动
)

var db *gorm.DB

// init 数据库初始化
func initDB() {

	// 获取配置
	conf := config.Get().Mysql

	fmt.Println(conf)

	// 初始化连接
	var err error
	db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", conf.User, conf.Password, conf.Host, conf.Port, conf.DB))
	if err != nil {
		log.Fatal("数据库连接错误：", err, conf)
	}
	db.DB().SetMaxOpenConns(20)
	db.DB().SetConnMaxLifetime(10 * time.Second)

	autoMigrate()
}

// 自动迁移
func autoMigrate() {
	db.AutoMigrate(
		&model.User{},
		&model.Blog{},
	)
}

// DB 返回db，如果不存在则初始化
func DB() *gorm.DB {
	if db == nil {
		initDB()
	}
	return db
}

// Close 关闭数据库连接
func Close() error {
	return DB().Close()
}


