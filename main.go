package main

import (
	"database/sql" //for the database sql
	_ "encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux" //for the routers
	"html/template"          //to render the html files
	"log"
	"net/http"
	"path"
	"runtime"
)

/*
- A user view that displays a list of users (data via api /users), each user should have a
method of clicking to viewing all the details of that user (/user/:id)
- A widget view that displays a list of widgets (/widgets), each widget should
have a method of clicking to view the details of that widget (/widget/:id)
- A search/filter on the user and widget list views
- A method of creating a new widget (POST /widgets)
- A method of updating an existing widget (PUT /widgets/:id)
*/

//GET /users http://spa.tglrw.com:4000/users
//GET /users/:id http://spa.tglrw.com:4000/users/:id
//GET /widgets http://spa.tglrw.com:4000/widgets
//GET /widgets/:id http://spa.tglrw.com:4000/widgets/:id
//POST /widgets for creating new widgets http://spa.tglrw.com:4000/widgets
//PUT /widgets/:id for updating existing widgets http://spa.tglrw.com:4000/widgets/:id

type User struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	//Description string `json:"description,omitempty"`
}

type Widget struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Color     string `json:"color,omitempty"`
	Price     string `json:"price,omitempty"`
	Melts     string `json:"melts,omitempty"`
	Inventory int    `json:"inventory,omitempty"`
}

type Data struct {
	User        []User
	Widget      []Widget
	UserCount   int `json:"usercount,omitempty"`
	WidgetCount int `json:"widgetcount,omitempty"`
}

var user = "root"
var password = "secret"
var ip = "172.17.0.2:3306"
var dbname = "testdb"

func indexHandler(w http.ResponseWriter, r *http.Request) {

	//db requests
	db, err := sql.Open("mysql", user+":"+password+"@tcp("+ip+")/"+dbname)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer db.Close()
	// make sure connection is available
	err = db.Ping()
	if err != nil {
		fmt.Print(err.Error())
	}

	//query
	rows, err := db.Query("select id, name from person")
	var userCount int
	_ = db.QueryRow("select count(*) from person").Scan(&userCount)

	tRes := User{}     //table response
	var results []User //array of users

	//fetch result
	for rows.Next() {
		var id, name string
		rows.Scan(&id, &name)
		tRes.ID = id
		tRes.Name = name
		results = append(results, tRes)
	}

	//query
	rows2, err := db.Query("select id, name, color, price, melts, inventory from widgets")
	var widgetCount int
	_ = db.QueryRow("select count(*) from widgets").Scan(&widgetCount)

	tRes2 := Widget{}     //table response
	var results2 []Widget //array of widgets

	//fetch result
	for rows2.Next() {
		var id, name, color, price, melts string
		var inventory int
		rows2.Scan(&id, &name, &color, &price, &melts, &inventory)
		tRes2.ID = id
		tRes2.Name = name
		tRes2.Color = color
		tRes2.Price = price
		tRes2.Melts = melts
		tRes2.Inventory = inventory
		results2 = append(results2, tRes2)
	}

	var data Data
	data.User = results
	data.Widget = results2
	data.UserCount = userCount
	data.WidgetCount = widgetCount

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		fmt.Println("No caller information")
	}

	tmpl := template.Must(template.ParseFiles(path.Dir(filename) + "/templates/index.html"))
	if r.Method != http.MethodPost {
		// render template
		tmpl.Execute(w, data) //passing the array data into the HTML file
		return
	}

}

//ROUTES HANDLERS
func getUsers(w http.ResponseWriter, r *http.Request) {

	//db requests
	db, err := sql.Open("mysql", user+":"+password+"@tcp("+ip+")/"+dbname)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer db.Close()
	// make sure connection is available
	err = db.Ping()
	if err != nil {
		fmt.Print(err.Error())
	}

	//query
	rows, err := db.Query("select id, name from person")

	tRes := User{}     //table response
	var results []User //array of users

	//fetch result
	for rows.Next() {
		var id, name string
		rows.Scan(&id, &name)
		tRes.ID = id
		tRes.Name = name
		results = append(results, tRes)
	}

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		fmt.Println("No caller information")
	}

	// get template from views directory
	tmpl := template.Must(template.ParseFiles(path.Dir(filename) + "/templates/user.html"))
	if r.Method != http.MethodPost {
		// render template
		tmpl.Execute(w, results)
		return
	}

	//fmt.Fprintf(w, "You hit the /users page ")
}

