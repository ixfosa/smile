// Rest 参数 ...
    //- 在 JavaScript 中，无论函数是如何定义的，你都可以使用任意数量的参数调用函数。

// 例如：
function sum(a, b) {
  return a + b;
}
// 虽然这里不会因为传入“过多”的参数而报错。但是当然，在结果中只有前两个参数被计算进去了。
console.log( sum(1, 2, 3, 4, 5) );


// Rest 参数可以通过使用三个点 ... 并在后面跟着包含剩余参数的数组名称，来将它们包含在函数定义中。
// 这些点的字面意思是“将剩余参数收集到一个 数组 中”。

// 例如，我们需要把所有的参数都放到数组 args 中：
function sumAll(...args) {
    let sum = 0;
    for (let arg of args) {
        sum += arg;
    }
    return sum;
}

console.log( sumAll(1) ); // 1
console.log( sumAll(1, 2) ); // 3
console.log( sumAll(1, 2, 3) ); // 6



// 下面的例子把前两个参数定义为变量，并把剩余的参数收集到 titles 数组中：
function showName(firstName, lastName, ...titles) {
    console.log( firstName + ' ' + lastName ); // Julius Caesar

    // 剩余的参数被放入 titles 数组中
    // i.e. titles = ["Consul", "Imperator"]
    console.log( titles[0] ); // Consul
    console.log( titles[1] ); // Imperator
    console.log( titles.length ); // 2
}

showName("Julius", "Caesar", "Consul", "Imperator");



// Rest 参数必须放到参数列表的末尾
    //- Rest 参数会收集剩余的所有参数，因此下面这种用法没有意义，并且会导致错误：

    //- function f(arg1, ...rest, arg2) {  // arg2 在 ...rest 后面？！
    //-     // error
    //- }

// ...rest 必须处在最后。





// -------------------------------------------------------------------------------------------





// “arguments” 变量
    //- 有一个名为 arguments 的特殊的类数组对象，该对象按参数索引包含所有参数。
function show() {
    console.log(arguments.length);

    console.log(arguments[0]);
    console.log(arguments[1]);

    // 它是可遍历的
    // for(let arg of arguments) alert(arg);
}
show("java", "python", "go");  // 3, java, python



// JavaScript 中没有 rest 参数，而使用 arguments 是获取函数所有参数的唯一方法。
// 现在它仍然有效，我们可以在一些老代码里找到它。

// 但缺点是，尽管 arguments 是一个类数组，也是可迭代对象，但它终究不是数组。
// 它不支持数组方法，因此我们不能调用 arguments.map(...) 等方法。

// 此外，它始终包含所有参数，我们不能像使用 rest 参数那样只截取入参的一部分。

// 因此，当我们需要这些功能时，最好使用 rest 参数。



// 箭头函数是没有 "arguments"
    //- 如果我们在箭头函数中访问 arguments，访问到的 arguments 并不属于箭头函数，
    //- 而是属于箭头函数外部的“普通”函数。

// 举个例子：
function f() {
  let showArg = () => console.log(arguments[0]);
  showArg();
}
f(1); // 1


function f2(i1, i2) {
    let showArg = (i1) => console.log(arguments[1]);
    showArg(i1);
}
f2(1, 2); // 2  arguments 属于箭头函数外部的“普通”函数。

// 箭头函数没有自身的 this。现在我们知道了它们也没有特殊的 arguments 对象。







// -------------------------------------------------------------------------------------------




// Spread 语法
    // 我们刚刚看到了如何从参数列表中获取数组。
    // 不过有时候我们也需要做与之相反的事儿。


// 例如，内建函数 Math.max 会返回参数中最大的值：
console.log( Math.max(3, 5, 1) ); // 5

// 假如我们有一个数组 [3, 5, 1]，我们该如何用它调用 Math.max 呢？

// 直接把数组“原样”传入是不会奏效的，因为 Math.max 希望你传入一个列表形式的数值型参数，而不是一个数组：
console.log( Math.max([3, 5, 1]) ); // NaN


