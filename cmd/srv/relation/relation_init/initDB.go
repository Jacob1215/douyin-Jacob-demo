package relation_init

import (
	global2 "douyin-Jacob/cmd/srv/relation/global"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

func InitDB() {
	//dsn := "root:root@tcp(192.168.1.104:3306)/mxshop_order_srv?charset=utf8mb4&parseTime=True&loc=Local" //虚拟机的地址
	c := global2.ServerConfig.MysqlInfo
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.User, c.Password, c.Host, c.Port, c.Name)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), //io wirter
		logger.Config{
			SlowThreshold: time.Second, //慢SQL阈值
			LogLevel:      logger.Info, //log level
			Colorful:      true,        //禁用彩色打印
		},
	)
	//全局模式
	var err error
	global2.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //使生成表的时候使user,不是users。
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
}
