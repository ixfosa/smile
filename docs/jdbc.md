## jdbc

JDBC API 允许用户访问任何形式的表格数据，尤其是存储在关系数据库中的数据。

执行流程：

- 连接数据源，如：数据库。
- 为数据库传递查询和更新指令。
- 处理数据库响应并返回的结果。

使用java代码（程序）`发送sql语句`的技术，就是jdbc技术！！！！



## JDBC 架构

#### 双层

![双层](E:/smile/java/images/双层.jpg)

作用：此架构中，Java Applet 或应用直接访问数据源。

条件：要求 Driver 能与访问的数据库交互。

机制：用户命令传给数据库或其他数据源，随之结果被返回。

部署：数据源可以在另一台机器上，用户通过网络连接，称为 `C/S配置`（可以是内联网或互联网）。

#### 三层

![三层](E:/smile/java/images/三层.jpg)



侧架构特殊之处在于，引入中间层服务。

流程：命令和结构都会经过该层。

吸引：可以增加企业数据的访问控制，以及多种类型的更新；另外，也可简化应用的部署，并在多数情况下有性能优势。

历史趋势： 以往，因性能问题，中间层都用 C 或 C++ 编写，随着优化编译器（将 Java 字节码 转为 高效的 特定机器码）和技术的发展，如EJB，Java 开始用于中间层的开发这也让 Java 的优势突显出现出来，使用 Java 作为服务器代码语言，JDBC随之被重视。



## JDBC 编程步骤

使用jdbc发送sql前提
		登录数据库服务器（连接数据库服务器）

   * 数据库的IP地址
   * 端口
   * 数据库用户名
   * 密码

jar: mysql-connector-java-5.1.7-bin.jar

1. 加载驱动程序：

```java
Class.forName(driverClass)
//加载MySql驱动
Class.forName("com.mysql.jdbc.Driver")
//加载Oracle驱动
Class.forName("oracle.jdbc.driver.OracleDriver")
```

2. 获得数据库连接：

```java
DriverManager.getConnection("jdbc:mysql://127.0.0.1:3306/ixfosa", "root", "root");
```

3. 创建Statement\PreparedStatement对象：

```java
conn.createStatement();
conn.prepareStatement(sql);
```

```java
package gz.ixfosa.jdbc;

import java.sql.Connection;
import java.sql.Driver;
import java.sql.DriverManager;
import java.util.Properties;

import org.junit.Test;
/**
 * jdbc连接数据库
 *
 */
public class Demo {
    
	//连接数据库的URL
	private String url = "jdbc:mysql://localhost:3306/ixfosa";
	                    // jdbc协议:数据库子协议:主机:端口/连接的数据库   //

	private String user = "root";//用户名
	private String password = "root";//密码
	
	/**
	 * 第一种方法
	 * @throws Exception
	 */
	@Test
	public void test1() throws Exception{
		//1.创建驱动程序类对象
		Driver driver = new com.mysql.jdbc.Driver(); //新版本
		//Driver driver = new org.gjt.mm.mysql.Driver(); //旧版本
		
		//设置用户名和密码
		Properties props = new Properties();
		props.setProperty("user", user);
		props.setProperty("password", password);
		
		//2.连接数据库，返回连接对象
		Connection conn = driver.connect(url, props);
		
		System.out.println(conn);
	}
	
	/**
	 * 使用驱动管理器类连接数据库(注册了两次，没必要)
	 * @throws Exception
	 */
	@Test
	public void test2() throws Exception{
        
		Driver driver = new com.mysql.jdbc.Driver();
		//Driver driver2 = new com.oracle.jdbc.Driver();
        
		//1.注册驱动程序(可以注册多个驱动程序)
		DriverManager.registerDriver(driver);
		//DriverManager.registerDriver(driver2);
		
		//2.连接到具体的数据库
		Connection conn = DriverManager.getConnection(url, user, password);
		System.out.println(conn);
		
	}
	
	/**
	 * （推荐使用这种方式连接数据库）
	 * 推荐使用加载驱动程序类  来 注册驱动程序 
	 * @throws Exception
	 */
	@Test
	public void test3() throws Exception{
        
		//Driver driver = new com.mysql.jdbc.Driver();
		
		//通过得到字节码对象的方式加载静态代码块，从而注册驱动程序
		Class.forName("com.mysql.jdbc.Driver");
		
		//Driver driver2 = new com.oracle.jdbc.Driver();
		//1.注册驱动程序(可以注册多个驱动程序)
		//DriverManager.registerDriver(driver);
		//DriverManager.registerDriver(driver2);
		
		//2.连接到具体的数据库
		Connection conn = DriverManager.getConnection(url, user, password);
		System.out.println(conn);
	}
}
```



## JDBC接口核心的API

```java
java.sql.*   和  javax.sql.*

|- Driver接口： 表示java驱动程序接口。所有的具体的数据库厂商要来实现此接口。
    |- connect(url, properties):  连接数据库的方法。
    	url: 连接数据库的URL 
        URL语法： jdbc协议:数据库子协议://主机:端口/数据库
		user： 数据库的用户名
  		  password： 数据库用户密码
            
|- DriverManager类： 驱动管理器类，用于管理所有注册的驱动程序
    |-registerDriver(driver)  : 注册驱动类对象
    |-Connection getConnection(url,user,password);  获取连接对象

|- Connection接口： 表示java程序和数据库的连接对象。
	|- Statement createStatement() ： 创建Statement对象
	|- PreparedStatement prepareStatement(String sql)  创建PreparedStatement对象
	|- CallableStatement prepareCall(String sql) 创建CallableStatement对象

|- Statement接口： 用于执行静态的sql语句
		|- int executeUpdate(String sql)  ： 执行静态的更新sql语句（DDL，DML）
		|- ResultSet executeQuery(String sql)  ：执行的静态的查询sql语句（DQL）
            
	|-PreparedStatement接口：用于执行预编译sql语句
			|- int executeUpdate() ： 执行预编译的更新sql语句（DDL，DML）
			|-ResultSet executeQuery()  ： 执行预编译的查询sql语句（DQL）

		|-CallableStatement接口：用于执行存储过程的sql语句（call xxx）
			|-ResultSet executeQuery()  ： 调用存储过程的方法
            
|- ResultSet接口：用于封装查询出来的数据
	|- boolean next() ： 将光标移动到下一行
	|-getXX() : 获取列的值
```



## Statement执行sql语句

### 执行DDL语句

```java
package gz.ixfosa.statement;

import java.sql.Connection;
import java.sql.DriverManager;
import java.sql.SQLException;
import java.sql.Statement;

import org.junit.Test;

/**
 * 使用Statement对象执行静态sql语句
 *
 */
public class Demo1 {

	private String url = "jdbc:mysql://localhost:3306/day17";
	private String user = "root";
	private String password = "root";
	/**
	 * 执行DDL语句(创建表)
	 */
	@Test
	public void test1(){
		Statement stmt = null;
		Connection conn = null;
		try {
			//1.驱动注册程序
			Class.forName("com.mysql.jdbc.Driver");
			
			//2.获取连接对象
			conn = DriverManager.getConnection(url, user, password);
			
			//3.创建Statement
			stmt = conn.createStatement();
			
			//4.准备sql
			String sql = "CREATE TABLE student(id INT PRIMARY KEY AUTO_INCREMENT,NAME VARCHAR(20),gender VARCHAR(2))";
			
			//5.发送sql语句，执行sql语句,得到返回结果
			int count = stmt.executeUpdate(sql);
			
			//6.输出
			System.out.println("影响了"+count+"行！");
		} catch (Exception e) {
			e.printStackTrace();
			throw new RuntimeException(e);
		} finally{
			//7.关闭连接(顺序:后打开的先关闭)
			if(stmt!=null)
				try {
					stmt.close();
				} catch (SQLException e) {
					e.printStackTrace();
					throw new RuntimeException(e);
				}
			if(conn!=null)
				try {
					conn.close();
				} catch (SQLException e) {
					e.printStackTrace();
					throw new RuntimeException(e);
				}
		}		
	}
}
```



### 执行DML语句

```java
package gz.ixfosa.statement;

import gz.itcast.util.JdbcUtil;

import java.sql.Connection;
import java.sql.DriverManager;
import java.sql.SQLException;
import java.sql.Statement;

import org.junit.Test;

/**
 * 使用Statement执行DML语句
 *
 */
public class Demo2 {
	private String url = "jdbc:mysql://localhost:3306/ixfosa";
	private String user = "root";
	private String password = "root";

	/**
	 * 增加
	 */
	@Test
	public void testInsert(){
		Connection conn = null;
		Statement stmt = null;
		try {
			//通过工具类获取连接对象
			conn = JdbcUtil.getConnection();
			
			//3.创建Statement对象
			stmt = conn.createStatement();
			
			//4.sql语句
			String sql = "INSERT INTO student(NAME,gender) VALUES('李四','女')";
			
			//5.执行sql
			int count = stmt.executeUpdate(sql);
			
			System.out.println("影响了"+count+"行");
			
		} catch (Exception e) {
			e.printStackTrace();
			throw new RuntimeException(e);
		} finally{
			//关闭资源
			/*if(stmt!=null)
				try {
					stmt.close();
				} catch (SQLException e) {
					e.printStackTrace();
					throw new RuntimeException(e);
				}
			if(conn!=null)
				try {
					conn.close();
				} catch (SQLException e) {
					e.printStackTrace();
					throw new RuntimeException(e);
				}*/
			JdbcUtil.close(conn, stmt);
		}
	}
	
	/**
	 * 修改
	 */
	@Test
	public void testUpdate(){
		Connection conn = null;
		Statement stmt = null;
		//模拟用户输入
		String name = "陈六";
		int id = 3;
		try {
			/*//1.注册驱动
			Class.forName("com.mysql.jdbc.Driver");
			
			//2.获取连接对象
			conn = DriverManager.getConnection(url, user, password);*/
			//通过工具类获取连接对象
			conn = JdbcUtil.getConnection();
			
			//3.创建Statement对象
			stmt = conn.createStatement();
			
			//4.sql语句
			String sql = "UPDATE student SET NAME='"+name+"' WHERE id="+id+"";
			
			System.out.println(sql);
			
			//5.执行sql
			int count = stmt.executeUpdate(sql);
			
			System.out.println("影响了"+count+"行");
			
		} catch (Exception e) {
			e.printStackTrace();
			throw new RuntimeException(e);
		} finally{
			//关闭资源
			/*if(stmt!=null)
				try {
					stmt.close();
				} catch (SQLException e) {
					e.printStackTrace();
					throw new RuntimeException(e);
				}
			if(conn!=null)
				try {
					conn.close();
				} catch (SQLException e) {
					e.printStackTrace();
					throw new RuntimeException(e);
				}*/
			JdbcUtil.close(conn, stmt);
		}
	}
	
	/**
	 * 删除
	 */
	@Test
	public void testDelete(){
		Connection conn = null;
		Statement stmt = null;
		//模拟用户输入
		int id = 3;
		try {
			/*//1.注册驱动
			Class.forName("com.mysql.jdbc.Driver");
			
			//2.获取连接对象
			conn = DriverManager.getConnection(url, user, password);*/
			//通过工具类获取连接对象
			conn = JdbcUtil.getConnection();
			
			//3.创建Statement对象
			stmt = conn.createStatement();
			
			//4.sql语句
			String sql = "DELETE FROM student WHERE id="+id+"";
			
			System.out.println(sql);
			
			//5.执行sql
			int count = stmt.executeUpdate(sql);
			
			System.out.println("影响了"+count+"行");
			
		} catch (Exception e) {
			e.printStackTrace();
			throw new RuntimeException(e);
		} finally{
			//关闭资源
			/*if(stmt!=null)
				try {
					stmt.close();
				} catch (SQLException e) {
					e.printStackTrace();
					throw new RuntimeException(e);
				}
			if(conn!=null)
				try {
					conn.close();
				} catch (SQLException e) {
					e.printStackTrace();
					throw new RuntimeException(e);
				}*/
			JdbcUtil.close(conn, stmt);
		}
	}
}
```



