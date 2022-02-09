package main

import (
	"Task_T0/pkg/db"
	"log"
	"net/http"
)

func main() {

	go Stan()
	orderRepo := db.New()
	orderRepo.FromDb()
	reg := Reg{orderRepo}
	TakeMessage("test", "test", "test-1", &reg)
	mux := http.NewServeMux()
	mux.HandleFunc("/Order", reg.OrderReg)

	log.Println("Server works on http://127.0.0.1:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
