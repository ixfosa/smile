// 添加/移除数组元素

// 数组的首端或尾端添加和删除元素的方法：
    //- arr.push(...items) —— 从尾端添加元素，
    //- arr.pop() —— 从尾端提取元素，
    //- arr.shift() —— 从首端提取元素，
    //- arr.unshift(...items) —— 从首端添加元素。




// -------------------------------------------------------------------------------------------




// splice
    // 如何从数组中删除元素？
        //- 数组是对象，所以我们可以尝试使用 delete：
let arr = ["I", "go", "home"];

delete arr[1]; // remove "go"

console.log( arr[1] ); // undefined

// now arr = ["I",  , "home"];
console.log( arr.length ); // 3

// 元素被删除了，但数组仍然有 3 个元素，我们可以看到 arr.length == 3。
// 因为 delete obj.key 是通过 key 来移除对应的值。对于对象来说是可以的。
// 但是对于数组来说，我们通常希望剩下的元素能够移动并占据被释放的位置。我们希望得到一个更短的数组。

// 所以应该使用特殊的方法。
// arr.splice 可以做所有事情：添加，删除和插入元素。
    //- arr.splice(start[, deleteCount, elem1, ..., elemN])
        //- 它从索引 start 开始修改 arr：
        //- 删除 deleteCount 个元素并在当前位置插入 elem1, ..., elemN。最后返回已被删除元素的数组。

// 删除：
let arr1 = ["I", "study", "JavaScript"];
arr1.splice(1, 1);   // 从索引 1 开始删除 1 个元素。
console.log( arr1 ); // [ 'I', 'JavaScript' ]


// 删除了 3 个元素，并用另外两个元素替换它们：

let arr2 = ["I", "study", "JavaScript", "right", "now"];

// remove 3 first elements and replace them with another
// splice 返回了已删除元素的数组：
let removed  = arr2.splice(0, 3, "Let's", "dance");
console.log( removed ) // [ 'I', 'study', 'JavaScript' ]
console.log( arr2 ) // [ "Let's", 'dance', 'right', 'now' ]


// 将 deleteCount 设置为 0，splice 方法就能够插入元素而不用删除任何元素：
let arr3 = ["I", "study", "JavaScript"];
// 从索引 2 开始
// 删除 0 个元素
// 然后插入 "complex" 和 "language"
arr3.splice(2, 0, "complex", "language")
console.log( arr3 )

// 允许负向索引
    //- 在这里和其他数组方法中，负向索引都是被允许的。它们从数组末尾计算位置，如下所示：

let arr4 = [1, 2, 5];

// 从索引 -1（尾端前一位）
// 删除 0 个元素，
// 然后插入 3 和 4
arr4.splice(-1, 0, 3, 4);
console.log( arr4 ); // 1,2,3,4,5





// ------------------------------------------------------------------------------------------



// slice

// 语法是：
    //- arr.slice([start], [end])
    //- - 它会返回一个新数组，将所有从索引 start 到 end（不包括 end）的数组项复制到一个新的数组。
    //- start 和 end 都可以是负数，在这种情况下，从末尾计算索引。
let arr5 = ["t", "e", "s", "t"];

console.log( arr5.slice(1, 3) ); // e,s（复制从位置 1 到位置 3 的元素）

console.log( arr5.slice(-2) ); // s,t（复制从位置 -2 到尾端的元素）

// 不带参数地调用它：arr.slice() 会创建一个 arr 的副本。
// 其通常用于获取副本，以进行不影响原始数组的进一步转换。
let a = arr5.slice();
console.log( arr5 ); // [ 't', 'e', 's', 't' ]
console.log( a ); // [ 't', 'e', 's', 't' ]
a[0] = 'T';
console.log( arr5 ); // [ 't', 'e', 's', 't' ]
console.log( a ); // [ 'T', 'e', 's', 't' ]





// -------------------------------------------------------------------------------------------





// concat
    //- arr.concat 创建一个新数组，其中包含来自于其他数组和其他项的值。

// 语法：
    //- arr.concat(arg1, arg2...)
        //- 它接受任意数量的参数 —— 数组或值都可以。
        //- 结果是一个包含来自于 arr，然后是 arg1，arg2 的元素的 新数组。
        //- 如果参数 argN 是一个数组，那么其中的所有元素都会被复制。否则，将复制参数本身。

let arr6 = [1, 2];

// create an array from: arr and [3,4]
console.log( arr6.concat([3, 4]) ); // 1,2,3,4

// create an array from: arr and [3,4] and [5,6]
console.log( arr6.concat([3, 4], [5, 6]) ); // 1,2,3,4,5,6

