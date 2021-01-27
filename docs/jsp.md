## jsp基本认识

什么是Java Server Pages

1. JSP全称Java Server Pages，是一种动态网页开发技术。它使用JSP标签在HTML网页中插入Java代码。标签通常以`<%开头以%>结束`。

2. `JSP是一种Java servlet`，主要用于实现Java web应用程序的用户界面部分。网页开发者们通过结合HTML代码、XHTML代码、XML元素以及嵌入JSP操作和命令来编写JSP。

3. JSP通过网页表单获取用户输入数据、访问数据库及其他数据源，然后动态地创建网页。

4. JSP标签有多种功能，比如访问数据库、记录用户选择信息、访问JavaBeans组件等，还可以在不同的网页中传递控制信息和共享信息。

   

Servlet的作用： 用java语言开发动态资源的技术！！！
Jsp的作用：用java语言（+html语言）开发动态资源的技术！！！ 



Jsp的特点

1. jsp的运行必须交给tomcat服务器！！！！
   tomcat的work目录： tomcat服务器存放jsp运行时的临时文件
2. jsp页面既可以写html代码，也可以写java代码。
   html页面不能写java代码 。而jsp页面可以写java代码
3. 体验jsp页面作用
   需求：显示当前时间到浏览器上



问题： 为什么Jsp就是servlet！！！

```java
jsp翻译的java文件时：
public final class _01_hello_jsp 
extends org.apache.jasper.runtime.HttpJspBase
implements org.apache.jasper.runtime.JspSourceDependent { }

HttpJspBase类：
public abstract class org.apache.jasper.runtime.HttpJspBase 
extends javax.servlet.http.HttpServlet 
implements javax.servlet.jsp.HttpJspPage { }
```

结论： Jsp就是一个servlet程序！！！
			servlet 的技术可以用在j sp 程序中
			jsp 的技术并不是全部适用于 servlet 程序！




## jsp实例

需求：显示当前时间到浏览器上

可以把jsp页面当做html页面在tomcat中访问！！！

```jsp
<%@ page language="java" import="java.util.*,java.text.*" pageEncoding="utf-8"%>

<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN">
<html>
  <head> 
    <title>第一个jsp页面</title>  
  </head>
  
  <body>
    <%
    	//写java代码
    	//获取当前时间
    	SimpleDateFormat sdf = new SimpleDateFormat();
    	String curDate = sdf.format(new Date());
    	//输出内容到浏览器
    	//response.getWriter().write("");
    	out.write("当前时间为2："+curDate);
     %>
  </body>
</html>
```



## Jsp的执行过程

问题： 访问http://localhost:8080/day/hello.jsp  如何显示效果？

1. 访问到01.hello.jsp页面，tomcat扫描到jsp文件，在%tomcat%/work把jsp文件翻译成java源文件
   `01.hello.jsp   ->   _01_hello_jsp.java` （翻译） 
2. tomcat服务器把java源文件编译成class字节码文件 （编译）
   `_01_hello_jsp.java  ->  _01_hello_jsp.class`
3. tomcat服务器构造`_01_hello_jsp`类对象
4. tomcat服务器调用`_01_hello_jsp`类里面方法，返回内容显示到浏览器。

注意：
		jsp文件修改了或jsp的临时文件被删除了，要重新走翻译（1）和编译（2）的过程


![jsp执行过程](E:\smile\java\images\jsp执行过程.jpg)



## jSP 处理

以下步骤表明了 Web 服务器是如何使用JSP来创建网页的：

- 就像其他普通的网页一样，浏览器发送一个 HTTP 请求给服务器。
- Web 服务器识别出这是一个对 JSP 网页的请求，并且将该请求传递给 JSP 引擎。通过使用 URL或者 .jsp 文件来完成。
- JSP 引擎从磁盘中载入 JSP 文件，然后将它们转化为 Servlet。这种转化只是简单地将所有模板文本改用 println() 语句，并且将所有的 JSP 元素转化成 Java 代码。
- JSP 引擎将 Servlet 编译成可执行类，并且将原始请求传递给 Servlet 引擎。
- Web 服务器的某组件将会调用 Servlet 引擎，然后载入并执行 Servlet 类。在执行过程中，Servlet 产
- HTML 格式的输出并将其内嵌于 HTTP response 中上交给 Web 服务器。
- Web 服务器以静态 HTML 网页的形式将 HTTP response 返回到您的浏览器中。
- 最终，Web 浏览器处理 HTTP response 中动态产生的HTML网页，就好像在处理静态网页一样。

一般情况下，JSP 引擎会检查 JSP 文件对应的 Servlet 是否已经存在，并且检查 JSP 文件的修改日期是否早于 Servlet。如果 JSP 文件的修改日期早于对应的 Servlet，那么容器就可以确定 JSP 文件没有被修改过并且 Servlet 有效。



## Jsp的生命周期

JSP生命周期就是从创建到销毁的整个过程，类似于servlet生命周期，区别在于JSP生命周期还包括将JSP文件编译成servlet。

以下是JSP生命周期中所走过的几个阶段：

- **编译阶段：**

  servlet容器编译servlet源文件，生成servlet类

- 初始化阶段：

  加载与JSP对应的servlet类，创建其实例，并调用它的初始化方法

- 执行阶段：

  调用与JSP对应的servlet实例的服务方法

- 销毁阶段：

  调用与JSP对应的servlet实例的销毁方法，然后销毁servlet实例

![JSP 生命周期](E:\smile\java\images\JSP 生命周期.jpg)



### JSP编译

当浏览器请求JSP页面时，JSP引擎会首先去检查是否需要编译这个文件。如果这个文件没有被编译过，或者在上次编译后被更改过，则编译这个JSP文件。

编译的过程包括三个步骤：

- 解析JSP文件。
- 将JSP文件转为servlet。
- 编译servlet。

------

### JSP初始化

容器载入JSP文件后，它会在为请求提供任何服务前调用jspInit()方法。如果您需要执行自定义的JSP初始化任务，复写jspInit()方法就行了，就像下面这样：

```java
public void jspInit(){
  // 初始化代码
}
```

一般来讲程序只初始化一次，servlet也是如此。通常情况下您可以在jspInit()方法中初始化数据库连接、打开文件和创建查询表。

------

### JSP执行

这一阶段描述了JSP生命周期中一切与请求相关的交互行为，直到被销毁。

当JSP网页完成初始化后，JSP引擎将会调用_jspService()方法。

_jspService()方法需要一个HttpServletRequest对象和一个HttpServletResponse对象作为它的参数，就像下面这样：

```java
void _jspService(HttpServletRequest request,
                 HttpServletResponse response)
{
   // 服务端处理代码
}
```

_jspService()方法在每个request中被调用一次并且负责产生与之相对应的response，并且它还负责产生所有7个HTTP方法的回应，比如GET、POST、DELETE等等。

------

### JSP清理

JSP生命周期的销毁阶段描述了当一个JSP网页从容器中被移除时所发生的一切。

jspDestroy()方法在JSP中等价于servlet中的销毁方法。当您需要执行任何清理工作时复写jspDestroy()方法，比如释放数据库连接或者关闭文件夹等等。

jspDestroy()方法的格式如下：

```java
public void jspDestroy()
{
   // 清理代码
}
```

### 实例

JSP生命周期代码实例如下所示：

```jsp
<%@ page language="java" contentType="text/html; charset=UTF-8"
    pageEncoding="UTF-8"%>
<html>
    <head>
    <title>life.jsp</title>
    </head>
<body>

        <%! 
          private int initVar=0;
          private int serviceVar=0;
          private int destroyVar=0;
        %>

        <%!
          public void jspInit(){
            initVar++;
            System.out.println("jspInit(): JSP被初始化了"+initVar+"次");
          }
          public void jspDestroy(){
            destroyVar++;
            System.out.println("jspDestroy(): JSP被销毁了"+destroyVar+"次");
          }
        %>

        <%
          serviceVar++;
          System.out.println("_jspService(): JSP共响应了"+serviceVar+"次请求");

          String content1="初始化次数 : "+initVar;
          String content2="响应客户请求次数 : "+serviceVar;
          String content3="销毁次数 : "+destroyVar;
        %>
        <p><%=content1 %></p>
        <p><%=content2 %></p>
        <p><%=content3 %></p>
    </body>
</html>

浏览器打开该页面，输出结果为：
    初始化次数 : 1
    响应客户请求次数 : 1
    销毁次数 : 0
```

