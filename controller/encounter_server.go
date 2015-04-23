package main

import (
	"log"
	"strings"
	"net/http"
	"html/template"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Monster struct  {
	Name string
	CR int
	Alignment string
	Size string
	Type string
	Init int
	Armor int
	Shield int
	Deflection int
	SizeAC int
	NaturalArmor int
	Dodge int
	MiscAC int
	HitDie int
	Fort int
	Reflex int
	Will int
	BaseSpeed int
	Space int
	Reach int
	SpellAbilities string
	Spell string
	Str int
	Dex int
	Con int
	Wis int
	Int int
	Cha int
	BAB int
	CMB int
	CMD int
	Feats string
	Lang string
	SpecAtt string
	Environment string
	Attack_1 string
	Attack_2 string
	Attack_3 string
	Attack_4 string
	Attack_5 string
	Defense string
	Offense string
	Stats string
	Special string
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("pathfinder_form.html")
	if err != nil {
		http.Error(w, "Interal Server Error", 500)
		return
	}
	t.Execute(w, nil)
}

func queryHandler(w http.ResponseWriter, r *http.Request) {
	var (
		name string
		alignment string
		cr string
		environment string
		type_ string = ""
	)

	db, err := sql.Open("mysql", "root:W3iRd$_1tR#y@tcp(127.0.0.1:3306)/PathfinderEncounter")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	query := "select Name, CR, Alignment, TypeName, Environment from PathfinderEncounter"

	alignment = r.FormValue("alignment")
	cr = r.FormValue("cr")
	environment = r.FormValue("environment")
	type_ = r.FormValue("type")

	if alignment == "any" && environment == "any" && cr == "any" && type_ == "any" {
		query = "select * from PathfinderEncounter"
	} else {
		if alignment != "any" {
			query += " where Alignment = '" + alignment + "'"
		}

		if  cr != "any" {
			if strings.Contains(query, "where") {
				query += " and CR = " + cr
			} else {
				query += " where CR = " + cr
			}
		}

		if environment != "any" {
			if strings.Contains(query, "where") {
				query += " and Environment = " + environment + "'"
			} else {
				query += " where Environment = " + environment + "'"
			}
		}

		if type_ != "any" {
			if strings.Contains(query, "where") {
				query += " and TypeName = '" + type_ + "'"
			} else {
				query += " where TypeName = '" + type_ + "'"
			}
		}
	}

	if strings.Contains(query, "drop") || strings.Contains(query, "insert") {
		table = ""
	} else {

		rows, err := db.Query(query)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		for rows.Next() {
			err = rows.Scan(&name, &cr, &alignment, &type_, &environment)
			if err != nil {
				log.Fatal(err)
			}
			table += "<tr><td><a href=\"/monster/" + name + "\">" + name + "</a></td><td>"
			table += cr + "</td><td>" + alignment + "</td><td>" + environment + "</td></tr>\n"
		}

		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}
	}

	t, _ := template.ParseFiles("monster_results.html")
	t.Execute(w, template.HTML(table))
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path[len("/monster/"):]

	s := loadMonster(name)

	t, _ := template.ParseFiles("pretty_result.html")
	t.Execute(w, s)
}

func loadMonster(name string) *Monster {
	var (
		major string
		year int
	)

	db, err := sql.Open("mysql", "root:W3iRd$_1tR#y@tcp(127.0.0.1:3306)/PathfinderEncounter")
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

	return &Monster{Name: name, Major: major, Year: year}
}

func main() {
	http.HandleFunc("/monster/", dataHandler)
	http.HandleFunc("/query", queryHandler)
	http.HandleFunc("/", viewHandler)
	http.ListenAndServe(":8080", nil)
}
