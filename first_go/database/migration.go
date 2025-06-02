package database

import (
	"fmt"
	"log"
)

func MigrateDB() {

	createUserTable :=
		`CREATE TABLE IF NOT EXISTS users (
	id INT AUTO_INCREMENT PRIMARY KEY,
	microsoft_id INT,
	name VARCHAR(255) NOT NULL,
	password VARCHAR(255) NOT NULL,
	email VARCHAR(255) NOT NULL UNIQUE,
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP);`

	createPostsTable :=
		`CREATE TABLE IF NOT EXISTS posts (
	id INT AUTO_INCREMENT PRIMARY KEY,
	title VARCHAR(255) NOT NULL,
	content TEXT NOT NULL,
	users_id INT NOT NULL,
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY (users_id) REFERENCES users(id)
	);`

	if _, err := DB.Exec(createUserTable); err != nil {
		log.Fatal("failed to create user table:", err)
	}

	if _, err := DB.Exec(createPostsTable); err != nil {
		log.Fatal("failed to create posts table:", err)
	}

	fmt.Println("Database Tables Migrated Successfully")
}
