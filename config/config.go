package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Environment struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPass     string
	DBName     string
	AppPort    string
	SecretKey  string
	JwtExpired int
}

var Env *Environment

func getEnv(key string, required bool) string {
	value, ok := os.LookupEnv(key)
	if !ok && required {
		log.Fatalf("Missing or invalid environment key: '%s'", key)
	}
	return value
}

func LoadEnvironment() {
	if Env == nil {
		Env = new(Environment)
	}

	JwtExpired, err := strconv.Atoi(getEnv("JWT_EXPIRED", true))
	if err != nil {
		log.Fatalf("jwt expired parsing error: '%s'", err.Error())
	}
	
	Env.DBHost = getEnv("DB_HOST", true)
	Env.DBPort = getEnv("DB_PORT", true)
	Env.DBUser = getEnv("DB_USER", true)
	Env.DBPass = getEnv("DB_PASSWORD", true)
	Env.DBName = getEnv("DB_NAME", true)

	Env.AppPort = getEnv("APP_PORT", true)
	Env.SecretKey = getEnv("SECRET_KEY", true)
	Env.JwtExpired = JwtExpired
}

func LoadEnvironmentFile(file string) {
	if err := godotenv.Load(file); err != nil {
		log.Printf("Warning: Error on load environment file: %s. Using environment variables.", file)
	}
	LoadEnvironment()
}
