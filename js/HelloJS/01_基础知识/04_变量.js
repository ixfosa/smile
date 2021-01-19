// 变量---变量就是用来储存这些信息的。


// 在 JavaScript 中创建一个变量，我们需要用到 let 关键字。
// 下面的语句创建（也可以称为 声明 或者 定义）了一个名称为 “message” 的变量
let name = "ixfsoa";


// 通过赋值运算符 = 为变量添加一些数据：
let message;
message = "hello js";   // 保存字符串
console.log(message);   // 显示变量内容 


// 可以将变量定义和赋值合并成一行：
let age = 22;
console.log(age);


// 在一行中声明多个变量：
// 不推荐这样。为了更好的可读性，请一行只声明一个变量。
let user = "ixfosa", sex = "man";
console.log(`user: ${user} , age: ${age}`);

// let user = 'John';
// let age = 25;
// let message = 'Hello';


let buf;
buf = "js";
buf = "JavaScript";   // 当值改变的时候，之前的数据就被从变量中删除了：
console.log("buf: " + buf);



// 声明两个变量，然后将其中一个变量的数据拷贝到另一个变量。
let str1 = "hello js";
let str2;
str2 = str1;
console.log(`str1: ${str1}`);
console.log(`str2: ${str2}`);


// 声明两次会触发 error
// 一个变量应该只被声明一次。

// 对同一个变量进行重复声明会触发 error：
// let info = "error";
// let info = "warn"   // SyntaxError: Identifier 'info' has already been declared


//变量命名
 /**
    JavaScript 的变量命名有两个限制：
        变量名称必须仅包含字母，数字，符号 $ 和 _。
        首字符必须非数字。
        有效的命名，例如：
            let userName;
            let test123;

    如果命名包括多个单词，通常采用驼峰式命名法（camelCase）。也就是，单词一个接一个，除了第一个单词，其他的每个单词都以大写字母开头：myVeryLongName。

    有趣的是，美元符号 '$' 和下划线 '_' 也可以用于变量命名。它们是正常的符号，就跟字母一样，没有任何特殊的含义。


    允许非英文字母，但不推荐
    可以使用任何语言，包括西里尔字母（cyrillic letters）甚至是象形文字，就像这样：
        let имя = '...';
        let 我 = '...';
  */

 // 下面的命名是有效的：
 let $ = 1; // 使用 "$" 声明一个变量
 let _ = 2; // 现在用 "_" 声明一个变量
console.log($ + _); // 3

// 下面的变量命名不正确：

// let 1a; // 不能以数字开始
// let my-name; // 连字符 '-' 不允许用于变量命名

// 区分大小写
// 命名为 apple 和 AppLE 的变量是不同的两个变量。

/**
    保留字
        有一张 保留字列表，这张表中的保留字无法用作变量命名，因为它们被用于编程语言本身了。
            比如，let、class、return、function 都被保留了。

    下面的代码将会抛出一个语法错误：
        let let = 5; // 不能用 "let" 来命名一个变量，错误！
        let return = 5; // 同样，不能使用 "return"，错误！


    未采用 use strict 下的赋值
    一般，我们需要在使用一个变量前定义它。但是在早期，我们可以不使用 let 进行变量声明，
    而可以简单地通过赋值来创建一个变量。现在如果我们不在脚本中使用 use strict 声明启用严格模式，
    这仍然可以正常工作，这是为了保持对旧脚本的兼容。

    注意：这个例子中没有 "use strict"
    num = 5; // 如果变量 "num" 不存在，就会被创建
    alert(num); // 5
    上面这是个糟糕的做法，严格模式下会报错。

    "use strict";
    num = 5; // 错误：num 未定义
 */