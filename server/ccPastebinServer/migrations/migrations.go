package migrations

import (
	"ccPasteBinServer/database"
	"ccPasteBinServer/model"
)

func Migrate() {
   
	database.DBConnection.AutoMigrate(model.Note{})


}