##  BeanUtils

 BeanUtils主要解决 的问题： 把对象的属性数据封装 到对象中。

使用BenUtils组件：

1. 引入commons-beanutils-1.8.3.jar核心包
2. 引入日志支持包: commons-logging-1.1.3.jar



如果缺少日志jar文件，报错：

```java
java.lang.NoClassDefFoundError: org/apache/commons/logging/LogFactory
at org.apache.commons.beanutils.ConvertUtilsBean.<init>(ConvertUtilsBean.java:157)
at org.apache.commons.beanutils.BeanUtilsBean.<init>(BeanUtilsBean.java:117)
at org.apache.commons.beanutils.BeanUtilsBean$1.initialValue(BeanUtilsBean.java:68)
```

 **BeanUtils的好处：**

      	1. BeanUtils设置属性值的时候，如果属性是基本数据 类型，BeanUtils会自动帮我转换数据类型。
   	2. BeanUtils设置属性值的时候底层也是依赖于get或者Set方法设置以及获取属性值的。
      	3. BeanUtils设置属性值,如果设置的属性是其他的引用 类型数据，那么这时候必须要注册一个类型转换器。

------

```java
package cn.ixfosa.beans;

import java.util.Date;

/**
 * 1. 实体类设计
 *
 */
public class Admin {

   private int id;
   private String userName;
   private String pwd;
   private int age;
   private Date birth;
   
    
   public Date getBirth() {
      return birth;
   }
   public void setBirth(Date birth) {
      this.birth = birth;
   }
   public int getAge() {
      return age;
   }
   public void setAge(int age) {
      this.age = age;
   }
   public String getPwd() {
      return pwd;
   }
   public void setPwd(String pwd) {
      this.pwd = pwd;
   }
   public int getId() {
      return id;
   }
   public void setId(int id) {
      this.id = id;
   }
   public String getUserName() {
      return userName;
   }
   public void setUserName(String userName) {
      this.userName = userName;
   }
   @Override
   public String toString() {
      return "Admin [age=" + age + ", birth=" + birth + ", id=" + id
            + ", pwd=" + pwd + ", userName=" + userName + "]";
   }
}
```

```java
package cn.ixfosa.beans;

import java.text.ParseException;
import java.text.SimpleDateFormat;
import java.util.Date;
import java.util.HashMap;
import java.util.Map;

import org.apache.commons.beanutils.BeanUtils;
import org.apache.commons.beanutils.ConvertUtils;
import org.apache.commons.beanutils.Converter;
import org.apache.commons.beanutils.locale.converters.DateLocaleConverter;
import org.junit.Test;

public class App {

	//1. 对javabean的基本操作
	@Test
	public void test1() throws Exception {
		
		// a. 基本操作
		Admin admin = new Admin();
		//admin.setUserName("Jack");
		//admin.setPwd("999");
		
		// b. BeanUtils组件实现对象属性的拷贝
		BeanUtils.copyProperty(admin, "userName", "jack");
		BeanUtils.setProperty(admin, "age", 18);
		
		// 总结1： 对于基本数据类型，会自动进行类型转换!
		
		
		// c. 对象的拷贝
		Admin newAdmin = new Admin();
		BeanUtils.copyProperties(newAdmin, admin);
		
		// d. map数据，拷贝到对象中
		Admin adminMap = new Admin();
		Map<String,Object> map = new HashMap<String,Object>();
		map.put("userName", "Jerry");
		map.put("age", 29);
		// 注意：map中的key要与javabean的属性名称一致
		BeanUtils.populate(adminMap, map);
		
		// 测试
		System.out.println(adminMap.getUserName());
		System.out.println(adminMap.getAge());
	}
	
	
	//2. 自定义日期类型转换器
	@Test
	public void test2() throws Exception {
		// 模拟表单数据
		String name = "jack";
		String age = "20";
		String birth = "   ";
		
		// 对象
		Admin admin = new Admin();
		
		// 注册日期类型转换器：1， 自定义的方式
		ConvertUtils.register(new Converter() {
			// 转换的内部实现方法，需要重写
			@Override
			public Object convert(Class type, Object value) {
				
				// 判断
				if (type != Date.class) {
					return null;
				}
				if (value == null || "".equals(value.toString().trim())) {
					return null;
				}
				
				try {
					// 字符串转换为日期
					SimpleDateFormat sdf = new SimpleDateFormat("yyyy-MM-dd");
					return sdf.parse(value.toString());
				} catch (ParseException e) {
					throw new RuntimeException(e);
				}
			}
		},Date.class);
		
		// 把表单提交的数据，封装到对象中
		BeanUtils.copyProperty(admin, "userName", name);
		BeanUtils.copyProperty(admin, "age", age);
		BeanUtils.copyProperty(admin, "birth", birth);
		
		//------ 测试------
		System.out.println(admin);
	}
	
	//2. 使用提供的日期类型转换器工具类
	@Test
	public void test3() throws Exception {
		// 模拟表单数据
		String name = "userName";
		String age = "20";
		String birth = null;
		
		// 对象
		Admin admin = new Admin();
		
		// 注册日期类型转换器：2， 使用组件提供的转换器工具类
		ConvertUtils.register(new DateLocaleConverter(), Date.class);
				
		// 把表单提交的数据，封装到对象中
		BeanUtils.copyProperty(admin, "userName", name);
		BeanUtils.copyProperty(admin, "age", age);
		BeanUtils.copyProperty(admin, "birth", birth);
		
		//------ 测试------
		System.out.println(admin);
	}
}
```



```java
public class WebUtils {
    
	@Deprecated
	public static <T> T copyToBean_old(HttpServletRequest request, Class<T> clazz) {
		try {
			// 创建对象
			T t = clazz.newInstance();
			
			// 获取所有的表单元素的名称
			Enumeration<String> enums = request.getParameterNames();
			// 遍历
			while (enums.hasMoreElements()) {
				// 获取表单元素的名称:<input type="password" name="pwd"/>
				String name = enums.nextElement();  // pwd
				// 获取名称对应的值
				String value = request.getParameter(name);
				// 把指定属性名称对应的值进行拷贝
				BeanUtils.copyProperty(t, name, value);
			}
			
			return t;
		} catch (Exception e) {
			throw new RuntimeException(e);
		}
	}
	
	/**
	 * 处理请求数据的封装
	 */
	public static <T> T copyToBean(HttpServletRequest request, Class<T> clazz) {
		try {
			// （注册日期类型转换器）
			// 创建对象
			T t = clazz.newInstance();
			BeanUtils.populate(t, request.getParameterMap());
			return t;
		} catch (Exception e) {
			throw new RuntimeException(e);
		}
	}
}
```



## junit-单元测试

**目前存在的问题**

	1. 目前的方法如果需要测试，都需要在main方法上调用
	2. 目前的结果都需要我们人工对比。



**junit要注意的细节：**

      	1.  如果使用junit测试一个方法的时候，在junit窗口上显示绿条那么代表测试正确，如果是出现了红条，则代表该方法测试出现了异常不通过。
   	2.  如果点击方法名、 类名、包名、 工程名运行junit分别测试的是对应的方法，类、 包中 的所有的test方法，工程中的所有test方法。
      	3.  @Test测试的方法不能是static修饰与不能带有形参。
         	4.  如果测试一个方法的时候需要准备测试的环境或者是清理测试的环境，那么可以@Before、         @After 、@BeforeClass、 @AfterClass这四个注解。
        @Before、 @After 是在每个测试方法测试的时候都会调用一次， @BeforeClass、 @AfterClass是在所有的测试方法测试之前与测试之后调用一次而已。

```java
@Test

​	表示单元测试方法。

@Before 

​	所修饰的方法应是非static的（且没有参数，返回值为void）。

​	表示这个方法会在本类中的每个单元测试方法之前都执行一次。

@After 

​	所修饰的方法应是非static的（且没有参数，返回值为void）。

​	表示这个方法会在本类中的每个单元测试方法之后都执行一次。

@BeforeClass 

​	所修饰的方法应是static的（且没有参数，返回值为void）。

​	表示这个方法会在本类中的所有单元测试方法之前执行，只执行一次。

@AfterClass 

​	所修饰的方法应是static的（且没有参数，返回值为void）。

​	表示这个方法会在本类中的所有单元测试方法之后执行，只执行一次。
```



**junit使用规范：**

      	1. 一个类如果需要测试，那么该类就应该对应着一个测试类，测试类的命名规范 ： 被测试类的类名+ Test.
   	2. 一个被测试的方法一般对应着一个测试的方法，测试的方法的命名规范是： test+ 被测试的方法的方法名



**Assert断言工具类**
        其中有一些静态的工具方法（不符合期望就抛异常）：

```java
	assertTrue(...)		参数的值应是true
    assertFalse(...)	参数的值应是false  

    assertNull(...)		应是null值
    assertNotNull(...)	应是非null的值

    assertSame(...)		使用==比较的结果为true（表示同一个对象）
    AssertNotSame(...)	使用==比较的结果为false

    assertEquals(...)	两个对象equals()方法比较结果为true
```

```java
package cn.ixfosa.junit;

import static org.junit.Assert.*;

import java.util.Arrays;

import javax.management.RuntimeErrorException;

import org.junit.Test;

public class Demo {
	
	@Test //注解
	public	 void getMax(int a, int b){
	/*	int a = 3;
		int b = 5 ;*/
		int max = a>b?a:b;
		System.out.println("最大值："+max);
	}

	@Test
	public void sort(){
		int[] arr = {12,4,1,19};
		for(int i = 0 ; i  < arr.length-1 ; i++){
			for(int j = i+1 ; j<arr.length ; j++){
				if(arr[i]>arr[j]){
					int temp = arr[i];
					arr[i] = arr[j];
					arr[j] = temp;
				}
			}
		}
		System.out.println("数组的元素："+Arrays.toString(arr));
	}
}
```



```java
import java.io.FileInputStream;
import java.io.IOException;

import org.junit.After;
import org.junit.AfterClass;
import org.junit.Before;
import org.junit.BeforeClass;
import org.junit.Test;

public class Demo {
	
	//准备测试的环境
	//@Before
	@BeforeClass
	public static void beforeRead(){
		System.out.println("准备测试环境成功...");
	}
	
	//读取文件数据，把把文件数据都
	@Test
	public void readFile() throws IOException{
		FileInputStream fileInputStream = new FileInputStream("F:\\a.txt");
		int content  = fileInputStream.read();
		System.out.println("内容："+content);
	}
	
	
	@Test
	public void sort(){
		System.out.println("读取文件数据排序..");
	}

	//清理测试环境的方法
	//@After 
	@AfterClass
	public static void afterRead(){
		System.out.println("清理测试环境..");
	}
}
```



```java
package cn.ixfosa.junit;

public class Tool {
	
	public static int getMax(){
		int a = 3;
		int b  =5; 
		int max = a>b?a:b;
		return max;
	}
	
	public static int getMin(){
		int a = 3;
		int b = 5; 
		int min = a<b?a:b;
		return min;
	}
}

/*-------------------------------------*/

package cn.ixfosa.junit;

import junit.framework.Assert;

import org.junit.Test;

//测试类
public class ToolTest {
	
	@Test
	public void testGetMax(){
		int max = Tool.getMax();
		if(max!=5){
			throw new RuntimeException();
		}else{
			System.out.println("最大值："+ max);
		}
		
		//断言
		//Assert.assertSame(5, max); // expected 期望   actual  真实     ==
		Assert.assertSame(new String("abc"), "abc");
		Assert.assertEquals(new String("abc"), "abc"); //底层是使用Equals方法比较的
		Assert.assertNull("aa");
		Assert.assertTrue(true);
	}
	
	@Test
	public void  testGetMin(){
		int min = Tool.getMin(); 
		if(min!=3){
			throw new RuntimeException();
		}else{
			System.out.println("最小值："+ min);
		}
	}
}
```



