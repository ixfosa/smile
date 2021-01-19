/**
    "switch" 语句
        - switch 语句可以替代多个 if 判断。
        - switch 语句为多分支选择的情况提供了一个更具描述性的方式。

    语法
        - switch 语句有至少一个 case 代码块和一个可选的 default 代码块。

        switch(x) {
        case 'value1':  // if (x === 'value1')
            ...
            [break]

        case 'value2':  // if (x === 'value2')
            ...
            [break]

        default:
            ...
            [break]
        }

    - 比较 x 值与第一个 case（也就是 value1）是否严格相等，然后比较第二个 case（value2）以此类推。
    - 如果相等，switch 语句就执行相应 case 下的代码块，直到 switch 语句末尾。
    - 如果没有符合的 case，则执行 default 代码块（如果 default 存在）。
 */


let a = 2 + 2;
switch (a) {
    case 3:
        console.log( 'Too small' );
        break;
    case 4:
        console.log( 'Exactly!' ); // Exactly!
        break;
    case 5:
        console.log( 'Too large' );
        break;
    default:
        console.log( "I don't know such values" );
}


// 如果没有 break，程序将不经过任何检查就会继续执行下一个 case。
let b = 2 + 2;
switch (b) {
    case 3:
        console.log( 'Too small' );
    case 4:
        console.log( 'Exactly!' );
    case 5:
        console.log( 'Too big' );
    default:
        console.log( "I don't know such values" );
}
// res: 
    //- Exactly!
    //- Too big
    //- I don't know such values



// 任何表达式都可以成为 switch/case 的参数
    //- switch 和 case 都允许任意表达式。
a = "1";
b = 0;
switch (+a) {
    case b + 1:
        console.log("this runs, because +a is 1, exactly equals b+1");
        break;
    default:
        console.log("this doesn't run");
}
// 这里 +a 返回 1，这个值跟 case 中 b + 1 相比较，然后执行对应的代码



// “case” 分组
    //- 共享同一段代码的几个 case 分支可以被分为一组：

// 比如，如果我们想让 case 3 和 case 5 执行同样的代码：
a = 3;
switch (a) {
  case 4:
    console.log('Right!');
    break;

  case 3: // (*) 下面这两个 case 被分在一组
  case 5:
    console.log('Wrong!');
    console.log("Why don't you take a math class?");
    break;
    // Wrong!
    // Why don't you take a math class?

  default:
    console.log('The result is strange. Really.');
}


// 类型很关键
    //- 强调一下，这里的相等是严格相等。被比较的值必须是相同的类型才能进行匹配。

// 比如，我们来看下面的代码：
let arg = '3'
switch (arg) {
  case '0':
  case '1':
    console.log( 'One or zero' );
    break;

  case '2':
    console.log( 'Two' );
    break;

  case 3:
    console.log( 'Never executes!' );
    break;
  default:
    console.log( 'An unknown value' ) // An unknown value  严格相等 
}