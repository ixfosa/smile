

// -------------------------------------------------------------------------------------------


// “class” 语法


/*
    基本语法是：
        class MyClass {
            // class 方法
            constructor() { ... }
            method1() { ... }
            method2() { ... }
            method3() { ... }
            ...
        }

    然后使用 new MyClass() 来创建具有上述列出的所有方法的新对象。
    new 会自动调用 constructor() 方法，因此我们可以在 constructor() 中初始化对象。
*/



// 例如： 
class User {
    constructor(name) {
        this.name = name;
    }

    sayHi() {
        console.log("Hello " + this.name);
    }
}

let user = new User("Zi Qing");

user.sayHi();  // Hello Zi Qing


// 类的方法之间没有逗号
    //- 不要把这里的符号与对象字面量相混淆。在类中，不需要逗号。





// -------------------------------------------------------------------------------------------




// 什么是 class？
    // 所以，class 到底是什么？
        //- 在 JavaScript 中，类是一种函数。

// 看看下面这段代码

// class 是一个函数
console.log(typeof User); // function


// class User {...} 构造实际上做了如下的事儿：
    //- 1. 创建一个名为 User 的函数，该函数成为类声明的结果。该函数的代码来自于 constructor 方法
    // （如果我们不编写这种方法，那么它就被假定为空）。

    //- 2. 存储类中的方法，例如 User.prototype 中的 sayHi。

// 当 new User 对象被创建后，当我们调用其方法时，它会从原型中获取对应的方法，
// 因此，对象 new User 可以访问类中的方法。


// 可以将 class User 声明的结果解释为：

// 更确切地说，是 constructor 方法
console.log(User === User.prototype.constructor); // true

// 方法在 User.prototype 中，例如：
User.prototype.sayHi(); // Hello undefined

// 在原型中实际上有两个方法
console.log(Object.getOwnPropertyNames(User.prototype)); // [ 'constructor', 'sayHi' ]






// -------------------------------------------------------------------------------------------




// 不仅仅是语法糖
    //- 人们常说 class 是一个语法糖（旨在使内容更易阅读，但不引入任何新内容的语法），
    // 因为我们实际上可以在没有 class 的情况下声明相同的内容：

// 用纯函数重写 class User
function User1(name) {
    this.name = name;
}
// 函数的原型（prototype）默认具有 "constructor" 属性，
// 所以，我们不需要创建它

// 2. 将方法添加到原型
User1.prototype.sayHi = function() {
    console.log("Hello " + this.name);
};

let user1 = new User1("Zi Qing");
user1.sayHi(); // Hello Zi Qing


// 这个定义的结果与使用类得到的结果基本相同。
// 因此，这确实是将 class 视为一种定义构造器及其原型方法的语法糖的理由。


// 尽管，它们之间存在着重大差异：
    // 1. 首先，通过 class 创建的函数具有特殊的内部属性标记 [[FunctionKind]]:"classConstructor"。
    //    因此，它与手动创建并不完全相同。

/*
    // 编程语言会在许多地方检查该属性。例如，与普通函数不同，必须使用 new 来调用它：
    class User {
        constructor() {}
    }
        
    alert(typeof User); // function

    User(); // Error: Class constructor User cannot be invoked without 'new'
    此外，大多数 JavaScript 引擎中的类构造器的字符串表示形式都以 “class…” 开头
        
    class User {
        constructor() {}
    }
    alert(User); // class User { ... }
*/

    // 2. 类方法不可枚举。 类定义将 "prototype" 中的所有方法的 enumerable 标志设置为 false。
    //    这很好，因为如果我们对一个对象调用 for..in 方法，我们通常不希望 class 方法出现。

    // 3. 类总是使用 use strict。 在类构造中的所有代码都将自动进入严格模式。





// -------------------------------------------------------------------------------------------





// 类表达式
    // 就像函数一样，类可以在另外一个表达式中被定义，被传递，被返回，被赋值等。

// 这是一个类表达式的例子：
let pserson = class {
    sayHi() {
        console.log("Hello js");
    }
}

// 类似于命名函数表达式（Named Function Expressions），类表达式可能也应该有一个名字。
// 如果类表达式有名字，那么该名字仅在类内部可见：

// “命名类表达式（Named Class Expression）”
// (规范中没有这样的术语，但是它和命名函数表达式类似)
let pserson2 = class Pserson2 {
    sayHi() {
        console.log(Pserson2); // MyClass 这个名字仅在类内部可见
    }
}

new pserson2().sayHi(); // [class Pserson2]

// console.log(Pserson2); // error，MyClass 在外部不可见 ReferenceError: Pserson2 is not defined


// 甚至可以动态地“按需”创建类，就像这样：
function makeClass(phrase) {
     // 声明一个类并返回它
    return class {
        sayHi() {
            console.log(phrase);
        }
    }
}

let myClass = makeClass("Hello");

new myClass().sayHi(); // Hello





// -------------------------------------------------------------------------------------------




// Getters/setters
    // 就像对象字面量，类可能包括 getters/setters，计算属性（computed properties）等。