## Domj4-xml解析

### XML

`XML` 指可`扩展标记语言`, 被设计用来传输和存储数据。

**作用**

**XML 作用**

1. 描述带关系的数据（软件的配置文件

2. 数据的载体（小型的“数据库”）



###  **XML语法**

​		xml文件以xml后缀名结尾。

​		xml文件需要使用xml解析器去解析。浏览器内置了xml解析器。

#### 标签

​		语法： <student></student>  开始标签  标签体内容  结束标签

​				1）<student/> 或 <student></student> 空标签。没有标签体内容

​				2）xml标签名称区分大小写。

​				3）xml标签一定要正确配对。

​				4）xml标签名中间不能使用空格

​				5）xml标签名不能以数字开头

​				**6）注意： 在一个xml文档中，有且仅有`一个根标签`**

#### **属性**

​			语法： <Student name="eric">student</Student>

​			注意：

​					1）属性值必须以引号包含，不能省略，也不能单双引号混用！！！

​					2）一个标签内可以有多个属性，但不能出现重复的属性名！！！

#### 注释

​		语言： <!--  xml注释 -->

​		练习：

​				通讯录系统

​				联系人数据：编号 （唯一的） 姓名  年龄  电话 邮箱  QQ 

​				 要求： 

​				contact.xml

​					1）设计一个xml文件，用于存储联系人数据

​					2）这个xml文件可以多个联系人。

```xml
<?xml version="1.0" encoding="utf-8"?>
<contactList>
	<contact id="001">
		<name>张三</name>
		<age>20</age>
		<phone>134222223333</phone>
		<email>zhangsan@qq.com</email>
		<qq>432221111</qq>
	</contact>
	<contact id="003">
		<name>lisi</name>
		<age>20</age>
		<phone>134222225555</phone>
		<email>lisi@qq.com</email>
		<qq>432222222</qq>
	</contact>
</contactList>
```



#### 文档声明

​		语法： <?xml version="1.0" encoding="utf-8"?>

​				`version`: xml的版本号

​				`encoding`： 解析xml文件时查询的码表（解码过程时查询的码表）

​		注意：

​				1）如果在ecplise工具中开发xml文件，保存xml文件时自动按照文档声明的encoding来保存文件。

​				2）如果用记事本工具修改xml文件，注意保存xml文件按照文档声明的encoding的码表来保存。

 

#### 转义字符

​		在xml中内置了一些特殊字符，这些特殊字符不能直接被浏览器原样输出。如果希望把这些特殊字符按照原样输出到浏览器，对这些特殊字符进行转义。转义之后的字符就叫转义字节。

​		特殊字符  转义字符

​				 <     <

​				 >     >

​				 "     "

​				 &     &

​				空格    &nsbp;

​				

#### CDATA块

​		作用： 可以让一些需要进行包含特殊字符的内容统一进行原样输出。

```xml
<?xml version="1.0" encoding="utf-8"?>
<codes>
	<code>&lt;p&gt;段落&lt;/p&gt;</code>
<code>
	<![CDATA[	<html><head>head</head><body>body</body></html>
	]]>
</code>
</codes>
```



#### 处理指令

​			作用： 告诉xml解析如何解析xml文档

​			案例： <?xml-stylesheet type="text/css" href="1.css"?> 告诉xml解析该xml文档引用了哪个css文件

​			需要提前xml内容可以使用xml-stylesheet指令指令

```xml
<?xml version="1.0" encoding="utf-8"?>
<?xml-stylesheet type="text/css" href="contact.css"?>
<contactList>
	<contact id="001">
		<name>张三</name>
		<age>20</age>
		<phone>134222223333</phone>
		<email>zhangsan@qq.com</email>
		<qq>432221111</qq>
	</contact>
	<contact id="003">
		<name>lisi</name>
		<age>20</age>
		<phone>134222225555</phone>
		<email>lisi@qq.com</email>
		<qq>432222222</qq>
	</contact>
</contactList>
```

```css
<!--  contact.css -->

contact{
	color:red;
	font-size:20px;
	width:150px;
	height:100px;
	display:block;/*以块状显示该标签*/
	margin-top:40px;
	background-color:green;
}
```



###  XML约束

XML约束要求：大家能够看懂约束内容，根据约束内容写出符合规则的xml文件。

**引入**
		XML语法： 规范的xml文件的基本编写规则。（由w3c组织制定的）
		XML约束： 规范XML文件数据内容格式的编写规则。（由开发者自行定义）

**XML约束技术**
		DTD约束：语法相对简单，功能也相对简单。学习成本也低。
		Schema约束：语法相对复杂，功能也相对强大。学习成本相对高！！！（名称空间）



#### **DTD约束**

**导入dtd方式**

1. 内部导入

```xml
<!DOCTYPE note [
  <!ELEMENT note (to,from,heading,body)>
  <!ELEMENT to      (#PCDATA)>
  <!ELEMENT from    (#PCDATA)>
  <!ELEMENT heading (#PCDATA)>
  <!ELEMENT body    (#PCDATA)>
]>
```

2. 本地文件系统-外部导入：

```dtd
  <!ELEMENT note (from?,to+,heading*,body+)>
  <!ELEMENT to      EMPTY>
  <!ELEMENT from    (#PCDATA)>
  <!ELEMENT heading (#PCDATA)>
  <!ELEMENT body    (#PCDATA)>
  <!ATTLIST to id ID #REQUIRED>
```

```xml
<?xml version="1.0"?>
<!DOCTYPE note SYSTEM "note.dtd">
<note>
  <to id="a1"></to>
  <to id="a2"></to>
  <to id="a3"></to>
  <heading>Reminder</heading>
  <heading>Reminder</heading>
  <heading>Reminder</heading>
  <body>Don't forget me this weekend</body>
  <body>Don't forget me this weekend</body>
  <body>Don't forget me this weekend</body>
</note>
```

3. 公共的外部导入-外部导入:

```xml
<!DOCTYPE 根元素 PUBLIC "http://gz.ixfosa.cn/itcast.dtd">
```

​		

DTD语法
		约束标签

