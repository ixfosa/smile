// 类继承
    // 类继承是一个类扩展另一个类的一种方式。
    // 因此，我们可以在现有功能之上创建新功能。
class Animal {
    constructor(name) {
        this.name = name;
    }

    eat() {
        console.log(this.name + " eat...");
    }
}

let animal = new Animal("My animal");

//  animal 和 class Animal 的图形化表示：

/*
            prototype
    Animal -----------> Animal.prototype:
                        constructor: Animal 
                        eat: function 
                             ^
                             |  [[Prototype]]
                             |  
                        new Animal:
                        name: "My animal"
*/

// 扩展另一个类的语法是：class Child extends Parent。

// 创建另一个 class Rabbit：
class Rabbit extends Animal {
    run() {
        console.log(this.name + " run...")
    }
}

let rabbit  = new Rabbit("White Rabbit");
rabbit.eat(); // White Rabbit eat...
rabbit.run(); // White Rabbit run...

// Class Rabbit 的对象可以访问例如 rabbit.run() 等 Rabbit 的方法，
// 还可以访问例如 rabbit.eat() 等 Animal 的方法。

// 在内部，关键字 extends 使用了很好的旧的原型机制进行工作。
// 它将 Rabbit.prototype.[[Prototype]] 设置为 Animal.prototype。
// 所以，如果在 Rabbit.prototype 中找不到一个方法，JavaScript 就会从 Animal.prototype 中获取该方法。
console.log(Rabbit.prototype.__proto__ === Animal.prototype); // true

console.log(rabbit.constructor === Rabbit); // true

console.log(rabbit.__proto__ === Rabbit.prototype); // true

console.log(rabbit.__proto__.__proto__ === Animal.prototype); // true

console.log(Rabbit.__proto__ === Animal);  // true


/*
            prototype
    Animal -----------> Animal.prototype:
                        constructor: Animal 
                        eat: function 
                             ^
                             |  [[Prototype]]   <----- extends
                             |  
             prototype    Rabbit.prototype:
    Rabbit: ----------->  constructor: Rabbit
    constructor            run: function
                             ^
                             |  [[Prototype]]
                             | 
                        new Rabbit
                        name: "White Rabbit"
*/

// 例如，要查找 rabbit.eat 方法，JavaScript 引擎会进行如下检查（如图所示从下到上）：
    //- 1. 查找对象 rabbit（没有 eat）。
    //- 2. 查找它的原型，即 Rabbit.prototype（有 run，但没有 eat）。
    //- 3. 查找它的原型，即（由于 extends）Animal.prototype，在这儿找到了 eat 方法。

// JavaScript 内建对象同样也使用原型继承。
// 例如，Date.prototype.[[Prototype]] 是 Object.prototype。这就是为什么日期可以访问通用对象的方法。





// -------------------------------------------------------------------------------------------




// 在 extends 后允许任意表达式
    // 类语法不仅允许指定一个类，在 extends 后可以指定任意表达式。

// 例如，一个生成父类的函数调用：
function f(phrase) {
  return class {
    sayHi() { console.log(phrase); }
  };
}

class User extends f("Hello") {}

new User().sayHi(); // Hello

// 这里 class User 继承自 f("Hello") 的结果。

// 这对于高级编程模式，例如当我们根据许多条件使用函数生成类，并继承它们时来说可能很有用。





// -------------------------------------------------------------------------------------------





// 重写方法
    // 默认情况下，所有未在 class Rabbit 中指定的方法均从 class Animal 中直接获取。
    // 但是如果我们在 Rabbit 中指定了我们自己的方法，例如 stop()，那么将会使用它：


// 但是通常来说，我们不希望完全替换父类的方法，而是希望在父类方法的基础上进行调整或扩展其功能。
// 我们在我们的方法中做一些事儿，但是在它之前或之后或在过程中会调用父类方法。

// Class 为此提供了 "super" 关键字。
    //- 执行 super.method(...) 来调用一个父类方法。
    //- 执行 super(...) 来调用一个父类 constructor（只能在我们的 constructor 中）。

// 例如，让我们的 rabbit 在停下来的时候自动 hide：
class Animal2 {
    constructor(name) {
        this.name = name;
    }

    run() {
        console.log(this.name + " run...");
    }

    stop() {
        console.log(this.name + " stop...");
    }
}


class Rabbit2 extends Animal2{
    constructor(name, isLovely) {
        super(name);
        this.isLovely = isLovely;
    }

    hide() {
        console.log(this.name + " hide...");
    }

