
// 循环 是一种重复运行同一代码的方法。

// “while” 循环
    //- while 循环的语法如下：
/*
    // 当 condition 为真时，执行循环体的 code。
    while (condition) {
        // 代码
        // 所谓的“循环体”
    }
*/

// 例如，以下将循环输出当 i < 3 时的 i 值：
let i = 0;
while (i < 3) { // 依次显示 0、1 和 2
    console.log( i );
  i++;
}
// 循环体的单次执行叫作 一次迭代。上面示例中的循环进行了三次迭代。

// while (i != 0) 可简写为 while (i)：
let j = 3;
while (j) { // 当 i 变成 0 时，条件为假，循环终止
    console.log( j );
  j--;
}

// 单行循环体不需要大括号
    //- 如果循环体只有一条语句，则可以省略大括号 {…}：
let k = 3;
while (k) console.log(k--);


// “do…while” 循环
    //- 使用 do..while 语法可以将条件检查移至循环体 下面：
/*
    // 循环首先执行循环体，然后检查条件，当条件为真时，重复执行循环体。
    do {
        // 循环体
    } while (condition);
*/
// ，循环体 至少执行一次。通常我们更倾向于使用另一个形式：while(…)
let p = 0;
do {
    console.log( p );
  p++;
} while (p < 3);


// “for” 循环
    //- for 循环更加复杂，但它是最常使用的循环形式。

/*
    // for 循环看起来就像这样：
    for (begin; condition; step) {
        // ……循环体……
    }
*/


//- 下述循环从 i 等于 0 到 3（但不包括 3）运行 alert(i)：
for (let i = 0; i < 3; i++) { // 结果为 0、1、2
    console.log(i);
}

/**
    逐个部分分析 for 循环：

    语句段		
        begin	i = 0	进入循环时执行一次。
        condition	i < 3	在每次循环迭代之前检查，如果为 false，停止循环。
        body（循环体）	alert(i)	条件为真时，重复运行。
        step	i++	在每次循环体迭代后执行。

    一般循环算法的工作原理如下：
        开始运行
        → (如果 condition 成立 → 运行 body 然后运行 step)
        → (如果 condition 成立 → 运行 body 然后运行 step)
        → (如果 condition 成立 → 运行 body 然后运行 step)
        → ...

    所以，begin 执行一次，然后进行迭代：每次检查 condition 后，执行 body 和 step。
 */

 // 内联变量声明
    //- 这里“计数”变量 i 是在循环中声明的。这叫做“内联”变量声明。这样的变量只在循环中可见。

for (let n = 0; n < 3; n++) {
    console.log(n); // 0, 1, 2
}
// console.log(n); // 错误，没有这个变量。 ReferenceError: n is not defined

// 除了定义一个变量，我们也可以使用现有的变量：
let m = 0;
for (m = 0; m < 3; m++) { // 使用现有的变量
    console.log(m); // 0, 1, 2
}
console.log(m); //3，可见，因为是在循环之外声明的


// 省略语句段
    //- for 循环的任何语句段都可以被省略。

//- 例如，如果我们在循环开始时不需要做任何事，我们就可以省略 begin 语句段

let x = 0; // 我们已经声明了 x 并对它进行了赋值
for (; x < 3; x++) { // 不再需要 "begin" 语句段
console.log( x ); // 0, 1, 2
}

// 我们也可以移除 step 语句段：
let y = 0;

for (; y < 3;) {
    console.log( y++ );
}
// 该循环与 while (i < 3) 等价。


// 实际上我们可以删除所有内容，从而创建一个无限循环：
for (;;) {
  // 无限循环
}
//- 请注意 for 的两个 ; 必须存在，否则会出现语法错误。


// 跳出循环
    //- 通常条件为假时，循环会终止。
    //- 随时都可以使用 break 指令强制退出。

