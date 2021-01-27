## java基本功

### Java 语言特点

1. 简单易学；
2. 面向对象（封装，继承，多态）；
3. 平台无关性（ Java 虚拟机实现平台无关性）；
4. 可靠性；
5. 安全性；
6. 支持多线程（ C++ 语言没有内置的多线程机制，因此必须调用操作系统的多线程功能来进行多线程程序设计，而 Java 语言却提供了多线程支持）；
7. 支持网络编程并且很方便（ Java 语言诞生本身就是为简化网络编程设计的，因此 Java 语言不仅支持网络编程而且很方便）；
8. 编译与解释并存；

### 关于 JVM JDK 和 JRE

#### JVM

Java 虚拟机（JVM）是运行 Java 字节码的虚拟机。JVM 有针对不同系统的特定实现（Windows，Linux，macOS），目的是使用相同的字节码，它们都会给出相同的结果。

**什么是字节码?采用字节码的好处是什么?**

> 在 Java 中，JVM 可以理解的代码就叫做`字节码`（即扩展名为 `.class` 的文件），它不面向任何特定的处理器，只面向虚拟机。Java 语言通过字节码的方式，在一定程度上解决了传统解释型语言执行效率低的问题，同时又保留了解释型语言可移植的特点。所以 Java 程序运行时比较高效，而且，由于字节码并不针对一种特定的机器，因此，Java 程序无须重新编译便可在多种不同操作系统的计算机上运行。

**Java 程序从源代码到运行一般有下面 3 步：**

![Java程序运行过程](images/Java 程序运行过程.png)

我们需要格外注意的是 .class->机器码 这一步。在这一步 JVM 类加载器首先加载字节码文件，然后通过解释器逐行解释执行，这种方式的执行速度会相对比较慢。而且，有些方法和代码块是经常需要被调用的(也就是所谓的热点代码)，所以后面引进了 JIT 编译器，而 JIT 属于运行时编译。当 JIT 编译器完成第一次编译后，其会将字节码对应的机器码保存下来，下次可以直接使用。而我们知道，机器码的运行效率肯定是高于 Java 解释器的。这也解释了我们为什么经常会说 Java 是编译与解释共存的语言。

> HotSpot 采用了惰性评估(Lazy Evaluation)的做法，根据二八定律，消耗大部分系统资源的只有那一小部分的代码（热点代码），而这也就是 JIT 所需要编译的部分。JVM 会根据代码每次被执行的情况收集信息并相应地做出一些优化，因此执行的次数越多，它的速度就越快。JDK 9 引入了一种新的编译模式 AOT(Ahead of Time Compilation)，它是直接将字节码编译成机器码，这样就避免了 JIT 预热等各方面的开销。JDK 支持分层编译和 AOT 协作使用。但是 ，AOT 编译器的编译质量是肯定比不上 JIT 编译器的。

**总结：**

Java 虚拟机（JVM）是运行 Java 字节码的虚拟机。JVM 有针对不同系统的特定实现（Windows，Linux，macOS），目的是使用相同的字节码，它们都会给出相同的结果。字节码和不同系统的 JVM 实现是 Java 语言“一次编译，随处可以运行”的关键所在。

#### JDK 和 JRE

JDK 是 Java Development Kit，它是功能齐全的 Java SDK。它拥有 JRE 所拥有的一切，还有编译器（javac）和工具（如 javadoc 和 jdb）。它能够创建和编译程序。

JRE 是 Java 运行时环境。它是运行已编译 Java 程序所需的所有内容的集合，包括 Java 虚拟机（JVM），Java 类库，java 命令和其他的一些基础构件。但是，它不能用于创建新程序。

如果你只是为了运行一下 Java 程序的话，那么你只需要安装 JRE 就可以了。如果你需要进行一些 Java 编程方面的工作，那么你就需要安装 JDK 了。但是，这不是绝对的。有时，即使您不打算在计算机上进行任何 Java 开发，仍然需要安装 JDK。例如，如果要使用 JSP 部署 Web 应用程序，那么从技术上讲，您只是在应用程序服务器中运行 Java 程序。那你为什么需要 JDK 呢？因为应用程序服务器会将 JSP 转换为 Java servlet，并且需要使用 JDK 来编译 servlet。

> ```
> jre:  java运行环境。  jre =  java虚拟机 + 核心类库(辅助java虚拟机运行的文件)
> 	
> jdk： java开发工具集   jdk = jre + java开发工具。
> ```

####  Oracle JDK 和 OpenJDK 的对比

> 问：OpenJDK 存储库中的源代码与用于构建 Oracle JDK 的代码之间有什么区别？
>
> 答：非常接近 - 我们的 Oracle JDK 版本构建过程基于 OpenJDK 7 构建，只添加了几个部分，例如部署代码，其中包括 Oracle 的 Java 插件和 Java WebStart 的实现，以及一些封闭的源代码派对组件，如图形光栅化器，一些开源的第三方组件，如 Rhino，以及一些零碎的东西，如附加文档或第三方字体。展望未来，我们的目的是开源 Oracle JDK 的所有部分，除了我们考虑商业功能的部分。



#### jdk目录的介绍：

> ​	bin: 存放的是java的开发工具。
> ​	db : JDK7附带的一个轻量级的数据库，名字叫做Derby。
> ​	include ：存放的调用系统资源的接口文件。
> ​	jre ： java运行环境
> ​	lib : 核心类库。
> ​	src.zip : java源代码 


java.exe 启动java虚拟机解释并执行指定的class文件。

javac.exe  启动java编译器对指定的java源文件进行编译。

#### 环境变量path和classpath的作用

**1. PATH环境变量**
		作用是指定命令搜索路径，在i命令行下面执行命令如javac编译java程序时，它会到PATH变量所指定的路径中查找看是否能找到相应的命令程序。
	需要把jdk安装目录下的bin目录增加到现有的PATH变量中，bin目录中包含经可执行文件javac/java/javadoc等待，设置PATH变量后，就可以在任何目录下执行javac/java等工具了。
**2. CLASSPATH环境变量。**
		作用是指定类搜索路径，要使用已经编写好的类，前提当然是能够找到它们了，JVM就是通过CLASSPTH来寻找类的。
		需要把jdk安装目录下的lib子目录中的dt.jar和tools.jar设置到CLASSPATH中，当然，当前目录“.”也必须加入到该变量中。
新版的软件，classpath没有什么用了，因为在安装的时候已经选了JDK而且可以添加 
**3. JAVA_HOME环境变量**
		它指向jdk的安装目录，Eclipse/NetBeans/Tomcat等软件就是通过搜索JAVA_HOME变量来找到并使用安装好的jdk。

####  Java 程序的主类

一个程序中可以有多个类，但只能有一个类是主类。在 Java 应用程序中，这个主类是指包含 main（）方法的类。而在 Java 小程序中，这个主类是一个继承自系统类 JApplet 或 Applet 的子类。应用程序的主类不一定要求是 public 类，但小程序的主类要求必须是 public 类。主类是 Java 程序执行的入口点
应用程序是主线启动（main()方法）,applet小程序没有main()方法，主要是潜在浏览器页面上通过init()、run()启动。

> 运行于Web浏览器之上的Java程序叫做applet，applet代码嵌在HTML代码中。applet采用带有按钮、文本字段、文本区域，单选按钮等的现代图形界面，与Web上的用户进行交互并处理请求。

#### import java 和 javax 区别

刚开始的时候 JavaAPI 所必需的包是 java 开头的包，javax 当时只是扩展 API 包来使用。然而随着时间的推移，javax 逐渐地扩展成为 Java API 的组成部分。但是，将扩展从 javax 包移动到 java 包确实太麻烦了，最终会破坏一堆现有的代码。因此，最终决定 javax 包将成为标准 API 的一部分。所以，实际上 java 和 javax 没有区别。这都是一个名字。

#### Java 语言“编译与解释并存”

高级编程语言按照程序的执行方式分为编译型和解释型两种。简单来说，编译型语言是指编译器针对特定的操作系统将源代码一次性翻译成可被该平台执行的机器码；解释型语言是指解释器对源程序逐行解释成特定平台的机器码并立即执行。比如，你想阅读一本英文名著，你可以找一个英文翻译人员帮助你阅读，
有两种选择方式，你可以先等翻译人员将全本的英文名著（也就是源码）都翻译成汉语，再去阅读，也可以让翻译人员翻译一段，你在旁边阅读一段，慢慢把书读完。

Java 语言既具有编译型语言的特征，也具有解释型语言的特征，因为 Java 程序要经过先编译，后解释两个步骤，由 Java 编写的程序需要先经过编译步骤，生成字节码（\*.class 文件），这种字节码必须由 Java 解释器来解释执行。因此，我们可以认为 Java 语言编译与解释并存。

## java语法

### 注释

注释就是使用文字对程序的说明，注释是写给程序员看的，编译器会忽略注释的 内容的。

**注释的类别：**

> ​		第一种： 单行注释。   // 注释的内容
>
> ​		第二种： 多行注释。   /* 注释的内容  */
>
> ​		第三种： 文档注释.    /** 注释的内容  */

​		文档注释也是一个多行注释。

**注释要注意的细节：**

	1. 单行注释可以嵌套使用，多行注意是不能嵌套使用的。

**注释的作用：**

 	1. 使用文字对程序说明。
 	2. 调试程序

**文档注释与多行注释的区别：**

多行注释与文档注释的区别： 多行注释的内容不能用于生成一个开发者文档，
而文档注释 的内容可以生产一个开发者文档。

**使用javadoc开发工具即可生成一个开发者文档。**
`javadoc`工具的使用格式：
`javadoc -d `   存放文档的路径   java的源文件 

**使用javadoc工具要注意细节：**

  		1. 如果一个类需要使用javadoc工具生成一个软件的开发者文档，那么该类必须使用public修饰。
 		2. 文档注释注释的内容一般都是位于类或者方法的上面的。

> 写注释的规范：一般**`单行注释是位于代码的右侧`，`多行注释与文档注释一般是写在类或者方法的上面的`。**

**代码的注释不是越详细越好。实际上好的代码本身就是注释，我们要尽量规范和美化自己的代码来减少不必要的注释。**

**若编程语言足够有表达力，就不需要注释，尽量通过代码来阐述。**

### 标识符和关键字

#### 标识符

在java程序中有些名字是可以自定义的，那么这些自定义的名字我们就称作为自定义的标识符。

**标识符要注意的细节：**

  1. 标识符的组成元素是由 字母（`a-zA-Z`） 、数字(`0-9`) 、 下划线(`_`)、 美元符号(`$`).
  2. 标识符不能以数字开头。
  3. 标识符是严格区分大小写的。
  4. 标识符的长度是没有长度限制的。
  5. 标识符的命名一般要有意义（要做到让人见名知意）
  6. 关键字、保留字不能用于自定义的标识符。

**自定义标识符的规范：**

 	1. 类名和接口名单词的首字母大写，其他单词小写。 比如： RunTime.
	2. 变量名与方法名首单词全部小写，其他单词首字母大写，其他小写。  比如： doCook()；
	3. 包名全部单词小写。
	4. 常量全部单词大写，单词与单词之间使用下划线分隔。  比如： UP_DIRECTION

**判断一下那些是 符合的标识符:**
		12abc_   不合法  数字不能开头
		_12abc    合法
		$ab12#    不合法  #号不属于标识符组成元素。
		abc@123   不合法  @号不属于标识符组成元素。 

#### 关键字

**关键字**：关键字就是在java程序中具备特殊含义的标识符。关键字一般用于描述一个程序的结构或者表示数据类型。

|       访问控制       | private  | protected  | public   |              |            |           |        |
| :------------------: | -------- | ---------- | -------- | ------------ | ---------- | --------- | ------ |
| 类，方法和变量修饰符 | abstract | class      | extends  | final        | implements | interface | native |
|                      | new      | static     | strictfp | synchronized | transient  | volatile  |        |
|       程序控制       | break    | continue   | return   | do           | while      | if        | else   |
|                      | for      | instanceof | switch   | case         | default    |           |        |
|       错误处理       | try      | catch      | throw    | throws       | finally    |           |        |
|        包相关        | import   | package    |          |              |            |           |        |
|       基本类型       | boolean  | byte       | char     | double       | float      | int       | long   |
|                      | short    | null       | true     | false        |            |           |        |
|       变量引用       | super    | this       | void     |              |            |           |        |
|        保留字        | goto     | const      |          |              |            |           |        |

### 常量和变量

#### 常量

**常量就是程序在运行过程中其值不能发生改变的量。**

常量的类别：
		整数常量      10 12
		小数常量	  3.14
		布尔常量      布尔常量只有两个值： true(正确). false(错误)
		字符常量      字符常量就是单个字符使用单引号引起来的内容我们称作为字符常量。
		字符串常量    字符串常量就是使用双引号引起来的内容称作为字字符串常量。


整数常量的表现形式：整数的表现形式主要是以不同的进制(二进制、八进制、十六进制)表现出来。

```java
class Demo{
	public static void main(String[] args){
		System.out.println(12); //整数常量
		System.out.println(3.14); //小数常量
		System.out.println(false); //布尔常量
		System.out.println('1'); //字符常量
		System.out.println("hello world"); //字符串常量

		//如果一个数据没有加上任何的标识之前，默认就是十进制的数据。
		System.out.println(10);//十进制  10
		System.out.println(0b10); //二进制 2
		System.out.println(010); //八进制数据， 八进制的数据需要以0开头  8 
		System.out.println(0x10); //十六进制数据， 十六进制的数据需要以0x开头    16
	}
}
```

#### 变量

变量就是在程序运行过程中其值可以发生变化的量,     变量就是一个存储数据的容器。

容器具备什么特点：
	1. 容量(大小)。
	2. 存储一定格式的数据。
	3. 名字.

声明（定义）一个变量的格式：
- 容量  变量名字 = 数据。   int age = 22;
- **变量名的命名规范： 首单词小写，其他单词首字母大写，其他小写**  



**声明变量的方式：**

​		声明变量的方式一：
​		
​				数据类型 变量名;

​		声明变量的方式二： 一次性声明了多个相同类型变量

​				据类型  变量名1 ， 变量2....

注意：

 	1. 变量是必须先声明再使用。
 	2. 在一个作用域中不能声明同名的变量。  

```java
class Demo {
	public static void main(String[] args) {	
        
		//int age = 12; //声明一个变量
	
		int age ,height ;  //声明了变量
        
		//给变量赋值
		age = 10;
		height = 175;
		
		System.out.println(age);
		System.out.println(height);

	}
}
```


### 基本数据类型

Java**中**有8种基本数据类型，分别为：

1. 6种数字类型 ：byte、short、int、long、float、double
2. 1种字符类型：char
3. 1中布尔型：boolean

这八种基本类型都有对应的包装类分别为：Byte、Short、Integer、Long、Float、Double、Character、Boolean

| 基本类型  | 位数 | 字节  | 默认值  |      |
| :------- | :--- | :--- | ------- | ---- |
| int      | 32   | 4    | 0       |      |
| short    | 16   | 2    | 0       |      |
| long     | 64   | 8    | 0L      |      |
| byte     | 8    | 1    | 0       |      |
| char     | 16   | 2    | 'u0000' |      |
| float    | 32   | 4    | 0f      |      |
| double   | 64   | 8    | 0d      |      |
| boolean  | 1    |      | false   |      |



**布尔类型： 布尔类型只有两个 值，true或者false。**1字节或者4个字节 

如果使用boolean声明一个基本类型的变量时，那么该变量占4个字节，
如果使用boolean声明一个数组类型 的时候，那么每个数组的元素占一个字节。

对于`boolean`，官方文档未明确定义，它依赖于 JVM 厂商的具体实现。逻辑上理解是占用 1位，但是实际中会考虑计算机高效存储因素。

注意：

1. Java 里使用 long 类型的数据一定要在数值后面加上 **L**，否则将作为整型解析：
2. `char a = 'h'`char :单引号，`String a = "hello"` :双引号



注意： 如果一个整数没有加上任何的标识的时候，那么默认是int类型的数据。
如果需要把该数据表示成一个`long`类型的数据，那么需要加数据后面加上`L`表示，L是不区分大小写的，但是建议使用大写。

注意： 如果一个小数没有加上任何标识的时候，那么该小数默认是`double`类型的数据，如果需要表示成`float`类型，那么需要在小数的后面加上`f`表示。f不区分大小写的。



字符串的数据类型是：`String `引用数据类型,不属于基本数据类型。

#### 数据类型转换 

- 小数据类型-------->大数据类型     自动类型转换

- 大数据类型--------->小数据类型    强制类型转换


强制类型转换的格式：
		**小数据类型  变量名 = (小数据类型)大数据类型**

数据类型转换要注意的细节：
	1. 凡是byte、short 、 char数据类型数据在运算的时候都会自动转换成int类型的数据再运算。
	2. 两个不同数据类型的数据在运算的时候，结果取决于大的数据类型。

```java
class Demo {
	public static void main(String[] args) {
		
		byte  b=  11;   // 一个字节   一两碗
		short s = b; // 把b变量存储的值赋予给s变量。 2字节 二两的碗
		int i = s;   // i是4字节
		long l = i;  // l 是8个字节。 
		System.out.println(l);
		
		int i = 128;  //4个字节
		byte b =(byte) i;  // 1个字节
		System.out.println(b);  // -128 
		
		//如果是负数，那么最高位肯定是1， 正数的二进制位的最高位是0。
        //sun给我们提供一个功能 Integer.tobinaryString()  查看一个数据的二进制数据形式的。
		System.out.println(Integer.toBinaryString(-7)); // 11111001（补码）

		//凡是byte、short 、 char数据类型数据在运算的时候都会自动转换成int类型的数据再运算。
		byte b1 =1;
		byte b2 = 2;
		byte b3 = (byte)(b1+b2);
		System.out.println(b3); //3  
		
		System.out.println('a'+1); //98 
				
		//两个不同数据类型的数据在运算的时候，结果取决于大的数据类型
		int i =10;
		long l = 20;
		i = (int)(i+l); 
		System.out.println(i);

		int i = 10;  //
		byte b = i;  // 一个整数没有加上任何标识的时候，默认是int类型的数据。
		//10 是一个常量， 编译器在编译的时候就能够确认常量的值了，byte b = 10,在编译到的时候
		//java编译器就会检查到10并没有超出byte的表示范围，所以允许赋值。
		//java编译器在编译 的时候并不能确认变量所存储的值，变量存储的值是在运行的时候才在内存中分配空间的。
		System.out.println(b);

	}
}
```





####  自动装箱与拆箱

- **装箱**：将基本类型用它们对应的引用类型包装起来；
- **拆箱**：将包装类型转换为基本数据类型；

#### 几种基本类型的包装类和常量池

**Java 基本类型的包装类的大部分都实现了常量池技术，即 Byte,Short,Integer,Long,Character,Boolean；前面 4 种包装类默认创建了数值[-128，127] 的相应类型的缓存数据，Character创建了数值在[0,127]范围的缓存数据，Boolean 直接返回True Or False。如果超出对应范围仍然会去创建新的对象。** 为啥把缓存设置为[-128，127]区间？性能和资源之间的权衡。

```java
public static Boolean valueOf(boolean b) {
    return (b ? TRUE : FALSE);
}
```

```java
private static class CharacterCache {         
    private CharacterCache(){}

    static final Character cache[] = new Character[127 + 1];          
    static {
        for (int i = 0; i < cache.length; i++)                 
            cache[i] = new Character((char)i);         
    }   
}
```

**两种浮点数类型的包装类 Float,Double 并没有实现常量池技术**

```java
		Integer i1 = 33;
		Integer i2 = 33;
		System.out.println(i1 == i2);// 输出 true
		Integer i11 = 333;
		Integer i22 = 333;
		System.out.println(i11 == i22);// 输出 false
		Double i3 = 1.2;
		Double i4 = 1.2;
		System.out.println(i3 == i4);// 输出 false
```

**Integer 缓存源代码：** 

```java
/**
*此方法将始终缓存-128 到 127（包括端点）范围内的值，并可以缓存此范围之外的其他值。
*/
    public static Integer valueOf(int i) {
        if (i >= IntegerCache.low && i <= IntegerCache.high)
            return IntegerCache.cache[i + (-IntegerCache.low)];
        return new Integer(i);
    }

```

**应用场景：**

1. Integer i1=40；Java 在编译的时候会直接将代码封装成 Integer i1=Integer.valueOf(40);，从而使用常量池中的对象。
2. Integer i1 = new Integer(40);这种情况下会创建新的对象。

```java
  Integer i1 = 40;
  Integer i2 = new Integer(40);
  System.out.println(i1==i2);//输出 false
```

**Integer 比较更丰富的一个例子:**

```java
  Integer i1 = 40;
  Integer i2 = 40;
  Integer i3 = 0;
  Integer i4 = new Integer(40);
  Integer i5 = new Integer(40);
  Integer i6 = new Integer(0);
  
  System.out.println("i1=i2   " + (i1 == i2));
  System.out.println("i1=i2+i3   " + (i1 == i2 + i3));
  System.out.println("i1=i4   " + (i1 == i4));
  System.out.println("i4=i5   " + (i4 == i5));
  System.out.println("i4=i5+i6   " + (i4 == i5 + i6));
  System.out.println("40=i5+i6   " + (40 == i5 + i6));
```

结果：

```
i1=i2   true
i1=i2+i3   true
i1=i4   false
i4=i5   false
i4=i5+i6   true
40=i5+i6   true
```

解释：

语句 i4 == i5 + i6，因为+这个操作符不适用于 Integer 对象，首先 i5 和 i6 进行自动拆箱操作，进行数值相加，即 i4 == 40。然后 Integer 对象无法与数值进行直接比较，所以 i4 自动拆箱转为 int 值 40，最终这条语句转为 40 == 40 进行数值比较。

### 运算符

#### 算术运算符:

`+ `  (正数、加法、连接符)
		连接符的作用： 让任何的数据都可以与字符串进行拼接。
		如果+号用于字符串的时候，那么+号就是一个连接符，并不是 做加法功能了
        连接符要注意：任何类型的数据与字符串使用连接符连接，那么结果都是字符串类型的数据。
`-`(减法)

`*`（乘法）

`/ ` (除法)

`%`(取模、取余数)

```java
class Demo {

	public static void main(String[] args){
	
		int i1 = 1;
		int i2 = 2;
		System.out.println(i1 +" world");  //    1 world
		
		// 计算机每次运算的时候只能取两个 数据运算。
		System.out.println(1+2+3 +" world"+1+2+3); // 6 world123

		double a= 12.0;
		int b = 3;
		System.out.println(a/b); // 4.0

		//在java中做取模 运算的时，结果的正负号是取决于被除数。
		System.out.println("结果:"+(10%3));  // 1
		System.out.println("结果:"+(10%-3)); // 1      
		System.out.println("结果:"+(-10%3)); // -1    
		System.out.println("结果:"+(-10%-3)); // -1 
	}
}
```

#### 自增和自减

**自增**

`++ `（自增） :  自增就是相当于操作数+1.

前自增：++位于操作数的前面。  比如： ++a；
前自增：先自增，后使用。

后自增：++位于操作数的后面。  比如： a++；
后自增： 先使用，后自增。

```java
class Demo {
	public static void main(String[] args) {
		
		int a = 0;
		int sum1 = ++a; //前自增。 a = a+1  ， sum = a
		System.out.println("sum1 = "+ sum1+ " a = "+ a);  //sum1 = 1 a = 1
		int sum2 = a++; //后自增  sum = 0 , a = a+1
		System.out.println("sum2 = "+ sum2+ " a = "+ a); //sum2 = 1 a = 2

   		//后自增在jvm的运行原理：
		//因为后自增要使用 到没有+1之前 的值，那么jvm会先声明一个变量用于保存没有+1之前的值。
		//int i = 0;
		//i = temp;

		/*原理：
			1. int temp = i; // 声明了一个临时变量用于记录了i没有加1之前的值。
			2. 自增。  i = i+1;   i = 1;
			3. temp把用作了表达式 的结果。
	
		i的值发生了几次变化：
			i = 0 -----> 1----> 0*/

		int i = 0;
		i = i++; // 后自增...  后自增要使用到没有+1之前的值。
		System.out.println("i= "+i);   //0
	}
}
```

**自减**

自减： 操作数-1.

前自减: --位于操作数的前面。  --操作数
前自减： 先自减，后使用。

后自减：--位于操作数的后面。 操作数--；
后自减： 先使用，后自减。

```java
class Demo {
	public static void main(String[] args) {	
		/*
		int i = 1;
		int sum = --i;  //前自减   i = i-1 , sum = i;
		int sum = i--; // 后自减   sum = i ; i = i-1;
		System.out.println("sum = "+ sum);  // 0  1 
		*/

		int num = 10;
		//int sum = 10* num++; //后自增
		int sum = 10* ++num; //前自增 
		System.out.println("sum = "+ sum);  //110
	}
}
```

#### 赋值运算符

赋值运算符：
	`= `  (赋值运算符)

​	`+= `

​	`-=`

​	`*=`

​	`/=`

​	`%=`

```java
class Demo {
	public static void main(String[] args) {
		
		int i = 10; // 把10赋予给i变量。
		i+=2;  // i = i+2; 
		System.out.println("i = "+i);  //12
	
		byte b1 = 1;
		byte b2 = 2;
		//b2 = b2+b1; //报错。， 需要强制类型转换   错误: 不兼容的类型: 从int转换到byte可能会有损失
		//b2 = (byte)(b2+b1);
		
		b2+=b1;       //b2 = b2+ b1;	 b2+=b1 在编译的时候，java编译器会进行强制类型转换，不需要我们手动转换了。
		System.out.println("b2 : "+ b2);  //3
	}
}
```

#### 比较运算符

比较运算符： 比较运算符的结果都是返回一个`布尔值`的。

`==` (判断是否等于)
==用于比较两个**基本数据类型数据**的时候，比较的是两个变量所**存储的值**是否一致.
==用于比较两个**引用类型变量的数据**时候，比较的是两个 引用类型变量所记录的**内存地址是否一致**. 

`!= ` (不等于)

`>`   (大于)

`<`   (小于)

`=`   (大于等于)

`<=` (小于等于)

```java
class Demo {
	public static void main(String[] args) {
		int a = 10;
		int b =10;
		System.out.println("10等于10吗？"+ (a==b));  //true  
		System.out.println("10不等于1吗？"+ (10!=1) ); //true  

		byte c  = 10;
		long l = 30;
		System.out.println(l>c); //true   两个不同类型的数据是否可以比较呢.,可以的，但是两个不                                  //同类型的数据必须是兼用的数据。
		//这个比较的过程会先把b转换了long类型的数据，然后再进行比较 。

		System.out.println('a'>50); //true  
	}
}
```

#### 位运算符

位运算符：位运算符就是直接操作二进制位的。
	`& ` (与)
	`| ` (或)
	`^` (异或)
规律： 如果操作数A连续异或同一个操作数两次，那么结果还是操作数A。

应用： 对数据加密.

   `~`  (取反)

```java
class Demo{
	public static void main(String[] args){
		System.out.println(6&3); // 2 
		System.out.println(6|3); // 7
		System.out.println(6^3); //  5
		System.out.println(~7);  // -8
	}
}
```

![alt 位运算](images/位运算符.png)



#### 逻辑运算符

逻辑运算符 ：逻辑运算符的**作用是用于连接布尔表达式**的。 

`& `（与,并且）
规律： 只有左右变量同时 为true，那么结果才是true，否则就为false。

`|`  (或,或者)
规律： 只要两边的布尔表达式有一边为true，那么结果就为true，只有两边同时为false 的时候，结果才是false.

`^ ` (异或)
规律： 只要两边的布尔表达式 结果不一致，那么结果就为true，如果左右两边 的布尔表达式一致，那么就为false.

`！` (非)

`&&` (短路与\双与)

短路与和单与符号的相同与不同点：
		相同点： 短路与和单与运算 的结果是一样的。
		不同点： 使用短路与的时候，如果左边的布尔表达式为false，则不会在运算右边的布尔表达式，从而提高了效率。使用单与的时候，即使发现左边的布尔表达式为false，还是会运算右边的布尔表达式的。

​		只有左边的布尔表达式为false时，双与的效率才要高于单与的.

`|| `(短路或\双或)

短路或与单或的相同点与不同点：	
		相同点：运算的结果是一致 的。
		不同点：使用短路或的时候，当发现左边的布尔表达式为true时，则不会运算右边的布尔表达式。
使用单或的时候 发现左边的布尔表达式为true，还是会运算右边布尔表达式。

```java
class Demo {
	public static void main(String[] args) {
	
		int workAge = 2;
		int age = 24;
		System.out.println(workAge>=2|age++>18);
		System.out.println("age:"+ age);
		
		
		System.out.println(true&true);  //true
		System.out.println(true&false); // false
		System.out.println(false&true); // false
		System.out.println(false&false); // false
		

		System.out.println(true|true);  // true
		System.out.println(true|false); // true
		System.out.println(false|true); // true
		System.out.println(false|false); // false

		

		System.out.println(true^true);  //  false
		System.out.println(true^false); //  true
		System.out.println(false^true); //  true
		System.out.println(false^false); //  false
		
		System.out.println(!true); // 
		
		System.out.println(true&&true);  //true
		System.out.println(true&&false); // false
		System.out.println(false&&true); // false
		System.out.println(false&&false); // false

		System.out.println(true||true);  // true
		System.out.println(true||false); // true
		System.out.println(false||true); // true
		System.out.println(false||false); // false
	}
}
```

> ```java
> 位运算符可能会出现的笔试题目：
>  	1. 交换两个变量的值,不准出现第三方变量。
>  	2. 取出一个二进制数据的指定位数。要求读取该二进制数据的低4位
> 		00000000-00000000-00010100-01001101
> 	&   00000000-00000000-00000000-00001111
> 	------------------------------------------
> 		00000000-00000000-00000000-00001101
>     
>     
> class Demo {
> 	public static void main(String[] args) {
> 		int a = 3;
> 		int b = 5;
> 		/*
> 		第一种方式： 定义第三方变量。
> 		int temp = a;  //3 
> 		a = b; //a = 5 
> 		b = temp; 
> 		
> 		方式2：相加法， 缺点： 两个int类型的数据相加，有可能会出现超出int的表示范围。
> 		a = a+b;  // a =8
> 		b = a-b; //b = 8 - 5 = 3
> 		a = a-b; // a = 8 - 3 = 5
> 		
> 		方式3： 可以使用异或。 缺点： 逻辑不清晰。
> 		*/
> 		a = a^b;  // a = 3^5
> 		b = a^b;  // b = (3^5)^5 = 3
> 		a = a^b; //  a = (5^3)^3 = 5 
> 		System.out.println("a = "+ a+" b="+b);
> 	}
> ```

#### 移位运算符

`<<`    (左移)
规律：一个操作数进行左移运算的时候，结果就是等于操作数乘以2的n次方，n就是左移 的位数.
	>3<<1 = 3 *2(1) = 6;
	>3<<2 = 3*2(2) = 12
	>3<<3 = 3*2(3) = 24

`>>`    (右移)

规律：一个操作数在做右移运算的时候，实际上就是等于该操作数除以2的n次方，n就是右移的位数。

	>3>>1 = 3 / 2(1) = 1
	>3>>2 = 3 / 2(2) = 0 


​	

`>>>`(无符号右移) ：

> ​		无符号右移与右移的区别：进行右移运算的时候，如果操作数是一个正数，那么左边的空缺位使用0补，如果操作数是一个负数，那么左边的空缺位使用1补。而使用无符号右移的时候，不管是正数还是负数都统一使用0补。

笔试题目：使用最高的效率算出2乘以8的结果。

		>2<<3 = 2*2(3) = 16;

```java
class Demo {
	public static void main(String[] args){
		/*
		左移
		System.out.println(3<<1); // 6 
		System.out.println(3<<2); // 12
		System.out.println(3<<3); // 24 
		
		右移：
		*/
		System.out.println(3>>>1); // 1 
		System.out.println(3>>>2);  //0 
	}
}
```

#### 三元运算符(三目运算符)

格式;

	>布尔表达式？值1:值2  ;

三元运算符要注意的细节：
	使用三元运算符的时候，一定要使用该表达式返回的结果，或者是定义一个变量接收该表达式返回的结果。

```java
class Demo {
	public static void main(String[] args) {
		
		int age = 26;
		String result = age>=18? "成年人":"未成年人";
		System.out.println(result);

		int a = 1;
		int b = 2;
		int c = 3;

		int result = a*b+c/2+2*(a+b)/c;
		int res = c / 2;
		System.out.println("result="+result); //5  2+1+3/3 
		System.out.println("res="+res);   //1
	}
}
```

#### 使用异或对图片数据进行加密

```java
import java.io.*;
class ImageTest {
	public static void main(String[] args)  throws Exception {
        
		//找到图片文件
		File inFile = new File("e:\\加密的图片.jpg");
		File outFile = new File("e:\\解密的图片.jpg");

		//建立数据通道，让图片的二进制数据流入
		FileInputStream input = new FileInputStream(inFile);
		FileOutputStream output = new FileOutputStream(outFile);
        
		//边读，把读到的数据异或一个数据，把把数据写出
		int content = 0; //该变量是用于存储读取到的数据
		while((content = input.read())!=-1){  // 如果没有到文件的末尾，那么继续读取数据，读取                                               //到的数据已经存储到content变量中了。
			output.write(content^12);
		}
	
		//关闭资源
		output.close();
		input.close();
	}
}
```

### 控制流程语句

#### if语句

格式一： 只适用于一种情况下去使用。

	if(判断条件){
		符合条件执行的代码;
	}
格式二：适用于两种情况下去使用


	if(判断条件){
		符合条件执行的代码
	}else{
		不符合条件执行 的 代码
	}

格式3： 适用于多种情况使用的

	if(判断条件1){
		符合条件1执行的 语句;
	}else if(判断条件2){
		符合条件2执行 的语句;
	}else if(判断条件3){
		符合条件3执行 的语句;
	}else if(判断条件4){
		符合条件4执行 的语句;
	}......else{
		都不符合上述 条件执行的代码...
	}
#### switch语句

switch语句的格式：
	switch(你的选择){
		
		case 值1：
			符合值1执行的代码
			break;
		case 值2：
			符合值 2执行的代码
			break;
		case 值3：
			符合值 3执行的代码
			break;
		......
		default: 
			你的选择都符合上述的选项时执行的代码;
			break;
	}
switch语句要注意的事项：

> 1. switch语句使用的变量只能是`byte`、 `char`、 `short`、`int`、 `String`数据类型，String数据类型是从jdk7.0的时候开始支持的。
> 2. case后面跟 的数据必须是一个`常量`。
>  3. switch的停止条件：
>    	 		`switch`语句一旦匹配上了其中的一个`case`语句，那么就会执行对应的case中的语句代码，执行完毕之后如果没有
>    	 		遇到`break`关键字或者是结束switch语句的大括号，那么switch语句不会再判断，按照代码的顺序从上往下执行
>    	 		所有的代码。直到遇到break或者是结束siwitch语句的大括号为止。
>
> 4. 在switch语句中不管代码的顺序如何，永远都是会先判断case语句，然后没有符合的情况下才会执行`default`语句。

```java
class Demo {
	public static void main(String[] args) {
        
		int option = 13;	//定义一个变量存储你的选择
		switch(option){
			case 1:
				System.out.println("java");
				
			case 2:
				System.out.println("C#");
				
			case 3:
				System.out.println("javascript");
				
			case 4:
				System.out.println("android");
			default:
				System.out.println("你的选择有误");
		}
```

```java
import java.util.*;
class Demo {
	public static void main(String[] args) {
		System.out.println("请输入一个月份：");
		//创建一个扫描器
		Scanner scanner = new Scanner(System.in);
		//调用扫描器的nextInt方法
		int month = scanner.nextInt();
		switch(month){
			case 3:
			case 4:
			case 5:
				System.out.println("春天");
				break;
			case 6:
			case 7:
			case 8:
				System.out.println("夏天");
				break;
			case 9:
			case 10:
			case 11:
				System.out.println("秋天");
				break;
			case 12:
			case 1:
			case 2:
				System.out.println("冬天");
				break;
			default:
				System.out.println("没有对应的季节");
				break;
		}
	}
}

```

### 循环语句

#### while循环语句

while循环 语句的格式:

	while(循环的条件){
		循环语句；
	}
while循环语句要注意的事项：
	>1. while循环语句一般是通过一个变量控制其循环的次数。
	2. while循环语句的循环体代码如果只有一个语句的时候，那么可以省略大括号。但是也是不建议大家省略。
	
	3. while循环语句的判断条件后面不能跟有分号，否则会影响到执行的效果。

```java
class Demo {
	public static void main(String[] args) {
		 int count = 0;
		 while(count<5){
			System.out.println("Hello World!");
			count++;
		 }
	}
}
```

```java
//需求： 计算1+2+3+....+ 100的总和。
class Demo {
	public static void main(String[] args) {
		int num = 1;
		int sum  = 0;	//定义一个变量用于保存每次相加的结果
		while(num<=100){
			sum = sum+num; // sum = 1  
			num++;
		}
		System.out.println("sum = "+ sum);
	}
}
```

```java
//计算1-100,7的倍数总和。 7  14 21
class Demo {
	public static void main(String[] args){
        
		int num = 1;
		int sum = 0;	//定义一个变量用于保存每次相加的总和。
	
		while(num<=100){ // num = 1
			
			if(num%7==0){
				sum = sum+num;
			}
			num++;
		}
		System.out.println("总和是："+ sum);
	}
}
```

```java

/*
需求： 实现猜数字游戏， 如果没有猜对可以继续输入你猜的数字，如果猜对了停止程序。
*/
import java.util.*;
class Demo {
	public static void main(String[] args) 
	{
		//创建一个随机数对象
		Random random = new Random();
		//调用随机数对象的nextInt方法产生一个随机数
		int randomNum = random.nextInt(10)+1; //要求随机数是 1~10
		//创建一个扫描器对象
		Scanner scanner = new Scanner(System.in);
		
		while(true){
			System.out.println("请输入你要猜的数字:");
			//调用扫描器的nextInt方法扫描一个数字
			int guessNum = scanner.nextInt();
			if (guessNum>randomNum){
				System.out.println("猜大了..");
			}else if(guessNum<randomNum){
				System.out.println("猜小了..");	
			}else{
				System.out.println("恭喜你，猜对了`..");	
				break;
			}
		}
	}
}
```

#### do while循环语句

格式：

	do{
		....
	}while(判断条件);

while循环语句与do-while循环语句的区别：

	>while循环语句是先判断后执行循环语句的，do-while循环语句
	>是先执行，后判断。不管条件是否满足至少会执行一次。

```java
class Demo{
	public static void main(String[] args) 
	{
		
		/*
		int count =0; 
		while(count<5){
			System.out.println("Hello World!");     //5次
			count++;
		}
		*/

		/*
		boolean flag = false;
		while(flag){
			System.out.println("Hello World!"); //0次
		}
		*/
		
		/*
		boolean flag = false;
		do{
			System.out.println("Hello World!");   //1次
		}while(flag);
		*/

		int count = 0;
		do
		{
			System.out.println("hello world"); //5次
			count++;
		}while(count<5);
	}
}
```

```java
/*
需求： 使用do-while算出1-100之间偶数的总和。
*/
class Demo {
	public static void main(String[] args) {
		int num = 1;
		int sum = 0;	//定义一个变量用于保存每次相加的总和
		do{
			if(num%2==0){
				sum += num;
			}
			num++;
		}while(num<101);
		System.out.println("sum = "+ sum);
	}
}
```

#### for循环语句

for循环语句的格式:

```java
for(初始化语句;判断语句;循环后的语句){
	循环语句;
}
```
for循环语句 要注意的事项：
	>1. `for(;;)`这种写法 是一个死循环语句，相当于`while(true)`;
	>2. for循环语句的初始化语句只会执行一次，只是在第一次循环的时候执行而已。
	>3. for循环语句的循环体语句只有一句的时候，可以省略大括号不写。但是不建议省略。

```java
class Demo
{
	public static void main(String[] args) 
	{
		int count = 0 ;
		for(System.out.println("初始化语句A");count<5 ;System.out.println("循环后的语句C")){
			System.out.println("循环体语句B");
			count++;
		}

		/*初始化语句A
		循环体语句B
		循环后的语句C
		循环体语句B
		循环后的语句C
		循环体语句B
		循环后的语句C
		循环体语句B
		循环后的语句C
		循环体语句B
		循环后的语句C
		*/

		for(count = 0 ; count<5;  count++){
			System.out.println("hello world");  //5次
		}	

		/*for(int count = 0 ; count<5;  count++){
			System.out.println("hello world");  //5次
		}	*/
	}
}
```

```java
/*
需求： 在控制台上打印一个 五行五列矩形/.
	
	*****
	*****
	*****
	*****
	*****
	
先打印一行
*/
class Demo {
	public static void main(String[] args) {
		for(int j = 0 ; j<5 ; j++){ //  控制行数
			for(int i = 0 ; i<5 ; i++){ // 控制列数
				System.out.print("*");
			}  // *****
			//换行
			System.out.println();
		}
	}
}
```

```java
/*
需求： 在控制台上打印一个正立的直角三角形 。

*
**
***
****
*****

多行多列的图形。

行数 5行

列数： 会发生变化的.

分析列数:
	i = 0 ; i<5; j=0 ; j<=i 	1个星号
	i = 1 ; i<5 ;j=0 ; j<=1 	2个星号
	i = 2 ; i<5; j=0 ; j<=2    3个星号
	.....
*/
class Demo {
	public static void main(String[] args) {
		for(int i = 0 ; i< 5 ; i++){
			for(int j = 0 ; j<=i ; j++){ //控制列数 
				System.out.print("*");
			}
			//换行
			System.out.println();
		}
	}
}
```

```java
/*
需求： 打印一个倒立的直角三角形。

*****
****
***
**
*

5行

列数会发生变化
				      j<(5-i)
	i= 0 ; i<5; j=0 ; j<5 ;	五个星号
	i = 1; i<5; j=0 ; j<4; 	四个星号
	i = 2; i<5; j=0 ; j<3; 	三个星号
*/

class Demo {
	public static void main(String[] args) {
		for(int i = 0 ; i<5;  i++){
			for (int j = 0 ; j<(5-i)  ;j++ ){
				System.out.print("*");
			}
			//换行
			System.out.println();
		}
	}
}
```

```java
/*
需求： 打印一个九九乘法表.
*/
class Demo {
	public static void main(String[] args) {
		for(int i = 1 ; i<=9 ; i++){
			for(int j = 1 ; j<=i ; j++){ //控制列数 
				System.out.print(i+"*"+j+"="+i*j+ "\t");
			}
			//换行
			System.out.println();
		}
	}
}
```

### break和continue

#### break关键字

break适用范围：只能用于`switch`或者是`循环语句`中。

break作用：
	1. break用于switch语句的作用是结束一个switch语句。
	2. break用于循环语句中的作用是`结束当前所在的循环语句`。

break目前位于内层的for循环，如何才能让break作用于外层 的for循环。
	  `可以标记解决`
标记的命名只要符合标识符的命名规则即可。

```java
class Demo {
	public static void main(String[] args) {
		aaa:for(int j = 0 ; j<3 ; j++){        // j=0 外层for循环
			bbb:for(int i = 0 ; i< 2 ; i++){  // i=0 内层for循环
				System.out.println("hello world"); // 1	
				break aaa;
			}
			
		}
	}
}
```

#### continue关键字

continue的适用范围： `continue`只能用于循环语句。

continue的作用：continue的作用是跳过本次的循环体内容。`继续下一次`。

continue要注意的事项：
	1. 在一种情况下，continue后面不能跟有其他语句，因为是永远都无法执行到。
	2. continue 也可以配合标记使用的。

```java
class Demo{
	public static void main(String[] args) {
		
		for(int i = 0 ; i<5 ; i++){   // i=1  2
			if(i==1){
				continue;
			}
			System.out.println("hello "+i);    
		}
	}
}
/*
hello 0
hello 2
hello 3
hello 4
*/
```

```java
class Demo {
	public static void main(String[] args) {
		
		outer:for(int i = 0 ; i<3; i++){       // i= 0;  i =1 i=2 3
			inner:for(int j = 0 ; j<2 ; j++){  //j=0
				System.out.println("hello j-i = " + j+"-"+i);   //1 2 3
				continue outer;
			}
		}
	}
}
/*
hello j-i = 0-0
hello j-i = 0-1
hello j-i = 0-2
*/
```

```java
class Demo {
	public static void main(String[] args) {
	
		int sum = 0 ;
		for(int num = 1 ; num<=100 ; num++){
			if(num%2!=0){
				continue;  //如果是奇数就跳过本次循环。
			}
			sum  = sum+num;
		}
		System.out.println("总和："+ sum);
	}
}
```

### 转义字符

特殊字符使用`”\”`把其转化成字符的本身输出，那么使用”\”的字符称作为转移字符。

常见的转义字符有：

> \b	Backspace （退格键）
> \t	Tab    制表符(制表符的作用就是为了让一列对齐)  一个tab一般等于四个空格。
> \n	换行
> \r	回车  把光标移动到一行的首位置上。

注意： 如果是在windows系统上操作文件的时候需要换行，是需要`\r\n`一起使用的。
如果是在其他的操作系统上需要换行，仅需要\n即可。

```java
import java.io.*;
class Demo {
	public static void main(String[] args) throws Exception{
        
		//System.out.println("Hello哈哈\rworld!");

		File file = new File("F:\\a.txt");
		FileWriter  out = new FileWriter(file);
		out.write("大家好\r\n");
		out.write("你们好");
		out.close();
	}
}
```

### 函数（方法）

函数的作用： 提高功能代码的复用性。

#### 函数的定义格式

```java
修饰符  返回值类型  函数名(形式参数..){
	需要被封装的功能代码；
	return 结果;
}
```
> 分析函数：
> 	public static int add(){
> 		int a =2;
> 		int b =3;
> 		return a+b;
> 	}
>
> `修饰符`: public static
>
> `返回值类型`： int 。 返回值类型就是指函数运行完毕后，返回的结果的数据类型。
> 注意： 某些函数是没有结果返回给调用者的,那么这时候返回值类型是void。
>
> `函数名`： add   函数名的作用：如果需要调用该函数就需要使用的函数名。 函数名只要符合标识符的命名规则即可。 
> 		 函数名的命名规范： `首单词全部小写，其他单词的首字母大写，其他小写`。
>
> `形式参数`： 如果一个函数在运行的时候，存在着数据是要调用者确定 的，那么这时候就应该定义形式参数。
>
> `return `： 把一个结果返回给调用者

函数的特点：

 	1. 函数定义好之后，是需要被调用才会执行的。 `main函数是有jvm调用的`，不需要我们手动调用。
 	2. 函数的作用就是把一个功能代码给封装起来，已达到提高功能代码的`复用性`。
 	3. 函数定义好之后是`需要被调用`才会执行的。
 	4. 如果一个函数`没有返回值`返回给调用者，那么返回值类型必须是使用`void`表示。

#### return关键字

return 关键字的作用：
	>1. 返回数据给函数的调用者。
	>2. 函数一旦执行到了return关键字，那么该函数马上结束。 (能结束一个函数)

break关键字与return关键字的区别：

> `break`关键字是结束一个`循环`。
>
> `return`关键字是结束一个`函数`。

**注意：一个函数的返回值类型 是`void`，那么也可以出现`return`关键字，但是return关键字的后面不能有数据。**

```java
class Demo {

	public static void main(String[] args) {
	    //String result = getGrade(10);	//调用函数
	    //ystem.out.println("对应的等级是："+ result );
		//add(0,2);
		print();
	}

	public static void print(){
		for(int i = 0 ; i < 5;  i++){
			System.out.println("hello world");
			///break; //结束了当前的循环
			return ; //结束当前的函数
		}
		System.out.println("哈哈我能执行吗??");
	}

	//目前该函数的返回值类型是void，那么是否可以存在return关键字呢？
	public static void add(int a , int b){
		if(a==0){
			return; //结束一个函数
		}
		System.out.println("总和："+(a+b));
	}
}

```



#### 方法的四种类型

1、无参数无返回值的方法

```java
// 无参数无返回值的方法(如果方法没有返回值，不能不写，必须写void，表示没有返回值)
public void f1() {
    System.out.println("无参数无返回值的方法");
}
```

2、有参数无返回值的方法

```java
/**
* 有参数无返回值的方法
* 参数列表由零组到多组“参数类型+形参名”组合而成，多组参数之间以英文逗号（,）隔开，形参类型和形参名之间以英文空格隔开
*/
public void f2(int a, String b, int c) {
    System.out.println(a + "-->" + b + "-->" + c);
}
```

3、有返回值无参数的方法

```java
// 有返回值无参数的方法（返回值可以是任意的类型,在函数里面必须有return关键字返回对应的类型）
public int f3() {
    System.out.println("有返回值无参数的方法");
    return 2;
}
```

4、有返回值有参数的方法

```java
// 有返回值有参数的方法
public int f4(int a, int b) {
    return a * b;
}
```

5、return 在无返回值方法的特殊使用

```java
// return在无返回值方法的特殊使用
public void f5(int a) {
    if (a>10) {
    return;//表示结束所在方法 （f5方法）的执行,下方的输出语句不会执行
}
    System.out.println(a);
}

```

#### example-Code

```java
class Demo {
	public static void main(String[] args){
        
		//int sum = add(2,3); //调用了add名字 的函数。
		//System.out.println("结果："+ sum);
		
		add(2,3);
	}
	
/*
	//做加法功能的函数
	public static int add(int a,int b){ //int a,int b形式参数: 形式参数的值是交给调用者确定                                            的。
		return a+b;
	}
*/
    
    //需求： 定义一个函数做加法功能，不需要返回结果给调用者，直接输出结果即可。
	public static  void add(int a ,int b){
		int sum = a+b;
		System.out.println("结果："+ sum);
	}
}
```



```java
class Demo {
	public static void main(String[] args) {
		int max = getMax(14,5); //调用了函数   实际参数
		System.out.println("最大值："+ max);
		getMax2(3,7);
	}
	
	//需求2：定义一个函数比较两个int类型的数据大小， 不需要把最大值返回给调用者,直接打印即可。
	public static void getMax(int a, int b){
		int max = 0; //定义一个变量用于保存最大值的
		if(a>b){
			max = a;
		}else{
			max = b;
		}
		System.out.println("最大值："+ max);
	}

	//需求1： 定义一个函数比较两个int类型的数据大小， 把最大值返回给调用者。
	public static int  getMax1(int a, int b){  // 形式参数
		int max = 0; //定义一个变量用于保存最大值的
		if(a>b){
			max = a;
		}else{
			max = b;
		}
		return max;	//把结果返回给调用者
	}
}
```

```java
class Demo {
	public static void main(String[] args) {
		//String result = getGrade(189);
		//System.out.println(result);
		print(7);
	}

	//需求2： 定义一个函数打印一个乘法表，不需要返回任何数据。 
	public static void  print(int row){
		for(int i = 1 ; i<= row ; i++){
			for (int j = 1 ;j<=i  ;j++ ){
				System.out.print(i+"*"+j+"="+i*j+"\t");
			}
			//换行
			System.out.println();
		}
	}

	//需求1： 定义一个函数判断一个分数的等级，把分数的等级返回给调用者。
	public static String getGrade2(int score){
		String grade = "";	//定义一个变量存储等级
		if(score>=90&&score<=100){
			grade = "A等级";
		}else if(score>=80&&score<=89){
			grade = "B等级";
		}else if(score>=70&&score<=79){
			grade = "C等级";
		}else if(score>=60&&score<=69){
			grade = "D等级";
		}else if(score>=0&&score<=59){
			grade = "E等级";
		}else{
			grade = "补考等级";
		}
		return grade;	//把等级返回给调用者
	}
    
    public static String  getGrade(int score){//未知的参数定义在形参中,由函数的调用者确定。
		if(score>=90&&score<=100){
			return "A等级";
		}else if(score>=80&&score<=89){
			return "B等级";
		}else if(score>=70&&score<=79){
			return "C等级";
		}else if(score>=60&&score<=69){
			return "D等级";
		}else if(score>=0&&score<=59){
			return "E等级";
		}		
	}
}
```

#### 函数的重载

`函数的重载`：在一个类中出现两个或者两个以上的同名函数，这个称作为函数的重载。

​                         重载就是同样的一个方法能够根据输入数据的不同，做出不同的处理

函数重载的要求：
	>1. **`函数名一致`。**
	>2. **`形参列表不一致`。（参数类型不同、个数不同、顺序不同）**
	>3. **`与函数的返回值类型是无关的`。**
	>4. **`与修饰符无关`**。

**函数重载的作用： 同一个函数名可以出现了不同的函数，以应对不同个数或者不同数据类型的参数。**

```java
class Demo {
	public static void main(String[] args) {
		//System.out.println("Hello World!");
		//add1(1,2);
		add(1,2.0);
	}

    // 这些函数都是在做加法运算。
	public static double add(int a, int b){
		System.out.println("两个参数的总和: "+ (a+b));
		return 3.14;
	}
	
	//重复定义
	public static int add(int a, double b){
		System.out.println("double参数的总和: "+ (a+b));
		return 12;
	}
	
	public static void add(int a , int b , int c){
		System.out.println("三个参数的总和: "+ (a+b+c));
	}
    
	public static void add(int a , int b , int c,int d){
		System.out.println("四个参数的总和: "+ (a+b+c+d));
	}
}
```

### main函数的详解

`public `： 公共的。 权限是最大，在任何情况下都可以访问。 private  
		原因： 为了保证让jvm在任何情况下都可以访问到main方法。

`static`:  静态。静态可以让jvm调用main函数的时候更加的方便。不需要通过对象调用。

`void`:  没有返回值。因为返回的数据是给 jvm，而jvm使用这个数据是没有意义的。所以就不要了。

`main`: 函数名。   注意： `main并不是关键字,`只不过是jvm能识别的一个特殊的函数名而已

`arguments` ：担心某些程序在启动需要参数。

```java
import java.util.*;
class Demo {
	public static  void main(String[] args) {
        
		System.out.println("数组的长度："+ args.length);
        
		for(int i = 0 ; i <args.length ; i++){
			System.out.print(args[i]+",");
		}
		Scanner scanner = new Scanner(System.in);
	}
}

E:\smile>java Demo  1 1 1 1
数组的长度：4
1,1,1,1,
```



### 数组

#### 数组定义

数组是存储`同一种数据类型`数据的集合容器。

- **动态方式（指定数组的长度）**

定义可以存储3个整数的数组容器，代码如下：

```java
数据类型[]  变量名 = new 数据类型[长度];

int[] arr = new int[3];
```

分析数组：   
左边： int[] arr    声明了一个int类型的的数组变量，变量名为arr。
		`int `: 表示该数组容器只能存储int类型的数据。
		`[] `： 这是一个数组类型。
		`arr` ： 变量名.

右边：new int[50]; 创建了一个长度为50的int类型数组对象。
		`new`： 创建数组对象的关键字。
		`int`:  表示该数组对象只能存储int类型数据。
		`[]`： 表示是数组类型。
		`50` : 该数组最多能存储50个数据。数组的容量。



- 静态方式（指定数组的元素)

定义存储1，2，3，4，5整数的数组容器。

```java
数据类型[] 数组名 = new 数据类型[]{元素1,元素2,元素3...};

int[] arr = new int[]{1,2,3,4,5};
```

**或者省略格式：（`不能先声明后赋值，只能声明的同时赋值`）**

定义存储1，2，3，4，5整数的数组容器。

```java
数据类型[] 数组名 = {元素1,元素2,元素3...};

int[] arr = {1,2,3,4,5};
```

```java
class Demo {
	public static void main(String[] args) {	
        
		//动态初始化
		//int[] arr = new int[10];
        
		//静态初始化
		//int[] arr = {10,20,30,40,50};
		//for(int index = 0 ; index<arr.length ; index++){
			 System.out.print(arr[index]+",");
		//}
        
		int[] arr = new int[50];
		Scanner scanner = new Scanner(System.in);
		for(int i  = 0 ; i< arr.length ; i++){
			arr[i] = scanner.nextInt();
		}
	}
}
```



**注意:**

> 1.如果使用静态方式创建数组,那么系统会根据元素的个数自动计算数组的长度
>
>  2.静态方式创建数组右边的中括号里面不能写长度
>
>  3.静态方式的省略格式创建数组不能先声明后赋值,只能声明的同时直接赋值

**数组的特点：**

	>1. 数组只能存储同一种 数据类型的数据。
	>2. 数组是会给存储到数组中 的元素分配一个索引值的，索引值从0开始，最大的索引值是length-1；
	>3. 数组一旦初始化，长度固定。
	>4. 数组中的元素与元素之间的内存地址是连续的。

数组的好处： 对分配到数组对象中每一个数据都分配一个编号（索引值、角标、下标）,索引值的范围是`从0开始`，最大是： `长度-1`

数组的有一个`length `的属性，可以查看数组 的容量。`arr.length`

```java
class Demo {
	public static void main(String[] args) {
		//定义一个数组
		int[] arr = new int[4];
		arr[0] = 10;
		arr[1] = 30;
		arr[2] = 50;
		arr[3] = 90; 
        
		//数组的有一个length 的属性，可以查看数组 的容量。
		//System.out.println("数组的容量："+ arr.length);	
		//System.out.println("arr[2] = "+ arr[2]);
		
		//查看数组中的所有数据。
		for(int index = 0 ; index<arr.length ; index++){
			System.out.println(arr[index]);
		}
	}
}

```

![数组的内存分析1](images\数组的内存分析1.png)

```java
class Demo {
	public static void main(String[] args) {
		
		int[] arr1 = new int[2];
		int[] arr2 = new int[2];
		arr1[1] = 10;
		arr2[1] = 20;
		System.out.println(arr1[1]);  
        
		/*
		
		int[] arr1 = new int[2];
		arr1[1] = 100;
		int[] arr2 = arr1;
		arr1[1] = 10;
		arr2[1] = 20;
		System.out.println(arr1[1]);  // 20 20 20

		*/

	}
}
```

![数组的内存分析2](images\数组的内存分析2.png)



#### 数组中最常见的问题

`NullPointerException `空指针异常

- 原因： 引用类型变量没有指向任何对象，而访问了对象的属性或者是调用了对象的方法。

`ArrayIndexOutOfBoundsException `索引值越界。

- 原因：访问了不存在的索引值。

```java
class Demo {
	public static void main(String[] args) 
	{
		/*
		
		int[] arr = new int[2];
		arr = null ;  //null 让该变量不要引用任何的对象。 不要记录任何 的内存地址。
		arr[1] = 10;
		System.out.println(arr[1]);
	
		*/

		int[] arr = new int[4];
		arr[0] = 10;
		arr[1] = 30;
		arr[2]  =40;
		arr[3] = 50;
        
		//System.out.println(arr[4]); //访问索引值为4的内存空间存储的值。
			
		for(int index = 0 ; index<=arr.length ; index++){
			System.out.print(arr[index]+",");
		}
	}
}
```

#### example

```java
class Demo {
	public static void main(String[] args) 
	{
		int[] arr = {-12,-14,-5,-26,-4};
		int max = getMax(arr);
		System.out.println("最大值："+ max); 
	}

	public static int  getMax(int[] arr){
		int max = arr[0]; //用于记录最大值
		for(int i = 1 ; i < arr.length ; i++){
			if(arr[i]>max){  //如果发现有元素比max大，那么max变量就记录该元素。
				max = arr[i];
			}
		}
		return max;
	}


}
```

##### 选择排序
`选择排序`(直接排序)：使用一个元素与其他 的元素挨个比较一次，符合条件交换位置。

```java
class Demo {
	
	public static void main(String[] args) {
		int[] arr = {12,5,17,8,9};  //对于5元素的数组，只需要找出4个最大值就可以排序了。
		selectSort(arr);
	}

	public static void selectSort(int[] arr){
		
		//把最大值放在首位置。
		for(int j = 0; j<arr.length-1; j++){  //  控制的是轮数。
			for(int i = j+1 ; i<arr.length ; i++){ // 找出最大值
				if(arr[i]>arr[j]){
					//交换位置
					int temp = arr[i];
					arr[i] = arr[j];
					arr[j] = temp;
				}
			}
		}
		

/*
		//把老二放在第二个位置
		for(int i = 2  ; i< arr.length ; i++){
			if(arr[i]>arr[1]){
				int temp = arr[i];
				arr[i] = arr[1];
				arr[1] = temp;
			}
		}
		
		//把老三放在第三个位置
		for(int i = 3  ; i< arr.length ; i++){
			if(arr[i]>arr[2]){
				int temp = arr[i];
				arr[i] = arr[2];
				arr[2] = temp;
			}
		}

		//把老四放在第四个位置
		for(int i = 4  ; i< arr.length ; i++){
			if(arr[i]>arr[3]){
				int temp = arr[i];
				arr[i] = arr[3];
				arr[3] = temp;
			}
		}

*/
		//遍历数组，查看效果
		System.out.print("目前的元素：");
		for (int i = 0 ; i<arr.length  ;i++){
			System.out.print(arr[i]+",");
		}
	}
}
```

##### 冒泡排序
`冒泡排序`：冒泡排序的思想就是使用相邻的两个 元素挨个比较一次，符合条件交换位置。

```java
class Demo {
	public static void main(String[] args) {
		int[] arr = {12,8,17,5,9}; // 最大的索引值: 4   容量：5 
 		bubbleSort(arr);
	}
    
	public static void bubbleSort(int[] arr){
		// 把最大值放在最后一个位置
		for(int j = 0 ; j<arr.length-1 ; j++){ //控制轮数
			for(int i = 0 ; i<arr.length-1-j  ; i++){  // 找出一个最大值  
				//相邻的元素比较
				if(arr[i]>arr[i+1]){
					int temp  = arr[i];
					arr[i] = arr[i+1];
					arr[i+1] = temp;
				}
			}
		}

/*
		//把老二放在倒数第二个位置上。
		for(int i = 0 ;  i <arr.length-1-1 ; i++){
			if(arr[i]>arr[i+1]){
				int temp = arr[i];
				arr[i] = arr[i+1];
				arr[i+1] = temp;
			}	
		}

	
		//把老三放在倒数第三个位置上。
		for(int i = 0 ;  i <arr.length-1-2 ; i++){
			if(arr[i]>arr[i+1]){
				int temp = arr[i];
				arr[i] = arr[i+1];
				arr[i+1] = temp;
			}	
		}

		//把老四放在倒数第四个位置上。
		for(int i = 0 ;  i <arr.length-1-3 ; i++){
			if(arr[i]>arr[i+1]){
				int temp = arr[i];
				arr[i] = arr[i+1];
				arr[i+1] = temp;
			}	
		}

*/

		//遍历数组，查看效果
		System.out.print("目前的元素：");
		for (int i = 0 ; i<arr.length  ;i++){
			System.out.print(arr[i]+",");
		}	
	}
}
```

##### 二分(折半)
折半查找法(二分法): 使用前提必需是有序的数组。

```java
//需求：定义一个函数接收一个数组对象和一个要查找的目标元素，函数要返回该目标元素在
//数组中的索引值，如果目标元素不存在数组中，那么返回-1表示。
class Demo {
	public static void main(String[] args) {
		int[] arr = {12,16,19,23,54};
		//int index = searchEle(arr,23);
		int index = halfSearch(arr,116);
		System.out.println("元素所在的索引值是："+ index);
	}
		
	public static int halfSearch(int[] arr, int target){
		//定义三个变量分别记录最大、最小、中间的查找范围索引值
		int max = arr.length-1;
		int min = 0;
		int mid = (max+min)/2;
		while(true){
			if(target>arr[mid]){
				min = mid+1;
			}else if(target<arr[mid]){
				max = mid -1;
			}else{
				//找到了元素
				return mid;
			}
			//没有找到的情况
			if (max<min){
				return -1;
			}
			//重新计算中间索引值
			mid = (min+max)/2;
		}
	
	}
	public static int searchEle(int[] arr, int target){
		for(int i = 0 ; i<arr.length ; i++){
			if(arr[i]==target){
				return i;
			}
		}
		return -1;
	}
}
```

##### 翻转数组

```java
/*
需求： 定义 一个函数接收一个char类型的数组对象,然后翻转数组中的元素。
	
char[] arr = {'a','b','c','d','e'};
*/
class Demo {
	public static void main(String[] args) {
		char[] arr = {'a','b','c','d','e'};
		reverse(arr);
	}

	public static void reverse(char[] arr){
		for(int startIndex = 0 ,endIndex = arr.length-1 ;   startIndex<endIndex ; startIndex++,endIndex--){
			char temp = arr[startIndex];
			arr[startIndex] = arr[endIndex];
			arr[endIndex] = temp;
		}
2w
		//遍历数组，查看效果
		System.out.print("目前的元素：");
		for (int i = 0 ; i<arr.length  ;i++){
			System.out.print(arr[i]+",");
		}

	}
}
```

##### 数组模拟堆栈的存储方式
```java
import java.util.Arrays;

/*
 内存泄露
 
需求：编写一个类使用数组模拟堆栈的存储方式。  

堆栈存储特点： 先进后出，后进先出。

注意： 不再使用的对象，应该不要让变量指向该对象，要让该对象尽快的被垃圾回收期回收。
 */

class StackList{

	Object[] elements;

	int index = 0 ; //当前的索引值

	public StackList(){
		this.elements = new Object[3];
	}

	//添加内容
	public void add(Object o){
		//添加元素之前应该要先检查是否容量够用。
		ensureCapcity();
		elements[index++] = o;
	}

	//出栈: 删除集合的元素，并且返回。
	public Object pop(){
		int tempIndex = --index;
		Object o = elements[tempIndex];
		elements[tempIndex] = null; //让该位置不再 引用着指定的对象,让垃圾回收期赶快回收该垃圾。
		return o;
	}

	//检查当前的数组使用够用。
	public void ensureCapcity(){
		if(index==elements.length){
			//计算一个新的长度
			int newLength =	elements.length*2;
			elements = Arrays.copyOf(elements, newLength);
		}
	}
	//获取当前的元素 个数
	public int size(){
		return index;
	}
}

public class Demo {
	public static void main(String[] args) {

		StackList list = new StackList();
		list.add("狗娃");
		list.add("狗剩");
		list.add("铁蛋");
		list.add("美美");
		int size = list.size();
		for(int i = 0 ; i<size ; i++){
			System.out.println(list.pop());
		}
	}
}
```

#### Arrays类的常见操作

##### 排序 : `sort()`

```java
		// *************排序 sort****************
		int[] a = { 1, 3, 2, 7, 6, 5, 4, 9 };
		// sort(int[] a)方法按照数字顺序排列指定的数组。
		Arrays.sort(a);
		System.out.println("Arrays.sort(a):");
		for (int i : a) {
			System.out.print(i);
		}
		// 换行
		System.out.println();

		// sort(int[] a,int fromIndex,int toIndex)按升序排列数组的指定范围
		int[] b = { 1, 3, 2, 7, 6, 5, 4, 9 };
		Arrays.sort(b, 2, 6);
		System.out.println("Arrays.sort(b, 2, 6):");
		for (int i : b) {
			System.out.print(i);
		}
		// 换行
		System.out.println();

		int[] c = { 1, 3, 2, 7, 6, 5, 4, 9 };
		// parallelSort(int[] a) 按照数字顺序排列指定的数组(并行的)。同sort方法一样也有按范围的排序
		Arrays.parallelSort(c);
		System.out.println("Arrays.parallelSort(c)：");
		for (int i : c) {
			System.out.print(i);
		}
		// 换行
		System.out.println();

		// parallelSort给字符数组排序，sort也可以
		char d[] = { 'a', 'f', 'b', 'c', 'e', 'A', 'C', 'B' };
		Arrays.parallelSort(d);
		System.out.println("Arrays.parallelSort(d)：");
		for (char d2 : d) {
			System.out.print(d2);
		}
		// 换行
		System.out.println();

```

在做算法面试题的时候，我们还可能会经常遇到对字符串排序的情况,`Arrays.sort()` 对每个字符串的特定位置进行比较，然后按照升序排序。

```java
String[] strs = { "abcdehg", "abcdefg", "abcdeag" };
Arrays.sort(strs);
System.out.println(Arrays.toString(strs));//[abcdeag, abcdefg, abcdehg]
```

##### 查找 : `binarySearch()`

```java
		// *************查找 binarySearch()****************
		char[] e = { 'a', 'f', 'b', 'c', 'e', 'A', 'C', 'B' };
		// 排序后再进行二分查找，否则找不到
		Arrays.sort(e);
		System.out.println("Arrays.sort(e)" + Arrays.toString(e));
		System.out.println("Arrays.binarySearch(e, 'c')：");
		int s = Arrays.binarySearch(e, 'c');
		System.out.println("字符c在数组的位置：" + s);
```

##### 比较: `equals()`

```java
		// *************比较 equals****************
		char[] e = { 'a', 'f', 'b', 'c', 'e', 'A', 'C', 'B' };
		char[] f = { 'a', 'f', 'b', 'c', 'e', 'A', 'C', 'B' };
		/*
		* 元素数量相同，并且相同位置的元素相同。 另外，如果两个数组引用都是null，则它们被认为是相等的 。
		*/
		// 输出true
		System.out.println("Arrays.equals(e, f):" + Arrays.equals(e, f));
```

##### 填充 : `fill()`

```java
		// *************填充fill(批量初始化)****************
		int[] g = { 1, 2, 3, 3, 3, 3, 6, 6, 6 };
		// 数组中所有元素重新分配值
		Arrays.fill(g, 3);
		System.out.println("Arrays.fill(g, 3)：");
		// 输出结果：333333333
		for (int i : g) {
			System.out.print(i);
		}
		// 换行
		System.out.println();

		int[] h = { 1, 2, 3, 3, 3, 3, 6, 6, };
		// 数组中指定范围元素重新分配值
		Arrays.fill(h, 0, 2, 9);
		System.out.println("Arrays.fill(h, 0, 2, 9);：");
		// 输出结果：993333666
		for (int i : h) {
			System.out.print(i);
		}
```

##### 转列表 `asList()`

```java
		// *************转列表 asList()****************
		/*
		 * 返回由指定数组支持的固定大小的列表。
		 * （将返回的列表更改为“写入数组”。）该方法作为基于数组和基于集合的API之间的桥梁，与Collection.toArray()相结合 。
		 * 返回的列表是可序列化的，并实现RandomAccess 。
		 * 此方法还提供了一种方便的方式来创建一个初始化为包含几个元素的固定大小的列表如下：
		 */
		List<String> stooges = Arrays.asList("Larry", "Moe", "Curly");
		System.out.println(stooges);
```

##### 转字符串 `toString()`

```java
		// *************转字符串 toString()****************
		/*
		* 返回指定数组的内容的字符串表示形式。
		*/
		char[] k = { 'a', 'f', 'b', 'c', 'e', 'A', 'C', 'B' };
		System.out.println(Arrays.toString(k));// [a, f, b, c, e, A, C, B]
```

##### 复制 `copyOf()`

```java
        // *************复制 copy****************
		// copyOf 方法实现数组复制,h为数组，6为复制的长度
		int[] h = { 1, 2, 3, 3, 3, 3, 6, 6, 6, };
		int[] i = Arrays.copyOf(h, 6);
		System.out.println("Arrays.copyOf(h, 6);：");
		// 输出结果：123333
		for (int j : i) {
			System.out.print(j);
		}
		// 换行
		System.out.println();
		// copyOfRange将指定数组的指定范围复制到新数组中
		int j[] = Arrays.copyOfRange(h, 6, 11);
		System.out.println("Arrays.copyOfRange(h, 6, 11)：");
		// 输出结果66600(h数组只有9个元素这里是从索引6到索引11复制所以不足的就为0)
		for (int j2 : j) {
			System.out.print(j2);
		}
		// 换行
		System.out.println();
```

#### 二维数组

 二维数组就是数组中的数组。

二维数组 的定义格式：
```java
数据类型[][] 变量名 = new 数据类型[长度1][长度2];

长度1：一条烟有多少盒。

长度2： 一盒烟有多少根。
```


二维数组 的初始化方式：
```java
动态初始化:
	
	数据类型[][] 变量名 = new 数据类型[长度1][长度2];
```


```java
静态初始化：

	数据类型[][]  变量名 = {{元素1,元素2...},{元素1,元素2...},{元素1,元素2...} ..}
```
```java
class Demo {
	public static void main(String[] args) {	
		//定义了一个二维数组
		int[][] arr = new int[3][4];
		arr[1][1] = 100;
		/*
		System.out.println("二维数组的长度："+ arr.length);  // 3
		System.out.println("二维数组的长度："+ arr[1].length); //
		*/
		System.out.println("数组的元素："+ arr[1][1]);
	}
}
```

```java
/*
静态初始化：
	数据类型[][]  变量名 = {{元素1,元素2...},{元素1,元素2...},{元素1,元素2...} ..}
*/
class Demo{
	public static void main(String[] args) {
		int[][] arr = {{10,11,9},{67,12},{33,35,39,40}};
		//遍历二维数组
		for(int i = 0;  i <arr.length ; i++){
			for(int j = 0 ; j<arr[i].length ; j++){
				System.out.print(arr[i][j]+",");
			}
			//换行
			System.out.println();
		}
	}
}
```

#### 编写一个数组的工具类

```java
//数组工具类
class ArrayTool{
	public static String toString(int[] arr){
		String result  = "";
		for(int i = 0;  i < arr.length ; i++){
			if (i==0){
				result+="["+arr[i]+",";
			}else if(i==(arr.length-1)){
				result+= arr[i]+"]";
			}else{
				result+=arr[i]+",";
			}
		}
		return result;
	}
    
	public static void sort(int[] arr){
		for(int i = 0; i < arr.length-1 ; i++){
			for(int j = i+1 ; j<arr.length ; j++){
				if(arr[i]>arr[j]){
					int temp = arr[i];
					arr[i] = arr[j];
					arr[j] = temp;
				}
			}
		}	
	}
}

class Demo {
	public static void main(String[] args) {
        
		int[] arr = {12,1,456,165};
        
		ArrayTool tool = new ArrayTool();

		ArrayTool.sort(arr);
		String info = ArrayTool.toString(arr);
		System.out.println("数组的元素："+ info);
	}
}

```



## 面向对象

### 类和对象

对象：真实存在唯一的事物。

类： 实际就是对某种类型事物的`共性属性`与`行为的抽取`。 `抽象的概念`

面向对象的计算机语言核心思想： 找适合的对象做适合的事情。

如何找适合的对象：

	>1.  sun已经定义好了很多的类，我们只需要认识这些类，我们就可以通过这些类创建对象使用。 
	>2.  自定义类，通过自定义类来创建对象使用。

自定义类创建对象的三步骤：
 	1. 自定义类。

```java
class 类名{
	事物的公共属性使用成员变量描述。
				
	事物的公共行为使用函数描述。
}

2. 通过自定义类创建对象。
		格式： 
			类名 变量名 =  new 类名();
		
3. 访问(设置)对象的属性或者调用对象的功能。
	1.访问对象属性的格式：
		对象.属性名.
	2.设置对象的属性：
		对象.属性名 = 数据。
	3. 调用对象的功能
		对象.函数名();
```
```java
//车类
class Car{

	//事物的公共属性使用成员变量描述。
	String	name; //名字的属性

	String	color; //颜色属性

	int wheel;	//轮子数

	//事物的公共行为使用函数描述。
	public void run(){
		System.out.println(name+"飞快的跑起来啦...");
	}
}

class Demo {
	public static void main(String[] args) {	
		/*
		//使用了Car类声明了一c变量， c变量指向了一个车对象。
		Car	c = new Car(); 
		//设置车对象的属性值。
		c.name = "BMW";
		c.color = "白色";
		c.wheel = 4;
		//访问车对象的属性值
		System.out.println("名字："+ c.name+" 颜色:"+ c.color+" 轮子数："+c.wheel);
		c.run();
		*/

		Car	c1 = new Car(); 
		c1.name = "宝马";
		Car c2 = new Car();
		c2.name = "大众";
		c1 = c2; 
		System.out.println("名字："+ c1.name);  //大众  大众 \大众
	}
}
```

需求： 使用java描述一个车与修车厂两个事物， 车具备的公共属性:轮子数、 名字、 颜色 ，还
具备跑的功能行为。跑之前要检测轮子是否少于了4个，如果少于了4个，那么要送到修车厂修理，
修车厂修理之后，车的轮子数要补回来4个。 然后车就继续的跑起来。

修车厂： 具备公共属性： 名字、 地址、 电话。
		       公共的行为： 修车。

```java
//车类
class Car{
	//事物的公共属性使用成员变量描述	
	String name ; // 名字
	String color; //名字
	int wheel; //轮子数
	
	//事物的公共行为使用函数描述
	public void run(){
		if(wheel>=4){
			System.out.println(name+wheel+"个轮子飞快跑起来..");
		}else{
			System.out.println(name+"不够4个轮子了，赶快去修理");
		}
	}
}

//修车厂
class CarFactory{
	String name;//名字
	String address ;	//地址
	String tel;	//电话

	//修车公共行为 ------ 返回值类型、 未知的参数
	public void repair(Car c){
		if(c.wheel>=4){
			System.out.println("告诉你，费了很大力气修好了，给钱");
		}else{
			c.wheel = 4;
			System.out.println("修好了，给钱!!");	
		}
	}
}

class Demo {
	public static void main(String[] args) {	
        
		Car c = new Car();
		//给车对象赋予属性值
		c.name = "陆丰";
		c.color = "黑色";
		c.wheel = 4;
		
		for(int i = 0 ; i<100 ; i++){
			c.run();
		}
		c.wheel = 3;
		c.run();

		//创建修车厂对象
		CarFactory f = new CarFactory();
		//给修车厂赋予属性值
		f.name = "集群宝修车厂";
		f.address = "韵泰商业广场一楼";
		f.tel = "020-1234567";

		//调用修车的修车
		f.repair(c);//0x98
		
		c.run();		
	}
}
```

### 成员变量与局部变量

成员变量与局部变量的区别:

- 定义的位置上区别：

  > 1. 成员变量是定义在`方法之外`，`类之内`的。
  > 2. 局部变量是定义在`方法之内`。

#### 作用上的区别：

  > 1. 成员变量的作用是用于描述一类事物的`公共属性`的。
  >
  > 2. 局部变量的作用就是提供一个变量给方法内部使用而已。

- 生命周期区别:

  >1. 随着对象 的创建而存在，随着对象的消失而消失。
  >2. 局部变量在调用了对应的方法时执行到了创建该变量的语句时存在，局部变量一旦出了自己的作用域那么马上从内存中消失。

- 初始值的区别：

  >1. 成员变量是有默认的初始值。
  >			数据类型     默认的初始值
  >			int                0
  >			float              0.0f
  >			double              0.0
  >			boolean            false
  >			char                 ' '
  >			String（引用数据类型） null
  >2. 局部变量是没有默认的初始值的，必须要先初始化才能使用。

```java
class Person {
	
	String name; //成员变量

	public void eat(){
		int age ;    //局部变量
		age = 12;
		for(int i =  0 ; i< 5 ; i++){
			System.out.println("hello world"+ age);
		}
	}
}

class Demo {
	public static void main(String[] args) {
		Person p1 =  new Person();
		p1.eat();
	}
}
```

### 匿名对象

`匿名对象`：没有引用类型变量指向的对象称作为匿名对象。

- 匿名对象要注意的事项：
  	>1. 我们一般不会给匿名对象赋予属性值，因为永远无法获取到。
    >2. 两个匿名对象永远都不可能是同一个对象。

- 匿名对象好处：`简化书写`。


- 匿名对象的应用场景：
  	>1. 如果一个对象需要调用一个方法一次的时候，而调用完这个方法之后，该对象就不再使用了，这时候可以使用匿名对象。
    	2. 可以作为实参调用一个函数。

  ```java
  //学生类
  class Student {
  	
  	int num; //学号
  
  	String name; //名字
  	
  	public void study(){
  		System.out.println("好好学习，为将来称为高帅富做准备!");
  	}
  }
  
  class Demo{
  	public static void main(String[] args) {
  		//创建一个学生对象
  		//Student s = new Student();
  		//new Student().name = "狗娃"; //匿名对象 
  		//System.out.println(Student().name); //null
          
  		System.out.println(new Student() == new Student()) ;  // "==" 用于引用类型变量时，比较的是内存地址。判断两个 对象是否为同一个对象
  		
  		需求： 调用Student的study方法。
  		Student s = new Student();
  		s.study();
  		new Student().study();
  	}
  }
  
  ```

  

### 面向对象三大特征

#### 封装

权限修饰符：权限修饰符就是控制变量可见范围的。

> `public` :  公共的。 public修饰的成员变量或者方法任何人都可以直接访问。
>
> `private` ： 私有的， private修饰的成员变量或者方法只能在本类中进行直接访问。

封装的步骤：
	>1. 使用private修饰需要被封装的属性。
	>2. 提供一个公共的方法设置或者获取该私有的成员属性。
			 命名规范：
					set属性名();
					get属性名(); 

规范 ： 在现实开发中一般实体类的所有成员属性（成员变量）都要封装起来。

实体类：实体类就是用于描述一类 事物的就称作为实体类。

封装的好处：
	>1. 提高数据的安全性。
	>2. 操作简单。 
	>3. 隐藏了实现。

```java
 class Member{
	
	public	String name; //名字

	private	String sex; //性别

	public	int salary; //薪水
	
	//定义一个公共的方法设置sex属性
	public void setSex(String s){
		if (s.equals("男")||s.equals("女")){ //注意： 如果比较两个字符串的内容是否一致，不要使用==比较， 使用equals方法。
			sex = s;
		}else{
			//默认是男
			sex = "男";
		}
	}

	//定义一个公共的方法获取sex属性
	public String getSex(){
		return sex;
	}

	//聊天
	public void talk(){
		System.out.println(name+"聊得非常开心");
	}
}

class Demo {
	public static void main(String[] args) {
		Member m = new Member();
		m.name="狗娃";
		m.setSex("女");
		m.salary  = 800;
		System.out.println("姓名："+ m.name+" 性别："+ m.getSex()+" 薪水："+ m.salary);
	}
}
```

```java
/计算器类
class Calculator {

	private int num1; //操作数1

	private int num2;  //操作数2
 
	private	char option ; //运算符
	
	//提供公共的方法设置属性值....					
	public void initCalculator(int n1 , int n2 , char o){
		num1 = n1;
		num2 = n2;
		if(o=='+'||o=='-'||o=='*'||o=='/'){
			option = o;
		}else{
			option = '+';	
		}
	}

	//计算的功能
	public void calculate(){
		switch(option){
			case '+':
				System.out.println("做加法运算,结果是："+(num1+num2));
				break;
			case '-':
				System.out.println("做减法运算,结果是："+(num1-num2));
				break;
			case '*':
				System.out.println("做乘法运算,结果是："+(num1*num2));
				break;
			case '/':
				System.out.println("做除法运算,结果是："+(num1/num2));
				break;
		}
	}
}

class Demo {
	public static void main(String[] args) {
		//创建了一个计算器对象
		Calculator c = new Calculator();
		//设置属性值
		c.initCalculator(1,2,'a');
		//调用计算器的计算功能
		c.calculate();
	}
}
```

需求： 目前存在数组：int[] arr = {0,0,12,1,0,4,6,0} ，编写一个函数
接收该数组，然后把该数组的0清空，然后返回一个不存在0元素的数组。

步骤：

	1. 计算机新数组的长度。 原来的数组长度-0的个数

```java
import java.util.*;
class Demo {
	public static void main(String[] args) {
		int[] arr = {0,0,12,1,0,4,6,0};
		arr = clearZero(arr);
		System.out.println("数组的元素："+Arrays.toString(arr));
	}

	public static  int[] clearZero(int[] arr){
		//统计0的个数
		int count = 0;	//定义一个变量记录0的个数
		for(int i = 0 ; i<arr.length ; i++){
			if(arr[i]==0){
				count++;
			}
		}

		//创建一个新的数组
		int[] newArr = new int[arr.length-count];
			
		int index  =0 ; //新数组使用的索引值
		//把非的数据存储到新数组中。
		for(int i = 0; i<arr.length ; i++){
			if(arr[i]!=0){
				newArr[index] = arr[i];
				index++;
			}
		}
		return newArr;
	}
}

```

#### 继承

整体与部分关系    has a 关系  

继承的关系       is a 关系


继承：继承是通过关键字`extends`体现的。

继承的格式：

```java
class 类名1 extends 类名2{

}
```

继承要注意的事项：
	1. 千万不要为了减少重复代码而去继承，只有真正存在着继承关系的时候才去继承。
	2. `父类私有的成员不能被继承`。
	3. `父类的构造函数不能被继承`。
	4. `创建子类对象时默认会先调用父类无参的构造函数`。

`java是单继承`
	一个类最多只能有一个直接的父类。但是有多个间接的父类。

```java
//人类 
class Person{
	
	String name;

	private	int age;

	public  Person(String name){
		this.name = name;
	}

	public Person(){
		System.out.println("Person类的构造方法被调用了....");
	}

	public void eat(){
		System.out.println(name+"在吃饭...");
	}
}

//学生类
class Student extends Person {  // Student 就称作为Person类的子类， Person类就称作为Student的父类(超类、基类)

	int num; //学号

	public Student(){
		System.out.println("Student类的构造方法被调用了....");
	}

	public void study(){
		System.out.println(name+"good good study , day day up");
	}	
}

class Demo {
	public static void main(String[] args) {
		Student s = new Student();
		/*
		s.name = "狗娃";
		System.out.println("名字："+ s.name);
		s.eat();
		*/
	}
}
```
疑问： 为什么要调用父类的构造方法啊？这样子做的意义在那？
	   调用父类 的构造方法是可以初始化从父类继承下去的属性的。
```java
class Fu{
	
	int x = 10;

	String name;

	public Fu(String name){
		this.name = name;
		System.out.println("Fu类d带参的构造方法...");
	}

	public Fu(){
		System.out.println("Fu类无参的构造方法...");
	}
}


class Zi extends Fu{

	int x = 20;

	public Zi(String name){
		super(name); //指定调用父类一个参数的构造函数。
	}

	public void print(){
		System.out.println("x1 = "+ x);
	}
}

class Demo {
	public static void main(String[] args) {
		Zi z = new Zi("大头儿子"); 
		System.out.println("name= "+z.name);
	}
}
```

	
```java
//球员类
class Player{
	
	int num; //编码

	String name;

	public Player(int num , String name){
		this.num = num;
		this.name = name;
	}

	public void run(){
		System.out.println(name+"开跑...");
	}
}

//球队类
class Team{

	String name;  //球队的名字

	Player p1;  //球员1
 
	Player p2;   //球员2

	Player p3;  //球员3

	public  Team(String name,Player p1,Player p2,Player p3){
		this.name = name;
		this.p1 = p1;
		this.p2 = p2;
		this.p3 = p3;
	}
	
	//开始比赛
	public void startGame(){
		System.out.println(name+"开赛啦！！");
	}
}

class Demo {
	public static void main(String[] args) {
        
		Player p1 = new Player(12,"梅西");
		Player p2 = new Player(7,"C罗");
		Player p3 = new Player(11,"内马尔");
		//球队
		Team t = new Team("恒大",p1,p2,p3);
		t.startGame();
		
		System.out.println("名字："+ t.p2.name);
	}
}
```

```java
class Ye{
	String name;
}

class Fu extends Ye{

}

class  Zi extends Fu{

}

class Demo {
	public static void main(String[] args) {
		Zi zi =new Zi();
		zi.name = "狗娃";
	}
}
```

#### 多态

多态：一个对象具备多种形态。(父类的引用类型变量指向了子类的对象

或者是接口的引用类型变量指向了接口实现类的对象)

多态的前提：必须存在`继承`或者`实现` 关系。

```java
动物  a  = new   狗();
```?

多态要注意 的细节：
	1.  多态情况下，子父类存在`同名的成员变量`时，访问的是`父类`的成员变量。
	2.  多态情况下，子父类存在`同名的非静态的成员函数`时，访问的是`子类`的成员函数。
	3.  多态情况下，子父类存在`同名的静态的成员函数`时，访问的是`父类`的成员函数。
	4.  多态情况下，不能访问子类特有的成员。

总结：多态情况下，子父类存在同名的成员时，访问的都是父类的成员，除了在同名非静态函数时才是访问子类的。

编译看左边，运行不一定看右边。

编译看左边：java编译器在编译的时候，会检查引用类型变量所属的类是否具备指定的成员，如果不具备马上编译报错。

```java
/动物类
abstract class Animal{
	String name;
	String  color = "动物色";
	public Animal(String name){
		this.name = name;
	}
	public abstract void run();

	public static void eat(){
		System.out.println("动物在吃..");
	}
}

//老鼠
class  Mouse extends Animal{
	String color = "黑色";
	public Mouse(String name){
		super(name);
	}
	public  void run(){
		System.out.println(name+"四条腿慢慢的走!");
	}
	public static void eat(){
		System.out.println("老鼠在偷吃..");
	}
	//老鼠特有方法---打洞
	public void dig(){
		System.out.println("老鼠在打洞..");
	}
}

class Fish extends Animal {
	public Fish(String name){
		super(name);
	}
	public  void run(){
		System.out.println(name+"摇摇尾巴游..");
	}
}


class Demo {
	public static void main(String[] args) {

		Mouse m = new Mouse("老鼠");
		System.out.println(m.color);
		
		//多态： 父类的引用类型变量指向子类的对象
		Animal a = new Mouse("老鼠");
		a.dig();
		a.eat();
	}	
}

```

多态的应用：
  1. 多态用于`形参类型`的时候，可以接收更多类型的数据 。
  2. 多态用于`返回值类型`的时候，可以返回更多类型的数据。

多态的好处： 提高了`代码的拓展性`。

需求1： 定义一个函数可以接收任意类型的图形对象，并且打印图形面积与周长。
```java
//图形类
abstract class MyShape{
	public abstract void getArea();
	public abstract void getLength();
}

class Circle extends MyShape{
	public static final double PI = 3.14;
	double r;
	public Circle(double r){
		this.r =r ;	
	}
	public  void getArea(){
		System.out.println("圆形的面积："+ PI*r*r);
	}
	public  void getLength(){
		System.out.println("圆形的周长："+ 2*PI*r);
	}
}

class Rect  extends MyShape{
	int width;
	int height;
	public Rect(int width , int height){
		this.width = width;
		this.height = height;
	}
	public  void getArea(){
		System.out.println("矩形的面积："+ width*height);
	}
	public  void getLength(){
		System.out.println("矩形的周长："+ 2*(width+height));
	}
}

class Demo {

	public static void main(String[] args) {
		
		Circle c = new Circle(4.0);
		print(c);

		Rect r = new Rect(3,4);
		print(r);

		MyShape m = getShape(0); //调用了使用多态的方法，定义的变量类型要与返回值类型一致。
		m.getArea();
		m.getLength();
	}

	//需求1： 定义一个函数可以接收任意类型的图形对象，并且打印图形面积与周长。
	public static void print(MyShape s){ // MyShpe s = new Circle(4.0);
		s.getArea();
		s.getLength();
	}

	// 需求2： 定义一个函数可以返回任意类型的图形对象。
	public static MyShape  getShape(int i){
		if (i==0){
			return new Circle(4.0);
		}else{
			return new Rect(3,4);
		}
	}
}
```

目前多态情况下不能访问子类特有的成员。

如果需要访问子类特有的成员，那么需要进行`类型强制转换`.

基本数据类型的转换
	小数据类型-------->大的数据类型      自动类型转换
	大数据类型--------->小数据类型       强制类型转换     小数据类型  变量名 = （小数据类型）大数据类型;

引用数据类型的转换
	小数据类型--------->大数据类型     自动类型转换。
	大数据类型--------->小数据类型      强制类型转换

类型转换最场景的问题： `java.lang.ClassCastException`。 强制类型转换失败。
```java
//动物类
abstract class Animal{
	String name;
	public Animal(String name){
		this.name = name;
	}
	public abstract void run();
}
//老鼠
class  Mouse extends Animal{
	public Mouse(String name){
		super(name);
	}
	public  void run(){
		System.out.println(name+"四条腿慢慢的走!");
	}
	//老鼠特有方法---打洞
	public void dig(){
		System.out.println("老鼠在打洞..");
	}
}
//鱼
class Fish extends Animal{
	public Fish(String name){
		super(name);
	}

	public  void run(){
		System.out.println(name+"摇摇尾巴游啊游 !");
	}
	//吹泡泡
	public void bubble(){
		System.out.println(name+"吹泡泡...!");
	}
}
class Demo {
	public static void main(String[] args) {
		/*
		Animal a = new Mouse("老鼠");  //多态
		//调用子类特有的方法
		Mouse m  = (Mouse)a;  //强制类型转换
		m.dig();
		*/

		Mouse m = new Mouse("米老鼠");
		Fish f = new Fish("草鱼");
		print(f);
	}

	//需求： 定义一个函数可以接收任意类型的动物对象，在函数内部要调用到动物特有的方法。
	public static void print(Animal a){ // Animal a   = new Mouse("米老鼠");
		if(a instanceof Fish){
			Fish f = (Fish)a;
			f.bubble();
		}else if(a instanceof Mouse){
			Mouse m = (Mouse)a;
			m.dig();
		}
	}
}
```

多态： 父类的引用类型变量指向了子类的对象或者是接口类型的引用类型变量指向了接口实现类的对象。

	实现关系下的多态：
```java
	接口  变量  = new  接口实现类的对象。	
```
```java
interface Dao{   //接口的方法全部都是非静态的方法。

	public void add();

	public void delete();
}
//接口的实现类
class UserDao implements Dao{
	public void add(){
		System.out.println("添加员工成功！！");
	}
	public void delete(){
		System.out.println("删除员工成功！！");
	}

}

class Demo {
	public static void main(String[] args) {
		//实现关系下的多态
		Dao d = new UserDao(); //接口的引用类型变量指向了接口实现类的对象。
		d.add();
	}
}
```
### 构造函数

**构造函数的作用： 给对应的对象进行`初始化`。**

- 构造函数的定义的格式：

```java
修饰符  函数名(形式参数){
	函数体...
}
```

- 构造函数要注意的细节：
  	>1. 构造函数 是没有返回值类型的。
  	>2. 构造函数的函数名必须要与类名一致。
  	>3. 构造函数并不是由我们手动调用的，而是在创建对应的对象时，jvm就会主动调用到对应的构造函数。
  	>4. 如果一个类没有显式的写上一个构造方法时，那么java编译器会为该类添加一个无参的构造函数的。
  	>5. 如果一个类已经显式的写上一个构造方法时,那么java编译器则 不会再为该类添加 一个无参 的构造方法。
  	>6. 构造函数是 可以在一个类中以函数重载的形式存在多个的。

疑问：创建对象时，jvm就会调用到对应的构造方法，那么我们以前没有学构造方法，那么
以前创建对象时，jvm是否 也会调用构造方法呢？如果有？构造方法从何而来呢？		

​		会调用， java编译器在编译的 时候给加上去的。

jdk提供了一个java开发工具(javap.exe)给我们进行反编译的。

- 
  javap 反编译工具的使用格式：
```java
javap -c -l -private 类名
```

疑问： java编译器添加 的无参构造方法的权限修饰符是什么？
		与类的权限修饰是一致的。

- 构造函数与普通 函数的区别：

  ​		1. 返回值类型的区别：

  ​				构造函数是`没有返回值类型`的，

  ​				普通函数是有返回值类型的，即使函数没有返回值，返回值类型也要写上void。

  ​		2. 函数名的区别：

  ​				构造函数的`函数名必须要与类名一致`，

  ​				普通函数的函数名只要符合标识符的命名规则即可。

  ​		3. 调用方式的区别：

  ​				构造函数是 在创建对象的时候由`jvm调用的`。

  ​				普通函数是由我们使用对象调用的，一个对象可以对象多次普通 的函数，

  ​		4. 作用上的区别：

  ​				构造函数 的作用用于`初始化一个对象`。

  ​				普通函数是用于描述一类事物的`公共行为`的。

```java
class Baby {
	
	int id; //身份证
 
	String  name;  //名字

	//构造函数
	public  Baby(int i , String n){
		id  = i;
		name = n;
		System.out.println("baby的属性初始化完毕！！");
	}

	//无参 的构造函数
	public Baby(){
		System.out.println("无参的构造函数被调用了..");
	}

	//哭
	public void cry(){
		System.out.println(name+"哇哇哭...");
	}	
}


class Demo {
	public static void main(String[] args) {	
        
		//创建一个baby对象
		Baby b1 = 	new Baby(110,"狗娃"); //婴儿诞生   白户
		System.out.println("编号："+ b1.id +" 姓名："+ b1.name);
		b1.cry();

		//黑户
		Baby b2 = new Baby();
		new Baby();

		b2.id = 112;
		b2.name = "狗剩";
		System.out.println("编号："+ b2.id +" 姓名："+ b2.name);
	}
}
```

```java
//要求：一旦创建一个员工对象 的时候，那么该员工对象就要对应 的属性值。
class Employee {
	
	int id;  //编号

	String name;  //名字

	int age; //年龄

	//构造函数
	public Employee(int a, String b , int c){
		id =a;
		name = b;
		age = c;
	}

	public void work(){
		System.out.println(name+"好好工作，努力挣钱!!");
	}
}

class Demo {
	public static void main(String[] args) 
	{
		//创建一个员工对象
		Employee e = new  Employee(110,"狗娃",20);
		System.out.println("编号："+ e.id+" 名字："+e.name +" 年龄："+ e.age);
	}
}
```

### 构造代码块

构造代码块的作用：`给对象进行统一的初始化`。

构造函数的作用： 给对应的`对象`进行初始化。


构造代码块的格式：
```java
{
	构造代码块
}
```

注意： 构造代码块的大括号必须位于`成员位置`上。

代码块的类别：
		1. 构造代码块。
		2. 局部代码块.   大括号位于`方法之内`。  作用：缩短局部 变量 的`生命周期`，节省一点点内存。
		3. 静态代码块  `static` 

```java
class Baby{
	
	int id; //身份证
 
	String  name;  //名字
	
	//构造代码块...
	{
		System.out.println("构造代码块的代码执行了......");  //
		
	}

	//带参构造函数
	public  Baby(int i , String n){
		id  = i;
		name = n;
	}
	
	//无参构造方法
	public Baby(){
	}

	public void cry(){
		System.out.println(name+"哇哇哭...");
	}	
}

class Demo {
	public static void main(String[] args) {
        
		Baby b1 = new Baby(110,"狗娃");  //  狗娃 狗剩 铁蛋
		System.out.println("编号："+ b1.id + " 名字："+b1.name);
		
		new Baby(112,"狗剩");
		new Baby();
	
		new Baby(110,"狗娃");
		new Baby(112,"狗剩");
		new Baby();
	}
}
```

结果

```java
构造代码块的代码执行了......
编号：110 名字：狗娃
构造代码块的代码执行了......
构造代码块的代码执行了......
构造代码块的代码执行了......
构造代码块的代码执行了......
构造代码块的代码执行了......
```

构造 代码块要注意的事项：
	1. java编译器编译一个java源文件的时候，会把`成员变量的声明语句提前至一个类的最前端`。
	2. `成员变量的初始化`工作其实都在在`构造函数`中执行的。
	3. 一旦经过java编译器编译后，那么`构造代码块`的代码块就会被`移动构造函数中`执行，是在构造函数之前执行的，`构造函数的中代码是最后执行 的`。
	4. 成员变量的显示初始化与构造代码块 的代码是按照当前代码的顺序执行的。

```java
class Demo  {	
	//构造函数
	public Demo(){   //构造函数
		i = 300000000;	
        System.out.println("hello"); 
	}
		
	//构造代码块   //构造代码块的初始化
	{
		i = 200000000;
        System.out.println("world"); 
	}
	
	int i = 100000000;	 //成员变量的显初始化

	public static void main(String[] args) {
		Demo d = new Demo();
		System.out.println("i = "+d.i); 
	}
}

---------------------
world
hello
i = 300000000
```

### this关键字

问题：存在`同名`的`成员变量`与`局部变量`时，在方法的内部访问的是局部变量(java 采取的是`就近原则的机制`访问的)

this关键字代表了`所属函数`的`调用者对象`。

* this关键字作用：
	1. 如果存在同名成员变量与局部变量时，在`方法内部默认是访问局部变量的数据`，可以通过`this关键字指定访问成员变量的数据`。
	2. **在一个`构造函数`中可以`调用另外一个`构造函数初始化对象。**

* this关键字调用其他的构造函数要注意的事项：

* this关键字要注意事项：
	1. 存在同名的成员变量与局部变量时，在方法的内部访问的是局部变量(java 采取的是“就近原则”的机制访问的。)
	2. 如果在一个方法中访问了一个变量，该变量只存在成员变量的情况下，那么`java编译器`会在该变量的 前面添加`this`关键字。

```java
class Animal{

	String name ;  //成员变量

	String color;

	public Animal(String n , String c){
		name = n;
		color = c;
        System.out.println("this:"+ this);
	}
	
	//this关键字代表了所属函数的调用者对象
	public void eat(){
		System.out.println("this:"+ this);
		String name = "老鼠"; //局部变量
		System.out.println(name+"在吃..."); //需求： 就要目前的name是成员变量的name.
		
	}
}

class Demo {
	public static void main(String[] args) {
        
		Animal dog = new Animal("狗","白色");  //现在在内存中存在两份name数据。
		Animal cat = new Animal("猫","黑色");
        dog.eat();
		cat.eat();
        System.out.println("dog:" + dog);
        System.out.println("cat:" + cat);
	}
}
```

结果

```java
this:Animal@15db9742
this:Animal@6d06d69c
this:Animal@15db9742
老鼠在吃...
this:Animal@6d06d69c
老鼠在吃...
dog:Animal@15db9742
cat:Animal@6d06d69c
```

```java
class Student{

	int id;  //身份证

	String name;  //名字

	//目前情况：存在同名的成员变量与局部变量，在方法内部默认是使用局部变量的。
	public Student(int id,String name){  //一个函数的形式参数也是属于局部变量。
        
		this(name); //调用了本类的一个参数的构造方法
		//this(); //调用了本类无参的 构造方法。
        
		this.id = id; // this.id = id 局部变量的id给成员变量的id赋值
		System.out.println("两个参数的构造方法被调用了...");
	}
	
	
	public Student(){
		System.out.println("无参的构造方法被调用了...");
	}

	public Student(String name){
		this.name = name;
		System.out.println("一个参数的构造方法被调用了...");
	}
}


class Demo {
	public static void main(String[] args) {
		Student s = new Student(110,"铁蛋");
		System.out.println("编号："+ s.id +" 名字：" + s.name);

	/*
		Student s2 = new Student("金胖子");
		System.out.println("名字：" + s2.name);
	*/
	}
}
```

结果

```java
一个参数的构造方法被调用了...
两个参数的构造方法被调用了...
编号：110 名字：铁蛋
```

```java
class Person{
		
	int id; //编号

	String name; //姓名
  
	int age;  //年龄

	//构造函数
	public Person(int id,String name ,int age){
		this.id  = id;
		this.name = name;
		this.age = age;
	}

	//比较年龄的方法
	public void compareAge(Person p2){
		if(this.age>p2.age){
			System.out.println(this.name+"大!");
		}else if(this.age<p2.age){
			System.out.println(p2.name+"大!");
		}else{
			System.out.println("同龄");
		}
	}
}

class Demo {

	public static void main(String[] args) {
		Person p1 = new Person(110,"狗娃",17);
		Person p2 = new Person(119,"铁蛋",9);
		p1.compareAge(p2);  //狗娃大!
	}
}
```

### static(静态)

`静态的成员变量`只会在`数据共享区`中`维护一份`，而`非静态成员变量`的数据会在`每个对象`中都维护一份的。。


```java
class Student{

	String name;
	
	//使用了static修饰country，那么这时候country就是一个共享的数据。

	static	String  country  = "中国";	//国籍
	
	//构造函数
	public Student(String name){
		this.name = name;
	}
}

class Demo {

	public static void main(String[] args) {
        
		Student s1 = new Student("张三");
		Student s2 = new Student("陈七");

		s1.country = "小日本";
		System.out.println("姓名："+s1.name+" 国籍："+ s1.country); //  中国   
		System.out.println("姓名："+s2.name+" 国籍："+ s2.country); // 小日本
	}
}
```

- static修饰成员变量 ：如果有数据需要被共享给所有对象使用时，那么就可以使用static修饰。


​		静态成员变量的访问方式：
​				方式1： 可以使用对象进行访问。
​						格式： `对象.变量名`。
​		
​				方式二： 可以使用`类名`进行访问。
​						格式： `类名.变量名`;

- 注意： 
           * 非静态的成员变量只能使用对象进行访问，不能使用类名进行访问。
           * 千万不要为了方便访问数据而使用static修饰成员变量，只有成员变量的数据是真正需要被共享的时候才使用static修饰。



static修饰成员变量的应用场景： 如果`一个数据需要被所有对象共享使用`的时候，这时候即可好实用static修饰。

- static修饰成员函数:

```java
class Student{

	static	String name;  //非静态成员变量
	
	static	String  country  = "中国";	  //静态的成员变量
	
	public Student(String name){
		this.name = name;
	}
}


class Demo {
	public static void main(String[] args) {
        
		Student s1 = new Student("狗娃");
		Student s2 = new Student("狗剩");
        
		System.out.println("国籍："+ Student.country);
		System.out.println("名字："+ s1.name);
		System.out.println("名字："+ s2.name);
	}
}

```

```java
/*
需求： 统计一个类被使用了多少次创建对象，该类对外显示被创建的次数。
*/
class Emp{
	
	//静态的成员变量。
	static	int count = 0;	//计数器

	String name;
	
	//构造代码块
	{
		count++;
	}

	public Emp(String name){
		this.name = name;

	}

	public Emp(){  //每创建一个对象的时候都会执行这里 的代码
		
	}
	
	public void showCount(){
		System.out.println("创建了"+ count+"个对象");
	}
}

class Demo {
	public static void main(String[] args) {
		Emp e1 = new Emp();
		Emp e2 = new Emp();
		Emp e3 =new Emp();
		e3.showCount();
	}
}
```

#### 静态的成员变量与非静态的成员变量的区别：

1. 作用上的区别：
    1. 静态的成员变量的作用共享一个 数据给所有的对象使用。
    2. 非静态的成员变量的作用是描述一类事物的公共属性。
2. 数量与存储位置上的区别：
    1. 静态成员变量是存储方法区内存中，而且只会存在一份数据。
    2. 非静态的成员变量是存储在堆内存中，有n个对象就有n份数据。
3. 生命周期的区别：
    1. 静态的成员变量数据是随着类的加载而存在，随着类文件的消失而消失。
    2. 非静态的成员数据是随着对象的创建而存在，随着 对象被垃圾回收器回收而消失。



#### 静态函数要注意的事项：

	1. 静态函数是可以调用类名或者对象进行调用的，而非静态函数只能使用对象进行调用。
 	2. 静态的函数可以直接访问静态的成员，但是不能直接访问非静态的成员。	
     		原因：静态函数是可以使用类名直接调用的，这时候可能还没有存在对象，而非静态的成员数据是随着对象的存在而存在的。
	2. 非静态的函数是可以直接访问静态与非静态的成员。
  		原因：非静态函数只能由对象调用，当对象存在的时候，静态数据老早就已经存在了，而非静态数据也随着对象的创建而存在了。
	4. `静态函数`不能出现`this`或者`super`关键字。
          原因：因为静态的函数是可以使用类名调用的，一旦使用类名调用这时候不存在对象，而this关键字是代表了一个函数 的调用者对象，这时候产生了冲突。



静态的数据的生命周期：`静态的成员变量数据是优先于对象存在的`。



- static什么时候修饰一个函数？
  		如果一个函数没有直接访问到非静态的成员时，那么就可以使用static修饰了。 一般用于`工具类型的方法`

- 静态函数不能访问非静态的成员？
  	   静态函数只要存在有对象，那么也可以访问非静态的数据。只是`不能直接访问而已`。

```java
class Student{

	String name; //名字

	static	String country = "中国"; //国籍

	//静态代码块 ：静态代码块是在Student.class文件加载到内存的时候就马上执行的。
	static{
		System.out.println("静态代码块执行了...");
	}

	//构造函数
	public Student(String name){
		this.name = name;
	}
	
	//非静态的成员函数
	public  void study(){
		System.out.println("好好学习"+this);
	}


	//静态函数
    //静态方法与非静态方法的字节码文件是同时存在内存中的。只是静态的成员变量数据是优先于对象存在而已。
	public static void sleep(){ 
		Student s= new Student("铁蛋");
		System.out.println(s.name+"呼呼大睡...");   //静态函数访问非静态的成员
	}
}

class Demo {
	public static void main(String[] args) 
	{
		Student.sleep();
	    Student s = new Student("狗娃");
	}
}
```

**构造代码块**：
	给对象进行统一的初始化工作。

应用场景： 如何创建任意对象的时候都需要调用某个方法为该对象进行初始化时，这时候就可以使用构造代码块。

**静态代码块:**
	静态代码块是静态代码块所属的类被加载到内存的时候执行的。

静态代码块的应用场景： 以后主要用于准备一个项目的初始化工作
	比如： 从配置配置文件中读取数据库用户名与密码。

```java
class Baby{

	int id;
	String name;

	//构造代码块的代码其实是在构造函数中执行的。
	{
		
		cry();
	}

	static{
		System.out.println("静态代码块执行了...");
	}

	public Baby(int id, String name) {
		this.id = id;
		this.name = name;
	}

	public Baby(){

	}

	public void cry(){
		System.out.println("哭...");
	}

	@Override
	public String toString() {
		return " 编号："+this.id+" 姓名："+ this.name;
	}
}

public class Demo {
	public static void main(String[] args) {

		Baby b1 = new Baby();
		Baby b2 = new Baby(110, "狗娃");
	}
}
```
### super关键字：

super关键字代表了`父类空间的引用`。

super关键字的 作用：
	1. 子父类存在着同名的成员时，在子类中默认是访问子类的成员，可以通过`super关键字指定访问父类的成员`。
	2. 创建子类对象时，`默认会先调用父类无参的构造方法`，可以通过super关键字指定调用父类的构造方法。

super关键字调用父类构造方法要注意的事项：
	1. 如果在子类的构造方法上没有指定调用父类的构造方法，那么`java编译器会`在`子类的构造方法上面加上super()`语句。
	2. super关键字调用父类的构造函数时，该语句必须要是`子类构造函数中的第一个语句`。
	3. super与this关键字不能同时出现在同一个构造函数中调用其他的构造函数。因为两个语句都需要第一个语句。

super关键字与this关键字的区别：
	1. 代表的事物不一致。
			1. super关键字代表的是`父类空间的引用`。	
			2. this关键字代表的是`所属函数的调用者对象`。
	2. 使用前提不一致。
			1. super关键字必须要有`继承`关系才能使用。
			2. this关键字不需要存在继承关系也可使用。
	3. 调用构造函数的区别：
			1. super关键字是调用`父类的构造函数`。
			2. this关键字是调用`本类的构造函数`。

```java
class Fu{
		
	int x = 10;

	String name;

	public Fu(){
		System.out.println("Fu类无参的构造方法..");
	}

	public Fu(String name){
		this.name = name;
		System.out.println("Fu类带参的构造方法..");
	}

	public void eat(){
		System.out.println("小头爸爸吃番薯..");
	}
}

class Zi extends Fu{
	
	int x = 20; 

	int num;
	
	public Zi(String name,int num){
		super(name); //指定调用了父类带参的 构造方法...
		this(); // 调用本类无参构造方法..
		//super(); //指定调用了父类无参构造方法。。。
		System.out.println("Zi类带参的构造方法..");
	}

	public Zi(){
		System.out.println("Zi类无参的构造方法..");
	}

	public void print(){
		System.out.println("x = " +super.x);
	}

	public void eat(){
		System.out.println("大头儿子吃龙虾..");
	}
}

class Demo {
	public static void main(String[] args) {
		Zi z = new Zi("狗娃");
	}
}
```
### 方法重写
方法重写的前提： 必须要`存在继承的关系`。

`方法的重写`: `子父类出了同名的函数`，这个我们就称作为方法的重写。

什么是时候要使用方法的重写：父类的功能无法满足子类的需求时。

方法重写要注意的事项：
	1.方法重写时， `方法名与形参列表必须一致`。
	2.方法重写时，`子类的权限修饰符`必须要`大于`或者`等于` `父类的权限修饰`符。
	3.方法重写时，`子类的返回值类型`必须要`小于`或者 `等于`父类的返回值类型。
	4.方法重写时， 子类抛出的`异常类型`要`小于`或者`等于`父类抛出的异常类型。
			Exception(最坏)
			RuntimeException(小坏)

`方法的重载`：在`一个类中 存在两个或者两个 以上的同名函数`,称作为方法重载。

方法重写的要求
	1. 函数名要一致。
	2. 形参列表不一致（形参的个数或形参 的类型不一致）
	3. 与返回值类型无关。

```java
class Animal{  //大的数据 类型 
}

class Fish extends Animal{  //Fish小 的数据类型。
}

class Fu{

	String name;

	public Fu(String name){
		this.name = name;
	}

	public Animal eat() throws RuntimeException {
		System.out.println(name+"吃番薯...");
		return new Animal();
	}
}


class Zi extends Fu{

	String num;
	
	public Zi(String name){
		super(name);//指定调用 父类带参的构造方法
	}

	//重写父类的eat方法
	public Animal eat() throws Exception{  //报错，子类抛出的异常比父类大
		System.out.println("吃点开胃菜..");
		System.out.println("喝点汤....");
		System.out.println("吃点龙虾....");
		System.out.println("吃青菜....");
		System.out.println("喝两杯....");
		System.out.println("吃点甜品....");	
		return new Animal();
	}
}

class Demo {
	public static void main(String[] args) {
		Zi z = new Zi("大头儿子");
		z.eat();
	}
}
```
```java
//普通的学生类
class Student{
	String name;
	//构造函数
	public Student(String name){
		this.name = name;
	}
	public void study(){
		System.out.println(name+"学习马克思列宁主义");
	}
}

//基础班的学生是属于学生中一种
class BaseStudent extends Student{
	public BaseStudent(String name){
		super(name);//指定调用父类构造函数
	}
	//重写
	public void study(){
		System.out.println(name+"学习javase..");
	}
}

//就业班学生 也是属于普通学生中一种
class WorkStudent extends Student{
	//构造 函数
	public WorkStudent(String name){
		super(name);
	}
	//重写
	public void study(){
		System.out.println(name+"学习javaee+android..");
	}
}

class Demo {
	public static void main(String[] args) {

		//System.out.println("Hello World!");
		
		BaseStudent s = new BaseStudent("东东");
		s.study();

		//创建一个就业班的学生
		WorkStudent w = new WorkStudent("张三");
		w.study();
	}
}
```
### instanceof关键字
instanceof关键字的作用：`判断一个对象是否属于指定的类别`。
instanceof关键字的使用前提：判断的对象与指定的类别必须要`存在继承`或者实现的关系。
instanceof关键字的使用格式：
		`对象 instanceof 类别`
```java
class Animal{
	String name;
	String color;
	public Animal(String name, String color){
		this.name = name;
		this.color = color;
	}
}

//狗是属于动物中一种
class Dog extends Animal {

	public Dog(String name,String color){
		super(name,color); //指定调用父类两个 参数的构造函数。
	}
	public void bite(){
		System.out.println(name+"咬人!!");
	}
}

//老鼠 也是属于动物中一种
class Mouse extends  Animal{
	public Mouse(String name,String color){
		super(name,color);
	}
	public void dig(){
		System.out.println(name+"打洞..");
	}
}

class Demo {
	public static void main(String[] args) {
		Dog d = new Dog("哈士奇","白色");
		System.out.println("狗是狗类吗？"+ (d instanceof Dog));
		System.out.println("狗是动物类吗？"+ (d instanceof Animal));

		//System.out.println("狗是老鼠类吗？"+ (d instanceof Mouse));		// false true
		
		Animal a = new Animal("狗娃","黄色"); //狗娃是人
		System.out.println("动物都是狗吗？"+ (a instanceof Dog));
	}
}
```
### final
final(最终、修饰符)  

final关键字的用法：
	1. final关键字修饰一个基本类型的变量时，该变量`不能重新赋值`，第一次的值为最终的。
	2. fianl关键字修饰一个引用类型变量时，该变量`不能重新指向新的对象`。
	3. final关键字修饰一个函数的时候，该`函数不能被重写`。
	4. final关键字修饰一个类的时候，该`类不能被继承`。

`public static final` 

```java
//圆形
class Circle{
	double r; //半径
	public static final double pi = 3.14; //固定不变的
	public Circle(double r){
		this.r = r;
	}
	//计算面积
	public final void getArea(){
		System.out.println("圆形的面积是："+r*r*pi);
	}
}

class Demo extends Circle{
	public Demo(double r){
		super(r);
	}
	public static void main(String[] args) 
	{
		/*
		final Circle c = new Circle(4.0);
		test(c);
		*/	
		Demo c = new Demo(4.0);
		c.getArea();
	}
	
	public static void test(Circle c){
		c = new Circle(5.0);  //c变量又重新指向了新的对象。
		c.getArea();
	}
}

```

常量的修饰符一般为： `public static final`

### 单例设计模式

- 单例设计模式：保证一个类在内存中只有一个对象。


- 模式：模式就是解决 一类 问题的固定步骤 。


- 单例设计模式的步骤：

  - 饿汉单例设计模式
    1. `私有化构造函数`。
    2. 声明`本类的引用`类型变量，并且使用`该变量指向本类对象`。
    3. 提供一个`公共静态的方法`获取本类的对象。

  - 懒汉单例设计模式：
    1. 私有化构造函数。
    2. 声明本类的引用类型变量，但是不要创建对象，
    3. 提供公共静态 的方法获取本类 的对象，获取之前先判断是否已经创建了本类 对象，如果已经创建了，那么直接返回对象即可，如果还没有创建，那么先创建本类的对象，
      然后再返回。


推荐使用： 饿汉单例设计模式。  因为懒汉单例设计模式会存在线程安全问题，目前还不能保证一类在内存中只有一个对象。

```java
//饿汉单例设计模式 ----> 保证Single在在内存中只有一个对象。
class Single{

	//声明本类的引用类型变量，并且使用该变量指向本类对象
	private static	Single s = new Single();

	//私有化构造函数
	private Single(){}

	//提供一个公共静态的方法获取本类的对象
	public	static  Single getInstance(){
		return s;
	}
}

//懒汉单例设计模式 ----> 保证Single在在内存中只有一个对象。
class Single2{
	
	//声明本类的引用类型变量，不创建本类的对象
	private static Single2 s;

	//私有化了构造函数
	private Single2(){}

	//
	public static Single2 getInstance(){
		if(s==null){
			s = new Single2();
		}
		return s;
	}
}

class Demo {
	public static void main(String[] args) {
        
		Single2 s1 = Single2.getInstance();
		Single2 s2 = Single2.getInstance();
		System.out.println("是同一个对象吗？"+ (s1==s2));
	}
}
```

### 抽象类
抽象类的应用场景
	我们在描述一类事物的时候，发现该种事物确实存在着某种行为，
但是这种行为目前是不具体的，那么我们可以抽取这种行为 的声明，但是不去实现该种行为，这时候这种行为我们称作为抽象的行为，我们就需要使用抽象类。

抽象类的好处: `强制要求子类一定要实现指定的方法`。

抽象类要注意的细节：
	1. 如果一个`函数没有方法体`，那么该函数必须要使用`abstract`修饰,把该函数修饰成抽象的函数。
	2. 如果一个类出现了`抽象的函数`，那么`该类也必须 使用abstract修饰`。
	3. `如果一个非抽象类继承了抽象类，那么必须要把抽象类的所有抽象方法全部实现`。
	4. 抽象类可以存在`非抽象方法`，也可以存在`抽象的方法`.
	5. 抽象类可以`不存在抽象方法`的。
	6. 抽象类是`不能创建对象`的。
		疑问：为什么抽象类不能创建对象呢？
		因为抽象类是存在抽象方法的，如果能让抽象类创建对象的话，那么使用抽象的对象调用抽象方法是没有任何意义的。
	7. 抽象类是`存在构造函数`的，其构造函数是提供给子类创建对象的时候初始化父类的属性的。

abstract不能与以下关键字共同修饰一个方法：
	1. abstract 不能与`private`共同修饰一个方法。
	2. abstract 不能与`static`共同修饰一个方法。
	3. abstract 不能与`final`共同修饰一个方法。

```java
abstract class Animal{
	String name;
	String  color;
	public Animal(String  name,String color){
		this.name = name;
		this.color = color;
	}
	//非抽象的方法
	public void eat(){
		System.out.println(name+"吃粮食");
	}
	//移动...
	public abstract void run();
}

//狗 是属于动物中一种 
class Dog extends Animal{
	public Dog(String name,String color){
		super(name,color);
	}
	public void run(){
		System.out.println(name+"四条腿跑得很快...");
	}
}

//鱼 是属于动物中一种
class Fish extends Animal{
	public Fish(String name,String color){
		super(name,color);
	}
	public void run(){
		System.out.println(name+"摇摇尾巴游啊游！");
	}
}

class Demo {
	public static void main(String[] args) {

		Dog d = new Dog("牧羊犬","棕色");
		d.run();

		//创建一个鱼对象
		Fish f  = new Fish("锦鲤","金黄色");
		f.run();

		//Animal a = new Animal();
	}
}
```
```java
//图形类
abstract class MyShape{
	String name;
	public MyShape(String name){
		this.name = name;
	}
	public  abstract void getArea();
	public abstract void getLength();
}

//圆形 是属于图形类的一种
class Circle extends MyShape{
	double r;
	public static final double PI = 3.14;
	public Circle(String name,double r){
		super(name);
		this.r =r;
	}
	public  void getArea(){
		System.out.println(name+"的面积是："+PI*r*r);
	}
	public void getLength(){
		System.out.println(name+"的周长是："+2*PI*r);
	}
}

//矩形 属于图形中的 一种
class Rect extends MyShape{
	int width;
	int height;
	public Rect(String name,int width, int height){
		super(name);
		this.width = width;
		this.height = height;
	}
	public  void getArea(){
		System.out.println(name+"的面积是："+width*height);
	}
	public  void getLength(){
		System.out.println(name+"的周长是："+2*(width+height));
	}
}

class Demo {
	public static void main(String[] args) {
	
		Circle c = new Circle("圆形",4.0);
		c.getArea();
		c.getLength();

		//矩形
		Rect r = new Rect("矩形",3,4);
		r.getArea();
		r.getLength();
	}
}
```

### 接口-interface

接口：拓展功能的。  usb接口.。。

接口的定义格式：
```java
	interface 接口名{
	
	}
```

接口要注意的事项 ：。
	1. 接口的`成员变量`默认的修饰符为： `public static final` 。那么也就是说接口中的成员变量都是`常量`。
	2. 接口中 的方法都是`抽象`的方法，默认的修饰符为： `public abstract`。
	3. 接口`不能创建对象`。
	4. 接口是`没有构造方法`的。
	5. 接口是给类去实现使用的，`非抽象类实现一个接口的时候，必须要把接口中所有方法全部实现`。

实现接口的格式：
```java
	class  类名 implements 接口名{
	
	}
```

接口的作用：
	1. 程序的解耦。  （低耦合）
	2. 定义约束规范。
	3. 拓展功能。

```java
interface A{
	//成员变量
	public static final int i=10;
	//成员函数
	public void print();
}

class Demo implements A{ // Demo就实现了A接口

	public static void main(String[] args) {
		Demo d = new Demo();
		d.print();
	}
	//实现接口中的方法
	public void print(){
		System.out.println("这个是接口中的print方法...");
	}
}
```
```java
//普通的铅笔类
class Pencil{
	String name;
	public Pencil(String name){
		this.name = name;
	}
	public void write(){
		System.out.println(name+"沙沙的写...");
	}
}

//橡皮接口
interface Eraser{
	public void remove();
}

//带橡皮的铅笔
class PencilWithEraser extends Pencil implements Eraser {

	public PencilWithEraser(String name){
		super(name);
	}

	public void remove(){
		System.out.println("涂改,涂改....");
	}
}

class Demo {
	public static void main(String[] args) {	
		PencilWithEraser p = new PencilWithEraser("2B铅笔");
		p.write();
		p.remove();
	}
}
```

类与接口之间关系： `实现关系`。

类与接口要注意的事项：
	1. 非抽象类实现一个接口时，必须要把接口中所有方法全部实现。
	2. 抽象类实现一个接口时，可以实现也可以不实现接口中的方法。
	3. `一个类可以实现多个接口`。
		
疑问： java为什么不支持多继承，而支持了多实现呢？
	类多继承情况下： 两个父类存在同名方法或变量，那么子类该调用谁就变得模糊不清了。
	接口多实现情况下： 两个接口存在同名方法，因为都是需要实现类去实现的，所以无关大雅；而接口的变量默认是static final类型的，也就是常量，根据JVM的底层机制，常量在编译期就确定了值，倘若两个同名常量不同值，编译期就直接编译不通过了。
```java
class A{
	public void print(){
		System.out.println("AAAAAA");
	}
}
class B{
	public void print(){
	    System.out.println("BBBBBB");
	}
}

class C extends A ,B{
	new C().print();
}
```

接口与接口之间关系： `继承关系`。

接口与接口之间要注意事项：
	1. 一个接口是可以继承多个接口的。

```java
interface A{
	public void print();
}

interface C{
	public void getArea();
}

interface B extends A,C{ // B接口继承A接口
	public void test();
}

class Demo implements B{
	public static void main(String[] args) {
		Demo d = new Demo10();
		d.print();
	}

	public void getArea(){}

	public void test(){}

	public void print(){
		System.out.println("这个是A接口的print方法...");
	}
}

```
### 内部类
内部类：`一个类定义在另外一个类的内部`，那么该类就称作为内部类。

内部类的class文件名： `外部类$内部类`.  好处：便于区分该class文件是属于哪个外部类的。

#### 成员内部类:
成员内部类的访问方式：
	方式一：在外部类提供一个方法创建内部类的对象进行访问。
	方式二：在其他类直接创建内部类的对象。 
		格式：`外部类.内部类  变量名 = new 外部类().new 内部类()`;
	
	注�����： 如果是一个`静态内部类`，那么在其他类创建 的格式：
			`外部类.内部类  变量名 = new 外部类.内部类()`;

内部类的应用场景： 我们在描述A事物的时候，发现描述的A事物内部还存在另外一个较复杂的事物B时候，而且这个比较复杂事物B还需要访问A事物的属性等数据，那么这时候我们就可以使用内部类描述B事物。
```java

比如：
人--->心脏

class 人{
	血

	等....
	class 心脏{

	}		
}
```

内部类的好处：`内部类可以直接访问外部类的所有成员`。

内部类要注意的细节：
	1. 如果外部类与内部类存在同名的成员变量时，在内部类中默认情况下是访问内部类的成员变量。
		可以通过"`外部类.this.成员变量名`" 指定访问外部类的成员。
	2. `私有的成员内部类`只能在外部类提供一个方法创建内部类的对象进行访问，不能在其他类创建对象了。
	3. 成员内部类一旦出现了`静态的成员`，那么该类也必须使用`static修饰`。

```java
//成员内部类
//外部类
class Outer{
	//成员变量
	int x = 100; // Outer.class文件被加载到内存的时候存在内存中。  静态的成员数据是不需要对象存在才能访问。

	//成员内部类
	static	class Inner{
		static	int i = 10;
		public void print(){
			System.out.println("这个是成员内部类的print方法！"+i);
		}
	}

	//在外部的方法中创建了内部类的对象，然后调用内部 方法。
	public void instance(){
		Inner inner = new Inner();
		inner.print();
	}
}

//其他类
class Demo {
	public static void main(String[] args) {
		/*
		System.out.println(Outer.Inner.i);
		Outer outer = new Outer();
		outer.instance();
		
		Outer.Inner inner = new Outer().new Inner();
		inner.print();
		*/

		Outer.Inner inner = new Outer.Inner();
		inner.print();
	}
}
```
#### 局部内部类：
在一个`类的方法内部定义另外一个类`，那么另外一个类就称作为局部内部类。

局部内部类要注意的细节：
	1. 如果局部 内部类访问了一个局部变量，那么该局部变量必须使用`final`修饰、
   
```java
//局部内部类
class  Outer{

	String name= "外部类的name";

	public void test(){

		//局部变量
		final	int y =100;  // y 什么时候从内存中消失？ 方法执行完毕之后y消失。

		//局部内部类
		class Inner{     /*
							当test方法执行完毕之后，那么y马上从内存中消失，而Inner对象在方法
							执行完毕的时候还没有从内存中消失，而inner对象的print方法还在访问着
							y变量，这时候的y变量已经消失了，那么就给人感觉y的生命变量已经被延长了
							.
							解决方案： 如果一个局部内部类访问一个局部变量的时候，那么就让该局部内部类
							访问这个局部 变量 的复制品。
						*/
			int x = 10;

			public void print(){
				System.out.println("这个是局部内部类的print方法.."+y);
			}
		}
		
		Inner inner = new Inner();  //这个inner对象什么时候消失？  Inner对象的生命周期比局部变量y的生命周期要长。
		inner.print();
	}
}

class Demo {
	public static void main(String[] args) {
		Outer outer = new Outer();
		outer.test();
	}
}
```
#### 匿名内部类
匿名内部类：`没有类名`的类就称作为匿名内部类。

匿名内部类的好处：简化书写。

匿名内部类的使用前提：必须存在`继承`或者`实现`关系才能使用。

匿名内部类一般是`用于实参`。

```java
//继承关系下匿名内部类
abstract class Animal{
	public abstract Animal run();
	public abstract void sleep();
}

class Outer{
	public void print(){
		//需求： 在方法内部定义一个类继承Animal类，然后调用run方法与sleep()。
		
		/*
		//局部内部类
		class Dog extends Animal{
			
			public void run(){
				System.out.println("狗在跑..");
			}

			public void sleep(){
				System.out.println("狗趴在睁开眼睛睡..");
			}
		}
		
		//创建对象
		Dog d = new Dog();
		d.run();	
		d.sleep();
		*/	
		//匿名内部类 ：匿名内部类只是没有类名，其他的一概成员都是具备的。
		//匿名内部类与Animal是继承的关系。目前是创建Animal子类的对象. 
		Animal	a = new Animal(){  //多态
			
				//匿名内部的成员 
				public Animal run(){
					System.out.println("狗在跑..");
					return this;
				}
				
				public void sleep(){
					System.out.println("狗趴在睁开眼睛睡..");
				}

				//特有的方法
				public void bite(){
					System.out.println("狗在咬人..");
				}
		
			};
	
		a.bite();
		a.run();
		a.sleep();
	}
}

class Demo {
	public static void main(String[] args) {
		Outer outer = new Outer();
		outer.print();
	}
}
```
```java
//实现关系下匿名内部类
interface Dao{
	public void add();
}

class Outer{

	public void print(){
		//创建一个匿名内部类的对象
		new Dao(){   //不是接口不能创建对象吗？怎么现在又可以了？
			 
			 public void add(){
				System.out.println("添加成功");

			 }
		}.add();
	}
}

class Demo {
	public static void main(String[] args) {
		test(new Dao(){
			
			public void add(){
				System.out.println("添加员工成功");
			}
		});
	}

	//调用这个方法...
	public static void  test(Dao d){  //形参类型是一个接口引用..
		d.add();
	}
}
```

### Java包(package)
包的作用：
	1. 解决类名重复产生冲突的问题。
	2. 便于软件版本的发布。

定义包的格式：
	`package 包名`;

包名命名规范：`包名全部小写`。

包语句要注意的事项：
	1. package语句必须位于java文件中中`第一个语句`。
	2. 如果一个类加上了包语句，那么该类的完整类名就是: `包名.类名`
	3. `一个java文件只能有一个包语句`。

问题：  每次编译的时候都需要自己创建一个文件夹，把对应的class文件存储到文件夹中。
		javac -d 指定类文件的存放路径   java源文件

```java
package aa;

class Demo {
	public static void main(String[] args) {
		System.out.println("这个是Demo的main方法...");
	}
}
```
有了包之后类与类之间的访问：

问题： 有了包之后类与类之间的访问每次都必须 要写上包名！

解决方案： sum提供导包语句让我们解决该问题。

导包语句作用：简化书写。 (误区： 把一个类导入到内存中)

导包语句的格式：
	`import 包名.类名;`   （导入xxx包中某个类）

导包语句要注意的细节：
	 1. 一个java文件中可以出现多句导包语句。
	 2. `"*"`是 导包语句的通配符。可以匹配任何的类名。
	 3. import aa.*; 是不会作用于aa包下面的子包的。
	
推荐使用：`import 包名.类名`;   因为使用*通配符会导致结构不清晰。

什么时候使用import语句: 
	1. 相互访问的两个类不是在同一个包下面，这时候就需要使用到导包语句。
	2. `java.lang` 是默认导入的，不需要我们自己导入。

```java
package hello;
public class Hello 
{
	/*
	static{
		System.out.println("这个是Dmeo3的静态代码块...");
	}
	*/
	public void print(){
		System.out.println("这个是Dmeo3的print方法...");
	}

	public void test(){
		System.out.println("这个是Demo5的test方法...");
	}
}

//-------------------------------------------------------

package world;
import hello.Hello;
import java.util.Arrays; 
class Demo {
	public static void main(String[] args) {

		int[] arr = {1,234,2};
		Arrays.sort(arr);
		Object o = new Object();

		Hello h = new Hello();
		h.print();

		Hello h2 = new Hello();
		h2.test();
	}
}
```

### 权限修饰符

权限修饰符： 权限修饰符就是控制被修饰的成员的范围可见性。

			public(公共)	protected(受保护)	default(缺省)	private (大到小)

同一个类      true            true                 true           true

同一个包      true            true                 true           false

子父类        true            true                 false          false

不同包        true            false                false          false

注意： 在`不同包`下面只有public 与 protected 可以访问，而且`protected必须是在继承关系下才能够访问`。

### 模板模式

模板模式 ：解决某类事情的步骤有些是固定的，有些是会发生变化的，那么这时候我们可以
为这类事情提供一个模板代码，从而提高效率 。

需求；编写一个计算程序运行时间 的模板。

模板模式的步骤：
	1. 先写出解决该类事情其中的一件的解决方案。
	2. 分析代码，把会发生变化的代码抽取出来独立成一个方法。把该方法描述成一个抽象的方法。
	3. 使用final修饰模板方法，防止别重写你的模板方法。

```java
abstract class MyRuntime{
	public final void getTime(){
		long startTime = System.currentTimeMillis();	//记录开始的时间
		code();
		long endTime = System.currentTimeMillis();  //记录结束的时间.
		System.out.println("运行时间 ："+ (endTime-startTime));
	}
	public abstract void code();
}

class Demo extends MyRuntime{
	public static void main(String[] args) {
		Demo d = new Demo();
		d.getTime();
	}
	//code方法内部就写要计算运行时间 的代码；
	public  void code(){
		int i = 0;
		while(i<100){
			System.out.println("i="+i);
			i++;
		}
	}
}
```

## 异常
### 异常的体系：
----------| `Throwable`    所有异常或者错误类的`超类`
--------------|`Error`     错误   错误一般是用于jvm或者是硬件引发的问题，所以我们一般不会通过代码去处理错误的。
--------------|`Exception` 异常  是需要通过代码去处理的。

如何区分错误与异常呢：
	如果程序出现了不正常的信息，如果不正常的信息的类名是以Error结尾的，那么肯定是一个错误。
	如果是以Exception结尾的，那么肯定就是一个异常。

Throwable常用的方法：
	`toString()`  返回当前异常对象的`完整类名+病态信息`。
	`getMessage()` 返回的是创建Throwable传入的字符串信息。
	`printStackTrace()` 打印异常的栈信息。


```java
class Demo {
	public static void main(String[] args) {
		
		//创建了一个Throwable对象。
		Throwable t = new Throwable("头晕，感冒..");
		String info = t.toString();
		String message = t.getMessage();
		System.out.println("toString: "+ info);  // java.lang.Throwable  包名+类名 = 完整类名
		System.out.println("message: "+ message);
		
		test();
	}

	public static void test(){
		Throwable t = new Throwable();
		t.printStackTrace();
	}
}
```

疑问：下面的信息是通过printStackTrace方法打印出来，那么异常对象从何而来呢？ 
	Exception in thread "main" java.lang.ArithmeticException: / by zero
        at Demo.div(Demo.java:10)
        at Demo.main(Demo.java:5)
	
jvm运行到a/b这个语句的时候，发现b为0，除数为0在我们现实生活中是属于不正常的情况，jvm一旦发现了这种不正常的情况时候，那么jvm就会马上创建一个对应的异常对象，并且会调用这个异常对象的`printStackTrace`的方法来处理。

### 异常的处理：

#### 捕获处理

捕获处理的格式：
```java
try{
	可能发生异常的代码；

}catch(捕获的异常类型 变量名){
	处理异常的代码....
}
```

捕获处理要注意的细节：
	1. 如果try块中代码出了异常经过了处理之后，那么`try-catch`块外面的代码可以`正常执行`。
	2. 如果try块中出了异常的代码，那么在`try块`中出现异常代码后面的代码是`不会执行了`。
	3. 一个try块后面是可以跟有多个catch块的，也就是一个try块可以捕获多种异常的类型。
	4. 一个try块`可以捕获多种异常的类型`,但是捕获的异常类型必须从`小到大进行捕获`，否则编译报错。

疑问： 以后捕获处理的时候是否就是捕获Exception即可？
	错的，因为我们在现实开发中遇到不同的异常类型的时候，我往往会有不同 的处理方式。
	所以要分开不同的异常类型处理。

```java
class Demo {
	public static void main(String[] args) {
		int[] arr = null;
		div(4,0,arr);
	}

	public static void div(int a , int b,int[] arr){
		int c = 0;
		try{
			c = a/b;      //jvm在这句话的时候发现了不正常的情况，那么就会创建一个对应的异常对象。
			System.out.println("数组的长度："+ arr.length);
		}catch(ArithmeticException e){
			//处理异常的代码....
			System.out.println("异常处理了....");
			System.out.println("toString:"+ e.toString());
		}catch(NullPointerException e){
			System.out.println("出现了空指针异常....");
		}catch(Exception e){  
			System.out.println("我是急诊室，包治百病！");
		}
		System.out.println("c="+c);
	}
}
```

#### 抛出处理
抛出处理---`throw， throws`

抛出处理要注意的细节：
	1. 如果一个`方法的内部抛出`了一个异常对象，那么必须要在`方法上声明抛出`。
	2. 如果调用了一个声明抛出异常的方法，那么`调用者必须要处理`异常。
	3. 如果一个方法内部抛出了一个异常对象，那么`throw语句`后面的代码都`不会再执行`了（一个方法遇到了throw关键字，该方法也会马上停止执行的）。


throw与throws两个关键字：
	1. `throw`关键字是用于`方法内部`的，`throws`是用于`方法声声明`上的。
	2. throw关键字是用于方法内部抛出一个异常对象的，throws关键字是用于在方法声明上声明抛出异常类型的。
	3. throw关键字后面`只能有一个异常对象`，throws后面一次`可以声明抛出多种类型的异常`。

疑问：何时使用抛出处理？何时捕获处理？原则是如何？
	如果你需要通知到调用者，你代码出了问题，那么这时候就使用抛出处理.
	如果代码是直接与用户打交道遇到了异常千万不要再抛，再抛的话，就给了用户了。
	这时候就应该使用捕获处理。
		
```java
class Demo {
	public static void main(String[] args) {
		try{
			int[] arr = null;
			div(4,0,arr);      //调用了一个 声明抛出异常类型 的方法
		}catch(Exception e){
			System.out.println("出现异常了...");
			e.printStackTrace();
		}
	}

	public static void div(int a, int b,int[] arr) throws Exception,NullPointerException {
		if(b==0){
			throw new Exception(); //抛出一个异常对象...
		}else if(arr==null){
			throw new  NullPointerException();
		}
		int c = a/b;
		System.out.println("c="+c);
	}
}
```

### 自定义异常类
自定义异常类的步骤：  自定义一个类继承`Exception`即可。
```java
class NoIpException extends Exception{
	public NoIpException(String message){
		super(message);  //调用了Exception一个参数的构造函数。
	}
}

class Demo {
	public static void main(String[] args) {
		String ip = "192.168.10.100";
		ip = null;
		try{
			feiQ(ip);  // 如果调用了一个声明抛出异常类型的方法，那么调用者必须要处理。
		}catch(NoIpException e){
			e.printStackTrace();
			System.out.println("马上插上网线！");
		}
	}

	public static void feiQ(String ip) throws NoIpException{
		if(ip==null){
			throw new  NoIpException("没有插网线啊，小白！");
		}
		System.out.println("正常显示好友列表..");
	}
}
```
```java
class NoMoneyException extends Exception {
	public NoMoneyException(String message){
		super(message);
	}
}
class Demo {
	public static void main(String[] args) {
		try{
			eat(9);

		}catch(NoMoneyException e){
			e.printStackTrace();
			System.out.println("跟我洗碗一个月！！");
		}
	}

	public static void eat(int money) throws NoMoneyException{
		if(money<10){
			throw new NoMoneyException("吃霸王餐");
		}
		System.out.println("吃上了香喷喷的地沟油木桶饭!!");
	}
}

```
### 运行时,编译时异常
`运行时异常`: 如果一个方法内部抛出了一个运行时异常，那么方法上可以声明也可以不 声明，调用者可以以处理也可以不处理。
	`RuntimeException`以及`RuntimeException`子类 都是属于运行时异常。
`编译时异常`(非运行时异常、受检异常):  如果一个方法内部抛出了一个编译时异常对象，那么方法上就必须要声明，而且调用者也必须要处理。
	除了运行时异常就是编译异常。

疑问： 为什么java编译器会如此严格要求编译时异常，对运行时异常如此宽松？
	运行时异常都是可以通过程序员良好的编程习惯去避免，所以java编译器就没有严格要求处理运行时异常。
```java

import java.security.acl.*;
class Demo {
	public static void main(String[] args) throws InterruptedException{

		int[] arr = null;
		div(4,0,arr);

		 Object o = new Object();
		 o.wait();
	}
	public static void div(int a , int b ,int[] arr) {
		if(b==0){
			return;
		}
		int c = a/b;
		System.out.println("c = "+c);

		if(arr!=null){
			System.out.println("数组的长度： "+arr.length);
		}
	}
}
```

### finally 块
finally块的 使用前提是必须要存在try块才能使用。

`finally块的代码在任何情况下都会执行的`，除了jvm退出的情况。

finally非常适合做资源释放的工作，这样子可以保证资源文件在任何情况下都 会被释放。

try块的三种组合方式：
第一种： 比较适用于有异常要处理，但是没有资源要释放的。
```java
try{

	可能发生异常的代码

}catch(捕获的异常类型 变量名){
	处理异常的代码
}
```

第二种：比较适用于既有异常要处理又要释放资源的代码。
```java
try{

	可能发生异常的代码

	}catch(捕获的异常类型 变量名){
		处理异常的代码
	}finally{ 
		释放资源的代码;
}
```

第三种： 比较适用于内部抛出的是运行时异常，并且有资源要被释放。
```java
try{

	可能发生异常的代码

	}finally{ 
		释放资源的代码;
}
```

```java
class Demo {
	public static void main(String[] args) {
		div(4,0);
	}
	public static void div(int a, int b){
		try{
			if(b==0){
				System.exit(0);//退出jvm
			}
			int c = a/b;
			System.out.println("c="+ c);
		}catch(Exception e){
			System.out.println("出了除数为0的异常...");
			throw e;
		}finally{
			System.out.println("finall块的代码执行了..");
		}
	}
}
```

```java
//fianlly释放资源的代码
import java.io.*;
class Demo {
	public static void main(String[] args) {
		FileReader fileReader = null;
		try{
			//找到目标文件
			File file = new File("f:\\a.txt");
			//建立程序与文件的数据通道
			fileReader = new FileReader(file);
			//读取文件
			char[] buf = new char[1024];
			int length = 0; 
			length = fileReader.read(buf);
			System.out.println("读取到的内容："+ new String(buf,0,length));
		}catch(IOException e){
			System.out.println("读取资源文件失败....");
		}finally{
			try{
				//关闭资源
				fileReader.close();
				System.out.println("释放资源文件成功....");
			}catch(IOException e){
				System.out.println("释放资源文件失败....");
			}
		}
	}
}
```

## 常用类
### Object
Object类是所有类的终极父类。 任何一个类都继承了`Object`类。

Object类常用的方法：
	`toString()`     返回该对象的字符串表示。 返回一个字符串用于描述该对象的。

	`equals(Object obj)`   用于比较两个对象的内存地址，判断两个对象是否为同一个对象。

	`hashCode()`   返回该对象的哈希码值(大家可以把哈希码就 理解成是对象的内存地址)

疑问：toString() 有何作用？	
	重写toString之后，我们直接输出一个对象的时候，就会输出符合我们所需求的格式数据。

java中的规范：一般我们重写了一个类的equals方法，我们都会重写它的hashCode方法。

```java
class Person{
	int id;
	String name;
	public Person(int id, String name) {
		this.id = id;
		this.name = name;
	}
	public Person() {
	}
	//目前我需要直接输出一个对象的时候，输出 的格式是： 编号：110 姓名： 狗娃  这种格式。 目前Object的
    //toString方法无法满足子类的需求，那么这时候我们就应该对Object类的toString进行重写。
	@Override
	public String toString() {
		return "编号："+ this.id + " 姓名："+this.name;
	}
	
	//为什么要重写：Object的equals方法默认比较的是两个对象的内存地址，我目前需要比较的是两个对象的身份证，所以Object类的equals方法不符合我现在的需求。
	@Override
	public boolean equals(Object obj) { 
		Person p  = (Person)obj;
		return this.id== p.id;
	}
	
	@Override
	public int hashCode() {
		return  this.id;
	}
		
}

public class Demo {

	public static void main(String[] args) {
		
		/*
		Object o = new Object();
		System.out.println(o.toString());  // java.lang.Object@18b3364  返回的字符串表示：完整类名+@+ 对象的哈希码

		System.out.println(o);  // 通过查看源代码得知，直接输出一个对象 的时候，实际上在println方法内部会调用这个 调用的toString方法，把toString方法返回的内容输出。

		//疑问： 为什么直接输出一个对象的时候和输出对象的toString方法返回的字符串结果是一样的呢？
		
		Person  p1 = new Person(110,"狗娃");
		System.out.println("p1:"+p1);  
		//如果我们能够输出一个p对象的时候，输出的格式： 编号：110 姓名： 狗娃..
		Person  p2 = new Person(112,"狗剩");
		System.out.println("p2:"+p2);  
		*/
		
		Person p1 = new Person(110,"狗娃");
		Person p2 = new Person(110,"陈富贵");
		//需求：在现实生活中只要两个人的身份证一致，那么就是同一个人。
		System.out.println("p1与p2是同一个对象吗？"+ p1.equals(p2));
		
		System.out.println("p1哈希码:"+ p1.hashCode());
		System.out.println("p2哈希码:"+ p2.hashCode());
		
	}
}
```

Object 类是一个特殊的类，是所有类的父类。它主要提供了以下 11 个方法：
`native`是一个计算机函数，一个Native Method就是一个Java调用非Java代码的接口。方法的实现由非Java语言实现，比如C或C++。

```java
//native方法，用于返回当前运行时对象的Class对象，使用了final关键字修饰，故不允许子类重写。
public final native Class<?> getClass()

//native方法，用于返回对象的哈希码，主要使用在哈希表中，比如JDK中的HashMap。
public native int hashCode() 

//用于比较2个对象的内存地址是否相等，String类对该方法进行了重写用户比较字符串的值是否相等。
public boolean equals(Object obj)

//naitive方法，用于创建并返回当前对象的一份拷贝。一般情况下，对于任何对象 x，表达式 x.clone() != x 为true，x.clone().getClass() == x.getClass() 为true。Object本身没有实现Cloneable接口，所以不重写clone方法并且进行调用的话会发生CloneNotSupportedException异常。
protected native Object clone() throws CloneNotSupportedException

//返回类的名字@实例的哈希码的16进制的字符串。建议Object所有的子类都重写这个方法。
public String toString()

//native方法，并且不能重写。唤醒一个在此对象监视器上等待的线程(监视器相当于就是锁的概念)。如果有多个线程在等待只会任意唤醒一个。
public final native void notify()

//native方法，并且不能重写。跟notify一样，唯一的区别就是会唤醒在此对象监视器上等待的所有线程，而不是一个线程。
public final native void notifyAll()

//native方法，并且不能重写。暂停线程的执行。注意：sleep方法没有释放锁，而wait方法释放了锁 。timeout是等待时间。
public final native void wait(long timeout) throws InterruptedException

//多了nanos参数，这个参数表示额外时间（以毫微秒为单位，范围是 0-999999）。 所以超时的时间还需要加上nanos毫秒。
public final void wait(long timeout, int nanos) throws InterruptedException

//跟之前的2个wait方法一样，只不过该方法一直等待，没有超时时间这个概念
public final void wait() throws InterruptedException

//实例被垃圾回收器回收的时候触发的操作
protected void finalize() throws Throwable { }

```

#### == 与 equals(重要)

**==** : 它的作用是判断两个对象的`地址`是不是相等。即，判断两个对象是不是同一个对象(`基本数据类型==比较的是值`，`引用数据类型==比较的是内存地址`)。

**equals()** : 它的作用也是判断两个`对象是否相等`。但它一般有两种使用情况：

- 情况 1：类没有覆盖 equals() 方法。则通过 equals() 比较该类的两个对象时，等价于通过“==”比较这两个对象。
- 情况 2：类覆盖了 equals() 方法。一般，我们都覆盖 equals() 方法来比较两个对象的内容是否相等；若它们的内容相等，则返回 `true` (即，认为这两个对象相等)。

**举个例子：**
```java
public class test1 {
    public static void main(String[] args) {
        String a = new String("ab"); // a 为一个引用
        String b = new String("ab"); // b为另一个引用,对象的内容一样
        String aa = "ab"; // 放在常量池中
        String bb = "ab"; // 从常量池中查找
        if (aa == bb) // true
            System.out.println("aa==bb");
        if (a == b) // false，非同一对象
            System.out.println("a==b");
        if (a.equals(b)) // true
            System.out.println("aEQb");
        if (42 == 42.0) { // true
            System.out.println("true");
        }
    }
}
```

**说明：**
- `String 中的 equals 方法是被重写过的`，因为 object 的 equals 方法是比较的对象的内存地址，而 String 的 equals 方法比较的是`对象的值`。
- 当创建 String 类型的对象时，虚拟机会在常量池中查找有没有已经存在的值和要创建的值相同的对象，如果有就把它赋给当前引用。如果没有就在常量池中重新创建一个 String 对象。


#### hashCode 与 equals (重要)

面试官可能会问你：“你重写过 hashcode 和 equals 么，为什么重写 equals 时必须重写 hashCode 方法？”

#### hashCode（）介绍

hashCode() 的作用是获取`哈希码`，也称为`散列码`；它实际上是返回一个 `int 整数`。这个哈希码的作用是确定该对象在`哈希表`中的`索引位置`。hashCode() 定义在 JDK 的 `Object.java` 中，这就意味着 Java 中的任何类都包含有 hashCode() 函数。

散列表存储的是键值对(key-value)，它的特点是：能根据“键”快速的检索出对应的“值”。这其中就利用到了散列码！（可以快速找到所需要的对象）

#### 为什么要有 hashCode

**我们先以“HashSet 如何检查重复”为例子来说明为什么要有 hashCode：** 当你把对象加入 HashSet 时，HashSet 会先计算对象的 hashcode 值来判断对象加入的位置，同时也会与该位置其他已经加入的对象的 hashcode 值作比较，如果没有相符的 hashcode，HashSet 会假设对象没有重复出现。但是如果发现有相同 hashcode 值的对象，这时会调用 `equals()`方法来检查 hashcode 相等的对象是否真的相同。如果两者相同，HashSet 就不会让其加入操作成功。如果不同的话，就会重新散列到其他位置。（摘自我的 Java 启蒙书《Head first java》第二版）。这样我们就大大减少了 equals 的次数，相应就大大提高了执行速度。

通过我们可以看出：`hashCode()` 的作用就是**获取哈希码**，也称为散列码；它实际上是返回一个 int 整数。这个**哈希码的作用**是确定该对象在哈希表中的索引位置。**`hashCode()`在散列表中才有用，在其它情况下没用**。在散列表中 hashCode() 的作用是获取对象的散列码，进而确定该对象在散列表中的位置。

#### hashCode（）与 equals（）的相关规定

1. 如果两个对象相等，则 hashcode 一定也是相同的
2. 两个对象相等,对两个对象分别调用 equals 方法都返回 true
3. 两个对象有相同的 hashcode 值，它们也不一定是相等的
4. **因此，equals 方法被覆盖过，则 hashCode 方法也必须被覆盖**
5. hashCode() 的默认行为是对堆上的对象产生独特值。如果没有重写 hashCode()，则该 class 的两个对象无论如何都不会相等（即使这两个对象指向相同的数据）


#### Java 序列化中如果有些字段不想进行序列化

对于不想进行序列化的变量，使用 transient 关键字修饰。

transient 关键字的作用是：阻止实例中那些用此关键字修饰的的变量序列化；当对象被反序列化时，被 transient 修饰的变量值不会被持久化和恢复。transient 只能修饰变量，不能修饰类和方法。

###  String
#### String类

笔试题目：new String("abc")创建了几个对象？
	两个对象， 一个对象是 位于`字符串常量池`中，一个对象是位于`堆内存`中。

`String类重写了Object的equals方法`，比较的是两个字符串对象 的内容 是否一致。

```java
public class Demo {
	public static void main(String[] args) {
		
		String str1 = "hello";
		String str2 = "hello";
		String str3 = new String("hello");
		String str4 = new String("hello");
		System.out.println("str1==str2?"+(str1==str2));  // true  
		System.out.println("str2==str3?"+(str2==str3));  //false
		System.out.println("str3==str4?"+(str3==str4));  // false
		System.out.println("str3.equals(str2)?"+(str3.equals(str4))); //true
		//是String类重写了Object的equals方法，比较的是两个字符串对象 的内容 是否一致。
		// "=="用于比较 引用数据类型数据的时候比较的是两个对象 的内存地址，equals方法默认情况下比较也是两个对象 的内存地址。
		
		test(null);
	}
	
	public static void test(String str){
		if("中国".equals(str))
		{
			System.out.println("回答正确");
		}else{
			System.out.println("回答错误");
		}
	}
}
```
#### String的构造方法：
 
 	String()  创建一个空内容的字符串对象。
 	String(byte[] bytes)  使用一个`字节数组`构建一个字符串对象
 	String(byte[] bytes, int offset, int length) 
 		bytes :  要解码的数组
 		offset： 指定从数组中那个索引值开始解码。
 		length： 要解码多个元素。
 	
 	String(char[] value)  使用一个`字符数组`构建一个字符串。	
 	String(char[] value, int offset, int count)  使用一个字符数组构建一个字符串，指定开始的索引值，与使用字符个数。
	String(int[] codePoints,int offset,int count)
	String(String original) 

记住： 使用字节数组或者字符数组都可以构建一个字符串对象。

`字节数组`与`字符数组`、`字符串`他们三者之间是可以互相转换。

```java
public class Demo {
	public static void main(String[] args) {

		String str = new String();
		byte[] buf = {97,98,99};
		
		str = new String(buf); //使用一个字节数组构建一个字符串对象
		str = new String(buf,1,2);   //使用一个字节数组构建一个字符串对象,指定开始解码 的索引值和解码的个数。
		
		char[] arr = {'明','天','是','圣','诞'};
		str = new String(arr); //使用字符数组构建一个字符串
		str = new String(arr,3,2);
		
		int[] 	buf2 = {65,66,67};
		str = new String(buf2,0,3);
		
		str = new String("abc");
	
		System.out.println("字符串的内容："+str);
		
	}
}
```

#### String常用方法
获取方法
```java
	int length()  获取字符串的长度
	char charAt(int index) 获取特定位置的字符 (角标越界)
	int indexOf(String str) 查找子串第一次出现的索引值,如果子串没有出现 在字符串中，那么则返回-1表示。
	int lastIndexOf(String str) 查找子串最后一次出现的索引值 , 如果子串没有出现 在字符串中，那么则返回-1表
```

```java
public class Demo {
	public static void main(String[] args) {
		String str = "abc中国ab中国";
		System.out.println("字符串的字符 个数: " + str.length() );
		System.out.println("根据索引值获取对应 的字符:"+ str.charAt(3));
		System.out.println("查找子串第一次出现的索引值：" + str.indexOf("中国"));
		System.out.println("查找子串最后一次出现的索引值：" + str.lastIndexOf("中国"));
		
	}
}
```
判断方法
```java
	boolean endsWith(String str) 是否以指定字符结束
	boolean isEmpty()是否长度为0 如：“” null V1.6
	boolean contains(CharSequences) 是否包含指定序列 应用：搜索
	boolean equals(Object anObject) 是否相等
	boolean equalsIgnoreCase(String anotherString) 忽略大小写是否相等

```

转换方法 
```java   
	char[] toCharArray()  将字符串转换为字符数组
	byte[]	getBytes();
```
```java
public class Demo4 {
	
	public static void main(String[] args) {
		String str = "Demo.java";
		System.out.println("是否以指定 的字符串结束:"+ str.endsWith("ja"));
		//str = "";
		System.out.println("判断字符串是否为空内容："+str.isEmpty());
		System.out.println("判断字符串是否包含指定的内容："+ str.contains("Demo"));
		System.out.println("判断两个 字符串的内容是否一致："+ "DEMO.JAVA".equals(str));
		System.out.println("判断两个字符串的内容是否一致(忽略大小写比较):"+ "DEMO.JAVA".equalsIgnoreCase(str));
		
		//转换的方法
		char[] buf = str.toCharArray();  //把字符串转换字符数组
		System.out.println("字符数组："+ Arrays.toString(buf));
		byte[] buf2 = str.getBytes();	//把字符串转字节数组
		System.out.println("字节数组："+ Arrays.toString(buf2));
	}
}
```
其他方法
```java
	String replace(String oldChar, String newChar) 替换
	String[] split(String regex) 切割
	
	String substring(int beginIndex)   指定开始 的索引值截取子串
	String substring(int beginIndex, int endIndex)指定开始 与结束的索引值截取子串 `包头不包尾`截取的位置是结束的索引值`-1`.
	
	String toUpperCase() 转大写
	String toLowerCase() 转小写
	String trim() 去除字符串首尾的空格
```

```java

import java.util.Arrays;

public class Demo {
	public static void main(String[] args) {
		String str = "今天晚上不考试";
		System.out.println("指定新内容替换旧 的内容:"+ str.replace("不", "要好好"));
		str = "大家-下-午-好";
		String[] arr = str.split("-"); //根据指定的字符进行切割 。
		System.out.println("字符串数组的内容："+ Arrays.toString(arr));
		str = "广州传智播客";
		System.out.println("指定开始的索引值截取子串:"+ str.substring(2));
		System.out.println("指定开始的索引值截取子串:"+ str.substring(2,6)); //包头不包尾  注意：截取的内容是包括开始的索引值，不包括结束的索引值， 截取的位置是结束的索引值-1.
		
		str = "abC中国";
		System.out.println("转大写：" + str.toUpperCase());
		str = "AbdfSDD"; 
		System.out.println("转小写："+ str.toLowerCase());
		
		str = "         大家最近        都非常努力            ";
		System.out.println("去除字符串首尾的空格："+ str.trim());
	}
}
```

```java
public class Demo {
	public static void main(String[] args) {
		String str  ="        傻逼        大傻逼             ";	
		System.out.println(myTrim(str));
		
		str =  "D:\\20120512\\day12\\Demo.java";
		getFileName(str);
		
		str = "新中国好";
		System.out.println("翻转后的字符串："+ reverse(str));
		
		str = "abcjavaabcjavaphpjava";  //java
		getCount(str, "java");
		
	}
	
	//统计子串出现 的次数
	public static void getCount(String str,String target){
		int count = 0 ; //用于记录出现的次数
		int fromIndex  = 0;  //记录从那个索引值开始寻找目标子串
		while((fromIndex = str.indexOf(target, fromIndex))!=-1){
			//如果indexof方法返回 的不是-1，那么就是已经找到了目标 元素。
			count++;
			fromIndex = fromIndex+target.length();
		}
		System.out.println("出现的次数："+ count);
	}
	
	
	
	public static String reverse(String str){
		char[] arr = str.toCharArray();
		for(int startIndex = 0 , endIndex=arr.length-1 ; startIndex<endIndex; startIndex++,endIndex--){
				char temp = arr[startIndex];
				arr[startIndex] = arr[endIndex];
				arr[endIndex] = temp;
		}
		//使用字符数组构建一个字符串。
		return new String(arr);
	}
	
	//需求2： 获取上传文件名  "D:\\20120512\\day12\\Demo1.java"。
	public static void getFileName(String path){
		int index = path.lastIndexOf("\\");
		String fileName = path.substring(index+1);
		System.out.println("文件名："+ fileName);
	}
	
	
	
//	需求1：自己实现trim的方法。
	public static String myTrim(String str){
		//先转换成字符 数组
		char[] arr = str.toCharArray();
		//定义两个 变量记录开始与结束 的索引值
		int startIndex = 0 ;
		int endIndex = arr.length -1;
		//确定开始 的索引值
		while(true){
			if(arr[startIndex]==' '){
				startIndex++;
			}else{
				break;
			}
		}
		//确定结束 的索引值：
		while(true){
			if(arr[endIndex]==' '){
				endIndex--;
			}else{
				break;
			}
		}
		//截取子串返回
		return str.substring(startIndex,endIndex+1);		
	}
}
```
### StringBuffer

 字符串特点：字符串是常量；它们的值在创建之后不能更改.
 字符串的内容一旦发生了变化，那么马上会创建一个新的对象。

 注意： 字符串的内容不适宜频繁修改，因为一旦修改马上就会创建一个新的对象。
 
 如果需要频繁修改字符串 的内容，建议使用字符串缓冲 类（StringBuffer）。
 
 `StringBuffer` 其实就是一个存储字符 的容器。

 笔试题目：使用Stringbuffer无 参的构造函数创建 一个对象时，默认的初始容量是多少？ 如果长度不够使用了，自动增长多少倍？
	StringBuffer 底层是依赖了一个`字符数组`才能存储字符数据 的，该字符串数组默的初始容量是`16`， 如果字符数组的长度不够使用，`自动增长1倍`。

 ```java
public class Demo {
	public static void main(String[] args) {
		//先使用StringBuffer无参的构造函数创建一个字符串缓冲类。
		StringBuffer sb = new StringBuffer(); 
		sb.append("java");
		sb.append("java");
		sb.append("java");
		sb.append("java");
		sb.append("java");
		
		System.out.println(sb);
	}
}

```java
增加
	append(boolean b)    可以添加任意类型 的数据到容器中
	insert(int offset, boolean b)  指定插入的索引值，插入对应的内容。 

删除
	delete(int start, int end)  根据指定的开始与结束的索引值删除对应的内容。
	deleteCharAt(int index)   根据指的索引值删除一个字符。


修改

	replace(int start, int end, String str) 根据指定 的开始与结束索引值替代成指定的内容。
	reverse()   翻转字符串缓冲类的内容。  abc--->cba
	
	setCharAt(int index, char ch)  把指定索引值的字符替换指定的字符。
	substring(int start, int end)  根据指定的索引值截取子串。
	ensureCapacity(int minimumCapacity)  指定StringBuffer内部的字符数组长度的。
	
查看
	indexOf(String str, int fromIndex) 查找指定的字符串第一次出现的索引值,并且指定开始查找的位置。
	lastIndexOf(String str) 
	
	capacity() 查看当前字符数组的长度。 
	length() 
	
	charAt(int index) 
	toString()   把字符串缓冲类的内容转成字符串返回。
```

```java
public class Demo {
	public static void main(String[] args) {

		//先使用StringBuffer无参的构造函数创建一个字符串缓冲类。
		StringBuffer sb = new StringBuffer(); 
		sb.append("abcjavaabc");
		
		//添加 
		sb.append(true);
		sb.append(3.14f);

		//插入
		sb.insert(2, "小明");
		
		
		
		//删除
		sb.delete(2, 4); //  删除的时候也是包头不包尾
		sb.deleteCharAt(3); //根据指定 的索引值删除一个字符
		
		//修改	
		sb.replace(2, 4, "陈小狗");
		
		// 翻转字符串的内容
		sb.reverse(); 
		
		//把指定索引值的字符替换指定的字符。 
		sb.setCharAt(3, '红');
		
		String subString = sb.substring(2, 4);
		System.out.println("子串的内容："+ subString);
		

		查看
		int index = sb.indexOf("abc", 3);
		System.out.println("索引值为："+index);
			
		sb.append("javajava");
		System.out.println("查看字符数组的长度："+ sb.capacity());
		
		System.out.println("存储的字符个数："+sb.length());
		System.out.println("索引指定的索引值查找字符："+sb.charAt(2) );
		System.out.println("字符串缓冲类的内容："+ sb);
		
		String content = sb.toString();
		test(content);
	}
	
	public static void test(String str){
	}
}
```

```java

import java.util.Random;

public class Demo {
	public static void main(String[] args) {

		char[] arr={'s','b','g','h','过','傻'};
		StringBuilder sb=new StringBuilder();
		Random  random=new Random();
		for(int i=0;i<4;i++){
			int index=random.nextInt(arr.length);
			sb.append(arr[index]);
		}
		System.out.println("验证码:"+sb);
	} 
}
```
StringBuffer 与 StringBuilder的相同处与不同处：
	相同点：
		1. 两个类都是字符串缓冲类。
		2. 两个类的方法都是一致的。
	不同点：
		1. `StringBuffer`是线程安全的,操作效率低 ，`StringBuilder`是线程非安全的,操作效率高。
		2. StringBuffer是jdk1.0出现 的，StringBuilder 是jdk1.5的时候出现的。	
推荐使用： StringBuilder，因为操作效率高。


String StringBuffer 和 StringBuilder 的区别是什么? String 为什么是不可变的?

简单的来说：`String` 类中使用 `final` 关键字修饰字符数组来保存字符串，`private final char value[]`，所以` String` 对象是不可变的。

而 `StringBuilder` 与 `StringBuffer` 都继承自 `AbstractStringBuilder` 类，在 `AbstractStringBuilder` 中也是使用字符数组保存字符串`char[]value` 但是没有用 `final` 关键字修饰，所以这两种对象都是可变的。

```java
abstract class AbstractStringBuilder implements Appendable, CharSequence {
    /**
     * The value is used for character storage.
     */
    char[] value;

    /**
     * The count is the number of characters used.
     */
    int count;

    AbstractStringBuilder(int capacity) {
        value = new char[capacity];
    }
}
```

**线程安全性**

`String` 中的对象是不可变的，也就可以理解为常量，线程安全。`AbstractStringBuilder` 是 `StringBuilder` 与 `StringBuffer` 的公共父类，定义了一些字符串的基本操作，如 `expandCapacity`、`append`、`insert`、`indexOf` 等公共方法。`StringBuffer` 对方法加了同步锁或者对调用的方法加了同步锁，所以是线程安全的。`StringBuilder` 并没有对方法进行加同步锁，所以是非线程安全的。

**性能**

每次对 `String` 类型进行改变的时候，都会生成一个新的 `String` 对象，然后将指针指向新的 `String` 对象。`StringBuffer` 每次都会对 `StringBuffer` 对象本身进行操作，而不是生成新的对象并改变对象引用。相同情况下使用 `StringBuilder` 相比使用 `StringBuffer` 仅能获得 10%~15% 左右的性能提升，但却要冒多线程不安全的风险。

**对于三者使用的总结：**

1. 操作少量的数据: 适用 `String`
2. 单线程操作字符串缓冲区下操作大量数据: 适用 `StringBuilder`
3. 多线程操作字符串缓冲区下操作大量数据: 适用 `StringBuffer`

### Scanner

键盘输入常用的两种方法
```java
//方法 1：通过 Scanner
Scanner input = new Scanner(System.in);
String s  = input.nextLine();
input.close();
```

```java
//方法 2：通过 BufferedReader
BufferedReader input = new BufferedReader(new InputStreamReader(System.in));
String s = input.readLine();
```

###  System 系统类

`System` 系统类 主要用于获取系统的属性数据。
System类常用的方法：
 	`arraycopy(Object src, int srcPos, Object dest, int destPos, int length)` 一般
 		src - 源数组。
		srcPos - 源数组中的起始位置。
		dest - 目标数组。
		destPos - 目标数据中的起始位置。
		length - 要复制的数组元素的数量。
		
	`currentTimeMillis()`  获取当前系统时间。  重点
```java
//System.currentTimeMillis()去代替new Date()，效率上会高一点。

//获得系统的时间，单位为毫秒,转换为秒
long totalMilliSeconds = System.currentTimeMillis();
long totalSeconds = totalMilliSeconds / 1000;
	
//求出现在的秒
long currentSecond = totalSeconds % 60;
	
//求出现在的分
long totalMinutes = totalSeconds / 60;
long currentMinute = totalMinutes % 60;
	
//求出现在的小时
long totalHour = totalMinutes / 60;
long currentHour = totalHour % 24;
	
//显示时间
System.out.println("总毫秒为： " + totalMilliSeconds);
System.out.println(currentHour + ":" + currentMinute + ":" + currentSecond + " GMT");
```

	`exit(int status)`  退出jvm  如果参数是`0`表示正常退出jvm，`非0`表示异常退出jvm。    一般

	`gc()`    建议jvm赶快启动垃圾回收期回收垃圾。

	`getenv(String name)` 根据环境变量的名字获取环境变量。

	`getProperty(key)`
	
	`finalize()`  如果一个对象被垃圾回收器回收的时候，会先调用对象的finalize()方法。

```java

package cn.ixfosa.other;

import java.util.Arrays;
import java.util.Properties;

class Person{
	
	String name;

	public Person(String name) {
		this.name = name;
	}
	
	@Override
	public void finalize() throws Throwable {
		super.finalize();
		System.out.println(this.name+"被回收了..");
	}
}

public class Demo {
	
	public static void main(String[] args) {
		
		int[] srcArr = {10,12,14,16,19};
		//把srcArr的数组元素拷贝 到destArr数组中。
		int[] destArr = new int[4];
		
		System.arraycopy(srcArr, 1, destArr, 0,4);
		//System.exit(0); //jvm退出..  注意： 0或者非0的 数据都可以退出jvm。对于用户而言没有任何区别。
		System.out.println("目标数组的元素："+ Arrays.toString(destArr)); // 0 14 16 0
		System.out.println("当前的系统时间：" + System.currentTimeMillis());
		System.out.println("环境变量："+System.getenv("JAVA_HOME"));
		
		for(int i = 0 ; i<4; i++){
			new Person("狗娃"+i);
			System.gc(); //建议马上启动垃圾回收期
		}
		
		Properties properties = System.getProperties();  //获取系统的所有属性。
		properties.list(System.out);
		*/
		String value = System.getProperty("os.name"); //根据系统的属性名获取对应的属性值
		System.out.println("当前系统："+value);
	}
}
```

###  RunTime类

 `RunTime`该类类主要代表了应用程序运行的环境。
 
 	`getRuntime()`  返回当前应用程序的运行环境对象。
 	`exec(String command)`  根据指定的路径执行对应的可执行文件。
 	`freeMemory()`   返回 Java 虚拟机中的空闲内存量。。 以字节为单位
 	`maxMemory()`    返回 Java 虚拟机试图使用的最大内存量。
 	`totalMemory()`    返回 Java 虚拟机中的内存总量

```java
package cn.ixfosa.other;

import java.io.IOException;
import javax.management.RuntimeErrorException;

public class Demo {

	public static void main(String[] args) throws IOException, InterruptedException {

		Runtime runtime = Runtime.getRuntime();
		Process process = runtime.exec("C:\\Windows\\notepad.exe");

		//Thread.sleep(3000); //让当前程序停止3秒。

		process.destroy();
		System.out.println(" Java虚拟机中的空闲内存量。"+runtime.freeMemory());
		System.out.println("Java 虚拟机试图使用的最大内存量:"+ runtime.maxMemory());
		System.out.println("返回 Java 虚拟机中的内存总量:"+ runtime.totalMemory());
	}
}
```

### Date 日期类

Date 日期类

Calendar 抽象类 用于设置和获取日期/时间数据的特定部分

SimpleDateFormat 日期格式化类

```java

package cn.ixfosa.other;

import java.text.ParseException;
import java.text.SimpleDateFormat;
import java.util.Calendar;
import java.util.Date;

public class Demo  {
	
	public static void main(String[] args) throws ParseException {

		Date date = new Date(); // 获取当前的系统时间
		System.out.println("年份："+ date.getYear());

		Calendar calendar = Calendar.getInstance(); //获取当前的系统时间。
		System.out.println("年："+ calendar.get(Calendar.YEAR));
		System.out.println("月："+ (calendar.get(Calendar.MONTH)+1)); //month是从0开始的，而月份是从1开始的，所以month需要加`1`。
		System.out.println("日："+ calendar.get(Calendar.DATE));
		
		System.out.println("时："+ calendar.get(Calendar.HOUR_OF_DAY));
		System.out.println("分："+ calendar.get(Calendar.MINUTE));
		System.out.println("秒："+ calendar.get(Calendar.SECOND));
		
		// 显示 当前系统时间: 2014年12月26日  xx时xx分xx秒   
		 /*  日期格式化类    SimpleDateFormat 
		 *  		作用1： 可以把日期转换转指定格式的字符串     format()
		 *  		作用2： 可以把一个 字符转换成对应的日期。    parse()
		 *  	
		 */
		Date date = new Date(); //获取当前的系统时间。
		SimpleDateFormat dateFormat = new SimpleDateFormat("yyyy年MM月dd日   HH:mm:ss") ; //使用了默认的格式创建了一个日期格式化对象。
		String time = dateFormat.format(date);  //可以把日期转换转指定格式的字符串
		System.out.println("当前的系统时间："+ time);
		
		String birthday = "1998年02月19日   11:11:11";
		Date date2 = dateFormat.parse(birthday);  //注意： 指定的字符串格式必须要与SimpleDateFormat的模式要一致。
		System.out.println(date2);
	}
}
```

###  Math 数学类

 Math 数学类， 主要是提供了很多的数学公式。
	`abs(double a)`  获取绝对值
	`ceil(double a)`  向上取整
	`floor(double a)`  向下取整
	`round(float a)`   四舍五入
	`random()`   产生一个随机数. 大于等于 0.0 且小于 1.0 的伪随机 `double` 值
 
```java

package cn.ixfosa.other;

public class Demo {
	public static void main(String[] args) {

		System.out.println("绝对值:"+Math.abs(-3)); //绝对值:3        
		System.out.println("向上取整："+Math.ceil(3.14)); //向上取整：4.0
		System.out.println("向下取整："+Math.floor(-3.14)); //向下取整：-4.0
		System.out.println("四舍五入:"+Math.round(3.54));//四舍五入:4
		System.out.println("随机数："+Math.random());//随机数：0.1694973050015567
	}
}
```

### Random 随机数类

1. 不带参数的nextInt()会生成所有有效的整数（包含正数，负数，0）

2. 带参的nextInt(int x)则会生成一个范围在`0~x`（不包含X）内的任意正整数
　　例如：
		int x=new Random.nextInt(100);    //则x为一个0~99的任意整数
		int i=rand.nextInt(100)+1;        //随机生成从1到100的随机整数
		int j=(int)(Math.random()*99)+1;  //随机生成从1到100的随机整数

3. 生成一个指定范围内的整数
```java
/*
* 生成[min, max]之间的随机整数
* @param min 最小整数
* @param max 最大整数
*/
public static int randomInt(int min, int max){
	return new Random().nextInt(max) % (max-min+1) + min;
}
//例如：调用方法randomInt（10,20）;则会生成一个[10,20]集合内的任意整数
```

```java
package cn.ixfosa.other;

import java.util.Random;

//需求： 编写一个函数随机产生四位的验证码。
public class Demo {

	public static void main(String[] args) {
		
		Random random0 = new Random();
		int randomNum = random0.nextInt(10)+1; //产生 的 随机数就是1-10之间
		System.out.println("随机数："+ randomNum);
		
		char[] arr = {'中','国','传','a','Q','f','B'};
		StringBuilder sb = new StringBuilder();
		Random random = new Random();
		//需要四个随机数，通过随机数获取字符数组中的字符，
		for(int i  = 0 ; i< 4 ; i++){
			int index = random.nextInt(arr.length);  //产生的 随机数必须是数组的索引值范围之内的。
			sb.append(arr[index]);
		}
		System.out.println("验证码："+ sb);
	}
}
```

## 多线程

### 进程和线程
`进程`：正在执行的程序称作为一个进程。进程负责了内存空间的划分。

问题：windows号称是多任务的操作系统,那么windows是同时运行多个应用程序吗？
	从宏观的角度： windows确实是在同时运行多个应用程序。
	从微观角度： cpu是做了一个快速切换执行的动作，由于速度态度，所以我感觉不到在切换而已。

`线程`：线程在一个进程中负责了代码的执行，就是进程中一个执行路径，
一个线程不能独立的存在，它必须是进程的一部分。一个进程一直运行，直到所有的非守护线程都结束运行后才能结束。

`多线程`： 在一个进程中有多个线程同时在执行不同的任务。


疑问：线程负责了代码的执行，我们之前没有学过线程，为什么代码可以执行呢？
	运行任何一个java程序，jvm在运行的时候都会创建一个`main线程`执行main方法中所有代码。

一个java应用程序至少有几个线程？
	`至少有两个线程`， 一个是主线程负责`main方法`代码的执行，一个是`垃圾回收器线程`,负责了回收垃圾。

多线程的好处：
	1. 解决了一个进程能同时执行多个任务的问题。
	2. 提高了资源的利用率。

多线程 的弊端：
	1. 增加cpu的负担。
	2. 降低了一个进程中线程的执行概率。
	3. 引发了线程安全问题。
	4. 出现了死锁现象。

### 一个线程的生命周期
线程是一个动态执行的过程，它也有一个从产生到死亡的过程。
![alt ](\images\java-thread.jpg)

新建状态:
	使用 new 关键字和 Thread 类或其子类建立一个线程对象后，该线程对象就处于新建状态。它保持这个状态直到程序 start() 这个线程。

就绪状态:
	当线程对象调用了start()方法之后，该线程就进入就绪状态。就绪状态的线程处于就绪队列中，要等待JVM里线程调度器的调度。

运行状态:
	如果就绪状态的线程获取 CPU 资源，就可以执行 run()，此时线程便处于运行状态。处于运行状态的线程最为复杂，它可以变为阻塞状态、就绪状态和死亡状态。

阻塞状态:
	如果一个线程执行了sleep（睡眠）、suspend（挂起）等方法，失去所占用资源之后，该线程就从运行状态进入阻塞状态。在睡眠时间已到或获得设备资源后可以重新进入就绪状态。可以分为三种：

	等待阻塞：运行状态中的线程执行 wait() 方法，使线程进入到等待阻塞状态。

	同步阻塞：线程在获取 synchronized 同步锁失败(因为同步锁被其他线程占用)。

	其他阻塞：通过调用线程的 sleep() 或 join() 发出了 I/O 请求时，线程就会进入到阻塞状态。当sleep() 状态超时，join() 等待线程终止或超时，或者 I/O 处理完毕，线程重新转入就绪状态。

死亡状态:
	一个运行状态的线程完成任务或者其他终止条件发生时，该线程就切换到终止状态。

![alt ](\images\线程的生命周期状态图.png)

### 创建多线程：

Java 提供了三种创建线程的方法：

	通过实现 `Runnable` 接口；
	通过继承 `Thread` 类本身；
	通过 `Callable` 和 `Future` 创建线程。

#### Thread 方式

1. 自定义一个类继承Thread类。

2. 重写Thread类的run方法 , 把自定义线程的任务代码写在run方法中
	疑问： 重写run方法的目的是什么？  
	每个线程都有自己的任务代码，jvm创建的主线程的任务代码就是main方法中的所有代码, 自定义线程的任务代码就写在run方法中，自定义线程负责了run方法中代码。

3. 创建Thread的子类对象，并且调用start方法开启线程。

注意：	一个线程一旦开启，那么线程就会执行run方法中的代码，run方法千万不能直接调用，直接调用run方法就相当调用了一个普通的方法而已
并没有开启新的线程。

Thread类 本质上也是实现了 `Runnable` 接口的一个实例。

```java
public class Demo extends Thread {

	@Override  //把自定义线程的任务代码写在run方法中。
	public void run() {
		for(int i  = 0 ; i < 100 ; i++){
			System.out.println("自定义线程："+i);
		}
	}

	public static void main(String[] args) {
		//创建了自定义的线程对象。
		Demo d = new Demo();
		//调用start方法启动线程
		d.start();
		
		for(int i  = 0 ; i < 100 ; i++){
			System.out.println("main线程："+i);
		}
	}
}
```

```java
/*
 需求： 模拟QQ视频与聊天同时在执行。
 */

class TalkThread extends Thread{

	@Override
	public void run() {
		while(true){
			System.out.println("hi，你好！开视频呗...");
		}
	}
}

class VideoThread extends Thread{

	@Override
	public void run() {
		while(true){
			System.out.println("视频视频....");
		}
	}
}

public class Demo {
	public static void main(String[] args) {

		TalkThread talkThread = new TalkThread();
		talkThread.start();
		VideoThread videoThread = new VideoThread();
		videoThread.start();
	}
}
```

#### Runnable接口

1. 自定义一个类实现Runnable接口。
2. 实现Runnable接口 的run方法，把自定义线程的任务定义在run方法上。
3. 创建Runnable实现类对象。
4. 创建Thread类的对象，并且把Runnable实现类的对象作为实参传递。
5. 调用Thread对象的start方法开启一个线程。

注意：千万不要直接调用run方法，调用start方法的时候线程就会开启，线程一旦开启就会执行run方法中代码，如果直接调用run方法，那么就相当于调用了一个普通的方法而已。


问题1：请问Runnable实现类的对象是线程对象吗？
	Runnable实现类的对象并不是一个线程对象，只不过是实现了Runnable接口的对象而已。
	只有是Thread或者是Thread的子类才是线程对象。

问题2：为什么要把Runnable实现类的对象作为实参传递给Thread对象呢？作用是什么？
	作用就是把Runnable实现类的对象的run方法作为了线程的任务代码去执行了。

推荐使用： 第二种。 实现Runable接口的。 
	原因： 因为java单继承 ,多实现的。

```java

//Thread类 的run方法
@Override
public void run() {
	if (target != null) {
		target.run();  //就相当于Runnable实现类的对象的run方法作为了Thread对象的任务代码了。
	}
}

```

```java

public class Demo implements Runnable{

	@Override
	public void run() {
		//System.out.println("this："+ this);
		System.out.println("当前线程："+ Thread.currentThread());
		for(int i = 0 ; i < 100 ; i++){
			System.out.println(Thread.currentThread().getName()+":"+i);
		}
	}

	public static void main(String[] args) {

		//创建Runnable实现类的对象
		Demo d = new Demo();
		//创建Thread类的对象， 把Runnable实现类对象作为实参传递。
		Thread thread = new Thread(d,"狗娃");  //Thread类使用Target变量记录了d对象，
		//调用thread对象的start方法开启线程。
		thread.start();
		
		for(int i = 0 ; i < 100 ; i++){
			System.out.println(Thread.currentThread().getName()+":"+i);
		}
	}
}
```

#### Callable 和 Future 创建线程

1. 创建 `Callable` 接口的实现类，并实现 `call()` 方法，该 call() 方法将作为线程执行体，并且有返回值。

2. 创建 Callable 实现类的实例，使用 `FutureTask` 类来包装 Callable 对象，该 FutureTask 对象封装了该 Callable 对象的 call() 方法的返回值。

3. 使用 FutureTask 对象作为 `Thread` 对象的 `target` 创建并启动新线程。

4. 调用 FutureTask 对象的 `get()` 方法来获得子线程执行结束后的返回值。

创建线程的三种方式的对比
   1. 采用实现 Runnable、Callable 接口的方式创建多线程时，线程类只是实现了 Runnable 接口或 Callable 接口，还可以继承其他类。
   2. 使用继承 Thread 类的方式创建多线程时，编写简单，如果需要访问当前线程，则无需使用 Thread.currentThread() 方法，直接使用 this 即可获得当前线程。

```java

public class CallableThreadTest implements Callable<Integer> {
    public static void main(String[] args)  {  

        CallableThreadTest ctt = new CallableThreadTest();  
        FutureTask<Integer> ft = new FutureTask<>(ctt);

        for(int i = 0;i < 100;i++)  {  

            System.out.println(Thread.currentThread().getName()+" 的循环变量i的值"+i);  
            if(i==20){  
                new Thread(ft,"有返回值的线程").start();  
            }  
        }  
        try {
            System.out.println("子线程的返回值："+ft.get());  
        } catch (InterruptedException e) {  
            e.printStackTrace();  
		} catch (ExecutionException e) {  
            e.printStackTrace();  
        }  
  
    }
    @Override  
    public Integer call() throws Exception {  
        int i = 0;  
        for(;i<100;i++) {  
            System.out.println(Thread.currentThread().getName()+" "+i);  
        }  
        return i;  
    }  
}

```
### 线程常用的方法：
	`Thread(String name)`     初始化线程的名字
	`setName(String name)`    设置线程对象名
	`getName()`               返回线程的名字
	
	`sleep()`                 线程睡眠指定的毫秒数。 静态的方法， 那个线程执行了sleep方法代码那么就是那个线程睡眠。
	
	`currentThread()`      返回`当前的线程对象`，该方法是一个静态的方法， 注意： 那个线程执行了currentThread()代码就返回那个线程的对象。
	
	`getPriority()`             返回当前线程对象的优先级默认线程的优先级是`5`
	`setPriority(int newPriority`) 设置线程的优先级    虽然设置了线程的优先级，但是具体的实现取决于底层的操作系统的实现（最大的优先级是10 ，最小的1 ， 默认是5）。

```java

public class Demo extends Thread {
	public Demo(String name){
		super(name);  //调用了Thread类的一个 参数的构造方法。
	}

	@Override
	public void run() {

		//System.out.println("this:"+ this);
		System.out.println("当前线程对象：" + Thread.currentThread());	
		
		for (int i = 0; i < 100 ; i++) {
			System.out.println(Thread.currentThread().getName()+":"+i);
			try {
				Thread.sleep(100);  //为什么在这里不能抛出异常，只能捕获？？ Thread类的run方法没有抛出异常类型，所以子类不能抛出异常类型。
			} catch (InterruptedException e) {
				e.printStackTrace();
			}
		}
	}

	public static void main(String[] args) throws InterruptedException {

		//创建了一个线程对象
		Demo d = new Demo("狗娃");
		d.setPriority(10); //设置线程 的优先级。 优先级的数字越大，优先级越高，优先级的范围是1~10
		d.start();
		
		for (int i = 0; i < 100 ; i++) {
			System.out.println(Thread.currentThread().getName()+":"+i);
		}
		
		System.out.println("自定义线程的优先级："+d.getPriority());  //线程的优先级����是5
		System.out.println("主线程的优先级："+Thread.currentThread().getPriority());
		
		d.start();
		
		d.setName("铁蛋"); //setName设置线程的名字
		d.start(); //开启线程
		
		Thread mainThread = Thread.currentThread();
		System.out.println("主线程的名字："+ mainThread.getName());
	}
}
```
### 线程安全问题

线程安全出现的根本原因：
	1. 存在两个或者两个以上的线程对象`共享同一个资源`。
	2. 多线程操作共享资源的代码有多句。

线程安全问题的解决方案:
	`同步代码块`: `synchronized(锁对象){ 需要被同步的代码 }`

	`同步函数`  ： 同步函数就是使用`synchronized`修饰一个函数。

#### 同步代码块

格式：
	synchronized(锁对象){
		需要被同步的代码
	}

同步代码块要注意的事项：
	1. 锁对象可以是任意的一个对象。
	2. 一个线程在同步代码块中sleep了，并不会释放锁对象。
	3. 如果不存在着线程安全问题，千万不要使用同步代码块，因为会降低效率。
	4. 锁对象必须是多线程共享的一个资源，否则锁不住。

```java
class SaleTicket extends Thread{

	static int num = 50; //票数  非静态的成员变量,非静态的成员变量数据是在每个对象中都会维护一份数据的。
	 
     static	Object o = new Object();

	 public SaleTicket(String name) {
		super(name);
	}

	@Override
	public void run() {
		while(true){
			//同步代码块
			synchronized ("锁") {
				if(num>0){
					System.out.println(Thread.currentThread().getName()+"售出了第"+num+"号票");
					try {
						Thread.sleep(100);
					} catch (InterruptedException e) {
						e.printStackTrace();
					}
					num--;
				}else{
					System.out.println("售罄了..");
					break;
				}
			}

		}
	}
}

public class Demo {
	public static void main(String[] args) {
		//创建三个线程对象，模拟三个窗口
		SaleTicket thread1 = new SaleTicket("窗口1");
		SaleTicket thread2 = new SaleTicket("窗口2");
		SaleTicket thread3 = new SaleTicket("窗口3");
		//开启线程售票
		thread1.start();
		thread2.start();
		thread3.start();
	}
}

```java
class SaleTicket implements Runnable{

	int  num = 50; // 票数

	@Override
	public void run() {
		while(true){
			synchronized ("锁") {
				if(num>0){
					System.out.println(Thread.currentThread().getName()+"售出了第"+ num+"号票");
					num--;
				}else{
					System.out.println("售罄了..");
					break;
				}
			}
		}		
	}
}

public class Demo {
	public static void main(String[] args) {
		//创建了一个Runnable实现类的对象
		SaleTicket saleTicket = new SaleTicket();
		//创建三个线程对象模拟三个窗口
		Thread thread1 = new Thread(saleTicket,"窗口1");
		Thread thread2 = new Thread(saleTicket,"窗口2");
		Thread thread3 = new Thread(saleTicket,"窗口3");
		//开启线程售票
		thread1.start();
		thread2.start();
		thread3.start();
	}
}

```

#### 同步函数

`同步函数` : 同步函数就是使用`synchronized`修饰一个函数。

同步函数要注意的事项 ：
	1. 如果是一个非静态的同步函数的锁 对象是this对象，如果是静态的同步函数的锁 对象是当前函数所属的类的字节码文件（class对象）。
	2. 同步函数的锁对象是固定的，不能由你来指定的。


推荐使用： 同步代码块。
	原因：
		1. 同步代码块的锁对象可以由我们随意指定，方便控制。同步函数的锁对象是固定的，不能由我们来指定。
		2. 同步代码块可以很`方便控制需要被同步代码的范围`，同步函数必须是`整个函数的所有代码都被同步`了。
```java

class BankThread extends Thread{
	static	int count = 5000;

	public BankThread(String name){
		super(name);
	}

	@Override  
	public synchronized  void run() {
		while(true){
			synchronized ("锁") {
				if(count>0){
					System.out.println(Thread.currentThread().getName()+"取走了1000块,还剩余"+(count-1000)+"元");
					count= count - 1000;
				}else{
					System.out.println("取光了...");
					break;
				}
			}
		}
	}

	//静态的函数---->函数所属的类的字节码文件对象--->BankThread.class  唯一的。
	public static synchronized  void getMoney(){
		
	}
}

public class Demo {
	public static void main(String[] args) {
		//创建两个线程对象
		BankThread thread1 = new BankThread("老公");
		BankThread thread2 = new BankThread("老婆");
		//调用start方法开启线程取钱
		thread1.start();
		thread2.start();
	}
}
```

### 死锁现象

java中同步机制解决了线程安全问题，但是也同时引发死锁现象。

死锁现象出现 的根本原因：
	1. 存在两个或者两个以上的线程。
	2. 存在两个或者两个以上的共享资源。

死锁现象的解决方案： 没有方案。只能尽量避免发生而已。

```java
class DeadLock extends Thread{
	public DeadLock(String name){
		super(name);
	}

	public void run() {
		if("张三".equals(Thread.currentThread().getName())){
			synchronized ("遥控器") {
				System.out.println("张三拿到了遥控器，准备 去拿电池!!");
				synchronized ("电池") {
					System.out.println("张三拿到了遥控器与电池了，开着空调爽歪歪的吹着...");
				}
			}
		}else if("狗娃".equals(Thread.currentThread().getName())){
			synchronized ("电池") { 
				System.out.println("狗娃拿到了电池，准备去拿遥控器!!");
				synchronized ("遥控器") {
					System.out.println("狗娃拿到了遥控器与电池了，开着空调爽歪歪的吹着...");
				}
			}
		}
	}
}

public class Demo {
	public static void main(String[] args) {

		DeadLock thread1 = new DeadLock("张三");
		DeadLock thread2 = new DeadLock("狗娃");
		//开启线程
		thread1.start();
		thread2.start();
	}
}

```

### 线程通讯

线程通讯： 一个线程完成了自己的任务时，要通知另外一个线程去完成另外一个任务.

生产者与消费者

`wait()`:  等待   如果线程执行了wait方法，那么该线程会进入等待的状态，等待状态下的线程必须要被其他线程调用notify方法才能唤醒。
`notify()`： 唤醒    唤醒线程池等待线程其中的一个。
`notifyAll()` : 唤醒线程池所有等待 线程。

wait与notify方法要注意的事项：
	1. wait方法与notify方法是属于Object对象的。
	2. wait方法与notify方法必须要在`同步代码块`或者是`同步函数`中才能使用。
	3. wait方法与notify方法必需要由`锁对象`调用。

```java
//产品类
class Product{
	String name;  //名字
	double price;  //价格
	boolean flag = false; //产品是否生产完毕的标识，默认情况是没有生产完成。
}

//生产者
class Producer extends Thread{
	Product  p ;  	//产品
	public Producer(Product p) {
		this.p  = p ;
	}

	@Override
	public void run() {
		int i = 0 ;
		while(true){
		 synchronized (p) {

			if(p.flag==false){
				 if(i%2==0){
					 p.name = "苹果";
					 p.price = 6.5;
				 }else{
					 p.name="香蕉";
					 p.price = 2.0;
				 }

				 System.out.println("生产者生产出了："+ p.name+" 价格是："+ p.price);

				 p.flag = true;
				 i++;

				 p.notifyAll(); //唤醒消费者去消费
			}else{
				//已经生产 完毕，等待消费者先去消费
				try {
					p.wait();   //生产者等待
				} catch (InterruptedException e) {
					e.printStackTrace();
				}
			}
		}
	  }
	}
}

//消费者
class Customer extends Thread{
	Product p;
	public  Customer(Product p) {
		this.p = p;
	}

	@Override
	public void run() {
		while(true){
			synchronized (p) {

				if(p.flag==true){  //产品已经生产完毕

					System.out.println("消费者消费了"+p.name+" 价格："+ p.price);
					p.flag = false;

					p.notifyAll(); // 唤醒生产者去生产

				}else{
					//产品还没有生产,应该 等待生产者先生产。
					try {
						p.wait(); //消费者也等待了...
					} catch (InterruptedException e) {
						e.printStackTrace();
					}
				}
			}
		}
	}
}

public class Demo {
	public static void main(String[] args) {

		Product p = new Product();  //产品
		//创建生产对象
		Producer producer = new Producer(p);
		//创建消费者
		Customer customer = new Customer(p);
		//调用start方法开启线程
		producer.start();
		customer.start();
		
	}
}
```

###  线程的停止

 线程的停止：
 	1. 停止一个线程 我们一般都会通过一个变量去控制的。
 	2. 如果需要停止一个处于等待状态下的线程，那么我们需要通过变量配合notify方法或者interrupt()来使用。

```java

public class Demo extends Thread {

	boolean flag = true;
	public Demo(String name){
		super(name);
	}

	@Override
	public synchronized void run() {
		int i = 0 ;
		while(flag){
			try {

				this.wait(); //狗娃等待..

			} catch (InterruptedException e) {

				System.out.println("接收到了异常了....");
			}

			System.out.println(Thread.currentThread().getName()+":"+i);
			i++;
		}
	}

	public static void main(String[] args) {

		Demo d = new Demo("狗娃");

		d.setPriority(10);
		d.start();
		
		for(int i = 0 ; i<100 ; i++){

			System.out.println(Thread.currentThread().getName()+":"+i);
			//当主线程的i是80的时候停止狗娃线程。
			//d.interrupt();  // interrupt()根本就是无法停止一个线程。

			if(i==80){
				d.flag = false;
				d.interrupt(); //把线程的等待状态强制清除，被清除状态的线程会接收到一个InterruptedException。
				/*synchronized (d) {
					d.notify();
				}*/
			}
		}
	}
}
```
### 守护线程（后台线程）

 守护线程（后台线程）:在一个进程中如果只剩下了守护线程，那么守护线程也会死亡。

 一个线程默认都不是守护线程。

`setDaemon()` 设置线程是否为守护线程，true为守护线程， false为非守护线程。
`isDaemon()`  判断线程是否为守护线程。

 ```java

public class Demo extends Thread {
	public Demo(String name){
		super(name);
	}

	@Override
	public void run() {
		for(int i = 1 ; i<=100 ; i++){
			System.out.println("更新包目前下载"+i+"%");
			if(i==100){
				System.out.println("更新包下载完毕,准备安装..");
			}
			try {
				Thread.sleep(100);
			} catch (InterruptedException e) {
				e.printStackTrace();
			}
		}
	}

	public static void main(String[] args) {
		
		 Demo d = new Demo("后台线程");

		d.setDaemon(true); //setDaemon() 设置线程是否为守护线程，true为守护线程， false为非守护线程。
		System.out.println("是守护线程吗？"+ d.isDaemon()); //判断线程是否为守护线程。
		d.start();

		for(int i = 1 ; i<=100 ; i++){
			System.out.println(Thread.currentThread().getName()+":"+i);
		}
	}
}
 ```

 ###  join方法

 join方法。 加入
一个线程如果执行join语句，那么就有新的线程加入，执行该语句的线程必须要让步给新加入的线程先完成任务，然后才能继续执行。

 ```java
//老妈
class  Mon extends Thread{
	public void run() {

		System.out.println("妈妈洗菜");
		System.out.println("妈妈切菜");
		System.out.println("妈妈准备炒菜，发现没有酱油了..");

		//叫儿子去打酱油
		Son s= new Son();
		s.start();

		try {
			s.join();  //加入。
		} catch (InterruptedException e) {
			e.printStackTrace();
		}
		
		System.out.println("妈妈继续炒菜");
		System.out.println("全家一起吃饭..");		
	}
}

class Son extends Thread{

	@Override
	public void run() {
		System.out.println("儿子下楼..");

		try {
			Thread.sleep(1000);
		} catch (InterruptedException e) {
			e.printStackTrace();
		}

		System.out.println("儿子一直往前走");
		System.out.println("儿子打完酱油了");
		System.out.println("上楼，把酱油给老妈");
	}
}

public class Demo {
	public static void main(String[] args) {

		Mon m = new Mon();
		m.start();
	}
}
 ```

 ## collection 单例集合

 数组： 存储`同一种数据类型`的集合容器.

数组的特点：
	1. 只能存储同一种数据类型的数据。
	2. 一旦初始化，`长度固定`。
	3. 数组中的元素与元素之间的`内存地址是连续`的。

注意： `Object`类型的数组可以存储任意类型的数据。

单例集合
----------| Collection 单列集合的根接口
----------------| List 如果是实现了List接口集合类具备的特点： 有序，可重复。
-------------------| ArrayList   底层使用Object数组实现的。 特点： 查询速度快，增删慢。
-------------------| LinkedList  底层是使用了链表数据数据结构实现的。 特点： 查询慢，增删快。
-------------------| Vector(了解)   底层使用Object数组实现的， 实现与ArrayList是一样，只不过是线程安全的，操作效率低。

----------------| Set  如果是实现了Set接口集合类具备的特点： 无序，不可重复。
------------------| HashSet  底层使用的是哈希表实现的。  
------------------| TreeSet  底层使用二叉数实现。 

双列集合：
--------| Map  (只需要把Map接口的方法全部练习一次即可。)
-----------| HashMap  底层使用的是哈希表实现的。  
-----------| TreeMap  底层使用二叉数实现
-----------| HashTable（了解）

```java
public class Demo {
	public static void main(String[] args) {

		Object[] arr = new Object[10];
		arr[1] = "abc";
		arr[2] = 'a';
		arr[3]  = 12;
	}
}
```
集合：集合是存储对象数据的集合容器。

集合比数组的优势：
	1. 集合可以存储任意类型的对象数据，数组只能存储同一种数据类型的数据。
	2. 集合的长度是会发生变化的，数组的长度是固定的。
   
-------| Collection  `单例集合`的根接口。 
 ----------| List  如果是实现了`List`接口的集合类，具备的特点： `有序`，`可重复`。
 ----------| Set   如果是实现了`Set`接口的集合类，具备特点： `无序`，`不可重复`。

Collection接口中的方法：
	增加
		`add(E e)`  添加成功返回true，添加 失败返回false.
		`addAll(Collection c)`  把一个集合的元素添加到另外一个集合中去。
	
	删除
		`clear()`   清空集合中的元素
		`remove(Object o)`  指定集合中的元素删除，删除成功返回true，删除失败返回false.
		
		`removeAll(Collection  c)`  删除c集合中与c2的交集元素。
		`retainAll(Collection  c)`  保留c集合与c2的交集元素，其他的元素一并删除。
	
	查看
		`size()` 

	
	判断
		`isEmpty()`  判断集合是否为空元素
		`contains(Object o)`    contains方法内部是依赖于equals方法进行比较的。
		`containsAll(Collection<?> c)` 
	
	迭代
		`toArray()` 
		`iterator()`
```java

package cn.icfosa.collection;

import java.util.ArrayList;
import java.util.Collection;

public class Demo {
	public static void main(String[] args) {

		Collection c = new ArrayList();

		c.add("令计划");
		c.add("徐才厚");
		c.add("周永康");

		System.out.println("添加成功吗？"+c.add("郭美美"));
		
		//创建集合
		Collection c2 = new ArrayList();
		c2.add("徐才厚");
		c2.add("郭美美");
		c2.add("狗娃");
		
		c.addAll(c2); //把c2的元素的添加到c集合 中去。
		
		//删除方法:
		c.clear(); //clear()清空集合中的元素
		System.out.println("删除成功吗？"+c.remove("美美"));  // remove 指定集合中的元素删除，删除成功返回true，删除失败返回false.
		
		c.removeAll(c2); //删除c集合中与c2的交集元素。
		
		c.retainAll(c2); //保留c集合与c2的交集元素，其他的元素一并删除。
		
		System.out.println("查看元素个数："+c.size());

		System.out.println("集合的元素："+ c);
	}
}
```

```java

import java.util.ArrayList;
import java.util.Collection;

/*
判断
	isEmpty()  
	contains(Object o) 
	containsAll(Collection<?> c)
*/

class Person{
	int id; 
	String name;
	public Person(int id, String name) {
		this.id = id;
		this.name = name;
	}

	@Override
	public String toString() {
		return "{编号："+this.id+" 姓名："+ this.name+"}";
	}

	@Override
	public boolean equals(Object obj) {
		Person p = (Person)obj;
		return this.id == p.id ;
	}

	//java规范： 一般重写equlas方法我们都会重写hashCode方法的。
	@Override
	public int hashCode() {
		return this.id;
	}
}

public class Demo {
	public static void main(String[] args) {

		Collection c = new ArrayList();

		c.add("令计划");
		c.add("徐才厚");
		c.add("周永康");

		System.out.println("判断集合是否为空元素："+ c.isEmpty());

		System.out.println("判断集合中是否存在指定的元素："+ c.contains("薄熙来"));
		
		//集合添加自定义的元素
		Collection c1 = new ArrayList();
		c1.add(new Person(110,"狗娃"));
		c1.add(new Person(119,"狗剩"));
		c1.add(new Person(120,"铁蛋"));
		
		
		Collection c2 = new ArrayList();
		c2.add(new Person(110,"狗娃"));
		c2.add(new Person(119,"狗剩"));
		c2.add(new Person(104,"陈狗剩"));
		
		System.out.println("c集合有包含c2集合中的所有元素吗？"+ c.containsAll(c2));
		
		//如果在现实生活中，只要身份证编号一致，那么就为同一个人。
		System.out.println("存在该元素吗？"+c.contains(new Person(120,"陈铁蛋"))); //其实contains方法内部是依赖于equals方法进行比较的。
		System.out.println("集合的元素："+ c);
}
```

```java
package cn.icfosa.collection;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collection;

/*
迭代
	toArray() 
*/
public class Demo {
	public static void main(String[] args) {

		Collection c = new ArrayList();

		c.add("令计划");
		c.add("徐才厚");
		c.add("周永康");

		Object[] arr = c.toArray(); //把集合中的元素全部 存储到一个Object的数组中返回。
		System.out.println("数组的元素："+Arrays.toString(arr));
		
		Collection c = new ArrayList();
		c.add(new Person(110,"狗娃"));
		c.add(new Person(119,"狗剩"));
		c.add(new Person(120,"铁蛋"));

		Object[] arr = c.toArray();

		//需求： 把编号是110的人信息 输出。
		for(int i = 0 ; i<arr.length ; i++){
			Person p = (Person) arr[i];  //从Object数组中取出的元素只能使用Object类型声明变量接收，如果需要其他的类型需要进行强制类型转换。
			if(p.id==110){
				System.out.println(p);
			}
		}
	}
}
```

### 迭代器

Collection---迭代的方法：
	`toArray()` 
	`iterator()`  返回的是iterator接口的实现类对象。

迭代器的作用：就是用于抓取集合中的元素。

迭代器的方法：

	`hasNext()`   问是否有元素可遍历。如果有元素可以遍历，返回true，否则返回false 。

 	`next()`    获取元素...
      	  
 	`remove()`  移除迭代器最后一次返回 的元素。

NoSuchElementException 没有元素的异常。
出现的原因： 没有元素可以被迭代了。。。

```java

import java.util.ArrayList;
import java.util.Collection;
import java.util.Iterator;

public class Demo {

	public static void main(String[] args) {

		Collection c = new ArrayList();
		c.add("狗娃");
		c.add("狗剩");
		c.add("铁蛋");
		c.add("美美");

		//遍历集合的元素------>方式一： 可以使用toArray方法。
		Object[] arr = c.toArray(); // toArray()  把集合 的元素存储到一个 Object的数组中 返回。
		for(int i = 0 ; i<arr.length ; i++){
			System.out.print(arr[i]+",");
		}
		
		//要求使用iterator迭代器遍历。
		Iterator it = c.iterator();  //返回一个迭代器    疑问：iterator()方法返回的是一个接口类型，为什么接口又可以调用方法可以使用呢？  iterator 实际 上返回的是iterator接口的实现类对象。
		
		while(it.hasNext()){ // hasNext() 问是否有元素可以遍历。
			System.out.println("元素："+ it.next()); //获取元素
		}

		//it.next();
		//it.remove();  //删除迭代器最后一次返回的元素。

		//清空集合 的元素
		while(it.hasNext()){
			it.next();
			it.remove();
		}
		System.out.println("集合的元素："+ c);
	}
}

```

```java

package cn.ixfosa.collelction;

import java.util.ArrayList;
import java.util.Collection;
import java.util.Iterator;
import java.util.Scanner;

class User{
	int id;  //账号
	String password;  //密码
	public int getId() {
		return id;
	}
	public void setId(int id) {
		this.id = id;
	}
	public String getPassword() {
		return password;
	}
	public void setPassword(String password) {
		this.password = password;
	}
	public User(int id, String password) {
		this.id = id;
		this.password = password;
	}

	@Override
	public boolean equals(Object obj) {
		User user = (User)obj;
		return this.id==user.id;
	}

	@Override
	public String toString() {
		return "{ 账号："+this.id+" 密码："+this.password+"}";
	}
}

public class Demo {

	static Scanner scanner = new Scanner(System.in);

	static Collection users = new ArrayList(); //使用该集合保存所有的用户信息..

	public static void main(String[] args) {

		while(true){
			System.out.println("请选择功能      A(注册 )    B(登陆)");
			String option = scanner.next();
			if("a".equalsIgnoreCase(option)){

				 reg();

			}else if("b".equalsIgnoreCase(option)){

				login();

			}else{
				System.out.println("你的选择有误,请重新输入");
			}
		}
	}

	public static void login() {
		System.out.println("请输入账号：");
		int id = scanner.nextInt();
		System.out.println("请输入密码:");
		String password = scanner.next();

		//判断集合的用户是否存在该用户名与密码
		//遍历集合的元素，查看是否存在该用户信息
		boolean isLogin = false; 	//定义变量用于记录是否登陆成功的信息  , 默认是没有登陆成功的
		Iterator it = users.iterator();
		while(it.hasNext()){
			User user = (User) it.next();
			if(user.id==id && user.password.equals(password)){
				//存在该用户信息，登陆成功...
				isLogin = true;
			}
		}

		if(isLogin==true){
			System.out.println("欢迎登陆...");
		}else{
			System.out.println("用户名或者密码错误，登陆失败...");
		}
	}

	public static void reg() {

		User user = null;

		while(true){
			System.out.println("请输入账号:");
			int id = scanner.nextInt();  
			user = new User(id,null);
			if(users.contains(user)){   //contains底层依赖了equals方法。
				//如果存在
				System.out.println("该账号已经存在，请重新输入账号");
			}else{
				//不存在
				break;
			}
		}
		
		System.out.println("请输入密码：");
		String password = scanner.next();
		user.setPassword(password);

		//把user对象保存到集合中
		users.add(user);
		System.out.println("注册成功!");
		System.out.println("当前注册的人员："+users);
	}

}
```

### List接口

集合的体系：
 
------| Collection 单列集合 的根接口
---------|  `List`  如果是实现了List接口的集合类，该集合类具备的特点：有序，可重复。
---------|  `Set`   如果是实现了Set接口的集合类，该集合类具备的特点： 无序，不可重复。

有序： 集合的有序不是指自然顺序，而是指`添加进去的顺序`与元素`出来的顺序`是一致的。

List接口中特有方法：
	
	添加
		`add(int index, E element)` 
		`addAll(int index, Collection<? extends E> c)` 
	获取：
		`get(int index)`   索引从0开始
		`indexOf(Object o)` 
		`lastIndexOf(Object o)` 
		`subList(int fromIndex, int toIndex)` 
	修改：
		`set(int index, E element)` 

	迭代
		`listIterator()` 

List接口中特有的方法具备的特点： 操作的方法都存在`索引值`。

只有List接口下面的集合类才具备索引值。其他接口下面的集合类都没有索引值。

```java
@SuppressWarnings("unchecked")   //使用了未经检查或不安全的操作。jdk1.5以后再定义容器的时候要加上泛型，可以忽视
public class Demo {
	
	public static void main(String[] args) {

		List list=  new ArrayList();
		list.add("狗娃");
		list.add("狗剩");
		list.add("铁蛋");  //把元素添加到集合的末尾处。
		list.add("狗娃");
		
	    //添加方法
		list.add(1, "赵本山"); // 把元素添加到集合中的指定索引值位置上。

		List list2 = new ArrayList();
		list2.add("本山");
		list2.add("小沈阳");

		list.addAll(2,list2); //把list2的元素添加到list集合指定索引值的位置上。在原来位置上的往后移
		
		
		//获取的方法 
		System.out.println("get方法获取元素："+list.get(1)); //根据索引值获取集合中的元素

		//使用get方法遍历集合的元素：
		for (int i = 0; i < list.size() ; i++) {
			System.out.print(list.get(i)+",");
		}
		
		System.out.println("找出指定元素第一次出现在集合中的索引值："+ list.indexOf("本山"));
		System.out.println("找指定的元素最后一次出现在集合中的索引值："+list.lastIndexOf("狗娃"));

		List subList = list.subList(1, 3); //指定开始与结束的索引值截取集合中的元素。
		System.out.println("子集合的元素是："+ subList);
	 
		list.set(3, "赵本山"); //使用指定的元素替换指定索引值位置的元素。

		System.out.println("集合的元素："+list);
	}
}
```

#### ListIterator特有的方法：
	
	添加：
		`hasPrevious()`  判断是否存在上一个元素。
		`previous()`    当前指针先向上移动一个单位，然后再取出当前指针指向的元素。
		
		`next()`  先取出当前指针指向的元素，然后指针向下移动一个单位。

---------------------------	

		`add(E e)`   把当前有元素插入到当前指针指向的位置上。
		`set(E e)`   替换迭代器最后一次返回的元素。

 迭代器在变量元素的时候要注意事项： 
 	在迭代器迭代元素的过程中，不允许使用集合对象改变集合中的元素个数，如果需要`添加`或者`删除`只能使用`迭代器的方法`进行操作。
 
 如果使用过了集合对象改变集合中元素个数那么就会出现`ConcurrentModificationException`异常。
 
 迭代元素的过程中: 迭代器创建到使用结束的时间。

```java

import java.util.ArrayList;
import java.util.List;
import java.util.ListIterator;

@SuppressWarnings("unchecked")
public class Demo {

	public static void main(String[] args) {

		List list = new ArrayList();
		list.add("狗娃");
		list.add("狗剩");
		list.add("铁蛋");
		list.add("美美");
		
		ListIterator it = list.listIterator(); //返回的是一个List接口中特有的迭代器

		System.out.println("有上一个元素吗？"+ it.hasPrevious());  //false
		//System.out.println("获取上一个元素："+it.previous());      // java.util.NoSuchElementException

		it.next();
		System.out.println("获取上一个元素："+ it.previous());  //狗娃
		
		while(it.hasNext()){
			it.next();
		}
			
		while(it.hasPrevious()){
			System.out.println("元素："+ it.previous()); ////元素：美美  元素：铁蛋 元素：狗剩 元素：狗娃
		}
		
		it.next();
		it.next();
		it.add("张三");

		it.next();
		it.next();
		it.set("张三");
		
		System.out.println("集合的元素："+ list);  //集合的元素：[狗娃, 狗剩, 张三, 铁蛋, 张三]
	}
}
```

```java

import java.util.List;
import java.util.ArrayList;
import java.util.ListIterator;

/*
练习： 使用三种方式遍历集合的元素. 	
第一种： 使用get方法遍历。
第二种： 使用迭代器正序遍历。
第三种： 使用迭代器逆序遍历。
*/

@SuppressWarnings("unchecked")
public class Demo {
	
	public static void main(String[] args) {

		List list = new ArrayList();
		list.add("张三");
		list.add("李四");
		list.add("王五");
		
		System.out.println("======get方法遍历=======");
		for(int i = 0 ; i<list.size() ; i++){
			System.out.print(list.get(i)+",");
		}
		
		System.out.println("\r\n======使用迭代器正序遍历==========");
		ListIterator it = list.listIterator();	//获取到迭代器
		while(it.hasNext()){
			System.out.print(it.next()+",");
		}
		
		System.out.println("\r\n======使用迭代器逆序遍历==========");
		while(it.hasPrevious()){
			System.out.print(it.previous()+",");
		}
	}
}
```

```java
import java.util.ArrayList;
import java.util.List;
import java.util.ListIterator;

@SuppressWarnings("unchecked")
public class Demo {
	public static void main(String[] args) {

		List list = new ArrayList();
		list.add("张三");
		list.add("李四");
		list.add("王五");
		
		ListIterator it = list.listIterator();	//获取到迭代器
		while(it.hasNext()){
			System.out.print(it.next()+",");
			//it.add("aa"); // 把元素添加到当前指针指向位置
			//list.add("aa");  // add方法是把元素添加到集合的末尾处的。//java.util.ConcurrentModificationException
			//list.remove("张三");
		}
		
		//list.add("aa");  //java.util.ConcurrentModificationException

		it.next();

		System.out.println("\r\n集合的元素："+ list);
	}

}
```

#### ArrayList 特有的方法

集合的体系：
----| Collection  单列集合的根接口
--------| List 如果实现了List接口的集合类，具备的特点： 有序，可重复。
-----------| `ArrayList`   ArrayList 底层是维护了一个`Object`数组实现 的， 特点: `查询速度快，增删慢`。
-----------| `LinkedList`
-----------| `Vector` (了解即可)  底层也是维护了一个Object的数组实现的，实现与ArrayList是一样的，但是Vector是线程安全的，操作效率低。
--------| Set  如果实现了Set接口的集合类， 具备的特点： 无序，不可重复。

ArrayList 特有的方法：
	`ensureCapacity(int minCapacity)`
	`trimToSize()`  

什么时候使用ArrayList: 如果目前的数据是`查询比较多`，`增删比较少`的时候，那么就使用ArrayList存储这批数据。

笔试题目： 使用ArrayList无参的构造函数创建一个 对象时，默认的容量是多少? 如果长度不够使用时又自增增长多少？
	ArrayList底层是维护了一个`Object数组`实现 的，使用无参构造函数时，Object数组默认的`容量是10`，当长度不够时，`自动增长0.5倍`。

```java
import java.util.ArrayList;
import java.util.Iterator;

class Book{
	int id;
	String name;// 名
	public Book(int id, String name) {
		this.id = id;
		this.name = name;
	}

	@Override
	public String toString() {
		return "{ 书号："+ this.id+" 书名："+ this.name+" }";
	}

	@Override
	public boolean equals(Object obj) {
		Book book =(Book)obj;
		return this.id==book.id;
	}
}

@SuppressWarnings("unchecked")
public class Demo {
	public static void main(String[] args) {

		ArrayList list=  new ArrayList();
		list.add(new Book(110,"java编程思想"));
		list.add(new Book(220,"java核心技术"));
		list.add(new Book(330,"深入javaweb"));
		list.add(new Book(110,"java神书"));
		
		ArrayList list2 = clearRepeat(list);
		System.out.println("新集合的元素是："+ list2);
	}

	public static ArrayList  clearRepeat(ArrayList list){

		//创建一个新的集合
		ArrayList newList = new ArrayList();

		//获取迭代器
		Iterator it = list.iterator();

		while(it.hasNext()){
			Book book = (Book) it.next();  //从旧集合中获取的元素
			if(!newList.contains(book)){
				//如果新集合没有包含该书籍，那么就存储到新集合中
				newList.add(book);
			}
		}
		return newList;
	}
}
```

#### Linkedlist特有的方法

1：方法介绍
	addFirst(E e)
	addLast(E e)

	getFirst() 
	getLast() 
	
	removeFirst() 
	removeLast() 

2：数据结构
	1：栈 （1.6）   主要是用于实现堆栈数据结构的存储方式。
		先进后出
		`push()` 
		`pop()`

	2：队列（双端队列1.5）： 主要是为了让你们可以使用LinkedList模拟队列数据结构的存储方式。
		先进先出
		`offer()`
		`poll()`

3：返回逆序的迭代器对象
		descendingIterator()   返回逆序的迭代器对象


机试题目： 使用LinkedList实现堆栈数据结构的存储方式与队列的数据结构存储方式。
```java 
import java.util.LinkedList;

// 使用LinkedList模拟堆栈的数据结构存储方式
class StackList{
	
	LinkedList list;

	public StackList(){
		list = new LinkedList();
	}

	//进栈
	public void add(Object o){
		list.push(o);
	}

	//弹栈 : 把元素删除并返回。
	public Object pop(){
		return list.pop();
	}

	//获取元素个数
	public int size(){
		return list.size();
	}
}

//使用LinkedList模拟队列的存储方式
class TeamList{

	LinkedList list;

	public TeamList(){
		list = new LinkedList();
	}

	public void add(Object o){
		list.offer(o);
	}

	public Object remove(){
		return list.poll();
	}

	//获取元素个数
	public int size(){
		return list.size();
	}
}

@SuppressWarnings("unchecked")
public class Demo {
	public static void main(String[] args) {

		TeamList list=  new TeamList();
		list.add("李嘉诚");
		list.add("马云");
		list.add("王健林");
		
		int size = list.size();
		for(int i = 0 ; i<size ; i++){
			System.out.println(list.remove());
		}
	}
}
```

```java
import java.util.Iterator;
import java.util.LinkedList;

@SuppressWarnings("unchecked")
public class Demo {
	public static void main(String[] args) {

		LinkedList list= new LinkedList();
		list.add("张三");
		list.add("李四");
		list.add("王五");

		list.addFirst("狗娃"); //把元素添加到集合的首位置上。
		list.addLast("狗剩");  //把元素添加到集合的末尾处。

		System.out.println("获取集合中首位置的元素:"+list.getFirst());  //获取集合中首位置的元素:狗娃
		System.out.println("获取集合中末尾的元素："+ list.getLast());  //获取集合中末尾的元素：狗剩

		System.out.println("删除集合中的首位置元素并返回："+ list.removeFirst());  //删除集合中的首位置元素并返回：狗娃
		System.out.println("删除集合中的末尾素并返回："+ list.removeLast());  //删除集合中的末尾素并返回：狗剩
		
		list.push("狗娃");   //将该元素插入此集合的开头处。 
		System.out.println("删除集合的首元素："+list.pop()); // 移除并返回集合中的第一个元素  删除集合的首元素：狗娃

		
		list.offer("狗剩");
		System.out.println("删除集合的首元素: "+list.poll());  //删除集合的首元素: 张三

		System.out.println("集合中的元素："+ list);  //集合中的元素：[李四, 王五, 狗剩]

		Iterator  it = list.descendingIterator();
		while(it.hasNext()){
			System.out.println(it.next());    //狗剩 王五 李四
		}
	}
}
```

```java
import java.util.LinkedList;
import java.util.Random;

//需求： 使用LinkedList存储一副扑克牌，然后实现洗牌功能。

//扑克类
class Poker{
	String  color; //花色
	String num;	//点数
	public Poker(String color, String num) {
		super();
		this.color = color;
		this.num = num;
	}
	
	@Override
	public String toString() {
		return "{"+color+num+"}";
	}
}

@SuppressWarnings("unchecked")
public class Demo {
	public static void main(String[] args) {

		LinkedList pokers = createPoker();
		shufflePoker(pokers);
		showPoker(pokers);

	}

	//洗牌的功能
	public static void shufflePoker(LinkedList pokers){

		//创建随机数对象
		Random random = new Random();

		for(int i = 0 ; i <100; i++){ 

			//随机产生两个索引值
			int index1 = random.nextInt(pokers.size());
			int index2 = random.nextInt(pokers.size());

			//根据索引值取出两张牌，然后交换两张牌的顺序
			Poker poker1 = (Poker) pokers.get(index1);
			Poker poker2 = (Poker) pokers.get(index2);

			pokers.set(index1, poker2);
			pokers.set(index2, poker1);
		}
		
	}

	//显示扑克牌
	public static void showPoker(LinkedList pokers){

		for(int i = 0 ; i<pokers.size() ; i++){
			System.out.print(pokers.get(i));
			//换行
			if(i%10==9){
				System.out.println();
			}
		}
	}

	//生成扑克牌的方法
	public static LinkedList createPoker(){

		//该集合用于存储扑克对象。
		LinkedList list = new LinkedList();	

		//定义数组存储所有的花色与点数
		String[] colors = {"黑桃","红桃","梅花","方块"};
		String[] nums = {"A","2","3","4","5","6","7","8","9","10","J","Q","K"};

		for(int i = 0 ; i < nums.length ; i++){
			for(int j = 0 ; j<colors.length ; j++){
				list.add(new Poker(colors[j], nums[i]));
			}
		}
		return list;
	}
}
```

```java
import java.util.LinkedList;
/*
编写一个函数根据人的年龄及逆行排序存储。
*/
class Person{
	String name;
	int age;

	public Person(String name, int age) {
		super();
		this.name = name;
		this.age = age;
	}

	@Override 
	public String toString() {
		return "{ 名字："+ this.name+" 年龄："+ this.age+"}";
	}
}

@SuppressWarnings("unchecked")
public class Demo {
	public static void main(String[] args) {

		LinkedList list = new LinkedList();

		list.add(new Person("狗娃", 7));
		list.add(new Person("狗剩", 17));
		list.add(new Person("铁蛋", 3));
		list.add(new Person("美美", 30));
		
		for(int i= 0 ; i<list.size() -1 ; i++){
			for(int j = i+1 ; j<list.size() ; j++){

				//符合条件交换位置
				Person p1 = (Person) list.get(i);
				Person p2 = (Person) list.get(j);

				if(p1.age>p2.age){
					//交换位置
					list.set(i, p2);
					list.set(j, p1);
				}
			}
		}
		System.out.println(list);
	}
}
```

#### Vector

`Vector`  
	底层也是维护了一个Object的数组实现的，实现与ArrayList是一样的，但是Vector是线程安全的，操作效率低。

说出ArrayLsit与Vector的区别?

	相同点： ArrayList与Vector底层都是使用了`Object数组`实现的。
	
	不同点： 
		1. ArrayList是`线程不同步`的，操作效率高。 
		   Vector是`线程同步`的，操作效率低。
		2. ArrayList是JDK1.2出现，Vector是jdk1.0的时候出现的。

```java
import java.util.Enumeration;
import java.util.Vector;

@SuppressWarnings("unchecked")
public class Demo {
	public static void main(String[] args) {

		Vector v  =  new Vector();

		//添加元素
		v.addElement("张三");
		v.addElement("李四");
		v.addElement("王五");

		//迭代该集合
		//Enumeration接口中定义了一些方法，通过这些方法可以枚举（一次获得一个）对象集合中的元素。
		Enumeration e = v.elements(); //获取迭代器
		while(e.hasMoreElements()){ // 测试此枚举是否包含更多的元素。
			System.out.println(e.nextElement()); //如果此枚举对象至少还有一个可提供的元素，则返回此枚举的下一个元素。
		}
	}
}
```

### Set 接口

集合 的体系：
-----| Collection 单例集合的根接口
--------| List  如果是实现了List接口的集合类，具备的特点： 有序，可重复。 
------------| ArrayList  ArrayList 底层是维护了一个Object数组实现的。 特点： 查询速度快，增删慢。
------------| LinkedList LinkedList 底层是使用了链表数据结构实现的， 特点： 查询速度慢，增删快。
------------| Vector	 底层也是维护了一个Object的数组实现的，实现与ArrayList是一样的，但是Vector是线程安全的，操作效率低。

--------| Set  如果是实现了Set接口的集合类，具备的特点： 无序，不可重复。
------------| HashSet  底层是使用了`哈希表`来支持的，特点： `存取速度快`. 
------------| TreeSet  如果元素具备自然顺序 的特性，那么就按照元素自然顺序的特性进行排序存储。


`无序`： 添加元素 的顺序与元素出来的顺序是不一致 的。

```java
import java.util.Set;
import java.util.HashSet;

@SuppressWarnings("unchecked")
public class Demo {
	public static void main(String[] args) {

		Set set = new HashSet();

		set.add("王五");
		set.add("张三");
		set.add("李四");

		System.out.println("添加成功吗？"+set.add("李四"));  //HashSet
		System.out.println(set); 
	}
}
```

#### HashSet

--------| Set  如果是实现了Set接口的集合类，具备的特点： 无序，不可重复。
------------| HashSet  底层是使用了`哈希表`来支持的，特点： `存取速度快`.
------------| TreeSet 如果元素具备自然顺序 的特性，那么就按照元素自然顺序的特性进行排序存储。

hashSet的实现原理：
	往Hashset添加元素的时候，`HashSet`会先调用元素的`hashCode`方法得到元素的`哈希值`，
	然后通过元素的哈希值经过移位等运算，就可以算出该元素在哈希表中的存储位置。

	情况1： 如果算出元素存储的位置目前没有任何元素存储，那么该元素可以直接存储到该位置上。

	情况2： 如果算出该元素的存储位置目前已经存在有其他的元素了，那么会调用该元素的`equals方法`与该位置的元素再比较一次
	，如果equals返回的是`true`，那么该元素与这个位置上的元素就视为重复元素，不允许添加，如果equals方法返回的是`false`，那么该元素运行添加。


```java
import java.util.HashSet;
import javax.print.attribute.HashAttributeSet;

class Person{

	int id;

	String name;

	public Person(int id, String name) {
		super();
		this.id = id;
		this.name = name;
	}

	@Override
	public String toString() {
		return "{ 编号:"+ this.id+" 姓名："+ this.name+"}";
	}

	@Override
	public int hashCode() {
		System.out.println("=======hashCode=====");
		return this.id;
	}

	@Override
	public boolean equals(Object obj) {
		System.out.println("======equals======");
		Person p = (Person)obj;
		return this.id==p.id;
	}
}

@SuppressWarnings("unchecked")
public class Demo {
	public static void main(String[] args) {

		HashSet set = new HashSet();

		set.add("狗娃");
		set.add("狗剩");
		set.add("铁蛋");

		System.out.println("集合的元素："+ set);

		HashSet set2 = new HashSet();
		set2.add(new Person(110,"狗娃"));
		set2.add(new Person(220,"狗剩"));
		set2.add(new Person(330,"铁蛋"));

		//在现实生活中只要编号一致就为同一个人.
		System.out.println("添加成功吗？"+set2.add(new Person(110,"狗娃")));   //添加成功吗？false
		System.out.println("集合的元素："+set2); //集合的元素：[{ 编号:330 姓名：铁蛋}, { 编号:220 姓名：狗剩}, { 编号:110 姓名：狗娃}]
	}
}
```

```java
import java.util.HashSet;
import java.util.Scanner;

/*
 需求： 接受键盘录入用户名与密码，如果用户名与密码已经存在集合中，那么就是视为重复元素，不允许添加到HashSet中。
 
 */
class User{

	String userName;

	String password;

	public User(String userName, String password) {
		super();
		this.userName = userName;
		this.password = password;
	}

	@Override
	public String toString() {
		return "{ 用户名："+this.userName+" 密码："+ this.password+"}";
	}

	@Override
	public boolean equals(Object obj) {
		User user = (User)obj;
		return this.userName.equals(user.userName)&&this.password.equals(user.password);
	}

	@Override
	public int hashCode() { //  abc 123   ， 123 abc
		return userName.hashCode()+password.hashCode();
	}
}

@SuppressWarnings("unchecked")
public class Demo {
	public static void main(String[] args) {

		Scanner scanner = new Scanner(System.in);

		HashSet set = new HashSet();
		
		while(true){

			System.out.println("请输入用户名:");
			String userName = scanner.next();

			System.out.println("请输入密码：");
			String password = scanner.next();

			//创建一个对象
			User user = new User(userName, password);

			if(set.add(user)){

				System.out.println("注册成功...");
				System.out.println("当前的用户有："+ set);

			}else{
				System.out.println("注册失败...");
			}
		}
	}
}
```

```java
@SuppressWarnings("unchecked")
public class Demo {
	public static void main(String[] args) {

		String str1 = "hello";
		String str2 = new String("hello");

		System.out.println("两个是同一个对象吗？"+(str1==str2));  //两个是同一个对象吗？false
		System.out.println("相等吗？"+str1.equals(str2));    //true

		System.out.println("str1的hashCode："+ str1.hashCode());  //str1的hashCode：99162322
		System.out.println("str2的hashCode:"+ str2.hashCode());  //str2的hashCode:99162322

		//HashCode默认情况下表示的是内存地址，String 类已经重写了Object的hashCode方法了。
		//注意： 如果两个字符串的内容一致，那么返回的hashCode 码肯定也会一致的。
	}
}
```

#### TreeSet

--------| Set  如果是实现了Set接口的集合类，具备的特点： 无序，不可重复。
------------| HashSet  底层是使用了`哈希表`来支持的，特点： `存取速度快`. 
------------| TreeSet  如果元素具备`自然顺序` 的特性，那么就`按照元素自然顺序的特性进行排序存储`。
			  TreeSet 底层是使用了红黑树（二叉树）数据结构实现的， 特点：会对元素进行排序存储。

```java

import java.util.TreeSet;
@SuppressWarnings("unchecked")
public class Demo {
	
	public static void main(String[] args) {

		TreeSet tree1 = new TreeSet();
		tree1.add(1);
		tree1.add(10);
		tree1.add(7);
		tree1.add(19);
		tree1.add(9);
		System.out.println("tree1: " + tree1); //tree1: [1, 7, 9, 10, 19]

		TreeSet tree2 = new TreeSet();
		tree2.add('b');
		tree2.add('f');
		tree2.add('a');
		tree2.add('c');
		System.out.println("tree2: " + tree2);	//tree2: [a, b, c, f]
	}
}
```

##### treeSet添加自定义元素-比较器
 
 treeSet要注意的事项：
 	1. 往TreeSet添加元素的时候，如果元素本身具备了自然顺序的特性，那么就按照元素自然顺序的特性进行排序存储。
   
 	2. 往TreeSet添加元素的时候，如果元素本身不具备自然顺序的特性，那么该元素所属的类`必须要实现Comparable接口`，把元素
 	的比较规则定义在`compareTo(T o)`方法上。 
 	
 	3. 如果比较元素的时候，`compareTo`方法返回的是`0`，那么该元素就被视为`重复元素`，`不允许添加`.
   		(注意：`TreeSet与HashCode、equals方法是没有任何关系`。)
 	
 	4. 往TreeSet添加元素的时候, 如果元素本身没有具备自然顺序的特性，而元素所属的类也没有实现Comparable接口，那么必须要在创建TreeSet的时候传入一个比较器。
 	
 	5.  往TreeSet添加元素的时候，如果元素本身不具备自然顺序的特性，而元素所属的类已经实现了Comparable接口， 在创建TreeSet对象的时候也传入了比较器那么是以`比较器`的比较规则优先使用。
   
 	
如何自定义定义比较器： 自定义一个类实现Comparator接口即可，把元素与元素之间的比较规则定义在compare方法内即可。
 		
 		自定义比较器的格式 ：
 			
 			class  类名  implements Comparator{
 			
 			}
 	
推荐使用：使用`比较器(Comparator)`。 

```java
import java.util.Comparator;
import java.util.TreeSet;

class  Emp implements Comparable<Emp>{

	int id;

	String name;

	int salary;

	public Emp(int id, String name, int salary) {
		super();
		this.id = id;
		this.name = name;
		this.salary = salary;
	}

	@Override
	public String toString() {
		return "{ 编号："+  this.id+" 姓名："+ this.name+" 薪水："+ this.salary+"}";
	}

	//@Override //元素与元素之间的比较规则。
	// 负整数、零 或 正整数，根据此对象是小于、等于还是大于指定对象。 
	public int compareTo(Emp o) {
		return this.salary- o.salary;
	}
}


//自定义一个比较器
class MyComparator implements Comparator<Emp>{

	//根据第一个参数小于、等于或大于第二个参数分别返回负整数、零或正整数。 
	/*@Override
	public int compare(Object o1, Object o2) {
		Emp e1 = (Emp) o1;
		Emp e2 = (Emp) o2;
		return e1.id - e2.id;
	}*/

	@Override
	public int compare(Emp o1, Emp o2) {
		return o1.id-o2.id;
	}
}

@SuppressWarnings("unchecked")
public class Demo {
	public static void main(String[] args) {

		//创建一个比较器对象
		MyComparator comparator = new MyComparator();

		//创建TreeSet的时候传入比较器
		TreeSet tree = new TreeSet(comparator);
		
		tree.add(new Emp(110, "老陆", 100));
		tree.add(new Emp(113, "老钟", 200));
		tree.add(new Emp(220, "老汤", 300));
		tree.add(new Emp(120, "老蔡", 500));

		System.out.println("集合的元素："+tree);
		//集合的元素：[{ 编号：110 姓名：老陆 薪水：100}, { 编号：113 姓名：老钟 薪水：200}, { 编号：120 姓名：老蔡 薪水：500}, { 编号：220 姓名：老汤 薪水：300}]
	}
}
```

```java
import java.util.TreeSet;
/*
TreeSet是可以对字符串进行排序的， 因为字符串已经实现了Comparable接口。

字符串的比较规则：
	情况一： 对应位置有不同的字符出现， 就比较的就是对应位置不同的字符。
	情况 二：对应位置上 的字符都一样，比较的就是字符串的长度。
 */

@SuppressWarnings("unchecked")
public class Demo {
	
	public static void main(String[] args) {

		TreeSet tree = new TreeSet();
		tree.add("abcccccccccccccccccc");
		tree.add("abc");
		System.out.println(tree);   //[abc, abcccccccccccccccccc]
		
		System.out.println("a".compareTo("c"));  //-2
		System.out.println("a".compareTo("f"));  //-5
		System.out.println("abw".compareTo("abcccccccccccc"));  //20
	}
}
```

```java

import java.util.Iterator;
import java.util.TreeSet;

/*
需求：将字符串中的数值进行排序。
	例如 String str="8 10 15 5 2 7"; ---->   "2 5 7 8 10 15"
*/

@SuppressWarnings("unchecked")
public class Demo {
	public static void main(String[] args) {

		String str="8 10 15 5 2 7";
		String[] datas = str.split(" ");
		
		TreeSet tree = new TreeSet();
		for(int i = 0 ; i<datas.length ; i++){
			tree.add(Integer.parseInt(datas[i])); // 字符串转int类型数据是需要使用Integer.parseInt()
		}
		
		//遍历treeSet的元素拼接成对应的字符串
		Iterator it = tree.iterator();
		while(it.hasNext()){
			System.out.print(it.next()+" ");    //2 5 7 8 10 15
		}
	}
}
```
##  泛型 genrictiry

泛型是jdk1.5使用的新特性。
 
**泛型的好处：**
   1. `将运行时的异常提前至了编译时`。
   2. `避免了无谓的强制类型转换` 。

**泛型在集合中的常见应用：**
  	
  	`ArrayList<String>  list = new ArrayList<String>();`  true     推荐使用。
  	
  	ArrayList<Object>  list = new ArrayList<String>();  false
  	ArrayList<String>  list = new ArrayList<Object>();  false
  	
  	//以下两种写法主要是为了兼顾新老系统的兼用性问题。
  
    ArrayList<String> list = new ArrayList();           true   
  
    ArrayList list = new ArrayList<String>();   		 true   
 
注意： `泛型没有多态的概念`，`左右两边的数据 类型必须要一致`，或者只是`写一边的泛型类型`。

推荐使用： `两边都写泛型`。



```java
import java.util.ArrayList;

//需求： 把一个集合中元素全部转成大写。 
public class Demo {
	public static void main(String[] args) {

		ArrayList<String>  list = new ArrayList<String>();  //<String> 表示该容器只能存储字符串类型 的数据。
		
		list.add("aa");
		list.add("bb");
		list.add("cc");
		
		for(int i = 0 ; i < list.size() ; i++){
			String str =  list.get(i);
			System.out.println("大写："+ str.toUpperCase());
		}
	}
}
```

### 自定义泛型

自定义泛型：  自定义泛型就是一个`数据类型的占位符`或者是一个`数据类型的变量`。


在泛型中`不能使用基本数据类型`，如果需要使用基本数据类型，那么就使用基本数据类型对应的`包装类型`。
 byte      Byte
 short     Short
 int       Integer
 long      Long
  
 double    Double
 float     Float

 boolean   Boolean

 char-------> Character

 #### 泛型方法
 
** 方法泛型注意的事项：**
 	1. 在`方法上`自定义泛型，这个自定义泛型的`具体数据类型`是在`调用该方法`的时候`传入实参时` `确定具体的数据类型`的。
 	2. 自定义泛型只要符合标识符的命名规则即可, 但是自定义泛型我们一般都习惯使用一个大写字母表示。  T Type  E Element

**泛型方法的定义格式：**

	修饰符  <声明自定义的泛型>返回值类型   函数名(使用自定义泛型 ...){
		
	}

```java
//需求： 定义一个方法可以接收任意类型的参数，而且返回值类型必须 要与实参的类型一致。
public class Demo {
	public static void main(String[] args) {

		String str = getData("abc");
		Integer i = getData(123);
	}

	public static <abc>abc getData(abc o){
		
		return o;
	}
}
```

#### 泛型类

泛型类的定义格式：

	class 类名<声明自定义泛型>{
	
	}
 
泛型类要注意的事项：
 	1. 在类上自定义泛型的具体数据类型是在使用该类的时候`创建对象时候确定`的。
 	2. 如果一个类在类上已经声明了自定义泛型，如果使用该类创建对象的时候没有指定 泛型的具体数据类型，那么`默认为Object类型`
 	3.`在类上自定义泛型不能作用于静态的方法`，如果静态的方法需要使用自定义泛型，那么需要在方法上自己声明使用。

```java
// 需求： 编写一个数组的工具类
class MyArrays<T>{

	//元素翻转
	public void reverse(T[] arr){
		for(int startIndex = 0, endIndex = arr.length-1 ; startIndex<endIndex ; startIndex++,endIndex--){
			T temp  = arr[startIndex];
			arr[startIndex] = arr[endIndex];
			arr[endIndex] = temp;
		}
	}

	public String toString(T[] arr){
		StringBuilder sb = new StringBuilder();
		for(int i = 0 ; i < arr.length ; i++){
			if(i==0){
				sb.append("["+arr[i]+",");
			}else if(i==arr.length-1){
				sb.append(arr[i]+"]");
			}else{
				sb.append(arr[i]+",");
			}
		}
		return sb.toString();
	}

	public static <T>void print(T[] t){
		
	}
}

public class Demo {
	public static void main(String[] args) {

		Integer[] arr = {10,12,14,19};
		
		MyArrays<Integer> tool = new MyArrays<Integer>();
		tool.reverse(arr);
		System.out.println("数组的元素："+tool.toString(arr));
		

		MyArrays<String> tool2 = new MyArrays<String>();
		String[] arr2 = {"aaa","bbb","ccc"};
		tool2.reverse(arr2);
	}
}
```

#### 泛型接口
 
泛型接口的定义格式:

	interface 接口名<声明自定义泛型>{
	
	}


泛型接口要注意的事项：
	1. 接口上自定义的`泛型的具体数据类型是在实现一个接口的时候指定的`。
	2. 在接口上自定义的泛型如果在`实现接口的时候没有指定具体的数据类型`，那么默认为`Object类型`。
	3. 如果需要在创建接口实现类对象的时候才指定接口上自定义泛型，那么需要以下格式：	class<T> 类名  implements  接口<T>

需求：目前我实现一个接口的时候，我还不明确我目前要操作的数据类型，我要等待创建接口实现类 对象的时候我才能指定泛型的具体数据类型。
	
如果要延长接口自定义泛型 的具体数据类型，那么格式如下：
	格式：  
		public class Demo<T> implements Dao<T>{
		
		}
```java
interface Dao<T>{
	public void add(T t);
}

public class Demo<T> implements Dao<T> {

	public static void main(String[] args) {
		Demo<String> d = new Demo<String>();
	}

	public void add(T t){
		
	}
}
```

### 泛型的上下限

泛型中通配符： `？`
	
	? super Integer : 只能存储Integer或者是Integer父类元素。      `泛型的下限`
 	
 	? extends Number ： 只能存储Number或者是Number类型的子类数据。 `泛型的上限`

 Number 是一个抽象类,也是一个超类（即父类）。
	Number 类属于 `java.lang` 包
	所有的包装类（如 `Double`、`Float`、`Byte`、`Short`、`Integer` 以及` Long`）都是`抽象类 Number` 的子类。
	
```java
import java.util.ArrayList;
import java.util.Collection;
import java.util.HashSet;

public class Demo {
	public static void main(String[] args) {

		ArrayList<Integer> list1 = new ArrayList<Integer>();
		ArrayList<Number> list2 = new ArrayList<Number>();
		
		HashSet<String> set = new HashSet<String>();
		//getData(set);
		
	}

	//泛型的上限
	public static void getData(Collection<? extends Number> c){
		
		
	}

	//泛型的下限
	public static void print(Collection<? super Integer> c){
		
	}
}
```

##  Map 双列集合

------| Map  如果是实现了Map接口的集合类，具备的特点： 存储的数据都是以键值对的形式存在的，键不可重复，值可以重复。
---------| HashMap
---------| TreeMap
---------| Hashtable
 
 Map接口的方法：
 	
 	添加：
 		`put(K key, V value)` 
 		`putAll(Map<? extends K,? extends V> m)` 
 	
 	删除
 		`remove(Object key)` 
 		`clear()` 

 	获取：
 		`get(Object key)` 
 		`size()` 
 	
 	判断：
 		`containsKey(Object key)` 
 		`containsValue(Object value)` 
 		`isEmpty()` 

```java

import java.util.HashMap;
import java.util.Map;

@SuppressWarnings("unchecked")
public class Demo {

	public static void main(String[] args) {

		Map<String,String> map = new HashMap<String, String>();

		//添加方法
		map.put("汪峰", "章子怡");
		map.put("文章", "马伊琍");
		map.put("谢霆锋","张柏芝");

		//添加
		//System.out.println("返回值："+map.put("ixfosa","long")); //没有存在该键，返回的是null 返回值：null
		System.out.println("返回值："+map.put("谢霆锋","黄菲"));  //已经存在该键了，返回该键之前对应的值。并替换原来的 返回值：张柏芝

		Map<String,String> map2 = new HashMap<String, String>();
		map2.put("杨振宁", "翁帆");
		map2.put("习总", "彭丽媛");
		map.putAll(map2); //把map2的元素添加到map集合中。
		
		System.out.println("删除的数据是:"+map.remove("汪峰")) ;  //根据键删除一条map中的数据，返回的是该键对应的值。 删除的数据是:章子怡

		//获取
		System.out.println("根据指定 的键获取对应的值:"+ map.get("文章"));  //根据指定 的键获取对应的值:马伊琍
		System.out.println("获取map集合键值对个数："+map.size());  /获取map集合键值对个数：4/
		
		//判断
		System.out.println("判断map集合是否包含指定的键："+ map.containsKey("文章"));  //判断map集合是否包含指定的键：true
		
		System.out.println("判断map集合中是否包含指定 的值："+ map.containsValue("张柏芝"));  //判断map集合中是否包含指定的值：false

		map.clear();

		System.out.println("判断map集合是否为空元素："+ map.isEmpty()); //判断map集合是否为空元素：true

		System.out.println("集合的元素："+ map); //集合的元素：{}
	}
}
```

### 迭代：

	keySet() 
	values() 
	entrySet()

```java
import java.util.Collection;
import java.util.HashMap;
import java.util.Iterator;
import java.util.Map;
import java.util.Map.Entry;
import java.util.Set;

@SuppressWarnings("unchecked")
public class Demo {

	public static void main(String[] args) {

		Map<String,String> map = new HashMap<String, String>();

		//添加方法
		map.put("汪峰", "章子怡");
		map.put("文章", "马伊琍");
		map.put("谢霆锋","张柏芝");
		map.put("成龙", "林凤娇");
		
		//map集合中遍历方式一： 使用keySet方法进行遍历       缺点： keySet方法只是返回了所有的键，没有值。 
		Set<String> keys = map.keySet();  //keySet() 把Map集合中的所有键都保存到一个Set类型 的集合对象中返回。
		Iterator<String> it1 = keys.iterator();
		while(it1.hasNext()){
			String key = it1.next();
			System.out.println("键："+ key+" 值："+ map.get(key));
		}
		
		
		//map集合的遍历方式二： 使用values方法进行 遍历。    缺点： values方法只能返回所有 的值，没有键。
		Collection<String>  c = map.values(); //values() 把所有的值存储到一个Collection集合中返回。
		Iterator<String> it2 = c.iterator();
		while(it2.hasNext()){
			System.out.println("值："+ it2.next());
		}
		
		//map集合的遍历方式三： entrySet方法遍历。
		Set<Map.Entry<String,String>>  entrys = map.entrySet();
		Iterator<Map.Entry<String,String>> it3 = entrys.iterator();
		while(it3.hasNext()){
			Map.Entry<String,String> entry = it3.next();
			System.out.println("键："+ entry.getKey()+" 值："+ entry.getValue());
		}

		/*class Map{
			//静态内部类 
			static class Entry<K ,V>{
				K  key;
				V value;
			} 
		}*/


		System.out.println("\r\n======entrySet方式遍历=========");
		HashMap<String, String> map = new HashMap<String, String>();
		map.put("张三","001");
		map.put("李四","002");
		map.put("王五","003");
		Set<Entry<String,String>> entrys = map.entrySet(); //
		for (Entry<String,String> entry : entrys) {
			System.out.println("键："+entry.getKey()+" 值："+entry.getValue());
		}
	}
}
```

```java
import java.util.Scanner;
import java.util.TreeMap;

/*
从键盘输入一个字母组成字符串，分别统计每个字母出现的次数
要求输出的效果按照字母的顺序输出  a(7)b(5)...
*/
public class Demo {
	public static void main(String[] args) {

		System.out.println("请输入一段字符串：");

		Scanner  scanner = new Scanner(System.in);
		String line = scanner.next();

		char[] arr = line.toCharArray();//先把字符串转换成字符数组。
		TreeMap<Character, Integer> map = new TreeMap<Character, Integer>(); 

		for(char c : arr){
			if(map.containsKey(c)){ //map集合已经包含了该字符
				int count = map.get(c);
				map.put(c, count+1);
			}else{  //没有包含
				map.put(c, 1);
			}
		}
		System.out.println(map);
	}
}
```
### HashMap

双列集合：
------| Map  如果是实现了Map接口的集合类，具备的特点： 存储的数据都是以键值对的形式存在的，键不可重复，值可以重复。
--------| HashMap  底层也是基于哈希表实现的。
--------| TreeMap
--------| Hashtable 

HashMap的存储原理：
	往HashMap添加元素的时候，首先会调用键的hashCode方法得到元素 的哈希码值，然后经过运算就可以算出该元素在哈希表中的存储位置。 

	情况1： 如果算出的位置目前没有任何元素存储，那么该元素可以直接添加到哈希表中。
	
	情况2：如果算出 的位置目前已经存在其他的元素，那么还会调用该元素的equals方法与这个位置上的元素进行比较，如果equals方法返回 的是false，那么该元素允许被存储，如果equals方法返回的是true，那么该元素被视为重复元素，不允存储。

```java

import java.util.HashMap;

class Person{

	int id;

	String name;

	public Person(int id, String name) {
		super();
		this.id = id;
		this.name = name;
	}

	@Override
	public String toString() {
		return  "[编号："+this.id+" 姓名："+ this.name+"]";
		
	}

	@Override
	public int hashCode() {
		return this.id;
	}

	@Override
	public boolean equals(Object obj) {
		Person p = (Person) obj;
		return this.id== p.id;
	}
}

public class Demo {
	public static void main(String[] args) {

		HashMap<Person, String> map = new HashMap<Person, String>();

		map.put(new Person(110,"狗娃"), "001");
		map.put(new Person(220,"狗剩"), "002");
		map.put(new Person(330,"铁蛋"), "003");
		map.put(new Person(110,"狗娃"), "007");  //如果出现了相同键，那么后添加的数据的值会取代之前 的值。

		System.out.println("集合的元素："+ map);
	}
	
}
```
### TreeMap

双列集合： 
-------------| Map  如果是实现了Map接口的集合类，具备的特点： 存储的数据都是以键值对的形式存在的，键不可重复，值可以重复。
----------------| HashMap  底层也是基于哈希表实现 的。
----------------| `TreeMap`   TreeMap也是基于红黑树（二叉树）数据结构实现 的， 特点：会对元素的键进行排序存储
----------------| Hashtable 

TreeMap 要注意的事项：
   1. 往TreeMap添加元素的时候，如果元素的键具备自然顺序，那么就会按照`键的自然顺序特性进行排序存储`。
   2. 往TreeMap添加元素的时候，如果元素的键不具备自然顺序特性， 那么键所属的类必须要实现`Comparable`接口，把键的比较规则定义在CompareTo方法上。
   3. 往TreeMap添加元素的时候，如果元素的键不具备自然顺序特性，而且键所属的类也没有实现`Comparable`接口，那么就必须在创建TreeMap对象的时候传入比较器。

```java

import java.util.Comparator;
import java.util.TreeMap;

class Emp {    //implements Comparable<Emp>{

	String name;

	int salary;

	public Emp(String name, int salary) {
		super();
		this.name = name;
		this.salary = salary;
	}

	@Override
	public String toString() {
		return "[姓名："+this.name+" 薪水："+ this.salary+"]";
	}

    /*
	@Override
	public int compareTo(Emp o) {
		return this.salary - o.salary;
	}*/
}

//自定义一个比较器
class MyComparator implements Comparator<Emp>{

	@Override
	public int compare(Emp o1, Emp o2) {
		return o1.salary - o2.salary;
	}
}

public class Demo {
	
	public static void main(String[] args) {

		TreeMap<Character, Integer> tree = new TreeMap<Character, Integer>();

		tree.put('c',10);
		tree.put('b',2);
		tree.put('a',5);
		tree.put('h',12);
		System.out.println(tree);
		
		//创建一个自定义比较器
		MyComparator comparator = new MyComparator();
		
		TreeMap<Emp, String> tree2 = new TreeMap<Emp, String>(comparator);
		tree2.put(new Emp("冰冰", 2000),"001");
		tree2.put(new Emp("家宝", 1000),"002");
		tree2.put(new Emp("习总", 3000),"003");
		tree2.put(new Emp("克强", 5000),"005");
		
		tree2.put(new Emp("财厚", 5000),"008");
		System.out.println(tree2);
	}
}
```

 ## jdk1.5新特性
 ### 静态导入

静态导入的作用： 简化书写, 使用import可以省略写包名；而使用`import static`，则连`类名都可以省略`。

静态导入可以作用一个类的所有静态成员。

 静态导入的格式：
 	`import static 包名.类名.静态的成员`；

静态导入要注意的事项：
 	1. 如果静态导入的成员与本类的成员存在同名的情况下，那么默认使用本类的静态成员，如果需要指定使用静态导入的成员，那么需要在静态成员前面加上类名

```java
import java.util.ArrayList;
import java.util.Collections;

import static java.util.Collections.sort;
import static java.util.Collections.binarySearch;
import static java.util.Collections.max;

import static java.lang.System.out;

public class Demo {

	public static void main(String[] args) {

		ArrayList<Integer> list = new ArrayList<Integer>();
		list.add(13);
		list.add(9);
		list.add(10);
		list.add(19);
		
		//排序
		Collections.sort(list);
		out.println("集合的元素："+ list);
		out.println("索引值："+ binarySearch(list,13));
		out.println("最大值："+ max(list));
	}

	public static void sort(ArrayList<Integer> list){
		System.out.println("本类 的sort方法.....");
	}
```

### 增强for循环

增强for循环的作用： 简化迭代器的书写格式。(注意：增强for循环的底层还是使用了`迭代器遍历`。)

增强for循环的适用范围： 如果是实现了`Iterable接口`的对象或者是`数组对象`都可以使用增强for循环。

增强for循环的格式：
 	
 	for(数据类型  变量名  :遍历的目标){
 	
 	}

增强for循环要注意的事项：
	1. 增强for循环底层也是使用了迭代器获取的，只不过获取迭代器由jvm完成，不需要我们获取迭代器而已，所以在使用增强for循环变量元素的过程中不准使用集合对象对集合的元素个数进行修改。
	2. 迭代器遍历元素与增强for循环变量元素的区别：使用迭代器遍历集合的元素时可以删除集合的元素，而增强for循环变量集合的元素时，不能调用迭代器的remove方法删除元素。
	3. 普通for循环与增强for循环的区别：普通for循环可以没有变量的目标，而增强for循环一定要有变量的目标。

```java

import java.util.HashMap;
import java.util.HashSet;
import java.util.Iterator;
import java.util.Map;
import java.util.Set;

public class Demo {
	public static void main(String[] args) {

		HashSet<String> set = new HashSet<String>();
		//添加元素
		set.add("狗娃");
		set.add("狗剩");
		set.add("铁蛋");
		
		
		//使用迭代器遍历Set的集合.
		Iterator<String> it  = set.iterator();
		while(it.hasNext()){
			String temp = it.next();
			System.out.println("元素："+ temp);
			//it.remove();   //迭代器遍历集合的元素时可以删除集合的元素
		}
		
		//使用增强for循环解决
		for(String item : set){
			System.out.println("元素："+ item);  //不能调用迭代器的remove方法删除元素。
		}
		

		int[] arr = {12,5,6,1};
		
	 	//普通for循环的遍历方式
	 	for(int i =  0 ; i<arr.length ; i++){
			System.out.println("元素："+ arr[i]);
		}
		
		//使用增强for循环实现
		for(int item : arr){
			System.out.println("元素："+ item);
		}
		
		
		//需求： 在控制台打印5句hello world.
		for(int i = 0 ; i < 5; i++){
			System.out.println("hello world");
		}
		
		
		//注意： Map集合没有实现Iterable接口，所以map集合不能直接使用增强for循环，如果需要使用增强for循环需要借助于Collection
		// 的集合。
		HashMap<String, String> map = new HashMap<String, String>();
		map.put("001","张三");
		map.put("002","李四");
		map.put("003","王五");
		map.put("004","赵六");

		Set<Map.Entry<String, String>> entrys = map.entrySet();

		for(Map.Entry<String, String> entry  :entrys){
			System.out.println("键："+ entry.getKey()+" 值："+ entry.getValue());
		}
	}
}
```

```java

import java.util.Iterator;

//自定一个类使用增强for循环
class MyList implements Iterable<String>{

	Object[] arr = new Object[10];

	int index = 0 ;	//当前的指针

	public void add(Object o){
		arr[index++] = o;  // 1
	}

	public int size(){
		return index;
	}

	@Override
	public Iterator<String> iterator() {

		return new Iterator<String>() {

			int cursor  = 0;

			@Override
			public boolean hasNext() {
				return cursor<index;
			}

			@Override
			public String next() {
				return (String) arr[cursor++];
			}

			@Override
			public void remove() {

			}
		};
	}
}

public class Demo {

	public static void main(String[] args) {
		MyList list = new MyList();
		list.add("张三");
		list.add("李四");
		list.add("王五");
		
		for(String item : list){
			System.out.println(item);
		}
	}
}
```

### 可变参数

可变参数的格式：

	`数据类型... 变量名`

可变参数要 注意的细节：
	1. 如果一个函数 的形参使用上了可变参数之后，那么调用该方法的时候`可以传递参数也可以不传递参数`。
	2. 可变参数实际上是一个`数组对象`。
	3. 可变参数必须位于形参中的`最后一个参数`。
	4. 一个函数最多只能有一个可变参数，因为可变参数要位于形参中`最后一个位置`上。

```java
//需求： 定义一个函数做加法功能（函数做几个数据 的加法功能是不确定）。
public class Demo {
	public static void main(String[] args) {

		int[] arr = {1,2,45,6,7};

		add(arr);   //总和：61

		add(1, 1, 1);  //总和：3
	}

	public static void add(int... arr){ //长度是0
		
		int result = 0;
		for(int item : arr){
			result+=item;
		}
		System.out.println("总和："+ result);
	}
}
```

### 自动装箱与自动拆箱。

 java是面向对象 的语言，任何事物都可以使用类进行描述，sun就使用了一些类描述java中八种基本数据类型数据

`自动装箱`： 自动把java的基本数据类型数据转换成对象类型数据。
`自动拆箱`： 把引用类型的数据转换成基本类型的数据
 
 	基本数据类型            包装类型
 	byte     	  		   Byte
	short      	  		   Short
	int          		   Integer
	long         		   Long 
	
	float         		   Float
	double         		   Double 

	boolean       	       Boolean 
	
	char          		   Character


基本数据类型数据有了对应 的包装 类型的好处：
	Integer.parseInt(str)   字符串转换成int类型数据
	Integer.toString(i)     数字转换成字符串
	toBinaryString(10)      整数转换成对应的进制形式

Integer类细节 :
	Integer类内部维护了缓冲数组，该缓冲数组存储的-128~127 这些数据在一个数组中。如果你获取的数据是落入到这个范围之内的，那么就直接从该缓冲区中获取对应的数据。

```java
public class Demo {
	public static void main(String[] args) {

		String str = "12";
		
		//字符串转换成int类型数据。 可以把字符串转换成对应的数字
		int i = Integer.parseInt(str);
		System.out.println(i+1);
		
		//把数字转换成字符串
		System.out.println("把整数转换成对应 的字符串："+Integer.toString(i));
		
		//把整数转换成对应的进制形式
		System.out.println("10的二进制："+ Integer.toBinaryString(10));
		System.out.println("10的二进制："+ Integer.toBinaryString(10));
		System.out.println("10的十六进制："+ Integer.toHexString(10));
		
		//可以把字符串当成对应的进行数据帮你转换
		String data = "10";
		int a = Integer.parseInt(data, 2);
		System.out.println("a="+a);
		
		//集合：集合是可以存储任意对象类型数据的容器。
		ArrayList list = new ArrayList();
		list.add(1);
		list.add(2);
		list.add(3);
		
		//自动装箱： 自动把java的基本数据类型数据转换成对象类型数据。
		int temp = 10;  //基本数据类型
		Integer b = temp; //把a存储的值赋予给b变量。
		
		//自动拆箱： 把引用类型的数据转换成基本类型的数据
		Integer c = new Integer(13);
		int d = c; //
		System.out.println(d);
		
		//引用的数据类型
		Integer e = 128;
		Integer f = 128; 
		System.out.println("同一个对象吗？"+(e==f)); // Integer类内部维护 了缓冲数组，该缓冲数组存储的-128~127 这些数据在一个数组中。如果你获取的数据是落入到这个范围之内的，那么就直接从该缓冲区中获取对应的数据。
	}
}
```

### 枚举

问题：某些方法所接收的数据必须是在固定范围之内的.
一些方法在运行时，它需要的数据不能是任意的，而必须是一定范围内的值，可以直接使用枚举予以解决。

解决方案： 这时候我们的解决方案就是自定义一个类,然后是私有化构造函数，在自定义类中创建本类的对象对外使用。

jdk1.5对以上问题提出了新的解决方案： 就是使用枚举类解决。

一些方法在运行时，它需要的数据不能是任意的，而必须是一定范围内的值，Java5以后可以直接使用枚举予以解决。
 	比如： 方向 , 性别 、 季节 、 星期......

**枚举类的定义格式：**
	enum 类名{
		//枚举值
	}
	
**枚举要注意的细节：**
	1. 枚举类也是一个特殊的类。
	2. 枚举值默认的修饰符是`public static final`。
	3. 枚举值就是是`枚举值所属的类的类型`， 然后枚举值是`指向了本类的对象`的。
	4. 枚举类的构造方法默认的修饰符是`private`的。
	5. 枚举类可以定义自己的成员变量与成员函数。
	6. 枚举类可以自定义构造函数，但是构造函数的修饰符必须是`private`。
	7. 枚举类可以存在抽象的方法，但是`枚举值必须要实现抽象`的方法。
	8. 枚举值必须要位置枚举类的`第一个语句`。
	9. 可以使用 `for` 语句来迭代枚举元素
	10. switch适用的数据类型： byte \ char \short \ int \ String\ 枚举类型
 		注意： case语句后面跟的枚举值，只需要单写枚举值即可，不需要再声明该枚举值是属于哪个枚举类的。

**values(), ordinal() 和 valueOf() 方法**
enum 定义的枚举类默认继承了 java.lang.Enum 类，并实现了 java.lang.Seriablizable 和 java.lang.Comparable 两个接口。

values(), ordinal() 和 valueOf() 方法位于 java.lang.Enum 类中：

`values()` 返回枚举类中所有的值。
`ordinal()`方法可以找到每个枚举常量的索引，就像数组索引一样。
`valueOf()`方法返回指定字符串值的枚举常量。


 ```java
//自定义一个性别类
class Gender{

	String value;

	public static final Gender man = new Gender("男");

	public static final Gender woman = new Gender("女");


	private Gender(String value) {
		this.value = value;
	}
}

public class Demo {
	public static void main(String[] args) {
		System.out.println("gender："+ Gender.man.value);
		System.out.println("gender："+ Gender.woman.value);
	}
}
 ```

 ```java
enum Gender{

	man("男"),woman("女");

	String value;

	private Gender(String value){
		this.value = value;
	}
}

public class Demo {
	public static void main(String[] args) {
		System.out.println("gender："+ Gender.man.value);
		System.out.println("gender："+ Gender.woman.value);
	}
}
 ```

 ```java
enum Color{
    RED, GREEN, BLUE;
}

public class Test{
    public static void main(String[] args){

        Color c1 = Color.RED;
        System.out.println(c1);  //RED
    }
}
 ```

```java

/*
内部类中使用枚举
枚举类也可以声明在内部类中：
*/
public class Test{

    enum Color{
        RED, GREEN, BLUE;
    }

    public static void main(String[] args){

        Color c1 = Color.RED;
        System.out.println(c1); //RED
    }
    }
}
```

```java

//可以使用 for 语句来迭代枚举元素：
enum Color{
    RED, GREEN, BLUE;
}
public class MyClass {
  public static void main(String[] args) {
    for (Color myVar : Color.values()) {
      System.out.println(myVar);
    }
  }
}
```

```java

//枚举类常应用于 switch 语句中：
//case语句后面跟的枚举值，只需要单写枚举值即可，不需要再声明该 枚举值是属于哪个枚举类的。
enum Color{
    RED, GREEN, BLUE;
}
public class MyClass {
  public static void main(String[] args) {

    Color myVar = Color.BLUE;

    switch(myVar) {
      case RED:
        System.out.println("红色");
        break;
      case GREEN:
         System.out.println("绿色");
        break;
      case BLUE:
        System.out.println("蓝色"); //蓝色
        break;
    }
  }
}
```

```java
//values(), ordinal() 和 valueOf() 方法
enum Color{
    RED, GREEN, BLUE;
}

public class Test{
    public static void main(String[] args){
        // 调用 values()
        Color arr[] = Color.values();

        // 迭代枚举
        for (Color col : arr){
            // 查看索引
            System.out.println(col + " at index " + col.ordinal());  //RED at index 0 | GREEN at index 1 | BLUE at index 2
        }

        // 使用 valueOf() 返回枚举常量，不存在的会报错 IllegalArgumentException
        System.out.println(Color.valueOf("RED"));  //RED
        System.out.println(Color.valueOf("WHITE"));  //WHITE
    }
}
```

```java
/*
枚举类成员
枚举跟普通类一样可以用自己的变量、方法和构造函数，构造函数只能使用 private 访问修饰符，所以外部无法调用。
枚举既可以包含具体方法，也可以包含抽象方法。 如果枚举类具有抽象方法，则枚举类的每个实例都必须实现它。
*/
enum Color{

    RED, GREEN, BLUE;

    // 构造函数
    private Color(){
        System.out.println("Constructor called for : " + this.toString());
    }

    public void colorInfo(){
        System.out.println("Universal Color");
    }
}

public class Test{
    public static void main(String[] args){

        Color c1 = Color.RED;
        System.out.println(c1);
        c1.colorInfo();
    }
}
	/*
	Constructor called for : RED
	Constructor called for : GREEN
	Constructor called for : BLUE
	RED
	Universal Color
	*/
```

```java
//自定义一个枚举类
enum Sex{

	man("男"){

		@Override
		public void run() {
			System.out.println("男人在跑...");
		}
		
	},woman("女"){

		@Override
		public void run() {
			System.out.println("女人在跑...");
		}
	}; //枚举值

	String value; //成员 变量

	//public static final Sex man = new Sex();

	//构造函数
	private Sex(String  value){
		this.value = value;
	}

	//成员函数
	public void getValue(){
		System.out.println("value :"+ value);
	}

	public abstract void run();

}

public class Demo {
	public static void main(String[] args) {

		Sex sex = Sex.man; //获取到了枚举类的对象
		sex.value = "男";

		sex.getValue();  //value :男

		sex.run();  //男人在跑...

		System.out.println("Sex.man.value :"+ Sex.man.value);  //Sex.man.value :男
	}
}
```

## IO流(Input Output)
IO技术主要的作用是解决设备与设备之间 的数据传输问题。
	比如： 硬盘--->内存      内存的数据---->硬盘上       把键盘的数据------->内存中

IO技术的应用场景：
	 导出报表、 上传、 下载 、

数据保存到硬盘上，该数据就可以做到永久性的保存。   数据一般是以文件的形式保存到硬盘上

### File类

File类可以描述一个文件或者一个文件夹。

File类的构造方法：

	`File(String pathname)`  指定文件或者文件夹的`路径`创建一个File文件。
	
	`File(File parent, String child)`   根据 parent 抽象路径名和 child 路径名字符串创建一个新 File 实例。 

	`File(String parent, String child)` 

目录分隔符：  在windows机器上 的目录分隔符是 `\`  ,在linux机器上的目录分隔符是`/` .

**注意**：  
	在windows上面`\ 与 /`都可以使用作为目录分隔符。 而且，`如果写 / 的时候只需要写一个即可`。

**路径问题：**
`绝对路径`： 该文件在硬盘上 的完整路径。绝对路径一般都是以盘符开头的。
`相对路径`:  相对路径就是资源文件相对于当前程序所在的路径。
注意： 如果程序当前所在的路径与资源文件不是在同一个盘下面，是没法写相对路径的。

 `.` 当前路径
 `..` 上一级路径

** File常用方法**

	getAbsolutePath()   目录分隔符

**创建：**
	createNewFile()		在指定位置创建一个空文件，成功就返回true，如果已存在就不创建然后返回false
	
	renameTo(File dest)	重命名文件或文件夹，也可以操作非空的文件夹。
					如果目标文件与源文件是在`同一个路径下`，那么renameTo的作用是`重命名`
					如果目标文件与源文件`不在同一个路径下`，那么renameTo的作用就是`剪切`，而且还不能操作文件夹。
					移动/重命名成功则返回true，失败则返回false。

**删除：**
delete()		删除文件或一个空文件夹，如果是文件夹且不为空，则不能删除，成功返回true，失败返回false。
deleteOnExit()	在虚拟机终止时，请求删除此抽象路径名表示的文件或目录，保证程序异常时创建的临时文件也可以被删除

**判断：**
	exists()		文件或文件夹是否存在。
	isFile()		是否是一个文件，如果不存在，则始终为false。
	isDirectory()	是否是一个目录，如果不存在，则始终为false。
	isHidden()		是否是一个隐藏的文件或是否是隐藏的目录。
	isAbsolute()	测试此抽象路径名是否为绝对路径名。

**获取：**
	getName()		获取文件或文件夹的名称，不包含上级路径。
	getPath()       返回绝对路径，可以是相对路径，但是目录要指定
	getAbsolutePath()	获取文件的绝对路径，与文件是否存在没关系
	length()		获取文件的大小（字节数），如果文件不存在则返回0L，如果是文件夹也返回0L。
	getParent()		返回此抽象路径名父目录的路径名字符串；如果此路径名没有指定父目录，则返回null。
	lastModified()	获取最后一次被修改的时间。

**文件夹相关：**
	mkdir()				在指定位置创建目录，这只会创建最后一级目录，如果上级目录不存在就抛异常。
	mkdirs()			在指定位置创建目录，这会`创建路径中所有不存在的目录`。

	staic File[] listRoots()	列出所有的根目录（Window中就是所有系统的盘符）
	list()						返回目录下的文件或者目录名，包含隐藏文件。对于文件这样操作会返回null。
	listFiles()					返回目录下的文件或者目录对象（File类实例），包含隐藏文件。对于文件这样操作会返回null。

	list(FilenameFilter filter)	把当前文件夹下面的所有子文件名与子文件夹名存储到一个String类型的数组中 返回。
	listFiles(FilenameFilter filter) 把当前文件夹下面的所有子文件与子文件夹都使用了一个Flie对象描述，然后把这些File对象存储到一个FIle数组中返回

```java
import java.io.File;

public class Demo {
	public static void main(String[] args) {

		File file1 = new File("E:/a.txt");
		
		File parentFile = new File("E:\\");
		File file2 = new File("F:\\","a.txt");
		
		System.out.println("存在吗？ "+ file.exists()); //存在吗？ false // exists 判断该文件是否存在，存在返回true，否则返回false。
		
		System.out.println("目录分隔符："+ File.separator); //目录分隔符：\
	}
}
```

```java
import java.io.File;

public class Demo {
	public static void main(String[] args) {
		
		File file = new  File(".");
		System.out.println("当前路径是："+ file.getAbsolutePath());    //当前路径是：E:\smile\.
		
		File file2 = new File("..\\..\\a.txt");
		System.out.println("存在吗？"+ file2.exists());  //存在吗？false
	}
}
```

```java
import java.io.File;
import java.io.IOException;

public class Demo {
	
	public static void main(String[] args) throws IOException {

		File file = new File("E:\\aa");

		////createNewFile 创建一个指定的文件，如果该文件存在了，则不会再创建，如果还没有存在则创建。创建成功返回true，否则返回false。
		System.out.println("创建成功了吗？"+file.createNewFile());  //创建成功了吗？true
		
		File dir = new  File("E:\\a.txt");
		System.out.println("创建文件夹成功吗？"+dir.mkdir()); //创建文件夹成功吗？true //mkdir 创建一个单级文件夹，

		dir = new File("F:\\aa\\bb");  //F不存在
		System.out.println("创建多级文件夹："+ dir.mkdirs()); //创建多级文件夹：false
		
		//renameTo()  
		//如果目标文件与源文件是在同一个路径下，那么renameTo的作用是重命名
		//如果目标文件与源文件不是在同一个路径下，那么renameTo的作用就是剪切，而且还不能操作文件夹。
		File destFile = new File("E:\\aaaaaaw");
		System.out.println("重命名成功吗？"+file.renameTo(destFile)) ; //重命名成功吗？true
	}
}
```

```java
import java.io.File;

public class Demo {
	public static void main(String[] args) {

		//删除的。
		File file = new File("E:\\a.txt");
		System.out.println("删除成功吗？ "+ file.delete()); //delete方法不能用于删除非空的文件夹。 
		file.deleteOnExit();  //jvm退出的时候删除文件。 一般用于删除临时 文件。		
		

		//判断
		File file = new File("..\\..\\a.txt");
		System.out.println("存在吗？"+ file.exists());
		System.out.println("判断是否是一个文件："+file.isFile()); //如果是文件返回true，否则返回false.
		System.out.println("判断是否是一个文件夹："+ file.isDirectory()); // 是文件夹返回ture，否则返回false.
		System.out.println("是隐藏的 文件吗："+ file.isHidden());
		System.out.println("是绝对路吗？"+ file.isAbsolute()); //测试此抽象路径名是否为绝对路径名。
	}
}
```

```java
import java.io.File;
import java.sql.Date;
import java.text.SimpleDateFormat;
import java.io.IOException;

public class  Demo {
	public static void main(String[] args) throws IOException{

		File file = new File("..\\c.txt");

		System.out.println("创建成功了吗？"+file.createNewFile());  //创建成功了吗？true

		System.out.println("存在吗？ "+ file.exists());  //存在吗？ true
		System.out.println("文件名："+ file.getName());   //文件名：c.txt
		System.out.println("获取相对路径："+ file.getPath()); 获取相对路径：..\c.txt
		System.out.println("getAbsolutePath获取绝对路径："+file.getAbsolutePath()); //getAbsolutePath获取绝对路径：E:\smile\..\c.txt 
		System.out.println("获取文件的父路径："+ file.getParent());  //获取文件的父路径：..
		System.out.println("获取文件的的大小(字节为单位)："+ file.length()); //获取文件的的大小(字节为单位)：0
		
		//使用毫秒值转换成Date对象
		long lastModified = file.lastModified();
		Date date = new Date(lastModified);	
		SimpleDateFormat dateFormat = new SimpleDateFormat("yyyy年MM月dd日  HH:mm:ss");  
		System.out.println("获取最后一次的修改时间(毫秒值)："+ dateFormat.format(date) );  //2020年07月24日  11:08:01
	}
}
```

```java
import java.io.File;

public class Demo {
	public static void main(String[] args) {

		File[] roots = File.listRoots(); //列出所有的根目录
		for(File file  : roots){
			System.out.println(file);    //C:\   D:\    E:\
		}
		
		File file = new File("E:\smile");
		String[] fileNames = file.list(); //把当前文件夹下面的所有子文件名与子文件夹名存储到一个String类型的数组中 返回。
		for(String fileName : fileNames){
			System.out.println(fileName);
		}
		/*
		Demo.class
		Demo.java
		ixfosa.txt
		java
		linux
		py
		*/

		File[] files = file.listFiles(); // 把当前文件夹下面的所有子文件与子文件夹都使用了一个File对象描述，然后把这些File对象存储到一个FIle数组中返回
		for(File fileItem : files){
			System.out.println("文件名："+ fileItem.getName());
		}
		/*
		文件名：Demo.class
		文件名：Demo.java
		文件名：ixfosa.txt
		文件名：java
		文件名：linux
		文件名：py
		 */
	}
}
```

```java
import java.io.File;
import java.io.FilenameFilter;

/*
需求1 ：  指定一个文件夹，然后该文件夹下面所有java文件。
需求2： 指定一个文件夹，然后列出文件夹下面的所有子文件与文件夹，但是格式要如下:
 
文件：
	文件名1
	文件名2
	文件名3
	..

文件夹：
	文件夹名1
	文件夹名2
	文件夹名3
	....
*/

// 自定义一个文件名过滤器
class MyFilter implements FilenameFilter{

	@Override
	public boolean accept(File dir, String name) {
		//System.out.println("文件夹:"+dir+" 文件名："+ name);
		/*文件夹:E:\smile 文件名：Demo.class
		文件夹:E:\smile 文件名：Demo.java
		文件夹:E:\smile 文件名：ixfosa.txt
		文件夹:E:\smile 文件名：java
		文件夹:E:\smile 文件名：linux
		文件夹:E:\smile 文件名：MyFilter.class
		文件夹:E:\smile 文件名：py*/
		return name.endsWith(".java");
	}
}

public class Demo {
	public static void main(String[] args) {

		File dir = new File("E:\\smile");
		//listJava2(dir);  //Demo.java
		//listJava(dir);  //Demo.java

		listFile(dir);
	}

	public static void listJava2(File dir){
		File[] files = dir.listFiles(new MyFilter()); //得到文件夹下面的所有子文件与文件夹。
		for(File file : files){
			System.out.println(file.getName());
		}
	}

	//列出所有的java文件
	public static void listJava(File dir){

		File[] files = dir.listFiles(); //获取到了所有的子文件

		for(File file : files){
			String fileName = file.getName();

			/*if(fileName.endsWith(".java")&&file.isFile()){
				System.out.println(fileName);
			}*/

			if(fileName.matches(".+\\.java")&&file.isFile()){
				System.out.println(fileName);
			}
		}
	}

	public static void listFile(File dir){
		File[] files = dir.listFiles();//获取到所有的子文件
		System.out.println("文件：");
		for(File fileItem : files){
			if(fileItem.isFile()){
				System.out.println("\t"+fileItem.getName());
			}
		}

		System.out.println("文件夹：");
		for(File fileItem : files){
			if(fileItem.isDirectory()){
				System.out.println("\t"+fileItem.getName());
			}
		}
	}
}
```

### 字节流

IO流分类：
	如果是按照数据的流向划分：
		输入流
		输出流
		
	如果按照处理的单位划分：
		字节流: 字节流读取得都是文件中二进制数据，读取到二进制数据不会经过任何的处理。
		字符流： 字符流读取的数据是以字符为单位的 。 字符流也是读取文件中的二进制数据，不过会把这些二进制数据转换成我们能 识别的字符。  
				`字符流 = 字节流 + 解码`
#### 输入字节流
--------| `InputStream` 所有输入字节流的基类  抽象类
------------| `FileInputStream`  读取文件数据的输入字节流 

使用FileInputStream读取文件数据的步骤：
	1. 找到目标文件
	2. 建立数据的输入通道。
	3. 读取文件中的数据。
	4. 关闭 资源.

 问题1： 读取完一个文件的数据的时候，我不关闭资源有什么影响?
 答案： 资源文件一旦 使用完毕应该马上释放，否则其他的程序无法对该资源文件进行其他 的操作。

```java
import java.io.File;
import java.io.FileInputStream;
import java.io.FileNotFoundException;
import java.io.IOException;

public class Demo {
	public static void main(String[] args) throws IOException {

		//readTest1();
		//readTest2();
		//readTest3();
		readTest4();
	}

	/*-----------------------------------------------------------------------------*/

	//方式4：使用缓冲数组配合循环一起读取。28
	public static void readTest4() throws IOException{
		long startTime = System.currentTimeMillis();
		//找到目标文件
		File file = new File("E:\\smile\\java\\images\\java-thread.jpg");
		//建立数据的输入通道
		FileInputStream fileInputStream = new FileInputStream(file);
		//建立缓冲数组配合循环读取文件的数据。
		int length = 0; //保存每次读取到的字节个数。
		byte[] buf = new byte[1024]; //存储读取到的数据    缓冲数组 的长度一般是1024的倍数，因为与计算机的处理单位。  理论上缓冲数组越大，效率越高
		
		while((length = fileInputStream.read(buf))!=-1){ // read方法如果读取到了文件的末尾，那么会返回-1表示。
			System.out.print(new String(buf,0,length));
		}
		
		//关闭资源
		fileInputStream.close();
		

		long endTime = System.currentTimeMillis();
		System.out.println("读取的时间是："+ (endTime-startTime)); //446
	}

    /*-----------------------------------------------------------------------------*/

	//方式3：使用缓冲数组 读取。    缺点： 无法读取完整一个文件的数据。     12G
	public static void readTest3() throws IOException{
		//找到目标文件
		File file = new File("E:\\a.txt");
		//建立数据的输入通道
		FileInputStream fileInputStream = new FileInputStream(file);
		//建立缓冲字节数组，读取文件的数据。
		byte[] buf = new byte[1024];  //相当于超市里面的购物车。
		int length = fileInputStream.read(buf); // 如果使用read读取数据传入字节数组，那么数据是存储到字节数组中的，而这时候read方法的返回值是表示的是本次读取了几个字节数据到字节数组中。
		System.out.println("length:"+ length);
		
		//使用字节数组构建字符串
		String content = new String(buf,0,length);
		System.out.println("内容："+ content);   //中文会乱码
		//关闭资源
		fileInputStream.close();
	}

	/*-----------------------------------------------------------------------------*/

	//方式2 ： 使用循环读取文件的数据
	public static void readTest2() throws IOException{
		long startTime = System.currentTimeMillis();
		//找到目标文件
		File file = new File("E:\\smile\\java\\images\\java-thread.jpg");
		//建立数据的输入通道
		FileInputStream fileInputStream = new FileInputStream(file);
		//读取文件的数据
		int content = 0; //声明该变量用于存储读取到的数据
		while((content = fileInputStream.read())!=-1){
			System.out.print(content);
		}
		//关闭资源
		fileInputStream.close();
		
		long endTime = System.currentTimeMillis();  
		System.out.println("读取的时间是："+ (endTime-startTime)); //446
	}

	/*-----------------------------------------------------------------------------*/

	//读取的方式一缺陷： 无法读取完整一个文件 的数据.
	public static void readTest1() throws IOException{
		//1. 找到目标文件
		File file = new File("E:\\a.txt");
		//建立数据的输入通道。
		FileInputStream fileInputStream = new FileInputStream(file);
		//读取文件中的数据
		int content = fileInputStream.read(); // read() 读取一个字节的数据，把读取的数据返回。
		System.out.println("读到的内容是："+ (char)content);
		//关闭资源    实际上就是释放资源。 
		 fileInputStream.close();
	}
}
```

```java
import java.io.File;
import java.io.FileInputStream;
import java.io.IOException;

import java.util.Arrays;

public class Demo {
	public static void main(String[] args) throws IOException {

		//找到目标文件
		File file = new File("E:\\a.txt");

		//建立数据的输入通道
		FileInputStream fileInputStream = new FileInputStream(file);

		//建立缓冲字节数组读取文件
		byte[] buf = new byte[8];
		int length = 0 ;
		
		while((length = fileInputStream.read(buf))!=-1){

			//System.out.print(new String(buf)); 

			/* 
			//a.txt 实际内容 ixfosalongzhong(15byte)  实际：ixfosalongzhongo
				ixfosalo
				12345678
				ngzhong        多输出 o
				分析： 建立缓冲字节数组byte[] buf = new byte[8]; 写满后数据并不会消失
			 */	

			//System.out.print(new String(buf,0,length));  //ixfosalongzhong
			
			//System.out.print(buf);  //[B@15db9742 [B@15db9742

			System.out.println(Arrays.toString(buf));
			//[105, 120, 102, 111, 115, 97, 108, 111]
			//[110, 103, 122, 104, 111, 110, 103, 111]
		}

		//释放文件资源
		fileInputStream.close();
	}
}
```

```java
import java.io.BufferedInputStream;
import java.io.BufferedOutputStream;
import java.io.File;
import java.io.FileInputStream;
import java.io.FileOutputStream;
import java.io.IOException;

/*
练习： 使用缓冲输入输出字节流拷贝一个图片。

*/
public class Demo {
	public static void main(String[] args) throws IOException {

		//找到目标文件
		File  inFile = new File("E:\\smile\\java\\images\\java-thread.jpg");
		File outFile = new File("E:\\1.jpg");

		//建立数据输入输出通道
		FileInputStream fileInputStream = new FileInputStream(inFile);
		FileOutputStream fileOutputStream = new FileOutputStream(outFile);

		//建立缓冲输入输出流
		BufferedInputStream bufferedInputStream = new BufferedInputStream(fileInputStream);
		BufferedOutputStream bufferedOutputStream = new BufferedOutputStream(fileOutputStream);

		//边读边写
		int content = 0; 

		//int length = bufferedInputStream.read(buf);   如果传入了缓冲数组，内容是存储到缓冲数组中，返回值是存储到缓冲数组中的字节个数。
		while((content = bufferedInputStream.read())!=-1){ // 如果使用read方法没有传入缓冲数组，那么返回值是读取到的内容。
			bufferedOutputStream.write(content);
			//bufferedOutputStream.flush();
		}
		//关闭资源
		bufferedInputStream.close();
		bufferedOutputStream.close();
	}
}
```

####  输出字节流

--------| `OutputStream` 是所有输出字节流的父类。  `抽象类`
 -----------| `FileOutStream` 向文件输出数据的输出字节流。

**FileOutputStream要注意的细节：**
	1. 使用FileOutputStream 的时候，如果目标文件不存在，那么会自动创建目标文件对象。 
	2. 使用FileOutputStream写数据的时候，如果目标文件已经存在，那么会先清空目标文件中的数据，然后再写入数据。
	3. 使用FileOutputStream写数据的时候, 如果目标文件已经存在，需要在原来数据基础上追加数据的时候应该使用`new FileOutputStream(file,true)构造函数，第二参数为true`。
	4. 使用FileOutputStream的write方法写数据的时候，虽然接收的是一个int类型的数据，但是真正写出的只是一个字节的数据，只是
	把低八位的二进制数据写出，其他二十四位数据全部丢弃。
	 
**FileOutputStream如何使用**
	1. 找到目标文件
	2. 建立数据的输出通道。
	3. 把数据转换成字节数组写出。
	4. 关闭资源

```java

import java.io.File;
import java.io.FileInputStream;
import java.io.FileNotFoundException;
import java.io.FileOutputStream;
import java.io.IOException;

public class Demo {
	
	public static void main(String[] args) throws IOException {

		writeTest3();
		//writeTest2();
		//writeTest1();
	}
	
	//使用字节数组把数据写出。
	public static void writeTest3() throws IOException{

		//找到目标文件
		File file = new File("E:\\b.txt");

		//建立数据输出通道
		FileOutputStream fileOutputStream = new FileOutputStream(file);

		//把数据写出。
		String data = "abc";
		byte[] buf = data.getBytes();

		fileOutputStream.write(buf, 0, 3); // 0 从字节数组的指定索引值开始写， 2：写出两个字节。

		//关闭资源
		fileOutputStream.close();
	}
		
	
	//使用字节数组把数据写出。
	public static void writeTest2() throws IOException{

		//找到目标文件
		File file = new File("E:\\b.txt");

		//建立数据输出通道
		FileOutputStream fileOutputStream = new FileOutputStream(file,true);

		//把数据写出。
		String data = "\r\nhello world";

		fileOutputStream.write(data.getBytes());

		//关闭资源
		fileOutputStream.close();
	}
	
	//每次只能写一个字节的数据出去。
	public static void writeTest1() throws IOException{

		//找到目标文件
		File file = new File("E:\\b.txt");

		//建立数据的输出通道
		FileOutputStream fileOutputStream = new FileOutputStream(file);

		//把数据写出
		fileOutputStream.write('h');
		fileOutputStream.write('e');
		fileOutputStream.write('l');
		fileOutputStream.write('l');
		fileOutputStream.write('o');

		//关闭资源
		fileOutputStream.close();
		
	}
}
```

```java
import java.io.File;
import java.io.FileInputStream;
import java.io.FileNotFoundException;
import java.io.FileOutputStream;
import java.io.IOException;
import java.util.Arrays;

/*
使用FileOutputStream的write方法写数据的时候，虽然接收的是一个int类型的数据，但是真正写出的只是一个字节的数据，只是
把低八位的二进制数据写出，其他二十四位数据全部丢弃。
 
 00000000-000000000-00000001-11111111   511
*/
public class Demo {
	public static void main(String[] args) throws IOException {

		readTest();
		writeTest();
	}
	
	public static void readTest() throws IOException{

		//找到目标文件
		File file = new File("E:\\c.txt");

		//建立数据的输入通道
		FileInputStream fileInputStream = new FileInputStream(file);

		//建立缓冲输入读取文件数据
		byte[] buf = new byte[4];

		//读取文件数据
		int length = fileInputStream.read(buf);

		System.out.println("字节数组的内容："+ Arrays.toString(buf));  //字节数组的内容：[-1, 0, 0, 0]

		//关闭资源
		fileInputStream.close();
		
	}
	
	public static void writeTest() throws FileNotFoundException, IOException {

		//找到目标文件
		File file = new File("E:\\c.txt");

		//建立数据的输出通道
		FileOutputStream fileOutputStream = new FileOutputStream(file);

		//把数据写出
		fileOutputStream.write(511);

		//关闭资源
		fileOutputStream.close();
	}
}
```

```java
import java.io.File;
import java.io.FileInputStream;
import java.io.FileOutputStream;
import java.io.IOException;

/*
需求： 拷贝一张图片。
*/
public class CopyImage {
	public static void main(String[] args) throws IOException {

		//找到目标文件
		File inFile = new File("E:\\smile\\java\\images\\java-thread.jpg");
		File destFile = new File("E:\\1.jpg");

		//建立数据的输入输出通道
		FileInputStream fileInputStream = new  FileInputStream(inFile);
		FileOutputStream fileOutputStream = new FileOutputStream(destFile); //追加数据....
		
		//每新创建一个FileOutputStream的时候，默认情况下FileOutputStream的指针是指向了文件的开始的位置。每写出一次，指向都会出现相应移动。
		//建立缓冲数据，边读边写
		byte[] buf = new byte[1024]; 
		int length = 0 ;

		while((length = fileInputStream.read(buf))!=-1){ //最后一次只剩下了824个字节
			fileOutputStream.write(buf,0,length); //写出很多次数据，所以就必须要追加。
		}

		//关闭资源 原则： 先开后关，后开先关。
		fileOutputStream.close();
		fileInputStream.close();
	}

}
```

#### 缓冲输入字节流

输入字节流体系： 
----| InputStream  输入字节流的基类。 抽象
----------| FileInputStream 读取文件数据的输入字节流
----------| BufferedInputStream 缓冲输入字节流    缓冲输入字节流的出现主要是为了提高读取文件数据的效率。
                                其实该类内部只不过是`维护了一个8kb的字节数组`而已。 

注意： `凡是缓冲流都不具备读写文件的能力`。

使用BufferedInputStream的步骤:
	1. 找到目标文件。
	2. 建立数据 的输入通道
	3. 建立缓冲 输入字节流流
	4. 关闭资源


```java

import java.io.BufferedInputStream;
import java.io.BufferedOutputStream;
import java.io.File;
import java.io.FileInputStream;
import java.io.FileOutputStream;
import java.io.IOException;

public class Demo {
	public static void main(String[] args) throws IOException {

		readTest2();
	}

	public static void readTest2() throws IOException{

		//找到目标文件
		File file = new File("E:\\a.txt");
		
		FileInputStream fileInputStream= new FileInputStream(file);
		BufferedInputStream bufferedInputStream= new BufferedInputStream(fileInputStream);
		bufferedInputStream.read();
		
		FileOutputStream fileOutputStream= new FileOutputStream(file);
		BufferedOutputStream bufferedOutputStream= new BufferedOutputStream(fileOutputStream);
		fileOutputStream.write(null);
		
		//读取文件数据
		int content = 0 ;

		//疑问二：BufferedInputStream出现 的目的是了提高读取文件的效率，但是BufferedInputStream的read方法每次读取一个字节的数据
		//而FileInputStreram每次也是读取一个字节的数据，那么BufferedInputStream效率高从何而来？
		while((content = fileInputStream.read())!=-1){
			System.out.print((char)content);
		}

		//关闭资源
		bufferedInputStream.close();//调用BufferedInputStream的close方法实际上关闭的是FileinputStream.
	}


	//读取文件的时候我们都是使用缓冲数组读取。效率会更加高
	public static void readTest() throws IOException{

		File file = new File("E:\\a.txt");

		//建立数组通道
		FileInputStream fileInputStream = new FileInputStream(file);

		//建立缓冲数组读取数据
		byte[] buf = new byte[1024*8];

		int length = 0;

		while((length = fileInputStream.read(buf))!=-1){
			System.out.print(new String(buf,0,length));
		}

		//关闭资源
		fileInputStream.close();
	}
}
```

#### 缓冲输出字节流

--------| OutputStream  所有输出字节流的基类  抽象类
------------| FileOutputStream 向文件 输出数据 的输出字节流
------------| Bufferedoutputstream  缓冲输出字节流    BufferedOutputStream出现的目的是为了提高写数据的效率。 
        						 内部也是维护了一个`8kb的字节数组`而已。
 
使用BufferedOutputStream的步骤：
 	1. 找到目标文件
 	2. 建立数据的输出通道

 
BufferedOutputStream 要注意的细节
	1. 使用BufferedOutStream写数据的时候，它的write方法是是先把数据写到它内部维护的字节数组中。
 	2. 使用BufferedOutStream写数据的时候，它的write方法是是先把数据写到它内部维护的字节数组中，如果需要把数据真正的写到硬盘上面，需要
 	调用flush方法或者是close方法、 或者是内部维护的字节数组已经填满数据的时候。

```java
import java.io.BufferedOutputStream;
import java.io.File;
import java.io.FileNotFoundException;
import java.io.FileOutputStream;
import java.io.IOException;

public class Demo {

	public static void main(String[] args) throws IOException {

		//找到目标文件
		File file = new File("E:\\a.txt");
		
		//建立数据的输出通道
		FileOutputStream  fileOutputStream = new FileOutputStream(file);

		//建立缓冲输出字节流对象
		BufferedOutputStream bufferedOutputStream  = new BufferedOutputStream(fileOutputStream);

		//把数据写出
		bufferedOutputStream.write("hello world".getBytes()); 

		//把缓冲数组中内部的数据写到硬盘上面。
		bufferedOutputStream.flush();

		bufferedOutputStream.close();
	}
}
```

```java
import java.io.BufferedInputStream;
import java.io.BufferedOutputStream;
import java.io.File;
import java.io.FileInputStream;
import java.io.FileOutputStream;
import java.io.IOException;

/*
练习： 使用缓冲输入输出字节流拷贝一个图片。
*/
public class CopyImage {
	public static void main(String[] args) throws IOException {

		//找到目标文件
		File  inFile = new File("E:\\smile\\java\\images\\java-thread.jpg");
		File outFile = new File("E:\\1.jpg");

		//建立数据输入输出通道
		FileInputStream fileInputStream = new FileInputStream(inFile);
		FileOutputStream fileOutputStream = new FileOutputStream(outFile);

		//建立缓冲输入输出流
		BufferedInputStream bufferedInputStream = new BufferedInputStream(fileInputStream);
		BufferedOutputStream bufferedOutputStream = new BufferedOutputStream(fileOutputStream);

		//边读边写
		int content = 0;

		// int length = bufferedInputStream.read(buf);   如果传入了缓冲数组，内容是存储到缓冲数组中，返回值是存储到缓冲数组中的字节个数。
		
		while((content = bufferedInputStream.read())!=-1){ // 如果使用read方法没有传入缓冲数组，那么返回值是读取到的内容。
			bufferedOutputStream.write(content);
			//bufferedOutputStream.flush();
		}
		//关闭资源
		bufferedInputStream.close();
		bufferedOutputStream.close();
	}
}

```

####  IO异常 的处理
```java
import java.io.File;
import java.io.FileInputStream;
import java.io.FileOutputStream;
import java.io.IOException;

import javax.management.RuntimeErrorException;

/*
 IO异常 的处理

 */
public class Demo {
	public static void main(String[] args) {

		//readTest();
		copyImage();
	}

	// 拷贝图片
	public static void copyImage() {

		FileInputStream fileInputStream = null;
		FileOutputStream fileOutputStream = null;

		try {
			// 找到目标文件
			File inFile = new File("E:\\smile\\java\\images\\java-thread.jpg");
			File outFile = new File("E:\\1.jpg");

			// 建立输入输出通道
			fileInputStream = new FileInputStream(inFile);
			fileOutputStream = new FileOutputStream(outFile);

			// 建立缓冲数组，边读边写
			byte[] buf = new byte[1024];

			int length = 0;

			while ((length = fileInputStream.read(buf)) != -1) {
				fileOutputStream.write(buf, 0, length);
			}

		} catch (IOException e) {

			System.out.println("拷贝图片出错...");
			throw new RuntimeException(e);

		} finally {
			// 关闭资源
			try {

				if (fileOutputStream != null) {
					fileOutputStream.close();
					System.out.println("关闭输出流对象成功...");
				}

			} catch (IOException e) {

				System.out.println("关闭输出流资源失败...");
				throw new RuntimeException(e);

			} finally {

				if (fileInputStream != null) {
					try {

						fileInputStream.close();
						System.out.println("关闭输入流对象成功...");

					} catch (IOException e) {

						System.out.println("关闭输入流对象失败...");
						throw new RuntimeException(e);
					}
				}

			}
		}
	}

	public static void readTest() {

		FileInputStream fileInputStream = null;
		try {

			// 找到目标文件
			File file = new File("F:\\aaaaa.txt");

			// 建立数据输入通道
			fileInputStream = new FileInputStream(file);

			// 建立缓冲数组读取数据
			byte[] buf = new byte[1024];

			int length = 0;

			while ((length = fileInputStream.read(buf)) != -1) {
				System.out.print(new String(buf, 0, length));
			}
		} catch (IOException e) {
			/*
			 * //处理的代码... 首先你要阻止后面的代码执行,而且要需要通知调用者这里出错了... throw new
			 * RuntimeException(e);
			 * //把IOException传递给RuntimeException包装一层，然后再抛出，这样子做的目的是
			 * 为了让调用者使用变得更加灵活。
			 */
			System.out.println("读取文件资源出错....");
			throw new RuntimeException(e);
		} finally {
			try {
				if (fileInputStream != null) {
					fileInputStream.close();
					System.out.println("关闭资源成功...");
				}
			} catch (IOException e) {
				System.out.println("关闭资源失败...");
				throw new RuntimeException(e);
			}
		}
	}
}
```

### 字符流

`字节流`：字节流读取的是文件中的二进制数据，读到的数据并不会帮你转换成你看得懂的字符。 
	**输入字节流**
	------| `InputStream` 输入字节流的基类   抽象类
	-----------| `FileInputStream` 读取文件数据的输入字节流。
	-----------| `BufferedInputStream` 缓冲输入字节流       缓冲输入字节流出现的目的： 为了提高读取文件数据的效率。 该类其实内部就是维护了一个8kb字节数组而已。
 
	**输出字节流：**
		------| `OutputStream` 输出字节流的基类。  抽象类。
		-----------| `FileOutStream`  向文件输出数据的输出字节流。
		-----------| `BufferedOutputStream` 缓冲输出字节流。  该类出现的目的是为了提高写数据的效率 。 其实该类内部也是维护了一个8kb的数组而已，当调用其write方法的时候数据默认是向它内部的数组中存储 的，只有调用flush方法或者是close方法或者是8kb的字节数组存储满数据的时候才会真正的向硬盘输出。


`字符流`：字符流会把读取到的二进制的数据进行对应 的编码与解码工作。   字符流 = 字节流 + 编码(解码)
	**输入字符流**
		------| `Reader` 输入字符流的基类   抽象类
		-----------|  `FileReader` 读取文件的输入字符流。

	**输出字符流**
		------| `Writer` 输出字符流的基类。 抽象类
		-----------| `FileWriter` 向文件数据数据的输出字符流

```java
import java.io.File;
import java.io.FileInputStream;
import java.io.FileOutputStream;
import java.io.IOException;
import java.util.Arrays;

public class Demo {
	public static void main(String[] args) throws IOException {
		readrTest();
	}

	//使用字节流读取中文
	public static void readrTest() throws IOException{

		//找到目标文件
		File file = new File("E:\\a.txt");

		//建立数据的输入通道
		FileInputStream fileInputStream = new FileInputStream(file);
		//读取内容
		//读取内容
		/*int content = 0;
		while((content = fileInputStream.read())!=-1){ //出现乱码的原因： 一个中文在gbk码表中默认是占两个字节，
					 										   //目前你只读取了一个字节而已，所以不是一个完整的中文。
			System.out.print((char)content);
		}*/
		
		byte[] buf = new byte[2];
		for(int i = 0 ; i < 3 ; i++){
			fileInputStream.read(buf);
			System.out.print(new String(buf));  //a.txt中 夏，  实际写出 夏夏夏
		}

		//关闭资源
		fileInputStream.close();
	}
}
```

#### 输入字符流-FileReader
----------| `Reader` 输入字符流的基类   抽象类
-------------|  `FileReader` 读取文件的输入字符流。


FileReader的用法：
	1. 找到目标文件
	2. 建立数据的输入通道
	3. 读取数据
	4. 关闭资源

```java
import java.io.File;
import java.io.FileNotFoundException;
import java.io.FileReader;
import java.io.IOException;

public class Demo {
	public static void main(String[] args) throws IOException {
		readTest2();
		//readTest1();
	}

	//使用缓冲字符数组读取文件。
	public static void readTest2() throws IOException{

		//找到目标文件
		File file = new File("E:\\smile\\Demo.java");

		// 建立数据的输入通道
		FileReader fileReader = new FileReader(file);

		//建立缓冲字符数组读取文件数据
		char[] buf = new char[1024];
		int length = 0;

		while((length = fileReader.read(buf))!=-1){
			System.out.print(new String(buf,0,length));  
		}

		//关闭资源
		fileReader.close();
	}
	
	public static void readTest1() throws IOException{

		//找到目标文件
		File file = new File("E:\\smile\\Demo.java");

		//建立数据的输入通道
		FileReader fileReader = new FileReader(file);

		int content = 0 ;

		while((content = fileReader.read())!=-1){ //每次只会读取一个字符，效率低。
			System.out.print((char)content);   
		}

		//关闭资源
		fileReader.close();
	}
}
```

#### 输出字符流-FileWriter

------| Writer 输出字符流的基类。 抽象类
-----------| FileWriter 向文件数据数据的输出字符流

FileWriter的使用步骤：
	1. 找到目标文件。
	2. 建立数据输出通道
	3. 写出数据。
	4. 关闭资源

FileWriter要注意的事项：
	1. 使用FileWriter写数据的时候，FileWriter内部是维护了一个1024个字符数组的，写数据的时候会先写入到它内部维护的字符数组中，如果需要
	把数据真正写到硬盘上，需要调用flush或者是close方法或者是填满了内部的字符数组。
	2. 使用FileWriter的时候，如果目标文件不存在，那么会自动创建目标文件。
	3. 使用FileWriter的时候， 如果目标文件已经存在了，那么默认情况会先情况文件中的数据，然后再写入数据 ， 如果需要在原来的基础上追加数据，
	需要使用“new FileWriter(File , boolean)”的构造方法，第二参数为true。

```java
import java.io.File;
import java.io.FileOutputStream;
import java.io.FileWriter;
import java.io.IOException;

public class Demo {
	public static void main(String[] args) throws IOException {
		writeTest();
	}

	public static void  writeTest() throws IOException{

		//找到目标文件E
		File file = new File("E:\\a.txt");

		//建立数据输出通道
		FileWriter fileWriter = new FileWriter(file,true);

		//准备数据，把数据写出
		String data = "辣条！！";

		fileWriter.write(data);  //字符流具备解码的功能。

		//刷新字符流
		//fileWriter.flush();

		//关闭资源
		fileWriter.close();
	}
}
```

```java
import java.io.File;
import java.io.FileInputStream;
import java.io.FileReader;
import java.io.IOException;

//使用字节流读取中文
public class Demo {

	public static void main(String[] args) throws IOException {

		File file = new File("E:\\a.txt");

		FileInputStream fileInputStream = new FileInputStream(file);

		byte[] buf = new byte[1024];
		int length = 0;

		while((length = fileInputStream.read(buf))!=-1){
			System.out.println(new String(buf,0,length)); //借用字符串的解码功能。 
		}
		fileInputStream.close();
	}
}
```

```java
import java.io.BufferedReader;
import java.io.BufferedWriter;
import java.io.File;
import java.io.FileNotFoundException;
import java.io.FileReader;
import java.io.FileWriter;
import java.io.IOException;
/*
何时使用字符流，何时使用字节流？依据是什么？

使用字符流的应用场景： 如果是读写字符数据的时候则使用字符流。

使用字节流的应用场景： 如果读写的数据都不需要转换成字符的时候，则使用字节流。 
 */
//使用字符流拷贝文件
public class Copy {
	public static void main(String[] args) throws IOException {
		BufferedReader bufferedReader = new BufferedReader(new FileReader("E:\\Test.txt"));
		BufferedWriter bufferedWriter = new BufferedWriter(new FileWriter("E:\\Test.exe"));
		String line=null;
		while((line = bufferedReader.readLine())!=null){
			bufferedWriter.write(line);
		}
		bufferedWriter.close();
		bufferedReader.close();
	}
}
```
#### 缓冲输入字符流-BufferedReader
-------| Reader 所有输入字符流的基类。 抽象类
----------| FileReader 读取文件字符串的输入字符流。
----------| BufferedReader   缓冲输入字符流  。 缓冲 输入字符流出现的目的是为了提高读取文件 的效率和拓展了FileReader的功能。  其实该类内部也是维护了一个字符数组

记住：缓冲流都不具备读写文件的能力。

BufferedReader的使用步骤：
	1.找到目标文件
	2 .建立数据的输入通道。

```java
import java.io.BufferedReader;
import java.io.File;
import java.io.FileReader;
import java.io.IOException;
import java.util.Arrays;

public class Demo {
	public static void main(String[] args) throws IOException {

		//readTest1();
		File file  = new File("E:\\smile\\Demo.java");
		//建立数据的输入通道。
		FileReader fileReader = new FileReader(file);
		String line =  null;
		
		while((line = myReadLine(fileReader))!=null){
			System.out.println(line);
		}
	}

	//自己去实现readLine方法。
	public static String myReadLine(FileReader fileReader) throws IOException{

		//创建一个字符串缓冲类对象
		StringBuilder sb = new StringBuilder();  //StringBuilder主要是用于存储读取到的数据,非线程安全

		int content = 0 ;

		while((content = fileReader.read())!=-1){
			if(content=='\r'){
				continue;
			}else if(content=='\n'){
				break;
			}else{
				//普通字符
				sb.append((char)content);
			}
		}
		//代表已经读取完毕了。
		if(content ==-1){
			return null;
		}
		
		return sb.toString();  
	}

	public static void readTest1() throws IOException{

		//找到目标文件
		File file  = new File("E:\\a.txt");

		//建立数据的输入通道。
		FileReader fileReader = new FileReader(file);

		//建立缓冲输入字符流
		BufferedReader bufferedReader = new BufferedReader(fileReader);

		//读取数据
		/*int content = bufferedReader.read(); //读到了一个字符。读取到的字符肯定也是从Bufferedreader内部的字符数组中获取的到,所以效率高。
		System.out.println((char)content);*/
		//使用BUfferedReader拓展的功能，readLine()  一次读取一行文本，如果读到了文件的末尾返回null表示。

		String line =  null;
		while((line = bufferedReader.readLine())!=null){ // 虽然readLine每次读取一行数据，但是但会的line是不包含\r\n的、
			System.out.println(Arrays.toString("aaa".getBytes()));
		}
		//关闭资源
		bufferedReader.close();
	}
}

```
#### 缓冲输出字符流-BufferedWriter
输出字符流
----------| Writer  所有输出字符流的基类，  抽象类。
--------------- | FileWriter 向文件输出字符数据的输出字符流。
----------------| BufferedWriter 缓冲输出字符流          缓冲输出字符流作用： 提高FileWriter的写数据效率与拓展FileWriter的功能。
BufferedWriter内部只不过是提供了一个8192长度的字符数组作为缓冲区而已，拓展了FileWriter的功能。

BufferedWriter如何使用？
	1. 找到目标文件
 	2. 建立数据的输出通道

```java
import java.io.BufferedWriter;
import java.io.File;
import java.io.FileWriter;
import java.io.IOException;
public class Demo {
	public static void main(String[] args) throws IOException {

		//找到目标文件
		File file = new File("E:\\a.txt");

		//建立数据的输出通道
		FileWriter fileWriter = new FileWriter(file,true);

		//建立缓冲输出流对象
		BufferedWriter bufferedWriter = new BufferedWriter(fileWriter);

		//写出数据
		//bufferedWriter.newLine(); //newLine() 换行。 实际上就是想文件输出\r\n.
		bufferedWriter.write("\r\n");
		bufferedWriter.write("前两天李克强来萝岗！！");

		//关闭资源
		bufferedWriter.flush();
		bufferedWriter.close();
	}
}
```

```java
import java.io.BufferedReader;
import java.io.BufferedWriter;
import java.io.File;
import java.io.FileReader;
import java.io.FileWriter;
import java.io.IOException;
import java.util.Scanner;

/*
练习： 缓冲输入输出字符流用户登陆注册...
	
*/
public class Login {

	static Scanner scanner = new Scanner(System.in);

	public static void main(String[] args) throws IOException {

		while(true){
			System.out.println("请选择功能： A(注册)   B(登陆)");
			String option =  scanner.next();
			if("a".equalsIgnoreCase(option)){
				//注册
				reg();
			}else if("b".equalsIgnoreCase(option)){
				//登陆
				login();
			}else{
				System.out.println("你的输入有误,请重新输入...");
			}
		}
	}

	//登陆
	public static void login() throws IOException{
		System.out.println("请输入用户名：");
		String userName = scanner.next();
		System.out.println("请输入密码：");
		String password = scanner.next();
		String info = userName+" "+ password;

		//读取文件的信息，查看文件否有该用户的信息存在，如果存在则登陆成功。
		//建立数据的输入通道
		//建立缓冲输入字符流
		BufferedReader bufferedReader = new BufferedReader(new FileReader("E:\\users.txt"));

		String line = null;
		
		boolean isLogin = false; // 用于记录是否登陆成功的标识， 默认是登陆失败的。

		//不断的读取文件的内容
		while((line = bufferedReader.readLine())!=null){
			if(info.equals(line)){
				isLogin = true;
				break;
			}
		}
		
		if(isLogin){
			System.out.println("欢迎"+userName+"登陆成功...");
		}else{
			System.out.println("不存在该用户信息，请注册!!");
		}
	}

	//注册
	public static void reg() throws IOException{

		System.out.println("请输入用户名：");
		String userName = scanner.next();
		System.out.println("请 输入密码：");
		String password = scanner.next();
		String info = userName+" "+ password;

		//把用户的注册的信息写到文件上
		File file = new File("F:\\users.txt");

		FileWriter fileWriter = new FileWriter(file,true);

		//建立缓冲输出字符流
		BufferedWriter bufferedWriter = new BufferedWriter(fileWriter);

		//把用户信息写出
		bufferedWriter.write(info);
		bufferedWriter.newLine();

		//关闭资源
		bufferedWriter.close();
	}
}
```

####  装饰者设计模式

装饰者设计模式：增强一个类的功能，而且还可以让这些装饰类互相装饰。

装饰者设计模式的步骤：
	1. 在装饰类的内部维护一个被装饰类的引用。
	2. 让装饰类有一个共同的父类或者是父接口。

需求1： 编写一个类拓展BufferedReader的功能， 增强readLine方法返回 的字符串带有行号。
 
需求2：编写一个类拓展BufferedReader的功能， 增强readLine方法返回 的字符串带有分号。
  
需求3： 编写一个类拓展BufferedReader的功能， 增强readLine方法返回 的字符串带有双引号。
  
需求4： 编写一个类拓展BufferedReader的功能， 增强readLine方法返回 的字符串带有行号+ 分号。
  
需求5： 编写一个类拓展BufferedReader的功能， 增强readLine方法返回 的字符串带有分号+ 双引号。

需求6： 编写一个类拓展BufferedReader的功能， 增强readLine方法返回 的字符串带有双引号+ 行号。

需求7： 编写一个类拓展BufferedReader的功能， 增强readLine方法返回 的字符串带有行号+ 分号+双引号。

继承实现的增强类和修饰模式实现的增强类有何区别？
	继承实现的增强类：
		优点：代码结构清晰，而且实现简单. 
		缺点：对于每一个的需要增强的类都要创建具体的子类来帮助其增强，这样会导致继承体系过于庞大。

修饰模式实现的增强类：
		优点：内部可以通过多态技术对多个需要增强的类进行增强， 可以是这些装饰类达到互相装饰的效果。使用比较灵活。
		缺点：需要内部通过多态技术维护需要被增强的类的实例。进而使得代码稍微复杂。

```java
import java.io.BufferedReader;
import java.io.File;
import java.io.FileNotFoundException;
import java.io.FileReader;
import java.io.IOException;
import java.io.Reader;
/*
	增强一个类的功能时候我们可以选择使用继承：
		通过继承实现增强一个类的功能
		优点：  码结构清晰，通俗易懂。
		缺点： 使用不灵活，会导致继承的体系过于庞大。
 */

 public class Demo {

	public static void main(String[] args) throws IOException {

		File file = new File("E:\\smile\\Demo.java");
		//建立数据的输入通道
		FileReader fileReader = new FileReader(file);
		//建立带行号的缓冲输入字符流
		BufferedLineNum bufferedLineNum = new BufferedLineNum(fileReader);
		
		//带有分号的缓冲输入字符流
		BufferedSemi bufferedSemi = new BufferedSemi(fileReader);
		
		//带有双引号的缓冲输入字符流
		BufferedQuto bufferedQuto = new BufferedQuto(fileReader);
		
		String line = null;
		while((line = bufferedQuto.readLine())!=null){
			System.out.println(line);
		}
	}
 }

class BufferedLineNum extends BufferedReader{

	//行号
	int count = 1 ;

	public BufferedLineNum(Reader in) {
		super(in);
	}

	@Override
	public String readLine() throws IOException {

		String line = super.readLine();

		if(line ==null){
			return null;
		}
		line = count+" "+ line;
		count++;
		return line;
	}
}

//带分号的缓冲输入字符流
class BufferedSemi extends BufferedReader{

	public BufferedSemi(Reader in) {
		super(in);
	}

	@Override
	public String readLine() throws IOException {
		String line =  super.readLine();
		if(line==null){
			return null;
		}
		line = line+";";
		return line;
	}
}

//带双引号的缓冲输入字符流
class  BufferedQuto extends BufferedReader{

	public BufferedQuto(Reader in) {
		super(in);
	}

	@Override
	public String readLine() throws IOException {
		String line = super.readLine();
		if(line==null){
			return null;
		}
		line = "\""+line+"\"";
		return line;
	}
}
```

```java

import java.io.BufferedReader;
import java.io.File;
import java.io.FileNotFoundException;
import java.io.FileReader;
import java.io.IOException;
/*
装饰者设计模式：增强一个类的功能，而且还可以让这些装饰类互相装饰。

装饰者设计模式的步骤：
	1. 在装饰类的内部维护一个被装饰类的引用。
	2. 让装饰类有一个共同的父类或者是父接口。

修饰模式实现的增强类：
		优点：内部可以通过多态技术对多个需要增强的类进行增强， 可以是这些装饰类达到互相装饰的效果。使用比较灵活。
		
		缺点：需要内部通过多态技术维护需要被增强的类的实例。进而使得代码稍微复杂。
*/

//带行号的缓冲输入字符流
class BufferedLineNum2  extends BufferedReader{  

	//在内部维护一个被装饰类的引用。
	BufferedReader bufferedReader;

	int count = 1;

	public BufferedLineNum(BufferedReader bufferedReader){
		super(bufferedReader);// 注意： 该语句没有任何的作用，只不过是为了让代码不报错。
		this.bufferedReader = bufferedReader;
	}

	public String readLine() throws IOException{
		String line = bufferedReader.readLine();
		if(line==null){
			return null;
		}
		line = count+" "+line;
		count++;
		return line;
	}
}

//带分号缓冲输入字符流
class BufferedSemi extends BufferedReader{  //为什么要继承?  是为了让这些装饰类的对象可以作为参数进行传递，达到互相装饰 的效果。

	//在内部维护一个被装饰类的引用。
	BufferedReader bufferedReader;
		
	public BufferedSemi2(BufferedReader bufferedReader){ // new BuffereLineNum();
		super(bufferedReader);// 注意： 该语句没有任何的作用，只不过是为了让代码不报错。
		this.bufferedReader = bufferedReader;
	}

	public String readLine() throws IOException{
		String line = bufferedReader.readLine();  //如果这里的ReadLine方法是调用了buffereLineNum的readLine方法，问题马上解决。
		if(line==null){
			return null;
		}
		line = line +";";
		return line;
	}

}

//缓冲类带双引号
class BufferedQuto extends BufferedReader{

	//在内部维护一个被装饰的类
	BufferedReader bufferedReader;

	public BufferedQuto2(BufferedReader bufferedReader){  //new  BufferedSemi2();
		super(bufferedReader) ; //只是为了让代码不报错..
		this.bufferedReader = bufferedReader;
	}

	public String readLine() throws IOException{
		String line = bufferedReader.readLine();  //如果这里的ReadLine方法是调用了buffereLineNum的readLine方法，问题马上解决。
		if(line==null){
			return null;
		}
		line = "\""+line +"\"";
		return line;
	}
}
```

```java
/*练习：
	一家三口每个人都会工作，儿子的工作就是画画，母亲的工作就是在儿子的基础上做一个增强，不单止可以画画，还可以上涂料。
	爸爸的工作就是在妈妈基础上做了增强，就是上画框。
*/

interface Work{
	public void work();
}

class Son implements Work{

	@Override
	public void work() {
		System.out.println("画画...");
	}
}

class Mather implements Work{

	//需要被增强的类。
	Work worker;

	public Mather(Work worker){
		this.worker = worker;
	}

	@Override
	public void work() {
		worker.work();
		System.out.println("给画上颜色..");
	}
}

class Father implements Work{

	//需要被增强的类的引用
	Work worker;

	public Father(Work worker){
		this.worker = worker;
	}

	@Override
	public void work() {
		worker.work();
		System.out.println("上画框...");
	}
}

public class Demo {
	public static void main(String[] args) {

		Son s = new Son();
//		s.work();
		Mather m = new Mather(s);
//		m.work();
		Father f = new Father(s);
		f.work();
	}
}
```

### 序列流-SequenceInputStream

#### 文件的内容合并

```java
import java.io.File;
import java.io.FileInputStream;
import java.io.FileOutputStream;
import java.io.IOException;
import java.io.SequenceInputStream;
import java.util.ArrayList;
import java.util.Enumeration;
import java.util.Vector;

/*
SequenceInputStream(序列流)

需求：把a.txt与b.txt 文件的内容合并。
 */
public class Demo {
	public static void main(String[] args) throws IOException {
		merge3();
	}

	//把三个文件合并成一个文件
	public static void merge3() throws IOException{
		//找到目标文件
		File file1 = new File("E:\\a.txt");
		File file2 = new File("E:\\b.txt");
		File file3 = new File("E:\\c.txt");
		File file4 = new File("E:\\d.txt");
		
		
		//建立对应 的输入输出流对象
		FileOutputStream fileOutputStream = new FileOutputStream(file4);
		FileInputStream fileInputStream1 = new FileInputStream(file1);
		FileInputStream fileInputStream2 = new FileInputStream(file2);
		FileInputStream fileInputStream3 = new FileInputStream(file3);

		//创建序列流对象
		Vector<FileInputStream> vector = new Vector<FileInputStream>();
		vector.add(fileInputStream1);
		vector.add(fileInputStream2);
		vector.add(fileInputStream3);
		Enumeration<FileInputStream> e = vector.elements();

		SequenceInputStream sequenceInputStream = new SequenceInputStream(e);
		
		//读取文件数据
		byte[] buf = new byte[1024];
		int length = 0;
		
		while((length = sequenceInputStream.read(buf))!=-1){
			fileOutputStream.write(buf,0,length);
		}
		
		//关闭资源
		sequenceInputStream.close();
		fileOutputStream.close();
	}

//	使用SequenceInputStream合并文件。
	public static void merge2() throws IOException{

		//找到目标文件
		File inFile1 = new File("E:\\a.txt");
		File inFile2 = new File("E:\\b.txt");
		File outFile = new File("E:\\c.txt");

		//建立数据的输入输出通道
		FileOutputStream fileOutputStream = new FileOutputStream(outFile);
		
		FileInputStream fileInputStream1 = new FileInputStream(inFile1);
		FileInputStream fileInputStream2 = new FileInputStream(inFile2);

		//建立序列流对象
		SequenceInputStream inputStream = new SequenceInputStream(fileInputStream1,fileInputStream2);

		byte[] buf = new byte[1024];
		int length = 0;
		
		while((length = inputStream.read(buf))!=-1){
			fileOutputStream.write(buf,0,length);
		}
		//关闭资源
		inputStream.close();
		fileOutputStream.close();
	}

	//需求：把a.txt与b.txt 文件的内容合并。
	public static void merge1() throws IOException{

		//找到目标文件
		File inFile1 = new File("F:\\a.txt");
		File inFile2 = new File("F:\\b.txt");
		File outFile = new File("F:\\c.txt");

		//建立数据的输入输出通道
		FileOutputStream fileOutputStream = new FileOutputStream(outFile);
		FileInputStream fileInputStream1 = new FileInputStream(inFile1);
		FileInputStream fileInputStream2 = new FileInputStream(inFile2);
		
		//把输入流存储到集合中，然后再从集合中读取
		ArrayList<FileInputStream> list = new ArrayList<FileInputStream>();
		list.add(fileInputStream1);
		list.add(fileInputStream2);
		
		//准备一个缓冲数组
		byte[] buf = new byte[1024];
		int length = 0 ;
		
		for(int i = 0 ; i< list.size() ; i++){
			FileInputStream fileInputStream = list.get(i);
			while((length = fileInputStream.read(buf))!=-1){
				fileOutputStream.write(buf,0,length);
			}
			//关闭资源
			fileInputStream.close();
		}
		fileOutputStream.close();
	}
}
```
#### 文件切割

```java
import java.io.File;
import java.io.FileInputStream;
import java.io.FileNotFoundException;
import java.io.FileOutputStream;
import java.io.IOException;
import java.io.SequenceInputStream;
import java.util.Enumeration;
import java.util.Vector;

/*
 需求： 把一首mp3先切割成n份，然后再把这些文件合并起来。
 */
public class Demo {

	public static void main(String[] args) throws IOException {
//		cutFile();
		mergeFlile();
	}

	//合并
	public static void mergeFlile() throws IOException{

		//找到目标文件
		File dir = new File("F:\\music");

		//通过目标文件夹找到所有的MP3文件，然后把所有的MP3文件添加到vector中。
		Vector<FileInputStream> vector = new Vector<FileInputStream>();
		File[] files = dir.listFiles();

		for(File file : files){
			if(file.getName().endsWith(".mp3")){
				vector.add(new FileInputStream(file));
			}
		}

		//通过Vector获取迭代器
		Enumeration<FileInputStream> e = vector.elements();

		//创建序列流
		SequenceInputStream inputStream = new SequenceInputStream(e);

		//建立文件的输出通道
		FileOutputStream fileOutputStream = new FileOutputStream("F:\\合并.mp3");

		//建立缓冲数组读取文件
		byte[] buf = new byte[1024];
		int length = 0 ;

		while((length =  inputStream.read(buf))!=-1){
			fileOutputStream.write(buf,0,length);
		}
		//关闭资源
		fileOutputStream.close();
		inputStream.close();
	}

	//切割MP3
	public static void cutFile() throws IOException{

		File file = new File("F:\\1.mp3");
		//目标文件夹
		File dir = new File("F:\\music");

		//建立数据的输入通道
		FileInputStream fileInputStream = new FileInputStream(file);

		//建立缓冲数组读取
		byte[] buf = new byte[1024*1024];
		int length = 0;

		for(int i = 0 ;  (length = fileInputStream.read(buf))!=-1 ; i++){
			FileOutputStream fileOutputStream =	new FileOutputStream(new File(dir,"part"+i+".mp3"));
			fileOutputStream.write(buf,0,length);
			fileOutputStream.close();
		}
		//关闭资源
		fileInputStream.close();
	}
}
```
###  对象的输入输出流- 序列化

对象的输入输出流 : 对象的输入输出流 主要的作用是用于写对象的信息与读取对象的信息。
	对象信息一旦写到文件上那么对象的信息就可以做到持久化了
 
对象的输出流： `ObjectOutputStream` 
 
对象的输入流:  `ObjectInputStream` 
 
**对象输入输出流要注意的细节：**
 	1. 如果对象需要被写出到文件上，那么对象所属的类必须要实现`Serializable接口`。 Serializable接口没有任何的方法，是一个`标识接口`而已。
   
 	2. 对象的反序列化创建对象的时候并不会调用到构造方法的、
   
 	3. `serialVersionUID` 是用于记录class文件的版本信息的，serialVersionUID这个数字是通过一个类的类名、成员、包名、工程名算出的一个数字。
   
 	4. 使用ObjectInputStream反序列化的时候，ObjectInputStream会先读取文件中的serialVersionUID，然后与本地的class文件的serialVersionUID
 	进行对比，如果这两个id不一致，那么反序列化就失败了。

 	5. 如果序列化与反序列化的时候可能会修改类的成员，那么最好一开始就给这个类指定一个serialVersionUID，如果一类已经指定的serialVersionUID，然后在序列化与反序列化的时候，jvm都不会再自己算这个 class的serialVersionUID了。

 	6. 如果一个对象某个数据不想被序列化到硬盘上，可以使用关键字`transient`修饰。
   
 	7. 如果一个类维护了另外`一个类的引用`，那么另外一个类也需要实现`Serializable接口`。
 	
 */
```java
import java.io.File;
import java.io.FileInputStream;
import java.io.FileOutputStream;
import java.io.IOException;
import java.io.ObjectInput;
import java.io.ObjectInputStream;
import java.io.ObjectOutputStream;
import java.io.Serializable;

class Address implements Serializable{

	String country;

	String city;

	public Address(String country,String city){
		this.country = country;
		this.city = city;
	}
}

class User implements Serializable{

	private static final long serialVersionUID = 1L;

	String userName ;

	String password;

	transient int age;  // transient 透明

	Address address ;

	public User(String userName , String passwrod) {
		this.userName = userName;
		this.password = passwrod;
	}

	public User(String userName , String passwrod,int age,Address address) {
		this.userName = userName;
		this.password = passwrod;
		this.age = age;
		this.address = address;
	}

	@Override
	public String toString() {
		return "用户名："+this.userName+ " 密码："+ this.password+" 年龄："+this.age+" 地址："+this.address.city;
	}
}

public class Demo {
	public static void main(String[] args) throws IOException, Exception {
		writeObj();
		//readObj();
	}

	//把文件中的对象信息读取出来-------->对象的反序列化
	public static void readObj() throws  IOException, ClassNotFoundException{

		//找到目标文件
		File file = new File("E:\\obj.txt");

		//建立数据的输入通道
		FileInputStream fileInputStream = new FileInputStream(file);

		//建立对象的输入流对象
		ObjectInputStream objectInputStream = new ObjectInputStream(fileInputStream);

		//读取对象信息
		User user = (User) objectInputStream.readObject(); //创建对象肯定要依赖对象所属 的class文件。
		System.out.println("对象的信息："+ user);
	}

	//定义方法把对象的信息写到硬盘上------>对象的序列化。
	public static void writeObj() throws IOException{

		//把user对象的信息持久化存储。
		Address address = new Address("中国","夏畈");
		User user = new User("admin","123",22,address);

		//找到目标文件F
		File file = new File("E:\\obj.txt");

		//建立数据输出流对象
		FileOutputStream fileOutputStream = new FileOutputStream(file);

		//建立对象的输出流对象
		ObjectOutputStream objectOutputStream = new ObjectOutputStream(fileOutputStream);

		//把对象写出
		objectOutputStream.writeObject(user);

		//关闭资源
		objectOutputStream.close();
	}
}
```
### 配置文件类-Properties

Properties（配置文件类）: 主要用于生产配置文件与读取配置文件的信息。

Properties 继承于 Hashtable 表示一个持久的属性集.属性列表中每个键及其对应值都是一个字符串。

Properties要注意的细节：
	1. 如果配置文件的信息一旦使用了`中文`，那么在使用`store`方法生成配置文件的时候只能使用`字符流解决`，如果使用字节流生成配置文件的话，
	默认使用的是iso8859-1码表进行编码存储，这时候会出现乱码。

	2. 如果Properties中的内容发生了变化，一定要重新使用Properties生成配置文件，否则配置文件信息不会发生变化。

```java
//获取JVM的系统属性
import java.util.Properties;
public class ReadJVM {
    public static void main(String[] args) {
        Properties pps = System.getProperties();
        pps.list(System.out);
    }
}
```

```java
import java.io.FileInputStream;
import java.io.FileNotFoundException;
import java.io.FileOutputStream;
import java.io.FileReader;
import java.io.FileWriter;
import java.io.IOException;
import java.util.Map;
import java.util.Map.Entry;
import java.util.Properties;
import java.util.Set;

public class Demo {
	public static void main(String[] args) throws IOException {
		creatProperties();
		
		//readProperties();
	}

	//读取配置文件的信息 
	public static void readProperties() throws IOException{

		//创建Properties对象
		Properties properties = new Properties();

		//加载配置文件信息到Properties中
		properties.load(new FileReader("F:\\persons.properties"));

		//遍历
		Set<Entry<Object, Object>> entrys = properties.entrySet();

		for(Entry<Object, Object> entry : entrys){
			System.out.println("键："+ entry.getKey() +" 值："+ entry.getValue());
		}

		//把修改后的Properties再生成一个配置文件
		properties.setProperty("狗娃", "007");
		properties.store(new FileWriter("E:\\persons.properties"), "hehe");
	}

	//保存配置文件文件的信息。
	public static void creatProperties() throws IOException{

		//创建Properties
		Properties properties = new Properties();

		properties.setProperty("狗娃", "123");
		properties.setProperty("狗剩","234");
		properties.setProperty("铁蛋","345");

		// 遍历Properties
		/*Set<Entry<Object, Object>> entrys = properties.entrySet();
		for(Entry<Object, Object> entry  :entrys){
			System.out.println("键："+ entry.getKey() +" 值："+ entry.getValue());
		}*/
		
		//使用Properties生产配置文件。
		//properties.store(new FileOutputStream("F:\\persons.properties"), "haha"); //第一个参数是一个输出流对象，第二参数是使用一个字符串描述这个配置文件的信息。
		properties.store(new FileWriter("E:\\persons.properties"), "hehe");
		properties.store(new FileWriter("E:\\persons.properties", true), "hehe");  //追加
	}
}
```

```java
import java.io.File;
import java.io.FileInputStream;
import java.io.FileOutputStream;
import java.io.IOException;
import java.util.Properties;

/*
 需求： 使用properties实现本软件只能 运行三次，超过了三次之后就提示购买正版，退jvm.
 */

public class Demo {
	public static void main(String[] args) throws IOException {

		File file = new File("E:\\count.properties");

		if(!file.exists()){
			//如果配置文件不存在，则创建该配置文件
			file.createNewFile();
		}
		
		//创建Properties对象。
		Properties properties = new Properties();

		//把配置文件的信息加载到properties中
		properties.load(new FileInputStream(file));

		FileOutputStream fileOutputStream = new FileOutputStream(file);

		int count = 0; //定义该变量是用于保存软件的运行次数的。

		//读取配置文件的运行次数
		String value = properties.getProperty("count");
		if(value!=null){
			count = Integer.parseInt(value);
		}
		
		//判断使用的次数是否已经达到了三次，
		if(count==3){
			System.out.println("你已经超出了试用次数，请购买正版软件！！");
			System.exit(0);
		}

		count++;
		System.out.println("你已经使用了本软件第"+count+"次");
		properties.setProperty("count",count+"");  //count+"" 转String

		//使用Properties生成一个配置文件
		properties.store(fileOutputStream,"runtime");
	}
}
```



### 打印流-printStream 

打印流可以打印任意类型的数据，而且打印数据之前都会先把数据转换成字符串再进行打印。

```java
import java.io.File;
import java.io.FileOutputStream;
import java.io.IOException;
import java.io.PrintStream;

class Animal{

	String name;

	String color;

	public Animal(String name,String color){
		this.name = name;
		this.color = color;
	}
	@Override
	public String toString() {
		return "名字："+this.name+ " 颜色："+ this.color;
	}
}

public class Demo {
	public static void main(String[] args) throws IOException {

		/*FileOutputStream fileOutputStream = new FileOutputStream("F:\\a.txt");
		fileOutputStream.write("97".getBytes());
		fileOutputStream.close();*/
		
		
		//打印流可以打印任何类型的数据，而且打印数据之前都会先把数据转换成字符串再进行打印。
		File file = new  File("E:\\a.txt");
		//创建一个打印流
		PrintStream printStream = new PrintStream(file);
		
		printStream.println(97);
		printStream.println(3.14);
		printStream.println('a');
		printStream.println(true);
		Animal a = new Animal("老鼠", "黑色");
		printStream.println(a);
		
		//默认标准的输出流就是向控制台输出的，
		System.setOut(printStream); //重新设置了标准的输出流对象
		System.out.println("哈哈，猜猜我在哪里！！");
	}
}
```

```java
import java.io.File;
import java.io.FileOutputStream;
import java.io.IOException;
import java.io.PrintStream;

//收集异常的日志信息。
public class Demo {
	public static void main(String[] args) throws IOException {
		
		File logFile = new File("E:\\2020年7月25日.log");

		PrintStream logPrintStream = new PrintStream( new FileOutputStream(logFile,true) );

		try{
			int c = 4/0;
			System.out.println("c="+c);
			int[] arr = null;
			System.out.println(arr.length);
		}catch(Exception e){
			e.printStackTrace(logPrintStream);
		}
	}
}
```

### 编码与解码

编码： 把看得懂的字符变成看不懂码值这个过程我们称作为编码。
 
解码： 把码值查找对应的字符，我们把这个过程称作为解码。

注意： 以后编码与解码一般我们都使用统一的码表。否则非常容易出乱码。

```java
import java.io.UnsupportedEncodingException;
import java.util.Arrays;

public class Demo {
	public static void main(String[] args) throws Exception {

		/*
		String str = "中国";
		byte[] buf = str.getBytes("utf-8");// 平台默认的编码表是gbk编码表。  编码
		System.out.println("数组的元素："+Arrays.toString(buf)); //
		
		str = new String(buf,"utf-8");  //默认使用了gbk码表去解码。 
		System.out.println("解码后的字符串："+ str);		
		*/
		
		/*String str = "a中国"; // ,0, 97, 78, 45, 86, -3
		byte[] buf = str.getBytes("unicode");  //编码与解码的时候指定 的码表是unicode实际上就是用了utf-16.
		System.out.println("数组的内容："+ Arrays.toString(buf));
		*/
		
		
		String str = "大家好";
		byte[] buf = str.getBytes(); //使用gbk进行编码
		System.out.println("字节数组："+ Arrays.toString(buf));  // -76, -13, -68, -46, -70, -61
		
		str = new String(buf,"iso8859-1");

		// 出现乱码之后都可以被还原吗？
		byte[] buf2 = str.getBytes("iso8859-1");

		str = new String(buf2,"gbk"); 

		System.out.println(str);
	}
}
```


### 转换流

输入字节流的转换流：
	`InputStreamReader` 是字节流通向字符流的桥

输出字节流的转换流：
	`OutputStreamWriter`   可以把输出字节流转换成输出字符流 。  
 	
转换流的作用：
	1. 如果目前所 获取到的是一个字节流需要转换字符流使用，这时候就可以使用转换流。  字节流----> 字符流
 	1. 使用转换流可以指定编码表进行读写文件。

```java
import java.io.BufferedReader;
import java.io.File;
import java.io.FileInputStream;
import java.io.FileOutputStream;
import java.io.FileReader;
import java.io.FileWriter;
import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.io.OutputStream;
import java.io.OutputStreamWriter;


public class Demo {
	public static void main(String[] args) throws IOException {

		//readTest();
		//writeTest();'
		
		//writeTest2();
		readTest2();
	}

	//使用输入字节流的转换流指定码表进行读取文件数据
	public static void readTest2() throws IOException{

		File file = new File("F:\\a.txt");

		FileInputStream fileInputStream = new FileInputStream(file);

		//创建字节流的转换流并且指定码表进行读取
		InputStreamReader inputStreamReader = new InputStreamReader(fileInputStream,"utf-8");

		char[] buf = new char[1024];
		int length = 0;

		while((length = inputStreamReader.read(buf))!=-1){
			System.out.println(new String(buf,0,length));
		}
	}

	//使用输出字节流的转换流指定码表写出数据
	public static void writeTest2() throws IOException{

		File file = new File("E:\\a.txt");

		//建立数据的输出通道
		FileOutputStream fileOutputStream = new FileOutputStream(file);

		//把输出字节流转换成字符流并且指定编码表。
		OutputStreamWriter outputStreamWriter = new OutputStreamWriter(fileOutputStream, "utf-8");

		outputStreamWriter.write("中国好啊");

		//关闭资源
		outputStreamWriter.close();
	}

	public static void writeTest() throws IOException{

		File file = new File("E:\\a.txt");

		FileOutputStream fileOutputStream = new FileOutputStream(file);

		//把输出字节流转换成输出字符流。
		OutputStreamWriter outputStreamWriter = new OutputStreamWriter(fileOutputStream);
		
		outputStreamWriter.write("大家好");
		outputStreamWriter.close();
	}

	public static void readTest() throws IOException{

		InputStream in = System.in; //获取了标准的输入流。
		//System.out.println("读到 的字符："+ (char)in.read());  //read()一次只能读取一个字节。

		//需要把字节流转换成字符流。
		InputStreamReader inputStreamReader = new InputStreamReader(in);

		//使用字符流的缓冲类
		BufferedReader bufferedReader = new BufferedReader(inputStreamReader);

		String line = null;

		while((line = bufferedReader.readLine())!=null){
			System.out.println("内容："+ line);
		}
	}
}
```

### 文件遍历-递归

递归：函数的自身调用函数的自身。

递归的使用前提： 必须要有条件的情况下调用。

```java
//需求： 算出5的阶乘。 5! = 5*4 *3 * 2*1
public class Demo {

	public static void main(String[] args) {
		int result = print(5);
		System.out.println("结果："+ result);
	}

	public static int print(int num){
		if(num==1){
			return 1;
		}else{
			return num*print(num-1);
		}
	}

	public static int test(int num){
		int result  =  1;
		while(num>0){
			result = result*num;
			num--;
		}
		return result;
	}
}
```

```java
import java.io.File;

/*
1：列出一个文件夹的子孙文件与目录。
 
2，列出指定目录中所有的子孙文件与子孙目录名，要求名称前面要有相应数量的空格：
		第一级前面有0个，第二级前面有1个，第三级前面有2个...，以此类推。

3，列出指定目录中所有的子孙文件与子孙目录名，要求要是树状结构，效果如下所示：
		|--src
		|   |--cn
		|   |   |--itcast
		|   |   |   |--a_helloworld
		|   |   |   |   |--HelloWorld.java
		|   |   |   |--b_for
		|   |   |   |   |--ForTest.java
		|   |   |   |--c_api
		|   |   |   |   |--Student.java
		|--bin
		|   |--cn
		|   |   |--itcast
		|   |   |   |--i_exception
		|   |   |   |   |--ExceptionTest.class
		|   |   |   |--h_linecount
		|   |   |   |   |--LineCounter3.class
		|   |   |   |   |--LineCounter2.class
		|   |   |   |   |--LineCounter.class
		|--lib
		|   |--commons-io.jar

需求4: 删除一个非空的文件夹。 
 */

public class Demo {
	public static void main(String[] args) {

	/*	File dir = new File("F:\\1208project\\day22");
		listFiles3(dir,"|--");*/
		
		File dir = new File("F:\\aa");
		deleteDir(dir);
	}

	//删除了一个非空的目录
	public static void deleteDir(File dir){ // bb
		File[] files = dir.listFiles(); //列出了所有的子文件
		for(File file : files){
			if(file.isFile()){
				file.delete();
			}else if(file.isDirectory()){
				deleteDir(file);
			}
		}
		dir.delete();
	}

	public static void listFiles3(File dir,String space){ 
		File[] files = dir.listFiles(); //列出所有 的子文件
		for(File file : files){
			if(file.isFile()){
				System.out.println(space+file.getName());
			}else if(file.isDirectory()){
				System.out.println(space+file.getName());
				listFiles3(file,"|   "+space);
			}
		}
	}

	//列出一个文件夹的子孙文件与目录。
	public static void listFiles2(File dir,String space){ //space 存储的是空格
		File[] files = dir.listFiles(); //列出所有 的子文件
		for(File file : files){
			if(file.isFile()){
				System.out.println(space+file.getName());
			}else if(file.isDirectory()){
				System.out.println(space+file.getName());
				listFiles2(file,"  "+space);
			}
		}
	}

	//列出一个文件夹的子孙文件与目录。
	public static void listFiles1(File dir){
		File[] files = dir.listFiles(); //列出所有 的子文件
		for(File file : files){
			if(file.isFile()){
				System.out.println("文件名："+file.getName());
			}else if(file.isDirectory()){
				System.out.println("文件夹："+file.getName());
				listFiles1(file);
			}
		}
	}
}
```

## 网络编程

网络编程： 网络编程主要用于解决计算机与计算机（手机、平板..）之间的数据传输问题。

网络编程: 不需要基于html页面就可以达到数据之间的传输。 比如： feiQ , QQ , 微信....

网页编程： 就是要基于html页面的基础上进行数据的交互的。 比如： oa(办公自动化)、 高考的报告系统...

计算机网络： 分布在不同地域的计算机通过外部设备链接起来达到了消息互通、资源共享的效果就称作为一个计算机网络。

网络通讯的三要素：
	1. `IP`
	2. `端口号`。
	3. `协议`.
	
192.168.10.1
IP地址：IP地址的本质就是一个由`32位的二进制数据`组成的数据。 后来别人为了方便我们记忆IP地址，就把IP地址切成了4份，每份8bit.   2^8 = 0~255
      00000000-00000000-00000000-00000000

`IP地址 = 网络号+ 主机号`。

IP地址的分类：
	A类地址 = 一个网络号 + 三个主机号     2^24   政府单位
	B类地址 =  两个网络号+ 两个主机号   2^16 事业单位（学校、银行..）
	C类地址= 三个网络号+ 一个主机号  2^8    私人使用..

### InetAddress(IP类)

常用的方法：
	getLocalHost(); 获取本机的IP地址对象
	getByName("IP或者主机名") 根据一个IP地址的字符串形式或者是一个主机名生成一个IP地址对象。 (用于获取别人的IP地址对象)

	getHostAddress()  返回一个IP地址的字符串表示形式。
	getHostName()  返回计算机的主机名。

端口号是没有类描述的。
	端口号的范围： 0~65535
		从0到1023，系统紧密绑定于一些服务。 
		1024~65535  我们可以使用....

网络通讯的协议：
	`udp`通讯协议
	`tcp`通讯协议。

```java
import java.net.InetAddress;
import java.net.UnknownHostException;

public class Demo {
	public static void main(String[] args) throws UnknownHostException {
		
		//getLocalHost 获取本机的IP地址对象
		InetAddress address = InetAddress.getLocalHost();
		System.out.println("IP地址："+address.getHostAddress());  //IP地址：192.168.183.1
		System.out.println("主机名："+address.getHostName());   //主机名：ixfosa
		

		//获取别人机器的IP地址对象。
		//可以根据一个IP地址的字符串形式或者是一个主机名生成一个IP地址对象。
		InetAddress address2 = InetAddress.getByName("ixfosa");
		System.out.println("IP地址："+address2.getHostAddress());  //IP地址：192.168.183.1
		System.out.println("主机名："+address2.getHostName()); //主机名：ixfosa


		InetAddress[]  arr = InetAddress.getAllByName("www.baidu.com");//域名*/
		System.out.println("主机名："+arr[0].getHostName()); //主机名：www.baidu.com
		System.out.println("IP地址："+arr[0].getHostAddress()); //IP地址：14.215.177.39

		System.out.println();

		for (InetAddress inetAddress : arr) {
			System.out.println("主机名："+inetAddress.getHostName()); //主机名：www.baidu.com
			System.out.println("主机名："+inetAddress.getHostAddress()); 
		}
	}
}
```

### tcp
UDP通讯协议的特点：
	1. 将数据极封装为数据包，面向无连接。
	2. 每个数据包大小限制在64K中
	3. 因为无连接，所以不可靠
	4. 因为不需要建立连接，所以速度快
	5. udp 通讯是不分服务端与客户端的，只分发送端与接收端。

TCP通讯协议特点：
	1. tcp是基于IO流进行数据 的传输 的，面向连接。
 	2. tcp进行数据传输的时候是没有大小限制的。
 	3. tcp是面向连接，通过三次握手的机制保证数据的完整性。 可靠协议。
 	4. tcp是面向连接的，所以速度慢。
 	5. tcp是区分客户端与服务端 的。
 	
 	比如： 打电话、 QQ\feiQ的文件传输、 迅雷下载.... 
 	
tcp协议下的Socket：
	Socket(客户端) , tcp的客户端一旦启动马上要与服务端进行连接。
	ServerSocket(服务端类)



```java
import java.io.IOException;
import java.io.InputStream;
import java.io.OutputStream;
import java.net.InetAddress;
import java.net.Socket;
import java.net.UnknownHostException;

/*
tcp的客户端使用步骤：
	1. 建立tcp的客户端服务。
	2. 获取到对应的流对象。
	3. 写出或读取数据
	4. 关闭资源。
*/

//tcp客户端
public class Clinet {
	
	public static void main(String[] args) throws IOException{

		//建立tcp的服务
		Socket socket  = new Socket(InetAddress.getLocalHost(),9090);

		//获取到Socket的输出流对象
		OutputStream outputStream = socket.getOutputStream();

		//利用输出流对象把数据写出即可。
		outputStream.write("服务端你好".getBytes());
		
		//获取到输入流对象，读取服务端回送的数据。
		InputStream inputStream = socket.getInputStream();

		byte[] buf = new byte[1024];
		int length = inputStream.read(buf);

		System.out.println("客户端接收到的数据："+ new String(buf,0,length));
		
		//关闭资源
		socket.close();		
	}
}

/*-----------------------------------------------------------------------------*/

import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.io.OutputStream;
import java.net.ServerSocket;
import java.net.Socket;

/*
java.net.BindException:   端口被占用。

ServerSocket的使用 步骤
	1. 建立tcp服务端 的服务。
	2. 接受客户端的连接产生一个Socket.
	3. 获取对应的流对象读取或者写出数据。
	4. 关闭资源。
*/

//tcp的服务端
public class Server {

	public static void main(String[] args) throws Exception {

		//建立Tcp的服务端,并且监听一个端口。
		ServerSocket serverSocket = new ServerSocket(9090);

		//接受客户端的连接
		Socket socket  =  serverSocket.accept(); //accept()  接受客户端的连接 该方法也是一个阻塞型的方法，没有客户端与其连接时，会一直等待下去。
		//获取输入流对象，读取客户端发送的内容。
		InputStream inputStream = socket.getInputStream();

		byte[] buf = new byte[1024];
		int length = 0;

		length = inputStream.read(buf);
		System.out.println("服务端接收："+ new String(buf,0,length));
		
		//获取socket输出流对象，想客户端发送数据
		OutputStream outputStream = socket.getOutputStream();
		outputStream.write("客户端你好啊！".getBytes());
		
		//关闭资源
		serverSocket.close();
	}
}
```

```java
import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.io.OutputStreamWriter;
import java.net.InetAddress;
import java.net.Socket;

/*
 需求： 客户端与服务端一问一答聊天。
 
 1.如果使用BuffrerdReader的readline方法一定要加上\r\n才把数据写出。
 2.使用字符流一定要调用flush方法数据才会写出。
 */

//聊天的客户端
public class ChatClient {
	public static void main(String[] args) throws IOException {

		//建立tcp的客户端服务
		Socket socket = new Socket(InetAddress.getLocalHost(),9090);

		//获取socket的输出流对象。
		OutputStreamWriter socketOut =	new OutputStreamWriter(socket.getOutputStream());

		//获取socket的输入流对象
		BufferedReader socketReader = new BufferedReader(new InputStreamReader(socket.getInputStream()));
		
		//获取键盘的输入流对象，读取数据
		BufferedReader keyReader = new BufferedReader(new InputStreamReader(System.in));
		
		String line = null;

		//不断的读取键盘录入的数据，然后把数据写出
		while((line = keyReader.readLine())!=null){

			socketOut.write(line+"\r\n");

			//刷新
			socketOut.flush();

			//读取服务端回送的数据
			line = socketReader.readLine();
			System.out.println("服务端回送的数据是："+line);
		}
		//关闭资源
		socket.close();
	}
}

/*-----------------------------------------------------------------------------*/

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.io.OutputStreamWriter;
import java.net.ServerSocket;
import java.net.Socket;

/*
 聊天的服务端
 */
public class ChatServer {
	public static void main(String[] args) throws IOException {

		//建立tcp的服务端
		ServerSocket serverSocket = new ServerSocket(9090);

		//接受客户端的连接，产生一个SOcket
		Socket socket = serverSocket.accept();

		//获取到Socket的输入流对象
		BufferedReader socketReader = new BufferedReader(new InputStreamReader(socket.getInputStream()));
		
		//获取到Socket输出流对象
		OutputStreamWriter socketOut =  new OutputStreamWriter(socket.getOutputStream());
		
		//获取键盘的输入流对象
		BufferedReader keyReader = new BufferedReader(new InputStreamReader(System.in));
		
		//读取客户端的数据
		String line = null;

		while((line = socketReader.readLine())!=null){

			System.out.println("服务端接收到的数据："+ line);

			System.out.println("请输入回送给客户端的数据：");

			line = keyReader.readLine();
			socketOut.write(line+"\r\n");

			socketOut.flush();
		}
		//关闭资源
		serverSocket.close();
	}
}
```

```java
import java.io.IOException;
import java.io.OutputStream;
import java.net.ServerSocket;
import java.net.Socket;

//模拟Tomcat服务器
public class TomcatDemo extends Thread {

	Socket socket;

	public TomcatDemo(Socket socket){
		this.socket = socket;
	}

	public void run() {
		try {
			//获取socket的输出流对象
			OutputStream outputStream = socket.getOutputStream();
			//把数据写到浏览器上
			outputStream.write("<html><head><title>aaa</title></head><body>你好啊浏览器</body></html>".getBytes());
			socket.close();
		} catch (IOException e) {
			e.printStackTrace();
		}
	}

	public static void main(String[] args) throws IOException {

		//建立tcp的服务端
		ServerSocket serverSocket = new ServerSocket(9090);

		//不断的接受客户端的连接
		while(true){
			Socket socket = serverSocket.accept();
			new TomcatDemo(socket).start();
		}
	}
}
```

### udp
在java中网络通讯业称作为Socket(插座)通讯，要求通讯 的两台器都必须要安装Socket。

不同的协议就有不同 的插座（Socket）

**UDP通讯协议的特点**：
	1. 将数据极封装为数据包，面向无连接。
	2. 每个数据包大小限制在64K中
	3.因为无连接，所以不可靠
	4. 因为不需要建立连接，所以速度快
	5.udp 通讯是不分服务端与客户端的，只分发送端与接收端。

	比如： 物管的对讲机, 飞Q聊天、 游戏...

udp协议下的Socket:
	`DatagramSocket`(udp插座服务)
	`DatagramPacket`(数据包类)
		DatagramPacket(buf, length, address, port)
		buf: 发送的数据内容
		length : 发送数据内容的大小。
		address : 发送的目的IP地址对象
		port : 端口号。

发送端的使用步骤：
	1. 建立udp的服务。
	2. 准备数据，把数据封装到数据包中发送。 发送端的数据包要带上ip地址与端口号。
	3. 调用udp的服务，发送数据。
	4. 关闭资源。

```java
import java.io.IOException;
import java.net.DatagramPacket;
import java.net.DatagramSocket;
import java.net.InetAddress;
import java.net.SocketException;

//发送端
public class Sender {
	public static void main(String[] args) throws IOException {

		//建立udp的服务
		DatagramSocket datagramSocket = new DatagramSocket();

		//准备数据，把数据封装到数据包中。
		String data = "这个是我第一个udp的例子..";

		//创建了一个数据包
		DatagramPacket packet = new DatagramPacket(data.getBytes(), data.getBytes().length,InetAddress.getLocalHost() , 9090);

		//调用udp的服务发送数据包
		datagramSocket.send(packet);

		//关闭资源 ---实际上就是释放占用的端口号
		datagramSocket.close();
	}

}

/*-----------------------------------------------------------------------------*/

import java.io.IOException;
import java.net.DatagramPacket;
import java.net.DatagramSocket;

//接收端
/*
 接收端的使用步骤
 	1. 建立udp的服务
 	2. 准备空 的数据 包接收数据。
 	3. 调用udp的服务接收数据。
 	4. 关闭资源
 
 */
public class Receive {
	public static void main(String[] args) throws IOException {

		//建立udp的服务 ，并且要监听一个端口。
		DatagramSocket  socket = new DatagramSocket(9090);
		
		//准备空的数据包用于存放数据。
		byte[] buf = new byte[1024];
		DatagramPacket datagramPacket = new DatagramPacket(buf, buf.length); // 1024

		//调用udp的服务接收数据
		socket.receive(datagramPacket);//receive是一个阻塞型的方法,没有接收到数据包之前会一直等待,数据实际上就是存储到了byte的自己数组中了。

		System.out.println("接收端接收到的数据："+ new String(buf,0,datagramPacket.getLength())); // getLength() 获取数据包存储了几个字节。

		//关闭资源
		socket.close();
	}
}
```

```java
import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.net.DatagramPacket;
import java.net.DatagramSocket;
import java.net.InetAddress;
import java.net.SocketException;
import java.net.UnknownHostException;

//群聊发送端
public class ChatSender extends Thread {

	@Override
	public void run() {
		try {

			//建立udp的服务
			DatagramSocket socket = new DatagramSocket();

			//准备数据，把数据封装到数据包中发送
			BufferedReader keyReader = new BufferedReader(new InputStreamReader(System.in));

			String line = null;
			DatagramPacket packet  = null;

			while((line = keyReader.readLine())!=null){
				//把数据封装 到数据数据包中，然后发送数据。
				packet = new DatagramPacket(line.getBytes(), line.getBytes().length, InetAddress.getByName("192.168.1.255"), 9090);

				//把数据发送出去
				socket.send(packet);
			}

			//关闭 资源
			socket.close();
		} catch (IOException e) {
			e.printStackTrace();
		}
	}
}

/*-----------------------------------------------------------------------------*/

import java.io.IOException;
import java.net.DatagramPacket;
import java.net.DatagramSocket;
import java.net.SocketException;

//群聊接收端
public class ChatReceive extends Thread {
	
	@Override
	public void run() {
		try {

			//建立udp的服务,要监听一个端口
			DatagramSocket socket = new DatagramSocket(9090);

			//准备空的数据包存储数据
			byte[] buf = new byte[1024];

			DatagramPacket packet = new DatagramPacket(buf, buf.length);

			boolean flag = true;

			while(flag){
				socket.receive(packet);
				// packet.getAddress() 获取对方数据 包的IP地址对象。
				System.out.println(packet.getAddress().getHostAddress()+"说:"+new String(buf,0,packet.getLength()));
			}

			//关闭资源
			socket.close();
		
		}catch (IOException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
		}
	}
}

/*-----------------------------------------------------------------------------*/

public class ChatMain {
	public static void main(String[] args) {

		ChatReceive chatReceive = new ChatReceive();
		chatReceive.start();
		
		ChatSender chatSender = new ChatSender();
		chatSender.start();
	}
}
```

```java
import java.net.DatagramPacket;
import java.net.DatagramSocket;
import java.net.InetAddress;
import java.net.UnknownHostException;

/*
udp是一个不可靠（数据包可能会丢失）的协议

什么情况下数据包会出现丢失呢？
	1.带宽不足 。
	2.cpu的处理能力不足。
 */

public class SafeSender {
	public static void main(String[] args) throws Exception {
		//建立udp的服务
		DatagramSocket socket = new DatagramSocket();
		//准备数据，数据封装到数据中发送
		DatagramPacket packet = null;

		for(int i =  0 ; i< 10; i++){  //连续发送10个数据包
			String data =i +"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa";

			packet = new DatagramPacket(data.getBytes(), data.getBytes().length, InetAddress.getLocalHost(), 9090);
			//发送数据包
			socket.send(packet);
		}
		//关闭资源
		socket.close();
	}
}

/*-----------------------------------------------------------------------------*/

import java.io.IOException;
import java.net.DatagramPacket;
import java.net.DatagramSocket;

public class SafeReceive {
	public static void main(String[] args) throws IOException, Exception {

		//建立udp的服务
		DatagramSocket socket = new DatagramSocket(9090);
		
		//建立空的数据包存储数据
		byte[] buf = new byte[1024];
		DatagramPacket packet = new DatagramPacket(buf, buf.length);
		
		//不断接收数据包
		while(true){
			socket.receive(packet);
			System.out.println(new String(buf,0,packet.getLength()));
			Thread.sleep(10);
		}
	}
}
```


### 例子

```java
import java.io.IOException;
import java.net.DatagramPacket;
import java.net.DatagramSocket;
import java.net.InetAddress;

/*
 每个网络程序都有自己所处理的特定格式数据,如果接收到的数据不符合指定的格式，那么就会被当成垃圾数据丢弃。(加密..)

 飞Q接收的数据格式：
 version:time :sender : ip: flag:content ;
 版本号          时间         发送人   :IP： 发送的标识符(32): 真正的内容;

 在udp协议中，有一个IP地址称作为广播地址，广播地址就是主机号为255地址。

 给广播IP地址发送消息的时候，在同一个网络段的机器都可以接收 到信息。
 192.168.15.255

 */

//使用udp协议给飞Q发送消息。
public class FeiQDemo {
	public static void main(String[] args) throws IOException {

		// 建立udp的服务
		DatagramSocket socket = new DatagramSocket();

		// 准备数据，把数据封装到数据包中
		String data = getData("feiQ你好！");

		DatagramPacket packet = new DatagramPacket(data.getBytes(),
				data.getBytes().length,
				InetAddress.getByName("192.168.15.255"), 2425);

		// 发送数据
		socket.send(packet);
		// 关闭资源
		socket.close();

	}

	// 把数据拼接成指定格式的数据
	public static String getData(String content) {
		StringBuilder sb = new StringBuilder();
		sb.append("1.0:");
		sb.append(System.currentTimeMillis() + ":");
		sb.append("ixfosa:");
		sb.append("192.168.10.1:");
		sb.append("32:");
		sb.append(content);

		return sb.toString();
	}
}
```

```java
import java.io.FileInputStream;
import java.io.FileNotFoundException;
import java.io.IOException;
import java.io.OutputStream;
import java.net.ServerSocket;
import java.net.Socket;
import java.util.HashSet;

/*
1. 编写一个服务端可以给多个客户端发送图片。 （多线程）
*/
public class ImageServer extends Thread {

	Socket socket ;

	//使用该集合是用于存储ip地址的。
	static HashSet<String> ips = new HashSet<String>();

	public  ImageServer(Socket socket) {
		this.socket = socket;
	}

	@Override
	public void run() {
		try {

			//获取到socket输出流对象
			OutputStream outputStream = socket.getOutputStream();

			//获取图片的输入流对象
			FileInputStream fileInputStream = new FileInputStream("E:\\3.jpg");

			//读取图片数据，把数据写出
			byte[] buf = new byte[1024];
			int length = 0 ;

			while((length = fileInputStream.read(buf))!=-1){

				outputStream.write(buf,0,length);
			}

			String ip = socket.getInetAddress().getHostAddress();   // socket.getInetAddress() 获取对方的IP地址
			if(ips.add(ip)){
				System.out.println("恭喜"+ip+"同学成功下载，当前下载的人数是："+ ips.size());
			}

			//关闭资源
			fileInputStream.close();
			socket.close();
		}catch (IOException e) {

		}
	}

	public static void main(String[] args) throws IOException {

		//建立tcp的服务 ,并且要监听一个端口
		ServerSocket serverSocket  = new ServerSocket(9090);
		while(true){

			//接受用户的链接。
			Socket socket = serverSocket.accept();
			new ImageServer(socket).start();
		}
	}
}

/*-----------------------------------------------------------------------------*/

import java.io.FileOutputStream;
import java.io.IOException;
import java.io.InputStream;
import java.net.InetAddress;
import java.net.Socket;
import java.net.UnknownHostException;

//下载图片的客户端
public class ImageClient {
	public static void main(String[] args) throws Exception{

		//建立tcp的服务
		Socket socket = new Socket(InetAddress.getLocalHost(),9090);

		//获取socket的输入流对象
		InputStream inputStream = socket.getInputStream();

		//获取文件的输出流对象
		FileOutputStream fileOutputStream = new FileOutputStream("D:\\3.jpg");

		//边读边写
		byte[] buf = new byte[1024];
		int length = 0 ;

		while((length = inputStream.read(buf))!=-1){
			fileOutputStream.write(buf,0,length);
		}

		//关闭资源
		fileOutputStream.close();
		socket.close();
	}
}
```

```java
import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.io.OutputStreamWriter;
import java.net.InetAddress;
import java.net.Socket;


/*
2. 实现登陆与注册 功能。  
	客户端与服务端连接的时候，就要提示客户端请选择功能。

	客户端注册的时候，用户名与密码都是发送给服务端 的，服务端需要把数据保存到服务端的文件上。

	登陆： 登陆的时候客户端输入用户名与密码发送给服务端，服务端需要校验，返回结果给客户端。
*/

public class LoginClinet {
	public static void main(String[] args) throws IOException {		

		Socket socket = new Socket(InetAddress.getLocalHost(),9090);

		//获取socket的输出流对象
		OutputStreamWriter  socketOut = new OutputStreamWriter(socket.getOutputStream());
		
		//获取到socket的输入流对象
		BufferedReader socketReader = new BufferedReader(new InputStreamReader(socket.getInputStream()));
		
		//获取到键盘的输入流对象
		BufferedReader keyReader = new BufferedReader(new InputStreamReader(System.in));

		while(true){

			System.out.println("请选择功能： A(登陆)  B(注册)");

			String option = keyReader.readLine();

			if("a".equalsIgnoreCase(option)){

				getInfo(socketOut, keyReader, option);

				//读取服务器反馈的信息
				String line = socketReader.readLine();

				System.out.println(line);

			}else if("b".equalsIgnoreCase(option)){

				getInfo(socketOut, keyReader, option);

				//读取服务器反馈的信息
				String line = socketReader.readLine();

				System.out.println(line);
			}
		}
	}

	public static void getInfo(OutputStreamWriter  socketOut,BufferedReader keyReader, String option)
			throws IOException {

		System.out.println("请输入用户名:");
		String userName = keyReader.readLine();

		System.out.println("请输入密码：");
		String password = keyReader.readLine();

		String info = option +" "+userName+" "+password+"\r\n";

		socketOut.write(info);
		socketOut.flush();
	}
}

/*-----------------------------------------------------------------------------*/

import java.io.BufferedReader;
import java.io.File;
import java.io.FileReader;
import java.io.FileWriter;
import java.io.IOException;
import java.io.InputStreamReader;
import java.io.OutputStreamWriter;
import java.net.ServerSocket;
import java.net.Socket;
import java.util.Properties;

public class LoginServer extends Thread {

	Socket socket;

	static File file = new File("E:\\users.properties");

	public LoginServer(Socket socket) {
		this.socket = socket;
	}

	static {
		try {
			if (!file.exists()) {
				file.createNewFile();
			}
		} catch (IOException e) {
			e.printStackTrace();
		}
	}

	@Override
	public void run() {
	while(true){
			try {

				// 获取socket的输入流对象
				BufferedReader bufferedReader = new BufferedReader(
						new InputStreamReader(socket.getInputStream()));

				// 获取socket的输出流对象
				OutputStreamWriter socketOut = new OutputStreamWriter(
						socket.getOutputStream());

				// 读取客户端输入的信息
				String info = bufferedReader.readLine();

				String[] datas = info.split(" ");

				// 获取到用户 的选择功能
				String option = datas[0];
				String userName = datas[1];
				String password = datas[2];

				if ("a".equalsIgnoreCase(option)) {
					// 登陆
					Properties properties = new Properties();
					// 加载配置文件
					properties.load(new FileReader(file));
					if (properties.containsKey(userName)) {

						String tempPass = properties.getProperty(userName);
						if (password.equals(tempPass)) {
							socketOut.write("欢迎" + userName + "登陆成功\r\n");

						} else {
							socketOut.write("密码错误\r\n");
						}

					} else {
						socketOut.write("用户名不存在，请重新输入...\r\n");
					}

					socketOut.flush();

				} else if ("b".equalsIgnoreCase(option)) {

					// 创建一个配置文件类
					Properties properties = new Properties();

					//加载原来的配置文件
					properties.load(new FileReader(file));

					if (!properties.containsKey(userName)) {

						// 不存在该用户名
						properties.setProperty(userName, password);

						// 生成一个配置文件
						properties.store(new FileWriter(file), "users");
						socketOut.write("注册成功..\r\n");

					} else {

						// 存在用户名
						socketOut.write("用户名已经被注册，请重新输入\r\n");
					}
					socketOut.flush();
				}
			} catch (Exception e) {
				e.printStackTrace();
			}
		}

	}

	public static void main(String[] args) throws IOException {
		ServerSocket serverSocket = new ServerSocket(9090);
		while (true) {
			Socket socket = serverSocket.accept();
			new LoginServer(socket).start();
		}
	}
}
```

## GUI

软件的交互方式：
	1. 控制台的交互方式、
	2. 图形化界面的交互方式 。

java使用到的图形类主要在java.awt 与javax.swing包中。

java.awt 与 javax.swing包的区别：
	`java.awt`中使用的图形类都是依赖于系统的图形库的。
	`javax.swing`包使用到的图形类都是sun自己实现，不需要依赖系统的图形库。

疑问： 既然swing包中的图形类已经取代awt包的图形类，为什么不删除awt呢？

在java中所有的图形类都被称作组件类。

组件的类别：
----------| 容器组件
----------| 非容器组件  

```java
import java.awt.Dimension;
import java.awt.Toolkit;

import javax.swing.JFrame;

public class Demo {
	public static void main(String[] args) {

		JFrame frame = new JFrame("这个是我第一个图形化界面的例子");

		//设置窗体的大小
		//frame.setSize(300,400);
		
		//设置窗体(左上角)出现的位置
		//frame.setBounds((1366-300)/2, (768-400)/2, 300,400); // 第一个参数是左上角的x轴坐标， 第二参数是左上角y的坐标。 第三个窗体宽， 第四窗体的高。

		initFrame(frame, 300,400);

		frame.setVisible(true); //setVisible 设置窗体的可见性。

		//设置窗体关闭事件
		frame.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
	}

	//获取屏幕的分辨率   设置窗体在屏幕的居中位置。
	public static void initFrame(JFrame frame,int width , int height){
		
		Toolkit toolkit = Toolkit.getDefaultToolkit(); //获取一个与系统相关工具类对象

		//获取屏幕的分辨率
		Dimension d = toolkit.getScreenSize();
		int x = (int) d.getWidth();
		int y = (int) d.getHeight();
		frame.setBounds((x-width)/2, (y-height)/2, width, height);
	}
}
```

### FrameUtil
```java
package cn.ixfosa.util;

import java.awt.Dimension;
import java.awt.Toolkit;

import javax.swing.JFrame;

//初始化窗体的工具类
public class FrameUtil {

	public static void initFrame(JFrame frame,int width , int height){

		Toolkit toolkit = Toolkit.getDefaultToolkit(); //获取一个与系统相关工具类对象

		//获取屏幕的分辨率
		Dimension d = toolkit.getScreenSize();
		
		int x = (int) d.getWidth();
		int y = (int) d.getHeight();

		frame.setBounds((x-width)/2, (y-height)/2, width, height);

		frame.setVisible(true);

		frame.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
	}
}

```

### 容器组件

 #### 对话框类(Dialog):

 	JDialog(Dialog owner, String title, boolean modal) 
 	
 	owner: 所有者
 	
 	title : 标题
 	
 	modal : modal

JOptionPane(对话框)
	消息对话框
	警告对话框
	错误对话框
	输入对话框
	确认对话框

```java
import java.awt.Dialog;

import javax.swing.JDialog;
import javax.swing.JFrame;
import javax.swing.JOptionPane;

import cn.ixfosa.util.FrameUtil;


public class Demo{
	public static void main(String[] args) {

		/*JFrame frame = new JFrame("窗体");
		//创建对话框
		JDialog dialog = new JDialog(frame, "对话框",true);

		//使用我自定义的 窗体工具类
		FrameUtil.initFrame(frame, 300, 400);
		
		dialog.setBounds(500,300, 100, 200);

		dialog.setVisible(true); //设置对话框的可见性*/
		
		
		JFrame frame = new JFrame("窗体");

		//显示一个对话框
		FrameUtil.initFrame(frame, 300, 400);

		//消息对话框
		/*JOptionPane.showMessageDialog(frame, "明天不用上课", "通知",JOptionPane.INFORMATION_MESSAGE);

		//警告对话框
		JOptionPane.showMessageDialog(frame,"警告某位同学晚上晚自习不要看动漫","警告",JOptionPane.WARNING_MESSAGE);*/

		//错误对话框
		//JOptionPane.showMessageDialog(frame,"扣6分","错误",JOptionPane.ERROR_MESSAGE);
		
		//输入框
		/*String moeny = JOptionPane.showInputDialog("请输入你要给我的金额($)");
		System.out.println("money:"+ moeny);*/
		
		//确认框
		int num = JOptionPane.showConfirmDialog(frame, "你确认要卸载吗？");
		System.out.println(num);
	}
}
```

 #### 文件对话框(FileDialog)

 	FileDialog(Dialog parent, String title, int mode) 
	parent： 对话框的所有者
	tiltle : 对话框的标题
	mode: load(打开) 、 save(保存)

```java
import java.awt.FileDialog;

import javax.swing.JFrame;

import cn.itcast.util.FrameUtil;

public class Demo {
	public static void main(String[] args) {

		JFrame frame = new JFrame("窗体");

		//创建一个文件对话框(初始也是不可见)
		//FileDialog fileDialog = new FileDialog(frame, "请选择打开的文件", FileDialog.LOAD);

		FileDialog fileDialog = new FileDialog(frame,"请选择保存的路径",FileDialog.SAVE);

		FrameUtil.initFrame(frame, 300,400);

		fileDialog.setVisible(true);

		System.out.println("文件所在的目录："+ fileDialog.getDirectory());
		System.out.println("文件名："+ fileDialog.getFile());
	}
}
```

#### 面板(JPanel)

```java
import java.awt.Color;

import javax.swing.JFrame;
import javax.swing.JPanel;

import cn.itcast.util.FrameUtil;

public class Demo {
	public static void main(String[] args) {

		JFrame frame = new JFrame("窗体");

		//创建一个面板
		JPanel panel = new JPanel();

		//设置面板的背景颜色
		panel.setBackground(Color.RED);

		//把面板添加到窗体
		frame.add(panel);

		FrameUtil.initFrame(frame, 400, 300);
	}
}
```

### 非容器组件

```java
import javax.swing.ButtonGroup;
import javax.swing.JCheckBox;
import javax.swing.JComboBox;
import javax.swing.JFrame;
import javax.swing.JLabel;
import javax.swing.JPanel;
import javax.swing.JPasswordField;
import javax.swing.JRadioButton;
import javax.swing.JTextArea;
import javax.swing.JTextField;

import cn.itcast.util.FrameUtil;

/*
 非容器组件:
 */

public class Demo {
	public static void main(String[] args) {

		JFrame frame= new JFrame("注册");

		//创建一个面板
		JPanel panel = new JPanel();
		frame.add(panel);

		//用户名
		JLabel nameLabel = new JLabel("用户名");
		//用户名的输入框
		JTextField nameField = new JTextField(12);
		//把用户名的组件添加到面板上
		panel.add(nameLabel);
		panel.add(nameField);
		
		//密码
		JLabel passLabel= new JLabel("密码");
		//密码框
		JPasswordField passField = new JPasswordField(12);
		//把密码的组件添加到面板
		panel.add(passLabel);
		panel.add(passField);
		
		//性别--单选框
		JLabel sexLabel = new JLabel("性别");
		JRadioButton man = new JRadioButton("男",true);
		JRadioButton woman = new JRadioButton("女");
		//如果是单选框必须要进行分组，同一个组的单选框只能选择其中的一个
		ButtonGroup group = new ButtonGroup();
		group.add(woman);
		group.add(man);
		//把性别组件添加到面板上
		panel.add(sexLabel);
		panel.add(man);
		panel.add(woman);
		
		//来自城市--->下拉框
		JLabel cityLabel = new JLabel("来自的城市");
		Object[]  arr = {"北京","上海","广州","深圳","湛江"};
		JComboBox citys = new JComboBox(arr);
		panel.add(cityLabel);
		panel.add(citys);
		
		//兴趣爱好---->复选框
		JLabel hobitLabel = new JLabel("兴趣爱好:");
		JCheckBox checkBox1 = new JCheckBox("篮球",true);
		JCheckBox checkBox2 = new JCheckBox("java",true);
		JCheckBox checkBox3 = new JCheckBox("javascript");
		JCheckBox checkBox4 = new JCheckBox("android");
		panel.add(hobitLabel);
		panel.add(checkBox1);
		panel.add(checkBox2);
		panel.add(checkBox3);
		panel.add(checkBox4);
		
		
		//个人简介
		JLabel jLabel = new JLabel("个人简介");
		JTextArea area = new JTextArea(20, 15);
		area.setLineWrap(true); //设置自动换行 
		panel.add(jLabel);
		panel.add(area);
		
		FrameUtil.initFrame(frame, 500, 400);
	}
}
```

### 菜单组件

菜单条(MenuBar) 、  菜单（Menu） 、 菜单项(MenuItem)

菜单条可以添加菜单

菜单可以添加菜单项

复选菜单：
	首先菜单添加菜单 ， 菜单添加菜单项。

```java
import java.awt.BorderLayout;

import javax.swing.JFrame;
import javax.swing.JMenu;
import javax.swing.JMenuBar;
import javax.swing.JMenuItem;

import javax.swing.JTextArea;

import cn.itcast.util.FrameUtil;


public class Demo {

	JFrame frame = new JFrame("记事本");

	//菜单条
	JMenuBar bar = new JMenuBar();

	//文件菜单
	JMenu fileMenu = new JMenu("文件");
	JMenu editMenu  = new JMenu("编辑");
	JMenu switchMenu = new JMenu("切换工作目录");

	//菜单项
	JMenuItem openMenu = new JMenuItem("打开");
	JMenuItem saveMenu = new JMenuItem("保存");
 
	JMenuItem aboutMenu = new JMenuItem("关于");
	JMenuItem closeMenu = new JMenuItem("关闭");

	JMenuItem  workMenu1 = new JMenuItem("0910project");
	JMenuItem  workMenu2 = new JMenuItem("1208project");
	JMenuItem  workMenu3 = new JMenuItem("1110project");

	JTextArea area = new JTextArea(20,30);

	public void initNotepad(){

		//菜单添加菜单项目
		fileMenu.add(openMenu);
		fileMenu.add(saveMenu);
		
		editMenu.add(aboutMenu);
		editMenu.add(closeMenu);
		
		//复选菜单
		switchMenu.add(workMenu1);
		switchMenu.add(workMenu2);
		switchMenu.add(workMenu3);

		//菜单添加菜单就是复选菜单
		fileMenu.add(switchMenu);
		
		//菜单条添加菜单
		bar.add(fileMenu);
		bar.add(editMenu);
		
		//添加菜单条
		frame.add(bar,BorderLayout.NORTH);
		frame.add(area);
		FrameUtil.initFrame(frame, 500, 600);
	}

	public static void main(String[] args) {
		new Demo().initNotepad();
	}
}
```

###  布局管理器
 布局管理器:布局管理就是用于指定组件的 摆放位置的。

#### BorderLayout

BorderLayout: 边框布局管理器
	摆放的风格： 上北  、 下南 、 左西、 右东 ， 中 

Borderlayout 要注意的事项：
	1. 使用Borderlayout添加组件的时候，如果没有指定组件的方位，那么默认添加到中间的位置上。
	2. 使用BorderLayout的时候，如果东南西北那个方向没有对应 的组件，那么中间位置的组件就会占据其空缺的位置。
	3. 窗体默认的布局管理器就是Borderlayout.


```java
import java.awt.BorderLayout;

import javax.swing.JButton;
import javax.swing.JFrame;

import cn.itcast.util.FrameUtil;

public class Demo {
	public static void main(String[] args) {

		JFrame frame = new JFrame("边框局部管理器");

		//创建一个边框布局管理器
		BorderLayout borderLayout = new BorderLayout();

		//让borderlayout管理frame窗体。
		frame.setLayout(borderLayout);
		
		frame.add(new JButton("北"),BorderLayout.NORTH);
		frame.add(new JButton("南"),BorderLayout.SOUTH);
		frame.add(new JButton("西"),BorderLayout.WEST);
		frame.add(new JButton("东"),BorderLayout.EAST);
		frame.add(new JButton("中"),BorderLayout.CENTER);

		//初始化窗体
		FrameUtil.initFrame(frame, 300, 300);
	}
}
```

#### FlowLayout
/*
FlowLayout: 流式布局管理器

流式布局管理器要注意的事项：
	1. 流式布局管理器默认情况是居中对齐的。
	2. panel默认的局部管理器就是FlowLayout.

```java
import java.awt.FlowLayout;

import javax.swing.JButton;
import javax.swing.JFrame;
import javax.swing.JPanel;

import cn.itcast.util.FrameUtil;

public class Demo {
	public static void main(String[] args) {

		JFrame frame = new JFrame("窗体");

		//创建面板
		JPanel panel = new JPanel();
		frame.add(panel);

		//创建一个流式布局管理器
		FlowLayout flowLayout = new FlowLayout(FlowLayout.LEFT, 0, 30);// FlowLayout.LEFT 指定对齐的方式。

		//让流式布局管理器管理frame窗体
		panel.setLayout(flowLayout);
		
		panel.add(new JButton("按钮1"));
		panel.add(new JButton("按钮2"));
		panel.add(new JButton("按钮3"));
		panel.add(new JButton("按钮4"));
		
		//初始化窗体
		FrameUtil.initFrame(frame, 300, 300);
	}
}
```

#### GridLayout
 GridLayout : 表格布局管理器

 注意的事项： 如果表格数量不够使用时，默认会多加一列。


```java
import java.awt.GridLayout;

import javax.swing.JButton;
import javax.swing.JFrame;

import cn.itcast.util.FrameUtil;

public class Demo {
	public static void main(String[] args) {

		JFrame frame = new JFrame("计算器");

		//创建表格布局管理器
		GridLayout gridLayout = new GridLayout(4, 4, 1, 2);
		
		//让窗体交给表格布局管理器管理
		frame.setLayout(gridLayout);

		for(int i = 0 ; i<10; i++){
			frame.add(new JButton(i+""));
		}

		frame.add(new JButton("+"));
		frame.add(new JButton("-"));
		frame.add(new JButton("*"));
		frame.add(new JButton("/"));
		frame.add(new JButton("="));
		frame.add(new JButton("."));
		
		//frame.add(new JButton("aa"));
		
		//初始化窗体
		FrameUtil.initFrame(frame, 300, 300);
	}
}
```

#### CardLayout

CardLayout: 卡片布局管理器

```java
import java.awt.CardLayout;
import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;

import javax.swing.JButton;
import javax.swing.JFrame;
import javax.swing.JPanel;

import cn.itcast.util.FrameUtil;

public class Demo {
	public static void main(String[] args) {

		JFrame frame = new JFrame("卡片布局管理器");
		final JPanel panel = new JPanel();
		frame.add(panel);
		
		//创建一个卡片布局管理器
		final CardLayout cardLayout = new CardLayout();
		panel.setLayout(cardLayout);
		
		//往面板添加数据
		JButton button = new JButton("黑桃A");
		panel.add(button);

		panel.add(new JButton("红桃K"));
		panel.add(new JButton("梅花6"));
		panel.add(new JButton("方块2"));
		
		button.addActionListener(new ActionListener() {

			@Override
			public void actionPerformed(ActionEvent e) {
				cardLayout.next(panel);  //下一张
				//cardLayout.previous(parent);  上一张
			}
		});
		
		//初始化窗体
		FrameUtil.initFrame(frame,300, 300);
	}
}
```

### 事件

事件： 当发生了某个事件的时候，就会有相应处理方案。
 
	事件源	监听器	事件	处理方案

```java
import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;

import javax.swing.JButton;
import javax.swing.JFrame;

import cn.itcast.util.FrameUtil;


public class Demo {
	public static void main(String[] args) {

		JFrame frame = new JFrame("窗体");
		JButton button = new JButton("点我啊");

		frame.add(button);

		//给按钮添加动作监听器    动作时间监听器对于鼠标点击以及空格键都会起作用的。
		button.addActionListener(new ActionListener() {

			//当按钮被点击的时候，就会调用actionPerformed的方法。
			@Override
			public void actionPerformed(ActionEvent e) {  // ActionEvent 当前按钮被点击的时候，jvm就会把对应 的时间传递ActionEvent，并且调用actionPerformed方法。

				//System.out.println("哎呀，被点了...");

				JButton button =(JButton) e.getSource(); //getSource() 获取到事件源
				if(button.getText().equals("点我啊")){
					button.setText("点他吧");
				}else{
					button.setText("点我啊");
				}
			}
		});
		FrameUtil.initFrame(frame, 200, 200);
	}
}
```

####  鼠标事件监听器

```java
import java.awt.event.MouseAdapter;
import java.awt.event.MouseEvent;
import java.awt.event.MouseListener;

import javax.swing.JButton;
import javax.swing.JFrame;

import cn.itcast.util.FrameUtil;

/*
 鼠标事件监听器
 */
public class Demo {
	public static void main(String[] args) {

		JFrame frame = new JFrame("鼠标事件监听器");
		JButton button = new JButton("按钮");
		frame.add(button);

		//给按钮添加鼠标事件监听器
		/*button.addMouseListener(new MouseListener() {

			@Override
			public void mouseReleased(MouseEvent e) {
				System.out.println("鼠标松开...");

			}

			@Override
			public void mousePressed(MouseEvent e) {
				System.out.println("鼠标按下..");
			}

			@Override
			public void mouseExited(MouseEvent e) {
				System.out.println("鼠标移出...");
			}

			@Override
			public void mouseEntered(MouseEvent e) {
				System.out.println("鼠标进入...");
			}

			@Override
			public void mouseClicked(MouseEvent e) {
				System.out.println("鼠标单击..");
			}
		});
		
		添加鼠标监听器的时候我只使用 到单击事件，但是目前要我实现所有的方法？？
		解决方案： 适配器。适配器是实现了MouseListener方法的所有方法，但是实现的方法全部都是空实现。
		*/
		
		button.addMouseListener(new MouseAdapter() {

			@Override
			public void mouseClicked(MouseEvent e) {

				//System.out.println("鼠标单击了..");
		
				if(e.getClickCount()==2){
					System.out.println("双击了..");
				}
			}
		});

		FrameUtil.initFrame(frame, 300, 300);
	}
}
```

####  键盘事件监听器
```java
package cn.itcast.event;

import java.awt.event.KeyAdapter;
import java.awt.event.KeyEvent;
import java.awt.event.KeyListener;

import javax.swing.JButton;
import javax.swing.JFrame;

import cn.itcast.util.FrameUtil;

/*
 键盘事件监听器
 */
public class Demo3 {

	public static void main(String[] args) {
		JFrame frame = new JFrame("键盘事件监听器");
		JButton button = new JButton("按钮");
		frame.add(button);

		//给按钮添加键盘事件监听器
		/*button.addKeyListener(new KeyListener() {

			@Override
			public void keyTyped(KeyEvent e) {
				System.out.println("键入某个键");
			}

			@Override
			public void keyReleased(KeyEvent e) {
				System.out.println("释放某个键");
			}

			@Override
			public void keyPressed(KeyEvent e) {
				System.out.println("按下某个键..");
			}
		});*/

		button.addKeyListener(new KeyAdapter() {

			@Override
			public void keyPressed(KeyEvent e) {

				//System.out.println("按下的字符："+e.getKeyChar());
				//System.out.println("获取键对应的数值："+ e.getKeyCode());

				int code = e.getKeyCode();
				switch (code) {
				case 38:
					System.out.println("上");
					break;
				case 40:
					System.out.println("下");
					break;
				case 37:
					System.out.println("左");
					break;
				case 39:
					System.out.println("右");
					break;
				default:
					break;
				}
			}
		});
		
		FrameUtil.initFrame(frame,300, 300);
	}
}
```

### 例子

```java
import java.awt.BorderLayout;
import java.awt.FileDialog;
import java.awt.TextArea;
import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;
import java.io.File;
import java.io.FileNotFoundException;
import java.io.FileOutputStream;
import java.io.IOException;

import javax.swing.JFrame;
import javax.swing.JMenu;
import javax.swing.JMenuBar;
import javax.swing.JMenuItem;
/*
 菜单组件
 	
 	菜单条(MenuBar) 、  菜单（Menu） 、 菜单项(MenuItem)
 	
 	菜单条可以添加菜单
 	
 	菜单可以添加菜单项
 	
 	复选菜单：
 		首先菜单添加菜单 ， 菜单添加菜单项。

 */

import javax.swing.JTextArea;

import cn.itcast.util.FrameUtil;


public class Notepad {

	JFrame frame = new JFrame("记事本");

	//菜单条
	JMenuBar bar = new JMenuBar();

	//文件菜单
	JMenu fileMenu = new JMenu("文件");
	JMenu editMenu  = new JMenu("编辑");

	JMenu switchMenu = new JMenu("切换工作目录");

	//菜单项
	JMenuItem openMenu = new JMenuItem("打开");
	JMenuItem saveMenu = new JMenuItem("保存");
 
	JMenuItem aboutMenu = new JMenuItem("关于");
	JMenuItem closeMenu = new JMenuItem("关闭");

	JMenuItem  workMenu1 = new JMenuItem("0910project");
	JMenuItem  workMenu2 = new JMenuItem("1208project");
	JMenuItem  workMenu3 = new JMenuItem("1110project");


	TextArea area = new TextArea(20,30);

	public void initNotepad(){
		//菜单添加菜单项目
		fileMenu.add(openMenu);
		fileMenu.add(saveMenu);
		
		//给保存添加事件
		saveMenu.addActionListener(new ActionListener() {

			@Override
			public void actionPerformed(ActionEvent e) {
				try {
					FileDialog fileDialog = new FileDialog(frame, "请选择保存的路径",FileDialog.SAVE);
					fileDialog.setVisible(true);

					//获取用户选择的路径与文件名
					String path = fileDialog.getDirectory();
					String fileName = fileDialog.getFile();

					//创建一个输入对象
					FileOutputStream fileOutputStream = new FileOutputStream(new File(path,fileName));

					//获取文本域的内容，把内容写出
					String content = area.getText();

					//content = content+"\r\n";
					content = content.replaceAll("\n", "\r\n");

					fileOutputStream.write(content.getBytes());

					//关闭资源
					fileOutputStream.close();
				} catch (IOException e1) {
					e1.printStackTrace();
				}
			}
		});
		
		editMenu.add(aboutMenu);
		editMenu.add(closeMenu);
		
		//复选菜单
		switchMenu.add(workMenu1);
		switchMenu.add(workMenu2);
		switchMenu.add(workMenu3);
		//菜单添加菜单就是复选菜单
		fileMenu.add(switchMenu);
		
		//菜单条添加菜单
		bar.add(fileMenu);
		bar.add(editMenu);
		
		//添加菜单条
		frame.add(bar,BorderLayout.NORTH);
		frame.add(area);
		FrameUtil.initFrame(frame, 500, 600);
	}

	public static void main(String[] args) {
		new Notepad().initNotepad();
	}
}
```

```java
import java.awt.BorderLayout;
import java.awt.ScrollPane;
import java.awt.Scrollbar;
import java.awt.TextField;
import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;
import java.awt.event.MouseAdapter;
import java.awt.event.MouseEvent;
import java.io.File;

import javax.swing.JButton;
import javax.swing.JFrame;
import javax.swing.JPanel;
import javax.swing.JScrollBar;
import javax.swing.JTextArea;
import javax.swing.JTextField;

import cn.itcast.util.FrameUtil;

public class FileSearch {

	JFrame frame = new JFrame("文件搜索器");

	JPanel panel = new JPanel();

	JTextField field = new JTextField("请输入目录名...",15);

	JButton button = new JButton("搜索");

	JTextArea area = new JTextArea(15,15);

	//滚动条
	ScrollPane bar = new ScrollPane();

	public void init(){
		
		//先把area添加 到滚动条上。
		bar.add(area);

		//先把组件添加到panel上
		panel.add(field);
		panel.add(button);
		
		//给输入框添加事件
		field.addMouseListener(new MouseAdapter() {

			@Override
			public void mouseClicked(MouseEvent e) {
				JTextField field =  (JTextField) e.getSource();
				if(field.getText().equals("请输入目录名...")){
					field.setText("");
				}
			}
		});
		
		//给按钮添加事件监听器
		button.addActionListener(new ActionListener() {

			@Override
			public void actionPerformed(ActionEvent e) {

				//获取输入框输入的路径
				String path = field.getText();

				//使用输入的路径构建一个FIle对象
				File dir = new File(path);

				//找到目录下的所有子 文件
				File[] files = dir.listFiles();

				for(File file : files){ 
					area.setText(area.getText()+ file.getName()+"\r\n");
				}
			}
		});
		
		//把面板添加到frame上
		frame.add(panel,BorderLayout.NORTH);
		frame.add(bar);		
		FrameUtil.initFrame(frame, 300, 400);
	}

	public static void main(String[] args) {
		new FileSearch().init();
	}
}
```

### SnakeGame

贪吃蛇游戏： 一个游戏最核心的部分是游戏的业务逻辑代码。 

分析里面的有几个事物：
	1. 地图
	2. 蛇
	3. 食物

了解游戏规则：
	1.蛇初始化的时候是三个节点，而且一开始蛇是出现在地图的中间位置。
	2. 蛇吃到东西会长长一节。
	3. 蛇咬到自己会死，蛇撞墙也会死。
	4. 食物是随机生成的，食物不能长在石头上，食物可以长在蛇身上。

建议：  由于目前我们经验不足，所以在做项目的时候往往会出现比较大幅度修改代码 。 自己建立自己的代码库。

地图：
		***********************************
		*                                 *
		*                      @          *
		*                                 *
		*            ##$                  *
		*                                 *
		*                                 *
		*                                 *
		***********************************

	char[][]  宽  高

把这个地图存储到一个二维数组中

```java
import java.awt.Point;
import java.util.LinkedList;
import java.util.Random;

public class SnakeGame {

	//地图的宽(列数)
	public static final int WIDTH = 35;

	//地图高（行数）
	public static final int HEIGHT = 9;	

	//地图
	private char[][] background = new char[HEIGHT][WIDTH];

	//使用集合保存蛇节点的所有信息
	LinkedList<Point>  snake = new LinkedList<Point>(); 

	//食物
	Point	food;

	//向上移动
	public void moveUp(){
		//获取原来蛇头 
		Point head = snake.getFirst();
		//添加新的蛇头
		snake.addFirst(new Point(head.x,head.y-1));
		//删除蛇尾
		snake.removeLast();
	}

	//向下走
	public void moveDown(){
		//获取到原来的蛇头
		Point  head = snake.getFirst();
		//添加新的蛇头
		snake.addFirst(new  Point(head.x,head.y+1));
		//删除蛇尾
		snake.removeLast();
	}

	//向左走
	public void moveLeft(){
		//获取到原来的蛇头
		Point  head = snake.getFirst();
		//添加新的蛇头
		snake.addFirst(new  Point(head.x-1,head.y));
		//删除蛇尾
		snake.removeLast();
	}
		
	//向右走
	public void moveRight(){
		//获取到原来的蛇头
		Point  head = snake.getFirst();
		//添加新的蛇头
		snake.addFirst(new  Point(head.x+1,head.y));
		//删除蛇尾
		snake.removeLast();
	}

	//生成食物 
	public void createFood(){
		//创建一个随机数对象
		Random random = new Random();
		while(true){
			int x = random.nextInt(WIDTH); 
			int y = random.nextInt(HEIGHT);
			if(background[y][x]!='*'){
				food = new Point(x,y);
				break;
			}
		}
	}

	//显示食物
	public void showFood(){
		background[food.y][food.x] ='@';
	}

	//初始化蛇
	public void initSnake(){
		int x = WIDTH/2;
		int y = HEIGHT/2;
		snake.addFirst(new Point(x-1,y));
		snake.addFirst(new Point(x,y));
		snake.addFirst(new Point(x+1,y));
	}

	//显示蛇--->实际上就是将蛇节点 的坐标信息反馈到地图上，在地图上画上对应的字符而已
	public void showSnake(){
		//取出蛇头
		Point head = snake.getFirst();
		background[head.y][head.x] = '$';
		//画蛇身
		for(int i =1; i<snake.size() ; i++ ){
			Point body = snake.get(i);
			background[body.y][body.x] = '#';
		}
	}

	//初始化地图
	public void initBackground(){
		for(int rows = 0 ; rows<background.length ; rows++){
			for(int cols = 0  ; cols<background[rows].length ; cols++ ){
				if(rows==0||rows==(HEIGHT-1)||cols==0||cols==(WIDTH-1)){  //第一行、最后一行、 第一列与最后一列
					background[rows][cols] = '*';
				}else{
					background[rows][cols] = ' ';
				}
			}
		}
	}

	//显示地图的
	public void showBackground() {
		//打印二维数组
		for(int rows = 0 ; rows<background.length ; rows++){ // rows =0 
			for(int cols = 0  ; cols<background[rows].length ; cols++ ){
				System.out.print(background[rows][cols]);
			}
			System.out.println(); //换行
		}
	}

	//刷新游戏 的状态 
	public void refrash(){
		//清空游戏之前的所有状态信息
		initBackground();
		//把蛇最新的状态反馈到地图上
		showSnake();
		//把食物的信息反馈到地图婚丧。
		showFood();
		//显示当前地图的信息
		showBackground();
		
	}

	public static void main(String[] args) throws Exception {

		SnakeGame snake = new SnakeGame();
		//初始化地图
		snake.initBackground();
		//初始化蛇
		snake.initSnake(); 
		//把蛇的信息反馈到地图上
		snake.showSnake();
		
		//初始化事物
		snake.createFood();
		//显示食物
		snake.showFood();
		snake.showBackground();
		
		//向上三步
		for(int i = 0;  i < 3; i++){
			snake.moveUp();
			snake.refrash();
			Thread.sleep(500);
		}
		
		//向下三步
		for(int i = 0;  i < 3; i++){
			snake.moveDown();
			snake.refrash();
			Thread.sleep(500);
		}
		
		//向右三步
		for(int i = 0;  i < 3; i++){
			snake.moveRight();
			snake.refrash();
			Thread.sleep(500);
		}
		
		//向左三步
		for(int i = 0;  i < 3; i++){
			snake.moveLeft();
			snake.refrash();
			Thread.sleep(500);
		}
	}
}
```

```java
import java.awt.Color;
import java.awt.Font;
import java.awt.Graphics;

import javax.swing.JFrame;
import javax.swing.JPanel;

import cn.itcast.util.FrameUtil;

public class Demo extends JPanel {  //Demo 也是一个面板
	
	@Override
	public void paint(Graphics g) {  // Graphics g 画笔  使用该画笔可以画任何的东西。

		//设置画笔的颜色
		g.setColor(Color.GRAY);
		//画矩形
		g.fill3DRect(0, 0, 20, 20, true);
		g.fill3DRect(20, 0, 20, 20, true);
		g.setColor(Color.GREEN);
		g.fill3DRect(20,20, 20, 20, true);
		
		//写字
		g.setColor(Color.RED);
		//设置画笔 的字体
		g.setFont(new Font("宋体", Font.BOLD, 30));
		g.drawString("GAME OVER", 300, 200);
		
	}

	public static void main(String[] args) {

		JFrame frame = new JFrame("测试");
		Demo d = new Demo();
		frame.add(d);
		FrameUtil.initFrame(frame,700, 500);
	}
}
```

```java
package cn.ixfosa.view;

import java.awt.Color;
import java.awt.Graphics;
import java.awt.Point;
import java.awt.event.KeyAdapter;
import java.awt.event.KeyEvent;
import java.util.LinkedList;

import javax.swing.JFrame;
import javax.swing.JPanel;

import cn.itcast.util.FrameUtil;

public class SnakeView  extends JPanel{

	//地图的宽(列数)
	public static final int WIDTH = 40;

	//地图高（行数）
	public static final int HEIGHT = 30;

	//格子宽
	public static final int CELLWIDTH = 20;

	//格子高
	public static final int CELLHEIGHT = 20;

	//地图
	private boolean[][] background = new boolean[HEIGHT][WIDTH];

	//使用集合保存蛇节点的所有信息
	LinkedList<Point>  snake = new LinkedList<Point>(); 

	//初始化蛇
	public void initSnake(){
		int x = WIDTH/2;
		int y = HEIGHT/2;
		snake.addFirst(new Point(x-1,y));
		snake.addFirst(new Point(x,y));
		snake.addFirst(new Point(x+1,y));
	}
		
	//向上移动
	public void moveUp(){
		//获取原来蛇头 
		Point head = snake.getFirst();
		//添加新的蛇头
		snake.addFirst(new Point(head.x,head.y-1));
		//删除蛇尾
		snake.removeLast();
	}
	
	//初始化地图
	public void initBackground(){
		for(int rows = 0 ; rows<background.length ; rows++){
			for(int cols = 0  ; cols<background[rows].length ; cols++ ){
				if(rows==0||rows==(HEIGHT-1)||cols==0||cols==(WIDTH-1)){  //第一行、最后一行、 第一列与最后一列
					background[rows][cols] = true;
				}
			}
		}
	}
		
	
	@Override
	public void paint(Graphics g) {
		//画地图		
		for(int rows = 0 ; rows<background.length ; rows++){ // rows =0 
			for(int cols = 0  ; cols<background[rows].length ; cols++ ){
				if(background[rows][cols]){
					//石头
					g.setColor(Color.GRAY);
				}else{
					//空地
					g.setColor(Color.WHITE);
				}
				//画矩形
				g.fill3DRect(cols*CELLWIDTH, rows*CELLHEIGHT, CELLWIDTH, CELLHEIGHT, true);
			}
		}
		
		//画蛇
	
		//取出蛇头
		Point head = snake.getFirst();
		g.setColor(Color.RED);
		g.fill3DRect(head.x*CELLWIDTH, head.y*CELLHEIGHT, CELLWIDTH, CELLHEIGHT, true);
		//画蛇身
		g.setColor(Color.GREEN);
		for(int i =1; i<snake.size() ; i++ ){
			Point body = snake.get(i);
			g.fill3DRect(body.x*CELLWIDTH, body.y*CELLHEIGHT, CELLWIDTH, CELLHEIGHT, true);
		}
	}
	
	public static void main(String[] args) {
		JFrame frame = new JFrame("贪吃蛇");
		final SnakeView snakeView = new SnakeView();
		snakeView.initBackground();
		snakeView.initSnake();
		frame.add(snakeView);
		FrameUtil.initFrame(frame, WIDTH*CELLWIDTH+20, HEIGHT*CELLHEIGHT+35);
		
		//给窗口添加事件监听
		frame.addKeyListener(new KeyAdapter() {

			@Override
			public void keyPressed(KeyEvent e) {
				int code = e.getKeyCode();
				switch (code) {
				case KeyEvent.VK_UP:
						snakeView.moveUp();
						//重画游戏
						snakeView.repaint(); //调用repaint方法的时候实际上就是调用了paint方法。
					break;

				default:
					break;
				}
			}
		});
	}
}

/*----------------------------------------------------------------------*/

package cn.ixfosa.snake;

import java.awt.BorderLayout;
import java.awt.Point;
import java.awt.event.KeyAdapter;
import java.awt.event.KeyEvent;
import java.util.LinkedList;
import java.util.Random;

import javax.swing.JButton;
import javax.swing.JFrame;

import cn.ixfosa.util.FrameUtil;

public class SnakeGame {
	
	//地图的宽(列数)
	public static final int WIDTH = 35;
	
	//地图高（行数）
	public static final int HEIGHT = 9;	
	
	//地图
	private char[][] background = new char[HEIGHT][WIDTH];
	
	//使用集合保存蛇节点的所有信息
	LinkedList<Point>  snake = new LinkedList<Point>(); 
	
	//食物
	Point	food;
	
	//使用四个常量表示四个方向
	public static final int UP_DIRECTION = 1;  //上
	
	public static final int DOWN_DIRECTION = -1;  //下
	
	public static final int LEFT_DIRECTION = 2;  //左
	
	public static final int RIGHT_DIRECTION =-2;  //右

	//蛇当前的方向
	int currentDrection = -2; // 蛇默认是向右行走
	
	//记录游戏是否结束
	static	boolean isGameOver = false; //默认游戏没有结束的。
	
	//蛇移动的方法
	public void move(){
		Point head = snake.getFirst();
		//蛇是根据当前的方向移动的
		switch (currentDrection) {
			case UP_DIRECTION:
				//添加新的蛇头
				snake.addFirst(new Point(head.x,head.y-1));
				break;
			case DOWN_DIRECTION:
				//添加新的蛇头
				snake.addFirst(new Point(head.x,head.y+1));
				break;
			case LEFT_DIRECTION:
				if(head.x==0){
					snake.addFirst(new Point(WIDTH-1,head.y));
				}else{
					//添加新的蛇头
					snake.addFirst(new Point(head.x-1,head.y));
				}
				break;
			case RIGHT_DIRECTION:
				if(head.x==WIDTH-1){
					snake.addFirst(new Point(0,head.y));
				}else{
					//添加新的蛇头
					snake.addFirst(new Point(head.x+1,head.y));
				}
				break;
			default:
				break;
		}
		
		if(eat()){
			//吃到了食物
			createFood();
		}else{
			//删除蛇尾
			snake.removeLast(); 		
		}
	}
	
	//改变当前方向的方法
	public  void changeDirection(int newDirection){
		//判断新方向是否与当前方向是否是相反方向，才允许其改变
		if(newDirection+currentDrection!=0){
			this.currentDrection = newDirection;
		}
	}

	//生成食物 
	public void createFood(){
		//创建一个随机数对象
		Random random = new Random();
		while(true){
			int x = random.nextInt(WIDTH); 
			int y = random.nextInt(HEIGHT);
			if(background[y][x]!='*'){
				food = new Point(x,y);
				break;
			}
		}
	}
	
	//显示食物
	public void showFood(){
		background[food.y][food.x] ='@';
	}
	
	//初始化蛇
	public void initSnake(){
		int x = WIDTH/2;
		int y = HEIGHT/2;
		snake.addFirst(new Point(x-1,y));
		snake.addFirst(new Point(x,y));
		snake.addFirst(new Point(x+1,y));
	}
	
	//显示蛇--->实际上就是将蛇节点 的坐标信息反馈到地图上，在地图上画上对应的字符而已
	public void showSnake(){
	
		//画蛇身
		for(int i =1; i<snake.size() ; i++ ){
			Point body = snake.get(i);
			background[body.y][body.x] = '#';
			
		}
		
		//取出蛇头
		Point head = snake.getFirst();
		background[head.y][head.x] = '$';
	}
	
	//初始化地图
	public void initBackground(){
		for(int rows = 0 ; rows<background.length ; rows++){
			for(int cols = 0  ; cols<background[rows].length ; cols++ ){
				if(rows==0||rows==(HEIGHT-1)){  //第一行、最后一行、 第一列与最后一列
					background[rows][cols] = '*';
				}else{
					background[rows][cols] = ' ';
				}
			}
		}
	}

	//显示地图的
	public void showBackground() {
		//打印二维数组
		for(int rows = 0 ; rows<background.length ; rows++){ // rows =0 
			for(int cols = 0  ; cols<background[rows].length ; cols++ ){
				System.out.print(background[rows][cols]);
			}
			System.out.println(); //换行
		}
	}
	
	//刷新游戏 的状态 
	public void refrash(){
		//清空游戏之前的所有状态信息
		initBackground();
		//把蛇最新的状态反馈到地图上
		showSnake();
		//把食物的信息反馈到地图婚丧。
		showFood();
		//显示当前地图的信息
		showBackground();
	}
	
	//吃食物
	public boolean eat(){
		//获取到原来的蛇头
		Point head = snake.getFirst();
		if(head.equals(food)){
			return true; 
		}
		return false;
	}
	
	//游戏结束的方法
	public void isGameOver(){
		//撞墙死亡
		Point head = snake.getFirst();
		if(background[head.y][head.x]=='*'){
			isGameOver = true;
		}
		
		//咬到自己  蛇身
		for(int i = 1; i<snake.size() ; i++){
			Point body = snake.get(i);
			if(head.equals(body)){
				isGameOver = true;
			}
		}
	}
	
	public static void main(String[] args) throws Exception {
		final SnakeGame snake = new SnakeGame();
		//初始化地图
		snake.initBackground();
		//初始化蛇
		snake.initSnake(); 
		//把蛇的信息反馈到地图上
		snake.showSnake();
		
		//初始化事物
		snake.createFood();
		//显示食物
		snake.showFood();
		snake.showBackground();
		
		JFrame frame = new JFrame("方向盘");
		frame.add(new JButton("↑"),BorderLayout.NORTH);
		frame.add(new JButton("↓"),BorderLayout.SOUTH);
		frame.add(new JButton("←"),BorderLayout.WEST);
		frame.add(new JButton("→"),BorderLayout.EAST);
		JButton button = new JButton("点击控制方向");
		frame.add(button);
		FrameUtil.initFrame(frame, 300, 300);
		
		//给按钮添加事件监听器
		button.addKeyListener(new KeyAdapter(){
			@Override
			public void keyPressed(KeyEvent e) {
				int code = e.getKeyCode();
				switch (code) {
					case KeyEvent.VK_UP:
						snake.changeDirection(UP_DIRECTION);
						break;
					case KeyEvent.VK_DOWN:
						snake.changeDirection(DOWN_DIRECTION);
						break;
					case KeyEvent.VK_LEFT:
						snake.changeDirection(LEFT_DIRECTION);
						break;
					case KeyEvent.VK_RIGHT:
						snake.changeDirection(RIGHT_DIRECTION);
						break;
					default:
						break;
				}
				snake.move();
				snake.isGameOver();//判断是否游戏结束
				snake.refrash();
				if(isGameOver){
					System.out.println("游戏结束了..");
					System.exit(0);
				}
			}
		});
	}
}
```

## reflect-反射


反射： 当一个字节码文件加载到内存的时候，jvm会对该字节码进行解剖，然后会创建一个对象的`Class对象`，把字节码文件的信息全部都
存储到该Class对象中，我们只要获取到Class对象，我们就可以使用字节码对象设置对象的属性或者调用对象的方法等操作....

注意： 在反射技术中一个类的任何成员都有对应 的类进行描述。  比如：  成员变量（Field）   方法----> Method类  


```java
package cn.ixfosa.reflect;

public class Person {
	
	private int id;
	
	String name;
	
	public Person(int id,String name){
		this.id = id;
		this.name = name;
	}
	
	public Person(){}
	
	
	public void eat(int num){
		System.out.println(name+"吃很"+num+"斤饭");
	}
	
	private static void sleep(int num){
		System.out.println("明天睡上"+num+"小时");
	}
	
	
	public static void  sum(int[] arr){
		System.out.println("长度是："+ arr.length);
	}
	
	@Override
	public String toString() {
		return " 编号："+ this.id +" 姓名："+ this.name;
	}
	
}

/*------------------------------------------------------------*/

package cn.ixfosa.reflect;

public class Demo {
	
	Person p;
	
	public static void main(String[] args) throws ClassNotFoundException {
		//Person p = new Person(110,"狗娃");
		
		//推荐使用： 获取Class对象的方式一
		Class clazz1 = Class.forName("cn.ixfosa.reflect.Person");
		System.out.println("clazz1:"+ clazz1);
		
		
		//获取Class对象的方式二： 通过类名获取
		Class clazz2 = Person.class;
		System.out.println("clazz1==clazz2?"+ (clazz1==clazz2));
		

		//获取Class对象的方式三 ：通过对象获取
		Class clazz3 = new Person(110,"狗娃").getClass();
		System.out.println("clazz2==clazz3?"+ (clazz2==clazz3));
	}
}

/*------------------------------------------------------------*/

package cn.ixfosa.reflect;

import java.lang.reflect.Constructor;

/*
 如何通过Class对象获取构造方法。
 */
public class Demo {
	public static void main(String[] args) throws Exception {

		//获取到对应的Class对象
		Class clazz = Class.forName("cn.ixfosa.reflect.Person");
		
		//通过Class对象获取对应的构造方法
		/*Constructor[] constructors = clazz.getConstructors();  // getConstructors()获取一个类的所有公共的构造方法
		for(Constructor constructor : constructors){
			System.out.println(constructor);
		}
		
		Constructor[] constructors =  clazz.getDeclaredConstructors(); //获取到一个类的所有构造方法，包括私有的在内 。
		for(Constructor constructor : constructors){
			System.out.println(constructor);
		}
		*/
		
		/*Constructor constructor = clazz.getConstructor(int.class,String.class);  // getConstructor 获取单个指定的构造方法。
		Person p  = (Person) constructor.newInstance(999,"ixfosa"); // newInstance()创建一个对象
		System.out.println(p);*/
		

		//获取私有的构造函数
		Constructor constructor =  clazz.getDeclaredConstructor(null);

		//暴力反射
		constructor.setAccessible(true);
		Person p  =(Person) constructor.newInstance(null);
		System.out.println(p);
	}
}

/*------------------------------------------------------------*/

package cn.ixfosa.reflect;

import java.lang.reflect.Method;

/*
 通过Class对象获取到对应的方法。
 
 在反射技术中使用了Method类描述了方法的。
 
 */
public class Demo3 {
	public static void main(String[] args) throws Exception {

		//获取到对应的Class对象
		Class clazz = Class.forName("cn.ioxfosa.reflect.Person");

		//获取到所有公共的方法
		/*Method[] methods = clazz.getMethods(); // getMethods() 获取所有的公共方法而已。

		Method[] methods = clazz.getDeclaredMethods(); //获取到所有的方法，但是不包含父类的方法。

		for(Method method : methods){
			System.out.println(method);
		}*/
		

		Person p = new Person(110,"狗娃");
		/*	
		Method m = clazz.getMethod("eat", int.class);
		m.invoke(p, 3); //invoke 执行一个方法。 第一个参数：方法的调用对象。 第二参数： 方法所需要的参数。
		
		//执行私有的方法
		Method m =clazz.getDeclaredMethod("sleep",int.class);
		//设置访问权限允许访问
		m.setAccessible(true);
		m.invoke(null, 6);*/
		
		Method m = clazz.getMethod("sum", int[].class);
		m.invoke(p, new int[]{12,5,9});
	}
}

/*------------------------------------------------------------*/

package cn.ioxfosa.reflect;

import java.lang.reflect.Field;
/*
 通过反射获取对应的成员变量
 
 在反射技术中使用了Field类描述了成员变量的。
 */

public class Demo4 {
	public static void main(String[] args) throws Exception {

		//获取到对应的Class对象
		Class clazz = Class.forName("cn.ioxfosa.reflect.Person");

		//获取 到所有的成员变量
		/*Field[] fields = clazz.getDeclaredFields();
		for(Field field  : fields){
			System.out.println(field);
		}*/

		Person p = new Person();

		Field field = clazz.getDeclaredField("id");

		//设置访问权限可以访问
		field.setAccessible(true);

		field.set(p, 110); //第一个参数： 设置该数据 的成员变量，第二个参数：属性值。

		System.out.println(p);
	}
}
```

## ntrospector-内省

内省--->一个变态的反射.

内省主要解决的问题： 把对象的属性数据封装到对象中。

```java
package cn.ixfosa.introspector;

//实体类---javaBean
public class Person {
	
	private int id;
	
	private String name;

	public Person(int id, String name) {
		super();
		this.id = id;
		this.name = name;
	}
	
	
	public Person(){}
	
	
	public int getId() {
		return id;
	}


	public void setId(int id) {
		this.id = id;
	}


	public String getName() {
		return name;
	}


	public void setName(String name) {
		this.name = name;
	}


	@Override
	public String toString() {
	
		return "编号："+ this.id+" 姓名："+ this.name;
	}
}

```

```java
package cn.ixfosa.introspector;

import java.beans.BeanInfo;
import java.beans.IntrospectionException;
import java.beans.Introspector;
import java.beans.PropertyDescriptor;
import java.lang.reflect.Method;

import org.junit.Test;

public class Demo {
	
	//2. 通过Introspector类获得Bean对象的 BeanInfo，然后通过 BeanInfo 来获取属性的描述器（ PropertyDescriptor ），通过这个属性描述器就可以获取某个属性对应的 getter/setter 方法，然后通过反射机制来调用这些方法。
	@Test
	public void getAllProperty() throws IntrospectionException{

		//Introspector 内省类
		BeanInfo beanInfo = Introspector.getBeanInfo(Person.class);

		//通过BeanInfo获取所有的属性描述其
		PropertyDescriptor[] descriptors = beanInfo.getPropertyDescriptors(); //获取一个类中的所有属性描述器

		for(PropertyDescriptor p : descriptors){
			System.out.println(p.getReadMethod()); //get方法
		}
		
		
	}
	
	

	//1. 通过PropertyDescriptor类操作Bean的属性.
	@Test
	public  void testProperty() throws Exception{

		Person p = new Person();

		//属性描述器
		PropertyDescriptor descriptor = new PropertyDescriptor("id", Person.class);

		//获取属性对应的get或者是set方法设置或者获取属性了。
		Method  m = descriptor.getWriteMethod();  //获取属性的set方法。

		//执行该方法设置属性值
		m.invoke(p,110);
		
		Method readMethod = descriptor.getReadMethod(); //是获取属性的get方法
		
		System.out.println(readMethod.invoke(p, null));
	}
}
```

```java
package cn.ixfosa.introspector;

import java.io.BufferedReader;

import java.io.FileReader;
import java.lang.reflect.Constructor;
import java.lang.reflect.Field;

/*
需求： 编写一个工厂方法根据配置文件的内容，工厂方法返回对应的对象，并且把对象要有对应的属性值。
 */

public class Demo {
	public static void main(String[] args) throws Exception {

		Person p = (Person) getInstance();
		System.out.println(p);
	}
	
	//根据配置文件的内容生产对象的对象并且要把对象的属性值封装到对象中。
	public static Object getInstance() throws Exception{

		BufferedReader bufferedReader = new BufferedReader(new FileReader("obj.txt"));

		String className =  bufferedReader.readLine(); //读取配置文件获取到完整的类名。

		Class clazz = Class.forName(className);

		//通过class对象获取到无参的构造方法
		Constructor constructor = clazz.getConstructor(null);

		//创建对象
		Object o  = constructor.newInstance(null);

		//读取属性值
		String line = null;
		while((line = bufferedReader.readLine())!=null){

			String[] datas = line.split("=");

			//通过属性名获取到对应的Field对象。
			Field field = clazz.getDeclaredField(datas[0]);

			if(field.getType()==int.class){
				field.set(o, Integer.parseInt(datas[1]));
			}else{
				field.set(o, datas[1]);
			}
		}
		return o;
	}
}

/*-----------------------------------------------*/

obj.txt
cn.itcast.introspector.Person
id=110
name=狗娃

```

## 文件路径

### 绝对路径
以根目录或某盘符开头的路径（或者说完整的路径）
例如：
	c:/a.txt （Windows操作系统中）
	c:/xxx/a.txt （Windows操作系统中）
	/var/xx/aa.txt （Linux操作系统中）

绝对路径的问题: 比如C:\abc\a.properties文件路径，该路径在windows上执行没有问题，但是如果把该项目移动到linux上面执行 ，该路径就会出现问题了，因为在linux上面没有c盘的，只有根目录/。 

### 相对路径
相对于当前路径的一个路径。例如当前文件夹为c:/abc时：相对路径a.txt表示c:/abc/a.txt，相对路径xx/a.txt = c:/abc/xx/a.txt
		
	.  表示当前文件夹
	.. 表示上级文件夹	

相对路径存在的问题:相对路径是相对于目前执行`class文件`的时候，`控制台所在的路径`，这样子也会导致出现问题。

### Java程序中的相对路径
在Java程序中使用File时写相对路径，是指相对于执行java命令时当前所在的文件夹。

### classpath路径
在Java程序中，一般情况下使用绝对路径还是相对路径都不太合适，因为Java程序的jar包所放的位置不确定，执行java程序时当前的路径也不确定，所以不合适。一般在Java程序中我们会把资源放到classpath中，然后使用classpath路径查找资源。

Classpath路径：就是使用classpath目前的路径。

```java
package cn.ixfosa.path;

import java.io.File;
import java.io.FileNotFoundException;
import java.io.FileReader;
import java.io.IOException;
import java.io.InputStream;
import java.util.Properties;
/*
 如果经常会发生变化的数据我们可以定义在配置文件上。 比如说：数据库的用户名与密码。
  
 配置文件的路径应该如何写呢？
 	
 	绝对路径：一个文件的完整路径信息。一般绝对路径是包含有盘符的。  绝对路径的缺陷： 因为绝对路径是有盘符开头的，有些系统是没有盘符的。
 	
 	相对路径: 相对路径是相对于当前程序的路径。当前路径就是执行java命令的时候，控制台所在的路径。
 	
 	类文件路径 :类文件路径就是使用了classpath的路径找对应的资源文件。
 	
 	如果需要使用到类文件路径首先先要获取到一个Class对象。
 */
public class DBUtil {
	
	static Properties properties ;
	
	//获取classpath中的资源（InputStream）
	static{
		try {
			properties = new Properties();
			//去加载配置文件  /
			Class clazz = DBUtil.class; 

			//  "/"代表了Classpath的路径。    getResourceAsStream 该方法里面使用的路径就是使用了类文件路径。
			InputStream inputStream = clazz.getResourceAsStream("/db.properties");
			
			properties.load(inputStream);
		} catch (IOException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
		}
	}
	
	public static void main(String[] args) {
		System.out.println("当前路径："+ new File(".").getAbsolutePath() );
		System.out.println("用户名："+ properties.getProperty("userName")+" 密码："+properties.getProperty("password"));
		
	}
}

/*-------------------------------------------------------*/

db.properties
userName=root
password=admin
```

## 注解
### 概述

注解与注释，
	注解： 告诉编译器如何运行程序！
	注释： 给程序员阅读，对编译、运行没有影响；

注解作用，
	1. 告诉编译器如何运行程序；
	2. 简化(取代)配置文件   【案例后再看】

### 常用的注解
```java
package cn.ixfosa.anno;

import java.util.List;

import org.junit.Test;

/**
 * 常用的注解
 *
 */
public class App {

	// 重写父类的方法
	@Override
	public String toString() {
		return super.toString();
	}

	// 抑制编译器警告
	@SuppressWarnings(value = {"unused","unchecked"})
	private void save() {
		List list = null;
	}

	// 标记方法以及过时
	@Deprecated
	private void save1() {
	}
	
	
	@Test
	public void testMain() throws Exception {
	}
}
```

### 自定义注解

通过自定义注解，可以给类、字段、方法上添加描述信息！

#### 注解基本写法

```java
/**
 * 自定义注解  (描述一个作者)
 *
 */
public @interface Author {

	/**
	 * 注解属性
	 * 	  1. 修饰为默认或public
	 *    2. 不能有主体
	 */
	String name();
	int age();
}

/**
 * 使用自定义注解  
 *
 */
@Author(name = "ixfosa", age = 22)
public void save() {

}
```

#### 带默认值的注解

```java
public @interface Author {

	/**
	 * 注解属性
	 * 	  1. 修饰为默认或public
	 *    2. 不能有主体
	 */
	String name();
	int age() default 30;   // 带默认值的注解;  使用的时候就可以不写此属性值
}
```

#### 默认名称的注解

注解属性名称为value，这就是默认名称

```java
public @interface Author {
	// 如果注解名称为value,使用时候可以省略名称，直接给值
	// (且注解只有一个属性时候才可以省略名称)
	String value();
}

//使用
@Author("ixfosa")
@Author(value = "ixfosa")
```

```java
//注解属性类型为数组：
public @interface Author {
	String[] value() default {"test1","test2"};
}

//使用：
@Author（{“”，“”}）
public void save() {

}
```

### 元注解

元注解，表示注解的注解！

指定注解的可用范围：

```java
	@Target({
	TYPE,        类
	FIELD,       字段
	METHOD,      方法
	PARAMETER,   参数
	CONSTRUCTOR, 构造器
	LOCAL_VARIABLE  局部变量
})


// 元注解 - 2. 指定注解的声明周期
@Retention(RetentionPolicy.SOURCE)     注解只在源码级别有效
@Retention(RetentionPolicy.CLASS)      注解在字节码即别有效默认值
@Retention(RetentionPolicy.RUNTIME)    注解在运行时期有效

```

### 注解反射

```java
package cn.ixfosa.anno;

import static java.lang.annotation.ElementType.CONSTRUCTOR;
import static java.lang.annotation.ElementType.FIELD;
import static java.lang.annotation.ElementType.LOCAL_VARIABLE;
import static java.lang.annotation.ElementType.METHOD;
import static java.lang.annotation.ElementType.PARAMETER;
import static java.lang.annotation.ElementType.TYPE;

import java.lang.annotation.Retention;
import java.lang.annotation.RetentionPolicy;
import java.lang.annotation.Target;
import java.lang.reflect.Type;

// 元注解 - 1. 定义注解的可用范围
@Target({TYPE,FIELD , METHOD, PARAMETER, CONSTRUCTOR, LOCAL_VARIABLE})
//@Target({METHOD,FIELD,TYPE})   指定只能在方法、字段、类上用；

// 元注解 - 2. 指定注解的声明周期
@Retention(RetentionPolicy.RUNTIME)   // 字节码级别有效
public @interface Author {
	
	String authorName() default "ixfosa";
	int age() default 22;	
	String remark();
}

```

```java
package cn.ixfosa.anno;

public @interface Id {

}
```


```java
package cn.ixfosa.anno;

import java.lang.reflect.Field;
import java.lang.reflect.Method;

import javax.persistence.Table;


import org.junit.Test;

/**
 * 测试：自定义注解的语法
 *
 */

public class App_2 {
	
	private String test;
	
	@Id
	@Author(remark = "保存信息！！！", age = 19)
	public void save() throws Exception {

		// 获取注解信息： name/age/remark
		// 1. 先获取代表方法的Method类型;
		Class clazz = App_2.class;
		Method m = clazz.getMethod("save");
		
		// 2. 再获取方法上的注解
		Author author = m.getAnnotation(Author.class);
		// 获取输出注解信息
		System.out.println(author.authorName());
		System.out.println(author.age());
		System.out.println(author.remark());
	}
	
	@Test
	public void testMain() throws Exception {
		save();
	}
}
```

###  注解，优化BaseDao的代码
当表名与数据库名称不一致、 字段与属性不一样、主键不叫id， 上面的BaseDao不能用！
这时，可以通过配置文件(XML) 解决！
	
注解：
	简化XML配置， 程序处理非常方便！
	(不便于维护： 例如修改字段名，要重新编译！)

XML
	便于维护！  需要些读取代码！

```java
package cn.ixfosa.utils;

import java.lang.annotation.Retention;
import java.lang.annotation.RetentionPolicy;
import java.lang.annotation.Target;
import static java.lang.annotation.ElementType.FIELD;

/**
 * 描述一个字段
 *
 */
@Target({FIELD})
@Retention(RetentionPolicy.RUNTIME)  // 指定注解在运行时期有效
public @interface Column {

	String columnName();
}

/**************************************************************/

package cn.ixfosa.utils;

import static java.lang.annotation.ElementType.FIELD;

import java.lang.annotation.Retention;
import java.lang.annotation.RetentionPolicy;
import java.lang.annotation.Target;

/**
 * 描述一个主键字段
 *
 */
@Target({FIELD})
@Retention(RetentionPolicy.RUNTIME)  // 指定注解在运行时期有效
public @interface Id {

}

/**************************************************************/

package cn.ixfosa.utils;

import static java.lang.annotation.ElementType.TYPE;

import java.lang.annotation.Retention;
import java.lang.annotation.RetentionPolicy;
import java.lang.annotation.Target;

/**
 * 注解，描述表名称
 *
 */
@Target({TYPE})
@Retention(RetentionPolicy.RUNTIME)  // 指定注解在运行时期有效
public @interface Table {

	String tableName();
}
```

```java
package cn.ixfosa.eg;

import cn.ixfosa.utils.Column;
import cn.ixfosa.utils.Id;
import cn.ixfosa.utils.Table;

// Admin=a_admin
@Table(tableName="a_admin")
public class Admin {

	@Id
	@Column(columnName = "a_id")
	private int id;
	
	@Column(columnName = "a_userName")
	private String userName;
	
	@Column(columnName = "a_pwd")
	private String pwd;
	
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
	public String getPwd() {
		return pwd;
	}
	public void setPwd(String pwd) {
		this.pwd = pwd;
	}
	@Override
	public String toString() {
		return "Admin [id=" + id + ", pwd=" + pwd + ", userName=" + userName
				+ "]";
	}
}
```

```java
package cn.ixfosa.eg;

import java.lang.reflect.Field;
import java.lang.reflect.ParameterizedType;
import java.lang.reflect.Type;
import java.sql.ResultSet;
import java.sql.SQLException;
import java.util.ArrayList;
import java.util.List;

import org.apache.commons.beanutils.BeanUtils;
import org.apache.commons.dbutils.ResultSetHandler;

import cn.ixfosa.b_reflect.JdbcUtils;
import cn.ixfosa.utils.Column;
import cn.ixfosa.utils.Id;
import cn.ixfosa.utils.Table;

/**
 * 解决优化的问题：
 * 	  1. 当数据库表名与类名不一致、 
 *    2. 字段与属性不一样、
 *    3. 主键不叫id 
 *    
 */
public class BaseDao<T> {
	
	// 当前运行类的类型
	private Class<T> clazz;
	// 表名
	private String tableName;
	// 主键
	private String id_primary;

	// 拿到当前运行类的参数化类型中实际的类型  ( BaseDao<Admin> ,  Admin.class)
	public BaseDao(){
		Type type = this.getClass().getGenericSuperclass();
		ParameterizedType pt = (ParameterizedType) type;
		Type[] types = pt.getActualTypeArguments();
		clazz = (Class<T>) types[0];
		
		//已经拿到：  Admin.class
		
		/*******1. 获取表名*******/
		Table table = clazz.getAnnotation(Table.class);
		tableName = table.tableName();
		
		/*******2. 获取主键字段*******/
		//获取当前运行类的所有字段、遍历、获取每一个字段上的id注解
		Field[] fs = clazz.getDeclaredFields();
		for (Field f : fs) {
			
			// 设置强制访问
			f.setAccessible(true);
			
			// 获取每一个字段上的id注解
			Id anno_id = f.getAnnotation(Id.class);
			
			// 判断
			if (anno_id != null) {
				// 如果字段上有id注解，当前字段(field)是主键； 再获取字段名称
				Column column = f.getAnnotation(Column.class);
				// 主键
				id_primary = column.columnName();
				// 跳出循环
				break;
			}
		}
		
		System.out.println("表：" + tableName);
		System.out.println("主键：" + id_primary);
	}
	
	
	public T findById(int id){
		try {
			String sql = "select * from " + tableName + " where " + id_primary +"=?";
			/*
			 * DbUtils的已经封装好的工具类：BeanHandler?   属性=字段
			 */
			return JdbcUtils.getQuerrRunner().query(sql, new BeanHandler<T>(clazz), id);
		} catch (Exception e) {
			throw new RuntimeException(e);
		}
	}
	
	public List<T> getAll(){
		try {
			String sql = "select * from " + tableName;
			return JdbcUtils.getQuerrRunner().query(sql, new BeanListHandler<T>(clazz));
		} catch (Exception e) {
			throw new RuntimeException(e);
		}
	}
}

/**
 * 自定义结果集：封装单个Bean对象
 */
class BeanHandler<T> implements ResultSetHandler<T>{
	// 保存传入的要封装的类的字节码
	private Class<T> clazz;
	public BeanHandler(Class<T> clazz) {
		this.clazz = clazz;
	}
	
	// 封装结果集的方法
	@Override
	public T handle(ResultSet rs) throws SQLException {
		try {
			// 创建要封装的对象  ‘1’
			T t = clazz.newInstance(); 
			// 向下读一行
			if (rs.next()) {
				
				// a. 获取类的所有的Field字段数组
				Field[] fs = clazz.getDeclaredFields();
				
				// b. 遍历， 得到每一个字段类型：Field
				for (Field f : fs) {
				
					// c. 获取”属性名称“
					String fieldName = f.getName();
					
					// e. 获取Field字段上注解   【@Column(columnName = "a_userName")】
					Column column =  f.getAnnotation(Column.class);
					
					// f. ”字段名“
					String columnName = column.columnName();        // 数据库中字段 a_userName
					
					// g. 字段值
					Object columnValue = rs.getObject(columnName);
					
					// 设置（BeanUtils组件）
					BeanUtils.copyProperty(t, fieldName, columnValue);
				}
			}
			return t;
		} catch (Exception e) {
			throw new RuntimeException(e);
		}
	}
}


/**
 * 自定义结果集：封装多个Bean对象到List集合
 */
class BeanListHandler<T> implements ResultSetHandler<List<T>>{
	
	// 要封装的单个对象
	private Class<T> clazz;
	public BeanListHandler(Class<T> clazz){
		this.clazz = clazz;
	}

	// 把从数据库查询到的没一行记录，封装为一个对象，再提交到list集合， 返回List<T>
	@Override
	public List<T> handle(ResultSet rs) throws SQLException {
		List<T> list = new ArrayList<T>();
		try {
			// 向下读一行
			while (rs.next()) {
				
				// 创建要封装的对象  ‘1’
				T t = clazz.newInstance(); 
				
				// a. 获取类的所有的Field字段数组
				Field[] fs = clazz.getDeclaredFields();
				
				// b. 遍历， 得到每一个字段类型：Field
				for (Field f : fs) {
				
					// c. 获取”属性名称“
					String fieldName = f.getName();
					
					// e. 获取Field字段上注解   【@Column(columnName = "a_userName")】
					Column column =  f.getAnnotation(Column.class);
					
					// f. ”字段名“
					String columnName = column.columnName();        // 数据库中字段 a_userName
					
					// g. 字段值
					Object columnValue = rs.getObject(columnName);
					
					// 设置（BeanUtils组件）
					BeanUtils.copyProperty(t, fieldName, columnValue);
				}
				// 对象添加到集合
				list.add(t);
			}
			return list;
		} catch (Exception e) {
			throw new RuntimeException(e);
		}
	}
}

/***************************************************************************/


package cn.ixfosa.eg;

public class AdminDao extends BaseDao<Admin> {

}
```

```java
package cn.ixfosa.eg;

import org.junit.Test;

public class App {

	@Test
	public void testDao() throws Exception {

		AdminDao adminDao = new AdminDao();
		//Admin admin = adminDao.findById(8);
		//System.out.println(admin);
		
		System.out.println(adminDao.findById(8));
		System.out.println(adminDao.getAll());
	}
}
```

## jdk1.8

### Lambda 表达式
#### 概述
Lambda 表达式，也可称为闭包
Lambda 允许把函数作为一个方法的参数（函数作为参数传递进方法中）。
使用 Lambda 表达式可以使代码变的更加简洁紧凑。

语法
```java
//lambda 表达式的语法格式如下：
(parameters) -> expression
//或
(parameters) ->{ statements; }
```

以下是lambda表达式的重要特征:
1. 可选类型声明：不需要声明参数类型，编译器可以统一识别参数值。
2. 可选的参数圆括号：一个参数无需定义圆括号，但多个参数需要定义圆括号。
3. 可选的大括号：如果主体包含了一个语句，就不需要使用大括号。
4. 可选的返回关键字：如果主体只有一个表达式返回值则编译器会自动返回值，大括号需要指定明表达式返回了一个数值。

Lambda 表达式的省略规则：
 1. 小括号中的参数类型可以省略。
 2. 如果小括号中只有一个参数，那么可以省略小括号。
 3. 如果大括号中只有一条语句，那么可以省略大括号，return，分号。
```java
public class Demo04SimpleLambda {

    //定义方法，使用接口当做参数
    public static void method(MyInterface m) {
        m.printStr("hello");
    }

    public static void main(String[] args) {
        //调用method方法，参数传递MyInterface实现类对象
        method(new MyInterface() {
            @Override
            public void printStr(String str) {
                System.out.println(str);
            }
        });
        //使用Lambda表达式的标准格式。
        method((String str) -> {
            System.out.println(str);
        });

        //1. 小括号中的参数类型可以省略。
        method((str) -> {
            System.out.println(str);
        });
        //2. 如果小括号中只有一个参数，那么可以省略小括号。
        method(str -> {
            System.out.println(str);
        });
        //3. 如果大括号中只有一条语句，那么可以省略大括号，return，分号。
        method(str -> System.out.println(str));
    }
}
```

#### Lambda 表达式实例
```java
//Lambda 表达式的简单例子:

// 1. 不需要参数,返回值为 5  
() -> 5  
  
// 2. 接收一个参数(数字类型),返回其2倍的值  
x -> 2 * x  
  
// 3. 接受2个参数(数字),并返回他们的差值  
(x, y) -> x – y  
  
// 4. 接收2个int型整数,返回他们的和  
(int x, int y) -> x + y  
  
// 5. 接受一个 string 对象,并在控制台打印,不返回任何值(看起来像是返回void)  
(String s) -> System.out.print(s)
```

```java
//在 Java8Tester.java 文件输入以下代码：
//Java8Tester.java 文件

public class Java8Tester {
   public static void main(String args[]){

      Java8Tester tester = new Java8Tester();
        
      // 类型声明
      MathOperation addition = (int a, int b) -> a + b;
        
      // 不用类型声明
      MathOperation subtraction = (a, b) -> a - b;
        
      // 大括号中的返回语句
      MathOperation multiplication = (int a, int b) -> { return a * b; };
        
      // 没有大括号及返回语句
      MathOperation division = (int a, int b) -> a / b;
        
      System.out.println("10 + 5 = " + tester.operate(10, 5, addition));
      System.out.println("10 - 5 = " + tester.operate(10, 5, subtraction));
      System.out.println("10 x 5 = " + tester.operate(10, 5, multiplication));
      System.out.println("10 / 5 = " + tester.operate(10, 5, division));
        
      // 不用括号
      GreetingService greetService1 = message ->
      System.out.println("Hello " + message);
        
      // 用括号
      GreetingService greetService2 = (message) ->
      System.out.println("Hello " + message);
        
      greetService1.sayMessage("Runoob");
      greetService2.sayMessage("Google");
   }
    
   interface MathOperation {
      int operation(int a, int b);
   }
    
   interface GreetingService {
      void sayMessage(String message);
   }
    
   private int operate(int a, int b, MathOperation mathOperation){
      return mathOperation.operation(a, b);
   }
}

///执行以上脚本，输出结果为：
$ javac Java8Tester.java 
$ java Java8Tester
10 + 5 = 15
10 - 5 = 5
10 x 5 = 50
10 / 5 = 2
Hello Runoob
Hello Google
```

使用 Lambda 表达式需要注意以下两点：
1. Lambda 表达式主要用来定义行内执行的方法类型接口，例如，一个简单方法接口。在上面例子中，我们使用各种类型的Lambda表达式来定义MathOperation接口的方法。然后我们定义了sayMessage的执行。
2. Lambda 表达式免去了使用匿名方法的麻烦，并且给予Java简单但是强大的函数化的编程能力。

#### 变量作用域
```java
lambda 表达式只能引用标记了 final 的外层局部变量，这就是说不能在 lambda 内部修改定义���域���的局部变量，否则会编译错误。

//在 Java8Tester.java 文件输入以下代码：
public class Java8Tester {
 
   final static String salutation = "Hello! ";
   
   public static void main(String args[]){
      GreetingService greetService1 = message -> 
     	 System.out.println(salutation + message); 
      greetService1.sayMessage("Runoob"); //Hello! Runoob
   }
    
   interface GreetingService {
      void sayMessage(String message);
   }
}
```

也可以直接在 lambda 表达式中访问外层的局部变量：
```java
public class Java8Tester {
    public static void main(String args[]) {
        final int num = 1;
        Converter<Integer, String> s = (param) -> System.out.println(String.valueOf(param + num));
        s.convert(2);  // 输出结果为 3
    }
 
    public interface Converter<T1, T2> {
        void convert(int i);
    }
}

```
lambda 表达式的局部变量可以不用声明为 `final`，但是必须不可被后面的代码修改（即隐性的具有 final 的语义）
```java
int num = 1;  
Converter<Integer, String> s = (param) -> System.out.println(String.valueOf(param + num));
s.convert(2);
num = 5;  
//报错信息：Local variable num defined in an enclosing scope must be final or effectively final 

 ```

 在 Lambda 表达式当中不允许声明一个与局部变量同名的参数或者局部变量。
 ```java
String first = "";  
Comparator<String> comparator = (first, second) -> Integer.compare(first.length(), second.length());  //编译会出错 
```

#### Lambda 使用条件
Lambda 表达式的使用前提:
1. 必须有接口（不能是抽象类），接口中有且仅有一个需要被重写的抽象方法。
2. 必须支持上下文推导，要能够推导出来 Lambda 表达式表示的是哪个接口中的内容。
可以使用接口当做参数，然后传递 Lambda 表达式(常用)，也可以将 Lambda 表达式赋值给一个接口类型的变量。
```java
public class Demo05BeforeLambda {

    //使用接口当做参数
    public static void method(MyInterface m) {//m = s -> System.out.println(s)
        m.printStr("HELLO");
    }

    public static void main(String[] args) {
        //使用接口当做参数，然后传递Lambda表达式。
        //method(s -> System.out.println(s));

        //使用匿名内部类方式创建对象
        /*
        MyInterface m = new MyInterface() {
            @Override
            public void printStr(String str) {
                System.out.println(str);
            }
        };
        */

        MyInterface m = str -> System.out.println(str);
        m.printStr("Hello");
    }
}
```

```java
public class LambdaDemo {

    final static String ST = "Hello!";
    String ss = "Hello!";

    public static void main(String[] args) {
        int c = 8;
        int d = 4;
        // int a = 10; 不允许声明一个与局部变量同名的参数或者局部变量。

        AddService add = (a, b) -> a + b;
        System.out.println(c + " + " + d + " = " + add.operation(c, d));

        add = (a, b) -> a - b;
        System.out.println(c + " - " + d + " = " + add.operation(c, d));

        add = (a, b) -> a * b;
        System.out.println(c + " * " + d + " = " + add.operation(c, d));

        add = (a, b) -> a / b;
        System.out.println(c + " / " + d + " = " + add.operation(c, d));

        add = (int a, int b) -> {
            return a + b;
        };
        System.out.println(c + " + " + d + " = " + add.operation(c, d));

        // 打印方法中可以获取ST,c,d 但是获取不到ss (且c,d不可更改，隐性的具有 final 的语义)
        GreetingService gs = str -> System.out.println(ST + str);
        gs.operation(String.valueOf(c));
    }

    /**
     * 实现方法可自定义
     */
    interface AddService {
        Integer operation(int a, int b);
    }

    interface GreetingService {
        void operation(String str);
    }
}
```

#### lambdas 实现 Runnable 接口
```java
// 1.1使用匿名内部类  
new Thread(new Runnable() {  
    @Override  
    public void run() {  
        System.out.println("Hello world !");  
    }  
}).start();  
  
// 1.2使用 lambda expression  
new Thread(() -> System.out.println("Hello world !")).start();  
  
// 2.1使用匿名内部类  
Runnable race1 = new Runnable() {  
    @Override  
    public void run() {  
        System.out.println("Hello world !");  
    }  
};  
  
// 2.2使用 lambda expression  
Runnable race2 = () -> System.out.println("Hello world !");  
   
// 直接调用 run 方法(没开新线程哦!)  
race1.run();  
race2.run();  
```

#### 使用 Lambdas 排序集合
在 Java 中，Comparator 类被用来排序集合。 
在下面的例子中，我们将根据球员的 name、surname、name 长度以及最后一个字母。 和前面的示例一样，先使用匿名内部类来排序，然后再使用 lambda 表达式精简我们的代码。
```java
//我们将根据name来排序list。 使用旧的方式，代码如下所示:
String[] players = {"Rafael Nadal", "Novak Djokovic",   
    "Stanislas Wawrinka", "David Ferrer",  
    "Roger Federer", "Andy Murray",  
    "Tomas Berdych", "Juan Martin Del Potro",  
    "Richard Gasquet", "John Isner"};  
   
// 1.1 使用匿名内部类根据 name 排序 players  
Arrays.sort(players, new Comparator<String>() {  
    @Override  
    public int compare(String s1, String s2) {  
        return (s1.compareTo(s2));  
    }  
});  


//使用 lambdas，可以通过下面的代码实现同样的功能:

// 1.2 使用 lambda expression 排序 players  
Comparator<String> sortByName = (String s1, String s2) -> (s1.compareTo(s2));  
Arrays.sort(players, sortByName);  
  
// 1.3 也可以采用如下形式:  
Arrays.sort(players, (String s1, String s2) -> (s1.compareTo(s2)));  


//其他的排序如下所示。 和上面的示例一样,代码分别通过匿名内部类和一些lambda表达式来实现Comparator :
// 1.1 使用匿名内部类根据 surname 排序 players  
Arrays.sort(players, new Comparator<String>() {  
    @Override  
    public int compare(String s1, String s2) {  
        return (s1.substring(s1.indexOf(" ")).compareTo(s2.substring(s2.indexOf(" "))));  
    }  
});  
  
// 1.2 使用 lambda expression 排序,根据 surname  
Comparator<String> sortBySurname = (String s1, String s2) ->   
    ( s1.substring(s1.indexOf(" ")).compareTo( s2.substring(s2.indexOf(" ")) ) );  
Arrays.sort(players, sortBySurname);  
  
// 1.3 或者这样,怀疑原作者是不是想错了,括号好多...  
Arrays.sort(players, (String s1, String s2) ->   
      ( s1.substring(s1.indexOf(" ")).compareTo( s2.substring(s2.indexOf(" ")) ) )   
    );  
  
// 2.1 使用匿名内部类根据 name lenght 排序 players  
Arrays.sort(players, new Comparator<String>() {  
    @Override  
    public int compare(String s1, String s2) {  
        return (s1.length() - s2.length());  
    }  
});  
  
// 2.2 使用 lambda expression 排序,根据 name lenght  
Comparator<String> sortByNameLenght = (String s1, String s2) -> (s1.length() - s2.length());  
Arrays.sort(players, sortByNameLenght);  
  
// 2.3 or this  
Arrays.sort(players, (String s1, String s2) -> (s1.length() - s2.length()));  
  
// 3.1 使用匿名内部类排序 players, 根据最后一个字母  
Arrays.sort(players, new Comparator<String>() {  
    @Override  
    public int compare(String s1, String s2) {  
        return (s1.charAt(s1.length() - 1) - s2.charAt(s2.length() - 1));  
    }  
});  
  
// 3.2 使用 lambda expression 排序,根据最后一个字母  
Comparator<String> sortByLastLetter =   
    (String s1, String s2) ->   
        (s1.charAt(s1.length() - 1) - s2.charAt(s2.length() - 1));  
Arrays.sort(players, sortByLastLetter);  
  
// 3.3 or this  
Arrays.sort(players, (String s1, String s2) -> (s1.charAt(s1.length() - 1) - s2.charAt(s2.length() - 1))); 
```

### 函数式接口
#### 概述
函数式接口(Functional Interface)就是`一个有且仅有一个抽象方法`，`但是可以有多个非抽象方法的接口`。
函数式接口可以被隐式转换为 lambda 表达式。
Lambda 表达式和方法引用（实际上也可认为是Lambda表达式）上。

如定义了一个函数式接口如下：
```java
@FunctionalInterface
interface GreetingService {
    void sayMessage(String message);
}
```

那么就可以使用`Lambda`表达式来表示该接口的一个实现(注：JAVA 8 之前一般是用匿名类实现的)：
```java
GreetingService greetService1 = message -> System.out.println("Hello " + message);
```

函数式接口可以对现有的函数友好地支持 lambda。

JDK 1.8 之前已有的函数式接口:
	java.lang.Runnable
	java.util.concurrent.Callable
	java.security.PrivilegedAction
	java.util.Comparator
	java.io.FileFilter
	java.nio.file.PathMatcher
	java.lang.reflect.InvocationHandler
	java.beans.PropertyChangeListener
	java.awt.event.ActionListener
	javax.swing.event.ChangeListener

#### 函数式接口的使用

```java
/*
    函数式接口:有且只有一个抽象方法的接口,称之为函数式接口
    当然接口中可以包含其他的方法(默认,静态,私有)

    @FunctionalInterface注解
    作用:可以检测接口是否是一个函数式接口
        是:编译成功
        否:编译失败(接口中没有抽象方法抽象方法的个数多余1个)
 */
@FunctionalInterface
public interface MyFunctionalInterface {
    //定义一个抽象方法
    public abstract void method();
}
```

```java
/*
    @Override注解
    检查方法是否为重写的方法
        是:编译成功
        否:编译失败
 */
public class MyFunctionalInterfaceImpl implements MyFunctionalInterface{
    @Override
    public void method() {

    }

    /*@Override
    public void method2() {

    }*/

    /*@Override
    public void method3() {

    }*/
}
```

```java
/*
    函数式接口的使用:一般可以作为方法的参数和返回值类型
 */
public class Demo {

    //定义一个方法,参数使用函数式接口MyFunctionalInterface
    public static void show(MyFunctionalInterface myInter){
        myInter.method();
    }

    public static void main(String[] args) {

        //调用show方法,方法的参数是一个接口,所以可以传递接口的实现类对象
        show(new MyFunctionalInterfaceImpl());

        //调用show方法,方法的参数是一个接口,所以我们可以传递接口的匿名内部类
        show(new MyFunctionalInterface() {
            @Override
            public void method() {
                System.out.println("使用匿名内部类重写接口中的抽象方法");
            }
        });

        //调用show方法,方法的参数是一个函数式接口,所以我们可以Lambda表达式
        show(()->{
            System.out.println("使用Lambda表达式重写接口中的抽象方法");
        });

        //简化Lambda表达式
        show(()-> System.out.println("使用Lambda表达式重写接口中的抽象方法"));
    }
}
```

#### Lambda延迟执行
```java
@FunctionalInterface
public interface MessageBuilder {
    //定义一个拼接消息的抽象方法,返回被拼接的消息
    public abstract String builderMessage();
}
```

```java
/*
    日志案例

    发现以下代码存在的一些性能浪费的问题
    调用showLog方法,传递的第二个参数是一个拼接后的字符串
    先把字符串拼接好,然后在调用showLog方法
    showLog方法中如果传递的日志等级不是1级
    那么就不会是如此拼接后的字符串
    所以感觉字符串就白拼接了,存在了浪费
 */
public class Demo01Logger {
    //定义一个根据日志的级别,显示日志信息的方法
    public static void showLog(int level, String message){
        //对日志的等级进行判断,如果是1级别,那么输出日志信息
        if(level==1){
            System.out.println(message);
        }
    }

    public static void main(String[] args) {
        //定义三个日志信息
        String msg1 = "Hello";
        String msg2 = "World";
        String msg3 = "Java";

        //调用showLog方法,传递日志级别和日志信息
        showLog(2,msg1+msg2+msg3);

    }
}
```

```java
/*
    使用Lambda优化日志案例
    Lambda的特点:延迟加载
    Lambda的使用前提,必须存在函数式接口
 */
public class Demo02Lambda {

    //定义一个显示日志的方法,方法的参数传递日志的等级和MessageBuilder接口
    public static void showLog(int level, MessageBuilder mb){
        //对日志的等级进行判断,如果是1级,则调用MessageBuilder接口中的builderMessage方法
        if(level==1){
            System.out.println(mb.builderMessage());
        }
    }

    public static void main(String[] args) {
        //定义三个日志信息
        String msg1 = "Hello";
        String msg2 = "World";
        String msg3 = "Java";

        //调用showLog方法,参数MessageBuilder是一个函数式接口,所以可以传递Lambda表达式
        /*showLog(2,()->{
            //返回一个拼接好的字符串
            return  msg1+msg2+msg3;
        });*/

        /*
            使用Lambda表达式作为参数传递,仅仅是把参数传递到showLog方法中
            只有满足条件,日志的等级是1级
                才会调用接口MessageBuilder中的方法builderMessage
                才会进行字符串的拼接
            如果条件不满足,日志的等级不是1级
                那么MessageBuilder接口中的方法builderMessage也不会执行
                所以拼接字符串的代码也不会执行
            所以不会存在性能的浪费
         */
        showLog(1,()->{
            System.out.println("不满足条件不执行");
            //返回一个拼接好的字符串
            return  msg1+msg2+msg3;
        });
    }
}
```

#### 常用的函数式接口

##### Supplier接口

生产型接口Supplier

java.util.function.Supplier<T> 接口仅包含一个无参的方法： T get() 。用来获取一个泛型参数指定类型的对象数据。由于这是一个函数式接口，这也就意味着对应的Lambda表达式需要“对外提供”一个符合泛型类型的对象数据。
```java
import java.util.function.Supplier;

/*
    常用的函数式接口
    java.util.function.Supplier<T>接口仅包含一个无参的方法：T get()。用来获取一个泛型参数指定类型的对象数据。

    Supplier<T>接口被称之为生产型接口,指定接口的泛型是什么类型,那么接口中的get方法就会生产什么类型的数据
 */
public class Demo01Supplier {
    //定义一个方法,方法的参数传递Supplier<T>接口,泛型执行String,get方法就会返回一个String
    public static String getString(Supplier<String> sup){
        return sup.get();
    }

    public static void main(String[] args) {
        //调用getString方法,方法的参数Supplier是一个函数式接口,所以可以传递Lambda表达式
        String s = getString(()->{
            //生产一个字符串,并返回
            return "胡歌";
        });
        System.out.println(s);

        //优化Lambda表达式
        String s2 = getString(()->"胡歌");
        System.out.println(s2);
    }
}
```

```java
import java.util.function.Supplier;

/*
    练习：求数组元素最大值
        使用Supplier接口作为方法参数类型，通过Lambda表达式求出int数组中的最大值。
        提示：接口的泛型请使用java.lang.Integer类。
 */
public class Demo02Test {
	
   //定义一个方法,用于获取int类型数组中元素的最大值,方法的参数传递Supplier接口,泛型使用Integer
   public static int getMax(Supplier<Integer> sup){
       return sup.get();
   }

    public static void main(String[] args) {

        //定义一个int类型的数组,并赋值
        int[] arr = {100,0,-50,880,99,33,-30};

        //调用getMax方法,方法的参数Supplier是一个函数式接口,所以可以传递Lambda表达式
        int maxValue = getMax(()->{
            //获取数组的最大值,并返回
            //定义一个变量,把数组中的第一个元素赋值给该变量,记录数组中元素的最大值
            int max = arr[0];
            //遍历数组,获取数组中的其他元素
            for (int i : arr) {
                //使用其他的元素和最大值比较
                if(i>max){
                    //如果i大于max,则替换max作为最大值
                    max = i;
                }
            }
            //返回最大值
            return max;
        });
        System.out.println("数组中元素的最大值是:"+maxValue);
    }
}
```

##### Consumer接口

消费型接口Consumer

java.util.function.Consumer<T> 接口则正好与Supplier接口相反，它不是生产一个数据，而是消费一个数据，其数据类型由泛型决定
	Consumer 接口中包含抽象方法 void accept(T t) ，意为消费一个指定泛型的数据
```java
import java.util.function.Consumer;

/*
    java.util.function.Consumer<T>接口则正好与Supplier接口相反，
        它不是生产一个数据，而是消费一个数据，其数据类型由泛型决定。
    Consumer接口中包含抽象方法void accept(T t)，意为消费一个指定泛型的数据。

   Consumer接口是一个消费型接口,泛型执行什么类型,就可以使用accept方法消费什么类型的数据
   至于具体怎么消费(使用),需要自定义(输出,计算....)
 */
public class Demo01Consumer {
    /*
        定义一个方法
        方法的参数传递一个字符串的姓名
        方法的参数传递Consumer接口,泛型使用String
        可以使用Consumer接口消费字符串的姓名
     */
    public static void method(String name, Consumer<String> con){
        con.accept(name);
    }

    public static void main(String[] args) {
        //调用method方法,传递字符串姓名,方法的另一个参数是Consumer接口,是一个函数式接口,所以可以传递Lambda表达式
        method("赵丽颖",(String name)->{
            //对传递的字符串进行消费
            //消费方式:直接输出字符串
            //System.out.println(name);

            //消费方式:把字符串进行反转输出
            String reName = new StringBuffer(name).reverse().toString();
            System.out.println(reName);
        });
    }
}
```

**andThen**
如果一个方法的参数和返回值全都是 Consumer 类型，那么就可以实现效果：消费数据的时候，首先做一个操作，然后再做一个操作，实现组合
```java
import java.util.function.Consumer;

/*
   Consumer接口的默认方法andThen
   作用:需要两个Consumer接口,可以把两个Consumer接口组合到一起,在对数据进行消费

   例如:
    Consumer<String> con1
    Consumer<String> con2
    String s = "hello";
    con1.accept(s);
    con2.accept(s);
    连接两个Consumer接口  再进行消费
    con1.andThen(con2).accept(s); 谁写前边谁先消费
*/
public class Demo02AndThen {
    //定义一个方法,方法的参数传递一个字符串和两个Consumer接口,Consumer接口的泛型使用字符串
    public static void method(String s, Consumer<String> con1 ,Consumer<String> con2){
        //con1.accept(s);
        //con2.accept(s);
        //使用andThen方法,把两个Consumer接口连接到一起,在消费数据
        con1.andThen(con2).accept(s);//con1连接con2,先执行con1消费数据,在执行con2消费数据
    }

    public static void main(String[] args) {
        //调用method方法,传递一个字符串,两个Lambda表达式
        method("Hello",
                (t)->{
                    //消费方式:把字符串转换为大写输出
                    System.out.println(t.toUpperCase());
                },
                (t)->{
                    //消费方式:把字符串转换为小写输出
                    System.out.println(t.toLowerCase());
                });
    }
}
```

```java
import java.util.function.Consumer;

/*
    练习:
        字符串数组当中存有多条信息，请按照格式“姓名：XX。性别：XX。”的格式将信息打印出来。
        要求将打印姓名的动作作为第一个Consumer接口的Lambda实例，
        将打印性别的动作作为第二个Consumer接口的Lambda实例，
        将两个Consumer接口按照顺序“拼接”到一起。
 */
public class Demo03Test {
    //定义一个方法,参数传递String类型的数组和两个Consumer接口,泛型使用String
    public static void printInfo(String[] arr, Consumer<String> con1,Consumer<String> con2){
        //遍历字符串数组
        for (String message : arr) {
            //使用andThen方法连接两个Consumer接口,消费字符串
            con1.andThen(con2).accept(message);
        }
    }

    public static void main(String[] args) {
        //定义一个字符串类型的数组
        String[] arr = { "迪丽热巴,女", "古力娜扎,女", "马尔扎哈,男" };

        //调用printInfo方法,传递一个字符串数组,和两个Lambda表达式
        printInfo(arr,(message)->{
            //消费方式:对message进行切割,获取姓名,按照指定的格式输出
            String name = message.split(",")[0];
            System.out.print("姓名: "+name);
        },(message)->{
            //消费方式:对message进行切割,获取年龄,按照指定的格式输出
            String age = message.split(",")[1];
            System.out.println("。年龄: "+age+"。");
        });
    }
}
```
##### Predicate接口
有时候我们需要对某种类型的数据进行判断，从而得到一个boolean值结果。这时可以使用java.util.function.Predicate<T> 接口	
```java
import java.util.function.Predicate;

/*
    java.util.function.Predicate<T>接口
    作用:对某种数据类型的数据进行判断,结果返回一个boolean值

    Predicate接口中包含一个抽象方法：
        boolean test(T t):用来对指定数据类型数据进行判断的方法
            结果:
                符合条件,返回true
                不符合条件,返回false
*/
public class Demo01Predicate {
    /*
        定义一个方法
        参数传递一个String类型的字符串
        传递一个Predicate接口,泛型使用String
        使用Predicate中的方法test对字符串进行判断,并把判断的结果返回
     */
    public static boolean checkString(String s, Predicate<String> pre){
        return  pre.test(s);
    }

    public static void main(String[] args) {
        //定义一个字符串
        String s = "abcdef";

        //调用checkString方法对字符串进行校验,参数传递字符串和Lambda表达式
        /*boolean b = checkString(s,(String str)->{
            //对参数传递的字符串进行判断,判断字符串的长度是否大于5,并把判断的结果返回
            return str.length()>5;
        });*/

        //优化Lambda表达式
        boolean b = checkString(s,str->str.length()>5);
        System.out.println(b);
    }
}
```

Predicate接口中有一个方法`and`,表示并且关系,也可以用于连接两个判断条件
```java
import java.util.function.Predicate;

/*
    逻辑表达式:可以连接多个判断的条件
    &&:与运算符,有false则false
    ||:或运算符,有true则true
    !:非(取反)运算符,非真则假,非假则真

    需求:判断一个字符串,有两个判断的条件
        1.判断字符串的长度是否大于5
        2.判断字符串中是否包含a
    两个条件必须同时满足,我们就可以使用&&运算符连接两个条件

    Predicate接口中有一个方法and,表示并且关系,也可以用于连接两个判断条件
    default Predicate<T> and(Predicate<? super T> other) {
        Objects.requireNonNull(other);
        return (t) -> this.test(t) && other.test(t);
    }
    方法内部的两个判断条件,也是使用&&运算符连接起来的
 */
public class Demo02Predicate_and {
    /*
        定义一个方法,方法的参数,传递一个字符串
        传递两个Predicate接口
            一个用于判断字符串的长度是否大于5
            一个用于判断字符串中是否包含a
            两个条件必须同时满足
     */
    public static boolean checkString(String s, Predicate<String> pre1,Predicate<String> pre2){
        //return pre1.test(s) && pre2.test(s);
        return pre1.and(pre2).test(s);//等价于return pre1.test(s) && pre2.test(s);
    }

    public static void main(String[] args) {
        //定义一个字符串
        String s = "abcdef";
        //调用checkString方法,参数传递字符串和两个Lambda表达式
        boolean b = checkString(s,(String str)->{
            //判断字符串的长度是否大于5
            return str.length()>5;
        },(String str)->{
            //判断字符串中是否包含a
            return str.contains("a");
        });
        System.out.println(b);
    }
}
```

Predicate接口中有一个方法`or`,表示或者关系,也可以用于连接两个判断条件
```java
import java.util.function.Predicate;

/*
     需求:判断一个字符串,有两个判断的条件
        1.判断字符串的长度是否大于5
        2.判断字符串中是否包含a
    满足一个条件即可,我们就可以使用||运算符连接两个条件

    Predicate接口中有一个方法or,表示或者关系,也可以用于连接两个判断条件
    default Predicate<T> or(Predicate<? super T> other) {
        Objects.requireNonNull(other);
        return (t) -> test(t) || other.test(t);
    }
    方法内部的两个判断条件,也是使用||运算符连接起来的
 */
public class Demo03Predicate_or {
    /*
            定义一个方法,方法的参数,传递一个字符串
            传递两个Predicate接口
                一个用于判断字符串的长度是否大于5
                一个用于判断字符串中是否包含a
                满足一个条件即可
         */
    public static boolean checkString(String s, Predicate<String> pre1, Predicate<String> pre2){
        //return pre1.test(s) || pre2.test(s);
        return  pre1.or(pre2).test(s);//等价于return pre1.test(s) || pre2.test(s);
    }

    public static void main(String[] args) {
        //定义一个字符串
        String s = "bc";
        //调用checkString方法,参数传递字符串和两个Lambda表达式
        boolean b = checkString(s,(String str)->{
            //判断字符串的长度是否大于5
            return str.length()>5;
        },(String str)->{
            //判断字符串中是否包含a
            return str.contains("a");
        });
        System.out.println(b);
    }
}
```

Predicate接口中有一个方法`negate`,表示取反的意思
```java
import java.util.function.Predicate;

/*
    需求:判断一个字符串长度是否大于5
        如果字符串的长度大于5,那返回false
        如果字符串的长度不大于5,那么返回true
    所以我们可以使用取反符号!对判断的结果进行取反

    Predicate接口中有一个方法negate,表示取反的意思
    default Predicate<T> negate() {
        return (t) -> !test(t);
    }
 */
public class Demo04Predicate_negate {
    /*
           定义一个方法,方法的参数,传递一个字符串
           使用Predicate接口判断字符串的长度是否大于5
    */
    public static boolean checkString(String s, Predicate<String> pre){
        //return !pre.test(s);
        return  pre.negate().test(s);//等效于return !pre.test(s);
    }

    public static void main(String[] args) {
        //定义一个字符串
        String s = "abc";
        //调用checkString方法,参数传递字符串和Lambda表达式
        boolean b = checkString(s,(String str)->{
            //判断字符串的长度是否大于5,并返回结果
            return str.length()>5;
        });
        System.out.println(b);
    }
}
```

 练习：集合信息筛选
```java
import java.util.ArrayList;
import java.util.function.Predicate;

/*
    练习：集合信息筛选
    数组当中有多条“姓名+性别”的信息如下，
    String[] array = { "迪丽热巴,女", "古力娜扎,女", "马尔扎哈,男", "赵丽颖,女" };
    请通过Predicate接口的拼装将符合要求的字符串筛选到集合ArrayList中，
    需要同时满足两个条件：
        1. 必须为女生；
        2. 姓名为4个字。

    分析:
        1.有两个判断条件,所以需要使用两个Predicate接口,对条件进行判断
        2.必须同时满足两个条件,所以可以使用and方法连接两个判断条件
 */
public class Demo05Test {
    /*
        定义一个方法
        方法的参数传递一个包含人员信息的数组
        传递两个Predicate接口,用于对数组中的信息进行过滤
        把满足条件的信息存到ArrayList集合中并返回
     */
    public static ArrayList<String> filter(String[] arr,Predicate<String> pre1,Predicate<String> pre2){
        //定义一个ArrayList集合,存储过滤之后的信息
        ArrayList<String> list = new ArrayList<>();
        //遍历数组,获取数组中的每一条信息
        for (String s : arr) {
            //使用Predicate接口中的方法test对获取到的字符串进行判断
            boolean b = pre1.and(pre2).test(s);
            //对得到的布尔值进行判断
            if(b){
                //条件成立,两个条件都满足,把信息存储到ArrayList集合中
                list.add(s);
            }
        }
        //把集合返回
        return list;
    }

    public static void main(String[] args) {
        //定义一个储存字符串的数组
        String[] array = { "迪丽热巴,女", "古力娜扎,女", "马尔扎哈,男", "赵丽颖,女" };
        //调用filter方法,传递字符串数组和两个Lambda表达式
        ArrayList<String> list = filter(array,(String s)->{
            //获取字符串中的性别,判断是否为女
           return s.split(",")[1].equals("女");
        },(String s)->{
            //获取字符串中的姓名,判断长度是否为4个字符
           return s.split(",")[0].length()==4;
        });
        //遍历集合
        for (String s : list) {
            System.out.println(s);
        }
    }
}
```


##### Function接口

转换型接口Function

java.util.function.Function<T,R> 接口用来根据一个类型的数据得到另一个类型的数据，前者称为前置条件，后者称为后置条件
Function 接口中最主要的抽象方法为： R apply(T t) ，根据类型T的参数获取类型R的结果。使用的场景例如：将 String 类型转换为 Integer 类型。

```java
import java.util.function.Function;
/*
    java.util.function.Function<T,R>接口用来根据一个类型的数据得到另一个类型的数据，
        前者称为前置条件，后者称为后置条件。
    Function接口中最主要的抽象方法为：R apply(T t)，根据类型T的参数获取类型R的结果。
        使用的场景例如：将String类型转换为Integer类型。
 */
public class Demo01Function {
    /*
        定义一个方法
        方法的参数传递一个字符串类型的整数
        方法的参数传递一个Function接口,泛型使用<String,Integer>
        使用Function接口中的方法apply,把字符串类型的整数,转换为Integer类型的整数
     */
    public static void change(String s, Function<String,Integer> fun){
        //Integer in = fun.apply(s);
        int in = fun.apply(s);//自动拆箱 Integer->int
        System.out.println(in);
    }

    public static void main(String[] args) {
        //定义一个字符串类型的整数
        String s = "1234";
        //调用change方法,传递字符串类型的整数,和Lambda表达式
        change(s,(String str)->{
            //把字符串类型的整数,转换为Integer类型的整数返回
            return Integer.parseInt(str);
        });
        //优化Lambda
        change(s,str->Integer.parseInt(str));
    }
}
```

Function接口中的默认方法`andThen`:用来进行组合操作
```java
import java.util.function.Function;

/*
    Function接口中的默认方法andThen:用来进行组合操作

    需求:
        把String类型的"123",转换为Inteter类型,把转换后的结果加10
        把增加之后的Integer类型的数据,转换为String类型

    分析:
        转换了两次
        第一次是把String类型转换为了Integer类型
            所以我们可以使用Function<String,Integer> fun1
                Integer i = fun1.apply("123")+10;
        第二次是把Integer类型转换为String类型
            所以我们可以使用Function<Integer,String> fun2
                String s = fun2.apply(i);
        我们可以使用andThen方法,把两次转换组合在一起使用
            String s = fun1.andThen(fun2).apply("123");
            fun1先调用apply方法,把字符串转换为Integer
            fun2再调用apply方法,把Integer转换为字符串
 */
public class Demo02Function_andThen {
    /*
        定义一个方法
        参数串一个字符串类型的整数
        参数再传递两个Function接口
            一个泛型使用Function<String,Integer>
            一个泛型使用Function<Integer,String>
     */
    public static void change(String s, Function<String,Integer> fun1,Function<Integer,String> fun2){
        String ss = fun1.andThen(fun2).apply(s);
        System.out.println(ss);
    }

    public static void main(String[] args) {
        //定义一个字符串类型的整数
        String s = "123";
        //调用change方法,传递字符串和两个Lambda表达式
        change(s,(String str)->{
            //把字符串转换为整数+10
            return Integer.parseInt(str)+10;
        },(Integer i)->{
            //把整数转换为字符串
            return i+"";
        });

        //优化Lambda表达式
        change(s,str->Integer.parseInt(str)+10,i->i+"");
    }
}
```

练习：自定义函数模型拼接
```java
import java.util.function.Function;

/*
    练习：自定义函数模型拼接
    题目
    请使用Function进行函数模型的拼接，按照顺序需要执行的多个函数操作为：
        String str = "赵丽颖,20";

    分析:
    1. 将字符串截取数字年龄部分，得到字符串；
        Function<String,String> "赵丽颖,20"->"20"
    2. 将上一步的字符串转换成为int类型的数字；
        Function<String,Integer> "20"->20
    3. 将上一步的int数字累加100，得到结果int数字。
        Function<Integer,Integer> 20->120
 */
public class Demo03Test {
    /*
        定义一个方法
        参数传递包含姓名和年龄的字符串
        参数再传递3个Function接口用于类型转换
     */
    public static int change(String s, Function<String,String> fun1,
                             Function<String,Integer> fun2,Function<Integer,Integer> fun3){
        //使用andThen方法把三个转换组合到一起
        return fun1.andThen(fun2).andThen(fun3).apply(s);
    }

    public static void main(String[] args) {
        //定义一个字符串
        String str = "赵丽颖,20";
        //调用change方法,参数传递字符串和3个Lambda表达式
        int num = change(str,(String s)->{
            //"赵丽颖,20"->"20"
           return s.split(",")[1];
        },(String s)->{
            //"20"->20
            return Integer.parseInt(s);
        },(Integer i)->{
            //20->120
            return i+100;
        });
        System.out.println(num);
    }
}
```

### Stream流
#### 概述

Java 8 API添加了一个新的抽象称为流Stream，可以让你以一种声明的方式处理数据。

Stream 使用一种类似用 SQL 语句从数据库查询数据的直观方式来提供一种对 Java 集合运算和表达的高阶抽象。

Stream API可以极大提高Java程序员的生产力，让程序员写出高效率、干净、简洁的代码。

这种风格将要处理的元素集合看作一种流， 流在管道中传输， 并且可以在管道的节点上进行处理， 比如筛选， 排序，聚合等。

元素流在管道中经过中间操作（intermediate operation）的处理，最后由最终操作(terminal operation)得到前面处理的结果。
```java
+--------------------+       +------+   +------+   +---+   +-------+
| stream of elements +-----> |filter+-> |sorted+-> |map+-> |collect|
+--------------------+       +------+   +------+   +---+   +-------+


//以上的流程转换为 Java 代码为：
List<Integer> transactionsIds = 
widgets.stream()
             .filter(b -> b.getColor() == RED)
             .sorted((x,y) -> x.getWeight() - y.getWeight())
             .mapToInt(Widget::getWeight)
             .sum();
```

什么是 Stream？
	Stream（流）是一个来自数据源的元素`队列`并支持聚合操作
	元素是特定类型的对象，形成一个队列。 Java中的Stream并不会存储元素，而是按需计算。
	`数据源`流的来源。 可以是`集合`，`数组`，`I/O` `channel`， 产生器generator 等。
	`聚合操作` 类似SQL语句一样的操作， 比如filter, map, reduce, find, match, sorted等。

和以前的Collection操作不同， Stream操作还有两个基础的特征：
`Pipelining`: 中间操作都会返回流对象本身。 这样多个操作可以串联成一个管道， 如同流式风格（fluent style）。 这样做可以对操作进行优化， 比如延迟执行(laziness)和短路( short-circuiting)。
`内部迭代`： 以前对集合遍历都是通过Iterator或者For-Each的方式, 显式的在集合外部进行迭代， 这叫做外部迭代。 Stream提供了内部迭代的方式， 通过访问者模式(Visitor)实现。

元素是特定类型的对象，形成一个队列。 Java中的Stream并不会存储元素，而是按需计算。
	数据源 流的来源。 可以是集合，数组 等

当使用一个流的时候，通常包括三个基本步骤：获取一个数据源（source）→ 数据转换→执行操作获取想要的结
	果，每次转换原有 Stream 对象不改变，返回一个新的 Stream 对象（可以有多次转换），这就允许对其操作可以
	像链条一样排列，变成一个管道。


**stream流的注意事项, 只能使用一次**

#### 传统的方式集合操作
```java
import java.util.ArrayList;
import java.util.List;

/*
    使用传统的方式,遍历集合,对集合中的数据进行过滤
 */
public class Demo01List {
    public static void main(String[] args) {
        //创建一个List集合,存储姓名
        List<String> list = new ArrayList<>();
        list.add("张无忌");
        list.add("周芷若");
        list.add("赵敏");
        list.add("张强");
        list.add("张三丰");

        //对list集合中的元素进行过滤,只要以张开头的元素,存储到一个新的集合中
        List<String> listA = new ArrayList<>();
        for(String s : list){
            if(s.startsWith("张")){
                listA.add(s);
            }
        }

        //对listA集合进行过滤,只要姓名长度为3的人,存储到一个新集合中
        List<String> listB = new ArrayList<>();
        for (String s : listA) {
            if(s.length()==3){
                listB.add(s);
            }
        }

        //遍历listB集合
        for (String s : listB) {
            System.out.println(s);
        }
    }
}
```

#### Stream流的方式集合操作

```java
import java.util.ArrayList;
import java.util.List;

/*
    使用Stream流的方式,遍历集合,对集合中的数据进行过滤
    Stream流是JDK1.8之后出现的
    关注的是做什么,而不是怎么做
 */
public class Demo02Stream {
    public static void main(String[] args) {

        //创建一个List集合,存储姓名
        List<String> list = new ArrayList<>();
        list.add("张无忌");
        list.add("周芷若");
        list.add("赵敏");
        list.add("张强");
        list.add("张三丰");

        //对list集合中的元素进行过滤,只要以张开头的元素,存储到一个新的集合中
        //对listA集合进行过滤,只要姓名长度为3的人,存储到一个新集合中
        //遍历listB集合
        list.stream()
                .filter(name->name.startsWith("张"))
            .filter(name->name.length()==3)
            .forEach(name-> System.out.println(name));
	}
}
```

#### 两种获取Stream流的方式
所有的 Collection 集合都可以通过 stream 默认方法获取流；
Stream 接口的静态方法 of 可以获取数组对应的流。 
```java
import java.util.*;
import java.util.stream.Stream;

/*
    java.util.stream.Stream<T>是Java 8新加入的最常用的流接口。（这并不是一个函数式接口。）
    获取一个流非常简单，有以下几种常用的方式：
        - 所有的Collection集合都可以通过stream默认方法获取流；
            default Stream<E> stream​()
        - Stream接口的静态方法of可以获取数组对应的流。
            static <T> Stream<T> of​(T... values)
            参数是一个可变参数,那么我们就可以传递一个数组
 */
public class Demo01GetStream {
    public static void main(String[] args) {

        //把集合转换为Stream流
        List<String> list = new ArrayList<>();
        Stream<String> stream1 = list.stream();

        Set<String> set = new HashSet<>();
        Stream<String> stream2 = set.stream();

        Map<String,String> map = new HashMap<>();
        //获取键,存储到一个Set集合中
        Set<String> keySet = map.keySet();
        Stream<String> stream3 = keySet.stream();

        //获取值,存储到一个Collection集合中
        Collection<String> values = map.values();
        Stream<String> stream4 = values.stream();

        //获取键值对(键与值的映射关系 entrySet)
        Set<Map.Entry<String, String>> entries = map.entrySet();
        Stream<Map.Entry<String, String>> stream5 = entries.stream();

        //把数组转换为Stream流
        Stream<Integer> stream6 = Stream.of(1, 2, 3, 4, 5);
        //可变参数可以传递数组
        Integer[] arr = {1,2,3,4,5};
        Stream<Integer> stream7 = Stream.of(arr);
        String[] arr2 = {"a","bb","ccc"};
        Stream<String> stream8 = Stream.of(arr2);
    }
}
```

#### Stream流中的常用方法
##### forEach
```java
import java.util.stream.Stream;

/*
    Stream流中的常用方法_forEach
    void forEach(Consumer<? super T> action);
    该方法接收一个Consumer接口函数，会将每一个流元素交给该函数进行处理。
    Consumer接口是一个消费型的函数式接口,可以传递Lambda表达式,消费数据

    简单记:
        forEach方法,用来遍历流中的数据
        是一个终结方法,遍历之后就不能继续调用Stream流中的其他方法
 */
public class Demo02Stream_forEach {
    public static void main(String[] args) {
        //获取一个Stream流
        Stream<String> stream = Stream.of("张三", "李四", "王五", "赵六", "田七");
        //使用Stream流中的方法forEach对Stream流中的数据进行遍历
        /*stream.forEach((String name)->{
            System.out.println(name);
        });*/

        stream.forEach(name->System.out.println(name));
    }
}
```

##### filter
```java
import java.util.stream.Stream;

/*
    Stream流中的常用方法_filter:用于对Stream流中的数据进行过滤
    Stream<T> filter(Predicate<? super T> predicate);
    filter方法的参数Predicate是一个函数式接口,所以可以传递Lambda表达式,对数据进行过滤
    Predicate中的抽象方法:
        boolean test(T t);
 */
public class Demo03Stream_filter {
    public static void main(String[] args) {
        //创建一个Stream流
        Stream<String> stream = Stream.of("张三丰", "张翠山", "赵敏", "周芷若", "张无忌");
        //对Stream流中的元素进行过滤,只要姓张的人
        Stream<String> stream2 = stream.filter((String name)->{return name.startsWith("张");});
        //遍历stream2流
        stream2.forEach(name-> System.out.println(name));

        /*
            Stream流属于管道流,只能被消费(使用)一次
            第一个Stream流调用完毕方法,数据就会流转到下一个Stream上
            而这时第一个Stream流已经使用完毕,就会关闭了
            所以第一个Stream流就不能再调用方法了
            IllegalStateException: stream has already been operated upon or closed
         */
        //遍历stream流
        stream.forEach(name-> System.out.println(name));
    }
}
```

##### map
```java
import java.util.stream.Stream;

/*
    Stream流中的常用方法_map:用于类型转换
    如果需要将流中的元素映射到另一个流中，可以使用map方法.
    <R> Stream<R> map(Function<? super T, ? extends R> mapper);
    该接口需要一个Function函数式接口参数，可以将当前流中的T类型数据转换为另一种R类型的流。
    Function中的抽象方法:
        R apply(T t);
 */
public class Demo04Stream_map {
    public static void main(String[] args) {
        //获取一个String类型的Stream流
        Stream<String> stream = Stream.of("1", "2", "3", "4");
        //使用map方法,把字符串类型的整数,转换(映射)为Integer类型的整数
        Stream<Integer> stream2 = stream.map((String s)->{
            return Integer.parseInt(s);
        });
        //遍历Stream2流
        stream2.forEach(i-> System.out.println(i));
    }
}
```

##### count
```java
import java.util.ArrayList;
import java.util.stream.Stream;

/*
    Stream流中的常用方法_count:用于统计Stream流中元素的个数
    long count();
    count方法是一个终结方法,返回值是一个long类型的整数
    所以不能再继续调用Stream流中的其他方法了
 */
public class Demo05Stream_count {
    public static void main(String[] args) {
        //获取一个Stream流
        ArrayList<Integer> list = new ArrayList<>();
        list.add(1);
        list.add(2);
        list.add(3);
        list.add(4);
        list.add(5);
        list.add(6);
        list.add(7);
        Stream<Integer> stream = list.stream();
        long count = stream.count();
        System.out.println(count);//7
    }
}
```

##### forEach
```java
import java.util.stream.Stream;

/*
    Stream流中的常用方法_limit:用于截取流中的元素
    limit方法可以对流进行截取，只取用前n个。方法签名：
    Stream<T> limit(long maxSize);
        参数是一个long型，如果集合当前长度大于参数则进行截取；否则不进行操作
    limit方法是一个延迟方法,只是对流中的元素进行截取,返回的是一个新的流,所以可以继续调用Stream流中的其他方法
 */
public class Demo06Stream_limit {
    public static void main(String[] args) {
        //获取一个Stream流
        String[] arr = {"美羊羊","喜洋洋","懒洋洋","灰太狼","红太狼"};
        Stream<String> stream = Stream.of(arr);
        //使用limit对Stream流中的元素进行截取,只要前3个元素
        Stream<String> stream2 = stream.limit(3);
        //遍历stream2流
        stream2.forEach(name-> System.out.println(name));
    }
}
```


##### forEach
```java
import java.util.stream.Stream;

/*
    Stream流中的常用方法_skip:用于跳过元素
    如果希望跳过前几个元素，可以使用skip方法获取一个截取之后的新流：
    Stream<T> skip(long n);
        如果流的当前长度大于n，则跳过前n个；否则将会得到一个长度为0的空流。
 */
public class Demo07Stream_skip {
    public static void main(String[] args) {
        //获取一个Stream流
        String[] arr = {"美羊羊","喜洋洋","懒洋洋","灰太狼","红太狼"};
        Stream<String> stream = Stream.of(arr);
        //使用skip方法跳过前3个元素
        Stream<String> stream2 = stream.skip(3);
        //遍历stream2流
        stream2.forEach(name-> System.out.println(name));
    }
}
```

##### concat
```java
import java.util.stream.Stream;

/*
    Stream流中的常用方法_concat:用于把流组合到一起
    如果有两个流，希望合并成为一个流，那么可以使用Stream接口的静态方法concat
    static <T> Stream<T> concat(Stream<? extends T> a, Stream<? extends T> b)
 */
public class Demo08Stream_concat {
    public static void main(String[] args) {
        //创建一个Stream流
        Stream<String> stream1 = Stream.of("张三丰", "张翠山", "赵敏", "周芷若", "张无忌");
        //获取一个Stream流
        String[] arr = {"美羊羊","喜洋洋","懒洋洋","灰太狼","红太狼"};
        Stream<String> stream2 = Stream.of(arr);
        //把以上两个流组合为一个流
        Stream<String> concat = Stream.concat(stream1, stream2);
        //遍历concat流
        concat.forEach(name-> System.out.println(name));
    }
}
```

#### 集合元素处理
```java
public class Person {
    private String name;

    public Person() {
    }

    public Person(String name) {
        this.name = name;
    }

    @Override
    public String toString() {
        return "Person{" +
                "name='" + name + '\'' +
                '}';
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }
}
```

**利用传统集合的方式, 对数据进行筛选、跳过..**
```java
import java.util.ArrayList;

/*
    练习:集合元素处理（传统方式）
        现在有两个ArrayList集合存储队伍当中的多个成员姓名，要求使用传统的for循环（或增强for循环）依次进行以下若干操作步骤：
        1. 第一个队伍只要名字为3个字的成员姓名；存储到一个新集合中。
        2. 第一个队伍筛选之后只要前3个人；存储到一个新集合中。
        3. 第二个队伍只要姓张的成员姓名；存储到一个新集合中。
        4. 第二个队伍筛选之后不要前2个人；存储到一个新集合中。
        5. 将两个队伍合并为一个队伍；存储到一个新集合中。
        6. 根据姓名创建Person对象；存储到一个新集合中。
        7. 打印整个队伍的Person对象信息。
 */
public class Demo01StreamTest {
    public static void main(String[] args) {
        //第一支队伍
        ArrayList<String> one = new ArrayList<>();
        one.add("迪丽热巴");
        one.add("宋远桥");
        one.add("苏星河");
        one.add("石破天");
        one.add("石中玉");
        one.add("老子");
        one.add("庄子");
        one.add("洪七公");
        //1. 第一个队伍只要名字为3个字的成员姓名；存储到一个新集合中。
        ArrayList<String> one1 = new ArrayList<>();
        for (String name : one) {
            if(name.length()==3){
                one1.add(name);
            }
        }
        //2. 第一个队伍筛选之后只要前3个人；存储到一个新集合中。
        ArrayList<String> one2 = new ArrayList<>();
        for (int i = 0; i <3 ; i++) {
            one2.add(one1.get(i));//i = 0,1,2
        }

        //第二支队伍
        ArrayList<String> two = new ArrayList<>();
        two.add("古力娜扎");
        two.add("张无忌");
        two.add("赵丽颖");
        two.add("张三丰");
        two.add("尼古拉斯赵四");
        two.add("张天爱");
        two.add("张二狗");
        //3. 第二个队伍只要姓张的成员姓名；存储到一个新集合中。
        ArrayList<String> two1 = new ArrayList<>();
        for (String name : two) {
            if(name.startsWith("张")){
                two1.add(name);
            }
        }
        //4. 第二个队伍筛选之后不要前2个人；存储到一个新集合中。
        ArrayList<String> two2 = new ArrayList<>();
        for (int i = 2; i <two1.size() ; i++) {
            two2.add(two1.get(i)); //i 不包含0 1
        }

        //5. 将两个队伍合并为一个队伍；存储到一个新集合中。
        ArrayList<String> all = new ArrayList<>();
        all.addAll(one2);
        all.addAll(two2);

        //6. 根据姓名创建Person对象；存储到一个新集合中。
        ArrayList<Person> list = new ArrayList<>();
        for (String name : all) {
            list.add(new Person(name));
        }

        //7. 打印整个队伍的Person对象信息。
        for (Person person : list) {
            System.out.println(person);
        }
    }
}
```

**集合元素处理（Stream方式）**
```java
import java.util.ArrayList;
import java.util.stream.Stream;

/*
    练习：集合元素处理（Stream方式）
    将上一题当中的传统for循环写法更换为Stream流式处理方式。
    两个集合的初始内容不变，Person类的定义也不变。
 */
public class Demo02StreamTest {
    public static void main(String[] args) {
        //第一支队伍
        ArrayList<String> one = new ArrayList<>();
        one.add("迪丽热巴");
        one.add("宋远桥");
        one.add("苏星河");
        one.add("石破天");
        one.add("石中玉");
        one.add("老子");
        one.add("庄子");
        one.add("洪七公");
        //1. 第一个队伍只要名字为3个字的成员姓名；存储到一个新集合中。
        //2. 第一个队伍筛选之后只要前3个人；存储到一个新集合中。
        Stream<String> oneStream = one.stream().filter(name -> name.length() == 3).limit(3);

        //第二支队伍
        ArrayList<String> two = new ArrayList<>();
        two.add("古力娜扎");
        two.add("张无忌");
        two.add("赵丽颖");
        two.add("张三丰");
        two.add("尼古拉斯赵四");
        two.add("张天爱");
        two.add("张二狗");
        //3. 第二个队伍只要姓张的成员姓名；存储到一个新集合中。
        //4. 第二个队伍筛选之后不要前2个人；存储到一个新集合中。
        Stream<String> twoStream = two.stream().filter(name -> name.startsWith("张")).skip(2);

        //5. 将两个队伍合并为一个队伍；存储到一个新集合中。
        //6. 根据姓名创建Person对象；存储到一个新集合中。
        //7. 打印整个队伍的Person对象信息。
        Stream.concat(oneStream,twoStream).map(name->new Person(name)).forEach(p-> System.out.println(p));
    }
}
```

### 方法引用
#### 引入
```java
/*
    定义一个打印的函数式接口
 */
@FunctionalInterface
public interface Printable {
    //定义字符串的抽象方法
    void print(String s);
}
```

```java
public class Demo01Printable {
    //定义一个方法,参数传递Printable接口,对字符串进行打印
    public static void printString(Printable p) {
        p.print("HelloWorld");
    }

    public static void main(String[] args) {
        //调用printString方法,方法的参数Printable是一个函数式接口,所以可以传递Lambda
        printString((s) -> {
            System.out.println(s);
        });

        /*
            分析:
                Lambda表达式的目的,打印参数传递的字符串
                把参数s,传递给了System.out对象,调用out对象中的方法println对字符串进行了输出
                注意:
                    1.System.out对象是已经存在的
                    2.println方法也是已经存在的
                所以我们可以使用方法引用来优化Lambda表达式
                可以使用System.out方法直接引用(调用)println方法
         */
        printString(System.out::println);
    }
}
```

#### 概述
方法引用通过方法的名字来指向一个方法。
方法引用可以使语言的构造更紧凑简洁，减少冗余代码。
方法引用使用一对冒号 `::` 。

```java
//我们在 Car 类中定义了 4 个方法作为例子来区分 Java 中 4 种不同方法的引用。

@FunctionalInterface
public interface Supplier<T> {
    T get();
}
 
class Car {
    //Supplier是jdk1.8的接口，这里和lamda一起使用了
    public static Car create(final Supplier<Car> supplier) {
        return supplier.get();
    }
 
    public static void collide(final Car car) {
        System.out.println("Collided " + car.toString());
    }
 
    public void follow(final Car another) {
        System.out.println("Following the " + another.toString());
    }
 
    public void repair() {
        System.out.println("Repaired " + this.toString());
    }
}
```

```java

//构造器引用：它的语法是Class::new，或者更一般的Class< T >::new实例如下：
final Car car = Car.create( Car::new );
final List< Car > cars = Arrays.asList( car );

//静态方法引用：它的语法是Class::static_method，实例如下：
cars.forEach( Car::collide );

//特定类的任意对象的方法引用：它的语法是Class::method实例如下：
cars.forEach( Car::repair );

//特定对象的方法引用：它的语法是instance::method实例如下：
final Car police = Car.create( Car::new );
cars.forEach( police::follow );
```

#### 对象名引用成员方法

对象名:: 方法名

```java
/*
    定义一个打印的函数式接口
 */
@FunctionalInterface
public interface Printable {
    //定义字符串的抽象方法
    void print(String s);
}
```

```java
public class MethodRerObject {
    //定义一个成员方法,传递字符串,把字符串按照大写输出
    public void printUpperCaseString(String str){
        System.out.println(str.toUpperCase());
    }
}
```

```java
/*
    通过对象名引用成员方法
    使用前提是对象名是已经存在的,成员方法也是已经存在
    就可以使用对象名来引用成员方法
 */
public class Demo01ObjectMethodReference {

    //定义一个方法,方法的参数传递Printable接口
    public static void printString(Printable p){
        p.print("Hello");
    }

    public static void main(String[] args) {

        //调用printString方法,方法的参数Printable是一个函数式接口,所以可以传递Lambda表达式
        printString((s)->{
            //创建MethodRerObject对象
            MethodRerObject obj = new MethodRerObject();
            //调用MethodRerObject对象中的成员方法printUpperCaseString,把字符串按照大写输出
            obj.printUpperCaseString(s);
        });

        /*
            使用方法引用优化Lambda
            对象是已经存在的MethodRerObject
            成员方法也是已经存在的printUpperCaseString
            所以我们可以使用对象名引用成员方法
         */
        //创建MethodRerObject对象
        MethodRerObject obj = new MethodRerObject();
        printString(obj::printUpperCaseString);
    }
}
```

#### 类名引用静态成员方法

类名::方法名

通过类名引用静态成员方法
类已经存在,静态成员方法也已经存在
就可以通过类名直接引用静态成员方法

```java
@FunctionalInterface
public interface Calcable {
    //定义一个抽象方法,传递一个整数,对整数进行绝对值计算并返回
    int calsAbs(int number);
}
```

```java
/*
    通过类名引用静态成员方法
    类已经存在,静态成员方法也已经存在
    就可以通过类名直接引用静态成员方法
 */
public class Demo01StaticMethodReference {
    //定义一个方法,方法的参数传递要计算绝对值的整数,和函数式接口Calcable
    public static int method(int number,Calcable c){
       return c.calsAbs(number);
    }

    public static void main(String[] args) {
        //调用method方法,传递计算绝对值得整数,和Lambda表达式
        int number = method(-10,(n)->{
            //对参数进行绝对值得计算并返回结果
            return Math.abs(n);
        });
        System.out.println(number);

        /*
            使用方法引用优化Lambda表达式
            Math类是存在的
            abs计算绝对值的静态方法也是已经存在的
            所以我们可以直接通过类名引用静态方法
         */
        int number2 = method(-10,Math::abs);
        System.out.println(number2);
    }
}
```

#### super引用父类的成员方法

super :: 方法名

```java
/*
    定义见面的函数式接口
 */
@FunctionalInterface
public interface Greetable {
    //定义一个见面的方法
    void greet();
}
```

```java
/*
    定义父类
 */
public class Human {
    //定义一个sayHello的方法
    public void sayHello(){
        System.out.println("Hello 我是Human!");
    }
}
```

```java
/*
    定义子类
 */
public class Man extends Human{
    //子类重写父类sayHello的方法
    @Override
    public void sayHello() {
        System.out.println("Hello 我是Man!");
    }

    //定义一个方法参数传递Greetable接口
    public void method(Greetable g){
        g.greet();
    }

    public void show(){
        //调用method方法,方法的参数Greetable是一个函数式接口,所以可以传递Lambda
        /*method(()->{
            //创建父类Human对象
            Human h = new Human();
            //调用父类的sayHello方法
            h.sayHello();
        });*/

        //因为有子父类关系,所以存在的一个关键字super,代表父类,所以我们可以直接使用super调用父类的成员方法
       /* method(()->{
            super.sayHello();
        });*/

      /*
           使用super引用类的成员方法
           super是���经存在的
           父类的成员方法sayHello也是已经存在的
           所以我们可以直接使用super引用父类的成员方法
       */
      method(super::sayHello);
    }

    public static void main(String[] args) {
        new Man().show();
    }
}
```

#### 通过this引用本类的成员方法
```java
/*
    定义一个富有的函数式接口
 */
@FunctionalInterface
public interface Richable {
    //定义一个想买什么就买什么的方法
    void buy();
}
```

```java
/*
    使用this引用本类的成员方法
 */
public class Husband {
    //定义一个买房子的方法
    public void buyHouse(){
        System.out.println("北京二环内买一套四合院!");
    }

    //定义一个结婚的方法,参数传递Richable接口
    public void marry(Richable r){
        r.buy();
    }

    //定义一个非常高兴的方法
    public void soHappy(){
        //调用结婚的方法,方法的参数Richable是一个函数式接口,传递Lambda表达式
       /* marry(()->{
            //使用this.成员方法,调用本类买房子的方法
            this.buyHouse();
        });*/

        /*
            使用方法引用优化Lambda表达式
            this是已经存在的
            本类的成员方法buyHouse也是已经存在的
            所以我们可以直接使用this引用本类的成员方法buyHouse
         */
        marry(this::buyHouse);
    }

    public static void main(String[] args) {
        new Husband().soHappy();
    }
}
```

#### 类的构造器(构造方法)引用

类名::new

```java
public class Person {
    private String name;

    public Person() {
    }

    public Person(String name) {
        this.name = name;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }
}
```

```java
/*
    定义一个创建Person对象的函数式接口
 */
@FunctionalInterface
public interface PersonBuilder {
    //定义一个方法,根据传递的姓名,创建Person对象返回
    Person builderPerson(String name);
}
```

```java
/*
    类的构造器(构造方法)引用
 */
public class Demo {
    //定义一个方法,参数传递姓名和PersonBuilder接口,方法中通过姓名创建Person对象
    public static void printName(String name,PersonBuilder pb){
        Person person = pb.builderPerson(name);
        System.out.println(person.getName());
    }

    public static void main(String[] args) {
        //调用printName方法,方法的参数PersonBuilder接口是一个函数式接口,可以传递Lambda
        printName("迪丽热巴",(String name)->{
            return new Person(name);
        });

        /*
            使用方法引用优化Lambda表达式
            构造方法new Person(String name) 已知
            创建对象已知 new
            就可以使用Person引用new创建对象
         */
        printName("古力娜扎",Person::new);//使用Person类的带参构造方法,通过传递的姓名创建对象
    }
}
```

#### 数组的构造器引用
```java
/*
    定义一个创建数组的函数式接口
 */
@FunctionalInterface
public interface ArrayBuilder {
    //定义一个创建int类型数组的方法,参数传递数组的长度,返回创建好的int类型数组
    int[] builderArray(int length);
}
```

```java
import java.util.Arrays;

/*
    数组的构造器引用
 */
public class Demo {
    /*
        定义一个方法
        方法的参数传递创建数组的长度和ArrayBuilder接口
        方法内部根据传递的长度使用ArrayBuilder中的方法创建数组并返回
     */
    public static int[] createArray(int length, ArrayBuilder ab){
        return  ab.builderArray(length);
    }

    public static void main(String[] args) {
        //调用createArray方法,传递数组的长度和Lambda表达式
        int[] arr1 = createArray(10,(len)->{
            //根据数组的长度,创建数组并返回
            return new int[len];
        });
        System.out.println(arr1.length);//10

        /*
            使用方法引用优化Lambda表达式
            已知创建的就是int[]数组
            数组的长度也是已知的
            就可以使用方法引用
            int[]引用new,根据参数传递的长度来创建数组
         */
        int[] arr2 =createArray(10,int[]::new);
        System.out.println(Arrays.toString(arr2));
        System.out.println(arr2.length);//10
    }
}
```

### 默认方法
#### 概述

默认方法就是接口可以有实现方法，而且不需要实现类去实现其方法。

我们只需在方法名前面加个 `default` 关键字即可实现默认方法。

> 为什么要有这个特性？

首先，之前的接口是个双刃剑，好处是面向抽象而不是面向具体编程，缺陷是，当需要修改接口时候，需要修改全部实现该接口的类，目前的 java 8 之前的集合框架没有 foreach 方法，通常能想到的解决办法是在JDK里给相关的接口添加新的方法及实现。然而，对于已经发布的版本，是没法在给接口添加新方法的同时不影响已有的实现。所以引进的默认方法。他们的目的是为了解决接口的修改与现有的实现不兼容的问题。

语法
```java
默认方法语法格式如下：
public interface Vehicle {
   default void print(){
      System.out.println("我是一辆车!");
   }
}
```

#### 多个默认方法

一个接口有默认方法，考虑这样的情况，一个类实现了多个接口，且这些接口有相同的默认方法，以下实例说明了这种情况的解决方法：
```java
public interface Vehicle {
   default void print(){
      System.out.println("我是一辆车!");
   }
}
 
public interface FourWheeler {
   default void print(){
      System.out.println("我是一辆四轮车!");
   }
}
```

第一个解决方案是创建自己的默认方法，来覆盖重写接口的默认方法：
```java
public class Car implements Vehicle, FourWheeler {
   default void print(){
      System.out.println("我是一辆四轮汽车!");
   }
}
```

第二种解决方案可以使用 super 来调用指定接口的默认方法：
```java
public class Car implements Vehicle, FourWheeler {
   public void print(){
      Vehicle.super.print();
   }
}
```

#### 静态默认方法
Java 8 的另一个特性是接口可以声明（并且可以提供实现）静态方法。例如：
```java
public interface Vehicle {
   default void print(){
      System.out.println("我是一辆车!");
   }
    // 静态方法
   static void blowHorn(){
      System.out.println("按喇叭!!!");
   }
}
```

#### 默认方法实例

我们可以通过以下代码来了解关于默认方法的使用，可以将代码放入 Java8Tester.java 文件中：
```java
public class Java8Tester {
   public static void main(String args[]){
      Vehicle vehicle = new Car();
      vehicle.print();
   }
}
 
interface Vehicle {
   default void print(){
      System.out.println("我是一辆车!");
   }
    
   static void blowHorn(){
      System.out.println("按喇叭!!!");
   }
}
 
interface FourWheeler {
   default void print(){
      System.out.println("我是一辆四轮车!");
   }
}
 
class Car implements Vehicle, FourWheeler {
   public void print(){
      Vehicle.super.print();
      FourWheeler.super.print();
      Vehicle.blowHorn();
      System.out.println("我是一辆汽车!");
   }
}
//执行以上脚本，输出结果为：
我是一辆车!
我是一辆四轮车!
按喇叭!!!
我是一辆汽车!
```
### Optional类
#### 概述

Optional 类是一个可以为`null`的容器对象。如果值存在则isPresent()方法会返回true，调用get()方法会返回该对象。

Optional 是个容器：它可以保存类型T的值，或者仅仅保存null。Optional提供很多有用的方法，这样我们就不用显式进行空值检测。

Optional 类的引入很好的解决空指针异常。

类声明
以下是一个 java.util.Optional<T> 类的声明：
```java
public final class Optional<T>
extends Object
```

#### Optional 实例

我们可以通过以下实例来更好的了解 Optional 类的使用：
```java

import java.util.Optional;
 
public class Java8Tester {
   public static void main(String args[]){
   
      Java8Tester java8Tester = new Java8Tester();
      Integer value1 = null;
      Integer value2 = new Integer(10);
        
      // Optional.ofNullable - 允许传递为 null 参数
      Optional<Integer> a = Optional.ofNullable(value1);
        
      // Optional.of - 如果传递的参数是 null，抛出异常 NullPointerException
      Optional<Integer> b = Optional.of(value2);
      System.out.println(java8Tester.sum(a,b));
   }
    
   public Integer sum(Optional<Integer> a, Optional<Integer> b){
    
      // Optional.isPresent - 判断值是否存在
        
      System.out.println("第一个参数值存在: " + a.isPresent());
      System.out.println("第二个参数值存在: " + b.isPresent());
        
      // Optional.orElse - 如果值存在，返回它，否则返回默认值
      Integer value1 = a.orElse(new Integer(0));
        
      //Optional.get - 获取值，值需要存在
      Integer value2 = b.get();
      return value1 + value2;
   }
}
//执行以上脚本，输出结果为：
第一个参数值存在: false
第二个参数值存在: true
10
````