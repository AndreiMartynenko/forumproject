package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

var tpl *template.Template

var user *User

// dummy user data
// var users = map[string]string{"useremail1": "password", "useremail2": "password"}

// store the secret key in env variable in production
// var store = sessions.NewCookieStore([]byte("my_secret_key"))

func statusCode(w http.ResponseWriter, statusCodeValue int) {
	// put something here to check if errorpage has a valid template dynamically instead of manually and if not return 500 status code maybe?
	w.WriteHeader(statusCodeValue)
	errorinfo := struct {
		StatusCode    int
		StatusMessage string
	}{
		statusCodeValue,
		http.StatusText(statusCodeValue),
	}
	tpl.ExecuteTemplate(w, "error.html", errorinfo)
}

// func init executues from start, defore exec func main
func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" && r.URL.Path != "/" {
		statusCode(w, 404)
		// return here will stop execution this function
		return
	}
	// Render the index.html template
	templatePath := "templates/index.html"
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		statusCode(w, 404)
		return
	}

	// err = t.Execute(w, nil)

	// if err != nil {
	// 	statusCode(w, 500)
	// 	return
	// }

	t.Execute(w, user)
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	templatePath := "templates/login.html"
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		statusCode(w, 404)
		return
	}

	t.Execute(w, nil)
}

func signOutHandler(w http.ResponseWriter, r *http.Request) {
	user = nil
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func signupResultHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	userName := r.FormValue("username")
	emailSignUp := r.FormValue("signupemail")
	pswd := r.FormValue("pswd")

	user = &User{}
	user.Id = 1
	user.Username = userName
	user.Password = pswd
	user.Email = emailSignUp

	saveUser(user)

	// Check for email and username
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)

	// formSignUp := struct {
	// 	UserN string
	// 	Email string
	// 	PSWD  string
	// }{
	// 	UserN: userName,
	// 	Email: emailSignUp,
	// 	PSWD:  pswd,
	// }

	// tpl.ExecuteTemplate(w, "resultSignup.html", formSignUp)
}

func loginResultHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// 1. Get email and password from the parsed form and add into a map
	userEmail := r.FormValue("loginemail")
	userPswd := r.FormValue("loginpswd")

	user = getUserByEmailAndPassword(userEmail, userPswd)

	http.Redirect(w, r, "/", http.StatusSeeOther)

	//userCred := make(map[int]string)

	//where key is an id from db, values : email and password
	/*
		var forumUser = user.NewUser(db)

		type Data struct {
			Username string
			Password string
		}

		var data Data

		username := r.FormValue("username")
		password := r.FormValue("password")
		user := getUsers(username)

		if user.Username == username {
		}
		IsAuthorized = true
		Username = username
		http.Redirect(w, r, "/", 302)
		return

		password = "this password is incorrect"
		tpl.ExecuteTemplate(w, "login.html", data)

		return
		data.Username = "no such user"
		tpl.ExecuteTemplate(w, "login.html", data)
	*/
}

func dbHandler(w http.ResponseWriter, r *http.Request) {
	users := getUsers()
	if users != nil {
		fmt.Fprint(w, users)
	}
}

// Cookies, Sessions
// Map -?? Not the best solution. For testing purposes only.
var sessions = map[string]string{}

func getCookeisHandler() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//cookie - method of http.Request
		sessionID, err := r.Cookie("session_id")
		templatePath := "templates/login.html"
		//ErrNoCookie - seperate error cookie is not existing
		if err == http.ErrNoCookie {
			w.Write([]byte(templatePath))
			return
		} else if err != nil {
			PanicOnErr(err)

		}

		username, ok := sessions[sessionID.Value]
		if !ok {
			fmt.Fprint(w, "Session not found")
		} else {
			fmt.Fprint(w, "Welcome, "+username)
		}

		fmt.Fprint(w, "Welcome, "+sessionID.Value)
	})

	http.HandleFunc("/resultLogin", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		inputLogin := r.Form["loginemail"][0]
		expiration := time.Now().Add(365 * 24 * time.Hour)
		// cookie := http.Cookie{
		// 	Name:  "session_id",
		// 	Value: inputLogin,
		// 	//Expires - for how long we provide cookies
		// 	Expires: expiration,
		// }
		sessionID := RandStringRunes(32)
		sessions[sessionID] = inputLogin
		cookie := http.Cookie{Name: "session_id", Value: sessionID, Expires: expiration}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", http.StatusFound)
	})
	// ???
	// http.ListenAndServe(":8080", nil)
}

// To make cookies/sesions more secure
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