### 执行DQL语句

```java
package gz.ixfosa.statement;

import gz.itcast.util.JdbcUtil;

import java.sql.Connection;
import java.sql.ResultSet;
import java.sql.Statement;

import org.junit.Test;

/**
 * 使用Statement执行DQL语句（查询操作）
 *
 */
public class Demo3 {

	@Test
	public void test1(){
		Connection conn = null;
		Statement stmt = null;
		try{
			//获取连接
			conn = JdbcUtil.getConnection();
			//创建Statement
			stmt = conn.createStatement();
			//准备sql
			String sql = "SELECT * FROM student";
			//执行sql
			ResultSet rs = stmt.executeQuery(sql);
			
			//移动光标
			/*boolean flag = rs.next();
			
			flag = rs.next();
			flag = rs.next();
			if(flag){
				//取出列值
				//索引
				int id = rs.getInt(1);
				String name = rs.getString(2);
				String gender = rs.getString(3);
				System.out.println(id+","+name+","+gender);
				
				//列名称
				int id = rs.getInt("id");
				String name = rs.getString("name");
				String gender = rs.getString("gender");
				System.out.println(id+","+name+","+gender);
			}*/
			
			//遍历结果
			while(rs.next()){
				int id = rs.getInt("id");
				String name = rs.getString("name");
				String gender = rs.getString("gender");
				System.out.println(id+","+name+","+gender);
			}
			
		}catch(Exception e){
			e.printStackTrace();
			throw new RuntimeException(e);
		}finally{
			JdbcUtil.close(conn, stmt);
		}
	}
}
```



## PreparedStatement执行sql语句

PreparedStatement vs Statment

1. 语法不同：PreparedStatement可以使用预编译的sql，而Statment只能使用静态的sql
2. 效率不同： PreparedStatement可以使用sql缓存区，效率比Statment高
3. 安全性不同： PreparedStatement可以有效防止sql注入，而Statment不能防止sql注入。

推荐使用PreparedStatement

```java
package gz.ixfosa.prepared;

import gz.itcast.util.JdbcUtil;

import java.sql.Connection;
import java.sql.PreparedStatement;
import java.sql.ResultSet;

import org.junit.Test;
/**
 * PreparedStatement執行sql語句
 *
 */
public class Demo1 {

	/**
	 * 增加
	 */
	@Test
	public void testInsert() {
		Connection conn = null;
		PreparedStatement stmt = null;
		try {
			//1.获取连接
			conn = JdbcUtil.getConnection();
			
			//2.准备预编译的sql
			String sql = "INSERT INTO student(NAME,gender) VALUES(?,?)"; //?表示一个参数的占位符
			
			//3.执行预编译sql语句(检查语法)
			stmt = conn.prepareStatement(sql);
			
			//4.设置参数值
			/**
			 * 参数一： 参数位置  从1开始
			 */
			stmt.setString(1, "李四");
			stmt.setString(2, "男");
			
			//5.发送参数，执行sql
			int count = stmt.executeUpdate();
			
			System.out.println("影响了"+count+"行");
			
		} catch (Exception e) {
			e.printStackTrace();
			throw new RuntimeException(e);
		} finally {
			JdbcUtil.close(conn, stmt);
		}
	}
	
	/**
	 * 修改
	 */
	@Test
	public void testUpdate() {
		Connection conn = null;
		PreparedStatement stmt = null;
		try {
			//1.获取连接
			conn = JdbcUtil.getConnection();
			
			//2.准备预编译的sql
			String sql = "UPDATE student SET NAME=? WHERE id=?"; //?表示一个参数的占位符
			
			//3.执行预编译sql语句(检查语法)
			stmt = conn.prepareStatement(sql);
			
			//4.设置参数值
			/**
			 * 参数一： 参数位置  从1开始
			 */
			stmt.setString(1, "王五");
			stmt.setInt(2, 9);
			
			//5.发送参数，执行sql
			int count = stmt.executeUpdate();
			
			System.out.println("影响了"+count+"行");
			
		} catch (Exception e) {
			e.printStackTrace();
			throw new RuntimeException(e);
		} finally {
			JdbcUtil.close(conn, stmt);
		}
	}
	
	/**
	 * 删除
	 */
	@Test
	public void testDelete() {
		Connection conn = null;
		PreparedStatement stmt = null;
		try {
			//1.获取连接
			conn = JdbcUtil.getConnection();
			
			//2.准备预编译的sql
			String sql = "DELETE FROM student WHERE id=?"; //?表示一个参数的占位符
			
			//3.执行预编译sql语句(检查语法)
			stmt = conn.prepareStatement(sql);
			
			//4.设置参数值
			/**
			 * 参数一： 参数位置  从1开始
			 */
			stmt.setInt(1, 9);
			
			//5.发送参数，执行sql
			int count = stmt.executeUpdate();
			
			System.out.println("影响了"+count+"行");
			
		} catch (Exception e) {
			e.printStackTrace();
			throw new RuntimeException(e);
		} finally {
			JdbcUtil.close(conn, stmt);
		}
	}
	
	/**
	 * 查询
	 */
	@Test
	public void testQuery() {
		Connection conn = null;
		PreparedStatement stmt = null;
		ResultSet rs = null;
		try {
			//1.获取连接
			conn = JdbcUtil.getConnection();
			
			//2.准备预编译的sql
			String sql = "SELECT * FROM student"; 
			
			//3.预编译
			stmt = conn.prepareStatement(sql);
			
			//4.执行sql
			rs = stmt.executeQuery();
			
			//5.遍历rs
			while(rs.next()){
				int id = rs.getInt("id");
				String name = rs.getString("name");
				String gender = rs.getString("gender");
				System.out.println(id+","+name+","+gender);
			}
			
		} catch (Exception e) {
			e.printStackTrace();
			throw new RuntimeException(e);
		} finally {
			//关闭资源
			JdbcUtil.close(conn,stmt,rs);
		}
	}
}
```

```java
package gz.ixfosa.prepared;

import gz.itcast.util.JdbcUtil;

import java.sql.Connection;
import java.sql.PreparedStatement;
import java.sql.ResultSet;
import java.sql.Statement;

import org.junit.Test;

/**
 * 模拟用户登录效果
 *
 */
public class Demo2 {
	//模拟用户输入
	//private String name = "ericdfdfdfddfd' OR 1=1 -- ";
	private String name = "eric";
	//private String password = "123456dfdfddfdf";
	private String password = "123456";

	/**
	 * Statment存在sql被注入的风险
	 */
	@Test
	public void testByStatement(){
		Connection conn = null;
		Statement stmt = null;
		ResultSet rs = null;
		try {
			//获取连接
			conn = JdbcUtil.getConnection();
			
			//创建Statment
			stmt = conn.createStatement();
			
			//准备sql
			String sql = "SELECT * FROM users WHERE NAME='"+name+"' AND PASSWORD='"+password+"'";
			
			//执行sql
			rs = stmt.executeQuery(sql);
			
			if(rs.next()){
				//登录成功
				System.out.println("登录成功");
			}else{
				System.out.println("登录失败");
			}
			
		} catch (Exception e) {
			e.printStackTrace();
			throw new RuntimeException(e);
		} finally {
			JdbcUtil.close(conn, stmt ,rs);
		}
		
	}
	
	/**
	 * PreparedStatement可以有效地防止sql被注入
	 */
	@Test
	public void testByPreparedStatement(){
		Connection conn = null;
		PreparedStatement stmt = null;
		ResultSet rs = null;
		try {
			//获取连接
			conn = JdbcUtil.getConnection();
			
			String sql = "SELECT * FROM users WHERE NAME=? AND PASSWORD=?";
			
			//预编译
			stmt = conn.prepareStatement(sql);
			
			//设置参数
			stmt.setString(1, name);
			stmt.setString(2, password);
			
			//执行sql
			rs = stmt.executeQuery();
			
			if(rs.next()){
				//登录成功
				System.out.println("登录成功");
			}else{
				System.out.println("登录失败");
			}
			
		} catch (Exception e) {
			e.printStackTrace();
			throw new RuntimeException(e);
		} finally {
			JdbcUtil.close(conn, stmt ,rs);
		}
		
	}
}
```



##  CallableStatement执行存储过程