// create an array from: arr and [3,4], then add values 5 and 6
console.log( arr6.concat([3, 4], 5, 6) ); // 1,2,3,4,5,6


// 通常，它只复制数组中的元素。其他对象，即使它们看起来像数组一样，但仍然会被作为一个整体添加：
let arr7 = [1, 2];

let arrayLike = {
  0: "something",
  length: 1
};

console.log( arr7.concat(arrayLike) ); // 1,2,[object Object]

// 但是，如果类似数组的对象具有 Symbol.isConcatSpreadable 属性，那么它就会被 concat 当作一个数组来处理：此对象中的元素将被添加：

let arrayLike2 = {
  0: "something",
  1: "else",
  [Symbol.isConcatSpreadable]: true,
  length: 2
};

console.log( arr7.concat(arrayLike2) ); // 1,2,something,else





// -------------------------------------------------------------------------------------------



// indexOf/lastIndexOf 和 includes
    //- arr.indexOf、arr.lastIndexOf 和 arr.includes 方法与字符串操作具有相同的语法，
    //-  并且作用基本上也与字符串的方法相同，只不过这里是对数组元素而不是字符进行操作：

    //- arr.indexOf(item, from) 从索引 from 开始搜索 item，如果找到则返回索引，否则返回 -1。
    //- arr.lastIndexOf(item, from) —— 和上面相同，只是从右向左搜索。
    //- arr.includes(item, from) —— 从索引 from 开始搜索 item，如果找到则返回 true（
        //- 译注：如果没找到，则返回 false）
let arr8 = [1, 0, false];
console.log( arr8.indexOf(0) ); // 1
console.log( arr8.indexOf(false) ); // 2
console.log( arr8.indexOf(null) ); // -console.log
console.log( arr8.includes(1) ); // true

// 这些方法使用的是严格相等 === 比较。所以如果我们搜索 false，会精确到的确是 false 而不是数字 0。

// 如果我们想检查是否包含某个元素，并且不想知道确切的索引，那么 arr.includes 是首选。

// 此外，includes 的一个非常小的差别是它能正确处理NaN，而不像 indexOf/lastIndexOf：
const arr9 = [NaN];
console.log( arr9.indexOf(NaN) ); // -1  （应该为 0，但是严格相等 === equality 对 NaN 无效）
console.log( arr9.includes(NaN) );// true（这个结果是对的）




// -------------------------------------------------------------------------------------------




// 遍历：forEach
    //- arr.forEach 方法允许为数组的每个元素都运行一个函数。

//语法：
    //- arr.forEach(function(item, index, array) {
    //-     // ... do something with item
    //- });


// 例如，下面这个程序显示了数组的每个元素：

// ixfosa 0 [ 'ixfosa', 'long', 'zhong' ]
// long 1 [ 'ixfosa', 'long', 'zhong' ]
// zhong 2 [ 'ixfosa', 'long', 'zhong' ]
["ixfosa", "long", "zhong"].forEach(console.log);

// 而这段代码更详细地介绍了它们在目标数组中的位置：
["ixfosa", "long", "zhong"].forEach((item, index, array) => {
    console.log(`${item} is at index ${index} in ${array}`);
});

// 该函数的结果（如果它有返回）会被抛弃和忽略。




// -------------------------------------------------------------------------------------------




// find 和 findIndex
    //- 想象一下，我们有一个对象数组。我们如何找到具有特定条件的对象？
    //- 这时可以用 arr.find 方法。

// 语法如下：

    //- let result = arr.find(function(item, index, array) {
    //-     // 如果返回 true，则返回 item 并停止迭代
    //-     // 对于假值（false）的情况，则返回 undefined
    //- });

    // 依次对数组中的每个元素调用该函数：
        //- item 是元素。
        //- index 是它的索引。
        //- array 是数组本身。
        //- 如果它返回 true，则搜索停止，并返回 item。如果没有搜索到，则返回 undefined。

// 例如，我们有一个存储用户的数组，每个用户都有 id 和 name 字段。让我们找到 id == 1 的那个用户：
let users = [
  {id: 1, name: "John"},
  {id: 2, name: "Pete"},
  {id: 3, name: "Mary"}
];

let user = users.find(item => item.id == 1);

console.log(user.name); // John

// 传给了 find 一个单参数函数 item => item.id == 1. 并且 find 方法的其他参数很少使用。

// arr.findIndex 方法（与 arr.find 方法）基本上是一样的，但它返回找到元素的索引，而不是元素本身。
// 并且在未找到任何内容时返回 -1。

let idx = users.findIndex(item => item.id == 2)
console.log(idx);  // 1




// -------------------------------------------------------------------------------------------




