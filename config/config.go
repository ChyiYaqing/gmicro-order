package config

import (
	"log"
	"os"
	"strconv"
)

func GetEnv() string {
	return getEnvironmentValue("ENV")
}

func GetSqliteDB() string {
	return getEnvironmentValue("SQLITE_DB")
}

func GetApplicationPort() int {
	portStr := getEnvironmentValue("APPLICATION_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("port: %s is invalid", portStr)
	}
	return port
}

func GetPaymentServiceUrl() string {
	return getEnvironmentValue("PAYMENT_SERVICE_URL")
}

func GetUserServiceUrl() string {
	return getEnvironmentValue("USER_SERVICE_URL")
}

func GetShippingServiceUrl() string {
	return getEnvironmentValue("USER_SERVICE_URL")
}

func getEnvironmentValue(key string) string {
	if os.Getenv(key) == "" {
		log.Fatalf("%s environment variable is missing.", key)
	}
	return os.Getenv(key)
}
