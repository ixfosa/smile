
// 对象
    //- 可以通过使用带有可选 属性列表 的花括号 {…} 来创建对象。
    //- 一个属性就是一个键值对（“key: value”），其中键（key）是一个字符串（也叫做属性名），值（value）可以是任何值。


// 下面两种语法中的任一种来创建一个空的对象（“空柜子”）：
let user1 = new Object(); // “构造函数” 的语法
let user2 = {};  // “字面量” 的语法



// 文本和属性
// 我们可以在创建对象的时候，立即将一些属性以键值对的形式放到 {...} 中。
let person = {
    name: "梓晴",
    age: 17
};                      // 冒号

// 属性有键（或者也可以叫做“名字”或“标识符”），位于冒号 ":" 的前面，值在冒号的右边。
    //- 在 user 对象中，有两个属性：
        //- 第一个的键是 "name"，值是 "梓晴"。
        //- 第二个的键是 "age"，值是 17

// 可以随时添加、删除和读取文件。

// 属性的值可以是任意类型
person.sex = "女";

// 小陷阱：一个名为 __proto__ 的属性。我们不能将它设置为一个非对象的值：
let obj0 = {};
obj0.__proto__ = 5; // 分配一个数字
console.log("__proto__" + obj0.__proto__); // __proto__[object Object] — 值为对象，与预期结果不同


// 读取
// 姓名：梓晴 - 年龄：17 - 性别：女
console.log(`姓名：${person.name} - 年龄：${person.age} - 性别：${person.sex}`);

// 用 delete 操作符移除属性：
delete person.age;

// 读取
// 姓名：梓晴 - 年龄：undefined - 性别：女
console.log(`姓名：${person.name} - 年龄：${person.age} - 性别：${person.sex}`);



// 可以用多字词语来作为属性名，但必须给它们加上引号：
let person2 = {
    name: "梓晴",
    age: 17,
   
    "likes birds": true,  // 多词属性名必须加引号, 列表中的最后一个属性应以逗号结尾：
};
// 尾随（trailing）或悬挂（hanging）逗号。这样便于我们添加、删除和移动属性，因为所有的行都是相似的。



// -------------------------------------------------------------------------------------------




// 使用 const 声明的对象是可以被修改的
let person3 = {
    name: "梓晴",
};
person3.name = "如也";
console.log(person3. name);


// const 声明仅固定了 user 的值，而不是值（该对象）里面的内容。
// 仅当我们尝试将 user=... 作为一个整体进行赋值时，const 会抛出错误。



// -------------------------------------------------------------------------------------------



// 方括号
    //- 对于多词属性，点操作就不能用了：

// 这将提示有语法错误
    //- user.likes birds = true

// 请注意方括号中的字符串要放在引号中，单引号或双引号都可以。
// 使用方括号，可用于任何字符串：
let person4 = {};

// 设置
// 请注意方括号中的字符串要放在引号中，单引号或双引号都可以。
person4["likes birds"] = true;

// 读取
console.log(person4["likes birds"]);

// 删除
delete person4["likes birds"];


// 方括号同样提供了一种可以通过任意表达式来获取属性名的方法 —— 跟语义上的字符串不同 —— 比如像类似于下面的变量：
let key  = "likes birds";
// 跟 user["likes birds"] = false; 一样
person4[key] = false;
// 访问
console.log(person4[key]);

// // 点符号不能以类似的方式使用：
console.log(person4.key); // undefined



// -------------------------------------------------------------------------------------------




// 计算属性
    //- 当创建一个对象时，我们可以在对象字面量中使用方括号。这叫做 计算属性。

let fruit = "apple";
let bag = {
    [fruit]: 5, // 属性名是从 fruit 变量中得到的
};
console.log( bag.apple ); // 5 如果 fruit="apple"

// 计算属性的含义很简单：[fruit] 含义是属性名应该从 fruit 变量中获取。
// 所以，如果一个用户输入 "apple"，bag 将变为 {apple: 5}。


// 本质上，这跟下面的语法效果相同：
let fruit2 = "apple";
let bag2 = {};
// 从 fruit 变量中获取值
bag2[fruit2] = 5;


// 可以在方括号中使用更复杂的表达式：

let fruit3 = 'apple';
let bag3 = {
  [fruit3 + 'Computers']: 5 // bag.appleComputers = 5
};



// -------------------------------------------------------------------------------------------




// 属性值简写
    //- 在实际开发中，我们通常用已存在的变量当做属性名。
function makeUser1(name, age) {
    return {
        name: name,
        age: age,
        // ……其他的属性
    };
}
    
let fo1 = makeUser1("fo1", 22);
console.log(fo1.name); // fo1

function makeUser2(name, age) {
    return {
      name, // 与 name: name 相同
      age,  // 与 age: age 相同
      // ...
    };
}
let fo2 = makeUser1("fo2", 22);
console.log(fo2.name); // fo2



// -------------------------------------------------------------------------------------------




// 属性名简写方式和正常方式混用：
let name = "fo3"
let fo3 = {
  name,  // 与 name:name 相同
  age: 30
};
console.log(fo3.name); // fo3



// 属性名称限制
    //- 变量名不能是编程语言的某个保留字，如 “for”、“let”、“return” 等……
    //- 但对象的属性名并不受此限制：

// 这些属性都没问题
let obj = {
    for: 1,
    let: 2,
    return: 3
};
  
console.log( obj.for + obj.let + obj.return );  // 6

