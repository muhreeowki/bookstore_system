package database

import "log"

func CreateBook(book Book) (Book, error) {
	tx := DB.Create(&book)
	if tx.Error != nil {
		log.Println("Error creating book", tx.Error)
		return Book{}, tx.Error
	}

	log.Printf("Book created: %v", book)
	log.Printf("Transaction: %v", tx)
	return book, nil
}

// Get a Book by ID
func GetBooks() ([]Book, error) {
	var books []Book
	result := DB.Find(&books)
	if result.Error != nil {
		log.Println("Error getting books", result.Error)
		return nil, result.Error
	}
	return books, nil
}