------



Servlet的生命周期：

```
1）构造方法（第1次访问）
2）init方法（第1次访问）
3）service方法
4）destroy方法		
```

Jsp的生命周期

```
1）翻译： jsp->java文件
2）编译： java文件->class文件（servlet程序）
3）构造方法（第1次访问）
4）init方法（第1次访问）：_jspInit()
5）service方法：_jspService()
6）destroy方法：_jspDestroy()
```



## Jsp语法

### Jsp模板

​		jsp页面中的html代码就是jsp的模板

###  Jsp表达式

​		语法：`<%=变量或表达式%>`
​		作用： 向浏览器`输出变量的值`或`表达式计算的结果`
​		注意：		
​			1）表达式的原理就是翻译成`out.print(“变量” )`;通过该方法向浏览器写出内容
​			2）表达式后面不需要带分号结束。

### Jsp的脚本

​		语法：`<%java代码 %>`
​		作用： 执行java代码	
​		注意：原理把脚本中java代码原封不动拷贝到_jspService方法中执行。



### Jsp的声明

​		语法：`<%! 变量或方法 %>`
​		作用： 声明jsp的变量或方法
​		注意:  变量翻译成成员变量，方法翻译成成员方法。
​				

### jsp的注释

​				语法：` <%!--  jsp注释  --%>`
​				注意:  html的注释会被翻译和执行。而jsp的注释不能被翻译和执行。					

### 实例

```jsp
<%@ page language="java" import="java.util.*" pageEncoding="utf-8"%>

<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN">
<html>
  <head> 
    <title>Jsp语法</title>  
  </head>
  
  <body>
      
  	<!-- jsp表达式  -->
  	<%
  		//变量
  		String name = "eric";
  		int a = 10;
  		int b =20;
  	%>
  	<%=name %>  
  	<br/>
  	<%=(a-b) %>  
      
  	<hr/>
      
  	<!-- jsp脚本  -->
  	<%
  		//生成随机数
  		Random ran = new Random();
  		float num = ran.nextFloat();
  	 %>
  	 随机小数：<%=num %>
  	 <hr/>
      
  	 <!-- 穿插html代码 -->
  	 <%
  	 	for(int i=1;i<=6;i++){ 	
  	 %>
  	 	<h<%=i %>>标题<%=i %></h<%=i %>>
  	 <%
  	  }
  	 %>
      
  	  <hr/>
      
  	  <!-- 练习： 使用脚本和html代码显示99乘法表 -->
  	 <%
  	 	for(int i=1;i<=9;i++){//行
  	 		for(int j=1;j<=i;j++){//公式
  	 %>			
  	 		<%=i %> x <%=j %>=<%=(i*j) %>&nbsp;
  	 	<%			
  	 		}
  	 	%>
  	 	   <br/>
  	 	<% 	
  	 	}
  	    %>
      
  	  <%
  	  	String age = "20";
  	  	
  	  	/* 脚本中不能声明方法
  	  	public String getAge(){
  	  		return age;
  	  	}
  	  	*/
  	   %>
  	  
  	  
  	  <!-- jsp声明 -->
  	  <%!
  	   //变量
  	  	String name = "jacky";
  	  	
  	  	public String getName(){
  	  		return name;
  	  	}
  	  	
  	  	/*jsp声明中不能重复定义翻译好的一些方法
  	  	public void _jspInit(){
  	  	
  	  	}
  	  	*/
  	   %>
  	   
  	  <!-- html注释 -->
  	  <%-- <jsp:forward page="/01.hello.jsp"></jsp:forward> --%>
  	  <%-- jsp注释 --%>
  	  
  </body>
</html>
```



## Jsp的三大指令

JSP指令用来设置与整个JSP页面相关的属性。

JSP指令语法格式：

```jsp
<%@ directive attribute="value" %>
```

这里有三种指令标签：

| **指令**           | **描述**                                                  |
| :----------------- | :-------------------------------------------------------- |
| <%@ page ... %>    | 定义页面的依赖属性，比如脚本语言、error页面、缓存需求等等 |
| <%@ include ... %> | 包含其他文件                                              |
| <%@ taglib ... %>  | 引入标签库的定义，可以是自定义标签                        |



### include指令

​		作用： 在当前页面用于包含其他页面
​		语法： `<%@include file="common/header.jsp"%>`
​		注意：
​			1）原理是把被包含的页面（header.jsp）的内容翻译到包含页面(index.jsp)中,合并成翻译成一					个java源文件，再编译运行！！，这种包含叫静态包含（源码包含）
​			2）如果使用静态包含，被包含页面中不需要出现全局的html标签了！！！（如html、head、						body）

```jsp
<%@ page language="java" import="java.util.*" pageEncoding="utf-8"%>

<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN">
<html>
  <head> 
    <title>网站首页</title>  
  </head>
  
  <body>
  <%@include file="common/header.jsp"%>
    首页的内容xxxxxxx
  </body>
</html>
```

WebRoot/common/header.jsp

```jsp
<%@ page language="java" import="java.util.*" pageEncoding="utf-8"%>
   头部页面的内容wwwwwww
```



### page指令

​		作用： 告诉tomcat服务器如何翻译jsp文件

```jsp
<%@ page 
    language="java"                 --告诉服务器使用什么动态语言来翻译jsp文件
    import="java.util.*"            --告诉服务器java文件使用什么包导入包，多个包之间用逗号分割
    
    pageEncoding="utf-8"            --告诉服务器使用什么编码翻译jsp文件（成java文件）
    contentType="text/html; charset=utf-8" -- 服务器发送浏览器的数据类型和内容编码
    											注意：在开发工具中，以后只需要设置														pageEncoding即可解决中文乱码问题
    
    errorPage="error.jsp"            -- 指定当前jsp页面的错误处理页面。
    isErrorPage="false"              -- 指定当前页面是否为错误处理页面。false，不是错误处理页										面，则不能使用exception内置对象；true，是错误处理											页面，可以使用exception内置对象。
    
    buffer="8kb"                     --  jsp页面的缓存区大小。  
    
    session="true"                   -- 是否开启session功能。false，不能用session内置对象；											true，可以使用session内置对象。
    
    isELIgnored="false"              -- 是否忽略EL表达式。
%>
```



```jsp
<!-- 全局错误处理页面配置 -->
<error-page>
    <error-code>500</error-code>
    <location>/common/500.jsp</location>
</error-page>
<error-page>
    <error-code>404</error-code>
    <location>/common/404.html</location>
</error-page>
```



```jsp
<%@ page 
	language="java" 
	import="java.util.*,java.text.*" 
	pageEncoding="utf-8"
	%>

<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN">

<html>
  <head> 
    <title>page指令</title>  
  </head>
  
  <body>
    <%
    	new Date();
    	new SimpleDateFormat();
     %>
    ixfosa
  </body>
</html>

```



### taglib指令

引入标签库的定义，可以是自定义标签



## JSP行为

JSP行为标签使用XML语法结构来控制servlet引擎。它能够动态插入一个文件，重用JavaBean组件，引导用户去另一个页面，为Java插件产生相关的HTML等等。

行为标签只有一种语法格式，它严格遵守XML标准：

```jsp
<jsp:action_name attribute="value" />
```

行为标签基本上是一些预先就定义好的函数，下表罗列出了一些可用的JSP行为标签：：

