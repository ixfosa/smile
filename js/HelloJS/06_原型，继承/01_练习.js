
// 使用原型

// 下面这段代码创建了一对对象，然后对它们进行修改。

// 过程中会显示哪些值？

let animal = {
  jumps: null
};

let rabbit = {
  __proto__: animal,
  jumps: true
};

console.log( rabbit.jumps ); // ? (1)    true

delete rabbit.jumps;

console.log( rabbit.jumps ); // ? (2)   null

delete animal.jumps;

console.log( rabbit.jumps ); // ? (3)    undefined






// -------------------------------------------------------------------------------------------




// 搜索算法

// 本题目有两个部分。

// 给定以下对象：

let head = {
  glasses: 1
};

let table = {
  pen: 3
};

let bed = {
  sheet: 1,
  pillow: 2
};

let pockets = {
  money: 2000
};

// 使用 __proto__ 来分配原型，以使得任何属性的查找都遵循以下路径：pockets → bed → table → head。
// 例如，pockets.pen 应该是 3（在 table 中找到），bed.glasses 应该是 1（在 head 中找到）。

// 回答问题：通过 pockets.glasses 或 head.glasses 获取 glasses，哪个更快？必要时需要进行基准测试。
pockets.__proto__ = bed;
bed.__proto__ = table;
table.__proto__ = head;

console.log(pockets.pen); // 3
console.log(bed.glasses); // 1

// 在现代引擎中，从性能的角度来看，我们是从对象还是从原型链获取属性都是没区别的。
// 它们（引擎）会记住在哪里找到的该属性，并在下一次请求中重用它。

// 例如，对于 pockets.glasses 来说，它们（引擎）会记得在哪里找到的 glasses（在 head 中），
// 这样下次就会直接在这个位置进行搜索。并且引擎足够聪明，一旦有内容更改，它们就会自动更新内部缓存，
// 因此，该优化是安全的。






// -------------------------------------------------------------------------------------------




// 写在哪里？

// 我们有从 animal 中继承的 rabbit。

// 如果我们调用 rabbit.eat()，哪一个对象会接收到 full 属性：animal 还是 rabbit？

let animal2 = {
  eat() {
    this.full = true;
  }
};

let rabbit2 = {
  __proto__: animal2
};

rabbit2.eat();

// rabbit 接收到 full 属性

// 这是因为 this 是点符号前面的这个对象，因此 rabbit.eat() 修改了 rabbit。
// 属性查找和执行是两回事儿。
// 首先在原型中找到 rabbit.eat 方法，然后在 this=rabbit 的情况下执行。





// -------------------------------------------------------------------------------------------




// 为什么两只仓鼠都饱了？

// 我们有两只仓鼠：speedy 和 lazy 都继承自普通的 hamster 对象。

// 当我们喂其中一只的时候，另一只也吃饱了。为什么？如何修复它？

let hamster = {
    stomach: [],

    eat(food) {
        this.stomach.push(food);
    }
};

let speedy = {
    __proto__: hamster
};

let lazy = {
    __proto__: hamster
};

// 这只仓鼠找到了食物
speedy.eat("apple");
console.log( speedy.stomach ); // apple

// 这只仓鼠也找到了食物，为什么？请修复它。
console.log( lazy.stomach ); // apple


let hamster2 = {
    stomach: [],
  
    eat(food) {
        // this.stomach= 不会执行对 stomach 的查找。该值会被直接写入 this 对象。
        //  this.stomach = [food];  // 分配给 this.stomach 而不是 this.stomach.push
        this.stomach.push(food);
    }
};

let speedy2 = {
    stomach: [],
    __proto__: hamster2
};

let lazy2 = {
    stomach: [],
    __proto__: hamster2
};


  // 这只仓鼠找到了食物
speedy2.eat("apple");
speedy2.eat("apple");
console.log( speedy2.stomach ); // apple

// 这只仓鼠也找到了食物，为什么？请修复它。
console.log( "lazy2:" + lazy2.stomach ); // undefined