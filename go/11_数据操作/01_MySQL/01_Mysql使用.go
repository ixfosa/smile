package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

/*
新建test数据库，person、place 表

CREATE TABLE `person` (
    `user_id` int(11) NOT NULL AUTO_INCREMENT,
    `username` varchar(260) DEFAULT NULL,
    `sex` varchar(260) DEFAULT NULL,
    `email` varchar(260) DEFAULT NULL,
    PRIMARY KEY (`user_id`)
  ) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

CREATE TABLE place (
    country varchar(200),
    city varchar(200),
    telcode int
)ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

 mysql> desc person;
    +----------+--------------+------+-----+---------+----------------+
    | Field    | Type         | Null | Key | Default | Extra          |
    +----------+--------------+------+-----+---------+----------------+
    | user_id  | int(11)      | NO   | PRI | NULL    | auto_increment |
    | username | varchar(260) | YES  |     | NULL    |                |
    | sex      | varchar(260) | YES  |     | NULL    |                |
    | email    | varchar(260) | YES  |     | NULL    |                |
    +----------+--------------+------+-----+---------+----------------+
    4 rows in set (0.00 sec)

    mysql> desc place;
    +---------+--------------+------+-----+---------+-------+
    | Field   | Type         | Null | Key | Default | Extra |
    +---------+--------------+------+-----+---------+-------+
    | country | varchar(200) | YES  |     | NULL    |       |
    | city    | varchar(200) | YES  |     | NULL    |       |
    | telcode | int(11)      | YES  |     | NULL    |       |
    +---------+--------------+------+-----+---------+-------+
    3 rows in set (0.01 sec)


 mysql使用

使用第三方开源的mysql库: github.com/go-sql-driver/mysql （mysql驱动） github.com/jmoiron/sqlx （基于mysql驱动的封装）

命令行输入 ：

    go get github.com/go-sql-driver/mysql
    go get github.com/jmoiron/sqlx

链接mysql
    database, err := sqlx.Open("mysql", "root:XXXX@tcp(127.0.0.1:3306)/test")
    //database, err := sqlx.Open("数据库类型", "用户名:密码@tcp(地址:端口)/数据库名")


type DB struct {
	driver driver.Driver
	dsn    string
	numClosed uint64
	mu           sync.Mutex // protects following fields
	freeConn     []*driverConn
	connRequests map[uint64]chan connRequest
	nextRequest  uint64 // Next key to use in connRequests.
	numOpen      int    // number of opened and pending open connections
	openerCh    chan struct{}
	closed      bool
	dep         map[finalCloser]depSet
	lastPut     map[*driverConn]string // stacktrace of last conn's put; debug only
	maxIdle     int                    // zero means defaultMaxIdleConns; negative means 0
	maxOpen     int                    // <= 0 means unlimited
	maxLifetime time.Duration          // maximum amount of time a connection may be reused
	cleanerCh   chan struct{}


	db.Prepare () 函数用来返回准备要执行的 sql 操作，然后返回准备完毕的执行状态。

	db.Query () 函数用来直接执行 Sql 返回 Rows 结果。

	stmt.Exec () 函数用来执行 stmt 准备好的 SQL 语句
}
*/

// insert 插入数据
func insertData(db *sql.DB, sql string, args ...interface{})  {
	// func (db *DB) Prepare(query string) (*Stmt, error)
	stmt, err := db.Prepare(sql)
	checkErr(err)

	// func (s *Stmt) Exec(args ...interface{}) (Result, error)
	result, err := stmt.Exec(args)
	checkErr(err)
	// type Result interface {
	// 	LastInsertId() (int64, error)
	// 	RowsAffected() (int64, error)
	// }
	id, _ := result.LastInsertId()
	affected, _ := result.RowsAffected()
	fmt.Println("id: ", id)
	fmt.Println("affected: ", affected)
}


// 更新数据
func updateData(db *sql.DB, sql string, args ...interface{})  {
	stmt, err := db.Prepare(sql)
	checkErr(err)
	result, err := stmt.Exec(args[0], args[1])
	checkErr(err)
	affected, err := result.RowsAffected()
	checkErr(err)
	fmt.Println("affected: ", affected)
}

// 更新数据
func delData(db *sql.DB, sql string, args ...interface{})  {
	stmt, err := db.Prepare(sql)
	checkErr(err)
	result, err := stmt.Exec(args[0])
	checkErr(err)
	affected, err := result.RowsAffected()
	checkErr(err)
	fmt.Println("affected: ", affected)
}

func selectData(db *sql.DB, sql string)  {
	// func (s *Stmt) Query(args ...interface{}) (*Rows, error)
	rows, err := db.Query(sql)
	checkErr(err)

	for rows.Next() {
		var id int
		var username string
		var sex string
		var email string
		err = rows.Scan(&id, &username, &sex, &email)
		checkErr(err)
		fmt.Println("id: ", id)
		fmt.Println("username: ", username)
		fmt.Println("sex: ", sex)
		fmt.Println("email: ", email)
	}
}

func main() {
	// func Open(driverName, dataSourceName string) (*DB, error)
	//database, err := sqlx.Open("数据库类型", "用户名:密码@tcp(地址:端口)/数据库名")
	db, err := sql.Open("mysql", "root:ixfosa@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println("sql.Open err: ", err)
		return
	}

	//insertSQL := "insert into person(id, username, sex, email) values (?, ?, ?, ?)"
	//insertData(db, insertSQL)

	//updateSQL := "UPDATE person SET email=? WHERE id=?"
	//email := "long@163,com"
	//id := 20
	//updateData(db, updateSQL, email, id)
	delSQL := "delete from person where id = ?"
	id := 21
	delData(db, delSQL, id)

	selectSQL := "select * from person"
	selectData(db, selectSQL)

	db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}