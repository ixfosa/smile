// 异步迭代和 generator
    // 异步迭代允许我们对按需通过异步请求而得到的数据进行迭代。
    // 例如，我们通过网络分段（chunk-by-chunk）下载数据时。异步生成器（generator）使这一步骤更加方便。

// 首先，让我们来看一个简单的示例以掌握语法，然后再看一个实际用例。



// 回顾可迭代对象

// 假设我们有一个对象，例如下面的 range：
let range = {
  from: 1,
  to: 5
};

// 想对它使用 for..of 循环，例如 for(value of range)，来获取从 1 到 5 的值。

// 换句话说，想向对象 range 添加 迭代能力。

// 这可以通过使用一个名为 Symbol.iterator 的特殊方法来实现：
    // 当循环开始时，该方法被 for..of 结构调用，并且它应该返回一个带有 next 方法的对象。
    // 对于每次迭代，都会为下一个值调用 next() 方法。\
    // next() 方法应该以 {done: true/false, value:<loop value>} 的格式返回一个值，其中 done:true 表示循环结束。

// 这是可迭代的 range 的一个实现
range = {
    from: 1,
    to: 5,

    [Symbol.iterator]() {
        return {
            current: this.from,
            last: this.to,
            next() {
                if (this.current <= this.last) {
                    return { done: false, value: this.current++ };
                } else {
                    return { done: true};
                }
            }
        }
    }
}

for(let value of range) {
    console.log(value); // 1，然后 2，然后 3，然后 4，然后 5
}






// -------------------------------------------------------------------------------------------






// 异步可迭代对象
    // 当值是以异步的形式出现时，例如在 setTimeout 或者另一种延迟之后，就需要异步迭代。
    // 最常见的场景是，对象需要发送一个网络请求以传递下一个值

// 要使对象异步迭代：
    //- 1. 使用 Symbol.asyncIterator 取代 Symbol.iterator。

    //- 2. next() 方法应该返回一个 promise（带有下一个值，并且状态为 fulfilled）。

        // 关键字 async 可以实现这一点，我们可以简单地使用 async next()。

    //- 3. 我们应该使用 for await (let item of iterable) 循环来迭代这样的对象。
        // 注意关键字 await。


// 创建一个可迭代的 range 对象，与前面的那个类似，不过现在它将异步地每秒返回一个值。
range = {
    from: 1,
    to: 5,

    [Symbol.asyncIterator]() {  // (1)
        return {
            current: this.from,
            last: this.to,
            async next() {     // (2)
                // 注意：我们可以在 async next 内部使用 "await"
                await new Promise(resolve => setTimeout(resolve, 1000)); // (3)
                if (this.current <= this.last) {
                    return { done: false, value: this.current++ };
                } else {
                    return { done: true};
                }
            }
        };
    }
};

/*
(async () => {
    for await (let value of range) {   // (4)
        console.log(value); // 1，然后 2，然后 3，然后 4，然后 5
    }
})();
*/

// 其结构与常规的 iterator 类似:
    //- 1. 为了使一个对象可以异步迭代，它必须具有方法 Symbol.asyncIterator (1)。

    //- 2. 这个方法必须返回一个带有 next() 方法的对象，next() 方法会返回一个 promise (2)。

    //- 3. 这个 next() 方法可以不是 async 的，它可以是一个返回值是一个 promise 的常规的方法，但是使用
    //     async 关键字可以允许我们在方法内部使用 await，所以会更加方便。这里只是用于延迟 1 秒的操作 (3)。

    //- 4. 使用 for await(let value of range) (4) 来进行迭代，也就是在 for 后面添加 await。
    //     它会调用一次 range[Symbol.asyncIterator]() 方法一次，然后调用它的 next() 方法获取值。

// 这是一个对比 Iterator 和异步 iterator 之间差异的表格：
/*
                                Iterator	        异步 iterator
    提供 iterator 的对象方法	  Symbol.iterator	  Symbol.asyncIterator
    next() 返回的值是	         任意值	              Promise
    要进行循环，使用	          for..of	          for await..of
*/






// -------------------------------------------------------------------------------------------






// Spread 语法 ... 无法异步工作
    // 需要常规的同步 iterator 的功能，无法与异步 iterator 一起使用。

// 例如，spread 语法无法工作：

// console.log( [...range] ); // Error, no Symbol.iterator

// 这很正常，因为它期望找到 Symbol.iterator，而不是 Symbol.asyncIterator。

// for..of 的情况和这个一样：没有 await 关键字时，则期望找到的是 Symbol.iterator。







// -------------------------------------------------------------------------------------------






// 回顾 generator
    // generator，它使我们能够写出更短的迭代代码。
    // 在大多数时候，当想要创建一个可迭代对象时，会使用 generator。

// 简单起见，这里省略了一些解释，即 generator 是“生成（yield）值的函数”。