| **语法**        | **描述**                                                   |
| :-------------- | :--------------------------------------------------------- |
| jsp:include     | 用于在当前页面中包含静态或动态资源                         |
| jsp:useBean     | 寻找和初始化一个JavaBean组件                               |
| jsp:setProperty | 设置 JavaBean组件的值                                      |
| jsp:getProperty | 将 JavaBean组件的值插入到 output中                         |
| jsp:forward     | 从一个JSP文件向另一个文件传递一个包含用户请求的request对象 |
| jsp:plugin      | 用于在生成的HTML页面中包含Applet和JavaBean对象             |
| jsp:element     | 动态创建一个XML元素                                        |
| jsp:attribute   | 定义动态创建的XML元素的属性                                |
| jsp:body        | 定义动态创建的XML元素的主体                                |
| jsp:text        | 用于封装模板数据                                           |



## JSP隐含(内置)对象

JSP支持九个自动定义的变量，江湖人称隐含对象。这九个隐含对象的简介见下表：

| **对象**    | **描述**                                                     |
| :---------- | :----------------------------------------------------------- |
| request     | **HttpServletRequest**类的实例                               |
| response    | **HttpServletResponse**类的实例                              |
| out         | **PrintWriter**类的实例，用于把结果输出至网页上              |
| session     | **HttpSession**类的实例                                      |
| application | **ServletContext**类的实例，与应用上下文有关                 |
| config      | **ServletConfig**类的实例                                    |
| pageContext | **PageContext**类的实例，提供对JSP页面所有对象以及命名空间的访问 |
| page        | 类似于Java类中的this关键字                                   |
| Exception   | **Exception**类的对象，代表发生错误的JSP页面中对应的异常对象 |

### Out对象

out对象类型，JspWriter类，相当于带缓存的PrintWriter

```jsp
PrintWriter： 
	wrier(内容)： 直接向浏览器写出内容。

JspWriter
	writer(内容): 向jsp缓冲区写出内容
```



```jsp
<%@ page language="java" import="java.util.*" pageEncoding="utf-8" buffer="1kb"%>
	<%
		for(int i=1;i<=1024;i++){
    		out.write("a");
    	}
    	//查看缓存区大小
    	System.out.println("当前缓存区大小："+out.getBufferSize());
    	//查看缓存区剩余大小
    	System.out.println("缓存区剩余大小："+out.getRemaining());
    	//刷新缓存
    	//out.flush();
    	response.getWriter().write("123");
     %>
  </body>
</html>
```

​			当满足以下条件之一，缓冲区内容写出：
​					1）缓冲区满了
​					2）刷新缓存区
​					3）关闭缓存区
​					4）执行完毕jsp页面



### pageContext对象

​	pageContext对象的类型是PageContext，叫jsp的`上下文对象`

1. 可以获取其他八个内置对象

```java
public class 01_hello_jsp {
    public void _jspService(request,response){
        创建内置对象
        HttpSession session =....;
        ServletConfig config = ....;

        把8个经常使用的内置对象封装到PageContext对象中
        PageContext pageContext  = 封装；
        调用method1方法
      }

    public void method1(PageContext pageContext){
        希望使用内置对象
        从PageContext对象中获取其他8个内置对象
        JspWriter out =pageContext.getOut();
        HttpServletRequest rquest = 	pageContext.getRequest();
        ........
    }
}
```

​			使用场景： 在自定义标签的时候，PageContext对象频繁使用到！

2. 本身是一个域对象

```java
ServletContext context域
HttpServletRequet  request域
HttpSession    session域     -- Servlet学习的
PageContext   page域	        -- jsp学习的
```

​	作用： 保存数据和获取数据，用于共享数据

保存数据		

```java
1）默认情况下，保存到page域
	pageContext.setAttribute("name");
	
2）可以向四个域对象保存数据
	pageContext.setAttribute("name",域范围常量)
```

获取数据

```
1）默认情况下，从page域获取
	pageContext.getAttribute("name")
	
2) 可以从四个域中获取数据
	pageContext.getAttribute("name",域范围常量)

	域范围常量:
        PageContext.PAGE_SCOPE
        PageContext.REQUEST_SCOPE
        PageContext.SESSION_SCOPE
        PageContext.APPLICATION_SCOPE
        
3）自动在四个域中搜索数据
	pageContext.findAttribute("name");
	
顺序： page域 -> request域 -> session域- > context域（application域）
```

```jsp
<%@ page language="java" import="java.util.*" pageEncoding="utf-8" 
isErrorPage="true" 
session="true"
%>

<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
<html>
  <head> 
    <title>pageContext对象</title>  
  </head>
  
  <body>

    <%
    	//1）可以获取其他8个内置对象
    	//response.getWriter().write("是否相等？"+(out==pageContext.getOut()) +"<br/>");
    	//response.getWriter().write("是否相等？"+(session==pageContext.getSession()));
     %>
     
     
     <%--
     	2)pageContext作为域对象使用
     		2.1 可以往不同域对象中存取数据
      --%>
      <%
      	//保存数据。默认情况下，保存在page域中
      	pageContext.setAttribute("message","page's message");

      	pageContext.setAttribute("message","request's messsage",PageContext.REQUEST_SCOPE);//保存到request域中

      	pageContext.setAttribute("message","session's messsage",PageContext.SESSION_SCOPE);//保存到sessio域中

      	pageContext.setAttribute("message","application's messsage",PageContext.APPLICATION_SCOPE);//保存到context域中

      	//request.setAttribute("message","request's message"); 等价于上面的代码
 
       %>
       
       <%
       	//获取数据
       	//String message = (String)pageContext.getAttribute("message");
       	//out.write(message);
        %>
        <%--从request域中取出数据 --%>
        <%--
     	      原则： 
           	1）在哪个域中保存数据，就必须从哪个域取出数据！！！
         --%>
        <%=pageContext.getAttribute("message",PageContext.PAGE_SCOPE) %><br/>
        <%=pageContext.getAttribute("message",PageContext.REQUEST_SCOPE) %><br/>
        <%=pageContext.getAttribute("message",PageContext.SESSION_SCOPE) %><br/>
        <%=pageContext.getAttribute("message",PageContext.APPLICATION_SCOPE) %><br/>
        <hr/>
        <%--
        	findAttribute(): 在四个域自动搜索数据
        		顺序： page域 -> request域  -> session域 -> context域
        		
         --%>
         <%=pageContext.findAttribute("message") %>

         <% //request.getAttribute("message") %><br/>
      
      
      <%
      	//转发
      	// request.getRequestDispatcher("/03.pageContext2.jsp").forward(request,response);
      	//重定向
      	response.sendRedirect(request.getContextPath()+"/03.pageContext2.jsp");
       %>
  </body>
</html>
```

03.pageContext2.jsp

```jsp
<%@ page language="java" import="java.util.*" pageEncoding="utf-8"%>

<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN">
<html>
  <head> 
    <title>从四个域中获取数据</title>  
  </head>
  
  <body>
    page域：<%=pageContext.getAttribute("message",PageContext.PAGE_SCOPE) %><br/>
    request域： <%=pageContext.getAttribute("message",PageContext.REQUEST_SCOPE) %><br/>
    session域： <%=pageContext.getAttribute("message",PageContext.SESSION_SCOPE) %><br/>
    context域：<%=pageContext.getAttribute("message",PageContext.APPLICATION_SCOPE) %><br/>
  </body>
</html>
```

案例-猜字

```java
package gz.ixfosa.entity;

public class Student {
	private String name1;
	private int age;
	public String getName() {
		return name1;
	}
	public void setName(String name) {
		this.name1 = name;
	}
	public int getAge() {
		return age;
	}
	public void setAge(int age) {
		this.age = age;
	}
	public Student(String name, int age) {
		super();
		this.name1 = name;
		this.age = age;
	}
	public Student() {
		super();
		// TODO Auto-generated constructor stub
	}
}
```

