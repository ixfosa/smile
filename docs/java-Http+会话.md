##  Http协议

### Http协议入门

​		`http协议`： http是一个简单的`请求-响应`协议，它通常运行在TCP之上。对浏览器客户端 和  服务器端 之间数据传输的格式规范。

​		HTTP是`应用层`协议，同其他应用层协议一样

​		HTTP协议是基于`C/S架构`进行通信的，而HTTP协议的服务器端实现程序有httpd、nginx等，其客户端的实现程序主要是Web浏览器，Web服务是基于`TCP`的，因此为了能够随时响应客户端的请求，Web服务器需要监听在80/TCP端口。

​		HTTP使用统一资源标识符（Uniform Resource Identifiers, `URI`）来传输数据和建立连接。



### Http协议工作原理

HTTP是基于客户/服务器模式，且面向连接的。典型的HTTP事务处理有如下的过程：
	（1）客户与服务器建立连接；
	（2）客户向服务器提出请求；
	（3）服务器接受请求，并根据请求返回相应的文件作为应答；
	（4）客户与服务器关闭连接



HTTP三点注意事项：

​		HTTP是`无连接`：无连接的含义是限制每次连接`只处理一个请求`。服务器处理完客户的请求，并收到客户的应答后，即断开连接。采用这种方式可以节省传输时间。

​		HTTP是`媒体独立的`一种`面向对象`的协议。允许传送任意类型的数据对象。它通过数据类型和长度来标识所传送的数据内容和大小，并允许对数据进行压缩传送。只要客户端和服务器知道如何处理的数据内容，任何类型的数据都可以通过HTTP发送。客户端以及服务器指定使用适合的`MIME-type`内容类型。

​		HTTP是`无状态`：HTTP协议是无状态协议。无状态是指协议对于事务处理没有记忆能力。缺少状态意味着如果后续处理需要前面的信息，则它必须重传，这样可能导致每次连接传送的数据量增大。另一方面，在服务器不需要先前信息时它的应答就较快。



例如：在浏览器地址栏键入URL，按下回车之后会经历以下流程：

1. 浏览器向 DNS 服务器请求解析该 URL 中的域名所对应的 IP 地址;
2. 解析出 IP 地址后，根据该 IP 地址和默认端口 80，和服务器建立TCP连接;
3. 浏览器发出读取文件(URL 中域名后面部分对应的文件)的HTTP 请求，该请求报文作为 TCP 三次握手的第三个报文的数据发送给服务器;
4. 服务器对浏览器请求作出响应，并把对应的 html 文本发送给浏览器;
5. 释放 TCP连接;
6. 浏览器将该 html 文本并显示内容; 　



### Http 请求

客户端发送一个HTTP请求到服务器的请求消息包括以下格式：

`请求行`（request line）

`请求头部`（header）

`空行`

``请求数据`

		1. get: 请求行中 

```http
GET /test/testMethod.html?name=ixfosa&password=123456 HTTP/1.1   请求行
```



2. post：一个空行之后, 实体内容中 

```http
                                        --一个空行
name=ixfosa&password=123456             --（可选）实体内容(post)（post时请求数据）
```



四个部分组成

```http
--GET /test/testMethod.html?name=ixfosa&password=123456 HTTP/1.1 -请求行（get时请求数据）
POST /hello/hello HTTP/1.1               -请求行
Host: localhost:8080                    --请求头（多个key-value对象）
User-Agent: Mozilla/5.0 (Windows NT 6.1; WOW64; rv:35.0) Gecko/20100101 Firefox/35.0
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8
Accept-Language: zh-cn,en-us;q=0.8,zh;q=0.5,en;q=0.3
Accept-Encoding: gzip, deflate
Connection: keep-alive
                                        --一个空行
name=ixfosa&password=123456             --（可选）实体内容(post)（post时请求数据）
```

![客户端请求消息](images\客户端请求消息.png)



#### 请求行

##### http协议版本

http1.0：当前浏览器客户端与服务器端建立连接之后，只能发送一次请求，一次请求之后连接关闭。http1.1：当前浏览器客户端与服务器端建立连接之后，可以在一次连接中发送多次请求。

##### 请求资源

`URL`:  统一资源定位符。http://localhost:8080/test/test.html。只能定位互联网资源。是URI的子集。
`URI`： 统一资源标记符。/hello/hello。用于标记任何资源。可以是本地文件系统，局域网的资源（//192.168.14.10/myweb/index.html），可以是互联网。

##### 请求方式

根据 HTTP 标准，HTTP 请求可以使用多种请求方法。

HTTP1.0 定义了三种请求方法： `GET`, `POST` 和 HEAD方法。

HTTP1.1 新增了六种请求方法：OPTIONS、PUT、PATCH、DELETE、TRACE 和 CONNECT 方法。

| 序 号 | 方法    | 描述                                                         |
| :---: | :------ | :----------------------------------------------------------- |
|   1   | GET     | 请求指定的页面信息，并返回实体主体。                         |
|   2   | HEAD    | 类似于 GET 请求，只不过返回的响应中没有具体的内容，用于获取报头 |
|   3   | POST    | 向指定资源提交数据进行处理请求（例如提交表单或者上传文件）。数据被包含在请求体中。POST 请求可能会导致新的资源的建立和/或已有资源的修改。 |
|   4   | PUT     | 从客户端向服务器传送的数据取代指定的文档的内容。             |
|   5   | DELETE  | 请求服务器删除指定的页面。                                   |
|   6   | CONNECT | HTTP/1.1 协议中预留给能够将连接改为管道方式的代理服务器。    |
|   7   | OPTIONS | 允许客户端查看服务器的性能。                                 |
|   8   | TRACE   | 回显服务器收到的请求，主要用于测试或诊断。                   |
|   9   | PATCH   | 是对 PUT 方法的补充，用来对已知资源进行局部更新 。           |



**常用的请求方式： `GET`  和 `POST`**	

```html
<form action="提交地址" method="GET/POST">
    
