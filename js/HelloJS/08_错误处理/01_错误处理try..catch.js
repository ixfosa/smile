// 错误处理，"try..catch"
    // 一种语法结构 try..catch，它使我们可以“捕获（catch）”错误，
    // 因此脚本可以执行更合理的操作，而不是死掉。



// “try…catch” 语法
    //- try..catch 结构由两部分组成：try 和 catch：
/*
        try {

            // 代码...

        } catch (err) {

            // 错误捕获

        }
*/

// 它按照以下步骤执行：
    //- 1. 首先，执行 try {...} 中的代码。
    //- 2. 如果这里没有错误，则忽略 catch(err)：执行到 try 的末尾并跳过 catch 继续执行。
    //- 3. 如果这里出现错误，则 try 执行停止，控制流转向 catch(err) 的开头。
    //     变量 err（我们可以使用任何名称）将包含一个 error 对象，该对象包含了所发生事件的详细信息。

// 所以，try {…} 块内的错误不会杀死脚本 — 我们有机会在 catch 中处理它。

// 没有 error 的例子：显示 console.log (1) 和 (2)：
try {
    console.log('Start of try runs');  // (1) <--
    // ...这里没有 error
    console.log('End of try runs');   // (2) <--
} catch(err) {
    console.log('Catch is ignored, because there are no errors'); // (3)
}
// Start of try runs
// End of try runs



// 包含 error 的例子：显示 (1)  (3)中的内容：
try {
    console.log('Start of try runs');  // (1) <--
    ZiQing  // Error，变量未定义！ SyntaxError: Unexpected identifier
    console.log('End of try (never reached)');   // (2) <--
} catch(err) {
    console.log(`Error has occurred!`); // (3) <--
}

// try..catch 仅对运行时的 error 有效
    // 要使得 try..catch 能工作，代码必须是可执行的。换句话说，它必须是有效的 JavaScript 代码。

// JavaScript 引擎首先会读取代码，然后运行它。在读取阶段发生的错误被称为“解析时间（parse-time）”错误，
// 并且无法恢复（从该代码内部）。这是因为引擎无法理解该代码。

// 所以，try..catch 只能处理有效代码中出现的错误。这类错误被称为“运行时的错误（runtime errors）”，
// 有时被称为“异常（exceptions）”。
console.log( 100 / 0); // Infinity

try {
    let Users = {};
    console.log(Users.user.name);
} catch(err) {
    console.log(err.toString()); // TypeError: Cannot read property 'name' of undefined
}




// -------------------------------------------------------------------------------------------




// try..catch 同步工作

// 如果在“计划的（scheduled）”代码中发生异常，例如在 setTimeout 中，则 try..catch 不会捕获到异常：
/*
try {
    setTimeout(function() {
        noSuchVariable; // 脚本将在这里停止运行
    }, 1000);
} catch (e) {
    console.log( "won't work" );
}
*/

// 因为 try..catch 包裹了计划要执行的函数，该函数本身要稍后才执行，这时引擎已经离开了 try..catch 结构。


// 为了捕获到计划的（scheduled）函数中的异常，那么 try..catch 必须在这个函数内：
setTimeout(function() {
    try {
        noSuchVariable; // try..catch 处理 error 了！
    } catch {
        console.log( "error is caught here!" ); // error is caught here!
    }
}, 1000);





// -------------------------------------------------------------------------------------------




// Error 对象
    // 发生错误时，JavaScript 生成一个包含有关其详细信息的对象。然后将该对象作为参数传递给 catch：

try {
    // ...
} catch(err) {         // <-- “error 对象”，也可以用其他参数名代替 err
    // ...
}


// 对于所有内建的 error，error 对象具有两个主要属性：
    //- name        Error 名称。例如，对于一个未定义的变量，名称是 "ReferenceError"。
    //- message     关于 error 的详细文字描述。


