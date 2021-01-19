
// “箭头函数”: 
    //- let func = (arg1, arg2, ...argN) => expression
    //- 创建了一个函数 func，它接受参数 arg1..argN，然后使用参数对右侧的 expression 求值并返回其结果。

    // 换句话说，它是下面这段代码的更短的版本：

    //- let func = function(arg1, arg2, ...argN) {
    //-     return expression;
    //- };




let sum1 = (a, b) => a + b;

// 这个箭头函数是下面这个函数的更短的版本：
let sum2 = function(a, b) {
  return a + b;
};

console.log( sum1(1, 2) ); // 


// 如果我们只有一个参数，还可以省略掉参数外的圆括号，使代码更短。
let isLove = love => love;
console.log(isLove(true));

// 如果没有参数，括号将是空的（但括号应该保留）：
let sayHi = () => "Hello";
console.log(sayHi());


// 箭头函数可以像函数表达式一样使用。
// 例如，动态创建一个函数：
let age = 18;

let welcome = (age < 18) ?
  () => console.log('Hello') :
  () => console.log("Greetings!");

welcome();


// 多行的箭头函数
    //- 上面的例子从 => 的左侧获取参数，然后使用参数计算右侧表达式的值。
    //- 但有时我们需要更复杂一点的东西，比如多行的表达式或语句。这也是可以做到的，但是我们应该用花括号括起来
    //- 然后使用一个普通的 return 将需要返回的值进行返回。

let sum3 = (a, b) => {  // 花括号表示开始一个多行函数
  let result = a + b;
  return result; // 如果我们使用了花括号，那么我们需要一个显式的 “return”
};

console.log( sum3(1, 2) ); // 3


// 总结
    //- 对于一行代码的函数来说，箭头函数是相当方便的。它具体有两种：
    //- 不带花括号：(...args) => expression — 右侧是一个表达式：函数计算表达式并返回其结果。
    //- 带花括号：(...args) => { body } — 花括号允许我们在函数中编写多个语句，但是我们需要显式地 return 来返回一些内容。




// 用箭头函数重写下面的函数表达式：
function ask(question, yes, no) {
    if (question) yes()
    else no();
}

ask(
    true,
    function() { console.log("You agreed."); },
    function() { console.log("You canceled the execution."); }
);

let ask2 = (question, yes, no) => {
    if (question) yes();
    else no();
};

ask2(
    true, 
    () => console.log("YES"),
    () => console.log("NO")
)