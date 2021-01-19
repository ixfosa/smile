package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
)

/***************** 启动外部命令和程序**********************/
/*
	os标准包，是一个比较重要的包，顾名思义，主要是在服务器上进行系统的基本操作，
如文件操作，目录操作，执行命令，信号与中断，进程，系统状态等等。在os包下，有 exec，signal，user三个子包。

可以通过变量Args来获取命令参数，os.Args返回一个字符串数
  	fmt.Println(os.Args)

在os包中，相关函数名字和作用有较重的UNIX风格，比如：
	func Chdir(dir string) error   //chdir将当前工作目录更改为dir目录
	func Getwd() (dir string, err error)    //获取当前目录
	func Chmod(name string, mode FileMode) error     //更改文件的权限
	func Chown(name string, uid, gid int) error  //更改文件拥有者owner
	func Chtimes(name string, atime time.Time, mtime time.Time) error
	func Clearenv()    //清除所有环境变量（慎用）
	func Environ() []string  //返回所有环境变量
	func Exit(code int)     //系统退出，并返回code，其中０表示执行成功并退出，非０表示错误并退出

在os包中，有关文件的处理也有很多方法，如：
	func Create(name string) (file *File, err error) // Create采用模式0666创建一个名为name的文件，如果文件已存在会截断它（为空文件）
	func Open(name string) (file *File, err error) // Open打开一个文件用于读取。
	func (f *File) Stat() (fi FileInfo, err error) // Stat返回描述文件f的FileInfo类型值
		type FileInfo interface {
			Name() string       // base name of the file
			Size() int64        // length in bytes for regular files; system-dependent for others
			Mode() FileMode     // file mode bits
			ModTime() time.Time // modification time
			IsDir() bool        // abbreviation for Mode().IsDir()
			Sys() interface{}   // underlying data source (can return nil)
		}
	func (f *File) Readdir(n int) (fi []FileInfo, err error) // Readdir读取目录f的内容，返回一个有n个成员的[]FileInfo
	func (f *File) Read(b []byte) (n int, err error) // Read方法从f中读取最多len(b)字节数据并写入b
	func (f *File) WriteString(s string) (ret int, err error) // 向文件中写入字符串
	func (f *File) Sync() (err error) // Sync递交文件的当前内容进行稳定的存储。
	func (f *File) Close() error // Close关闭文件f

	在os 包中有一个 StartProcess 函数可以调用或启动外部系统命令和二进制可执行文件；
它的第一个参数是要运行的进程，第二个参数用来传递选项或参数，第三个参数是含有系统环境基本信息的结构体。
这个函数返回被启动进程的 id（pid），或者启动失败返回错误。

	type ProcAttr struct {
		Dir string
		Env []string
		Files []*File
		Sys *syscall.SysProcAttr
	}
 */

func main131() {
	env := os.Environ()
	fmt.Println("env: ", env)


	procAttr := &os.ProcAttr{
		Env: env,
		Files: []*os.File{
			os.Stdin,
			os.Stdout,
			os.Stderr,
		},
	}
	fmt.Println("procAttr: ", procAttr)
	// func StartProcess(name string, argv []string, attr *ProcAttr) (*Process, error)
}


/******************os/signal 信号处理*********************/
/*
	一个运行良好的程序在退出（正常退出或者强制退出，如ctrl+c，kill等）时是可以执行一段清理代码，将收尾工作做完后再真正退出。
一般采用系统Signal来通知系统退出，如kill pid。在程序中针对一些系统信号设置了处理函数，当收到信号后，
会执行相关清理程序或通知各个子进程做自清理。

golang中对信号的处理主要使用os/signal包中的两个方法：一个是notify方法用来监听收到的信号；一个是 stop方法用来取消监听。
	func Stop(c chan<- os.Signal)
	func Notify(c chan<- os.Signal, sig ...os.Signal)


	func Notify(c chan<- os.Signal, sig ...os.Signal)
第一个参数表示接收信号的channel, 第二个及后面的参数表示设置要监听的信号，如果不设置表示监听所有的信号。

该函数会将进程收到的系统Signal转发给channel c。如果没有传入sig参数，那么Notify会将系统收到的所有信号转发给channel c。

Notify会根据传入的os.Signal，监听对应Signal信号，Notify()方法会将接收到对应os.Signal往一个channel c中发送。

*/


func main132() {
	 c := make(chan os.Signal)

	//os.Interrupt 表示中断
	//os.Kill 杀死退出进程
	//signal.Notify(c, os.Interrupt, os.Kill);
	 signal.Notify(c)

	// Block until a signal is received.
	 s := <- c

	fmt.Println("Got signal:", s) //Got signal: terminated
}

/***************************************/

func main133() {
	ch := make(chan os.Signal)
	//如果不指定要监听的信号，那么默认是所有信号
	signal.Notify(ch);

	//停止向ch转发信号，ch将不再收到任何信号
	signal.Stop(ch);
	//ch将一直阻塞在这里，因为它将收不到任何信号
	//所以下面的exit输出也无法执行
	<-ch;
	fmt.Println("exit")
}

/*
type Cmd struct {

	Path string   //运行命令的路径，绝对路径或者相对路径

	Args []string  // 命令参数

	Env []string    //进程环境，如果环境为空，则使用当前进程的环境

	Dir string 　//指定command的工作目录，如果dir为空，则comman在调用进程所在当前目录中运行

	Stdin io.Reader //标准输入，如果stdin是nil的话，进程从null device中读取

	Stdout io.Writer
	Stderr io.Writer

	ProcessState *os.ProcessState

	ctx             context.Context // nil means none
	lookPathErr     error           // LookPath error, if any.
	finished        bool            // when Wait was called
	childFiles      []*os.File
	closeAfterStart []io.Closer
	closeAfterWait  []io.Closer
	goroutine       []func() error
	errch           chan error // one send per goroutine
	waitDone        chan struct{}
}
*/

// 打开 POWERPNT
func main() {
	cmd := exec.Command("D:\\Program Files\\Microsoft Office\\Office16\\POWERPNT.EXE")
	fmt.Println(cmd)
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}