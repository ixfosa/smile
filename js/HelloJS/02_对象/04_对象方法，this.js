// 对象方法，"this"
    //- 通常创建对象来表示真实世界中的实体：
// 'use strict'
let user = {
  name: "ixfosa",
  age: 22
};

// 并且，在现实世界中，用户可以进行 操作：从购物车中挑选某物、登录和注销等。
// 在 JavaScript 中，行为（action）由属性中的函数来表示


// 方法示例
// 刚开始，我们来教 user 说 hello：
user.sayHi = function() {
    console.log("Hello JS");
}

user.sayHi();

// 使用预先声明的函数作为方法，就像这样：
// // 首先，声明函数
function sayHello() {
    console.log("Hello ixfosa");
}

// 然后将其作为一个方法添加
user.sayHello = sayHello;

user.sayHello();  



// -------------------------------------------------------------------------------------------



// 方法简写
let person1 = {
    sayHi: function() {
        console.log("Hello person1");
    }
}

let person2 = {
    sayHi() {
        console.log("Hello person2");   // 等同  sayHi: function() {  }
    }
}

person1.sayHi();
person2.sayHi();




// -------------------------------------------------------------------------------------------



// 方法中的 “this”
    // 在 JavaScript 中，this 是“自由”的，它的值是在调用时计算出来的，
    // 它的值并不取决于方法声明的位置，而是取决于在“点符号前”的是什么对象。

let user1 = {
    name: "ixfosa",
    age: 22,
    sayHi: function() {
        // "this" 指的是“当前的对象”
        console.log("Hi " + this.name);
    },
    sayHello: function() {
        // "this" 指的是“当前的对象”
        console.log("Hello " + user1.name);
    }
};

// user.sayHi() 执行过程中，this 的值是 user。
user1.sayHi();


// 技术上讲，也可以在不使用 this 的情况下，通过外部变量名来引用它：
user1.sayHello();

// 但这样的代码是不可靠的。如果将 user 复制给另一个变量，
// 例如 admin = user，并赋另外的值给 user，那么它将访问到错误的对象。

let admin = user1;
user1 = null;
admin.sayHi();

// this.name 替换 user.name，那么代码就会正常运行。
// admin.sayHello(); // TypeError: Cannot read property 'name' of null



// -------------------------------------------------------------------------------------------



// “this” 不受限制
    //- 在 JavaScript 中，this 关键字与其他大多数编程语言中的不同。
    //- JavaScript 中的 this 可以用于任何函数，即使它不是对象的方法。

// 下面这样的代码没有语法错误：
function sayHi() {
  console.log( this.name );
}

// this 的值是在代码运行时计算出来的，它取决于代码上下文。

// 例如，这里相同的函数被分配给两个不同的对象，在调用中有着不同的 “this” 值：
let student1 = {name: "long"};
let student2 = {name: "zhong"};

student1.sayHi = sayHi;
student2.sayHi = sayHi;

// 这两个调用有不同的 this 值
// 函数内部的 "this" 是“点符号前面”的那个对象
student1.sayHi();
student2.sayHi();




// 在没有对象的情况下调用：this == undefined

// 我们甚至可以在没有对象的情况下调用函数：
function say() {
    'use strict'
  console.log(this);
}

say(); // undefined  'use strict'下.

// 在这种情况下，严格模式下的 this 值为 undefined。如果我们尝试访问 this.name，将会报错。

// 在非严格模式的情况下，this 将会是 全局对象（浏览器中的 window）。这是一个历史行为，"use strict" 已经将其修复了。

// 通常这种调用是程序出错了。如果在一个函数内部有 this，那么通常意味着它是在对象上下文环境中被调用的。




// -------------------------------------------------------------------------------------------


// 箭头函数没有自己的 “this”
    //- 箭头函数有些特别：它们没有自己的 this。如果我们在这样的函数中引用 this，this 值取决于外部“正常的”函数。

// 举个例子，这里的 arrow() 使用的 this 来自于外部的 user.sayHi() 方法：

let user3 = {
    firstName: "fo",
    sayHi() {
        let arrow = () => console.log(this.firstName);
        arrow();
    },
   
    sayHello() {
        let that = this;
        function arrow() {
            // console.log(this.firstName); // TypeError: Cannot read property 'firstName' of undefined
            console.info("this " + this); // strict下 (this undefined); 非strict (this [object global])
            console.log(that.firstName);
        }
            arrow();
    },
};

user3.sayHi(); // fo
user3.sayHello(); 







// -------------------------------------------------------------------------------------------


// 总结
    //- 存储在对象属性中的函数被称为“方法”。
    //- 方法允许对象进行像 object.doSomething() 这样的“操作”。
    //- 方法可以将对象引用为 this。

    //- this 的值是在程序运行时得到的。
        //- 一个函数在声明时，可能就使用了 this，但是这个 this 只有在函数被调用时才会有值。
        //- 可以在对象之间复制函数。
        //- 以“方法”的语法调用函数时：object.method()，调用过程中的 this 值是 object。
        //- 请注意箭头函数有些特别：它们没有 this。在箭头函数内部访问到的 this 都是从外部获取的。

