// 解构赋值
    //- JavaScript 中最常用的两种数据结构是 Object 和 Array。

    // 对象让我们能够创建通过键来存储数据项的单个实体，数组则让我们能够将数据收集到一个有序的集合中。

    //  但是，当我们把它们传递给函数时，它可能不需要一个整体的对象/数组，而是需要单个块。

// 解构赋值 是一种特殊的语法，它使我们可以将数组或对象“拆包”为到一系列变量中，因为有时候使用变量更加方便。
// 解构操作对那些具有很多参数和默认值等的函数也很奏效。




// -------------------------------------------------------------------------------------------



"use strict";
// 数组解构

//下面是一个将数组解构到变量中的例子：
let user = ["fo", "long", "zhong"];

// let name1 = user[0];
// let name2 = user[1];

// 解构
let [name1, name2] = user;

console.log(`name1: ${name1} - name2: ${name2}`) // name1: fo - name2: long


// “解构”并不意味着“破坏”
    //- 这种语法叫做“解构赋值”，因为它通过将结构中的各元素复制到变量中来达到“解构”的目的。
    //- 但数组本身是没有被修改的。

// 这只是下面这些代码的更精简的写法而已：
// let [firstName, surname] = arr;
// let firstName = arr[0];
// let surname = arr[1];



// 忽略使用逗号的元素
    //- 数组中不想要的元素也可以通过添加额外的逗号来把它丢弃：
// 不需要第二个元素
let [name, , title] = ["Julius", "Caesar", "Consul", "of the Roman Republic"];
console.log( title ); // Consul



// 等号右侧可以是任何可迭代对象
    //- 实际上，我们可以将其与任何可迭代对象一起使用，而不仅限于数组：
let [a, b, c] = "abc"; // ["a", "b", "c"]
let [one, two, three] = new Set([1, 2, 3]);




// 赋值给等号左侧的任何内容
    //- 我们可以在等号左侧使用任何“可以被赋值的”东西。

// 例如，一个对象的属性：
let users = {};
[users.name, users.surname] = "Ilya Kantor".split(' ');
console.log(users.name); // Ilya




// 与 .entries() 方法进行循环操作

// 我们可以将 .entries() 方法与解构语法一同使用，来遍历一个对象的“键—值”对：

let user1 = {
  name: "John",
  age: 30
};

// 循环遍历键—值对
for (let [key, value] of Object.entries(user1)) {
  console.log(`${key}:${value}`); // name:John, then age:30
}

//对于 map 对象也类似：
let user2 = new Map();
user2.set("name", "John");
user2.set("age", "30");

for (let [key, value] of user2) {
    console.log(`${key}:${value}`); // name:John, then age:30
}


// 交换变量值的技巧
// 一个用于交换变量值的典型技巧：
let guest = "Jane";
let admin = "Pete";
// 交换值：让 guest=Pete, admin=Jane
[guest, admin] = [admin, guest];

console.log(`${guest} ${admin}`); // Pete Jane（成功交换！）





// -------------------------------------------------------------------------------------------




// 剩余的 ‘…’
    // 如果我们不只是要获得第一个值，还要将后续的所有元素都收集起来 — 
    // 我们可以使用三个点 "..." 来再加一个参数来接收“剩余的”元素：

let [name11, name12, ...rest] = ["Julius", "Caesar", "Consul", "of the Roman Republic"];

console.log(name11); // Julius
console.log(name12); // Caesar

// 请注意，`rest` 的类型是数组
console.log(rest[0]); // Consul
console.log(rest[1]); // of the Roman Republic
console.log(rest.length); // 2

// rest 的值就是数组中剩下的元素组成的数组。
// 不一定要使用变量名 rest，我们也可以使用其他的变量名，只要确保它前面有三个点，
// 并且在解构赋值的最后一个参数位置上就行了。






// -------------------------------------------------------------------------------------------




// 默认值
    //- 如果赋值语句中，变量的数量多于数组中实际元素的数量，赋值不会报错。未赋值的变量被认为是 undefined：


let [firstName, surname] = [];
console.log(firstName); // undefined
console.log(surname); // undefined

// 如果我们想要一个“默认”值给未赋值的变量，我们可以使用 = 来提供：
// 默认值
let [name0 = "Guest", surname0 = "Anonymous"] = ["Julius"];

console.log(name0);    // Julius（来自数组的值）
console.log(surname0); // Anonymous（默认值被使用了）

// 默认值可以是更加复杂的表达式甚至可以是函数调用，这些表达式或函数只会在这个变量未被赋值的时候才会被计算。






// -------------------------------------------------------------------------------------------




// 对象解构
    //- 解构赋值同样适用于对象。

// 基本语法是：
    //- let {var1, var2} = {var1:…, var2:…}
    // 在等号右侧有一个已经存在的对象，我们想把它拆开到变量中。
    // 等号左侧包含了对象相应属性的一个“模式（pattern）”。在简单的情况下，等号左侧的就是 {...} 中的变量名列表。

