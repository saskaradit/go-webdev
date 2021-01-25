package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func main() {
	// user:password@tcp(localhost:5555)/dbname?charset=utf8
	db, err = sql.Open("mysql", "root:Saskarad0520@tcp(localhost:3306)/test?charset=utf8")
	check(err)
	defer db.Close()

	err = db.Ping()
	check(err)

	http.HandleFunc("/", index)
	http.HandleFunc("/rad", rad)
	http.HandleFunc("/create", create)
	http.HandleFunc("/insert", insert)
	http.HandleFunc("/read", read)
	http.HandleFunc("/update", update)
	http.HandleFunc("/delete", delete)
	http.HandleFunc("/drop", drop)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err := http.ListenAndServe(":8080", nil)
	check(err)
}

func index(w http.ResponseWriter, req *http.Request) {
	_, err = io.WriteString(w, " at Index")
	check(err)
}

func rad(res http.ResponseWriter, req *http.Request) {
	rows, err := db.Query(`SELECT Name FROM Users;`)
	check(err)
	defer rows.Close()

	var s, name string
	s = "RETRIEVED RECORDS:\n"

	// Query
	for rows.Next() {
		err = rows.Scan(&name)
		check(err)
		s += name + "\n"
	}
	fmt.Fprintln(res, s)
}

func create(res http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`CREATE TABLE Items (Name VARCHAR(20));`)
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Fprintln(res, "CREATED TABLE Items", n)
}

func insert(res http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`INSERT INTO Items Values("Rad")`)
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Fprintln(res, "INSERTED RECORD", n)
}

func read(res http.ResponseWriter, req *http.Request) {
	rows, err := db.Query(`SELECT * FROM Items;`)
	check(err)
	defer rows.Close()

	var name string
	for rows.Next() {
		err = rows.Scan(&name)
		check(err)
		fmt.Fprintln(res, "RETRIEVED RECORD:", name)
	}
}

func update(res http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`UPDATE Items SET Name="Yang Mulia" WHERE name="Rad";`)
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)

	fmt.Fprintln(res, "RETRIEVED RECORDS", n)
}

func delete(res http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`DELETE FROM Items WHERE name="Yang Mulia";`)
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec()
	check(err)

	n, err := r.RowsAffected()
	check(err)
	fmt.Fprintln(res, "DELETED RECORDS", n)
}

func drop(res http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`DROP TABLE Items;`)
	check(err)
	defer stmt.Close()

	_, err = stmt.Exec()
	check(err)

	fmt.Fprintln(res, "DROPPED TABLE Items")
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
