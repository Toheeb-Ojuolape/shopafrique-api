package initializers

import "github.com/Toheeb-Ojuolape/shopafrique-api/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})

}
