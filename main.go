package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	name                 string
	age                  uint16
	money                int16
	avg_grades, happines float64
}

func (u User) getAllInfo() string {
	return fmt.Sprintf("Fuck this nigger:%s. Ego ebali v jopy %d negrov", u.name, u.age)
}

func (u *User) setNewName(newName string) {
	u.name = newName
}

func home_page(page http.ResponseWriter, r *http.Request) {
	//rero := User{"rero", 25, -50, 4.2, 0.8}
	//rero.setNewName("Pidor")
	//fmt.Fprintf(page, rero.getAllInfo())
	tmpl, _ := template.ParseFiles("index.html")
	tmpl.Execute(page, nil)
}

func handleRequest() {
	http.HandleFunc("/", home_page)
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleRequest()
}
