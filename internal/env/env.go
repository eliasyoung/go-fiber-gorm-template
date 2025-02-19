package env

import (
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
)

func GetString(key, fallback string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	return val
}

func GetInt(key string, fallback int) int {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	valAsInt, err := strconv.Atoi(val)
	if err != nil {
		return fallback
	}

	return valAsInt
}

func GetDotEnvConfig(key string) string {
	wd, err := os.Getwd()

	if err != nil {
		log.Fatal("Can't read work directory!")
	}

	envPath := filepath.Join(wd, ".env")

	err = godotenv.Load(envPath)
	if err != nil {
		log.Fatal("ENV file not set!")
	}

	val, ok := os.LookupEnv(key)
	if !ok || len(val) == 0 {
		log.Fatalf("%s in .env not found!", key)
	}

	return val

}

func GetDotEnvConfigWithFallback(key, fallback string) string {
	// wd, err := os.Getwd()

	// if err != nil {
	// 	return fallback
	// }

	// envPath := filepath.Join(wd, ".env")

	// err = godotenv.Load(envPath)
	// if err != nil {
	// 	return fallback
	// }

	// val, ok := os.LookupEnv(key)
	// if !ok || len(val) == 0 {
	// 	return fallback
	// }

	// return val

	val, ok := os.LookupEnv(key)
	if ok && len(val) > 0 {
		return val
	}

	wd, err := os.Getwd()
	if err != nil {
		return fallback
	}

	envPath := filepath.Join(wd, ".env")

	err = godotenv.Load(envPath)
	if err == nil {
		valDotEnv, okDotEnv := os.LookupEnv(key)
		if okDotEnv && len(valDotEnv) > 0 {
			return valDotEnv
		}
	}
	return fallback
}

func GetIntDotEnvConfigWithFallback(key string, fallback int) int {
	val, ok := os.LookupEnv(key)
	if ok && len(val) > 0 {
		valAsInt, error := strconv.Atoi(val)
		if error == nil {
			return valAsInt
		}
	}

	wd, err := os.Getwd()
	if err != nil {
		return fallback
	}

	envPath := filepath.Join(wd, ".env")

	err = godotenv.Load(envPath)
	if err != nil {
		return fallback
	}

	err = godotenv.Load(envPath)
	if err == nil {
		valDotEnv, okDotEnv := os.LookupEnv(key)
		if okDotEnv && len(valDotEnv) > 0 {
			valAsInt, error := strconv.Atoi(valDotEnv)
			if error == nil {
				return valAsInt
			}
		}
	}
	return fallback
}