// 属性命名没有限制。属性名可以是任何字符串或者 symbol。

// 其他类型会被自动地转换为字符串。

// 例如，当数字 0 被用作对象的属性的键时，会被转换为字符串 "0"：
let obj2 = {
    0: "test" // 等同于 "0": "test"
};
  
  // 都会输出相同的属性（数字 0 被转为字符串 "0"）
console.log( obj2["0"] ); // test
console.log( obj2[0] ); // test (相同的属性)



// -------------------------------------------------------------------------------------------



// 属性存在性测试，“in” 操作符

// JavaScript 的对象有一个需要注意的特性：能够被访问任何属性。即使属性不存在也不会报错！

// 读取不存在的属性只会得到 undefined。所以我们可以很容易地判断一个属性是否存在：
let o = {};
console.log(o.no === undefined); // true


//检查属性是否存在的操作符 "in"。

// 语法是：
    //- "key" in object

let o1 = { name: "fo", age: 22 };

// in 的左边必须是 属性名。通常是一个带引号的字符串。
console.log( "age" in o1 ); // true，o1.age 存在
console.log( "sex" in o1 ); // false，o1.sex 不存在。

// 如果省略引号，就意味着左边是一个变量，它应该包含要判断的实际属性名。例如：
let age = "age";
console.log( age in o1 ); // true



// 为何会有 in 运算符呢？与 undefined 进行比较来判断还不够吗？
    //- 大部分情况下与 undefined 进行比较来判断就可以了。
    //- 但有一个例外情况，这种比对方式会有问题，但 in 运算符的判断结果仍是对的。

    //-那就是属性存在，但存储的值是 undefined 的时候：

let o2 = {
  test: undefined
};

console.log( o2.test ); // 显示 undefined，所以属性不存在？

console.log( "test" in o2 ); // true，属性存在！

// 在上面的代码中，属性 obj.test 事实上是存在的，所以 in 操作符检查通过。
// 这种情况很少发生，因为通常情况下不应该给对象赋值 undefined。我们通常会用 null 来表示未知的或者空的值。
// 因此，in 运算符是代码中的特殊来宾。



// -------------------------------------------------------------------------------------------




// “for…in” 循环
    //- 为了遍历一个对象的所有键（key），可以使用一个特殊形式的循环：for..in。
    //- 这跟我们在前面学到的 for(;;) 循环是完全不一样的东西。

    // 语法：

        //- for (key in object) {
        //-     // 对此对象属性中的每个键执行的代码
        //- }

// 让我们列出 admin 所有的属性：
let admin = {
    name: "fo",
    age: 22,
    gender: "男",
}

for (let key in admin) {
    // admin.key 方式  err  
    // console.log(`${key}: ${admin.key}`);
    console.log(`${key}: ${admin[key]}`);
}



// -------------------------------------------------------------------------------------------



// 像对象一样排序
    //- 对象有顺序吗？换句话说，如果我们遍历一个对象，我们获取属性的顺序是和属性添加时的顺序相同吗？这靠谱吗？
    //- 简短的回答是：“有特别的顺序”：整数属性会被进行排序，其他属性则按照创建的顺序显示。详情如下：

//- 例如，让我们考虑一个带有电话号码的对象：
let codes = {
  "49": "Germany",
  "41": "Switzerland",
  "44": "Great Britain",
  // ..,
  "1": "USA"
};

for(let code in codes) {
  console.log(code); // 1, 41, 44, 49
}


// 整数属性？那是什么？
    //- 这里的“整数属性”指的是一个可以在不做任何更改的情况下与一个整数进行相互转换的字符串。
    //- 所以，“49” 是一个整数属性名，因为我们把它转换成整数，再转换回来，它还是一样的。但是 “+49” 和 “1.2” 就不行了：

// Math.trunc 是内置的去除小数部分的方法。
console.log( String(Math.trunc(Number("49"))) ); // "49"，相同，整数属性
console.log( String(Math.trunc(Number("+49"))) ); // "49"，不同于 "+49" ⇒ 不是整数属性
console.log( String(Math.trunc(Number("1.2"))) ); // "1"，不同于 "1.2" ⇒ 不是整数属性
console.log(+"+66") // 66

// 如果属性名不是整数，那它们就按照创建时的顺序来排序，例如：
let user3 = {
    name: "fo",
    surname: "lian"
};

user3.age = 22; // 增加一个

// 非整数属性是按照创建的顺序来排列的
for (let prop in user3) {
    console.log( prop ); // name, surname, age
}



// 所以，为了解决电话号码的问题，我们可以使用非整数属性名来 欺骗 程序。只需要给每个键名加一个加号 "+" 前缀就行了。


let codes2 = {
    "+49": "Germany",
    "+41": "Switzerland",
    "+44": "Great Britain",
    // ..,
    "+1": "USA"
};

for (let code in codes2) {
    console.log( +code ); // 49, 41, 44, 1
}
// -------------------------------------------------------------------------------------------



/**
    总结
        对象是具有一些特殊特性的关联数组。

    它们存储属性（键值对），其中：
        属性的键必须是字符串或者 symbol（通常是字符串）。
        值可以是任何类型。

    我们可以用下面的方法访问属性
        点符号: obj.property。
        方括号 obj["property"]，方括号允许从变量中获取键，例如 obj[varWithKey]。
        
    其他操作：
        删除属性：delete obj.prop。
        检查是否存在给定键的属性："key" in obj。
        遍历对象：for(let key in obj) 循环。
 */