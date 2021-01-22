// 在对象字面量中使用 "this"

// 这里 makeUser 函数返回了一个对象。

// 访问 ref 的结果是什么？为什么？
'use strict'
function makeUser() {
  return {
    name: "John",
    ref: this
  };
}

let user = makeUser();

// console.log( user.ref.name ); // 结果是什么？ // TypeError: Cannot read property 'name' of undefined


// 这是因为设置 this 的规则不考虑对象定义。只有调用那一刻才重要。
// 这里 makeUser() 中的 this 的值是 undefined，因为它是被作为函数调用的，而不是通过点符号被作为方法调用。
// this 的值是对于整个函数的，代码段和对象字面量对它都没有影响。
// 所以 ref: this 实际上取的是当前函数的 this。
// 我们可以重写这个函数，并返回和上面相同的值为 undefined 的 this：

function makeUser1(){
  return this; // 这次这里没有对象字面量
}
console.log( makeUser1()); // undefined
// console.log( makeUser1().name ); // Error: Cannot read property 'name' of undefined



// 反例
function makeUser2() {
    return {
      name: "John",
      ref() {
        return this;
      }
    };
}
  
  let user2 = makeUser2();
  
  // 因为 user.ref() 是一个方法。this 的值为点符号 . 前的这个对象。
  console.log( user2.ref().name ); // John






// 创建一个计算器

// 创建一个有三个方法的 calculator 对象：
    //- read() 提示输入两个值，并将其保存为对象属性。
    //- sum() 返回保存的值的和。
    //- mul() 将保存的值相乘并返回计算结果。

let calculator = {

    read(num1, num2) {
        this.num1 = num1;
        this.num2 = num2;
    },

    sum() {
        return this.num1 + this.num2;
    },

    mul() {
        return this.num1 * this.num2;
    },

};

calculator.read(1, 2);
console.log( calculator.sum() );
console.log( calculator.mul() );







// 链式（调用）

//有一个可以上下移动的 ladder 对象：
let ladder = {
    step: 0,
    up() {
        this.step++;
    },
    down() {
        this.step--;
    },
    showStep: function() { // 显示当前的 step
        console.log( this.step );
    }
};

// 现在，如果我们要按顺序执行几次调用，可以这样做：
ladder.up();
ladder.up();
ladder.down();
ladder.showStep(); // 1



let ladder2 = {
    step: 0,
    up() {
        this.step++;
        return this
    },
    down() {
        this.step--;
        return this
    },
    showStep: function() { // 显示当前的 step
        console.log( this.step );
    }
};

// 修改 up，down 和 showStep 的代码，让调用可以链接，就像这样：
ladder2.up().up().down().showStep(); // 1

ladder2
  .up()
  .up()
  .down()
  .up()
  .down()
  .showStep(); // 2 = 1+ 1