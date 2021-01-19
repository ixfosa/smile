"use strict";
// 函数表达式
    //- 在 JavaScript 中，函数不是“神奇的语言结构”，而是一种特殊的值。

//函数声明：
function sayHi1() {
  console.log( "Hello" );
}


//函数表达式。
let sayHi2 = function() {
    console.log( "Hello" );
};

console.log(sayHi2); // [Function: sayHi2]
console.log(sayHi2()); // Hello, undefined



function sayHi3() {   // (1) 创建
    console.log( "Hello" );
  }
  
  let func = sayHi3;    // (2) 复制
  
  func(); // Hello     // (3) 运行复制的值（正常运行）！
  sayHi3(); // Hello    //     这里也能运行



//-------------------------------------------------------------------------------
// 为什么这里末尾会有个分号？
// 你可能想知道，为什么函数表达式结尾有一个分号 ;，而函数声明没有：

function fn1() {
  // ...
}

let fn2 = function() {
  // ...
};

// 答案很简单：
    //- 在代码块的结尾不需要加分号 ;，像 if { ... }，for { }，function f { } 等语法结构后面都不用加。
    //- 函数表达式是在语句内部的：let sayHi = ...;，作为一个值。它不是代码块而是一个赋值语句。
    //- 不管值是什么，都建议在语句末尾添加分号 ;。所以这里的分号与函数表达式本身没有任何关系，它只是用于终止语句。

//-------------------------------------------------------------------------------




// 回调函数
    //- 将函数作为值来传递以及如何使用函数表达式。


// 一个包含三个参数的函数 ask(question, yes, no)：
    //- question: 关于问题的文本
    //- yes: 当回答为 “Yes” 时，要运行的脚本
    //- no: 当回答为 “No” 时，要运行的脚本
    
 // 函数需要提出 question（问题），并根据用户的回答，调用 yes() 或 no()：

function ask(question, yes, no) {
    if (question) yes()
    else no();
}

function showOk() {
  console.log( "You agreed." );
}

function showCancel() {
    console.log( "You canceled the execution." );
}

// 用法：函数 showOk 和 showCancel 被作为参数传入到 ask
// ask 的两个参数值 showOk 和 showCancel 可以被称为 回调函数 或简称 回调。
ask(true, showOk, showCancel);




// 用函数表达式对同样的函数进行大幅简写：
function ask(question, yes, no) {
    if (question) yes()
    else no();
}
// 直接在 ask(...) 调用内进行函数声明。这两个函数没有名字，所以叫 匿名函数。
// 这样的函数在 ask 外无法访问（因为没有对它们分配变量），不过这正是我们想要的。
ask(
    true,
    function() { console.log("You agreed."); },
    function() { console.log("You canceled the execution."); }
);




// 函数表达式 vs 函数声明


// 语法：
//- 函数声明：在主代码流中声明为单独的语句的函数。
// 函数声明
function sum1(a, b) {
  return a + b;
}

// 函数表达式：在一个表达式中或另一个语法结构中创建的函数。下面这个函数是在赋值表达式 = 右侧创建的：
// 函数表达式
let sum2 = function(a, b) {
  return a + b;
};

// JavaScript 引擎会在 什么时候 创建函数。
    //- 函数表达式是在代码执行到达时被创建，并且仅从那一刻起可用。
    //- 一旦代码执行到赋值表达式 let sum2 = function… 的右侧，此时就会开始创建该函数，
    //- 并且可以从现在开始使用（分配，调用等）。

    // 函数声明则不同
        //- 在函数声明被定义之前，它就可以被调用。
        //- 例如，一个全局函数声明对整个脚本来说都是可见的，无论它被写在这个脚本的哪个位置。
        //- 这是内部算法的原故。当 JavaScript 准备 运行脚本时，首先会在脚本中寻找全局函数声明，并创建这些函数。
        //- 我们可以将其视为“初始化阶段”。
        //- 在处理完所有函数声明后，代码才被执行。所以运行时能够使用这些函数。

// 例如下面的代码会正常工作：
sayHi3("ixfosa"); // Hello, ixfosa
function sayHi3(name) {
    console.log( `Hello, ${name}` );
}
// 函数声明 sayHi 是在 JavaScript 准备运行脚本时被创建的，在这个脚本的任何位置都可见。

// ……如果它是一个函数表达式，它就不会工作：
// sayHi4("ixfosa"); // error!    ReferenceError: Cannot access 'sayHi4' before initialization
let sayHi4 = function(name) {  // (*) no magic any more
    console.log( `Hello, ${name}` );
};


// ----------------------------------------------------------------------------

// 函数声明的另外一个特殊的功能是它们的块级作用域。
    //- 严格模式下，当一个函数声明在一个代码块内时，它在该代码块内的任何位置都是可见的。但在代码块外不可见。
    //- 例如，我们需要依赖于在代码运行过程中获得的变量 age 声明一个函数 welcome()。并且我们计划在之后的某个时间使用它。

// 如果我们使用函数声明，则以下代码无法像预期那样工作：
let age1 = 18;

// 有条件地声明一个函数
if (age1 < 18) {
    function welcome() {
        console.log("Hello!");
    }
} else {
    function welcome() {
        console.log("Greetings!");
    }

}

// ……稍后使用
// welcome(); // Error: welcome is not defined

// 这是因为函数声明只在它所在的代码块中可见。

// 下面是另一个例子：
let age2 = 16; // 拿 16 作为例子

if (age2 < 18) {
  welcome();               // \   (运行)
                           //  |
  function welcome() {     //  |
    console.log("Hello!");       //  |  函数声明在声明它的代码块内任意位置都可用
  }                        //  |
                           //  |
  welcome();               // /   (运行)

} else {

  function welcome() {
    alert("Greetings!");
  }
}

// 在这里，我们在花括号外部调用函数，我们看不到它们内部的函数声明。
// welcome(); // Error: welcome is not defined

// 我们怎么才能让 welcome 在 if 外可见呢？
    //- 正确的做法是使用函数表达式，并将 welcome 赋值给在 if 外声明的变量，并具有正确的可见性。

// 下面的代码可以如愿运行：
let age3 = 18;

let welcome1;

if (age3 < 18) {

  welcome1 = function() {
    console.log("Hello!");
  };

} else {

  welcome1 = function() {
    console.log("Greetings!");
  };

}
welcome1(); // 现在可以了

// 或者我们可以使用问号运算符 ? 来进一步对代码进行简化：
let age4 = 18;

let welcome2 = (age4 < 18) ?
  function() { console.log("Hello!"); } :
  function() { console.log("Greetings!"); };

welcome2(); // 现在可以了



// 总结
    //- 函数是值。它们可以在代码的任何地方被分配，复制或声明。
    //- 如果函数在主代码流中被声明为单独的语句，则称为“函数声明”。
    //- 如果该函数是作为表达式的一部分创建的，则称其“函数表达式”。
    //- 在执行代码块之前，内部算法会先处理函数声明。所以函数声明在其被声明的代码块内的任何位置都是可见的。
    //- 函数表达式在执行流程到达时创建。
    //- 在大多数情况下，当我们需要声明一个函数时，最好使用函数声明，因为函数在被声明之前也是可见的。