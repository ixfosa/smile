// 创建实例时出错

// 这里有一份 Rabbit 扩展 Animal 的代码。

// 不幸的是，Rabbit 对象无法被创建。是哪里出错了呢？请解决它。

class Animal {

    constructor(name) {
        this.name = name;
    }

}

class Rabbit extends Animal {
    constructor(name) {
        // this.name = name;
        super(name);
        this.created = Date.now();
    }
}

let rabbit = new Rabbit("White Rabbit"); // Error: this is not defined
console.log(rabbit.name);




// -------------------------------------------------------------------------------------------





// 扩展 clock

// 我们获得了一个 Clock 类。到目前为止，它每秒都会打印一次时间。

class Clock {
    constructor({ template }) {
        this.template = template;
    }

    render() {
        let date = new Date();

        let hours = date.getHours();
        if (hours < 10) hours = '0' + hours;

        let mins = date.getMinutes();
        if (mins < 10) mins = '0' + mins;

        let secs = date.getSeconds();
        if (secs < 10) secs = '0' + secs;

        let output = this.template
        .replace('h', hours)
        .replace('m', mins)
        .replace('s', secs);

        console.log(output);
    }

    stop() {
        clearInterval(this.timer);
    }

    start() {
        this.render();
        this.timer = setInterval(() => this.render(), 1000);
    }
}

// 创建一个继承自 Clock 的新的类 ExtendedClock，并添加参数 precision — 每次 “ticks” 之间间隔的毫秒数，
// 默认是 1000（1 秒）。

// 你的代码应该在 extended-clock.js 文件里。
// 不要修改原有的 clock.js。请扩展它。

class ExtendedClock extends Clock {
    /*
    constructor(template, precision = 1000) {
        super(template);
        this.precision =  precision;
    }
    */

   constructor(options) {
    super(options);
    let { precision = 1000 } = options;
    this.precision =  precision;
}


    start() {
        super.render();
        this.timer = setInterval(() => super.render(), this.precision);
    }
}

// let eClock = new ExtendedClock({template: 'h:m:s'}, 5000);

let eClock = new ExtendedClock({template: 'h:m:s', precision: 5000});
eClock.start()

