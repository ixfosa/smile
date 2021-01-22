"user strict";

// 可选链 "?."
    // 可选链 ?. 是一种访问嵌套对象属性的安全的方式。即使中间的属性不存在，也不会出现错误。

// “不存在的属性”的问题
    // 举个例子，假设我们有很多个 user 对象，其中存储了我们的用户数据。
    // 大多数用户的地址都存储在 user.address 中，街道地址存储在 user.address.street 中，但有些用户没有提供这些信息。
    // 在这种情况下，当我们尝试获取 user.address.street，而该用户恰好没提供地址信息，我们则会收到一个错误：

let user = {}; // 一个没有 "address" 属性的 user 对象
console.log(user.address); // undefined
// console.log(user.address.street); // TypeError: Cannot read property 'street' of undefined

// 因为 user.address 为 undefined，尝试读取 user.address.street 会失败，并收到一个错误。


// 可以在访问该值的属性之前，使用 if 或条件运算符 ? 对该值进行检查，像这样：
console.log( user.address ? user.address.street : undefined ); //undefined

// "user.address" 在代码中出现了两次。对于嵌套层次更深的属性就会出现更多次这样的重复，这就是问题了。
// 例如，让我们尝试获取 user.address.street.name。
// 我们既需要检查 user.address，又需要检查 user.address.street：
console.log(user.address ? user.address.street ? user.address.street.name : null : null); // null



// 一种更好的实现方式，就是使用 && 运算符：
console.log( user.address && user.address.street && user.address.street.name ); // undefined（不报错）

// 依次对整条路径上的属性使用与运算进行判断，以确保所有节点是存在的（如果不存在，则停止计算），但仍然不够优雅。
// 就像你所看到的，在代码中我们仍然重复写了好几遍对象属性名。例如在上面的代码中，user.address 被重复写了三遍。




// -------------------------------------------------------------------------------------------





// 可选链
    //- 如果可选链 ?. 前面的部分是 undefined 或者 null，它会停止运算并返回该部分。

// 现在如果一个属性既不是 null 也不是 undefined，那么它就“存在”。

//换句话说，例如 value?.prop：
    // 如果 value 存在，则结果与 value.prop 相同，
    // 否则（当 value 为 undefined/null 时）则返回 undefined。

    // 请注意：?. 语法使其前面的值成为可选值，但不会对其后面的起作用。

// 下面这是一种使用 ?. 安全地访问 user.address.street 的方式：
let user2 = {}; // user 没有 address 属性
console.log( user2?.address?.street ); // undefined

// 即使 对象 user 不存在，使用 user?.address 来读取地址也没问题：
user2 = null;
console.log( user2?.address ); // undefined
// console.log( user2?.address.street ); // TypeError: Cannot read property 'street' of undefined


// ?. 前的变量必须已声明
    //- 如果未声明变量 user，那么 user?.anything 会触发一个错误：

// ReferenceError: user is not defined
//- user3?.address;  
// ?. 前的变量必须已声明（例如 let/const/var user 或作为一个函数参数）。可选链仅适用于已声明的变量。



// 短路效应
    //- 正如前面所说的，如果 ?. 左边部分不存在，就会立即停止运算（“短路效应”）。

// 所以，如果后面有任何函数调用或者副作用，它们均不会执行。
let user3 = null;
let x = 0;

user3?.sayHi(x++); // 没有 "sayHi"，因此代码执行没有触达 x++

console.log(x); // 0，值没有增加




// -------------------------------------------------------------------------------------------



// 其它变体：?.()，?.[]
    //- 可选链 ?. 不是一个运算符，而是一个特殊的语法结构。它还可以与函数和方括号一起使用。


// 例如，将 ?.() 用于调用一个可能不存在的函数。
// 在下面这段代码中，有些用户具有 admin 方法，而有些没有：
let userAdmin = {
    sayHi() {
        console.log("userAdmin hello!!!");
    }
};

let userGuest = {};

userAdmin.sayHi?.(); // userAdmin hello!!!

userGuest.sayHi?.(); // 啥都没有（没有这样的方法）


// 使用方括号 [] 而不是点符号 . 来访问属性，语法 ?.[] 也可以使用。
// 跟前面的例子类似，它允许从一个可能不存在的对象上安全地读取属性。

let person = {
    name: "long",
}
let key = "name";

console.log(person?.["name"]); // long
console.log(person?.[key]); // long

console.log(person?.["age"]); //undefined



// 将 ?. 跟 delete 一起使用：
delete person?.name; 


// 使用 ?. 来安全地读取或删除，但不能写入
    //- 可选链 ?. 不能用在赋值语句的左侧

// let user = null;
// user?.name = "long"; // Error，不起作用

// 因为它在计算的是 undefined = "John"



// -------------------------------------------------------------------------------------------


// 总结
    //- 可选链 ?. 语法有三种形式：
        //- obj?.prop —— 如果 obj 存在则返回 obj.prop，否则返回 undefined。
        //- obj?.[prop] —— 如果 obj 存在则返回 obj[prop]，否则返回 undefined。
        //- obj.method?.() —— 如果 obj.method 存在则调用 obj.method()，否则返回 undefined。

    //- 正如我们所看到的，这些语法形式用起来都很简单直接。?. 检查左边部分是否为 null/undefined，
    // 如果不是则继续运算。

// ?. 链使我们能够安全地访问嵌套属性。

// 应该谨慎地使用 ?.，仅在当左边部分不存在也没问题的情况下使用为宜。
// 以保证在代码中有编程上的错误出现时，也不会对我们隐藏。

