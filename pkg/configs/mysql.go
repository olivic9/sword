package configs

type MysqlConfig struct {
	Url      string
	Port     string
	Database string
	UserName string
	Password string
}

func initializeMysqlConfig() {
	if MysqlCfg == nil {
		MysqlCfg = &MysqlConfig{
			Url:      getEnv("MYSQL_URL", "localhost"),
			Port:     getEnv("MYSQL_PORT", "3306"),
			Database: getEnv("MYSQL_DATABASE", ""),
			UserName: getEnv("MYSQL_USER", ""),
			Password: getEnv("MYSQL_PASSWORD", ""),
		}
	}
}
