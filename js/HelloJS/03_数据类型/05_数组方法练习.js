// 将 border-left-width 转换成 borderLeftWidth

// 编写函数 camelize(str) 将诸如 “my-short-string” 之类的由短划线分隔的单词变成骆驼式的 “myShortString”。

// 即：删除所有短横线，并将短横线后的每一个单词的首字母变为大写。

// 示例：
// console.log( camelize("background-color") == 'backgroundColor' );
// console.log( camelize("list-style-image") == 'listStyleImage' );
// console.log( camelize("-webkit-transition") == 'WebkitTransition' );
// 提示：使用 split 将字符串拆分成数组，对其进行转换之后再 join 回来。
function camelize(str) {
    return str
    .split('-') 
    .map(
      (word, index) => index == 0 ? word : word[0].toUpperCase() + word.slice(1)
        
     
    )
    .join(''); // joins ['my', 'Long', 'Word'] into 'myLongWord'
}

console.log( camelize("-webkit-transition"));

console.log("-webkit-transition".split('-') ); // [ '', 'webkit', 'transition' ]
console.log("background-color".split('-') ); // [ 'background', 'color' ]
// -------------------------------------------------------------------------------------------



// 写一个函数 filterRange(arr, a, b)，该函数获取一个数组 arr，在其中查找数值大于或等于 a，且小于或等于 b 的元素，并将结果以数组的形式返回。

// 该函数不应该修改原数组。它应该返回新的数组。

// 例如：

let arr = [5, 3, 8, 1];

let filtered = filterRange(arr, 1, 4);

console.log( filtered ); // 3,1（匹配值）

console.log( arr ); // 5,3,8,1（未修改）

function filterRange(arr, a, b) {
    return arr.filter(item => item >= a && item <= b)
}



// -------------------------------------------------------------------------------------------



// 原位（in place）过滤范围

// 写一个函数 filterRangeInPlace(arr, a, b)，该函数获取一个数组 arr，并删除其中介于 a 和 b 区间以外的所有值。检查：a ≤ arr[i] ≤ b。

// 该函数应该只修改数组。它不应该返回任何东西。

function filterRangeInPlace(arr, a, b) {
    arr.forEach( (item, idx) => {
        (item >= a && item <= b) ? "" : arr.splice(idx, 1)
    });
}
filterRangeInPlace(arr, 1, 4);
console.log( arr ); // [ 3, 1 ]



// -------------------------------------------------------------------------------------------



let arr1 = [5, 2, 1, -10, 8];

// 以降序对其进行排序
acsSort(arr1)

console.log( arr1 ); // 8, 5, 2, 1, -10

function acsSort(arr) {
    arr.sort((item1, item2) => item1 - item2); 
}

// 复制和排序数组
    // 我们有一个字符串数组 arr。我们希望有一个排序过的副本，但保持 arr 不变。

// 创建一个函数 copySorted(arr) 返回这样一个副本。

let arr2 = ["HTML", "JavaScript", "CSS"];

let sorted = copySorted(arr2);

console.log( sorted ); // CSS, HTML, JavaScript
console.log( arr2 ); // HTML, JavaScript, CSS (no changes)

function copySorted(arr) {
    let t = arr.slice(0);
    t.sort((a, b) => a.localeCompare(b) ); 
    return t;
    // return arr.slice().sort();
}



// -------------------------------------------------------------------------------------------




// 创建一个可扩展的 calculator
    // 创建一个构造函数 Calculator，以创建“可扩展”的 calculator 对象。

// 该任务由两部分组成。

// 首先，实现 calculate(str) 方法，该方法接受像 "1 + 2" 这样格式为“数字 运算符 数字”
// （以空格分隔）的字符串，并返回结果。该方法需要能够理解加号 + 和减号 -。

// 用法示例：
let calc = new Calculator;

console.log( calc.calculate("3 + 7") ); // 10


// 然后添加方法 addMethod(name, func)，该方法教 calculator 进行新操作。
// 它需要运算符 name 和实现它的双参数函数 func(a,b)。

// 例如，我们添加乘法 *，除法 / 和求幂 **：
let powerCalc = new Calculator;
powerCalc.addMethod("*", (a, b) => a * b);
powerCalc.addMethod("/", (a, b) => a / b);
powerCalc.addMethod("**", (a, b) => a ** b);

let result = powerCalc.calculate("2 ** 3");
console.log( result ); 

// 此任务中没有括号或复杂的表达式。
// 数字和运算符之间只有一个空格。
// 你可以自行选择是否添加错误处理功