```java
package gz.ixfosa.callable;

import gz.itcast.util.JdbcUtil;

import java.sql.CallableStatement;
import java.sql.Connection;
import java.sql.ResultSet;

import org.junit.Test;

/**
 * 使用CablleStatement调用存储过程
 *
 */
public class Demo1 {

	/**
	 * 调用带有输入参数的存储过程
	 * CALL pro_findById(4);
	 */
	@Test
	public void test1(){
		Connection conn = null;
		CallableStatement stmt = null;
		ResultSet rs = null;
		try {
			//获取连接
			conn = JdbcUtil.getConnection();
			
			//准备sql
			String sql = "CALL pro_findById(?)"; //可以执行预编译的sql
			
			//预编译
			stmt = conn.prepareCall(sql);
			
			//设置输入参数
			stmt.setInt(1, 6);
			
			//发送参数
			rs = stmt.executeQuery(); //注意： 所有调用存储过程的sql语句都是使用executeQuery方法执行！！！
			
			//遍历结果
			while(rs.next()){
				int id = rs.getInt("id");
				String name = rs.getString("name");
				String gender = rs.getString("gender");
				System.out.println(id+","+name+","+gender);
			}
			
		} catch (Exception e) {
			e.printStackTrace();
			throw new RuntimeException(e);
		} finally {
			JdbcUtil.close(conn, stmt ,rs);
		}
	}
	
	/**
	 * 执行带有输出参数的存储过程
	 * CALL pro_findById2(5,@NAME);
	 */
	@Test
	public void test2(){
		Connection conn = null;
		CallableStatement stmt = null;
		ResultSet rs = null;
		try {
			//获取连接
			conn = JdbcUtil.getConnection();
			//准备sql
			String sql = "CALL pro_findById2(?,?)"; //第一个？是输入参数，第二个？是输出参数
			
			//预编译
			stmt = conn.prepareCall(sql);
			
			//设置输入参数
			stmt.setInt(1, 6);
			//设置输出参数(注册输出参数)
			/**
			 * 参数一： 参数位置
			 * 参数二： 存储过程中的输出参数的jdbc类型    VARCHAR(20)
			 */
			stmt.registerOutParameter(2, java.sql.Types.VARCHAR);
			
			//发送参数，执行
			stmt.executeQuery(); //结果不是返回到结果集中，而是返回到输出参数中
			
			//得到输出参数的值
			/**
			 * 索引值： 预编译sql中的输出参数的位置
			 */
			String result = stmt.getString(2); //getXX方法专门用于获取存储过程中的输出参数
			
			System.out.println(result);

		} catch (Exception e) {
			e.printStackTrace();
			throw new RuntimeException(e);
		} finally {
			JdbcUtil.close(conn, stmt ,rs);
		}
	}
}
```

## utils

```java
package gz.ixfosa.util;

import java.io.InputStream;
import java.sql.Connection;
import java.sql.DriverManager;
import java.sql.ResultSet;
import java.sql.SQLException;
import java.sql.Statement;
import java.util.Properties;

/**
 * jdbc工具类
 *
 */
public class JdbcUtil {
	private static String url = null;
	private static String user = null;
	private static String password = null;
	private static String driverClass = null;
	
	/**
	 * 静态代码块中（只加载一次）
	 */
	static{
		try {
			//读取db.properties文件
			Properties props = new Properties();
			/**
			 *  . 代表java命令运行的目录
			 *  在java项目下，. java命令的运行目录从项目的根目录开始
			 *  在web项目下，  . java命令的而运行目录从tomcat/bin目录开始
			 *  所以不能使用点.
			 */
			//FileInputStream in = new FileInputStream("./src/db.properties");
			
			/**
			 * 使用类路径的读取方式
			 *  / : 斜杠表示classpath的根目录
			 *     在java项目下，classpath的根目录从bin目录开始
			 *     在web项目下，classpath的根目录从WEB-INF/classes目录开始
			 */
			InputStream in = JdbcUtil.class.getResourceAsStream("/db.properties");
			
			//加载文件
			props.load(in);
			//读取信息
			url = props.getProperty("url");
			user = props.getProperty("user");
			password = props.getProperty("password");
			driverClass = props.getProperty("driverClass");
			
			
			//注册驱动程序
			Class.forName(driverClass);
		} catch (Exception e) {
			e.printStackTrace();
			System.out.println("驱程程序注册出错");
		}
	}

	/**
	 * 抽取获取连接对象的方法
	 */
	public static Connection getConnection(){
		try {
			Connection conn = DriverManager.getConnection(url, user, password);
			return conn;
		} catch (SQLException e) {
			e.printStackTrace();
			throw new RuntimeException(e);
		}
	}
	
	
	/**
	 * 释放资源的方法
	 */
	public static void close(Connection conn,Statement stmt){
		if(stmt!=null){
			try {
				stmt.close();
			} catch (SQLException e) {
				e.printStackTrace();
				throw new RuntimeException(e);
			}
		}
		if(conn!=null){
			try {
				conn.close();
			} catch (SQLException e) {
				e.printStackTrace();
				throw new RuntimeException(e);
			}
		}
	}
	
	public static void close(Connection conn,Statement stmt,ResultSet rs){
		if(rs!=null)
			try {
				rs.close();
			} catch (SQLException e1) {
				e1.printStackTrace();
				throw new RuntimeException(e1);
			}
		if(stmt!=null){
			try {
				stmt.close();
			} catch (SQLException e) {
				e.printStackTrace();
				throw new RuntimeException(e);
			}
		}
		if(conn!=null){
			try {
				conn.close();
			} catch (SQLException e) {
				e.printStackTrace();
				throw new RuntimeException(e);
			}
		}
	}
}
```

src/db.properties

```properties
url=jdbc:mysql://localhost:3306/ixfosa
user=root
password=root
driverClass=com.mysql.jdbc.Driver
```

```java
package cn.ixfosa.utils;

import java.sql.Connection;
import java.sql.DriverManager;
import java.sql.ResultSet;
import java.sql.SQLException;
import java.sql.Statement;

/**
 * 1. 返回连接 2. 关闭
 * 
 */
public class JdbcUtil {

	// 连接参数
	// private String url = "jdbc:mysql://localhost:3306/jdbc_demo";
	private static String url = "jdbc:mysql:///jdbc_demo";
	private static String user = "root";
	private static String password = "root";

	/**
	 * 返回连接对象
	 */
	public static Connection getConnection() {
		try {
			Class.forName("com.mysql.jdbc.Driver");
			return DriverManager.getConnection(url, user, password);
		} catch (Exception e) {
			throw new RuntimeException(e);
		}
	}

	/**
	 * 关闭
	 */
	public static void closeAll(Connection con, Statement stmt, ResultSet rs) {
		try {
			if (rs != null) {
				rs.close();  // 快速异常捕获 Alt + shift + z 
				rs = null;   // 建议垃圾回收期回收资源
			}
			if (stmt != null) {
				stmt.close();
				stmt = null;
			}
			if (con != null && !con.isClosed()) {
				con.close();
				con = null;
			}
		} catch (SQLException e) {
			throw new RuntimeException(e);
		}
	}
}
```



## 预编译sql处理(防止sql注入)

```sql
-- 创建数据库
CREATE DATABASE jdbc_demo DEFAULT CHARACTER SET utf8;i
-- 创建表
USE jdbc_demo;
CREATE TABLE admin(
    id INT PRIMARY KEY AUTO_INCREMENT,
    userName VARCHAR(20),
    pwd VARCHAR(20)
)
```

```java
|--Statement      执行SQL命令
	|-- CallableStatement,     执行存储过程
	|-- PreparedStatement    预编译SQL语句执行
```

使用预编译SQL语句的命令对象，好处：

1. 避免了频繁sql拼接 (可以使用占位符)
2. 可以防止sql注入

```java
package cn.ixfosa.pstmt;

import java.sql.Connection;
import java.sql.DriverManager;
import java.sql.ResultSet;
import java.sql.SQLException;
import java.sql.Statement;

import org.junit.Test;

import cn.ixfosa.utils.JdbcUtil;

import java.sql.PreparedStatement;

public class App {
	
	// 全局参数
	private Connection con;
	private Statement stmt;
	private PreparedStatement pstmt;
	private ResultSet rs;
	

	// 1. 没有使用防止sql注入的案例
	@Test
	public void testLogin() {
		
		// 1.0 模拟登陆的用户名，密码
		String userName = "tom";
		//String pwd = "8881";
		String pwd = " ' or 1=1 -- ";
		
		// SQL语句
		String sql = "select * from admin where userName='"+userName+"'  and pwd='"+pwd+"' ";
		System.out.println(sql);
		try {
			// 1.1 加载驱动，创建连接
			con = JdbcUtil.getConnection();
			// 1.2 创建stmt对象
			stmt = con.createStatement();
			// 1.3 执行查询
			rs = stmt.executeQuery(sql);
			// 业务判断
			if (rs.next()) {
				System.out.println("登陆成功, 编号：" + rs.getInt("id"));
			}
		} catch (Exception e) {
			e.printStackTrace();
		} finally {
			// 1.4 关闭
			try {
				rs.close();
				stmt.close();
				con.close();
			} catch (Exception e) {
				e.printStackTrace();
			}
		}
	}
	
	
	// 2. 使用PreparedStatement, 防止sql注入
	@Test
	public void testLogin2() {
		
		// 1.0 模拟登陆的用户名，密码
		String userName = "tom";
		//String pwd = "8881";
		String pwd = " ' or 1=1 -- ";
		
		// SQL语句
		String sql = "select * from admin where userName=?  and pwd=? ";
		try {
			// 1.1 加载驱动，创建连接
			con = JdbcUtil.getConnection();
			// 1.2 创建pstmt对象
			pstmt = con.prepareStatement(sql);   // 对sql语句预编译
			// 设置占位符值
			pstmt.setString(1, userName);
			pstmt.setString(2, pwd);
			
			// 1.3 执行
			rs = pstmt.executeQuery();
			if (rs.next()) {
				System.out.println("登陆成功，" + rs.getInt("id"));
			}
			
		} catch (Exception e) {
			e.printStackTrace();
		} finally {
			// 1.4 关闭
			try {
				rs.close();
				pstmt.close();
				con.close();
			} catch (Exception e) {
				e.printStackTrace();
			}
		}
	}
}
```

## 批处理

需求：批量保存信息！  
设计：
	

```java
//AdminDao
Public  void  save(List<Admin list){    // 目前用这种方式
	// 循环
	// 保存  (批量保存)
}

Public  void  save(Admin  admin ){
	// 循环
	// 保存
}
```

技术：

```java
|-- Statement
	批处理相关方法
	void addBatch(String sql)     添加批处理
	void clearBatch()             清空批处理
	int[] executeBatch()          执行批处理
```

实现：
	Admin.java         实体类封装数据
	AdminDao.java      封装所有的与数据库的操作
	App.java      测试

