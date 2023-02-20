package test

import (
	"testing"

	"github.com/svipvm/VM-IM-Web/app/utils"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type Student struct {
	Id   int32
	Name string
	Age  int
}

func getConnection() *xorm.Engine {
	// username := "root"
	// password := "shadow24"
	// con_addr := "127.0.0.1"
	// con_port := "3306"
	// database := "test"

	// con_url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
	// 	username, password, con_addr, con_port, database)
	// engine, err := xorm.NewEngine("mysql", con_url)
	// if err != nil {
	// 	panic(err)
	// }
	// engine.Ping()

	// return engine

	return utils.GetEngine()
}

func TestConnection(t *testing.T) {
	engine := getConnection()

	if engine != nil {
		engine.Ping()
	}
	// engine.ShowSQL(true)
	// engine, err = xorm.NewEngine("mysql", "root:3306@/test?charset=utf8")
	// t.Log(engine.DBMetas())

	defer engine.Close()
}

func TestQuertOne(t *testing.T) {
	engine := getConnection()
	student := new(Student)
	has, err := engine.Where("name=?", "WHQ").Get(student)
	if err != nil {
		panic(err)
	}
	if has {
		t.Log(student)
	}
	defer engine.Close()
}

func TestInsertOne(t *testing.T) {
	engine := getConnection()
	student := new(Student)
	// student.Id = 3
	student.Name = "VYW"
	student.Age = 24

	affected, err := engine.Insert(student)
	if err != nil {
		panic(err)
	}
	if affected != 0 {
		t.Log("Insert success")
	}
	defer engine.Close()
}

// func TestUpdateOne(t *testing.T) {
// 	engine := getConnection()
// 	student := new(Student)
// 	gain, err := engine.Id(1).Get(student)
// 	if err != nil {
// 		panic(err)
// 	}
// 	if gain {
// 		t.Log(student)
// 		student.Age = 99
// 		t.Log(student)
// 		affected, err := engine.Id(1).Update(student)
// 		if err != nil {
// 			panic(err)
// 		}
// 		if affected != 0 {
// 			t.Log("Update success")
// 		}
// 	}
// 	defer engine.Close()
// }
