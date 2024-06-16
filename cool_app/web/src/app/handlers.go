package app

import "net/http"

func (a *App) Index(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("<h1>Hello</h1><p>Welcome to the home page</p>"))
	//t := NewTemplates()
	a.templates.Render(w, "index.html", map[string]any{
		"title":  "Title from data",
		"header": "Header from data",
	})
}

func (a *App) About(w http.ResponseWriter, r *http.Request) {
	//t := NewTemplates()
	a.templates.Render(w, "about.html", nil)
}

type User struct {
	userID   int
	Username string
	Email    string
}

// create a slice (array) of users
// var users = []User{
// 	{userID: 1, Username: "john", Email: "john@example.com"},
// 	{userID: 2, Username: "jane", Email: "jane@example.com"},
// 	{userID: 3, Username: "mary", Email: "mary@example.com"},
// 	{userID: 4, Username: "peter", Email: "peter@example.com"},
// }

func (a *App) Users(w http.ResponseWriter, r *http.Request) {
	//t := NewTemplates()
	// a.templates.Render(w, "users.html", map[string]any{
	// 	//"username": users[0].Username,
	// 	//"email":    users[0].Email,
	// 	"users": users,
	// })
	res, err := a.db.Query("SELECT id, username, email FROM users")
	if err != nil {
		http.Error(w, err.Error(), 500) // http.StatusInternalServerError = 500
		return
	}

	var users []User
	// Iterate over the rows of sql query result
	// append each user to the users slice
	for res.Next() {
		var user User
		res.Scan(&user.userID, &user.Username, &user.Email)
		users = append(users, user)
	}
	a.templates.Render(w, "users.html", map[string]any{
		"users": users,
	})
}

// func (a *App) User(w http.ResponseWriter, r *http.Request) {
// 	//t := NewTemplates()

// 	// get the username from the URL
// 	//this works starting from Go 1.22
// 	username := r.PathValue("username")

// 	index := -1
// 	for i, user := range users {
// 		if user.Username == username {
// 			index = i
// 			break
// 		}
// 	}

// 	// if the user is not found
// 	if index < 0 {
// 		http.NotFound(w, r)
// 		return
// 	}

// 	a.templates.Render(w, "user.html", map[string]any{
// 		//"user": users[],
// 		"user": users[index],
// 	})
//}
