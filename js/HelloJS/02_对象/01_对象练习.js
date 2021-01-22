
// 按下面的要求写代码，一条对应一行代码：
    // -创建一个空的对象 user。
    // -为这个对象增加一个属性，键是 name，值是 小佛。
    // -再增加一个属性，键是 surname，值是 夏。
    // -把键为 name 的属性的值改成 小小佛。
    // -删除这个对象中键为 name 的属性。

let user = {};

user.name = "小佛";
user.surname = "夏";
console.log(`name: ${user.name} - surname: ${user.surname}`);

// 把键为 name 的属性的值改成 小小佛。
user.name = "小小佛";
console.log(`name: ${user.name} - surname: ${user.surname}`);

// 删除这个对象中键为 name 的属性
delete user.name;
console.log(`name: ${user.name} - surname: ${user.surname}`);




// 检查空对象
// 写一个 isEmpty(obj) 函数，当对象没有属性的时候返回 true，否则返回 false。

let schedule = {};

console.log( isEmpty(schedule) ); // true

schedule["8:30"] = "get up";

console.log( isEmpty(schedule) ); // false

function isEmpty(obj) {
    for (let key in obj) {
        // 如果进到循环里面，说明有属性。
        return false;
    }
    return true;
}



// 对象属性求和

// 我们有一个保存着团队成员工资的对象：

let salaries = {
  John: 100,
  Ann: 160,
  Pete: 130
};

// 写一段代码求出我们的工资总和，将计算结果保存到变量 sum。从所给的信息来看，结果应该是 390。
// 如果 salaries 是一个空对象，那结果就为 0。

function sumSalary(obj) {
    let sum = 0;
    for (let key in obj) {
        sum += obj[key];
    }
    return sum;
}

console.log(sumSalary(salaries)); // 390




// 将数值属性值都乘以 2

// 创建一个 multiplyNumeric(obj) 函数，把 obj 所有的数值属性值都乘以 2。

// 在调用之前
let menu = {
  width: 200,
  height: 300,
  title: "My menu"
};

multiplyNumeric(menu);


// 调用函数之后
menu = {
  width: 400,
  height: 600,
  title: "My menu"
};

console.log(menu.width);
console.log(menu.height);
console.log(menu.title);

// 注意 multiplyNumeric 函数不需要返回任何值，它应该就地修改对象。

// P.S. 用 typeof 检查值类型。
function multiplyNumeric(obj) {
    for (let key in obj) {
        if (typeof obj[key] === "number") {
            obj[key] *= 2;
        }
    }
}