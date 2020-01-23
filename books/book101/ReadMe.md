

## book101

《Go并发编程实战》 郝林

<https://github.com/hyper0x/go_command_tutorial> , ...

<https://github.com/hyper0x/goc2p> ， 本书源码

 本地: `D:\temp2\goc2p`

### 类型

**基本类型**也可以称为预定义类型

|序号|类型|
|---|---|
|1|bool|
|2|byte|
|3|rune|
|4|int/uint|
|5|int8/uint8|
|6|int16/uint16|
|7|int32/uint32|
|8|int64/uint64|
|9|float32|
|10|float64|
|11|complex64|
|12|complex128|

**复合类型**

|序号|类型|名称|
|---|---|---|
|1|Array|数组|
|2|Struct|结构体|
|3|Function|函数|
|4|Interface|接口|
|5|Slice|切片|
|6|Map|字典|
|7|Channel|通道|
|8|Pointer|指针|

### 自动解引用

如果x是与一个结构体类型对应的指针类型的值，那么表达式`x.f`就是表达式`(*x).f`的一个速记法。

### 类型断言

```
x.(T)
```

- `int(123).(int)` 会引发一个编译错误
- `interface{}(123).(int)` 应该这样写。`interface{}`是一个特殊的接口类型，代表空接口。所有类型都是它的实现类型。

```
v, ok := x.(T)
```

### 数组

```
[...]string{"Go", "Python", "Java", "C", "C++", "PHP"}
```

### 切片

- 一个切片值总会持有一个对某个数组值的引用。
- 事实上，一个切片值一旦被初始化，就会与一个包含了其中元素的数组值相关联。
- 这个数组值被称为引用它的切片值的底层数组。
- 多个切片值可能会共用同一个底层数组。
- 例如，如果我们把一个切片复制成多个，或者针对其中的某个连续片段再切片成新的切片值，那么这些切片值所引用的都会是同一个底层数组。
- 对切片值中的元素值的修改，实质上就是对其底层数组上的对应元素的修改。
- 从这个角度看，切片值类似于指向底层数组的指针。

-----
- 我们可以把切片值想象成朝向其底层数组的一个窗口。
- 这个窗口是我们查看底层数组中的元素值的途径。
- 这个值的长度就是我们当前可以看到的底层数组中的元素值的数量，而它的容量则表示了我们最多能够看到多少个当前底层数组中的元素值。

|序号|内建函数|
|---|---|
|1|`append`|
|2|`copy`|

### 字典

|序号|内建函数|
|---|---|
|1|`delete`|

### 函数

```
func Module(x, y int) (result int) {
    result = x % y
    return
}
```

### 结构体 | 嵌入类型

- 嵌入类型所附带的方法都会无条件地与被嵌入的结构体类型关联在一起，
- 即它们也成为了被嵌入的结构体类型的方法。
- 这意味着，结构体类型自动地实现了它包含的所有嵌入类型所实现的接口类型。
- 但是，请注意，嵌入类型的方法的接收者类型仍然是该嵌入类型，而不是那个被嵌入的结构体类型。
- 当我们在被嵌入的结构体上调用实际上属于其中某个嵌入类型的方法的时候，这一调用会被自动转发到这个嵌入类型的值上。

|需要强调的是|
|---|
|在Go语言中，只存在嵌入而不存在继承的概念。|

### 数据初始化

|序号|内建函数|
|---|---|
|1|`new`|
|2|`make`|

- `new`函数用于为值分配内存。但是与其他编程语言中的new不同的是，它并不会去初始化分配到的内存，而只会清零它。
- 调用表达式`new(T)`被求值时，所做的是为T类型的新值分配并清零一块内存空间，然后将这块内存空间的地址作为结果返回。而这个结果就是指向这个新的T类型值的指针值。它的类型为`*T`。
- `make`函数只能被用于创建切片类型、字典类型和通道类型的值，并返回一个已被初始化的对应类型的值。

### 返回多个结果的表达式

`4`种情况
- 返回多个结果的函数或方法。
    - 比如： `func Save(p Person) (id int, done bool)`
    - `id, done := Save(Person{})`
- 应用于字典值之上的索引表达式。
    - `v, ok := map["k1"]`
- 类型断言表达式。
    - `v, ok := x.(string)`
- 由接收操作符和通道类型值组成的表达式。
    - `v, ok := <-ch`

### iota

- D:\wxg104_Go\src\books\book101\ch03\iota_test.go

### 字面常量与变量类型的对应关系

- 当省略类型时，默认的类型。比如 `var v = 0`

|序号|字面常量的种类|变量的类型|
|---|---|---|
|1|布尔常量|`bool`|
|2|字符常量|`rune`|
|3|整数常量|`int`|
|4|浮点数常量|`float64`|
|5|复数常量|`complex128`|
|6|字符串常量|`string`|

### 短变量声明

- 我们在函数体内部声明变量的时候可以采用一种更加简短的声明方式，如：
- `v6 := true`
- 我们并不需要也不能显式地给定变量的类型。
- 也就是说，在这种声明中的变量的类型总是根据它的值推导得到的。
- 另外，短变量声明也不需要以`var`开始。
- 短变量声明可以出现在`if`, `for` 和 `switch` 等语句的初始化器中。

### switch

