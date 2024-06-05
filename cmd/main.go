package main

import (
	"log"
	"net/http"

	"github.com/dev-soubhagya/employee_management/internal/db"
	"github.com/dev-soubhagya/employee_management/internal/handler"
)

func main() {
	db.InitDB()
	defer db.CloseDB()

	router := handler.SetupRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
