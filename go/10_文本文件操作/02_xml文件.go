package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

/*******************基于标记的XML解码********************/
/*
	XML（extensible Markup Language）格式被广泛用作一种数据交换格式，并且自成一种文件格式。
	encoding/xml 包也提供了一个更低层的基于标记的API用于XML解码。在基于标记的样式中，解析器消费输入并产生一个标记流；
四个主要的标记类型－StartElement，EndElement，CharData，和Comment－每一个都是encoding/xml包中的具体类型。
每一个对(*xml.Decoder).Token的调用都返回一个标记。

Go语言内置的 encoding/xml 包可以用在结构体和 XML 格式之间进行编解码

package xml

type Name struct {
    Local string // e.g., "Title" or "id"
}
type Attr struct { // e.g., name="value"
    Name  Name
    Value string
}
// A Token includes StartElement, EndElement, CharData,
// and Comment, plus a few esoteric types (not shown).
type Token interface{}
type StartElement struct { // e.g., <name>
    Name Name
    Attr []Attr
}

type EndElement struct { Name Name } // e.g., </name>
type CharData []byte                 // e.g., <p>CharData</p>
type Comment []byte                  // e.g., <!-- Comment -->

type Decoder struct{   }
func NewDecoder(io.Reader) *Decoder
func (*Decoder) Token() (Token, error) // returns next Token in sequence


func Marshal(v interface{}) ([]byte, error)
func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error)
Marshal函数和MarshalIndent函数返回XML编码。MarshalIndent函数功能类似于Marshal函数。区别在于有无缩进。


	Marshal处理数组或者切片时会序列化每一个元素。Marshal处理指针时，会序列化其指向的值；如果指针为nil，则啥也不输出。
Marshal处理接口时，会序列化其内包含的具体类型值，如果接口值为nil，也是不输出。Marshal处理其余类型数据时，
会输出一或多个包含数据的XML元素。


将结构体转换为XML输出时，需要注意以下规则：
	-  XMLName字段，会省略 xml.Name
	- 具有标签"-"的字段会省略
	- 具有标签"name,attr"的字段会成为该XML元素的名为name的属性
	- 具有标签",attr"的字段会成为该XML元素的名为字段名的属性
	- 具有标签",chardata"的字段会作为字符数据写入，而非XML元素
	- 具有标签",innerxml"的字段会原样写入，而不会经过正常的序列化过程
	- 具有标签",comment"的字段作为XML注释写入，而不经过正常的序列化过程，该字段内不能有"--"字符串
	- 标签中包含"omitempty"选项的字段如果为空值会省略
	  空值为false、0、nil指针、nil接口、长度为0的数组、切片、映射
	- 匿名字段（其标签无效）会被处理为其字段是外层结构体的字段
	- 如果一个字段的标签为"a>b>c"，则元素c将会嵌套进其上层元素a和b中。如果该字段相邻的字段标签指定了同样的上层元素，则会放在同一个XML元素里。


func Unmarshal(data []byte, v interface{}) error
Unmarshal解析XML编码的数据并将结果存入v指向的值。v只能指向结构体、切片或者和字符串。

解析XML编码时，需要遵守以下规则：
	- 如果结构体字段的类型为字符串或者[]byte，且标签为",innerxml"，
	  Unmarshal函数直接将对应原始XML文本写入该字段，其余规则仍适用。
	- 如果结构体字段类型为xml.Name且名为XMLName，Unmarshal会将元素名写入该字段
	- 如果字段XMLName的标签的格式为"name"或"namespace-URL name"，
	  XML元素必须有给定的名字（以及可选的名字空间），否则Unmarshal会返回错误。
	- 如果XML元素的属性的名字匹配某个标签",attr"为字段的字段名，或者匹配某个标签为"name,attr"的字段的标签名，Unmarshal会将该属性的值写入该字段。
	- 如果XML元素包含字符数据，该数据会存入结构体中第一个具有标签",chardata"的字段中，
	  该字段可以是字符串类型或者[]byte类型。如果没有这样的字段，字符数据会丢弃。
	- 如果XML元素包含注释，该数据会存入结构体中第一个具有标签",comment"的字段中，
	  该字段可以是字符串类型或者[]byte类型。如果没有这样的字段，字符数据会丢弃。
	- 如果XML元素包含一个子元素，其名称匹配格式为"a"或"a>b>c"的标签的前缀，反序列化会深入XML结构中寻找具有指定名称的元素，并将最后端的元素映射到该标签所在的结构体字段。
	  以">"开始的标签等价于以字段名开始并紧跟着">" 的标签。
	- 如果XML元素包含一个子元素，其名称匹配某个结构体类型字段的XMLName字段的标签名，
	  且该结构体字段本身没有显式指定标签名，Unmarshal会将该元素映射到该字段。
	- 如果XML元素的包含一个子元素，其名称匹配够格结构体字段的字段名，且该字段没有任何模式选项（",attr"、",chardata"等），Unmarshal会将该元素映射到该字段。
	- 如果XML元素包含的某个子元素不匹配以上任一条，而存在某个字段其标签为",any"，
	  Unmarshal会将该元素映射到该字段。
	- 匿名字段被处理为其字段好像位于外层结构体中一样。
	- 标签为"-"的结构体字段永不会被反序列化填写。
 */

