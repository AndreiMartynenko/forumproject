package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"net/http"
// 	"text/template"
// 	"unicode"

// 	_ "github.com/mattn/go-sqlite3"
// 	"golang.org/x/crypto/bcrypt"
// )

// var tpl *template.Template
// var db *sql.DB
// var port = ":8080"

// func main() {
// 	tpl, _ = template.ParseGlob("templates/*.html")
// 	var err error
// 	db, err = sql.Open("sqlite3", "table.db")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	defer db.Close()
// 	http.HandleFunc("/register", registerHandler)
// 	http.HandleFunc("/registerauth", registerAuthHandler)
// 	fmt.Println("Started on port: 8080")
// 	http.ListenAndServe(port, nil)
// }

// func registerHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("registerHandler is running")
// 	tpl.ExecuteTemplate(w, "register.html", nil)
// }

// func registerAuthHandler(w http.ResponseWriter, r *http.Request) {

// 	/*

// 	   1. check username criteria
// 	   2. check password criteria
// 	   3. check if username is already exists in database
// 	   4. create bcrypt hash from password
// 	   5. insert username and password hash in database

// 	*/

// 	fmt.Println("registerAuthHandler is running")
// 	r.ParseForm()
// 	username := r.FormValue("username")
// 	//check username for only alphanumeric characters
// 	var nameAlphaNumeric = true
// 	for _, char := range username {
// 		// func Isletter(r rune) bool, func IsNUmber (r rune) bool
// 		// if !unicode.IsLetter(char) && !unicode.IsNumber(char) {
// 		if unicode.IsLetter(char) == false && unicode.IsNumber(char) == false {
// 			nameAlphaNumeric = false
// 		}
// 	}
// 	// check username length
// 	var nameLength bool
// 	if 5 <= len(username) && len(username) <= 50 {
// 		nameLength = true
// 	}
// 	// check password criteria
// 	password := r.FormValue("password")
// 	fmt.Println("password:", password, "\npswdLength:", len(password))
// 	// variables that must pass for password creation criteria
// 	var pswdLowercase, pswdUppercase, pswdNumber, pswdSpecial, pswdLength, pswdNoSpaces bool
// 	pswdNoSpaces = true
// 	for _, char := range password {
// 		switch {
// 		// func IsLower(r rune) bool
// 		case unicode.IsLower(char):
// 			pswdLowercase = true
// 		// func IsUpper(r rune) bool
// 		case unicode.IsUpper(char):
// 			pswdUppercase = true
// 		// func IsNumber(r rune) bool
// 		case unicode.IsNumber(char):
// 			pswdNumber = true
// 		// func IsPunct(r rune) bool, func IsSymbol(r rune) bool
// 		case unicode.IsPunct(char) || unicode.IsSymbol(char):
// 			pswdSpecial = true
// 		// func IsSpace(r rune) bool, type rune = int32
// 		case unicode.IsSpace(int32(char)):
// 			pswdNoSpaces = false
// 		}
// 	}
// 	if 11 < len(password) && len(password) < 60 {
// 		pswdLength = true
// 	}
// 	fmt.Println("pswdLowercase:", pswdLowercase, "\npswdUppercase:", pswdUppercase, "\npswdNumber:", pswdNumber, "\npswdSpecial:", pswdSpecial, "\npswdLength:", pswdLength, "\npswdNoSpaces:", pswdNoSpaces, "\nnameAlphaNumeric:", nameAlphaNumeric, "\nnameLength:", nameLength)
// 	if !pswdLowercase || !pswdUppercase || !pswdNumber || !pswdSpecial || !pswdLength || !pswdNoSpaces || !nameAlphaNumeric || !nameLength {
// 		tpl.ExecuteTemplate(w, "register.html", "please check username and password criteria")
// 		return
// 	}
// 	// check if username already exists for availability
// 	stmt := "SELECT UserID FROM bcrypt WHERE username = ?"
// 	row := db.QueryRow(stmt, username)
// 	var uID string
// 	err := row.Scan(&uID)
// 	if err != sql.ErrNoRows {
// 		fmt.Println("username already exists, err:", err)
// 		tpl.ExecuteTemplate(w, "register.html", "username already taken")
// 		return
// 	}
// 	// create hash from password
// 	var hash []byte
// 	// func GenerateFromPassword(password []byte, cost int) ([]byte, error)
// 	hash, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// 	if err != nil {
// 		fmt.Println("bcrypt err:", err)
// 		tpl.ExecuteTemplate(w, "register.html", "there was a problem registering account")
// 		return
// 	}
// 	fmt.Println("hash:", hash)
// 	fmt.Println("string(hash):", string(hash))
// 	// func (db *DB) Prepare(query string) (*Stmt, error)
// 	var insertStmt *sql.Stmt
// 	insertStmt, err = db.Prepare("INSERT INTO bcrypt (Username, Hash) VALUES (?, ?);")
// 	if err != nil {
// 		fmt.Println("error preparing statement:", err)
// 		tpl.ExecuteTemplate(w, "register.html", "there was a problem registering account")
// 		return
// 	}
// 	defer insertStmt.Close()
// 	var result sql.Result
// 	//  func (s *Stmt) Exec(args ...interface{}) (Result, error)
// 	result, err = insertStmt.Exec(username, hash)
// 	rowsAff, _ := result.RowsAffected()
// 	lastIns, _ := result.LastInsertId()
// 	fmt.Println("rowsAff:", rowsAff)
// 	fmt.Println("lastIns:", lastIns)
// 	fmt.Println("err:", err)
// 	if err != nil {
// 		fmt.Println("error inserting new user")
// 		tpl.ExecuteTemplate(w, "register.html", "there was a problem registering account")
// 		return
// 	}
// 	fmt.Fprint(w, "congrats, your account has been successfully created")
// }
