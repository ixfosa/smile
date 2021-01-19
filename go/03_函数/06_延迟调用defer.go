package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

/*****************Golang延迟调用：**********************/
/*
defer特性：
    1. 关键字 defer 用于注册延迟调用。
    2. 这些调用直到 return 前才被执。因此，可以用来做资源清理。
    3. 多个defer语句，按先进后出的方式执行。
    4. defer语句中的变量，在defer声明时就决定了。

defer用途：
    1. 关闭文件句柄
    2. 锁资源释放
    3. 数据库连接释放

go语言 defer
	go 语言的defer功能强大，对于资源管理非常方便，但是如果没用好，也会有陷阱。
defer 是先进后出
	后面的语句会依赖前面的资源，因此如果先前面的资源先释放了，后面的语句就没法执行了。
 */
func main61() {
	var whatever [5]struct{}
	for i := range whatever{
		defer fmt.Println(i)
	}
	//4
	//3
	//2
	//1
	//0
}

/******************defer 碰上闭包*********************/
func main62() {
	var whatever [5]struct{}
	for i := range whatever{
		defer func() {
			fmt.Println(i)
		}()
	}
	//也就是说函数正常执行,由于闭包用到的变量 i 在执行的时候已经变成4,所以输出全都是4.

	for i := range whatever{
		defer func(i int) {
			fmt.Println("带参数", i)
		}(i)
	}
	//带参数 4
	//带参数 3
	//带参数 2
	//带参数 1
	//带参数 0
	//4
	//4
	//4
	//4
	//4
}


/******************defer f.Close*********************/
type Test struct {
	name string
}

func (t *Test) Close()  {
	fmt.Println(t.name, "Close")
}

func Close(t Test)  {
	t.Close()
}
func main63()  {
	//defer后面的语句在执行的时候，函数调用的参数会被保存起来，但是不执行。
	//也就是复制了一份。但是并没有说struct这里的this指针如何处理，通过这个例子可以看出go语言并没有把这个明确写出来的this指针当作参数来看待。
	a := []Test{{"a"}, {"b"}, {"c"}}
	for _, v := range a {
		defer v.Close()
		fmt.Printf("v(%p) = %v | ", &v, v) //v(0xc0420441c0) = {a} | v(0xc0420441c0) = {b} | v(0xc0420441c0) = {c}
	}
	//c Close
	//c Close
	//c Close

	for _, v := range a {
		defer Close(v)
	}
	//c Close
	//b Close
	//a Close

	//不想多写一个函数,也很简单,可以像下面这样,同样会输出c b a
	for _, v := range a {
		t := v
		defer t.Close()
	}
	//c Close
	//b Close
	//a Close

	fmt.Printf("-------")
	for _, v := range a {
		defer func(v Test) {
			v.Close()
		}(v)
	}
	//c Close
	//b Close
	//a Close
	fmt.Printf("-------")

}




//多个 defer 注册，按 FILO 次序执行 ( 先进后出 )。哪怕函数或某个延迟调用发生错误，这些调用依旧会被执行。
func test11(x int)  {
	defer println("a")
	defer println("b")

	defer func() {
		println(100/x) // div0 异常未被捕获，逐步往外传递，最终终止进程。
	}()

	defer println("c")

	//c
	//b
	//a
	//panic: runtime error: integer divide by zero
}

//延迟调用参数在注册时求值或复制，可用指针或闭包 "延迟" 读取。
func test12() {
	x, y := 10, 20

	defer func(i int) {
		println("defer", i, y) //defer 10 120
	}(x)

	x += 10
	y += 100
	println("x:", x, "y:", y) //x: 20 y: 120
}


func main64()  {
	//test11(0)
	test12()
}




//滥用 defer 可能会导致性能问题，尤其是在一个 "大循环" 里。
var lock sync.Mutex
func test13()  {
	lock.Lock()
	lock.Unlock()
}

func testDefer() {
	lock.Lock()
	defer lock.Unlock()
}

func main65() {
	func() {
		t1 := time.Now()
		for i := 0; i <100000; i++ {
			test13()
		}
		elapsed := time.Since(t1)
		fmt.Println("elapsed:", elapsed)
	}()

	func() {
		t1 := time.Now()
		for i := 0; i <100000; i++ {
			testDefer()
		}
		elapsed := time.Since(t1)
		fmt.Println("defer elapsed:", elapsed)
	}()
	//elapsed: 1.998ms
	//defer elapsed: 12.5153ms
}


/*******************defer 与 defer 与 closure********************/

func foo(a, b int) (i int, err error) {
	defer fmt.Printf("first defer err %v\n", err)
	defer func(err error) {
		fmt.Printf("second defer err %v\n", err)
	}(err)

	defer func() {
		fmt.Printf("third defer err %v\n", err)
	}()

	if b == 0 {
		err = errors.New("divided by zero!")
		return
	}

	i = a / b
	return

	//third defer err divided by zero!
	//second defer err <nil>
	//first defer err <nil>

	//解释：如果 defer 后面跟的不是一个 closure 最后执行的时候我们得到的并不是最新的值。
}

func main66() {
	foo(8, 0)
}


