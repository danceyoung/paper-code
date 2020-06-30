# 面向包设计

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