</form>
```

**GET   vs  POST 区别**

1. GET方式提交 
   		a）地址栏（URI）会跟上参数数据。以？开头，多个参数之间以&分割。

```http
GET /test/testMethod.html?name=ixfosa&password=123456 HTTP/1.1
Host: localhost:8080
User-Agent: Mozilla/5.0 (Windows NT 6.1; WOW64; rv:35.0) Gecko/20100101 Firefox/35.0
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8
Accept-Language: zh-cn,en-us;q=0.8,zh;q=0.5,en;q=0.3
Accept-Encoding: gzip, deflate
Referer: http://localhost:8080/day09/testMethod.html
Connection: keep-alive
```

​				b）GET提交参数数据有限制，不超过1KB。
​				c）GET方式不适合提交敏感密码。
​				d）注意： 浏览器直接访问的请求，默认提交方式是GET方式

	2. POST方式提交
			a）参数不会跟着URI后面。参数而是跟在请求的实体内容中。没有？开头，多个参数之间以&分割。

```http
POST /day09/testMethod.html HTTP/1.1
Host: localhost:8080
User-Agent: Mozilla/5.0 (Windows NT 6.1; WOW64; rv:35.0) Gecko/20100101 Firefox/35.0
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8
Accept-Language: zh-cn,en-us;q=0.8,zh;q=0.5,en;q=0.3
Accept-Encoding: gzip, deflate
Referer: http://localhost:8080/day09/testMethod.html
Connection: keep-alive

