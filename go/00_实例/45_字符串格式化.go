package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}
func main() {
	p := point{1, 2}

	fmt.Printf("%v\n", p) // {1 2}

	// 如果值是一个结构体，%+v 的格式化输出内容将包括结构体的字段名。
	fmt.Printf("%+v\n", p) // {x:1 y:2}

	// %#v 形式则输出这个值的 Go 语法表示。例如，值的运行源代码片段。
	fmt.Printf("%#v\n", p) // main.point{x:1, y:2}

	// 需要打印值的类型，使用 %T。
	fmt.Printf("%T\n", p) // main.point

	fmt.Printf("%t\n", true) // true


	// 使用 %d进行标准的十进制格式化。
	fmt.Printf("%d\n", 123) // 123

	//输出二进制表示形式
	fmt.Printf("%b\n", 14) // 1110

	// 输出给定整数的对应字符。
	fmt.Printf("%c\n", 33) //!

	// %x 提供十六进制编码
	fmt.Printf("%x\n", 456) // 1c8

	// 使用 %f 进行最基本的十进制格式化。
	fmt.Printf("%f\n", 78.9) // 78.900000

	// %e 和 %E 将浮点型格式化为（稍微有一点不同的）科学技科学记数法表示形式。
	fmt.Printf("%e\n", 123400000.0) // 1.234000e+08
	fmt.Printf("%E\n", 123400000.0) // 1.234000E+08

	// 使用 %s 进行基本的字符串输出
	fmt.Printf("%s\n", "\"string\"") // "string"

	// 像 Go 源代码中那样带有双引号的输出，使用 %q。
	fmt.Printf("%q\n", "\"string\"") // "\"string\""

	// %x 输出使用 base-16 编码的字符串，每个字节使用 2 个字符表示。
	fmt.Printf("%x\n", "hex this") // 6865782074686973

	// 要输出一个指针的值，使用 %p。
	fmt.Printf("%p\n", &p) // 0xc0420080c0

	// 使用在 % 后面使用数字来控制输出宽度。默认结果使用右对齐并且通过空格来填充空白部分。
	fmt.Printf("|%6d|%6d|\n", 12, 345) // |    12|   345|

	// 也可以指定浮点型的输出宽度，同时也可以通过 宽度.精度 的语法来指定输出的精度。
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45) // |  1.20|  3.45|

	// 要左对齐，使用 - 标志
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45) //|1.20  |3.45  |

	// 控制字符串输出时的宽度，特别是要确保他们在类表格输出时的对齐。这是基本的右对齐宽度表示。
	fmt.Printf("|%6s|%6s|\n", "foo", "b") // |   foo|     b|

	// 要左对齐，和数字一样，使用 - 标志。
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b") // |foo   |b     |

	// Sprintf 则格式化并返回一个字符串而不带任何输出。
	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s) // a string

	// 使用 Fprintf 来格式化并输出到 io.Writers而不是 os.Stdout。
	fmt.Fprintf(os.Stderr, "an %s\n", "error") // an error

}
