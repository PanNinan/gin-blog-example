package models

import (
	"example.com/example/pkg/setting"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

var db *gorm.DB

type Model struct {
	ID         int        `gorm:"primary_key" json:"id"`
	CreatedAt  int64      `json:"created_at"`
	ModifiedAt *time.Time `gorm:"type:timestamp" json:"modified_at"`
}

type Database struct {
	Type        string
	dbName      string
	User        string
	Password    string
	Host        string
	Port        string
	TablePrefix string
}

var DBSetting = &Database{}

func Setup() {
	err := setting.Cfg.Section("database").MapTo(DBSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo DatabaseSetting err: %v", err)
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DBSetting.User, DBSetting.Password, DBSetting.Host, DBSetting.Port, DBSetting.dbName)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   DBSetting.TablePrefix, // 表前缀
			SingularTable: true,                  // 使用单数表名
		},
		Logger: logger.Default.LogMode(logger.Info),
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
	})
	if err != nil {
		log.Println(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get generic database object: %v", err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(10 * time.Minute)
}

func closeDB() {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get generic database object: %v", err)
	}
	sqlDB.SetMaxIdleConns(0)
	if err = sqlDB.Close(); err != nil {
		log.Fatalf("failed to close database connection: %v", err)
	}
	*db = gorm.DB{}
}