```java
package gz.ixfosa.exec;

import java.io.IOException;
import java.util.Random;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

public class GuessServlet extends HttpServlet {
    
	//产生一个幸运数字
	int answer;
	
	/**
	 * 新游戏方法。产生一个新的幸运数字
	 */
	public void newGame(){
		Random random = new Random();
		answer = random.nextInt(30);
	}
	
	public GuessServlet(){
		//第一次访问
		newGame();
	}
	
	public void doGet(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, IOException {
		response.setContentType("text/html;charset=utf-8");
		
		//1.接收输入的数字
		String lucyNoStr = request.getParameter("lucyNo");
		System.out.println("答案："+answer);
		Integer lucyNo = null;
		//2.判断幸运数字和用户的数字
		//2.1 把用户输入的数字转成整数
		if(lucyNoStr!=null || !lucyNoStr.equals("")){
			lucyNo = Integer.parseInt(lucyNoStr);
		}
		
		//标记记录当前竞猜的次数
		Integer times = 1;//初始值
		
		//接收客户当前竞猜次数
		String timesStr = request.getParameter("times");
		if(timesStr!=null && !timesStr.equals("")){
			times = Integer.parseInt(timesStr)+1;
		}
		
		
		if(times<5){
			String msg = "";
			//比较
			if(lucyNo>answer){
				//大了
				msg = "可惜，大了点";
			}else if(lucyNo<answer){
				//小了
				msg = "可惜，小了点";
			}else if(lucyNo==answer){
				//等于,中奖
				msg = "恭喜你，中得1000000元现金大奖,请带身份证到xxx地方领奖！";
				times = null;
			}
			//把当前竞猜的次数放入域对象
			request.setAttribute("times", times);
			//把信息放入域对象中
			request.setAttribute("msg", msg);
		}else{
			//产生新的幸运数字
			newGame();
			//游戏结束
			response.getWriter().write("游戏结束。<a href='"+request.getContextPath()+"/05.guess.jsp'>再来一盘</a>");
			return;
		}
		//转发
		request.getRequestDispatcher("/05.guess.jsp").forward(request, response);
	}

	public void doPost(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, IOException {
		doGet(request, response);
	}
}
```



```jsp
<%@ page language="java" import="java.util.*" pageEncoding="utf-8"%>

<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN">
<html>
  <head> 
    <title>百万富翁数字竞猜游戏</title>  
  </head>
  
  <body>
  	<%
  		//从request域中取出信息
  		String msg = (String)request.getAttribute("msg");
  		if(msg!=null){
  			out.write("<font color='red'>"+msg+"</font>");
  		}
  	 %>
  	 <%
  	 	//获取竞猜次数
  	 	Integer times = (Integer)request.getAttribute("times");
  	 	if(times!=null){
  	 		out.write(",你还有"+(5-times)+"次机会！");
  	 	}
  	  %>
  	  
    <form action="/day13/GuessServlet" method="post">
    	请输入30以下的整数：<input type="text" name="lucyNo"/><br/>
    	<%
    		if(times!=null){
    	 %>
    	<input type="hidden" name="times" value="<%=times %>"/>
    	<%
    	}
    	 %>
    	<input type="submit" value="开始竞猜"/>
    </form>
  </body>
</html>
```

```xml
<?xml version="1.0" encoding="UTF-8"?>
<web-app version="2.5" 
	xmlns="http://java.sun.com/xml/ns/javaee" 
	xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" 
	xsi:schemaLocation="http://java.sun.com/xml/ns/javaee 
	http://java.sun.com/xml/ns/javaee/web-app_2_5.xsd">
  <servlet>
    <servlet-name>GuessServlet</servlet-name>
    <servlet-class>gz.ixfosa.exec.GuessServlet</servlet-class>
  </servlet>

  <servlet-mapping>
    <servlet-name>GuessServlet</servlet-name>
    <url-pattern>/GuessServlet</url-pattern>
  </servlet-mapping>
  <!-- 全局错误处理页面配置 -->
  <error-page>
  	<error-code>500</error-code>
  	<location>/common/500.jsp</location>
  </error-page>
  <error-page>
  	<error-code>404</error-code>
  	<location>/common/404.html</location>
  </error-page>
</web-app>
```



## EL表达式

#### EL作用

jsp的核心语法： jsp表达式 `<%=%>`和 jsp脚本`<%  %>`。
以后开发jsp的原则： 尽量在jsp页面中少写甚至不写java代码。

使用`EL表达式替换掉 jsp 表达式`

EL表达式作用： 向浏览器输出域对象中的变量值或表达式计算的结果！！！

语法： `${变量或表达式}`

```jsp
<%@ page language="java" import="java.util.*" pageEncoding="utf-8"%>
<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN">

<html>
  <head> 
    <title>EL语法</title>  
  </head>
  <body>
  
    <%
	 String name = "rose";  
	 //放入域中
	 //pageContext.setAttribute("name",name);
	 pageContext.setAttribute("name",name,PageContext.REQUEST_SCOPE); 
	  %>
	  <%=name %>
	  <br/>
      
	  <%--
	  1)从四个域自动搜索
	   --%>
      
	  EL表达式： ${name }
      
	  <%--
	  	${name } 等价于
	  	 <%=pageContext.findAttribute("name")%>
	   --%>
      
	   <%--
	    2） 从指定的域中获取数据
	    --%>
      
	    EL表达式： ${pageScope.name }
      
	    <%--
	    	${pageScope.name } 等价于
	    	 <%= pageContext.getAttribute("name",PageContext.PAGE_SCOPE)%>
	     --%>
      
  </body>
</html>
```



#### EL语法

1. 输出基本数据类型变量
   1.1 从四个域获取
     			${name}
     	1.2 指定域获取
     			${pageScope.name}
              域范围： pageScoep / requestScope / sessionScope / applicationScope
2. 输出对象的属性值
   Student

3. 输出集合对象
   	   List  和 Map
4. EL表达式计算
5. 实例

```jsp
<%@ page language="java" import="java.util.*,gz.ixfosa.entity.*" pageEncoding="utf-8"%>

<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN">
<html>
  <head> 
    <title>EL输出不同类型的数据</title>  
  </head>
  
  <body>
    <%--
		1)EL输出对象的属性    
     --%>
     <%
     	//保存数据
     	Student student = new Student("eric",20);
     	//放入域中
     	pageContext.setAttribute("student",student);
     	
     	//List
     	List<Student>  list = new ArrayList<Student>();
     	list.add(new Student("rose",18));
     	list.add(new Student("jack",28));
     	list.add(new Student("lucy",38));
     	//放入域中
     	pageContext.setAttribute("list",list);
     	
     	//Map
     	Map<String,Student> map = new HashMap<String,Student>();
     	map.put("100",new Student("mark",20));
     	map.put("101",new Student("maxwell",30));
     	map.put("102",new Student("narci",40));
     	//放入域中
     	pageContext.setAttribute("map",map);
     	
      %>
      
      <%--使用EL获取对象 --%>
      ${student.name} - ${student.age }
      <%--
       ${student.name} 等价于     （点相对于调用getXX（）方法）
          <%=((Student)pageContext.findAttribute("student")).getName()%>
       --%>
       
       <hr/>
       <%--使用EL获取List对象 --%>
       ${list[0].name } - ${list[0].age }<br/>
       ${list[1].name } - ${list[1].age }<br/>
       ${list[2].name } - ${list[2].age }
       <%--
       list[0]等价于       (中括号相对于调用get(参数)方法)
       		((List)pageContext.findAttribute("list")).get(0)
        --%>
        <hr/>
        <%--使用EL获取Map对象 --%>
        ${map['100'].name } -  ${map['100'].age }<br/>
        ${map['101'].name } -  ${map['101'].age }<br/>
        ${map['102'].name } -  ${map['102'].age }<br/>
  </body>
</html>
```

