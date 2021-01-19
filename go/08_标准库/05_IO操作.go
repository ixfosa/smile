package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

/*******************输入输出的底层原理********************/
/*
终端其实是一个文件，相关实例如下：
	os.Stdin：标准输入的文件实例，类型为*File
	os.Stdout：标准输出的文件实例，类型为*File
	os.Stderr：标准错误输出的文件实例，类型为*File
 */

// 以文件的方式操作终端:
func main051() {
	var buf = [1024]byte{}
	n, err := os.Stdin.Read(buf[:])

	if err != nil {
		fmt.Println("os.Stdin.Read err: ", err)
		return
	}
	os.Stdout.WriteString(string(buf[:n]))
}

/*******************文件操作相关API********************/
/*
	func Create(name string) (file *File, err Error)
		根据提供的文件名创建新的文件，返回一个文件对象，默认权限是0666

	func NewFile(fd uintptr, name string) *File
		根据文件描述符创建相应的文件，返回一个文件对象

	func Open(name string) (file *File, err Error)
		只读方式打开一个名称为name的文件

	func OpenFile(name string, flag int, perm uint32) (file *File, err Error)
		打开名称为name的文件，flag是打开的方式，只读、读写等，perm是权限

	func (file *File) Write(b []byte) (n int, err Error)
		写入byte类型的信息到文件

	func (file *File) WriteAt(b []byte, off int64) (n int, err Error)
		在指定位置开始写入byte类型的信息

	func (file *File) WriteString(s string) (ret int, err Error)
		写入string信息到文件

	func (file *File) Read(b []byte) (n int, err Error)
		读取数据到b中

	func (file *File) ReadAt(b []byte, off int64) (n int, err Error)
		从off开始读取数据到b中

	func Remove(name string) Error
		删除文件名为name的文件
 */

/********************打开和关闭文件*******************/
/*
	os.Open()函数能够打开一个文件，返回一个*File和一个err。对得到的文件实例调用close()方法能够关闭文件。
 */
func main052() {
	file, err := os.Open("./ixfosa.log")

	if err != nil {
		fmt.Println("os.Open failed: ", err)
		return
	}
	defer file.Close()

	buf := [1024]byte{}
	n, _ := file.Read(buf[:])

	fmt.Println("---------")
	fmt.Println(string(buf[:n]))
}

/*******************写文件********************/
func main053() {
	// 新建文件\
	f, err := os.Create("./long.txt")

	if err != nil {
		fmt.Println("os.Create failed:", err)
		return
	}
	defer f.Close()
	for i := 0; i < 6; i++ {
		f.WriteString("long\r\n")
		f.Write([]byte("ixfosa\r\n"))
	}
}


/******************读文件*********************/
/*
	文件读取可以用file.Read()和file.ReadAt()，读到文件末尾会返回io.EOF的错误
 */
func main054()  {
	f, err := os.Open("./01_go基础/02_内置类型和函数.go")

	if err != nil {
		fmt.Println("os.Open failed: ", err)
		return
	}
	defer f.Close()

	/*buf := [1024]byte{}

	for {
		if n, _ := f.Read(buf[:]); n != 0 {
			fmt.Print(string(buf[:n]))
		} else {
			break
		}
	}*/

	buf := [128]byte{}
	content := []byte{}

	for {
		n, err := f.Read(buf[:])

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("read file err ", err)
			return
		}
		content = append(content, buf[:n]...)
	}
	fmt.Println(string(content))
}

/*******************拷贝文件********************/
func main055()  {
	srcFile, err := os.Open("./01_go基础/02_内置类型和函数.go")
	if err != nil {
		fmt.Print("os.Open failed: ", err)
		return
	}
	defer srcFile.Close()
	dstFile, err := os.Create("./ixfosa.go")

	if err != nil {
		fmt.Print("os.Create failed: ", err)
		return
	}
	defer dstFile.Close()

	buf := [1024]byte{}
	for {
		n, err := srcFile.Read(buf[:])

		if err == io.EOF {
			fmt.Println("读取完毕")
			break
		}

		if err != nil {
			fmt.Println("srcFile.Read err: ", err)
			return
		}
		//写出去
		dstFile.Write(buf[:n])
	}
}


