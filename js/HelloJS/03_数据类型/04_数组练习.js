// 数组被拷贝了吗?

// 下面的代码将会显示什么？
    // - 这是因为数组是对象。所以 shoppingCart 和 fruits 是同一数组的引用。
let fruits = ["Apples", "Pear", "Orange"];

// 在“副本”里 push 了一个新的值
let shoppingCart = fruits;
shoppingCart.push("Banana");

// fruits 里面是什么？
console.log( fruits.length ); // 4






// 数组操作。
    //- 创建一个数组 styles，里面存储有 “Jazz” 和 “Blues”。
    //- 将 “Rock-n-Roll” 从数组末端添加进去。
    //- 用 “Classics” 替换掉数组最中间的元素。查找数组最中间的元素的代码应该适用于任何奇数长度的数组。
    //- 去掉数组的第一个值并显示它。
    //- 在数组前面添加 Rap 和 Reggae。

    //过程中的数组：
        //- Jazz, Blues
        //- Jazz, Bues, Rock-n-Roll
        //- Jazz, Classics, Rock-n-Roll
        //- Classics, Rock-n-Roll
        //- Rap, Reggae, Classics, Rock-n-Roll

let styles = ["Jazz", "Blues"];
console.log( styles ); // Jazz, Blues

styles.push("Rock-n-Roll") 
console.log( styles );  // Jazz, Bues, Rock-n-Roll

styles[1] = "Classics";
console.log( styles ); // Jazz, Classics, Rock-n-Roll

styles.shift();
console.log( styles ); // Classics, Rock-n-Roll

styles.unshift("Rap", "Reggae");
console.log( styles ); // Rap, Reggae, Classics, Rock-n-Roll





// 在数组上下文调用

// 结果是什么？为什么？

let arr = ["a", "b"];

arr.push(function() {
  console.log( this );
})

arr[2](); // [ 'a', 'b', [Function (anonymous)] ]

// arr[2]() 调用从句法来看可以类比于 obj[method]()，与 obj 对应的是 arr，与 method 对应的是 2。
// 所以调用 arr[2] 函数也就是调用对象函数。自然地，它接收 this 引用的对象 arr 




// 输入数字求和

// 写出函数 sumInput()，要求如下：
    //- 使用 prompt 向用户索要值，并存在数组中。
    //- 当用户输入了非数字、空字符串或者点击“取消”按钮的时候，问询结束。
    //- 计算并返回数组所有项之和。
    //- P.S. 0 是有效的数字，不要因为是 0 就停止问询。
function sumInput() {

    let numbers = [];

    while (true) {

        let value = prompt("A number please?", 0);

        // 应该结束了吗？
        if (value === "" || value === null || !isFinite(value)) break;

        numbers.push(+value);
    }

    let sum = 0;
    for (let number of numbers) {
        sum += number;
    }
    return sum;
}

alert( sumInput() );