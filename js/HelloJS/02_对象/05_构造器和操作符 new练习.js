// 两个函数 — 一个对象

// 是否可以创建像 new A()==new B() 这样的函数 A 和 B？
let o = {};

function A() { 
    return o;
}
function B() { 
    return o;
}

let a = new A;
let b = new B;

console.log( a == b ); // true





// 创建 new Calculator
    //- 创建一个构造函数 Calculator，它创建的对象中有三个方法：

        //- read() 使有两个参数并把它们记录在对象的属性中。
        //- sum() 返回这些属性的总和。
        //- mul() 返回这些属性的乘积。

function Calculator() {

    this.read = function(num1, num2) {
        this.num1 = num1;
        this.num2 = num2;
    };

    this.sum = function() {
        return  this.num1 +  this.num2
    };

    this.mul = function() {
        return  this.num1 *  this.num2
    };
}


let calculator = new Calculator();
calculator.read(1, 2);

console.log( "Sum=" + calculator.sum() );
console.log( "Mul=" + calculator.mul() );





// 创建 new Accumulator
    //创建一个构造函数 Accumulator(startingValue)。

//它创建的对象应该：
    //- 将“当前 value”存储在属性 value 中。起始值被设置到构造器 startingValue 的参数。
    //- read() 方法应该使用一个参数来传递给一个新的数字，并将其添加到 value 中。
    //- 换句话说，value 属性是所有用户输入值与初始值 startingValue 的总和。

function Accumulator(value) {

    this.value = value;

    this.read = function(value) {

        this.value += value;

    }
}

let accumulator = new Accumulator(1); // 初始值 1

accumulator.read(1); // 添加用户输入的 value
accumulator.read(1); // 添加用户输入的 value

console.log(accumulator.value); // 显示这些值的总和


