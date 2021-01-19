// 现代模式，"use strict"


/**
    长久以来，JavaScript 不断向前发展且并未带来任何兼容性问题。新的特性被加入，旧的功能也没有改变。

    这么做有利于兼容旧代码，但缺点是 JavaScript 创造者的任何错误或不完善的决定也将永远被保留在 JavaScript 语言中。

    这种情况一直持续到 2009 年 ECMAScript 5 (ES5) 的出现。ES5 规范增加了新的语言特性并且修改了一些已经存在的特性。
    为了保证旧的功能能够使用，大部分的修改是默认不生效的。你需要一个特殊的指令 —— "use strict" 来明确地激活这些特性。

    没有办法取消 use strict
    没有类似于 "no use strict" 这样的指令可以使程序返回默认模式。
    一旦进入了严格模式，就没有回头路了。
 */

// 欢迎将 "use strict"; 写在脚本的顶部。稍后，当你的代码全都写在了 class 和 module 中时，
// 你则可以将 "use strict"; 这行代码省略掉。

//"use strict";
    // 代码以现代模式工作
    // ...


// 确保 “use strict” 出现在最顶部
// 这里的严格模式就没有被启用：

alert("some code");
// 下面的 "use strict" 会被忽略，必须在最顶部。
"use strict";
// 严格模式没有被激活




(function() {
    'use strict';
  
    // ...你的代码...
  })()



  /**
    浏览器控制台
    当你使用 开发者控制台 运行代码时，请注意它默认是不启动 use strict 的。

    有时，当 use strict 会对代码产生一些影响时，你会得到错误的结果。

    那么，怎么在控制台中启用 use strict 呢？

    首先，你可以尝试搭配使用 Shift+Enter 按键去输入多行代码，然后将 use strict 放在代码最顶部，就像这样：

    'use strict'; <Shift+Enter 换行>
    //  ...你的代码
    <按下 Enter 以运行>
   */

