// Promise 链
    // 简介：回调 一章中提到的问题：我们有一系列的异步任务要一个接一个地执行 — 
    // 例如，加载脚本。我们如何写出更好的代码呢？

    // Promise 提供了一些方案来做到这一点。

// 它看起来就像这样：
new Promise(resolve =>{
    setTimeout(() => resolve(1), 1000) // (*)
})
.then(res => {
    console.log("res1: " + res); // (**)
    return res * 2;
})
.then(res => {
    console.log("res2: " + res); // (***)
    return res * 2;
})
.then(res => {
    console.log("res3: " + res);
    return res * 2;
})

// res1: 1
// res2: 2
// res3: 4


// 它的理念是将 res 通过 .then 处理程序（handler）链进行传递。

// 运行流程如下：
    //- 1. 初始 promise 在 1 秒后进行 resolve (*)，
    //- 2. 然后 .then 处理程序（handler）被调用 (**)。
    //- 3. 它返回的值被传入下一个 .then 处理程序（handler）(***)
    //- 4. ……依此类推。

/* 
                resolve(1)         return 2          return 4
    new Promise -----------> then -----------> then -----------> 

*/


// 之所以这么运行，是因为对 promise.then 的调用会返回了一个 promise，
// 所以我们可以在其之上调用下一个 .then。

// 当处理程序（handler）返回一个值时，它将成为该 promise 的 result，所以将使用它调用下一个 .then。


// 经典错误：从技术上讲，也可以将多个 .then 添加到一个 promise 上。但这并不是 promise 链（chaining）。
// 例如：
let promise = new Promise(resolve => {
    setTimeout(() => resolve(1), 1000)
});

promise.then(res => {
    console.log("res1: " + res); // 1
    return res * 2;
});

promise.then(res => {
    console.log("res2: " + res); // 1
    return res * 2;
});

promise.then(res => {
    console.log("res3: " + res); // 1
    return res * 2;
});

// 在这里所做的只是一个 promise 的几个处理程序（handler）。
// 它们不会相互传递 result；相反，它们之间彼此独立运行处理任务。

// 这里它的一张示意图（你可以将其与上面的链式调用做一下比较）：
/*
            ------------ new Promise(resolve(1))  ------------
            |                     |                          |      
          .then                  .then                     .then
*/

// 在同一个 promise 上的所有 .then 获得的结果都相同 — 该 promise 的结果。
// 所以，在上面的代码中，所有 console.log 都显示相同的内容：1。





// -------------------------------------------------------------------------------------------





// 返回 promise

    // .then(handler) 中所使用的处理程序（handler）可以创建并返回一个 promise。

    // 在这种情况下，其他的处理程序（handler）将等待它 settled 后再获得其结果（result）。

// 例如：
new Promise((resolve, reject) => {
    setTimeout(() => resolve(1), 1000)
}).then(res => {
    console.log("res1: " + res);
    return new Promise(resolve => {             // (*)
        setTimeout(() => resolve(res * 2));
    });
}).then(res => {                                // (**)
    console.log("res2: " + res);
    return new Promise(resolve => {
        setTimeout(() => resolve(res * 2));
    });
}).then(res => {
    console.log("res3: " + res);
    return new Promise(resolve => {
        setTimeout(() => resolve(res * 2));
    });
})

// 这里第一个 .then 显示 1 并在 (*) 行返回 new Promise(…)。
// 1 秒后它会进行 resolve，然后 result（resolve 的参数，在这里它是 result*2）
// 被传递给第二个 .then 的处理程序（handler）。这个处理程序（handler）位于 (**) 行，
// 它显示 2，并执行相同的动作（action）。

// 所以输出与前面的示例相同：1 → 2 → 4，但是现在在每次 alert 调用之间会有 1 秒钟的延迟。

// 返回 promise 使我们能够构建异步行为链。





// -------------------------------------------------------------------------------------------





// 示例：loadScript
function loadScript(src) {
    new Promise((resolve, reject) => {
        let script = document.createElement("script");
        script.src = src;
        script.onload = resolve(script);
        script.onload = reject(new Error("load error!"));
        document.head.append(script);
    });
}
/*
loadScript("somepath/one.js")
.then(script => loadScript("somepath/two.js"))
.then(script => loadScript("somepath/three.js"))
.then(() => {
    // 脚本加载完成，我们可以在这儿使用脚本中声明的函数
    one();
    two();
    three();
})
*/

// 在这，每个 loadScript 调用都返回一个 promise，并且在它 resolve 时下一个 .then 开始运行。
// 然后，它启动下一个脚本的加载。所以，脚本是一个接一个地加载的。


// 可以向链中添加更多的异步行为（action）。请注意，代码仍然是“扁平”的 — 它向下增长，而不是向右。
// 这里没有“厄运金字塔”的迹象。

// 从技术上讲，我们可以向每个 loadScript 直接添加 .then，就像这样：
/*
loadScript("/somepath/one.js").then(script1 => {
    loadScript("/somepath/two.js").then(script2 => {
        loadScript("/somepath/three.js").then(script3 => {
            // 此函数可以访问变量 script1，script2 和 script3
            one();
            two();
            three();
        });
    });
});
*/

// 这段代码做了相同的事：按顺序加载 3 个脚本。但它是“向右增长”的。所以会有和使用回调函数一样的问题。

// 刚开始使用 promise 的人可能不知道 promise 链，所以他们就这样写了。通常，链式是首选。

// 有时候直接写 .then 也是可以的，因为嵌套的函数可以访问外部作用域。
// 在上面的例子中，嵌套在最深层的那个回调（callback）可以访问所有变量 script1，script2 和 script3。
// 但这是一个例外，而不是一条规则。






