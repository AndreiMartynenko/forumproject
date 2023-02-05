package main

import (
	"fmt"
	"net/http"
)

func main() {

	const portNumber = ":8080"
	createUsersTable()
	getCookeisHandler()
	// a multiplexer, which redirects the request to the correct handler to process the request.
	//Static
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	// // Static
	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	// fmt.Printf("Starting application on port %s", portNumber)
	// err := http.ListenAndServe(portNumber, nil)
	// if err != nil {
	// 	fmt.Println("\nCannot start server")
	// }

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/signup", signupHandler)
	mux.HandleFunc("/signout", signOutHandler)
	mux.HandleFunc("/resultSignup", signupResultHandler)
	mux.HandleFunc("/resultLogin", loginResultHandler)

	mux.HandleFunc("/db", dbHandler)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	fmt.Printf("Starting application on port %s", portNumber)
	server.ListenAndServe()

}

// PanicOnErr panics on error
func PanicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
