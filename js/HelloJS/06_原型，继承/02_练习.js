// 修改 "prototype"

//  new Rabbit，然后尝试修改它的 prototype。

// 最初，我们有以下代码：
function Rabbit() {}
Rabbit.prototype = {
    eats: true
};

let rabbit = new Rabbit();

console.log( rabbit.eats ); // true




// 1. 我们增加了一个字符串（强调）。现在 console.log 会显示什么？
function Rabbit1() {}
Rabbit1.prototype = {
    eats: true
};

let rabbit1 = new Rabbit1();

Rabbit1.prototype = {};

console.log( rabbit1.eats ); // ?  true 

// Rabbit1.prototype 的赋值操作为新对象设置了 [[Prototype]]，但它不影响已有的对象。



// 2. 如果代码是这样的（修改了一行）？
function Rabbit2() {}
Rabbit2.prototype = {
    eats: true
};

let rabbit2 = new Rabbit2();

Rabbit2.prototype.eats = false;

console.log( rabbit2.eats ); // ? false

// 对象通过引用被赋值。来自 Rabbit2.prototype 的对象并没有被复制，
// 它仍然是被 Rabbit2.prototype 和 rabbit2 的 [[Prototype]] 引用的单个对象。

// 所以当我们通过一个引用更改其内容时，它对其他引用也是可见的。




// 3. 像这样呢（修改了一行）？
function Rabbit3() {}
Rabbit3.prototype = {
  eats: true
};

let rabbit3 = new Rabbit3();

delete rabbit3.eats;

console.log( rabbit3.eats ); // ? true

// 所有 delete 操作都直接应用于对象。这里的 delete rabbit3.eats 试图从 rabbit 中删除 eats 属性，
// 但 rabbit3 对象并没有 eats 属性。所以这个操作不会有任何影响。





// 4. 最后一种变体：
function Rabbit4() {}
Rabbit4.prototype = {
  eats: true
};

let rabbit4 = new Rabbit4();

delete Rabbit4.prototype.eats;

console.log( rabbit4.eats ); // ? undefined
// 属性 eats 被从 prototype 中删除，prototype 中就没有这个属性了。







// -------------------------------------------------------------------------------------------





// 使用相同的构造函数创建一个对象

// 想象一下，我们有一个由构造函数创建的对象 obj 
// —— 我们不知道使用的是哪个构造函数，但是我们想使用它创建一个新对象。

// 我们可以这样做吗？

// let obj2 = new obj.constructor();

// 请给出一个可以使这样的代码正常工作的 obj 的构造函数的例子。再给出会导致这样的代码无法正确工作的例子。


// 代码正常工作的 obj 的构造函数的例子
function Obj() {
    console.log("new Obj");
}

let obj = new Obj(); // new Obj
let obj2 = new obj.constructor(); // new Obj
console.log(obj == obj2); // false
console.log(obj2.constructor === Obj); // true


// 代码无法正确工作的例子
function Obj2() {
    console.log("new Obj2");
}

Obj2.prototype = {};

let obj3 = new Obj2(); // new Obj

let obj4 = new obj3.constructor();  // Nothing output.

console.log(obj4.constructor === Obj2); // false
console.log(obj3 == obj4); // false