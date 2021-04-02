package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/google/uuid"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func RootHandler(w http.ResponseWriter, r *http.Request) {
	data := UserPageData{
		PageHeader: PageHeader{
			Title:       "Evil Corp",
			Description: "Welcome to Evil Corp's Internal Portal!",
		},
	}
	err := templates.ExecuteTemplate(w, "index", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	searchQuery := r.PostFormValue("search")
	if searchQuery != "" {
		data := UserPageData{
			PageHeader: PageHeader{
				Title:       "Evil Corp Internal Directory",
				Description: "Search for Evil Corp employees and contractors.",
			},
			Users: getUsers(searchQuery),
		}
		err := templates.ExecuteTemplate(w, "users", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		data := UserPageData{
			PageHeader: PageHeader{
				Title:       "Evil Corp Internal Directory",
				Description: "Search for Evil Corp employees and contractors.",
			},
			Users: nil,
		}
		err := templates.ExecuteTemplate(w, "users", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func SubscribeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		data := SubscriberPageData{
			PageHeader: PageHeader{
				Title:       "Evil Corp Newsletter",
				Description: "Subscribe to recieve updates about internal company events.",
			},
			Subscribed: false,
			Name:       "",
			Email:      "",
		}
		err := templates.ExecuteTemplate(w, "subscribe", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else if r.Method == "POST" {
		name := r.PostFormValue("name")
		email := r.PostFormValue("email")
		data := SubscriberPageData{
			PageHeader: PageHeader{
				Title:       "Evil Corp Newsletter",
				Description: "Subscribe to recieve updates about internal company events.",
			},
			Subscribed: true,
			Name:       name,
			Email:      email,
		}
		err := templates.ExecuteTemplate(w, "subscribe", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		fmt.Println("Unsupported Method")
	}
}

func WikiHandler(w http.ResponseWriter, r *http.Request) {

	wikiPage := WikiPage{
		PageHeader: PageHeader{
			Title:       "Evil Corp Internal Wiki",
			Description: "View and create content to be shared with fellow Evil Corp employees.",
		},
	}

	// // view specific id
	// id := r.URL.Query().Get("id")
	// if id == "" {
	// 	return
	// } else {
	// 	id = r.FormValue("id")
	// 	data, err := ioutil.ReadFile(("data/wiki-topics/" + id + ".json"))
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}

	// 	var article WikiArticle
	// 	err = json.Unmarshal(data, &article)

	// 	wikiPage.SelectedArticle = article
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}

	// 	err = templates.ExecuteTemplate(w, "wiki", wikiPage)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// }

	switch method := r.Method; method {
	case "HEAD":
		fmt.Println(method)
	case "GET":
		items, err := ioutil.ReadDir("data/wiki-topics/")
		if err != nil {
			fmt.Println(err)
		}
		for _, item := range items {
			filename := item.Name()
			var p WikiPage
			if filepath.Ext(filename) == ".json" {
				data, err := ioutil.ReadFile((filename))
				if err != nil {
					json.Unmarshal(data, &p)
				}
			}
		}
		err = templates.ExecuteTemplate(w, "wiki", wikiPage)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case "POST":
		fmt.Println(method)
		uuid := uuid.New().String()
		subject := r.FormValue("subject")
		content := r.FormValue("content")

		if subject == "" || content == "" {
			return
		}
		article := WikiArticle{
			UUID: uuid, Subject: subject, Content: content,
		}

		data, err := json.Marshal(&article)
		if err != nil {
			fmt.Println(err)
		}
		filename := "data/wiki-topics/" + uuid + ".json"
		ioutil.WriteFile(filename, []byte(data), 0600)

		http.Redirect(w, r, "/wiki?id="+uuid, http.StatusSeeOther)
	case "PUT":
		fmt.Println(method)
	case "DELETE":
		fmt.Println(method)
	case "OPTIONS":
		fmt.Println(method)
	default:
		fmt.Println(method)
	}

}