```java
<%@ page language="java" import="java.util.*" pageEncoding="utf-8"%>

<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN">
<html>
  <head> 
    <title>EL表达式计算</title>  
  </head>
  
  <body>
    <%--
    	1)算术表达式
    	  +  -  *  /
     --%>
     ${10+5 }<br/>
     ${10*5 }
     <hr/>
     <%--
    	2)比较运算
    	>  <  >=  <=  ==   !=
     --%>
     ${10>5 }<br/>
     ${10<5 }<br/>
     ${10!=10 }
     <hr/>
     <%--
    	3)逻辑运算
    	 &&  ||  !
     --%>
     ${true && false }<br/>
     ${true || false }<br/>
     ${!false }<br/>
     <hr/>
      <%--
    	4)判空
    	null 或 空字符串:  empty
     --%>
     <%
     	//String name = "eric";
     	//String name = null;
     	String name = "";
     	pageContext.setAttribute("name",name);
      %>
      判断null： ${name==null }<br/>
      判断空字符： ${name=="" }<br/>
     判空：  ${name==null || name=="" }
     另一种判空写法： ${empty name }
  
  </body>
</html>
```



##  jsp标签

jsp标签作用：替换jsp脚本



#### jsp标签分类

​			1）内置标签（动作标签）： 不需要在jsp页面导入标签
​			2）jstl标签： 需要在jsp页面中导入标签
​			3）自定义标签 ： 开发者自行定义，需要在jsp页面导入标签



#### 动作标签

 			转发标签：`<jsp:forward />`
            参数标签： ` <jsp:pararm/>`
			包含标签： ` <jsp:include/`>
			原理： 包含与被包含的页面先各自翻译成java源文件，然后再运行时合并在一起。
						（先翻译再合并），动态包含

​	静态包含  vs  动态包含的区别？

			1.  语法不同
			静态包含语法： <%@inclue file="被包含的页面"%>
			动态包含语法：　<jsp:include page="被包含的页面">
	
			2. 参数传递不同
			静态包含不能向被包含页面传递参数
			动态包含可以向被包含页面传递参数
	
	3. 原理不同
			静态包含： 先合并再翻译
			动态包含： 先翻译再合并

```jsp
<%@ page language="java" import="java.util.*" pageEncoding="utf-8"%>
<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN">

<html>
  <head> 
    <title>动作标签</title>  
  </head>
  
  <body>
    <%--转发 --%>
    <%
    	//request.getRequestDispatcher("/09.action2.jsp?name=eric").forward(request,response);
     %>
      
    <%-- 参数 --%>
    <%--
    <jsp:forward page="/09.action2.jsp">
    	<jsp:param value="jacky" name="name"/>
    	<jsp:param value="123456" name="password"/>
    </jsp:forward>
        
      --%>
      
      <%--包含 --%>
      <%--
   <jsp:include page="/common/header.jsp">
   		<jsp:param value="lucy" name="name"/>
   	</jsp:include>
   	 --%>
   	 <%@include file="common/header.jsp" %>
      主页的内容
  </body>
</html>


/09.action2.jsp
-----------------------------------------------------------------

<%@ page language="java" import="java.util.*" pageEncoding="utf-8"%>
<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN">
<html>
  <head> 
    <title>动作标签</title>  
  </head>
  
  <body>
    <%=request.getParameter("name")%><br/>
    <%=request.getParameter("password")%>
  </body>
</html>
```

​				

#### JSTL标签

​		JSTL (全名：java  standard  tag  libarary   -  java标准标签库  )

​		核心标签库 （c标签库） 天天用
​		国际化标签（fmt标签库）
​		EL函数库（fn函数库）
​		xml标签库（x标签库）
​		sql标签库（sql标签库）			

##### 使用JSTL标签步骤

1. 导入jstl支持的jar包（标签背后隐藏的java代码）

2. 使用taglib指令导入标签库 

   `<%@taglib uri="tld文件的uri名称" prefix="简写" %>`

3. 在jsp中使用标签		

```jsp
<%--导入标签库 --%>
<%@taglib uri="http://java.sun.com/jsp/jstl/core" prefix="c" %>
```

##### 核心标签库的重点标签：		

​	保存数据：
​			<c:set></c:set>  
​	获取数据： 
​             <c:out value=""></c:out>
​	单条件判断
​            <c:if test=""></c:if>
​	多条件判断
​          <c:choose></c:choose>
​    	  <c:when test=""></c:when>
​          <c:otherwise></c:otherwise>
​    循环数据
​          <c:forEach></c:forEach>
​          <c:forTokens items="" delims=""></c:forTokens>
​	重定向
​          <c:redirect></c:redirect>

```jsp
<%@ page language="java" import="java.util.*,gz.ixfosa.entity.*" pageEncoding="utf-8"%>
<%--导入标签库 --%>
<%@taglib uri="http://java.sun.com/jsp/jstl/core" prefix="c" %>
<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN">
<html>
  <head> 
    <title>核心标签库</title>  
  </head>
  
  <body>
    <%--使用标签 --%>
    <%--set标签 :保存数据(保存到域中)默认保存到page域 --%>
    <c:set var="name" value="rose" scope="request"></c:set>
    
    <%
    	String msg = null;
    	pageContext.setAttribute("msg",msg);
     %>
    ${msg }
      
    <br/>
      
    <%--out标签： 获取数据（从域中） 
    default： 当value值为null时，使用默认值
    escapeXml: 是否对value值进行转义，false，不转义，true，转义（默认）
    --%>
    <c:out value="${msg}" default="<h3>标题3</h3>" escapeXml="true"></c:out>
    
    <hr/>
    
    <%--if标签 ：单条件判断--%>
    <c:if test="${!empty msg}">
    	条件成立
    </c:if>
    
    <hr/>
    <%--choose标签+when标签+otherwirse标签: 多条件判断 --%>
    <c:set var="score" value="56"></c:set>
    
    <c:choose>
    	<c:when test="${score>=90 && score<=100}">
    		优秀
    	</c:when>
    	<c:when test="${score>=80 && score<90}">
    		良好
    	</c:when>
    	<c:when test="${score>=70 && score<80}">
    		一般
    	</c:when>
    	<c:when test="${score>=60 && score<70}">
    		及格
    	</c:when>
    	<c:otherwise>
    		不及格
    	</c:otherwise>
    </c:choose>
    
    <%-- forEach标签：循环 --%>
    <%
    	//List
     	List<Student>  list = new ArrayList<Student>();
     	list.add(new Student("rose",18));
     	list.add(new Student("jack",28));
     	list.add(new Student("lucy",38));
     	//放入域中
     	pageContext.setAttribute("list",list);
     	
     	//Map
     	Map<String,Student> map = new HashMap<String,Student>();
     	map.put("100",new Student("mark",20));
     	map.put("101",new Student("maxwell",30));
     	map.put("102",new Student("narci",40));
     	//放入域中
     	pageContext.setAttribute("map",map);
     %>
     <hr/>
     <%--
      begin="" : 从哪个元素开始遍历，从0开始.默认从0开始
      end="":     到哪个元素结束。默认到最后一个元素
      step="" ： 步长    (每次加几)  ，默认1
      items=""： 需要遍历的数据（集合） 
      var=""： 每个元素的名称 
      varStatus=""： 当前正在遍历元素的状态对象。（count属性：当前位置，从1开始）
      
     --%>
    <c:forEach items="${list}" var="student" varStatus="varSta">
    	序号：${varSta.count} - 姓名：${student.name } - 年龄：${student.age}<br/>
    </c:forEach>
    
    <hr/>
    
    <c:forEach items="${map}" var="entry">
    	${entry.key } - 姓名： ${entry.value.name } - 年龄：${entry.value.age }<br/>
    </c:forEach>
    <hr/>
    <%-- forToken标签： 循环特殊字符串 --%>
    <%
    	String str = "java-php-net-平面";
    	pageContext.setAttribute("str",str);
     %>
    
    <c:forTokens items="${str}" delims="-" var="s">
    	${s }<br/>
    </c:forTokens>
    
    <%--redrict:重定向 --%>
    <c:redirect url="http://www.baidu.com"></c:redirect>
  </body>
</html>
```



#### 自定义标签

##### 自定义标签开发步骤

