package serialize

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/qkgo/scaff/pkg/cfg"
	"github.com/qkgo/scaff/pkg/log"
	"github.com/qkgo/scaff/pkg/util/system"
	"time"
)

func ConfigXDatabase(dialect string, url string) {
	if dialect == "" || url == "" {
		//fmt.Println("connect database param is null.")
		cfg.Log.Fatalf("connect database param is null.")
		system.Exit(-100)
		return
	}

	engine, err := xorm.NewEngine(dialect, url)
	if err != nil {
		//println("Got error when connect database, the error is '%v'", err)
		cfg.Log.Fatalf("Got error when connect database, the error is '%v'", err)
		time.Sleep(time.Second)
		system.Exit(-200)
		return
	}
	cfg.XOrmLog = &log.SqlLogger{Log: cfg.SqlLog}
	engine.SetLogger(cfg.XOrmLog)
	engine.ShowSQL()
	XDB = engine
	return
}

func ConfigSecondXDatabase(dialect string, url string) {
	if dialect == "" || url == "" {
		//fmt.Println("connect database param is null.")
		cfg.Log.Fatalf("connect database param is null.")
		system.Exit(-100)
		return
	}
	engine, err := xorm.NewEngine(dialect, url)
	if err != nil {
		//println("Got error when connect database, the error is '%v'", err)
		cfg.Log.Fatalf("Got error when connect database, the error is '%v'", err)
		time.Sleep(time.Second)
		system.Exit(-200)
		return
	}
	cfg.XOrmLog = &log.SqlLogger{Log: cfg.SqlLog}
	engine.SetLogger(cfg.XOrmLog)
	engine.ShowSQL()
	SecondXDB = engine
	return
}