name=ixfosa&password=123456
```

​				b）POST提交的参数数据没有限制。
​				c）POST方式提交敏感数据。

#### 请求头

多个`key-value`结构

```http
Accept: text/html,image/*                           -- 浏览器接受的数据类型
Accept-Charset: ISO-8859-1                          -- 浏览器接受的编码格式
Accept-Encoding: gzip,compress                      --浏览器接受的数据压缩格式
Accept-Language: en-us,zh-                           --浏览器接受的语言
Host: www.ixfosa.org:80          --（必须的）当前请求访问的目标地址（主机:端口）
If-Modified-Since: Tue, 11 Jul 2000 18:23:51 GMT                 --浏览器最后的缓存时间
Referer: http://www.it315.org/index.jsp                          -- 当前请求来自于哪里
User-Agent: Mozilla/4.0 (compatible; MSIE 5.5; Windows NT 5.0)   --浏览器类型
Cookie:name=eric                        -- 浏览器保存的cookie信息
Connection: close/Keep-Alive            -- 浏览器跟服务器连接状态。close: 连接关闭  keep-alive：                                 --保存连接。
Date: Tue, 11 Jul 2000 18:23:51 GMT     -- 请求发出的时间
```

#### 实体内容

只有`POST`提交的参数会放到实体内容中



### Http 响应

```http
HTTP/1.1 200 OK                           --响应行
Server: Apache-Coyote/1.1                 --响应头（key-vaule）
Content-Length: 24 
Date: Fri, 30 Jan 2015 01:54:57 GMT
                                          --一个空行
this is hello servlet!!!                  --实体内容
```

HTTP响应也由四个部分组成，分别是：`状态行`、`消息报头`、`空行`和`响应正文`。

![服务器响应消息](images\服务器响应消息.png)



#### 响应行

1. http协议版本:

2. 状态码: 服务器处理请求的结果（状态）

​		常见的状态：

​				200 ：  表示请求处理完成并完美返回
​				302：  表示请求需要进一步细化。
​				404：  表示客户访问的资源找不到。
​				500：  表示服务器的资源发送错误。（服务器内部错误）

3. 状态描述	



#### 响应头

```http
Location: http://www.ixfosa.org/index.jsp   -- 表示重定向的地址，该头和302的状态码一起使用。
Server:apache tomcat                        -- 表示服务器的类型
Content-Encoding: gzip                      -- 表示服务器发送给浏览器的数据压缩类型
Content-Length: 80                          --表示服务器发送给浏览器的数据长度
Content-Language: zh-cn                     --表示服务器支持的语言
Content-Type: text/html; charset=GB2312     --表示服务器发送给浏览器的数据类型及内容编码
Last-Modified: Tue, 11 Jul 2000 18:23:51 GMT  --表示服务器资源的最后修改时间
Refresh: 1;url=http://www.ixfosa.org           --表示定时刷新
Content-Disposition: attachment; filename=aaa.zip --表示告诉浏览器以下载方式打开资源（下载文												    件时用到）
Transfer-Encoding: chunked
Set-Cookie:SS=Q0=5Lb_nQ; path=/search    --表示服务器发送给浏览器的cookie信息（会话管理用到）
Expires: -1                              --表示通知浏览器不进行缓存
Cache-Control: no-cache
Pragma: no-cache
Connection: close/Keep-Alive             -- 表示服务器和浏览器的连接状态。close：关闭连接 											    keep-alive:保存连接
```



### 状态消息

常见的HTTP状态码：
		`200 `- 请求成功
		`301` - 资源（网页等）被永久转移到其它URL
		`404` - 请求的资源（网页等）不存在
		`500 `- 内部服务器错误

#### 1xx:信息

| 消息                    | 描述                                                         |
| ----------------------- | ------------------------------------------------------------ |
| 100 Continue            | 服务器仅接收到部分请求，但是一旦服务器并没有拒绝该请求，客户端应该继续发送其余的请求。 |
| 101 Switching Protocols | 服务器转换协议：服务器将遵从客户的请求转换到另外一种协议。   |

#### 2xx:成功

| 消息                              | 描述                                                         |
| --------------------------------- | ------------------------------------------------------------ |
| `200 OK`                          | 请求成功（其后是对GET和POST请求的应答文档。）                |
| 201 Created                       | 请求被创建完成，同时新的资源被创建。                         |
| 202 Accepted                      | 供处理的请求已被接受，但是处理未完成。                       |
| 203 Non-authoritative Information | 文档已经正常地返回，但一些应答头可能不正确，因为使用的是文档的拷贝。 |
| 204 No Content                    | 没有新文档。浏览器应该继续显示原来的文档。如果用户定期地刷新页面，而Servlet可以确定用户文档足够新，这个状态代码是很有用的。 |
| 205 Reset Content                 | 没有新文档。但浏览器应该重置它所显示的内容。用来强制浏览器清除表单输入内容。 |
| 206 Partial Content               | 客户发送了一个带有Range头的GET请求，服务器完成了它。         |

#### 3xx:重定向

| 消息                   | 描述                                                         |
| ---------------------- | ------------------------------------------------------------ |
| 300 Multiple Choices   | 多重选择。链接列表。用户可以选择某链接到达目的地。最多允许五个地址。 |
| 301 Moved Permanently  | 所请求的页面已经转移至新的url。                              |
| `302Found`             | 所请求的页面已经临时转移至新的url。                          |
| 303 See Other          | 所请求的页面可在别的url下被找到。                            |
| 304 Not Modified       | 未按预期修改文档。客户端有缓冲的文档并发出了一个条件性的请求（一般是提供If-Modified-Since头表示客户只想比指定日期更新的文档）。服务器告诉客户，原来缓冲的文档还可以继续使用。 |
| 305 Use Proxy          | 客户请求的文档应该通过Location头所指明的代理服务器提取。     |
| 306 *Unused*           | 此代码被用于前一版本。目前已不再使用，但是代码依然被保留。   |
| 307 Temporary Redirect | 被请求的页面已经临时移至新的url。                            |

#### 4xx:客户端错误

| 消息                                | 描述                                                         |
| ----------------------------------- | ------------------------------------------------------------ |
| `400 Bad Request`                   | 服务器未能理解请求。                                         |
| 401 Unauthorized                    | 被请求的页面需要用户名和密码。                               |
| 402 Payment Required                | 此代码尚无法使用。                                           |
| 403 Forbidden                       | 对被请求页面的访问被禁止。                                   |
| `404 Not Found`                     | 服务器无法找到被请求的页面。                                 |
| 405 Method Not Allowed              | 请求中指定的方法不被允许。                                   |
| 406 Not Acceptable                  | 服务器生成的响应无法被客户端所接受。                         |
| 407 Proxy Authentication Required   | 用户必须首先使用代理服务器进行验证，这样请求才会被处理。     |
| 408 Request Timeout                 | 请求超出了服务器的等待时间。                                 |
| 409 Conflict                        | 由于冲突，请求无法被完成。                                   |
| 410 Gone                            | 被请求的页面不可用。                                         |
| 411 Length Required                 | "Content-Length" 未被定义。如果无此内容，服务器不会接受请求。 |
| 412 Precondition Failed             | 请求中的前提条件被服务器评估为失败。                         |
| 413 Request Entity Too Large        | 由于所请求的实体的太大，服务器不会接受请求。                 |
| 414 Request-url Too Long            | 由于url太长，服务器不会接受请求。当post请求被转换为带有很长的查询信息的get请求时，就会发生这种情况。 |
| 415 Unsupported Media Type          | 由于媒介类型不被支持，服务器不会接受请求。                   |
| 416 Requested Range Not Satisfiable | 服务器不能满足客户在请求中指定的Range头。                    |
| 417 Expectation Failed              | 执行失败。                                                   |
| 423                                 | 锁定的错误。                                                 |

#### 5xx:服务器错误

|      | 消息                           | 描述                                               |
| ---- | ------------------------------ | -------------------------------------------------- |
|      | `500 Internal Server Error`    | 请求未完成。服务器遇到不可预知的情况。             |
|      | 501 Not Implemented            | 请求未完成。服务器不支持所请求的功能。             |
|      | 502 Bad Gateway                | 请求未完成。服务器从上游服务器收到一个无效的响应。 |
|      | 503 Service Unavailable        | 请求未完成。服务器临时过载或宕机。                 |
|      | 504 Gateway Timeout            | 网关超时。                                         |
|      | 505 HTTP Version Not Supported | 服务器不支持请求中指明的HTTP协议版本。             |



### HTTP content-type

Content-Type（内容类型），一般是指网页中存在的 Content-Type，用于定义网络文件的类型和网页的编码，决定浏览器将以什么形式、什么编码读取这个文件，这就是经常看到一些 PHP 网页点击的结果却是下载一个文件或一张图片的原因。

`Content-Type` 标头告诉客户端实际返回的内容的内容类型。

语法格式：

```http
Content-Type: text/html; charset=utf-8
Content-Type: multipart/form-data; boundary=something
```

![HTTP content-type](images\HTTP content-type.png)



常见的媒体格式类型如下：

- text/html ： HTML格式
- text/plain ：纯文本格式
- text/xml ： XML格式
- image/gif ：gif图片格式
- image/jpeg ：jpg图片格式
- image/png：png图片格式

以application开头的媒体格式类型：

- `application/xhtml+xml` ：XHTML格式
- `application/xml`： XML数据格式
- application/atom+xml ：Atom XML聚合格式
- `application/json`： JSON数据格式
- application/pdf：pdf格式
- application/msword ： Word文档格式
- application/octet-stream ： 二进制流数据（如常见的文件下载）
- application/x-www-form-urlencoded ： <form encType=””>中默认的encType，form表单数据被编码为key/value格式发送到服务器（表单默认的提交数据的格式）

另外一种常见的媒体格式是上传文件之时使用的：

- `multipart/form-data` ： 需要在表单中进行文件上传时，就需要使用该格式







## 会话管理

一次会话： 打开浏览器 -> 访问一些服务器内容 -> 关闭浏览器

会话管理： 管理浏览器客户端 和 服务器端之间会话过程中产生的会话数据。

`Cookie`技术：会话数据保存在`浏览器客户端`。

`Session`技术：会话数据保存在`服务器端`。

域对象： 实现资源之间的数据共享。

### Cooke技术

Cookie技术：会话数据保存在浏览器客户端。



#### Cookie Api

​			Cookie类：用于存储会话数据

```java
1. 构造Cookie对象
	Cookie(java.lang.String name, java.lang.String value)
    
2. 设置cookie
	void setPath(java.lang.String uri);       ：设置cookie的有效访问路径
	void setMaxAge(int expiry);               ：设置cookie的有效时间
	void setValue(java.lang.String newValue); ：设置cookie的值
    
3. 发送cookie到浏览器端保存
    void response.addCookie(Cookie cookie);   : 发送cookie
    
4. 服务器接收cookie
    Cookie[] request.getCookies();            : 接收cookie
```

- Cookie原理

```java
1.服务器创建cookie对象，把会话数据存储到cookie对象中。
  	new Cookie("name","value");
  					
2.服务器发送cookie信息到浏览器
  	response.addCookie(cookie);
	举例： set-cookie: name=eric  (隐藏发送了一个set-cookie名称的响应头)
	
3.浏览器得到服务器发送的cookie，然后保存在浏览器端。

4.浏览器在下次访问服务器时，会带着cookie信息
	举例： cookie: name=eric  (隐藏带着一个叫cookie名称的请求头)
	
5.服务器接收到浏览器带来的cookie信息
	request.getCookies();
```

#### 创建，接受cookie

```java
package gz.ixfosa.cookie;

import java.io.IOException;

import javax.servlet.ServletException;
import javax.servlet.http.Cookie;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
/**
 * 第一个cookie的程序
 *
 */
