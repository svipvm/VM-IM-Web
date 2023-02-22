package test

import (
	"os"
	"reflect"
	"testing"

	"github.com/svipvm/VM-IM-Web/boots"
	"gopkg.in/yaml.v3"
)

func TestReadConfig(t *testing.T) {
	// yamlFile, err := os.ReadFile("../../conf/config.yaml")
	// if err != nil {
	// 	t.Log(err)
	// }
	// // t.Log(yamlFile)
	// config := new(boots.Config)
	// err = yaml.Unmarshal(yamlFile, &config)

	// _, filename, _, _ := runtime.Caller(0)
	// t.Log(os.Getwd())
	// t.Log(os.Args)
	config, err := boots.ReadConfig()
	if err != nil {
		t.Log(err)
	}
	if config != nil {
		t.Log(config)
		t.Logf("%s, %d", config.Mysql.Host, config.Mysql.Port)
	}
}

func TestReadOneNodeConfig(t *testing.T) {
	yamlFile, err := os.ReadFile("../../conf/config.yaml")
	if err != nil {
		t.Log(err)
	}
	// t.Log(yamlFile)
	config := new(boots.Config)
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		t.Log(err)
	}
	ref_type := reflect.TypeOf(*config)
	// t.Log(ref_type.NumField())
	for i := 0; i < ref_type.NumField(); i++ {
		field := ref_type.Field(i)
		t.Log(field)
		t.Log(field.Name, field.Index, field.Type, field.Tag)
		t.Log(field.Name == "mysql")
		t.Log(field.Name == "Mysql")
	}

}
