package seed

import (
	"gin-seed/models"

	"gorm.io/gorm"
)

var users = []models.User{
	{
		Name: "Lorem",
	},
	{
		Name: "Ipsum",
	},
}

func Load(db *gorm.DB) {
	db.Create(&users)
}