/*******************defer 与 return********************/
func foo1() (i int) {
	i = 0;
	defer func() {
		fmt.Println(i) //2
	}()

	return 2

	//解释：在有具名返回值的函数中（这里具名返回值为 i），执行 return 2 的时候实际上已经将 i 的值重新赋值为 2。
	//所以defer closure 输出结果为 2 而不是 1。
}

func main67()  {
	foo1()
}

/*******************defer nil 函数********************/
func foo2()  {
	var run func() = nil
	defer run()
	fmt.Println("runs")
}

func main68() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	foo2()
	//runs
	//runtime error: invalid memory address or nil pointer dereference

	//解释：名为 foo2 的函数一直运行至结束，然后 defer 函数会被执行且会因为值为 nil 而产生 panic 异常。
	//然而值得注意的是，run() 的声明是没有问题，因为在foo2函数运行完成后它才会被调用。
}


/*******************在错误的位置使用 defer********************/
//当 http.Get 失败时会抛出异常
func do() error {
	res, err := http.Get("http://www.google.com")
	defer res.Body.Close()

	if err != nil {
		fmt.Println("http.Get err", err)
		return err
	}
	return nil
}

func do1() error {
	res, err := http.Get("http://www.google.com")

	if res != nil {
		defer res.Body.Close()
	}

	if err != nil {
		fmt.Println("http.Get err", err)
		return err
	}
	return nil
	//在上述的代码中，当有错误的时候，err 会被返回，否则当整个函数返回的时候，会关闭 res.Body 。
	//解释：在这里，你同样需要检查 res 的值是否为 nil ，这是 http.Get 中的一个警告。
	//通常情况下，出错的时候，返回的内容应为空并且错误会被返回，可当你获得的是一个重定向 error 时，
	//res 的值并不会为 nil ，但其又会将错误返回。上面的代码保证了无论如何 Body 都会被关闭，
	//如果你没有打算使用其中的数据，那么你还需要丢弃已经接收的数据。
}


func main69()  {
	do() //panic: runtime error: invalid memory address or nil pointer dereference
	//因为在这里我们并没有检查我们的请求是否成功执行，当它失败的时候，我们访问了 Body 中的空变量 res ，因此会抛出异常

	//解决方案
	//总是在一次成功的资源分配下面使用 defer ，对于这种情况来说意味着：当且仅当 http.Get 成功执行时才使用 defer
	do1()
}
/*******************不检查错误********************/
//在这里，f.Close() 可能会返回一个错误，可这个错误会被我们忽略掉
func testFileCLose1() error {
	f, err := os.Open("go.txt")

	if err != nil {
		fmt.Println("os.Open err", err)
		return err
	}

	if f != nil {
		defer f.Close()
	}

	// code...

	return nil
}

//改进
func testFileCLose2() error {
	f, err := os.Open("go.txt")

	if err != nil {
		fmt.Println("os.Open err", err)
		return err
	}

	if f != nil {
		defer func() {
			if err := f.Close(); err != nil {
				// log
			}
		}()
	}

	//code...

	return nil
}

//再改进一下
//通过命名的返回变量来返回 defer 内的错误。
func testFileCLose3() (err error) {
	f, err := os.Open("go.txt")

	if err != nil {
		fmt.Println("os.Open err", err)
		return err
	}

	if f != nil {
		defer func() {
			if ferr := f.Close(); ferr != nil {
				err = ferr
			}
		}()
	}
	//code...
	return nil
}

//释放相同的资源
//如果你尝试使用相同的变量释放不同的资源，那么这个操作可能无法正常执行。
func testFileCLose4() error {
	f, err := os.Open("go.txt")

	if err != nil {
		fmt.Println("os.Open err", err)
		return err
	}

	if f != nil {
		defer func() {
			if ferr := f.Close(); ferr != nil {
				fmt.Printf("defer close go.txt err %v\n", err)
			}
		}()
	}

	f, err = os.Open("book.txt")

	if err != nil {
		fmt.Println("os.Open err", err)
		return err
	}

	if f != nil {
		defer func() {
			if err := f.Close(); err != nil {
				fmt.Printf("defer close another-book.txt err %v\n", err)
			}
		}()
	}

	//code...
	return nil

	//输出结果： defer close go.txt err close ./book.txt: file already closed
	//当延迟函数执行时，只有最后一个变量会被用到，因此，f 变量 会成为最后那个资源 (another-book.txt)。
	//而且两个 defer 都会将这个资源作为最后的资源来关闭

}
//解决方案：
func testFileCLose5() error {
	f, err := os.Open("go.txt")

	if err != nil {
		fmt.Println("os.Open err", err)
		return err
	}

	if f != nil {
		defer func(f io.Closer) {
			if ferr := f.Close(); ferr != nil {
				fmt.Printf("defer close go.txt err %v\n", err)
			}
		}(f)
	}

	f, err = os.Open("book.txt")

	if err != nil {
		fmt.Println("os.Open err", err)
		return err
	}

	if f != nil {
		defer func(f io.Closer) {
			if err := f.Close(); err != nil {
				fmt.Printf("defer close another-book.txt err %v\n", err)
			}
		}(f)
	}

	//code...
	return nil

	//输出结果： defer close go.txt err close ./book.txt: file already closed
	//当延迟函数执行时，只有最后一个变量会被用到，因此，f 变量 会成为最后那个资源 (another-book.txt)。
	//而且两个 defer 都会将这个资源作为最后的资源来关闭

}