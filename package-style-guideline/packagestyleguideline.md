# Go 包的规范

原文：https://rakyll.org/style-packages/

Go和其他语言一样，也涉及到命名和代码的组织。组织良好的代码具有很强的交流性、可读性和易用性，和设计良好的API一样重要。当其他用户阅读你代码的时候，包的位置、命名和结构分层的设计是他首先接触的东西。

本文通过比较通用且很好的实践例子来为你介绍Go包的规范，并不是一定要遵守的规则。实际中你要根据自己的需要来挑选最适合你项目的解决方案。

## 包

所有的go代码都被组织在包中。一个包其实就是包含很多`.go`文件的一个路径，这样就实现了代码的组织和互相隔离，阅读代码从一个包开始。要想写好go代码，要不断理解包和基于包多实践。

## 包的组织

我们从这2方面的建议开始：怎么组织go代码和包位置设计的约定。

### 使用多个文件

一个包是包含了很多`.go`文件的路径。你要根据功能逻辑，充分地把代码分在不同的文件中，这样才有更好的可读性。

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
