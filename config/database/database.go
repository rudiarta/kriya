package database

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func InitDatabase() (*gorm.DB, error) {
	var conn string
	var driver string
	conn = "host=" + os.Getenv("DB_HOST") + " port=" + os.Getenv("DB_PORT") + " user=" + os.Getenv("DB_USER") + " dbname=" + os.Getenv("DB_NAME") + " password= " + os.Getenv("DB_PASSWORD")
	driver = os.Getenv("DB_DRIVER")

	return gorm.Open(driver, conn)
}
