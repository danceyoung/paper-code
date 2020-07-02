# 面向包设计

原文 https://www.ardanlabs.com/blog/2017/02/package-oriented-design.html

非直译，有取舍

## 序

This post is part of a series of posts designed to make you think about your own design philosophy on different topics. If you haven’t read these posts yet, please do so first:
Develop Your Design Philosophy
Design Philosophy On Packaging

这篇文章仅仅是关于面向包设计理念中的一篇。如果你还没有看过下面这些内容，请先看看，有助于你理解本篇文章内容。

## 前

Introduction
Package Oriented Design allows a developer to identify where a package belongs inside a Go project and the design guidelines the package must respect. It defines what a Go project is and how a Go project is structured. Finally, it improves communication between team members and promotes clean package design and project architecture that is discussable.
Package oriented design is NOT bound to a single project structure, but states that a project structure is paramount to applying guidelines for good package design. Moving forward, I will present one possible project structure and the guidelines to follow based on the design philosophies presented earlier.

面向包设计的理念让开发者在一个 go 项目中确定包的组织和必须要遵守的设计准则。它定义了一个 go 项目应该是什么样的及怎么架构和分层一个 go 项目。它最终的目的是为了提高项目的易懂性、整洁和可讨论性，便于团队成员沟通。

面向包设计不局限于项目本身的结构，更多为了表达一个实现合理面向包设计的项目结构是多么的重要。下面我将介绍一个面向包设计的项目和之前提到过的相关的准则。

## 项目结构

I believe that every company should establish a single Kit project and then multiple Application projects for the different sets of programs that get deployed together.

每个公司都会有一个工具包的项目和不同业务的应用项目

### 工具包项目

Kit Projects
Think of the Kit project as a company’s standard library, so there should only be one. The packages that belong to the Kit project need to be designed with the highest levels of portability in mind. These packages should be usable across multiple Application projects and provide a very specific but foundational domain of functionality. To this end, the Kit project is not allowed to have a vendor folder. If any of packages are dependent on 3rd party packages, they must always build against the latest version of those dependences.

考虑到工具包作为公司的一个标准类库，所以应该仅有一个。里面的所有包都需要设计为高可移植性。这些包可以在任何一个项目中都能使用，并且提供的都是很实用、具体的但又非常基础的功能。为了达到这样的目标，工具包项目不能有一个包依赖三方的 vendor。因为如果有包依赖三方包，那就得不断的构建编译随着那些三方包的更新。

同时也不建议把工具包项目的部分包直接复制到你的应用项目中，因为这样本身增加了你对这些包管理、更新的工作，当然你如果真这样做也没毛病。

### 应用项目

Application projects contain the set of programs that get deployed together. The set of programs can include services, cli tooling and background programs. Each Application project is bound to a single repo that contains all the source code for that project, including all the source code for the 3rd party dependencies. How many Application projects you need is up to you, but always take a less is more approach.

Each Application project contains three root level folders. These are cmd/, internal/ and vendor/. There is also a platform/ folder inside of the internal/ folder, which has different design constraints from the other packages that live inside of internal/.

应用项目是包含了很多需要部署在一起的程序集，包括服务、命令行工具和后台运行的程序。每个项目都对应一个含有其所有源代码的仓库，包括所有依赖的三方包。你需要几个应用项目，视情况以你而定，当然是越少越好。

每个应用项目通常包含三个根文件夹，分别是 cmd，internal，vendor。在 internal 文件里也会包含 platform 文件夹，但是它和 internal 里其他的包有着不同的设计约束。

A typical Application project might look like this:

一个典型的应用项目结构应该是这样的：

```

github.com/servi-io/api
├── cmd/
│   ├── servi/
│   │   ├── cmdupdate/
│   │   ├── cmdquery/
│   │   └── servi.go
│   └── servid/
│       ├── routes/
│       │   └── handlers/
│       ├── tests/
│       └── servid.go
├── internal/
│   ├── attachments/
│   ├── locations/
│   ├── orders/
│   │   ├── customers/
│   │   ├── items/
│   │   ├── tags/
│   │   └── orders.go
│   ├── registrations/
│   └── platform/
│       ├── crypto/
│       ├── mongo/
│       └── json/
└── vendor/
    ├── github.com/
    │   ├── ardanlabs/
    │   ├── golang/
    │   ├── prometheus/
    └── golang.org/
```

#### vendor/

Good documentation for the `vendor/` folder can be found in this Gopher Academy [post](https://blog.gopheracademy.com/advent-2015/vendor-folder) by Daniel Theophanes. For the purpose of this post, all the source code for 3rd party packages need to be vendored (or copied) into the `vendor/` folder. This includes packages that will be used from the company `Kit` project. Consider packages from the `Kit` project as 3rd party packages.

vendor 文件夹包含了所有依赖的三方的源代码，它是 go 项目最早的依赖包的管理方式。目前大都用的 go mod 的依赖包管理，相对 vendor，能指定版本，并且你不用特意手动下载更新依赖包，通过正常的 go build, go run 命令会自动处理。这样会减少项目本身的容量大小。

> 这里不过多介绍 go mod 的用法和特性。
