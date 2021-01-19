
/**
    大于 / 小于：a > b，a < b。
    大于等于 / 小于等于：a >= b，a <= b。
    检查两个值的相等：a == b，请注意双等号 == 表示相等性检查，而单等号 a = b 表示赋值。
    检查两个值不相等。不相等在数学中的符号是 ≠，但在 JavaScript 中写成 a != b。
 */



 // 比较结果为 Boolean 类型
// 所有`比较运算`符均返回布尔值：
    //- true —— 表示“yes（是）”，“correct（正确）”或“the truth（真）”。
    //- false —— 表示“no（否）”，“wrong（错误）”或“not the truth（非真）”。

console.log( 2 > 1 );  // true（正确）
console.log( 2 == 1 ); // false（错误）
console.log( 2 != 1 ); // true（正确）

// 和其他类型的值一样，比较的结果可以被赋值给任意变量：
let result = 5 > 4; // 把比较的结果赋值给 result
console.log( result ); // true


// 字符串比较
// 在比较字符串的大小时，JavaScript 会使用“字典（dictionary）”或“词典（lexicographical）”顺序进行判定。
// 换言之，字符串是按字符（母）逐个进行比较的。
console.log( 'Z' > 'A' ); // true
console.log( 'Glow' > 'Glee' ); // true
console.log( 'Bee' > 'Be' ); // true

/**
    字符串的比较算法非常简单：
        -首先比较两个字符串的首位字符大小。
        -如果一方字符较大（或较小），则该字符串大于（或小于）另一个字符串。算法结束。
        -否则，如果两个字符串的首位字符相等，则继续取出两个字符串各自的后一位字符进行比较。
        -重复上述步骤进行比较，直到比较完成某字符串的所有字符为止。
        -如果两个字符串的字符同时用完，那么则判定它们相等，否则未结束（还有未比较的字符）的字符串更大。

    非真正的字典顺序，而是 Unicode 编码顺序
        在上面的算法中，比较大小的逻辑与字典或电话簿中的排序很像，但也不完全相同。
        比如说，字符串比较对字母大小写是敏感的。大写的 "A" 并不等于小写的 "a"。哪一个更大呢？
        实际上小写的 "a" 更大。这是因为在 JavaScript 使用的内部编码表中（Unicode），
        小写字母的字符索引值更大。我们会在 字符串 这章讨论更多关于字符串的细节。
 */


// 不同类型间的比较
// 当对不同类型的值进行比较时，JavaScript 会首先将其转化为数字（number）再判定大小。
console.log( '01' == 1 ); // true，字符串 '01' 会被转化为数字 1

//对于布尔类型值，true 会被转化为 1、false 转化为 0。
console.log( true == 1 ); // true
console.log( false == 0 ); // true


// 有时候，以下两种情况会同时发生：
    //- 若直接比较两个值，其结果是相等的。
    //- 若把两个值转为布尔值，它们可能得出完全相反的结果，即一个是 true，一个是 false。

let a = 0;
alert( Boolean(a) ); // false

let b = "0";
console.log( Boolean(b) ); // true

aleconsole.logrt(a == b); // true!
// 因为 JavaScript 会把待比较的值转化为数字后再做比较（因此 "0" 变成了 0）。
// 若只是将一个变量转化为 Boolean 值，则会使用其他的类型转换规则。



// 严格相等
// 普通的相等性检查 == 存在一个问题，它不能区分出 0 和 false：

console.log( 0 == false ); // true

//也同样无法区分空字符串和 false：
console.log( '' == false ); // true

// 这是因为在比较不同类型的值时，处于相等判断符号 == 两侧的值会先被转化为数字。
// 空字符串和 false 也是如此，转化后它们都为数字 0。


// 严格相等运算符 === 在进行比较时不会做任何的类型转换。
// 如果 a 和 b  a ==属于不同的数据类型，那么= b 不会做任何的类型转换而立刻返回 false。
console.log( 0 === false ); // false，因为被比较值的数据类型不同

// 同样的，与“不相等”符号 != 类似，“严格不相等”表示为 !==。






// 对 null 和 undefined 进行比较
// 当使用 null 或 undefined 与其他值进行比较时，其返回结果常常出乎你的意料。

//- 当使用严格相等 === 比较二者时
//- 它们不相等，因为它们属于不同的类型。
console.log( null === undefined ); // false

// 当使用非严格相等 == 比较二者时
// JavaScript 存在一个特殊的规则，会判定它们相等。
// 它们俩就像“一对恋人”，仅仅等于对方而不等于其他任何的值（只在非严格相等下成立）。

console.log( null == undefined ); // true


//当使用数学式或其他比较方法 < > <= >= 时：
    // -null/undefined 会被转化为数字：null 被转化为 0，undefined 被转化为 NaN。


// 奇怪的结果：null vs 0

//-通过比较 null 和 0 可得：
console.log( null > 0 );  // (1) false
console.log( null == 0 ); // (2) false
console.log( null >= 0 ); // (3) true
// 为什么会出现这种反常结果，这是因为相等性检查 == 和普通比较符 > < >= <= 的代码逻辑是相互独立的。
// 进行值的比较时，null 会被转化为数字，因此它被转化为了 0。
// 这就是为什么（3）中 null >= 0 返回值是 true，（1）中 null > 0 返回值是 false。

// 另一方面，undefined 和 null 在相等性检查 == 中不会进行任何的类型转换，它们有自己独立的比较规则，
// 所以除了它们之间互等外，不会等于任何其他的值。这就解释了为什么（2）中 null == 0 会返回 false。

// 特立独行的 undefined
    //- undefined 不应该被与其他值进行比较：
console.log( undefined > 0 ); // false (1)
console.log( undefined < 0 ); // false (2)
console.log( undefined == 0 ); // false (3)

// 为何它看起来如此厌恶 0？返回值都是 false！
//- 原因如下：

// (1) 和 (2) 都返回 false 是因为 undefined 在比较中被转换为了 NaN，而 NaN 是一个特殊的数值型值，
// 它与任何值进行比较都会返回 false。
// (3) 返回 false 是因为这是一个相等性检查，而 undefined 只与 null 相等，不会与其他值相等。


//- 总结
    //- 比较运算符始终返回布尔值。
    //- 字符串的比较，会按照“词典”顺序逐字符地比较大小。
    //- 当对不同类型的值进行比较时，它们会先被转化为数字（不包括严格相等检查）再进行比较。
    //- 在非严格相等 == 下，null 和 undefined 相等且各自不等于任何其他的值。
    //- 在使用 > 或 < 进行比较时，需要注意变量可能为 null/undefined 的情况。
    //- 比较好的方法是单独检查变量是否等于 null/undefined。


/**
    5 > 4 → true
    "apple" > "pineapple" → false
    "2" > "12" → true
    undefined == null → true
    undefined === null → false
    null == "\n0\n" → false
    null === +"\n0\n" → false
    
    结果的原因：
        数字间比较大小，显然得 true。
        按词典顺序比较，得 false。"a" 比 "p" 小。
        与第 2 题同理，首位字符 "2" 大于 "1"。
        null 只与 undefined 互等。
        严格相等模式下，类型不同得 false。
        与第 4 题同理，null 只与 undefined 相等。
        不同类型严格不相等。
*/





















