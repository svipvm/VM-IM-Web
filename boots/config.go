package boots

import (
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

// const CONFIG_FILE string = "../conf/config.yaml"
const EXE_ROOT_NAME = "VM-IM-Web"

// exe_path, _ := os.Getwd()
// const CONFIG_FILE string = exe_path

var (
	config *Config
)

// func GetConfig(nodeName string, nodeData interface{}) error {
// 	err := readConfig()
// 	if err != nil {
// 		return err
// 	}
// 	ref_type := reflect.TypeOf(*config)
// 	for i := 0; i < ref_type.NumField(); i++ {
// 		field := ref_type.Field(i)
// 		if strings.ToLower(field.Name) == strings.ToLower(nodeName) {

// 		}
// 	}
// 	// nodeData = yaml
// }

func getConfigPath() (string, error) {
	exec_path, err := os.Getwd()
	if err != nil {
		return "", err
	}
	root_index := strings.Index(exec_path, EXE_ROOT_NAME)
	conf_path := exec_path[:root_index+len(EXE_ROOT_NAME)] + "/conf/config.yaml"
	return conf_path, nil
	// if len(os.Args) < 2 {
	// 	return "", errors.New("application root error")
	// }
	// if os.Args[1] == "test" {
	// 	return "../../conf/config.yaml", nil
	// } else {
	// 	return "../conf/config.yaml", nil
	// }
}

func ReadConfig() (*Config, error) {
	if config != nil {
		return config, nil
	}
	config_ := new(Config)
	// fmt.Printf(os.Getwd())
	exe_path, err := getConfigPath()
	if err != nil {
		return nil, err
	}
	yamlFile, err := os.ReadFile(exe_path)
	// yamlFile, err := os.ReadFile(CONFIG_FILE)
	if err != nil {
		return nil, err
	}
	// t.Log(yamlFile)
	err = yaml.Unmarshal(yamlFile, &config_)
	if err != nil {
		return nil, err
	}
	config = config_
	return config, nil
}
