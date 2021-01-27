## Tomcat

### 基本使用

下载并按照

  1. 到apache官网。www.apache.org     http://jakarta.apache.org(产品的主页)

  2. 安装版：window （exe、msi） linux（rmp）
     压缩版：window（rar，zip） linux（tar，tar.gz）学习时候使用

  3. 运行和关闭tomcat

     1. 启动软件
        				a）找到%tomcat%/bin/startup.bat ，双击这个文件
        				b）弹出窗口，显示信息（不要关闭次窗口）
        				c）打开浏览器，输出以下地址   `http://localhost:8080`
        				d）看到一只猫画面，证明软件启动成功！

          	2. 关闭软件

        	​		a）找到%tomcat%/bin/shutdown.bat，双击这个文件即可！
        c）打开浏览器，输出以下地址。看到“无法连接”（最好先清空浏览器缓存）

        4. tomcat软件使用的常见问题

           1. 闪退问题
              				原因：tomcat软件是java语言开发的。 tomcat软件启动时，会默认到系统的环境变量中查找一个名称叫JAVA_HOME的变量。这个变量的作用找到tomcat启动所需的jvm。

              解决办法； 到环境变量中设置JAVA_HOME的变量

​          JAVA_HOME= C:\Program Files\Java\jdk1.6.0_30  (注意别配置到bin目录下)

 2. 端口占用的错误
    原因： tomcat启动所需的端口被其他软件占用了！
     		解决办法： 
     				a）关闭其他软件程序，释放所需端口
     				b）修改tomcat软件所需端口
     			       找到并修改%tomcat%/conf/server.xml文件

    ```xml
    <Connector port="8080" protocol="HTTP/1.1" 
                   connectionTimeout="20000" 
                   redirectPort="8443" />
    ```

    3. CATALINA环境变量问题
       原因： tomcat软件启动后，除了查找JAVA_HOME后，还会再查找一个叫CATALINA_HOME变量，这个变量的作用是设置tomcat的根目录。
       			解决办法：建议不要设置CATALINA_HOME变量。检查如果有的话，清除掉！



### Tomcat的目录结构

​		|-bin: 存放tomcat的命令。
​				 catalina.bat 命令：
​						`startup`.bat  -> catalina.bat start	
​						`shutdown.bat` - > catalina.bat stop
​	    |- conf: 存放tomcat的配置信息。其中server.xml文件是核心的配置文件。
​		|-lib：支持tomcat软件运行的jar包。其中还有技术支持包，如servlet，jsp
​		|-logs：运行过程的日志信息
​		|-temp: 临时目录
​		|-`webapps`： 共享资源目录。web应用目录。（注意不能以单独的文件进行共享）
​		|-work： tomcat的运行目录。jsp运行时产生的临时文件就存放在这里



## Web开发入门

### 基本介绍 

之前的程序： java桌面程序，控制台控制，socket gui界面。javase规范
现在和以后的程序：java web程序。浏览器控制。javaee规范



**软件的结构**
		`C/S` (Client - Server  客户端-服务器端)
				典型应用：QQ软件 ，飞秋，红蜘蛛。
				特点：
					    1）必须下载特定的客户端程序。
						2）服务器端升级，客户端升级。

​		`B/S`（Broswer -Server 浏览器端- 服务器端）
​				典型应用： 腾讯官方（www.qq.com）  163新闻网站， 传智官网（俗称：网站）
​				特点：
​						1）不需要安装特定的客户端（只需要安装浏览器即可！！）
​						2）服务器端升级，浏览器不需要升级！！！！

​					javaweb的程序就是b/s软件结构！！！



**服务器**
		从物理上来说，服务器就是一台PC机器。8核，8G以上，T来计算，带宽100M

​		web服务器：PC机器安装一个具有web服务的软件，称之为web服务器
​				数据库服务器：PC机器安装一个具有数据管理件服务的软件，称之为数据库服务器。
​				邮件服务器：PC机器安装一个具有发送邮件服务的软件，称之为邮件服务器。

web服务软件
		web服务软件的作用：把本地的资源共享给外部访问。



常见的市面上web服务软件
		javase的规范,包含IO流，线程，集合，socket编程。
		WebLogic: BEA公司的产品。 收费的。支持JavaEE规范。
		WebSphere： IBM公司的产品。收费的。支持JavaEE规范
		JBoss: Redhat公司的产品。收费的。支持JavaEE规范
		`Tomcat`： 开源组织Apache的产品。免费的。支持部分的JavaEE规范。
							（支持`servlet`、`jsp`。`jdbc`，但ejb， rmi不支持）



### Web应用的目录结构

​		|- WebRoot :   web应用的根目录
​				|- `静态资源`（html+css+js+image+vedio）
​				|- `WEB-INF` ： 固定写法。
​					|-`classes`： （可选）固定写法。存放class字节码文件
​					|-`lib`： （可选）固定写法。存放`jar包`文件。
​					|-`web.xml`    
​	

注意：

     		1. WEB-INF目录里面的资源不能通过浏览器直接访问
     		2. 如果希望访问到WEB-INF里面的资源，就必须把资源配置到一个叫web.xml的文件中。

练习：
			1）在webapps下建立一个test目录
			2）创建两个文件
					2.1 index.html  里面随便写内容	，有超链接-连接到test.html	
					2.2 test.html   里面随便写
			3）通过浏览器访问到。 

​					 `http://localhost:8080/test/index.html`



### 手动开发动态资源

1. 静态资源和动态资源的区别
   静态资源： 当用户多次访问这个资源，资源的源代码永远不会改变的资源。
    		动态资源：当用户多次访问这个资源，资源的源代码可能会发送改变。

2. 动态资源的开发技术
   Servlet : 用java语言来编写动态资源的开发技术。

   ​		Servlet特点：
   ​				1）普通的java类，继承HttpServlet类，覆盖doGet方法
   ​				2）Servlet类只能交给tomcat服务器运行！！！！（开发者自己不能运行！！！）

   

Servlet手动编写步骤：
	     1）编写一个servlet程序，继承HttpServlet    jar包：  `servlet-api.jar`

```java
import java.io.IOException;
import java.util.Date;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
/**
 * 第一个servlet程序
 *
 */
public class HelloServlet extends HttpServlet{

	@Override
	protected void doGet(HttpServletRequest req, HttpServletResponse resp)
			throws ServletException, IOException {
		//解决中文乱码问题
		resp.setContentType("text/html;charset=utf-8");
		//向浏览器输出内容
		resp.getWriter().write("这是第一个servlet程序。当前时间为："+new Date());
	}
}
```

​		2）找到HelloServlet类的class字节码，然后把拷贝到tomcat的一个web应用中

​			`WEB-INF/classes`目录下。 还有   `WEB-INF/lib`servlet-api.jar`

​		3）在当前web应用下的`web.xml`文件配置Servlet。

```xml
  <!-- 配置一个servlet程序 -->
  <servlet>
    <!-- servlet的内部名称 ，可以自定义-->
    <servlet-name>HelloServlet</servlet-name>
    <!-- servlet类名： 包名+简单类名-->
    <servlet-class>gz.ixfosa.servlet.HelloServlet</servlet-class>
  </servlet>

  <servlet-mapping>
    <!-- servlet的内部名称，和上面的名称保持一致！！！-->
    <servlet-name>HelloServlet</servlet-name>
    <!-- servlet的访问名称： /名称 -->
    <url-pattern>/hello</url-pattern>
  </servlet-mapping>
```

​		4）启动tomcat服务器，运行访问

​				访问servlet:  http://localhost:8080/ hello



## Servlet 实例

开发一个Servlet步骤：
				1）编写java类，继承`HttpServlet`类
				2）重新doGet和doPost方法
				3）Servlet程序交给tomcat服务器运行！！
						3.1 servlet程序的class码拷贝到WEB-INF/classes目录
						3.2 在web.xml文件中进行配置

```java
package gz.ixfosa.servlet;

import java.io.IOException;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

public class FirstServlet extends HttpServlet{
	
	@Override
	protected void doGet(HttpServletRequest req, HttpServletResponse response)
			throws ServletException, IOException {
		//向浏览器输出内容
		response.getWriter().write("this is first servlet！");
	}
}
```

WebRoot\WEB-INF\web.xml

```xml
<!-- 配置一个servlet -->
  <!-- servlet的配置 -->
  <servlet>
  	<!-- servlet的内部名称，自定义。尽量有意义 -->
  	<servlet-name>FirstServlet</servlet-name>
  	<!-- servlet的类全名： 包名+简单类名 -->
  	<servlet-class>gz.ixfosa.servlet.FirstServlet</servlet-class>
  </servlet>
  
  
  <!-- servlet的映射配置 -->
  <servlet-mapping>
  	<!-- servlet的内部名称，一定要和上面的内部名称保持一致！！ -->
  	<servlet-name>FirstServlet</servlet-name>
  	<!-- servlet的映射路径（访问servlet的名称） -->
  	<url-pattern>/first</url-pattern>
  </servlet-mapping>
```

访问次URL：  `http://localhost:8080/test/first`  

前提： tomcat服务器启动时，首先加载`webapps`中的每个web应用的`web.xml`配置文件。	
			http://: http协议
			localhost： 到本地的hosts文件中查找是否存在该域名对应的IP地址  127.0.0.1		
			8080：    找到tomcat服务器
			/test 在 tomcat 的 webapps 目录下找 test 的目录
			/first    资源名称。
				1）在 test 的 web.xml 中查找是否有匹配的 url-pattern 的内容（/first）
				2）如果找到匹配的 `url-pattern`,则使用当前 `servlet-name` 的名称到web.xml文件中查询是否相同名称的servlet配置
				3）如果找到，则取出对应的servlet配置信息中的servlet-class内容：	
					  字符串： gz.ixfosa.servlet.FirstServlet

​				通过反射：
​						a）构造FirstServlet的对象
​						b）然后调用FirstServlet里面的方法



## Servlet的映射路径

```xml
<servlet-mapping>
  	<!-- servlet的内部名称，一定要和上面的内部名称保持一致！！ -->
  	<servlet-name>FirstServlet</servlet-name>
  	<!-- servlet的映射路径（访问servlet的名称） -->
  	<url-pattern>/first</url-pattern>
</servlet-mapping>
```

|          | url-pattern  | 浏览器输入                                 |
| -------- | ------------ | ------------------------------------------ |
| 精确匹配 | /first       | http://localhost:8080/test/first           |
|          | /ixfosa/demo | http://localhost:8080/test/ixfosa/demo     |
|          |              |                                            |
| 模糊匹配 | /*           | http://localhost:8080/test/任意路径        |
|          | /ixfosa/*    | http://localhost:8080/test/ixfosa/任意路径 |
|          | *.后缀名*    | http://localhost:8080/test/任意路径.后缀名 |
|          | *.do         | http://localhost:8080/test/任意路径.do     |

注意：
		1）url-pattern要么以 `/` 开头，要么以`*`开头。  例如， ixfosa是非法路径。
		2）不能同时使用两种模糊匹配，例如 /itcast/*.do是非法路径
		3）当有输入的URL有多个servlet同时被匹配的情况下：
				3.1 精确匹配优先。（长的最像优先被匹配）
				3.2 以后缀名结尾的模糊url-pattern优先级最低！！

```xml
<servlet>
    <servlet-name>Demo1</servlet-name>
    <servlet-class>gz.ixfosa.mapping.Demo1</servlet-class>
</servlet>
<servlet>
    <servlet-name>Demo2</servlet-name>
    <servlet-class>gz.ixfosa.mapping.Demo2</servlet-class>
</servlet>



<servlet-mapping>
  	<servlet-name>Demo1</servlet-name>
  	<url-pattern>/ixfosa/*</url-pattern>
</servlet-mapping>

<servlet-mapping>
    <servlet-name>Demo2</servlet-name>
    <url-pattern>/demo2</url-pattern>
</servlet-mapping>
```

servlet缺省路径

​		servlet的缺省路径（<url-pattern>/</url-pattern>）是在tomcat服务器内置的一个路径。该路径对应的是一个DefaultServlet（缺省Servlet）。这个缺省的Servlet的作用是用于解析web应用的静态资源文件。

问题： URL输入http://localhost:8080/test/index.html 如何读取文件？？？？

​		1）到当前 test 应用下的web.xml文件查找是否有匹配的url-pattern。
​		2）如果没有匹配的url-pattern，则交给tomcat的内置的DefaultServlet处理
​		3）DefaultServlet程序到test应用的根目录下查找是存在一个名称为index.html的静态文件。
​		4）如果找到该文件，则读取该文件内容，返回给浏览器。
​		5）如果找不到该文件，则返回404错误页面。

结论： `先找动态资源，再找静态资源`。

静态的web资源是需要通过`DefaultServlet`程序去读取的



## Sevlet的生命周期

​	`Servlet的生命周期`： servlet类对象什么时候创建，什么时候调用什么方法，什么时候销毁。



​	以前的对象： new Student（）； stu.study();   stu=null;

​	Servlet程序的生命周期由tomcat服务器控制的！！！！
​			

#### Servlet重要的四个生命周期方法

构造方法     ： 创建servlet对象的时候调用。默认情况下，第一次访问servlet的时候创建servlet对象						只调用1次。证明servlet对象在tomcat是`单实例`的。
init方法       ： 创建完servlet对象的时候调用。只调用1次。
service方法： 每次发出请求时调用。调用n次。
destroy方法： 销毁servlet对象的时候调用。停止服务器或者重新部署web应用时销毁servlet对象。
						只调用1次。

```java
package gz.ixfosa.life;
import java.io.IOException;

import javax.servlet.ServletConfig;
import javax.servlet.ServletException;
import javax.servlet.ServletRequest;
import javax.servlet.ServletResponse;
import javax.servlet.http.HttpServlet;

public class LifeDemo extends HttpServlet {

	/**
	 * 1.构造方法
	 */
	public LifeDemo(){
		System.out.println("1.servlet对象被创建了。");
	}

	/**
	 * 2.init方法
	 */
	@Override
	public void init(ServletConfig config) throws ServletException {
		System.out.println("2.init方法被调用");
	}
	
	/**
	 * 3.service方法
	 */
	@Override			
	public void service(ServletRequest req, ServletResponse res)
			throws ServletException, IOException {
		System.out.println("3.service方法被调用");
	}
	
	/**
	 * 4.destroy方法
	 */
	@Override
	public void destroy() {
		System.out.println("4.servlet对象销毁了");
	}
}
```

```xml
<servlet>
    <servlet-name>LifeDemo</servlet-name>
    <servlet-class>gz.ixfosa.life.LifeDemo</servlet-class>
</servlet>

<servlet-mapping>
<servlet-name>LifeDemo</servlet-name>
<url-pattern>/LifeDemo</url-pattern>
</servlet-mapping>
```



####  伪代码演示servlet的生命周期

​	Tomtcat内部代码运行：

      			1. 通过映射找到到servlet-class的内容，字符串：gz.ixfosa.life.LifeDemo
                			2. 通过反射构造FirstServlet对象
            2.1 得到字节码对象
                             		Class clazz = class.forName("gz.ixfosa.life.LifeDemo");
            2.2 调用无参数的构造方法来构造对象
                             		Object obj = clazz.newInstance();     ---1.servlet的构造方法被调用
                			3. 创建ServletConfig对象，通过反射调用init方法
    
       3.1 得到方法对象
       		Method m = clazz.getDeclareMethod("init",ServletConfig.class);
       3.2 调用方法
       		m.invoke(obj,config);              --2.servlet的init方法被调用
    		4. 创建request，response对象，通过反射调用service方法
     4.1 得到方法对象
     		Methodm m =clazz.getDeclareMethod("service",HttpServletRequest.class,HttpServletResponse.class);
     4.2 调用方法
     		m.invoke(obj,request,response);   --3.servlet的service方法被调用
    		5. 当tomcat服务器停止或web应用重新部署，通过反射调用destroy方法
     5.1 得到方法对象
     		Method m = clazz.getDeclareMethod("destroy",null);
     5.2 调用方法
     		m.invoke(obj,null);            --4.servlet的destroy方法被调用

#### servlet生命周期时序图

![servlet时序图](E:\smile\java\images\servlet时序图.png)

## Servlet的自动加载

​			默认情况下，第一次访问servlet的时候创建servlet对象。如果servlet的构造方法或init方法中执行了比较多的逻辑代码，那么导致用户第一次访问sevrlet的时候比较慢。

​		改变servlet创建对象的时机： 提前到加载web应用的时候！！！

​		在servlet的配置信息中，加上一个`<load-on-startup>`即可！！

```xml
<servlet>
    <servlet-name>LifeDemo</servlet-name>
    <servlet-class>gz.ixfosa.life.LifeDemo</servlet-class>
    <!-- 让servlet对象自动加载 -->
    <load-on-startup>1</load-on-startup>  注意：整数值越大，创建优先级越低！！
</servlet>
```

## 有参init()和无参init()

```java
package gz.ixfosa.init;


import javax.servlet.ServletConfig;
import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;

/**
 * 有参数的init和无参的init方法
 *
 */
public class InitDemo extends HttpServlet {
		
	/**
	 * 有参数的init方法
	 * 该方法是servlet的生命周期方法，一定会被tomcat服务器调用
	 */
    
	/**
	 * 注意：如果要编写初始代码，不需要覆盖有参数的init方法
	 */
    
	/*@Override
	public void init(ServletConfig config) throws ServletException {
		System.out.println("有参数的init方法");
	}*/
	
	/**
	 * 无参数的init方法
	 * 该方法是servlet的编写初始化代码的方法。是Sun公司设计出来专门给开发者进行覆盖，然后在里面编写servlet的初始逻辑代码的方法。
	 */
	@Override
	public void init() throws ServletException {
		System.out.println("无参数的init方法");
	}
}
```

```xml
<servlet>
    <servlet-name>InitDemo</servlet-name>
    <servlet-class>gz.ixfosa.init.InitDemo</servlet-class>
    <load-on-startup>3</load-on-startup>
</servlet>

<servlet-mapping>
    <servlet-name>InitDemo</servlet-name>
    <url-pattern>/InitDemo</url-pattern>
</servlet-mapping>
```

## Servlet的多线程并发

`servlet`对象在tomcat服务器是`单实例多线程`的。

​		因为servlet是多线程的，所以当多个servlet的线程同时访问了servlet的共享数据，如成员变量，可能会引发`线程安全`问题。

​	解决办法：
​			1）把使用到共享数据的代码块进行同步（使用`synchronized`关键字进行同步）
​			2）建议在servlet类中尽量不要使用成员变量。如果确实要使用成员，必须同步。而且尽量缩小同步代码块的范围。以避免因为同步而导致并发效率降低。

```java
package gz.ixfosa.thread;

import java.io.IOException;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
/**
 * servlet的多线程并发问题
 *
 */
public class TheradDemo extends HttpServlet {
	
	int count = 1;

