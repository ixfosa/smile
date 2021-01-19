package main

import (
	"encoding/json"
	"fmt"
	"os"
)

/******************* JSON ********************/
/*
	JSON（JavaScript 对象表示，JavaScript Object Notation）一种轻量级的数据交换格式

	Go 语言通过 encoding/json 对外提供标准的 JSON 序列化和反序列化方法，
即 encoding/json.Marshal 和 encoding/json.Unmarshal，它们也是包中最常用的两个方法。
*/

/*****************序列化和反序列化**********************/
/*
	序列化和反序列化的开销完全不同，JSON 反序列化的开销是序列化开销的好几倍。
Go 语言中的 JSON 序列化过程不需要被序列化的对象预先实现任何接口，它会通过反射获取结构体或者数组中的值并以树形的结构递归地进行编码，
标准库也会根据 encoding/json.Unmarshal 中传入的值对 JSON 进行解码。Go 语言 JSON 标准库编码和解码的过程大量地运用了反射这一特性，

接口：
	JSON 标准库中提供了 encoding/json.Marshaler 和 encoding/json.Unmarshaler 两个接口分别可以影响 JSON 的序列化和反序列化结果：

	type Marshaler interface {
		MarshalJSON() ([]byte, error)
	}

	type Unmarshaler interface {
		UnmarshalJSON([]byte) error
	}

	在 JSON 序列化和反序列化的过程中，它们会使用反射判断结构体类型是否实现了上述接口，如果实现了上述接口就会优先使用对应的方法进行
编码和解码操作，除了这两个方法之外，Go 语言其实还提供了另外两个用于控制编解码结果的方法，
即 encoding.TextMarshaler 和 encoding.TextUnmarshaler：
	type TextMarshaler interface {
		MarshalText() (text []byte, err error)
	}

	type TextUnmarshaler interface {
		UnmarshalText(text []byte) error
	}
	一旦发现 JSON 相关的序列化方法没有被实现，上述两个方法会作为候选方法被 JSON 标准库调用，参与编解码的过程。
总得来说，我们可以在任意类型上实现上述这四个方法自定义最终的结果，后面的两个方法的适用范围更广，但是不会被 JSON 标准库优先调用。

标签：
	Go 语言的结构体标签也是一个比较有趣的功能，在默认情况下，当我们在序列化和反序列化结构体时，
标准库都会认为字段名和 JSON 中的键具有一一对应的关系，然而 Go 语言的字段一般都是驼峰命名法，JSON 中下划线的命名方式相对比较常见，
所以使用标签这一特性直接建立键与字段之间的映射关系是一个非常方便的设计。

结构体与 JSON 的映射：
	JSON 中的标签由两部分组成，如下所示的 name 和 age 都是标签名，后面的所有的字符串是标签选项，即 encoding/json.tagOptions，
标签名和字段名会建立一一对应的关系，后面的标签选项也会影响编解码的过程：
	type Author struct {
		Name string `json:"name,omitempty"`
		Age  int32  `json:"age,string,omitempty"`
	}

	常见的两个标签是 string 和 omitempty，前者表示当前的整数或者浮点数是由 JSON 中的字符串表示的，
而另一个字段 omitempty 会在字段为零值时，直接在生成的 JSON 中忽略对应的键值对，例如："age": 0、"author": "" 等。
标准库会使用 encoding/json.parseTag 函数来解析标签：

	func parseTag(tag string) (string, tagOptions) {
		if idx := strings.Index(tag, ","); idx != -1 {
			return tag[:idx], tagOptions(tag[idx+1:])
		}
		return tag, tagOptions("")
	}
	从该方法的实现中，我们能分析出 JSON 标准库中的合法标签是什么形式的 — 标签名和标签选项都以 , 连接，最前面的字符串为标签名，
后面的都是标签选项。

 */


/*
	encoding/json.Marshal 是 JSON 标准库中提供的最简单的序列化函数，它会接收一个 interface{} 类型的值作为参数，
这也意味着几乎全部的 Go 语言变量都可以被 JSON 标准库序列化，为了提供如此复杂和通用的功能，在静态语言中使用反射是常见的选项，
我们来深入了解一下该方法的实现：
	func Marshal(v interface{}) ([]byte, error) {
		e := newEncodeState()
		err := e.marshal(v, encOpts{escapeHTML: true})
		if err != nil {
			return nil, err
		}
		buf := append([]byte(nil), e.Bytes()...)
		encodeStatePool.Put(e)
		return buf, nil
	}
	上述方法会调用 encoding/json.newEncodeState 从全局的编码状态池中获取 encoding/json.encodeState，
随后的序列化过程都会使用这个编码状态，该结构体也会在编码结束后被重新放回池中以便重复利用。
 */



/*
结构体必须是大写字母开头的成员才会被JSON处理到，小写字母开头的成员不会有影响。
Mashal时，结构体的成员变量名将会直接作为JSON Object的key打包成JSON；Unmashal时，会自动匹配对应的变量名进行赋值，大小写不敏感。
Unmarshal时，如果JSON中有多余的字段，会被直接抛弃掉；如果JSON缺少某个字段，则直接忽略不对结构体中变量赋值，不会报错。

StructTag
如果希望手动配置结构体的成员和JSON字段的对应关系，可以在定义结构体的时候给成员打标签：
使用 omitempty，如果该字段为nil或0值（数字0,字符串"",空数组[]等），则打包的JSON结果不会有这个字段。

interface{}
interface{}类型在Unmarshal时，会自动将JSON转换为对应的数据类型：
	JSON的boolean 转换为bool
	JSON的数值 转换为float64
	JSON的字符串 转换为string
	JSON的Array 转换为[]interface{}
	JSON的Object 转换为map[string]interface{}
	JSON的null 转换为nil
注意
	一个是所有的JSON数值自动转换为float64类型，使用时需要再手动转换为需要的int，int64等类型。
	第二个是JSON的object自动转换为map[string]interface{}类型，访问时直接用JSON Object的字段名作为key进行访问。
	再不知道JSON数据的格式时，可以使用interface{}。
 */