// filter
    //- find 方法搜索的是使函数返回 true 的第一个（单个）元素。
    //- 如果需要匹配的有很多，我们可以使用 arr.filter(fn)。

// 语法与 find 大致相同，但是 filter 返回的是所有匹配元素组成的数组：

    //- let results = arr.filter(function(item, index, array) {
        //- // 如果 true item 被 push 到 results，迭代继续
        //- // 如果什么都没找到，则返回空数组
    //- });


let users2 = [
  {id: 1, name: "John"},
  {id: 2, name: "Pete"},
  {id: 3, name: "Mary"}
];

// 返回前两个用户的数组
let someUsers = users2.filter(item => item.id < 3);

console.log(someUsers.length); // 2




// -------------------------------------------------------------------------------------------




// map
    //- arr.map 方法是最有用和经常使用的方法之一。
    //- 它对数组的每个元素都调用函数，并返回结果数组。

// 语法：

    //- let result = arr.map(function(item, index, array) {
    //-     // 返回新值而不是当前元素
    //- })

//- 在这里我们将每个元素转换为它的字符串长度：

let lengths = ["Bilbo", "Gandalf", "Nazgul"].map(item => item.length);
console.log(lengths); // 5,7,6






// -------------------------------------------------------------------------------------------



// sort(fn)
    //- arr.sort 方法对数组进行 原位（in-place） 排序，更改元素的顺序。
    //- (译注：原位是指在此数组内，而非生成一个新数组。)

    //- 它还返回排序后的数组，但是返回值通常会被忽略，因为修改了 arr 本身。

let arr01 = [ 1, 2, 15 ];

// 该方法重新排列 arr 的内容
arr01.sort();

console.log( arr );  // 1, 15, 2

// 结果有什么奇怪的地方？
// 顺序变成了 1, 15, 2。不对，但为什么呢？
    //- 这些元素默认情况下被按字符串进行排序。
    // 从字面上看，所有元素都被转换为字符串，然后进行比较。对于字符串，按照词典顺序进行排序，
    // 实际上应该是 "2" > "15"。

// 要使用我们自己的排序顺序，我们需要提供一个函数作为 arr.sort() 的参数。

// 该函数应该比较两个任意值并返回：
function compare(a, b) {
  if (a > b) return 1; // 如果第一个值比第二个值大
  if (a == b) return 0; // 如果两个值相等
  if (a < b) return -1; // 如果第一个值比第二个值小
}

// 例如，按数字进行排序：
function compareNumeric(a, b) {
  if (a > b) return 1;
  if (a == b) return 0;
  if (a < b) return -1;
}

// let arr = [ 1, 2, 15 ];
arr01.sort(compareNumeric);
console.log(arr01);  // 1, 2, 15


// 思考一下这儿发生了什么。
    //- arr 可以是由任何内容组成的数组，对吗？它可能包含数字、字符串、对象或其他任何内容。
    //- 我们有一组 一些元素。要对其进行排序，我们需要一个 排序函数 来确认如何比较这些元素。
    //- 默认是按字符串进行排序的。

// arr.sort(fn) 方法实现了通用的排序算法。
// 我们不需要关心它的内部工作原理（大多数情况下都是经过 快速排序 或 Timsort 算法优化的）。
// 它将遍历数组，使用提供的函数比较其元素并对其重新排序，我们所需要的就是提供执行比较的函数 fn。

// 如果我们想知道要比较哪些元素
let b = [1, -2, 15, 2, 0, 8].sort(function(a, b) {
  console.log( a + " <> " + b );
  return a - b;
});
console.log("b: " + b) // b: -2,0,1,2,8,15
//- 2 <> 1
//- 15 <> -2
//- 15 <> 1
//- 2 <> 1
//- 2 <> 15
//- 0 <> 2
//- 0 <> 1
//- 0 <> -2
//- 8 <> 1
//- 8 <> 15
//- 8 <> 2


// 比较函数可以返回任何数字
    //- 实际上，比较函数只需要返回一个正数表示“大于”，一个负数表示“小于”。

// 通过这个原理我们可以编写更短的函数：
// let arr01 = [ 1, 2, 15 ];

arr01.sort(function(a, b) { return a - b; });

console.log(arr01);  // 1, 2, 15

// 使用箭头函数会更加简洁：
arr01.sort( (a, b) => a - b ); 


// 使用 localeCompare for strings
    //- 默认情况下，它通过字母的代码比较字母。
    // 对于许多字母，最好使用 str.localeCompare 方法正确地对字母进行排序，例如 Ö。

// 例如，让我们用德语对几个国家/地区进行排序：
let countries = ['Österreich', 'Andorra', 'Vietnam'];