```java
package cn.ixfosa.batch;

//实体类封装数据
public class Admin {

	private String userName;
	private String pwd;
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
package cn.ixfosa.batch;

import java.sql.Connection;
import java.sql.PreparedStatement;
import java.sql.ResultSet;
import java.sql.Statement;
import java.util.List;

import cn.ixfosa.utils.JdbcUtil;

// 封装所有的与数据库的操作
public class AdminDao {
	
	// 全局参数
	private Connection con;
	private PreparedStatement pstmt;
	private ResultSet rs;

	// 批量保存管理员
	public void save(List<Admin> list) {
		// SQL
		String sql = "INSERT INTO admin(userName,pwd) values(?,?)";
		
		try {
			
			// 获取连接
			con = JdbcUtil.getConnection();
			// 创建stmt 
			pstmt = con.prepareStatement(sql);   		// 【预编译SQL语句】
			
			for (int i=0; i<list.size(); i++) {
				Admin admin = list.get(i);
				// 设置参数
				pstmt.setString(1, admin.getUserName());
				pstmt.setString(2, admin.getPwd());
				
				// 添加批处理
				pstmt.addBatch();						// 【不需要传入SQL】
				
				// 测试：每5条执行一次批处理
				if (i % 5 == 0) {
					// 批量执行 
					pstmt.executeBatch();
					// 清空批处理
					pstmt.clearBatch();
				}
				
			}
			
			// 批量执行 
			pstmt.executeBatch();
			// 清空批处理
			pstmt.clearBatch();
			
		} catch (Exception e) {
			e.printStackTrace();
		} finally {
			JdbcUtil.closeAll(con, pstmt, rs);
		}
	}
}

```

```java
package cn.ixfosa.batch;

import java.util.ArrayList;
import java.util.List;

import org.junit.Test;

public class App {

	// 测试批处理操作
	@Test
	public void testBatch() throws Exception {
		
		// 模拟数据
		List<Admin> list = new ArrayList<Admin>();
		for (int i=1; i<21; i++) {
			Admin admin = new Admin();
			admin.setUserName("Jack" + i);
			admin.setPwd("888" + i);
			list.add(admin);
		}
		
		// 保存
		AdminDao dao = new AdminDao();
		dao.save(list);
	}
}
```

## 插入数据，获取自增长值

部门与员工

​	一对多的关系

### 设计数据库

​	员工表 （外键表） 【员工表有一个外键字段，引用了部门表的主键】

​	部门表（主键表）

```sql
部门
CREATE TABLE dept(
   deptId INT PRIMARY KEY AUTO_INCREMENT,
   deptName VARCHAR(20)
);

-- 员工
CREATE TABLE employee(
   empId INT PRIMARY KEY AUTO_INCREMENT,
   empName VARCHAR(20),
   dept_id  INT   --  外键字段   
);
-- 给员工表添加外键约束
ALTER TABLE employee ADD CONSTRAINT FK_employee_dept_deptId
	FOREIGN KEY(dept_id) REFERENCES dept(deptId) ;


```

编码总体思路

​	保存员工及其对应的部门！

​	步骤：

1. 先保存部门

2. 再得到部门主键，再保存员

开发具体步骤：

1. 设计javabean

2. 设计dao

3. 测试

### 设计javabean

```java
package cn.itcast.auto;

public class Dept {

	private int id;
	private String deptName;
	public int getId() {
		return id;
	}
	public void setId(int id) {
		this.id = id;
	}
	public String getDeptName() {
		return deptName;
	}
	public void setDeptName(String deptName) {
		this.deptName = deptName;
	}
}
```

```java
package cn.ixfosa.auto;

public class Employee {

	private int empId;
	private String empName;
	// 关联的部门
	private Dept dept;
	
	
	public Dept getDept() {
		return dept;
	}
	public void setDept(Dept dept) {
		this.dept = dept;
	}
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
}
```



### 设计dao

```java
package cn.ixfosa.auto;

import java.sql.Connection;
import java.sql.PreparedStatement;
import java.sql.ResultSet;
import java.sql.SQLException;
import java.sql.Statement;

import cn.ixfosa.utils.JdbcUtil;


public class EmpDao {
	
	private Connection con;
	private PreparedStatement pstmt;
	private ResultSet rs;

	// 保存员工，同时保存关联的部门
	public void save(Employee emp){
		
		// 保存部门
		String sql_dept = "insert into dept(deptName) values(?)";
		// 保存员工
		String sql_emp = "INSERT INTO employee (empName,dept_id) VALUES (?,?)";
		// 部门id
		int deptId = 0;
		
		try {
			// 连接
			con = JdbcUtil.getConnection();
			
			/*****保存部门，获取自增长*******/
			// 【一、需要指定返回自增长标记】
			pstmt = con.prepareStatement(sql_dept,Statement.RETURN_GENERATED_KEYS);
			// 设置参数
			pstmt.setString(1, emp.getDept().getDeptName());
			// 执行
			pstmt.executeUpdate();
			
			// 【二、获取上面保存的部门子增长的主键】
			rs =  pstmt.getGeneratedKeys();
			// 得到返回的自增长字段
			if (rs.next()) {
				deptId = rs.getInt(1);
			}
			
			/*****保存员工*********/
			pstmt = con.prepareStatement(sql_emp);
			// 设置参数
			pstmt.setString(1, emp.getEmpName());
			pstmt.setInt(2, deptId);
			pstmt.executeUpdate();
			
		} catch (Exception e) {
			e.printStackTrace();
		} finally {
			JdbcUtil.closeAll(con, pstmt, rs);
		}
	}
}
```



### 测试

```java
package cn.ixfosa.auto;

import org.junit.Test;

public class App {

	// 保存员工
	@Test
	public void testSave() throws Exception {
		// 模拟数据
		Dept d = new Dept();
		d.setDeptName("应用开发部");
		Employee emp = new Employee();
		emp.setEmpName("ixfosa");
		emp.setDept(d);   // 关联
		
		// 调用dao保存
		EmpDao empDao = new EmpDao();
		empDao.save(emp);
		
	}
}
```



## 事务

### 基本概念

事务使指一组`最小逻辑操作单元`，里面有多个操作组成。 组成事务的每一部分必须要`同时提交成功`，如果有一个操作失败，整个操作就`回滚`。

### 事务ACID特性

原子性，是一个最小逻辑操作单元 !
一致性，事务过程中，数据处于一致状态。
持久性， 事务一旦提交成功，对数据的更改会反映到数据库中。
隔离性， 事务与事务之间是隔离的

#### 原子性（Atomicity）

原子性是指事务是一个不可分割的工作单位，事务中的操作要么都发生，要么都不发生。 

#### 一致性（Consistency）

事务必须使数据库从一个一致性状态变换到另外一个一致性状态。

#### 隔离性（Isolation）

事务的隔离性是多个用户并发访问数据库时，数据库为每一个用户开启的事务，不能被其他事务的操作数据所干扰，多个并发事务之间要相互隔离。

#### 持久性（Durability）

持久性是指一个事务一旦被提交，它对数据库中数据的改变就是永久性的，接下来即使数据库发生故障也不应该对其有任何影响



### 案例

需求： 张三给李四转账
设计： 账户表

```sql
-- 账户表
CREATE TABLE account(
   id INT PRIMARY KEY AUTO_INCREMENT,
   accountName VARCHAR(20),
   money DOUBLE
);

-- 转账
UPDATE account SET money=money-1000 WHERE accountName='张三';
UPDATE account SET money=money+1000 WHERE accountName='李四';
```

技术：

```java
|-- Connection
	void setAutoCommit(boolean autoCommit) ;  设置事务是否自动提交
								              如果设置为false，表示手动提交事务。
	void commit() ();						  手动提交事务
	void rollback() ;						  回滚（出现异常时候，所有已经执行成功的代码需要回												退到事务开始前的状态。）
	Savepoint setSavepoint(String name) 
```



```java
package cn.ixfosa.tx;

import java.sql.Connection;
import java.sql.PreparedStatement;
import java.sql.SQLException;
import java.sql.Savepoint;

import cn.ixfosa.utils.JdbcUtil;

public class AccountDao {

	// 全局参数
	private Connection con;
	private PreparedStatement pstmt;

	// 1. 转账，没有使用事务
	public void trans1() {

		String sql_zs = "UPDATE account SET money=money-1000 WHERE accountName='张三';";
        
		String sql_ls = "UPDATE account SET money=money+1000 WHERE accountName='李四';";

		try {
			con = JdbcUtil.getConnection(); // 默认开启的隐士事务
			con.setAutoCommit(true);

			/*** 第一次执行SQL ***/
			pstmt = con.prepareStatement(sql_zs);
			pstmt.executeUpdate();

			/*** 第二次执行SQL ***/
			pstmt = con.prepareStatement(sql_ls);
			pstmt.executeUpdate();

		} catch (Exception e) {
			e.printStackTrace();
		} finally {
			JdbcUtil.closeAll(con, pstmt, null);
		}

	}

	// 2. 转账，使用事务
	public void trans2() {

		String sql_zs = "UPDATE account SET money=money-1000 WHERE accountName='张三';";
		String sql_ls = "UPDATE1 account SET money=money+1000 WHERE accountName='李四';";

		try {
			con = JdbcUtil.getConnection(); // 默认开启的隐士事务
			// 一、设置事务为手动提交
			con.setAutoCommit(false);

			/*** 第一次执行SQL ***/
			pstmt = con.prepareStatement(sql_zs);
			pstmt.executeUpdate();

			/*** 第二次执行SQL ***/
			pstmt = con.prepareStatement(sql_ls);
			pstmt.executeUpdate();

		} catch (Exception e) {
			try {
				// 二、 出现异常，需要回滚事务
				con.rollback();
			} catch (SQLException e1) {
			}
			e.printStackTrace();
		} finally {
			try {
				// 三、所有的操作执行成功, 提交事务
				con.commit();
				JdbcUtil.closeAll(con, pstmt, null);
			} catch (SQLException e) {
			}
		}

	}

	// 3. 转账，使用事务， 回滚到指定的代码段
	public void trans() {
		// 定义个标记
		Savepoint sp = null;
		
		// 第一次转账
		String sql_zs1 = "UPDATE account SET money=money-1000 WHERE accountName='张三';";
		String sql_ls1 = "UPDATE account SET money=money+1000 WHERE accountName='李四';";
		
		// 第二次转账
		String sql_zs2 = "UPDATE account SET money=money-500 WHERE accountName='张三';";
		String sql_ls2 = "UPDATE1 account SET money=money+500 WHERE accountName='李四';";

		try {
			con = JdbcUtil.getConnection(); // 默认开启的隐士事务
			con.setAutoCommit(false);       // 设置事务手动提交

			/*** 第一次转账 ***/
			pstmt = con.prepareStatement(sql_zs1);
			pstmt.executeUpdate();
			pstmt = con.prepareStatement(sql_ls1);
			pstmt.executeUpdate();
			
			// 回滚到这个位置？
			sp = con.setSavepoint(); 
			
			
			/*** 第二次转账 ***/
			pstmt = con.prepareStatement(sql_zs2);
			pstmt.executeUpdate();
			pstmt = con.prepareStatement(sql_ls2);
			pstmt.executeUpdate();
			

		} catch (Exception e) {
			try {
				// 回滚 (回滚到指定的代码段)
				con.rollback(sp);
			} catch (SQLException e1) {
			}
			e.printStackTrace();
		} finally {
			try {
				// 提交
				con.commit();
			} catch (SQLException e) {
			}
			JdbcUtil.closeAll(con, pstmt, null);
		}

	}
}
```

