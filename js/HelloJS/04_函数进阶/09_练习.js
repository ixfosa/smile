// 间谍装饰器

// 创建一个装饰器 spy(func)，它应该返回一个包装器，该包装器将所有对函数的调用保存在其 calls 属性中。

// 每个调用都保存为一个参数数组。

// 例如：
function work(a, b) {
  console.log( a + b ); // work 是一个任意的函数或方法
}

work = spy(work);

work(1, 2); // 3
work(4, 5); // 9

for (let args of work.calls) {
    console.log( 'call:' + args.join() ); // "call:1,2", "call:4,5"
}

// P.S. 该装饰器有时对于单元测试很有用。它的高级形式是 Sinon.JS 库中的 sinon.spy。
function spy(func) {
    function wrapper(...args) {
        wrapper.calls.push(args);
        // console.log(this);
        // return func.apply(this, args);
        return func(...args);
    }
    wrapper.calls = [];
    return wrapper;
}



// -------------------------------------------------------------------------------------------


// 延时装饰器

// 创建一个装饰器 delay(f, ms)，该装饰器将 f 的每次调用延时 ms 毫秒。

// 例如：
function f(x) {
  console.log(x);
}

// create wrappers
let f1000 = delay(f, 1000 * 3);
let f1500 = delay(f, 1500);

f1000("test"); // 在 1000ms 后显示 "test"
//f1500("test"); // 在 1500ms 后显示 "test"

// 换句话说，delay(f, ms) 返回的是延迟 ms 后的 f 的变体。

// 在上面的代码中，f 是单个参数的函数，但是你的解决方案应该传递所有参数和上下文 this。
function delay(f, ms) {
    return function() {
        setTimeout(f, ms, ...arguments);
        // setTimeout(() => f.apply(this, arguments), ms);
    }
}








