package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Id       string
	UserName string
	Email    string
	Password string
}

/*sql.DB performs some important tasks for you behind the scenes:

It opens and closes connections to the actual underlying database, via the driver.
It manages a pool of connections as needed, which may be a variety of things as mentioned.
*/

type ForumUser struct {
	DB *sql.DB
}

var db *sql.DB

func createUsersTable() {
	// db, err := sql.Open("sqlite3", "table.db")
	// checkErr(err)
	//insert

	users_table := `CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY,
        username TEXT NOT NULL,
        email TEXT NOT NULL,
        password TEXT NOT NULL);`
	query, err := db.Prepare(users_table)
	checkErr(err)
	query.Exec()
	fmt.Println("Table created successfully!")

	query.Close()
	//db.Close()

}

func insertDataTable() {
	q, err := db.Prepare("INSERT INTO users (username, email, password) values(?,?,?)")
	checkErr(err)

	data, err := q.Exec("Andrew", "email@gmail.com", "qwerty123")
	checkErr(err)

	id, err := data.LastInsertId()
	checkErr(err)

	fmt.Println(id)

}

// func getNewUser() {

// }

// Query
/*
We’re using db.Query() to send the query to the database. We check the error, as usual.
We defer rows.Close(). This is very important.
We iterate over the rows with rows.Next().
We read the columns in each row into variables with rows.Scan().
We check for errors after we’re done iterating over the rows.
*/

func getDataUsersfromTable() (users []User, err error) {
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		return users, err
	}
	defer rows.Close()
	//webUser := User{}
	var id string
	var email string
	var username string
	var password string
	//webUsers := []string{}
	for rows.Next() {
		err = rows.Scan(&(id), &(email), &(username), &(password))
		if err != nil {
			return users, err
		}
		tableUser := User{
			Id:       id,
			Email:    email,
			UserName: username,
			Password: password,
		}
		users = append(users, tableUser)

		// err := rows.Scan(&(user.Id), &(user.UserName), &(user.Email), &(user.Password))
		// checkErr(err)
		// webUsers = append(webUsers, user.UserName)
	}
	//rows.Close()
	fmt.Println(users)
	return users, err
	//return webUsers
}

// Adding New User
func newUser(email, username, password string) *User {
	return &User{Email: email, UserName: username, Password: password}
}

// Get User by ID
func getUserByEmailandPassword(userEmail, userPassword string) *User {
	// rows, err := db.Query("SELECT * FROM users")
	// checkErr(err)
	// err = db.Get(&(u), &(rows), &(userId))
	db, err := sql.Open("sqlite3", "table.db")
	if err != nil {
		db.Close()
		fmt.Println(err)
		return nil
	}
	row, err := db.Query("SELECT * FROM users WHERE email = ? AND password = ?", userEmail, userPassword)
	if err != nil {
		row.Close()
		db.Close()
		fmt.Println(err)
		return nil
	}
	for row.Next() {
		user := User{}
		row.Scan(&(user.Id), &(user.UserName), &(user.Email), &(user.Password))
		row.Close()
		db.Close()
		return &user
	}
	row.Close()
	db.Close()
	return nil

}

//Getting the data

// func getEmailFromTable() []string {

// 	rows, err := db.Query("SELECT * FROM users")
// 	checkErr(err)
// 	defer rows.Close()

// 	userEmail := []string{}
// 	for rows.Next() {
// 		err := rows.Scan(&(user.Email))
// 		checkErr(err)
// 		fmt.Println(user.Email)
// 		userEmail = append(userEmail, user.Email)
// 	}
// 	err = rows.Err()
// 	checkErr(err)
// 	rows.Close()
// 	return userEmail
// }

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	email := "yahoo@yahoo.com"
	password := "qwerty"
	username := "Mr Smith"
	//To create a sql.DB, you use sql.Open(). This returns a *sql.DB:
	db, _ = sql.Open("sqlite3", "table.db")
	//checkErr(err)
	defer db.Close()

	createUsersTable()
	getDataUsersfromTable()
	//getEmailFromTable()
	newUser(email, username, password)
	getUserByEmailandPassword(email, password)

}
