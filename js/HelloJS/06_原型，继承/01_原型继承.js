// 原型继承
    // 在编程中，我们经常会想获取并扩展一些东西。

    // 例如，我们有一个 user 对象及其属性和方法，并希望将 admin 和 guest 作为基于 user 稍加修改的变体。
    // 我们想重用 user 中的内容，而不是复制/重新实现它的方法，而只是在其之上构建一个新的对象。

// 原型继承（Prototypal inheritance） 这个语言特性能够帮助我们实现这一需求。



// [[Prototype]]
    // 在 JavaScript 中，对象有一个特殊的隐藏属性 [[Prototype]]（如规范中所命名的），
    // 它要么为 null，要么就是对另一个对象的引用。该对象被称为“原型”：
/* 
    prototype object
        ^
        |  [[Prototype]]
      object
*/

// 当我们从 object 中读取一个缺失的属性时，JavaScript 会自动从原型中获取该属性。
// 在编程中，这种行为被称为“原型继承”。

// 属性 [[Prototype]] 是内部的而且是隐藏的，但是这儿有很多设置它的方式。

// 其中之一就是使用特殊的名字 __proto__，就像这样：

let animal = {
    eats: true
};

let cat = {
    jumps: false 
};

cat.__proto__ = animal; // (*)   // 设置 cat.[[Prototype]] = animal

// 如果从 rabbit 中读取一个它没有的属性，JavaScript 会自动从 animal 中获取。
console.log(cat.eats); // true
console.log(cat.jumps); // false


// 这里的 (*) 行将 animal 设置为 rabbit 的原型。

// 当 console.log 试图读取 cat.eats (**) 时，因为它不存在于 cat 中，
// 所以 JavaScript 会顺着 [[Prototype]] 引用，在 animal 中查找（自下而上）


// 在这儿可以说 "animal 是 cat 的原型"，或者说 "cat 的原型是从 animal 继承而来的"。

// 因此，如果 animal 有许多有用的属性和方法，那么它们将自动地变为在 cat 中可用。这种属性被称为“继承”。



// 如果我们在 animal 中有一个方法，它可以在 cat 中被调用：
let animal2 = {
    eats: true,
    sleep() {
        console.log("animal sleep...");
    }
};

let cat2 = {
    jumps: false,
    __proto__ : animal2
};

// sleep 方法是从原型中获得的
cat2.sleep(); // animal sleep...



// 原型链可以很长： 
let grandfather = {
    surname: "xia",
};

let father = {
    name: "xin",
    __proto__: grandfather
};


let son = {
    name: "fo",
    __proto__: father
};

console.log(son.surname); // xia
console.log(son.name); // fo


// 现在，如果我们从 son 中读取一些它不存在的内容，JavaScript 会先在 father 中查找，然后在 grandfather 中查找。

// 这里只有两个限制：
    //- 引用不能形成闭环。如果我们试图在一个闭环中分配 __proto__，JavaScript 会抛出错误。
    //- __proto__ 的值可以是对象，也可以是 null。而其他的类型都会被忽略。

//- 只能有一个 [[Prototype]]。一个对象不能从其他两个对象获得继承。



// __proto__ 是 [[Prototype]] 的因历史原因而留下来的 getter/setter
    // __proto__ 和 [[Prototype]] 的区别
    // 请注意，__proto__ 与内部的 [[Prototype]] 不一样。 __proto__ 是 [[Prototype]] 的 getter/setter。
    // __proto__ 属性有点过时了。它的存在是出于历史的原因，
    // 现代编程语言建议我们应该使用函数 Object.getPrototypeOf/Object.setPrototypeOf 
    // 来取代 __proto__ 去 get/set 原型。

    // 根据规范，__proto__ 必须仅受浏览器环境的支持。但实际上，包括服务端在内的所有环境都支持它，
    // 因此我们使用它是非常安全的。




// -------------------------------------------------------------------------------------------




// 写入不使用原型
    // 原型仅用于读取属性。

    // 对于写入/删除操作可以直接在对象上进行。

// 在下面的示例中，我们将为 cat 分配自己的 walk：
let animal3 = {
  eats: true,
  walk() {
    /* cat 不会使用此方法 */
    console.log("animal walk...");
  }
};

let cat3 = {
  __proto__: animal3
};

cat3.walk = function() {
  console.log("cat walk!");
};

// cat3.walk() 将立即在对象中找到该方法并执行，而无需使用原型
cat3.walk(); // cat walk!



// 访问器（accessor）属性是一个例外，因为分配（assignment）操作是由 setter 函数处理的。
// 因此，写入此类属性实际上与调用函数相同。

// 也就是这个原因，所以下面这段代码中的 admin.fullName 能够正常运行：
let user = {
    name: "Fosong",
    surname: "Xia",

    set fullName(value) {
        [this.name, this.surname] = value.split(" ");
    },

    get fullName() {
        return this.surname + " " + this.name;
    }
};

let admin = {
    __proto__: user,
    isAdmin: true
};

console.log(admin.fullName); // Xia Fosong 属性 admin.fullName 在原型 user 中有一个 getter，因此它会被调用。
 
admin.fullName = "Qing Zi";  // 属性在原型中有一个 setter，因此它会被调用。
console.log(admin.fullName); // Zi Qing
console.log(user.fullName); // Xia Fosong