```java
package cn.ixfosa.tx;

import org.junit.Test;

public class App {

	@Test
	public void testname() throws Exception {
		
		// 转账
		AccountDao accountDao = new AccountDao();
		accountDao.trans();
	}
}
```



## 大文本, 二进制类型处理

Oracle中大文本数据类型，
		Clob    长文本类型   （MySQL中不支持，使用的是text）
		Blob    二进制类型

MySQL数据库，
		Text    长文本类型
		Blob    二进制类型

需求： jdbc中操作长文本数据。

编码：
		保存大文本数据类型
		读取大文本数据类型

​		保存二进制数据
​		读取二进制数据



### 设计表

```sql
-- 测试大数据类型
CREATE TABLE test(
     id INT PRIMARY KEY AUTO_INCREMENT,
     content LONGTEXT,
     img LONGBLOB
);
```



### 长文本类型

```java
package cn.ixfosa.longtext;

import java.io.File;
import java.io.FileReader;
import java.io.Reader;
import java.sql.Connection;
import java.sql.PreparedStatement;
import java.sql.ResultSet;
import java.sql.Statement;

import org.junit.Test;

import cn.ixfosa.utils.JdbcUtil;

public class App_text {
	
	// 全局参数
	private Connection con;
	private Statement stmt;
	private PreparedStatement pstmt;
	private ResultSet rs;
	

	@Test
	// 1. 保存大文本数据类型   ( 写longtext)
	public void testSaveText() {
		String sql = "insert into test(content) values(?)";
		try {
			// 连接
			con = JdbcUtil.getConnection();
			// pstmt 对象
			pstmt = con.prepareStatement(sql);
			// 设置参数
			// 先获取文件路径
            //src/cn/ixfosa/longtext/tips.txt
			String path = App_text.class.getResource("tips.txt").getPath();
			FileReader reader = new FileReader(new File(path));
			pstmt.setCharacterStream(1, reader);
			
			// 执行sql
			pstmt.executeUpdate();
			
			// 关闭
			reader.close();
		} catch (Exception e) {
			e.printStackTrace();
		} finally {
			JdbcUtil.closeAll(con, pstmt, null);
		}
	}
	
	@Test
	// 2. 读取大文本数据类型   ( 读longtext)
	public void testGetAsText() {
		String sql = "select * from  test;";
		try {
			// 连接
			con = JdbcUtil.getConnection();
			// pstmt 对象
			pstmt = con.prepareStatement(sql);
			// 读取
			rs = pstmt.executeQuery();
			if (rs.next()) {
				// 获取长文本数据, 方式1:
				//Reader r = rs.getCharacterStream("content");
				
				// 获取长文本数据, 方式2:
				System.out.print(rs.getString("content"));
			}
			
		} catch (Exception e) {
			e.printStackTrace();
		} finally {
			JdbcUtil.closeAll(con, pstmt, null);
		}
	}
}
```



### 二进制类型

```java
package cn.ixfosa.longtext;

import java.io.File;
import java.io.FileOutputStream;
import java.io.FileReader;
import java.io.InputStream;
import java.io.Reader;
import java.sql.Connection;
import java.sql.PreparedStatement;
import java.sql.ResultSet;
import java.sql.Statement;

import org.junit.Test;

import cn.ixfosa.utils.JdbcUtil;

public class App_blob {
	
	// 全局参数
	private Connection con;
	private Statement stmt;
	private PreparedStatement pstmt;
	private ResultSet rs;
	

	@Test
	// 1. 二进制数据类型   ( 写longblob)
	public void testSaveText() {
		String sql = "insert into test(img) values(?)";
		try {
			// 连接
			con = JdbcUtil.getConnection();
			// pstmt 对象
			pstmt = con.prepareStatement(sql);
			// 获取图片流   src/cn/ixfosa/longtext/7.jpg
			InputStream in = App_text.class.getResourceAsStream("7.jpg");
			pstmt.setBinaryStream(1, in);
			
			// 执行保存图片
			pstmt.execute();
			
			// 关闭
			in.close();
			
		} catch (Exception e) {
			e.printStackTrace();
		} finally {
			JdbcUtil.closeAll(con, pstmt, null);
		}
	}
	
	@Test
	// 2. 读取大文本数据类型   ( 读longblob)
	public void testGetAsText() {
		String sql = "select img from  test where id=2;";
		try {
			// 连接
			con = JdbcUtil.getConnection();
			// pstmt 对象
			pstmt = con.prepareStatement(sql);
			// 读取
			rs = pstmt.executeQuery();
			if (rs.next()) {
				// 获取图片流
				InputStream in = rs.getBinaryStream("img");
				// 图片输出流
				FileOutputStream out = new FileOutputStream(new File("c://1.jpg"));
				int len = -1;
				byte b[] = new byte[1024];
				while ((len = in.read(b)) != -1) {
					out.write(b, 0, len);
				}
				// 关闭
				out.close();
				in.close();
			}
			
		} catch (Exception e) {
			e.printStackTrace();
		} finally {
			JdbcUtil.closeAll(con, pstmt, null);
		}
	}
}
```

## 综合案例

需求分析：
		登陆、注册、注销；
		登陆成功，
			显示所有的员工
设计
		数据库设计：
			Admin,  存放所有的登陆用户
			Employee, 存放所有的员工信息
系统设计
		a.系统结构
			分层： 基于mvc模式的分层
		b.项目用到的公用组件、类 (了解)

![综合案例](E:\smile\java\images\综合案例.jpg)



### utils

```java
package cn.ixfosa.utils;

import java.sql.Connection;
import java.sql.DriverManager;
import java.sql.ResultSet;
import java.sql.SQLException;
import java.sql.Statement;

/**
 * 1. 返回连接 2. 关闭
 * 
 */
public class JdbcUtil {

	// 连接参数
	// private String url = "jdbc:mysql://localhost:3306/jdbc_demo";
	private static String url = "jdbc:mysql:///jdbc_demo";
	private static String user = "root";
	private static String password = "root";

	/**
	 * 返回连接对象
	 */
	public static Connection getConnection() {
		try {
			Class.forName("com.mysql.jdbc.Driver");
			return DriverManager.getConnection(url, user, password);
		} catch (Exception e) {
			throw new RuntimeException(e);
		}
	}

	/**
	 * 关闭
	 */
	public static void closeAll(Connection con, Statement stmt, ResultSet rs) {
		try {
			if (rs != null) {
				rs.close();  // 快速异常捕获 Alt + shift + z 
				rs = null;   // 建议垃圾回收期回收资源
			}
			if (stmt != null) {
				stmt.close();
				stmt = null;
			}
			if (con != null && !con.isClosed()) {
				con.close();
				con = null;
			}
		} catch (SQLException e) {
			throw new RuntimeException(e);
		}
	}
}

```

### exception

```java
package cn.ixfosa.exception;

public class UserExistsException extends Exception {

	public UserExistsException() {
		// TODO Auto-generated constructor stub
	}

	public UserExistsException(String message) {
		super(message);
		// TODO Auto-generated constructor stub
	}

	public UserExistsException(Throwable cause) {
		super(cause);
		// TODO Auto-generated constructor stub
	}

	public UserExistsException(String message, Throwable cause) {
		super(message, cause);
		// TODO Auto-generated constructor stub
	}

}
```



### entity

```java
package cn.ixfosa.entity;

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
}
```

### dao

```java
package cn.ixfosa.dao;

import cn.ixfosa.entity.Admin;

/**
 * 2. 数据访问层, 接口设计
 *
 */
public interface IAdminDao {

	/**
	 * 保存
	 * @param admin
	 */
	void save(Admin admin);
	
	/**
	 * 根据用户名密码查询
	 */
	Admin findByNameAndPwd(Admin admin);
	
	/**
	 * 检查用户名是否存在
	 * @param name   要检查的用户名
	 * @return  true表示用户名已经存在； 否则用户名不存在
	 */
	boolean userExists(String name);
}

```

```java
package cn.ixfosa.dao.impl;

import java.sql.Connection;
import java.sql.PreparedStatement;
import java.sql.ResultSet;
import java.sql.SQLException;

import cn.ixfosa.dao.IAdminDao;
import cn.ixfosa.entity.Admin;
import cn.ixfosa.utils.JdbcUtil;

/**
 * 2. 数据访问层接口的实现类
 *
 */
public class AdminDao implements IAdminDao {
	
	private Connection con;
	private PreparedStatement pstmt;
	private ResultSet rs;
	

	@Override
	public Admin findByNameAndPwd(Admin admin) {
		String sql = "select * from admin where userName=? and pwd=?";
		Admin ad = null;
		try {
			con = JdbcUtil.getConnection();
			pstmt = con.prepareStatement(sql);
			pstmt.setString(1, admin.getUserName());
			pstmt.setString(2, admin.getPwd());
			// 执行
			rs = pstmt.executeQuery();
			// 遍历
			if (rs.next()) {
				ad = new Admin();
				ad.setId(rs.getInt("id"));
				ad.setUserName(rs.getString("userName"));
				ad.setPwd(rs.getString("pwd"));
			}
			return ad;
		} catch (Exception e) {
			throw new RuntimeException(e);
		} finally {
			JdbcUtil.closeAll(con, pstmt, null);
		}
	}

	@Override
	public void save(Admin admin) {
		String sql = "INSERT INTO admin(userName,pwd) VALUES(?,?);";
		try {
			con = JdbcUtil.getConnection();
			pstmt = con.prepareStatement(sql);
			// 设置参数
			pstmt.setString(1, admin.getUserName());
			pstmt.setString(2, admin.getPwd());
			// 执行更新
			pstmt.executeUpdate();		
		} catch (Exception e) {
			throw new RuntimeException(e);
		} finally {
			JdbcUtil.closeAll(con, pstmt, null);
		}
		
	}

	@Override
	public boolean userExists(String name) {
		String sql = "select id from admin where userName=?";
		try {
			con = JdbcUtil.getConnection();
			pstmt = con.prepareStatement(sql);
			// 设置参数
			pstmt.setString(1, name);
			// 执行查询
			rs = pstmt.executeQuery();
			// 判断
			if (rs.next()) {
				int id = rs.getInt("id");
				if (id > 0) {
					// 用户名已经存在
					return true;
				}
			}
			return false;
		} catch (SQLException e) {
			throw new RuntimeException(e);
		} finally {
			JdbcUtil.closeAll(con, pstmt, rs);
		}
	}
}
```

