# Go 包的命名和组织

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

有些情况下，你并不能在一个独立的包中提供所有要使用的类型。比如你想在一个包中实现一个接口，但是这个接口所在的包和你这个包不是同一个包，或者有些类型是在第三方包中，这样就会很不清晰。此时就需要提供例子来说明这些包是怎么一起用的。

如果你的代码中引用来一些不标准的包，通常需要提供一些代码例子来说明一下。例子会提高那些不太容易被发现包的透明度。

### 不要在main包中导出任何标识符

一个标识符可以被导出（首字母大写），以允许其他包应用他。

但是`main`包不允许其他包导入，所以在`main`包导出标识符，没有任何意义。也有些例外，就是`main`包将要被构建入a .so, or a .a or Go plugin.

## 包的命名

包的名称和导入路径是你包的重要标识，并且它代表了你的包包含的所有东西。所以包的命名一定要“见文识意“，因为这样不仅提高了代码的质量，也为你的用户提高了代码的可读性。

### 名称要小写

包的名称应该要小写，不要用下划线your_package或驼峰yourPackage的命名样式。更详细的介绍可以参考[官方博客](https://blog.golang.org/package-names)，专门介绍包的命名规则。

### 精、简、短、见文识意

包名应该精简短，且名称有唯一性并能见文识意。就是你的用户能从包名里看出包的代码逻辑和使用意图。

避免一个泛泛的名称，比如`common`,`utils`.

```
import "pkgs.org/common" // DON'T!!!
```

避免命名重复，避免用户引用里同一个名称的包，但却是2个不同的包，这样容易产生混淆。

如果你起不好一个名称，有可能从业务需求、整个架构或代码的组织层面就出现了问题。


coming soon...
