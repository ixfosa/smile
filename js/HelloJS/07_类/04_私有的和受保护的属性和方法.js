// 私有的和受保护的属性和方法
    // 面向对象编程最重要的原则之一 —— 将内部接口与外部接口分隔开来。


// 内部接口和外部接口
    // 在面向对象的编程中，属性和方法分为两组：
        //- 内部接口 —— 可以通过该类的其他方法访问，但不能从外部访问的方法和属性。
        //- 外部接口 —— 也可以从类的外部访问的方法和属性。


    // 在 JavaScript 中，有两种类型的对象字段（属性和方法）：
        //- 公共的：可从任何地方访问。它们构成了外部接口。到目前为止，我们只使用了公共的属性和方法。
        //- 私有的：只能从类的内部访问。这些用于内部接口。


// 受保护的字段不是在语言级别的 Javascript 中实现的，但实际上它们非常方便，
// 因为它们是在 Javascript 中模拟的类定义语法。




    
// -------------------------------------------------------------------------------------------




// 受保护的 

// 首先，让我们做一个简单的咖啡机类：
class CoffeeMachine {
    waterAmount = 0; // 内部的水量
    constructor(power) {
        this.power = power;
        console.log(`Created a coffee-machine, power: ${power}`);
    }
}

// 创建咖啡机
let coffeeMachine = new CoffeeMachine(100); // Created a coffee-machine, power: 100
console.log(coffeeMachine.power); // 100

// 现在，属性 waterAmount 和 power 是公共的。我们可以轻松地从外部将它们 get/set 成任何值。

// 让我们将 waterAmount 属性更改为受保护的属性，以对其进行更多控制。
// 例如，我们不希望任何人将它的值设置为小于零的数。

// 受保护的属性通常以下划线 _ 作为前缀。

// 这不是在语言级别强制实施的，但是程序员之间有一个众所周知的约定，即不应该从外部访问此类型的属性和方法。

// 所以我们的属性将被命名为 _waterAmount：
class CoffeeMachine1 {
    _waterAmount = 0; // 内部的水量
    constructor(power) {
        this._power = power;
    }

    set waterAmount(value) {
        if (value < 0) {
            throw new Error("Negative water");
        }
        this._waterAmount = value;
    }
}

// 创建咖啡机
let coffeeMachine1 = new CoffeeMachine1(100);

// 加水
// coffeeMachine1.waterAmount = -10; // Error: Negative water
// 现在访问已受到控制，因此将水量的值设置为小于零的数将会失败。




// -------------------------------------------------------------------------------------------



// 只读的
    // 对于 power 属性，让我们将它设为只读。有时候一个属性必须只能被在创建时进行设置，之后不再被修改。

// 咖啡机就是这种情况：功率永远不会改变。

// 要做到这一点，我们只需要设置 getter，而不设置 setter：
class CoffeeMachine2 {

    // ...

    constructor(power) {
        this._power = power;
    }

    get power() {
        return this._power;
    }
}

// 创建咖啡机
let coffeeMachine2 = new CoffeeMachine2(100);

console.log(`Power is: ${coffeeMachine2.power}W`); // 功率是：100W

coffeeMachine.power = 25; // Error（没有 setter）

console.log(`Power is: ${coffeeMachine2.power}W`); // Power is: 100W



/*

Getter/setter 函数
    这里我们使用了 getter/setter 语法。

但大多数时候首选 get.../set... 函数，像这样：

class CoffeeMachine {
    _waterAmount = 0;

    setWaterAmount(value) {
        if (value < 0) throw new Error("Negative water");
        this._waterAmount = value;
    }

    getWaterAmount() {
        return this._waterAmount;
    }
}

new CoffeeMachine().setWaterAmount(100);

这看起来有点长，但函数更灵活。它们可以接受多个参数（即使我们现在还不需要）。

另一方面，get/set 语法更短，所以最终没有严格的规定，而是由你自己来决定。

*/


// 受保护的字段是可以被继承的
    // 如果我们继承 class MegaMachine extends CoffeeMachine，
    // 那么什么都无法阻止我们从新的类中的方法访问 this._waterAmount 或 this._power。

    // 所以受保护的字段是自然可被继承的。与我们接下来将看到的私有字段不同





// -------------------------------------------------------------------------------------------




// 私有的

// 这儿有一个马上就会被加到规范中的已完成的 Javascript 提案，它为私有属性和方法提供语言级支持。

// 私有属性和方法应该以 # 开头。它们只在类的内部可被访问。

// 例如，这儿有一个私有属性 #waterLimit 和检查水量的私有方法 #checkWater：
class CoffeeMachine3 {

    #waterLimit = 200;

    #checkWater(value) {
        if (value < 0) throw new Error("Negative water");
        if (value > this.#waterLimit)  throw new Error("Too much water");
    }
}

let coffeeMachine3 = new CoffeeMachine3();

// 不能从类的外部访问类的私有属性和方法
// SyntaxError: Private field '#checkWater' must be declared in an enclosing class
//- coffeeMachine3.#checkWater(100);
//- console.log(coffeeMachine3.#waterLimit);


// 在语言级别，# 是该字段为私有的特殊标志。我们无法从外部或从继承的类中访问它。

// 私有字段与公共字段不会发生冲突。我们可以同时拥有私有的 #waterAmount 和公共的 waterAmount 字段。

// 例如，让我们使 waterAmount 成为 #waterAmount 的一个访问器：
class CoffeeMachine5 {

    #waterAmount  = 0;

    get waterAmount() {
        return this.#waterAmount;
    }

    set waterAmount(value) {
        if (value < 0) throw new Error("Negative water");
        this.#waterAmount = value;
    }
}

let coffeeMachine4 = new CoffeeMachine5();

coffeeMachine4.waterAmount = 100;
// console.log(coffeeMachine4.#waterAmount) // Error
console.log(coffeeMachine4.waterAmount); // 100


// 与受保护的字段不同，私有字段由语言本身强制执行。这是好事儿。

// 但是如果我们继承自 CoffeeMachine，那么我们将无法直接访问 #waterAmount。
//我们需要依靠 waterAmount getter/setter：
class MegaCoffeeMachine extends CoffeeMachine5 {

    method() {
        // console.log( this.#waterAmount ); // Error: can only access from CoffeeMachine
    }
}


// 在许多情况下，这种限制太严重了。如果我们扩展 CoffeeMachine，则可能有正当理由访问其内部。
// 这就是为什么大多数时候都会使用受保护字段，即使它们不受语言语法的支持。






// 私有字段不能通过 this[name] 访问

// 通常我们可以使用 this[name] 访问字段：
/*
class User {
    // ...
    sayHi() {
        let fieldName = "name";
        console.log(`Hello, ${this[fieldName]}`);
    }
}
*/

// 对于私有字段来说，这是不可能的：this['#name'] 不起作用。这是确保私有性的语法限制。



