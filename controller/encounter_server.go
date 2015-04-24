package main

import (
//	"fmt"
	"log"
	"strings"
	"strconv"
	"net/http"
	"html/template"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type MonsterData struct  {
	Name string
	Basic template.HTML
	Defense template.HTML
	Offense template.HTML
	Stats template.HTML
	Special template.HTML
}

func newMonsterHandler(w http.ResponseWriter, r *http.Request) {

	var (
		max_id int
		sizeVal int
		sizeAC int
	)

	name := r.FormValue("name")
	cr := r.FormValue("cr")
	alignment := r.FormValue("alignment")
	size := r.FormValue("size")
	class := r.FormValue("class")
	type_ := r.FormValue("type_")
	armor := r.FormValue("armor")
	shield := r.FormValue("shield")
	deflection := r.FormValue("deflection")
	naturalArmor := r.FormValue("naturalArmor")
	dodge := r.FormValue("dodge")
	miscAC := r.FormValue("miscAC")
	hitdie := r.FormValue("hitdie")
	fort := r.FormValue("fort")
	reflex := r.FormValue("reflex")
	will := r.FormValue("will")
	speed := r.FormValue("speed")
	space := r.FormValue("space")
	reach := r.FormValue("reach")
	spellabl := r.FormValue("spellabl")
	spell := r.FormValue("spell")
	str := r.FormValue("str")
	dex := r.FormValue("dex")
	con := r.FormValue("con")
	wis := r.FormValue("wis")
	int_ := r.FormValue("int_")
	cha	 := r.FormValue("cha")
	feat := r.FormValue("feat")
	skill := r.FormValue("skill")
	lang := r.FormValue("lang")
	specatt := r.FormValue("specatt")
	environment := r.FormValue("environment")
	bab := r.FormValue("bab")
	attack_1 := r.FormValue("att_1")
	attack_2 := r.FormValue("att_2")
	attack_3 := r.FormValue("att_3")
	attack_4 := r.FormValue("att_4")
	attack_5 := r.FormValue("att_5")
	book := r.FormValue("book")

	switch {
		case size == "F":
			sizeAC = 8
			sizeVal = -4
		case size == "D":
			sizeAC = 4
			sizeVal = -3
		case size == "T":
			sizeAC = 2
			sizeVal = -2
		case size == "S":
			sizeAC = 1
			sizeVal = -1
		case size == "M":
			sizeAC = 0
			sizeVal = 0
		case size == "L":
			sizeAC = -1
			sizeVal = 1
		case size == "H":
			sizeAC = -2
			sizeVal = 2
		case size == "C":
			sizeAC = -4
			sizeVal = 3
		case size == "G":
			sizeAC = -8
			sizeVal = 4
		default:
			sizeAC = 0
			sizeVal = 0
	}

	bab_n, _ := strconv.Atoi(bab)
	str_n, _ := strconv.Atoi(str)
	dex_n, _ := strconv.Atoi(dex)

	cmb := bab_n + str_n
	cmd := bab_n + str_n + dex_n + 10
	init := (dex_n - 10)/2

	db, err := sql.Open("mysql", "root:W3iRd$_1tR#y@tcp(127.0.0.1:3306)/PathfinderEncounter")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	query := "select max(idMonster) from Monster"
	err = db.QueryRow(query).Scan(&max_id)
	if err != nil {
		log.Fatal(err)
	}

	max_id++

	book_insert := "INSERT INTO `PathfinderEncounter`.`Book` (`idBook`, `MonsterName`, `BookName`, `ThirdParty`) VALUES (?,?,?, 0)"

	stmt, err := db.Prepare(book_insert)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(max_id, name, book)
	if err != nil {
		log.Fatal(err)
	}

	insert := "INSERT INTO `PathfinderEncounter`.`Monster` (`idMonster`,`Name`,`CR`,`Alignment`,`Size`,`Class`,`TypeName`,`Initiative`,`Armor`,`Shield`,`Deflection`,`SizeAC`,`NaturalArmor`,`Dodge`,`MiscAC`,`HitDie`,`Fort`,`Reflex`,`Will`,`BaseSpeed`,`Space`,`Reach`,`Spell-Like Abilities`,`Spells`,`Str`,`Dex`,`Con`,`Inte`,`Wis`,`Cha`,`BaseAttack`,`CMB`,`CMD`,`Feats`,`Skills`,`Languages`,`Special Attacks`,`Environment`,`Attack1`,`Attack2`,`Attack3`,`Attack4`,`Attack5`,`Book_idBook`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"

	stmt, err = db.Prepare(insert)
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(max_id,
						name,
						cr,
						alignment,
						sizeVal,
						class,
						type_,
						init,
						armor,
						shield,
						deflection,
						sizeAC,
						naturalArmor,
						dodge,
						miscAC,
						hitdie,
						fort,
						reflex,
						will,
						speed,
						space,
						reach,
						spellabl,
						spell,
						str,
						dex,
						con,
						int_,
						wis,
						cha,
						bab,
						cmb,
						cmd,
						feat,
						skill,
						lang,
						specatt,
						environment,
						attack_1,
						attack_2,
						attack_3,
						attack_4,
						attack_5,
						max_id)
	if err != nil {
		log.Fatal(err)
	}

	t, err := template.ParseFiles("confirm.html")
	if err != nil {
		http.Error(w, "Internal Sever Error", 500)
		return
	}
	t.Execute(w, template.HTML(name))
}

func landingHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("landing.html")
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
	t.Execute(w, nil)
}

func iformHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("insert_form.html")
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
	t.Execute(w, nil)
}

func pformHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("pathfinder_form.html")
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
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
		table string
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

	query := "select Name, CR, Alignment, TypeName, Environment from Monster"

	alignment = r.FormValue("alignment")
	cr = r.FormValue("cr")
	environment = r.FormValue("environment")
	type_ = r.FormValue("type")

	if alignment == "any" && environment == "any" && cr == "any" && type_ == "any" {
		query = "select Name, CR, Alignment, TypeName, Environment from Monster"
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
				query += " and Environment = '" + environment + "'"
			} else {
				query += " where Environment = '" + environment + "'"
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
			table += cr + "</td><td>" + alignment + "</td><td>" + type_ + "</td><td>" + environment + "</td></tr>\n"
		}

		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}
	}

	t, err := template.ParseFiles("monster_results.html")
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
	t.Execute(w, template.HTML(table))
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path[len("/monster/"):]

	s := loadMonsterData(name)

	t, err := template.ParseFiles("pretty_results.html")
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
	t.Execute(w, s)
}

func loadMonsterData(m_name string) *MonsterData {
	var(
		id []byte
		name []byte
		cr []byte
		alignment []byte
		size []byte
		class []byte
		type_ []byte
		init []byte
		armor []byte
		shield []byte
		deflection []byte
		sizeAC []byte
		naturalArmor []byte
		dodge []byte
		miscAC []byte
		hitDie []byte
		fort []byte
		reflex []byte
		will []byte
		baseSpeed []byte
		space []byte
		reach []byte
		spellAbilities []byte
		spell []byte
		str []byte
		dex []byte
		con []byte
		wis []byte
		int_ []byte
		cha	[]byte
		bAB []byte
		cMB []byte
		cMD []byte
		feats []byte
		skills []byte
		lang []byte
		specAtt []byte
		environment []byte
		attack_1 []byte
		attack_2 []byte
		attack_3 []byte
		attack_4 []byte
		attack_5 []byte
		book []byte
		basic string = ""
		defense string = ""
		offense string = ""
		stats string = ""
		special string = ""
		hd_size int
		m_size string
		att_1 []byte
		att_2 []byte
		att_3 []byte
		att_4 []byte
		att_5 []byte
		book_name []byte
	)

	//connect to database
	db, err := sql.Open("mysql", "root:W3iRd$_1tR#y@tcp(127.0.0.1:3306)/PathfinderEncounter")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	//query for monster base information
	query := "select * from Monster where name = '" + m_name + "'"

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	//grab data from query
	rows.Next()
	err = rows.Scan(&id,
					&name,
					&cr,
					&alignment,
					&size,
					&class,
					&type_,
					&init,
					&armor,
					&shield,
					&deflection,
					&sizeAC,
					&naturalArmor,
					&dodge,
					&miscAC,
					&hitDie,
					&fort,
					&reflex,
					&will,
					&baseSpeed,
					&space,
					&reach,
					&spellAbilities,
					&spell,
					&str,
					&dex,
					&con,
					&int_,
					&wis,
					&cha,
					&bAB,
					&cMB,
					&cMD,
					&feats,
					&skills,
					&lang,
					&specAtt,
					&environment,
					&attack_1,
					&attack_2,
					&attack_3,
					&attack_4,
					&attack_5,
					&book)

	if err != nil {
		log.Fatal(err)
	}

	//convert size from int to letter size
	switch {
		case string(size) == "-4":
			m_size = "F"
		case string(size) == "-3":
			m_size = "D"
		case string(size) == "-2":
			m_size = "T"
		case string(size) == "-1":
			m_size = "S"
		case string(size) == "0":
			m_size = "M"
		case string(size) == "1":
			m_size = "L"
		case string(size) == "2":
			m_size = "H"
		case string(size) == "3":
			m_size = "C"
		case string(size) == "4":
			m_size = "G"
		default:
			m_size = "M"
	}

	//format basic information for pretty_results
	basic += "<b>CR:</b> " + string(cr) + "<br>"
	basic += string(name) + " " + string(type_) + "<br>"
	basic += string(alignment) + " " + m_size + " " + string(type_) + "<br>"
	basic += "<b>Initiative:</b> " + string(init) + "<br>"

	//calculate values for Armor class
	dex_base, err := strconv.Atoi(string(dex))
	if err != nil {
		dex_base = 0
	}

	dex_mod := (dex_base - 10) / 2
	armor_mod, err := strconv.Atoi(string(armor))
	if err != nil {
		armor_mod = 0
	}

	shiel_mod, err := strconv.Atoi(string(shield))
	if err != nil {
		shiel_mod = 0
	}

	defle_mod, err := strconv.Atoi(string(deflection))
	if err != nil {
		defle_mod = 0
	}

	size_mod, err := strconv.Atoi(string(sizeAC))
	if err != nil {
		size_mod = 0
	}

	nata_mod, err := strconv.Atoi(string(naturalArmor))
	if err != nil {
		nata_mod = 0
	}

	dodge_mod, err := strconv.Atoi(string(dodge))
	if err != nil {
		dodge_mod = 0
	}

	misc_mod, err := strconv.Atoi(string(miscAC))
	if err != nil {
		misc_mod = 0
	}

	ac := armor_mod + shiel_mod + defle_mod + size_mod + nata_mod + dodge_mod + misc_mod + dex_mod + 10
	touch := ac - armor_mod - shiel_mod
	flatfoot := ac - dex_mod

	//compile and format information for monster defense
	defense += "<b>AC:</b> " + strconv.Itoa(ac) + " <b>Touch:</b> " + strconv.Itoa(touch) + " <b>Flat-footed:</b> " + strconv.Itoa(flatfoot) + "<br>"

	query = "select HitDie from Type where TypeName = '" + string(type_) + "'"
	rows, err = db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	rows.Next()
	err = rows.Scan(&hd_size)
	if err != nil {
		log.Fatal(err)
	}

	defense += "<b>HP:</b> " + string(hitDie) + "d" + strconv.Itoa(hd_size) + "<br>"
	defense += "<b>Fort:</b> " + string(fort) + " <b>Reflex:</b> " + string(reflex) + " <b>Will:</b> " + string(will) + "<br>"

	//compile and format information for monster attack
	offense += "<b>Speed:</b> " + string(baseSpeed) + "<br>"
	offense += "<b>Attacks:</b><br>"

	if string(attack_1) != "" {
		query = "select " + m_size + " from Attacks where AttackName = '" + string(attack_1) + "'"
		rows, err = db.Query(query)
		if err != nil {
			log.Fatal(err)
		}

		rows.Next()
		err = rows.Scan(&att_1)
		if err != nil {
			log.Fatal(err)
		}

		if string(att_1) != "" {
			offense += string(attack_1) + " " + string(att_1) + "<br>"
		}
	}

	if string(attack_2) != "" {
		query = "select " + m_size + " from Attacks where AttackName = '" + string(attack_2) + "'"
		rows, err = db.Query(query)
		if err != nil {
			log.Fatal(err)
		}

		rows.Next()
		err = rows.Scan(&att_2)
		if err != nil {
			log.Fatal(err)
		}

		if string(att_2) != "" {
			offense += string(attack_2) + " " + string(att_2) + "<br>"
		}
	}

	if string(attack_3) != "" {

		query = "select " + m_size + " from Attacks where AttackName = '" + string(attack_3) + "'"
		rows, err = db.Query(query)
		if err != nil {
			log.Fatal(err)
		}

		rows.Next()
		err = rows.Scan(&att_3)
		if err != nil {
			log.Fatal(err)
		}

		if string(att_3) != "" {
			offense += string(attack_3) + " " + string(att_3) + "<br>"
		}
	}

	if string(attack_4) != "" {
		query = "select " + m_size + " from Attacks where AttackName = '" + string(attack_4) + "'"
		rows, err = db.Query(query)
		if err != nil {
			log.Fatal(err)
		}

		rows.Next()
		err = rows.Scan(&att_4)
		if err != nil {
			log.Fatal(err)
		}

		if string(att_4) != "" {
			offense += string(attack_4) + " " + string(att_4) + "<br>"
		}
	}

	if string(attack_5) != "" {
		query = "select " + m_size + " from Attacks where AttackName = '" + string(attack_5) + "'"
		rows, err = db.Query(query)
		if err != nil {
			log.Fatal(err)
		}

		rows.Next()
		err = rows.Scan(&att_5)
		if err != nil {
			log.Fatal(err)
		}

		if string(att_5) != "" {
			offense += string(attack_5) + " " + string(att_5) + "<br>"
		}
	}

	offense += "<b>Special Attacks</b>: " + string(specAtt) + "<br>"
	offense += "<b>Spell-like Abilities:</b> " + string(spellAbilities) + "<br>"
	offense += "<b>Spells:</b> " + string(spell) + "<br>"

	//compile and format monster statistics
	stats += "<b>Str</b> " + string(str) + " "
	stats += "<b>Dex</b> " + string(dex) + " "
	stats += "<b>Con</b> " + string(con) + " "
	stats += "<b>Int</b> " + string(int_) + " "
	stats += "<b>Wis</b> " + string(wis) + " "
	stats += "<b>Cha</b> " + string(cha) + "<br>"
	stats += "<b>BAB</b> " + string(bAB) + " "
	stats += "<b>CMB</b> " + string(cMB) + " "
	stats += "<b>CMD</b> " + string(cMD) + "<br>"
	stats += "<b>Feats:</b> " + string(feats) + "<br>"
	stats += "<b>Skills:</b> " + string(skills) + "<br>"
	stats += "<b>Languages:</b> " + string(lang) + "<br>"

	query = "select BookName from Book where MonsterName = '" + string(name) + "'"
	rows, err = db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	rows.Next()
	err = rows.Scan(&book_name)
	if err != nil {
		log.Fatal(err)
	}

	special += "Information provided by Pathfinder Role Playing Game " + string(book_name) + "<br>"
	special += "Available under the OGL (Open Gaming License)"

	return &MonsterData{Name: string(name),
					Basic: template.HTML(basic),
					Defense: template.HTML(defense),
					Offense: template.HTML(offense),
					Stats: template.HTML(stats),
					Special: template.HTML(special)}
}

func main() {
	http.HandleFunc("/confirm", newMonsterHandler)
	http.HandleFunc("/monster/", dataHandler)
	http.HandleFunc("/query", queryHandler)
	http.HandleFunc("/find", pformHandler)
	http.HandleFunc("/insert", iformHandler)
	http.HandleFunc("/", landingHandler)
	http.ListenAndServe(":8080", nil)
}
