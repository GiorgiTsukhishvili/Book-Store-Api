package initializers

import "log"

func EnumsInitializer() {
	sql := `
	DO $$ BEGIN
		IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'user_type') THEN
			CREATE TYPE user_type AS ENUM ('user', 'business', 'admin');
		END IF;
	END $$;
	`
	err := DB.Exec(sql).Error

	if err != nil {
		log.Fatal("Failed to initialize enums:", err)
	}

	log.Println("Database enums inserted successfully")
}
