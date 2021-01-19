//- 空值合并运算符 '??'
    //- 在本文中，将值既不是 null 也不是 undefined 的表达式称为“已定义的（defined）”。

// 空值合并运算符（nullish coalescing operator）的写法为两个问号 ??。

// a ?? b 的结果是：
    //- 如果 a 是已定义的，则结果为 a，
    //- 如果 a 不是已定义的，则结果为 b。

// 换句话说，如果第一个参数不是 null/undefined，则 ?? 返回第一个参数。否则，返回第二个参数。
// 空值合并运算符并不是什么全新的东西。它只是一种获得两者中的第一个“已定义的”值的不错的语法。
// 我们可以使用我们已知的运算符重写 result = a ?? b，像这样：
    //-  result = (a !== null && a !== undefined) ? a : b;


// 通常 ?? 的使用场景是，为可能是未定义的变量提供一个默认值。
// 例如，在这里，如果 user 是未定义的，我们则显示 Anonymous：
let user;
console.log(user ?? "long"); // long

// 如果 user 的值为除 null/undefined 外的任意值，那么我们看到的将是它：
user = "zhong";
console.log(user ?? "long"); // zhong



// 使用 ?? 序列从一系列的值中选择出第一个非 null/undefined 的值。
let firstName = null;
let lastName = null;
let nickName = "梓晴";
// 显示第一个已定义的值：
console.log(firstName ?? lastName ?? nickName ?? "如也"); // 梓晴


// 与 || 比较
// 或运算符 || 可以以与 ?? 运算符相同的方式使用。

// 例如，在上面的代码中，我们可以用 || 替换掉 ??，也可以获得相同的结果：
console.log(firstName || lastName || nickName || "如也"); // 梓晴


//- 它们之间重要的区别是：
    //- || 返回第一个 真 值。
    //- ?? 返回第一个 已定义的 值。

// 换句话说，|| 无法区分 false、0、空字符串 "" 和 null/undefined。
// 它们都一样 —— 假值（falsy values）。如果其中任何一个是 || 的第一个参数，
// 那么我们将得到第二个参数作为结果。

// 不过在实际中，我们可能只想在变量的值为 null/undefined 时使用默认值。
// 也就是说，当该值确实未知或未被设置时。

// 如果高度 0 为有效值，则不应将其替换为默认值，所以 ?? 能够得出正确的结果。
let height = 0;
console.log(height || 100); // 100
console.log(height ?? 100); // 0



// 优先级
let h = null;
let w = null;
let area;
// 重要：使用括号
area = (h ?? 100) * (w ?? 50);
console.log(area); // 5000

// 否则，如果我们省略了括号，则由于 * 的优先级比 ?? 高，它会先执行，进而导致错误的结果。
// 没有括号
// area = h1 ?? 100 * w1 ?? 50; // ReferenceError: h1 is not defined
// console.log(area);

// ……与下面这行代码的计算方式相同（应该不是我们所期望的）：
// area = h2 ?? (100 * w2) ?? 50; // ReferenceError: h2 is not defined
// console.log(area);


// ?? 与 && 或 || 一起使用
    //- 出于安全原因，JavaScript 禁止将 ?? 运算符与 && 和 || 运算符一起使用，除非使用括号明确指定了优先级。

//下面的代码会触发一个语法错误：
let x = 1 && 2 ?? 3; // Syntax error
// 这个限制无疑是值得商榷的，但它被添加到语言规范中是为了避免人们从 || 切换到 ?? 时的编程错误。

// 可以明确地使用括号来解决这个问题：
let y = (1 && 2) ?? 3; // 正常工作了

console.log(y); // 2


// 总结
    //- 空值合并运算符 ?? 提供了一种从列表中选择第一个“已定义的”值的简便方式。

// 它被用于为变量分配默认值：
// 当 height 的值为 null 或 undefined 时，将 height 的值设置为 100
    //- height = height ?? 100;

// ?? 运算符的优先级非常低，仅略高于 ? 和 =，因此在表达式中使用它时请考虑添加括号。
// 如果没有明确添加括号，不能将其与 || 或 && 一起使用。


