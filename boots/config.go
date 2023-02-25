package boots

import (
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

const EXE_ROOT_NAME = "VM-IM-Web"

var (
	config *Config
)

func getConfigPath() (string, error) {
	exec_path, err := os.Getwd()
	if err != nil {
		return "", err
	}
	root_index := strings.Index(exec_path, EXE_ROOT_NAME)
	conf_path := exec_path[:root_index+len(EXE_ROOT_NAME)] + "/conf/config.yaml"
	return conf_path, nil
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