	public void doGet(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, IOException {
		response.setContentType("text/html;charset=utf-8");
		
		synchronized (TheradDemo.class) {//锁对象必须唯一。建议使用类对象
			response.getWriter().write("你现在是当前网站的第"+count+"个访客");   //线程1执行完  ， 线程2执行

		//线程1还没有执行count++
		/*try {
			Thread.sleep(5000);
		} catch (InterruptedException e) {
			e.printStackTrace();
		}*/
			count++;
		}
	}
}
```

## ServletConfig对象

作用:
		`ServletConfig`对象: 主要是用于加载`servlet的初始化参数`。在一个web应用可以存在多个ServletConfig对象（一个Servlet对应一个ServletConfig对象）

对象创建和得到
			创建时机： 在创建完servlet对象之后，在调用init方法之前创建。
			得到对象： 直接从有参数的init方法中得到！



 注意： servlet的参数只能由当前的这个sevlet获取！！！



**ServletConfig的API：**

```java
java.lang.String getInitParameter(java.lang.String name)  根据参数名获取参数值
java.util.Enumeration getInitParameterNames()             获取所有参数
ServletContext getServletContext()                        得到servlet上下文对象
java.lang.String getServletName()                         得到servlet的名称
```

**web.xml**

```xml
servlet的初始化参数配置
<servlet>
    <servlet-name>ConfigDemo</servlet-name>
    <servlet-class>gz.ixfosa.config.ConfigDemo</servlet-class>
    <!-- 初始参数： 这些参数会在加载web应用的时候，封装到ServletConfig对象中 -->
    <init-param>
        <param-name>path</param-name>
        <param-value>e:/b.txt</param-value>
    </init-param>
    <init-param>
        <param-name>BBB</param-name>
        <param-value>BBB's value</param-value>
    </init-param>
    <init-param>
        <param-name>CCCC</param-name>
        <param-value>CCCC's value</param-value>
    </init-param>
</servlet>

<servlet-mapping>
    <servlet-name>ConfigDemo</servlet-name>
    <url-pattern>/ConfigDemo</url-pattern>
</servlet-mapping>
```

```java
package gz.ixfosa.config;

import java.io.BufferedReader;
import java.io.File;
import java.io.FileReader;
import java.io.IOException;
import java.util.Enumeration;

import javax.servlet.ServletConfig;
import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

public class ConfigDemo extends HttpServlet {
	/**
	 * 以下两段代码GenericServlet已经写了，我们无需编写！！
	 */
	//private ServletConfig config;
	
	/**
	 *  1）tomcat服务器把这些参数会在加载web应用的时候，封装到ServletConfig对象中 
	 *  2）tomcat服务器调用init方法传入ServletConfig对象
	 */
	/*@Override
	public void init(ServletConfig config) throws ServletException {
		this.config = config;
	}*/
	
	
	public void doGet(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, IOException {
		/**
		 * 读取servlet的初始参数
		 */
		String path = this.getServletConfig().getInitParameter("path");

		File file = new File(path);
		//读取内容
		BufferedReader br = new BufferedReader(new FileReader(file));
		String str = null;
		while( (str=br.readLine())!=null ){
			System.out.println(str);
		}
		
		//查询当前servlet的所有初始化参数
		Enumeration<String> enums = this.getServletConfig().getInitParameterNames();
		while(enums.hasMoreElements()){
			String paramName = enums.nextElement();
			String paramValue = this.getServletConfig().getInitParameter(paramName);
			System.out.println(paramName+"="+paramValue);
		}
		//得到servlet的名称
		String servletName = this.getServletConfig().getServletName();
		System.out.println(servletName);
	}
}
```



## ServletContext对象

`ServletContext`对象 ,叫做Servlet的上下文对象。表示一个当前的web应用环境。一个web应用中只有一个ServletContext对象。

对象创建和得到
			创建时机：加载web应用时创建ServletContext对象。
			得到对象： 从ServletConfig对象的getServletContext方法得到
			

我们设计：
		创建ServletConfig对象

```java
public void init( ServletConfig config，ServletContext context ){  //多了一个参数
	得到ServletConfig对象
	得到ServletContext对象；
}
```



Sun公司设计

     		1. 创建ServletContext对象	  ServletContext  context = new ServletContext()		
                 		2. 创建ServletConfig对象        ServetConfig config = new ServletConfig();
                                                    config.setServletContxt(context);

```java
class  ServletConfig{
    ServletContext context;
    public ServletContext getServletContxt(){
        return contxt;
    }
} 

public void init( ServletConfig config ){
    得到ServletConfig对象
    从ServletConfig对象中得到ServletContext对象
    SerlvetContext context = config.getServletContext();
}
```



#### 核心API

```java
java.lang.String getContextPath()                               --得到当前web应用的路径

java.lang.String getInitParameter(java.lang.String name)        --得到web应用的初始化参数
java.util.Enumeration getInitParameterNames()  

RequestDispatcher getRequestDispatcher(java.lang.String path)    --转发（类似于重定向）

java.lang.String getRealPath(java.lang.String path)              --得到web应用的资源文件
ava.io.InputStream getResourceAsStream(java.lang.String path)  

void setAttribute(java.lang.String name, java.lang.Object object)--域对象有关的方法
java.lang.Object getAttribute(java.lang.String name)  
void removeAttribute(java.lang.String name)  
```

```java
package gz.ixfosa.context;

import java.io.IOException;

import javax.servlet.ServletContext;
import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
/**
 * 得到web应用路径
 *
 */
public class ContextDemo extends HttpServlet {

	public void doGet(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, IOException {
        
		//1.得到ServletContext对象
		//ServletContext context = this.getServletConfig().getServletContext();
		ServletContext context = this.getServletContext(); //（推荐使用）
		
		
		//2.得到web应用路径  /test
		/**
		 * web应用路径：部署到tomcat服务器上运行的web应用名称
		 */
		String contextPath = context.getContextPath();
		System.out.println(contextPath);
		
		
		/**
		 * 案例：应用到请求重定向
		 */
		response.sendRedirect(contextPath+"/index.html");
	}
}

```



#### web应用参数获取

web应用参数可以让当前web应用的所有servlet获取！！！

```xml
<!-- 配置web应用参数 -->
<context-param>
    <param-name>AAA</param-name>
    <param-value>AAA's value</param-value>
</context-param>
<context-param>
    <param-name>BBB</param-name>
    <param-value>BBB's value</param-value>
</context-param>
<context-param>
    <param-name>CCC</param-name>
    <param-value>CCC's value</param-value>
</context-param>

<servlet>
    <servlet-name>ContextDemo</servlet-name>
    <servlet-class>gz.ixfosa.context.ContextDemo</servlet-class>
</servlet>

<servlet-mapping>
    <servlet-name>ContextDemo</servlet-name>
    <url-pattern>/ContextDemo</url-pattern>
</servlet-mapping>
```

```java
package gz.ixfosa.context;

import java.io.IOException;
import java.util.Enumeration;

import javax.servlet.ServletContext;
import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
/**
 * 得到web应用参数
 *
 */
public class ContextDemo extends HttpServlet {

	public void doGet(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, IOException {
		//得到SErvletContext对象
		ServletContext context = this.getServletContext();
		
		System.out.println("参数"+context.getInitParameter("AAA"));
		
		Enumeration<String> enums = context.getInitParameterNames();
		while(enums.hasMoreElements()){
			String paramName = enums.nextElement();
			String paramValue  =context.getInitParameter(paramName);
			System.out.println(paramName+"="+paramValue);
		}
		
		//尝试得到ConfigDemo中的servlet参数
		String path = this.getServletConfig().getInitParameter("path");
		System.out.println("path="+path);
	}
}
```



#### 域对象

域对象：作用是用于保存数据，获取数据。可以在不同的动态资源之间共享数据。

- 所有域对象

```java
HttpServletRequet     域对象
ServletContext        域对象
HttpSession           域对象
PageContext           域对象	
```

ServletContext域对象：作用范围在`整个web应用`中有效！！！

- 域对象有关api

```java
保存数据：void setAttribute(java.lang.String name, java.lang.Object object)
    
获取数据： java.lang.Object getAttribute(java.lang.String name)  
    
删除数据： void removeAttribute(java.lang.String name)  
```

- 案例：   
  Servlet1, Servlet2 之间传递   name = ixfosa

  ​     response.sendRedirect("/Servlet2?name=ixfosa")        String request.getParameter("name");
  ​     保存到域对象中            从域对象获取

  

  ​	 方案1： 可以通过传递参数的形式，共享数据。局限：只能传递字符串类型。

  ​	 方案2： 可以使用域对象共享数据，好处：可以共享任何类型的数据！！！！！

​	

```java
package gz.ixfosa.context;

import java.io.IOException;

import javax.servlet.ServletContext;
import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
/**
 * 保存数据
 *
 */
public class ContextDemo extends HttpServlet {

	public void doGet(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, IOException {
        
		//1.得到域对象
		ServletContext context = this.getServletContext();
		
		//2.把数据保存到域对象中
		//context.setAttribute("name", "ixfosa");
		context.setAttribute("student", new Student("jacky",20));
		System.out.println("保存成功");
	}
}

class Student{
	private String name;
	private int age;
	public String getName() {
		return name;
	}
	public void setName(String name) {
		this.name = name;
	}
	public int getAge() {
		return age;
	}
	public void setAge(int age) {
		this.age = age;
	}
	public Student(String name, int age) {
		super();
		this.name = name;
		this.age = age;
	}
	@Override
	public String toString() {
		return "Student [age=" + age + ", name=" + name + "]";
	}
}
```

```java
package gz.ixfosa.f_context;

import java.io.IOException;

import javax.servlet.ServletContext;
import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

/**
 * 获取数据
 *
 */
public class ContextDemo2 extends HttpServlet {

	public void doGet(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, IOException {
        
		//1.得到域对象
		ServletContext context = this.getServletContext();
		
		//2.从域对象中取出数据
		String name = (String)context.getAttribute("name");
		Student student = (Student)context.getAttribute("student");
        
		System.out.println("name="+name);
		System.out.println(student);
	}
}
```



#### 转发和重定向

`转发`
		`RequestDispatcher` getRequestDispatcher(java.lang.String path)
		a）地址栏不会改变
		b）转发只能转发到当前web应用内的资源
		c）可以在转发过程中，可以把数据保存到request域对象中

`重定向`	`sendRedirect`		
			a）地址栏会改变，变成重定向到地址。
			b）重定向可以跳转到当前web应用，或其他web应用，甚至是外部域名网站。
			c）不能在重定向的过程，把数据保存到request中。

结论： 如果要使用request域对象进行数据共享，只能用转发技术！！！

- **转发**

```java
package gz.ixfosa.forward;

import java.io.IOException;

import javax.servlet.RequestDispatcher;
import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
/**
 * 转发（效果：跳转页面）
 *
 */
public class ForwardDemo extends HttpServlet {

   public void doGet(HttpServletRequest request, HttpServletResponse response)
         throws ServletException, IOException {
      
      /**
       * 保存数据到request域对象
       */
      request.setAttribute("name", "rose");
      
      //转发   
      /**
       * 注意：不能转发当前web应用以外的资源。
       */
       
      /*RequestDispatcher rd = this.getServletContext().getRequestDispatcher("/GetDataServlet");
      rd.forward(request, response);*/
       
    	this.getServletContext(). getRequestDispatcher("/GetDateServlet")
                                .forward(request, response);                        	     
   }
}

package gz.ixfosa.forward;

import java.io.IOException;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

public class GetDataServlet extends HttpServlet {

	public void doGet(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, IOException {

		/**
		 * 从request域对象中获取数据
		 */
		String name = (String)request.getAttribute("name");
		System.out.println("name="+name);
	}
}
```

- **重定向**

```java
package gz.ixfosa.forward;

import java.io.IOException;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

public class RedirectDemo extends HttpServlet {

	public void doGet(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, IOException {
		/**
		 * 保存数据到request域对象
		 */
		request.setAttribute("name", "rose");
		
		//重定向
		/**
		 * 注意：可以跳转到web应用内，或其他web应用，甚至其他外部域名。
		 */
		//response.sendRedirect("/test/adv.html");
		response.sendRedirect("/test/GetDataServlet");
	}
}
```

- **GetDataServlet**

```java
package gz.ixfosa.forward;

import java.io.IOException;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

public class RedirectDemo1 extends HttpServlet {

   public void doGet(HttpServletRequest request, HttpServletResponse response)
         throws ServletException, IOException {
      /**
       * 保存数据到request域对象
       */
      request.setAttribute("name", "rose");
      
      //重定向
      /**
       * 注意：可以跳转到web应用内，或其他web应用，甚至其他外部域名。
       */
      //response.sendRedirect("/day09/adv.html");
      response.sendRedirect("/day10/GetDataServlet");
   }

}
```

doPost和doGet

WebRoot\testMethod.html

```html
  <body>
  <h3>GET方式提交</h3>
    <form action="/test/RequestDemo" method="GET">
    	用户名：<input type="text" name="name"/><br/>
    	密码：<input type="password" name="password"/><br/>
    	<input type="submit" value="提交"/>
    </form>
    <hr/>
    
    <h3>POST方式提交</h3>
    <form action="/test/RequestDemo" method="POST">
    	用户名：<input type="text" name="name"/><br/>
    	密码：<input type="password" name="password"/><br/>
    	<input type="submit" value="提交"/>
    </form>
  </body>
```

```java
package gz.ixfosa.request;

import java.io.IOException;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

public class RequestDemo extends HttpServlet {
	
	@Override
	protected void service(HttpServletRequest req, HttpServletResponse resp)
			throws ServletException, IOException {
		System.out.println(req.getMethod());
		System.out.println("service方法被调用");
	}
	
