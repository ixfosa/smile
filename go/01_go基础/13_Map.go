package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

/*----------------Map--------------------*/
/*
map是一种无序的基于key-value的数据结构，Go语言中的map是引用类型，必须初始化才能使用。
 */


/*-----------------map定义-------------------*/
/*
Go语言中 map的定义语法如下
 	map[KeyType]ValueType
其中，
    KeyType:表示键的类型。
	ValueType:表示键对应的值的类型。

map类型的变量默认初始值为nil，需要使用make()函数来分配内存。语法为：
    make(map[KeyType]ValueType, [cap])
其中cap表示map的容量，该参数虽然不是必须的，但是我们应该在初始化map的时候就为其指定一个合适的容量。

其中K对应的key必须是支持 == 比较运算符的数据类型，所以map可以通过测试key是否相等来判断是否已经存在。
虽然浮点数类型也是支持相等运算符比较的，但是将浮点数用做key类型则是一个坏的想法，最坏的情况是可能出现的NaN和任何浮点数都不相等。
对于V对应的value数据类型则没有任何的限制。

`Map` 是一种`无序`的`键值对`的集合。Map 最重要的一点是通过 key 来快速检索数据，key 类似于索引，指向数据的值。

在一个map里所有的键都是`唯一`的，而且必须是支持`==`和`!=`操作符的类型，
切片、函数以及包含切片的结构类型这些类型由于具有`引用语义`，不能作为映射的键，使用这些类型会造成编译错误：
 	dict := map[ []string ]int{} //err, invalid map key type []string
map值可以是任意类型，没有限制。map里所有键的数据类型必须是相同的，值也必须如何，但键和值的数据类型可以不相同。
	注意：map是无序的，我们无法决定它的返回顺序，所以，每次打印结果的顺利有可能不同。

*/
func main131() {
	scoreMap := make(map[string]int, 10)
	scoreMap["java"] = 99
	scoreMap["goland"] = 88
	fmt.Println(scoreMap) //map[goland:88 java:99]

	//map也支持在声明的时候填充元素，例如：
	userInfo := map[string]string {
		"name" : "ixfosa",
		"sex" : "man",
	}
	fmt.Println("name = ", userInfo["name"]) //name =  ixfosa
	fmt.Println("userInfo = ", userInfo) //userInfo =  map[name:ixfosa sex:man]
}

/*-----------------判断某个键是否存在-------------------*/
/*
Go语言中有个判断map中键是否存在的特殊写法，格式如下:
    value, ok := map[key]
 */
func main132() {
	userInfo := map[string]string {
		"name" : "ixfosa",
		"gender" : "man", //英文逗号别忘了
	}
	v, ok := userInfo["name"]
	fmt.Printf("v:%v, ok:%v\n", v, ok) //v:ixfosa, ok:true

	if _, okk := userInfo["sex"]; okk {
		fmt.Println("人妖")
	} else {
		fmt.Println("渣渣") //渣渣
	}
}

/*----------------map的遍历--------------------*/
func mai133n() {
	scoreMap := make(map[string]int)
	scoreMap["java"] = 1
	scoreMap["html"] = 2
	scoreMap["python"] = 3
	scoreMap["javascript"] = 4
	scoreMap["go"] = 99
	for key, value := range scoreMap {
		fmt.Printf("key:%v, value:%v \n", key, value)
	}

	//但我们只想遍历key的时候，可以按下面的写法：
	for k := range scoreMap	{
		fmt.Printf("k:%v\n", k)
	}
}


/*--------------- 使用delete()函数删除键值对---------------------*/
/*
使用delete()内建函数从map中删除一组键值对，delete()函数的格式如下：
    delete(map, key)
其中，
    map:表示要删除键值对的map
    key:表示要删除的键值对的键
*/
func main134() {
	scoreMap := make(map[string]int)
	scoreMap["java"] = 1
	scoreMap["html"] = 2
	scoreMap["python"] = 3
	scoreMap["javascript"] = 4
	scoreMap["go"] = 99
	for key := range scoreMap {
		if key == "html" {
			delete(scoreMap, key)  //删除键html:2
		}
	}
	fmt.Println(scoreMap) //map[java:1 python:3 javascript:4 go:99]


}


/*------------------按照指定顺序遍历map------------------*/
func main135() {
	rand.Seed(time.Now().UnixNano()) //初始化随机数种子
	scoreMap := make(map[string]int, 10)
	fmt.Println(len(scoreMap))

	for i := 1; i <= 10; i++ {
		key := fmt.Sprintf("stu%02d", i) ////生成stu开头的字符串
		value := rand.Intn(101) ////生成0~100的随机整数
		scoreMap[key] = value
	}
	fmt.Println(scoreMap)

	keys := make([]string, 10)
	// //取出map中的所有key存入切片keys
	for k := range scoreMap{
		keys = append(keys, k)
	}
	//对切片进行排序
	sort.Strings(keys)

	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}
}


/*-----------------元素为map类型的切片-------------------*/
func main136() {
	var mapSlice = make([]map[string]string, 3)
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
		//index:0 value:map[]
		//index:1 value:map[]
		//index:2 value:map[]
	}

	// 对切片中的map元素进行初始化
	mapSlice[0] = make(map[string]string, 3)
	mapSlice[0]["java"] = "java"
	mapSlice[0]["goland"] = "goland"
	mapSlice[0]["python"] = "python"
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
		//index:0 value:map[python:python java:java goland:goland]
		//index:1 value:map[]
		//index:2 value:map[]
	}
}


/*-----------------值为切片类型的map-------------------*/
func mai137() {
	sliceMap := make(map[string][]string, 3)
	fmt.Println(sliceMap)

	key := "中国"
	value, ok := sliceMap[key]
	if !ok {
		value = make([]string, 0, 2)
	}
	value = append(value, "北京", "上海")
	sliceMap[key] = value
	fmt.Println(sliceMap)
}


/*-----------------值为map类型的map-------------------*/
func main() {
	stuList := make(map[string]map[string]string)
	value, ok := stuList["ixfosa"]
	if !ok {
		value = make(map[string]string)
	}
	value["name"] = "ixfosa"
	value["sex"] = "man"
	stuList["ixfosa"] = value
	//fmt.Println(stuList)  /
	fmt.Println(stuList["ixfosa"]["name"]) //ixfosa
}