func getUser(w http.ResponseWriter, r *http.Request) {

	db, err := sql.Open("mysql", user+":"+password+"@tcp("+ip+")/"+dbname)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer db.Close()
	// make sure connection is available
	err = db.Ping()
	if err != nil {
		fmt.Print(err.Error())
	}

	params := mux.Vars(r)
	var idRequestedPerson = params["id"]
	//log.Println("ID: " + idRequestedPerson)

	tRes := User{}

	row := db.QueryRow("select * from person where id = ?;", idRequestedPerson)
	err = row.Scan(&tRes.ID, &tRes.Name)
	log.Println(err)

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		fmt.Println("No caller information")
	}

	//log.Println("Filename: " + filename)

	tmpl := template.Must(template.ParseFiles(path.Dir(filename) + "/templates/userDetail.html"))
	tmpl.Execute(w, tRes)
	//fmt.Fprintf(w, "You hit the /users/:id page")
}

func getWidgets(w http.ResponseWriter, r *http.Request) {

	//db requests
	db, err := sql.Open("mysql", user+":"+password+"@tcp("+ip+")/"+dbname)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer db.Close()
	// make sure connection is available
	err = db.Ping()
	if err != nil {
		fmt.Print(err.Error())
	}

	//query
	rows2, err := db.Query("select id, name, color, price, melts, inventory from widgets")

	tRes2 := Widget{}     //table response
	var results2 []Widget //array of widgets

	//fetch result
	for rows2.Next() {
		var id, name, color, price, melts string
		var inventory int
		rows2.Scan(&id, &name, &color, &price, &melts, &inventory)
		tRes2.ID = id
		tRes2.Name = name
		tRes2.Color = color
		tRes2.Price = price
		tRes2.Melts = melts
		tRes2.Inventory = inventory
		results2 = append(results2, tRes2)
	}

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		fmt.Println("No caller information")
	}

	// get template from views directory
	tmpl := template.Must(template.ParseFiles(path.Dir(filename) + "/templates/widget.html"))
	if r.Method != http.MethodPost {
		// render template
		tmpl.Execute(w, results2)
		return
	}
}

func getWidget(w http.ResponseWriter, r *http.Request) {
	//db requests
	//log.Println("/widgets/:id page -> " + r.Method)
	db, err := sql.Open("mysql", user+":"+password+"@tcp("+ip+")/"+dbname)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer db.Close()
	// make sure connection is available
	err = db.Ping()
	if err != nil {
		fmt.Print(err.Error())
	}

	params := mux.Vars(r)
	var idRequestedWidget = params["id"]
	//log.Println("ID: " + idRequestedWidget)

	tRes := Widget{}

	row := db.QueryRow("select * from widgets where id = ?;", idRequestedWidget)
	err = row.Scan(&tRes.ID, &tRes.Name, &tRes.Color, &tRes.Price, &tRes.Melts, &tRes.Inventory)
	log.Println(err)

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		fmt.Println("No caller information")
	}

	//log.Println("Filename: " + filename)

	tmpl := template.Must(template.ParseFiles(path.Dir(filename) + "/templates/widgetDetail.html"))
	tmpl.Execute(w, tRes)

	//http.Redirect(w, r, "/widgets/", http.StatusMovedPermanently)

	//fmt.Fprintf(w, "You hit the /widgets/:id page")

}

func createWidget(w http.ResponseWriter, r *http.Request) {

	//db requests
	db, err := sql.Open("mysql", user+":"+password+"@tcp("+ip+")/"+dbname)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer db.Close()
	// make sure connection is available
	err = db.Ping()
	if err != nil {
		fmt.Print(err.Error())
	}

	name := r.FormValue("widget-name")
	color := r.FormValue("widget-color")
	price := r.FormValue("widget-price")
	melts := r.FormValue("widget-properties")
	inventory := r.FormValue("widget-count")

	if melts == "melts" {
		melts = "yes"
	} else {
		melts = "no"
	}

	// log.Println("CREATE: %s %s %s %s %d", name, color, price, melts, inventory)

	stmt, err := db.Prepare("insert into widgets (id, name, color, price, melts, inventory) values(?,?,?,?,?,?);")
	if err != nil {
		fmt.Print(err.Error())
	}
	_, err = stmt.Exec(nil, name, color, price, melts, inventory)

	if err != nil {
		fmt.Print(err.Error())
	}

	defer stmt.Close()

	// _, filename, _, ok := runtime.Caller(0)
	// if !ok {
	// 	fmt.Println("No caller information")
	// }

	http.Redirect(w, r, "/widgets", http.StatusMovedPermanently)

	// params := mux.Vars(r)
	// var person Person
	// _ = json.NewDecoder(r.Body).Decode(&person)
	// person.ID = params["id"]
	// people = append(people, person)
	// json.NewEncoder(w).Encode(people)

	//fmt.Fprintf(w, "You hit the /widget/ creation page")
}

