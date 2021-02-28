package main

import (
	"fmt"
	"math"
)

// 接口 是方法特征的命名集合。
// 要在 Go 中实现一个接口，只需要实现接口中的所有方法。
type geometry interface {
	area() float64
	perim() float64
}

type square struct {
	side float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.side * s.side
}

func (s square) perim() float64 {
	return s.side * 4
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry)  {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	squ := square{side: 6}
	cir := circle{radius: 6}

	// 结构体类型 circle 和 square 都实现了 geometry接口，所以我们可以使用它们的实例作为 measure 的参数。
	measure(squ)
	measure(cir)
}
