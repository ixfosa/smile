// 类检查："instanceof"
    // instanceof 操作符用于检查一个对象是否属于某个特定的 class。同时，它还考虑了继承。

    // 在许多情况下，可能都需要进行此类检查。
    // 例如，它可以被用来构建一个 多态性（polymorphic） 的函数，该函数根据参数的类型对参数进行不同的处理。



// instanceof 操作符
    // 语法：
        //- obj instanceof Class
            //- 如果 obj 隶属于 Class 类（或 Class 类的衍生类），则返回 true。

// 例如:
class Rabbit {}
let rabbit = new Rabbit();
// // rabbit 是 Rabbit class 的对象吗？
console.log(rabbit instanceof Rabbit); // true


// 它还可以与构造函数一起使用：
function Rabbit1() {}   // 这里是构造函数，而不是 class
console.log(new Rabbit1() instanceof Rabbit1); // true

// 诸如 Array 之类的内建 class 一起使用：
let arr = [1, 2, 3];
console.log(arr instanceof Array);  // true
console.log(arr instanceof Object); // true


// arr 同时还隶属于 Object 类。因为从原型上来讲，Array 是继承自 Object 的。

// 通常，instanceof 在检查中会将原型链考虑在内。
// 此外，还可以在静态方法 Symbol.hasInstance 中设置自定义逻辑。


// obj instanceof Class 算法的执行过程大致如下：

// 1. 如果这儿有静态方法 Symbol.hasInstance，那就直接调用这个方法：
class Animal {
    static [Symbol.hasInstance](obj) {
        if (obj.canEat) return true;
    }
}

let obj = { canEat: true };

console.log(obj instanceof Animal); //true


// 2. 大多数 class 没有 Symbol.hasInstance。
//    在这种情况下，标准的逻辑是：
//    使用 obj instanceOf Class 检查 Class.prototype 是否等于 obj 的原型链中的原型之一。

// 换句话说就是，一个接一个地比较：
/*
    obj.__proto__ === Class.prototype?
    obj.__proto__.__proto__ === Class.prototype?
    obj.__proto__.__proto__.__proto__ === Class.prototype?
    ...
    // 如果任意一个的答案为 true，则返回 true
    // 否则，如果我们已经检查到了原型链的尾端，则返回 false
*/



// 而在继承的例子中，匹配将在第二步进行：
class Animal2 {}
class Rabbit2 extends Animal2 {}

console.log( new Rabbit2() instanceof Animal2); // true

console.log(new Rabbit2().__proto__ === Rabbit2.prototype); // true
console.log(new Rabbit2().__proto__.__proto__ === Animal2.prototype); // （匹配！）


// 这里还要提到一个方法 objA.isPrototypeOf(objB)，如果 objA 处在 objB 的原型链中，则返回 true。
// 所以，可以将 obj instanceof Class 检查改为 Class.prototype.isPrototypeOf(obj)。

// 但是 Class 的 constructor 自身是不参与检查的！检查过程只和原型链以及 Class.prototype 有关。

// 创建对象后，如果更改 prototype 属性，可能会导致有趣的结果。

// 就像这样：
class Rabbit3  {}
let rabbit3 = new Rabbit3();
console.log(rabbit3 instanceof Rabbit3); // true

Rabbit3.prototype = {};
console.log(rabbit3 instanceof Rabbit3);  // true
 

function Rabbit4() {}
let rabbit4 = new Rabbit4();
console.log(rabbit4 instanceof Rabbit4); // true

Rabbit4.prototype = {};
console.log(rabbit4 instanceof Rabbit4);  // false 





// -------------------------------------------------------------------------------------------






// 使用 Object.prototype.toString 方法来揭示类型

// 一个普通对象被转化为字符串时为 [object Object]：
/*
let o = {};
alert(obj); // [object Object]
alert(obj.toString()); // 同上
*/

// 这是通过 toString 方法实现的。但是这儿有一个隐藏的功能，该功能可以使 toString 实际上比这更强大。
// 我们可以将其作为 typeof 的增强版或者 instanceof 的替代方法来使用。


// 按照 规范 所讲，内建的 toString 方法可以被从对象中提取出来，并在任何其他值的上下文中执行。
// 其结果取决于该值。
    //- 对于 number 类型，结果是 [object Number]
    //- 对于 boolean 类型，结果是 [object Boolean]
    //- 对于 null：[object Null]
    //- 对于 undefined：[object Undefined]
    //- 对于数组：[object Array]
    //- ……等（可自定义）


// 演示:
// 将 toString 方法复制到一个变量中
let objectToString  = Object.prototype.toString;

// 它是什么类型的？
let array = [];

//  call 方法来在上下文 this=arr 中执行函数 objectToString。
console.log(objectToString.call(array)); // [object Array]


// 在内部，toString 的算法会检查 this，并返回相应的结果。再举几个例子：
console.log( objectToString.call(123) ); // [object Number]
console.log( objectToString.call(null) ); // [object Null]
console.log( objectToString.call(console.log) ); // [object Function]





// -------------------------------------------------------------------------------------------




// Symbol.toStringTag

// 可以使用特殊的对象属性 Symbol.toStringTag 自定义对象的 toString 方法的行为。
let User = {
    [Symbol.toStringTag] : "User"
};

// 对于大多数特定于环境的对象，都有一个这样的属性。下面是一些特定于浏览器的示例：
// 特定于环境的对象和类的 toStringTag：

    //- alert( window[Symbol.toStringTag]); // Window
    //- alert( XMLHttpRequest.prototype[Symbol.toStringTag] ); // XMLHttpRequest

    //- alert( {}.toString.call(window) ); // [object Window]
    //- alert( {}.toString.call(new XMLHttpRequest()) ); // [object XMLHttpRequest]


// 输出结果恰好是 Symbol.toStringTag（如果存在），只不过被包裹进了 [object ...] 里。

// 这样一来，就有了个“ typeof ”，不仅能检查原始数据类型，而且适用于内建对象，更可贵的是还支持自定义。

// 所以，如果想要获取内建对象的类型，并希望把该信息以字符串的形式返回，而不只是检查类型的话，
// 我们可以用 {}.toString.call 替代 instanceof。






// -------------------------------------------------------------------------------------------




// 总结
    // 类型检查方法：

/*
                    用于	                                                     返回值
    typeof	        原始数据类型	                                             string
    {}.toString	    原始数据类型，内建对象，包含 Symbol.toStringTag 属性的对象	    string
    instanceof	    对象	                                                     true/false
*/


// 从技术上讲，{}.toString 是一种“更高级的” typeof。

// 当使用类的层次结构（hierarchy），并想要对该类进行检查，同时还要考虑继承时，
// 这种场景下 instanceof 操作符确实很出色。




// -------------------------------------------------------------------------------------------




// 不按套路出牌的 instanceof

// 在下面的代码中，为什么 instanceof 会返回 true？我们可以明显看到，a 并不是通过 B() 创建的。

function A() {}
function B() {}

A.prototype = B.prototype = {};

let a = new A();

console.log( a instanceof B ); // true


// instanceof 并不关心函数，而是关心函数的与原型链匹配的 prototype。

// 并且，这里 a.__proto__ == B.prototype，所以 instanceof 返回 true。

// 总之，根据 instanceof 的逻辑，真正决定类型的是 prototype，而不是构造函数。







