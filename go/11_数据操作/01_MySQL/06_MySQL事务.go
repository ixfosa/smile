package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

/*
mysql事务特性：
    1) 原子性
    2) 一致性
    3) 隔离性
    4) 持久性

golang MySQL事务应用：
    1） import (“github.com/jmoiron/sqlx")
    2)  Db.Begin()        开始事务
    3)  Db.Commit()        提交事务
    4)  Db.Rollback()     回滚事务
*/
var DB06 *sqlx.DB

func init() {
	db, err := sqlx.Open("mysql", "root:ixfosa@tcp(127.0.0.1)/test")
	if err != nil {
		fmt.Println("open mysql failed: ", err)
		return
	}
	DB06 = db
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

	conn, err := DB06.Begin()
	if err != nil {
		fmt.Println("DB06.Begin() err: ", err)
		return
	}

	r1, err := conn.Exec("insert into person(username, sex, email) values(?, ?, ?)", "ixfosa", "man", "ixfosa@qq.com")
	if err != nil {
		fmt.Println("conn.Exec err: ", err)
		conn.Rollback()
		return
	}
	id1, err := r1.LastInsertId()
	if err != nil {
		fmt.Println("r.LastInsertId() err: ", err)
		conn.Rollback()
		return
	}
	fmt.Println("insert succ:", id1)

	r2, err := conn.Exec("update person set username=? where id=?", "ixfosa", 23)
	if err != nil {
		fmt.Println("conn.Exec err: ", err)
		conn.Rollback()
		return
	}
	id2, err := r2.LastInsertId()
	if err != nil {
		fmt.Println("r.LastInsertId() err: ", err)
		conn.Rollback()
		return
	}
	fmt.Println("insert succ:", id2)

	conn.Commit()
}
