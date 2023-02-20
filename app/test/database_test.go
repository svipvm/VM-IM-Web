package test

import (
	"testing"

	"github.com/svipvm/VM-IM-Web/app/utils"

	"github.com/go-xorm/xorm"
)

func getConnection() *xorm.Engine {
	return utils.GetEngine()
}

func TestConnection(t *testing.T) {
	engine := getConnection()

	if engine != nil {
		engine.Ping()
	}

	defer engine.Close()
}