public class CookieDemo extends HttpServlet {

	public void doGet(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, IOException {
        
		//1.创建Cookie对象
		Cookie cookie1 = new Cookie("name","ixfosa");
		//Cookie cookie2 = new Cookie("email","ixfosa@qq.com");
		
		
		/**
		 * 1）设置cookie的有效路径。默认情况：有效路径在当前web应用下。 /test
		 */
		//cookie1.setPath("/test");
		//cookie2.setPath("/test2");
		
		/**
		 * 2)设置cookie的有效时间
		 * 正整数 ：表示cookie数据保存浏览器的缓存目录（硬盘中），数值表示保存的时间。
		   负整数 ：表示cookie数据保存浏览器的内存中。浏览器关闭cookie就丢失了！！
		   零    ：表示删除同名的cookie数据

		 */
		//cookie1.setMaxAge(20); //20秒，从最后不调用cookie开始计算
		cookie1.setMaxAge(-1); //cookie保存在浏览器内存（会话cookie）
		//cookie1.setMaxAge(0);
		

		//2.把cookie数据发送到浏览器（通过响应头发送： set-cookie名称）
		//response.setHeader("set-cookie", cookie.getName()+"="+cookie.getValue()+",email=ixfosa@qq.com");
        
		//推荐使用这种方法，避免手动发送cookie信息
		response.addCookie(cookie1);
		//response.addCookie(cookie2);
		

		
		//3.接收浏览器发送的cookie信息
		/*String name = request.getHeader("cookie");
		System.out.println(name);*/
		Cookie[] cookies = request.getCookies();
        
		//注意：判断null,否则空指针
		if(cookies!=null){
			//遍历
			for(Cookie c:cookies){
				String name = c.getName();
				String value = c.getValue();
				System.out.println(name+"="+value);
			}
		}else{
			System.out.println("没有接收cookie数据");
		}
	}
}
```



Cookie的细节

1. void setPath(java.lang.String uri)   ：设置cookie的有效访问路径。

   有效路径指的是cookie的有效路径保存在哪里，那么浏览器在有效路径下访问服务器时就会带着cookie信息，否则不带cookie信息。
   		

2. void setMaxAge(int expiry) ： 设置cookie的有效时间。

​                    `正整数`：表示cookie数据保存浏览器的缓存目录（硬盘中），数值表示保存的时间。

​					`负整数`：表示cookie数据保存浏览器的内存中。浏览器关闭cookie就丢失了！！

​					`零`        ：表示删除同名的cookie数据

3. Cookie数据类型只能保存非中文字符串类型的。可以保存多个cookie，但是浏览器一般只允许存放300个Cookie，每个站点最多存放20个Cookie，每个·Cookie的大小限制为`4KB`。



#### 删除cookie

```java
package gz.ixfosa.cookie;

import java.io.IOException;

import javax.servlet.ServletException;
import javax.servlet.http.Cookie;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

public class DeleteCookie extends HttpServlet {

	public void doGet(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, IOException {
        
		/**
		 * 需求： 删除cookie
		 */
		Cookie cookie = new Cookie("name","xxxx");
		cookie.setMaxAge(0);  //删除同名的cookie
		response.addCookie(cookie);
		System.out.println("删除成功");
	}
}
```



#### 案例- 显示用户上次访问的时间



```java
package gz.ixfosa.cookie;

import java.io.IOException;
import java.text.SimpleDateFormat;
import java.util.Date;

import javax.servlet.ServletException;
import javax.servlet.http.Cookie;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
/**
 * 案例-用户上次访问时间
 *
 */
public class HistServlet extends HttpServlet {

