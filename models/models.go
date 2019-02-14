package models

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"

	"github.com/jinzhu/gorm"
	// 仅初始化
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

// Model 基础模型
type Model struct {
	ID int `gorm:"primary_key" json:"id"`
	// CreatedAt  int `json:"created_at"`
	CreatedAt time.Time `json:"created_at" time_format:"20060102150405"`
	// ModifiedAt int `json:"modified_at"`
}

func init() {
	var (
		err                                               error
		dbType, dbName, user, password, host, tablePrefix string
	)

	dbType = viper.GetString("database.TYPE")
	dbName = viper.GetString("database.NAME")
	user = viper.GetString("database.USER")
	password = viper.GetString("database.PASSWORD")
	host = viper.GetString("database.HOST")
	tablePrefix = viper.GetString("database.TABLE_PREFIX")

	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))

	if err != nil {
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}
	// 数据库表名 单数
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	defer db.Close()
}
