// Promise

//- 1. “生产者代码（producing code）”会做一些事儿，并且会需要一些时间。例如，通过网络加载数据的代码。
//- 2. “消费者代码（consuming code）”想要在“生产者代码”完成工作的第一时间就能获得其工作成果。许多函数可能都需要这个结果。
//- 3. Promise 是将“生产者代码”和“消费者代码”连接在一起的一个特殊的 JavaScript 对象。

// 这种类比并不十分准确，因为 JavaScipt 的 promise 比简单的订阅列表更加复杂：它们还拥有其他的功能和局限性。


// Promise 对象的构造器（constructor）语法如下：
//-   let promise = new Promise(function(resolve, reject) {
//-       // executor（生产者代码）
//-   });


// 传递给 new Promise 的函数被称为 executor。当 new Promise 被创建，executor 会自动运行。
// 它包含最终应产出结果的生产者代码

// 它的参数 resolve 和 reject 是由 JavaScript 自身提供的回调。我们的代码仅在 executor 的内部。


// 当 executor 获得了结果，无论是早还是晚都没关系，它应该调用以下回调之一：
    //- resolve(value) — 如果任务成功完成并带有结果 value。
    //- reject(error) — 如果出现了 error，error 即为 error 对象。


// 所以总结一下就是：executor 会自动运行并尝试执行一项工作。
// 尝试结束后，如果成功则调用 resolve，如果出现 error 则调用 reject。


// 由 new Promise 构造器返回的 promise 对象具有以下内部属性：
    //- state  —  最初是 "pending"，然后在 resolve 被调用时变为 "fulfilled"，
    //  或者在 reject 被调用时变为 "rejected"。

    //- result — 最初是 undefined，然后在 resolve(value) 被调用时变为 value，
    //  或者在 reject(error) 被调用时变为 error。

//

/*

                               resolve(value)   > state: "fulfilled"
                            |> -------------    > result: value
    new Promise(executor):  |
    > state: "pending"      |
    > result: undefined     |
                            |> -------------    > state: "rejected" 
                               reject(error)    > result: error
*/


// 下面是一个 promise 构造器和一个简单的 executor 函数，
// 该 executor 函数具有包含时间（即 setTimeout）的“生产者代码”：
let promise1 = new Promise(function(resolve, reject) {
    // 当 promise 被构造完成时，自动执行此函数
  
    // 1 秒后发出工作已经被完成的信号，并带有结果 "done"
    setTimeout(() => resolve("done"), 1000);
});

// 通过运行上面的代码，我们可以看到两件事儿：
    //- 1. executor 被自动且立即调用（通过 new Promise）。

    //- 2. executor 接受两个参数：resolve 和 reject。这些函数由 JavaScipt 引擎预先定义，
    //     因此我们不需要创建它们。我们只需要在准备好（译注：指的是 executor 准备好）时调用其中之一即可。

// 经过 1 秒的“处理”后，executor 调用 resolve("done") 来产生结果。这将改变 promise 对象的状态：
/*

    new Promise(executor):    resolve("done")
    state: "pending"       -------------------> state: "fulfilled" 
    result: undefined                           result: "done"
*/




// 一个 executor 以 error 拒绝 promise 的示例：
/*
    let promise2 = new Promise(function(resolve, reject) {
        // 1 秒后发出工作已经被完成的信号，并带有 error
        setTimeout(() => reject(new Error("Whoops!")), 1000); 
    });
*/  

// (node:26536) UnhandledPromiseRejectionWarning: Error: Whoops!

// 对 reject(...) 的调用将 promise 对象的状态移至 "rejected"：
/*

    new Promise(executor):    reject(error)
    state: "pending"       -------------------> state: "rejected" 
    result: undefined                           result: error
*/

// 总而言之，executor 应该执行一项工作（通常是需要花费一些时间的事儿），
// 然后调用 resolve 或 reject 来改变对应的 promise 对象的状态。

// 与最初的 “pending” promise 相反，一个 resolved 或 rejected 的 promise 都会被称为 “settled”。






// -------------------------------------------------------------------------------------------





// 这儿只能有一个结果或一个 error
    // executor 只能调用一个 resolve 或一个 reject。任何状态的更改都是最终的。

// 所有其他的再对 resolve 和 reject 的调用都会被忽略：
let promise3 = new Promise(function(resolve, reject) {
    resolve("done");

    reject(new Error("…")); // 被忽略
    setTimeout(() => resolve("…")); // 被忽略
});

// 这儿的宗旨是，一个被 executor 完成的工作只能有一个结果或一个 error。
// 并且，resolve/reject 只需要一个参数（或不包含任何参数），并且将忽略额外的参数。



// 以 Error 对象 reject
    // 如果什么东西出了问题， executor 应该调用 reject。这可以使用任何类型的参数来完成（就像 resolve 一样）。
    // 但是建议使用 Error 对象（或继承自 Error 的对象）。




// Resolve/reject 可以立即进行
    // 实际上，executor 通常是异步执行某些操作，并在一段时间后调用 resolve/reject，但这不是必须的。
    // 还可以立即调用 resolve 或 reject，就像这样：

