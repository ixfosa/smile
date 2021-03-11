
// 什么是模块？
    // 一个模块（module）就是一个文件。一个脚本就是一个模块。

// 模块可以相互加载，并可以使用特殊的指令 export 和 import 来交换功能，从另一个模块调用一个模块的函数：
    //- export 关键字标记了可以从当前模块外部访问的变量和函数。
    //- import 关键字允许从其他模块导入功能。


// 例如，我们有一个 sayHi.js 文件导出了一个函数：
/*
    // 📁 sayHi.js
    export function sayHi(user) {
        console.log(`Hello, ${user}!`);
    }

    // 然后另一个文件可能导入并使用了这个函数：
    // 📁 main.js
    import {sayHi} from './sayHi.js';

    console.log(sayHi); // function...
    sayHi('fo'); // Hello, John!
*/


// import 指令通过相对于当前文件的路径 ./sayHi.js 加载模块，并将导入的函数 sayHi 分配（assign）给相应的变量。

// 在浏览器中运行一下这个示例。

// 由于模块支持特殊的关键字和功能，因此我们必须通过使用 
// <script type="module"> 特性（attribute）来告诉浏览器，此脚本应该被当作模块（module）来对待。

/*
    // index.html
    <!doctype html>
    <script type="module">
    import {sayHi} from './say.js';

    document.body.innerHTML = sayHi('John');
    </script>
*/

// 浏览器会自动获取并解析（evaluate）导入的模块（如果需要，还可以分析该模块的导入），然后运行该脚本。


// 模块只通过 HTTP(s) 工作，在本地文件则不行


// 注意直接在js文件里run报错， 需要 package.json 文件， 并set "type": "module"
    // Warning: To load an ES module, set "type": "module" in the package.json or use the .mjs extension.







// -------------------------------------------------------------------------------------------




// 始终使用 “use strict”
    // 模块始终默认使用 use strict，例如，对一个未声明的变量赋值将产生错误（
/*
    <script type="module">
    a = 5; // error
    </script>
*/







// -------------------------------------------------------------------------------------------




// 模块级作用域
    // 每个模块都有自己的顶级作用域（top-level scope）。换句话说，一个模块中的顶级作用域变量和函数在其他脚本中是不可见的。

// 在下面这个例子中，我们导入了两个脚本，hello.js 尝试使用在 user.js 中声明的变量 user，失败了：
/*
    index.html
        <!doctype html>
        <script type="module" src="user.js"></script>
        <script type="module" src="hello.js"></script>


    模块期望 export 它们想要被外部访问的内容，并 import 它们所需要的内容。

    所以，我们应该将 user.js 导入到 hello.js 中，并从中获取所需的功能，而不要依赖于全局变量。

    这是正确的变体：
    hello.js
        import {user} from './user.js';
        document.body.innerHTML = user; 



    在浏览器中，每个 <script type="module"> 也存在独立的顶级作用域（译注：在浏览器控制台可以看到 error 信息）。
    script type="module">
        // 变量仅在这个 module script 内可见
        let user = "fo";
    </script>

    <script type="module">
        alert(user); // Error: user is not defined
    </script>

    如果真的需要创建一个 window-level 的全局变量，可以将其明确地赋值给 window，
    并以 window.user 来访问它。但是这需要你有足够充分的理由，否则就不要这样做。

*/







// -------------------------------------------------------------------------------------------





// 模块代码仅在第一次导入时被解析
    // 如果同一个模块被导入到多个其他位置，那么它的代码仅会在第一次导入时执行，
    // 然后将导出（export）的内容提供给所有的导入（importer）。


// 首先，如果执行一个模块中的代码会带来副作用（side-effect），
// 例如显示一条消息，那么多次导入它只会触发一次显示 —— 即第一次：
/*
    // 📁 showMsg.js
        console.log("Module is evaluated!");

    // 在不同的文件中导入相同的模块

    // 📁 1.js
        import `./showMsg.js`; // Module is evaluated!

    // 📁 2.js
        import `./showMsg.js`; // (什么都不显示)


    
    在实际开发中，顶级模块代码主要用于初始化，内部数据结构的创建，并且如果我们希望某些东西可以重用 — 请导出它。

    下面是一个高级点的例子。

    假设一个模块导出了一个对象：

    // 📁 admin.js
        export let admin = {
            name: "fo"
        };




    如果这个模块被导入到多个文件中，模块仅在第一次被导入时被解析，并创建 admin 对象，然后将其传入到所有的导入。

    所有的导入都只获得了一个唯一的 admin 对象：
    // 📁 1.js
        import {admin} from './admin.js';
        admin.name = "zi";

    // 📁 2.js
        import {admin} from './admin.js';
        alert(admin.name); // zi

    // 1.js 和 2.js 导入的是同一个对象
    // 在 1.js 中对对象做的更改，在 2.js 中也是可见的

    所以，让我们重申一下 —— 模块只被执行一次。生成导出，然后它被分享给所有对其的导入，
    所以如果某个地方修改了 admin 对象，其他的模块也能看到这个修改。

    这种行为让我们可以在首次导入时 设置 模块。我们只需要设置其属性一次，然后在进一步的导入中就都可以直接使用了。


    例如，下面的 admin.js 模块可能提供了特定的功能，但是希望凭证（credential）从外部进入 admin 对象：
    // 📁 admin.js
        export let admin = { };

        export function sayHi() {
            alert(`Ready to serve, ${admin.name}!`);
        }

    在 init.js 中 ——  APP 的第一个脚本，设置了 admin.name。现在每个位置都能看到它，包括在 admin.js 内部的调用。
    // 📁 init.js
        import {admin} from './admin.js';
        admin.name = "Zi";

    另一个模块也可以看到 admin.name：
    // 📁 other.js
        import {admin, sayHi} from './admin.js';
        alert(admin.name); // Zi
        sayHi(); // Ready to serve, Zi!

*/






