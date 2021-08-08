package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"path/filepath"
	"time"
)

var mySqlConnPool *gorm.DB

func init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", "root",
		"123456", "172.17.0.2", "3306", "hrs")
	//gormLogger := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
	//	SlowThreshold:             200 * time.Millisecond,
	//	LogLevel:                  logger.Warn,
	//	Colorful:                  true,
	//	IgnoreRecordNotFoundError: true,
	//})
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to open mysql,", err.Error())
	}
	sqlDB, _ := db.DB()
	sqlDB.SetConnMaxLifetime(time.Second * time.Duration(28000))
	mySqlConnPool = db
	log.Println("init mysql connection successful")
}

type gormTestTable struct {
	Id   int64
	Name string
}

func (gormTestTable) TableName() string {
	return "t_gorm_test"
}

func queryTask() {
	ti := time.NewTicker(time.Millisecond * 500)
	for {
		<-ti.C
		info := gormTestTable{}
		mySqlConnPool.Model(&info).Where("id = 2").First(&info)
	}
}

func main() {
	os.Create(filepath.Join(filepath.Dir(os.Args[0]), "start_successful.txt"))
	queryTask()
}
