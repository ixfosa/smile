// 函数对象，NFE
    //- 在 JavaScript 中，函数就是值。

    // JavaScript 中的每个值都有一种类型，那么函数是什么类型呢？
        //- 在 JavaScript 中，函数就是对象。

    // 一个容易理解的方式是把函数想象成可被调用的“行为对象（action object）”。
    // 我们不仅可以调用它们，还能把它们当作对象来处理：增/删属性，按引用传递等。

    // 命名函数表达式(NFE,Named Function Expression),指带有名字的函数表达式的术语。

// -------------------------------------------------------------------------------------------




// 属性 “name”
    //- 函数对象包含一些便于使用的属性。


// 比如，一个函数的名字可以通过属性 “name” 来访问：
function sayHi() {
  console.log("Hi");
}

console.log(sayHi.name); // sayHi
// 更有趣的是，名称赋值的逻辑很智能。即使函数被创建时没有名字，
// 名称赋值的逻辑也能给它赋予一个正确的名字，然后进行赋值：

let sayHii = function() {
  console.log("Hii");
};
console.log(sayHii.name); // sayHii（有名字！）

// 当以默认值的方式完成了赋值时，它也有效：
function f(sayHi = function() {}) {
  console.log(sayHi.name); // sayHi（生效了！）
}
f();
// 规范中把这种特性叫做「上下文命名」。如果函数自己没有提供，那么在赋值中，会根据上下文来推测一个。

// 对象方法也有名字：
let user = {
    sayHi() {
        // ...
    },

    sayBye: function() {
        // ...
    }
}

console.log(user.sayHi.name); // sayHi
console.log(user.sayBye.name); // sayBye

// 有时会出现无法推测名字的情况。此时，属性 name 会是空，像这样：
// 函数是在数组中创建的
let arr = [function() {}];

console.log( arr[0].name ); // <空字符串>
// 引擎无法设置正确的名字，所以没有值




// -------------------------------------------------------------------------------------------




// 属性 “length”
    //- 内置属性 “length”，它返回函数入参的个数，比如：

function f1(a) {}
function f2(a, b) {}
function many(a, b, ...more) {}

console.log(f1.length); // 1
console.log(f2.length); // 2
console.log(many.length); // 2   rest 参数不参与计数。


// 属性 length 有时在操作其它函数的函数中用于做 内省/运行时检查（introspection）。


// 比如，下面的代码中函数 ask 接受一个询问答案的参数 question 和可能包含任意数量 handler 的参数 ...handlers。
// 当用户提供了自己的答案后，函数会调用那些 handlers。我们可以传入两种 handlers：
    //- 一种是无参函数，它仅在用户回答给出积极的答案时被调用。
    //- 一种是有参函数，它在两种情况都会被调用，并且返回一个答案。

// 为了正确地调用 handler，我们需要检查 handler.length 属性。

// 我们的想法是，我们用一个简单的无参数的 handler 语法来处理积极的回答（最常见的变体），
// 但也要能够提供通用的 handler：

function ask(question, ...handlers) {

    let isYes = false;

    for(let handler of handlers) {
        if (handler.length == 0) {
            if (isYes) handler();
        } else {
            handler(isYes);
        }
    }
}

// 对于积极的回答，两个 handler 都会被调用
// 对于负面的回答，只有第二个 handler 被调用
ask("Question?", () => console.log('You said yes'), result => console.log(result)); // false

// 这种特别的情况就是所谓的 多态性 —— 根据参数的类型，或者根据在我们的具体情景下的 length 来做不同的处理。




// -------------------------------------------------------------------------------------------




// 自定义属性
    //- 我们也可以添加我们自己的属性。

// 这里我们添加了 counter 属性，用来跟踪总的调用次数：
function sayHi1() {
  console.log("Hi");

  // 计算调用次数
  sayHi.counter++;
}
sayHi.counter = 0; // 初始值
sayHi1(); // Hi
sayHi1(); // Hi
console.log( `Called ${sayHi.counter} times` ); // Called 2 times


