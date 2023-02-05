package main

import (
	"database/sql"
	"fmt"

	//driver for sqlite3
	_ "github.com/mattn/go-sqlite3"
)

// _ is here to stop alerting "I am not using this driver"

//Creating database

/* We’re using db.Query() to send the query to the database. We check the error, as usual.
We defer rows.Close(). This is very important.
We iterate over the rows with rows.Next().
We read the columns in each row into variables with rows.Scan().
*/

func createUsersTable() /**NewUserDataBase*/ {
	db, err := sql.Open("sqlite3", "forum.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	users_table := `CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY,
        username TEXT NOT NULL,
        email TEXT NOT NULL,
        password TEXT NOT NULL);`
	query, err := db.Prepare(users_table)
	if err != nil {
		fmt.Println(err)
		return
	}
	query.Exec()
	fmt.Println("Table created successfully!")

	query.Close()
	db.Close()
}

func getUsers() *[]User {
	db, err := sql.Open("sqlite3", "forum.db")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	row, err := db.Query("SELECT * FROM users")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer row.Close()

	users := []User{}
	for row.Next() { // Iterate and fetch the records from result cursor
		user := User{}
		row.Scan(&(user.Id), &(user.Username), &(user.Email), &(user.Password))
		users = append(users, user)
		fmt.Println(users)
	}
	return &users
}

func saveUser(user *User) {
	db, err := sql.Open("sqlite3", "forum.db")
	if err != nil {
		fmt.Println(err)
		return
	}

	insertUser := "INSERT INTO users (username, email, password) VALUES(?, ?, ?)"
	_, err = db.Exec(insertUser, user.Username, user.Email, user.Password)
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Close()

}

func getUserByEmailAndPassword(userEmail string, userPswd string) *User {
	// Establish connection. To connect with the database, we call the Open() function on the sql instance like so:
	db, err := sql.Open("sqlite3", "forum.db")
	if err != nil {
		db.Close()
		fmt.Println(err)
		return nil
	}
	row, err := db.Query("SELECT * FROM users WHERE email = ? AND password = ?", userEmail, userPswd)
	if err != nil {
		row.Close()
		db.Close()
		fmt.Println(err)
		return nil
	}
	for row.Next() {
		user := User{}
		row.Scan(&(user.Id), &(user.Username), &(user.Email), &(user.Password))
		row.Close()
		db.Close()
		return &user
	}
	row.Close()
	db.Close()
	return nil
}

// Check if username, email, password exist.
// func validEmail(email string) bool {
// 	_, err := mail.ParseAddress(email)
// 	return err == nil

// }

//requesting data from db users for email and pswd to check for their existing
