package boots

type Config struct {
	Mysql MysqlNode `yaml:"mysql"`
}

type MysqlNode struct {
	Host         string `yaml:"host"`
	Port         int    `yaml:"port"`
	Database     string `yaml:"database"`
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
	Charset      string `yaml:"charset"`
	Debug        bool   `yaml:"debug"`
	LogLevel     string `yaml:"logLevel"`
	MaxOpenConns int    `yaml:"maxOpenConns"`
	MaxIdleConns int    `yaml:"maxIdleConns"`
	Timeout      string `yaml:"timeout"`
	ParseTime    bool   `yaml:"parseTime"`
}