// -------------------------------------------------------------------------------------------






// import.meta
    // import.meta 对象包含关于当前模块的信息。

    // 它的内容取决于其所在的环境。在浏览器环境中，它包含当前脚本的 URL，或者如果它是在 HTML 中的话，
    // 则包含当前页面的 URL。
/*
    <script type="module">
        alert(import.meta.url); // 脚本的 URL（对于内嵌脚本来说，则是当前 HTML 页面的 URL）
    </script>
*/





// -------------------------------------------------------------------------------------------




// 在一个模块中，“this” 是 undefined
    // 在一个模块中，顶级 this 是 undefined。

// 将其与非模块脚本进行比较会发现，非模块脚本的顶级 this 是全局对象：
/*
    <script>
        alert(this); // window
    </script>

    <script type="module">
        alert(this); // undefined
    </script>
*/








// -------------------------------------------------------------------------------------------







// 浏览器特定功能
    // 与常规脚本相比，拥有 type="module" 标识的脚本有一些特定于浏览器的差异。


/*
    1. 模块脚本是延迟的
        模块脚本 总是 被延迟的，与 defer 特性对外部脚本和内联脚本（inline script）的影响相同。

        也就是说：
            下载外部模块脚本 <script type="module" src="..."> 不会阻塞 HTML 的处理，它们会与其他资源并行加载。
            模块脚本会等到 HTML 文档完全准备就绪（即使它们很小并且比 HTML 加载速度更快），然后才会运行。
            保持脚本的相对顺序：在文档中排在前面的脚本先执行。

        它的一个副作用是，模块脚本总是会“看到”已完全加载的 HTML 页面，包括在它们下方的 HTML 元素。
        例如：
            <script type="module">
                alert(typeof button); // object：脚本可以“看见”下面的 button
                // 因为模块是被延迟的（deferred，所以模块脚本会在整个页面加载完成后才运行
            </script>

        相较于下面这个常规脚本：
            <script>
                alert(typeof button); // button 为 undefined，脚本看不到下面的元素
                // 常规脚本会立即运行，常规脚本的运行是在在处理页面的其余部分之前进行的
            </script>
            <button id="button">Button</button>

        请注意：上面的第二个脚本实际上要先于前一个脚本运行！所以我们会先看到 undefined，然后才是 object。

        这是因为模块脚本是被延迟的，所以要等到 HTML 文档被处理完成才会执行它。
        而常规脚本则会立即运行，所以我们会先看到常规脚本的输出。

        当使用模块脚本时，应该知道 HTML 页面在加载时就会显示出来，在 HTML 页面加载完成后才会执行 JavaScript 
        模块，因此用户可能会在 JavaScript 应用程序准备好之前看到该页面。某些功能那时可能还无法正使用。
        应该放置“加载指示器（loading indicator）”，否则，请确保不会使用户感到困惑。
*/


/*
    2. Async 适用于内联脚本（inline script）
        对于非模块脚本，async 特性（attribute）仅适用于外部脚本。
        异步脚本会在准备好后立即运行，独立于其他脚本或 HTML 文档。

        对于模块脚本，它也适用于内联脚本。

        例如，下面的内联脚本具有 async 特性，因此它不会等待任何东西。

        它执行导入（fetch ./analytics.js），并在准备导入完成时运行，即使 HTML 文档还未完成，或者其他脚本仍在等待处理中。

        这对于不依赖任何其他东西的功能来说是非常棒的，例如计数器，广告，文档级事件监听器。

        <!-- 所有依赖都获取完成（analytics.js）然后脚本开始运行 -->
        <!-- 不会等待 HTML 文档或者其他 <script> 标签 -->
        <script async type="module">
            import {counter} from './analytics.js';
            counter.count();
        </script>

*/


