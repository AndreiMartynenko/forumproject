package main

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func printUsers() {
	//fmt.Println("Printing Users")
	sql := "SELECT username FROM users"
	rows, _ := db.Query(sql)
	for rows.Next() {
		//fmt.Println("User")
		user := User{}
		rows.Scan(&(user.UserName))
		fmt.Println(user)
	}
	rows.Close()

}
func createUsersTable() /**NewUserDataBase*/ {
	db, _ = sql.Open("sqlite3", "forum.db")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
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
	fmt.Println("Table for users created successfully!")
	query.Close()
	//db.Close()
}

// insert data
func insertUserinTable(user *User) {
	db, _ := sql.Open("sqlite3", "forum.db")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	statement, _ := db.Prepare("INSERT INTO users (username, email, password) VALUES(?, ?, ?)")
	_, err := statement.Exec(user.UserName, user.Email, user.Password)
	if err != nil {
		fmt.Println(err)
		return
	}
	statement.Close()

	defer db.Close()
}

func getUserFromTable() *[]User {
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
		row.Scan(&(user.UserName), &(user.Email), &(user.Password))
		users = append(users, user)
		//fmt.Println(users)
	}
	return &users
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
		row.Scan(&(user.UserName), &(user.Email), &(user.Password))
		row.Close()
		db.Close()
		return &user
	}
	row.Close()
	db.Close()
	return nil
}

func userExists(user User) bool {

	//stmt := `SELECT username, email FROM users WHERE username = ?, email = ?`
	rows, _ := db.Query("SELECT * FROM users WHERE username = ? OR email = ?", user.UserName, user.Email)

	for rows.Next() {
		rows.Close()
		return true
	}
	return false
}

func userLogin(user User) *User {
	//fmt.Println(user)
	rows, _ := db.Query("SELECT * FROM users WHERE email = ? AND password = ?", strings.TrimSpace(user.Email), strings.TrimSpace(user.Password))
	loginUser := User{}
	for rows.Next() { // Iterate and fetch the records from result cursor
		rows.Scan(&(loginUser.Id), &(loginUser.UserName), &(loginUser.Email), &(loginUser.Password))
		rows.Close()
		return &loginUser
	}
	return nil

}