需求： 向浏览器输出当前客户的IP地址 （只能使用jsp标签）

第一个自定义标签开发步骤

1. 编写一个普通的java类，继承SimpleTagSupport类，叫标签处理器类

```java
package gz.ixfosa.tag;

import java.io.IOException;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.jsp.JspContext;
import javax.servlet.jsp.JspException;
import javax.servlet.jsp.JspWriter;
import javax.servlet.jsp.PageContext;
import javax.servlet.jsp.tagext.SimpleTagSupport;

/**
 * 标签处理器类
 * 1）继承SimpleTagSupport
 *
 */
public class ShowIpTag extends SimpleTagSupport{
	/**
	 * 以下屏蔽的代码在SimpleTagSupport代码中已经做了！这里不需要重复再做！
	 */
	/*private JspContext context;
	
	*//**
	 * 传入pageContext
	 *//*
	@Override
	public void setJspContext(JspContext pc) {
		this.context = pc;
	}*/


	/**
	 * 2）覆盖doTag方法
	 */
	@Override
	public void doTag() throws JspException, IOException {
		//向浏览器输出客户的ip地址
		PageContext pageContext = (PageContext)this.getJspContext();
		
		HttpServletRequest request = (HttpServletRequest)pageContext.getRequest();
		
		String ip = request.getRemoteHost();
		
		JspWriter out = pageContext.getOut();
		
		out.write("使用自定义标签输出客户的IP地址："+ip);
		
	}
}
```



2. 在web项目的WEB-INF目录下建立ixfosa.tld文件，这个tld叫标签库的声明文件。

   （参考核心标签库的tld文件)

```xml
<?xml version="1.0" encoding="UTF-8" ?>

<taglib xmlns="http://java.sun.com/xml/ns/javaee"
    xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
    xsi:schemaLocation="http://java.sun.com/xml/ns/javaee http://java.sun.com/xml/ns/javaee/web-jsptaglibrary_2_1.xsd"
    version="2.1">
    
  <!-- 标签库的版本 -->
  <tlib-version>1.1</tlib-version>
  <!-- 标签库前缀 -->
  <short-name>ixfosa</short-name>
  <!-- tld文件的唯一标记 -->
  <uri>http://gz.ixfosa.cn</uri>

  <!-- 一个标签的声明 -->
  <tag>
    <!-- 标签名称 -->
    <name>showIp</name>
    <!-- 标签处理器类的全名 -->
    <tag-class>gz.itcast.a_tag.ShowIpTag</tag-class>
    <!-- 输出标签体内容格式 -->
    <body-content>scriptless</body-content>
  </tag>

</taglib>
```

3. 在jsp页面的头部导入自定义标签库
   <%@taglib uri="http://gz.ixfosa.cn" prefix="ixfosa"%>
4. 在jsp中使用自定义标签
     	<itcast:showIp></ifosa:showIp>

01.hellotag.jsp 

```jsp
<%@ page language="java" import="java.util.*" pageEncoding="utf-8"%>
<%@taglib uri="http://gz.ixfosa.cn" prefix="ixfosa"%>
<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN">
<html>
  <head> 
    <title>第一个自定义标签</title>  
  </head>
  
  <body>
    <%
    	//获取当前客户的IP地址
    	//String ip = request.getRemoteHost();
    	//输出到浏览器
    	//out.write("当前客户的IP地址是："+ip);
     %>
     <%--使用标签库中的标签--%>
     <ixfosa:showIp>xxxxx</ixfosa:showIp>
     
  </body>
</html>
```



#####  自定义标签的执行过程	

问题： http://localhost:8080/day14/01.hellotag.jsp  如何访问到自定义标

前提： tomcat服务器启动时，加载到每个web应用，加载每个web应用的WEB-INF目录下的所有文件！！！例如。web.xml, tld文件！！！

1. 访问01.hellotag.jsp资源

2. tomcat服务器把jsp文件翻译成java源文件->编译class->构造类对象->调用_jspService（）方法

3. 检查jsp文件的taglib指令，是否存在一个名为http://gz.itcast.cn的tld文件。如果没有，则报错

4. 上一步已经读到itcast.tld文件

5. 读到`<ifosa:showIp> `到itcast.tld文件中查询是否存在<name>为showIp的<tag>标签

6. 找到对应的<tag>标签，则读到<tag-class>内容

7. 得到 gz.ixfosa.a_tag.ShowIpTag
       构造ShowIpTag对象，然后调用ShowIpTag里面的方法

   

#####  自定义标签处理器类生命周期

```java
SimpleTag接口： 
void setJspContext(JspContext pc)  --设置pageContext对象，传入pageContext（一定调用）
     通过getJspCotext()方法得到pageContext对象

void setParent(JspTag parent)   -- 设置父标签对象，传入父标签对象，如果没有父标签，则不										   调用此方法。通过getParent()方法得到父标签对象。
	void  setXXX(值)                     --设置属性值。
	void setJspBody(JspFragment jspBody) --设置标签体内容。标签体内容封装到JspFragment对象											中，然后传入JspFragment对象。通过getJspBody()											  方法得到标签体内容。如果没有标签体内容，则不会调用                                             此方法
    void doTag()                         --执行标签时调用的方法。（一定调用）
```

##### 自定义标签的作用

     		1. 控制标签体内容是否输出
    		2. 控制标签余下内容是否输出
         3. 控制重复输出标签体内容
            4. 改变标签体内容
            5. 带属性的标签
               步骤： 在标签处理器中添加一个成语变量和setter方法

```java
package gz.ixfosa.tag;

import java.io.IOException;
import java.io.StringWriter;

import javax.servlet.jsp.JspException;
import javax.servlet.jsp.SkipPageException;
import javax.servlet.jsp.tagext.JspFragment;
import javax.servlet.jsp.tagext.SimpleTagSupport;
/**
 * 标签处理器类
 * @author APPle
 *
 */
public class DemoTag extends SimpleTagSupport{
	//1.声明属性的成员变量
	private Integer num;
	
	//2.关键点： 必须提供公开的setter方法，用于给属性赋值
	public void setNum(Integer num) {
		this.num = num;
	}

	@Override
	public void doTag() throws JspException, IOException {
		System.out.println("执行了标签");
		
		/**
		 * 1）控制标签内容是否输出
		 *    输出： 调用jspFrament.invoke();
		 *    不输出： 不调用jspFrament.invoke();
		 */
		//1.1 得到标签体内容
		JspFragment jspBody = this.getJspBody();
		
		/**
		 * 执行invoke方法： 把标签体内容输出到指定的Writer对象中
		 */
		//1.2 往浏览器输出内容，writer为null就是默认往浏览器输出
		//JspWriter out = this.getJspContext().getOut();
		//jspBody.invoke(out);
		jspBody.invoke(null);//等价于上面的代码
		
		/**
		 * 3)控制重复输出标签体内容
		 *     方法： 执行多次jspBody.invoke()方法
		 */
		/*for(int i=1;i<=num;i++){
			jspBody.invoke(null);
		}*/
		
		/**
		 * 4)改变标签体内容
		 */
		//4.1 创建StringWriter临时容器
		/*StringWriter sw = new StringWriter();
		//4.2 把标签体拷贝到临时容器
		jspBody.invoke(sw);
		//4.3 从临时容器中得到标签体内容
		String content = sw.toString();
		//4.4 改变内容
		content = content.toLowerCase();
		//System.out.println(content);
		//4.5 把改变的内容输出到浏览器
		//jspBody.invoke(null); 不能使用此方式输出，因为jsbBody没有改变过
		this.getJspContext().getOut().write(content);*/
		
		
		/**
		 * 2)控制标签余下内容是否输出
		 *   输出: 什么都不干！
		 *   不输出： 抛出SkipPageException异常
		 */
		throw new SkipPageException();
		
	}
}
```

