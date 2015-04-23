package main

import (
	"log"
	"strings"
	"net/http"
	"html/template"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Student struct  {
	Name string
	Major string
	Year int
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("form.html")
	if err != nil {
		http.Error(w, "Interal Server Error", 500)
		return
	}
	t.Execute(w, nil)
}

func queryHandler(w http.ResponseWriter, r *http.Request) {
	var (
		name string
		major string
		year string
		table string = ""
	)

	db, err := sql.Open("mysql", "root:W3iRd$_1tR#y@tcp(127.0.0.1:3306)/Test")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	query := "select * from test where"

	name = r.FormValue("name")
	major = r.FormValue("major")
	year = r.FormValue("year")

	if name == "any" && major == "any" && year == "0" {
		query = "select * from test"
	} else {
		if name != "any" {
			query += " name = '" + name + "'"
		}

		if major != "any" {
			if strings.Contains(query, "name") {
				query += " and major = '" + major + "'"
			} else {
				query += " major = '" + major + "'"
			}
		}

		if year != "0" {
			if strings.Contains(query, "name") || strings.Contains(query, "major") {
				query += " and year = " + year
			} else {
				query += " year = " + year
			}
		}
	}

	if strings.Contains(query, "drop") {
		table = ""
	} else {

		rows, err := db.Query(query)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		for rows.Next() {
			err = rows.Scan(&name, &major, &year)
			if err != nil {
				log.Fatal(err)
			}
			table += "<tr><td><a href=\"/student/" + name + "\">" + name + "</a></td><td>"
			table += major + "</td><td>" + year + "</td></tr>\n"
		}

		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}
	}

	t, _ := template.ParseFiles("results.html")
	t.Execute(w, template.HTML(table))
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path[len("/student/"):]

	s := loadStudent(name)

	t, _ := template.ParseFiles("pretty_test.html")
	t.Execute(w, s)
}

func loadStudent(name string) *Student {
	var (
		major string
		year int
	)

	db, err := sql.Open("mysql", "root:W3iRd$_1tR#y@tcp(127.0.0.1:3306)/Test")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	query := "select * from test where name = '" + name + "'"

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	rows.Next()
	err = rows.Scan(&name, &major, &year)
	if err != nil {
		log.Fatal(err)
	}

	return &Student{Name: name, Major: major, Year: year}
}

func main() {
	http.HandleFunc("/student/", dataHandler)
	http.HandleFunc("/query", queryHandler)
	http.HandleFunc("/", viewHandler)
	http.ListenAndServe(":8080", nil)
}
