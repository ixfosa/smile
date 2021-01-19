package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB03 *sqlx.DB
func init() {
	db, err := sqlx.Open("mysql", "root:ixfosa@tcp(127.0.0.1)/test")
	if err != nil {
		fmt.Println("open mysql failed: ", err)
		return
	}
	DB03 = db
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


	r, err := DB03.Exec("update person set username=? where id=?", "ixfosa", 23)
	if err != nil {
		fmt.Println("DB.Exec err: ", err)
		return
	}

	row, err := r.RowsAffected()
	if err != nil {
		fmt.Println("r.LastInsertId() err: ", err)
		return
	}
	fmt.Println("update succ:",row)
	defer DB03.Close()
}
