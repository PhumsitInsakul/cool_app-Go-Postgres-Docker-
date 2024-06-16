package main

import (
	"log"

	"github.com/pharick/cool_app/app"
)

// func index(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("<h1>Hello</h1><p>Welcome to the home page</p>"))
// }

func main() {
	a, err := app.NewApp()
	if err != nil {
		log.Fatal(err)
	}
	// create a router
	//r := http.NewServeMux()

	// register the handler
	a.RegisterHandler("/{$}", a.Index)
	a.RegisterHandler("/about", a.About)
	a.RegisterHandler("/users", a.Users)

	a.RegisterHandler("/users/{username}", a.Users)

	// start the server
	//log.Printf("Server started on port 3000\n")
	//http.ListenAndServe(":3000", app.router)
	a.Serve(3000)
}
