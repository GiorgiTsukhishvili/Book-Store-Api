package factories

import (
	"log"

	"github.com/GiorgiTsukhishvili/BookShelf-Api/models"
	"gorm.io/gorm"
)

func BookFactory(db *gorm.DB) {
	books := []models.Book{
		{Name: "Pride and Prejudice", Description: "A romantic novel of manners.", Price: "9.99", Image: "pride_and_prejudice.jpg", AuthorID: 1, UserID: 1},
		{Name: "Adventures of Huckleberry Finn", Description: "A novel about the journey of a young boy and a runaway slave.", Price: "12.99", Image: "huck_finn.jpg", AuthorID: 2, UserID: 2},
	}
	for _, book := range books {
		if err := db.Create(&book).Error; err != nil {
			log.Fatalf("failed to seed books: %v", err)
		}
	}

	log.Println("Book seeded successfully")
}
