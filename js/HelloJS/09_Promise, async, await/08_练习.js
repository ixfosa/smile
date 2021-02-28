


// 用 async/await 来重写

// 重写下面这个来自 Promise 链 一章的示例代码，使用 async/await 而不是 .then/catch：

function loadJson(url) {
    return fetch(url)
        .then(response => {
        if (response.status == 200) {
            return response.json();
        } else {
            throw new Error(response.status);
        }
    });
}

// loadJson('no-such-user.json').catch(alert); // Error: 404


async function myLoadJson(url) {
    let response = await fetch(url);

    if (response.status == 200) {
        let json = await response.json();
        return json;

        // return response.json();  // 外部的代码就必须 await 这个 promise resolve。
    } else {
        throw new Error(response.status);
    }
}
   





// -------------------------------------------------------------------------------------------




// 使用 async/await 重写 "rethrow"

// 下面你可以看到 “rethrow” 的例子。让我们来用 async/await 重写它，而不是使用 .then/catch。

// 同时，我们可以在 demoGithubUser 中使用循环以摆脱递归：在 async/await 的帮助下很容易实现。

class HttpError extends Error {
    constructor(response) {
        super(`${response.status} for ${response.url}`);
        this.name = 'HttpError';
        this.response = response;
    }
}

function loadJson(url) {
     return fetch(url)
        .then(response => {
        if (response.status == 200) {
            return response.json();
        } else {
            throw new HttpError(response);
        }
    });
}


async function loadJson2(url) {
    let response = await fetch(url);

    if (response.status == 200) {
        return response.json();
    } else {
        throw new HttpError(response);
    }
}


// 询问用户名，直到 github 返回一个合法的用户
function demoGithubUser() {
    let name = prompt("Enter a name?", "iliakan");

    return loadJson(`https://api.github.com/users/${name}`)
            .then(user => {
            alert(`Full name: ${user.name}.`);
            return user;
        })
        .catch(err => {
            if (err instanceof HttpError && err.response.status == 404) {
                alert("No such user, please reenter.");
                return demoGithubUser();
            } else {
                throw err;
            }
        });
}

// demoGithubUser();

async function demoGithubUser2() {

    let user;
    while(true) {
      let name = prompt("Enter a name?", "iliakan");
  
      try {
        user = await loadJson(`https://api.github.com/users/${name}`);
        break; // 没有 error，退出循环
      } catch(err) {
        if (err instanceof HttpError && err.response.status == 404) {
          // 循环将在 alert 后继续
          alert("No such user, please reenter.");
        } else {
          // 未知的 error，再次抛出（rethrow）
          throw err;
        }
      }
    }

    alert(`Full name: ${user.name}.`);
    return user;
}






// -------------------------------------------------------------------------------------------






// 在非 async 函数中调用 async 函数

// 我们有一个名为 fn 的“普通”函数。你会怎样调用 async 函数 wait() 并在 f 中使用其结果？

async function wait() {
  await new Promise(resolve => setTimeout(resolve, 1000));

  return 10;
}

function fn() {
    // 这里你应该怎么写？
    // 我们需要调用 async wait() 并等待以拿到结果 10
    // 记住，我们不能使用 "await"

    // console.log(wait()); // Promise { <pending> }

    wait().then(console.log); // 10

    (async () => {
        let res = await wait();
        console.log(res);
    })() // 10

}
fn();
// P.S. 这个任务其实很简单，但是对于 async/await 新手开发者来说，这个问题却很常见。