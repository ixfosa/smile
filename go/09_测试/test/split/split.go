package split

import "strings"

/*
	就像细胞是构成我们身体的基本单位，一个软件程序也是由很多单元组件构成的。单元组件可以是函数、结构体、方法和最终用户可能依赖的任意东西。
总之我们需要确保这些组件是能够正常运行的。单元测试是一些利用各种方法测试单元组件的程序，它会将结果与预期输出进行比较。

接下来，我们定义一个split的包，包中定义了一个Split函数，具体实现如下：
*/

// Split slices s into all substrings separated by sep and
// returns a slice of the substrings between those separators.
// Split函数并没有考虑到sep为多个字符的情况，我们来修复下这个Bug：
func Split1(s, sep string) (result []string){
	i := strings.Index(s, sep)

	for i > -1 {
		result = append(result, s[:i])
		s = s[i+1:]
		i = strings.Index(s, sep)
	}
	result = append(result, s)
	return
}

// 修复
func Split(s, sep string) (result []string) {
	i := strings.Index(s, sep)

	for i > -1 {
		result = append(result, s[:i])
		s = s[i+len(sep):] // 这里使用len(sep)获取sep的长度
		i = strings.Index(s, sep)
	}
	result = append(result, s)
	return
}
 // 在当前目录下，我们创建一个split_test.go的测试文件，并定义一个测试函数如下：