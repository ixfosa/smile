

// 字符串
    //- 在 JavaScript 中，文本数据被以字符串形式存储，单个字符没有单独的类型。 
    //- 字符串的内部格式始终是 UTF-16，它不依赖于页面编码。



// -------------------------------------------------------------------------------------------



///引号（Quotes）
    //- 引号的种类。

// 字符串可以包含在单引号、双引号或反引号中：

let single = 'single-quoted';
let double = "double-quoted";

let backticks = `backticks`;

// 单引号和双引号基本相同。但是，反引号允许我们通过 ${…} 将任何表达式嵌入到字符串中：

function sum(a, b) {
  return a + b;
}
console.log(`1 + 2 = ${sum(1, 2)}.`); // 1 + 2 = 3.


// 使用反引号的另一个优点是它们允许字符串跨行：
let guestList = `Guests:
 * John
 * Pete
 * Mary
`;

console.log(guestList); // 客人清单，多行

// 如果我们使用单引号或双引号来实现字符串跨行的话，则会出现错误：
/*
    let guestList = "Guests: // Error: Unexpected token ILLEGAL
    * John";
*/


// 反引号还允许我们在第一个反引号之前指定一个“模版函数”。
    //- 语法是：func`string`。函数 func 被自动调用，接收字符串和嵌入式表达式，并处理它们。




// -------------------------------------------------------------------------------------------




// 特殊字符

// 我们可以通过使用“换行符（newline character）”，以支持使用单引号和双引号来创建跨行字符串。
// 换行符写作 \n，用来表示换行：

let guestList2 = "Guests:\n * John\n * Pete\n * Mary";

console.log(guestList2); // 一个多行的客人列表

// 例如，这两行描述的是一样的，只是书写方式不同：
let str11 = "Hello\nWorld"; // 使用“换行符”创建的两行字符串

// 使用反引号和普通的换行创建的两行字符串
let str22 = `Hello
World`;

console.log(str11 == str22); // true

// 还有其他不常见的“特殊”字符。

// 这是完整列表：
    //- 字符            	 描述
    //- \n	                换行
    //- \r	                回车：不单独使用。Windows 文本文件使用两个字符 \r\n 的组合来表示换行。
    //- \',                 \"	引号
    //- \\	                反斜线
    //- \t	                制表符
    //- \b, \f, \v	        退格，换页，垂直标签 —— 为了兼容性，现在已经不使用了。
    //- \xXX	            具有给定十六进制 Unicode XX 的 Unicode 字符，例如：'\x7A' 和 'z' 相同。
    //- \uXXXX	            以 UTF-16 编码的十六进制代码 XXXX 的 unicode 字符，例如 \u00A9 —— 是版权符号 © 的 unicode。它必须正好是 4 个十六进制数字。
    //- \u{X…XXXXXX}        （1 到 6 个十六进制字符）	具有给定 UTF-32 编码的 unicode 符号。一些罕见的字符用两个 unicode 符号编码，占用 4 个字节。这样我们就可以插入长代码了。

// unicode 示例： 
console.log( "\u00A9" );    // ©
console.log( "\u{20331}" ); // 佫，罕见的中国象形文字（长 unicode）
console.log( "\u{1F60D}" ); // 😍，笑脸符号（另一个长 unicode）

// 所有的特殊字符都以反斜杠字符 \ 开始。它也被称为“转义字符”。

// 如果我们想要在字符串中插入一个引号，我们也会使用它。
console.log('I\'m the Walrus!' ); // I'm the Walrus!
// 必须在内部引号前加上反斜杠 \'，否则它将表示字符串结束。

// 只有与外部闭合引号相同的引号才需要转义。因此，作为一个更优雅的解决方案，我们可以改用双引号或者反引号：
console.log( `I'm the Walrus!` ); // I'm the Walrus!

// 注意反斜杠 \ 在 JavaScript 中用于正确读取字符串，然后消失。内存中的字符串没有 \。

// 但是如果需要在字符串中显示一个实际的反斜杠 \ 应该怎么做？
// 可以这样做，只需要将其书写两次 \\：
console.log( `The backslash: \\` ); // The backslash: \



// -------------------------------------------------------------------------------------------



// 字符串长度
    //- length 属性表示字符串长度：

