package initializers

import "jwt-auth-gin-app/models"

func migrateDB() {
	DB.AutoMigrate(&models.User{})
}