### service

```java
package cn.ixfosa.service;

import cn.ixfosa.entity.Admin;
import cn.ixfosa.exception.UserExistsException;

/**
 * 业务逻辑层接口设计
 * @author Jie.Yuan
 *
 */
public interface IAdminService {

	/**
	 * 注册
	 */
	void register(Admin admin) throws UserExistsException;
	
	/**
	 * 登陆
	 */
	Admin login(Admin admin);
}
```

```java
package cn.ixfosa.service.impl;

import cn.itcast.dao.IAdminDao;
import cn.itcast.dao.impl.AdminDao;
import cn.itcast.entity.Admin;
import cn.itcast.exception.UserExistsException;
import cn.itcast.service.IAdminService;

/**
 * 3. 业务逻辑层实现
 *
 */
public class AdminService implements IAdminService {

	// 调用的dao
	private IAdminDao adminDao = new AdminDao();
	
	@Override
	public Admin login(Admin admin) {
		try {
			return adminDao.findByNameAndPwd(admin);
		} catch (Exception e) {
			throw new RuntimeException(e);
		}
	}

	@Override
	public void register(Admin admin) throws UserExistsException  {
		
		try {
			// 1. 先根据用户名查询用户是否存在
			boolean flag = adminDao.userExists(admin.getUserName());
			
			// 2. 如果用户存在，不允许注册
			if (flag){
				// 不允许注册, 给调用者提示
				throw new UserExistsException("用户名已经存在，注册失败！");
			}
			
			// 3. 用户不存在，才可以注册
			adminDao.save(admin);
		
		} catch (UserExistsException e) {
			throw e;
		} catch (Exception e) {
			throw new RuntimeException(e);
		}
	}
}
```



```java
package cn.ixfosa.servlet;

import java.io.IOException;
import java.io.PrintWriter;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import cn.ixfosa.entity.Admin;
import cn.ixfosa.exception.UserExistsException;
import cn.ixfosa.service.IAdminService;
import cn.ixfosa.service.impl.AdminService;

/**
 * 4. 控制层开发
 * 
 */
public class AdminServlet extends HttpServlet {
	
	// 调用的service
	private IAdminService adminService = new AdminService();
	
	public void doPost(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, IOException {
		this.doGet(request, response);
	}

	public void doGet(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, IOException {

		// 获取操作类型
		String method = request.getParameter("method");
		if ("register".equals(method)) {
			register(request,response);
		}
	}

	/**
	 * 注册处理方法
	 * @param request
	 * @param response
	 * @throws IOException 
	 * @throws ServletException 
	 */
    
	private void register(HttpServletRequest request,
			HttpServletResponse response) throws IOException, ServletException {
		
		//1. 获取请求参数
		String userName = request.getParameter("userName");
		String pwd = request.getParameter("pwd");
		// 封装
		Admin admin = new Admin();
		admin.setUserName(userName);
		admin.setPwd(pwd);
		
		//2. 调用Service处理注册的业务逻辑
		try {
			adminService.register(admin);
			
			// 注册成功，跳转到首页
			request.getRequestDispatcher("/index.jsp").forward(request, response);
			
		} catch (UserExistsException e) {
			// 用户名存在，注册失败(跳转到注册页面)
			request.setAttribute("message", "用户名已经存在");
			// 转发
			request.getRequestDispatcher("/register.jsp").forward(request, response);
		} catch (Exception e) {
			e.printStackTrace();  // 测试时候用
			// 其他错误, 跳转到错误页面
			response.sendRedirect(request.getContextPath() + "/error/error.jsp");
		}
	}
}
```

### web.xml

```xml
<?xml version="1.0" encoding="UTF-8"?>
<web-app version="2.5" 
	xmlns="http://java.sun.com/xml/ns/javaee" 
	xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" 
	xsi:schemaLocation="http://java.sun.com/xml/ns/javaee 
	http://java.sun.com/xml/ns/javaee/web-app_2_5.xsd">
  <servlet>
    <servlet-name>AdminServlet</servlet-name>
    <servlet-class>cn.ixfosa.servlet.AdminServlet</servlet-class>
  </servlet>

  <servlet-mapping>
    <servlet-name>AdminServlet</servlet-name>
    <url-pattern>/admin</url-pattern>
  </servlet-mapping>
  <welcome-file-list>
    <welcome-file>index.jsp</welcome-file>
  </welcome-file-list>
</web-app>
```



### jsp

WebRoot/error/error.jsp

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
  	 系统忙，稍后再试！
  </body>
</html>
```

WebRoot/index.jsp

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
  	success!
  </body>
</html>
```

WebRoot/register.jsp

```jsp
<%@ page language="java" import="java.util.*" pageEncoding="UTF-8"%>

<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN">
<html>
  <head>
    
    <title>注册</title>
	<meta http-equiv="pragma" content="no-cache">
	<meta http-equiv="cache-control" content="no-cache">
	<meta http-equiv="expires" content="0">    
  </head>
  
  <body>
  	<form name="frm1" action="${pageContext.request.contextPath }/admin?method=register" method="post" >
  		<table>
  			<tr>
  				<td>用户名</td>
  				<td>
  					<input type="text" name="userName"/>
  					${requestScope.message } <!-- 如果用户名存在注册失败，给用户提示 -->
  				</td>
  			</tr>
  			<tr>
  				<td>密码</td>
  				<td><input type="password" name="pwd"/></td>
  			</tr>
  			<tr>
  				<td colspan="2">
  					<input type="submit" value="亲，点我注册！">
  				</td>
  			</tr>
  		</table>
  	
  	</form>
  </body>
</html>
```



## Dao操作的抽取,  BaseDao

### entity

```java
package cn.ixfosa.cdbc;

import java.util.Date;

/**
 * 1. 实体类设计
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



### BaseDao

```java
package cn.ixfosa.jdbc;

import java.sql.Connection;
import java.sql.PreparedStatement;
import java.sql.ResultSet;
import java.sql.ResultSetMetaData;
import java.util.ArrayList;
import java.util.List;

import org.apache.commons.beanutils.BeanUtils;

import cn.ixfosa.utils.JdbcUtil;

/**
 * 通用的dao，自己写的所有的dao都继承此类;
 * 此类定义了2个通用的方法：
 * 	1. 更新
 *  2. 查询
 *
 */
public class BaseDao {
	
	// 初始化参数
	private Connection con;
	private PreparedStatement pstmt;
	private ResultSet rs;

	/**
	 * 更新的通用方法
	 * @param sql   更新的sql语句(update/insert/delete)
	 * @param paramsValue  sql语句中占位符对应的值(如果没有占位符，传入null)
	 */
	public void update(String sql,Object[] paramsValue){
		
		try {
			// 获取连接
			con = JdbcUtil.getConnection();
			// 创建执行命令的stmt对象
			pstmt = con.prepareStatement(sql);
			// 参数元数据： 得到占位符参数的个数
			int count = pstmt.getParameterMetaData().getParameterCount();
			
			// 设置占位符参数的值
			if (paramsValue != null && paramsValue.length > 0) {
				// 循环给参数赋值
				for(int i=0;i<count;i++) {
					pstmt.setObject(i+1, paramsValue[i]);
				}
			}
			// 执行更新
			pstmt.executeUpdate();
			
		} catch (Exception e) {
			throw new RuntimeException(e);
		} finally {
			JdbcUtil.closeAll(con, pstmt, null);
		}
	}
    
	
	/**
	 * 查询的通用方法
	 * @param sql
	 * @param paramsValue
	 */
	public <T> List<T> query(String sql, Object[] paramsValue,Class<T> clazz){
		
		try {
			// 返回的集合
			List<T> list = new ArrayList<T>();
			// 对象
			T t = null;
			
			// 1. 获取连接
			con = JdbcUtil.getConnection();
            
			// 2. 创建stmt对象
			pstmt = con.prepareStatement(sql);
            
			// 3. 获取占位符参数的个数， 并设置每个参数的值
			//int count = pstmt.getParameterMetaData().getParameterCount();
			if (paramsValue != null && paramsValue.length > 0) {
				for (int i=0; i<paramsValue.length; i++) {
					pstmt.setObject(i+1, paramsValue[i]);
				}
			}
            
			// 4. 执行查询
			rs = pstmt.executeQuery();
            
			// 5. 获取结果集元数据
			ResultSetMetaData rsmd = rs.getMetaData();
            
			// ---> 获取列的个数
			int columnCount = rsmd.getColumnCount();
			
			// 6. 遍历rs
			while (rs.next()) {
				// 要封装的对象
				t = clazz.newInstance();
				
				// 7. 遍历每一行的每一列, 封装数据
				for (int i=0; i<columnCount; i++) {
					// 获取每一列的列名称
					String columnName = rsmd.getColumnName(i + 1);
					// 获取每一列的列名称, 对应的值
					Object value = rs.getObject(columnName);
					// 封装： 设置到t对象的属性中  【BeanUtils组件】
					BeanUtils.copyProperty(t, columnName, value);				
				}
				
				// 把封装完毕的对象，添加到list集合中
				list.add(t);
			}
			return list;
		} catch (Exception e) {
			throw new RuntimeException(e);
		} finally {
			JdbcUtil.closeAll(con, pstmt, rs);
		}
	}
}
```

### Dao

```java
package cn.ixfosa.jdbc;

import java.util.List;

public class AdminDao extends BaseDao {

	// 删除
	public void delete(int id) {
		String sql = "delete from admin where id=?";
		Object[] paramsValue = {id};
		super.update(sql, paramsValue);
	}

