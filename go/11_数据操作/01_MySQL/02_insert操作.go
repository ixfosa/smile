package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB02 *sqlx.DB

func init() {
	db, err := sqlx.Open("mysql", "root:ixfosa@tcp(127.0.0.1)/test")
	if err != nil {
		fmt.Println("open mysql failed: ", err)
		return
	}
	DB02 = db
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


	r, err := DB02.Exec("insert into person(username, sex, email)values(?, ?, ?)", "19", "woman", "zhong@qq.com")
	if err != nil {
		fmt.Println("DB.Exec err: ", err)
		return
	}

	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("r.LastInsertId() err: ", err)
		return
	}
	fmt.Println("insert succ:", id)
	defer DB02.Close()
}