// -------------------------------------------------------------------------------------------






// Thenables
    // 确切地说，处理程序（handler）返回的不完全是一个 promise，而是返回的被称为 “thenable” 对象 
    // 一个具有方法 .then 的任意对象。它会被当做一个 promise 来对待。

    // 这个想法是，第三方库可以实现自己的“promise 兼容（promise-compatible）”对象。
    // 它们可以具有扩展的方法集，但也与原生的 promise 兼容，因为它们实现了 .then 方法。

// 这是一个 thenable 对象的示例：
/*
class Thenable {
    constructor(num) {
        this.num = num;
    }

    then(resolve, reject) {
        console.log(resolve); // function() { native code }
        // 1 秒后使用 this.num*2 进行 resolve
        setTimeout(() => resolve(this.num * 2), 1000); // (**)
    }
}
  
new Promise(resolve => resolve(1))
.then(result => {
    return new Thenable(result); // (*)
})
.then(console.log); // 1000ms 后显示 2
*/

// JavaScript 检查在 (*) 行中由 .then 处理程序（handler）返回的对象：
// 如果它具有名为 then 的可调用方法，那么它将调用该方法并提供原生的函数 resolve 和 reject 作为参数
// （类似于 executor），并等待直到其中一个函数被调用。在上面的示例中，resolve(2) 在 1 秒后被调用 (**)。
// 然后，result 会被进一步沿着链向下传递。
  
// 这个特性允许我们将自定义的对象与 promise 链集成在一起，而不必继承自 Promise。








// -------------------------------------------------------------------------------------------






// 更复杂的示例：fetch
    // 在前端编程中，promise 通常被用于网络请求。

// 将使用 fetch 方法从远程服务器加载用户信息。它有很多可选的参数，基本语法很简单：
    // let promise = fetch(url);

    // 执行这条语句，向 url 发出网络请求并返回一个 promise。
    // 当远程服务器返回 header（是在 全部响应加载完成前）时，该 promise 使用一个 response 对象来进行 resolve。

// 为了读取完整的响应，我们应该调用 response.text() 方法：
    // 当全部文字（full text）内容从远程服务器下载完成后，它会返回一个 promise，
    // 该 promise 以刚刚下载完成的这个文本作为 result 进行 resolve。

// 下面这段代码向 发送请求，并从服务器加载该文本：
/*
fetch("https://infinity-api.infinitynewtab.com/baidu_suggestion?keyword=%E7%86%8A%E5%87%BA%E6%B2%A1")
.then(response => {
    return response.blob();
}) // 当远程服务器响应时，下面的 .then 开始执行
.then(blob => {
    let reader = new FileReader();
    reader.onload = function(e) {
        let text = reader.result;
        console.log(text);
        
    }
    reader.readAsText(blob, 'GBK');
});
*/



// 例如，我们可以再向 gitee 发送一个请求，加载用户个人资料并显示头像：
function bindGithubUser(username) {
    fetch(`https://api.github.com/users/${username}`)
    .then((response) => {
        return response.json();
    })
    .then(githubUser  => {
        let img = document.createElement("img");
        img.src = githubUser.avatar_url;
        document.body.append(img);

        setTimeout(() => img.remove(), 5000);
    })
}

// bindGithubUser("ixfos");

/*
fetch('user.json')
    .then(response => response.json())
    .then(user => fetch(`https://api.github.com/users/${user.name}`))
    .then(response => response.json())
    .then(githubUser => new Promise(function(resolve, reject) { // (*)
        let img = document.createElement('img');
        img.src = githubUser.avatar_url;
        img.className = "promise-avatar-example";
        document.body.append(img);

        setTimeout(() => {
            img.remove();
            resolve(githubUser); // (**)
        }, 3000);
}))
  // 3 秒后触发
.then(githubUser => alert(`Finished showing ${githubUser.name}`));
*/



// 将代码拆分为可重用的函数：
function loadJson(url) {
    return fetch(url)
        .then(response => response.json());
}
  
function loadGithubUser(name) {
    return fetch(`https://api.github.com/users/${name}`)
        .then(response => response.json());
}
  
function showAvatar(githubUser) {
    return new Promise(function(resolve, reject) {
        let img = document.createElement('img');
        img.src = githubUser.avatar_url;
        img.className = "promise-avatar-example";
        document.body.append(img);

        setTimeout(() => {
            img.remove();
            resolve(githubUser);
        }, 3000);
    });
}
  
// 使用它们：
//- loadJson('/article/promise-chaining/user.json')
//- .then(user => loadGithubUser(user.name))
//- .then(showAvatar)
//- .then(githubUser => alert(`Finished showing ${githubUser.name}`));
// ...







// -------------------------------------------------------------------------------------------




// Promise：then 对比 catch
// 这两个代码片段是否相等？换句话说，对于任何处理程序（handler），它们在任何情况下的行为都相同吗？
/*

    promise.then(f1).catch(f2);

    // 对比：

    promise.then(f1, f2);


简要回答就是：不，它们不相等：

不同之处在于，如果 f1 中出现 error，那么在这儿它会被 .catch 处理：
promise
  .then(f1)
  .catch(f2);

  
在这儿则不会：
promise
  .then(f1, f2);
这是因为 error 是沿着链传递的，而在第二段代码中，f1 下面没有链。

换句话说，.then 将 result/error 传递给下一个 .then/.catch。所以在第一个例子中，在下面有一个 catch，
而在第二个例子中并没有 catch，所以 error 未被处理。

*/