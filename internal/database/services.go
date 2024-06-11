package database

// Create a new book
func CreateBook(book Book) (Book, error) {
	tx := DB.Create(&book)
	if tx.Error != nil {
		return Book{}, tx.Error
	}

	return book, nil
}

// Get a list of all books
func GetBooks() ([]Book, error) {
	var books []Book

	result := DB.Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}

	return books, nil
}

// Get a single book by ID
func GetBookByID(book Book) (Book, error) {
	tx := DB.First(&book)

	if tx.Error != nil {
		return Book{}, tx.Error
	}

	return book, nil
}

// Delete a new book
func DeleteBookByID(book Book) error {
	tx := DB.Delete(&book)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