console.log( `My\n`.length ); // 3

// 注意 \n 是一个单独的“特殊”字符，所以长度确实是 3。


//length 是一个属性
    //- 掌握其他编程语言的人，有时会错误地调用 str.length() 而不是 str.length。这是行不通的。

// 请注意 str.length 是一个数字属性，而不是函数。后面不需要添加括号。




// -------------------------------------------------------------------------------------------



// 访问字符
    //- 要获取在 pos 位置的一个字符，可以使用方括号 [pos] 或者调用 str.charAt(pos) 方法。
    //- 第一个字符从零位置开始：

let str1 = `Hello`;
// 第一个字符
console.log( str1[0] );         // H    
console.log( str1.charAt(0) );  // H

// 最后一个字符
console.log( str1[str1.length - 1] ); // o

// 方括号是获取字符的一种现代化方法，而 charAt 是历史原因才存在的。

// 它们之间的唯一区别是，如果没有找到字符，[] 返回 undefined，而 charAt 返回一个空字符串：

let str2 = `Hello`;

console.log( str2[1000] ); // undefined
console.log( str2.charAt(1000) ); // ''（空字符串）

// 可以使用 for..of 遍历字符：
for (let char of "Hello") {
  console.log(char); // H,e,l,l,o（char 变为 "H"，然后是 "e"，然后是 "l" 等）
}



// -------------------------------------------------------------------------------------------



// 字符串是不可变的
    //- 在 JavaScript 中，字符串不可更改。改变字符是不可能的。


let s = 'Hi';

s[0] = 'h'; // error
console.log( s[0] ); // 无法运行


// 通常的解决方法是创建一个新的字符串，并将其分配给 str 而不是以前的字符串。
// 例如：
let s2 = 'Hi';
s2 = 'h' + s2[1];  // 替换字符串

console.log( s2 ); // hi



// -------------------------------------------------------------------------------------------



// toLowerCase() 和 toUpperCase() 方法可以改变大小写：

console.log( 'ixfosa'.toUpperCase() ); // IXFOSA
console.log( 'ixfosa'.toLowerCase() ); // ixfosa

// 或者我们想要使一个字符变成小写：
console.log( 'ixfosa'[0].toUpperCase() ); // 'i'



// -------------------------------------------------------------------------------------------



// 查找子字符串

// str.indexOf
    //- 第一个方法是 str.indexOf(substr, pos)。
    //- 它从给定位置 pos 开始，在 str 中查找 substr，如果没有找到，则返回 -1，否则返回匹配成功的位置。

let str3 = 'Widget with id';

console.log( str3.indexOf('Widget') ); // 0，因为 'Widget' 一开始就被找到
console.log( str3.indexOf('widget') ); // -1，没有找到，检索是大小写敏感的

console.log( str3.indexOf("id") ); // 1，"id" 在位置 1 处（……idget 和 id）

// 可选的第二个参数允许我们从给定的起始位置开始检索。

// 例如，"id" 第一次出现的位置是 1。查询下一个存在位置时，我们从 2 开始检索：

let str4 = 'Widget with id';
console.log( str4.indexOf('id', 2) ) // 12

// 如果我们对所有存在位置都感兴趣，可以在一个循环中使用 indexOf。每一次新的调用都发生在上一匹配位置之后：

let str5 = 'As sly as a fox, as strong as an ox';
let target5 = 'as'; // 这是我们要查找的目标
let pos5 = 0;
while (true) {
  let foundPos = str5.indexOf(target5, pos5);
  if (foundPos == -1) break;

  console.log( `Found at ${foundPos}` );
  pos5 = foundPos + 1; // 继续从下一个位置查找
}

// 相同的算法可以简写：
let str6 = "As sly as a fox, as strong as an ox";
let target6 = "as";

let pos6 = -1;
while ((pos6 = str6.indexOf(target6, pos6 + 1)) != -1) {
    console.log( pos6 );
}

// str.lastIndexOf(substr, pos)
    // 还有一个类似的方法 str.lastIndexOf(substr, position)，它从字符串的末尾开始搜索到开头。

// 在 if 测试中 indexOf 有一点不方便。我们不能像这样把它放在 if 中：
let str7 = "Widget with id";
if (str7.indexOf("Widget")) {
    console.log("We found it"); // 不工作！
}

