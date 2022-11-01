package model

import (
	_ "database/sql"
	"time"

	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 数据库的一些初始化工作

// 数据库连接实例
var DB *gorm.DB

func Initial(dsn string) error {
	var (
		db *gorm.DB
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return errors.Wrap(err, "open mysql error. ")
	}

	// 设置连接池
	sqlDb, err := db.DB()
	if err != nil {
		return errors.Wrap(err, "get sql db error. ")
	}
	//空闲
	sqlDb.SetMaxIdleConns(50)
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetConnMaxLifetime(time.Hour)

	DB = db
	// 执行数据库迁移
	if err := Migrate(); err != nil {
		return err
	}
	return nil
}

func Migrate() error {
	var err error
	// 1.确认是否需要执行迁移
	if err = DB.AutoMigrate(&WorkUser{}); err != nil {
		return err
	}

	return nil
}
