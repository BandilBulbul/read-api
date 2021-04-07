package main

import(
"database/sql"
_ "github.com/go-sql-driver/mysql"
	"html/template"
	"log"
"net/http"

)
type Employee struct{
	Id int
	Name string
	City string
}

func dbConn() (db *sql.DB){
	dbDriver:="mysql"
	dbUser :="root"
	dbPass :="admin"
	dbName :="goblog"
	db,err :=sql.Open(dbDriver,dbUser+":"+dbPass+"@/"+dbName)
	if err !=nil{
		panic(err.Error())
	}
	return db
}

func Insert(w http.ResponseWriter,r *http.Request){
	db :=dbConn()
	if r.Method =="POST"{
		name :=r.FormValue("name")
		city :=r.FormValue("city")
		insForm,_ :=db.Prepare("INSERT INTO Employee(name,city) Values(?,?)")
		insForm.Exec(name, city)
		log.Println("INSERT: Name: " + name + " | City: " + city)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}
func Index(w http.ResponseWriter,r *http.Request){
	db :=dbConn()
	selDB,err :=db.Query("Select * from Employee")
	p,_:=template.ParseFiles("C:\\Users\\Gaurav\\githubProject\\read-api\\CrudApplication\\src\\html\\template\\show.html")
	if err !=nil{
		panic(err.Error())
	}
	emp :=Employee{}
	res :=[]Employee{}
	for selDB.Next(){
		var id int
		var name,city string
		err :=selDB.Scan(&id,&name,&city)
		if err !=nil{
			panic(err.Error())
		}
		emp.Id=id
		emp.Name=name
		emp.City=city
		res=append(res,emp)
	}
	p.Execute(w,res)

}

func New(w http.ResponseWriter, r *http.Request) {
	News :="RegistrationPage"
	p,err:=template.ParseFiles("C:\\Users\\Gaurav\\githubProject\\read-api\\CrudApplication\\src\\html\\template\\neww.html")
	if err != nil{
		log.Fatal(err)
	}
	p.Execute(w,News)
}

func main() {
	log.Println("Server started on: http://localhost:8080")
	http.HandleFunc("/", New)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/index",Index)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

/*
built an application on crud operation with using mysql database.

*/