```dtd
	<!ELEMENT 元素名称 类别>  或 <!ELEMENT 元素名称 (元素内容)>
```

			*  **类别：**
			空标签： EMPTY。 表示元素一定是空元素。
			普通字符串： (#PCDATA)。表示元素的内容一定是普通字符串（不能含有子标签）。
			任何内容： ANY。表示元素的内容可以是任意内容（包括子标签） 


​	

			*  **元素内容:**
					顺序问题：

```dtd
<!ELEMENT 元素名称 (子元素名称 1,子元素名称 2,.....)>
按顺序出现子标签
```

​					次数问题：
​							标签    ：  必须且只出现1次。
​							标签+  ： 至少出现1次
​						    标签*   ： 0或n次。
​							标签？ ： 0 或1次。

			* **约束属性**

```dtd
<!ATTLIST 元素名称 属性名称 属性类型 默认值>
```

​				默认值：
​						#REQUIRED 属性值是必需的 
​						#IMPLIED   属性不是必需的 
​						#FIXED value 属性不是必须的，但属性值是固定的

​				属性类型：控制属性值的
​					CDATA ：表示普通字符串 
​					(en1|en2|..)： 表示一定是任选其中的一个值
​					ID：表示在一个xml文档中该属性值必须唯一。值不能以数字开头



#### Schema约束

​		名称空间：告诉xml文档的哪个元素被哪个schema文档约束。 在一个xml文档中，不同的标签可以受到不同的schema文档的约束。

​		 1）一个名称空间受到schema文档约束的情况
​		 2）多个名称空间受到多个schema文档约束的情况
​		 3）默认名称空间的情况
​		 4）没有名称空间的情况

```xsd
<?xml version="1.0" encoding="UTF-8" ?> 
<xs:schema xmlns:xs="http://www.w3.org/2001/XMLSchema"
					  targetNamespace="http://www.ixfosa.cn"
					  elementFormDefault="qualified">
	<xs:element name='书架' >
		<xs:complexType>
			<xs:sequence maxOccurs='unbounded' >
				<xs:element name='书' >
					<xs:complexType>
						<xs:sequence>
							<xs:element name='书名' type='xs:string' />
							<xs:element name='作者' type='xs:string' />
							<xs:element name='售价' type='xs:integer' />
						</xs:sequence>
					</xs:complexType>
				</xs:element>
			</xs:sequence>
		</xs:complexType>
	</xs:element>
</xs:schema>
```

```xml
<?xml version="1.0" encoding="UTF-8"?>
<ixfosa:书架 xmlns:ixfosa="http://www.ixfosa.cn"
				xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
				xsi:schemaLocation="http://www.ixfosa.cn book.xsd">
	<ixfosa:书>
		<ixfosa:书名>JavaScript网页开发</ixfosa:书名>
		<ixfosa:作者>张孝祥</ixfosa:作者>
		<ixfosa:售价>28</ixfosa:售价>
	</ixfosa:书>

</ixfosa:书架>
```



### XML解析

#### 引入

​		xml文件除了给开发者看，更多的情况使用程序读取xml文件的内容。这叫做xml解析

​		XML解析方式（原理不同）

​				`DOM解析`

​				`SAX解析`



#### 	XML解析工具

​				**DOM解析原理:**

​						1）JAXP （oracle-Sun公司官方）

​						2）JDOM工具（非官方）

​						3）Dom4J工具（非官方）

​							三大框架（默认读取xml的工具就是Dom4j）

​						.......

 

​				**SAX解析原理：**

​						1）Sax解析工具（oracle-sun公司官方）

 

#### 什么是DOM解析

​			DOM解析原理：xml解析器一次性把整个xml文档加载进内存，然后在内存中构建一颗`Document的对象树`，通过Document对象，得到树上的节点对象，通过`节点对象`访问（操作到

xml文档的内容。



### Dom4j工具

​			非官方，不在jdk中。

​			使用步骤：

​				1）导入dom4j的核心包。`dom4j-1.6.1.jar`

​				2）编写Dom4j读取xml文件代码

```xml
<!-- /src/contact.xml-->
<?xml version="1.0" encoding="utf-8"?>
<contactList>
	<contact id="001">
		<name>张三</name>
		<age>20</age>
		<phone>134222223333</phone>
		<email>zhangsan@qq.com</email>
		<qq>432221111</qq>
	</contact>
	<contact id="002">
		<name>李四</name>
		<age>20</age>
		<phone>134222225555</phone>
		<email>lisi@qq.com</email>
		<qq>432222222</qq>
	</contact>
</contactList>
```



```java
import java.io.File;

import org.dom4j.Document;
import org.dom4j.DocumentException;
import org.dom4j.io.SAXReader;

/**
 * 第一个Dom4j读取xml文档的例子
 */
public class Demo {
	public static void main(String[] args) {
        
		try {
			//1.创建一个xml解析器对象
			SAXReader reader = new SAXReader();
            
			//2.读取xml文档，返回Document对象
			Document doc = reader.read(new File("./src/contact.xml"));
			
			System.out.println(doc);
            
		} catch (DocumentException e) {
			e.printStackTrace();
			throw new RuntimeException(e);
		}
	}
}
```



#### Domj4读取xml文件

##### 节点

​		`Iterator  Element.nodeIterator()`    获取当前标签节点下的所有子节点

```java
import java.io.File;
import java.util.ArrayList;
import java.util.Iterator;
import java.util.List;

import org.dom4j.Document;
import org.dom4j.Element;
import org.dom4j.io.SAXReader;

public class Demo {
    
    /**
	 * 得到节点信息
	 */
	@Test
	public void test1() throws Exception{
        
		//1.读取xml文档，返回Document对象
		SAXReader reader = new SAXReader();
		Document doc = reader.read(new File("./src/contact.xml"));
		
		//2.nodeIterator: 得到当前节点下的所有子节点对象(不包含孙以下的节点)
		Iterator<Node> it = doc.nodeIterator();
        
		while(it.hasNext()){//判断是否有下一个元素
            
			Node node = it.next();//取出元素
			String name = node.getName();//得到节点名称
            
			//System.out.println(name);
			//System.out.println(node.getClass());
            
			//继续取出其下面的子节点
			//只有标签节点才有子节点
			//判断当前节点是否是标签节点
			if(node instanceof Element){
				Element elem = (Element)node;
				Iterator<Node> it2 = elem.nodeIterator();
				while(it2.hasNext()){
					Node n2 = it2.next();
					System.out.println(n2.getName());
				}
			}
		}
	}
	
	/**
	 * 遍历xml文档的所有节点
	 * @throws Exception
	 */
	@Test
	public void test2() throws Exception{
        
		//1.读取xml文档，返回Document对象
		SAXReader reader = new SAXReader();
		Document doc = reader.read(new File("./src/contact.xml"));
		
		//得到根标签
		Element rooElem = doc.getRootElement();
		
		getChildNodes(rooElem);
	}
	
	/**
	 * 获取 传入的标签下的所有子节点
	 * @param elem
	 */
	private void getChildNodes(Element elem){
        
		System.out.println(elem.getName());
		
		//得到子节点
		Iterator<Node> it = elem.nodeIterator();
        
		while(it.hasNext()){
            
			Node node = it.next();
			
			//1.判断是否是标签节点
			if(node instanceof Element){
				Element el = (Element)node;
				//递归
				getChildNodes(el);
			}
		}
	}
}
```



##### 标签

​		`Element  Document.getRootElement()`   获取xml文档的根标签		

​		`Element  ELement.element("标签名") `    指定名称的第一个子标签

​		`Iterator<Element> Element.elementIterator("标签名")`指定名称的所有子标签

​		`List<Element>	 Element.elements()` 获取所有子标签

```java
import java.io.File;
import java.util.ArrayList;
import java.util.Iterator;
import java.util.List;

import org.dom4j.Document;
import org.dom4j.Element;
import org.dom4j.io.SAXReader;

public class Demo {
    
	/**
	 * 获取标签
	 */
	@Test
	public void test3() throws Exception{
        
		//1.读取xml文档，返回Document对象
		SAXReader reader = new SAXReader();
		Document doc = reader.read(new File("./src/contact.xml"));
		
		//2.得到根标签
		Element  rootElem = doc.getRootElement();
        
		//得到标签名称
		String name = rootElem.getName();
		System.out.println(name);
		
		//3.得到当前标签下指定名称的第一个子标签
		/*
		Element contactElem = rootElem.element("contact");
		System.out.println(contactElem.getName());
		*/
		
		//4.得到当前标签下指定名称的所有子标签
		/*
		Iterator<Element> it = rootElem.elementIterator("contact");
		while(it.hasNext()){
			Element elem = it.next();
			System.out.println(elem.getName());
		}
		*/
		
		//5.得到当前标签下的的所有子标签
		List<Element> list = rootElem.elements();
		//遍历List的方法
		//1)传统for循环  2）增强for循环 3）迭代器
		/*for(int i=0;i<list.size();i++){
			Element e = list.get(i);
			System.out.println(e.getName());
		}*/
		
	   /*for(Element e:list){
			System.out.println(e.getName());
		}*/
        
		/*
		Iterator<Element> it = list.iterator(); //ctrl+2 松开 l
		while(it.hasNext()){
			Element elem = it.next();
			System.out.println(elem.getName());
		}*/
		
		//获取更深层次的标签(方法只能一层层地获取)
		Element nameElem = doc.getRootElement().
					element("contact").element("name");
		System.out.println(nameElem.getName());
		
	}
}
```





##### 属性

​		`String  Element.attributeValue("属性名")` 获取指定名称的属性值

​		`Attribute   Element.attribute("属性名")` 获取指定名称的属性对象	

​				`attribute.getName() ` 获取属性名称

​				`Attibute.getValue()`  获取属性值

​		`List<Attribute>	 Element.attributes()`  获取所有属性对象

​		`Iterator<Attribute>	Element.attibuteIterator()`  获取所有属性对象

  ```java
public class Demo {
     
	/**
	 * 获取属性
	 */
	@Test
	public void test4() throws Exception{
        
		//1.读取xml文档，返回Document对象
		SAXReader reader = new SAXReader();
		Document doc = reader.read(new File("./src/contact.xml"));
		
		//获取属性：（先获的属性所在的标签对象，然后才能获取属性）
		//1.得到标签对象
		Element contactElem = doc.getRootElement().element("contact");
        
		//2.得到属性
		//2.1  得到指定名称的属性值
		/*
		String idValue = contactElem.attributeValue("id");
		System.out.println(idValue);
		*/
        
		
		//2.2 得到指定属性名称的属性对象
		/*Attribute idAttr = contactElem.attribute("id");
		
		//getName： 属性名称    getValue：属性值
		System.out.println(idAttr.getName() +"=" + idAttr.getValue());*/
        
		
		//2.3 得到所有属性对象,返回LIst集合
		/*List<Attribute> list = contactElem.attributes();
		//遍历属性
		for (Attribute attr : list) {
			System.out.println(attr.getName()+"="+attr.getValue());
		}*/
		
        
		//2.4 得到所有属性对象，返回迭代器
		Iterator<Attribute> it = contactElem.attributeIterator();
		while(it.hasNext()){
			Attribute attr = it.next();
			System.out.println(attr.getName()+"="+attr.getValue());
		}
	}
}
  ```



##### 文本

​		`Element.getText()` 获取当前标签的文本

​		`Element.elementText("标签名") `获取当前标签的指定名称的子标签的文本内容

```java
import java.io.File;
import java.util.ArrayList;
import java.util.Iterator;
import java.util.List;

import org.dom4j.Document;
import org.dom4j.Element;
import org.dom4j.io.SAXReader;

public class Demo {
   
	/**
	 * 获取文本
	 */
	@Test
	public void test5() throws Exception{
        
		//1.读取xml文档，返回Document对象
		SAXReader reader = new SAXReader();
		Document doc = reader.read(new File("./src/contact.xml"));
		
        
		/**
		 * 注意: 空格和换行也是xml的内容
		 */
		String content = doc.getRootElement().getText();
		System.out.println(content);
		
		
		//获取文本（先获取标签，再获取标签上的文本）
		Element nameELem = 
			doc.getRootElement().element("contact").element("name");
        
		//1. 得到文本
		String text = nameELem.getText();
		System.out.println(text);
		
		//2. 得到指定子标签名的文本内容
		String text2 = 
			doc.getRootElement().element("contact").elementText("phone");
		System.out.println(text2);
	}
}
```

#### 案例

##### 实体对象

```java
package gz.ixfosa.a_dom4j_read;

/**
 * 实体到对象
 *
 */
public class Contact {
	private String id;
	private String name;
	private String age;
	private String phone;
	private String email;
	private String qq;
	public String getId() {
		return id;
	}
	public void setId(String id) {
		this.id = id;
	}
	public String getName() {
		return name;
	}
	public void setName(String name) {
		this.name = name;
	}
	public String getAge() {
		return age;
	}
	public void setAge(String age) {
		this.age = age;
	}
	public String getPhone() {
		return phone;
	}
	public void setPhone(String phone) {
		this.phone = phone;
	}
	public String getEmail() {
		return email;
	}
	public void setEmail(String email) {
		this.email = email;
	}
	public String getQq() {
		return qq;
	}
	public void setQq(String qq) {
		this.qq = qq;
	}
	@Override
	public String toString() {
		return "Contact [age=" + age + ", email=" + email + ", id=" + id
				+ ", name=" + name + ", phone=" + phone + ", qq=" + qq + "]";
	}
}
```



##### 把xml文档信息封装到对象中

```java
package gz.ixfosa.a_dom4j_read;

import java.io.File;
import java.util.ArrayList;
import java.util.Iterator;
import java.util.List;

import org.dom4j.Document;
import org.dom4j.Element;
import org.dom4j.io.SAXReader;

/**
 * 把xml文档信息封装到对象中
 *
 */
public class Demo {
	public static void main(String[] args) throws Exception{
            
        List<Contact> list = new ArrayList<Contact>();

        //读取xml，封装对象
        SAXReader reader = new SAXReader();
        Document doc = reader.read(new File("./src/contact.xml"));

        //读取contact标签
        Iterator<Element> it = doc.getRootElement().elementIterator("contact");

        while(it.hasNext()){
            Element elem = it.next();
            //创建Contact
            Contact contact = new Contact();
            contact.setId(elem.attributeValue("id"));
            contact.setName(elem.elementText("name"));
            contact.setAge(elem.elementText("age"));
            contact.setPhone(elem.elementText("phone"));
            contact.setEmail(elem.elementText("email"));
            contact.setQq(elem.elementText("qq"));
            list.add(contact);
        }

        for (Contact contact : list) {
            System.out.println(contact);
        }
    }
}
```

##### 完整读取xml文档内容

```java
package gz.ixfosa.a_dom4j_read;

import java.io.File;
import java.util.Iterator;
import java.util.List;

import org.dom4j.Attribute;
import org.dom4j.Document;
import org.dom4j.Element;
import org.dom4j.Node;
import org.dom4j.Text;
import org.dom4j.io.SAXReader;
import org.junit.Test;

/**
 * 练习-完整读取xml文档内容
 *
 */
public class Demo {

	@Test
	public void test() throws Exception{
        
		//读取xml文档
		SAXReader reader = new SAXReader();
		Document doc = reader.read(new File("./src/contact.xml"));
				
		//读取根标签
		Element rootELem = doc.getRootElement();
		
		StringBuffer sb = new StringBuffer();

		getChildNodes(rootELem,sb);
		
		System.out.println(sb.toString());
		
	}
	
	/**
	 * 获取当前标签的所有子标签
	 */
	private void getChildNodes(Element elem,StringBuffer sb){
        
		//System.out.println(elem.getName());
		
		//开始标签
		sb.append("<"+elem.getName());
		
		//得到标签的属性列表
		List<Attribute> attrs = elem.attributes();
		if(attrs!=null){
			for (Attribute attr : attrs) {
				//System.out.println(attr.getName()+"="+attr.getValue());
				sb.append(" "+attr.getName()+"=\""+attr.getValue()+"\"");
			}
		}
		sb.append(">");
		
		//得到文本
		//String content = elem.getText();
		//System.out.println(content);
		Iterator<Node> it = elem.nodeIterator();
		while(it.hasNext()){
            
			Node node = it.next();
			
			//标签
			if(node instanceof Element){
				Element el = (Element)node;
				getChildNodes(el,sb);
			}
			
			//文本
			if(node instanceof Text){
				Text text = (Text)node;
				sb.append(text.getText());
			}
		}
		
		//结束标签 
		sb.append("</"+elem.getName()+">");
	}
}
```



#### Dom4j修改xml文档

##### 写出内容到xml文档

​		`XMLWriter writer = new XMLWriter(OutputStream, OutputForamt)`

​				`OutputForamt.createCompactFormat() `  紧凑的格式.去除空格换行.项目上线时
​				`OutputForamt.createPrettyPrint()`   漂亮的格式.有空格和换行.开发调试时

​		`wirter.write(Document)`

```java
import java.io.File;
import java.io.FileOutputStream;

import org.dom4j.Document;
import org.dom4j.io.SAXReader;
import org.dom4j.io.XMLWriter;

/**
 * 第一个写出内容到xml文档
 *
 */
public class Demo {
	public static void main(String[] args) throws Exception{
        
		//一、读取或创建一个Document对象
		//读取项目 下的xm；（/src/contact.xml）文件
		Document doc = new SAXReader().read(new File("./src/contact.xml"));
		
		
		//二、修改Document对象内容
		
		
		//三、把修改后的Document对象写出到xml文档中
		//指定文件输出的位置
		FileOutputStream out = new FileOutputStream("e:/contact.xml");
        
		//1.创建写出对象
		XMLWriter writer = new XMLWriter(out);
		
		//2.写出对象
		writer.write(doc);
        
		//3.关闭流
		writer.close();
	}
}
```

```java
import java.io.File;
import java.io.FileOutputStream;

import org.dom4j.Document;
import org.dom4j.io.OutputFormat;
import org.dom4j.io.SAXReader;
import org.dom4j.io.XMLWriter;
/**
 * 讨论写出内容到xml文档的细节
 *
 */
public class Demo {

	/**
	 * @param args
	 */
	public static void main(String[] args) throws Exception{
        
		Document doc = new SAXReader().read(new File("./src/contact.xml"));
        
		//指定文件输出的位置
		FileOutputStream out = new FileOutputStream("e:/contact.xml");
        
		/**
		 * 1.指定写出的格式
		 */
		OutputFormat format = OutputFormat.createCompactFormat(); //紧凑的格式
		//OutputFormat format = OutputFormat.createPrettyPrint(); //漂亮的格式
        
        
		/**
		 * 2.指定生成的xml文档的编码
		 *    同时影响了xml文档保存时的编码  和  xml文档声明的encoding的编码（xml解析时的编码）
		 *    结论： 使用该方法生成的xml文档避免中文乱码问题。
		 */
		format.setEncoding("utf-8");
		
        
		//1.创建写出对象
		XMLWriter writer = new XMLWriter(out,format);
		
		//2.写出对象
		writer.write(doc);
        
		//3.关闭流
		writer.close();
	}
}
```



##### 修改xml文档的API

###### 增加： 

​		`DocumentHelper.createDocument()`  增加文档
​		`addElement("名称")`  增加标签
​		`addAttribute("名称"，“值”)`  增加属性

```java
import java.io.File;
import java.io.FileOutputStream;

import org.dom4j.Attribute;
import org.dom4j.Document;
import org.dom4j.DocumentHelper;
import org.dom4j.Element;
import org.dom4j.io.OutputFormat;
import org.dom4j.io.SAXReader;
import org.dom4j.io.XMLWriter;
import org.junit.Test;

/**
 * 修改xml内容
 * 增加：文档，标签 ，属性
 *
 */
public class Demo {

	/**
	 * 增加：文档，标签 ，属性
	 */
	@Test
	public void test1() throws Exception{
        
		/**
		 * 1.创建文档
		 */
		Document doc = DocumentHelper.createDocument();
        
		/**
		 * 2.增加标签
		 */
		Element rootElem = doc.addElement("contactList");
		Element contactElem = rootElem.addElement("contact");
		contactElem.addElement("name");
        
		/**
		 * 3.增加属性
		 */
		contactElem.addAttribute("id", "001");
		contactElem.addAttribute("name", "eric");
		
		//把修改后的Document对象写出到xml文档中
		FileOutputStream out = new FileOutputStream("e:/contact.xml");
        
		OutputFormat format = OutputFormat.createPrettyPrint();
		format.setEncoding("utf-8");
        
		XMLWriter writer = new XMLWriter(out,format);
		writer.write(doc);
		writer.close();
	}
}
```



###### 修改：

​		`Attribute.setValue("值")`  修改属性值
​		`Element.addAtribute("同名的属性名","值")`  修改同名的属性值
​		`Element.setText("内容")`  修改文本内容

```java
import java.io.File;
import java.io.FileOutputStream;

import org.dom4j.Attribute;
import org.dom4j.Document;
import org.dom4j.DocumentHelper;
import org.dom4j.Element;
import org.dom4j.io.OutputFormat;
import org.dom4j.io.SAXReader;
import org.dom4j.io.XMLWriter;
import org.junit.Test;

/**
 * 修改xml内容
 * 修改：属性值，文本
 *
 */
public class Demo {

	/**
	 * 修改：属性值，文本
	 * @throws Exception
	 */
	@Test
	public void test2()	throws Exception{
        
		Document doc = new SAXReader().read(new File("./src/contact.xml"));
		
		/**
		 * 方案一： 修改属性值   1.得到标签对象 2.得到属性对象 3.修改属性值
		 */
        
		//1.1  得到标签对象
		/*
		Element contactElem = doc.getRootElement().element("contact");
		//1.2 得到属性对象
		Attribute idAttr = contactElem.attribute("id");
		//1.3 修改属性值
		idAttr.setValue("003");
		*/
        
		/**
		 * 方案二： 修改属性值
		 */
		//1.1 得到标签对象
		/*
		Element contactElem = doc.getRootElement().element("contact");
		//1.2 通过增加同名属性的方法，修改属性值
		contactElem.addAttribute("id", "004");
		*/
		
		/**
		 * 修改文本 1.得到标签对象 2.修改文本
		 */
		Element nameElem = doc.getRootElement().element("contact").element("name");
		nameElem.setText("李四");
		
        
		FileOutputStream out = new FileOutputStream("e:/contact.xml");
      
		OutputFormat format = OutputFormat.createPrettyPrint();
		format.setEncoding("utf-8");
        
		XMLWriter writer = new XMLWriter(out,format);
		writer.write(doc);
		writer.close();
	}
}
```



###### 删除

​		`Element.detach();`  删除标签  
​		`Attribute.detach(); ` 删除属性

```java
import java.io.File;
import java.io.FileOutputStream;

import org.dom4j.Attribute;
import org.dom4j.Document;
import org.dom4j.DocumentHelper;
import org.dom4j.Element;
import org.dom4j.io.OutputFormat;
import org.dom4j.io.SAXReader;
import org.dom4j.io.XMLWriter;
import org.junit.Test;

/**
 * 修改xml内容
 * 删除：标签，属性
 *
 */
public class Demo {

	/**
	 * 删除：标签，属性
	 * @throws Exception
	 */
	@Test
	public void test3() throws Exception{
        
		Document doc = new SAXReader().read(new File("./src/contact.xml"));
		
		/**
		 * 1.删除标签     1.1 得到标签对象  1.2 删除标签对象    
		 */
        
		// 1.1 得到标签对象
		/*
		Element ageElem = doc.getRootElement().element("contact").element("age");
					
		//1.2 删除标签对象
		ageElem.detach();
		//ageElem.getParent().remove(ageElem);
		*/
        
        
        
		/**
		 * 2.删除属性   2.1得到属性对象  2.2 删除属性
		 */
        
		//2.1得到属性对象
		//得到第二个contact标签
		Element contactElem = (Element)doc.getRootElement().elements().get(1);
			
		//2.2 得到属性对象
		Attribute idAttr = contactElem.attribute("id");
        
		//2.3 删除属性
		idAttr.detach();
		//idAttr.getParent().remove(idAttr);
		
        
		FileOutputStream out = new FileOutputStream("e:/contact.xml");
     
		OutputFormat format = OutputFormat.createPrettyPrint();
		format.setEncoding("utf-8");
        
		XMLWriter writer = new XMLWriter(out,format);
		writer.write(doc);
		writer.close();
	}
}
```

###### 练习

```xml

1.使用dom4j的api来生成以下的xml文件
<Students>
    <Student id="1">
        <name>张三</name>
        <gender>男</gender>
        <grade>计算机1班</grade>
        <address>广州天河</address>
    </Student>
    <Student id="2">
        <name>李四</name>
        <gender>女</gender>
        <grade>计算机2班</grade>
        <address>广州越秀</address>
    </Student>
</Students>


2.修改id为2的学生的姓名，改为“王丽”

3.删除id为2的学生

```





```java
import java.io.File;
import java.io.FileOutputStream;
import java.util.Iterator;

import org.dom4j.Document;
import org.dom4j.DocumentHelper;
import org.dom4j.Element;
import org.dom4j.io.OutputFormat;
import org.dom4j.io.SAXReader;
import org.dom4j.io.XMLWriter;
import org.junit.Test;

public class Demo {

	/**
	 * 1.生成指定xml文档
	 * @throws Exception
	 */
	@Test
	public void test1() throws Exception{
        
		//1.内存创建xml文档
		Document doc = DocumentHelper.createDocument();
		
		//2.写入内容
		Element rootElem = doc.addElement("Students");
		
		//2.1 增加标签
		Element studentElem1 = rootElem.addElement("Student");
		//2.2 增加属性
		studentElem1.addAttribute("id", "1");
		//2.3 增加标签，同时设置文本
		studentElem1.addElement("name").setText("张三");
		studentElem1.addElement("gender").setText("男");
		studentElem1.addElement("grade").setText("计算机1班");
		studentElem1.addElement("address").setText("广州天河");
		
		//2.1 增加标签
		Element studentElem2 = rootElem.addElement("Student");
		//2.2 增加属性
		studentElem2.addAttribute("id", "2");
		//2.3 增加标签，同时设置文本
		studentElem2.addElement("name").setText("李四");
		studentElem2.addElement("gender").setText("女");
		studentElem2.addElement("grade").setText("计算机2班");
		studentElem2.addElement("address").setText("广州越秀");
		
		
		//3.内容写出到xml文件
		//3.1 输出位置
		FileOutputStream out = new FileOutputStream("e:/student.xml");
		//3.2 指定格式
		OutputFormat format = OutputFormat.createPrettyPrint();
		// 设置编码
		format.setEncoding("utf-8");
        
		XMLWriter writer = new XMLWriter(out,format);
		//3.3 写出内容
		writer.write(doc);
		//3.4关闭资源
		writer.close();
		
	}
	
	/**
	 * 2.修改id为2的学生姓名
	 * @throws Exception
	 */
	@Test
	public void test2() throws Exception{
		//1.查询到id为2的学生
		Document doc = new SAXReader().read(new File("e:/student.xml"));
        
		//1.1 找到所有的Student标签
		Iterator<Element> it = doc.getRootElement().elementIterator("Student");
		while(it.hasNext()){
			Element stuElem = it.next();
			//1.2 查询id为id的学生标签
			if(stuElem.attributeValue("id").equals("2")){
				stuElem.element("name").setText("王丽");
				break;
			}
		}
	
		
		//3.1 输出位置
		FileOutputStream out = new FileOutputStream("e:/student.xml");
		//3.2 指定格式
		OutputFormat format = OutputFormat.createPrettyPrint();
		// 设置编码
		format.setEncoding("utf-8");
		XMLWriter writer = new XMLWriter(out,format);
		//3.3 写出内容
		writer.write(doc);
		//3.4关闭资源
		writer.close();
	}
	
	/**
	 * 3.删除id为2的学生
	 * @throws Exception
	 */
	@Test
	public void test3() throws Exception{
        
		//1.查询到id为2的学生
		Document doc = new SAXReader().read(new File("e:/student.xml"));
        
		//1.1 找到所有的Student标签
		Iterator<Element> it = doc.getRootElement().elementIterator("Student");
        
		while(it.hasNext()){
			Element stuElem = it.next();
			//1.2 查询id为id的学生标签
			if(stuElem.attributeValue("id").equals("2")){
				//1.3 删除该学生标签
				stuElem.detach();
				break;
			}
		}

		
		//3.1 输出位置
		FileOutputStream out = new FileOutputStream("e:/student.xml");
        
		//3.2 指定格式
		OutputFormat format = OutputFormat.createPrettyPrint();
        
		// 设置编码
		format.setEncoding("utf-8");
        
		XMLWriter writer = new XMLWriter(out,format);
        
		//3.3 写出内容
		writer.write(doc);
        
		//3.4关闭资源
		writer.close();
	}
}
```



#### xPath技术  

##### 引入

​		问题：当使用dom4j查询比较深的层次结构的节点（标签，属性，文本），比较麻烦！！！

​		xPath作用:主要是用于快速获取所需的节点对象。



##### 在dom4j中使用xPath技术

​		1）导入xPath支持jar包 。  jaxen-1.1-beta-6.jar
​		2）使用xpath方法
​					`List<Node>  selectNodes("xpath表达式")`;   查询多个节点对象
​					`Node  selectSingleNode("xpath表达式")`;  查询一个节点对象

```java
import java.io.File;
import java.io.FileOutputStream;

import org.dom4j.Document;
import org.dom4j.Element;
import org.dom4j.io.OutputFormat;
import org.dom4j.io.SAXReader;
import org.dom4j.io.XMLWriter;

/**
 * 第一个xpath程序
 *
 */
public class Demo {
	public static void main(String[] args) throws Exception{
        
		/**
		 * 需求: 删除id值为2的学生标签
		 */
		Document doc = new SAXReader().read(new File("e:/student.xml"));
		
		//1.查询id为2的学生标签
		//使用xpath技术
		Element stuElem = (Element)doc.selectSingleNode("//Student[@id='2']");

		//2.删除标签
		stuElem.detach();
		
		//3.写出xml文件
		FileOutputStream out = new FileOutputStream("e:/student.xml");
		OutputFormat format = OutputFormat.createPrettyPrint();
		format.setEncoding("utf-8");
		XMLWriter writer = new XMLWriter(out,format);
		writer.write(doc);
		writer.close();
	}

}
```





##### xPath语法

​	    `/`      绝对路径       表示从xml的根位置开始或子元素（一个层次结构）
​        ` // `    相对路径       表示不分任何层次结构的选择元素。

​		`*`			    通配符         表示匹配所有元素
​		`[] `             条件             表示选择什么条件下的元素
​		`@`               属性             表示选择属性节点
​		`and `           关系             表示条件的与关系（等价于&&）
​		`text()`    文本             表示选择文本内容

```java
import java.io.File;
import java.util.List;

import org.dom4j.Document;
import org.dom4j.Node;
import org.dom4j.io.SAXReader;

/**
 * 学习xPath表达式语法
 *
 */
public class Demo {
	public static void main(String[] args) throws Exception {
        
		Document doc = new SAXReader().read(new File("./src/contact.xml"));
		
		String xpath = "";
		
		/**
		 * 1.  	/      绝对路径  表示从xml的根位置开始或子元素（一个层次结构）
		 */
		xpath = "/contactList";
		xpath = "/contactList/contact";
		
		/**
		 * 2. //     相对路径   表示不分任何层次结构的选择元素。
		 */
		xpath = "//contact/name";
		xpath = "//name";
		
        
		/**
		 * 3. *      通配符   表示匹配所有元素
		 */
		xpath = "/contactList/*"; //根标签contactList下的所有子标签
		xpath = "/contactList//*";//根标签contactList下的所有标签（不分层次结构）
		
		/**
		 * 4. []      条件           表示选择什么条件下的元素
		 */
		//带有id属性的contact标签
		xpath = "//contact[@id]";
		//第二个的contact标签
		xpath = "//contact[2]";
		//选择最后一个contact标签
		xpath = "//contact[last()]";
		
		/**
		 * 5. @        属性            表示选择属性节点
		 */
		xpath = "//@id"; //选择id属性节点对象，返回的是Attribute对象
		xpath = "//contact[not(@id)]";//选择不包含id属性的 contact标签节点
		xpath = "//contact[@id='002']";//选择id属性值为002的contact标签
		xpath = "//contact[@id='001' and @name='eric']";//选择id属性值为001，且name属性为eric的contact标签
		
		/**
		 *6.  text()   表示选择文本内容
		 */
		//选择name标签下的文本内容，返回Text对象
		xpath = "//name/text()";
		xpath = "//contact/name[text()='张三']";//选择姓名为张三的name标签
		
		
		List<Node> list = doc.selectNodes(xpath);
		for (Node node : list) {
			System.out.println(node);
		}
	}
}
```



##### 案例

###### 		用户登录功能：

​			用户输入用户名和密码 -> 到“数据库”查询是否有对应的用户 -> 
​					有： 则表示登录成功
​					没有： 则表示登录失败

​			用xml当做数据库

​						user.xml  用来存储用户的数据

**user.xml**

```xml
<?xml version="1.0" encoding="utf-8"?>
<users>
	<user id="001" name="eric" password="123456"></user>
	<user id="002" name="rose" password="123456"></user>
	<user id="003" name="jack" password="123456"></user>
</users>
```



```java
import java.io.BufferedReader;
import java.io.File;
import java.io.InputStreamReader;

import org.dom4j.Document;
import org.dom4j.Element;
import org.dom4j.io.SAXReader;

/**
 * xpath案例： 模拟用户登录效果
 *
 */
public class Demo {
	public static void main(String[] args)throws Exception{
        
		//1.获取用户输入的用户名和密码
		BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
				
		System.out.println("请输入用户名：");
		String name = br.readLine();
		
		System.out.println("请输入密码：");
		String password = br.readLine();
		
		//2.到“数据库”中查询是否有对应的用户
		//对应的用户：  在user.xml文件中找到一个
			//name属性值为‘用户输入’，且password属性值为‘用户输入’的user标签
		Document doc = new SAXReader().read(new File("./src/user.xml"));
        
		Element userElem = (Element)doc.selectSingleNode("//user[@name='" +name +"' and @password='"+password+"']");
		
		if(userElem!=null){
			//登录成功
			System.out.println("登录成功");
		}else{
			//登录失败
			System.out.println("登录失败");
		}
	}
}
```



###### xpath读取html

**/src/personList.html**

```html
<html>
	<head>
		<title>通讯录</title>
		<meta http-equiv="content-type" content="text/html; charset=UTF-8" />
	</head>
	<body>
		<center><h1>通讯录</h1></center>
		<table border="1" align="center" id="contactForm">
			<thead>	
<tr><th>编号</th><th>姓名</th><th>性别</th><th>年龄</th><th>地址</th><th>电话</th></tr>
			</thead>
			<tbody>
				<tr>
                    <td>001</td>
                    <td>张三</td>
                    <td>男</td>
                    <td>18</td>
                    <td>广州市天河区</td>
                    <td>134000000000</td>
				</tr>
				<tr>
                    <td>002</td>
                    <td>李四</td>
                    <td>女</td>
                    <td>20</td>
                    <td>广州市越秀区</td>
                    <td>13888888888</td>
				</tr>
			</tbody>
		</table>
	</body>
</html>
```

```java
import java.io.File;
import java.util.Iterator;
import java.util.List;

import org.dom4j.Document;
import org.dom4j.Element;
import org.dom4j.io.SAXReader;

/**
 * 使用xpath技术读取一个规范的html文档
 *
 */
public class Demo {
	public static void main(String[] args) throws Exception{
        
		Document doc = new SAXReader().read(new File("./src/personList.html"));
		
		//读取title标签
		Element titleElem = (Element)doc.selectSingleNode("//title");
		String title = titleElem.getText();
		System.out.println(title);
		
		/**
		 * 练习：读取联系人的所有信息
		 * 按照以下格式输出：
		 * 		 编号:001 姓名:张三 性别:男 年龄:18 地址：xxxx 电话： xxxx
		 *       编号:002 姓名:李四 性别:女 年龄:20 地址：xxxx 电话： xxxx
		 *       ......
		 */
		//1.读取出所有tbody中的tr标签
		List<Element> list = (List<Element>)doc.selectNodes("//tbody/tr");
        
		//2.遍历
		for (Element elem : list) {
			//编号
			//String id = ((Element)elem.elements().get(0)).getText();
			String id = elem.selectSingleNode("td[1]").getText();
			//姓名
			String name = ((Element)elem.elements().get(1)).getText();
			//性别
			String gender = ((Element)elem.elements().get(2)).getText();
			//年龄
			String age = ((Element)elem.elements().get(3)).getText();
			//地址
			String address = ((Element)elem.elements().get(4)).getText();
			//电话
			String phone = ((Element)elem.elements().get(5)).getText();
			
			System.out.println("编号："+id+"\t姓名："+name+"\t性别："+
								gender+"\t年龄："+
								age+"\t地址："+address+
								"\t电话："+phone);
		}
	}
}
```



### SAX解析

#### 回顾DOM解析

​		`DOM解析`原理：一次性把xml文档加载进内存，然后在内存中构建Document树,对内存要求比较高。不适合读取大容量的xml文件，容易导致内存溢出。						

​		`SAX解析`原理： 加载一点，读取一点，处理一点。对内存要求比较低。

 

#### SAX解析工具

​		SAX解析工具-  Sun公司提供的。内置在jdk中。org.xml.sax.*

​		核心的API：

​				SAXParser类： 用于读取和解析xml文件对象

​						parse（[File](mk:@MSITStore:E:\study\API帮助文档\jdk6.chm::/j2se6/api/java/io/File.html) f, [DefaultHandler](mk:@MSITStore:E:\study\API帮助文档\jdk6.chm::/j2se6/api/org/xml/sax/helpers/DefaultHandler.html) dh）方法： 解析xml文件

​								参数一： File：表示 读取的xml文件。

​					  		  参数二： DefaultHandler： SAX事件处理程序。使用DefaultHandler的子类



**DOM解析   vs  SAX解析**

| DOM解析                                                      | SAX解析                                                   |
| ------------------------------------------------------------ | --------------------------------------------------------- |
| 原理： 一次性加载xml文档，不适合大容量的文件读取             | 原理： 加载一点，读取一点，处理一点。适合大容量文件的读取 |
| DOM解析可以任意进行增删改成                                  | SAX解析只能读取                                           |
| DOM解析任意读取任何位置的数据，甚至往回读                    | SAX解析只能从上往下，按顺序读取，不能往回读               |
| DOM解析面向对象的编程方法（Node，Element，Attribute）,Java开发者编码比较简单。 | SAX解析基于事件的编程方法。java开发编码相对复杂。         |

**/src/contact.xml**

```xml
<?xml version="1.0" encoding="utf-8"?>
<contactList>
	<contact id="001" name="eric">
		<name>张三</name>
		<age>20</age>
		<phone>134222223333</phone>
		<email>zhangsan@qq.com</email>
		<qq>432221111</qq>
	</contact>
	<contact id="002" name="jacky">
		<name>eric</name>
		<age>20</age>
		<phone>134222225555</phone>
		<email>lisi@qq.com</email>
		<qq>432222222</qq>
	</contact>
</contactList>
```



```java
import java.io.File;

import javax.xml.parsers.SAXParser;
import javax.xml.parsers.SAXParserFactory;

/**
 * 第一个SAX读取xml文件程序
 *
 */
public class Demo {
	public static void main(String[] args) throws Exception{
        
		//1.创建SAXParser对象
		SAXParser parser = SAXParserFactory.newInstance().newSAXParser();
		
		//2.调用parse方法
		/**
		 * 参数一： xml文档
		 * 参数二： DefaultHandler的子类
		 */
		parser.parse(new File("./src/contact.xml"), new MyDefaultHandler());
	}
}
```

一个类继承class 类名（extends `DefaultHandler`） 在调用是创建传进去

```java
import org.xml.sax.Attributes;
import org.xml.sax.SAXException;
import org.xml.sax.helpers.DefaultHandler;
/**
 * SAX处理程序（如何解析xml文档）
 *
 */
public class MyDefaultHandler extends DefaultHandler {
	
	/**
	 * 开始文档时调用
	 */
	@Override
	public void startDocument() throws SAXException {
		System.out.println("MyDefaultHandler.startDocument()");
	}
	
	/**
	 * 开始标签时调用
	 * @param qName: 表示开始标签的标签名
	 * @param attributes: 表示开始标签内包含的属性列表
	 */
	@Override
	public void startElement(String uri, String localName, String qName,
			Attributes attributes) throws SAXException {
		System.out.println("MyDefaultHandler.startElement()-->"+qName);
	}
	
	/**
	 * 结束标签时调用
	 * @param qName: 结束标签的标签名称
	 */
	@Override
	public void endElement(String uri, String localName, String qName)
			throws SAXException {
		System.out.println("MyDefaultHandler.endElement()-->"+qName);
	}
	
	/**
	 * 读到文本内容的时调用
	 * @param ch: 表示当前读完的所有文本内容
	 * @param start： 表示当前文本内容的开始位置
	 * @param length: 表示当前文本内容的长度
	 * char[]（        张三 20 ...）    
	 */ 
	@Override
	public void characters(char[] ch, int start, int length)
			throws SAXException {
		//得到当前文本内容
		String content = new String(ch,start,length);
		System.out.println("MyDefaultHandler.characters()-->"+content);
	}
	
	/**
	 * 结束文档时调用
	 */
	@Override
	public void endDocument() throws SAXException {
		System.out.println("MyDefaultHandler.endDocument()");
	}
}
```





#### 完整输出文档内容

DefaultHandler类的API:

​		void startDocument()  :  在读到文档开始时调用

​		void endDocument()  ：在读到文档结束时调用

​		void startElement(String uri, String localName, String qName, Attributes attributes)  ：读到开始标签时调用				

​		void endElement(String uri, String localName, String qName)  ：读到结束标签时调用

​		void characters(char[] ch, int start, int length)  ： 读到文本内容时调用

```java
import java.io.File;

import javax.xml.parsers.SAXParser;
import javax.xml.parsers.SAXParserFactory;

/**
 * 读取contact.xml文件，完整输出文档内容
 *
 */
public class Demo2 {
	public static void main(String[] args)throws Exception {
        
		//1.创建SAXParser
		SAXParser parser = SAXParserFactory.newInstance().newSAXParser();
        
		//2.读取xml文件
		MyDefaultHandler2 handler = new MyDefaultHandler2();
        
		parser.parse(new File("./src/contact.xml"), handler);
        
		String content = handler.getContent();
		System.out.println(content);
	}
}
```

**MyDefaultHandler2**

```java
import org.xml.sax.Attributes;
import org.xml.sax.SAXException;
import org.xml.sax.helpers.DefaultHandler;

/**
 * SAX处理器程序
 */
public class MyDefaultHandler2 extends DefaultHandler {
    
	//存储xml文档信息
	private StringBuffer sb = new StringBuffer();
	
	//获取xml信息
	public String getContent(){
		return sb.toString();
	}
	
	/**
	 * 开始标签
	 */
	@Override
	public void startElement(String uri, String localName, String qName,
			Attributes attributes) throws SAXException {
		sb.append("<"+qName);
		//判断是否有属性
		if(attributes!=null){
			for(int i=0;i<attributes.getLength();i++){
				//得到属性名称
				String attrName = attributes.getQName(i);
				//得到属性值
				String attrValue = attributes.getValue(i);
				sb.append(" "+attrName+"=\""+attrValue+"\"");
			}
		}
		sb.append(">");
	}
	
	/**
	 * 文本内容
	 */
	@Override
	public void characters(char[] ch, int start, int length)
			throws SAXException {
		//得到当前读取的文本
		String content = new String(ch,start,length);
		sb.append(content);
	}
	
	/**
	 * 结束标签
	 */
	@Override
	public void endElement(String uri, String localName, String qName)
			throws SAXException {
		sb.append("</"+qName+">");
	}
}
```



#### 封装到实体对象

```java
public class Contact {
	private String id;
	private String name;
	private String age;
	private String phone;
	private String email;
	private String qq;
	public String getId() {
		return id;
	}
	public void setId(String id) {
		this.id = id;
	}
	public String getName() {
		return name;
	}
	public void setName(String name) {
		this.name = name;
	}
	public String getAge() {
		return age;
	}
	public void setAge(String age) {
		this.age = age;
	}
	public String getPhone() {
		return phone;
	}
	public void setPhone(String phone) {
		this.phone = phone;
	}
	public String getEmail() {
		return email;
	}
	public void setEmail(String email) {
		this.email = email;
	}
	public String getQq() {
		return qq;
	}
	public void setQq(String qq) {
		this.qq = qq;
	}
	@Override
	public String toString() {
		return "Contact [age=" + age + ", email=" + email + ", id=" + id
				+ ", name=" + name + ", phone=" + phone + ", qq=" + qq + "]";
	}
	
}
```

**MyDefaultHandler3**

```java
import java.util.ArrayList;
import java.util.List;

import org.xml.sax.Attributes;
import org.xml.sax.SAXException;
import org.xml.sax.helpers.DefaultHandler;

/**
 * SAX处理程序
 *
 */
public class MyDefaultHandler3 extends DefaultHandler {
	//存储所有联系人对象
	private List<Contact> list = new ArrayList<Contact>();
	
	public List<Contact> getList(){
		return list;
	}
    
	//保存一个联系人信息
	private Contact contact;
	/**
	 * 思路： 
	 * 	1）创建Contact对象
	 *  2）把每个contact标签内容存入到Contact对象
	 *  3）把Contact对象放入List中
	 */
	//用于临时存储当前读到的标签名
	private String curTag;

	@Override
	public void startElement(String uri, String localName, String qName,
			Attributes attributes) throws SAXException {
		curTag = qName;
		//读取到contact的开始标签创建Contact对象
		if("contact".equals(qName)){
			contact = new Contact();
			
			//设置id值
			contact.setId(attributes.getValue("id"));
		}
	}
	
	@Override
	public void characters(char[] ch, int start, int length)
			throws SAXException {
		//当前文本内容
		String content = new String(ch,start,length);
		
		if("name".equals(curTag)){
			contact.setName(content);
		}
		
		if("age".equals(curTag)){
			contact.setAge(content);
		}
		
		if("phone".equals(curTag)){
			contact.setPhone(content);
		}
		
		if("email".equals(curTag)){
			contact.setEmail(content);
		}
		
		if("qq".equals(curTag)){
			contact.setQq(content);
		}
	}
	
	@Override
	public void endElement(String uri, String localName, String qName)
			throws SAXException {
		//设置空时为了避免空格换行设置到对象的属性中
		curTag = null;
		//读到contact的结束标签放入List中
		if("contact".equals(qName)){
			list.add(contact);
		}
	}
}
```



```java
import java.io.File;

import javax.xml.parsers.SAXParser;
import javax.xml.parsers.SAXParserFactory;

/**
 * 读取contact.xml文件，封装到实体对象
 *
 */
public class Demo3 {
	public static void main(String[] args)throws Exception {
        
		//1.创建SAXParser
		SAXParser parser = SAXParserFactory.newInstance().newSAXParser();
        
		//2.读取xml文件
		MyDefaultHandler3 handler = new MyDefaultHandler3();
        
		parser.parse(new File("./src/contact.xml"), handler);
        
		List<Contact> list = handler.getLisr();
		System.out.println(list);
	}
}
```



### 案例-通讯录-增删改查

#### 所需 jar 文件

dom4j-1.6.1.jar

jaxen-1.1-beta-6.jar



#### 联系人实体对象

```java
package gz.ixfosa.contact;
/**
 * 联系人实体对象
 *
 */
public class Contact {

	private String id;
	private String name;
	private String gender;ixfosa
	private int age;
	private String phone;
	private String email;
	private String qq;
	public String getId() {
		return id;
	}
	public void setId(String id) {
		this.id = id;
	}
	public String getName() {
		return name;
	}
	public void setName(String name) {
		this.name = name;
	}
	public String getGender() {
		return gender;
	}
	public void setGender(String gender) {
		this.gender = gender;
	}
	public int getAge() {
		return age;
	}
	public void setAge(int age) {
		this.age = age;
	}
	public String getPhone() {
		return phone;
	}
	public void setPhone(String phone) {
		this.phone = phone;
	}
	public String getEmail() {
		return email;
	}
	public void setEmail(String email) {
		this.email = email;
	}
	public String getQq() {
		return qq;
	}
	public void setQq(String qq) {
		this.qq = qq;
	}
	@Override
	public String toString() {
		return "Contact [age=" + age + ", email=" + email + ", gender="
				+ gender + ", id=" + id + ", name=" + name + ", phone=" + phone
				+ ", qq=" + qq + "]";
	}
}
```

#### xml操作的工具类

```java
package gz.ixfosa.contact;

import java.io.File;
import java.io.FileNotFoundException;
import java.io.FileOutputStream;
import java.io.IOException;
import java.io.UnsupportedEncodingException;

import org.dom4j.Document;
import org.dom4j.DocumentException;
import org.dom4j.io.OutputFormat;
import org.dom4j.io.SAXReader;
import org.dom4j.io.XMLWriter;

/**
 * xml操作的工具类
 *
 */
public class XMLUtil {
	
	/**
	 * 用于读取xml文件
	 * @return
	 */
	public static Document getDocument(){
		try {
			Document doc = new SAXReader().read(new File("e:/contact.xml"));
			return doc;
		} catch (DocumentException e) {
			e.printStackTrace();
			throw new RuntimeException(e);
		}
	}
	

	/**
	 * 写出xml文件
	 */
	public static void write2xml(Document doc){
		try {
			FileOutputStream out = new FileOutputStream("e:/contact.xml");
			OutputFormat format = OutputFormat.createPrettyPrint();
			format.setEncoding("utf-8");
			XMLWriter writer = new XMLWriter(out,format);
			writer.write(doc);
			writer.close();
		} catch (Exception e) {
			e.printStackTrace();
			throw new RuntimeException(e);
		}
	}
}
```

#### 定义接口存放联系人相关操作

```java
package gz.ixfosa.contact;

import java.util.List;

/**
 * 该接口用于存放联系人相关操作的方法。
 *
 */
public interface ContactOperator {
	public void addContact(Contact contact);//增加联系人
	public void updateContact(Contact contact);//修改联系人
	public void deleteContact(String id);//根据ID删除联系人
	public List<Contact> findAll();//查询所有联系人
}
```

#### 实现联系人操作接口

```java
package gz.ixfosa.contact;

import java.io.File;
import java.util.ArrayList;
import java.util.List;

import org.dom4j.Document;
import org.dom4j.DocumentHelper;
import org.dom4j.Element;
/**
 * 实现联系人操作接口
 *
 */
public class ContactOperatorImpl implements ContactOperator {

	/**
	 * 添加联系人
	 */
	@Override
	public void addContact(Contact contact) {
		
		try {
			File file = new File("e:/contact.xml");
			Document doc = null;
			Element rootElem = null;
			if(!file.exists()){
				/**
				 * 需求： 把contact对象保存到xml文件中
				 */
				//如果没有xml文件，则创建xml文件
				doc = DocumentHelper.createDocument();
				//创建根标签
				rootElem = doc.addElement("contactList");
			}else{
				//如果有xml文件，则读取xml文件
				doc = XMLUtil.getDocument();
				//如果有xml文件，读取根标签
				rootElem = doc.getRootElement();
			}

			//添加contact标签
			/**
			 * <contact id="1">
					<name>eric</name>
					<gender>男</gender>
					<age>20</age>
					<phone>1343333</phone>
					<email>eric@qq.com</email>
					<qq>554444</qq>
				</contact>
			 */
			Element contactElem = rootElem.addElement("contact");
			contactElem.addAttribute("id", contact.getId());
			contactElem.addElement("name").setText(contact.getName());
			contactElem.addElement("gender").setText(contact.getGender());
			contactElem.addElement("age").setText(contact.getAge()+"");
			contactElem.addElement("phone").setText(contact.getPhone());
			contactElem.addElement("email").setText(contact.getEmail());
			contactElem.addElement("qq").setText(contact.getQq());
			
			//把Document写出到xml文件
			XMLUtil.write2xml(doc);
		} catch (Exception e) {
			e.printStackTrace();
			throw new RuntimeException(e);
		}
	}

	/**
	 * 修改联系人
	 */
	@Override
	public void updateContact(Contact contact) {
		/**
		 * 需求： 修改id值为2的联系人
		 * 	1）查询id值为2的contact标签
		 *  2）修改contact标签的内容
		 */
		try {
			//1.读取xml文件
			Document doc = XMLUtil.getDocument();
			
			Element contactElem = (Element)doc.selectSingleNode("//contact[@id='"+contact.getId()+"']");
			
			//2.修改contact标签内容
			contactElem.element("name").setText(contact.getName());
			contactElem.element("gender").setText(contact.getGender());
			contactElem.element("age").setText(contact.getAge()+"");
			contactElem.element("phone").setText(contact.getPhone());
			contactElem.element("email").setText(contact.getEmail());
			contactElem.element("qq").setText(contact.getQq());
			
			//3.把Document写出到xml文件
			XMLUtil.write2xml(doc);
		} catch (Exception e) {
			e.printStackTrace();
			throw new RuntimeException(e);
		}
	}
	
	/**
	 * 删除联系人
	 */
	@Override
	public void deleteContact(String id) {
		try {
			//1.读取xml文件
			Document doc = XMLUtil.getDocument();
			//2.查询需要删除id的contact
			Element contactElem = (Element)doc.selectSingleNode("//contact[@id='"+id+"']");
			//删除标签
			contactElem.detach();
			
			//3.把Document写出到xml文件
			XMLUtil.write2xml(doc);
		} catch (Exception e) {
			e.printStackTrace();
			throw new RuntimeException(e);
		}
	}

	/**
	 * 查询所有联系人
	 */
	@Override
	public List<Contact> findAll() {
		//1.读取xml文件
		Document doc = XMLUtil.getDocument();
		
		//2.创建List对象
		List<Contact> list = new ArrayList<Contact>();
		//3.读取contact标签
		List<Element> conList = (List<Element>)doc.selectNodes("//contact");
		for(Element e:conList){
			//创建COntact对象
			Contact c = new Contact();
			c.setId(e.attributeValue("id"));
			c.setName(e.elementText("name"));
			c.setGender(e.elementText("gender"));
			c.setAge(Integer.parseInt(e.elementText("age")));
			c.setPhone(e.elementText("phone"));
			c.setEmail(e.elementText("email"));
			c.setQq(e.elementText("qq"));
			//把Contact放入list中
			list.add(c);
		}
		return list;
	}
}
```

#### 主程序

```java
package gz.ixfosa.contact;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.List;

/**
 * 主程序
 *
 */
public class MainProgram {

	public static void main(String[] args) {
        
		try {
            
			/*Scanner scanner = new Scanner(System.in);
			String command = scanner.next();*/
			
			BufferedReader br = 
							new BufferedReader(new InputStreamReader(System.in));
			
			ContactOperator operator = new ContactOperatorImpl();
			//不断接收输入
			while(true){
                
				//1.看到菜单
				printMenu();
				
				//2.接收用户输入的命令
				String command = br.readLine();
				
				//3.根据用户的输入执行不同的操作
				if("1".equals(command)){
					//添加联系人
					//获取用输入的数据
					Contact contact  = new Contact();
					//获取ID
					System.out.println("请输入编号：");
					String id = br.readLine();
					//获取姓名
					System.out.println("请输入姓名：");
					String name = br.readLine();
					//获取性别
					System.out.println("请输入性别：");
					String gender = br.readLine();
					//获取年龄
					System.out.println("请输入年龄：");
					String age = br.readLine();
					//获取电话
					System.out.println("请输入电话：");
					String phone = br.readLine();
					//获取邮箱
					System.out.println("请输入邮箱：");
					String email = br.readLine();
					//获取QQ
					System.out.println("请输入QQ：");
					String qq = br.readLine();
					
					contact.setId(id);
					contact.setName(name);
					contact.setGender(gender);
					contact.setAge(Integer.parseInt(age));
					contact.setPhone(phone);
					contact.setEmail(email);
					contact.setQq(qq);
					
					operator.addContact(contact);
				}else if("2".equals(command)){
					//修改联系人
					//获取用输入的数据
					Contact contact  = new Contact();

					//获取ID
					System.out.println("请输入需要修改的编号：");
					String id = br.readLine();
					//获取姓名
					System.out.println("请输入修改后的姓名：");
					String name = br.readLine();
					//获取性别
					System.out.println("请输入修改后的性别：");
					String gender = br.readLine();
					//获取年龄
					System.out.println("请输入修改后的年龄：");
					String age = br.readLine();
					//获取电话
					System.out.println("请输入修改后的电话：");
					String phone = br.readLine();
					//获取邮箱
					System.out.println("请输入修改后的邮箱：");
					String email = br.readLine();
					//获取QQ
					System.out.println("请输入修改后的QQ：");
					String qq = br.readLine();
					
					contact.setId(id);
					contact.setName(name);
					contact.setGender(gender);
					contact.setAge(Integer.parseInt(age));
					contact.setPhone(phone);
					contact.setEmail(email);
					contact.setQq(qq);
					
					
					operator.updateContact(contact);
				}else if("3".equals(command)){
					//删除联系人
					//获取需要删除的ID
					System.out.println("请输入删除的编号：");
					String id = br.readLine();
					
					operator.deleteContact(id);
				}else if("4".equals(command)){
					//查询所有联系人
					List<Contact> list = operator.findAll();
					//打印输出
					for (Contact con : list) {
						System.out.println(con);
					}
				}else if("Q".equals(command)){
					break;
				}else{
					System.out.println("输入的命令有误！");
				}
			}
		} catch (IOException e) {
			e.printStackTrace();
			throw new  RuntimeException(e);
		}
			
		
	}

	private static void printMenu() {
		System.out.println("====================");
		System.out.println("[1]添加联系人");
		System.out.println("[2]修改联系人");
		System.out.println("[3]删除联系人");
		System.out.println("[4]查看所有联系人");
		System.out.println("[Q]退出系统");
		System.out.println("====================");
	}

}
```



## DbUtils-JDBC

​		commons-dbutils 是 Apache 组织提供的一个开源 JDBC工具类库，它是对JDBC的简单封装，学习成本极低，并且使用dbutils能极大简化jdbc编码的工作量，同时也不会影响程序的性能。

 

DbUtils组件，

1. 简化jdbc操作

2. 下载组件，引入jar文件 : `commons-dbutils-1.6.jar`

### Api

```java
|-- DbUtils   关闭资源、加载驱动
|-- QueryRunner   组件的核心工具类：定义了所有的与数据库操作的方法(查询、更新)
    	//执行更新带一个占位符的sql
		Int  update(Connection conn, String sql, Object param); 

		//执行更新带多个占位符的sql
		Int  update(Connection conn, String sql, Object…  param); 
			
		//批处理
		Int[]  batch(Connection conn, String sql, Object[][] params)      
           
        //查询方法
		T query(Connection conn ,String sql, ResultSetHandler<T> rsh, 
                Object... params)   	

		Int  update( String sql, Object param);  
		Int  update( String sql, Object…  param); 
		Int[]  batch( String sql, Object[][] params)    
            
	注意： 如果调用DbUtils组件的操作数据库方法，没有传入连接对象，那么在实例化QueryRunner对象的时	候需要传入数据源对象： QueryRunner qr = new QueryRunner(ds);
```

DbUtils提供的封装结果的一些对象：

1. BeanHandler: 查询返回单个对象
2. BeanListHandler: 查询返回list集合，集合元素是指定的对象
3. ArrayHandler, 查询返回结果记录的第一行，封装对对象数组, 即返回：Object[]
4. ArrayListHandler, 把查询的每一行都封装为对象数组，再添加到list集合中
5.  ScalarHandler 查询返回结果记录的第一行的第一列  (在聚合函数统计的时候用)
6.  MapHandler  查询返回结果的第一条记录封装为map



### 实例

```java
package cn.ixfosa.dbUtils;

import java.util.Date;

/**
 * 1. 实体类设计
 *
 */
public class Admin {

	private int id;
	private String userName;
	private String pwd;
	
	
	
	public String getPwd() {
		return pwd;
	}
	public void setPwd(String pwd) {
		this.pwd = pwd;
	}
	public int getId() {
		return id;
	}
	public void setId(int id) {
		this.id = id;
	}
	public String getUserName() {
		return userName;
	}
	public void setUserName(String userName) {
		this.userName = userName;
	}
	@Override
	public String toString() {
		return "Admin [id=" + id + ", pwd=" + pwd + ", userName=" + userName
				+ "]";
	}
}
```

```java
package cn.ixfosa.dbUtils;

import java.sql.Connection;
import java.sql.ResultSet;
import java.sql.SQLException;
import java.util.List;
import java.util.Map;

import org.apache.commons.dbutils.QueryRunner;
import org.apache.commons.dbutils.ResultSetHandler;
import org.apache.commons.dbutils.handlers.BeanHandler;
import org.apache.commons.dbutils.handlers.BeanListHandler;
import org.apache.commons.dbutils.handlers.MapHandler;
import org.junit.Test;

import cn.ixfosa.dbUtils.Admin;
import cn.ixfosa.utils.JdbcUtil;

public class App_query {

	private Connection conn;

	// 一、查询， 自定义结果集封装数据
	@Test
	public void testQuery() throws Exception {
        
		String sql = "select * from admin where id=?";
		// 获取连接
		conn = JdbcUtil.getConnection();
		// 创建DbUtils核心工具类对象
		QueryRunner qr = new QueryRunner();
		// 查询
		Admin admin = qr.query(conn, sql, new ResultSetHandler<Admin>() {

			// 如何封装一个Admin对象
			public Admin handle(ResultSet rs) throws SQLException {
				if (rs.next()) {
					Admin admin = new Admin();
					admin.setId(rs.getInt("id"));
					admin.setUserName(rs.getString("userName"));
					admin.setPwd(rs.getString("pwd"));
					return admin;
				}
				return null;
			}

		}, 29);

		// 测试
		System.out.println(admin);
		// 关闭
		conn.close();

	}
	
	// 二、查询， 使用组件提供的结果集对象封装数据
	
	// 1）BeanHandler: 查询返回单个对象
	@Test
	public void testQueryOne() throws Exception {
		String sql = "select * from admin where id=?";
		// 获取连接
		conn = JdbcUtil.getConnection();
		// 创建DbUtils核心工具类对象
		QueryRunner qr = new QueryRunner();
		// 查询返回单个对象
		Admin admin =  qr.query(conn, sql, new BeanHandler<Admin>(Admin.class), 29);
		
		System.out.println(admin);
		conn.close();
	}
	
	// 2）BeanListHandler: 查询返回list集合，集合元素是指定的对象
	@Test
	public void testQueryMany() throws Exception {
		String sql = "select * from admin";
		conn = JdbcUtil.getConnection();
		QueryRunner qr = new QueryRunner();
		// 查询全部数据
		List<Admin> list = qr.query(conn, sql, new BeanListHandler<Admin>(Admin.class));
		
		System.out.println(list);
		conn.close();
	}
    
	@Test
	//	3) ArrayHandler, 查询返回结果记录的第一行，封装对对象数组, 即返回：Object[]
	//	4) ArrayListHandler, 把查询的每一行都封装为对象数组，再添加到list集合中
	//	5) ScalarHandler 查询返回结果记录的第一行的第一列  (在聚合函数统计的时候用)
	//	6) MapHandler  查询返回结果的第一条记录封装为map
	public void testArray() throws Exception {
        
		String sql = "select * from admin";
		conn = JdbcUtil.getConnection();
		QueryRunner qr = new QueryRunner();
        
		// 查询
		//Object[] obj = qr.query(conn, sql, new ArrayHandler());
		//List<Object[]> list = qr.query(conn, sql, new ArrayListHandler());
		//Long num = qr.query(conn, sql, new ScalarHandler<Long>());
		Map<String, Object> map = qr.query(conn,sql, new MapHandler());
		
		conn.close();
	}
	
	
	
}
```

```java
package cn.ixfosa.dbUtils;

import java.sql.Connection;

import org.apache.commons.dbutils.DbUtils;
import org.apache.commons.dbutils.QueryRunner;
import org.junit.Test;

import cn.ixfosa.utils.JdbcUtil;

public class App_update {

	private Connection conn;

	// 1. 更新
	@Test
	public void testUpdate() throws Exception {
		String sql = "delete from admin where id=?";
		// 连接对象
		conn = JdbcUtil.getConnection();

		// 创建DbUtils核心工具类对象
		QueryRunner qr = new QueryRunner();
		qr.update(conn, sql, 26);

		// 关闭
		DbUtils.close(conn);
	}

	// 2. 批处理
	@Test
	public void testBatch() throws Exception {
		String sql = "insert into admin (userName, pwd) values(?,?)";
		conn = JdbcUtil.getConnection();
		QueryRunner qr = new QueryRunner();
		// 批量删除
		qr.batch(conn, sql, new Object[][]{ {"jack1","888"},{"jack2","999"}  });
		
		// 关闭
		conn.close();
	}
}
```



## 开源的连接池技术

### DBCP连接池：

DBCP 是 Apache 软件基金组织下的开源连接池实现，使用DBCP数据源，应用程序应在系统中增加如下两个 jar 文件：
		Commons-dbcp.jar：连接池的实现
		Commons-pool.jar：连接池实现的依赖库

Tomcat 的连接池正是采用该连接池来实现的。该数据库连接池既可以与应用服务器整合使用，也可由应用程序独立使用。

核心类：`BasicDataSource`

使用步骤
		引入jar文件
				commons-dbcp-1.4.jar
				commons-pool-1.5.6.jar

src/cn/ixfosa/dbcp/db.properties

```properties
url=jdbc:mysql:///jdbc_demo
driverClassName=com.mysql.jdbc.Driver
username=root
password=root
initialSize=3
maxActive=6
maxIdle=3000
```



```java
package cn.ixfosa.dbcp;

import java.io.InputStream;
import java.sql.Connection;
import java.util.Properties;

import javax.sql.DataSource;

import org.apache.commons.dbcp.BasicDataSource;
import org.apache.commons.dbcp.BasicDataSourceFactory;
import org.junit.Test;

public class App_DBCP {

	// 1. 硬编码方式实现连接池
	@Test
	public void testDbcp() throws Exception {
		// DBCP连接池核心类
		BasicDataSource dataSouce = new BasicDataSource();
		// 连接池参数配置：初始化连接数、最大连接数 / 连接字符串、驱动、用户、密码
		dataSouce.setUrl("jdbc:mysql:///jdbc_demo");			//数据库连接字符串
		dataSouce.setDriverClassName("com.mysql.jdbc.Driver");  //数据库驱动
		dataSouce.setUsername("root");							//数据库连接用户
		dataSouce.setPassword("root"); 							//数据库连接密码
		dataSouce.setInitialSize(3);  // 初始化连接
		dataSouce.setMaxActive(6);	  // 最大连接
		dataSouce.setMaxIdle(3000);   // 最大空闲时间
		
		// 获取连接
		Connection con = dataSouce.getConnection();
		con.prepareStatement("delete from admin where id=3").executeUpdate();
		// 关闭
		con.close();
	}
	
    
	@Test
	// 2. 【推荐】配置方式实现连接池  ,  便于维护
	public void testProp() throws Exception {
		// 加载prop配置文件
		Properties prop = new Properties();
		// 获取文件流   src/cn/ixfosa/dbcp/db.properties
		InputStream inStream = App_DBCP.class.getResourceAsStream("db.properties");
		// 加载属性配置文件
		prop.load(inStream);
		// 根据prop配置，直接创建数据源对象
		DataSource dataSouce = BasicDataSourceFactory.createDataSource(prop);
		
		// 获取连接
		Connection con = dataSouce.getConnection();
		con.prepareStatement("delete from admin where id=4").executeUpdate();
		// 关闭
		con.close();
	}
}
```



### C3P0连接池

最常用的连接池技术！Spring框架，默认支持C3P0连接池技术！
C3P0连接池，核心类：
		`CombopooledDataSource ds`;

使用：
		1.下载，引入jar文件:  c3p0-0.9.1.2.jar
		2.使用连接池，创建连接
				a)硬编码方式
				b)配置方式(xml)

