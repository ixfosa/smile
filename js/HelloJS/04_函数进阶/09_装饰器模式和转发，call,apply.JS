
// 透明缓存
    // 假设我们有一个 CPU 重负载的函数 slow(x)，但它的结果是稳定的。换句话说，对于相同的 x，它总是返回相同的结果。

    // 如果经常调用该函数，我们可能希望将结果缓存（记住）下来，以避免在重新计算上花费额外的时间。

    // 但是我们不是将这个功能添加到 slow() 中，而是创建一个包装器（wrapper）函数，该函数增加了缓存功能。

// 下面是代码和解释：
function slow(x) {
    // 这里可能会有重负载的 CPU 密集型工作
    console.log(`Called with ${x}`);
    return x;
}

function cachingDecorator(func) {
    let cache = new Map();

    return function(x) {
        if (cache.has(x)) {         // 如果缓存中有对应的结果
            return cache.get(x);    // 从缓存中读取结果
        } 

        let res = func(x);   // 否则就调用 func

        cache.set(x, res);
        return res;
    };
}



slow = cachingDecorator(slow);
console.log( slow(1) );              // slow(1) 被缓存下来了
console.log( "Again: " + slow(1) );  
console.log( slow(2) );              // slow(2) 被缓存下来了
console.log( "Again: " + slow(2) );  // 和前面一行结果相同

//- Called with 1
//- 1
//- Again: 1
//- Called with 2
//- 2
//- Again: 2


// cachingDecorator 是一个 装饰器（decorator）：一个特殊的函数，它接受另一个函数并改变它的行为。

// 其思想是，我们可以为任何函数调用 cachingDecorator，它将返回缓存包装器。
// 因为我们有很多函数可以使用这样的特性，而我们需要做的就是将 cachingDecorator 应用于它们。

// 通过将缓存与主函数代码分开，我们还可以使主函数代码变得更简单。

// cachingDecorator(func) 的结果是一个“包装器”：function(x) 将 func(x) 的调用“包装”到缓存逻辑中：
// 从外部代码来看，包装的 slow 函数执行的仍然是与之前相同的操作。它只是在其行为上添加了缓存功能。

// 总而言之，使用分离的 cachingDecorator 而不是改变 slow 本身的代码有几个好处：
    //- cachingDecorator 是可重用的。我们可以将它应用于另一个函数。
    //- 缓存逻辑是独立的，它没有增加 slow 本身的复杂性（如果有的话）。
    //- 如果需要，我们可以组合多个装饰器（其他装饰器将遵循同样的逻辑）。





// -------------------------------------------------------------------------------------------




// 使用 “func.call” 设定上下文
    // 上面提到的缓存装饰器不适用于对象方法。

// 例如，在下面的代码中，worker.slow() 在装饰后停止工作：

let worker = {
    someMethod() {
        return 1;
    },

    slow(x) {
        // CPU 过载任务
        console.log("Called with " + x);
        return x * this.someMethod();   // (*)
      },
};

// 和之前例子中的代码相同
function cachingDecorator2(func) {
    let cache = new Map();
    return function(x) {
        if (cache.has(x)) {
            return cache.get(x);
        }
        let result = func(x); // (**)
        cache.set(x, result);
        return result;
    };
}

console.log( worker.slow(1) ); // 原始方法有效

// worker.slow = cachingDecorator2(worker.slow); // 现在对其进行缓存
// console.log( worker.slow(2) );  // TypeError: this.someMethod is not a function

// 原因是包装器将原始函数调用为 (**) 行中的 func(x)。并且，当这样调用时，函数将得到 this = undefined。

// 如果尝试运行下面这段代码，我们会观察到类似的问题：
//- let func = worker.slow;
//- func(2);  // TypeError: this.someMethod is not a function


// 因此，包装器将调用传递给原始方法，但没有上下文 this。因此，发生了错误。


// 有一个特殊的内置函数方法 func.call(context, …args)，它允许调用一个显式设置 this 的函数。
// 语法如下：
    //- func.call(context, arg1, arg2, ...)

// 它运行 func，提供的第一个参数作为 this，后面的作为参数（arguments）。

// 简单地说，这两个调用几乎相同：
    //- func(1, 2, 3);
    //- func.call(obj, 1, 2, 3)
// 它们调用的都是 func，参数是 1、2 和 3。唯一的区别是 func.call 还会将 this 设置为 obj。


// 例如，在下面的代码中，我们在不同对象的上下文中调用 sayHi：
    // sayHi.call(user) 运行 sayHi 并提供了 this=user，然后下一行设置 this=admin：

