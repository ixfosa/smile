// Map and Set（映射和集合）

// 以下复杂的数据结构：
    //- 存储带键的数据（keyed）集合的对象。
    //- 存储有序集合的数组。
    //- Map 和 Set。

// Map
    //- Map 是一个带键的数据项的集合，就像一个 Object 一样。 
    //- 但是它们最大的差别是 Map 允许任何类型的键（key）。

// 它的方法和属性如下：
    //- new Map() —— 创建 map。
    //- map.set(key, value) —— 根据键存储值。
    //- map.get(key) —— 根据键来返回值，如果 map 中不存在对应的 key，则返回 undefined。
    //- map.has(key) —— 如果 key 存在则返回 true，否则返回 false。
    //- map.delete(key) —— 删除指定键的值。
    //- map.clear() —— 清空 map。
    //- map.size —— 返回当前元素个数。


// 举个例子：
let map = new Map();
map.set(1, 1);          // 字符串键
map.set("2", "String"); // 字符串键
map.set(true, false);   // 布尔值键

// 链式调用
// 每一次 map.set 调用都会返回 map 本身，所以我们可以进行“链式”调用：
//- map.set('1', 'str1')
//-   .set(1, 'num1')
//-   .set(true, 'bool1');

map.set(1, "ixfosa");   // 覆盖

console.log(map); // Map(3) { 1 => 'ixfosa', '2' => 'String', true => false }

// map与对象不同，键不会被转换成字符串。键可以是任何类型。
console.log(map.get(1));  // ixfosa
console.log(map.get("2")); // String
console.log(map.size); // 3


// map[key] 不是使用 Map 的正确方式
    //- 虽然 map[key] 也有效，例如我们可以设置 map[key] = 2，
    //  这样会将 map 视为 JavaScript 的 plain object，因此它暗含了所有相应的限制（没有对象键等）。

    //- 所以我们应该使用 map 方法：set 和 get 等。


// Map 还可以使用对象作为键。
let fo = { name: "ixfosa" };

// 存储每个用户的来访次数
let visitsCountMap = new Map();

// john 是 Map 中的键
visitsCountMap.set(fo, 123);

console.log( visitsCountMap.get(fo) ); // 123



// 使用对象作为键是 Map 最值得注意和重要的功能之一。对于字符串键，Object（普通对象）也能正常使用，但对于对象键则不行。
let long = { name: "long" };
let obj = {}; // 尝试使用对象
obj[long] = 123; // 尝试将 john 对象作为键
// 是写成了这样!
console.log( obj["[object Object]"] ); // 123


// 添加 toString()
let zhong = {
    name: "zhong",
    toString() {
        return "zhong"
    }
}
obj[zhong] = 123;

console.log(obj["zhong"]); // 123


// 因为 obj 是一个对象，它会将所有的键如 fo 转换为字符串，所以我们得到字符串键 "[object Object]"。
// 这显然不是我们想要的结果。


// Map 是怎么比较键的？
    //- Map 使用 SameValueZero 算法来比较键是否相等。
    //- 它和严格等于 === 差不多，但区别是 NaN 被看成是等于 NaN。所以 NaN 也可以被用作键。
// 这个算法不能被改变或者自定义。





// -------------------------------------------------------------------------------------------





// Map 迭代
    //- 如果要在 map 里使用循环，可以使用以下三个方法：
        //- map.keys() —— 遍历并返回所有的键（returns an iterable for keys），
        //- map.values() —— 遍历并返回所有的值（returns an iterable for values），
        //- map.entries() —— 遍历并返回所有的实体（returns an iterable for entries）[key, value]，
        //  for..of 在默认情况下使用的就是这个。

let recipeMap = new Map([
    ['cucumber', 500],
    ['tomatoes', 350],
    ['onion',    50]
]);

// 遍历所有的键（vegetables）
for (let vegetable of recipeMap.keys()) {
    console.log(vegetable); // cucumber, tomatoes, onion
}

// 遍历所有的值（amounts）
for (let amount of recipeMap.values()) {
    console.log(amount); // 500, 350, 50
}


// 遍历所有的实体 [key, value]
for (let entry of recipeMap.entries()) {
    console.log(entry); 
    // [ 'cucumber', 500 ]
    // [ 'tomatoes', 350 ]
    // [ 'onion', 50 ]
}


// 遍历所有的实体 [key, value]
for (let entry of recipeMap) { // 与 recipeMap.entries() 相同
    console.log(entry); 
    // [ 'cucumber', 500 ]
    // [ 'tomatoes', 350 ]
    // [ 'onion', 50 ]
}


// 使用插入顺序
    //- 迭代的顺序与插入值的顺序相同。与普通的 Object 不同，Map 保留了此顺序。

// 除此之外，Map 有内置的 forEach 方法，与 Array 类似：

// 对每个键值对 (key, value) 运行 forEach 函数
recipeMap.forEach( (value, key, map) => {
    console.log(`${key}: ${value}`); 
    // cucumber: 500
    // tomatoes: 350
    // onion: 50
});
    


// -------------------------------------------------------------------------------------------




// Object.entries：从对象创建 Map
    //- 当创建一个 Map 后，我们可以传入一个带有键值对的数组（或其它可迭代对象）来进行初始化，如下所示：

// 键值对 [key, value] 数组
let map1 = new Map([
  ['1',  'str1'],
  [1,    'num1'],
  [true, 'bool1']
]);

console.log( map1.get('1') ); // str1


// 如果我们想从一个已有的普通对象（plain object）来创建一个 Map，
// 那么我们可以使用内建方法 Object.entries(obj)，该返回对象的键/值对数组，
// 该数组格式完全按照 Map 所需的格式。

// 所以可以像下面这样从一个对象创建一个 Map：
obj = {
    name: "fo",
    age: 30
};

map = new Map(Object.entries(obj));

console.log( map.get('name') ); // fo

// 这里，Object.entries 返回键/值对数组：[ ["name","John"], ["age", 30] ]。
// 这就是 Map 所需要的格式。





// -------------------------------------------------------------------------------------------




// Object.fromEntries：从 Map 创建对象
    //- 使用 Object.entries(obj) 从普通对象（plain object）创建 Map。
    //- Object.fromEntries 方法的作用是相反的：给定一个具有 [key, value] 键值对的数组，
    //  它会根据给定数组创建一个对象：

let prices = Object.fromEntries([
  ['banana', 1],
  ['orange', 2],
  ['meat', 4]
]);

// 现在 prices = { banana: 1, orange: 2, meat: 4 }

console.log(prices.orange); // 2

// 可以使用 Object.fromEntries 从 Map 得到一个普通对象（plain object）。
// 例如，我们在 Map 中存储了一些数据，但是我们需要把这些数据传给需要普通对象（plain object）的第三方代码。

let map2 = new Map();
map2.set('banana', 1);
map2.set('orange', 2);
map2.set('meat', 4);

let obj2 = Object.fromEntries(map2.entries()); // 创建一个普通对象（plain object）(*)

// obj2 = { banana: 1, orange: 2, meat: 4 }

console.log(obj2.orange); // 2


// 调用 map.entries() 将返回一个可迭代的键/值对，这刚好是 Object.fromEntries 所需要的格式。
// 我们可以把带 (*) 这一行写得更短：
obj = Object.fromEntries(map2); // 省掉 .entries()
// 上面的代码作用也是一样的，因为 Object.fromEntries 期望得到一个可迭代对象作为参数，
// 而不一定是数组。并且 map 的标准迭代会返回跟 map.entries() 一样的键/值对。
// 因此，我们可以获得一个普通对象（plain object），其键/值对与 map 相同。

