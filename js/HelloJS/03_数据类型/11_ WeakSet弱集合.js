// WeakSet

// WeakSet 的表现类似：

    //- 1. 与 Set 类似，但是我们只能向 WeakSet 添加对象（而不能是原始值）。
    //- 2. 对象只有在其它某个（些）地方能被访问的时候，才能留在 set 中。
    //- 3. 跟 Set 一样，WeakSet 支持 add，has 和 delete 方法，但不支持 size 和 keys()，并且不可迭代。

//- 变“弱（weak）”的同时，它也可以作为额外的存储空间。但并非针对任意数据，而是针对“是/否”的事实。
//  WeakSet 的元素可能代表着有关该对象的某些信息。


// 例如，我们可以将用户添加到 WeakSet 中，以追踪访问过我们网站的用户：
let visitedSet = new WeakSet();

let john = { name: "John" };
let pete = { name: "Pete" };
let mary = { name: "Mary" };

visitedSet.add(john); // John 访问了我们
visitedSet.add(pete); // 然后是 Pete
visitedSet.add(john); // John 再次访问

// visitedSet 现在有两个用户了

// 检查 John 是否来访过？
console.log(visitedSet.has(john)); // true

// 检查 Mary 是否来访过？
console.log(visitedSet.has(mary)); // false

john = null; 

console.log(visitedSet.has(john));   // false

console.log(visitedSet); // WeakSet { <items unknown> }
// visitedSet 将被自动清理
// WeakMap 和 WeakSet 最明显的局限性就是不能迭代，并且无法获取所有当前内容。
//那样可能会造成不便，但是并不会阻止 WeakMap/WeakSet 完成其主要工作 — 成为在其它地方管理/存储“额外”的对象数据。




// -------------------------------------------------------------------------------------------


// 总结
    //- WeakMap 是类似于 Map 的集合，它仅允许对象作为键，并且一旦通过其他方式无法访问它们，
    //  便会将它们与其关联值一同删除。

    // WeakSet 是类似于 Set 的集合，它仅存储对象，并且一旦通过其他方式无法访问它们，便会将其删除。

    // 它们都不支持引用所有键或其计数的方法和属性。仅允许单个操作。

    // WeakMap 和 WeakSet 被用作“主要”对象存储之外的“辅助”数据结构。一旦将对象从主存储器中删除，
    // 如果该对象仅被用作 WeakMap 或 WeakSet 的键，那么它将被自动清除。




// -------------------------------------------------------------------------------------------



// 存储 "unread" 标识

// 这里有一个 messages 数组：

let messages = [
  {text: "Hello", from: "John"},
  {text: "How goes?", from: "John"},
  {text: "See you soon", from: "Alice"}
];

// 你的代码可以访问它，但是 message 是由其他人的代码管理的。
// 该代码会定期添加新消息，删除旧消息，但是你不知道这些操作确切的发生时间。

// 现在，你应该使用什么数据结构来保存关于消息“是否已读”的信息？
// 该结构必须很适合对给定的 message 对象给出“它读了吗？”的答案。

// P.S. 当一个消息被从 messages 中删除后，它应该也从你的数据结构中消失。

// P.S. 我们不能修改 message 对象，例如向其添加我们的属性。
// 因为它们是由其他人的代码管理的，我们修改该数据可能会导致不好的后果。


// 让我们将已读消息存储在 WeakSet 中：


let readMessages = new WeakSet();

// 两个消息已读
readMessages.add(messages[0]);
readMessages.add(messages[1]);

// readMessages 包含两个元素

//  让我们再读一遍第一条消息！
readMessages.add(messages[0]);

// readMessages 仍然有两个不重复的元素

// 回答：message[0] 已读？
console.log("Read message 0: " + readMessages.has(messages[0])); // true

messages.shift();

// 现在 readMessages 有一个元素（技术上来讲，内存可能稍后才会被清理）
// WeakSet 允许存储一系列的消息，并且很容易就能检查它是否包含某个消息。

// 它会自动清理自身。代价是，我们不能对它进行迭代，也不能直接从中获取“所有已读消息”。
// 但是，我们可以通过遍历所有消息，然后找出存在于 set 的那些消息来完成这个功能。

// 另一种不同的解决方案可以是，在读取消息后向消息添加诸如 message.isRead=true 之类的属性。
// 由于 messages 对象是由另一个代码管理的，因此通常不建议这样做，但是我们可以使用 symbol 属性来避免冲突。

// 像这样：
// symbol 属性仅对于我们的代码是已知的
let isRead = Symbol("isRead");
messages[0][isRead] = true;

// 现在，第三方代码可能看不到我们的额外属性。

// 尽管 symbol 可以降低出现问题的可能性，但从架构的角度来看，还是使用 WeakSet 更好。






// -------------------------------------------------------------------------------------------




// 保存阅读日期

// 这儿有一个和 上一个任务 类似的 messages 数组。场景也相似。

/*
let messages = [
  {text: "Hello", from: "John"},
  {text: "How goes?", from: "John"},
  {text: "See you soon", from: "Alice"}
];
*/

// 现在的问题是：你建议采用什么数据结构来保存信息：“消息是什么时候被阅读的？”。

// 在前一个任务中我们只需要保存“是/否”。现在我们需要保存日期，并且它应该在消息被垃圾回收时也被从内存中清除。

// P.S. 日期可以存储为内建的 Date 类的对象

let readMap = new WeakMap();

readMap.set(messages[0], new Date(2021, 1, 25));





