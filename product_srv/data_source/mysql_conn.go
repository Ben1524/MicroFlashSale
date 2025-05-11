package data_source

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
	"user_srv/models"
)

var Db *gorm.DB
var err error

func init() {
	InitConf()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		MysqlInfoSource.User,
		MysqlInfoSource.Pwd,
		MysqlInfoSource.Host,
		MysqlInfoSource.Port,
		MysqlInfoSource.Database,
	)
	log.Println(dsn)
	// Db,err = gorm.Open("mysql",data_source)
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	log.Println("mysql conn success")

	// 获取底层数据库连接对象
	sqlDB, err := Db.DB()
	if err != nil {
		log.Fatalf("failed to get underlying database connection: %v", err)
	}

	// 配置最大打开连接数
	sqlDB.SetMaxOpenConns(100)
	// 配置最大空闲连接数
	sqlDB.SetMaxIdleConns(20)
	// 配置连接的最大生命周期
	sqlDB.SetConnMaxLifetime(time.Minute * 30)

	log.Println("Connected to the database and configured connection limits!")
	if err := Db.AutoMigrate(&models.FrontUser{}, &models.AdminUser{}); err != nil {
		log.Fatal(err)
	}
}
