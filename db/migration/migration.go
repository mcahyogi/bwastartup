package migration

import (
	"bwastartup/db"
	"bwastartup/model"
)

func RunMigration() {
	conn := db.SetupConnection()

	conn.AutoMigrate(&model.User{})
}
