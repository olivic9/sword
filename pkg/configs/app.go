package configs

import (
	"log"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

const (
	Development = "development"
	Staging     = "staging"
	Production  = "production"
)

type ApplicationConfig struct {
	Env                string
	AppName            string
	HttpServerPort     string
	CorsTrustedOrigins []string
	JwtIssuer          string
	JwtSecret          string
}

func initializeApplicationConfigs() {
	if ApplicationCfg == nil {
		ApplicationCfg = &ApplicationConfig{
			Env:                getEnv("ENV", ""),
			AppName:            getEnv("APP_NAME", ""),
			HttpServerPort:     getEnv("HTTP_SERVER_PORT", "8080"),
			CorsTrustedOrigins: getEnvAsSlice("CORS_TRUSTED_ORIGINS", []string{}, ","),
			JwtIssuer:          getEnv("JWT_ISSUER", ""),
			JwtSecret:          getEnv("JWT_SECRET", ""),
		}
	}
}

var (
	ApplicationCfg *ApplicationConfig
	MysqlCfg       *MysqlConfig
	KafkaCfg       *KafkaConfig
)

func InitializeConfigs() {
	initializeApplicationConfigs()
	initializeMysqlConfig()
	initializeKafkaConfig()
}
