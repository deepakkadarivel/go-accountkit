package factory

import (
	"github.com/jinzhu/gorm"
	"Accountkit/config"
	"log"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/golang/glog"
)

type DatabaseFactory interface {
	InitializeDatabaseConnection()
	DBConnection() *gorm.DB
	CloseConnection()
}

type PostgresDatabaseFactory struct {
	appConfig *config.Config
	db *gorm.DB
}

func NewPostgresDatabaseFactory(appConfig *config.Config) *PostgresDatabaseFactory {
	return &PostgresDatabaseFactory{
		appConfig: appConfig,
	}
}

func (factory *PostgresDatabaseFactory) InitializeDatabaseConnection()  {
	if factory.db == nil {
		db, err := gorm.Open("postgres", factory.appConfig.DB.ConnectionString())
		if err != nil {
			message := fmt.Sprintf("DB Connection establishment failed: %v", err)
			glog.Fatalln(message)
			log.Fatalf(message)
		}
		factory.db = db
		glog.Infoln("Postgres DB Initialized.")
	}
}

func (factory *PostgresDatabaseFactory) DBConnection() *gorm.DB {
	return factory.db
}

func (factory *PostgresDatabaseFactory) CloseConnection() {
	err := factory.db.Close()
	if err != nil {
		log.Fatalf("Error closing database connection: %v", err.Error())
	}
}