- 语句`switch`可以使用表达式或者类型说明符作为`case`判定方法。
- 因此，`switch`语句也就可以被分为两类
    - 表达式`switch`语句
        - `src\books\book101\ch04\switch_test.go`
    - 类型`switch`语句
        - `src\books\book101\ch04\switch2_test.go`
        - `fallthrough` 语句不允许出现在类型`switch`语句中的任何`case`语句列表中。

-------

### for

- `for`子句的三个部分是有固定排列顺序的。
    - 初始化子句在左、条件在中、后置子句在右。
    - 并且它们之间需要用分号`;`来分隔。
    - 可以省略掉其中的任何部分。
- `range`

-------

### golang不支持参数默认值

- 参考: <https://www.zhihu.com/question/24368980> , 知乎
- 参考: <https://cloud.tencent.com/developer/article/1453857> , 按这篇文章的说法，需要提供多个参数来代替默认值, 代码如下

```go
package default_arg_test
import "testing"
type Image struct{
	width, height int
    bgColor uint32
}
func NewImageWithBgColor(w, h int, bgColor uint32) *Image {
    return &Image{width: w, height: h, bgColor: bgColor}
}
func NewImage(w, h int) *Image {
    return NewImageWithBgColor(w, h, 0xFFFFFF)
}
func TestGetImage(t *testing.T) {
    image1 := NewImage(100, 100)
    image2 := NewImageWithBgColor(120, 250, 0xFF00FF)
    _, _ = image1, image2
}
```

-------

### panic

- 标准库代码中，往往都会有自己的error接口类型的实现。
    - 只有当调用recover函数得到的结果值类型是自己定义的error类型时，才会处理这个恐慌。
    - 否则就会重新引发一个运行时恐慌并携带相同的值。
- 标准库代码包fmt中的Token函数处理恐慌的代码如下：

```
func (s *ss) Token(skipSpace bool, f func(rune) bool) (tok []byte, err error) {
    defer func() {
        if e := recover(); e != nil {
            if se, ok := e.(scanError); ok {
                err = se.err
            } else {
                panic(e)
            }
        }
    }()
    // 省略若干条语句 
}
```

### HashSet

- 本地: `D:\temp2\goc2p`
- 官方: <https://github.com/hyper0x/goc2p> ， `src/basic/set/hash_set.go`

### 测试

测试文件必须以`_test.go`结尾

- 功能测试
    - 以`TestXxx`命名
- 基准测试
    - 没看懂
- 样本测试
    - 以`ExampleXxx`命名
    - 另外还可以添加形如`// Output: xxx`的注释行。
    - 命名规范还很多哦。`Example`, `ExampleF`, `ExampleT`, `ExampleT_M`, `ExampleT_M_basic`
- 收集资源使用情况
    - `-cpuprofile cpu.out`
    - `-memprofile mem.out`
    - `-memprofilerate n`
    - 比如: `go test -cpuprofile cpu.out net` 运行标准库的net代码包中的测试，记录CPU使用情况，并写入到cpu.out文件中。
- 记录程序阻塞事件
    - `-blockprofile block.out`
    - `-blockprofilerate n`
- 测试覆盖率
    - `-cover`
    - `-covermode=set` : `set`, `count`, `atomic`
    - `-coverpkg bufio,net`
    - `-coverprofile cover.out`

### 程序文档

- `godoc -http=:9090 -index`
    - 使用上面命令可以在本地搭建一个Web服务，用于查看本机所有工作区中的所有代码包文档。
- 程序注释： 类似java的单行注释和多行注释
- 代码包的注释
    - 当前代码包目录下的`doc.go`文件中。


### Goroutine

不要用共享内存的方式通信。作为替代，应该以通信作为手段来共享内存。

Go语言不推荐以共享内存区的方式传递数据。作为替代，我们应该优先使用Channel。

Channel主要被用来在多个Goroutine之间传递数据，并且还会保证其过程的同步。

|缩写|代表|含义|
|---|---|---|
|M|Machine|代表了一个内核线程。|
|P|Processor|代表了M所需的上下文环境。|
|G|Goroutine|代表了对一段需要被并发执行的Go语言代码的封装。|

**M、P和G的容器**

|中文名称|源码中的名称|作用域|简要说明|
|---|---|---|---|
|全局M列表|runtime.allm|运行时系统|存放所有M的列表|
|全局P列表|runtime.allp|运行时系统|存放所有P的列表|
|全局G列表|runtime.allg|运行时系统|存放所有G的列表|
|调度器的空闲M列表|runtime.sched.midle|调度器|存放空闲M的列表|
|调度器的空闲P列表|runtime.sched.pidle|调度器|存放空闲P的列表|
|调度器的可运行G队列|runtime.sched.runq|调度器|存放可运行G的队列|
|调度器的自由G列表|runtime.sched.gfree|调度器|存放自由G的列表|
|P的可运行G队列|runq|本地P|存放当前P中的可运行G的队列|
|P的自由G列表|gfree|本地P|存放当前P中的自由G的列表|

### Channel

- 基于通道的通讯是在多个Goroutine之间进行同步的重要手段。
- 针对通道的操作本身也是同步的。
- 在同一个时刻，仅有一个Goroutine能向一个通道发送元素值，同时也仅有一个Goroutine能从它那里接受元素值。
- 在通道中，各个元素值都是严格按照被发送至此的先后顺序排列的。
- 通道相当于一个FIFO（先进先出）的消息队列。
- 此外，通道中的元素值都具有原子性。它们是不可被分割的。通道中的每一个元素值都只可能被某一个Goroutine接收。已被接收的元素值会立刻被从通道中删除。
