// 根据需要，"无限循环 + break" 的组合非常适用于不必在循环开始/结束时检查条件，
// 但需要在中间甚至是主体的多个位置进行条件检查的情况。

// 例如，下面这个循环要求用户输入一系列数字，在输入的内容不是数字时“终止”循环。
let sum = 0;
while (true) {
  let value = +prompt("Enter a number", '');
  if (!value) break; // (*)
  sum += value;

}
console.log( 'Sum: ' + sum );


// 继续下一次迭代
    //- continue 指令是 break 的“轻量版”。它不会停掉整个循环。
    //- 而是停止当前这一次迭代，并强制启动新一轮循环（如果条件允许的话）。


// 下面这个循环使用 continue 来只输出奇数：
// 对于偶数的 i 值，continue 指令会停止本次循环的继续执行，
// 将控制权传递给下一次 for 循环的迭代（使用下一个数字）。因此 console.log 仅被奇数值调用。
for (let i = 0; i < 10; i++) {

  //如果为真，跳过循环体的剩余部分。
  if (i % 2 == 0) continue;

  console.log(i); // 1，然后 3，5，7，9
}


// 禁止 break/continue 在 ‘?’ 的右边
    //- 请注意非表达式的语法结构不能与三元运算符 ? 一起使用。
    //- 特别是 break/continue 这样的指令是不允许这样使用的。
if (i > 5) {
  console.log(i);
} else {
  continue;
}

// ……用问号重写：
// 代码会停止运行，并显示有语法错误。
// 这是不（建议）使用问号 ? 运算符替代 if 语句的另一个原因。
    //- (i > 5) ? console.log(i) : continue; // continue 不允许在这个位置




// break/continue 标签
    //- 有时候我们需要从一次从多层嵌套的循环中跳出来。

// 例如，下述代码中我们的循环使用了 i 和 j，从 (0,0) 到 (3,3) 提示坐标 (i, j)：
    //- 我们需要提供一种方法，以在用户取消输入时来停止这个过程。 
    //- 在 input 之后的普通 break 只会打破内部循环。这还不够 —— 标签可以实现这一功能！
for (let i = 0; i < 3; i++) {

    for (let j = 0; j < 3; j++) {

    let input = prompt(`Value at coords (${i},${j})`, '');

        // 如果我想从这里退出并直接执行 console.log('Done!')
    }
}
console.log('Done!');

    
/**
    标签 是在循环之前带有冒号的标识符：
        labelName: for (...) {
        ...
        }
    break <labelName> 语句跳出循环至标签处：
 */
    

// break outer 向上寻找名为 outer 的标签并跳出当前循环。
outer: for (let i = 0; i < 3; i++) {

    for (let j = 0; j < 3; j++) {

        let input = prompt(`Value at coords (${i},${j})`, '');

        // 如果是空字符串或被取消，则中断并跳出这两个循环。
        if (!input) break outer; // (*)

        // 用得到的值做些事……
    }
}
console.log('Done!');
  
    
// 控制权直接从 (*) 转至 alert('Done!')。
    
// 可以将标签移至单独一行：
outer:
for (let i = 0; i < 3; i++) { 
    //... 
}


/**
    标签并不允许“跳到”所有位置

    例如，这样做是不可能的：

        break label;  // 无法跳转到这个标签
        label: for (...)

    只有在循环内部才能调用 break/continue，并且标签必须位于指令上方的某个位置。


    三种循环：
        while —— 每次迭代之前都要检查条件。
        do..while —— 每次迭代后都要检查条件。
        for (;;) —— 每次迭代之前都要检查条件，可以使用其他设置。
        通常使用 while(true) 来构造“无限”循环。这样的循环和其他循环一样，都可以通过 break 指令来终止。

        如果不想在当前迭代中做任何事，并且想要转移至下一次迭代，那么可以使用 continue 指令。

        break/continue 支持循环前的标签。标签是 break/continue 跳出嵌套循环以转到外部的唯一方法。
 */