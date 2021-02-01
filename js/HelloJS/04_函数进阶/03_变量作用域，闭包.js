

//  let/const
    //-  在 JavaScript 中，有三种声明变量的方式：let，const（现代方式），var（过去留下来的方式）。
    //-  用 const 声明的变量的行为也相同（译注：与 let 在作用域等特性上是相同的），因此，本文也涉及用 const 进行变量声明。
    //-  旧的 var 与上面两个有着明显的区别，我们将在 旧时的 "var" 中详细介绍。



// -------------------------------------------------------------------------------------------



// 代码块
    // 如果在代码块 {...} 内声明了一个变量，那么这个变量只在该代码块内可见。
    {
        let msg = "hello js";
        console.log(msg) // hello js
    }
    // console.log(msg); // ReferenceError: msg is not defined
    
    
    // 可以使用它来隔离一段代码，该段代码执行自己的任务，并使用仅属于自己的变量：
    {
        let msg = "hello js";
        console.log(msg) // hello js
    }
    
    {
        let msg = "hello js";
        console.log(msg) // hello js
    }
    
    
    
    // 对于 if，for 和 while 等，在 {...} 中声明的变量也仅在内部可见：
    if (true) {
        let message = "hello java";
        console.log(message); // hello java
    }
    // console.log(message); // ReferenceError: message is not defined
    
    
    // 对于 for 和 while 循环也是如此：
    
    for (let i = 0; i < 3; i++) {
        // 变量 i 仅在这个 for 循环的内部可见
        console.log(i); // 0，然后是 1，然后是 2
    }
    // console.log(i); // Error, no such variable
    
    
    
    
    
    // -------------------------------------------------------------------------------------------
    
    
    
    
    
    // 嵌套函数
    
    // 当一个函数是在另一个函数中创建的时，那么该函数就被称为“嵌套”的。
    
    function sayHiBye(name) {
    
        // 辅助嵌套函数使用如下
        function getName() {
            return name;
        }
        console.log( "Hello, " + getName() );
        console.log( "Bye, " + getName() );
    }
    
    sayHiBye("梓晴");  // Hello, 梓晴, Bye, 梓晴
    
     
    
    
    
    
    
    // -------------------------------------------------------------------------------------------
    
    
    
    
    // 词法环境
    
    
    // Step 1. 变量
        //- 在 JavaScript 中，每个运行的函数，代码块 {...} 以及整个脚本，
        //- 都有一个被称为 词法环境（Lexical Environment） 的内部（隐藏）的关联对象。
    
    // 词法环境对象由两部分组成：
        //- 环境记录（Environment Record） —— 一个存储所有局部变量作为其属性（包括一些其他信息，例如 this 的值）的对象。
        //- 对外部词法环境 的引用，与外部代码相关联。
        //- 一个“变量”只是 环境记录 这个特殊的内部对象的一个属性。“获取或修改变量”意味着“获取或修改词法环境的一个属性”。
    
    /*
        举个例子，这段没有函数的简单的代码中只有一个词法环境：
        let phrase = "hello";      ---  phrase: "Hello"(Lexical Environment)  ---outer--> null
        alert(phrase);
        这就是所谓的与整个脚本相关联的 全局 词法环境。
        （Lexical Environment） 表示环境记录（变量存储），箭头表示外部引用。全局词法环境没有外部引用，所以箭头指向了 null。
    
    
        随着代码开始并继续运行，词法环境发生了变化。
        execution start             ----  phrase: <uninitialized> (Lexical Environment) ---outer--> null
        let phrase;                 ----  phrase: undefined (Lexical Environment)
        phrase = "hello";           ----  phrase: "Hello"  (Lexical Environment)
        phrase = "bye";             ----  phrase: "bye"  (Lexical Environment)
    
        右侧的示了执行过程中全局词法环境的变化：
            1. 当脚本开始运行，词法环境预先填充了所有声明的变量。
                最初，它们处于“未初始化（Uninitialized）”状态。这是一种特殊的内部状态，
                这意味着引擎知道变量，但是在用 let 声明前，不能引用它。几乎就像变量不存在一样。
            2. 然后 let phrase 定义出现了。它尚未被赋值，因此它的值为 undefined。从这一刻起，我们就可以使用变量了。
            3. phrase 被赋予了一个值。
            4. phrase 的值被修改。
           
    
        变量是特殊内部对象的属性，与当前正在执行的（代码）块/函数/脚本有关。
        操作变量实际上是操作该对象的属性。
    
        词法环境是一个规范对象
            “词法环境”是一个规范对象（specification object）：
                它仅仅是存在于 编程语言规范 中的“理论上”存在的，用于描述事物如何运作的对象。
                我们无法在代码中获取该对象并直接对其进行操作。
    
        但 JavaScript 引擎同样可以优化它，比如清除未被使用的变量以节省内存和执行其他内部技巧等，
        但显性行为应该是和上述的无差。
    
    */
    
    
    
    // -------------------------------------------------------------------------------------------
    
    
    
    
    // Step 2. 函数声明
        //- 一个函数其实也是一个值，就像变量一样。
        //- 不同之处在于函数声明的初始化会被立即完成。
        //- 当创建了一个词法环境（Lexical Environment）时，函数声明会立即变为即用型函数（不像 let 那样直到声明处才可用）。
        //- 这就是为什么我们可以在（函数声明）的定义之前调用函数声明。
        //- 例如，这是添加一个函数时全局词法环境的初始状态：
    
    
    /*
        例如，这是添加一个函数时全局词法环境的初始状态：
    
        execution start                 ---    phrase: <uninitialized> say: function  ---outer--> null
    
        let phrase = "hello";           ---
    
        function say(name) {
            alert(`${name ${name}}`);
        }
    
        正常来说，这种行为仅适用于函数声明，而 不适用于我们将函数分配给变量的函数表达式，
        例如 let say = function(name)...。
    
    */
    
        
    
    
    
    // -------------------------------------------------------------------------------------------
    
    
    
    
    // Step 3. 内部和外部的词法环境
        // 在一个函数运行时，在调用刚开始时，会自动创建一个新的词法环境以存储这个调用的局部变量和参数。
    
    /*
        例如，对于 say("John")，它看起来像这样（当前执行位置在箭头标记的那一行上）：
    
            execution start                 ---    phrase: <uninitialized> say: function  ---outer--> null
    
            let phrase = "hello";           ---
                                                                                     |
            function say(name) {            ----  (Lexical Environment of the call)  |  say: function phrase: "Hello"  ---outer--> null
        >       alert(`${name ${name}}`);      |    name: "梓晴"  ---outer-->        |
            }                               ---- 
            say("梓晴");
    
        在这个函数调用期间，我们有两个词法环境：内部一个（用于函数调用）和外部一个（全局）：
            内部词法环境与 say 的当前执行相对应。它具有一个单独的属性：name，函数的参数。我们调用的是 say("梓晴")，所以 name 的值为 "梓晴"。
            外部词法环境是全局词法环境。它具有 phrase 变量和函数本身。
    
        内部词法环境引用了 outer。
    
        当代码要访问一个变量时 —— 首先会搜索内部词法环境，然后搜索外部环境，然后搜索更外部的环境，以此类推，直到全局词法环境。
    
        如果在任何地方都找不到这个变量，那么在严格模式下就会报错（在非严格模式下，为了向下兼容，给未定义的变量赋值会创建一个全局变量）。
    
        在这个示例中，搜索过程如下：
            对于 name 变量，当 say 中的 alert 试图访问 name 时，会立即在内部词法环境中找到它。
            当它试图访问 phrase 时，然而内部没有 phrase，所以它顺着对外部词法环境的引用找到了它。
    */
    
    
    
    
    
    // -------------------------------------------------------------------------------------------
    
    
    
    
    // Step 4. 返回函数
    
    // makeCounter 这个例子
    function makeCounter() {
      let count = 0;
    
      return function() {
        return count++;
      };
    }
    
    let counter = makeCounter();
    
    // 在每次 makeCounter() 调用的开始，都会创建一个新的词法环境对象，以存储该 makeCounter 运行时的变量。
    // 因此，有两层嵌套的词法环境，
    // 不同的是，在执行 makeCounter() 的过程中创建了一个仅占一行的嵌套函数：return count++。我们尚未运行它，仅创建了它。
    
    // 所有的函数在“诞生”时都会记住创建它们的词法环境。从技术上讲，这里没有什么魔法：
        //- 所有函数都有名为 [[Environment]] 的隐藏属性，该属性保存了对创建该函数的词法环境的引用。
    
    // 现在，当 counter() 中的代码查找 count 变量时，它首先搜索自己的词法环境（为空，因为那里没有局部变量），
    // 然后是外部 makeCounter() 的词法环境，并且在哪里找到就在哪里修改。
    
    // 在变量所在的词法环境中更新变量。
    
    //如果我们调用 counter() 多次，count 变量将在同一位置增加到 2，3 等。
    
    
    // 闭包
        // 闭包 是指内部函数总是可以访问其所在的外部函数中声明的变量和参数，即使在其外部函数被返回（寿命终结）了之后。
        // 在 JavaScript 中，所有函数都是天生闭包的（只有一个例外，将在 "new Function" 语法 中讲到）。
    
        // 也就是说：JavaScript 中的函数会自动通过隐藏的 [[Environment]] 属性记住创建它们的位置，所以它们都可以访问外部变量。
    
    
    
    
    // -------------------------------------------------------------------------------------------
    
    
    
    // 垃圾收集
        //- 通常，函数调用完成后，会将词法环境和其中的所有变量从内存中删除。因为现在没有任何对它们的引用了。
        // 与 JavaScript 中的任何其他对象一样，词法环境仅在可达时才会被保留在内存中。
    
        // 但是，如果有一个嵌套的函数在函数结束后仍可达，则它将具有引用词法环境的 [[Environment]] 属性。
    
    
    // 在下面这个例子中，即使在函数执行完成后，词法环境仍然可达。因此，此嵌套函数仍然有效。
    function f() {
      let value = 123;
    
      return function() {
        console.log(value);
      }
    }
    
    let g = f(); // g.[[Environment]] 存储了对相应 f() 调用的词法环境的引用
    
    // 请注意，如果多次调用 f()，并且返回的函数被保存，那么所有相应的词法环境对象也会保留在内存中。
    // 下面代码中有三个这样的函数：
    let arr = [f(), f(), f()];
    
    // 当词法环境对象变得不可达时，它就会死去（就像其他任何对象一样）。换句话说，它仅在至少有一个嵌套函数引用它时才存在。
    
    // 在下面的代码中，嵌套函数被删除后，其封闭的词法环境（以及其中的 value）也会被从内存中删除：
    g = null;
    
    
    