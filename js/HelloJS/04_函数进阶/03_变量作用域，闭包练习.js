
// 函数会选择最新的内容吗？

// 函数 sayHi 使用外部变量。当函数运行时，将使用哪个值？
/*
let name = "zhong";

function sayHi() {
    console.log("hello " + name);
}
name = "zhong";

sayHi(); // hello zhong
*/

// 函数将从内到外依次在对应的词法环境中寻找目标变量，它使用最新的值。

// 旧变量值不会保存在任何地方。当一个函数想要一个变量时，它会从自己的词法环境或外部词法环境中获取当前值。






// -------------------------------------------------------------------------------------------




// 哪些变量可用呢？

// 下面的 makeWorker 函数创建了另一个函数并返回该函数。可以在其他地方调用这个新函数。

// 它是否可以从它被创建的位置或调用位置（或两者）访问外部变量？

function makeWorker() {
    let name = "Pete";
  
    return function() {
        console.log(name);
    };
}
  
let name = "John";

// create a function
let work = makeWorker();

// call it
work(); // 会显示什么？ Pete.



// 但如果在 makeWorker() 中没有 let name，那么将继续向外搜索并最终找到全局变量，
// 在这种情况下，结果将是 "John"。






// -------------------------------------------------------------------------------------------





// Counter 是独立的吗？

// 在这儿我们用相同的 makeCounter 函数创建了两个计数器（counters）：counter 和 counter2。

// 它们是独立的吗？第二个 counter 会显示什么？0,1 或 2,3 还是其他？

function makeCounter() {
    let count = 0;
  
    return function() {
        return count++;
    };
}
  
let counter = makeCounter();
let counter2 = makeCounter();

console.log( counter() ); // 0
console.log( counter() ); // 1

console.log( counter2() ); // ?  0
console.log( counter2() ); // ?  1


// 函数 counter 和 counter2 是通过 makeCounter 的不同调用创建的。
// 因此，它们具有独立的外部词法环境，每一个都有自己的 count。




// -------------------------------------------------------------------------------------------





// Counter 对象

// 这里通过构造函数创建了一个 counter 对象。

// 它能正常工作吗？它会显示什么呢？

function Counter() {
    let count = 0;
  
    this.up = function() {
        return ++count;
    };
    this.down = function() {
        return --count;
    };
}
  
counter = new Counter();

// 这两个嵌套函数都是在同一个词法环境中创建的，所以它们可以共享对同一个 count 变量的访问： 
console.log( counter.up() ); // ?     1
console.log( counter.up() ); // ?     2
console.log( counter.down() ); // ?   1




// -------------------------------------------------------------------------------------------



// if 内的函数
// 看看下面这个代码。最后一行代码的执行结果是什么？

let phrase = "Hello";

if (true) {
  let user = "John";

  function sayHi() {
    console.log(`${phrase}, ${user}`);
  }
}

sayHi(); // Hello, John





// -------------------------------------------------------------------------------------------





// 闭包 sum

// 编写一个像 sum(a)(b) = a+b 这样工作的 sum 函数。

// 是的，就是这种通过双括号的方式（并不是错误）。

// 举个例子：

// sum(1)(2) = 3
// sum(5)(-1) = 4


function sum(a) {
    return function(b) {
        return a + b;
    }
}

console.log( sum(1)(2) );   
console.log( sum(5)(-1) );






// -------------------------------------------------------------------------------------------





// 变量可见吗？

// 下面这段代码的结果会是什么？

let x = 1;

function func() {
  console.log(x); // ReferenceError: Cannot access 'x' before initialization

  let x = 2;
}
// func();  // 答案：error。


// 在这个例子中，我们可以观察到“不存在”的变量和“未初始化”的变量之间的特殊差异。

// 从程序执行进入代码块（或函数）的那一刻起，变量就开始进入“未初始化”状态。
// 它一直保持未初始化状态，直至程序执行到相应的 let 语句。

// 换句话说，一个变量从技术的角度来讲是存在的，但是在 let 之前还不能使用。

// 下面的这段代码证实了这一点。
function func2() {

  // 引擎从函数开始就知道局部变量 x，
  // 但是变量 x 一直处于“未初始化”（无法使用）的状态，直到结束 let（“死区”）
  // 因此答案是 error

  // console.log(x); // ReferenceError: Cannot access 'x' before initialization

  let x = 2;
}
// 变量暂时无法使用的区域（从代码块的开始到 let）有时被称为“死区”。