	// 插入
	public void save(Admin admin) {
		String sql = "insert into admin (userName,pwd) values (?,?)";
		Object[] paramsValue = {admin.getUserName(),admin.getPwd()};
		super.update(sql, paramsValue);
	}
	
	// 查询全部
	public List<Admin> getAll(){
		String sql = "select * from admin";
		List<Admin> list = super.query(sql, null, Admin.class);
		return list;
	}
	
	// 根据条件查询(主键)
	public Admin findById(int id){
		String sql = "select * from admin where id=?";
		List<Admin> list = super.query(sql, new Object[]{id}, Admin.class);
		return  (list!=null&&list.size()>0) ? list.get(0) : null;
	}
}
```

### Test

```JAVA
package cn.IXFOSA.jdbc;

import java.util.List;

import org.junit.Test;

public class AdminDaoTest {

	@Test
	public void testUpdate() throws Exception {
		AdminDao adminDao = new AdminDao();
		//adminDao.delete(2);
		//adminDao.save(new Admin());
		
		// 测试查询
		List<Admin> list = adminDao.getAll();
		System.out.println(list);
	}
}
```



## 元数据

在jdbc中获取数据库的定义，例如：数据库、表、列的定义信息。就用到元数据。
在jdbc中可以使用： 数据库元数据、参数元数据、结果集元数

```java
package cn.ixfosa.metadata;

import java.sql.Connection;
import java.sql.DatabaseMetaData;
import java.sql.ParameterMetaData;
import java.sql.PreparedStatement;
import java.sql.ResultSet;
import java.sql.ResultSetMetaData;

import org.junit.Test;

import cn.ixfosa.utils.JdbcUtil;

public class App {

	//1. 数据库元数据
	@Test
	public void testDB() throws Exception {
		// 获取连接
		Connection conn = JdbcUtil.getConnection();
		// 获取数据库元数据  // alt + shift + L  快速获取方法返回值
		DatabaseMetaData metaData = conn.getMetaData();
		
		System.out.println(metaData.getUserName());
		System.out.println(metaData.getURL());
		System.out.println(metaData.getDatabaseProductName());
	}
	
	//2. 参数元数据
	@Test
	public void testParams() throws Exception {
		// 获取连接
		Connection conn = JdbcUtil.getConnection();
		// SQL
		String sql = "select * from dept where deptid=? and deptName=?";
		// Object[] values = {"tom","888"};
		
		PreparedStatement pstmt = conn.prepareStatement(sql);
		// 参数元数据
		ParameterMetaData p_metaDate = pstmt.getParameterMetaData();
		// 获取参数的个数
		int count = p_metaDate.getParameterCount();
		
		
		// 测试
		System.out.println(count);
	}
	
	// 3. 结果集元数据
	@Test
	public void testRs() throws Exception {
		String sql = "select * from dept ";
		
		// 获取连接
		Connection conn = JdbcUtil.getConnection();
		PreparedStatement pstmt = conn.prepareStatement(sql);
		ResultSet rs = pstmt.executeQuery();
		// 得到结果集元数据(目标：通过结果集元数据，得到列的名称)
		ResultSetMetaData rs_metaData = rs.getMetaData();
		
		// 迭代每一行结果
		while (rs.next()) {
            
			// 1. 获取列的个数
			int count = rs_metaData.getColumnCount();
            
			// 2. 遍历，获取每一列的列的名称
			for (int i=0; i<count; i++) {
                
				// 得到列的名称
				String columnName = rs_metaData.getColumnName(i + 1);
                
				// 获取每一行的每一列的值
				Object columnValue = rs.getObject(columnName);
                
				// 测试
				System.out.print(columnName + "=" + columnValue + ",");
			}
			System.out.println();
		}
	}
}
```



## 连接池



思考：
		程序中连接如何管理？

1. 连接资源宝贵；需要对连接管理

2. 连接：
           a)操作数据库，创建连接
            b)操作结束，  关闭！

3. 分析：
   	涉及频繁的连接的打开、关闭，影响程序的运行效率！
   连接管理：
   	预先创建一组连接，有的时候每次取出一个； 用完后，放回；

   

学习连接池：
       a.自定义一个连接池
        b.学习优秀的连接池组件 DBCP, C3P0

### 代理：

如果对某个接口中的某个指定的方法的功能进行扩展，而不想实现接口里所有方法，可以使用(动态)代理模式!
		Java中代理模式：静态/动态/Cglib代理(spring)
		使用动态代理，可以监测接口中方法的执行！

如何对Connection对象，生成一个代理对象：

```java
|--Proxy
    
	static Object newProxyInstance(
		ClassLoader loader,    		当前使用的类加载器
    	Class<?>[] interfaces,   	目标对象(Connection)实现的接口类型
                InvocationHandler h 事件处理器：当执行上面接口中的方法的时候，就会自动触发											事件处理器代码，把当前执行的方法(method)作为参数传入。
	)  
```



### 自定义连接池

```java
package cn.ixfosa.pool;

import java.lang.reflect.InvocationHandler;
import java.lang.reflect.Method;
import java.lang.reflect.Proxy;
import java.sql.Connection;
import java.sql.DriverManager;
import java.sql.SQLException;
import java.util.LinkedList;

/**
 * 自定义连接池, 管理连接
 * 代码实现：
	1.  MyPool.java  连接池类，   
	2.  指定全局参数：  初始化数目、最大连接数、当前连接、   连接池集合
	3.  构造函数：循环创建3个连接
	4.  写一个创建连接的方法
	5.  获取连接
	------>  判断： 池中有连接， 直接拿
	 ------>                池中没有连接，
	------>                 判断，是否达到最大连接数； 达到，抛出异常；没有达到最大连接数，
			创建新的连接
	6. 释放连接
	 ------->  连接放回集合中(..)
 *
 */
public class MyPool {

	private int init_count = 3;		// 初始化连接数目
	private int max_count = 6;		// 最大连接数
	private int current_count = 0;  // 记录当前使用连接数
	// 连接池 （存放所有的初始化连接）
	private LinkedList<Connection> pool = new LinkedList<Connection>();
	
	
	//1.  构造函数中，初始化连接放入连接池
	public MyPool() {
		// 初始化连接
		for (int i=0; i<init_count; i++){
			// 记录当前连接数目
			current_count++;
			// 创建原始的连接对象
			Connection con = createConnection();
			// 把连接加入连接池
			pool.addLast(con);
		}
	}
	
	//2. 创建一个新的连接的方法
	private Connection createConnection(){
		try {
			Class.forName("com.mysql.jdbc.Driver");
			// 原始的目标对象
			final Connection con = DriverManager.getConnection("jdbc:mysql:///jdbc_demo", "root", "root");
			
			/**********对con对象代理**************/
			
			// 对con创建其代理对象
			Connection proxy = (Connection) Proxy.newProxyInstance(
					
					con.getClass().getClassLoader(),    // 类加载器
					//con.getClass().getInterfaces(),   // 当目标对象是一个具体的类的时候 
					new Class[]{Connection.class},      // 目标对象实现的接口
					
					new InvocationHandler() {// 当调用con对象方法的时候， 自动触发事务处理器
						@Override
						public Object invoke(Object proxy, Method method, Object[] args)
								throws Throwable {
							// 方法返回值
							Object result = null;
							// 当前执行的方法的方法名
							String methodName = method.getName();
							
							// 判断当执行了close方法的时候，把连接放入连接池
							if ("close".equals(methodName)) {
								System.out.println("begin:当前执行close方法开始！");
								// 连接放入连接池 (判断..)
								pool.addLast(con);
								System.out.println("end: 当前连接已经放入连接池了！");
							} else {
								// 调用目标对象方法
								result = method.invoke(con, args);
							}
							return result;
						}
					}
			);
			return proxy;
		} catch (Exception e) {
			throw new RuntimeException(e);
		}
	}
	
	//3. 获取连接
	public Connection getConnection(){
		
		// 3.1 判断连接池中是否有连接, 如果有连接，就直接从连接池取出
		if (pool.size() > 0){
			return pool.removeFirst();
		}
		
		// 3.2 连接池中没有连接： 判断，如果没有达到最大连接数，创建；
		if (current_count < max_count) {
			// 记录当前使用的连接数
			current_count++;
			// 创建连接
			return createConnection();
		}
		
		// 3.3 如果当前已经达到最大连接数，抛出异常
		throw new RuntimeException("当前连接已经达到最大连接数目 ！");
	}
	
	
	//4. 释放连接
	public void realeaseConnection(Connection con) {
		// 4.1 判断： 池的数目如果小于初始化连接，就放入池中
		if (pool.size() < init_count){
			pool.addLast(con);
		} else {
			try {
				// 4.2 关闭 
				current_count--;
				con.close();
			} catch (SQLException e) {
				throw new RuntimeException(e);
			}
		}
	}
	
	public static void main(String[] args) throws SQLException {
		MyPool pool = new MyPool();
		System.out.println("当前连接: " + pool.current_count);  // 3
		
		// 使用连接
		pool.getConnection();
		pool.getConnection();
		Connection con4 = pool.getConnection();
		Connection con3 = pool.getConnection();
		Connection con2 = pool.getConnection();
		Connection con1 = pool.getConnection();
		
		// 释放连接, 连接放回连接池
		//pool.realeaseConnection(con1);
		/*
		 * 希望：当关闭连接的时候，要把连接放入连接池！【当调用Connection接口的close方法时候，希望		   *触发pool.addLast(con);操作】把连接放入连接池
		 * 解决1：实现Connection接口，重写close方法
		 * 解决2：动态代理
		 */
		con1.close();
		
		// 再获取
		pool.getConnection();
		
		System.out.println("连接池：" + pool.pool.size());      // 0
		System.out.println("当前连接: " + pool.current_count);  // 3
	}
}
```

### DBCP连接池

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



## 分页技术

### jar

- c3p0-0.9.1.2.jar
- commons-dbutils-1.6.jar
- mysql-connector-java-5.1.12-bin.jar

### utils

src/c3p0-config.xml

```java
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
package cn.ixfosa.utils;

import javax.sql.DataSource;

import org.apache.commons.dbutils.QueryRunner;

import com.mchange.v2.c3p0.ComboPooledDataSource;

/**
 * 工具类
 * 1. 初始化C3P0连接池
 * 2. 创建DbUtils核心工具类对象
 * @author Jie.Yuan
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

```java
package cn.ixfosa.utils;

import java.util.List;

import cn.ixfosa.entity.Employee;

/**
 * 封装分页的参数
 * 
 */