	public void doGet(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, IOException {
        
		response.setContentType("text/html;charset=utf-8");
		
		//获取当前时间
		SimpleDateFormat format = new SimpleDateFormat("yyyy-MM-dd hh:mm:ss");
		String curTime = format.format(new Date());

		
		//取得cookie
		Cookie[] cookies = request.getCookies();
		String lastTime = null;
		if(cookies!=null){
			for (Cookie cookie : cookies) {
				if(cookie.getName().equals("lastTime")){
                    
					//有lastTime的cookie，已经是第n次访问
					lastTime = cookie.getValue();//上次访问的时间
					//第n次访问
					//1.把上次显示时间显示到浏览器
					response.getWriter().write("欢迎回来，你上次访问的时间为："+lastTime+",当前时间为："+curTime);
					//2.更新cookie
					cookie.setValue(curTime);
					cookie.setMaxAge(1*30*24*60*60);
					//3.把更新后的cookie发送到浏览器
					response.addCookie(cookie);
					break;
				}
			}
		}
		
		/**
		 * 第一次访问（没有cookie 或 有cookie，但没有名为lastTime的cookie）
		 */
		if(cookies==null || lastTime==null){
			//1.显示当前时间到浏览器
			response.getWriter().write("你是首次访问本网站，当前时间为："+curTime);
			//2.创建Cookie对象
			Cookie cookie = new Cookie("lastTime",curTime);
			cookie.setMaxAge(1*30*24*60*60);//保存一个月
			//3.把cookie发送到浏览器保存
			response.addCookie(cookie);
		}
	}
}
```



### Session技术

Cookie的局限：

1. Cookie只能存字符串类型。不能保存对象
2. 只能存非中文。
3. 1个Cookie的容量不超过4KB。

​		如果要保存非字符串，超过4kb内容，只能使用session技术！！！



Session特点：
				会话数据保存在服务器端。（内存中）



#### Session Api

​		`HttpSession`类：用于保存会话数据

```java
1.创建或得到session对象
	HttpSession getSession()  
	ttpSession getSession(boolean create)  
    
2.设置session对象
	void setMaxInactiveInterval(int interval)      ： 设置session的有效时间
	oid invalidate()                               ： 销毁session对象
	java.lang.String getId()                       ： 得到session编号
  
3.保存会话数据到session对象
	void setAttribute(java.lang.String name, java.lang.Object value)  ： 保存数据
	Object getAttribute(java.lang.String name)                        ： 获取数据
	void removeAttribute(java.lang.String name)                       ： 清除数据
```



- Session原理 
  		问题： 服务器能够识别不同的浏览者！！！

     	 前提： 在哪个session域对象保存数据，就必须从哪个域对象取出！

```java
浏览器1：(给s1分配一个唯一的标记：s001,把s001发送给浏览器)

创建session对象，保存会话数据
	HttpSession session = request.getSession();   --保存会话数据 s1
	
浏览器1的新窗口:（带着s001的标记到服务器查询，s001->s1,返回s1）	
	得到session对象的会话数据
	HttpSession session = request.getSession();   --可以取出  s1

新的浏览器1：(没有带s001,不能返回s1)
	得到session对象的会话数据
	HttpSession session = request.getSession();   --不可以取出  s2

浏览器2：(没有带s001,不能返回s1)
	得到session对象的会话数据
	HttpSession session = request.getSession();   --不可以取出  s3
```

​	

- 代码解读：HttpSession session = request.getSession();

```java
1）第一次访问创建session对象，给session对象分配一个唯一的ID，叫JSESSIONID
	new HttpSession();
				
2）把JSESSIONID作为Cookie的值发送给浏览器保存
	Cookie cookie = new Cookie("JSESSIONID", sessionID);
	response.addCookie(cookie);
	
3）第二次访问的时候，浏览器带着JSESSIONID的cookie访问服务器

4）服务器得到JSESSIONID，在服务器的内存中搜索是否存放对应编号的session对象。
				if(找到){
					return map.get(sessionID);
				}
				Map<String,HttpSession>
                    
                <"s001", s1>
				<"s001,"s2>
                    
5）如果找到对应编号的session对象，直接返回该对象
                    
6）如果找不到对应编号的session对象，创建新的session对象，继续走1的流程

		
```

**结论：通过`JSESSIONID`的cookie值在服务器找session对象！**

- 
  Sesson细节

```java
1）java.lang.String getId()  ： 得到session编号
    
2）两个getSession方法：
	getSession(true) / getSession(); : 创建或得到session对象。没有匹配的session编号，自动创										建新的session对象。
	getSession(false)                : 得到session对象。没有匹配的session编号，返回null
    
3）void setMaxInactiveInterval(int interval)  ：设置session的有效时间
	session对象销毁时间：
		3.1 默认情况30分服务器自动回收
		3.2 修改session回收时间
		3.3 全局修改session有效时间
```



#### 全局修改session有效时间

```xml
<!-- 修改session全局有效时间:分钟 -->
<session-config>
    <session-timeout>1</session-timeout>
</session-config>
```

#### 保存会话数据到session域对象

```java
package gz.ixfosa.session;

import java.io.IOException;

import javax.servlet.ServletException;
import javax.servlet.http.Cookie;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import javax.servlet.http.HttpSession;
/**
 * 保存会话数据到session域对象
 *
 */
public class SessionDemo extends HttpServlet {

