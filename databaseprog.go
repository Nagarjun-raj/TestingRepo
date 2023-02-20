package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "mydb"
)

type Student struct {
	Name    string `json:"name"`
	RollNum int    `json:"roll_num"`
	Class   int    `json:"class"`
}

func addStudent(w http.ResponseWriter, req *http.Request) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	checkError(err)
	_, er := db.Exec("CREATE TABLE IF NOT EXISTS studetails (name VARCHAR,roll INT,class INT);")
	checkError(er)
	if req.Method == http.MethodPost {
		defer db.Close()
		var stuDetails Student
		er := json.NewDecoder(req.Body).Decode(&stuDetails)
		checkError(er)
		insertStmt := `insert into studetails(name,roll,class) values($1,$2,$3)`
		_, err := db.Exec(insertStmt, stuDetails.Name, stuDetails.RollNum, stuDetails.Class)
		checkError(err)
	} else {
		http.Error(w, "Bad request", http.StatusBadRequest)
	}

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func getStudent(w http.ResponseWriter, req *http.Request) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	checkError(err)
	if req.Method == http.MethodGet {
		defer db.Close()
		var stuDetails Student
		er := json.NewDecoder(req.Body).Decode(&stuDetails)
		checkError(er)
		sqlStmt := `select * from studetails where roll=$1`
		rows, e := db.Query(sqlStmt, stuDetails.RollNum)
		checkError(e)
		defer rows.Close()
		for rows.Next() {
			var name string
			var roll int
			var class int
			r := rows.Scan(&name, &roll, &class)
			checkError(r)
			fmt.Println(name, roll, class)
		}
		if err := rows.Err(); err != nil {
			panic(err)
		}

	} else {
		http.Error(w, "Bad request", http.StatusBadRequest)
	}
}

func getAll(w http.ResponseWriter, req *http.Request) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	checkError(err)
	if req.Method == http.MethodGet {
		defer db.Close()
		sqlStmt1 := `select * from studetails`
		rows, e := db.Query(sqlStmt1)
		checkError(e)
		defer rows.Close()
		for rows.Next() {
			var name string
			var roll int
			var class int
			r := rows.Scan(&name, &roll, &class)
			checkError(r)
			fmt.Println(name, roll, class)
		}
		if err := rows.Err(); err != nil {
			panic(err)
		}

	} else {
		http.Error(w, "Bad request", http.StatusBadRequest)
	}

}

func main() {
	http.HandleFunc("/addstudent", addStudent)
	http.HandleFunc("/getstudent", getStudent)
	http.HandleFunc("/getallstudents", getAll)
	http.ListenAndServe(":8080", nil)

}