console.log( countries.sort( (a, b) => a > b ? 1 : -1) ); // Andorra, Vietnam, Österreich（错的）

console.log( countries.sort( (a, b) => a.localeCompare(b) ) ); // Andorra,Österreich,Vietnam（对的！）




// -------------------------------------------------------------------------------------------



// reverse
    //- arr.reverse 方法用于颠倒 arr 中元素的顺序。

let arr02 = [1, 2, 3, 4, 5];
arr02.reverse();

console.log( arr02 ); // 5,4,3,2,1
// 它也会返回颠倒后的数组 arr。




// -------------------------------------------------------------------------------------------




// split 和 join
    //- str.split(delim) 方法可以做到。它通过给定的分隔符 delim 将字符串分割成一个数组。

// 在下面的例子中，我们用“逗号后跟着一个空格”作为分隔符：
let names = 'Bilbo, Gandalf, Nazgul, Saruman';
let arr0 = names.split(', ');

for (let name of arr0) {
    console.log( `A message to ${name}.` ); 
}

// split 方法有一个可选的第二个数字参数 —— 对数组长度的限制。
// 如果提供了，那么额外的元素会被忽略。但实际上它很少使用：
console.log(names.split(', ', 2)); // Bilbo, Gandalf


// 拆分为字母
// 调用带有空参数 s 的 split(s)，会将字符串拆分为字母数组：

let str1 = "test";

console.log( str1.split('') ); // t,e,s,t



// arr.join(glue) 与 split 相反。它会在它们之间创建一串由 glue 粘合的 arr 项。

let arr03 = ['Bilbo', 'Gandalf', 'Nazgul'];

let str2 = arr03.join(';'); // 使用分号 ; 将数组粘合成字符串

console.log( str2 ); // Bilbo;Gandalf;Nazgul





// -------------------------------------------------------------------------------------------




// reduce/reduceRight
    //- 当我们需要遍历一个数组时 —— 我们可以使用 forEach，for 或 for..of。
    //- 当我们需要遍历并返回每个元素的数据时 —— 我们可以使用 map。

    //- arr.reduce 方法和 arr.reduceRight 方法和上面的种类差不多，但稍微复杂一点。
    //- 它们用于根据数组计算单个值。

// 语法是：
    //- let value = arr.reduce(function(accumulator, item, index, array) {
    //-     // ...
    //- }, [initial]);

// 该函数一个接一个地应用于所有数组元素，并将其结果“搬运（carry on）”到下一个调用。

// 参数：

    //- accumulator —— 是上一个函数调用的结果，第一次等于 initial（如果提供了 initial 的话）。
    //- item —— 当前的数组元素。
    //- index —— 当前索引。
    //- arr —— 数组本身。

// 应用函数时，上一个函数调用的结果将作为第一个参数传递给下一个函数。
// 因此，第一个参数本质上是累加器，用于存储所有先前执行的组合结果。最后，它成为 reduce 的结果。




// 通过一行代码得到一个数组的总和：
let arr04 = [1, 2, 3, 4, 5];

// 也可以删除 reduce 的初始值（没有 0），reduce 会将数组的第一个元素作为初始值，并从第二个元素开始迭代。
// 但当 arr=[]时， 那么在没有初始值的情况下调用 reduce 会导致错误。
let result = arr04.reduce((sum, current) => sum + current, 0);

console.log(result); // 15

// 传递给 reduce 的函数仅使用了 2 个参数，通常这就足够了。

// 让我们看看细节，到底发生了什么。
    //- 在第一次运行时，sum 的值为初始值 initial（reduce 的最后一个参数），等于 0，current 是第一个数组元素，等于 1。所以函数运行的结果是 1。
    //- 在第二次运行时，sum = 1，我们将第二个数组元素（2）与其相加并返回。
    //- 在第三次运行中，sum = 3，我们继续把下一个元素与其相加，以此类推……




// -------------------------------------------------------------------------------------------




// Array.isArray
    //- 数组是基于对象的，不构成单独的语言类型。
    //- 所以 typeof 不能帮助从数组中区分出普通对象：

console.log(typeof {}); // object
console.log(typeof []); // same

// 但是数组经常被使用，因此有一种特殊的方法用于判断：Array.isArray(value)。
// 如果 value 是一个数组，则返回 true；否则返回 false。

console.log(Array.isArray({})); // false

console.log(Array.isArray([])); // true   





// -------------------------------------------------------------------------------------------




// 大多数方法都支持 “thisArg”
    //- 几乎所有调用函数的数组方法 —— 比如 find，filter，map，除了 sort 是一个特例，
    // 都接受一个可选的附加参数 thisArg。