func editWidget(w http.ResponseWriter, r *http.Request) {

	//log.Println("/widgets/:id page" + r.Method)
	//db requests
	db, err := sql.Open("mysql", user+":"+password+"@tcp("+ip+")/"+dbname)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer db.Close()
	// make sure connection is available
	err = db.Ping()
	if err != nil {
		fmt.Print(err.Error())
	}

	params := mux.Vars(r)
	var idRequestedWidget = params["id"]
	//log.Println("ID: " + idRequestedWidget)

	tRes := Widget{}

	row := db.QueryRow("select * from widgets where id = ?;", idRequestedWidget)
	err = row.Scan(&tRes.ID, &tRes.Name, &tRes.Color, &tRes.Price, &tRes.Melts, &tRes.Inventory)
	log.Println(err)

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		fmt.Println("No caller information")
	}

	//log.Println("Filename: " + filename)

	tmpl := template.Must(template.ParseFiles(path.Dir(filename) + "/templates/edit.html"))
	tmpl.Execute(w, tRes)

	//fmt.Fprintf(w, "You hit the /widget/:id/edit edit page")
}

func updateWidget(w http.ResponseWriter, r *http.Request) {

	//log.Println("/widgets/:id update page ---> " + r.Method)

	//db requests
	db, err := sql.Open("mysql", user+":"+password+"@tcp("+ip+")/"+dbname)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer db.Close()
	// make sure connection is available
	err = db.Ping()
	if err != nil {
		fmt.Print(err.Error())
	}

	params := mux.Vars(r)
	var widgetID = params["id"]

	name := r.FormValue("widget-name")
	color := r.FormValue("widget-color")
	price := r.FormValue("widget-price")
	melts := r.FormValue("widget-properties")
	inventory := r.FormValue("widget-count")

	if melts == "melts" {
		melts = "yes"
	} else {
		melts = "no"
	}

	stmt, err := db.Prepare("update widgets set name=?, color=?, price=?, melts=?, inventory=? where id=?;")
	if err != nil {
		fmt.Print(err.Error())
	}
	_, err = stmt.Exec(name, color, price, melts, inventory, widgetID)

	if err != nil {
		fmt.Print(err.Error())
	}

	defer stmt.Close()

	http.Redirect(w, r, "/widgets", http.StatusMovedPermanently)

	//fazer update aqui
	//fmt.Fprintf(w, "You hit the /widget/:id update page")
}

func main() {

	migrateDocker() //username, password, tcp ip are defined above in the code

	//making sure we can load css and images
	// cssHandler := http.FileServer(http.Dir("./templates/components/"))
	// imagesHandler := http.FileServer(http.Dir("./templates/img/"))
	// http.Handle("/templates/components/", http.StripPrefix("/templates/components/", cssHandler))
	// http.Handle("/templates/img/", http.StripPrefix("/templates/img/", imagesHandler))

	//route handlers
	router := mux.NewRouter()

	router.PathPrefix("/templates/components/").Handler(http.StripPrefix("/templates/components/", http.FileServer(http.Dir("./templates/components/"))))
	router.PathPrefix("/templates/img/").Handler(http.StripPrefix("/templates/img/", http.FileServer(http.Dir("./templates/img/"))))

	router.HandleFunc("/", indexHandler).Methods("GET")
	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/user.html", getUsers).Methods("GET")

	router.HandleFunc("/users/{id}", getUser).Methods("GET")
	router.HandleFunc("/user.html/{id}", getUser).Methods("GET")

	router.HandleFunc("/widgets", getWidgets).Methods("GET")
	router.HandleFunc("/widget.html", getWidgets).Methods("GET")

	router.HandleFunc("/widgets/{id}", getWidget).Methods("GET")
	router.HandleFunc("/widget.html/{id}", getWidget).Methods("GET")

	router.HandleFunc("/widgets", createWidget).Methods("POST")
	router.HandleFunc("/widgets/{id}", updateWidget).Methods("POST")
	// router.HandleFunc("/widgets/{id}", updateWidget).Methods("PUT")
	router.HandleFunc("/widgets/{id}/edit", editWidget).Methods("GET")

	//http.ListenAndServe(":8000", nil)
	log.Println("Server running on :8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
