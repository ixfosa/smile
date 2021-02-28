function loadScript1(src, callback) {
    let script = document.createElement("script");
    script.src = src;

    script.onload = () => callback(script);
    script.onerror = () => callback(new Error(`Script load error for ${src}`));

    document.head.append(script);
}

// 用 promise 重写它。

// 新函数 loadScript 将不需要回调。
// 取而代之的是，它将创建并返回一个在加载完成时解析（resolve）的 promise 对象。
// 外部代码可以使用 .then 向其添加处理程序（订阅函数）：
function loadScript(src, callback) {
    return new Promise((resolve, reject) => {
        let script = document.createElement("script");
        script.src = src;

        script.onload = () => resolve(script);
        script.onerror = () => reject(new Error(`Script load error for ${src}`));

        document.head.append(script);
    });
}

// let promise = loadScript("https://cdnjs.cloudflare.com/ajax/libs/lodash.js/4.17.11/lodash.js");

// promise.then(
//     script => alert(`${script.src} is loaded!`),
//     error => alert(`Error: ${error.message}`)
// )




// ------------------------------------------------------------------------------------





// 用 promise 重新解决？

// 下列这段代码会输出什么？

let promise = new Promise(function(resolve, reject) {
  resolve(1);

  setTimeout(() => resolve(2), 1000);
});

promise.then(console.log); 1






// ------------------------------------------------------------------------------------





// 基于 promise 的延时

// 内建函数 setTimeout 使用了回调函数。请创建一个基于 promise 的替代方案。

// 函数 delay(ms) 应该返回一个 promise。
// 这个 promise 应该在 ms 毫秒后被 resolve，所以我们可以向其中添加 .then，像这样：

function delay(ms) {
    // 你的代码
    return new Promise((resolve, reject) => {
        setTimeout(() => resolve(), ms)
    })
}

delay(3000).then(() => console.log('runs after 3 seconds'));