src/c3p0-config.xml

```xml
<c3p0-config>
	<default-config>
		<property name="jdbcUrl">jdbc:mysql://localhost:3306/jdbc_demo
		</property>
		<property name="driverClass">com.mysql.jdbc.Driver</property>
		<property name="user">root</property>
		<property name="password">root</property>
		<property name="initialPoolSize">3</property>
		<property name="maxPoolSize">6</property>
		<property name="maxIdleTime">1000</property>
	</default-config>


	<named-config name="oracle_config">
		<property name="jdbcUrl">jdbc:mysql://localhost:3306/jdbc_demo</property>
		<property name="driverClass">com.mysql.jdbc.Driver</property>
		<property name="user">root</property>
		<property name="password">root</property>
		<property name="initialPoolSize">3</property>
		<property name="maxPoolSize">6</property>
		<property name="maxIdleTime">1000</property>
	</named-config>
</c3p0-config>
```

```java
package cn.ixfosa.c3p0;

import java.sql.Connection;
import java.sql.PreparedStatement;

import org.junit.Test;

import com.mchange.v2.c3p0.ComboPooledDataSource;

public class App {

	@Test
	//1. 硬编码方式，使用C3P0连接池管理连接
	public void testCode() throws Exception {
		// 创建连接池核心工具类
		ComboPooledDataSource dataSource = new ComboPooledDataSource();
		// 设置连接参数：url、驱动、用户密码、初始连接数、最大连接数
		dataSource.setJdbcUrl("jdbc:mysql://localhost:3306/jdbc_demo");
		dataSource.setDriverClass("com.mysql.jdbc.Driver");
		dataSource.setUser("root");
		dataSource.setPassword("root");
		dataSource.setInitialPoolSize(3);
		dataSource.setMaxPoolSize(6);
		dataSource.setMaxIdleTime(1000);
		
		// ---> 从连接池对象中，获取连接对象
		Connection con = dataSource.getConnection();
		// 执行更新
		con.prepareStatement("delete from admin where id=7").executeUpdate();
		// 关闭
		con.close();
	}
	
	@Test
	//2. XML配置方式，使用C3P0连接池管理连接
	public void testXML() throws Exception {
		// 创建c3p0连接池核心工具类
		// 自动加载src下c3p0的配置文件【c3p0-config.xml】
		ComboPooledDataSource dataSource = new ComboPooledDataSource();// 使用默认的配置
		PreparedStatement pstmt = null;
		
		// 获取连接
		Connection con = dataSource.getConnection();
		for (int i=1; i<11;i++){
			String sql = "insert into employee(empName,dept_id) values(?,?)";
			// 执行更新
			pstmt = con.prepareStatement(sql);
			pstmt.setString(1, "Rose" + i);
			pstmt.setInt(2, 1);
			pstmt.executeUpdate();
		}
		pstmt.close();
		// 关闭
		con.close();
	}
}
```



