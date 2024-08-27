---
title: Interfaces
tags: [Go]
draft: true
---

Go 語言的介面在設計上跟許多物件導向程式語言的介面用法有很大的不同，可能是這個原因，使得介面經常被誤用或濫用。

Rob Pike（Go 語言的其中一名主要開發者）曾說過：

> Don’t design with interfaces, discover them.

意思是，只有當我們發現這裡或那裡需要一個介面會比較好的時候，才去定義介面。這可以避免我們太早傷腦筋去設計「想像中的」介面，也能避免過度設計。

- 太早或者過度使用介面，很容易會加入一堆用處不大的抽象層，讓程式更複雜、更難維護。
- 如果沒有強烈的理由、或者不大確定增加一個介面能夠帶來明顯好處，就應該再三斟酌。
- 不要太擔心直接呼叫實作會造成什麼嚴重後果。與其用更多抽象層來應對未來可能的狀況，不如先解決眼下的需求。

## 介面該放在哪一邊？ {#where-should-interface-live}

關於介面要定義在哪裡，基本上有兩種做法：

- **Producer side** - An interface defined in the same package as the concrete implementation.
- **Consumer side** - An interface defined in an external package where it’s used

熟悉 C# 或 Java 的人通常會把介面定義在生產端（服務端）。然而，這跟 Go 的設計哲學大相逕庭。

一般而言，Go 的介面大多應該由用戶端來定義。這是指大多數的情況，但也有少部分特例是把介面寫在服務端。



