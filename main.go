package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]
	password := vars["password"]
	fmt.Println("Endpoint Hit: login")

	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/securitytest")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	var userId int
	err = db.QueryRow("SELECT id FROM users WHERE username='" + username + "' AND password='" + password + "'").Scan(&userId)
	//err = db.QueryRow("SELECT id FROM users WHERE username=? AND password=?", username, password).Scan(&userId)
	if err != nil {
		w.WriteHeader(400)
	}
	w.WriteHeader(200)
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/login", loginHandler).Methods("POST")
	router.HandleFunc("/", homePage).Methods("GET")
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8888", nil))
}

func main() {
	handleRequests()
}
