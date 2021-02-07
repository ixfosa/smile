// 扩展内建类
    // 内建的类，例如 Array，Map 等也都是可以扩展的（extendable）。


// 例如，这里有一个继承自原生 Array 的类 PowerArray：
class PowerArray extends Array {
    isEmpty() {
        return this.length === 0;
    }
}

let arr = new PowerArray(1, 2, 3, 4, 5);
console.log(arr.isEmpty()); // false

let filteredArr = arr.filter( item => item > 5);
console.log(filteredArr.isEmpty()); // true


// 内建的方法例如 filter，map 等 — 返回的正是子类 PowerArray 的新对象。
// 它们内部使用了对象的 constructor 属性来实现这一功能。
console.log(arr.constructor === PowerArray); // true


// 当 arr.filter() 被调用时，它的内部使用的是 arr.constructor 来创建新的结果数组，
// 而不是使用原生的 Array。因为我们可以在结果数组上继续使用 PowerArray 的方法。

// 甚至，我们可以定制这种行为。

// 可以给这个类添加一个特殊的静态 getter Symbol.species。如果存在，
// 则应返回 JavaScript 在内部用来在 map 和 filter 等方法中创建新实体的 constructor。

// 如果希望像 map 或 filter 这样的内建方法返回常规数组，我们可以在 Symbol.species 中返回 Array，
// 就像这样：
class PowerArray1 extends Array {
    isEmpty() {
        return this.length === 0;
    }

     // 内建方法将使用这个作为 constructor
    static get [Symbol.species]() {
        return Array;   
    }
}

let arr1 = new PowerArray1(1, 2, 3, 4, 5);
console.log(arr1.isEmpty()); // false

// filter 使用 arr.constructor[Symbol.species] 作为 constructor 创建新数组
let filteredArr1 = arr1.filter( item => item > 5);

// // filteredArr 不是 PowerArray，而是 Array
// console.log(filteredArr1.isEmpty()); // TypeError: filteredArr1.isEmpty is not a function

// 现在 .filter 返回 Array。所以扩展的功能不再传递。





// -------------------------------------------------------------------------------------------



// 内建类没有静态方法继承
    // 内建对象有它们自己的静态方法，例如 Object.keys，Array.isArray 等。

    // ，原生的类互相扩展。例如，Array 扩展自 Object。

    // 通常，当一个类扩展另一个类时，静态方法和非静态方法都会被继承。

    // 但内建类却是一个例外。它们相互间不继承静态方法。
        // 例如，Array 和 Date 都继承自 Object，所以它们的实例都有来自 Object.prototype 的方法。
        // 但 Array.[[Prototype]] 并不指向 Object，所以它们没有例如 Array.keys()（或 Date.keys()）这些静态方法。

console.log(new Object().__proto__ === Object.prototype); // true
console.log(new Object().constructor === Object); // true

console.log(new Date().__proto__ === Date.prototype); // true
console.log( Date.prototype.__proto__  === Object.prototype); // true

console.log(Object.__proto__ === Function.prototype); // true

console.log(new Function().prototype.__proto__ === Object.prototype); // true

console.log(new Function().__proto__ === Function.prototype); // true