let options0 = {
  title0: "Menu",
  width0: 100,
  height0: 200
};

let {title0, width0, height0} = options0;

console.log(title0);  // Menu
console.log(width0);  // 100
console.log(height0); // 200

// 属性 options.title、options.width 和 options.height 值被赋给了对应的变量。

// 变量的顺序并不重要，下面这个代码也奏效：
// 改变 let {...} 中元素的顺序
    //- let {height, width, title} = { title: "Menu", height: 200, width: 100 }

// 等号左侧的模式（pattern）可以更加复杂，并且指定了属性和变量之间的映射关系。

// 如果我们想把一个属性赋值给另一个名字的变量，比如把 options.width 属性赋值给变量 w，
// 那么我们可以使用冒号来指定：

let options1 = {
  width1: 100,
  height1: 200
};

// { sourceProperty: targetVariable }
let {width1: w, height1} = options1;

// width -> w
// height1 -> height1

console.log(w);      // 100
console.log(height1);     // 200

// 冒号表示“什么值：赋值给谁”。上面的例子中，属性 width 被赋值给了 w，属性 height 被赋值给了 h，
// 属性 title1 被赋值给了同名变量。


// 对于可能缺失的属性，我们可以使用 "=" 设置默认值，如下所示：
let options2 = {
  title2: "Menu"
};

let {width2 = 100, height2 = 200, title2} = options2;

console.log(title2);  // Menu
console.log(width2);  // 100
console.log(height2); // 200




// 就像数组或函数参数一样，默认值可以是任意表达式甚至可以是函数调用。
// 它们只会在未提供对应的值时才会被计算/调用。

// 在下面的代码中，prompt 提示输入 width 值，但不会提示输入 title 值：
/*
let options = {
  title: "Menu"
};

let {width = prompt("width?"), title = prompt("title?")} = options;

console.log(title);  // Menu
console.log(width);  //（无论 prompt 的结果是什么）
*/



// 我们还可以将冒号和等号结合起来：
let options3 = {
  title3: "Menu"
};
let {width3: w3 = 100, height3: h3 = 200, title3} = options3;

console.log(title3);  // Menu
console.log(w3);      // 100
console.log(h3);      // 200



// 如果我们有一个具有很多属性的复杂对象，那么我们可以只提取所需的内容：
let options4 = {
  title4: "Menu",
  width4: 100,
  height4: 200
};

// 仅提取 title 作为变量
let { title4 } = options4;

console.log(title4); // Menu




// -------------------------------------------------------------------------------------------




// 剩余模式（pattern）"…"
    //- 如果对象拥有的属性数量比我们提供的变量数量还多，该怎么办？
    //- 我们可以只取其中的某一些属性，然后把“剩余的”赋值到其他地方吗？

// 我们可以使用剩余模式（pattern），就像我们对数组那样。
// 一些较旧的浏览器不支持此功能（例如，使用 Babel 对其进行填充），但可以在现代浏览器中使用。

// 看起来就像这样：
let options01 = {
  title01: "Menu",
  height01: 200,
  width01: 100
};

// title = 名为 title 的属性
// rest = 存有剩余属性的对象
let {title01, ...rest01} = options01;

// 现在 title="Menu", rest={height: 200, width: 100}
console.log(rest01.height01);  // 200
console.log(rest01.width01);   // 100


// 不使用 let 时的陷阱
// 在上面的示例中，变量都是在赋值中通过正确方式声明的：let {…} = {…}。
// 当然，我们也可以使用已有的变量，而不用 let，但这里有一个陷阱。

/*
// 以下代码无法正常运行：
let title, width, height;

// 这一行发生了错误
{title, width, height} = {title: "Menu", width: 200, height: 100};
// 问题在于 JavaScript 把主代码流（即不在其他表达式中）的 {...} 当做一个代码块。
// 这样的代码块可以用于对语句分组，如下所示：
*/

{
  // 一个代码块
  let message = "Hello";
  // ...
  console.log( message );
}
// 因此，这里 JavaScript 假定我们有一个代码块，这就是报错的原因。我们需要解构它。


// 为了告诉 JavaScript 这不是一个代码块，我们可以把整个赋值表达式用括号 (...) 包起来：
let title02, width02, height02;

// 现在就可以了
({title02, width02, height02} = {title02: "Menu", width02: 200, height02: 100});
console.log( title02 ); // Menu



// -------------------------------------------------------------------------------------------




// 嵌套解构
    //- 如果一个对象或数组嵌套了其他的对象和数组，我们可以在等号左侧使用更复杂的模式（pattern）来提取更深层的数据。

// 在下面的代码中，options 的属性 size 是另一个对象，属性 items 是另一个数组。
// 赋值语句中等号左侧的模式（pattern）具有相同的结构以从中提取值：

