---
title: 認識 Tailwind CSS
slug: introduction-to-tailwind-css
date: 2024-06-26
---

Hugo 昨天（2024-06-25）發布了 [v0.128.0](https://github.com/gohugoio/hugo/releases/tag/v0.128.0)，其 release note 提到這次加入了 TailwindCSS v4 (alpha)。我對前端技術相當陌生，以前聽過 Tailwind，但從未花時間去了解它。由於我的文件網站和部落格都使用 Hugo，所以這次花了一點時間去了解什麼是 Tailwind CSS，以及它有什麼特別之處。

首先，我從 [Tailwind CSS 官網](https://tailwindcss.com/)看到這句話：

> "Best practices" don't actually work.

這引發了我的好奇心，便順著那句話底下提供的連結點進去這篇文章：

[CSS Utility Classes and "Separation of Concerns"](https://adamwathan.me/css-utility-classes-and-separation-of-concerns/)

該文的發表日期是 2017 年 8 月 7 日，作者便是 Tailwind CSS 的創始人 Adam Wathan。我仔細讀完後，覺得這篇文章清楚闡述了 Tailwind CSS 的由來和設計理念，頗有收穫。

作者用一個實際的例子來解釋他是如何一步一步地降低 HTML 和 CSS 之間的耦合，最終達到他覺得比較理想的樣子：許多精心規劃和命名的 CSS classes，方便以組合的方式重複使用樣式；除了能讓版面風格趨向一致，也能減少日後維護 CSS 的成本、避免 CSS 繼承造成的混亂。

## "Separation of Concerns" 是個迷思

作者在文中提到，"Separation of Concerns" 在這裡是個迷思（原文為 straw man，即稻草人謬誤）。這句話僅只是針對 HTML 與 CSS 之間的依賴關係來說的。他的意思是，SoC 並不適合用於思考 HTML 與 CSS 之間的關係，否則容易陷入非黑即白的陷阱——只要沒有符合 SoC 原則就會被認定是糟糕的設計。

作者認為應該要思考的是**依賴的方向**（dependency direction）。有兩個方向：

1. **CSS 依賴 HTML**<br />這種設計方式，HTML 裡面主要是文件內容，至於樣式的部分，則會基於 HTML 的文件結構來撰寫 CSS 類別。最終的成品傾向於讓 HTML 盡量單純（很少 CSS），以便只要更換一整套的 CSS 就能切換網頁的風格樣式。這種設計大抵可以宣稱符合 Separation of Concerns 原則。
2. **HTML 依賴 CSS**<br />撰寫 CSS 的時候，樣式名稱會以比較通用的「功能」來命名，例如 `.btn`、`.btn-primary`、`.card` 等等，而不會涉及特定內容或領域知識（例如 `.author-form`）。

Bootstrap 和 Tailwind CSS 的設計都屬於第二種，也就是 HTML 依賴 CSS。

## 延伸閱讀

我也一併讀了文中引用的其他文章，其中有兩篇我覺得也值得一讀（特別是像我這種不熟前端技術的人）：

- [About HTML semantics and front-end architecture](https://nicolasgallagher.com/about-html-semantics-front-end-architecture/) by Nicolas Gallagher (2012-03-15)
- [Block Element Modifier](http://getbem.com/introduction/) (BEM)

這篇筆記只是記錄自己的一點閱讀心得，就先寫到這裡吧。如果你也有興趣了解，網路上可以找到許多前輩分享的文章，例如：[Tailwind CSS 新手上路：概念、安裝與工具推薦](https://medium.com/@Kelly_CHI/tailwind-css-introduction-and-tools-68e770b2bf7f)。之所以提這篇文章，是因為我覺得它開頭那張圖特別好笑，反映了一些人對 Tailwind CSS 的想法轉變歷程——從一開始的嫌惡之情，到使用之後發覺「真香」，從此成為愛用者。

Keep learning!