//上述示例中的 alert 不会显示，因为 str.indexOf("Widget") 返回 0（意思是它在起始位置就查找到了匹配项）。是的，但是 if 认为 0 表示 false。

// 因此应该检查 -1，像这样：
let str8 = "Widget with id";
if (str8.indexOf("Widget") != -1) {
    console.log("We found it"); // 现在工作了！
}





// -------------------------------------------------------------------------------------------



// 按位（bitwise）NOT 技巧
    // 这里使用的一个老技巧是 bitwise NOT ~ 运算符。
    // 它将数字转换为 32-bit 整数（如果存在小数部分，则删除小数部分），
    // 然后对其二进制表示形式中的所有位均取反。

// 实际上，这意味着一件很简单的事儿：对于 32-bit 整数，~n 等于 -(n+1)。
console.log( ~2 ); // -3，和 -(2+1) 相同
console.log( ~1 ); // -2，和 -(1+1) 相同
console.log( ~0 ); // -1，和 -(0+1) 相同
console.log( ~-1 ); // 0，和 -(-1+1) 相同

// 只有当 n == -1 时，~n 才为零（适用于任何 32-bit 带符号的整数 n）。

// 因此，仅当 indexOf 的结果不是 -1 时，检查 if ( ~str.indexOf("...") ) 才为真。换句话说，当有匹配时。

// 用它来简写 indexOf 检查：
let str9 = "Widget";
if (~str9.indexOf("Widget")) {
  console.log( 'Found it!' ); // 正常运行
}

// 只要记住：if (~str.indexOf(...)) 读作 “if found”。

// 确切地说，由于 ~ 运算符将大数字截断为 32 位，因此存在给出 0 的其他数字，
// 最小的数字是 ~4294967295=0。这使得这种检查只有在字符串没有那么长的情况下才是正确的。




// -------------------------------------------------------------------------------------------



//includes，startsWith，endsWith
    //- 更现代的方法 str.includes(substr, pos) 根据 str 中是否包含 substr 来返回 true/false。

// 如果我们需要检测匹配，但不需要它的位置，那么这是正确的选择：

console.log( "Widget with id".includes("Widget") ); // true
console.log( "Hello".includes("Bye") ); // false

// str.includes 的第二个可选参数是开始搜索的起始位置：
console.log( "Midget".includes("id") ); // true
console.log( "Midget".includes("id", 3) ); // false, 从位置 3 开始没有 "id"

// 方法 str.startsWith 和 str.endsWith 的功能与其名称所表示的意思相同：
console.log( "Widget".startsWith("Wid") ); // true，"Widget" 以 "Wid" 开始
console.log( "Widget".endsWith("get") ); // true，"Widget" 以 "get" 结束




// -------------------------------------------------------------------------------------------




// 获取子字符串
    //- JavaScript 中有三种获取字符串的方法：substring、substr 和 slice。

// str.slice(start [, end])
    //- 返回字符串从 start 到（但不包括）end 的部分。

let str01 = "stringify";
console.log( str01.slice(0, 5) ); // 'strin'，从 0 到 5 的子字符串（不包括 5）
console.log( str01.slice(0, 1) ); // 's'，从 0 到 1，但不包括 1，所以只有在 0 处的字符

// 如果没有第二个参数，slice 会一直运行到字符串末尾：
console.log( str01.slice(2) ); // 从第二个位置直到结束

// start/end 也有可能是负值。它们的意思是起始位置从字符串结尾计算：
// 从右边的第四个位置开始，在右边的第一个位置结束
console.log( str01.slice(-4, -1) ); // 'gif'





// str.substring(start [, end])
    //- 返回字符串在 start 和 end 之间 的部分。

// 这与 slice 几乎相同，但它允许 start 大于 end。
// 这些对于 substring 是相同的
console.log( str01.substring(2, 6) ); // "ring"
console.log( str01.substring(6, 2) ); // "ring"

// ……但对 slice 是不同的：
console.log( str01.slice(2, 6) ); // "ring"（一样）
console.log( str01.slice(6, 2) ); // ""（空字符串）

// 不支持负参数（不像 slice），它们被视为 0。




// str.substr(start [, length])
    // 返回字符串从 start 开始的给定 length 的部分。

