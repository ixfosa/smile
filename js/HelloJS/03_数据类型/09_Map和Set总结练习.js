// 总结
    //- Map —— 是一个带键的数据项的集合。

// 方法和属性如下：
    //- new Map([iterable]) —— 创建 map，可选择带有 [key,value] 对的 iterable（例如数组）来进行初始化。
    //- map.set(key, value) —— 根据键存储值。
    //- map.get(key) —— 根据键来返回值，如果 map 中不存在对应的 key，则返回 undefined。
    //- map.has(key) —— 如果 key 存在则返回 true，否则返回 false。
    //- map.delete(key) —— 删除指定键的值。
    //- map.clear() —— 清空 map 。
    //- map.size —— 返回当前元素个数。

// 与普通对象 Object 的不同点：
    //- 任何键、对象都可以作为键。
    //- 有其他的便捷方法，如 size 属性。
    //- Set —— 是一组唯一值的集合。

// 方法和属性：
    //- new Set([iterable]) —— 创建 set，可选择带有 iterable（例如数组）来进行初始化。
    //- set.add(value) —— 添加一个值（如果 value 存在则不做任何修改），返回 set 本身。
    //- set.delete(value) —— 删除值，如果 value 在这个方法调用的时候存在则返回 true ，否则返回 false。
    //- set.has(value) —— 如果 value 在 set 中，返回 true，否则返回 false。
    //- set.clear() —— 清空 set。
    //- set.size —— 元素的个数。

// 在 Map 和 Set 中迭代总是按照值插入的顺序进行的，所以我们不能说这些集合是无序的，
// 但是我们不能对元素进行重新排序，也不能直接按其编号来获取元素。






// ---------------------------------------------------------------








// 过滤数组中的唯一元素

// 定义 arr 为一个数组。


// 创建一个函数 unique(arr)，该函数返回一个由 arr 中所有唯一元素所组成的数组。

// P.S. 这里用到了 string 类型，但其实可以是任何类型的值。

// P.S. 使用 Set 来存储唯一值。

function unique(arr) {
    /* 你的代码 */
    let set = new Set();
    arr.forEach(element => {
        (!set.has(element)) ? set.add(element) : "";
    });
    return set
}

let values = ["Hare", "Krishna", "Hare", "Krishna",
  "Krishna", "Krishna", "Hare", "Hare", ":-O"
];

console.log( unique(values) ); // Hare, Krishna, :-O




// ---------------------------------------------------------------




// 过滤字谜（anagrams）

// Anagrams 是具有相同数量相同字母但是顺序不同的单词。

// 例如：
    //- nap - pan
    //- ear - are - era
    //- cheaters - hectares - teachers

//- 写一个函数 aclean(arr)，它返回被清除了字谜（anagrams）的数组。

// 例如：

let arr = ["nap", "teachers", "cheaters", "PAN", "ear", "era", "hectares"];

// Set(3) { 'anp', 'aceehrst', 'aer' }
console.log( aclean(arr) ); // "nap,teachers,ear" or "PAN,cheaters,era"

// 对于所有的字谜（anagram）组，都应该保留其中一个词，但保留的具体是哪一个并不重要。

function aclean(arr) {
    let set = new Set();
    arr.forEach((itme) => set.add(itme.toLowerCase().split("").sort().join("")));
    return set;
}


// ---------------------------------------------------------------



// 迭代键

// 我们期望使用 map.keys() 得到一个数组，然后使用特定的方法例如 .push 等，对其进行处理。

// 但是运行不了：

let map = new Map();

map.set("name", "John");

console.log(map); // Map(1) { 'name' => 'John' }

// let keys = map.keys();

let keys = Array.from(map.keys());

// Error: keys.push is not a function
keys.push("more"); 

console.log(keys); // [ 'name', 'more' ]


// 为什么？我们应该如何修改代码让 keys.push 工作？
    //- let keys = Array.from(map.keys());