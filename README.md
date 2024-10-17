# bookstore-go

A simple CRUD API for managing a bookstore, built using Golang and MySQL.

## Features

- **Create a new book** (`POST /books`): Add a book by inserting its title, author, and publication date.
- **Get all books** (`GET /books`): Retrieve all the books in the bookstore.
- **Get a book by ID** (`GET /books/{id}`): Retrieve details of a specific book by its ID.
- **Update a book** (`PUT /books/{id}`): Update the details (title, author, publication) of a book by its ID.
- **Delete a book** (`DELETE /books/{id}`): Remove a book from the database by its ID.

## Technology Stack

- **Backend**: Golang
- **Database**: MySQL
- **Library**: Gorilla Mux (Router) and GORM 
