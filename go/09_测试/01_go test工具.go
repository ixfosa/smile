package main

/*
	Go语言中的测试依赖go test命令。编写测试代码和编写普通的Go代码过程是类似的，并不需要学习新的语法、规则或工具。

	go test命令是一个按照一定约定和组织的测试代码的驱动程序。在包目录内，所有以_test.go为后缀名的源代码文件都是go test测试的一部分，
不会被go build编译到最终的可执行文件中。

在*_test.go文件中有三种类型的函数，单元测试函数、基准测试函数和示例函数。
	类型			格式						作用
	测试函数		函数名前缀为Test			测试程序的一些逻辑行为是否正确
	基准函数		函数名前缀为Benchmark		测试函数的性能
	示例函数		函数名前缀为Example		为文档提供示例文档

	go test命令会遍历所有的 *_test.go文件中符合上述命名规则的函数，然后生成一个临时的main包用于调用相应的测试函数，
然后构建并运行、报告测试结果，最后清理测试中生成的临时文件。

Golang单元测试对文件名和方法名，参数都有很严格的要求。
	 1、文件名必须以xx_test.go命名
    2、方法必须是Test[^a-z]开头
    3、方法参数必须 t *testing.T
    4、使用go test执行单元测试

go test的参数解读：
	go test是go语言自带的测试工具，其中包含的是两类，单元测试和性能测试

通过go help test可以看到go test的使用说明：
	格式形如： go test [-c] [-i] [build flags] [packages] [flags for test binary]

参数解读：
-c : 编译go test成为可执行的二进制文件，但是不运行测试。

-i : 安装测试包依赖的package，但是不运行测试。

关于build flags，调用go help build，这些是编译运行过程中需要使用到的参数，一般设置为空
关于packages，调用go help packages，这些是关于包的管理，一般设置为空
关于flags for test binary，调用go help testflag，这些是go test过程中经常使用到的参数

-test.v : 是否输出全部的单元测试用例（不管成功或者失败），默认没有加上，所以只输出失败的单元测试用例。

-test.run pattern: 只跑哪些单元测试用例

-test.bench patten: 只跑那些性能测试用例

-test.benchmem : 是否在性能测试的时候输出内存情况

-test.benchtime t : 性能测试运行的时间，默认是1s

-test.cpuprofile cpu.out : 是否输出cpu性能分析文件

-test.memprofile mem.out : 是否输出内存性能分析文件

-test.blockprofile block.out : 是否输出内部goroutine阻塞的性能分析文件

-test.memprofilerate n : 内存性能分析的时候有一个分配了多少的时候才打点记录的问题。这个参数就是设置打点的内存分配间隔，也就是profile中一个sample代表的内存大小。默认是设置为512 * 1024的。如果你将它设置为1，则每分配一个内存块就会在profile中有个打点，那么生成的profile的sample就会非常多。如果你设置为0，那就是不做打点了。

你可以通过设置memprofilerate=1和GOGC=off来关闭内存回收，并且对每个内存块的分配进行观察。

-test.blockprofilerate n: 基本同上，控制的是goroutine阻塞时候打点的纳秒数。默认不设置就相当于-test.blockprofilerate=1，每一纳秒都打点记录一下

-test.parallel n : 性能测试的程序并行cpu数，默认等于GOMAXPROCS。

-test.timeout t : 如果测试用例运行时间超过t，则抛出panic

-test.cpu 1,2,4 : 程序运行在哪些CPU上面，使用二进制的1所在位代表，和nginx的nginx_worker_cpu_affinity是一个道理

-test.short : 将那些运行时间较长的测试用例运行时间缩短

目录结构：
    test
      |
       —— calc.go
      |
       —— calc_test.go
*/
func main() {
	
}