    stop() {
        super.stop();
        this.hide();
    }
}

let rabbit2 = new Rabbit2("Yellow Rabbit", true);

rabbit2.stop(); // Yellow Rabbit stop..., Yellow Rabbit hide...






// -------------------------------------------------------------------------------------------



/*
    // 箭头函数没有 super
        // 箭头函数没有 super。

         // 如果被访问，它会从外部函数获取。例如：

        class Rabbit extends Animal {
            stop() {
                setTimeout(() => super.stop(), 1000); // 1 秒后调用父类的 stop
            }
        }

        // 箭头函数中的 super 与 stop() 中的是一样的，所以它能按预期工作。
        // 如果我们在这里指定一个“普通”函数，那么将会抛出错误：

        // 意料之外的 super
        setTimeout(function() { super.stop() }, 1000);
*/




// -------------------------------------------------------------------------------------------




// 重写 constructor

// 根据 规范，如果一个类扩展了另一个类并且没有 constructor，那么将生成下面这样的“空” constructor：
/*
    class Rabbit extends Animal {
        // 为没有自己的 constructor 的扩展类生成的
        constructor(...args) {
            super(...args);
        }
    }

   它调用了父类的 constructor，并传递了所有的参数。
   如果我们没有写自己的 constructor，就会出现这种情况。
*/


// 给 Rabbit 添加一个自定义的 constructor。除了 name 之外，它还会指定 isLovely。
class Animal3 {
    constructor(name, color){
        this.name == name;
        this.color = color;
    }
    // ...
}

class Rabbit3 extends Animal3{
    constructor(name, isLovely) {
        // super.name = name;           // ReferenceError
        // super.color = "yellow";      // ReferenceError
        // this.name = name;            // ReferenceError
        // this.color = "yellow";       // ReferenceError

        // 需要在使用 this 之前调用 super()
        super(name, "yellow");  // true
        this.isLovely = isLovely;  
    }
    // ...
}

let rabbit3 = new Rabbit3("hah", false);


// 继承类的 constructor 必须调用 super(...)，并且 (!) 一定要在使用 this 之前调用。

// 但这是为什么呢？这里发生了什么？确实，这个要求看起来很奇怪。
    // 在 JavaScript 中，继承类（所谓的“派生构造器”，英文为 “derived constructor”）
    // 的构造函数与其他函数之间是有区别的。派生构造器具有特殊的内部属性 [[ConstructorKind]]:
    // "derived"。这是一个特殊的内部标签。

    // 该标签会影响它的 new 行为：
        //- 当通过 new 执行一个常规函数时，它将创建一个空对象，并将这个空对象赋值给 this。
        //- 但是当继承的 constructor 执行时，它不会执行此操作。它期望父类的 constructor 来完成这项工作。

    // 因此，派生的 constructor 必须调用 super 才能执行其父类（base）的 constructor，
    // 否则 this 指向的那个对象将不会被创建。并且我们会收到一个报错。






// -------------------------------------------------------------------------------------------





// 重写类字段: 一个棘手的注意要点

// 不仅可以重写方法，还可以重写类字段。

// 不过，当我们访问在父类构造器中的一个被重写的字段时，这里会有一个诡异的行为，
// 这与绝大多数其他编程语言都很不一样。

// 请思考此示例：

class Animal4 {
  name = 'animal';

  constructor() {
    console.log(this.name); // (*)
  }
}

class Rabbit4 extends Animal4 {
  name = 'rabbit';
}

new Animal4(); // animal
new Rabbit4(); // animal
console.log(new Rabbit4().name); // animal， rabbit

// 这里，Rabbit 继承自 Animal，并且用它自己的值重写了 name 字段。

// 因为 Rabbit 中没有自己的构造器，所以 Animal 的构造器被调用了。

// 在这两种情况下：new Animal4() 和 new Rabbit4()，在 (*) 行的 console.log 都打印了 animal。

// 换句话说， 父类构造器总是会使用它自己字段的值，而不是被重写的那一个。



// 古怪的是什么呢？

// 这里是相同的代码，但是我们调用 this.showName() 方法而不是 this.name 字段：
class Animal5 {
    showName() {  // 而不是 this.name = 'animal'
        console.log('animal');
    }

    constructor() {
        this.showName(); // 而不是 console。log(this.name);
    }
}

class Rabbit5 extends Animal5 {
    showName() {
        console.log('rabbit');
    }
}

new Animal5(); // animal
new Rabbit5(); // rabbit

