package main

import (
	"strings"
	"net/http"
	"html/template"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("test_form.html")
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
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
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

	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&name, &major, &year)
		if err != nil {
			panic(err)
		}
		table += "<tr><td>" + name + "</td><td>" + major + "</td><td>" + year + "</td></tr>\n"
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	t, _ := template.ParseFiles("results.html")
	t.Execute(w, template.HTML(table))
}

func main() {
	http.HandleFunc("/query", queryHandler)
	http.HandleFunc("/", viewHandler)
	http.ListenAndServe(":8080", nil)
}