/*
    3. 外部脚本
        具有 type="module" 的外部脚本（external script）在两个方面有所不同：

        具有相同 src 的外部脚本仅运行一次：

        <!-- 脚本 my.js 被加载完成（fetched）并只被运行一次 -->
        <script type="module" src="my.js"></script>
        <script type="module" src="my.js"></script>

        从另一个源（例如另一个网站）获取的外部脚本需要 CORS header，如我们在 Fetch：跨源请求 一章中所讲的那样。
        换句话说，如果一个模块脚本是从另一个源获取的，则远程服务器必须提供表示允许获取的 
        header Access-Control-Allow-Origin。

        <!-- another-site.com 必须提供 Access-Control-Allow-Origin -->
        <!-- 否则，脚本将无法执行 -->
        <script type="module" src="http://another-site.com/their.js"></script>

        默认这样做可以确保更好的安全性。

*/


/*
    4. 不允许裸模块（“bare” module）
        在浏览器中，import 必须给出相对或绝对的 URL 路径。没有任何路径的模块被称为“裸（bare）”模块。
        在 import 中不允许这种模块。

        例如，下面这个 import 是无效的：

        import {sayHi} from 'sayHi'; // Error，“裸”模块
        // 模块必须有一个路径，例如 './sayHi.js' 或者其他任何路径

        某些环境，像 Node.js 或者打包工具（bundle tool）允许没有任何路径的裸模块，
        因为它们有自己的查找模块的方法和钩子（hook）来对它们进行微调。但是浏览器尚不支持裸模块。

*/


/*
    5. 兼容性，“nomodule”
        旧时的浏览器不理解 type="module"。未知类型的脚本会被忽略。
        对此，我们可以使用 nomodule 特性来提供一个后备：

        <script type="module">
            alert("Runs in modern browsers");
        </script>

        <script nomodule>
            alert("Modern browsers know both type=module and nomodule, so skip this")
            alert("Old browsers ignore script with unknown type=module, but execute this.");
        </script>

*/





// -------------------------------------------------------------------------------------------






// 构建工具
    // 在实际开发中，浏览器模块很少被以“原始”形式进行使用。
    // 通常，我们会使用一些特殊工具，例如 Webpack，将它们打包在一起，然后部署到生产环境的服务器。

    // 使用打包工具的一个好处是 —— 它们可以更好地控制模块的解析方式，允许我们使用裸模块和更多的功能，
    // 例如 CSS/HTML 模块等。

    // 构建工具做以下这些事：
        //- 1. 从一个打算放在 HTML 中的 <script type="module"> “主”模块开始。

        //- 2. 分析它的依赖：它的导入，以及它的导入的导入等。

        //- 3. 使用所有模块构建一个文件（或者多个文件，这是可调的），并用打包函数（bundler function）替代原生的
        //     import 调用，以使其正常工作。还支持像 HTML/CSS 模块等“特殊”的模块类型。

        //- 4. 在处理过程中，可能会应用其他转换和优化：
            //- 删除无法访问的代码。
            //- 删除未使用的导出（“tree-shaking”）。
            //- 删除特定于开发的像 console 和 debugger 这样的语句。
            //- 可以使用 Babel 将前沿的现代的 JavaScript 语法转换为具有类似功能的旧的 JavaScript 语法。
            //- 压缩生成的文件（删除空格，用短的名字替换变量等）。

    // 如果使用打包工具，那么脚本会被打包进一个单一文件（或者几个文件），
    // 在这些脚本中的 import/export 语句会被替换成特殊的打包函数（bundler function）。
    // 因此，最终打包好的脚本中不包含任何 import/export，它也不需要 type="module"，
    // 可以将其放入常规的 <script>：

    // <!-- 假设我们从诸如 Webpack 这类的打包工具中获得了 "bundle.js" 脚本 -->
    // <script src="bundle.js"></script>

    // 也就是说，原生模块也是可以使用的。




// -------------------------------------------------------------------------------------------





// 总结

    // 1. 一个模块就是一个文件。浏览器需要使用 <script type="module"> 以使 import/export 可以工作。
    // 模块（译注：相较于常规脚本）有几点差别：
        //- 默认是延迟解析的（deferred）。
        //- Async 可用于内联脚本。
        //- 要从另一个源（域/协议/端口）加载外部脚本，需要 CORS header。
        //- 重复的外部脚本会被忽略

    // 2. 模块具有自己的本地顶级作用域，并可以通过 import/export 交换功能。
    // 3. 模块始终使用 use strict。
    // 4. 模块代码只执行一次。导出仅创建一次，然后会在导入之间共享。

    // 当我们使用模块时，每个模块都会实现特定功能并将其导出。然后我们使用 import 将其直接导入到需要的地方即可。
    // 浏览器会自动加载并解析脚本。

    // 在生产环境中，出于性能和其他原因，开发者经常使用诸如 Webpack 之类的打包工具将模块打包到一起。


