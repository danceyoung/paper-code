# 面向包设计
原文 https://www.ardanlabs.com/blog/2017/02/package-oriented-design.html

非直译，有取舍
## 序
This post is part of a series of posts designed to make you think about your own design philosophy on different topics. If you haven’t read these posts yet, please do so first:
Develop Your Design Philosophy
Design Philosophy On Packaging

这篇文章仅仅是关于面向包设计理念中的一篇。如果你还没有看过下面这些内容，请先看看，有助于你理解本篇文章内容。

## 前言
Introduction
Package Oriented Design allows a developer to identify where a package belongs inside a Go project and the design guidelines the package must respect. It defines what a Go project is and how a Go project is structured. Finally, it improves communication between team members and promotes clean package design and project architecture that is discussable.
Package oriented design is NOT bound to a single project structure, but states that a project structure is paramount to applying guidelines for good package design. Moving forward, I will present one possible project structure and the guidelines to follow based on the design philosophies presented earlier.

面向包设计的理念让开发者在一个go项目中确定包的组织和必须要遵守的设计准则。它定义了一个go项目应该是什么样的及怎么架构和分层一个go项目。它最终的目的是为了提高项目的易懂性、整洁和可讨论性，便于团队成员沟通。

面向包设计不局限于项目本身的结构，更多为了表达一个实现合理面向包设计的项目结构是多么的重要。下面我将介绍一个面向包设计的项目和之前提到过的相关的准则。



## 项目结构
I believe that every company should establish a single Kit project and then multiple Application projects for the different sets of programs that get deployed together.

每个公司都会有一个工具包的项目和不同业务的应用项目
### 工具包项目
Kit Projects
Think of the Kit project as a company’s standard library, so there should only be one. The packages that belong to the Kit project need to be designed with the highest levels of portability in mind. These packages should be usable across multiple Application projects and provide a very specific but foundational domain of functionality. To this end, the Kit project is not allowed to have a vendor folder. If any of packages are dependent on 3rd party packages, they must always build against the latest version of those dependences.

考虑到工具包作为公司的一个标准类库，所以应该仅有一个。里面的所有包都需要设计为高级的可移植性。这些包可以在任何一个项目中都能使用，并且提供的都是很实用、具体的但又非常基础的功能。为了达到这样的目标，工具包项目不能有一个包依赖三方的vendor。因为如果有包依赖三方包，那就得不断的构建编译随着那些三方包的更新。

同时也不建议把工具包项目的部分包直接复制到你的应用项目中，因为这样本身增加了你对这些包管理、更新的工作，当然你如果真这样做也没毛病。