/*******************bufio********************/
/*
bufio包实现了带缓冲区的读写，是对文件读写的封装
bufio缓冲写数据
	模式			含义
	os.O_WRONLY	只写
	os.O_CREATE	创建文件
	os.O_RDONLY	只读
	os.O_RDWR	读写
	os.O_TRUNC	清空
	os.O_APPEND	追加

func OpenFile(name string, flag int, perm FileMode) (*File, error)

四位数代表意思
	特殊权限位，拥有者位，同组用户位，其余用户位
每位值代表意思
	读是4，写是2，执行是1
		7=4+2+1
		6=4+2
例如：0666表示，特殊权限没有，拥有者可以读写，同组用户可以读写，其余用户可以读写

const (
	O_RDONLY int = syscall.O_RDONLY // open the file read-only.
	O_WRONLY int = syscall.O_WRONLY // open the file write-only.
	O_RDWR   int = syscall.O_RDWR   // open the file read-write.
	O_APPEND int = syscall.O_APPEND // append data to the file when writing.
	O_CREATE int = syscall.O_CREAT  // create a new file if none exists.
	O_EXCL   int = syscall.O_EXCL   // used with O_CREATE, file must not exist
	O_SYNC   int = syscall.O_SYNC   // open for synchronous I/O.
	O_TRUNC  int = syscall.O_TRUNC  // if possible, truncate file when opened.
)
 */
func wr()  {
	// 参数2：打开模式，所有模式d都在上面
	// 参数3是权限控制
	// w写 r读 x执行   w  2   r  4   x  1
	file, err := os.OpenFile("./zhong.txt", os.O_CREATE | os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("os.OpenFile failed:", err)
		return
	}
	defer  file.Close()

	write := bufio.NewWriter(file)

	for i := 0; i < 10; i++ {
		write.WriteString("hello zhong! \n")
	}
	write.Flush()
}

func re()  {
	file, err := os.OpenFile("./zhong.txt", os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println("os.OpenFile failed:", err)
		return
	}
	defer  file.Close()
	read := bufio.NewReader(file)
	for {
		line, _, err := read.ReadLine()
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("read.ReadLine failed:", err)
			return
		}
		fmt.Print(string(line))
	}

}

func main056()  {
	//wr()
	re()
}


/******************* ioutil工具包********************/
/*
工具包写文件
	func WriteFile(filename string, data []byte, perm os.FileMode) error

工具包读取文件
	func ReadFile(filename string) ([]byte, error)
 */
func ioutilTestW()  {
	data := "hello go"


	err := ioutil.WriteFile("./ixfosa.log", []byte(data), 0666)

	if err != nil {
		fmt.Println("ioutil.WriteFile err:", err)
		return
	}
}

func ioutilTestR()  {
	content, err := ioutil.ReadFile("./ixfosa.log")
	if err != nil {
		fmt.Println("ioutil.ReadFile err:", err)
		return
	}
	fmt.Println(string(content))
}

func main057()  {
	ioutilTestW()
	ioutilTestR()
}


/******************实现一个cat命令*********************/

// 模拟实现linux平台cat命令的功能。

func cat(r *bufio.Reader)  {
	for {
		// ReadBytes读取，直到输入中第一次出现 delim。
		// func (b *Reader) ReadBytes(delim byte) ([]byte, error)
		buf, err := r.ReadBytes('\n') //注意是字符 ''
		if err == io.EOF {
			fmt.Println("over!")
			break
		}
		if err != nil {
			fmt.Println("r.ReadBytes err: ", err)
			return
		}
		fmt.Fprintf(os.Stdout, "%s", buf)
	}
}

func main()  {
	flag.Parse()
	if flag.NArg() == 0 {
		// 如果没有参数默认从标准输入读取内容
		cat(bufio.NewReader(os.Stdin))
	}

	// 依次读取每个指定文件的内容并打印到终端
	for i := 0; i < flag.NArg(); i++ {
		f, err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Fprintf(os.Stdout, "reading from %s failed, err:%v\n", flag.Arg(i), err)
			continue
		}
		cat(bufio.NewReader(f))
	}
}