```xml
<?xml version="1.0" encoding="UTF-8" ?>

<taglib xmlns="http://java.sun.com/xml/ns/javaee"
    xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
    xsi:schemaLocation="http://java.sun.com/xml/ns/javaee http://java.sun.com/xml/ns/javaee/web-jsptaglibrary_2_1.xsd"
    version="2.1">
  <!-- 标签库的版本 -->
  <tlib-version>1.1</tlib-version>
  <!-- 标签库前缀 -->
  <short-name>ixfosa</short-name>
  <!-- tld文件的唯一标记 -->
  <uri>http://gz.ixfosa.cn</uri>

  <tag>
    <name>demoTag</name>
    <tag-class>gz.ixfosa.tag.DemoTag</tag-class>
    <body-content>scriptless</body-content>
    <!-- 属性声明 -->
    <attribute>
    	<!-- 属性名称 -->
    	<name>num</name>
    	<!-- 是否必填 -->
    	<required>true</required>
    	<!-- 是否支持EL表达式 -->
    	<rtexprvalue>false</rtexprvalue>
    </attribute>
  </tag>
 
</taglib>

```



##### 输出标签体内容格式

​		JSP：   在传统标签中使用的。可以写和执行jsp的java代码。
​		scriptless:  标签体不可以写jsp的java代码
​		empty:    必须是空标签。
​		tagdependent : 标签体内容可以写jsp的java代码，但不会执行。

##### 案例

​		核心标签库： c:if   c:choose+c:when+c:otherwise   c:forEach
​		高仿核心标签库

###### 登录标签

```xml
<%@ page language="java" import="java.util.*" pageEncoding="utf-8"%>
<%@taglib uri="http://gz.ixfosa.cn" prefix="ixfosa" %>
<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN">
<html>
  <head> 
    <title>登陆页面</title>  
  </head>
  
  <body>
  <form action="" method="post">
   <ixfosa:login password="pwd" username="user"></ixfosa:login> 
   </form>
  </body>
</html>
```

```java
package gz.ixfosa.cases;

import java.io.IOException;

import javax.servlet.http.HttpServletResponse;
import javax.servlet.jsp.JspException;
import javax.servlet.jsp.JspWriter;
import javax.servlet.jsp.PageContext;
import javax.servlet.jsp.tagext.SimpleTagSupport;

/**
 * 自定义登陆页面标签
 *
 */
public class LoginTag extends SimpleTagSupport{
	private String username;
	private String password;

	public void setUsername(String username) {
		this.username = username;
	}

	public void setPassword(String password) {
		this.password = password;
	}

	@Override
	public void doTag() throws JspException, IOException {
		 HttpServletResponse response = (HttpServletResponse)((PageContext)this.getJspContext()).getResponse();
		 //设置输出内容类型和编码
		 response.setContentType("text/html;charset=utf-8");
		 String html = "";
		
		 html += "<center><h3>用户登陆页面</h3></center>";
		 html += "<table border='1' align='center' width='400px'>";
		 html += "	<tr>";
		 html += "		<th>用户名：</th>";
		 html += "		<td><input type='text' name='"+username+"'/></td>";
		 html += "	</tr>";
		 html += "	<tr>";
		 html += "		<th>密码：</th>";
		 html += "		<td><input type='password' name='"+password+"'/></td>";
		 html += "	</tr>";
		 html += "	<tr>";
		 html += "		<td colspan='2' align='center'><input type='submit' value='登陆'/>&nbsp;<input type='reset' value='重置'/></td>";
		 html += "	</tr>";
		 html += "</table>";
		
		JspWriter out = this.getJspContext().getOut();
		out.write(html);
	}
}
```

###### c:if 

```jsp
<%@ page language="java" import="java.util.*" pageEncoding="utf-8"%>
<%@taglib uri="http://gz.ixfosa.cn" prefix="ixfosa" %>
<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN">
<html>
  <head> 
    <title>if标签</title>  
  </head>
  
  <body>
    <ixfosa:if test="${10>5}">
    	条件成立
    </ixfosa:if>
  </body>
</html>
```

```java
package gz.ixfosa.cases;

import java.io.IOException;

import javax.servlet.jsp.JspException;
import javax.servlet.jsp.tagext.SimpleTagSupport;

public class IfTag extends SimpleTagSupport {
	private boolean test;
	
	public void setTest(boolean test) {
		this.test = test;
	}

	@Override
	public void doTag() throws JspException, IOException {
		//根据test的返回值决定是否输出标签体内容
		if(test){
			this.getJspBody().invoke(null);
		}
	}
}
```



###### c:choose+c:when+c:otherwise

```jsp
<%@ page language="java" import="java.util.*" pageEncoding="utf-8"%>
<%@taglib uri="http://gz.ixfosa.cn" prefix="ixfosa" %>
<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN">
<html>
  <head> 
    <title>choose标签</title>  
  </head>
  
  <body>
    <ixfosa:choose>
		<ixfosa:when test="${10<5}">
			条件成立
		</ixfosa:when>
		<ixfosa:otherwise>
			条件不成立
		</ixfosa:otherwise>    
    </ixfosa:choose>
  </body>
</html>
```

```java
package gz.ixfosa.cases;

import java.io.IOException;

import javax.servlet.jsp.JspException;
import javax.servlet.jsp.tagext.SimpleTagSupport;

public class ChooseTag extends SimpleTagSupport {
	//不是属性，而是临时变量
	private boolean flag;

	public boolean isFlag() {
		return flag;
	}

	public void setFlag(boolean flag) {
		this.flag = flag;
	}

	@Override
	public void doTag() throws JspException, IOException {
		//输出标签体内容
		this.getJspBody().invoke(null);
	}
}
```

```java
package gz.ixfosa.cases;

import java.io.IOException;

import javax.servlet.jsp.JspException;
import javax.servlet.jsp.tagext.SimpleTagSupport;

public class WhenTag extends SimpleTagSupport {
	private boolean test;
	
	public void setTest(boolean test) {
		this.test = test;
	}

	@Override
	public void doTag() throws JspException, IOException {
		//根据test的返回值决定是否输出标签体内容
		if(test){
			this.getJspBody().invoke(null);
		}
		
		//获取父标签
		ChooseTag parent = (ChooseTag)this.getParent();
		parent.setFlag(test);
		
	}
}

```



```java
package gz.ixfosa.cases;

import java.io.IOException;

import javax.servlet.jsp.JspException;
import javax.servlet.jsp.tagext.SimpleTagSupport;

public class ChooseTag extends SimpleTagSupport {
	//不是属性，而是临时变量
	private boolean flag;

	public boolean isFlag() {
		return flag;
	}

	public void setFlag(boolean flag) {
		this.flag = flag;
	}

	@Override
	public void doTag() throws JspException, IOException {
		//输出标签体内容
		this.getJspBody().invoke(null);
	}
}
```



------



```jsp
<%@ page language="java" import="java.util.*" pageEncoding="utf-8"%>
<%@taglib uri="http://gz.ixfosa.cn" prefix="ixfosa" %>
<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN">
<html>
  <head> 
    <title>choose标签</title>  
  </head>
  
  <body>
    <ixfosa:choose>
		<ixfosa:when test="${10<5}">
			条件成立
		</ixfosa:when>
		<ixfosa:otherwise>
			条件不成立
		</ixfosa:otherwise>    
    </ixfosa:choose>
  </body>
</html>
```

```java
package gz.ixfosa.cases;

import java.io.IOException;

import javax.servlet.jsp.JspException;
import javax.servlet.jsp.tagext.SimpleTagSupport;

public class OtherwiseTag extends SimpleTagSupport {
	@Override
	public void doTag() throws JspException, IOException {
		//通过父标签传递，when标签中test的值
		//获取父标签
		ChooseTag parent = (ChooseTag)this.getParent();
		boolean test = parent.isFlag();
		
		if(!test){
			this.getJspBody().invoke(null);
		}
	}
}
```



###### c:forEach