function Calculator() {

    this.operators = {
        "+": (a, b) => a + b, 
        "-": (a, b) => a - b, 
    }

    this.calculate = function(str) {
        let strArr = str.split(" ");
       
        let operator = strArr[1];

        return this.operators[operator]?.(+strArr[0], +strArr[2]);
    };

    this.addMethod = function(a, b) {
        if (b && a) {
            this.operators[a] = b;
        }
    };
}





// -------------------------------------------------------------------------------------------



// 映射到 names

// 你有一个 user 对象数组，每个对象都有 user.name。编写将其转换为 names 数组的代码。

// 例如：
let john = { name: "John", age: 25 };
let pete = { name: "Pete", age: 30 };
let mary = { name: "Mary", age: 28 };

let users = [ john, pete, mary ];

let names = users.map(item => item.name);

console.log( names ); // John, Pete, Mary






// -------------------------------------------------------------------------------------------




// 映射到对象

// 你有一个 user 对象数组，每个对象都有 name，surname 和 id。

// 编写代码以该数组为基础，创建另一个具有 id 和 fullName 的对象数组，
// 其中 fullName 由 name 和 surname 生成。

// 例如：

john = { name: "John", surname: "Smith", id: 1 };
pete = { name: "Pete", surname: "Hunt", id: 2 };
mary = { name: "Mary", surname: "Key", id: 3 };

users = [ john, pete, mary ];

let usersMapped = users.map( item => ( {fullName: item.surname + " " + item.name, id: item.id} ));

/*
usersMapped = [
  { fullName: "John Smith", id: 1 },
  { fullName: "Pete Hunt", id: 2 },
  { fullName: "Mary Key", id: 3 }
]
*/

console.log( usersMapped[0].id ) // 1
console.log( usersMapped[0].fullName ) // John Smith
//所以，实际上你需要将一个对象数组映射到另一个对象数组。在这儿尝试使用箭头函数 => 来编写。




// -------------------------------------------------------------------------------------------



// 按年龄对用户排序

// 编写函数 sortByAge(users) 获得对象数组的 age 属性，并根据 age 对这些对象数组进行排序。

// 例如：
john = { name: "John", age: 25 };
pete = { name: "Pete", age: 30 };
mary = { name: "Mary", age: 28 };

arr = [ pete, john, mary ];

sortByAge(arr);

// now: [john, mary, pete]
console.log(arr[0].name); // John
console.log(arr[1].name); // Mary
console.log(arr[2].name); // Pete


function sortByAge(arr) {
    arr.sort((a, b) => a.age - b.age);
}



// -------------------------------------------------------------------------------------------



// 获取平均年龄

// 编写 getAverageAge(users) 函数，该函数获取一个具有 age 属性的对象数组，并返回平均年龄。

// 平均值的计算公式是 (age1 + age2 + ... + ageN) / N。

// 例如：
john = { name: "John", age: 25 };
pete = { name: "Pete", age: 30 };
mary = { name: "Mary", age: 29 };

arr = [ john, pete, mary ];

console.log( getAverageAge(arr) ); // (25 + 30 + 29) / 3 = 28

function getAverageAge(arr) {
    return arr.reduce( (sum, item) => (sum + item.age), 0) / arr.length;
}




// -------------------------------------------------------------------------------------------




// 数组去重

// arr 是一个数组。

// 创建一个函数 unique(arr)，返回去除重复元素后的数组 arr。

// 例如：

function unique(arr) {
    /* your code */
    let newArr = [];
    arr.forEach(item => (!newArr.includes(item)) ? newArr.push(item) : "" );
    return newArr;
}

let strings = ["Hare", "Krishna", "Hare", "Krishna",
  "Krishna", "Krishna", "Hare", "Hare", ":-O"
];

console.log( unique(strings) ); // Hare, Krishna, :-O




// -------------------------------------------------------------------------------------------





// 从数组创建键（值）对象

// 假设我们收到了一个用户数组，形式为：{id:..., name:..., age... }。

// 创建一个函数 groupById(arr) 从该数组创建对象，以 id 为键（key），数组项为值。

// 例如:

users = [
  {id: 'john', name: "John Smith", age: 20},
  {id: 'ann', name: "Ann Smith", age: 24},
  {id: 'pete', name: "Pete Peterson", age: 31},
];

let usersById = groupById(users);

console.log(usersById)

function groupById(users) {

    let usersById = {};

    users.forEach( item => usersById[item.id] = item );

    return usersById;
}
/*
// 调用函数后，我们应该得到：

usersById = {
  john: {id: 'john', name: "John Smith", age: 20},
  ann: {id: 'ann', name: "Ann Smith", age: 24},
  pete: {id: 'pete', name: "Pete Peterson", age: 31},
}

*/
// 处理服务端数据时，这个函数很有用。

// 在这个任务里我们假设 id 是唯一的。没有两个具有相同 id 的数组项。

// 请在解决方案中使用数组的 .reduce 方法。