// 使用 encoidng/xml 包可以很方便的将 xml 数据存储到文件中
func main021() {
	type Addr struct {
		City string
		PostCode string
	}
	type Person struct {
		XMLName xml.Name `xml:"person"`
		Id int `xml:"id,attr"`
		Name string `xml:"name"`
		Age int `xml:"age,omitempty"`
		Addr
		Friend []string `xml:"friends>friend"`
		Married bool `xml:"-"`
		Comment string `xml:",comment"`
		Desc string `xml:",innerxml"`
	}

	type People struct {
		PeopleList []Person `xml:"person"`
	}

	p := []Person{
		{
			Name: "ixfosa",
			Id: 21,
			Age: 22,
			Addr: Addr{City: "NanChang", PostCode: "332200"},
			Friend: []string{"long", "zhong"},
			Married: false,
			Comment: "this is Comment",
			Desc: "hahaha",
		},
		{
			Name: "long",
			Id: 21,
			Addr: Addr{City: "NanChang", PostCode: "332200"},
			Friend: []string{"ixfosa"},
			Married: false,
			Comment: "this is Comment",
			Desc: "lalalaala",
		},
	}

	pp := People{
		PeopleList: p,
	}

	dataXML, err := xml.MarshalIndent(pp, "", "    ")
	if err != nil {
		fmt.Println("xml.Marshal err:", err)
		return
	}

	fmt.Println(string(dataXML))

	file, err := os.Create("./person.xml")
	if err != nil {
		fmt.Println("os.Create err:", err)
		return
	}
	defer file.Close()
	encoder := xml.NewEncoder(file)
	encoder.Indent("", "    ")
	err = encoder.Encode(pp)
	if err != nil {
		fmt.Println("编码错误：", err.Error())
		return
	} else {
		fmt.Println("编码成功")
	}
}


type SConfig struct {
	XMLName  xml.Name `xml:"config"` // 指定最外层的标签为config
	SmtpServer string `xml:"smtpServer"` // 读取smtpServer配置项，并将结果保存到SmtpServer变量中
	SmtpPort int `xml:"smtpPort"`
	Sender string `xml:"sender"`
	SenderPasswd string `xml:"senderPasswd"`
	Receivers SReceivers `xml:"receivers"` // 读取receivers标签下的内容，以结构方式获取
}

type SReceivers struct {
	Flag string `xml:"flag,attr"` // 读取flag属性
	User []string `xml:"user"` // 读取user数组
}

func main022() {
	file, err := os.Open("default.xml") // For read access.
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	v := SConfig{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Println(v)
	fmt.Println("SmtpServer : ",v.SmtpServer)
	fmt.Println("SmtpPort : ",v.SmtpPort)
	fmt.Println("Sender : ",v.Sender)
	fmt.Println("SenderPasswd : ",v.SenderPasswd)
	fmt.Println("Receivers.Flag : ",v.Receivers.Flag)
	for i,element := range v.Receivers.User {
		fmt.Println(i,element)
	}
}


func main023() {
	dec := xml.NewDecoder(os.Stdin)
	var stack []string // stack of element names
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			fmt.Println("StartElement: ", stack, tok.Name.Local)
			stack = append(stack, tok.Name.Local) // push
			fmt.Println("StartElement: ", stack, tok.Name.Local)
		case xml.EndElement:
			fmt.Println("EndElement: ", stack, tok.Name.Local)
			stack = stack[:len(stack)-1] // pop
			fmt.Println("EndElement: ", stack, tok.Name.Local)
		case xml.CharData:
			if containsAll(stack, os.Args[1:]) {
				fmt.Printf("%s: %s\n", strings.Join(stack, " "), tok)
			}
		}
	}
}

func containsAll(x, y []string) bool {
	fmt.Println(x, y)
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if x[0] == y[0] {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}

func main()  {
	file, err := os.Open("./default.xml")
	if err != nil {
		fmt.Println("os.Open err:", err)
		return
	}

	dec := xml.NewDecoder(file)
	//var stack []string // stack of element names
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			fmt.Println("StartElement: ", tok.Name.Local)
		case xml.EndElement:
			fmt.Println("EndElement: ", tok.Name.Local)
		case xml.CharData:
			fmt.Println("CharData: ", string(tok))
		}
	}
}