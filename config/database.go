package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func InitDatabase() (*gorm.DB, error) {
	var conn string
	var driver string
	conn = "host=" + DotEnvVariable("DB_HOST") + " port=" + DotEnvVariable("DB_PORT") + " user=" + DotEnvVariable("DB_USER") + " dbname=" + DotEnvVariable("DB_NAME") + " password= " + DotEnvVariable("DB_PASSWORD")
	driver = DotEnvVariable("DB_DRIVER")

	return gorm.Open(driver, conn)
}
