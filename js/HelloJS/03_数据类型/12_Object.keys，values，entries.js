
// Object.keys，values，entries
    // 我们认识了 map.keys()，map.values() 和 map.entries() 方法。

    // 它们支持：
        //- Map
        //- Set
        //- Array

    // 普通对象也支持类似的方法，但是语法上有一些不同。



// -------------------------------------------------------------------------------------------




// 对于普通对象，下列这些方法是可用的：
    //- Object.keys(obj) —— 返回一个包含该对象所有的键的数组。
    //- Object.values(obj) —— 返回一个包含该对象所有的值的数组。
    //- Object.entries(obj) —— 返回一个包含该对象所有 [key, value] 键值对的数组。

// 但是请注意区别（比如说跟 map 的区别）：
    /*
                Map	           Object
        调用语法	map.keys()	    Object.keys(obj)，而不是 obj.keys()
        返回值	    可迭代项	    “真正的”数组

        第一个区别是，对于对象我们使用的调用语法是 Object.keys(obj)，而不是 obj.keys()。

        为什么会这样？
            主要原因是灵活性。请记住，在 JavaScript 中，对象是所有复杂结构的基础。
            因此，我们可能有一个自己创建的对象，比如 data，并实现了它自己的 data.values() 方法。
            同时，我们依然可以对它调用 Object.values(data) 方法。

        第二个区别是 Object.* 方法返回的是“真正的”数组对象，而不只是一个可迭代项。这主要是历史原因。

    */

// 举个例子：

//- let user = {
//-   name: "John",
//-   age: 30
//- };

//- Object.keys(user) = ["name", "age"]
//- Object.values(user) = ["John", 30]
//- Object.entries(user) = [ ["name","John"], ["age",30] ]

// 这里有一个使用 Object.values 来遍历属性值的例子：

let user = {
  name: "John",
  age: 30
};

// 遍历所有的值
for (let value of Object.values(user)) {
  console.log(value); // John, 30
}

// Object.keys/values/entries 会忽略 symbol 属性
// 就像 for..in 循环一样，这些方法会忽略使用 Symbol(...) 作为键的属性。

// 通常这很方便。但是，如果我们也想要 Symbol 类型的键，
// 那么这儿有一个单独的方法 Object.getOwnPropertySymbols，它会返回一个只包含 Symbol 类型的键的数组。
// 另外，还有一种方法 Reflect.ownKeys(obj)，它会返回 所有 键。




// -------------------------------------------------------------------------------------------




// 转换对象
    //- 对象缺少数组存在的许多方法，例如 map 和 filter 等。

// 如果我们想应用它们，那么我们可以使用 Object.entries，然后使用 Object.fromEntries：

// 使用 Object.entries(obj) 从 obj 获取由键/值对组成的数组。

// 对结果数组使用 Object.fromEntries(array) 方法，将结果转回成对象。

// 例如，我们有一个带有价格的对象，并想将它们加倍
let prices = {
  banana: 1,
  orange: 2,
  meat: 4,
};

let doublePrices = Object.fromEntries(
  // 转换为数组，之后使用 map 方法，然后通过 fromEntries 再转回到对象
  Object.entries(prices).map(([key, value]) => [key, value * 2])
);

console.log(doublePrices.meat); // 8




// -------------------------------------------------------------------------------------------




// 属性求和

// 有一个带有任意数量薪水的 salaries 对象。

// 编写函数 sumSalaries(salaries)，该函数使用 Object.values 和 for..of 循环返回所有薪水的总和。

// 如果 salaries 是空对象，那么结果必须是 0。


let salaries = {
  "John": 100,
  "Pete": 300,
  "Mary": 250
};

console.log( sumSalaries(salaries) ); // 650

function  sumSalaries(salaries) {
    let sum = 0;
    for (let val of Object.values(salaries)) {
        sum += val;
    }
    return sum;
    // return Object.values(salaries).reduce((a, b) => a + b, 0) // 650
}




// -------------------------------------------------------------------------------------------



// 计算属性数量

// 写一个函数 count(obj)，该函数返回对象中的属性的数量：

user = {
  name: 'John',
  age: 30
};

console.log( count(user) ); // 2

// 试着使代码尽可能简短。

// P.S. 忽略 Symbol 类型属性，只计算“常规”属性。

function count(obj) {
    return Object.keys(obj).length;
}













