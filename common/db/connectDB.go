package db

import (
	"fmt"
	"github.com/Carrotxyy/syncSense/common/setting"
	"log"
	"os"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DB struct {
	Conn *gorm.DB
}

// 链接数据库
func (d *DB) Connect(config setting.Config) error {
	// 获取配置
	dbType := config.DbType
	dbName := config.DbName
	user := config.DbUser
	pwd := config.DbPassword
	host := config.DbIP

	db, err := gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, pwd, host, dbName))
	if err != nil {
		log.Fatal("connecting DB error: ", err)
		os.Exit(1)
	}
	// 数据库表扩展名
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return config.TablePrefix + defaultTableName
	}
	db.LogMode(true) //打印SQL语句
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	d.Conn = db
	return nil
}
