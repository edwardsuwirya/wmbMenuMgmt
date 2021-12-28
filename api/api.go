package api

import (
	"database/sql"
	"github.com/edwardsuwirya/wmbMenuMgmt/config"
	"github.com/edwardsuwirya/wmbMenuMgmt/delivery"
	"github.com/edwardsuwirya/wmbMenuMgmt/entity"
	"github.com/edwardsuwirya/wmbMenuMgmt/manager"
	"log"
)

type Server interface {
	Run()
}

type server struct {
	config  *config.Config
	infra   manager.Infra
	usecase manager.UseCaseManager
}

func NewApiServer() Server {
	appConfig := config.NewConfig()
	infra := manager.NewInfra(appConfig)
	repo := manager.NewRepoManager(infra)
	usecase := manager.NewUseCaseManger(repo)
	return &server{
		config:  appConfig,
		infra:   infra,
		usecase: usecase,
	}
}

func (s *server) Run() {
	if !(s.config.RunMigration == "Y" || s.config.RunMigration == "y") {
		db, _ := s.infra.SqlDb().DB()
		defer func(db *sql.DB) {
			err := db.Close()
			if err != nil {
				log.Fatalln(err)
			}
		}(db)
		delivery.NewServer(s.config.RouterEngine, s.usecase)
		err := s.config.RouterEngine.Run(s.config.ApiBaseUrl)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		db := s.infra.SqlDb()
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
