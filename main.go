package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

type User struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	ScreenName string `json:"screen_name"`
	Email string `json:"email"`
	EmailVerifiedAt string `json:"email_verified_at"`
	Password string `json:"password"`
	RememberToken string `json:"remember_token"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}


func allArticles(w http.ResponseWriter, r *http.Request) {

	articles := Articles{
		Article{Title: "Test Title", Desc: "Test Description", Content: "Hello World"},
	}
	fmt.Println("Endpoint Hit: All Article Endpoint")
	json.NewEncoder(w).Encode(articles)
}

func testPostArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test Post Endpoint Work")
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "homepage Endpoint")
}

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/articles", testPostArticles).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	fmt.Println("Go MySQL Tutorial")

	db, err := sql.Open("mysql", "twiclo:twicloPass@tcp(127.0.0.1:23306)/twiclo")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Successfully connected to MySQL database")

	result, err := db.Query("SELECT id, name, screen_name, email, password, created_at, updated_at FROM users;")
	if err != nil {
		panic(err.Error())
	}

	var user User
	for result.Next() {

		err = result.Scan(
			&user.Id,
			&user.Name,
			&user.ScreenName,
			&user.Email,
			&user.Password,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
	}
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("successfully select %s", user)

	result.Close()

	handleRequests()
}