// 还有其他非标准的属性在大多数环境中可用。其中被最广泛使用和支持的是：
    //- stack   当前的调用栈：用于调试目的的一个字符串，其中包含有关导致 error 的嵌套调用序列的信息。


try {
    ZiQing  
} catch(err) {
    console.log("err.name: " + err.name); // err.name: ReferenceError
    console.log("err.messgge: " + err.messgge); // err.messgge: undefined
    // console.log("err.stack: " + err.stack); // err.stack: ReferenceError: ZiQing is not defined
    console.log("err: " + err); // err: ReferenceError: ZiQing is not defined
}

// -------------------------------------------------------------------------------------------





// 可选的 “catch” 绑定
    // 如果我们不需要 error 的详细信息，catch 也可以忽略它：

try {
    // ...
} catch { // <-- 没有 (err)
    // ...
}






// -------------------------------------------------------------------------------------------




// 使用 “try…catch”

// JavaScript 支持 JSON.parse(str) 方法来解析 JSON 编码的值。

// 通常，它被用来解析从网络，从服务器或是从其他来源接收到的数据。

// 收到数据后，然后像下面这样调用 JSON.parse：
let json = '{"name":"John", "age": 30}'; // 来自服务器的数据

let user = JSON.parse(json); // 将文本表示转换成 JS 对象

// 现在 user 是一个解析自 json 字符串的有自己属性的对象
console.log( user.name ); // John
console.log( user.age );  // 30


// 如果 json 格式错误，JSON.parse 就会生成一个 error，因此脚本就会“死亡”。

// 如果这样，当拿到的数据出了问题，那么访问者永远都不会知道原因（除非他们打开开发者控制台）。
// 代码执行失败却没有提示信息，这是很糟糕的用户体验。

// 用 try..catch 来处理这个 error：
json = "{ bad json }";
try {

    let user = JSON.parse(json); // <-- 当出现一个 error 时...
    console.log( user.name ); // 不工作

} catch (e) {
  // ...执行会跳转到这里并继续执行
  console.log( "Our apologies, the data has errors, we'll try to request it one more time." );
  console.log( e.name ); // SyntaxError
  console.log( e.message ); // Unexpected token b in JSON at position 2
}



// -------------------------------------------------------------------------------------------




// 抛出我们自定义的 error
    // 如果这个 json 在语法上是正确的，但是没有所必须的 name 属性该怎么办？

// 像这样：
json = '{ "age": 30 }'; // 不完整的数据
try {

  let user = JSON.parse(json); // <-- 没有 error
  console.log( user.name ); // 没有 name！

} catch (e) {
  console.log( "doesn't execute" );
}

// 这里 JSON.parse 正常执行，但是缺少 name 属性对我们来说确实是个 error。

// 为了统一进行 error 处理，我们将使用 throw 操作符。



// “Throw” 操作符
    // throw 操作符会生成一个 error 对象。

// 语法如下：
    // throw <error object>

// 技术上讲，可以将任何东西用作 error 对象。甚至可以是一个原始类型数据，
// 例如数字或字符串，但最好使用对象，最好使用具有 name 和 message 属性的对象
// （某种程度上保持与内建 error 的兼容性）。

// JavaScript 中有很多内建的标准 error 的构造器：
    //- Error，SyntaxError，ReferenceError，TypeError 等。我们也可以使用它们来创建 error 对象。

// 它们的语法是：
    //- let error = new Error(message);
    //- 或
    //- let error = new SyntaxError(message);
    //- let error = new ReferenceError(message);


// 对于内建的 error（不是对于其他任何对象，仅仅是对于 error），name 属性刚好就是构造器的名字。
// message 则来自于参数（argument）。

// 例如：
let error = new Error("Things happen o_O");

console.log(error.name); // Error
console.log(error.message); // Things happen o_O