// 毫无疑问，不可能手动地去一一设置参数 Math.max(arg[0], arg[1], arg[2])，
// 因为我们不确定这儿有多少个。在脚本执行时，可能参数数组中有很多个元素，也可能一个都没有。

// Spread 语法 来帮助你, 它看起来和 rest 参数很像，也使用 ...，但是二者的用途完全相反。

// 当在函数调用中使用 ...arr 时，它会把可迭代对象 arr “展开”到参数列表中。

// 以 Math.max 为例：
let arr = [3, 5, 1];
console.log( Math.max( ...arr ) )  // 5


// 还可以通过这种方式传递多个可迭代对象：
let arr1 = [1, -2, 3, 4];
let arr2 = [8, 3, -8, 1];
console.log( Math.max(...arr1, ...arr2) ); // 8

// 还可以将 spread 语法与常规值结合使用：
console.log( Math.max(1, ...arr1, 2, ...arr2, 25) ); // 25


// 还可以使用 spread 语法来合并数组：
arr = [3, 5, 1];
arr2 = [8, 9, 15];
let merged = [0, ...arr, 2, ...arr2];
console.log(merged); // 0,3,5,1,2,8,9,15（0，然后是 arr，然后是 2，然后是 arr2）



// 用数组展示了 spread 语法，其实 `任何可迭代对象` 都可以。

// 例如，在这儿我们使用 spread 语法将字符串转换为字符数组：
let str = "Hello";
console.log(...str);  // H,e,l,l,o


// Spread 语法内部使用了迭代器来收集元素，与 for..of 的方式相同。
// 因此，对于一个字符串，for..of 会逐个返回该字符串中的字符，...str 也同理会得到 "H","e","l","l","o" 
// 这样的结果。随后，字符列表被传递给数组初始化器 [...str]。


// 对于这个特定任务，我们还可以使用 Array.from 来实现，因为该方法会将一个可迭代对象（如字符串）转换为数组：
// Array.from 将可迭代对象转换为数组
console.log( Array.from(str) ); // [ 'H', 'e', 'l', 'l', 'o' ]


// 运行结果与 [...str] 相同。
// 不过 Array.from(obj) 和 [...obj] 存在一个细微的差别：
// Array.from 适用于类数组对象也适用于可迭代对象。
// Spread 语法只适用于可迭代对象。
// 因此，对于将一些“东西”转换为数组的任务，Array.from 往往更通用。





// -------------------------------------------------------------------------------------------





// 获取一个 array/object 的副本
    //- Object.assign()
    //- 使用 spread 语法也可以做同样的事情。

arr = [1, 2, 3];
let arrCopy = [...arr]; // 将数组 spread 到参数列表中
                        // 然后将结果放到一个新数组

// 两个数组中的内容相同吗？
console.log(JSON.stringify(arr) === JSON.stringify(arrCopy)); // true

// 两个数组相等吗？
console.log(arr === arrCopy); // false（它们的引用是不同的）

// 修改我们初始的数组不会修改副本：
arr.push(4);
console.log(arr); // 1, 2, 3, 4
console.log(arrCopy); // 1, 2, 3

// 并且，也可以通过相同的方式来复制一个对象：
let obj = { a: 1, b: 2, c: 3 };
let objCopy = { ...obj }; // 将对象 spread 到参数列表中
                          // 然后将结果返回到一个新对象

// 两个对象中的内容相同吗？
console.log(JSON.stringify(obj) === JSON.stringify(objCopy)); // true

// 两个对象相等吗？
console.log(obj === objCopy); // false (not same reference)

// 修改我们初始的对象不会修改副本：
obj.d = 4;
console.log(JSON.stringify(obj)); // {"a":1,"b":2,"c":3,"d":4}
console.log(JSON.stringify(objCopy)); // {"a":1,"b":2,"c":3}


// 这种方式比使用 let arrCopy = Object.assign([], arr); 来复制数组，
// 或使用 let objCopy = Object.assign({}, obj); 来复制对象写起来要短得多。