// 属性不是变量
    //- 被赋值给函数的属性，比如 sayHi.counter = 0，不会 在函数内定义一个局部变量 counter。换句话说，
    //- 属性 counter 和变量 let counter 是毫不相关的两个东西。

// 我们可以把函数当作对象，在它里面存储属性，但是这对它的执行没有任何影响。
// 变量不是函数属性，反之亦然。它们之间是平行的。

// 函数属性有时会用来替代闭包。例如，我们可以使用函数属性将 变量作用域，闭包

function makeCounter() {
    // 不需要这个了
    // let count = 0

    function counter() {
        return counter.count++;
    };

  counter.count = 0;  // 现在 count 被直接存储在函数里，而不是它外部的词法环境。

  return counter;
}

let counter = makeCounter();
console.log( counter() ); // 0
console.log( counter() ); // 1


// 那么它和闭包谁好谁赖？
    //- 两者最大的不同就是如果 count 的值位于外层（函数）变量中，那么外部的代码无法访问到它，
    //- 只有嵌套的函数可以修改它。而如果它是绑定到函数的，那么就很容易：
    //- 所以，选择哪种实现方式取决于我们的需求是什么。

function makeCounter2() {

  function counter() {
    return counter.count++;
  };

  counter.count = 0;

  return counter;
}

let counter2 = makeCounter2();

counter.count = 10;
console.log( counter2() ); // 10






// -------------------------------------------------------------------------------------------




// 命名函数表达式
    //- 命名函数表达式（NFE，Named Function Expression），指带有名字的函数表达式的术语。

// 例如，让我们写一个普通的函数表达式：
let sayHi01 = function(who) {
  console.log(`Hello, ${who}`);
};

// 然后给它加一个名字：
let sayHi02 = function func(who) {  // 添加这个名字当然也没有打破任何东西。
    console.log(`Hello, ${who}`);
};
// 函数依然可以通过 sayHi02() 来调用：
sayHi02("ixfosa");  // Hello, ixfosa

// 我们这里得到了什么吗？为它添加一个 "func" 名字的目的是什么？
    //- 首先请注意，它仍然是一个函数表达式。在 function 后面加一个名字 "func" 没有使它成为一个函数声明，
    //- 因为它仍然是作为赋值表达式中的一部分被创建的。


// 关于名字 func 有两个特殊的地方，这就是添加它的原因：
    //- 它允许函数在内部引用自己。
    //- 它在函数外是不可见的。

// 例如，下面的函数 sayHi 会在没有入参 who 时，以 "Guest" 为入参调用自己：

let sayHi03 = function func03(who) {
  if (who) {
    console.log(`Hello, ${who}`);
  } else {
    func03("Guest"); // 使用 func 再次调用函数自身
  }
};

sayHi03(); // Hello, Guest

// 但这不工作：
//- func03(); // Error, func is not defined（在函数外不可见）


// 我们为什么使用 func 呢？为什么不直接使用 sayHi 进行嵌套调用
let sayHi04 = function(who) {
  if (who) {
    console.log(`Hello, ${who}`);
  } else {
    sayHi04("Guest");
  }
};
// 上面这段代码的问题在于 sayHi 的值可能会被函数外部的代码改变。

// 如果该函数被赋值给另外一个变量（译注：也就是原变量被修改），那么函数就会开始报错：
let sayHi05 = function(who) {
  if (who) {
    console.log(`Hello, ${who}`);
  } else {
    sayHi05("Guest"); // Error: sayHi is not a function
  }
};

let welcome = sayHi05;
sayHi05 = null;

// welcome(); // Error，嵌套调用 sayHi 不再有效！

// 发生这种情况是因为该函数从它的外部词法环境获取 sayHi。没有局部的 sayHi 了，所以使用外部变量。
// 而当调用时，外部的 sayHi 是 null。


