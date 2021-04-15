package main

import (
	"fmt"
	"log"
	"net/http"
	"html/template"
)
func toDoList(w http.ResponseWriter,r *http.Request){
	Title :="To Do List App"
	p,err:=template.ParseFiles("C:\\Users\\Gaurav\\githubProject\\read-api\\toDoList\\static\\todoListHtml.html")
	if err != nil{
		log.Fatal(err)
	}
	p.Execute(w,Title)
	fmt.Println("Done")

}

func main(){
	log.Println("Server started on: http://localhost:9080")
	http.HandleFunc("/", toDoList)
	log.Fatal(http.ListenAndServe(":9080", nil))
}