package main

import (
	"fmt"
	"net/http"
	"text/template"
)

var user *User
var post *Post

// var db *sql.DB

func indexHandler(w http.ResponseWriter, r *http.Request) {
	//user := User{"Andrei", "yahoo@mail.com", "qwerty123", []string{"Music", "Coding", "Hacking"}}
	t, err := template.ParseFiles("templates/index.html")

	if err != nil {
		fmt.Println(err)
	}
	t.ExecuteTemplate(w, "index.html", nil)

}
func createPostHandler(w http.ResponseWriter, r *http.Request) {
	templatePath := "templates/createpost.html"
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		showError(w, "404 PAGE NOT FOUND", http.StatusNotFound)
		return
	}
	t.Execute(w, nil)
}

/* GET -  entering a page
   POST - when data transferring from 'form' form
*/

func savePostHandler(w http.ResponseWriter, r *http.Request) {

	title := r.FormValue("title")
	postText := r.FormValue("post")

	if title == "" || postText == "" {
		fmt.Fprintf(w, "Empty field. Please check.")
	}

	post = &Post{}
	post.Title = title
	post.Text = postText

	insertPostinTable(post)
	http.Redirect(w, r, "mainpage", http.StatusSeeOther)

}

func checkPostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "mainpage", http.StatusSeeOther)
		return
	}

	title := r.FormValue("title")
	postText := r.FormValue("post")

	post1 := &Post{}
	post1.Title = title
	post1.Text = postText

	postUser := userPost(*post1)
	if postUser == nil {
		http.Redirect(w, r, "createpost", http.StatusSeeOther)
		return
	}
	post = &Post{}
	post.Title = postUser.Title
	post.Text = postUser.Text

	http.Redirect(w, r, "mainpage", http.StatusSeeOther)
}
func signupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "mainpage", http.StatusSeeOther)
		return
	}
	userName := r.FormValue("username")
	emailSignUp := r.FormValue("email")
	pswd := r.FormValue("password")

	user = &User{}
	user.UserName = userName
	user.Password = pswd
	user.Email = emailSignUp

	insertUserinTable(user)
	http.Redirect(w, r, "mainpage", http.StatusSeeOther)
}
func signOutHandler(w http.ResponseWriter, r *http.Request) {
	user = nil
	http.Redirect(w, r, "login", http.StatusTemporaryRedirect)
}
func mainPageHandler(w http.ResponseWriter, r *http.Request) {
	// if r.Method != "POST" {
	// 	http.Redirect(w, r, "signup", http.StatusSeeOther)
	// 	return
	// }
	// userName := r.FormValue("username")
	// emailSignUp := r.FormValue("email")
	// pswd := r.FormValue("password")

	// user := &User{}
	// //user.Id = "1"
	// user.UserName = userName
	// user.Password = pswd
	// user.Email = emailSignUp

	// userExists := userExists(*user)

	// if userExists {
	// 	http.Redirect(w, r, "signup", http.StatusTemporaryRedirect)
	// 	return
	// }
	// insertDatainTable(user)

	templatePath := "templates/mainpage.html"
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		showError(w, "404 PAGE NOT FOUND", http.StatusNotFound)
		return
	}
	t.Execute(w, user)
}
func loginHandler(w http.ResponseWriter, r *http.Request) {

	templatePath := "templates/login.html"
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		showError(w, "404 PAGE NOT FOUND", http.StatusNotFound)
		return
	}
	t.Execute(w, nil)
}
func loginCheckHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "mainpage", http.StatusSeeOther)
		return
	}
	loginEmail := r.FormValue("loginemail")
	loginPswd := r.FormValue("loginpswd")
	user1 := &User{}
	user1.Email = loginEmail
	user1.Password = loginPswd

	loginUser := userLogin(*user1)
	if loginUser == nil {
		http.Redirect(w, r, "login", http.StatusSeeOther)
		return
	}
	//fmt.Println(loginUser)
	user = &User{}
	user.UserName = loginUser.UserName
	user.Email = loginUser.Email

	http.Redirect(w, r, "mainpage", http.StatusSeeOther)
	//return
}

// Render the error.html template
func showError(w http.ResponseWriter, message string, statusCode int) {
	t, err := template.ParseFiles("templates/error.html")
	if err == nil {
		w.WriteHeader(statusCode)
		t.Execute(w, message)
	}
}
