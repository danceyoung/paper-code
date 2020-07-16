# 面向包的设计和架构分层

## 序

本篇内容主要讲解golang项目的面向包设计准则和基础的架构分层。

信息来自原文

* [Ardan Labs: Package-Oriented-Design](https://www.ardanlabs.com/blog/2017/02/package-oriented-design.html),
* [Github: golang standard project layout](https://github.com/golang-standards/project-layout),
* [Microsoft: Design Fundamentals - Layout Application Guideline](https://docs.microsoft.com/en-us/previous-versions/msp-n-p/ee658109(v=pandp.10))

内容进行翻译、加工、整合及结合个人的实践经验，并附有一个真实的例子来解释本篇内容。

* [group event](https://github.com/danceyoung/paper-code/tree/master/examples/groupevent)

当然你也可以直接阅读英文原文。

当然高手如云，只是懒得写罢了。

百年太久，只争朝夕，不负韶华，不枉少年，来日怎方长。

## 前

一个基本的go项目一般会有`cmd`, `internal`, `pkg`三个基础目录来分层，当然这不是官方`go`核心开发团队定义的标准。但这个确实是目前`go`生态系统中比较常见的布局形式，不管从之前的和还是现在开发项目的分层来看。这些基础目录同样适用更大的项目，并且还有一些小的增强功能。

如果你创建一个项目来学习go或你开发的是一个PoC或很小的项目，这种分层就没必要使用了，可能一个`main.go`文件就够了，即把数据、业务逻辑、规则、路由等等全部放在这个文件即可，也是所谓的**反模式**。但是随着业务不断变化而让你的项目也不断扩大，你就有必要考虑这种分层模式了，否则你会欠下凌乱无序、不易扩张、不可维护的**技术债**。当你有更多的团队成员开发同一个项目，就更需要更多的架构了，这个时候介绍一个常用的基本模式来管理包和库显得那么重要的原因了。如果你有一个开源的项目，并且别人会导入你项目的代码，这时你的项目中有一个私有（通常叫`internal`）的包是非常重要的，这时clone下来你的项目，就仅仅保留自己想要的，删除其他不用的包或代码，比如`internal`文件夹的内容。至于用到那些模式和目录，视你的项目情况而定，比如`vendor`就不是一定有的。

之前go项目的三方依赖包的管理最早有`vendor, go dep`等等，但都不是官方的，使用起来也不是尽善尽美，并不能像java项目maven那样的粒度管理三方依赖包。但随着go 1.14正式发布，`go modules`管理三方依赖包的工具也正式发布了。请尽量使用go modules, 除非你有一定不用他的理由。用go modules，你就不用关心GOPATH和非要把你的项目放在go workspace文件夹了。

这样的项目架构分层只是一种通用的模式，他不会给go的面向包的设计强加什么东西。面向包设计的理念让开发者在一个 go 项目中确定包的组织和必须要遵守的设计准则。它定义了一个 go 项目应该是什么样的及怎么架构和分层一个 go 项目。它最终的目的是为了提高项目的可读性、代码整洁性和可交流性，便于团队成员沟通。一个很好的大家都理解的架构本身就是一种通用的沟通语言。

面向包设计不局限于项目本身的结构，更多为了表达一个实现合理面向包设计的项目结构是多么的重要。下面将介绍一个面向包设计的项目、之前提到过的相关的准则和基础的架构分层。

## 项目架构分层

每个公司都会有一个工具包的项目和不同业务的应用项目

### 工具包项目

考虑到工具包作为公司的一个标准类库，所以应该仅有一个。里面的所有包都需要设计为高可移植性。这些包可以在任何一个项目中都能使用，并且提供的都是很实用、具体的但又非常基础的功能。为了达到这样的目标，工具包项目不能有一个包依赖三方的 vendor。因为如果有包依赖三方包，那就得不断的构建编译随着那些三方包的更新。

同时也不建议把工具包项目的部分包直接复制到你的应用项目中，因为这样本身增加了你对这些包管理、更新的工作，当然你如果真这样做也没毛病。

### 应用项目

应用项目是包含了很多需要部署在一起的程序集，包括服务、命令行工具和后台运行的程序。每个项目都对应一个含有其所有源代码的仓库，包括所有依赖的三方包。你需要几个应用项目，视情况以你而定，当然是越少越好。

每个应用项目通常包含三个根目录，分别是 `cmd, internal, pkg, vendor`。在 internal 文件里也会包含 `pkg` 目录，但是它和 internal 里其他的包有着不同的设计约束。

一个典型的应用项目结构应该是这样的：

```

paper-code/examples/groupevent
├── cmd/
│   └── eventtimer/
│       └── update/
│       └── main.go
│   └── eventserver/
│       └── router/
│           └── handler/
│           └── router.go
│       └── tests/
│       └── main.go
├── internal/
│   └── eventserver/
│       └── biz/
│           └── event/
│           └── member/
│       └── data/
│           └── service/
│   └── eventpopdserver/
│       └── event/
│       └── member/
│   └── pkg/
│       └── cfg/
│       └── db/
│       └── log/
└── vendor/
│   ├── github.com/
│   │   ├── ardanlabs/
│   │   ├── golang/
│   │   ├── prometheus/
│   └── golang.org/
├── go.mod
├── go.sum
```

#### cmd/

项目中的所有你将要编译成可执行程序的入口代码都放在`cmd/` 文件夹里，这些代码和业务没有关系。每个程序对应一个文件夹，文件夹的名称应该以程序的名称命名。一般在名称后面加上`d` 代表该程序是一个守护进程运行。
每个文件夹必须有一个`main`包的源文件，该源文件的名称也最好命名成可执行程序的名称，当然也可以保留main文件名。在此会导入和调用`internal/`和`pkg/`等其他文件夹中相关的代码。

示例

```
├── cmd/
│   └── eventtimer/
│       └── update/
│       └── main.go
│   └── eventserver/
│       └── router/
│           └── handler/
│           └── router.go
│       └── tests/
│       └── main.go
```

* 该项目包含线上业务服务eventserver（提供restful API）、定时器eventtimer（定时更新数据的状态）二个应用程序。`cmd`文件夹对应有2个文件夹，并且每个文件夹下面都有一个`main`包的源文件，至于名称可以直接用main，也可以对应文件夹的名称。
* 每个文件夹下的源文件里的代码和业务逻辑基本没任何关系。比如rest ful的eventserver，里面仅包含router的配置和相关的handler。

#### internal/

在go语言中，变量，函数，方法等的存取权限只有exported(全局)和unexported(包可见，局部)2种。

在项目中不被复用，也不能被其他项目导入，仅被本项目内部使用的代码包即私有的代码包都应该放在`internal`文件夹下。该文件夹下的所有包及相应文件都有一个项目保护级别，即其他项目是不能导入这些包的，仅仅是该项目内部使用。

如果你在其他项目中导入另一个项目的`internal`的代码包，保存或`go build` 编译时会报错`use of internal package ... not allowed`，该特性是在go 1.4版本开始支持的，编译时强行校验。

```
1 package main
2
3 import (
4	"paper-code/examples/groupevent/cmd/eventserver/router/handler"
5	"paper-code/examples/groupevent/cmd/internal"
6	"paper-code/examples/groupevent/internal/eventpopdserver/event"
7	"paper-code/examples/groupevent/pkg/middleware"
8 )
9
10 func main() {
11 	middleware.HandlerConv(nil)
12
13	event.EventsBy("")
14
15	eh := new(handler.EventHandler)
16	eh.Events(nil, nil)
17
18	internal.CmdInternalFunc()
19 }

```

> 此代码片段为另一个项目导入paper-code/example/groupevent的代码包

> 第6行的导入就会提示`use of internal package paper-code/examples/groupevent/internal/eventpopdserver/event not allowed`

> 第5行的导入也会提示同样的错误

> 第7行导入就可以的，因为导入的pkg代码包

当然你也不要局限根目录下的`internal`目录，你也可以在任何一个目录中创建`internal`，规则都适用。比如上面的例子`第5行的导入也会提示同样的错误:use of internal package paper-code/examples/groupevent/cmd/internal not allowed`

你可以在`internal`文件夹添加其他的架构分层目录来区分可分享、不可分享的代码，比如`internal/myapp`是你项目中某个程序的不可分享的代码；`internal/pkg/`是你项目中的程序都可以分享的代码。也可以添加数据层、业务逻辑层的代码，这个属于在项目中更通用的一个架构分层，和这里的包设计并不冲突，即上层模块可以直接访问下层模块，反之不然。

#### internal/pkg/

在同一个项目中不同程序需要访问，但又不能让其他项目访问的代码包，需要放在这里。这些包是比较基础但又提供了很特殊的功能，比如数据库、日志、用户验证等功能。

#### pkg/

如果你把代码包放在根目录的`pkg`下，其他项目是可以直接导入`pkg`下的代码包的，即这里的代码包是开放的，当然你的项目本身也可以直接访问的。但是如果你要把代码放在`pkg`下，还想需要三思而后行吧，有没必要这样做，毕竟`internal`目录是最好的方式保护你的代码并且被go编译器强制校验`internal`的代码包不可分享的。如果你的项目是一个开源的并且让其他人使用你封装的一些函数等，这样做是合适的，如果你自己或公司的某一个项目，个人的经验，基本上用不上`pkg`

#### vendor/

vendor文件夹包含了所有依赖的三方的源代码，它是go项目最早的依赖包的管理方式。目前大都用的go mod的依赖包管理，相对vendor，能指定版本，并且你不用特意手动下载更新依赖包，通过正常的go build, go run命令会自动处理。这样会减少项目本身的容量大小。

你可以用命令 `go mod vendor`来创建你项目的vendor目录。如果你项目中既要用到之前的vendor,又要用到go mod，你可以使用 `-mod=vendor`参数进行编译，但是在go1.14就不用了，当你用go build时，会自动检查项目根目录下有无vendor，并进行编译。

这里不过多介绍go mod的用法和特性。

## 面向包的设计和验证

面向包设计的准则可以验证项目中包设计的是否合理，下面这些步骤可以帮你发现包设计的问题。

### 包的位置

* `kit`
  被不同应用项目导入的基础包
* `cmd`
  支持编译不同二进制程序的包，比如Restful路由程序，需要相关router, handler包和main入口包。
* `internal`
  项目内部使用的包，包括crud, service(facade)和业务逻辑的包。
* `internal/pkg`
  为本项目内部使用的基础包，包括数据库、认证和序列化等操作的包。
* `pkg` 其他项目可以访问pkg的代码包

### 依赖包导入

* 根据业务合理设计包的粒度。
* 在一个包中导入另一个包中的类型，是不合适的。
  go源码里面的网络方面的`Request, Response, Header`等都在`http`包下面

  go的设计本身不建议建一个model模块，里面全是一个个结构体。因为这样设计，让其他人看代码，可能不知道这些结构体在哪被使用，修改了结构体，也不知道影响面有多大。
* 在同一个目录级别下的包互相导入，是不合适的。

  go更多是按照功能职责进行包的设计，所以同一目录级别下的包是不能互相导入的。除非你采用了在其他语言的架构分层是可以导入的，但也仅限上层可以导入下层的代码包，比如服务层、展现层、业务逻辑层、数据持久化层。

  ```
  ├── internal/
  │   └── eventserver/
  │       └── biz/
  │           └── event/
  │           └── member/
  │       └── data/
  │           └── service/
  │   └── eventpopdserver/
  │       └── event/
  │       └── member/
  ```


  > eventserver下的biz, data就是按照业务逻辑层、数据层这样的架构分层进行的设计。这样biz里面的代码包就可以导入同一目录级别下的data下的代码包，反之不然。
  >

  > eventpopdserver下的event, member是按照功能职责进行的设计，2者不能互相导入。
  >

  > 架构大致上分2种，一个就是通用分层（presenter, service, business, data ...）的架构分层，另一种就是按照功能职责进行分层，go倾向于后者。
  >
* 如果真有上面的需求

  请检查你对领域知识的理解、领域模型设计和包的设计。

  如果情非得已，那么将被导入的包移动到你的包里面。
* `cmd/`可以导入其他目录中的代码包。
* `internal/`中的包不能导入`cmd/`中的包。
* `internal/pkg/`中的包不能导入`cmd/`, `internal/`中的包。
* `pkg/`中的包不能导入`cmd/`, `internal/`中的包。

### 应用级别的策略

比如给restful api的handler写中间件、定时更新等策略。

在`Kit`, `internal/pkg/, pkg/`中是不允许写这些策略的，也不允许日志的打印，因为这些都是某种意义上共用通用的代码包。在这里数据库的配置、日志文件的配置应该和运行时环境的改变是松散耦合的，可以通过环境变量来修改配置。

在`cmd/`, `internal/`是可以写中间件和定时器等。

### 数据的发送和接收

* 在语意上要确定好一个类型发送和接受的方式，即值类型还是引用类型。
  比如golang的`http`包中的`Request`结构体，在http中是以引用类型使用的。可以查看`http`包下面的`server`源码，里面包含了各种用法，如果你想自己写路由，server的几个函数和类型是必须要用的，这里不过多介绍。
* 如果你用一个接口类型的变量接收一个返回值，则更多的目的应该是调用接口的方法即行为，而不是值本身。如果不是这样，请直接用具体的类型。

### 错误处理

错误处理包括错误信息的日志输出，分析和解决错误，并且保证程序能恢复如果发生了错误。

* `Kit`
  <br/>  <br/>  不允许使用`panic`终止程序或抛出错误。<br/>不允许再次包装错误信息，原本原样的把系统错误或框架的错误返回即可。<br/>  <br/>
* `cmd/`
  <br/>  <br/>  允许使用`panic`终止程序或抛出错误。<br/>如果有错误发生且不处理，可以根据此时的业务或逻辑上下文包装一下错误，让更上层的处理错误的函数能知道是哪里抛出的错误。<br/>当然大多数的错误都应该在这里处理。<br/>  <br/>
* `internal/`
  <br/>  <br/>  不允许使用`panic`终止程序或抛出错误。<br/>如果有错误发生且不处理，可以根据此时的业务或逻辑上下文包装一下错误，让更上层的处理错误的函数能知道是哪里抛出的错误。<br/>当然大多数的错误都应该在这里处理。<br/>  <br/>
* `internal/pkg/`
  <br/>  <br/>  不允许使用`panic`终止程序或抛出错误。<br/>
  不允许再次包装错误信息，原本原样的把系统错误或框架的错误返回即可。<br/>
* `pkg/`不允许使用`panic`终止程序或抛出错误。<br/>不允许再次包装错误信息，原本原样的把系统错误或框架的错误返回即可。<br/>

### 测试

* `cmd/`

  允许使用第三方的测试包。<br/>
  可以独立创建一个test包来管理单元测试的文件。<br/>
  这里更多是集成测试而不是单元测试。
* `kit/`, `internal/`, `internal/pkg/,pkg/`

  强烈推荐使用golang的testing包。<br/>
  test文件可以直接创建在对应包下面。<br/>
  这里更多是单元测试而不是集成测试。

### 捕获错误

* `cmd/`

  可以捕获任何错误，且保证程序100%能恢复。
* `kit/, internal/, internal/pkg/,pkg/`

  不能捕获错误，除非发生错误时，有对应的线程可以处理，或通知到程序。

## 不建议的目录

* `src/`

  src目录在java开发语言的项目中是一个常用的模式，但是在go开发项目中，尽量不要使用src目录。
* `model/`

  在其他语言开发中一个非常通用的模块叫model，把所有类型都放在model里。但是在go里不建议的，因为go的包设计是根据功能职责划分的。比如一个User 模型，应该声明在他被用的功能模块里。
* `xxs/`

  带复数的目录或包。虽然go源码中有strings包，但更多都是用单数形式。

## 结论

在实际go项目开发中，一定要灵活运用，当然也可以完全不按照这样架构分层、包设计的规则，一切以项目的大小、业务的复杂度、个人专业技能认知的广度和深度、时间的紧迫度为准。

最后以软件大师 [Kent Beck](https://baike.baidu.com/item/Kent%20Beck/13006051?fr=aladdin) 在《重构Refactoring》一书中描述的结尾。

* 先让代码工作起来-如果代码不能工作，就不能产生价值
* 然后再试图将它变好-通过对代码进行重构，让我们自己和其他人更好地理解代码，并能按照需求不断地修改代码。
* 最后再试着让它运行得更快-按照性能提升的需求来重构代码。

谢谢