// 请注意：这时的输出是不同的。

// 这才是我们本来所期待的结果。当父类构造器在派生的类中被调用时，它会使用被重写的方法。

// 但对于类字段并非如此。父类构造器总是使用父类的字段。

// 这里为什么会有这样的区别呢？

// 实际上，原因在于字段初始化的顺序。类字段是这样初始化的：
    //- 对于基类（还未继承任何东西的那种），在构造函数调用前初始化。
    //- 对于派生类，在 super() 后立刻初始化。


// 在我们的例子中，Rabbit 是派生类，里面没有 constructor()。
// 正如先前所说，这相当于一个里面只有 super(...args) 的空构造器。

// 所以，new Rabbit() 调用了 super()，因此它执行了父类构造器，并且（根据派生类规则）只有在此之后，它的
// 类字段才被初始化。在父类构造器被执行的时候，Rabbit 还没有自己的类字段，这就是为什么 Animal 类字段被使用了。


// 这种行为仅在一个被重写的字段被父类构造器使用时才会显现出来。

// 如果出问题了，我们可以通过使用方法或者 getter/setter 替代类字段，来修复这个问题。





// -------------------------------------------------------------------------------------------





// 深入：内部探究和 [[HomeObject]]

// 让我们更深入地研究 super。。

// 首先要说的是，从我们迄今为止学到的知识来看，super 是不可能运行的。

// 的确是这样，让我们问问自己，以技术的角度它是如何工作的？
// 当一个对象方法执行时，它会将当前对象作为 this。随后如果我们调用 super.method()，
// 那么引擎需要从当前对象的原型中获取 method。但这是怎么做到的？

// 这个任务看起来是挺容易的，但其实并不简单。引擎知道当前对象的 this，
// 所以它可以获取父 method 作为 this.__proto__.method。不幸的是，这个“天真”的解决方法是行不通的。

// 让我们演示一下这个问题。简单起见，我们使用普通对象而不使用类。


// 在下面的例子中，cat.__proto__ = animal。现在让我们尝试一下：
// 在 cat.eat() 我们将会使用 this.__proto__ 调用 animal.eat()：
animal = {
    name: "animal",
    eat: function() {
        console.log(this.name + " eat...");
    }
}

let cat = {
    __proto__: animal,

    eat: function() {
        this.__proto__.eat.call(this); // (*)
    }
}

let longEar = {
    __proto__: cat,

    eat: function() {
        this.__proto__.eat.call(this); // (**)
    }
}

// longEar.eat(); // RangeError: Maximum call stack size exceeded

// 代码无法再运行了！我们可以看到，在试图调用 longEar.eat() 时抛出了错误。

// 原因可能不那么明显，但是如果我们跟踪 longEar.eat() 调用，就可以发现原因。
// 在 (*) 和 (**) 这两行中，this 的值都是当前对象（longEar）。
// 这是至关重要的一点：所有的对象方法都将当前对象作为 this，而非原型或其他什么东西。

// 因此，在 (*) 和 (**) 这两行中，this.__proto__ 的值是完全相同的：
// 都是 cat。它们俩都调用的是 cat.eat，它们在不停地循环调用自己，而不是在原型链上向上寻找方法。

/*
    1. 在 longEar.eat() 中，(**) 这一行调用 rabbit.eat 并为其提供 this=longEar。

        // 在 longEar.eat() 中我们有 this = longEar
        this.__proto__.eat.call(this) // (**)
        // 变成了
        longEar.__proto__.eat.call(this)
        // 也就是
        rabbit.eat.call(this);

    2. 之后在 rabbit.eat 的 (*) 行中，希望将函数调用在原型链上向更高层传递，但是 this=longEar，
       所以 this.__proto__.eat 又是 rabbit.eat！

        // 在 rabbit.eat() 中我们依然有 this = longEar
        this.__proto__.eat.call(this) // (*)
        // 变成了
        longEar.__proto__.eat.call(this)
        // 或（再一次）
        rabbit.eat.call(this);

    3.所以 rabbit.eat 在不停地循环调用自己，因此它无法进一步地提升。


    这个问题没法仅仅通过使用 this 来解决

*/




// -------------------------------------------------------------------------------------------





// [[HomeObject]]
    // 为了提供解决方法，JavaScript 为函数添加了一个特殊的内部属性：[[HomeObject]]。

    // 当一个函数被定义为类或者对象方法时，它的 [[HomeObject]] 属性就成为了该对象。

    // 然后 super 使用它来解析（resolve）父原型及其方法。