user.fullName = "Qing Zi";
console.log(user.fullName); // Zi Qing




// -------------------------------------------------------------------------------------------






// “this” 的值
    // 在上面的例子中可能会出现一个有趣的问题：在 set fullName(value) 中 this 的值是什么？
    // 属性 this.name 和 this.surname 被写在哪里：在 user 还是 admin？

    // 答案很简单：this 根本不受原型的影响。

    // 无论在哪里找到方法：在一个对象还是在原型中。在一个方法调用中，this 始终是点符号 . 前面的对象。

    // 因此，setter 调用 admin.fullName= 使用 admin 作为 this，而不是 user。

    // 这是一件非常重要的事儿，因为我们可能有一个带有很多方法的大对象，并且还有从其继承的对象。
    // 当继承的对象运行继承的方法时，它们将仅修改自己的状态，而不会修改大对象的状态。


// 例如，这里的 animal 代表“方法存储”，cat 在使用其中的方法。
// 调用 cat.sleep() 会在 cat 对象上设置 this.isSleeping：

// animal 有一些方法
let animal4 = {
  walk() {
    if (!this.isSleeping) {
      console.log(`I walk`);
    }
  },
  sleep() {
    this.isSleeping = true;
  }
};

let cat4 = {
  name: "White cat",
  __proto__: animal4
};

// 修改 cat.isSleeping
cat4.sleep();

console.log(cat4.isSleeping); // true
console.log(animal4.isSleeping); // undefined（原型中没有此属性）



// 如果我们还有从 animal 继承的其他对象，像 bird 和 snake 等，它们也将可以访问 animal 的方法。
// 但是，每个方法调用中的 this 都是在调用时（点符号前）评估的对应的对象，而不是 animal。
// 因此，当我们将数据写入 this 时，会将其存储到这些对象中。

// 所以，方法是共享的，但对象状态不是。




// -------------------------------------------------------------------------------------------




// for…in 循环
    //- for..in 循环也会迭代继承的属性。

// 例如：
let animal5 = {
  eats: true
};

let cat5 = {
  jumps: false,
  __proto__: animal5
};

// Object.keys 只返回自己的 key
console.log(Object.keys(cat5)); // [ 'jumps' ]

// for..in 会遍历自己以及继承的键
for(let prop in cat5) console.log(prop); // jumps，eats



// 如果这不是我们想要的，并且我们想排除继承的属性，那么这儿有一个内建方法:
    //- obj.hasOwnProperty(key)：如果 obj 具有自己的（非继承的）名为 key 的属性，则返回 true。

// 因此，我们可以过滤掉继承的属性（或对它们进行其他操作）：
for (let prop in cat5) {
    let isOwn = cat5.hasOwnProperty(prop);

    if(isOwn) {
        console.log("own: " + prop); // own: jumps
    } else {
        console.log("other: " + prop); // other: eats
    }
}


// cat 从 animal 中继承，animal 从 Object.prototype 中继承
// （因为 animal 是对象字面量 {...}，所以这是默认的继承），然后再向上是 null
console.log(cat5.__proto__ === animal5); // true
console.log(animal5.__proto__); // [Object: null prototype] {}


// 方法 rabbit.hasOwnProperty 来自哪儿？
// rabbit([[Prototype]]) --->  animal([[Prototype]])  --->  Object([[Prototype]])  ---> null
// 我们并没有定义它。该方法是 Object.prototype.hasOwnProperty 提供的。换句话说，它是继承的。


// 如果 for..in 循环会列出继承的属性，
// 那为什么 hasOwnProperty 没有像 eats 和 jumps 那样出现在 for..in 循环中？
    //- 答案很简单：它是不可枚举的。就像 Object.prototype 的其他属性，hasOwnProperty 有 enumerable:false 标志。
    //- 并且 for..in 只会列出可枚举的属性。这就是为什么它和其余的 Object.prototype 属性都未被列出。



// 几乎所有其他键/值获取方法都忽略继承的属性
    //- 几乎所有其他键/值获取方法，例如 Object.keys 和 Object.values 等，都会忽略继承的属性。
    //- 它们只会对对象自身进行操作。不考虑 继承自原型的属性。





// -------------------------------------------------------------------------------------------




// 总结
    //- JavaScript 中，所有的对象都有一个隐藏的 [[Prototype]] 属性，它要么是另一个对象，要么就是 null。

    //- 可以使用 obj.__proto__ 访问它（历史遗留下来的 getter/setter，这儿还有其他方法）。
    
    //- 通过 [[Prototype]] 引用的对象被称为“原型”。

    //- 如果想要读取 obj 的一个属性或者调用一个方法，并且它不存在，那么 JavaScript 就会尝试在原型中查找它。

    //- 写/删除操作直接在对象上进行，它们不使用原型（假设它是数据属性，不是 setter）。

    //- 如果调用 obj.method()，而且 method 是从原型中获取的，this 仍然会引用 obj。
    //  因此，方法始终与当前对象一起使用，即使方法是继承的。
    
    //- for..in 循环在其自身和继承的属性上进行迭代。所有其他的键/值获取方法仅对对象本身起作用。