// 我们给函数表达式添加的可选的名字，正是用来解决这类问题的。
let sayHi06 = function func06(who) {
  if (who) {
    console.log(`Hello, ${who}`);
  } else {
    func06("Guest"); // 现在一切正常
  }
};
welcome = sayHi06;
sayHi06 = null;
welcome(); // Hello, Guest（嵌套调用有效）

// 现在它可以正常运行了，因为名字 func 是函数局部域的。它不是从外部获取的（而且它对外部也是不可见的）。
// 规范确保它只会引用当前函数。

// 外部代码仍然有该函数的 sayHi 或 welcome 变量。而且 func 是一个“内部函数名”，
// 可用于函数在自身内部进行自调用。

// 函数声明没有这个东西
    //- 这里所讲的“内部名”特性只针对函数表达式，而不是函数声明。对于函数声明，没有用来添加“内部”名的语法。

// 有时，当我们需要一个可靠的内部名时，这就成为了你把函数声明重写成函数表达式的理由了。





// -------------------------------------------------------------------------------------------



// 总结
    // 函数就是对象。

    // 函数的一些属性：
        //- name —— 函数的名字。通常取自函数定义，但如果函数定义时没设定函数名，JavaScript 会尝试通过函数的上下文猜一个函数名（例如把赋值的变量名取为函数名）。
        //- length —— 函数定义时的入参的个数。Rest 参数不参与计数。

    // 如果函数是通过函数表达式的形式被声明的（不是在主代码流里），并且附带了名字，
    // 那么它被称为命名函数表达式（Named Function Expression）。
    // 这个名字可以用于在该函数内部进行自调用，例如递归调用等。

    // 此外，函数可以带有额外的属性。很多知名的 JavaScript 

    // 它们创建一个“主”函数，然后给它附加很多其它“辅助”函数。
    // 例如，jQuery 库创建了一个名为 $ 的函数。lodash 库创建一个 _ 函数，然后为其添加了 _.add、_.keyBy 以及其它属性（想要了解更多内容，参查阅 docs）。实际上，它们这么做是为了减少对全局空间的污染，这样一个库就只会有一个全局变量。这样就降低了命名冲突的可能性。

    // 所以，一个函数本身可以完成一项有用的工作，还可以在自身的属性中附带许多其他功能。





// -------------------------------------------------------------------------------------------




// 为 counter 添加 set 和 decrease 方法

// 修改 makeCounter() 代码，使得 counter 可以进行减一和设置值的操作：
   //- counter() 应该返回下一个数字（与之前的逻辑相同）。
   //- counter.set(value) 应该将 count 设置为 value。
   //- counter.decrease(value) 应该把 count 减 1。

// P.S. 你可以使用闭包或者函数属性来保持当前的计数，或者两种都写。


function myMakeCounter() {
    
    function counter() {
        return counter.count++;
    };

    counter.count = 0;

    counter.set = function(value) {
        counter.count = value;
    }

    counter.decrease = () => count--;

    return counter;
    
}
counter = myMakeCounter();
console.log(counter());
console.log(counter());
console.log(counter());


// -------------------------------------------------------------------------------------------




// 任意数量的括号求和

// 写一个函数 sum，它有这样的功能：
    //- sum(1)(2) == 3; // 1 + 2
    //- sum(1)(2)(3) == 6; // 1 + 2 + 3
    //- sum(5)(-1)(2) == 6
    //- sum(6)(-1)(-2)(-3) == 0
    //- sum(0)(1)(2)(3)(4)(5) == 15

// P.S. 提示：你可能需要创建    来为你的函数提供基本类型转换。
"use strict";
function sum(a) {

    let currentSum = a;
  
    function f(b) {
        currentSum += b;
        return f;
    }
  
    f.toString = function() {
        return currentSum;
    };
  
    return f;
}
  
console.log( sum(1)(2) ); // 3
console.log( sum(5)(-1)(2) ); // 6
console.log( sum(6)(-1)(-2)(-3) ); // 0
console.log( sum(0)(1)(2)(3)(4)(5) ); // 15