// JSON.parse 会生成什么样的 error：
try {
    JSON.parse("{ bad json o_O }");
} catch(e) {
    console.log(e.name);        // SyntaxError
    console.log(e.message);     // Unexpected token b in JSON at position 2
}

// 那是一个 SyntaxError。示例中，缺少 name 属性就是一个 error，因为用户必须有一个 name。
// 抛出这个 error。

json = '{ "age": 30 }'; // 不完整的数据

try {

    let user = JSON.parse(json); // <-- 没有 error

    if (!user.name) {
        throw new SyntaxError("Incomplete data: no name"); // (*)
    }

    console.log( user.name );

} catch(e) {
    console.log( "JSON Error: " + e.message ); // JSON Error: Incomplete data: no name
}

// 在 (*) 标记的这一行，throw 操作符生成了包含着我们所给定的 message 的 SyntaxError，
// 与 JavaScript 自己生成的方式相同。try 的执行立即停止，控制流转向 catch 块。

// 现在，catch 成为了所有 error 处理的唯一场所：对于 JSON.parse 和其他情况都适用。





// -------------------------------------------------------------------------------------------





// 再次抛出（Rethrowing）

// 上面的例子中，我们使用 try..catch 来处理不正确的数据。
// 但是在 try {...} 块中是否可能发生 另一个预料之外的 error？
// 例如编程错误（未定义变量）或其他错误，而不仅仅是这种“不正确的数据”。


// 例如：
json = '{ "age": 30 }'; // 不完整的数据
try {
    user = JSON.parse(json); // <-- 忘记在 user 前放置 "let"

    // ...
} catch(err) {
    console.log("JSON Error: " + err); // JSON Error: ReferenceError: user is not defined
  // (实际上并没有 JSON Error)
}

// 例子中，try..catch 旨在捕获“数据不正确”的 error。
// 但是实际上，catch 会捕获到 所有 来自于 try 的 error。在这儿，它捕获到了一个预料之外的 error，
// 但是仍然抛出的是同样的 "JSON Error" 信息。这是不正确的，并且也会使代码变得更难以调试。

// 为了避免此类问题，可以采用“重新抛出”技术。规则很简单：

// catch 应该只处理它知道的 error，并“抛出”所有其他 error。

// “再次抛出（rethrowing）”技术可以被更详细地解释为：
    //- 1.  Catch 捕获所有 error。
    //- 2. 在 catch(err) {...} 块中，我们对 error 对象 err 进行分析。
    //- 3. 如果我们不知道如何处理它，那我们就 throw err。

// 通常，可以使用 instanceof 操作符判断错误类型：
try {
    user = { /*...*/ };
} catch(err) {
    if (err instanceof ReferenceError) {
        console.log('ReferenceError'); // 访问一个未定义（undefined）的变量产生了 "ReferenceError"
    }
}


// 还可以从 err.name 属性中获取错误的类名。所有原生的错误都有这个属性。
// 另一种方式是读取 err.constructor.name。

// 在下面的代码中，使用“再次抛出”，以达到在 catch 中只处理 SyntaxError 的目的：
json = '{ "age": 30 }'; // 不完整的数据
try {

    let user = JSON.parse(json);

    if (!user.name) {
        throw new SyntaxError("Incomplete data: no name");
    }

    blabla(); // 预料之外的 error

    console.log( user.name );

} catch(e) {

    if (e instanceof SyntaxError) {
        console.log( "JSON Error: " + e.message );
    } else {
        throw e; // 再次抛出 (*)
    }

}
// 如果 (*) 标记的这行 catch 块中的 error 从 try..catch 中“掉了出来”，
// 那么它也可以被外部的 try..catch 结构（如果存在）捕获到，如果外部不存在这种结构，那么脚本就会被杀死。

// 所以，catch 块实际上只处理它知道该如何处理的 error，并“跳过”所有其他的 error。

