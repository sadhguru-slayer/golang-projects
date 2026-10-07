package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from Go API!")
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Users")
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error:", http.StatusBadRequest)
		return
	}
	fmt.Println("Name", user.Name)
	fmt.Println("Email:", user.Email)

	fmt.Fprintln(w, "User saved successfully")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", helloHandler)
	mux.HandleFunc("GET /users", getUsers)
	mux.HandleFunc("POST /add-user", createUser)
	fmt.Println("Server running on http://localhost:8000")
	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		fmt.Println(err)
	}
}
