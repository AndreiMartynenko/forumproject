package main

import (
	"fmt"
	"net/http"
	"text/template"
)

var serverPort = ":8080"

// type User struct {
// 	Id       string
// 	UserName string
// 	Email    string
// 	Password string
// 	Hobbies  []string
// }

//var user User{}

//Methods --------------------------------

func (u *User) getAllInfo() string {
	return fmt.Sprintf("User name is: %s. His email and password are %s, "+
		"£%s", u.UserName, u.Email, u.Password)
}

func (u *User) setNewName(newName string) {
	u.UserName = newName
}

// //Methods --------------------------------

func homePage(w http.ResponseWriter, r *http.Request) {

	//user := User{"Andrei", "yahoo@mail.com", "qwerty123", []string{"Music", "Coding", "Hacking"}}
	//user.setNewName("Andrew")
	//fmt.Fprintf(w, "User name is: " + user.User)
	//fmt.Fprintf(w, bob.getAllInfo())
	//fmt.Fprintf(w, `<h1>Hello User</h1><b>Main Text</b>`)
	tmpl, err := template.ParseFiles("templates/login.html")
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Fprintf(w, user.getAllInfo())
	tmpl.Execute(w, user)
}

func newPostPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "New Post Page")
}

func handleRequest() {
	// a multiplexer, which redirects the request to the correct handler to process the request.
	//Static
	const portNumber = ":8080"
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/signup", signupHandler)
	mux.HandleFunc("/login", loginHandler)
	mux.HandleFunc("/mainpage", mainPageHandler)
	mux.HandleFunc("/checkuserlogin", loginCheckHandler)
	mux.HandleFunc("/createpost", createPostHandler)
	mux.HandleFunc("/savepost", savePostHandler)
	// mux.HandleFunc("/releasepost", checkPostHandler)

	//http.HandleFunc("/", homePage)
	//http.HandleFunc("/newpost/", newPostPage)
	// fmt.Printf("Application started on port %v", serverPort)
	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	fmt.Printf("Starting application on port %s", portNumber)
	server.ListenAndServe()
}

// func main() {
// 	//bob := User{name: "Bob", age: 25, money: -50, avg_grades: 4.2, happiness: 0.8}

// 	handleRequest()
// }

func main() {
	createPostsTable()
	createUsersTable()
	//printUsers()
	//printPosts()
	handleRequest()
}
