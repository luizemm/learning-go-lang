package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

type Post struct {
	Id int
	Title string
	Body string
}

var db, _ = sql.Open("sqlite3", "file:../tmp/database/web.db")

func main() {
	r := mux.NewRouter()

	r.PathPrefix("/static").
		Handler(
			http.StripPrefix(
				"/static",
				http.FileServer(http.Dir("static/")),
			),
		)
	
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/{id}/view", ViewHandler)

	fmt.Println(http.ListenAndServe(":8080", r))
}

func ListPosts() []Post {
	rows, err := db.Query("select * from posts")
	checkErr(err)

	items := []Post{}

	for rows.Next() {
		post := Post{}

		rows.Scan(&post.Id, &post.Title, &post.Body)

		items = append(items, post)
	}

	return items
}

func GetPostById(id string) Post {
	row := db.QueryRow("select * from posts where id = ?", id)

	post := Post{}

	row.Scan(&post.Id, &post.Title, &post.Body)

	return post
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	items := ListPosts()

	t := template.Must(template.ParseFiles("templates/layout.html", "templates/list.html"))
		
	if err := t.ExecuteTemplate(w, "layout.html", items); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	t := template.Must(template.ParseFiles("templates/layout.html", "templates/view.html"))
		
	t.ExecuteTemplate(w, "layout.html", GetPostById(id))
}

func checkErr(err error) {
	if err != nil{
		panic(err)
	}
}