	public void doGet(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, IOException {
        
		//1.创建session对象
		HttpSession session = request.getSession();
		
		/**
		 * 得到session编号
		 */
		System.out.println("id="+session.getId());
		
		/**
		 * 修改session的有效时间
		 */
		//session.setMaxInactiveInterval(20);
		
		/**
		 * 手动发送一个硬盘保存的cookie给浏览器
		 */
		Cookie c = new Cookie("JSESSIONID",session.getId());
		c.setMaxAge(60*60);
		response.addCookie(c);
		
		//2.保存会话数据
		session.setAttribute("name", "rose");
	}
}
```

#### 从session域对象中取出会话数据

```java
package gz.ixfosa.session;

import java.io.IOException;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import javax.servlet.http.HttpSession;
/**
 * 从session域对象中取出会话数据
 *
 */
public class SessionDemo2 extends HttpServlet {

	public void doGet(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, IOException {
        
		//1.到session对象。没有匹配的session编号，返回null
		HttpSession session = request.getSession(false);
		
		if(session==null){
			System.out.println("没有找到对应的sessino对象");
			return;
		}
		
		/**
		 * 得到session编号
		 */
		System.out.println("id="+session.getId());
		
		//2.取出数据
		String name = (String)session.getAttribute("name");
		System.out.println("name="+name);
	}
}
```



#### 手动销毁session对象

```java
void invalidate()     ： 销毁session对象
```

```java
package gz.ixfosa.session;

import java.io.IOException;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import javax.servlet.http.HttpSession;

/**
 * 销毁session对象
 *
 */
public class DeleteSession extends HttpServlet {

	public void doGet(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, IOException {
		HttpSession session = request.getSession(false);
		if(session!=null){
			session.invalidate();//手动销毁
		}
		System.out.println("销毁成功");
	}
}
```

​	

如何避免浏览器的JSESSIONID的cookie随着浏览器关闭而丢失的问题

手动发送一个硬盘保存的cookie给浏览器

```java
Cookie c = new Cookie("JSESSIONID",session.getId());
c.setMaxAge(60*60);
response.addCookie(c);
```

​	

### 案例-cookie-商品推荐

#### 实体

src/gz/ixfosa/entity

```java
package gz.ixfosa.entity;

public class Product {

	private String id;
	private String proName;
	private String proType;
	private double price;
	public String getId() {
		return id;
	}
	public void setId(String id) {
		this.id = id;
	}
	public String getProName() {
		return proName;
	}
	public void setProName(String proName) {
		this.proName = proName;
	}
	public String getProType() {
		return proType;
	}
	public void setProType(String proType) {
		this.proType = proType;
	}
	public double getPrice() {
		return price;
	}
	public void setPrice(double price) {
		this.price = price;
	}
	public Product(String id, String proName, String proType, double price) {
		super();
		this.id = id;
		this.proName = proName;
		this.proType = proType;
		this.price = price;
	}
	public Product() {
		super();
		// TODO Auto-generated constructor stub
	}
	@Override
	public String toString() {
		return "Product [id=" + id + ", price=" + price + ", proName="
				+ proName + ", proType=" + proType + "]";
	}
}
```



#### DAO

src/gz/ixfosa/dao/ProductDao.java

```java
package gz.ixfosa.dao;

import gz.itcast.entity.Product;

import java.util.ArrayList;
import java.util.List;

/**
 * 该类中存放对Prodcut对象的CRUD方法
 *
 */
public class ProductDao {
    
	//模拟"数据库",存放所有商品数据
	private static List<Product> data = new ArrayList<Product>();
	
	/**
	 * 初始化商品数据
	 */
	static{
		//只执行一次
		for(int i=1;i<=10;i++){
			data.add(new Product(""+i,"笔记本"+i,"LN00"+i,34.0+i));
		}
	}
		
	
	/**
	 * 提供查询所有商品的方法
	 */
	public List<Product> findAll(){
		return data;
	}
	
	/**
	 * 提供根据编号查询商品的方法
	 */
	public Product findById(String id){
		for(Product p : data){
			if(p.getId().equals(id)){
				return p;
			}
		}
		return null;
	}
}
```

#### Servlet

```java
package gz.ixfosa.servlet;

import gz.itcast.dao.ProductDao;
import gz.itcast.entity.Product;

import java.io.IOException;
import java.io.PrintWriter;
import java.util.List;

import javax.servlet.ServletException;
import javax.servlet.http.Cookie;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

/**
 * 查询所有商品的servlet
 *
 */
public class ListServlet extends HttpServlet {