// 它是怎么工作的，首先，对于普通对象：
let animal6 = {
    name: "Animal",
    eat() {     // animal.eat.[[HomeObject]] == animal
        console.log(this.name + " eat...");
    }
}

let rabbit6 = {
    __proto__: animal6,
    name: "Rabbit",
    eat() {     // rabbit.eat.[[HomeObject]] == rabbit
        super.eat();
    }
}

let longEar6 = {    
    __proto__: rabbit6,
    name: "Long Ear",
    eat() {     // longEar.eat.[[HomeObject]] == longEar
        super.eat();
    }
}

// 正确执行
longEar6.eat();  // Long Ear eat...


// 它基于 [[HomeObject]] 运行机制按照预期执行。
// 一个方法，例如 longEar.eat，知道其 [[HomeObject]] 并且从其原型中获取父方法。并没有使用 this。





// -------------------------------------------------------------------------------------------





// 方法并不是“自由”的
    // 正如我们之前所知道的，函数通常都是“自由”的，并没有绑定到 JavaScript 中的对象。
    // 正因如此，它们可以在对象之间复制，并用另外一个 this 调用它。

    // [[HomeObject]] 的存在违反了这个原则，因为方法记住了它们的对象。[[HomeObject]] 不能被更改，
    //所以这个绑定是永久的。

    // 在 JavaScript 语言中 [[HomeObject]] 仅被用于 super。所以，如果一个方法不使用 super，
    // 那么我们仍然可以视它为自由的并且可在对象之间复制。但是用了 super 再这样做可能就会出错。


// 下面是复制后错误的 super 结果的示例：
let fruit = {
    say() {
        console.log(`I'm an fruit`);
    }
}

// apple 继承自 fruit
let apple = {
    __proto__: fruit,
    say() {
       super.say();
    }
}


let plant = {
    say() {
        console.log(`I'm an plant`);
    }
}

// tree 继承自 plant
let tree = {
    __proto__: plant,
    say: apple.say   //  (*)
}

tree.say();  // I'm an fruit

// 调用 tree.say() 显示 “I'm an fruit”。这绝对是错误的。

// 原因很简单：
    //- 在 (*) 行，tree.say 方法是从 apple 复制而来。也许我们只是想避免重复代码？
    //- 它的 [[HomeObject]] 是 apple，因为它是在 apple 中创建的。没有办法修改 [[HomeObject]]。
    //- tree.say() 内具有 super.say()。它从 apple 中上溯，然后从 animal 中获取方法。

// 这是发生的情况示意图：
/*

    fruit(say)                  plant(say)
        ^                          ^
        |        [[HomeObject]]    |
    apple(say) <---------------- tree(say)

*/




// -------------------------------------------------------------------------------------------




// 方法，不是函数属性
    // [[HomeObject]] 是为类和普通对象中的方法定义的。但是对于对象而言，方法必须确切指定为 method()，
    // 而不是 "method: function()"。

// 在下面的例子中，使用非方法（non-method）语法进行了比较。未设置 [[HomeObject]] 属性，并且继承无效：

let animal7 = {
    
    eat: function() { // 这里是故意这样写的，而不是 eat() { }   
        console.log("eat ... ");
        // ...
    }
}   

let rabbit7 = {
    __proto__: animal7,

    /*
    eat() {
        super.eat();   // eat ... 
    }
    */

    eat: function() {
        super.eat();    // SyntaxError: 'super' keyword unexpected here
    }
}

// rabbit7.eat();  //  // 错误调用 super（因为这里没有 [[HomeObject]]）






// -------------------------------------------------------------------------------------------






// 总结
    // 1. 想要扩展一个类：class Child extends Parent：
        //- 这意味着 Child.prototype.__proto__ 将是 Parent.prototype，所以方法会被继承。

    // 2. 重写一个 constructor：
        //- 在使用 this 之前，我们必须在 Child 的 constructor 中将父 constructor 调用为 super()。

    // 3. 重写一个方法：
        //- 我们可以在一个 Child 方法中使用 super.method() 来调用 Parent 方法。

    // 4. 内部：
        //- 方法在内部的 [[HomeObject]] 属性中记住了它们的类/对象。这就是 super 如何解析父方法的。
        //- 因此，将一个带有 super 的方法从一个对象复制到另一个对象是不安全的。

    // 补充：
        //- 箭头函数没有自己的 this 或 super，所以它们能融入到就近的上下文中，像透明似的。
