# Go 包的组织和命名

> https://rakyll.org/style-packages/
>
> https://blog.golang.org/package-names
>
> 当看了很多不同地方的英文blog，发现知识的思想、原理还是来自官方的英文文档，就如中文搜索出来的知识，最终发现都是来自某篇文章，而某篇文章又是来自官方的内容，说这话的目的，是希望开发者要提高英文水平，不断研读官方的英文文献和不断实践。本篇文章特别对包的命名、组织，基于官方文档进行整理。



Go和其他语言一样，也涉及到命名和代码的组织。组织良好的代码具有很强的交流性、可读性和易用性，和设计良好的API一样重要。当其他用户阅读你代码的时候，包的位置、命名和结构分层的设计是他首先接触的东西。

本文通过比较通用且很好的实践例子来为你介绍Go包的规范，包括包的组织和包的命名，并不是一定要遵守的规则。实际中你要根据自己的需要来挑选最适合你项目的解决方案。

## 包的组织

所有的go代码都被组织在包中。一个包其实就是包含很多`.go`文件的一个路径，这样就实现了代码的组织和互相隔离，阅读代码从一个包开始。要想写好go代码，要不断理解包和基于包多实践。

### 使用多个文件

你要根据业务功能逻辑，充分地把代码分在不同的文件中，然后再根据整体的上下文放在同一个包中，这样才有更好的可读性。

比如，`http`包就是根据http的不同功能分成不同的对应文件，就像下面所示的四个文件一样。

```
- doc.go       // package documentation
- headers.go   // HTTP headers types and code
- cookies.go   // HTTP cookies types and code
- http.go      // HTTP client implementation, request and response types, etc.
```

### 类“以类聚“

一个小的规范，尽量把类型声明在被使用的文件中，这样就会让其他人根据功能更容易的找到对应的类型，而不是放在一个叫`model`的包中。

比如`header`的结构体应该声明在`header.go`文件中。

```
package http

// Header represents an HTTP header.
type Header struct {...}
```

另外注意到`Header`的结构体在`header.go`文件的最上面，golang本身并没有强性要求这样做，但把基础的核心的类型放在对应文件的最上面，是一个很好的做法。

### 按照功能职责设计

在golang中是按照功能职责设计的，即每个包都要有一定的真实业务含义。

但在其他语言中会把类型统一放在`model`包中，甚至只有类型的定义，这样我们很难找到某个类型是在哪里、哪些文件中被使用，且一旦修改了某个类型，我们不知道影响的范围。

所以在golang中我们不会如下这样设计的

```
package models // DON'T DO IT!!!

// User represents a user in the system.
type User struct {...}
```

在该示例中，我们应该把`User`声明在他被使用的包中，像下面这样

```
package mngtservice

// User represents a user in the system.
type User struct {...}

func UsersByQuery(ctx context.Context, q *Query) ([]*User, *Iterator, error)

func UserIDByEmail(ctx context.Context, email string) (int64, error)
```

### 提供例子来解释不清晰的引用包

有些情况下，你并不能在一个独立的包中提供所有要使用的类型。比如你想在一个包中实现一个接口，但是这个接口所在的包和你这个包不是同一个包，或者有些类型是在第三方包中，这样就会很不清晰。此时就需要提供例子来说明这些包是怎么一起用的。例子会提高那些不太容易被发现包的透明度。

### 不要在main包中导出任何标识符

一个标识符可以被导出（首字母大写），以允许其他包应用他。

但是`main`包不允许其他包导入，所以在`main`包导出标识符，没有任何意义。也有些例外，就是`main`包将要被构建入a .so, or a .a or Go plugin.

## 包的命名

一个好的包名称，会有很好的可读性，要见文识义，就是使用者能从包名里看出包的代码逻辑和使用意图。因为包的名称表达了包内容的上下文语意，让使用者很容易理解这个包是干什么的及怎么用的。对代码的维护者来说，也更易维护。

### 包应该被命名什么样

包的名称应该要小写，不要用下划线your_package或驼峰yourPackage的命名样式。且应该是一个简短的名词，比如 time，list，http

所以在其他编程语言中的命名规范可能就不适合Go了，比如 computeServiceClient，priority_queue这样的

如果一个单词无法表达包的含义，可以用多个单词，但需要简写每个单词，且简写后的名称对编程人员来说，都是耳闻能详的，如果不能，就不要这样做了。比如 strconv(string conversion)，syscall(system call)，fmt(formatted i/o)

同时也要尽量避免这样的单词，即编程人员经常用到的单词，比如编程人员声明一个Buffer类型的变量buf。

不必担忧包名和其他源码库冲突，唯一性不是必须的，包名仅仅是导入后默认使用的名称。虽然不太容易发生引入不同地方的同一个名称的包名，但如果真发生了冲突，你可以在导入时重命名一下包名，比如下面这样。

