package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: login")
	var u User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		panic(err)
	}

	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/securitytest")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	var userId int
	err = db.QueryRow("SELECT id FROM users WHERE username='" + u.Username + "' AND password='" + u.Password + "'").Scan(&userId)
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

type User struct {
	Username string
	Password string
}
