package config

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

type Config struct {
	RouterEngine   *gin.Engine
	ApiBaseUrl     string
	RunMigration   string
	DataSourceName string
}

func NewConfig() *Config {
	config := new(Config)
	runMigration := os.Getenv("DB_MIGRATION")
	apiHost := os.Getenv("API_HOST")
	apiPort := os.Getenv("API_PORT")

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", dbHost, dbUser, dbPassword, dbName, dbPort)
	config.DataSourceName = dsn
	r := gin.Default()
	config.RouterEngine = r

	config.ApiBaseUrl = fmt.Sprintf("%s:%s", apiHost, apiPort)
	config.RunMigration = runMigration

	return config
}
