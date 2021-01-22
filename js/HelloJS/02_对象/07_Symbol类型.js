//Symbol 类型
    //- 根据规范，对象的属性键只能是字符串类型或者 Symbol 类型。
    //- 不是 Number，也不是 Boolean，只有字符串或 Symbol 这两种类型。


// “Symbol” 值表示唯一的标识符。


// 可以使用 Symbol() 来创建这种类型的值：
"use strict";
let id = Symbol();  // id 是 symbol 的一个实例化对象
console.log(id);  // Symbol()

// 可以给 Symbol 一个描述（也称为 Symbol 名），这在代码调试时非常有用：
let id0 = Symbol("id"); // id 是描述为 "id" 的 Symbol
console.log(id0.description) // id


// Symbol 保证是唯一的。即使我们创建了许多具有相同描述的 Symbol，它们的值也是不同。
// 描述只是一个标签，不影响任何东西。

// 例如，这里有两个描述相同的 Symbol —— 它们不相等：
let id1 = Symbol("id");
let id2 = Symbol("id");

console.log(id1 == id2); // false





// -------------------------------------------------------------------------------------------



// Symbol 不会被自动转换为字符串
    //- JavaScript 中的大多数值都支持字符串的隐式转换。
    //- 例如，我们可以 alert 任何值，都可以生效。Symbol 比较特殊，它不会被自动转换。

let id3 = Symbol("id");
// alert(id3); // 类型错误：无法将 Symbol 值转换为字符串。
// 这是一种防止混乱的“语言保护”，因为字符串和 Symbol 有本质上的不同，不应该意外地将它们转换成另一个。

// 如果我们真的想显示一个 Symbol，我们需要在它上面调用 .toString()，如下所示：
console.log(id3.toString()); // Symbol(id)，现在它有效了

// 或者获取 symbol.description 属性，只显示描述（description）：
console.log(id3.description); // id





// -------------------------------------------------------------------------------------------




// “隐藏”属性
    //- Symbol 允许我们创建对象的“隐藏”属性，代码的任何其他部分都不能意外访问或重写这些属性。

// 例如，如果我们使用的是属于第三方代码的 user 对象，我们想要给它们添加一些标识符。
// 我们可以给它们使用 Symbol 键：
let user = {
    name: "ixfsoa",
};

let sign = Symbol("sign");

user[sign] = 1;

console.log(user[sign]); // 我们可以使用 Symbol 作为键来访问数据




// -------------------------------------------------------------------------------------------



// 对象字面量中的 Symbol
    //- 如果我们要在对象字面量 {...} 中使用 Symbol，则需要使用方括号把它括起来

let flag = Symbol("flag");

let person = {
    name: "ixfosa",
    age: 22,
    [flag]: 666,       // 这是因为我们需要变量 id 的值作为键，而不是字符串 “id”。
}


// Symbol 在 for…in 中会被跳过
    //- Symbol 属性不参与 for..in 循环。
for (let key in person) {
    console.log(person[key]);  // ixfosa, 22
}

// 使用 Symbol 任务直接访问
console.log( "flag: " + person[flag] );



// Object.keys(person) 也会忽略它们。这是一般“隐藏符号属性”原则的一部分。
// 如果另一个脚本或库遍历我们的对象，它不会意外地访问到符号属性。

// 相反，Object.assign 会同时复制字符串和 symbol 属性：
let clone = Object.assign({}, person);
console.log(clone[flag]); // 666




// -------------------------------------------------------------------------------------------




//全局 symbol
    //- 通常所有的 Symbol 都是不同的，即使它们有相同的名字。但有时我们想要名字相同的 Symbol 具有相同的实体。
    //- 例如，应用程序的不同部分想要访问的 Symbol "id" 指的是完全相同的属性。

    //- 为了实现这一点，这里有一个 全局 Symbol 注册表。我们可以在其中创建 Symbol 并在稍后访问它们，
    //- 它可以确保每次访问相同名字的 Symbol 时，返回的都是相同的 Symbol。

    //- 要从注册表中读取（不存在则创建）Symbol，请使用 Symbol.for(key)。
    //- 该调用会检查全局注册表，如果有一个描述为 key 的 Symbol，则返回该 Symbol，
    //- 否则将创建一个新 Symbol（Symbol(key)），并通过给定的 key 将其存储在注册表中。

