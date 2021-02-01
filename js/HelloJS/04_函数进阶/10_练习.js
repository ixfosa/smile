// 二次 bind

// 我们可以通过额外的绑定改变 this 吗？

// 输出将会是什么？

function f() {
  console.log(this.name);
}

f = f.bind( {name: "John"} ).bind( {name: "Ann" } );

f(); // John

// f.bind(...) 返回的外来（exotic）绑定函数 对象仅在创建的时候记忆上下文（以及参数，如果提供了的话）。

// 一个函数不能被重绑定（re-bound）。




// -------------------------------------------------------------------------------------------



// bind 后的函数属性

// 函数的属性中有一个值。bind 之后它会改变吗？为什么，阐述一下？

function sayHi() {
  console.log( this.name );
}
sayHi.test = 5;

let bound = sayHi.bind({
  name: "John"
});

console.log( bound.test ); // 输出将会是什么？为什么？ undefined

// bind 的结果是另一个对象。它并没有 test 属性。





// -------------------------------------------------------------------------------------------




// 修复丢失了 "this" 的函数

// 下面代码中对 askPassword() 的调用将会检查 password，然后基于结果调用 user.loginOk/loginFail。

// 但是它导致了一个错误。为什么？

// 修改高亮的行，以使所有内容都能正常工作（其它行不用修改）。


/*


function askPassword(pwd, ok, fail) {
  let password = pwd;
  if (password == "rockstar") ok();
  else fail();
}

let user = {
  name: 'ixfosa',

    loginOk() {
        console.log(`${this.name} logged in`);
    },

    loginFail() {
        console.log(`${this.name} failed to log in`);
    },

};

// askPassword("ixfosa", user.loginOk, user.loginFail); //undefined failed to log in
askPassword("ixfosa",  user.loginOk.bind(user), user.loginFail.bind(user)); // ixfosa failed to log in

askPassword(() => "rockstar", user.loginOk(), () => user.loginFail());
//通常这也能正常工作，也看起来挺好的。

// 但是可能会在更复杂的场景下失效，例如变量 user 在调用 askPassword 之后但在访问者应答和调用 () => user.loginOk() 之前被修改。

*/



// -------------------------------------------------------------------------------------------




// 偏函数在登录中的应用

// 这个任务是比 修复丢失了 "this" 的函数 略微复杂的变体。

// user 对象被修改了。现在不是两个函数 loginOk/loginFail，现在只有一个函数 user.login(true/false)。

// 在下面的代码中，我们应该向 askPassword 传入什么参数，以使得 user.login(true) 结果是 ok，
// user.login(fasle) 结果是 fail？

function askPassword(pwd, ok, fail) {
  let password = pwd;
  if (password == "ixfosa") ok();
  else fail();
}

let user = {
  name: 'ixfosa',

  login(result) {
    console.log( this.name + (result ? ' logged in' : ' failed to log in') );
  }
};

// askPassword(?, ?); // ?
// 你只能修改高亮部分的代码。

askPassword("ixfosa", user.login.bind(user, true), user.login.bind(user, false)); // ixfosa logged in
