
// 字符串转换
// 当我们需要一个字符串形式的值时，就会进行字符串转换。
// 比如，alert(value) 将 value 转换为字符串类型，然后显示这个值。


// 可以显式地调用 String(value) 来将 value 转换为字符串类型：
let isLove = true;
console.log(typeof isLove);

isLove = String(isLove);
console.log(typeof isLove);
// 字符串转换最明显。false 变成 "false"，null 变成 "null" 等。






// 数字型转换
// 在算术函数和表达式中，会自动进行 number 类型转换。

// 比如，当把除法 / 用于非 number 类型：
console.log("6" / "2");
console.log(6 / "2");
console.log("6" / 2);

// 使用 Number(value) 显式地将这个 value 转换为 number 类型。
let strNumber = "66";
console.log(typeof strNumber);   // string

strNumber = Number(strNumber);
console.log(typeof strNumber);   // number

// 当我们从 string 类型源（如文本表单）中读取一个值，但期望输入一个数字时，通常需要进行显式转换。

// 如果该字符串不是一个有效的数字，转换的结果会是 NaN。例如：
let age = "a123"
age = Number(age);
console.log(age);   // NaN
console.log(typeof age);   // number

age = "123a"
age = Number(age);
console.log(typeof age); // number



// number 类型转换规则：
//   值	                变成……
// - undefined	        NaN
// - null	            0
// - true 和 false	    1 and 0
// - string	           去掉首尾空格后的纯数字字符串中含有的数字。如果剩余字符串为空，则转换结果为 0。否则，将会从剩余字符串中“读取”数字。当类型转换出现 error 时返回 NaN。
console.log( Number("   123   ") ); // 123
console.log( Number("123z") );      // NaN（从字符串“读取”数字，读到 "z" 时出现错误）
console.log( Number(true) );        // 1
console.log( Number(false) );       // 0

// 请注意 null 和 undefined 在这有点不同：null 变成数字 0，undefined 变成 NaN。
console.log(Number(undefined));  // NaN
console.log(Number(null)); // 0
console.log(Number(NaN));  // NaN




// 布尔型转换
// 它发生在逻辑运算中（稍后我们将进行条件判断和其他类似的东西），
// 但是也可以通过调用 Boolean(value) 显式地进行转换。

// 转换规则如下：
// 直观上为“空”的值（如 0、空字符串、null、undefined 和 NaN）将变为 false。
// 其他值变成 true。

console.log( Boolean(1) ); // true
console.log( Boolean(0) ); / // false
console.log( Boolean("hello") ); // true
console.log( Boolean("") ); // false
console.log( Boolean(" ") ); // true
console.log( Boolean([])); // true
console.log( Boolean({}) ); // true


/**
    总结
        有三种常用的类型转换：转换为 string 类型、转换为 number 类型和转换为 boolean 类型。


字符串转换 —— 转换发生在输出内容的时候，也可以通过 String(value) 进行显式转换。原始类型值的 string 类型转换通常是很明显的。
数字型转换 —— 转换发生在进行算术操作时，也可以通过 Number(value) 进行显式转换。

数字型转换遵循以下规则：
    值	            变成
    undefined	    NaN
    null	        0
    true / false	1 / 0
    string      	“按原样读取”字符串，两端的空白会被忽略。空字符串变成 0。转换出错则输出 NaN。


布尔型转换 —— 转换发生在进行逻辑操作时，也可以通过 Boolean(value) 进行显式转换。

布尔型转换遵循以下规则：
值	                             变成……
0, null, undefined, NaN, ""	     false
其他值	                         true
上述的大多数规则都容易理解和记忆。人们通常会犯错误的值得注意的例子有以下几个：
    对 undefined 进行数字型转换时，输出结果为 NaN，而非 0。
    对 "0" 和只有空格的字符串（比如：" "）进行布尔型转换时，输出结果为 true。
    
 */