	/**
	 * 该方法用于接收浏览器的Get方式提交的请求
	 */
	public void doGet(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, IOException {
		System.out.println("GET方式提交");
	}

	/**
	 * 该方法用于接收浏览器的Post方式提交的请求
	 */
	public void doPost(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, IOException {
		System.out.println("Post方式提交");
	}
}
```





## service 和 doXx方法

public abstract class `HttpServlet`

tomcat服务器首先会调用servlet的`service`方法，然后在service方法中再根据请求方式来分别调用对应的doXx方法

```java
protected void service(HttpServletRequest req, HttpServletResponse resp) throws ServletException, IOException {
    String method = req.getMethod();
    long lastModified;
    if (method.equals("GET")) {
        lastModified = this.getLastModified(req);
        if (lastModified == -1L) {
            this.doGet(req, resp);
        } else {
            long ifModifiedSince;
            try {
                ifModifiedSince = req.getDateHeader("If-Modified-Since");
            } catch (IllegalArgumentException var9) {
                ifModifiedSince = -1L;
            }

            if (ifModifiedSince < lastModified / 1000L * 1000L) {
                this.maybeSetLastModified(resp, lastModified);
                this.doGet(req, resp);
            } else {
                resp.setStatus(304);
            }
        }
    } else if (method.equals("HEAD")) {
        lastModified = this.getLastModified(req);
        this.maybeSetLastModified(resp, lastModified);
        this.doHead(req, resp);
    } else if (method.equals("POST")) {
        this.doPost(req, resp);
    } else if (method.equals("PUT")) {
        this.doPut(req, resp);
    } else if (method.equals("DELETE")) {
        this.doDelete(req, resp);
    } else if (method.equals("OPTIONS")) {
        this.doOptions(req, resp);
    } else if (method.equals("TRACE")) {
        this.doTrace(req, resp);
    } else {
        String errMsg = lStrings.getString("http.method_not_implemented");
        Object[] errArgs = new Object[]{method};
        errMsg = MessageFormat.format(errMsg, errArgs);
        resp.sendError(501, errMsg);
    }

}
```



## HttpServletRequest对象

HttpServletRequest对象作用是用于`获取请求数据`。



####  **核心的API：**

​		请求行： 

```java
request.getMethod();                                请求方式
    
request.getRequetURI();  / request.getRequetURL();  请求资源
    
request.getProtocol();                              请求http协议版本
```

​		请求头：

```java
request.getHeader("名称");   根据请求头获取请求值

request.getHeaderNames();   获取所有的请求头名称
```

​		实体内容:

```java
request.getInputStream();  获取实体内容数据
```

#### 获取请求行,头,实体

```java
package gz.ixfosa.request;

import java.io.IOException;
import java.io.InputStream;
import java.util.Enumeration;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

/**
 * 请求数据的获取
 *
 */
public class RequestDemo extends HttpServlet {

	/**
	 * 1)tomcat服务器接收到浏览器发送的请求数据，然后封装到HttpServetRequest对象
	 * 2）tomcat服务器调用doGet方法，然后把request对象传入到servlet中。
	 */
	public void doGet(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, IOException {
		/**
		 * 3)从request对象取出请求数据。
		 */
		//t1(request);
		
		t2(request); 
	}
	
	// 为了接收POST方式提交的请求
	@Override
	protected void doPost(HttpServletRequest request, HttpServletResponse resp)
			throws ServletException, IOException {
        
		/**
		 * 3.3 请求的实体内容
		 */
		InputStream in = request.getInputStream(); //得到实体内容
		byte[] buf = new byte[1024];
		int len = 0;
		while(  (len=in.read(buf))!=-1 ){
			String str = new String(buf,0,len);
			System.out.println(str);
		}
	}

	private void t2(HttpServletRequest request) {
        
		/**
		 * 3.2 请求头
		 */
		String host = request.getHeader("Host"); //根据头名称的到头的内容
		System.out.println(host);
		
		//遍历所有请求头
		Enumeration<String> enums = request.getHeaderNames(); //得到所有的请求头名称列表
		while(enums.hasMoreElements()){//判断是否有下一个元素
			String headerName = enums.nextElement(); //取出下一个元素
			String headerValue = request.getHeader(headerName);
			System.out.println(headerName+":"+headerValue);
		}
	}

	private void t1(HttpServletRequest request) {
        
		/**
		 * 3.1 请求行   格式：（GET /day09/hello HTTP/1.1）
		 */
		System.out.println("请求方式："+request.getMethod());//请求方式
		System.out.println("URI:"+request.getRequestURI());//请求资源
		System.out.println("URL:"+request.getRequestURL());
		System.out.println("http协议版本："+request.getProtocol());//http协议
	}
}
```





#### 请求参数获取	

GET方式： 参数放在URI后面
POST方式： 参数放在实体内容中

获取GET方式参数：

```java
request.getQueryString();  
```

获取POST方式参数：

```java
request.getInputStream();
```



问题：但是以上两种不通用，而且获取到的参数还需要进一步地解析。所以可以使用统一方便的获取参数的方式：
		 核心的API：

```java
request.getParameter("参数名");  根据参数名获取参数值（注意，只能获取一个值的参数）
    
request.getParameterValue("参数名“)；根据参数名获取参数值（可以获取多个值的参数）

request.getParameterNames();   获取所有参数名称列表  
```

WebRoot/testParameter.html

```html
<body>
  <h3>GET方式提交</h3>
    <form action="/tset/RequestDemo" method="GET">
    	用户名：<input type="text" name="name"/><br/>
    	密码：<input type="password" name="password"/><br/>
    	性别：
    	<input type="radio" name="gender" value="男"/>男
    	<input type="radio" name="gender" value="女"/>女<br/>
    	籍贯：
    		<select name="jiguan">
    			<option value="广东">广东</option>
    			<option value="广西">广西</option>
    			<option value="湖南">湖南</option>
    		</select>
    		<br/>
    	爱好：
    	<input type="checkbox" name="hobit" value="篮球"/>篮球
    	<input type="checkbox" name="hobit" value="足球"/>足球
    	<input type="checkbox" name="hobit" value="羽毛球"/>羽毛球<br/>
    	个人简介：
    	<textarea rows="5" cols="10" name="info"></textarea><br/>
    	<!-- 隐藏域 -->
    	<input type="hidden" name="id" value="001"/>
    	<input type="submit" value="提交"/>
    </form>
    <hr/>
    
    
    <h3>POST方式提交</h3>
    <form action="/day09/RequestDemo" method="POST">
    	用户名：<input type="text" name="name"/><br/>
    	密码：<input type="password" name="password"/><br/>
    	性别：
    	<input type="radio" name="gender" value="男"/>男
    	<input type="radio" name="gender" value="女"/>女<br/>
    	籍贯：
    		<select name="jiguan">
    			<option value="广东">广东</option>
    			<option value="广西">广西</option>
    			<option value="湖南">湖南</option>
    		</select>
    		<br/>
    	爱好：
    	<input type="checkbox" name="hobit" value="篮球"/>篮球
    	<input type="checkbox" name="hobit" value="足球"/>足球
    	<input type="checkbox" name="hobit" value="羽毛球"/>羽毛球<br/>
    	个人简介：
    	<textarea rows="5" cols="10" name="info"></textarea><br/>
    	<!-- 隐藏域 -->
    	<input type="hidden" name="id" value="001"/>
    	<input type="submit" value="提交"/>
    </form>
  </body>
```



```java
package gz.ixfosa.request;

import java.io.IOException;
import java.io.InputStream;
import java.util.Enumeration;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

/**
 * 获取GET方式和Post方式提交的参数
 *
 */
public class RequestDemo extends HttpServlet {

	public void doGet(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, IOException {
        
		/**
		 * 设置参数查询的编码
		 * 该方法只能对请求实体内容的数据编码起作用。
		 * POST提交的数据在实体内容中，所以该方法对POST方法有效！
		 * GET方法的参数放在URI后面，所以对GET方式无效！！！
		 */
		request.setCharacterEncoding("utf-8");
		
		
		//接收GET方式提交的参数
		String value = request.getQueryString();
		System.out.println(value);
		
        
		/**
		 * 统一方便地获取请求参数的方法
		 */
		System.out.println(request.getMethod()+"方式");
        
		//getParameter(name): 根据参数名得到参数值(只能获取一个值的参数)
		String name = request.getParameter("name");
		/**
		 * 手动重新解码(iso-8859-1 字符串-> utf-8 字符串)
		 */
		String password = request.getParameter("password");
		
		if("GET".equals(request.getMethod())){
            name = new String(name.getBytes("iso-8859-1"),"utf-8");
			password = new String(password.getBytes("iso-8859-1"),"utf-8");
		}
		System.out.println(name+"="+password);
		
        
		Enumeration<String> enums = request.getParameterNames();
		while( enums.hasMoreElements() ){
            
			String paramName = enums.nextElement();
			
			//如果参数名是hobit，则调用getParameterValues
			if("hobit".equals(paramName)){
				/**
				 * getParameterValues(name): 根据参数名获取参数值（可以获取多个值的同名参数）
				 */
				System.out.println(paramName+":");
				String[] hobits = request.getParameterValues("hobit");
				for(String h: hobits){
					if("GET".equals(request.getMethod())){
						h = new String(h.getBytes("iso-8859-1"),"utf-8");
					}
					System.out.print(h+",");
				}
				System.out.println();
                
				//如果不是hobit，则调用getParameter
			}else{
				String paramValue = request.getParameter(paramName);
				
				if("GET".equals(request.getMethod())){
					paramValue = new String(paramValue.getBytes("iso-8859-1"),"utf-8");
				}
				System.out.println(paramName+"="+paramValue);
			}
		}
	}
    
	public void doPost(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, IOException {
        
		//接收POST方式提交的参数
		InputStream in = request.getInputStream();
		byte[] buf = new byte[1024];
		int len = 0;
		while(  (len=in.read(buf))!=-1 ){
			System.out.println(new String(buf,0,len));
		}
        
		
		/**
		 * 统一方便地获取请求参数的方法
		 */
        
		//根据参数名得到参数值
		String name = request.get("name");
		String password = request.getParameter("password");
		System.out.println(name+"="+password);
		
        
		Enumeration<String> enums = request.getParameterNames();
		while( enums.hasMoreElements() ){
			String paramName = enums.nextElement();
			String paramValue = request.getParameter(paramName);
			System.out.println(paramName+"="+paramValue);
		}
		//一定调用doGet方式
		this.doGet(request, response);
	}
}
```

#### 请求参数编码

修改POST方式参数编码：

```java
request.setCharacterEncoding("utf-8");
```

修改GET方式参数编码：

```java
String name = new String(name.getBytes("iso-8859-1"),"utf-8");  //	手动解码：
```

​					

#### 浏览器类型（user-agent）

```java
package gz.ixfosa.request;

import java.io.IOException;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
/**
 * 案例-获取浏览器的类型
 *
 */
public class RequestDemo extends HttpServlet {

	public void doGet(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, IOException {
        
		response.setContentType("text/html;charset=utf-8");
        
		//获取请求头： user-agent
		String userAgent = request.getHeader("user-agent");
		System.out.println(userAgent);
		
		//判断用户使用的浏览器类型
		if(userAgent.contains("Firefox")){
			response.getWriter().write("你正在使用火狐浏览器");
		}else if(userAgent.contains("Chrome")){
			response.getWriter().write("你正在使用谷歌浏览器");
		}else if(userAgent.contains("Trident")){
			response.getWriter().write("你正在使用IE浏览器");
		}else{
			response.getWriter().write("地球上没有这个浏览器，建议使用火狐浏览器");
		}
	}
}
```



#### 案防止非法链接(referer)

非法链接：							
		1）直接访问下载的资源
		2）不是从广告页面过来的链接

`referer`： 当前请求来自于哪里。



WebRoot/adv.html

```html
  <body>
    广告内容，请猛戳这里。<br/>
    <a href="/test/RequestDemo">点击此处下载</a>
  </body>
```

```java
package gz.ixfosa.request;

import java.io.IOException;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
/**
 * 案例- 防止非法链接
 * 这是需要下载的资源
 *
 */
public class RequestDemo extends HttpServlet {

	public void doGet(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, IOException {
        
		response.setContentType("text/html;charset=utf-8");
		
		//得到referer头
		String referer = request.getHeader("referer");
		System.out.println("referer="+referer);
		
		/**
		 * 判断非法链接：
		 * 	1）直接访问的话referer=null
		 *  2）如果当前请求不是来自广告   
		 */
		if(referer==null || !referer.contains("/t/adv.html")){
			response.getWriter().write("当前是非法链接，请回到首页。<a href='/day09/adv.html'>首页</a>");
		}else{
			//正确的链接
			response.getWriter().write("资源正在下载...");
		}
	
	}

}

```



##  HttpServletResponse对象

`HttpServletResponse`对象: 修改响应信息：

#### 核心的API：

响应行

```java
response.setStatus()  设置状态码
```

响应头： 

```java
response.setHeader("name","value")  设置响应头
```

实体内容：

```java
response.getWriter().writer();   发送字符实体内容
response.getOutputStream().writer()  发送字节实体内容 
```

```java
package gz.ixfosa.response;

import java.io.IOException;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
/**
 * 设置响应信息
 *
 */
public class ResponseDemo extends HttpServlet {

	/**
	 * 1)tomcat服务器把请求信息封装到HttpServletRequest对象，且把响应信息封装到
	 *   HttpServletResponse
	 * 2）tomcat服务器调用doGet方法，传入request，和response对象
	 */
	public void doGet(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, IOException {
		/**
		 * 3）通过response对象改变响应信息
		 */
		/**
		
		 * 3.1 响应行
		 */
		response.setStatus(404);//修改状态码
		response.sendError(404); //发送404的状态码+404的错误页面
		
	
		/**
		 * 3.2 响应头
		 */
		response.setHeader("server", "JBoss");
		
		
		/**
		 * 3.3 实体内容(浏览器直接能够看到的内容就是实体内容)
		 */
		//response.getWriter().write("01.hello world"); //字符内容。
		response.getOutputStream().write("02.hello world".getBytes());//字节内容
		 
	}
	
	/**
	 * 4)tomcat服务器把response对象的内容转换成响应格式内容，再发送给浏览器解析。
	 */
}
```

#### 请求重定向（Location）

```java
package gz.ixfosa.response;

import java.io.IOException;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
/**
 * 案例- 请求重定向
 * （相当于超链接跳转页面）
 *
 */
public class ResponseDemo extends HttpServlet {

	public void doGet(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, IOException {
        
		/**
		 * 需求： 跳转到adv.html
		 * 使用请求重定向： 发送一个302状态码+location的响应头
		 */
        
		response.setStatus(302);//发送一个302状态码
		response.setHeader("location", "/test/adv.html"); //location的响应头
        
	
		//请求重定向简化写法
		//response.sendRedirect("/test/adv.html");
	}
}
```

#### 定时刷新(refresh)

```java
package gz.ixfosa.response;

import java.io.IOException;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
/**
 * 案例- 定时刷新
 *
 */
public class ResponseDemo extends HttpServlet {

	public void doGet(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, IOException {
        
		/**
		 * 定时刷新
		 * 原理：浏览器认识refresh头，得到refresh头之后重新请求当前资源
		 */
		//response.setHeader("refresh", "1"); //每隔1秒刷新次页面
		
		/**
		 * 隔n秒之后跳转另外的资源
		 */
		response.setHeader("refresh", "3;url=/test/adv.html");//隔3秒之后跳转到adv.html
	}

}
```

#### content-Type作用

```java
package gz.ixfosa.response;

import java.io.File;
import java.io.FileInputStream;
import java.io.IOException;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
/**
 * 案例- content-Type作用
 *
 */
public class ResponseDemo extends HttpServlet {

	public void doGet(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, IOException {
		/**
		 * 设置响应实体内容编码
		 */
		response.setCharacterEncoding("utf-8");
		
		/**
		 * 1. 服务器发送给浏览器的数据类型和内容编码
		 */
		//response.setHeader("content-type", "text/html");
		//response.setContentType("text/html;charset=utf-8");//和上面代码等价。推荐使用此方法
		//response.setContentType("text/xml");
		response.setContentType("image/jpg");

		
		//response.getWriter().write("<html><head><title>this is tilte</title></head><body>中国</body></html>");
        
		//response.getOutputStream().write("<html><head><title>this is tilte</title></head><body>中国</body></html>".getBytes("utf-8"));

        
		 File file = new File("e:/mm.jpg");
		/**
		 * 设置以下载方式打开文件
		 */
		response.setHeader("Content-Disposition", "attachment; filename="+file.getName());
       
		/**
		 * 发送图片
		 */
		FileInputStream in = new FileInputStream(file);
		byte[] buf = new byte[1024];
		int len = 0;
		
		//把图片内容写出到浏览器
		while( (len=in.read(buf))!=-1 ){
			response.getOutputStream().write(buf, 0, len);
		}
	}
}
```

## web应用中路径问题

```java
package gz.ixfosa.path;

import java.io.IOException;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
/**
 * web应用中路径问题
 *
 */
public class PathDemo extends HttpServlet {

	public void doGet(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, IOException {
        
		response.setContentType("text/html;charset=utf-8");
        
		//目标资源： target.html
		/**
		 * 思考： 目标资源是给谁使用的。
		 * 		给服务器使用的：   / 表示在当前web应用的根目录（webRoot下）
		 * 		给浏览器使用的：   /  表示在webapps的根目录下
		 */
		/**
		 * 1.转发
		 */
		//request.getRequestDispatcher("/target.html").forward(request, response);
		
		
		/**
		 * 2.请求重定向
		 */
		//response.sendRedirect("/test/target.html");
		
		/**
		 * 3.html页面的超连接href
		 */
		response.getWriter().write("<html><body><a href='/test/target.html'>超链接</a></body></html>");
		
		/**
		 * 4.html页面中的form提交地址
		 */
		response.getWriter().write("<html><body><form action='/test/target.html'><input type='submit'/></form></body></html>");
	}
}
```

```xml
<servlet>
    <servlet-name>PathDemo</servlet-name>
    <servlet-class>gz.ixfosa.path.PathDemo</servlet-class>
</servlet>

<servlet-mapping>
    <servlet-name>PathDemo</servlet-name>
    <url-pattern>/PathDemo</url-pattern>
</servlet-mapping>
```



## 读取web应用下的资源文件

src/db.properties

```properties
user=root
password=root
```

```java
package gz.ixfosa.resource;

import java.io.IOException;
import java.io.InputStream;
import java.util.Properties;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
/**
 * 读取web应用下的资源文件（例如properties）
 */

public class ResourceDemo extends HttpServlet {

	public void doGet(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, IOException {
        
		/**
		 *  '.' 代表java命令运行目录。java运行命令在哪里？？ 在tomcat/bin目录下
		 *    结论： 在web项目中， '.' 代表在tomcat/bin目录下开始，所以不能使用这种相对路径。
		 */
		
		
		//读取文件。在web项目下不要这样读取。因为.表示在tomcat/bin目录下
		/*File file = new File("./src/db.properties");
		FileInputStream in = new FileInputStream(file);*/
		
		/**
		 * 使用web应用下加载资源文件的方法
		 */
		/**
		 * 1. getRealPath读取,返回资源文件的绝对路径
		 */
		/*String path = this.getServletContext().getRealPath("/WEB-INF/classes/db.properties");
		System.out.println(path);
		File file = new File(path);
		FileInputStream in = new FileInputStream(file);*/
		
		/** 
		 * 2. getResourceAsStream() 得到资源文件，返回的是输入流
		 */
		InputStream in = this.getServletContext().getResourceAsStream("/WEB-INF/classes/db.properties");
		
		
		Properties prop = new Properties();
		//读取资源文件
		prop.load(in);
		
		String user = prop.getProperty("user");
		String password = prop.getProperty("password");
		System.out.println("user="+user);
		System.out.println("password="+password);
	}
}
```

```xml
<servlet>
    <servlet-name>ResourceDemo</servlet-name>
    <servlet-class>gz.ixfosa.resource.ResourceDemo</servlet-class>
</servlet>

<servlet-mapping>
    <servlet-name>ResourceDemo</servlet-name>
    <url-pattern>/ResourceDemo</url-pattern>
</servlet-mapping>
```

## Servlet-通讯录

### 1. 需求分析

功能分析：

​		1）添加联系人

​		2）修改联系人

​		3）删除联系人

​		4）查询所有联系人

![功能流转](E:/smile/java/images/功能流转.png)

### 2. 需求设计

1. 设计实体（抽象实体）

```java
//联系人实体：
    class Contact{
        private String id;
        private String name;
        private String gender;
        private int age;
        private String phone;
        private String email;
        private String qq;
    }
```

2. 设计“数据库”，（xml代替"数据库"）

```xml
contact.xml
<contactList>
    <contact id="1">
        <name>张三</name>
        <gender>男</gender>
        <age>20</age>
        <phone>13433334444</phone>
        <email>zs@qq.com</email>
        <qq>43222222<qq>
     </contact>
</contactList>
```

3. 设计涉及的接口

   ​		DAO接口（数据访问对象）：实体对象的CRUD方法。
   项目原则： 通常一个实体对象就会对应一个DAO接口和一个DAO实现类

```java
interface ContactDao{
    public void addContact(Contact contact);//添加联系人
    public void updateContact(Contact contact);//修改联系人
    public void deleteContact(String id);//删除联系人
    public List<Contact> findAll();  //查询所有联系人
    public Contact findById(String id);//根据编号查询联系人
}
```

4. 设计项目的目录结构

```
项目名称： contactSys_web
目录结构：
|- contactSys_web
	|-src
		|-gz.ixfosa.contactSys_web.entity       : 存放实体对象
		|-gz.ixfosa.contactSys_web.dao          : 存放dao的接口
        |-gz.ixfosa.contactSys_web.dao.impl     : 存放dao的实现类
        |-gz.ixfosa.contactSys_web.servlet      : 存放servlet的类
        |-gz.ixfosa.contactSys_web.test         : 存放单元测试类
        |-gz.ixfosa.contactSys_web.util         : 存放工具类
        |-gz.ixfosa.contactSys_web.exception    : 存放自定义异常类
    |-WebRoot
        |-html    文件
        |-images：目录。存放图片资源		
        |-css：   目录。存放css资源
        |-js：    目录。存放js资源
```

### 3. 编码实现

#### jar

WebRoot/WEB-INF/lib/

```
dom4j-1.6.1.jar
jaxen-1.1-beta-6.jar
```



#### util

e:/contact.xml

```xml
xml<contactList>
    <contact id="1">
        <name>张三</name>
        <gender>男</gender>
        <age>20</age>
        <phone>13433334444</phone>
        <email>zs@qq.com</email>
        <qq>43222222<qq>
     </contact>
</contactList>
```

src/gz/ixfosa/contactSys_web/util/XMLUtil.java

```java
package gz.ixfosa.contactSys_web.util;

import java.io.File;
import java.io.FileOutputStream;

import org.dom4j.Document;
import org.dom4j.DocumentException;
import org.dom4j.io.OutputFormat;
import org.dom4j.io.SAXReader;
import org.dom4j.io.XMLWriter;

/**
 * xml的工具类
 *
 */
public class XMLUtil {
	
	/**
	 * 读取xml文档方法
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
	 * 写出到xml文档中
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



#### entity

src/gz/ixfosa/contactSys_web/entity/Contact.java

```java
package gz.ixfosa.contactSys_web.entity;
/**
 * 实体对象
 *
 */
public class Contact {

	private String id;
	private String name;
	private String gender;
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

#### dao

src/gz/ixfosa/contactSys_web/dao/ContactDao.java

```java
package gz.ixfosa.contactSys_web.dao;

import gz.itcast.contactSys_web.entity.Contact;

import java.util.List;

/**
 * 联系人的DAO接口
 *
 */
public interface ContactDao {
	public void addContact(Contact contact);//添加联系人
	public void updateContact(Contact contact);//修改联系人
	public void deleteContact(String id);//删除联系人
	public List<Contact> findAll();  //查询所有联系人
	public Contact findById(String id);//根据编号查询联系人
}

```

src/gz/ixfosa/contactSys_web/dao/impl/ContactDaoImpl.java

```java
package gz.ixfosa.contactSys_web.dao.impl;

import gz.itcast.contactSys_web.dao.ContactDao;
import gz.itcast.contactSys_web.entity.Contact;
import gz.itcast.contactSys_web.util.XMLUtil;

import java.io.File;
import java.util.ArrayList;
import java.util.List;
import java.util.UUID;

import org.dom4j.Document;
import org.dom4j.DocumentHelper;
import org.dom4j.Element;

public class ContactDaoImpl implements ContactDao {

	/**
	 * 添加联系人
	 */
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
			
			/**
			 * 由系统自动生成随机且唯一的ID值，赋值给联系人
			 */
			String uuid = UUID.randomUUID().toString().replace("-","");
			
			contactElem.addAttribute("id", uuid);
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
	 * 删除联系人
	 */
	public void deleteContact(String id) {
		try {
			//1.读取xml文件
			Document doc = XMLUtil.getDocument();
			//2.查询需要删除id的contact
			Element contactElem = (Element)doc.selectSingleNode("//contact[@id='"+id+"']");
			//删除标签
			if(contactElem!=null){
				contactElem.detach();
			}
			
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

	/**
	 * 根据编号查询联系人
	 */
	public Contact findById(String id) {
		Document doc = XMLUtil.getDocument();
		Element e = (Element)doc.selectSingleNode("//contact[@id='"+id+"']");
		
		Contact c = null;
		if(e!=null){
			//创建COntact对象
			c = new Contact();
			c.setId(e.attributeValue("id"));
			c.setName(e.elementText("name"));
			c.setGender(e.elementText("gender"));
			c.setAge(Integer.parseInt(e.elementText("age")));
			c.setPhone(e.elementText("phone"));
			c.setEmail(e.elementText("email"));
			c.setQq(e.elementText("qq"));
		}
		return c;
	}

	/**
	 * 修改联系人
	 */
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
	
	public static void main(String[] args) {
		//测试UUID
		String uuid = UUID.randomUUID().toString().replace("-","");
		System.out.println(uuid);
	}

}
```

#### servlet

src/gz/ixfosa/contactSys_web/servlet/ListContactServlet.java

```java
package gz.ixfosa.contactSys_web.servlet;

import gz.itcast.contactSys_web.dao.ContactDao;
import gz.itcast.contactSys_web.dao.impl.ContactDaoImpl;
import gz.itcast.contactSys_web.entity.Contact;

import java.io.IOException;
import java.io.PrintWriter;
import java.util.List;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
/**
 * 显示所有联系人的逻辑
 *
 */
public class ListContactServlet extends HttpServlet {

	public void doGet(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, IOException {
		//1.从xml中读取出联系人数据
		ContactDao dao = new ContactDaoImpl();
		List<Contact> list = dao.findAll();
		
		//2.显示到浏览器
		response.setContentType("text/html;charset=utf-8");
		PrintWriter writer = response.getWriter();
		
		String html = "";
		
		//shift+alt+A   ^(.*)$  \1";
		html += "<!DOCTYPE html PUBLIC '-//W3C//DTD XHTML 1.0 Transitional//EN' 'http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd'>";
		html += "<html xmlns='http://www.w3.org/1999/xhtml'>";
		html += "<head>";
		html += "<meta http-equiv='Content-Type' content='text/html; charset=utf-8' />";
		html += "<title>查询所有联系人</title>";
		html += "<style type='text/css'>";
		html += "	table td{";
		html += "		/*文字居中*/";
		html += "		text-align:center;";
		html += "	}";
		html += "	";
		html += "	/*合并表格的边框*/";
		html += "	table{";
		html += "		border-collapse:collapse;";
		html += "	}";
		html += "</style>";
		html += "</head>";
		html += "";
		html += "<body>";
		html += "<center><h3>查询所有联系人</h3></center>";
		html += "<table align='center' border='1' width='800px'>";
		html += "	<tr>";
		html += "    	<th>编号</th>";
		html += "        <th>姓名</th>";
		html += "        <th>性别</th>";
		html += "        <th>年龄</th>";
		html += "        <th>电话</th>";
		html += "        <th>邮箱</th>";
		html += "        <th>QQ</th>";
		html += "        <th>操作</th>";
		html += "    </tr>";
		if(list!=null){
			for (Contact contact : list) {
				html += "    <tr>";
				html += "    	<td>"+contact.getId()+"</td>";
				html += "        <td>"+contact.getName()+"</td>";
				html += "        <td>"+contact.getGender()+"</td>";
				html += "        <td>"+contact.getAge()+"</td>";
				html += "        <td>"+contact.getPhone()+"</td>";
				html += "        <td>"+contact.getEmail()+"</td>";
				html += "        <td>"+contact.getQq()+"</td>";
				html += "        <td><a href='"+request.getContextPath()+"/QueryContactServlet?id="+contact.getId()+"'>修改</a>&nbsp;<a href='"+request.getContextPath()+"/DeleteContactServlet?id="+contact.getId()+"'>删除</a></td>";
				html += "    </tr>";
			}
		}
		html += "    <tr>";
		html += "    	<td colspan='8' align='center'><a href='"+request.getContextPath()+"/addContact.html'>[添加联系人]</a></td>";
		html += "    </tr>";
		html += "</table>";
		html += "</body>";
		html += "</html>";

		
		writer.write(html);
		
	}

	public void doPost(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, IOException {
		doGet(request, response);
	}

}
```

src/gz/ixfosa/contactSys_web/servlet/QueryContactServlet.java

```java
package gz.ixfosa.contactSys_web.servlet;

import gz.itcast.contactSys_web.dao.ContactDao;
import gz.itcast.contactSys_web.dao.impl.ContactDaoImpl;
import gz.itcast.contactSys_web.entity.Contact;

import java.io.IOException;
import java.io.PrintWriter;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
/**
 * 修改前查询联系人的逻辑
 *
 */
public class QueryContactServlet extends HttpServlet {

	public void doGet(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, IOException {
		//1.接收id
		String id = request.getParameter("id");
		
		//2.调用dao根据id查询联系人的方法
		ContactDao dao = new ContactDaoImpl();
		Contact contact = dao.findById(id);
		
		//3.把联系人显示到浏览器中
		response.setContentType("text/html;charset=utf-8");
		PrintWriter writer = response.getWriter();
		
		String html = "";
		
		html += "<!DOCTYPE html PUBLIC '-//W3C//DTD XHTML 1.0 Transitional//EN' 'http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd'>";
		html += "<html xmlns='http://www.w3.org/1999/xhtml'>";
		html += "<head>";
		html += "<meta http-equiv='Content-Type' content='text/html; charset=utf-8' />";
		html += "<title>修改联系人</title>";
		html += "</head>";
		html += "";
		html += "<body>";
		html += "<center><h3>修改联系人</h3></center>";
		html += "<form action='"+request.getContextPath()+"/UpdateContactServlet' method='post'>";
		//注意：添加id的隐藏域
		html += "<input type='hidden' name='id' value='"+contact.getId()+"'/>";
		html += "<table align='center' border='1' width='300px'>";
		html += "    <tr>";
		html += "    	<th>姓名</th>";
		html += "        <td><input type='text' name='name' value='"+contact.getName()+"'/></td>";
		html += "    </tr>";
		html += "    <tr>";
		html += "    	<th>性别</th>";
		html += "        <td>";
		
		if(contact.getGender().equals("男")){
			html += "        <input type='radio' name='gender' value='男' checked='checked'/>男";
			html += "        <input type='radio' name='gender' value='女'/>女";
		}else if(contact.getGender().equals("女")){
			html += "        <input type='radio' name='gender' value='男'/>男";
			html += "        <input type='radio' name='gender' value='女' checked='checked'/>女";
		}else{
			html += "        <input type='radio' name='gender' value='男' checked='checked'/>男";
			html += "        <input type='radio' name='gender' value='女'/>女";
		}
	
		html += "        </td>";
		html += "    </tr>";
		html += "    <tr>";
		html += "    	<th>年龄</th>";
		html += "        <td><input type='text' name='age' value='"+contact.getAge()+"'/></td>";
		html += "    </tr>";
		html += "    <tr>";
		html += "    	<th>电话</th>";
		html += "        <td><input type='text' name='phone' value='"+contact.getPhone()+"'/></td>";
		html += "    </tr>";
		html += "    <tr>";
		html += "    	<th>邮箱</th>";
		html += "        <td><input type='text' name='email' value='"+contact.getEmail()+"'/></td>";
		html += "    </tr>";
		html += "    <tr>";
		html += "    	<th>QQ</th>";
		html += "        <td><input type='text' name='qq' value='"+contact.getQq()+"'/></td>";
		html += "    </tr>";
		html += "    <tr>";
		html += "        <td colspan='2' align='center'>";
		html += "        <input type='submit' value='保存'/>&nbsp;";
		html += "        <input type='reset' value='重置'/></td>";
		html += "    </tr>";
		html += "</table>";
		html += "</form>";
		html += "</body>";
		html += "</html>";

		
		writer.write(html);
	}

	public void doPost(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, IOException {
		doGet(request, response);
	}

}
```

src/gz/ixfosa/contactSys_web/servlet/AddContactServlet.java

```java
package gz.itcixfosaast.contactSys_web.servlet;

import gz.itcast.contactSys_web.dao.ContactDao;
import gz.itcast.contactSys_web.dao.impl.ContactDaoImpl;
import gz.itcast.contactSys_web.entity.Contact;

import java.io.IOException;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
/**
 * 添加联系人的逻辑
 *
 */
public class AddContactServlet extends HttpServlet {

	public void doGet(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, IOException {
		request.setCharacterEncoding("utf-8");
		//1.接收参数
		String name = request.getParameter("name");
		String gender = request.getParameter("gender");
		String age = request.getParameter("age");
		String phone = request.getParameter("phone");
		String email = request.getParameter("email");
		String qq = request.getParameter("qq");
		
		//封装成Contact对象
		Contact contact = new Contact();
		contact.setName(name);
		contact.setGender(gender);
		contact.setAge(Integer.parseInt(age));
		contact.setPhone(phone);
		contact.setEmail(email);
		contact.setQq(qq);
		
		//2.调用dao类的添加联系人的方法
		ContactDao dao = new ContactDaoImpl();
		dao.addContact(contact);
		
		//3.跳转到查询联系人的页面
		response.sendRedirect(request.getContextPath()+"/ListContactServlet");
	}

	public void doPost(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, IOException {
		doGet(request, response);
	}
}
```

src/gz/ixfosa/contactSys_web/servlet/DeleteContactServlet.java

```java
package gz.ixfosa.contactSys_web.servlet;

import gz.itcast.contactSys_web.dao.ContactDao;
import gz.itcast.contactSys_web.dao.impl.ContactDaoImpl;

import java.io.IOException;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
/**
 * 删除联系人的逻辑
 *
 */
public class DeleteContactServlet extends HttpServlet {

	public void doGet(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, IOException {
		//在火狐浏览器中以Get方式提交带参数的数据，会重复提交两次。
		System.out.println("删除联系人");
		//1.接收id
		String id = request.getParameter("id");
		
		//2.调用dao删除联系人的方法
		ContactDao dao = new ContactDaoImpl();
		dao.deleteContact(id);
		
		//3.跳转到查询联系人的页面
		response.sendRedirect(request.getContextPath()+"/ListContactServlet");
	}

	public void doPost(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, IOException {
		doGet(request, response);
	}
}
```

src/gz/ixfosa/contactSys_web/servlet/UpdateContactServlet.java

```java
package gz.ixfosa.contactSys_web.servlet;

import gz.itcast.contactSys_web.dao.ContactDao;
import gz.itcast.contactSys_web.dao.impl.ContactDaoImpl;
import gz.itcast.contactSys_web.entity.Contact;

import java.io.IOException;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
/**
 * 修改联系人的逻辑
 *
 */
public class UpdateContactServlet extends HttpServlet {

	public void doGet(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, IOException {
		request.setCharacterEncoding("utf-8");
		//1.接收参数
		String id = request.getParameter("id");
		String name = request.getParameter("name");
		String gender = request.getParameter("gender");
		String age = request.getParameter("age");
		String phone = request.getParameter("phone");
		String email = request.getParameter("email");
		String qq = request.getParameter("qq");
		
		//封装成Contact对象
		Contact contact = new Contact();
		contact.setId(id);
		contact.setName(name);
		contact.setGender(gender);
		contact.setAge(Integer.parseInt(age));
		contact.setPhone(phone);
		contact.setEmail(email);
		contact.setQq(qq);
		
		//2.调用dao修改联系人的方法
		ContactDao dao = new ContactDaoImpl();
		dao.updateContact(contact);
		
		//3.跳转到查询联系人的页面
		response.sendRedirect(request.getContextPath()+"/ListContactServlet");
		
	}

	public void doPost(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, IOException {
		doGet(request, response);
	}
}
```

#### test

src/gz/ixfosa/contactSys_web/test/TestContactOperatorImpl.java

```java
package gz.ixfosa.contactSys_web.test;

import gz.itcast.contactSys_web.dao.ContactDao;
import gz.itcast.contactSys_web.dao.impl.ContactDaoImpl;
import gz.itcast.contactSys_web.entity.Contact;

import java.util.List;

import org.junit.Before;
import org.junit.Test;

/**
 * 联系人操作实现类的测试类
 *
 */
public class TestContactOperatorImpl {
	ContactDao operator = null;
	
	/**
	 * 初始化对象实例
	 */
	@Before
	public void init(){
		operator = new ContactDaoImpl();
	}
	

	@Test
	public void testAddContact(){
		Contact contact = new Contact();
		//contact.setId("2");
		contact.setName("张三2");
		contact.setGender("男");
		contact.setAge(20);
		contact.setPhone("134222233333");
		contact.setEmail("ixfosa@qq.com");
		contact.setQq("33334444");
		operator.addContact(contact);
	}
	
	@Test
	public void testUpdateContact(){
		Contact contact = new Contact();
		contact.setId("1"); //修改的ID
		contact.setName("李四");
		contact.setGender("女");
		contact.setAge(30);
		contact.setPhone("135222233333");
		contact.setEmail("zhangsan@qq.com");
		contact.setQq("33334444");
		operator.updateContact(contact);
	}
	
	@Test
	public void testDeleteContact(){
		operator.deleteContact("2");
	}
	
	@Test
	public void testFindAll(){
		List<Contact> list = operator.findAll();
		for (Contact contact : list) {
			System.out.println(contact);
		}
	}
	
	@Test
	public void testFindById(){
		Contact contact = operator.findById("1");
		System.out.println(contact);
	}
}
```



##### jsp

WebRoot/listContact.jsp

```jsp
<%@ page language="java" import="java.util.*,gz.itcast.contactSys_web.entity.*" pageEncoding="utf-8"%>
<%@ taglib uri="http://java.sun.com/jsp/jstl/core" prefix="c" %>
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
<title>查询所有联系人</title>
<style type="text/css">
	table td{
		/*文字居中*/
		text-align:center;
	}
	
	/*合并表格的边框*/
	table{
		border-collapse:collapse;
	}
</style>
</head>

<body>
<center><h3>查询所有联系人(jsp版本)</h3></center>
<table align="center" border="1" width="700px">
	<tr>
    	<th>编号</th>
        <th>姓名</th>
        <th>性别</th>
        <th>年龄</th>
        <th>电话</th>
        <th>邮箱</th>
        <th>QQ</th>
        <th>操作</th>
    </tr>
    <c:forEach items="${contacts}" var="con" varStatus="varSta">
    <tr>
    	<td>${varSta.count }</td>
        <td>${con.name }</td>
        <td>${con.gender }</td>
        <td>${con.age }</td>
        <td>${con.phone }</td>
        <td>${con.email }</td>
        <td>${con.qq }</td>
        <td><a href="${pageContext.request.contextPath }/QueryContactServlet?id=${con.id}">修改</a>&nbsp;<a href="${pageContext.request.contextPath }/DeleteContactServlet?id=${con.id}">删除</a></td>
    </tr>
    </c:forEach>
    <tr>
    	<td colspan="8" align="center"><a href="${pageContext.request.contextPath }/addContact.jsp">[添加联系人]</a></td>
    </tr>
</table>
</body>
</html>
```

WebRoot/addContact.jsp

```jsp
<%@ page language="java" import="java.util.*,gz.itcast.contactSys_web.entity.*" pageEncoding="utf-8"%>
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
<title>添加联系人</title>
</head>

<body>
<center><h3>添加联系人</h3></center>
<form action="${pageContext.request.contextPath }/AddContactServlet" method="post">
<table align="center" border="1" width="400px">
    <tr>
    	<th>姓名</th>
        <td><input type="text" name="name"/><font color="red">${msg }</font></td>
    </tr>
    <tr>
    	<th>性别</th>
        <td>
        <input type="radio" name="gender" value="男"/>男
        <input type="radio" name="gender" value="女"/>女
        </td>
    </tr>
    <tr>
    	<th>年龄</th>
        <td><input type="text" name="age"/></td>
    </tr>
    <tr>
    	<th>电话</th>
        <td><input type="text" name="phone"/></td>
    </tr>
    <tr>
    	<th>邮箱</th>
        <td><input type="text" name="email"/></td>
    </tr>
    <tr>
    	<th>QQ</th>
        <td><input type="text" name="qq"/></td>
    </tr>
    <tr>
        <td colspan="2" align="center">
        <input type="submit" value="保存"/>&nbsp;
        <input type="reset" value="重置"/></td>
    </tr>
</table>
</form>
</body>
</html>
```

WebRoot/updateContact.jsp

```jsp
<%@ page language="java" import="java.util.*,gz.itcast.contactSys_web.entity.*" pageEncoding="utf-8"%>
<%@ taglib uri="http://java.sun.com/jsp/jstl/core" prefix="c" %>
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
<title>查询所有联系人</title>
<style type="text/css">
	table td{
		/*文字居中*/
		text-align:center;
	}
	
	/*合并表格的边框*/
	table{
		border-collapse:collapse;
	}
</style>
</head>

<body>
<center><h3>查询所有联系人(jsp版本)</h3></center>
<table align="center" border="1" width="700px">
	<tr>
    	<th>编号</th>
        <th>姓名</th>
        <th>性别</th>
        <th>年龄</th>
        <th>电话</th>
        <th>邮箱</th>
        <th>QQ</th>
        <th>操作</th>
    </tr>
    <c:forEach items="${contacts}" var="con" varStatus="varSta">
    <tr>
    	<td>${varSta.count }</td>
        <td>${con.name }</td>
        <td>${con.gender }</td>
        <td>${con.age }</td>
        <td>${con.phone }</td>
        <td>${con.email }</td>
        <td>${con.qq }</td>
        <td><a href="${pageContext.request.contextPath }/QueryContactServlet?id=${con.id}">修改</a>&nbsp;<a href="${pageContext.request.contextPath }/DeleteContactServlet?id=${con.id}">删除</a></td>
    </tr>
    </c:forEach>
    <tr>
    	<td colspan="8" align="center"><a href="${pageContext.request.contextPath }/addContact.jsp">[添加联系人]</a></td>
    </tr>
</table>
</body>
</html>
```



### 4. 功能测试

### 5. 性能测试

### 6. 部署上线

### 7. 维护阶段

------







## 过滤器-Filter 

### 基本概念

为什么需用到过滤器？
项目开发中，经常会涉及到重复代码的实现！
		注册 ---- Servlet 【1. 设置编码】 ----  JSP
		修改 ----Servlet 【1. 设置编码】 ---  JSP
其他
		如判断用户是否登陆，只有登陆才能有操作权限！
		涉及到重复判断： 获取session，取出session数据，判断是否为空,为空说明没有登陆，不能操作； 只有登陆后，才能操作！



如何解决：

1. 抽取重复代码，封装
2. 每个用到重复代码的地方，手动的调用！



过滤器，设计执行流程：

1. 用户访问服务器
2. 过滤器： 对Servlet请求进行拦截
3. 先进入过滤器， 过滤器处理
4. 过滤器处理完后， 在放行， 此时，请求到达Servlet/JSP
5. Servlet处理
6. Servlet处理完后，再回到过滤器, 最后在由tomcat服务器相应用户；



### 过滤器相关Api

```java
Javax.servlet.*;

|-- interface  Filter	过滤器核心接口
	Void  init(filterConfig);    初始化方法，在服务器启动时候执行
	Void  doFilter(request,response,filterChain);   过滤器拦截的业务处理方法
	Void destroy();   销毁过滤器实例时候调用

|-- interface  FilterConfig   获取初始化参数信息
	String	getInitParameter(java.lang.String name) 
	Enumeration	getInitParameterNames() 

|-- interface  FilterChain     过滤器链参数；一个个过滤器形成一个执行链；
	void doFilter(ServletRequest request, ServletResponse response)  ;  执行下一个过滤器或放行
```

开发步骤：

1. 写一个普通java类，实现Filter接口

```java
package cn.ixfosa.filter_hello;

import java.io.IOException;
import java.util.Enumeration;

import javax.servlet.Filter;
import javax.servlet.FilterChain;
import javax.servlet.FilterConfig;
import javax.servlet.ServletException;
import javax.servlet.ServletRequest;
import javax.servlet.ServletResponse;

/**
 * 过滤器，测试
 *
 */
public class HelloFilter implements Filter{
	
	// 创建实例
	public HelloFilter(){
		System.out.println("1. 创建过滤器实例");
	}

	@Override
	public void init(FilterConfig filterConfig) throws ServletException {
		System.out.println("2. 执行过滤器初始化方法");
		
		// 获取过滤器在web.xml中配置的初始化参数
		String encoding = filterConfig.getInitParameter("encoding");
		System.out.println(encoding);
		
		// 获取过滤器在web.xml中配置的初始化参数 的名称
		Enumeration<String> enums =  filterConfig.getInitParameterNames();
		while (enums.hasMoreElements()){
			// 获取所有参数名称：encoding、path
			String name = enums.nextElement();
			// 获取名称对应的值
			String value = filterConfig.getInitParameter(name);
			System.out.println(name + "\t" + value);
		}
	}

	// 过滤器业务处理方法： 在请求到达servlet之前先进入此方法处理公用的业务逻辑操作
	@Override
	public void doFilter(ServletRequest request, ServletResponse response,
			FilterChain chain) throws IOException, ServletException {
        
		System.out.println("3. 执行过滤器业务处理方法");
		// 放行 (去到Servlet)
		// 如果有下一个过滤器，进入下一个过滤器，否则就执行访问servlet
		chain.doFilter(request, response);
		
		System.out.println("5. Servlet处理完成，又回到过滤器");
	}

	@Override
	public void destroy() {
		System.out.println("6. 销毁过滤器实例");
	}
}
```



1. 配置过滤器

```xml
<!-- 过滤器配置 -->
<filter>
    <!-- 配置初始化参数 -->
    <init-param>
        <param-name>encoding</param-name>
        <param-value>UTF-8</param-value>
    </init-param>
    <init-param>
        <param-name>path</param-name>
        <param-value>c:/...</param-value>
    </init-param>

    <!-- 内部名称 -->
    <filter-name>hello_filter</filter-name>
    <!-- 过滤器类的全名 -->
    <filter-class>cn.ixfosa.filter_hello.HelloFilter</filter-class>
</filter>
<filter-mapping>
    <!-- filter内部名称 -->
    <filter-name>hello_filter</filter-name>
    <!-- 拦截所有资源 -->
    <url-pattern>/*</url-pattern>
</filter-mapping>
```



### 过滤器执行流程

![过滤器执行流程](images\过滤器执行流程.jpg)





### 对指定的请求拦截

| `/*` 表示拦截所有的请求                                      |
| ------------------------------------------------------------ |
| <filter-mapping>	<br />	<filter-name>hello_filter2</filter-name>	<br />	<url-pattern>/*</url-pattern>	<br /></filter-mapping> |

```xml
默认拦截的类型：(直接访问或者重定向)
	<dispatcher>REQUEST</dispatcher>

拦截转发：
	<dispatcher>FORWARD</dispatcher>

拦截包含的页面(RequestDispatcher.include(/page.jsp);    对page.jsp也执行拦截)
	<dispatcher>INCLUDE</dispatcher>

拦截声明式异常信息：
	<dispatcher>ERROR</dispatcher>
```

------



```java
package cn.ixfosa.filter_hello;

import java.io.IOException;

import javax.servlet.Filter;
import javax.servlet.FilterChain;
import javax.servlet.FilterConfig;
import javax.servlet.ServletException;
import javax.servlet.ServletRequest;
import javax.servlet.ServletResponse;

public class HelloFilter2 implements Filter{

	@Override
	public void doFilter(ServletRequest request, ServletResponse response,
			FilterChain chain) throws IOException, ServletException {
		System.out.println("第二个过滤器");
		// 放心
		chain.doFilter(request, response);
		System.out.println("第二个过滤器执行结束");
	}

	@Override
	public void init(FilterConfig filterConfig) throws ServletException {
		// TODO Auto-generated method stub
		
	}
	
	@Override
	public void destroy() {
		// TODO Auto-generated method stub
	}
}
```



```xml
<!-- 配置第二个过滤器 -->
<!-- 演示： 拦截指定的请求 -->
<filter>
    <filter-name>hello_filter2</filter-name>
    <filter-class>cn.ixfosa.filter_hello.HelloFilter2</filter-class>
</filter>

<filter-mapping>
    <filter-name>hello_filter2</filter-name>
    
    <!-- 1. 拦截所有
      <url-pattern>/*</url-pattern>
       -->

    <!-- 2. 拦截指定的jsp 
       <url-pattern>/index.jsp</url-pattern>
       <url-pattern>/list.jsp</url-pattern>
       -->
    
    <!-- 拦截所有的jsp
       <url-pattern>*.jsp</url-pattern>
        -->
    
    <!-- 3. 根据servlet的内部名称拦截
        <servlet-name>IndexServlet</servlet-name>
         -->
    
    <!-- 拦截指定的servlet 
        <url-pattern>/index</url-pattern>
        -->

    <!-- 4. 指定拦截指定的类型 -->
    <url-pattern>/*</url-pattern>
    <!-- 拦截直接访问的请求或者重定向的资源 -->
    <dispatcher>REQUEST</dispatcher>
    <!--<dispatcher>FORWARD</dispatcher>-->
</filter-mapping>
```



### 案例

#### 编码统一处理

```java
// 设置POST提交的请求的编码
request.setCharacterEncoding("UTF-8");

// 设置相应体的编码
response.setCharacterEncoding("UTF-8");

// 设置页面打开时候时候的编码格式、 设置相应体的编码
response.setContentType("text/html;charset=UTF-8");
```

几乎每一个Servlet都要涉及编码处理：处理请求数据中文问题！
【GET/POST】
每个servlet都要做这些操作，把公用的代码抽取-过滤器实现！

代码实现思路:

1. Login.jsp  登陆，输入“中文”
2. LoginServlet.java   直接处理登陆请求
3. EncodingFilter.java   过滤器处理请求数据编码：GET/POST

```jsp
<%@ page language="java" import="java.util.*" pageEncoding="UTF-8"%>

<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN">
<html>
  <head>
    <title>My JSP 'index.jsp' starting page</title>
	<meta http-equiv="pragma" content="no-cache">
	<meta http-equiv="cache-control" content="no-cache">
	<meta http-equiv="expires" content="0">    
  </head>
  
  <body>
  	<form name="frmLogin" action="${pageContext.request.contextPath }/login" method="post">
  	   用户名: <input type="text" name="userName"><br/>
  	  <input type="submit" value="POST提交" >
  	</form>
  	<hr/>
  	<form name="frmLogin" action="${pageContext.request.contextPath }/login" method="get">
  	   用户名: <input type="text" name="userName"><br/>
  	  <input type="submit" value="GET提交" >
  	</form>
  </body>
</html>
```

```java
package cn.ixfosa.loginFilter;

import java.io.IOException;
import java.io.PrintWriter;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

public class LoginServlet extends HttpServlet {

	public void doGet(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, IOException {

		// 获取请求数据 
		String name = request.getParameter("userName");
		System.out.println("用户：" + name);
	}
    
	public void doPost(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, IOException {
        
		this.doGet(request, response);
	}
}
```

```java
package cn.ixfosa.loginFilter;

import java.io.IOException;
import java.lang.reflect.InvocationHandler;
import java.lang.reflect.Method;
import java.lang.reflect.Proxy;

import javax.servlet.Filter;
import javax.servlet.FilterChain;
import javax.servlet.FilterConfig;
import javax.servlet.ServletException;
import javax.servlet.ServletRequest;
import javax.servlet.ServletResponse;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

/**
 * 编码处理统一写到这里(servlet中不需要再处理编码)
 *
 */
public class EncodingFilter implements Filter {

	// 过滤器业务处理方法：处理的公用的业务逻辑操作
	@Override
	public void doFilter(ServletRequest req, ServletResponse res,
			FilterChain chain) throws IOException, ServletException {
		
		// 转型
		final HttpServletRequest request = (HttpServletRequest) req;    
		HttpServletResponse response = (HttpServletResponse) res;
		
		// 一、处理公用业务
		request.setCharacterEncoding("UTF-8");					// POST提交有效
		response.setContentType("text/html;charset=UTF-8");
		
		/*
		 * 出现GET中文乱码，是因为在request.getParameter方法内部没有进行提交方式判断并处理。
		 * String name = request.getParameter("userName");
		 * 
		 * 解决：对指定接口的某一个方法进行功能扩展，可以使用代理!
		 *      对request对象(目标对象)，创建代理对象！
		 */
		HttpServletRequest proxy =  (HttpServletRequest) Proxy.newProxyInstance(
				request.getClass().getClassLoader(), 		// 指定当前使用的类加载器
				new Class[]{HttpServletRequest.class}, 		// 对目标对象实现的接口类型
				new InvocationHandler() {					// 事件处理器
					@Override
					public Object invoke(Object proxy, Method method, Object[] args)
							throws Throwable {
                        
						// 定义方法返回值
						Object returnValue = null;
                        
						// 获取方法名
						String methodName = method.getName();
                        
						// 判断：对getParameter方法进行GET提交中文处理
						if ("getParameter".equals(methodName)) {
							
							// 获取请求数据值【 <input type="text" name="userName">】
							String value = request.getParameter(args[0].toString());	// 调用目标对象的方法
							
							// 获取提交方式
							String methodSubmit = request.getMethod(); // 直接调用目标对象的方法
							
							// 判断如果是GET提交，需要对数据进行处理  (POST提交已经处理过了)
							if ("GET".equals(methodSubmit)) {
								if (value != null && !"".equals(value.trim())){
									// 处理GET中文
									value = new String(value.getBytes("ISO8859-1"),"UTF-8");
								}
							} 
							return value;
						}else {
							// 执行request对象的其他方法
							returnValue = method.invoke(request, args);
						}
						
						return returnValue;
					}
				});
		
		// 二、放行 (执行下一个过滤器或者servlet)
		chain.doFilter(proxy, response);		// 传入代理对象
	}

    
	@Override
	public void init(FilterConfig filterConfig) throws ServletException {
		
	}
	
	@Override
	public void destroy() {
		
	}
}
```

```xml
<!--1.  编码处理过滤器配置 -->
<filter>
    <filter-name>encoding</filter-name>
    <filter-class>cn.ixfosa.loginFilter.EncodingFilter</filter-class>
</filter>
<filter-mapping>
    <filter-name>encoding</filter-name>
    <url-pattern>/*</url-pattern>
</filter-mapping>

<servlet-mapping>
    <servlet-name>LoginServlet</servlet-name>
    <url-pattern>/login</url-pattern>
</servlet-mapping>
```



#### 无效数据过滤

模拟：论坛过滤敏感词汇！

实现思路：

1. Dis.jsp    讨论区页面
2. DisServlet.java    处理提交
   - 获取请求参数
   - 保存到request域
   - 跳转dis.jsp  【从request取数据显示(处理后)】
3. DataFilter.java   过滤器
   * 编码
   * 无效数据处理
           即： 在上一个案例基础上，再添加无效数据过滤的相关代码!




JSP引入ckeditor组件：客户端组件，便于用户输入内容！

```jsp
<%@ page language="java" import="java.util.*" pageEncoding="UTF-8"%>

<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN">
<html>
  <head>
  
    <title>My JSP 'index.jsp' starting page</title>
	<meta http-equiv="pragma" content="no-cache">
	<meta http-equiv="cache-control" content="no-cache">
	<meta http-equiv="expires" content="0">   
	<!-- 引入ckeditor组件(给用户输入提供方便) --> 
	<script src="${pageContext.request.contextPath }/ckeditor/ckeditor.js"></script>
	<link rel="stylesheet" href="${pageContext.request.contextPath }/ckeditor/samples/sample.css">
	
  </head>
  
   <body>
  	${requestScope.content }
  
  	<form name="frmDis" action="${pageContext.request.contextPath }/dis" method="post">	
  	  发表评论: <textarea class="ckeditor" rows="6" cols="30" name="content"></textarea>
  	  
  	  <br/>
  	  <input type="submit" value="评论" >
  	</form>
  </body>
   
</html>
```



```java
package cn.ixfosa.filter_data;

import java.io.IOException;
import java.io.PrintWriter;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

public class DisServlet extends HttpServlet {

	public void doGet(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, IOException {

		// 获取请求数据 
		String content = request.getParameter("content");
		// 保存到request
		request.setAttribute("content", "Content：" + content);
		// 转发
		request.getRequestDispatcher("/dis.jsp").forward(request, response);
		
	}

	public void doPost(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, IOException {
		this.doGet(request, response);
	}

}
```

```java
package cn.ixfosa.filter_data;

import java.io.IOException;
import java.lang.reflect.InvocationHandler;
import java.lang.reflect.Method;
import java.lang.reflect.Proxy;
import java.util.ArrayList;
import java.util.List;

import javax.servlet.Filter;
import javax.servlet.FilterChain;
import javax.servlet.FilterConfig;
import javax.servlet.ServletException;
import javax.servlet.ServletRequest;
import javax.servlet.ServletResponse;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

/**
 * 无效数据过滤
 *
 */
public class DateFilter implements Filter {
	
	// 初始化无效数据
	private List<String> dirtyData;
    
	@Override
	public void init(FilterConfig filterConfig) throws ServletException {
		// 模拟几个数据
		dirtyData = new ArrayList<String>();
		dirtyData.add("tmd");
		dirtyData.add("sb");
	}

	@Override
	public void doFilter(ServletRequest req, ServletResponse res,
			FilterChain chain) throws IOException, ServletException {
		
		// 转型
		final HttpServletRequest request = (HttpServletRequest) req;    
		HttpServletResponse response = (HttpServletResponse) res;
		
		// 一、处理公用业务
		request.setCharacterEncoding("UTF-8");					// POST提交有效
		response.setContentType("text/html;charset=UTF-8");
		
		HttpServletRequest proxy =  (HttpServletRequest) Proxy.newProxyInstance(
				request.getClass().getClassLoader(), 		// 指定当前使用的类加载器
				new Class[]{HttpServletRequest.class}, 		// 对目标对象实现的接口类型
				new InvocationHandler() {					// 事件处理器
					@Override
					public Object invoke(Object proxy, Method method, Object[] args)
							throws Throwable {
						// 定义方法返回值
						Object returnValue = null;
						// 获取方法名
						String methodName = method.getName();
						// 判断：对getParameter方法进行GET提交中文处理
						if ("getParameter".equals(methodName)) {
							
							// 获取请求数据值
							String value = request.getParameter(args[0].toString());	// 调用目标对象的方法
							
							// 获取提交方式
							String methodSubmit = request.getMethod(); // 直接调用目标对象的方法
							
							// 判断如果是GET提交，需要对数据进行处理  (POST提交已经处理过了)
							if ("GET".equals(methodSubmit)) {
								if (value != null && !"".equals(value.trim())){
									// 处理GET中文
									value = new String(value.getBytes("ISO8859-1"),"UTF-8");
								}
							} 
							
							// 中文数据已经处理完： 下面进行无效数据过滤   
							//【如何value中出现dirtyData中数据，用****替换】  
							for (String data : dirtyData) {
								// 判断当前输入数据(value), 是否包含无效数据
								if (value.contains(data)){
									value = value.replace(data, "*****");
								}
							}
							// 处理完编码、无效数据后的正确数据
							return value;
						}
						else {
							// 执行request对象的其他方法
							returnValue = method.invoke(request, args);
						}
						
						return returnValue;
					}
				});
		
		// 二、放行 (执行下一个过滤器或者servlet)
		chain.doFilter(proxy, response);		// 传入代理对象
	}

	@Override
	public void destroy() {
		
	}
}
```

```xml
<!--无效数据过滤器配置 -->
<filter>
    <filter-name>dataFilter</filter-name>
    <filter-class>cn.ixfosa.filter_data.DateFilter</filter-class>
</filter>
<filter-mapping>
    <filter-name>dataFilter</filter-name>
    <url-pattern>/*</url-pattern>
</filter-mapping>


<servlet>
	<servlet-name>DisServlet</servlet-name>
	<servlet-class>cn.ixfosa.filter_data.DisServlet</servlet-class>
</servlet>
<servlet-mapping>
    <servlet-name>DisServlet</servlet-name>
    <url-pattern>/dis</url-pattern>
 </servlet-mapping>
```



#### 登陆权限判断

登陆， 提交到登陆Servlet处理其业务！
		登陆成功, 跳转到首页，显示欢迎信息 + 列表信息
		登陆失败，跳转到登陆！

要求：
		只有登陆后，才可以访问首页, 显示列表
		如果没有登陆，直接访问首页列表，要跳转到登陆!

实现思路：

1. Login.jsp   登陆页面
2. List.jsp     列表显示
3. LoginServlet.java   登陆处理servlet
4. IndexServlet.java   首页列表查询Servlet
5. LoginFilter.java     登陆验证过滤器



http://localhost:8080/emp_sys/login.jsp   可以直接访问
http://localhost:8080/emp_sys/login      可以直接访问
http://localhost:8080/emp_sys/index   不能直接访问
http://localhost:8080/emp_sys/list.jsp   不能直接访问

##### 引入jar

- mysql-connector-java-5.1.12-bin.jar
- commons-dbutils-1.6.jar
- c3p0-0.9.1.2.jar



##### utils

src/c3p0-config.xml

```xml
<c3p0-config>
  <default-config>
     <property name="driverClass">com.mysql.jdbc.Driver</property> 
     <property name="jdbcUrl">jdbc:mysql:///jdbc_demo</property> 
     <property name="user">root</property> 
     <property name="password">root</property> 
     <property name="initialPoolSize">5</property> 
     <property name="maxPoolSize">10</property> 

  </default-config>


  <named-config name="oracleConfig">
    <property name="driverClass">com.mysql.jdbc.Driver</property> 
     <property name="jdbcUrl">jdbc:mysql:///day17</property> 
     <property name="user">root</property> 
     <property name="password">root</property> 
     <property name="initialPoolSize">5</property> 
     <property name="maxPoolSize">10</property> 
   </named-config>

</c3p0-config>
```

```java
package cn.ixfosa.utils;

import javax.sql.DataSource;

import org.apache.commons.dbutils.QueryRunner;

import com.mchange.v2.c3p0.ComboPooledDataSource;

/**
 * 工具类
 * 1. 初始化C3P0连接池
 * 2. 创建DbUtils核心工具类对象
 *
 */
public class JdbcUtils {

	/**
	 *  1. 初始化C3P0连接池
	 */
	private static  DataSource dataSource;
	static {
		dataSource = new ComboPooledDataSource();
	}
	
	/**
	 * 2. 创建DbUtils核心工具类对象
	 */
	public static QueryRunner getQueryRuner(){
		// 创建QueryRunner对象，传入连接池对象
		// 在创建QueryRunner对象的时候，如果传入了数据源对象；
		// 那么在使用QueryRunner对象方法的时候，就不需要传入连接对象；
		// 会自动从数据源中获取连接(不用关闭连接)
		return new QueryRunner(dataSource);
	}
}
```



##### entity

```java
package cn.ixfosa.entity;

/**
 * 1. 管理员实体类开发
 *
 */
public class Admin {

	private int id; 
	private String userName;
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
}
```

```java
package cn.ixfosa.entity;

/**
 * 1. 员工
 *
 */
public class Employee {

	private int empId;
	private String empName;
	private int dept_id;
	
	public int getEmpId() {
		return empId;
	}
	public void setEmpId(int empId) {
		this.empId = empId;
	}
	public String getEmpName() {
		return empName;
	}
	public void setEmpName(String empName) {
		this.empName = empName;
	}
	public int getDept_id() {
		return dept_id;
	}
	public void setDept_id(int deptId) {
		dept_id = deptId;
	}
}
```



##### Dao

- interface

```java
package cn.ixfosa.dao;

import cn.ixfosa.entity.Admin;

/**
 * 2. 管理员数据访问层接口设计
 *
 */
public interface IAdminDao {

	/**
	 * 根据用户名密码查询
	 * @param admin
	 * @return
	 */
	Admin findByNameAndPwd(Admin admin);
}

```

```java
package cn.ixfosa.dao;

import java.util.List;

import cn.ixfosa.entity.Employee;

/**
 * 2. 员工数据访问层接口设计
 *
 */
public interface IEmployeeDao {

	/**
	 * 查询所有的员工
	 * @return
	 */
	List<Employee> getAll();
}
```

- impl

```java
package cn.ixfosa.dao.impl;

import java.sql.SQLException;

import org.apache.commons.dbutils.handlers.BeanHandler;

import cn.ixfosa.dao.IAdminDao;
import cn.ixfosa.entity.Admin;
import cn.ixfosa.utils.JdbcUtils;

public class AdminDao implements IAdminDao {

	@Override
	public Admin findByNameAndPwd(Admin admin) {
		try {
			String sql = "select * from admin where userName=? and pwd=?";
			return JdbcUtils.getQueryRuner()//
					.query(sql, 
							new BeanHandler<Admin>(Admin.class),
							admin.getUserName(),
							admin.getPwd());
		} catch (SQLException e) {
			throw new RuntimeException(e);
		}
	}
}
```

```java
package cn.itcast.dao.impl;

import java.sql.SQLException;
import java.util.List;

import org.apache.commons.dbutils.handlers.BeanListHandler;

import cn.ixfosa.dao.IEmployeeDao;
import cn.ixfosa.entity.Employee;
import cn.ixfosa.utils.JdbcUtils;

public class EmployeeDao implements IEmployeeDao {

	@Override
	public List<Employee> getAll() {
		String sql = "select * from employee";
		try {
			return JdbcUtils.getQueryRuner()//
				.query(sql, new BeanListHandler<Employee>(Employee.class));
		} catch (SQLException e) {
			throw new RuntimeException(e);
		}
	}

}
```



##### Servcie

- interface

```java
package cn.ixfosa.service;

import cn.ixfosa.entity.Admin;

/**
 * 3. 管理员业务逻辑层
 *
 */
public interface IAdminService {

	/**
	 * 根据用户名密码查询
	 * @param admin
	 * @return
	 */
	Admin findByNameAndPwd(Admin admin);
}
```

```java
package cn.ixfosa.dao;

import java.util.List;

import cn.ixfosa.entity.Employee;

/**
 * 2. 员工数据访问层接口设计
 *
 */
public interface IEmployeeDao {

	/**
	 * 查询所有的员工
	 * @return
	 */
	List<Employee> getAll();
}
```



- impl

```java
package cn.ixfosa.dao.impl;

import java.sql.SQLException;

import org.apache.commons.dbutils.handlers.BeanHandler;

import cn.ixfosa.dao.IAdminDao;
import cn.ixfosa.entity.Admin;
import cn.ixfosa.utils.JdbcUtils;

public class AdminDao implements IAdminDao {

	@Override
	public Admin findByNameAndPwd(Admin admin) {
		try {
			String sql = "select * from admin where userName=? and pwd=?";
			return JdbcUtils.getQueryRuner()//
					.query(sql, 
							new BeanHandler<Admin>(Admin.class),
							admin.getUserName(),
							admin.getPwd());
		} catch (SQLException e) {
			throw new RuntimeException(e);
		}
	}

}
```

```java
package cn.ixfosa.dao.impl;

import java.sql.SQLException;
import java.util.List;

import org.apache.commons.dbutils.handlers.BeanListHandler;

import cn.ixfosa.dao.IEmployeeDao;
import cn.ixfosa.entity.Employee;
import cn.ixfosa.utils.JdbcUtils;

public class EmployeeDao implements IEmployeeDao {

   @Override
   public List<Employee> getAll() {
      String sql = "select * from employee";
      try {
         return JdbcUtils.getQueryRuner()//
            .query(sql, new BeanListHandler<Employee>(Employee.class));
      } catch (SQLException e) {
         throw new RuntimeException(e);
      }
   }
}
```



##### servlet

```java
package cn.ixfosa.servlet;

import java.io.IOException;
import java.io.PrintWriter;
import java.util.List;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import cn.ixfosa.entity.Admin;
import cn.ixfosa.entity.Employee;
import cn.ixfosa.service.IAdminService;
import cn.ixfosa.service.IEmployeeService;
import cn.ixfosa.service.impl.AdminService;
import cn.ixfosa.service.impl.EmployeeService;

public class IndexServlet extends HttpServlet {

	// Service实例
	private IEmployeeService employeeService = new EmployeeService();
	// 跳转资源
	private String uri;

	public void doGet(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, IOException {
		
		try {
			// 调用service查询所有
			List<Employee> list = employeeService.getAll();
			request.setAttribute("listEmp", list);
			// 进入首页jsp
			uri = "/list.jsp";
		} catch (Exception e) {
			e.printStackTrace();
			uri = "/error/error.jsp";
		}
		// 转发
		request.getRequestDispatcher(uri).forward(request, response);

	}

	public void doPost(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, IOException {
		this.doGet(request, response);
	}
}
```

```java
package cn.itcast.servlet;

import java.io.IOException;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import cn.ixfosa.entity.Admin;
import cn.ixfosa.service.IAdminService;
import cn.ixfosa.service.impl.AdminService;

/**
 * 处理登陆请求
 *
 */
public class LoginServlet extends HttpServlet {
	// Service实例
	private IAdminService adminService = new AdminService();
	// 跳转资源
	private String uri;

	public void doGet(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, IOException {

		request.setCharacterEncoding("UTF-8");
		
		//1. 获取参数
		String userName = request.getParameter("userName");
		String pwd = request.getParameter("pwd");
		// 封装
		Admin admin = new Admin();
		admin.setUserName(userName);
		admin.setPwd(pwd);
		
		try {
			//2. 调用service处理业务
			Admin loginInfo = adminService.findByNameAndPwd(admin);
			// 判断:
			if (loginInfo == null){
				// 登陆失败
				uri = "/login.jsp";
			} else {
				// 登陆成功
				// 先，保存数据到session
				request.getSession().setAttribute("loginInfo", loginInfo);
				// 再，跳转到首页显示servlet(/index)
				uri = "/index";
			}
		} catch (Exception e) {
			// 测试
			e.printStackTrace();
			// 错误
			uri = "/error/error.jsp";
		}
		
		//3. 跳转
		request.getRequestDispatcher(uri).forward(request, response);
		
	}

	public void doPost(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, IOException {

		this.doGet(request, response);
	}
}
```



##### filter

```java
package cn.ixfosa.filter;

import java.io.IOException;

import javax.servlet.Filter;
import javax.servlet.FilterChain;
import javax.servlet.FilterConfig;
import javax.servlet.ServletException;
import javax.servlet.ServletRequest;
import javax.servlet.ServletResponse;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import javax.servlet.http.HttpSession;

/**
 * 登陆验证过滤器
 * 
 *  http://localhost:8080/emp_sys/login.jsp   可以直接访问
	http://localhost:8080/emp_sys/login      可以直接访问
	http://localhost:8080/emp_sys/index   不能直接访问
	http://localhost:8080/emp_sys/list.jsp   不能直接访问
 *
 */
public class LoginFilter implements Filter {
	
	private String uri;

	/**
	 * 分析：
	 * 
		1. 先指定放行的资源，哪些资源不需要拦截：
		      login.jsp   +    /login  (request对象可以获取)
		2. 获取session，从session中获取登陆用户
		3. 判断是否为空：
		      为空， 说明没有登陆， 跳转到登陆
		       不为空， 已经登陆，放行！
	 */
	@Override
	public void doFilter(ServletRequest req, ServletResponse res,
			FilterChain chain) throws IOException, ServletException {
		
		//0. 转换
		HttpServletRequest request = (HttpServletRequest) req;
		HttpServletResponse response = (HttpServletResponse) res;
		
		//1. 获取请求资源，截取  
		String uri = request.getRequestURI();   // /emp_sys/login.jsp
		// 截取 【login.jsp或login】
		String requestPath = uri.substring(uri.lastIndexOf("/") + 1, uri.length());  
		
		//2. 判断： 先放行一些资源：/login.jsp、/login
		if ("login".equals(requestPath) || "login.jsp".equals(requestPath)) {
			// 放行
			chain.doFilter(request, response);
		}
		else {
			//3. 对其他资源进行拦截
			//3.1 先获取Session、获取session中的登陆用户(loginInfo)
			HttpSession session = request.getSession(false);
			// 判断
			if (session != null) {
				
				Object obj = session.getAttribute("loginInfo");
				
				//3.2如果获取的内容不为空，说明已经登陆，放行
				if (obj != null) {
					// 放行
					chain.doFilter(request, response);
				} else {
					//3.3如果获取的内容为空，说明没有登陆； 跳转到登陆
					uri = "/login.jsp";
				}
				
			} else {
				// 肯定没有登陆
				uri = "/login.jsp";
			}
			request.getRequestDispatcher(uri).forward(request, response);
		}
	}

	@Override
	public void init(FilterConfig filterConfig) throws ServletException {
	}

	@Override
	public void destroy() {
	}
}
```



##### Jsp

```xml
<%@ page language="java" import="java.util.*" pageEncoding="UTF-8"%>

<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN">
<html>
  <head>
    <title>My JSP 'index.jsp' starting page</title>
	<meta http-equiv="pragma" content="no-cache">
	<meta http-equiv="cache-control" content="no-cache">
	<meta http-equiv="expires" content="0">    
  </head>
  
  <body>
  <form name="frmLogin" action="${pageContext.request.contextPath }/login" method="post">
  	<table align="center" border="1">
  		<tr>
  			<td>用户名</td>
  			<td>
  				<input type="text" name="userName">
  			</td>
  		</tr>
  		<tr>
  			<td>密码</td>
  			<td>
  				<input type="password" name="pwd">
  			</td>
  		</tr>
  		<tr>
  			<td>
  				<input type="submit" value="亲，点我登陆！">
  			</td>
  		</tr>
  	</table>
  </form>
  </body>
</html>

```

```jsp
<%@ page language="java" import="java.util.*" pageEncoding="UTF-8"%>
<!-- 引入jstl核心标签库 -->
<%@taglib uri="http://java.sun.com/jsp/jstl/core" prefix="c" %>
<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN">
<html>
  <head>
    <title>My JSP 'index.jsp' starting page</title>
	<meta http-equiv="pragma" content="no-cache">
	<meta http-equiv="cache-control" content="no-cache">
	<meta http-equiv="expires" content="0">    
  </head>
  
  <body>
  	<h1>欢迎你，${sessionScope.loginInfo.userName }</h1>
  	<!-- 列表展示数据 -->
  	<table align="center" border="1" width="80%" cellpadding="3" cellspacing="0">
  		<tr>
  			<td>序号</td>
  			<td>编号</td>
  			<td>员工名称</td>
  		</tr>
  		<c:if test="${not empty requestScope.listEmp}">
	  		<c:forEach var="emp" items="${requestScope.listEmp}" varStatus="vs">
		  		<tr>
		  			<td>${vs.count }</td>
		  			<td>${emp.empId }</td>
		  			<td>${emp.empName }</td>
		  		</tr>
	  		</c:forEach>
  		</c:if>
  	</table>
  </body>
</html>
```

##### web.xml

```xml
<?xml version="1.0" encoding="UTF-8"?>
<web-app version="2.5" xmlns="http://java.sun.com/xml/ns/javaee"
	xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
	xsi:schemaLocation="http://java.sun.com/xml/ns/javaee 
	http://java.sun.com/xml/ns/javaee/web-app_2_5.xsd">

	<!-- 登陆验证过滤器 -->
	<filter>
		<filter-name>loginFilter</filter-name>
		<filter-class>cn.ixfosa.filter.LoginFilter</filter-class>
	</filter>
	<filter-mapping>
		<filter-name>loginFilter</filter-name>
		<url-pattern>/*</url-pattern>
	</filter-mapping>
	
	<servlet>
		<servlet-name>LoginServlet</servlet-name>
		<servlet-class>cn.ixfosa.servlet.LoginServlet</servlet-class>
	</servlet>
	<servlet>
		<servlet-name>IndexServlet</servlet-name>
		<servlet-class>cn.ixfosa.servlet.IndexServlet</servlet-class>
	</servlet>


	<servlet-mapping>
		<servlet-name>LoginServlet</servlet-name>
		<url-pattern>/login</url-pattern>
	</servlet-mapping>
	<servlet-mapping>
		<servlet-name>IndexServlet</servlet-name>
		<url-pattern>/index</url-pattern>
	</servlet-mapping>
</web-app>
```

------



## 监听器-listener

### 概述

监听器： 主要是用来监听特定对象的创建或销毁、属性的变化的！
		 是一个实现特定接口的普通java类！



Servlet中哪些对象需要监听？
		分别对应的是`request`监听器、`session`相关监听器、`servletContext`监听器



监听器接口：
一、监听对象创建/销毁的监听器接口

```java
Interface ServletRequestListener     监听request对象的创建或销毁
Interface HttpSessionListener        监听session对象的创建或销毁
Interface ServletContextListener     监听servletContext对象的创建或销毁
```

二、监听对象属性的变化

```java
Interface ServletRequestAttributeListener 监听request对象属性变化: 添加、移除、修改
Interface HttpSessionAttributeListener    监听session对象属性变化: 添加、移除、修改
Interface ServletContextAttributeListener  监听servletContext对象属性变化
```

三、session相关监听器

```java
Interface HttpSessionBindingListener      监听对象绑定到session上的事件	
Interface HttpSessionActivationListener   监听session序列化及反序列化的事件
```

404(路径写错)
500(服务器错误，调试)

### 生命周期监听器

声明周期监听器： 监听对象的创建、销毁的过程！
监听器开发步骤：

1. 写一个普通java类，实现相关接口；
2. 配置(web.xml)

#### ServletRequestListener

监听request对象的创建或销毁

```java
package cn.ixfosa.life;

import javax.servlet.ServletRequestEvent;
import javax.servlet.ServletRequestListener;

/**
 *  监听request对象的创建或销毁
 *
 */
public class MyRequestListener implements ServletRequestListener{

	// 对象销毁
	@Override
	public void requestDestroyed(ServletRequestEvent sre) {
		// 获取request中存放的数据
		Object obj = sre.getServletRequest().getAttribute("cn");
		System.out.println(obj);
		System.out.println("MyRequestListener.requestDestroyed()");
	}

	// 对象创建
	@Override
	public void requestInitialized(ServletRequestEvent sre) {
		System.out.println("MyRequestListener.requestInitialized()");
	}
}
```

```xml
<!-- 监听request对象创建、销毁 -->
<listener>
	<listener-class>cn.ixfosa.life.MyRequestListener</listener-class>
</listener>
```



#### HttpSessionListener

监听session对象的创建或销毁

```java
package cn.itcast.life;

import javax.servlet.ServletContextEvent;
import javax.servlet.ServletContextListener;

/**
 * 监听ServletContext对象创建或销毁
 *
 */
public class MyServletContextListener implements ServletContextListener{

	@Override
	public void contextDestroyed(ServletContextEvent sce) {
		System.out.println("MyServletContextListener.contextDestroyed()");
	}

	@Override
	public void contextInitialized(ServletContextEvent sce) {
		System.out.println("1..........MyServletContextListener.contextInitialized()");
	}

}
```

```xml
<!-- session的最大活跃时间 -->
<session-config>
    <session-timeout>60</session-timeout>
</session-config>

<!-- 监听session对象创建、销毁 -->
<listener>
    <listener-class>cn.ixfosa.life.MySessionListener</listener-class>
</listener>
```



#### ServletContextListener

监听servletContext对象的创建或销毁

```java
package cn.package cn.itcast.a_life;

import javax.servlet.http.HttpSessionEvent;
import javax.servlet.http.HttpSessionListener;

/**
 * 监听Session对象创建、销毁
 *
 */
public class MySessionListener implements HttpSessionListener{

	// 创建
	@Override
	public void sessionCreated(HttpSessionEvent se) {
		System.out.println("MySessionListener.sessionCreated()");
	}

	// 销毁
	@Override
	public void sessionDestroyed(HttpSessionEvent se) {
		System.out.println("MySessionListener.sessionDestroyed()");
	}
}
```

```xml
<!-- 监听servletContext对象创建、销毁 -->
<listener>
    <listener-class>cn.ixfosa.life.MyServletContextListener</listener-class>
</listener>
```

### 属性监听器

监听:request/session/servletContext对象属性的变化！

#### ServletRequestAttributeListener



#### HttpSessionAttributeListener

```java
package cn.ixfosa.attr;

import javax.servlet.http.HttpSession;
import javax.servlet.http.HttpSessionAttributeListener;
import javax.servlet.http.HttpSessionBindingEvent;

/**
 * 监听session对象属性的变化
 *
 */
public class MySessionAttrListener implements HttpSessionAttributeListener {

	// 属性添加
	@Override
	public void attributeAdded(HttpSessionBindingEvent se) {
		// 先获取session对象
		HttpSession session = se.getSession();
		// 获取添加的属性
		Object obj = session.getAttribute("userName");
		// 测试
		System.out.println("添加的属性：" + obj);
	}

	// 属性移除
	@Override
	public void attributeRemoved(HttpSessionBindingEvent se) {
		System.out.println("属性移除");
	}

	// 属性被替换
	@Override
	public void attributeReplaced(HttpSessionBindingEvent se) {
		// 获取sesison对象
		HttpSession session = se.getSession();
		
		// 获取替换前的值
		Object old = se.getValue();
		System.out.println("原来的值：" + old);
		
		// 获取新值
		Object obj_new = session.getAttribute("userName");
		System.out.println("新值：" + obj_new);
		
	}
}
```



```xml
<!-- 属性监听器 -->
<listener>
    <listener-class>cn.ixfosa.attr.MySessionAttrListener</listener-class>
</listener>
```

```jsp
<%@ page language="java" import="java.util.*" pageEncoding="UTF-8"%>
<%@page import="cn.ixfosa.session.Admin"%>

<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN">
<html>
  <head>
    
    <title>My JSP 'index.jsp' starting page</title>
	<meta http-equiv="pragma" content="no-cache">
	<meta http-equiv="cache-control" content="no-cache">
	<meta http-equiv="expires" content="0">    
  </head>
  
  <body>
  	test!
  	<%
  	
  		request.setAttribute("cn","China");
  		session.invalidate();  //销毁session
  		
  		session.setAttribute("userName","Jack");
  		session.removeAttribute("userName");
  		session.setAttribute("userName","Jack_new");
  		
  	%>
  </body>
</html>
```



#### ServletContextAttributeListener



### 其他监听器： session相关监听器

#### HttpSessionBindingListener 

监听对象绑定/解除绑定到sesison上的事件!

步骤：
		对象实现接口； 再把对象绑定/解除绑定到session上就会触发监听代码。



这个session监听器，和上面的声明周期、属性监听器区别？
		不用再web.xml配置
		因为监听的对象是自己创建的对象，不是服务器对象！

```java
package cn.ixfosa.session;

import javax.servlet.http.HttpSessionBindingEvent;
import javax.servlet.http.HttpSessionBindingListener;

/**
 * 监听此对象绑定到session上的过程，需要实现session特定接口
 *
 */
public class Admin implements HttpSessionBindingListener {

	private int id;
	private String name;
	
	public Admin() {
		super();
	}
	public Admin(int id, String name) {
		super();
		this.id = id;
		this.name = name;
	}
	
	
	// 构造函数
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
	
	// 对象放入session
	@Override
	public void valueBound(HttpSessionBindingEvent event) {
		System.out.println("Admin对象已经放入session");
	}
    
	// 对象从session中移除
	@Override
	public void valueUnbound(HttpSessionBindingEvent event) {
		System.out.println("Admin对象从session中移除！");
	}
}
```

```jsp
<%@ page language="java" import="java.util.*" pageEncoding="UTF-8"%>
<%@page import="cn.ixfosa.session.Admin"%>

<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN">
<html>
  <head>
    
    <title>My JSP 'index.jsp' starting page</title>
	<meta http-equiv="pragma" content="no-cache">
	<meta http-equiv="cache-control" content="no-cache">
	<meta http-equiv="expires" content="0">    
  </head>
  
  <body>
  	test!
  	<%  		
  		session.setAttribute("userInfo",new Admin());
  		session.removeAttribute("userInfo");
  	%>
  </body>
</html>
```



### 案例

需求：做一个在线列表提醒的功能!
      用户-- 登陆
	             显示登陆信息，列表展示。(list.jsp)
				 显示在线用户列表        (list.jsp)
				 列表点击进入“在线列表页面”   onlineuser.jsp

实现：

1. 先增加退出功能；  再把session活跃时间1min;
2. 写监听器，监听servletContext对象的创建： 初始化集合(onlineuserlist)
3. 登	陆功能： 用户登陆时候，把数据保存到servletContext中
4. List.jsp  增加超链接， 点击时候提交直接跳转到online.jsp
5. 写监听器： 监听session销毁，把当前登陆用户从onlineuserlist移除！

```java
package cn.ixfosa.listener;

import java.util.ArrayList;
import java.util.List;

import javax.servlet.ServletContext;
import javax.servlet.ServletContextEvent;
import javax.servlet.ServletContextListener;

import cn.ixfosa.entity.Admin;

/**
 * 初始化在线列表集合监听器
 * 
 */
public class OnlineAdminListener implements ServletContextListener {

	//1. ServletContext对象创建
	@Override
	public void contextInitialized(ServletContextEvent sce) {
		// 创建集合：存放在线用户
		// 每次当用户登陆后，就往这个集合中添加添加当前登陆者
		List<Admin> onlineList = new ArrayList<Admin>();
		// 放入ServletContext中
		sce.getServletContext().setAttribute("onlineList", onlineList);
	}

	//2. ServletContext对象销毁
	@Override
	public void contextDestroyed(ServletContextEvent sce) {
		// 获取ServletContext
		ServletContext sc = sce.getServletContext();
		// 获取在线列表
		Object obj = sc.getAttribute("onlineList");
		// 移除在线列表集合
		if (obj != null) {
			sc.removeAttribute("onlineList");
		}
	}
}
```

```java
package cn.ixfosa.listener;

import java.util.List;

import javax.servlet.ServletContext;
import javax.servlet.http.HttpSession;
import javax.servlet.http.HttpSessionEvent;
import javax.servlet.http.HttpSessionListener;

import cn.ixfosa.entity.Admin;

/**
 * 监听Session销毁的动作：
 *  当服务器销毁session的时候，从在线列表集合中移除当亲的登陆用户
 *
 */
public class SessionListener implements HttpSessionListener{

	@Override
	public void sessionDestroyed(HttpSessionEvent se) {
		//1. 获取Session对象、ServletContext对象
		HttpSession session = se.getSession();
		ServletContext sc = session.getServletContext();
		
		//2. 获取Session中存储的当前登陆用户
		Object obj = session.getAttribute("loginInfo");//?
		
		//3. 获取ServletContext中存储的在线用户列表集合
		List<Admin> list = (List<Admin>) sc.getAttribute("onlineList");
		// 先判断
		if (obj != null){
			//4. 把“当前登陆用户”从在线列表集合中移除
			list.remove(obj);
		}
	}
	
	@Override
	public void sessionCreated(HttpSessionEvent se) {
	}
}
```

```xml
<!-- 监听器： 监听ServletContext对象 -->
<listener>
    <listener-class>cn.ixfosa.listener.OnlineAdminListener</listener-class>
</listener>

<!-- 监听器：监听session对象 -->
<listener>
    <listener-class>cn.ixfosa.listener.SessionListener</listener-class>
</listener>
```



```jsp
<!-- 在线用户 -->
<table align="center" border="1" width="80%" cellpadding="3" cellspacing="0">
    <tr>
        <td colspan="2" align="center"><h3>在线列表展示：</h3></td>
    </tr>
    <tr>
        <td>编号</td>
        <td>员工名称</td>
    </tr>
    <c:if test="${not empty applicationScope.onlineList}">
        <c:forEach var="admin" items="${applicationScope.onlineList}">
            <tr>
                <td>${admin.id }</td>
                <td>${admin.userName }</td>
            </tr>
        </c:forEach>
    </c:if>
</table>
```



## 国际化

国际化又简称为 i18n：internationalization

### Locale  本地化

Java提供了一个本地化的对象！封装当前语言、国家、环境等特征！

```java
package cn.ixfosa.locale;

import java.util.Locale;

import org.junit.Test;

public class App {

	@Test
	//1. 本地化对象:Locale
	// 封装语言、国家信息的对象，有java.util提供
	public void testLocale() throws Exception {
		// 模拟中国语言等环境
		//Locale locale = Locale.CHINA;
		Locale locale = Locale.getDefault();			// 当前系统默认的语言环境
		System.out.println(locale.getCountry());   		// CN  国家的简称
		System.out.println(locale.getDisplayCountry()); // 国家名称
		System.out.println(locale.getLanguage());		// zh 语言简称
		
		// 模拟美国国家
		Locale l_us = Locale.US;
		System.out.println(l_us.getCountry());
		System.out.println(l_us.getDisplayCountry());
	}
}
```

### 国际化



#### 资源文件

src/cn/ixfosa/i18n/msg.properties

```properties
hello=\u4F60\u597D
username=\u7528\u6237\u540D
pwd=\u5BC6\u7801
title=\u767B\u9646\u9875\u9762
submit=\ \u767B\u9646 
```

src/cn/ixfosa/i18n/msg_en_US.properties

```properties
hello=Hello
username=User Name
pwd=Password
title=Login Page
submit=Submit \!
```

#### 静态数据国际化

国际化的软件：

1. 存储所有国家显示的文本的字符串
         a. 文件: properties
          b. 命名：  基础名_语言简称_国家简称.properties

   例如：msg_zh_CN.properties     存储所有中文
         Msg_en_US.properties    存储所有英文

   

2. 程序中获取
   ResourceBundle类，可以读取国际化的资源文件!

```java
// 国际化 - 静态数据
@Test
public void testI18N() throws Exception {

    // 中文语言环境
    Locale locale = Locale.US;

    // 创建工具类对象ResourceBundle
    ResourceBundle bundle = ResourceBundle.getBundle("cn.ixfosa.f_i18n.msg", locale);
    // 根据key获取配置文件中的值
    System.out.println(bundle.getString("hello"));
    System.out.println(bundle.getString("username"));
    System.out.println(bundle.getString("pwd"));

}
```



#### 动态文本国际化

数值，货币，时间，日期等数据由于可能在程序运行时动态产生，所以无法像文字一样简单地将它们从应用程序中分离出来，而是需要特殊处理。Java 中提供了解决这些问题的 API 类(位于 java.util 包和 java.text 包中)



中文：1987-09-19    ￥1000
英文： Sep/09 1987  $100

```java
// 国际化 - 动态文本 - 0. 概述
@Test
public void testI18N2() throws Exception {
    // 国际化货币
    NumberFormat.getCurrencyInstance();
    // 国际化数字
    NumberFormat.getNumberInstance();
    // 国际化百分比
    NumberFormat.getPercentInstance();  
    // 国际化日期
    //DateFormat.getDateTimeInstance(dateStyle, timeStyle, aLocale)
}
	
// 国际化 - 动态文本 - 1. 国际化货币
@Test
public void testI18N3() throws Exception {
    // 模拟语言环境
    Locale locale = Locale.CHINA;
    // 数据准备
    double number = 100;
    // 工具类
    NumberFormat nf = NumberFormat.getCurrencyInstance(locale);
    // 国际化货币
    String m = nf.format(number);
    // 测试
    System.out.println(m);
}
	
//面试题：  代码计算：  $100 * 10  
@Test
public void eg() throws Exception {
    String str = "$100";
    int num = 10;

    // 1. 分析str值是哪一国家的货币
    Locale us = Locale.US;
    // 2. 国际化工具类
    NumberFormat nf = NumberFormat.getCurrencyInstance(us);
    // 3. 解析str国币
    Number n = nf.parse(str);

    System.out.println(n.intValue() * num);
}

// 国际化 - 动态文本 - 2. 国际化数值
@Test
public void testI18N4() throws Exception {
    // 模拟语言环境
    Locale locale = Locale.CHINA;
    NumberFormat nf = NumberFormat.getNumberInstance(Locale.US);
    String str = nf.format(1000000000);
    System.out.println(str);
}
	
// 国际化 - 动态文本 - 3. 国际化日期
/*
	 * 日期
	 * 	  FULL   2015年3月4日 星期三
	 * 	  LONG   2015年3月4日
	 * 	  FULL   2015年3月4日 星期三
	 *    MEDIUM 2015-3-4
	 *    SHORT  15-3-4
	 *    
	 * 时间
	 * 	  FULL   下午04时31分59秒 CST
	 * 	  LONG   下午04时32分37秒
	 *    MEDIUM 16:33:00
	 *    SHORT  下午4:33
	 *    
	 * 
	 */
@Test
public void testI18N5() throws Exception {

    // 日期格式
    int dateStyle = DateFormat.SHORT;
    // 时间格式
    int timeStyle = DateFormat.SHORT;

    // 工具类
    DateFormat df = 
        DateFormat.getDateTimeInstance(dateStyle, timeStyle, Locale.CHINA);		
    String date = df.format(new Date());

    System.out.println(date);
}
	
// 面试2： 请将时间值：09-11-28 上午10时25分39秒 CST，反向解析成一个date对象。
@Test
public void eg2() throws Exception {
    String str = "09-11-28 上午10时25分39秒 CST";
    // 创建DateFormat工具类，国际化日期
    DateFormat df = DateFormat.getDateTimeInstance(
        DateFormat.SHORT, DateFormat.FULL, Locale.getDefault());
    Date d = df.parse(str);

    System.out.println(d);
}
```



#### Jsp页面国际化

```
<%@ page language="java" import="java.util.*" pageEncoding="UTF-8"%>

<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN">
<html>
  <head>
  	<%
  	ResourceBundle bundle = ResourceBundle.getBundle("cn.ixfosa.i18n.msg",request.getLocale());
  	%>
    <title><%=bundle.getString("title") %></title>
	<meta http-equiv="pragma" content="no-cache">
	<meta http-equiv="cache-control" content="no-cache">
	<meta http-equiv="expires" content="0">    
  </head>
  
  <body>
  <form name="frmLogin" action="${pageContext.request.contextPath }/admin?method=login" method="post">
  	<table align="center" border="1">
  		<tr>
  			<td><%=bundle.getString("username") %></td>
  			<td>
  				<input type="text" name="userName">
  			</td>
  		</tr>
  		<tr>
  			<td><%=bundle.getString("pwd") %></td>
  			<td>
  				<input type="password" name="pwd">
  			</td>
  		</tr>
  		<tr>
  			<td>
  				<input type="submit" value="<%=bundle.getString("submit") %>">
  			</td>
  		</tr>
  	</table>
  </form>
  </body>
</html>
```



#### jsp页面国际化-jstl标签

JSTL标签：
		核心标签库
		`国际化与格式化标签库`
		数据库标签库(没用)
		函数库

<fmt:setLocale value=""/>        设置本地化对象
  	<fmt:setBundle basename=""/>     设置工具类
  	<fmt:message></fmt:message>     显示国际化文本

格式化数值
		<fmt:formatNumber pattern="#.##" value="100.99"></fmt:formatNumber>
格式化日期:
		<fmt:formatDate pattern="yyyy-MM-dd" value="${date}"/>

```jsp
<%@ page language="java" import="java.util.*" pageEncoding="UTF-8"%>
<%--引入jstl国际化与格式化标签库 --%>
<%@taglib uri="http://java.sun.com/jsp/jstl/fmt" prefix="fmt" %>

<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN">
<html>
  <head>
  	<!-- 一、设置本地化对象 -->
  	<fmt:setLocale value="${pageContext.request.locale}"/>
      
  	<!-- 二、设置工具类 -->
  	<fmt:setBundle basename="cn.ixfosa.i18n.msg" var="bundle"/>

    <title><fmt:message key="title" bundle="${bundle}"></fmt:message></title>
	<meta http-equiv="pragma" content="no-cache">
	<meta http-equiv="cache-control" content="no-cache">
	<meta http-equiv="expires" content="0">    
  </head>
  
  <body>
  <form name="frmLogin" action="${pageContext.request.contextPath }/admin?method=login" method="post">
  	<table align="center" border="1">
  		<tr>
  			<td><fmt:message key="username" bundle="${bundle}"></fmt:message></td>
  			<td>
  				<input type="text" name="userName">
  			</td>
  		</tr>
  		<tr>
  			<td><fmt:message key="pwd" bundle="${bundle}"></fmt:message></td>
  			<td>
  				<input type="password" name="pwd">
  			</td>
  		</tr>
  		<tr>
  			<td>
  				<input type="submit" value="<fmt:message key="submit" bundle="${bundle}"/>">
  			</td>
  		</tr>
  	</table>
  </form>
  </body>
</html>
```





## 文件上传下载

前台：

 	1. 提交方式：post
	2. 表单中有文件上传的表单项： <input type="file" />
	3. 指定表单类型:
  默认类型：enctype="application/x-www-form-urlencoded"
  文件上传类型：multipart/form-data



### 手动实现文件上传

```jsp
<body>	
    <form name="frm_test" action="${pageContext.request.contextPath }/upload" method="post" enctype="multipart/form-data">
        用户名：<input type="text" name="userName">  <br/>
        文件：   <input type="file" name="file_img">   <br/>

        <input type="submit" value="注册">
    </form>
</body>
```

```java
package cn.ixfosa.upload;

import java.io.BufferedReader;
import java.io.File;
import java.io.FileInputStream;
import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.io.UnsupportedEncodingException;
import java.util.List;
import java.util.UUID;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;


public class UploadServlet extends HttpServlet {
	
	// 手动实现过程
	private void upload(HttpServletRequest request) throws IOException,
			UnsupportedEncodingException {
		/*
		request.getParameter(""); // GET/POST
		request.getQueryString(); // 获取GET提交的数据 
		request.getInputStream(); // 获取post提交的数据   */
		
		/***********手动获取文件上传表单数据************/
		
		//1. 获取表单数据流
		InputStream in =  request.getInputStream();
		//2. 转换流
		InputStreamReader inStream = new InputStreamReader(in, "UTF-8");
		//3. 缓冲流
		BufferedReader reader = new BufferedReader(inStream);
		// 输出数据
		String str = null;
		while ((str = reader.readLine()) != null) {
			System.out.println(str);
		}
		
		// 关闭
		reader.close();
		inStream.close();
		in.close();
	}

	public void doPost(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, IOException {
		this.doGet(request, response);
	}
}
```

### FileUpload组件

文件上传功能开发中比较常用，apache也提供了文件上传组件！
FileUpload组件: 
 	1. 下载源码
	2. 项目中引入jar文件
           commons-fileupload-1.2.1.jar  【文件上传组件核心jar包】
    	   commons-io-1.4.jar                  【封装了对文件处理的相关工具类】

   

```java
package cn.ixfosa.upload;

import java.io.BufferedReader;
import java.io.File;
import java.io.FileInputStream;
import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.io.UnsupportedEncodingException;
import java.util.List;
import java.util.UUID;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import org.apache.commons.fileupload.FileItem;
import org.apache.commons.fileupload.FileItemFactory;
import org.apache.commons.fileupload.FileUploadException;
import org.apache.commons.fileupload.disk.DiskFileItemFactory;
import org.apache.commons.fileupload.servlet.ServletFileUpload;
import org.apache.commons.io.FileUtils;

public class UploadServlet extends HttpServlet {

	// upload目录，保存上传的资源  WebRoot/upload
	public void doGet(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, IOException {
		
		/*********文件上传组件： 处理文件上传************/
		
		try {
			// 1. 文件上传工厂
			FileItemFactory factory = new DiskFileItemFactory();
			// 2. 创建文件上传核心工具类
			ServletFileUpload upload = new ServletFileUpload(factory);
		
			// 一、设置单个文件允许的最大的大小： 30M
			upload.setFileSizeMax(30*1024*1024);
			// 二、设置文件上传表单允许的总大小: 80M
			upload.setSizeMax(80*1024*1024);
			// 三、 设置上传表单文件名的编码
			// 相当于：request.setCharacterEncoding("UTF-8");
			upload.setHeaderEncoding("UTF-8");
			
			
			// 3. 判断： 当前表单是否为文件上传表单
			if (upload.isMultipartContent(request)){
				// 4. 把请求数据转换为一个个FileItem对象，再用集合封装
				List<FileItem> list = upload.parseRequest(request);
				// 遍历： 得到每一个上传的数据
				for (FileItem item: list){
					// 判断：普通文本数据
					if (item.isFormField()){
						// 普通文本数据
						String fieldName = item.getFieldName();	// 表单元素名称
						String content = item.getString();		// 表单元素名称， 对应的数据
						//item.getString("UTF-8");  指定编码
						System.out.println(fieldName + " " + content);
					}
					// 上传文件(文件流) ----> 上传到upload目录下
					else {
						// 普通文本数据
						String fieldName = item.getFieldName();	// 表单元素名称
						String name = item.getName();			// 文件名				
						String content = item.getString();		// 表单元素名称， 对应的数据
						String type = item.getContentType();	// 文件类型
						InputStream in = item.getInputStream(); // 上传文件流
						
						/*
						 *  四、文件名重名
						 *  对于不同用户readme.txt文件，不希望覆盖！
						 *  后台处理： 给用户添加一个唯一标记!
						 */
						// a. 随机生成一个唯一标记
						String id = UUID.randomUUID().toString();
						// b. 与文件名拼接
						name = id +"#"+ name;
						
						// 获取上传基路径
						String path = getServletContext().getRealPath("/upload");
						// 创建目标文件
						File file = new File(path,name);
						
						// 工具类，文件上传
						item.write(file);
						item.delete();   //删除系统产生的临时文件
						
						System.out.println();
					}
				}
			}
			else {
				System.out.println("当前表单不是文件上传表单，处理失败！");
			}
		} catch (Exception e) {
			e.printStackTrace();
		}
	}
}
```

### 案例

步骤：
1. 文件上传
2. 列表下载

#### Servlet

```java
package cn.ixfosa.servlet;

import java.io.File;
import java.io.FileInputStream;
import java.io.IOException;
import java.io.InputStream;
import java.io.OutputStream;
import java.net.URLEncoder;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.UUID;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import org.apache.commons.fileupload.FileItem;
import org.apache.commons.fileupload.FileItemFactory;
import org.apache.commons.fileupload.disk.DiskFileItemFactory;
import org.apache.commons.fileupload.servlet.ServletFileUpload;

/**
 * 处理文件上传与下载
 *
 */
public class FileServlet extends HttpServlet {

	public void doGet(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, IOException {

		// 获取请求参数： 区分不同的操作类型
		String method = request.getParameter("method");
		if ("upload".equals(method)) {
			// 上传
			upload(request,response);
		}
		
		else if ("downList".equals(method)) {
			// 进入下载列表
			downList(request,response);
		}
		
		else if ("down".equals(method)) {
			// 下载
			down(request,response);
		}
	}
	
	
	/**
	 * 1. 上传
	 */
	private void upload(HttpServletRequest request, HttpServletResponse response)
	throws ServletException, IOException {
		
		try {
			// 1. 创建工厂对象
			FileItemFactory factory = new DiskFileItemFactory();
			// 2. 文件上传核心工具类
			ServletFileUpload upload = new ServletFileUpload(factory);
			// 设置大小限制参数
			upload.setFileSizeMax(10*1024*1024);	// 单个文件大小限制
			upload.setSizeMax(50*1024*1024);		// 总文件大小限制
			upload.setHeaderEncoding("UTF-8");		// 对中文文件编码处理

			// 判断
			if (upload.isMultipartContent(request)) {
				// 3. 把请求数据转换为list集合
				List<FileItem> list = upload.parseRequest(request);
				// 遍历
				for (FileItem item : list){
					// 判断：普通文本数据
					if (item.isFormField()){
						// 获取名称
						String name = item.getFieldName();
						// 获取值
						String value = item.getString();
						System.out.println(value);
					} 
					// 文件表单项
					else {
						/******** 文件上传 ***********/
						// a. 获取文件名称
						String name = item.getName();
						// ----处理上传文件名重名问题----
						// a1. 先得到唯一标记
						String id = UUID.randomUUID().toString();
						// a2. 拼接文件名
						name = id + "#" + name;						
						
						// b. 得到上传目录 WebRoot/upload
						String basePath = getServletContext().getRealPath("/upload");
						// c. 创建要上传的文件对象
						File file = new File(basePath,name);
						// d. 上传
						item.write(file);
						item.delete();  // 删除组件运行时产生的临时文件
					}
				}
			}
		} catch (Exception e) {
			e.printStackTrace();
		}
	}

	
	/**
	 * 2. 进入下载列表
	 */
	private void downList(HttpServletRequest request, HttpServletResponse response)
	throws ServletException, IOException {
		
		// 实现思路：先获取upload目录下所有文件的文件名，再保存；跳转到down.jsp列表展示
		
		//1. 初始化map集合Map<包含唯一标记的文件名, 简短文件名>  ;
		Map<String,String> fileNames = new HashMap<String,String>();
		
		//2. 获取上传目录，及其下所有的文件的文件名
		String bathPath = getServletContext().getRealPath("/upload");
		// 目录
		File file = new File(bathPath);
		// 目录下，所有文件名
		String list[] = file.list();
		// 遍历，封装
		if (list != null && list.length > 0){
			for (int i=0; i<list.length; i++){
				// 全名
				String fileName = list[i];
				// 短名
				String shortName = fileName.substring(fileName.lastIndexOf("#")+1);
				// 封装
				fileNames.put(fileName, shortName);
			}
		}
		
		// 3. 保存到request域
		request.setAttribute("fileNames", fileNames);
		// 4. 转发
		request.getRequestDispatcher("/downlist.jsp").forward(request, response);

	}

	
	/**
	 *  3. 处理下载
	 */
	private void down(HttpServletRequest request, HttpServletResponse response)
	throws ServletException, IOException {
		
		// 获取用户下载的文件名称(url地址后追加数据,get)
		String fileName = request.getParameter("fileName");
		fileName = new String(fileName.getBytes("ISO8859-1"),"UTF-8");
		
		// 先获取上传目录路径
		String basePath = getServletContext().getRealPath("/upload");
		// 获取一个文件流
		InputStream in = new FileInputStream(new File(basePath,fileName));
		
		// 如果文件名是中文，需要进行url编码
		fileName = URLEncoder.encode(fileName, "UTF-8");
		// 设置下载的响应头
		response.setHeader("content-disposition", "attachment;fileName=" + fileName);
		
		// 获取response字节流
		OutputStream out = response.getOutputStream();
		byte[] b = new byte[1024];
		int len = -1;
		while ((len = in.read(b)) != -1){
			out.write(b, 0, len);
		}
		// 关闭
		out.close();
		in.close();
	}
	
	public void doPost(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, IOException {
		this.doGet(request, response);
	}
}
```

#### web.xml

```xml
<?xml version="1.0" encoding="UTF-8"?>
<web-app version="2.5" 
	xmlns="http://java.sun.com/xml/ns/javaee" 
	xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" 
	xsi:schemaLocation="http://java.sun.com/xml/ns/javaee 
	http://java.sun.com/xml/ns/javaee/web-app_2_5.xsd">
  <servlet>
    <servlet-name>FileServlet</servlet-name>
    <servlet-class>cn.ixfosa.servlet.FileServlet</servlet-class>
  </servlet>

  <servlet-mapping>
    <servlet-name>FileServlet</servlet-name>
    <url-pattern>/fileServlet</url-pattern>
  </servlet-mapping>
  <welcome-file-list>
    <welcome-file>index.jsp</welcome-file>
  </welcome-file-list>
</web-app>
```

#### jsp-入口

```jsp
<%@ page language="java" import="java.util.*" pageEncoding="UTF-8"%>

<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN">
<html>
  <head>
    
    <title>index</title>
	<meta http-equiv="pragma" content="no-cache">
	<meta http-equiv="cache-control" content="no-cache">
	<meta http-equiv="expires" content="0">    
  </head>
  
  <body>	
  		<a href="${pageContext.request.contextPath }/upload.jsp">文件上传</a> &nbsp;&nbsp;&nbsp;
  		<a href="${pageContext.request.contextPath }/fileServlet?method=downList">文件下载</a> 
  		
  </body>
</html>
```



#### jsp-上传

```jsp
<%@ page language="java" import="java.util.*" pageEncoding="UTF-8"%>

<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN">
<html>
  <head>
    
    <title>upload</title>
	<meta http-equiv="pragma" content="no-cache">
	<meta http-equiv="cache-control" content="no-cache">
	<meta http-equiv="expires" content="0">    
  </head>
  
  <body>	
  	 <form name="frm_test" action="${pageContext.request.contextPath }/fileServlet?method=upload" method="post" enctype="multipart/form-data">
  	 	 <%--<input type="hidden" name="method" value="upload">--%>
  	 	 
  	 	 用户名：<input type="text" name="userName">  <br/>
  	 	 文 件：<input type="file" name="file_img">   <br/>
  	 	 
  	 	<input type="submit" value="提交">
   	 </form>
  </body>
</html>
```

#### jsp-下载

```jsp
<%@ page language="java" import="java.util.*" pageEncoding="UTF-8"%>
<%@taglib uri="http://java.sun.com/jsp/jstl/core" prefix="c" %>

<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN">
<html>
  <head>
    
    <title>下载列表</title>
	<meta http-equiv="pragma" content="no-cache">
	<meta http-equiv="cache-control" content="no-cache">
	<meta http-equiv="expires" content="0">    
  </head>
  
  <body>	
	<table border="1" align="center">
		<tr>
			<th>序号</th>
			<th>文件名</th>
			<th>操作</th>
		</tr>
		<c:forEach var="en" items="${requestScope.fileNames}" varStatus="vs">
			<tr>
				<td>${vs.count }</td>
				<td>${en.value }</td>
				<td>
					<%--<a href="${pageContext.request.contextPath }/fileServlet?method=down&..">下载</a>--%>
					<!-- 构建一个地址  -->
					<c:url var="url" value="fileServlet">
						<c:param name="method" value="down"></c:param>
						<c:param name="fileName" value="${en.key}"></c:param>
					</c:url>
					<!-- 使用上面地址 -->
					<a href="${url }">下载</a>
				</td>
			</tr>
		</c:forEach>
	</table>  		
  </body>
</html>
```



## JavaMail邮件

### 邮件开发准备

准备工作, 环境搭建：

1. 本地搭建一个邮件服务器

   ​		易邮服务器，eyoumailserversetup.exe

2. 新建邮箱账号
     张三给李四发邮件。
       步骤1：
                新建域名： 工具， 服务器设置， 单域名框中输入 ixfosa.com
       步骤2：
     		   新建邮箱账号:  zhangsan@ixfosa.com
                                         lisi@ixfosa.com

3. 安装foxmail
           配置邮件发送服务器(smtp)：  localhost      25
           邮件接收服务器(pop3)：    localhost     110
      再新建账号，就可以接收邮件了！

4. 注意
  Java project。
  如果是web项目，因为javaee自带的有邮件功能，可能存在问题！
  我们要用自己的mail.jar文件功能！  需要删除javaee中mail包!

5. ![mail](images\mail.jpg)



### 普通邮件

JavaMail开发，先引入jar文件：
	activation.jar   【如果使用jdk1.6或以上版本，可以不用这个jar文件】
	mail.jar           【邮件发送核心包】

```java
package cn.ixfosa.mail;

import java.util.Date;
import java.util.Properties;

import javax.mail.Session;
import javax.mail.Transport;
import javax.mail.internet.InternetAddress;
import javax.mail.internet.MimeMessage;
import javax.mail.internet.MimeMultipart;
import javax.mail.internet.MimeMessage.RecipientType;

import org.junit.Test;

/**
 * 1. 发送一封普通邮件
 *
 */
public class App_1SendMail {

	@Test
	public void testSend() throws Exception {
		
		//0. 邮件参数
		Properties prop = new Properties();
		prop.put("mail.transport.protocol", "smtp");	// 指定协议
		prop.put("mail.smtp.host", "localhost");		// 主机   stmp.qq.com
		prop.put("mail.smtp.port", 25);					// 端口
		prop.put("mail.smtp.auth", "true");				// 用户密码认证
		prop.put("mail.debug", "true");					// 调试模式
		
		//1. 创建一个邮件的会话
		Session session = Session.getDefaultInstance(prop);
		//2. 创建邮件体对象 (整封邮件对象)
		MimeMessage message = new MimeMessage(session);
		//3. 设置邮件体参数: 
		//3.1 标题
		message.setSubject("我的第一封邮件	");
		//3.2 邮件发送时间
		message.setSentDate(new Date());
		//3.3 发件人
		message.setSender(new InternetAddress("zhangsan@ixfosa.com"));
		//3.4 接收人
		message.setRecipient(RecipientType.TO, new InternetAddress("lisi@ixfosa.com"));
		//3.5内容
		//message.setText("你好，已经发送成功！  正文....");  // 简单纯文本邮件
		// 邮件中含有超链接
		//message.setText("<a href='#'>百度</a>");
		message.setContent("<a href='#'>百度</a>", "text/html;charset=UTF-8");
		
		message.saveChanges();   // 保存邮件(可选)
		
		//4. 发送
		Transport trans = session.getTransport();
		trans.connect("zhangsan", "888");
		// 发送邮件
		trans.sendMessage(message, message.getAllRecipients());
		trans.close();
	}
}
```



### 带图片

```java
package cn.ixfosa.mail;

import java.io.File;
import java.util.Date;
import java.util.Properties;

import javax.activation.DataHandler;
import javax.activation.DataSource;
import javax.activation.FileDataSource;
import javax.mail.Session;
import javax.mail.Transport;
import javax.mail.internet.AddressException;
import javax.mail.internet.InternetAddress;
import javax.mail.internet.MimeBodyPart;
import javax.mail.internet.MimeMessage;
import javax.mail.internet.MimeMultipart;
import javax.mail.internet.MimeMessage.RecipientType;

import org.junit.Test;

/**
 * 带图片资源的邮件
 *
 */
public class App_2SendWithImg {
	
	// 初始化参数
	private static Properties prop;
	// 发件人
	private static InternetAddress sendMan = null;
	static {
		prop = new Properties();
		prop.put("mail.transport.protocol", "smtp");	// 指定协议
		prop.put("mail.smtp.host", "localhost");		// 主机   stmp.qq.com
		prop.put("mail.smtp.port", 25);					// 端口
		prop.put("mail.smtp.auth", "true");				// 用户密码认证
		prop.put("mail.debug", "true");					// 调试模式
		try {
			sendMan = new InternetAddress("zhangsan@ixfosa.com");
		} catch (AddressException e) {
			throw new RuntimeException(e);
		}
	}

	@Test
	public void testSend() throws Exception {
		// 1. 创建邮件会话
		Session session = Session.getDefaultInstance(prop);
		// 2. 创建邮件对象
		MimeMessage message = new MimeMessage(session);
		// 3. 设置参数：标题、发件人、收件人、发送时间、内容
		message.setSubject("带图片邮件");
		message.setSender(sendMan);
		message.setRecipient(RecipientType.TO, new InternetAddress("lisi@ixfosa.com"));
		message.setSentDate(new Date());
		
		/***************设置邮件内容: 多功能用户邮件 (related)*******************/
		// 4.1 构建一个多功能邮件块
		MimeMultipart related = new MimeMultipart("related");
		// 4.2 构建多功能邮件块内容 = 左侧文本 + 右侧图片资源
		MimeBodyPart content = new MimeBodyPart();
		MimeBodyPart resource = new MimeBodyPart();
		
		// 设置具体内容: a.资源(图片)  src/cn/ixfosa/mail/8.jpg 
		String filePath = App_2SendWithImg.class.getResource("8.jpg").getPath();
		DataSource ds = new FileDataSource(new File(filePath));
		DataHandler handler = new DataHandler(ds);
		resource.setDataHandler(handler);
		resource.setContentID("8.jpg");   // 设置资源名称，给外键引用
		
		// 设置具体内容: b.文本
		content.setContent("<img src='cid:8.jpg'/>  好哈哈！", "text/html;charset=UTF-8");
		
		related.addBodyPart(content);
		related.addBodyPart(resource);
		
		/*******4.3 把构建的复杂邮件快，添加到邮件中********/
		message.setContent(related);
		
		
		// 5. 发送
		Transport trans = session.getTransport();
		trans.connect("zhangsan", "888");
		trans.sendMessage(message, message.getAllRecipients());
		trans.close();
	}
}
```



### 带图片 + 附件

```java
package cn.itcast.mail;

import java.io.File;
import java.util.Date;
import java.util.Properties;

import javax.activation.DataHandler;
import javax.activation.DataSource;
import javax.activation.FileDataSource;
import javax.mail.Session;
import javax.mail.Transport;
import javax.mail.internet.AddressException;
import javax.mail.internet.InternetAddress;
import javax.mail.internet.MimeBodyPart;
import javax.mail.internet.MimeMessage;
import javax.mail.internet.MimeMultipart;
import javax.mail.internet.MimeMessage.RecipientType;

import org.junit.Test;

/**
 * 3. 带图片资源以及附件的邮件
 *
 */
public class App_3ImgAndAtta {
	
	// 初始化参数
	private static Properties prop;
	// 发件人
	private static InternetAddress sendMan = null;
	static {
		prop = new Properties();
		prop.put("mail.transport.protocol", "smtp");	// 指定协议
		prop.put("mail.smtp.host", "localhost");		// 主机   stmp.qq.com
		prop.put("mail.smtp.port", 25);					// 端口
		prop.put("mail.smtp.auth", "true");				// 用户密码认证
		prop.put("mail.debug", "true");					// 调试模式
		try {
			sendMan = new InternetAddress("zhangsan@ixfosa.com");
		} catch (AddressException e) {
			throw new RuntimeException(e);
		}
	}

	@Test
	public void testSend() throws Exception {
		// 1. 创建邮件会话
		Session session = Session.getDefaultInstance(prop);
		// 2. 创建邮件对象
		MimeMessage message = new MimeMessage(session);
		// 3. 设置参数：标题、发件人、收件人、发送时间、内容
		message.setSubject("带图片邮件");
		message.setSender(sendMan);
		message.setRecipient(RecipientType.TO, new InternetAddress("lisi@ixfosa.com"));
		message.setSentDate(new Date());
		
		/*
		 * 带附件(图片)邮件开发
		 */
		// 构建一个总的邮件块
		MimeMultipart mixed = new MimeMultipart("mixed");
		// ---> 总邮件快，设置到邮件对象中
		message.setContent(mixed);
		// 左侧： （文本+图片资源）
		MimeBodyPart left = new MimeBodyPart();
		// 右侧： 附件
		MimeBodyPart right = new MimeBodyPart();
		// 设置到总邮件块
		mixed.addBodyPart(left);
		mixed.addBodyPart(right);
		
		/******附件********/   src/cn/ixfosa/mail/a.docx
		String attr_path = this.getClass().getResource("a.docx").getPath();
		DataSource attr_ds = new FileDataSource(new File(attr_path));
		DataHandler attr_handler = new DataHandler(attr_ds);
		right.setDataHandler(attr_handler);
		right.setFileName("a.docx");
		
		
		/***************设置邮件内容: 多功能用户邮件 (related)*******************/
		// 4.1 构建一个多功能邮件块
		MimeMultipart related = new MimeMultipart("related");
		// ----> 设置到总邮件快的左侧中
		left.setContent(related);
		
		// 4.2 构建多功能邮件块内容 = 左侧文本 + 右侧图片资源
		MimeBodyPart content = new MimeBodyPart();
		MimeBodyPart resource = new MimeBodyPart();
		
		// 设置具体内容: a.资源(图片)  src/cn/ixfosa/mail/8.jpg
		String filePath = App_3ImgAndAtta.class.getResource("8.jpg").getPath();
		DataSource ds = new FileDataSource(new File(filePath));
		DataHandler handler = new DataHandler(ds);
		resource.setDataHandler(handler);
		resource.setContentID("8.jpg");   // 设置资源名称，给外键引用
		
		// 设置具体内容: b.文本
		content.setContent("<img src='cid:8.jpg'/>  好哈哈！", "text/html;charset=UTF-8");
		
		related.addBodyPart(content);
		related.addBodyPart(resource);
		
		
		// 5. 发送
		Transport trans = session.getTransport();
		trans.connect("zhangsan", "888");
		trans.sendMessage(message, message.getAllRecipients());
		trans.close();
	}
}
```





## other

### factory

src/instance.properties

```properties
类名=完整类名
```

```
package cn.ixfosa.factory;

import java.util.ResourceBundle;

public class BeanFactory {
	
	private static ResourceBundle bundle ;
	static{
		bundle = ResourceBundle.getBundle("instance");
	}
	
	public static <T> T getInstance(String key,Class<T> clazz){
		try {
			String className = bundle.getString(key);
			return (T) Class.forName(className).newInstance();
		} catch (Exception e) {
			throw new RuntimeException(e);
		}
		
	}
}
```



### 跳转的通用方法

```java
package cn.itcast.utils;

import java.io.IOException;

import javax.servlet.RequestDispatcher;
import javax.servlet.ServletException;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

public class WebUtils {

	/**
	 * 跳转的通用方法
	 */
	public static void goTo(HttpServletRequest request, HttpServletResponse response, Object uri)
			throws ServletException, IOException {
		if (uri instanceof RequestDispatcher){
			((RequestDispatcher)uri).forward(request, response);
		} else if (uri instanceof String) {
			response.sendRedirect(request.getContextPath() + uri);
		} 
	}
}
```



### 全站乱码问题

```java

import javax.servlet.*;
import javax.servlet.annotation.WebFilter;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.io.IOException;

/**
 * 解决全站乱码问题，处理所有的请求
 */
@WebFilter("/*")
public class CharchaterFilter implements Filter {
    @Override
    public void init(FilterConfig filterConfig) throws ServletException {

    }

    @Override
    public void doFilter(ServletRequest req, ServletResponse rep, FilterChain filterChain) throws IOException, ServletException {
        //将父接口转为子接口
        HttpServletRequest request = (HttpServletRequest) req;
        HttpServletResponse response = (HttpServletResponse) rep;
        //获取请求方法
        String method = request.getMethod();
        //解决post请求中文数据乱码问题
        if(method.equalsIgnoreCase("post")){
            request.setCharacterEncoding("utf-8");
        }
        //处理响应乱码
        response.setContentType("text/html;charset=utf-8");
        filterChain.doFilter(request,response);
    }

    @Override
    public void destroy() {

    }
}
```

### BaseServlet

```java

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.io.IOException;
import java.lang.reflect.InvocationTargetException;
import java.lang.reflect.Method;


public class BaseServlet extends HttpServlet {


    @Override
    protected void service(HttpServletRequest req, HttpServletResponse resp) throws ServletException, IOException {
        //System.out.println("baseServlet的service方法被执行了...");

        //完成方法分发
        //1.获取请求路径
        String uri = req.getRequestURI(); //   /travel/user/add
        System.out.println("请求uri:"+uri);//  /travel/user/add
        //2.获取方法名称
        String methodName = uri.substring(uri.lastIndexOf('/') + 1);
        System.out.println("方法名称："+methodName);
        //3.获取方法对象Method
        //谁调用我？我代表谁
        System.out.println(this);//UserServlet的对象cn.itcast.travel.web.servlet.UserServlet@4903d97e
        try {
            //获取方法
            Method method = this.getClass().getMethod(methodName, HttpServletRequest.class, HttpServletResponse.class);
            //4.执行方法
            //暴力反射
            //method.setAccessible(true);
            method.invoke(this,req,resp);
        } catch (NoSuchMethodException e) {
            e.printStackTrace();
        } catch (IllegalAccessException e) {
            e.printStackTrace();
        } catch (InvocationTargetException e) {
            e.printStackTrace();
        }
    }

    /**
     * 直接将传入的对象序列化为json，并且写回客户端
     * @param obj
     */
    public void writeValue(Object obj,HttpServletResponse response) throws IOException {
        ObjectMapper mapper = new ObjectMapper();
        response.setContentType("application/json;charset=utf-8");
        mapper.writeValue(response.getOutputStream(),obj);
    }

    /**
     * 将传入的对象序列化为json，返回
     * @param obj
     * @return
     */
    public String writeValueAsString(Object obj) throws JsonProcessingException {
        ObjectMapper mapper = new ObjectMapper();
        return mapper.writeValueAsString(obj);
    }

}

```



### LoginFilter

```java


import javax.servlet.*;
import javax.servlet.annotation.WebFilter;
import javax.servlet.http.HttpServletRequest;
import java.io.IOException;

/**
 * 登录验证的过滤器
 */
@WebFilter("/*")
public class LoginFilter implements Filter {


    public void doFilter(ServletRequest req, ServletResponse resp, FilterChain chain) throws ServletException, IOException {
        System.out.println(req);
        //0.强制转换
        HttpServletRequest request = (HttpServletRequest) req;

        //1.获取资源请求路径
        String uri = request.getRequestURI();
        //2.判断是否包含登录相关资源路径,要注意排除掉 css/js/图片/验证码等资源
        if(uri.contains("/login.jsp") || uri.contains("/loginServlet") || uri.contains("/css/") || uri.contains("/js/") || uri.contains("/fonts/") || uri.contains("/checkCodeServlet")  ){
            //包含，用户就是想登录。放行
            chain.doFilter(req, resp);
        }else{
            //不包含，需要验证用户是否登录
            //3.从获取session中获取user
            Object user = request.getSession().getAttribute("user");
            if(user != null){
                //登录了。放行
                chain.doFilter(req, resp);
            }else{
                //没有登录。跳转登录页面

                request.setAttribute("login_msg","您尚未登录，请登录");
                request.getRequestDispatcher("/login.jsp").forward(request,resp);
            }
        }


        // chain.doFilter(req, resp);
    }

    public void init(FilterConfig config) throws ServletException {

    }

    public void destroy() {
    }

}
```

