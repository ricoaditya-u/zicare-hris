package main

import (
	"github.com/ricoaditya-u/hris-master/db"
	"github.com/ricoaditya-u/hris-master/models"
)

func init() {
	db.LoadEnvVariables()
	db.ConnectToDB()
}

func main() {
	db.DB.AutoMigrate(
		// &models.Grade{},
		// &models.JobDescription{},
		// &models.Level{},
		// &models.Division{},
		// &models.Department{},
		// &models.Supervision{},
		// &models.Employee{},
		// &models.User{},
		// &models.Mpp{},
		&models.Reqheadcount{},
		// &models.Casbin_rule{},
	)
}
