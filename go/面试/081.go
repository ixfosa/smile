package main

import "fmt"

// 下面代码输出什么？

type ConfigOne struct {
		Deamon string
}

func (cfg *ConfigOne) String() string {
	return fmt.Sprintf("print: %v", cfg)
}

func main() {
	c := &ConfigOne{}
	c.String()
}

// 参考答案及解析：运行时错误。
	// 如果类型实现 String() 方法，当格式化输出时会自动使用 String() 方法。
	// 上面这段代码是在该类型的 String() 方法内使用格式化输出，导致递归调用，最后抛错。

// runtime: goroutine stack exceeds 1000000000-byte limit
// fatal error: stack overflow