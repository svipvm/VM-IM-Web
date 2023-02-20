package utils

import (
	"fmt"
	"sync"

	"github.com/go-xorm/xorm"
)

var (
	engine *xorm.Engine
	lock   sync.Mutex
)

func GetEngine() *xorm.Engine {
	lock.Lock()
	defer lock.Unlock()

	if engine != nil {
		return engine
	}

	username := "root"
	password := "shadow24"
	con_addr := "127.0.0.1"
	con_port := "3306"
	database := "test"

	con_url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		username, password, con_addr, con_port, database)
	engine_, err := xorm.NewEngine("mysql", con_url)
	if err != nil {
		fmt.Printf("%s", err)
		return nil
	}
	engine_.ShowSQL(true)

	engine = engine_
	return engine
}