// 例如：
// 从全局注册表中读取
let mark  = Symbol.for("mark"); // 如果该 Symbol 不存在，则创建它

// 再次读取（可能是在代码中的另一个位置）
let again = Symbol.for("mark");

console.log(mark === again); // true

// 注册表内的 Symbol 被称为 全局 Symbol。
// 如果我们想要一个应用程序范围内的 Symbol，可以在代码中随处访问 —— 这就是它们的用途。




// Symbol.keyFor
    //- 对于全局 Symbol，不仅有 Symbol.for(key) 按名字返回一个 Symbol，
    //- 还有一个反向调用：Symbol.keyFor(sym)，它的作用完全反过来：通过全局 Symbol 返回一个名字。

// 通过 name 获取 Symbol
let sym = Symbol.for("name");
let sym2 = Symbol.for("id");

// 通过 Symbol 获取 name
console.log( Symbol.keyFor(sym) ); // name
console.log( Symbol.keyFor(sym2) ); // id


// Symbol.keyFor 内部使用全局 Symbol 注册表来查找 Symbol 的键。所以它不适用于非全局 Symbol。
// 如果 Symbol 不是全局的，它将无法找到它并返回 undefined。

//也就是说，任何 Symbol 都具有 description 属性。
let globalSymbol = Symbol.for("name");
let localSymbol = Symbol("name");

console.log( Symbol.keyFor(globalSymbol) ); // name，全局 Symbol
console.log( Symbol.keyFor(localSymbol) ); // undefined，非全局

console.log( localSymbol.description ); // name




// -------------------------------------------------------------------------------------------




// 系统 Symbol
    //- JavaScript 内部有很多“系统” Symbol，我们可以使用它们来微调对象的各个方面。

// 它们都被列在了 众所周知的 Symbol 表的规范中：
    //- Symbol.hasInstance
    //- Symbol.isConcatSpreadable
    //- Symbol.iterator
    //- Symbol.toPrimitive  允许我们将对象描述为原始值转换



// 总结
    //- Symbol 是唯一标识符的基本类型

    //- Symbol 是使用带有可选描述（name）的 Symbol() 调用创建的。

    //- Symbol 总是不同的值，即使它们有相同的名字。
    //- 如果我们希望同名的 Symbol 相等，那么我们应该使用全局注册表：Symbol.for(key) 返回（如果需要的话则创建）一个以 key 作为名字的全局 Symbol。
    //- 使用 Symbol.for 多次调用 key 相同的 Symbol 时，返回的就是同一个 Symbol。

    //- Symbol 有两个主要的使用场景：

        //- “隐藏” 对象属性。 如果我们想要向“属于”另一个脚本或者库的对象添加一个属性，我们可以创建一个 Symbol 并使用它作为属性的键。
        //- Symbol 属性不会出现在 for..in 中，因此它不会意外地被与其他属性一起处理。并且，它不会被直接访问，因为另一个脚本没有我们的 symbol。
        //- 因此，该属性将受到保护，防止被意外使用或重写。

        //- 因此我们可以使用 Symbol 属性“秘密地”将一些东西隐藏到我们需要的对象中，但其他地方看不到它。

        //- JavaScript 使用了许多系统 Symbol，这些 Symbol 可以作为 Symbol.* 访问。我们可以使用它们来改变一些内置行为。
        //- 例如，在本教程的后面部分，我们将使用 Symbol.iterator 来进行 迭代 操作，使用 Symbol.toPrimitive 来设置 对象原始值的转换 等等。

    //- 从技术上说，Symbol 不是 100% 隐藏的。有一个内置方法 Object.getOwnPropertySymbols(obj) 允许我们获取所有的 Symbol。
    //- 还有一个名为 Reflect.ownKeys(obj) 的方法可以返回一个对象的 所有 键，包括 Symbol。所以它们并不是真正的隐藏。
    //- 但是大多数库、内置方法和语法结构都没有使用这些方法。
