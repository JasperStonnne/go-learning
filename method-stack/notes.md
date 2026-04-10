# Go 方法与接口学习笔记

## 一、用指针接收者写一个栈

### 结构体定义

```go
type Stack struct {
    items []int
}
```

- `Stack` 是栈本身（盒子），`items` 是栈里存数据的地方（盒子里的东西）
- 定义结构体时 `items []int` 只是模板，不需要关键字创建变量

### 创建实例

```go
myStack := Stack{}
```

用 `Stack` 模板创建一个实例，`{}` 里是初始值，现在为空，`items` 默认是空切片。

---

## 二、方法结构拆解

```go
func  (s *Stack)  Push  (val int)  {
 ↑        ↑        ↑       ↑
函数关键字  接收者   方法名   参数
```

### 为什么 Push 用指针接收者

`Push` 需要修改 `Stack` 内部的 `items`，必须用指针接收者：

- 值接收者 `(s Stack)`：修改的是副本，原栈不变
- 指针接收者 `(s *Stack)`：直接操作原栈，修改生效

### *Stack 和 s.items 的关系

`(s *Stack)` 里 `*` 表示"直接操作原来那个栈，不复制一份"。用的时候写 `s.items` 即可，Go 自动处理 `*` 的细节。

---

## 三、输入参数 vs 返回值

```go
func (s *Stack) Push(val int)      // val int 是输入参数，调用时传给它
func (s *Stack) Pop() (int, bool)  // (int, bool) 是返回值，执行完返回两个结果
```

- 没有返回值 → 后面什么都不写（Push）
- 多个返回值 → 用括号包起来（Pop、Peek）

---

## 四、切片操作

### 添加元素

```go
s.items = append(s.items, val)
```

### 删除最后一个元素（不用 append）

```go
s.items = s.items[:len(s.items)-1]
```

### [:n] 和 [n] 的区别

```go
s.items[:len(s.items)-1]  // 取前面所有元素，结果是切片
s.items[len(s.items)-1]   // 取最后那一个元素，结果是单个值
```

---

## 五、其他语法

### 不同函数里同名变量互不影响

`Push` 里的 `val` 和 `Pop` 里的 `val` 是两个独立的变量，只在各自函数内存在。

### _ 占位符

Go 要求返回值必须接收，不需要某个值时用 `_` 占位：

```go
val, _ := myStack.Peek()  // 只要第一个返回值，忽略第二个
```

---

## 六、接口

### 什么是接口

接口定义行为规范，不提供具体实现。任何类型只要实现了接口里的所有方法，就自动实现了该接口，不需要 `implements` 关键字。

### 例子

```go
type Animal interface {
    Speak()
}

type Dog struct{ name string }
func (d Dog) Speak() { fmt.Println("汪汪") }

type Cat struct{ name string }
func (c Cat) Speak() { fmt.Println("喵喵") }

// 函数只依赖接口，不依赖具体类型
func MakeSpeak(a Animal) {
    a.Speak()
}
```

新增 `Bird` 只需实现 `Speak()`，`MakeSpeak` 一行不用改。

### 接口的意义

函数只依赖行为，不依赖具体类型。新增类型不需要修改已有函数，代码更灵活。

### 注意：接口要用指针赋值

方法用了指针接收者，实现接口时必须传指针：

```go
var s Stacker = &myStack  // 正确
var s Stacker = myStack   // 编译错误
```
