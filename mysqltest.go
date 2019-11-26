package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)
func main() {
	db, err := sql.Open("mysql", "golang:zzxiaomao@tcp(47.101.66.37:3306)/cs17")
	if err != nil {
		fmt.Print(err)
		return
	}
	Get(db)
	Insert(db)
	Delete(db)
	fmt.Print("main...over...")
}

func Get(db *sql.DB) {
	rows, err := db.Query("select * from std where class=?", 1706)
	if err != nil {
		panic(err)
	}
	var arr []map[string]interface{}
	for rows.Next() {
		m := map[string]interface{}{}
		var (
			id int64
			name string
			birthday string
			class	int64
		)
		if err := rows.Scan(&id, &name, &birthday, &class); err != nil {
			panic(err)
		}
		m["id"] = id; m["name"] = name; m["birthday"] = birthday; m["class"] = class
		arr = append(arr, m)
	}
	fmt.Println(arr)
}

func Insert(db *sql.DB) {
	stmt, err := db.Prepare("insert into std (name, birthday, class) values (?, ?, ?)")
	log(err)
	res, err := stmt.Exec("vjudge", "1999-9-9", 1702)
	log(err)
	id, err := res.LastInsertId()
	log(err)
	fmt.Println(id)
}

func Delete(db *sql.DB) {
	stmt, err := db.Prepare("delete from std where name=?")
	log(err)
	res, err := stmt.Exec("vjudge")
	row, err := res.RowsAffected()
	log(err)
	id, err := res.LastInsertId()
	log(err)
	fmt.Println(row, id)
}
func log(err error) {
	if err != nil {
		panic(err)
	}
}