//- 上面的部分中没有解释该参数，因为该参数很少使用。但是为了完整性，我们需要讲讲它。

// 以下是这些方法的完整语法
    //- arr.find(func, thisArg);
    //- arr.filter(func, thisArg);
    //- arr.map(func, thisArg);
    // ...


// thisArg 是可选的最后一个参数
    //- thisArg 参数的值在 func 中变为 this。

// 例如，在这里我们使用 army 对象方法作为过滤器，thisArg 用于传递上下文（passes the context）：

let army = {
  minAge: 18,
  maxAge: 27,
  canJoin(user) {
    return user.age >= this.minAge && user.age < this.maxAge;
  }
};

let user = [
  {age: 16},
  {age: 20},
  {age: 23},
  {age: 30}
];

// 找到 army.canJoin 返回 true 的 user
let soldiers = user.filter(army.canJoin, army);

console.log(soldiers.length); // 2
console.log(soldiers[0].age); // 20
console.log(soldiers[1].age); // 23

// 使用了 user.filter(army.canJoin)，那么 army.canJoin 将被作为独立函数调用，
// 并且这时 this=undefined，从而会导致即时错误。

// 可以用 user.filter(user => army.canJoin(user)) 替换对 users.filter(army.canJoin, army) 的调用。
// 前者的使用频率更高，因为对于大多数人来说，它更容易理解。





// -------------------------------------------------------------------------------------------





// 总结

// 添加/删除元素：
    //- push(...items) —— 向尾端添加元素，
    //- pop() —— 从尾端提取一个元素，
    //- shift() —— 从首端提取一个元素，
    //- unshift(...items) —— 向首端添加元素，
    //- splice(pos, deleteCount, ...items) —— 从 pos 开始删除 deleteCount 个元素，并插入 items。
    //- slice(start, end) —— 创建一个新数组，将从索引 start 到索引 end（但不包括 end）的元素复制进去。
    //- concat(...items) —— 返回一个新数组：复制当前数组的所有元素，并向其中添加 items。如果 items 中的任意一项是一个数组，那么就取其元素。

// 搜索元素：
    //- indexOf/lastIndexOf(item, pos) —— 从索引 pos 开始搜索 item，搜索到则返回该项的索引，否则返回 -1。
    //- includes(value) —— 如果数组有 value，则返回 true，否则返回 false。
    //- find/filter(func) —— 通过 func 过滤元素，返回使 func 返回 true 的第一个值/所有值。
    //- findIndex 和 find 类似，但返回索引而不是值。

// 遍历元素：
    //- forEach(func) —— 对每个元素都调用 func，不返回任何内容。

// 转换数组：
    //- map(func) —— 根据对每个元素调用 func 的结果创建一个新数组。
    //- sort(func) —— 对数组进行原位（in-place）排序，然后返回它。
    //- reverse() —— 原位（in-place）反转数组，然后返回它。
    //- split/join —— 将字符串转换为数组并返回。
    //- reduce/reduceRight(func, initial) —— 通过对每个元素调用 func 计算数组上的单个值，并在调用之间传递中间结果。

// 其他：
    //- Array.isArray(arr) 检查 arr 是否是一个数组。

// 请注意，sort，reverse 和 splice 方法修改的是数组本身。

// 这些是最常用的方法，它们覆盖 99％ 的用例。但是还有其他几个：

// arr.some(fn)/arr.every(fn) 检查数组。
// 与 map 类似，对数组的每个元素调用函数 fn。如果任何/所有结果为 true，则返回 true，否则返回 false。
// 这两个方法的行为类似于 || 和 && 运算符：
    // 如果 fn 返回一个真值，arr.some() 立即返回 true 并停止迭代其余数组项；
    // 如果 fn 返回一个假值，arr.every() 立即返回 false 并停止对其余数组项的迭代。

// 我们可以使用 every 来比较数组：

function arraysEqual(arr1, arr2) {
  return arr1.length === arr2.length && arr1.every((value, index) => value === arr2[index]);
}

console.log( arraysEqual([1, 2], [1, 2])); // true

// arr.fill(value, start, end) —— 从索引 start 到 end，用重复的 value 填充数组。

// arr.copyWithin(target, start, end) —— 将从位置 start 到 end 的所有元素复制到 自身 的 target 位置（覆盖现有元素）。

// arr.flat(depth)/arr.flatMap(fn) 从多维数组创建一个新的扁平数组。

// arr.of(element0[, element1[, …[, elementN]]]) 基于可变数量的参数创建一个新的 Array 实例，而不需要考虑参数的数量或类型。