// -------------------------------------------------------------------------------------------




// 通过函数筛选

// 我们有一个内建的数组方法 arr.filter(f)。它通过函数 f 过滤元素。
// 如果它返回 true，那么该元素会被返回到结果数组中。

// 制造一系列“即用型”过滤器：
    //- inBetween(a, b) —— 在 a 和 b 之间或与它们相等（包括）。
    //- inArray([...]) —— 包含在给定的数组中。

// 用法如下所示：
    //- arr.filter(inBetween(3,6)) —— 只挑选范围在 3 到 6 的值。
    //- arr.filter(inArray([1,2,3])) —— 只挑选与 [1,2,3] 中的元素匹配的元素。

// 例如：
/* .. inBetween 和 inArray 的代码 */
let arr = [1, 2, 3, 4, 5, 6, 7];

console.log( arr.filter(inBetween(3, 6)) ); // 3,4,5,6

console.log( arr.filter(inArray([1, 2, 10])) ); // 1,2


function inBetween(a, b) {
    return item => item >= a && item <= b;
}

function inArray(arr) {
    return item => arr.includes( item )
}



// -------------------------------------------------------------------------------------------




// 按字段排序

// 我们有一组要排序的对象：

let users = [
  { name: "John", age: 20, surname: "Johnson" },
  { name: "Pete", age: 18, surname: "Peterson" },
  { name: "Ann", age: 19, surname: "Hathaway" }
];

// 通常的做法应该是这样的：

// 通过 name (Ann, John, Pete)
// users.sort((a, b) => a.name > b.name ? 1 : -1);

// 通过 age (Pete, Ann, John)
// users.sort((a, b) => a.age > b.age ? 1 : -1);


// 我们可以让它更加简洁吗，比如这样？
users.sort(byField('name'));
console.log(users);
users.sort(byField('age'));
console.log(users);

// 这样我们就只需要写 byField(fieldName)，而不是写一个函数。
// 编写函数 byField 来实现这个需求

function byField(fieldName) {
    return (a, b) => a[fieldName] > b[fieldName] ? 1 : -1
}





// -------------------------------------------------------------------------------------------




// 函数大军

// 下列的代码创建了一个 shooters 数组。

// 每个函数都应该输出其编号。但好像出了点问题……

function makeArmy() {
  let shooters = [];

  let i = 0;
  while (i < 10) {
    let shooter = function() { // 创建一个 shooter 函数，
      console.log( i ); // 应该显示其编号
    };
    shooters.push(shooter); // 将此 shooter 函数添加到数组中
    i++;
  }

  // 返回 shooters 数组
  return shooters;
}

let army = makeArmy();

// 所有的 shooter 显示的都是 10，而不是它们的编号 0, 1, 2, 3...
army[0](); // 编号为 0 的 shooter 显示的是 10
army[1](); // 编号为 1 的 shooter 显示的是 10
army[2](); // 10，其他的也是这样。

// 为什么所有的 shooter 显示的都是同样的值？

// 修改代码以使得代码能够按照我们预期的那样工作。


function makeArmy2() {
    let shooters = [];
  
    let i = 0;

    while (i < 10) {
        let shooter = function(i) { // 创建一个 shooter 函数，
            return function() {
                console.log( i ); // 应该显示其编号
            }
        };
        shooters.push(shooter(i)); // 将此 shooter 函数添加到数组中
        i++;
    }
  
    // 返回 shooters 数组
    return shooters;
}
  
let army2 = makeArmy2();
army2[0](); 
army2[1](); 
army2[2]();



function makeArmy3() {
    let shooters = [];
  
    let i = 0;

    while (i < 10) {
        
        // 在 while {...} 块的每次迭代中，都会创建一个新的词法环境。
        // 因此，要解决此问题，我们可以将 i 的值复制到 while {...} 块内的变量中，如下所示：
        let j = i;

        let shooter = function() { // 创建一个 shooter 函数，
            console.log( j ); // 应该显示其编号
        };
        shooters.push(shooter); // 将此 shooter 函数添加到数组中
        i++;
    }
  
    // 返回 shooters 数组
    return shooters;
}
  
let army3 = makeArmy3();
army3[0](); 
army3[1](); 
army3[2]();