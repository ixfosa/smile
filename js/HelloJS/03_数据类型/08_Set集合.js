// Set
    //- Set 是一个特殊的类型集合 —— “值的集合”（没有键），它的每一个值只能出现一次。

// 它的主要方法如下：
    //- new Set(iterable) —— 创建一个 set，如果提供了一个 iterable 对象（通常是数组），
    //  将会从数组里面复制值到 set 中。

    //- set.add(value) —— 添加一个值，返回 set 本身
    //- set.delete(value) —— 删除值，如果 value 在这个方法调用的时候存在则返回 true ，否则返回 false。
    //- set.has(value) —— 如果 value 在 set 中，返回 true，否则返回 false。
    //- set.clear() —— 清空 set。
    //- set.size —— 返回元素个数。

// 它的主要特点是，重复使用同一个值调用 set.add(value) 并不会发生什么改变。
// 这就是 Set 里面的每一个值只出现一次的原因。

let set = new Set();

let u1 = {name: "fo"};
let u2 = {name: "long"};
let u3 = {name: "zhong"};

set.add(u1);
set.add(u2);
set.add(u3);
set.add(u2);
set.add(u3);

console.log(set.size);

console.log(set) // Set(3) { { name: 'fo' }, { name: 'long' }, { name: 'zhong' } }



for (let entry of set.entries()) {
    console.log(entry);
}

// set.keys() —— 遍历并返回所有的值
for (let key of set.keys()) {
    console.log(key.name);
   
}

// set.values() —— 与 set.keys() 作用相同
for (let val of set.values()) {
    console.log(val.name);
}




// -------------------------------------------------------------------------------------------





// Set 迭代（iteration）
    //- 我们可以使用 for..of 或 forEach 来遍历 Set：

let set2 = new Set(["oranges", "apples", "bananas"]);

for (let value of set2) console.log(value);

// 与 forEach 相同：
set2.forEach((value, valueAgain, set) => {
  console.log(value);
});


// 注意。forEach 的回调函数有三个参数：一个 value，然后是 同一个值 valueAgain，最后是目标对象。
// 没错，同一个值在参数里出现了两次。

// forEach 的回调函数有三个参数，是为了与 Map 兼容。
// 当然，这看起来确实有些奇怪。但是这对在特定情况下轻松地用 Set 代替 Map 很有帮助，反之亦然。

// Map 中用于迭代的方法在 Set 中也同样支持：
    //- set.keys() —— 遍历并返回所有的值（returns an iterable object for values），
    //- set.values() —— 与 set.keys() 作用相同，这是为了兼容 Map，
    //- set.entries() —— 遍历并返回所有的实体
    // （returns an iterable object for entries）[value, value]，它的存在也是为了兼容 Map。