// 与以前的方法相比，这个允许我们指定 length 而不是结束位置：

// let str01 = "stringify";
console.log( str01.substr(2, 4) ); // 'ring'，从位置 2 开始，获取 4 个字符
// 第一个参数可能是负数，从结尾算起：

// let str01 = "stringify";
console.log( str01.substr(-4, 2) ); // 'gi'，从第 4 位获取 2 个字符

// 回顾一下这些方法，以免混淆：
//- 方法	                    选择方式……	                                   负值参数
//- slice(start, end)	        从 start 到 end（不含 end）	允许
//- substring(start, end)	    start 与 end 之间（包括 start，但不包括 end）	负值代表 0
//- substr(start, length)	    从 start 开始获取长为 length 的字符串	        允许 start 为负数

// 它们可以完成这项工作。形式上，substr 有一个小缺点：它不是在 JavaScript 核心规范中描述的，
// 而是在附录 B 中，它涵盖了主要由于历史原因而存在的仅浏览器特性。
// 因此，非浏览器环境可能无法支持它。但实际上它在任何地方都有效。

// 相较于其他两个变体，slice 稍微灵活一些，它允许以负值作为参数并且写法更简短。
// 因此仅仅记住这三种方法中的 slice 就足够了。





// -------------------------------------------------------------------------------------------




// 比较字符串
    //- 字符串按字母顺序逐字比较。
    //- 不过，也有一些奇怪的地方。

// 小写字母总是大于大写字母：
console.log( 'a' > 'Z' ); // true

// 带变音符号的字母存在“乱序”的情况：
console.log( 'Österreich' > 'Zealand' ); // true

// 如果我们对这些国家名进行排序，可能会导致奇怪的结果。
// 通常，人们会期望 Zealand 在名单中的 Österreich 之后出现。

// 为了明白发生了什么，我们回顾一下在 JavaScript 中字符串的内部表示。

// 所有的字符串都使用 UTF-16 编码。即：每个字符都有对应的数字代码。
// 有特殊的方法可以获取代码表示的字符，以及字符对应的代码。

// str.codePointAt(pos)
    //- 返回在 pos 位置的字符代码 :

// 不同的字母有不同的代码
console.log( "z".codePointAt(0) ); // 122
console.log( "Z".codePointAt(0) ); // 90


// String.fromCodePoint(code)
    //- 通过数字 code 创建字符
console.log( String.fromCodePoint(90) ); // Z

// 还可以用 \u 后跟十六进制代码，通过这些代码添加 unicode 字符：

// 在十六进制系统中 90 为 5a
console.log( '\u005a' ); // Z

// 现在我们看一下代码为 65..220 的字符（拉丁字母和一些额外的字符），方法是创建一个字符串：
let str02 = '';

for (let i = 65; i <= 220; i++) {
  str02 += String.fromCodePoint(i);
}
console.log( str02 );
// ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_`abcdefghijklmnopqrstuvwxyz{|}~
// ¡¢£¤¥¦§¨©ª«¬­®¯°±²³´µ¶·¸¹º»¼½¾¿ÀÁÂÃÄÅÆÇÈÉÊËÌÍÎÏÐÑÒÓÔÕÖ×ØÙÚÛÜ

// 看到没？先是大写字符，然后是一些特殊字符，然后是小写字符，而 Ö 几乎是最后输出。

// 现在很明显为什么 a > Z。

// 字符通过数字代码进行比较。越大的代码意味着字符越大。a（97）的代码大于 Z（90）的代码。




// -------------------------------------------------------------------------------------------




// 代理对
    // 所有常用的字符都是一个 2 字节的代码。
    // 大多数欧洲语言，数字甚至大多数象形文字中的字母都有 2 字节的表示形式。

    // 但 2 字节只允许 65536 个组合，这对于表示每个可能的符号是不够的。
    // 所以稀有的符号被称为“代理对”的一对 2 字节的符号编码。

// 这些符号的长度是 2：
console.log( '𝒳'.length ); // 2，大写数学符号 X
console.log( '😂'.length ); // 2，笑哭表情
console.log( '𩷶'.length ); // 2，罕见的中国象形文字

// 注意，代理对在 JavaScript 被创建时并不存在，因此无法被编程语言正确处理！