public class PageBean<T> {
	private int currentPage = 1; // 当前页, 默认显示第一页
	private int pageCount = 4;   // 每页显示的行数(查询返回的行数), 默认每页显示4行
	private int totalCount;      // 总记录数
	private int totalPage;       // 总页数 = 总记录数 / 每页显示的行数  (+ 1)
	private List<T> pageData;       // 分页查询到的数据
	
	// 返回总页数
	public int getTotalPage() {
		if (totalCount % pageCount == 0) {
			totalPage = totalCount / pageCount;
		} else {
			totalPage = totalCount / pageCount + 1;
		}
		return totalPage;
	}
	public void setTotalPage(int totalPage) {
		this.totalPage = totalPage;
	}
	
	public int getCurrentPage() {
		return currentPage;
	}
	public void setCurrentPage(int currentPage) {
		this.currentPage = currentPage;
	}
	public int getPageCount() {
		return pageCount;
	}
	public void setPageCount(int pageCount) {
		this.pageCount = pageCount;
	}
	public int getTotalCount() {
		return totalCount;
	}
	public void setTotalCount(int totalCount) {
		this.totalCount = totalCount;
	}
	
	public List<T> getPageData() {
		return pageData;
	}
	public void setPageData(List<T> pageData) {
		this.pageData = pageData;
	}
}
```



### entity

```java
package cn.ixfosa.entity;

/**
 * 1. 实体类设计 (因为用了DbUtils组件，属性要与数据库中字段一致)
 *
 */
public class Employee {

	private int empId;			// 员工id
	private String empName;		// 员工名称
	private int dept_id;		// 部门id
	
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

### dao

```java
package cn.ixfosa.dao;

import cn.ixfosa.entity.Employee;
import cn.ixfosa.utils.PageBean;

/**
 * 2. 数据访问层，接口设计
 *
 */
public interface IEmployeeDao {

	/**
	 * 分页查询数据
	 */
	public void getAll(PageBean<Employee> pb);
	
	/**
	 * 查询总记录数
	 */
	public int getTotalCount();
}
```

```java
package cn.ixfosa.dao.impl;

import java.sql.SQLException;
import java.util.List;

import org.apache.commons.dbutils.QueryRunner;
import org.apache.commons.dbutils.handlers.BeanListHandler;
import org.apache.commons.dbutils.handlers.ScalarHandler;

import cn.ixfosa.dao.IEmployeeDao;
import cn.ixfosa.entity.Employee;
import cn.ixfosa.utils.JdbcUtils;
import cn.ixfosa.utils.PageBean;

/**
 * 2. 数据访问层实现
 *
 */
public class EmployeeDao implements IEmployeeDao {

	@Override
	public void getAll(PageBean<Employee> pb) {
		
		//2. 查询总记录数;  设置到pb对象中
		int totalCount = this.getTotalCount();
		pb.setTotalCount(totalCount);
		
		/*
		 * 问题： jsp页面，如果当前页为首页，再点击上一页报错！
		 *              如果当前页为末页，再点下一页显示有问题！
		 * 解决：
		 * 	   1. 如果当前页 <= 0;       当前页设置当前页为1;
		 * 	   2. 如果当前页 > 最大页数；  当前页设置为最大页数
		 */
		// 判断
		if (pb.getCurrentPage() <=0) {
			pb.setCurrentPage(1);					    // 把当前页设置为1
		} else if (pb.getCurrentPage() > pb.getTotalPage()){
			pb.setCurrentPage(pb.getTotalPage());		// 把当前页设置为最大页数
		}
		
		//1. 获取当前页： 计算查询的起始行、返回的行数
		int currentPage = pb.getCurrentPage();
		int index = (currentPage -1 ) * pb.getPageCount();		// 查询的起始行
		int count = pb.getPageCount();							// 查询返回的行数
		
		
		//3. 分页查询数据;  把查询到的数据设置到pb对象中
		String sql = "select * from employee limit ?,?";
		
		try {
			// 得到Queryrunner对象
			QueryRunner qr = JdbcUtils.getQueryRuner();
			// 根据当前页，查询当前页数据(一页数据)
			List<Employee> pageData = qr.query(sql, new BeanListHandler<Employee>(Employee.class), index, count);
			// 设置到pb对象中
			pb.setPageData(pageData);
			
		} catch (Exception e) {
			throw new RuntimeException(e);
		}
		
	}

	@Override
	public int getTotalCount() {
		String sql = "select count(*) from employee";
		try {
			// 创建QueryRunner对象
			QueryRunner qr = JdbcUtils.getQueryRuner();
			// 执行查询， 返回结果的第一行的第一列
			Long count = qr.query(sql, new ScalarHandler<Long>());
			return count.intValue();
		} catch (Exception e) {
			throw new RuntimeException(e);
		}
	}

}
```

### service

```java
package cn.ixfosa.service;

import cn.service.entity.Employee;
import cn.service.utils.PageBean;

/**
 * 3. 业务逻辑层接口设计
 *
 */
public interface IEmployeeService {

	/**
	 * 分页查询数据
	 */
	public void getAll(PageBean<Employee> pb);
}
```

```java
package cn.ixfosa.service.impl;

import cn.ixfosa.dao.IEmployeeDao;
import cn.ixfosa.dao.impl.EmployeeDao;
import cn.ixfosa.entity.Employee;
import cn.ixfosa.service.IEmployeeService;
import cn.ixfosa.utils.PageBean;

/**
 * 3. 业务逻辑层，实现
 *
 */
public class EmployeeService implements IEmployeeService {
	
	// 创建Dao实例
	private IEmployeeDao employeeDao = new EmployeeDao();

	@Override
	public void getAll(PageBean<Employee> pb) {
		try {
			employeeDao.getAll(pb);
		} catch (Exception e) {
			throw new RuntimeException(e);
		}
	}
}
```



### 	servlet

```java
package cn.ixfosa.servlet;

import java.io.IOException;
import java.io.PrintWriter;

import javax.servlet.ServletException;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import cn.ixfosa.entity.Employee;
import cn.ixfosa.service.IEmployeeService;
import cn.ixfosa.service.impl.EmployeeService;
import cn.ixfosa.utils.PageBean;

/**
 * 4. 控制层开发
 *
 */
public class IndexServlet extends HttpServlet {
	// 创建Service实例
	private IEmployeeService employeeService = new EmployeeService();
	// 跳转资源
	private String uri;

	public void doGet(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, IOException {
		
		try {
			//1. 获取“当前页”参数；  (第一次访问当前页为null) 
			String currPage = request.getParameter("currentPage");
			// 判断
			if (currPage == null || "".equals(currPage.trim())){
				currPage = "1";  	// 第一次访问，设置当前页为1;
			}
			// 转换
			int currentPage = Integer.parseInt(currPage);
			
			//2. 创建PageBean对象，设置当前页参数； 传入service方法参数
			PageBean<Employee> pageBean = new PageBean<Employee>();
			pageBean.setCurrentPage(currentPage);
			
			//3. 调用service  
			employeeService.getAll(pageBean);    // 【pageBean已经被dao填充了数据】
			
			//4. 保存pageBean对象，到request域中
			request.setAttribute("pageBean", pageBean);
			
			//5. 跳转 
			uri = "/WEB-INF/list.jsp";
		} catch (Exception e) {
			e.printStackTrace();  // 测试使用
			// 出现错误，跳转到错误页面；给用户友好提示
			uri = "/error/error.jsp";
		}
		request.getRequestDispatcher(uri).forward(request, response);
		
	}

	public void doPost(HttpServletRequest request, HttpServletResponse response)
			throws ServletException, IOException {
		this.doGet(request, response);
	}
}
```



### jsp

WebRoot/WEB-INF/list.jsp

```jsp
<%@ page language="java" import="java.util.*" pageEncoding="UTF-8"%>
<!-- 引入jstl核心标签库 -->
<%@taglib uri="http://java.sun.com/jsp/jstl/core" prefix="c" %>

<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN">
<html>
  <head>
    <title>分页查询数据</title>
	<meta http-equiv="pragma" content="no-cache">
	<meta http-equiv="cache-control" content="no-cache">
	<meta http-equiv="expires" content="0">    
  </head>
  
  <body>
  	<table border="1" width="80%" align="center" cellpadding="5" cellspacing="0">
  		<tr>
  			<td>序号</td>
  			<td>员工编号</td>
  			<td>员工姓名</td>
  		</tr>
  		<!-- 迭代数据 -->
  		<c:choose>
  			<c:when test="${not empty requestScope.pageBean.pageData}">
  				<c:forEach var="emp" items="${requestScope.pageBean.pageData}" varStatus="vs">
  					<tr>
  						<td>${vs.count }</td>
  						<td>${emp.empId }</td>
  						<td>${emp.empName }</td>
  					</tr>
  				</c:forEach>
  			</c:when>
  			<c:otherwise>
  				<tr>
  					<td colspan="3">对不起，没有你要找的数据</td>
  				</tr>
  			</c:otherwise>
  		</c:choose>
  		
  		<tr>
  			<td colspan="3" align="center">
  				当前${requestScope.pageBean.currentPage }/${requestScope.pageBean.totalPage }页     &nbsp;&nbsp;
  				
  				<a href="${pageContext.request.contextPath }/index?currentPage=1">首页</a>
  				<a href="${pageContext.request.contextPath }/index?currentPage=${requestScope.pageBean.currentPage-1}">上一页 </a>
  				<a href="${pageContext.request.contextPath }/index?currentPage=${requestScope.pageBean.currentPage+1}">下一页 </a>
  				<a href="${pageContext.request.contextPath }/index?currentPage=${requestScope.pageBean.totalPage}">末页</a>
  			</td>
  		</tr>
  	</table>
  </body>
</html>
```

### web.xml

```xml
<?xml version="1.0" encoding="UTF-8"?>
<web-app version="2.5" 
	xmlns="http://java.sun.com/xml/ns/javaee" 
	xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" 
	xsi:schemaLocation="http://java.sun.com/xml/ns/javaee 
	http://java.sun.com/xml/ns/javaee/web-app_2_5.xsd">
	
  <servlet>
    <servlet-name>IndexServlet</servlet-name>
    <servlet-class>cn.ixfosa.servlet.IndexServlet</servlet-class>
  </servlet>

  <servlet-mapping>
    <servlet-name>IndexServlet</servlet-name>
    <url-pattern>/index</url-pattern>
  </servlet-mapping>
  <welcome-file-list>
    <welcome-file>index.jsp</welcome-file>
  </welcome-file-list>
  
</web-app>
```