// 下面这个示例演示了这种类型的 error 是如何被另外一级 try..catch 捕获的：
function readData() {
    let json = '{ "age": 30 }';

    try {
        // ...
        blabla(); // error!
    } catch (e) {
        // ...
        if (!(e instanceof SyntaxError)) {
             throw e; // 再次抛出（不知道如何处理它）
        }
    }
}

try {
    readData();
} catch (e) {
     console.log( "External catch got: " + e ); // 捕获了它！
}

// 上面这个例子中的 readData 只知道如何处理 SyntaxError，而外部的 try..catch 知道如何处理任意的 error。







// -------------------------------------------------------------------------------------------




// try…catch…finally

// try..catch 结构可能还有一个代码子句（clause）：finally。
    // 如果它存在，它在所有情况下都会被执行：

// try 之后，如果没有 error，
// catch 之后，如果没有 error。

// 该扩展语法如下所示：
/*
        try {
        ... 尝试执行的代码 ...
        } catch(e) {
        ... 处理 error ...
        } finally {
        ... 总是会执行的代码 ...
        }
*/

try {
    console.log( 'try' );
    ZiQing; // Error
} catch (e) {
    console.log( 'catch' );
} finally {
    console.log( 'finally' );
}
// try
// catch
// finally


try {
    console.log( 'try' );
} catch (e) {
    console.log( 'catch' );
} finally {
    console.log( 'finally' );
}
// try
// finally




// 变量和 try..catch..finally 中的局部变量
    // 如果使用 let 在 try 块中声明变量，那么该变量将只在 try 块中可见。



// finally 和 return
    // finally 子句适用于 try..catch 的 任何 出口。这包括显式的 return。

// 在下面这个例子中，在 try 中有一个 return。在这种情况下，finally 会在控制转向外部代码前被执行。
function func() {

    try {
        return 1;

    } catch (e) {
        /* ... */
    } finally {
        console.log( 'finally' );
    }
}

console.log( func() ); // 先执行 finally 中的  console.log，然后执行这个  console.log





// try..finally
    // 没有 catch 子句的 try..finally 结构也很有用。
    // 当我们不想在这儿处理 error（让它们 fall through），但是需要确保我们启动的处理需要被完成。
/*
    function func() {
        // 开始执行需要被完成的操作（比如测量）
        try {
            // ...
        } finally {
            // 完成前面我们需要完成的那件事儿，即使 try 中的执行失败了
        }
    }
*/
// 上面的代码中，由于没有 catch，所以 try 中的 error 总是会使代码执行跳转至函数 func() 外。
// 但是，在跳出之前需要执行 finally 中的代码。




// -------------------------------------------------------------------------------------------



// 全局 catch

// 设想一下，在 try..catch 结构外有一个致命的 error，然后脚本死亡了。
// 这个 error 就像编程错误或其他可怕的事那样。

// 有什么办法可以用来应对这种情况吗？我们可能想要记录这个 error，
// 并向用户显示某些内容（通常用户看不到错误信息）等。

// 规范中没有相关内容，但是代码的执行环境一般会提供这种机制，因为它确实很有用。
// 例如，Node.JS 有 process.on("uncaughtException")。
// 在浏览器中，我们可以将将一个函数赋值给特殊的 window.onerror 属性，该函数将在发生未捕获的 error 时执行。

// 语法如下：
   //-   window.onerror = function(message, url, line, col, error) {
   //-       // ...
   //-   };

        //- message     Error 信息。
        //- url         发生 error 的脚本的 URL。
        //- line，col   发生 error 处的代码的行号和列号。
        //- error       Error 对象。


// 例如：
/*
    <script>
        window.onerror = function(message, url, line, col, error) {
            console.log(`${message}\n At ${line}:${col} of ${url}`);
        };

        function readData() {
            badFunc(); // 啊，出问题了！
        }

        readData();
    </script>
*/

// 全局错误处理程序 window.onerror 的作用通常不是恢复脚本的执行 — 如果发生编程错误，
// 那这几乎是不可能的，它的作用是将错误信息发送给开发者。

