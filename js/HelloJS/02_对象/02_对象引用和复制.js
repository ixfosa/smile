// 对象引用和复制
    //- 与原始类型相比，对象的根本区别之一是对象是“通过引用”被存储和复制的，
    //- 与原始类型值相反：字符串，数字，布尔值等 —— 始终是以“整体值”的形式被复制的。



// 从原始类型开始，例如一个字符串。
// 将 message 复制到 phrase：
let message = "Hello!";
let phrase = message;
//结果我们就有了两个独立的变量，每个都存储着字符串 "Hello!"。


// 赋值了对象的变量存储的不是对象本身，而是该对象“在内存中的地址”，换句话说就是对该对象的“引用”。
let user = {
    name: "ixfosa"
};

// 当一个对象变量被复制 —— 引用则被复制，而该对象并没有被复制。
let admin = user; // 复制引用

// 这里仍然只有一个对象，现在有两个引用它的变量。
// 可以通过其中任意一个变量来访问该对象并修改它的内容：
admin.name = "梓晴";
console.log(user.name); // 梓晴




// -------------------------------------------------------------------------------------------



// 通过引用来比较
    //- 仅当两个对象为同一对象时，两者才相等。

// 例如，这里 a 和 b 两个变量都引用同一个对象，所以它们相等：
let a = {};
let b = a; // 复制引用

console.log( a == b ); // true，都引用同一对象
console.log( a === b ); // true

// 而这里两个独立的对象则并不相等，即使它们看起来很像（都为空）：
let a2 = {};
let b2 = {}; // 两个独立的对象

console.log( a2 == b2 ); // false

// 对于类似 obj1 > obj2 的比较，或者跟一个原始类型值的比较 obj == 5，对象都会被转换为原始值。



// -------------------------------------------------------------------------------------------



// 克隆与合并，Object.assign
    //- 拷贝一个对象变量会又创建一个对相同对象的引用。
    //- 但是，如果我们想要复制一个对象，那该怎么做呢？创建一个独立的拷贝，克隆？


// 但是，如果我们真的想要这样做，那么就需要创建一个新对象，并通过遍历现有属性的结构，
// 在原始类型值的层面，将其复制到新对象，以复制已有对象的结构。

// 就像这样：
let user2 = {
    name: "ixfosa",
    age: 30
  };
  
  let clone = {}; // 新的空对象
  
  // 将 user 中所有的属性拷贝到其中
  for (let key in user2) {
    clone[key] = user2[key];
  }
  
  // 现在 clone 是带有相同内容的完全独立的对象
  clone.name = "xfos"; // 改变了其中的数据
  
  console.log( user2.name ); // 原来的对象中的 name 属性依然是 ixfosa


// 也可以使用 Object.assign 方法来达成同样的效果。
// 语法是：
    //- Object.assign(dest, [src1, src2, src3...])
        //- 第一个参数 dest 是指目标对象。
        //- 更后面的参数 src1, ..., srcN（可按需传递多个参数）是源对象。
        //- 该方法将所有源对象的属性拷贝到目标对象 dest 中。换句话说，从第二个开始的所有参数的属性都被拷贝到第一个参数的对象中。
        //- 调用结果返回 dest。

let user3 = { name: "ixfosa" };

let permissions1 = { canView: true };
let permissions2 = { canEdit: true };

// 将 permissions1 和 permissions2 中的所有属性都拷贝到 user 中
Object.assign(user3, permissions1, permissions2);

// 现在 user = { name: "ixfosa", canView: true, canEdit: true }
console.log(`name:${user3.name}, canView:${user3.canView}, canEdit:${user3.canEdit}`);

// 如果被拷贝的属性的属性名已经存在，那么它会被覆盖：
Object.assign(user3, {name: "xfos"});
// name:xfos, canView:true, canEdit:true
console.log(`name:${user3.name}, canView:${user3.canView}, canEdit:${user3.canEdit}`);


// 可以用 Object.assign 代替 for..in 循环来进行简单克隆：
let user4 = {
  name: "ixfsoa",
  age: 30
};

let clone2 = Object.assign({}, user4);
console.log(`name:${clone2.name}, age:${clone2.age}`); //name:ixfsoa, age:30



// -------------------------------------------------------------------------------------------




// 深层克隆
    //- 到现在为止，我们都假设 user 的所有属性均为原始类型。但属性可以是对其他对象的引用。那应该怎样处理它们呢？

let user5 = {
    name: "John",
    sizes: {
        height: 182,
        width: 50
    }
};
    
console.log( user5.sizes.height ); // 182
// 拷贝 clone.sizes = user.sizes 已经不足够了，因为 user.sizes 是个对象，它会以引用形式被拷贝。
// 因此 clone 和 user 会共用一个 sizes：


let clone5 = Object.assign({}, user5);

console.log( user5.sizes === clone5.sizes ); // true，同一个对象

// user 和 clone 分享同一个 sizes
user5.sizes.width++;       // 通过其中一个改变属性值
console.log(clone5.sizes.width); // 51，能从另外一个看到变更的结果

// 应该使用会检查每个 user[key] 的值的克隆循环，如果值是一个对象，那么也要复制它的结构。这就叫“深拷贝”。





// -------------------------------------------------------------------------------------------



//总结
    //- 对象通过引用被赋值和拷贝。换句话说，一个变量存储的不是“对象的值”，而是一个对值的“引用”（内存地址）。
    //- 因此，拷贝此类变量或将其作为函数参数传递时，所拷贝的是引用，而不是对象本身。

    //- 所有通过被拷贝的引用的操作（如添加、删除属性）都作用在同一个对象上。

    //- 为了创建“真正的拷贝”（一个克隆），我们可以使用 Object.assign 来做所谓的“浅拷贝”（嵌套对象被通过引用进行拷贝）
    //- 或者使用“深拷贝”函数，例如 _.cloneDeep(obj)。 https://lodash.com/docs/4.17.15



