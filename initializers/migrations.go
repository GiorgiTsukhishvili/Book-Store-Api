package initializers

import "log"

func MigrationsInitializer() {
	err := DB.AutoMigrate()

	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Database migrated successfully")
}