let promise4 = new Promise(function(resolve, reject) {
    // 不花时间去做这项工作
    resolve(123); // 立即给出结果：123
});


// 例如，当我们开始做一个任务时，但随后看到一切都已经完成并已被缓存时，可能就会发生这种情况。


// state 和 result 都是内部的
    // Promise 对象的 state 和 result 属性都是内部的。无法直接访问它们。
    // 但可以对它们使用 .then/.catch/.finally 方法。





// -------------------------------------------------------------------------------------------





// 消费者：then，catch，finally
    // Promise 对象充当的是 executor（“生产者代码”）和消费函数（“粉丝”）之间的连接，
    // 后者将接收结果或 error。可以通过使用 .then、.catch 和 .finally 方法为消费函数进行注册。


// then
    // 最重要最基础的一个就是 .then。

// 语法如下：
    //      promise.then(
    //          function(result) { /* handle a successful result */ },
    //          function(error) { /* handle an error */ }
    //      );

        // .then 的第一个参数是一个函数，该函数将在 promise resolved 后运行并接收结果。
        // .then 的第二个参数也是一个函数，该函数将在 promise rejected 后运行并接收 error。


// 例如，以下是对成功 resolved 的 promise 做出的反应：
let promise5 = new Promise(function(resolve, reject) {
    setTimeout(() => resolve("done"), 1000);
});

// resolve 运行 .then 中的第一个函数
promise5.then(
    result => console.log(result),  // 1 秒后显示 "done!"
    error => console.log(error)     // 不运行
)



// 在 reject 的情况下，运行第二个：
let promise6 = new Promise(function(resolve, reject) {
    setTimeout(() => reject(new Error("Error")), 1000);
});

// reject 运行 .then 中的第二个函数
promise6.then(
    result => console.log(result),  // 不运行
    error => console.log(error)   // 1 秒后显示 Error: Error
)








// -------------------------------------------------------------------------------------------


// 如果只对成功完成的情况感兴趣，那么我们可以只为 .then 提供一个函数参数：
promise5.then(console.log); // done





// catch
    //如果只对 error 感兴趣，那么我们可以使用 null 作为第一个参数：
    // .then(null, errorHandlingFunction)。或者我们也可以使用 .catch(errorHandlingFunction)，其实是一样的：


// .catch(f) 与 promise.then(null, f) 一样
promise6.catch(
    err => console.log(err) // 1 秒后显示 Error: Error
); 

// .catch(f) 调用是 .then(null, f) 的完全的模拟，它只是一个简写形式。







// -------------------------------------------------------------------------------------------






// finally
    // 就像常规 try {...} catch {...} 中的 finally 子句一样，promise 中也有 finally。

    // .finally(f) 调用与 .then(f, f) 类似，在某种意义上，f 总是在 promise 被 settled 时运行：
    // 即 promise 被 resolve 或 reject。

    // finally 是执行清理（cleanup）的很好的处理程序（handler），
    // 例如无论结果如何，都停止使用不再需要的加载指示符（indicator）。

// 像这样：
/*
    new Promise((resolve, reject) => {
        // 做一些需要时间的事儿，然后调用 resolve/reject
    })
    .finally(() => stop loading indicator)  // 在 promise 为 settled 时运行，无论成功与否
    .then(result => show result, err => show error)
   
    // 所以，加载指示器（loading indicator）始终会在我们处理结果/错误之前停止
    
*/


// 也就是说，finally(f) 其实并不是 then(f,f) 的别名。它们之间有一些细微的区别：
    // finally 处理程序（handler）没有参数。在 finally 中，我们不知道 promise 是否成功。
    // finally 处理程序将结果和 error 传递给下一个处理程序。


// 例如，结果被从 finally 传递给了 then：
new Promise((resolve, reject) => {
    setTimeout( () => resolve("Zi Qing"), 1000)
}) 
.finally( () =>console.log("finally") ) // finally
.then( res => console.log(res) ); //  .then 对结果进行处理   Zi Qing



// promise 中有一个 error，这个 error 被从 finally 传递给了 catch：
new Promise((resolve, reject) => {
    throw new Error("JavaScript")
})
.finally( () =>console.log("finally") ) // finally
.catch( err => console.log(err) ); //  .then 对结果进行处理   Error: JavaScript




// 可以对 settled 的 promise 附加处理程序
    // 如果 promise 为 pending 状态，.then/catch/finally 处理程序（handler）将等待它。
    // 否则，如果 promise 已经是 settled 状态，它们就会运行

// 下面这 promise 在被创建后立即变为 resolved 状态
let promise7 = new Promise(resolve => resolve("done!"));

promise7.then(console.log); // done!（现在显示）





// -------------------------------------------------------------------------------------------





// -------------------------------------------------------------------------------------------





// -------------------------------------------------------------------------------------------




// -------------------------------------------------------------------------------------------





// -------------------------------------------------------------------------------------------





// -------------------------------------------------------------------------------------------





// -------------------------------------------------------------------------------------------