function sayHi() {
    console.log(this.name);
}

let student1 = {name: "Xia Fosong"};
let student2 = {name: "Zi Qing"};

sayHi.call(student1); // Xia Fosong
sayHi.call(student2); //Zi Qing


// 用带有给定上下文和 phrase 的 call 调用 say：
function say(phrase) {
    console.log(`${phrase} ${this.name}`)
}

let user = {name: "Zi Qing"};

// user 成为 this，"Hello" 成为第一个参数
say.call(user, "hello"); // hello Zi Qing


// 可以在包装器中使用 call 将上下文传递给原始函数：
let worker3 = {
    someMethod() {
        return 1;
    },

    slow(x) {
        // CPU 过载任务
        console.log("Called with " + x);
        return x * this.someMethod();   // (*)
      },
};


function cachingDecorator3(func) {

    let cache = new Map();

    return function(x) {

        if (cache.has(x)) {
            return cache.get(x);
        }

        // let result = func(x); // (**)
        let result = func.call(this, x); // (**)

        cache.set(x, result);

        return result;
    };
}

worker3.slow = cachingDecorator(worker3.slow); // 现在对其进行缓存

console.log( worker.slow(2) ); // 工作正常
console.log( worker.slow(2) ); // 工作正常，没有调用原始函数（使用的缓存）
// Called with 2
// 2
// Called with 2
// 2

//- slow = cachingDecorator(worker3.slow);
//- console.log( slow(2) ); // TypeError: this.someMethod is not a function

// this 是如何被传递的：
    //- 在经过装饰之后，worker.slow 现在是包装器 function (x) { ... }。
    //- 因此，当 worker.slow(2) 执行时，包装器将 2 作为参数，并且 this=worker（它是点符号 . 之前的对象）。
    //- 在包装器内部，假设结果尚未缓存，func.call(this, x) 将当前的 this（=worker）和当前的参数（=2）传递给原始方法。





// -------------------------------------------------------------------------------------------



    
// 传递多个参数
    //- 现在让我们把 cachingDecorator 写得更加通用。到现在为止，它只能用于单参数函数。

    // 之前，对于单个参数 x，我们可以只使用 cache.set(x, result) 来保存结果，
    // 并使用 cache.get(x) 来检索并获取结果。但是现在，我们需要记住 参数组合 (min,max) 的结果。
    // 原生的 Map 仅将单个值作为键（key）。

    // 有许多解决方案可以实现：

        //- 1. 实现一个新的（或使用第三方的）类似 map 的更通用并且允许多个键的数据结构。

        //- 2. 使用嵌套 map：cache.set(min) 将是一个存储（键值）对 (max, result) 的 Map。
        //     所以我们可以使用 cache.get(min).get(max) 来获取 result。
        
        //- 3. 将两个值合并为一个。为了灵活性，我们可以允许为装饰器提供一个“哈希函数”，
        //     该函数知道如何将多个值合并为一个值。

    // 当然，我们需要传入的不仅是 x，还需要传入 func.call 的所有参数。
    // 在 function() 中我们可以得到一个包含所有参数的伪数组（pseudo-array）arguments，
    // 那么 func.call(this, x) 应该被替换为 func.call(this, ...arguments)。

let worker4 = {
    slow(...rest) {
        let sum = 0;
        
        for (let val of rest) {
            sum += val;
        }
       
        return sum;
    }
}

function hash(args) {
    let res = "";
    for (let arg of args) {
        res += "," + arg;
    }
    return res;
}

function cachingDecorator4(func, hash) {
    let cache = new Map();

    return function() {

        let key = hash(arguments);

        if (cache.has(key)) {
            return cache.get(key);
        }

        let res = func.call(this, ...arguments);

        cache.set(key, res);

        return res;
    };
}

// 应该记住相同参数的调用
worker4.slow = cachingDecorator4(worker4.slow, hash);

console.log( worker4.slow(3, 5) ); // works
console.log( "Again " + worker4.slow(3, 5) ); // same (cached)

console.log( worker4.slow(1, 3, 5, 7, 9) ); // works
console.log( "Again " + worker4.slow( 1, 3, 5, 7, 9 )); // same (cached)






// -------------------------------------------------------------------------------------------



