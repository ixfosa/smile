package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	// 直接使用字符串，但是对于一些其他的正则任务，你需要 Compile 一个优化的 Regexp 结构体。
	matched, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(matched) // true

	r, _ := regexp.Compile("p([a-z]+)ch")

	fmt.Println(r.MatchString("peach")) // true

	// 查找匹配字符串的
	fmt.Println(r.FindString("peach punch")) // peach

	// 查找第一次匹配的字符串的，但是返回的匹配开始和结束位置索引，而不是匹配的内容。
	fmt.Println(r.FindStringIndex("peach punch")) // [0 5]

	// Submatch 返回完全匹配和局部匹配的字符串。例如，这里会返回 p([a-z]+)ch 和 `([a-z]+) 的信息。
	fmt.Println(r.FindStringSubmatch("peach punch")) // [peach ea]

	// 返回完全匹配和局部匹配的索引位置。
	fmt.Println(r.FindStringSubmatchIndex("peach punch")) // [0 5 1 3]

	// 带 All 的这个函数返回所有的匹配项，而不仅仅是首次匹配项。例如查找匹配表达式的所有项。
	fmt.Println(r.FindAllString("peach punch pinch", -1)) // [peach punch pinch]

	// All 同样可以对应到上面的所有函数。
	fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch", -1)) // [[0 5 1 3] [6 11 7 9] [12 17 13 15]]

	// 提供一个正整数来限制匹配次数
	fmt.Println(r.FindAllString("peach punch pinch", 2)) // [peach punch]

	// 用了字符串作为参数，并使用了如 MatchString 这样的方法。也可以提供 []byte参数并将 String 从函数命中去掉。
	fmt.Println(r.Match([]byte("peach"))) // true

	// 创建正则表示式常量时，可以使用 Compile 的变体MustCompile 。因为 Compile 返回两个值，不能用于常量。
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r) // p([a-z]+)ch

	fmt.Println(r.ReplaceAllString("a peach", "<fruit>")) // a <fruit>

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out)) // a PEACH

	fmt.Println(in) // [97 32 112 101 97 99 104]
	fmt.Println(bytes.ToUpper(in)) // [65 32 80 69 65 67 72]





}