```jsp
<%@ page language="java" import="java.util.*,gz.ixfosa.cases.*" pageEncoding="utf-8"%>
<%@taglib uri="http://gz.ixfosa.cn" prefix="ixfosa" %>
<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN">
<html>
  <head> 
    <title>forEach标签</title>  
  </head>
  
  <body>
    <%
       //保存数据
       //List
     	List<Student>  list = new ArrayList<Student>();
     	list.add(new Student("rose",18));
     	list.add(new Student("jack",28));
     	list.add(new Student("lucy",38));
     	//放入域中
     	pageContext.setAttribute("list",list);
     	
     	//Map
     	Map<String,Student> map = new HashMap<String,Student>();
     	map.put("100",new Student("mark",20));
     	map.put("101",new Student("maxwell",30));
     	map.put("102",new Student("narci",40));
     	//放入域中
     	pageContext.setAttribute("map",map);
     %>
     
     <ixfosa:forEach items="${list}" var="student">
     		姓名：${student.name } - 年龄:${student.age }<br/>
     </ixfosa:forEach>
     
     <hr/>
     
     <ixfosa:forEach items="${map}" var="entry">
     	  编号：${entry.key} - 姓名：${entry.value.name} - 年龄：${entry.value.age }<br/>
     </ixfosa:forEach>
  </body>
</html>
```

```java
package gz.ixfosa.cases;

import java.io.IOException;
import java.util.Collection;
import java.util.List;
import java.util.Map;
import java.util.Set;
import java.util.Map.Entry;

import javax.servlet.jsp.JspException;
import javax.servlet.jsp.PageContext;
import javax.servlet.jsp.tagext.SimpleTagSupport;

public class ForEachTag extends SimpleTagSupport {
	private Object items;//需要遍历的数据.List和map
	private String var;//每个元素的名称
	
	public void setItems(Object items) {
		this.items = items;
	}

	public void setVar(String var) {
		this.var = var;
	}

	@Override
	public void doTag() throws JspException, IOException {
        
		//遍历items数据
		//List
		/*PageContext pageContext = (PageContext)this.getJspContext();
		if(items instanceof List){
			List list = (List)items;
			for (Object object : list) {
				//把每个对象放入域对象中(pageContext)
				pageContext.setAttribute(var, object);
				//显示标签体内容
				this.getJspBody().invoke(null);
			}
		}
		
		//Map
		if(items instanceof Map){
			Map map = (Map)items;
			Set<Entry> entrySet = map.entrySet();
			for(Entry entry :entrySet){
				//把每个对象放入域对象中(pageContext)
				pageContext.setAttribute(var, entry);
				//显示标签体内容
				this.getJspBody().invoke(null);
			}
		}*/
		
		
		//简化代码
		//思路： 
			//1）list         -> Collection
			//2) map.entrySet -> Collection
		PageContext pageContext = (PageContext)this.getJspContext();
		Collection colls = null;
		if(items instanceof List){
			colls = (List)items;
		}
		if(items instanceof Map){
			Map map = (Map)items;
			colls = map.entrySet();
		}
		     
		for(Object object:colls){
			//把每个对象放入域对象中(pageContext)
			pageContext.setAttribute(var, object);
			//显示标签体内容
			this.getJspBody().invoke(null);
		}
	}
}
```

###### ixfosa.tld

```xml
<?xml version="1.0" encoding="UTF-8" ?>

<taglib xmlns="http://java.sun.com/xml/ns/javaee"
    xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
    xsi:schemaLocation="http://java.sun.com/xml/ns/javaee http://java.sun.com/xml/ns/javaee/web-jsptaglibrary_2_1.xsd"
    version="2.1">
    
  <!-- 标签库的版本 -->
  <tlib-version>1.1</tlib-version>
  <!-- 标签库前缀 -->
  <short-name>ixfosa</short-name>
  <!-- tld文件的唯一标记 -->
  <uri>http://gz.ixfosa.cn</uri>
  
  <tag>
    <name>login</name>
    <tag-class>gz.ixfosa.cases.LoginTag</tag-class>
    <body-content>scriptless</body-content>
    <attribute>
    	<name>username</name>
    	<required>true</required>
    	<rtexprvalue>false</rtexprvalue>
    </attribute>
    <attribute>
    	<name>password</name>
    	<required>true</required>
    	<rtexprvalue>false</rtexprvalue>
    </attribute>
  </tag>
  
  <tag>
    <name>if</name>
    <tag-class>gz.ixfosa.cases.IfTag</tag-class>
    <body-content>scriptless</body-content>
    <attribute>
    	<name>test</name>
    	<required>true</required>
    	<rtexprvalue>true</rtexprvalue>
    </attribute>
  </tag>
  
  <tag>
    <name>choose</name>
    <tag-class>gz.ixfosa.cases.ChooseTag</tag-class>
    <body-content>scriptless</body-content>
  </tag>
  
  <tag>
    <name>when</name>
    <tag-class>gz.ixfosa.cases.WhenTag</tag-class>
    <body-content>scriptless</body-content>
    <attribute>
    	<name>test</name>
    	<required>true</required>
    	<rtexprvalue>true</rtexprvalue>
    </attribute>
  </tag>
  
  <tag>
    <name>otherwise</name>
    <tag-class>gz.ixfosa.cases.OtherwiseTag</tag-class>
    <body-content>scriptless</body-content>
  </tag>
  
  <tag>
    <name>forEach</name>
    <tag-class>gz.ixfosa.cases.ForEachTag</tag-class>
    <body-content>scriptless</body-content>
    <attribute>
    	<name>items</name>
    	<required>true</required>
    	<rtexprvalue>true</rtexprvalue>
    </attribute>
    <attribute>
    	<name>var</name>
    	<required>true</required>
    	<rtexprvalue>false</rtexprvalue>
    </attribute>
  </tag>
    
</taglib>
```



## JavaBean

JavaBean,  咖啡豆。 JavaBean是一种开发规范，可以说是一种技术。

JavaBean就是一个普通的java类。只有符合以下规定才能称之为javabean：

	1. 必须提供无参数的构造方法
	2. 类中属性都必须私有化(private)
	3. 该类提供公开的getter 和 setter方法



JavaBean的作用： 用于封装数据，保存数据。
			访问javabean只能使用`getter`和`setter`方法



JavaBean的使用场景：

1. 项目中用到实体对象（entity）符合javabean规范

2. EL表达式访问对象属性。${student.name}  调用getName()方法，符合javabean规范。

3. jsp标签中的属性赋值。 setNum（Integer num）。符合javabean规范。

4. jsp页面中使用javabean。符合javabean规范

   

​	问题： 以下方法哪些属于javabean的规范的方法？ 答案 ：( 1，3，5，6  )
​			    注意： boolean类型的get方法名称叫 isXXX()方法

1. getName()    

2. getName(String name)
3. setName(String name) 
4. setName()
5. setFlag(boolean flag)   6)isFlag()

```jsp
<%@ page language="java" import="java.util.*,gz.ixfosa.cases.*" pageEncoding="utf-8"%>

<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN">
<html>
  <head> 
    <title>jsp页面中使用javabean</title>  
  </head>
  
  <body>
    <%
    	//Student stu = new Student();
    	//stu.setName("rose");
    	//stu.setAge(20);
    	
    	//stu.getName();
     %>
     <%---创建对象 --%>
     <jsp:useBean id="stu" class="gz.itcast.b_cases.Student"></jsp:useBean>
     
     
     <%--赋值 --%>
     <jsp:setProperty property="name" name="stu" value="jacky"/>
     
       <%--获取--%>
    <jsp:getProperty property="name" name="stu"/>
  </body>
</html>
```



## web开发模式


MVC开发模式：

```java
Model        - JavaBean实现。用于封装业务数据
View         - Jsp实现。用于显示数据
Controller   - servlet实现。用于控制model和view
```

三层结构：

```java
dao层： 和数据访问相关的操作
service层： 和业务逻辑相关的操作
web层： 和用户直接交互相关的操作（传接参数，跳转页面）
```

![web开发模式](E:/smile/java/images/web开发模式.jpg)




