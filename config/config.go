package config

import "github.com/eliasyoung/fiber-flavor/internal/env"

type Config struct {
	Addr     string
	DBConfig DBConfig
	Env      string
	ApiURL   string
}

type DBConfig struct {
	DBHost     string
	DBName     string
	DBUser     string
	DBPort     int
	DBPassword string
}

func InitConfig() Config {
	cfg := Config{
		Addr: env.GetDotEnvConfigWithFallback("ADDR", "8080"),
		DBConfig: DBConfig{
			DBHost:     env.GetDotEnvConfigWithFallback("DB_HOST", "127.0.0.1"),
			DBName:     env.GetDotEnvConfigWithFallback("DB_NAME", "whatever"),
			DBUser:     env.GetDotEnvConfigWithFallback("DB_USER", "admin"),
			DBPassword: env.GetDotEnvConfigWithFallback("DB_PASSWORD", "adminpassword"),
			DBPort:     env.GetIntDotEnvConfigWithFallback("DB_PORT", 8848),
		},
		Env:    env.GetDotEnvConfigWithFallback("ENV", "development"),
		ApiURL: env.GetDotEnvConfigWithFallback("EXTERNAL_URL", "localhost:8080"),
	}

	return cfg
}