/*
let options = {
  size: {
    width: 100,
    height: 200
  },
  items: ["Cake", "Donut"],
  extra: true
};

// 为了清晰起见，解构赋值语句被写成多行的形式
let {
  size: { // 把 size 赋值到这里
    width,
    height
  },
  items: [item1, item2], // 把 items 赋值到这里
  title = "Menu" // 在对象中不存在（使用默认值）
} = options;

console.log(title);  // Menu
console.log(width);  // 100
console.log(height); // 200
console.log(item1);  // Cake
console.log(item2);  // Donut

*/

// -------------------------------------------------------------------------------------------



// 智能函数参数
    //- 有时，一个函数有很多参数，其中大部分的参数都是可选的。对用户界面来说更是如此。
    //- 想象一个创建菜单的函数。它可能具有宽度参数，高度参数，标题参数和项目列表等。


// 下面是实现这种函数的一个很不好的写法：
    //- function showMenu(title = "Untitled", width = 200, height = 100, items = []) {
    //-     // ...
    //- }


// 在实际开发中存在一个问题就是你怎么记得住这么多参数的顺序。
// 通常集成开发环境工具（IDE）会尽力帮助我们，特别是当代码有良好的文档注释的时候，
// 但是…… 另一个问题就是，当大部分的参数采用默认值就好的情况下，怎么调用这个函数。

// 难道像这样？
    // 在采用默认值就可以的位置设置 undefined
    //- showMenu("My Menu", undefined, undefined, ["Item1", "Item2"])


// 解构赋值语法前来救援
// 我们可以把所有参数当作一个对象来传递，然后函数马上把这个对象解构成多个变量：

// 我们传递一个对象给函数
let options = {
  title: "My menu",
  items: ["Item1", "Item2"]
};

// ……然后函数马上把对象展开成变量
function showMenu1({title = "Untitled", width = 200, height = 100, items = []}) {
  // title, items – 提取于 options，
  // width, height – 使用默认值
  console.log( `${title} ${width} ${height}` ); // My Menu 200 100
  console.log( items ); // Item1, Item2
}

showMenu1(options);


// 我们同样可以使用带有嵌套对象和冒号映射的更加复杂的解构：
options = {
  title: "My menu",
  items: ["Item1", "Item2"]
};

function showMenu2({
  title = "Untitled",
  width: w = 100,  // width goes to w
  height: h = 200, // height goes to h
  items: [item1, item2] // items first element goes to item1, second to item2
}) {
  console.log( `${title} ${w} ${h}` ); // My Menu 100 200
  console.log( item1 ); // Item1
  console.log( item2 ); // Item2
}
showMenu2(options);


// 完整语法和解构赋值是一样的：
    //- function({
    //-     incomingProperty: varName = defaultValue
    //-     // ...
    //- })

// 对于参数对象，属性 incomingProperty 对应的变量是 varName，默认值是 defaultValue。

// 请注意，这种解构假定了 showMenu() 函数确实存在参数。如果我们想让所有的参数都使用默认值，那我们应该传递一个空对象：

// showMenu({}); // 不错，所有值都取默认值

// showMenu(); // 这样会导致错误

// 我们可以通过指定空对象 {} 为整个参数对象的默认值来解决这个问题：

function showMenu3({ title = "Menu", width = 100, height = 200 } = {}) {
  console.log( `${title} ${width} ${height}` );
}

showMenu3(); // Menu 100 200

// 在上面的代码中，整个参数对象的默认是 {}，因此总会有内容可以用来解构。



// -------------------------------------------------------------------------------------------



// 我们有一个对象：
user = {
    sex: "man",
    years: 22
};

// 写一个解构赋值语句使得：

// name 属性赋值给变量 name。
// years 属性赋值给变量 age。
// isAdmin 属性赋值给变量 isAdmin（如果属性缺失则取默认值 false）。

// 下面是赋值完成后的值的情况：
  
// 等号左侧是你的代码
// ... = user
let {sex, years: age, isAdmin = false} = user
console.log( sex); // man
console.log( age ); // 22
console.log( isAdmin ); // false




// --------------------------------------------------



// 最高薪资

// 这儿有一个 salaries 对象：
let salaries = {
  "John": 100,
  "Pete": 300,
  "Mary": 250
};
// 新建一个函数 topSalary(salaries)，返回收入最高的人的姓名。

// 如果 salaries 是空的，函数应该返回 null。
// 如果有多个收入最高的人，返回其中任意一个即可。
// P.S. 使用 Object.entries 和解构语法来遍历键/值对。

function topSalary(salaries) {
    let max = 0;
    let maxName = null;
    for (let [name, sal] of Object.entries(salaries)) {
        if (max < sal) {
            max = sal;
            maxName = name;
        }
    }
    return maxName;
}

console.log( topSalary(salaries) ); // Pete