// 我们实际上在上面的每个字符串中都有一个符号，但 length 显示长度为 2。

// String.fromCodePoint 和 str.codePointAt 是几种处理代理对的少数方法。
// 它们最近才出现在编程语言中。在它们之前，只有 String.fromCharCode 和 str.charCodeAt。
// 这些方法实际上与 fromCodePoint/codePointAt 相同，但是不适用于代理对。

// 获取符号可能会非常麻烦，因为代理对被认为是两个字符：
console.log( '𝒳'[0] ); // 奇怪的符号……
console.log( '𝒳'[1] ); // ……代理对的一块

// 请注意，代理对的各部分没有任何意义。因此，上述示例中的显示的实际上是垃圾信息。

// 技术角度来说，代理对也是可以通过它们的代码检测到的：
    // 如果一个字符的代码在 0xd800..0xdbff 范围内，那么它是代理对的第一部分。
    // 下一个字符（第二部分）必须在 0xdc00..0xdfff 范围中。这些范围是按照标准专门为代理对保留的。


// charCodeAt 不理解代理对，所以它给出了代理对的代码
console.log( '𝒳'.charCodeAt(0).toString(16) ); // d835，在 0xd800 和 0xdbff 之间
console.log( '𝒳'.charCodeAt(1).toString(16) ); // dcb3, 在 0xdc00 和 0xdfff 之间




// -------------------------------------------------------------------------------------------




// 变音符号与规范化
    //- 在许多语言中，都有一些由基本字符组成的符号，在其上方/下方有一个标记。

// 例如，字母 a 可以是 àáâäãåā 的基本字符。最常见的“复合”字符在 UTF-16 表中都有自己的代码。
// 但不是全部，因为可能的组合太多。

// 为了支持任意组合，UTF-16 允许我们使用多个 unicode 字符：基本字符紧跟“装饰”它的一个或多个“标记”字符。

// 例如，如果我们 S 后跟有特殊的 “dot above” 字符（代码 \u0307），则显示 Ṡ。

console.log( 'S\u0307' ); // Ṡ

// 如果我们需要在字母上方（或下方）添加额外的标记 —— 没问题，只需要添加必要的标记字符即可。

// 例如，如果我们追加一个字符 “dot below”（代码 \u0323），那么我们将得到“S 上面和下面都有点”的字符：Ṩ。
console.log( 'S\u0307\u0323' ); // Ṩ

// 这在提供良好灵活性的同时，也存在一个有趣的问题：
    //- 两个视觉上看起来相同的字符，可以用不同的 unicode 组合表示。

let s01 = 'S\u0307\u0323'; // Ṩ，S + 上点 + 下点
let s02 = 'S\u0323\u0307'; // Ṩ，S + 下点 + 上点

console.log( `s01: ${s01}, s02: ${s02}` );

console.log( s01 == s02 ); // false，尽管字符看起来相同（?!）

// 为了解决这个问题，有一个 “unicode 规范化”算法，它将每个字符串都转化成单个“通用”格式。

// 它由 str.normalize() 实现。
console.log( "S\u0307\u0323".normalize() == "S\u0323\u0307".normalize() ); // true

// 有趣的是，在实际情况下，normalize() 实际上将一个由 3 个字符组成的序列合并为一个：\u1e68（S 有两个点）。

console.log( "S\u0307\u0323".normalize().length ); // 1

console.log( "S\u0307\u0323".normalize() == "\u1e68" ); // true





// -------------------------------------------------------------------------------------------



// 总结
    //- 有 3 种类型的引号。反引号允许字符串跨越多行并可以使用 ${…} 在字符串中嵌入表达式。
    //- JavaScript 中的字符串使用的是 UTF-16 编码。
    //- 我们可以使用像 \n 这样的特殊字符或通过使用 \u... 来操作它们的 unicode 进行字符插入。
    //- 获取字符时，使用 []。
    //- 获取子字符串，使用 slice 或 substring。
    //- 字符串的大/小写转换，使用：toLowerCase/toUpperCase。
    //- 查找子字符串时，使用 indexOf 或 includes/startsWith/endsWith 进行简单检查。
    //- 根据语言比较字符串时使用 localeCompare，否则将按字符代码进行比较。


