package main

import "fmt"

type Person struct {
	age int
}

func main() {
	/*
	person := &Person{23}

	defer fmt.Println("1: ", person.age) // 1:  23

	defer func(p *Person) {
		fmt.Println("2: ", p.age) // 3:  18
	}(person)

	defer func() {
		fmt.Println("3: ", person.age) // 3:  18
	}()

	person.age = 18

	// 3:  18
	// 2:  18
	// 1:  23
	 */

	person := &Person{23}

	defer fmt.Printf("1: %d, %p \n", person.age, person) // 1: 23, 0xc0420080d8

	defer func(p *Person) {
		fmt.Println("2: ", p.age) // 2:  23
		fmt.Printf("2.ptr: %p \n", p) // 2.ptr: 0xc0420080d8
	}(person)

	defer func() {
		fmt.Println("3: ", person.age) // 3:  18
	}()

	person = &Person{18}
	fmt.Printf("person ptr: %p \n", person) // person ptr: 0xc0420080f8

	// person ptr: 0xc0420080f8
	// 3:  18
	// 2:  23
	// 2.ptr: 0xc0420080d8
	// 1: 23, 0xc0420080d8
}

