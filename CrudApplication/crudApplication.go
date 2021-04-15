package main


/*
built an application on crud operation with using mysql database.

*/
import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql" // import database/sql
)

//create Employee struct that has field-id,name,city
type Employee struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	City string `json:"city"`
}

//create Function dbConn opens connection with MySql driver
func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "admin"
	dbName := "goblog"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}
//Create Function Insert to save employee information to db
func Insert(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		name := r.FormValue("name")
		city := r.FormValue("city")
		insForm, _ := db.Prepare("INSERT INTO Employee(name,city) Values(?,?)")
		insForm.Exec(name, city)
		log.Println("INSERT: Name: " + name + " | City: " + city)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}
//Create Index Function to fetch all employee from db into table form
func Index(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	selDB, err := db.Query("Select * from Employee")
	p, _ := template.ParseFiles("C:\\Users\\Gaurav\\githubProject\\read-api\\CrudApplication\\src\\html\\template\\show.html")
	if err != nil {
		panic(err.Error())
	}
	emp := Employee{}
	res := []Employee{}
	for selDB.Next() {
		var id int
		var name, city string
		err := selDB.Scan(&id, &name, &city)
		if err != nil {
			panic(err.Error())
		}
		emp.Id = id
		emp.Name = name
		emp.City = city
		res = append(res, emp)
	}
	p.Execute(w, res)

}

//function New Execute at home page
func New(w http.ResponseWriter, r *http.Request) {
	News := "RegistrationPage"
	p, err := template.ParseFiles("C:\\Users\\Gaurav\\githubProject\\read-api\\CrudApplication\\src\\html\\template\\neww.html")
	if err != nil {
		log.Fatal(err)
	}
	p.Execute(w, News)
}
//function delete the employee acc to id number
func Delete(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	emp := r.URL.Query().Get("id")
	delForm, err := db.Prepare("DELETE FROM Employee WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(emp)
	log.Println("DELETE")
	defer db.Close()
	http.Redirect(w, r, "/index", 301) //StatusMovedPermanently
}
//func getEmployees perform get all employees information into json format.
func getEmployees(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	selDB, err := db.Query("Select * from Employee")
	if err != nil {
		panic(err.Error())
	}
	emp := Employee{}
	res := []Employee{}
	for selDB.Next() {
		var id int
		var name, city string
		err := selDB.Scan(&id, &name, &city)
		if err != nil {
			panic(err.Error())
		}
		emp.Id = id
		emp.Name = name
		emp.City = city
		res = append(res, emp)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
	fmt.Println("Done")
}
//create getEmployeeById function get employee information by id number.
func getEmployeesById(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	params := mux.Vars(r)
    employeeId:=params["id"]
	selDB, err := db.Query("Select * from Employee",employeeId)
	if err != nil {
		panic(err.Error())
	}
	emp := Employee{}
	//res := []Employee{}
	for selDB.Next() {
		var id int
		var name, city string
		err := selDB.Scan(&id, &name, &city)
		if err != nil {
			panic(err.Error())
		}
			emp.Id = id
			emp.Name = name
			emp.City = city
			//res = append(res, emp)

	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(emp)
}
func swaggerMethod(){

	router := mux.NewRouter()
	router.HandleFunc("/get", getEmployees).Methods("GET")
	router.HandleFunc("/get/{id}",getEmployeesById).Methods("GET")
	// Swagger
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	log.Fatal(http.ListenAndServe(":8080", router))

}
func handlerMethod(){
	log.Println("Server started on: http://localhost:8080")
	http.HandleFunc("/", New)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/index", Index)
	http.HandleFunc("/delete", Delete)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
//main function
func main() {
	var input int
	fmt.Println("handlerMethod 1 and Router 2")
	fmt.Scanf("%v",&input)
	if input==1 {
		handlerMethod()
	}else{
		swaggerMethod()

		}



}

/*
built an application  on crud operation with using mysql database.

*/
