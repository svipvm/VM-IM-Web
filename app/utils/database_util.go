package utils

import (
	"fmt"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/svipvm/VM-IM-Web/boots"
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

	config, err := boots.ReadConfig()
	if err != nil {

	}
	username := config.Mysql.Username
	password := config.Mysql.Password
	con_addr := config.Mysql.Host
	con_port := config.Mysql.Port
	database := config.Mysql.Database

	con_url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
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
