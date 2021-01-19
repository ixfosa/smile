// “if” 语句
    //- if(...) 语句计算括号里的条件表达式，如果计算结果是 true，就会执行对应的代码块。

if (year == 2021) {
  console.log( "That's correct!" );
  console.log( "You're so smart!" );
}
// 建议每次使用 if 语句都用大括号 {} 来包装代码块，即使只有一条语句。这样可以提高代码可读性。



// 布尔转换
    //- if (…) 语句会计算圆括号内的表达式，并将计算结果转换为布尔型。

// 数字 0、空字符串 ""、null、undefined 和 NaN 都会被转换成 false。因为它们被称为“假值（falsy）”值。
// 其他值被转换为 true，所以它们被称为“真值（truthy）”。
// 所以，下面这个条件下的代码永远不会执行：

if (0) { // 0 是假值（falsy）
  //- ...
}

// ……但下面的条件 —— 始终有效：
if (1) { // 1 是真值（truthy）
  //- ...
}

//也可以将未计算的布尔值传入 if 语句，像这样：
let cond = (year == 2015); // 相等运算符的结果是 true 或 false
if (cond) {
  //- ...
}


// “else” 语句
    //- if 语句有时会包含一个可选的 “else” 块。如果判断条件不成立，就会执行它内部的代码。
if (year == 2021) {
    console.log( 'You guessed it right!' );
} else {
    console.log( 'How can you be so wrong?' ); // 2015 以外的任何值
}

// 多个条件：“else if”
    //- 有时我们需要测试一个条件的几个变体。我们可以通过使用 else if 子句实现。
if (year < 2021) {
    console.log( 'Too early...' );
} else if (year > 2015) {
    console.log( 'Too late' );
} else {
    console.log( 'Exactly!' );
}

// 可以有更多的 else if 块。结尾的 else 是可选的。


// 条件运算符 ‘?’
let accessAllowed;
let age = prompt('How old are you?', '');
if (age > 18) {
  accessAllowed = true;
} else {
  accessAllowed = false;
}
console.log(accessAllowed);



// 这个运算符通过问号 ? 表示。有时它被称为三元运算符，被称为“三元”是因为该运算符中有三个操作数。
// 实际上它是 JavaScript 中唯一一个有这么多操作数的运算符。
// 语法：
let result = condition ? value1 : value2;
// 计算条件结果，如果结果为真，则返回 value1，否则返回 value2。

let accessAllowed = (age > 18) ? true : false;
// 第一个问号检查 age < 3。
// 如果为真 — 返回 'Hi, baby!'。否则，会继续执行冒号 ":" 后的表达式，检查 age < 18。
// 如果为真 — 返回 'Hello!'。否则，会继续执行下一个冒号 ":" 后的表达式，检查 age < 100。
// 如果为真 — 返回 'Greetings!'。否则，会继续执行最后一个冒号 ":" 后面的表达式，返回 'What an unusual age!'。
// 这是使用 if..else 实现上面的逻辑的写法：

if (age < 3) {
  message = 'Hi, baby!';
} else if (age < 18) {
  message = 'Hello!';
} else if (age < 100) {
  message = 'Greetings!';
} else {
  message = 'What an unusual age!';
}

// ‘?’ 的非常规使用
// 有时可以使用问号 ? 来代替 if 语句：
let company = prompt('Which company created JavaScript?', '');
(company == 'Netscape') ?
   alert('Right!') : alert('Wrong.');

// 三元运算符写法比 if 语句更短，对一些程序员很有吸引力。但它的可读性差。
let company = prompt('Which company created JavaScript?', '');
if (company == 'Netscape') {
  alert('Right!');
} else {
  alert('Wrong.');
}