	public void doGet(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, IOException {
        
		response.setContentType("text/html;charset=utf-8");
        
		//1.读取数据库，查询商品列表
		ProductDao dao = new ProductDao();
		List<Product> list = dao.findAll();
		
		
		//2.把商品显示到浏览器
		PrintWriter writer = response.getWriter();
		String html = "";
		
		html += "<html>";
		html += "<head>";
		html += "<title>显示商品列表</title>";
		html += "</head>";
		html += "<body>";
		html += "<table border='1' align='center' width='600px'>";
		html += "<tr>";
		html += "<th>编号</th><th>商品名称</th><th>商品型号</th><th>商品价格</th>";
		html += "</tr>";
		//遍历商品
		if(list!=null){
			for(Product p:list){
				html += "<tr>";
				// /hist/DetailServlet?id=1 访问DetailSErvlet的servlet程序，同时传递 名为id，值为1 的参数
				html += "<td>"+p.getId()+"</td><td><a href='"+request.getContextPath()+"/DetailServlet?id="+p.getId()+"'>"+p.getProName()+"</a></td><td>"+p.getProType()+"</td><td>"+p.getPrice()+"</td>";
				html += "<tr>";
			}
		}
		html += "</table>";
		
		/**
		 * 显示浏览过的商品
		 */
		html += "最近浏览过的商品：<br/>";
		//取出prodHist的cookie
		Cookie[] cookies = request.getCookies();
		if(cookies!=null){
			for (Cookie cookie : cookies) {
				if(cookie.getName().equals("prodHist")){
					String prodHist = cookie.getValue(); // 3,2,1
					String[] ids = prodHist.split(",");
					//遍历浏览过的商品id
					for (String id : ids) {
						//查询数据库，查询对应的商品
						Product p = dao.findById(id);
						//显示到浏览器
						html += ""+p.getId()+"&nbsp;"+p.getProName()+"&nbsp;"+p.getPrice()+"<br/>";
					}
				}
			}
		}

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



```java
package gz.ixfosa.servlet;

import gz.itcast.dao.ProductDao;
import gz.itcast.entity.Product;

import java.io.IOException;
import java.io.PrintWriter;
import java.util.Arrays;
import java.util.Collection;
import java.util.LinkedList;

import javax.servlet.ServletException;
import javax.servlet.http.Cookie;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
/**
 * 显示商品详细
 *
 */
public class DetailServlet extends HttpServlet {

	public void doGet(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, IOException {
        
		response.setContentType("text/html;charset=utf-8");
        
		//1.获取编号
		String id = request.getParameter("id");
		
		//2.到数据库中查询对应编号的商品
		ProductDao dao = new ProductDao();
		Product product = dao.findById(id);
		
		//3.显示到浏览器
		PrintWriter writer = response.getWriter();
		String html = "";
		
		html += "<html>";
		html += "<head>";
		html += "<title>显示商品详细</title>";
		html += "</head>";
		html += "<body>";
		html += "<table border='1' align='center' width='300px'>";
		if(product!=null){
			html += "<tr><th>编号:</th><td>"+product.getId()+"</td></tr>";
			html += "<tr><th>商品名称:</th><td>"+product.getProName()+"</td></tr>";
			html += "<tr><th>商品型号:</th><td>"+product.getProType()+"</td></tr>";
			html += "<tr><th>商品价格:</th><td>"+product.getPrice()+"</td></tr>";
		}
		
		html += "</table>";
		html += "<center><a href='"+request.getContextPath()+"/ListServlet'>[返回列表]</a></center>";
		html += "</body>";
		html += "</html>";
		
		writer.write(html);
		
		
		/**
		 * 创建cookie，并发送
		 */
		//1.创建cookie
		Cookie cookie = new Cookie("prodHist",createValue(request,id));
		cookie.setMaxAge(1*30*24*60*60);//一个月
		//2.发送cookie
		response.addCookie(cookie);
	}

/**
	 * 生成cookie的值
	 * 分析：
	 * 	当前cookie值       传入商品id    最终cookie值
	 *  null或没有prodHist     1         1     算法： 直接返回传入的id 
	 *         1               2        2,1   没有重复且小于3个。算法：直接把传入的id放最前面 
	 *         2,1             1        1,2   有重复且小于3个。算法：去除重复id，把传入的id放	 *									      最前面
	 *         3,2,1           2        2,3,1 有重复且3个。算法：去除重复id，把传入的id放最前      *                                        面
	 *         3,2,1           4        4,3,2 没有重复且3个。算法：去最后的id，把传入的id放最      *                                         前面
	 * @return
	 */
	private String createValue(HttpServletRequest request,String id) {
		
		Cookie[] cookies = request.getCookies();
		String prodHist = null;
		if(cookies!=null){
			for (Cookie cookie : cookies) {
				if(cookie.getName().equals("prodHist")){
					prodHist = cookie.getValue();
					break;
				}
			}
		}
		
		// null或没有prodHist
		if(cookies==null || prodHist==null){
			//直接返回传入的id
			return id;
		}
		
		// 3,21          2
		//String -> String[] ->  Collection :为了方便判断重复id
		String[] ids = prodHist.split(",");
		Collection colls = Arrays.asList(ids); //<3,21>
		// LinkedList 方便地操作（增删改元素）集合
		// Collection -> LinkedList
		LinkedList list = new LinkedList(colls);
		
		
		//不超过3个
		if(list.size()<3){
			//id重复
			if(list.contains(id)){
				//去除重复id，把传入的id放最前面
				list.remove(id);
				list.addFirst(id);
			}else{
				//直接把传入的id放最前面
				list.addFirst(id);
			}
		}else{
			//等于3个
			//id重复
			if(list.contains(id)){
				//去除重复id，把传入的id放最前面
				list.remove(id);
				list.addFirst(id);
			}else{
				//去最后的id，把传入的id放最前面
				list.removeLast();
				list.addFirst(id);
			}
		}
		
		// LinedList -> String 
		StringBuffer sb = new StringBuffer();
		for (Object object : list) {
			sb.append(object+",");
		}
		//去掉最后的逗号
		String result = sb.toString();
		result = result.substring(0, result.length()-1);
		return result;
	}

	public void doPost(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, IOException {
		doGet(request, response);
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
    <servlet-name>ListServlet</servlet-name>
    <servlet-class>gz.ixfosa.servlet.ListServlet</servlet-class>
  </servlet>
  <servlet>
    <servlet-name>DetailServlet</servlet-name>
    <servlet-class>gz.ixfosa.servlet.DetailServlet</servlet-class>
  </servlet>


  <servlet-mapping>
    <servlet-name>ListServlet</servlet-name>
    <url-pattern>/ListServlet</url-pattern>
  </servlet-mapping>
  <servlet-mapping>
    <servlet-name>DetailServlet</servlet-name>
    <url-pattern>/DetailServlet</url-pattern>
  </servlet-mapping>
    
  <welcome-file-list>
    <welcome-file>index.jsp</welcome-file>
  </welcome-file-list>
</web-app>
```

### 案例-Session-登录

#### 静态页面

`WebRoot/login.html`

```html
<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN">
<html>
  <head>
    <title>登录页面</title>
	
    <meta http-equiv="keywords" content="keyword1,keyword2,keyword3">
    <meta http-equiv="description" content="this is my page">
    <meta http-equiv="content-type" content="text/html; charset=UTF-8">
    
  </head>
  
  <body>
    <form action="/login/LoginServlet" method="post">
    	用户名:<input type="text" name="userName"/>
    	<br/>
    	密码：<input type="password" name="userPwd"/>
    	<br/>
    	<input type="submit" value="登录"/>
    </form>
  </body>
</html>
```

`WebRoot/fail.html`

```html
<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN">
<html>
  <head>
    <title>信息提示页面</title>
	
    <meta http-equiv="keywords" content="keyword1,keyword2,keyword3">
    <meta http-equiv="description" content="this is my page">
    <meta http-equiv="content-type" content="text/html; charset=UTF-8">
    
  </head>
  
  <body>
    <font color='red' size='3'> 你的用户名或密码输入有误！请重新输入!</font><br/>
    <a href="/login/login.html">返回登录页面</a>
  </body>
</html>
```

#### Servlet

##### 用户主页的逻辑

```java
package gz.ixfosa;

import java.io.IOException;
import java.io.PrintWriter;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import javax.servlet.http.HttpSession;
/**
 * 用户主页的逻辑
 *
 */
public class IndexServlet extends HttpServlet {

	public void doGet(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, IOException {
		response.setContentType("text/html;charset=utf-8");
		PrintWriter writer = response.getWriter();
		
		
		String html = "";
		
		/**
		 * 接收request域对象的数据
		 */
		/*
		String loginName = (String)request.getAttribute("loginName");
		*/
		
		/**
		 * 二、在用户主页，判断session不为空且存在指定的属性才视为登录成功！才能访问资源。
		 * 从session域中获取会话数据
		 */
		//1.得到session对象
		HttpSession session = request.getSession(false);
		if(session==null){
			//没有登录成功，跳转到登录页面
			response.sendRedirect(request.getContextPath()+"/login.html");
			return;
		}
		//2.取出会话数据
		String loginName = (String)session.getAttribute("loginName");
		if(loginName==null){
			//没有登录成功，跳转到登录页面
			response.sendRedirect(request.getContextPath()+"/login.html");
			return;
		}
		
		html = "<html><body>欢迎回来，"+loginName+"，<a href='"+request.getContextPath()+"/LogoutServlet'>安全退出</a></body></html>";
		
		writer.write(html);
	}

	public void doPost(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, IOException {
		doGet(request, response);
	}
}
```



##### 处理登录的逻辑

```java
package gz.ixfosa;

import java.io.IOException;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import javax.servlet.http.HttpSession;
/**
 * 处理登录的逻辑
 *
 */
public class LoginServlet extends HttpServlet {

	public void doGet(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, IOException {
        
		request.setCharacterEncoding("utf-8");
        
		//1.接收参数
		String userName = request.getParameter("userName");
		String userPwd = request.getParameter("userPwd");
		
		//2.判断逻辑
		if("ixfosa".equals(userName)
				 && "123456".equals(userPwd)){
			//登录成功
			/**
			 * 分析：
			 * 	  context域对象：不合适，可能会覆盖数据。
			 *    request域对象： 不合适，整个网站必须得使用转发技术来跳转页面
			 *    session域对象：合适。
			 */
			/*
			request.setAttribute("loginName", userName);
			//request.getRequestDispatcher("/IndexServlet").forward(request, response);
			response.sendRedirect(request.getContextPath()+"/IndexServlet");
			*/
			
			/**
			 * 一、登录成功后，把用户数据保存session对象中
			 */
			//1.创建session对象
			HttpSession session = request.getSession();
			//2.把数据保存到session域中
			session.setAttribute("loginName", userName);
			//3.跳转到用户主页
			response.sendRedirect(request.getContextPath()+"/IndexServlet");
			
		}else{
			//登录失败
			//请求重定向
			response.sendRedirect(request.getContextPath()+"/fail.html");
		}
	}

	public void doPost(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, IOException {
		doGet(request, response);
	}
}
```

##### 退出逻辑

```java
package gz.ixfosa;

import java.io.IOException;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import javax.servlet.http.HttpSession;

/**
 * 退出逻辑
 *
 */
public class LogoutServlet extends HttpServlet {

	public void doGet(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, IOException {
		/**
		 * 三、安全退出：
		 * 		删除掉session对象中指定的loginName属性即可！  
		 */
		//1.得到session对象
		HttpSession session = request.getSession(false);
		if(session!=null){
			//2.删除属性
			session.removeAttribute("loginName");
		}
		
		//2.回来登录页面
		response.sendRedirect(request.getContextPath()+"/login.html");
	}

	public void doPost(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, IOException {
		doGet(request, response);
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
    <servlet-name>LoginServlet</servlet-name>
    <servlet-class>gz.ixfosa.LoginServlet</servlet-class>
  </servlet>
  <servlet>
    <servlet-name>IndexServlet</servlet-name>
    <servlet-class>gz.ixfosa.IndexServlet</servlet-class>
  </servlet>
  <servlet>
    <servlet-name>LogoutServlet</servlet-name>
    <servlet-class>gz.ixfosa.LogoutServlet</servlet-class>
  </servlet>


  <servlet-mapping>
    <servlet-name>LoginServlet</servlet-name>
    <url-pattern>/LoginServlet</url-pattern>
  </servlet-mapping>
  <servlet-mapping>
    <servlet-name>IndexServlet</servlet-name>
    <url-pattern>/IndexServlet</url-pattern>
  </servlet-mapping>
  <servlet-mapping>
    <servlet-name>LogoutServlet</servlet-name>
    <url-pattern>/LogoutServlet</url-pattern>
  </servlet-mapping>
  <welcome-file-list>
    <welcome-file>index.jsp</welcome-file>
  </welcome-file-list>
</web-app>
```