```bash
import (
  "interview/log"  //自己项目中的log包
  golog "log"     //为了避免冲突，可以把标准log包起个别名golog
)
```

### 包内容的命名

包的名字和包里面内容的名字是耦合的（包括类型、变量、常量、方法、函数等），是被使用者放在一起使用的，所以当你设计包的时候，尽量站在使用者的角度。

#### 避免重复或不清晰

使用者使用某个类型、函数等时，是把包名作为前缀的，比如 http.Server。这里http是包名，Server是类型名，所以没必要把Server命名成HTTPServer。直接使用http.Server更清晰，也不重复的表达。

#### 函数名

通常，如果在包pkg里有一个函数的返回值是pkg.Pkg(*pkg.Pkg)，那么这个函数的名称不应该含有Pkg字样，比如标准包的

```bash
start := time.Now()                                  // start is a time.Time
t, err := time.Parse(time.Kitchen, "6:06PM")         // t is a time.Time
ctx = context.WithTimeout(ctx, 10*time.Millisecond)  // ctx is a context.Context
ip, ok := userip.FromContext(ctx)                    // ip is a net.IP
```

包time的函数Now，Parse返回Time类型的值，就不应该写成NowTime, ParseTime,可以比较一下time.Now() 和time.NowTime()，很明显time.Now()更直接、简练。

如果包pkg有一个函数New返回pkg.Pkg，这个函数就是很好很标准的函数名字。

```bash
q := list.New()  // q is a *list.List
```

如果包pkg有一个函数返回的是类型T，而不是Pkg，这个函数名应该包含类型T的字样，这样对使用者来说更清晰明了。

```bash
d, err := time.ParseDuration("10s")  // d is a time.Duration
elapsed := time.Since(start)         // elapsed is a time.Duration
ticker := time.NewTicker(d)          // ticker is a *time.Ticker
timer := time.NewTimer(d)            // timer is a *time.Timer
```

不要担忧在不同的包里起了相同的类型名称，因为在使用某个类型的时候是需要包名作为前缀的，这样是不太会引起混淆或歧义的，但是也有尽量的避免。就像标准包的不同包里都有Reader类型的，jpeg.Reader, bufio.Reader, csv.Reader.

如果你起不好一个名称，可能你的逻辑边界是不清晰不准确的，你需要重新梳理你的业务需求、整个架构或代码的组织，直到让使用者或代码维护者能更容易理解和维护。

### 不好的包命名及解决方案

不好的命名让使用者很难理解这个包是什么，该怎么使用，同时也让代码维护者难于维护。

避免一个泛泛的名称，比如`common`,`utils`.

```
package util
func NewStringSet(...string) map[string]bool {...}
func SortStringSet(map[string]bool) []string {...}
```

使用者使用的时候，是这样的

```bash
set := util.NewStringSet("c", "a", "b")
fmt.Println(util.SortStringSet(set))
```

我们应该把这2个函数从util里分离出来单独成一个包，比如strset，这样更能表达这2个函数的意思。

```bash
package strset
func New(...string) map[string]bool {...}
func Sort(map[string]bool) []string {...}
```

此时使用者使用的时候，是这样的

```bash
set := stringset.New("c", "a", "b")
fmt.Println(stringset.Sort(set))
```

是不是很清晰了，然后还可以进一步提高，让代码看起来更简短

```bash
package strset
type Set map[string]bool
func New(...string) Set {...}
func (s Set) Sort() []string {...}
```

使用者也更简单了

```bash
set := stringset.New("c", "a", "b")
fmt.Println(set.Sort())
```

**不要把千万逻辑代码归一个包**。 有许多其他编程思想（或有惯性思维）的程序员会把一些逻辑上没相关，但属于某个架构层次的代码放在一块，比如都放在model，interface，api这样的包中。这样一个仅有一个的好处就是你能很快的找到代码的入口，但是这样的写法和上面util包没什么区别，随着代码量的增大，业务逻辑或代码逻辑的边界会越来越混乱，使用者也更难使用。此时需要按照业务的功能职责或纯技术的职责把那些包里的文件分离开来。

**避免复数的形式。** 在golang中，尽量避免以复数的形式命名，虽然go内置的包也有复数形式strings, errors。这样的规范会让其他开发语言的程序员感到很奇怪，但这就是go的设计哲学。比如不能命名为`httputils`,而要命名为`httputil`

```
package httputils  // DON'T DO IT, USE SINGULAR FORM!!
```

## 总结

在Go中，命名一个好的名称是很重要的一件事，希望开发者花时间来理解业务逻辑，组织代码结果，命名一个简短，简练，见文识义的名称。这样对使用者或代码维护者都是有好处的。
