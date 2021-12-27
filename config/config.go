package config

import (
	"database/sql"
	"fmt"
	"github.com/edwardsuwirya/wmbMenuMgmt/delivery"
	"github.com/edwardsuwirya/wmbMenuMgmt/entity"
	"github.com/edwardsuwirya/wmbMenuMgmt/manager"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

type Config struct {
	InfraManager   manager.Infra
	RepoManager    manager.RepoManager
	UseCaseManager manager.UseCaseManager
	RouterEngine   *gin.Engine
	ApiBaseUrl     string
	runMigration   string
}

func NewConfig() *Config {
	runMigration := os.Getenv("DB_MIGRATION")
	apiHost := os.Getenv("API_HOST")
	apiPort := os.Getenv("API_PORT")

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", dbHost, dbUser, dbPassword, dbName, dbPort)
	infraManager := manager.NewInfra(dsn)
	repoManager := manager.NewRepoManager(infraManager)
	useCaseManager := manager.NewUseCaseManger(repoManager)

	config := new(Config)
	config.InfraManager = infraManager
	config.RepoManager = repoManager
	config.UseCaseManager = useCaseManager

	r := gin.Default()
	delivery.NewServer(r, useCaseManager)
	config.RouterEngine = r

	config.ApiBaseUrl = fmt.Sprintf("%s:%s", apiHost, apiPort)
	config.runMigration = runMigration

	return config
}

func (c *Config) RunMigration() {
	if c.runMigration == "Y" || c.runMigration == "y" {
		db := c.InfraManager.SqlDb()
		err := db.AutoMigrate(&entity.Menu{})
		db.Unscoped().Where("id like ?", "%%").Delete(entity.Menu{})
		db.Model(&entity.Menu{}).Save([]entity.Menu{
			{
				ID:       "M0001",
				MenuName: "Sayur Lodeh",
				Price:    2000,
			},
			{
				ID:       "M0002",
				MenuName: "Perkedel",
				Price:    1000,
			},
			{
				ID:       "M0003",
				MenuName: "Nasi Putih",
				Price:    5000,
			},
			{
				ID:       "M0004",
				MenuName: "Tempe Orek",
				Price:    1000,
			},
			{
				ID:       "M0005",
				MenuName: "Tumis Kangkung",
				Price:    2000,
			},
			{
				ID:       "M0006",
				MenuName: "Es Teh Tawar",
				Price:    1000,
			},
			{
				ID:       "M0007",
				MenuName: "Es Teh Manis",
				Price:    1500,
			},
		})

		if err != nil {
			log.Fatalln(err)
		}
	}
}

func (c *Config) StartEngine() {
	if !(c.runMigration == "Y" || c.runMigration == "y") {
		db, _ := c.InfraManager.SqlDb().DB()
		defer func(db *sql.DB) {
			err := db.Close()
			if err != nil {
				log.Fatalln(err)
			}
		}(db)
		err := c.RouterEngine.Run(c.ApiBaseUrl)
		if err != nil {
			log.Fatal(err)
		}
	}
}
