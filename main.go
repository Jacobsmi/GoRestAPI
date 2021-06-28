package main

import (
	"encoding/json"
	"fmt"
	"goRestAPI/database"
	"log"
	"net/http"
)

type User struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var newUser User
	json.NewDecoder(r.Body).Decode(&newUser)
	fmt.Fprintln(w, newUser)
	database.CreateUser(newUser.Name, newUser.Username, newUser.Password)
}

func handleRequests() {
	fmt.Println("Running the API at http://localhost:5000")

	http.HandleFunc("/", homePage)
	http.HandleFunc("/createuser", createUser)

	log.Fatal(http.ListenAndServe(":5000", nil))
}

func main() {
	handleRequests()
}
