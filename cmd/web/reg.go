package main

import (
	"Task/pkg/db"
	"html/template"
	"log"
	"net/http"
)

type Register struct {
	order db.DB
}

func (re *Register) OrderReg(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	order, _ := re.order.FindById(id)
	tmpl, _ := template.ParseFiles("ui/static/order.html")
	err := tmpl.Execute(w, order)
	if err != nil {
		log.Fatalf("Error with template: %v", err)
	}

}