## Log4J-日志

Log4j,  log for java, 开源的日志组件！

使用步骤：
	1. 下载组件，引入jar文件;
		log4j-1.2.11.jar
	2. 配置 :  src/log4j.properties
	3. 使用

src/log4j.properties

```properties
# 通过根元素指定日志输出的级别、目的地： 
#  日志输出优先级： debug < info < warn < error 
log4j.rootLogger=info,console,file

############# 日志输出到控制台 #############
# 日志输出到控制台使用的api类
log4j.appender.console=org.apache.log4j.ConsoleAppender
# 指定日志输出的格式： 灵活的格式
log4j.appender.console.layout=org.apache.log4j.PatternLayout
# 具体格式内容
log4j.appender.console.layout.ConversionPattern=%d %p %c.%M()-%m%n


############# 日志输出到文件 #############
log4j.appender.file=org.apache.log4j.RollingFileAppender
# 文件参数： 指定日志文件路径
log4j.appender.file.File=../logs/MyLog.log
# 文件参数： 指定日志文件最大大小
log4j.appender.file.MaxFileSize=5kb
# 文件参数： 指定产生日志文件的最大数目
log4j.appender.file.MaxBackupIndex=100
# 日志格式
log4j.appender.file.layout=org.apache.log4j.PatternLayout
log4j.appender.file.layout.ConversionPattern=%d %c.%M()-%m%n
```

```java
package cn.ixfosa.log4j;

import java.io.IOException;
import java.io.PrintWriter;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;

public class Index extends HttpServlet {
	
	Log log =  LogFactory.getLog(Index.class);

	public void doGet(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, IOException {
		try {
			log.info("进入servlet");
			int i = 1/0;
			log.info("进入servlet结束");
		} catch (Exception e) {
			log.error("计算异常！",e);
		}
	}
}
```

```java
package cn.ixfosa.log4j;

import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
import org.junit.Test;


public class App {
	
	Log log = LogFactory.getLog(App.class);
	
	@Test
	public void save() {
		try {
            
			log.info("保存： 开始进入保存方法");

			int i = 1/0;
			
			log.info("保存： 执行保存结束，成功");
            
		} catch (Exception e) {
			
			log.error("执行App类Save()方法出现异常！");  // 异常
			e.printStackTrace();
		}
	}
	
	/*
	 * 思考： 日志的输出级别作用？
	 * 	 ----> 控制日志输出的内容。
	 */
	@Test
	public void testLog() throws Exception {
		// 输出不同级别的提示
		log.debug("调试信息");
		log.info("信息提示");
		log.warn("警告");
		log.error("异常");
	}
}
```