// 这是一个使用 get/set 实现 user.name 的示例：
class User3 {
    constructor(name) {
         // 调用 setter
        this.name = name;
    }

    get name() {
        return this._name;
    }

    set name(value) {
        if (value.length < 4) {
            console.log("Name is too short.");
            return
        }
        this._name = value;
    }
}

let user3 = new User3("Zi Qing");

console.log(user3.name); // Zi Qing

user3 = new User3("fo"); // Name is too short.

// 从技术上来讲，这样的类声明可以通过在 User.prototype 中创建 getters 和 setters 来实现。





// -------------------------------------------------------------------------------------------




// 计算属性名称 […]

// 这里有一个使用中括号 [...] 的计算方法名称示例：
class User4 {

  ['say' + 'Hi']() {
    console.log("Hello");
  }

}

new User4().sayHi(); // Hello

// 它们和对象字面量类似。





// -------------------------------------------------------------------------------------------





// Class 字段
    // 旧的浏览器可能需要 polyfill, 类字段（field）是最近才添加到语言中的。
    // 之前，我们的类仅具有方法。

    // “类字段”是一种允许添加任何属性的语法。

// 例如，让我们在 class User 中添加一个 name 属性：    
class User5 {
    name = "Zi Qing";

    sayHi() {
        console.log("Hello " + this.name);
    }
}

let user5 = new User5();
user5.sayHi();  // Hello Zi Qing

// 所以，我们就只需在表达式中写 " = "，就这样。
// 类字段重要的不同之处在于，它们会在每个独立对象中被设好，而不是设在 User.prototype：
console.log(user5.name); //Zi Qing
console.log(User5.prototype.name); // undefined


// 我们也可以在赋值时使用更复杂的表达式和函数调用：
    //-     class User {
    //-         name = prompt("Name, please?", "John");
    //-     }
    //-     
    //-     let user = new User();
    //-     alert(user.name); // John





// -------------------------------------------------------------------------------------------





// 使用类字段制作绑定方法
    // JavaScript 中的函数具有动态的 this。它取决于调用上下文。

    // 因此，如果一个对象方法被传递到某处，或者在另一个上下文中被调用，则 this 将不再是对其对象的引用。

// 例如，此代码将显示 undefined：
class Button {
    constructor(value) {
        this.value = value;
    }

    click() {
        console.log(this.value);
    }
}

let button = new Button("hello");

// 这个问题被称为“丢失 this”。
setTimeout(button.click, 2000); // undefined

// 修复它的方式：
    //- 传递一个包装函数，例如 setTimeout(() => button.click(), 1000)。
    //- 将方法绑定到对象，例如在 constructor 中。

setTimeout( () => button.click(), 2000 ); // hello

// 类字段提供了另一种非常优雅的语法：
class Button2 {
    constructor(value) {
        this.value = value;
    }

    click = () => {
        console.log(this.value);
    }
}

let button2 = new Button2("smile");
setTimeout(button2.click, 2000) // smile


// 类字段 click = () => {...} 是基于每一个对象被创建的，在这里对于每一个 Button 对象都有一个独立的方法，
// 在内部都有一个指向此对象的 this。我们可以把 button.click 传递到任何地方，而且 this 的值总是正确的。

// 在浏览器环境中，它对于进行事件监听尤为有用。





// -------------------------------------------------------------------------------------------





// 总结


// 基本的类语法看起来像这样：

/*
class MyClass {

    prop = value; // 属性

    constructor(...) { // 构造器
        // ...
    }

    method(...) {} // method

    get something(...) {} // getter 方法
    set something(...) {} // setter 方法

    [Symbol.iterator]() {} // 有计算名称（computed name）的方法（此处为 symbol）
    // ...
}

*/

// 技术上来说，MyClass 是一个函数（我们提供作为 constructor 的那个），
// 而 methods、getters 和 settors 都被写入了 MyClass.prototype。







// -------------------------------------------------------------------------------------------




// 重写为 class

// Clock 类（请见沙箱）是以函数式编写的。请以 “class” 语法重写它。

// P.S. 时钟在控制台（console）中滴答，打开控制台即可查看。 







/*
class Clock {
    count = 0;

    clock() {
        setInterval(() => console.log(this.count++), 1000);
    }
}

new Clock().clock();
*/



class Clock {
    constructor({ template }) {
        this.template = template;
    }
  
    render() {
        let date = new Date();
    
        let hours = date.getHours();
        if (hours < 10) hours = '0' + hours;
    
        let mins = date.getMinutes();
        if (mins < 10) mins = '0' + mins;
    
        let secs = date.getSeconds();
        if (secs < 10) secs = '0' + secs;
    
        let output = this.template
            .replace('h', hours)
            .replace('m', mins)
            .replace('s', secs);
    
        console.log(output);
    }
  
    stop() {
        clearInterval(this.timer);
    }
  
    start() {
        this.render();
        this.timer = setInterval(() => this.render(), 1000);
    }
  }
  
  
let clock = new Clock({template: 'h:m:s'});
clock.start();