package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var DB *gorm.DB

func ConnectToDB()  {
	// postgres://wkcysntv:mL7ws0aotcvU56zxi9Iw3XXYBuNycShG@raja.db.elephantsql.com/wkcysntv
	var err error
	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to DB")
	}
}