
// 类扩展自对象？

// 正如我们所知道的，所有的对象通常都继承自 Object.prototype，
// 并且可以访问“通用”对象方法，例如 hasOwnProperty 等。


// 例如：
class Rabbit {
    constructor(name) {
        this.name = name;
  }
}

let rabbit = new Rabbit("Rab");

// hasOwnProperty 方法来自于 Object.prototype
console.log( rabbit.hasOwnProperty('name') ); // true

// 但是，如果我们像这样 "class Rabbit extends Object" 把它明确地写出来，
// 那么结果会与简单的 "class Rabbit" 有所不同么？

// 不同之处在哪里？


// 下面是此类的示例代码（它无法正常运行 — 为什么？修复它？）：
// constructor 增加 super();
class Rabbit1 extends Object {
    constructor(name) {
        super();   // 需要在继承时调用父类的 constructor
        this.name = name;
    }
}

let rabbit1 = new Rabbit1("Rab");

console.log( rabbit1.hasOwnProperty('name') ); 



// 但这还不是全部原因。
    // 即便修复了它，"class Rabbit extends Object" 和 class Rabbit 之间仍存在着重要差异。

// “extends” 语法会设置两个原型：
    // 1. 在构造函数的 "prototype" 之间设置原型（为了获取实例方法）。
    // 2. 在构造函数之间会设置原型（为了获取静态方法）。

// 在我们的例子里，对于 class Rabbit extends Object，它意味着：

console.log( Rabbit1.prototype.__proto__ === Object.prototype ); // (1) true
console.log( Rabbit1.__proto__ === Object ); // (2) true

// 所以，现在 Rabbit 可以通过 Rabbit 访问 Object 的静态方法，像这样：

// class Rabbit extends Object {}

// 通常我们调用 Object.getOwnPropertyNames
console.log ( Rabbit1.getOwnPropertyNames({a: 1, b: 2})); // a,b


// 但是如果没有 extends Object，那么 Rabbit.__proto__ 将不会被设置为 Object。
// 下面是示例：
class Rabbit2 {}
console.log( Rabbit2.prototype.__proto__ === Object.prototype ); // (1) true
console.log( Rabbit2.__proto__ === Object ); // (2) false (!)
console.log( Rabbit2.__proto__ === Function.prototype ); // true，所有函数都是默认如此

// error，Rabbit 中没有这样的函数
// TypeError: Rabbit2.getOwnPropertyNames is not a function
// console.log( Rabbit2.getOwnPropertyNames({a: 1, b: 2})); // Error

// 所以，在这种情况下，Rabbit 没有提供对 Object 的静态方法的访问。

// 顺便说一下，Function.prototype 有一些“通用”函数方法，例如 call 和 bind 等。
// 在上述的两种情况下它们都是可用的，因为对于内建的 Object 构造函数而言，
// Object.__proto__ === Function.prototype。
console.log(Object.__proto__ === Function.prototype); // true