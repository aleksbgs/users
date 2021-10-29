package database

import (
	"github.com/aleksbgs/users/src/models"
	"github.com/aleksbgs/users/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error
//	host := utils.ViperEnvVariable("DBDOCKERHOST")
	dbuser := utils.ViperEnvVariable("DBUSER")
	password := utils.ViperEnvVariable("DBPASSWORD")
	dbname := utils.ViperEnvVariable("DB_USERS")
    //microservice dsn
	//dsn := "host=" + host + " " + "user=" + dbuser + " " + "password=" + password + " " + "dbname=" + dbname + " " + "port=5432 sslmode=disable"

	//local dsn
	dsn := "host=" + "localhost" + " " + "user=" + dbuser + " " + "password=" + password + " " + "dbname=" + dbname + " " + "port=5434 sslmode=disable"

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Could not connect with")
	}
}

func AutoMigrate() {
	DB.AutoMigrate(models.User{})
}