// Generator 是标有 function*（注意星号）的函数，它使用 yield 来生成值，
// 并且我们可以使用 for..of 循环来遍历它们。


// 下面这例子生成了从 start 到 end 的一系列值：
function* generateSequence0(start, end) {
    for (let i = start; i <= end; i++) {
        yield i;
    }
}

for (let val of generateSequence0(1, 6)) {
    console.log("val: " + val);
}


// 要使一个对象可迭代，我们需要给它添加 Symbol.iterator。
/*
    let range = {
        from: 1,
        to: 5,
        [Symbol.iterator]() {
        return <带有 next 方法的对象，以使对象 range 可迭代>
        }
    }
*/


// 对于 Symbol.iterator 来说，一个通常的做法是返回一个 generator，这样可以使代码更短，如下所示：
range = {
    from: 1,
    to:5, 

    *[Symbol.iterator]() {
        for (let i = this.from; i <= this.to; i++) {
            yield i;
        }
    }
}
console.log([...range]) // [ 1, 2, 3, 4, 5 ]








// -------------------------------------------------------------------------------------------






// 异步 generator (finally)
    // 对于大多数的实际应用程序，当我们想创建一个异步生成一系列值的对象时，我们都可以使用异步 generator。

// 语法很简单：在 function* 前面加上 async。这即可使 generator 变为异步的。

// 然后使用 for await (...) 来遍历它，像这样：
async function* generateSequence(start, end) {
    for (let i = start; i <= end; i++) {
        await new Promise(resolve => setTimeout(resolve, 1000));
        yield i;
    }
}

(async () => {
    for await (let val of generateSequence(1, 10)) {
        console.log("await val: " + val);
    }
})();

// 因为此 generator 是异步的，所以我们可以在其内部使用 await，依赖于 promise，执行网络请求等任务。



// 引擎盖下的差异
    // 如果你还记得我们在前面章节中所讲的关于 generator 的细节知识，那你应该知道，
    // 从技术上讲，异步 generator 和常规的 generator 在内部是有区别的。

    // 对于异步 generator，generatr.next() 方法是异步的，它返回 promise。

    // generator 中，我们使用 result = generator.next() 来获得值。
    // 但在一个异步 generator 中，我们应该添加 await 关键字，像这样：
        //- result = await generator.next(); // result = {value: ..., done: true/false}
        //- 这就是为什么异步 generator 可以与 for await...of 一起工作。





// -------------------------------------------------------------------------------------------





// 异步的可迭代对象 range
    // 常规的 generator 可用作 Symbol.iterator 以使迭代代码更短。

    // 与之类似，异步 generator 可用作 Symbol.asyncIterator 来实现异步迭代。


// 例如，我们可以通过将同步的 Symbol.iterator 替换为异步的 Symbol.asyncIterator，
// 来使对象 range 异步地生成值，每秒生成一个：
range = {
    from: 1,
    to: 6,

    // 这一行等价于 [Symbol.asyncIterator]: async function*() {
    async *[Symbol.asyncIterator]() {
        for (let i = this.from; i <= this.to; i++) {
            await new Promise(resolve => setTimeout(resolve, 1000));
            yield i;
        }
    }
};

(async () => {
    for await (let val of range) {
        console.log("range val: " + val);
    }
})();


// 请注意：
    // 从技术上讲，我们可以把 Symbol.iterator 和 Symbol.asyncIterator 都添加到对象中，
    // 因此它既可以是同步的（for..of）也可以是异步的（for await..of）可迭代对象。







// -------------------------------------------------------------------------------------------






// 总结

// 常规的 iterator 和 generator 可以很好地处理那些不需要花费时间来生成的的数据。

// 当我们期望异步地，有延迟地获取数据时，可以使用它们的异步版本，并且使用 for await..of 替代 for..of。


/*
    // 异步 iterator 与常规 iterator 在语法上的区别：

                                Iterable	                    异步 Iterable
    提供 iterator 的对象方法	 Symbol.iterator	             Symbol.asyncIterator
    next() 返回的值是	         {value:…, done: true/false}	 resolve 成 {value:…, done: true/false} 的 
                                                                Promise
     

    // 异步 generator 与常规 generator 在语法上的区别：

                        Generator	                    异步 generator
    声明方式	         function*	                     async function*
    next() 返回的值是	 {value:…, done: true/false}	 resolve 成 {value:…, done: true/false} 的 
                                                        Promise
*/

// 在 Web 开发中，我们经常会遇到数据流，它们分段流动（flows chunk-by-chunk）。例如，下载或上传大文件。

// 可以使用异步 generator 来处理此类数据。值得注意的是，在一些环境，
// 例如浏览器环境下，还有另一个被称为 Streams 的 API，它提供了特殊的接口来处理此类数据流，
// 转换数据并将数据从一个数据流传递到另一个数据流（例如，从一个地方下载并立即发送到其他地方）。