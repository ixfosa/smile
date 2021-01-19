package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB05 *sqlx.DB

func init() {
	db, err := sqlx.Open("mysql", "root:ixfosa@tcp(127.0.0.1)/test")
	if err != nil {
		fmt.Println("open mysql failed: ", err)
		return
	}
	DB05 = db
	//defer db.Close()  // 注意这行代码要写在上面err判断的下面
}


func main() {
	type Person struct {
		Id int `db:"id"`
		Username string `db:"username"`
		Sex string `db:"sex"`
		Email string `db:"email"`
	}

	type Place struct {
		Country string `db:"country"`
		City string `db:"citiy"`
		Telcode int `db:"telcode"`
	}

	p := []Person{}

	// func (db *DB) Select(dest interface{}, query string, args ...interface{}) error
	//err := DB05.Select(&p, "select * from person where username = ?", "long")
	err := DB05.Select(&p, "select * from person")
	if err != nil {
		fmt.Println("DB05.Select err: ", err)
		return
	}
	fmt.Printf("%v", p)
}