// 也有针对这种情况提供错误日志的 Web 服务，例如 https://errorception.com 或 http://www.muscula.com。

// 它们会像这样运行：
    //- 1. 注册该服务，并拿到一段 JS 代码（或脚本的 URL），然后插入到页面中。
    //- 2. 该 JS 脚本设置了自定义的 window.onerror 函数。
    //- 3. 当发生 error 时，它会发送一个此 error 相关的网络请求到服务提供方。
    //- 4. 可以登录到服务方的 Web 界面来查看这些 error。





// -------------------------------------------------------------------------------------------





// 总结

    // try..catch 结构允许我们处理执行过程中出现的 error。从字面上看，
    // 它允许“尝试”运行代码并“捕获”其中可能发生的错误。

    // 语法如下：
/*
        try {
            // 执行此处代码
        } catch(err) {
            // 如果发生错误，跳转至此处
            // err 是一个 error 对象
        } finally {
            // 无论怎样都会在 try/catch 之后执行
        }
*/

    // 这儿可能会没有 catch 部分或者没有 finally，所以 try..catch 或 try..finally 都是可用的。

    // Error 对象包含下列属性：
       //- message — 人类可读的 error 信息。
       //- name — 具有 error 名称的字符串（Error 构造器的名称）。
       //- stack（没有标准，但得到了很好的支持）— Error 发生时的调用栈。

    // 如果我们不需要 error 对象，我们可以通过使用 catch { 而不是 catch(err) { 来省略它。

    // 也可以使用 throw 操作符来生成自定义的 error。从技术上讲，throw 的参数可以是任何东西，
    // 但通常是继承自内建的 Error 类的 error 对象。下一章我们会详细介绍扩展 error。

    // 再次抛出（rethrowing）是一种错误处理的重要模式：
        // catch 块通常期望并知道如何处理特定的 error 类型，因此它应该再次抛出它不知道的 error。

    // 即使没有 try..catch，大多数执行环境也允许设置“全局”错误处理程序来捕获“掉出（fall out）”的 error。
    // 在浏览器中，就是 window.onerror。





// -------------------------------------------------------------------------------------------





// 使用 finally 还是直接放在代码后面？

// 比较下面两个代码片段。

/*
    // 第一个代码片段，使用 finally 在 try..catch 之后执行代码：
    try {
        work work
    } catch (e) {
        handle errors
    } finally {
        cleanup the working space
    }

    // 第二个代码片段，将清空工作空间的代码放在了 try..catch 之后：
    try {
    work work
    } catch (e) {
    handle errors
    }
    cleanup the working space
*/

// 我们肯定需要在工作后进行清理，无论工作过程中是否有 error 都不影响。

// 在这儿使用 finally 更有优势，还是说两个代码片段效果一样？如果在这儿有这样的优势，如果需要，请举例说明。



// 当我们看函数中的代码时，差异就变得很明显了。
// 如果在这儿有“跳出” try..catch 的行为，那么这两种方式的表现就不同了。

// 例如，当 try..catch 中有 return 时。finally 子句会在 try..catch 的 任意 出口处起作用，
// 即使是通过 return 语句退出的也是如此：在 try..catch 刚刚执行完成后，但在调用代码获得控制权之前。

function f() {
    try {
        console.log('start');
        return "result";
    } catch (e) {
        // ...
    } finally {
        console.log('cleanup!');
    }
}

f(); // cleanup!

// 或者当有 throw 时，如下所示：
function f2() {
    try {
        console.log('start');
        throw new Error("an error");
    } catch (e) {
            // ...
            if("can't handle the error") {
            throw e;
        }

    } finally {
        console.log('cleanup!')
    }
}

f2(); // cleanup!

// 正是这里的 finally 保证了 cleanup。如果我们只是将代码放在函数 f 的末尾，则在这些情况下它不会运行。