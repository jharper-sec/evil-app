package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"text/template"

	_ "github.com/mattn/go-sqlite3"

	"github.com/labstack/echo/v4"
)

type UserPageData struct {
	PageTitle string
	Users     []User
}

type SubscriberPageData struct {
	Subscribed bool
	Name       string
	Email      string
}

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Company   string `json:"company"`
	Title     string `json:"title"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	DOB       string `json:"dob"`
	SSN       string `json:"ssn"`
	Salary    int    `json:"salary"`
	Admin     string `json:"admin"`
}

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func seedUserData(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)

	var users []User

	json.Unmarshal(byteValue, &users)

	db, err := sql.Open("sqlite3", "data/users.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	drop_table_query := "DROP TABLE IF EXISTS user;"
	_, err = db.Exec(drop_table_query)
	if err != nil {
		panic(err)
	}

	create_table_query := "CREATE TABLE IF NOT EXISTS user(id INTEGER PRIMARY KEY, first_name TEXT, last_name TEXT, company TEXT, title TEXT, email TEXT, phone TEXT, dob TEXT, ssn TEXT, salary NUMERIC, admin BOOLEAN);"
	_, err = db.Exec(create_table_query)
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(users); i++ {
		user := users[i]
		insert_user_query := fmt.Sprintf("INSERT INTO user VALUES (%d, '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%d', '%s');", i+1, user.FirstName, user.LastName, user.Company, user.Title, user.Email, user.Phone, user.DOB, user.SSN, user.Salary, user.Admin)
		_, err = db.Exec(insert_user_query)
		if err != nil {
			panic(err)
		}
	}
}

func getUsers(search string) []User {
	db, err := sql.Open("sqlite3", "data/users.db")
	if err != nil {
		panic(err)
	}
	rows, err := db.Query("SELECT * FROM user WHERE (first_name LIKE '%" + search + "%' OR last_name LIKE '%" + search + "%') AND admin == 'false';")
	if err != nil {
		panic(err)
	}
	var results []User

	var id int
	var first_name string
	var last_name string
	var company string
	var title string
	var email string
	var phone string
	var dob string
	var ssn string
	var salary int
	var admin string
	for rows.Next() {
		_ = rows.Scan(&id, &first_name, &last_name, &company, &title, &email, &phone, &dob, &ssn, &salary, &admin)
		user := User{
			FirstName: first_name,
			LastName:  last_name,
			Company:   company,
			Title:     title,
			Email:     email,
			Phone:     phone,
			DOB:       dob,
			SSN:       ssn,
			Salary:    salary,
			Admin:     admin,
		}
		results = append(results, user)
	}
	rows.Close()
	return results
}

func vulnerable_echo_lib() {
	_ = echo.New()
}

func main() {
	seedUserData("data/user_seed_data.json")
	vulnerable_echo_lib()

	index_template := template.Must(template.ParseFiles("views/index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		index_template.Execute(w, nil)
	})

	users_template := template.Must(template.ParseFiles("views/users.html"))

	http.HandleFunc("/users", func(w http.ResponseWriter, req *http.Request) {
		searchQuery := req.PostFormValue("search")
		if searchQuery != "" {
			data := UserPageData{
				PageTitle: "Users",
				Users:     getUsers(searchQuery),
			}
			users_template.Execute(w, data)
		} else {
			users_template.Execute(w, nil)
		}
	})

	contact_template := template.Must(template.ParseFiles("views/subscribe.html"))
	http.HandleFunc("/subscribe", func(w http.ResponseWriter, req *http.Request) {
		if req.Method == "GET" {
			data := SubscriberPageData{
				Subscribed: false,
				Name:       "",
				Email:      "",
			}
			contact_template.Execute(w, data)
		} else if req.Method == "POST" {
			name := req.PostFormValue("name")
			email := req.PostFormValue("email")
			data := SubscriberPageData{
				Subscribed: true,
				Name:       name,
				Email:      email,
			}
			contact_template.Execute(w, data)
		} else {
			fmt.Println("Unsupported Method")
		}
	})

	help_template := template.Must(template.ParseFiles("views/wiki.html"))

	http.HandleFunc("/wiki", func(w http.ResponseWriter, req *http.Request) {
		help_template.Execute(w, nil)
	})

	http.HandleFunc("/wiki/view/", func(w http.ResponseWriter, req *http.Request) {
	})
	http.HandleFunc("/wiki/edit/", func(w http.ResponseWriter, req *http.Request) {
	})
	http.HandleFunc("/wiki/save/", func(w http.ResponseWriter, req *http.Request) {
	})
	http.HandleFunc("/wiki/delete/", func(w http.ResponseWriter, req *http.Request) {
	})

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	http.ListenAndServe(":8080", nil)
}
