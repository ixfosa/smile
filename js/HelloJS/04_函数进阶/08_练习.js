
// 每秒输出一次

// 编写一个函数 printNumbers(from, to)，使其每秒输出一个数字，数字从 from 开始，到 to 结束。

// 使用以下两种方法来实现。
    //- 使用 setInterval。
    //- 使用嵌套的 setTimeout。
/*
// 使用嵌套的 setTimeout。
let timer1 = setTimeout(function printNumbers(from, to) {
    console.log(from++);
    timer1 = setTimeout(printNumbers, 1000, from, to);
    if (from == (to + 1)) {
        clearTimeout(timer1);
    }
}, 1000, 1, 6);
*/

function printNumbers1(from, to) {
    let timer1 = setTimeout(function run() {
        console.log(from++);
        timer1 = setTimeout(run, 1000);
        if (from == (to + 1)) {
            clearTimeout(timer1);
        }
    }, 1000);
} 
// printNumbers1(6, 10)


// 使用 setInterval。
function printNumbers2(from, to) {
    let timer = setInterval(function() {
        console.log(from++);
        if (from == to + 1) {
            clearInterval(timer);
        }
    }, 1000);
}   

// printNumbers2(6, 10)





// setTimeout 会显示什么？

// 下面代码中使用 setTimeout 调度了一个调用，然后需要运行一个计算量很大的 for 循环，
// 这段运算耗时超过 100 毫秒。

// 调度的函数会在何时运行？
    //- 循环执行完成后。
    //- 循环执行前。
    //- 循环刚开始时。

// console.log 会显示什么？

let i = 0;

setTimeout(() => console.log(i), 100); // ? 100000000

// 假设这段代码的运行时间 >100ms
for(let j = 0; j < 100000000; j++) {
  i++;
}

// 任何 setTimeout 都只会在当前代码执行完毕之后才会执行。
// 所以 i 的取值为：100000000。