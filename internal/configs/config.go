package configs

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	RateLimit     int
	MongoURI      string
	MongoDB       string
	MongoUsername string
	MongoPassword string
	JWTSecret     string
	JWTExpiration int
}

var AppConfig *Config

func LoadConfigs() *Config {
	// Load .env file if it exists
	_ = godotenv.Load()

	AppConfig = &Config{
		MongoURI:      getString("MONGO_URI", "mongodb://localhost:27017"),
		MongoDB:       getString("MONGO_DB", "users"),
		MongoUsername: getString("MONGO_USERNAME", "admin"),
		MongoPassword: getString("MONGO_PASSWORD", "admin"),
		JWTSecret:     getString("JWT_SECRET", "secret"),
		JWTExpiration: getInt("JWT_EXPIRATION", 3600),
	}

	return AppConfig
}

func getString(key, fallback string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	return val
}

func getInt(key string, fallback int) int {
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
