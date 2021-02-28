package main

import "fmt"

func main() {
	// 使用 range 来统计一个 slice 的元素的和。数组也可以采用这种方法。
	nums := []int{1, 2, 3}
	sum := 0
	for _, num := range nums{
		sum += num
	}
	fmt.Println("sum = ", sum) //sum =  6

	for idx, num := range nums{
		if num == 2 {
			fmt.Println("idx = ", idx) // idx =  1
		}
	}

	// range 在 map 中迭代键值对
	kvs := map[string]string{"name": "ixfosa", "age": "22"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}


	for i, c := range "go" {
		fmt.Println(i, c)
		fmt.Println(i, string(c)) // 0 g, 1 o
	}

	str1 := "go"
	for i := 0; i < len(str1); i++ {
		fmt.Println(str1[i])
		fmt.Println(i, string(str1[i])) // 0 g, 1 o
	}

	for i, c := range "小佛" {
		fmt.Println(i, string(c)) // 0 小, 3 佛
	}

	str2 := "小佛"
	for i := 0; i < len(str2); i++ {
		fmt.Println(i, string(str2[i])) // 乱码
	}

	str3 := []rune(str2)
	for i := 0; i < len(str3); i++ {
		fmt.Println(i, string(str3[i])) // // 0 小, 3 佛
	}

}
