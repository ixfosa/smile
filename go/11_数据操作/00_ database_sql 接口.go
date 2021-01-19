package main

/*
						database/sql 接口
 	Go 官方没有提供数据库驱动，而是为开发数据库驱动定义了一些标准接口，开发者可以根据定义的接口来开发相应的数据库驱动，
这样做有一个好处，只要是按照标准接口开发的代码， 以后需要迁移数据库时，不需要任何修改。

						sql.Register
	这个存在于 database/sql 的函数是用来注册数据库驱动的，当第三方开发者开发数据库驱动时，
都会实现 init 函数，在 init 里面会调用这个 Register(name string, driver driver.Driver) 完成本驱动的注册。

	// https://github.com/mattn/go-sqlite3 驱动
	func init() {
		sql.Register("sqlite3", &SQLiteDriver{})
	}

	// https://github.com/mikespook/mymysql 驱动
	// Driver automatically registered in database/sql
	var d = Driver{proto: "tcp", raddr: "127.0.0.1:3306"}
	func init() {
		Register("SET NAMES utf8")
		sql.Register("mymysql", &d)
	}
	第三方数据库驱动都是通过调用这个函数来注册自己的数据库驱动名称以及相应的 driver 实现。
在 database/sql 内部通过一个 map 来存储用户定义的相应驱动。
	var drivers = make(map[string]driver.Driver)
	drivers[name] = driver
因此通过 database/sql 的注册函数可以同时注册多个数据库驱动，只要不重复。

在我们使用 database/sql 接口和第三方库的时候经常看到如下:
  import (
      "database/sql"
      _ "github.com/mattn/go-sqlite3"
  )
 _ 的意思是引入后面的包名而不直接使用这个包中定义的函数，变量等资源。


						driver.Driver
Driver 是一个数据库驱动的接口，他定义了一个 method： Open (name string)，这个方法返回一个数据库的 Conn 接口。
	type Driver interface {
		Open(name string) (Conn, error)
	}
返回的 Conn 只能用来进行一次 goroutine 的操作，也就是说不能把这个 Conn 应用于 Go 的多个 goroutine 里面。如下代码会出现错误
...
go goroutineA (Conn)  // 执行查询操作
go goroutineB (Conn)  // 执行插入操作
...
	上面这样的代码可能会使 Go 不知道某个操作究竟是由哪个 goroutine 发起的，从而导致数据混乱，
比如可能会把 goroutineA 里面执行的查询操作的结果返回给 goroutineB 从而使 B 错误地把此结果当成自己执行的插入数据。

第三方驱动都会定义这个函数，它会解析 name 参数来获取相关数据库的连接信息，解析完成后，它将使用此信息来初始化一个 Conn 并返回它。


									driver.Conn
	Conn 是一个数据库连接的接口定义，他定义了一系列方法，这个 Conn 只能应用在一个 goroutine 里面，
不能使用在多个 goroutine 里面
	type Conn interface {
		Prepare(query string) (Stmt, error)
		Close() error
		Begin() (Tx, error)
	}
Prepare 函数返回与当前连接相关的执行 Sql 语句的准备状态，可以进行查询、删除等操作。

Close 函数关闭当前的连接，执行释放连接拥有的资源等清理工作。因为驱动实现了 database/sql 里面建议的 conn pool，
所以你不用再去实现缓存 conn 之类的，这样会容易引起问题。

Begin 函数返回一个代表事务处理的 Tx，通过它你可以进行查询，更新等操作，或者对事务进行回滚、递交。


							driver.Stmt
Stmt 是一种准备好的状态，和 Conn 相关联，而且只能应用于一个 goroutine 中，不能应用于多个 goroutine。
	type Stmt interface {
		Close() error
		NumInput() int
		Exec(args []Value) (Result, error)
		Query(args []Value) (Rows, error)
	}
Close 函数关闭当前的链接状态，但是如果当前正在执行 query，query 还是有效返回 rows 数据。

NumInput 函数返回当前预留参数的个数，当返回 >=0 时数据库驱动就会智能检查调用者的参数。当数据库驱动包不知道预留参数的时候，返回 -1。

Exec 函数执行 Prepare 准备好的 sql，传入参数执行 update/insert 等操作，返回 Result 数据

Query 函数执行 Prepare 准备好的 sql，传入需要的参数执行 select 操作，返回 Rows 结果集


						driver.Tx
事务处理一般就两个过程，递交或者回滚。数据库驱动里面也只需要实现这两个函数就可以
	type Tx interface {
		Commit() error
		Rollback() error
	}
这两个函数一个用来递交一个事务，一个用来回滚事务。


								driver.Execer
这是一个 Conn 可选择实现的接口
	type Execer interface {
		Exec(query string, args []Value) (Result, error)
	}
如果这个接口没有定义，那么在调用 DB.Exec, 就会首先调用 Prepare 返回 Stmt，然后执行 Stmt 的 Exec，然后关闭 Stmt。


								driver.Result
这个是执行 Update/Insert 等操作返回的结果接口定义
	type Result interface {
		LastInsertId() (int64, error)
		RowsAffected() (int64, error)
	}
LastInsertId 函数返回由数据库执行插入操作得到的自增 ID 号。
RowsAffected 函数返回 query 操作影响的数据条目数。


								driver.Rows
Rows 是执行查询返回的结果集接口定义
	type Rows interface {
		Columns() []string
		Close() error
		Next(dest []Value) error
	}
Columns 函数返回查询数据库表的字段信息，这个返回的 slice 和 sql 查询的字段一一对应，而不是返回整个表的所有字段。
Close 函数用来关闭 Rows 迭代器。
Next 函数用来返回下一条数据，把数据赋值给 dest。dest 里面的元素必须是 driver.Value 的值除了 string，
返回的数据里面所有的 string 都必须要转换成 [] byte。如果最后没数据了，Next 函数最后返回 io.EOF。


						driver.RowsAffected
RowsAffected 其实就是一个 int64 的别名，但是他实现了 Result 接口，用来底层实现 Result 的表示方式
	type RowsAffected int64

func (RowsAffected) LastInsertId() (int64, error)
func (v RowsAffected) RowsAffected() (int64, error)


						driver.Value
Value 其实就是一个空接口，他可以容纳任何的数据
	type Value interface{}
drive 的 Value 是驱动必须能够操作的 Value，Value 要么是 nil，要么是下面的任意一种
	int64
	float64
	bool
	[]byte
	string   [*]除了Rows.Next 返回的不能是 string.
	time.Time
	driver.ValueConverter
	ValueConverter 接口定义了如何把一个普通的值转化成 driver.Value 的接口

	type ValueConverter interface {
		ConvertValue(v interface{}) (Value, error)
	}
在开发的数据库驱动包里面实现这个接口的函数在很多地方会使用到，这个 ValueConverter 有很多好处：

转化 driver.value 到数据库表相应的字段，例如 int64 的数据如何转化成数据库表 uint16 字段
把数据库查询结果转化成 driver.Value 值
在 scan 函数里面如何把 driver.Value 值转化成用户定义的值

Valuer 接口定义了返回一个 driver.Value 的方式
	type Valuer interface {
		Value() (Value, error)
	}
很多类型都实现了这个 Value 方法，用来自身与 driver.Value 的转化。


								database/sql
	database/sql 在 database/sql/driver 提供的接口基础上定义了一些更高阶的方法，用以简化数据库操作，
同时内部还建议性地实现一个 conn pool。
	type DB struct {
		driver   driver.Driver
		dsn      string
		mu       sync.Mutex // protects freeConn and closed
		freeConn []driver.Conn
		closed   bool
	}
	我们可以看到 Open 函数返回的是 DB 对象，里面有一个 freeConn，它就是那个简易的连接池。
它的实现相当简单或者说简陋，就是当执行 db.prepare -> db.prepareDC 的时候会 defer dc.releaseConn，
然后调用 db.putConn，也就是把这个连接放入连接池，每次调用 db.conn 的时候会先判断 freeConn 的长度是否大于 0，
大于 0 说明有可以复用的 conn，直接拿出来用就是了，如果不大于 0，则创建一个 conn，然后再返回之。

*/
func main() {
	
}