// 结构体序列化
func main011() {
	type Person struct {
		Name string `json:"name"`
		Socre map[string]int `json:"socre"`
		Age int `json:"age,string,omitempty"`
		Friend []string `json:"friend"`
		addr string `json:addr`
		Phone string "json:phone"
	}


	p := []Person{
		{
			Name: "ixfosa",
			Socre: map[string]int{"java":100, "go": 99},
			Age: 22,
			Friend: []string{"long", "zhong"},
			addr: "NanChang",
			Phone: "13879261463",
		},
		{
			Name: "long",
			Socre: map[string]int{"math":100, "chinese": 99},
			Friend: []string{"ixfosa"},
			addr: "NanChang",
			Phone: "666666666",
		},
		{
			Name: "zhong",
			Socre: map[string]int{"math":100, "ps": 99},
			Friend: []string{"ixfosa"},
			addr: "NanChang",
			Phone: "999999999",
		},
	}
	// json.MarshalIndent函数将产生整齐缩进的输出
	// data, err := json.MarshalIndent(p, "", "    ")
	// func Marshal(v interface{}) ([]byte, error)
	data, err := json.Marshal(p)
	if err != nil {
		fmt.Println("json.Marshal err:", err)
	}
	fmt.Println(string(data))
	// [{"name":"ixfosa","socre":{"go":99,"java":100},"age":"22","friend":["long","zhong"],"Phone":"13879261463"},
	//{"name":"long","socre":{"chinese":99,"math":100},"friend":["ixfosa"],"Phone":"666666666"},
	//{"name":"zhong","socre":{"math":100,"ps":99},"friend":["ixfosa"],"Phone":"999999999"}]

	file, err := os.Create("./person.json")
	if err != nil {
		fmt.Println("os.Create err:", err)
	}
	defer file.Close()
	// func NewEncoder(w io.Writer) *Encoder
	encoder := json.NewEncoder(file)

	err = encoder.Encode(p)
	if err != nil {
		fmt.Println("encoder.Encode err:", err)
	} else {
		fmt.Println("编码成功")
	}
}

// 结构体反序列化
func main012() {
	type Person struct {
		Name string `json:"name"`
		Socre map[string]int `json:"socre"`
		Age int `json:"age,string,omitempty"`
		Friend []string `json:"friend"`
		addr string `json:addr`
		Phone string "json:phone"
	}

	p := []Person{}
	// "desc":"hahaha" Person中没有desc字段
	data := `[{"name":"ixfosa","socre":{"go":99,"java":100},"age":"22","friend":["long","zhong"],"Phone":"13879261463", "desc":"hahaha"},{"name":"long","socre":{"chinese":99,"math":100},"friend":["ixfosa"],"Phone":"666666666"},{"name":"zhong","socre":{"math":100,"ps":99},"friend":["ixfosa"],"Phone":"999999999"}]`
	// func Marshal(v interface{}) ([]byte, error)
	err := json.Unmarshal([]byte(data), &p)
	if err != nil {
		fmt.Println("json.Unmarshal err:", err)
	}
	fmt.Println(p)
	// [{ixfosa map[go:99 java:100] 22 [long zhong]  13879261463}
	// {long map[chinese:99 math:100] 0 [ixfosa]  666666666}
	// {zhong map[math:100 ps:99] 0 [ixfosa]  999999999}]

	pp := []Person{}
	file, err := os.Open("./person.json")
	if err != nil {
		fmt.Println("os.Open err:", err)
	}
	defer file.Close()
	// func NewEncoder(w io.Writer) *Encoder
	encoder := json.NewDecoder(file)

	err = encoder.Decode(&pp)
	if err != nil {
		fmt.Println("encoder.Encode err:", err)
	} else {
		fmt.Println("编码成功")
		fmt.Println(pp)
	}
}


/***************************************/
// map转json示例:
func main()  {
	m := make(map[string]interface{})
	m["name"] = "long"
	m["age"] = 22
	m["addr"] = "NanChang"

	data, err := json.Marshal(m)

	if err != nil {
		fmt.Println("json.Marshal err: ", err)
		return
	}

	fmt.Println(string(data)) // {"addr":"NanChang","age":22,"name":"long"}

	mm := make(map[string]interface{})

	err = json.Unmarshal([]byte(data), &mm)
	if err != nil {
		fmt.Println("json.Unmarshal err:", err)
		return
	}
	fmt.Println(mm) // map[addr:NanChang age:22 name:long]
}

/***************************************/
/*

 */


/***************************************/
/*

 */

/***************************************/
/*

 */


/***************************************/
/*

 */


/***************************************/
/*

 */

/***************************************/
/*

 */


/***************************************/
/*

 */


/***************************************/
/*

 */

/***************************************/
/*

 */


/***************************************/
/*

 */


/***************************************/
/*

 */

/***************************************/
/*

 */


/***************************************/
/*

 */


/***************************************/
/*

 */

/***************************************/
/*

 */


/***************************************/
/*

 */

/***************************************/
/*

 */

/***************************************/
/*

 */


/***************************************/
/*

 */


/***************************************/
/*

 */

/***************************************/
/*

 */


/***************************************/
/*

 */

/***************************************/
/*

 */

/***************************************/
/*

 */


/***************************************/
/*

 */