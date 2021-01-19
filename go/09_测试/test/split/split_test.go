package split

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) { // 测试函数名必须以Test开头，必须接收一个*testing.T类型参数
	got := Split("a:b:c", ":")  // 程序输出的结果
	want := []string{"a", "b", "c"}    // 期望的结果
	fmt.Println(got)
	if !reflect.DeepEqual(got, want) { // 因为slice不能比较直接，借助反射包中的方法比较
		t.Errorf("excepted:%v, got:%v", want, got) // 测试失败输出错误提示
	}
	fmt.Println(got)
}

func TestMoreSplit(t *testing.T) {
	got := Split("abcd", "bc")
	want := []string{"a", "d"}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("excepted:%v, got:%v", want, got)
	}
}

/*
在split包路径下，执行go test命令，可以看到输出结果如下：
	D:\code\GoProjects\src\smile\09_测试\test\split>go test
	[a b c]
	[a b c]
	PASS
	ok      smile/09_测试/test/split        3.069s


// 这一次，测试失败了。我们可以为go test命令添加-v参数，查看测试函数名称和运行时间：
	D:\code\GoProjects\src\smile\09_测试\test\split>go test
	[a b c]
	[a b c]
	PASS
	ok      smile/09_测试/test/split        3.069s

	D:\code\GoProjects\src\smile\09_测试\test\split>go test
	[a b c]
	[a b c]
	--- FAIL: TestMoreSplit (0.00s)
			split_test.go:23: excepted:[a d], got:[a cd]
	FAIL
	exit status 1
	FAIL    smile/09_测试/test/split        0.642s


能清楚的看到是TestMoreSplit这个测试没有成功。 还可以在go test命令后添加-run参数，它对应一个正则表达式，
只有函数名匹配上的测试函数才会被go test命令执行。
	D:\code\GoProjects\src\smile\09_测试\test\split>go test -v
	=== RUN   TestSplit
	[a b c]
	[a b c]
	--- PASS: TestSplit (0.00s)
	=== RUN   TestMoreSplit
	--- FAIL: TestMoreSplit (0.00s)
			split_test.go:23: excepted:[a d], got:[a cd]


	D:\code\GoProjects\src\smile\09_测试\test\split>go test -v -run="More"
	=== RUN   TestMoreSplit
	--- FAIL: TestMoreSplit (0.00s)
			split_test.go:23: excepted:[a d], got:[a cd]
	FAIL
	exit status 1
	FAIL    smile/09_测试/test/split        0.436s


	D:\code\GoProjects\src\smile\09_测试\test\split>go test -v
	=== RUN   TestSplit
	[a b c]
	[a b c]
	--- PASS: TestSplit (0.00s)
	=== RUN   TestMoreSplit
	--- PASS: TestMoreSplit (0.00s)
	PASS
	ok      smile/09_测试/test/split        0.670s



*/