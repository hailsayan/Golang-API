package migrations

import (
	"github.com/hailsayan/Golang-API/config"
	"github.com/hailsayan/Golang-API/data/db"
	"github.com/hailsayan/Golang-API/data/models"
	"github.com/hailsayan/Golang-API/pkg/logging"
)

var logger = logging.NewLogger(config.GetConfig())

func Up_1() {
	database := db.GetDb()

	tables := []interface{}{}

	country := models.Country{}
	city := models.City{}

	if !database.Migrator().HasTable(country) {
		tables = append(tables, country)
	}

	if !database.Migrator().HasTable(city) {
		tables = append(tables, city)
	}

	database.Migrator().CreateTable(tables...)
	logger.Info(logging.Postgres, logging.Migration, "tables created", nil)

}

func Down_1() {

}
