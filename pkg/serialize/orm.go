package serialize

import (
	"github.com/jinzhu/gorm"
	"github.com/qkgo/scaff/pkg/cfg"
	"github.com/qkgo/scaff/pkg/util/system"
	"time"
)

func ConfigDatabase(dialect string, url string) {
	if dialect == "" || url == "" {
		println("connect database param is null.")
		cfg.Log.Fatalf("connect database param is null.")
		system.Exit(-100)
		return
	}
	db, err := gorm.Open(dialect, url)
	if err != nil {
		println("Got error when connect database, the error is ", err.Error())
		cfg.Log.Fatalf("Got error when connect database, the error is '%v'", err)
		time.Sleep(time.Second)
		system.Exit(-200)
		return
	}
	db.LogMode(true)
	db.SetLogger(cfg.SqlLog)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	DB = db
	return
}

func ConfigSecondDatabase(dialect string, url string) {
	if dialect == "" || url == "" {
		println("connect database param is null.")
		cfg.Log.Fatalf("connect database param is null.")
		system.Exit(-100)
		return
	}
	db, err := gorm.Open(dialect, url)
	if err != nil {
		println("Got error when connect database, the error is ", err.Error())
		cfg.Log.Fatalf("Got error when connect database, the error is %v", err)
		time.Sleep(time.Second)
		system.Exit(-200)
		return
	}
	db.LogMode(true)
	db.SetLogger(cfg.SqlLog)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	SecondDB = db
	return
}