// func.apply
    //- 可以使用 func.apply(this, arguments) 代替 func.call(this, ...arguments)。

    // 内建方法 func.apply 的语法是：
        //- func.apply(context, args)
        //- 它运行 func 设置 this=context，并使用类数组对象 args 作为参数列表（arguments）。

    // call 和 apply 之间唯一的语法区别是，call 期望一个参数列表，而 apply 期望一个包含这些参数的类数组对象。

    // 因此，这两个调用几乎是等效的：
        //- func.call(context, ...args); // 使用 spread 语法将数组作为列表传递
        //- func.apply(context, args);   // 与使用 call 相同

    // 这里只有很小的区别：
        //- Spread 语法 ... 允许将 可迭代对象 args 作为列表传递给 call。
        //- apply 仅接受 类数组对象 args。

// 因此，当我们期望可迭代对象时，使用 call，当我们期望类数组对象时，使用 apply。

// 对于即可迭代又是类数组的对象，例如一个真正的数组，我们使用 call 或 apply 均可，
// 但是 apply 可能会更快，因为大多数 JavaScript 引擎在内部对其进行了优化。

// 将所有参数连同上下文一起传递给另一个函数被称为“呼叫转移（call forwarding）”。

// 这是它的最简形式：
let wrapper = function() {
  return func.apply(this, arguments);
};
// 当外部代码调用这种包装器 wrapper 时，它与原始函数 func 的调用是无法区分的。





// -------------------------------------------------------------------------------------------



// 借用一种方法

// 对哈希函数再做一个较小的改进：
/*
function hash(args) {
    let res = "";
    for (let arg of args) {
        res += "," + arg;
    }
    return res;
}


使用 arr.join 方法：
function hash(args) {
  return args.join();  // // Error: arguments.join is not a function
}

这不行。因为我们正在调用 hash(arguments)，arguments 对象既是可迭代对象又是类数组对象，但它并不是真正的数组。

*/

// 有一种简单的方法可以使用数组的 join 方法：
// 这个技巧被称为 方法借用（method borrowing）。
function hash2() {
    console.log( [].join.call(arguments) ); // 1,2
}
  
hash2(1, 2);

// 我们从常规数组 [].join 中获取（借用）join 方法，并使用 [].join.call 在 arguments 的上下文中运行它。

// 原生方法 arr.join(glue) 的内部算法非常简单。
// 从规范中几乎“按原样”解释如下：
    //- 1. 让 glue 成为第一个参数，如果没有参数，则使用逗号 ","。
    //- 2. 让 result 为空字符串。
    //- 3. 将 this[0] 附加到 result。
    //- 4. 附加 glue 和 this[1]。
    //- 5. 附加 glue 和 this[2]。
    //- 6. 以此类推，直到 this.length 项目被粘在一起。
    //- 7. 返回 result。

// 因此，从技术上讲，它需要 this 并将 this[0]，this[1] ……等 join 在一起。
// 它的编写方式是故意允许任何类数组的 this 的（不是巧合，很多方法都遵循这种做法）。
// 这就是为什么它也可以和 this=arguments 一起使用。




// -------------------------------------------------------------------------------------------



// 装饰器和函数属性
    // 通常，用装饰的函数替换一个函数或一个方法是安全的，除了一件小东西。
    // 如果原始函数有属性，例如 func.calledCount 或其他，则装饰后的函数将不再提供这些属性。
    // 因为这是装饰器。因此，如果有人使用它们，那么就需要小心。

    // 例如，在上面的示例中，如果 slow 函数具有任何属性，而 cachingDecorator(slow) 则是一个没有这些属性的包装器。

    // 一些包装器可能会提供自己的属性。例如，装饰器会计算一个函数被调用了多少次以及花费了多少时间，
    // 并通过包装器属性公开（expose）这些信息。


// -------------------------------------------------------------------------------------------





// 总结
    // 装饰器 是一个围绕改变函数行为的包装器。主要工作仍由该函数来完成。

    // 装饰器可以被看作是可以添加到函数的 “(特征;特色;容貌)features” 或 “(外观; 外表)aspects”。
    // 可以添加一个或添加多个。而这一切都无需更改其代码！

// 为了实现 cachingDecorator，我们研究了以下方法：
    //- func.call(context, arg1, arg2…) —— 用给定的上下文和参数调用 func。
    //- func.apply(context, args) —— 调用 func 将 context 作为 this 和类数组的 args 传递给参数列表。

// 通用的 呼叫转移（call forwarding） 通常是使用 apply 完成的：
    //-   let wrapper = function() {
    //-      return original.apply(this, arguments);
    //-   };

// 我们也可以看到一个 方法借用（method borrowing） 的例子，就是我们从一个对象中获取一个方法，
// 并在另一个对象的上下文中“调用”它。采用数组方法并将它们应用于参数 arguments 是很常见的。
// 另一种方法是使用 Rest 参数对象，该对象是一个真正的数组。



