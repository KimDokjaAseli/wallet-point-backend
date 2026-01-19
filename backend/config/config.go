package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerHost         string
	ServerPort         string
	ServerAddress      string
	GinMode            string
	DBHost             string
	DBPort             string
	DBUser             string
	DBPassword         string
	DBName             string
	JWTSecret          string
	JWTExpiryHours     int
	AllowedOrigins     string
	MaxUploadSize      int64
	UploadPath         string
	ExternalAPITimeout int
}

func LoadConfig() *Config {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}

	// Parse JWT expiry hours
	jwtExpiry, err := strconv.Atoi(getEnv("JWT_EXPIRY_HOURS", "24"))
	if err != nil {
		jwtExpiry = 24
	}

	// Parse max upload size
	maxUploadSize, err := strconv.ParseInt(getEnv("MAX_UPLOAD_SIZE", "10485760"), 10, 64)
	if err != nil {
		maxUploadSize = 10485760 // 10MB default
	}

	// Parse external API timeout
	apiTimeout, err := strconv.Atoi(getEnv("EXTERNAL_API_TIMEOUT", "30"))
	if err != nil {
		apiTimeout = 30
	}

	serverHost := getEnv("SERVER_HOST", "0.0.0.0")
	serverPort := getEnv("SERVER_PORT", "8080")

	return &Config{
		ServerHost:         serverHost,
		ServerPort:         serverPort,
		ServerAddress:      serverHost + ":" + serverPort,
		GinMode:            getEnv("GIN_MODE", "debug"),
		DBHost:             getEnv("DB_HOST", "gateway01.ap-southeast-1.prod.aws.tidbcloud.com"),
		DBPort:             getEnv("DB_PORT", "4000"),
		DBUser:             getEnv("DB_USER", "3Rtj2MNUvxPhvBM.root"),
		DBPassword:         getEnv("DB_PASSWORD", ""),
		DBName:             getEnv("DB_NAME", "walletpointapp_wallet_point"),
		JWTSecret:          getEnv("JWT_SECRET", "change-this-secret-key-in-production"),
		JWTExpiryHours:     jwtExpiry,
		AllowedOrigins:     getEnv("ALLOWED_ORIGINS", "*"),
		MaxUploadSize:      maxUploadSize,
		UploadPath:         getEnv("UPLOAD_PATH", "./public/uploads"),
		ExternalAPITimeout: apiTimeout,
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
