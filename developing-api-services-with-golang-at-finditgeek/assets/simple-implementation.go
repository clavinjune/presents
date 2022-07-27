package main

import (
	"database/sql"
	"net/http"
)

// START OMIT
type Repository struct {
	db *sql.DB // output
}

type Service struct {
	repo *Repository // need repo
}

type Transport struct {
	svc *Service // need svc
}

func main() {
	db, _ := sql.Open("sqlite3", ":memory:") // open database connection
	repo := &Repository{db: db}              // create repository
	svc := &Service{repo: repo}              // create service
	trp := &Transport{svc: svc}              // create transport

	http.ListenAndServe(":8000", trp) // listen for input at port 8000
}

// END OMIT

func (t *Transport) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}
