package main

import "fmt"

/******************多态*********************/
/*
Go 通过接口来实现多态。在 Go 语言中，我们是隐式地实现接口。
一个类型如果定义了接口所声明的全部方法，那它就实现了该接口。
 */

//所有实现了接口的类型，都可以把它的值保存在一个接口类型的变量中。在 Go 中，我们使用接口的这种特性来实现多态。
//定义了接口 Interface，它包含了两个方法：calculate() 计算并返回项目的收入，而 source() 返回项目名称。
type Income interface {
	calculate() int
	source() string
}


//下面我们定义一个表示 FixedBilling 项目的结构体类型。
//项目 FixedBillin 有两个字段：projectName 表示项目名称，而 biddedAmount 表示组织向该项目投标的金额。
type FixedBilling struct {
	projectName string
	biddedAmount int
}


//TimeAndMaterial 结构体用于表示项目 Time and Material。
//结构体 TimeAndMaterial 拥有三个字段名：projectName、noOfHours 和 hourlyRate。
type TimeAndMaterial struct {
	projectName string
	noOfHours  int
	hourlyRate int
}


//下一步我们给这些结构体类型定义方法，计算并返回实际收入和项目名称。
func (fb FixedBilling) calculate() int {
	return fb.biddedAmount
}

func (fb FixedBilling) source() string {
	return fb.projectName
}

func (tm TimeAndMaterial) calculate() int {
	return tm.noOfHours * tm.hourlyRate
}

func (tm TimeAndMaterial) source() string {
	return tm.projectName
}

func calculateNetIncome(ic []Income) {
	var netincome int = 0
	for _, income := range ic {
		fmt.Printf("Income From %s = $%d\n", income.source(), income.calculate())
		netincome += income.calculate()
	}
	fmt.Printf("Net income of organisation = $%d", netincome)
}

func main() {
	project1 := FixedBilling{projectName: "Project 1", biddedAmount: 5000}
	project2 := FixedBilling{projectName: "Project 2", biddedAmount: 10000}
	project3 := TimeAndMaterial{projectName: "Project 3", noOfHours: 160, hourlyRate: 25}
	incomeStreams := []Income{project1, project2, project3}
	calculateNetIncome(incomeStreams)
}
