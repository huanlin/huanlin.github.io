---
title: 00 關於這個筆記
tags: [Go]
---

## Why Golang?

我學習 Go 語言的原因是：需要用到，覺得不學不行了。

無論是我的個人網站 [huanlin.cc](https://huanlin.cc/docs/) 還是其他場合所建立的文件網站，都是用 Hugo 搭建，而我逐漸發現需要一些強化的功能，要寫一些程式來實現才行。Hugo 是基於 Go 語言，故學習 Go 語言來自行開發文件網站的一些自訂或加強功能，應是最直截了當的選擇。這是我學習 Go 語言最主要的原因。

至於能學到什麼程度，自己也沒把握，畢竟每個人一天都是 24 小時，生活中經常要分心去處理很多待辦事項，很難專精。而且，這個領域變化很快，明天會發生什麼事情都很難說。

## About this note

由於我有一點點其他程式語言的經驗，所以學習 Golang 的時候會跳過一些基礎語法，只留意我不熟悉的部分。因此，我的筆記也會跳過許多基礎的東西，甚至有可能遺漏一些重要的觀念。

一如以往，我只是一邊學習，一邊寫點筆記罷了。

基於上述理由，我不認為我的筆記對別人能有多少幫助。也因為這個緣故，我把學習 Go 語言的相關資源整理在下一節，如果有人正好也要學習 Go，建議優先參考那些更專業可靠的學習資源。

## Learning resources

本節整理一些學習 Go 程式設計的資源，主要是網站和書籍。

### Free resources

這個時代，免費資源超多，不怕沒得學，只怕難選擇。這裡只列出一些我有看的，也覺得不錯的：

| Title                                                              | Remark                                 |
| ------------------------------------------------------------------ | -------------------------------------- |
| [使用 Golang 打造 Web 應用程式](https://willh.gitbook.io/build-web-application-with-golang-zhtw) by 保哥 | 最後一次更新似乎是 2021 年，仍有許多入門和通用知識值得參考。 |
| [A Tour of Go](https://go.dev/tour/)                               | 一小塊接著一小塊的學習方式，容易消化。 |
| [Go by Example](https://gobyexample.com/)                          | 有許多範例，每個範例都有解說。         |
| [Go Tutorial](https://www.w3schools.com/go/index.php)              | Go 語法的分類編排很清楚、方便查找。 |
| [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests) | 以撰寫測試來學習 Go 語言。（我覺得是很棒的學習方式） |
| [Effective Go](https://go.dev/doc/effective_go)                    | [進階] 由 Golang core team 撰寫的技術文章。 |
| [The Go Programming Language Specification](https://go.dev/ref/spec) | 官方的語法規格。 |

另外，我也會用一些零碎時間看 Youtube 教學影片。碰到想要更深入了解的部分，就針對特定議題來閱讀 Go 的官方文章，或者找書來讀。

### Books

跟著一本書來學習的好處是比較有系統地、按步就班地學習。

以下書單，前面四本我都有買。有給分的是我對那本書的喜愛程度（滿分是 5 分），不代表看完了。

| Title <div style="width: 140px;"></div>| Level<div style="width: 80px;"></div> | Rate | Description |
|-----------------------|------------------------|---------|------------------------------------|
| [Go in Practice, 2nd Edition](https://www.manning.com/books/go-in-practice-second-edition) | 中高階 | 5/5 | 著重實戰所需的知識與技術，不會詳細介紹基礎語法，例如變數如何宣告與賦值、如何寫迴圈等等，但重要的語言特性都有涵蓋，像是介面、泛型、concurrency、錯誤處理等等。後面幾章還介紹了單元測試、Web API、雲端應用程式設計等議題。|
| [Go by Example](https://www.manning.com/books/go-by-example) | 中高階 | 5/5 | 適合 Go 語言的初學者，但必須具備其他程式語言的開發經驗。 |
| [100 Go Mistakes and How to Avoid Them](https://100go.co/book/) | 中高階 | 5/5 | Go 專業開發人員必讀。 |
| [Go in Action, 2nd Edition](https://www.manning.com/books/go-in-action-second-edition) | 入門至中階 | 3/5 | 可以當作語法參考手冊，需要時再查閱。最後一章介紹單元測試。未包含 web 程式設計相關議題。|
| [Learning Go 2nd Edition](https://www.amazon.com/Learning-Go-Jon-Bodner-ebook/dp/B0CS5DY1VN) | 入門至進階 |  |
| [Let's Go](https://lets-go.alexedwards.net/) | 入門至實戰 |  | 從 Hello World 到開發 Web 應用程式。 |

### Courses

從缺，因為我沒有購買線上課程。

## 版本歷史

以下表格是我這份 Go 學習筆記的修改歷史的摘要。

| 日期        | 更新了什麼                                                 |
|------------|-----------------------------------------------------------|
| 2024-09-02 | 增加〈雜七雜八但是重要〉、〈結構〉。 |
| 2024-08-31 | 初次發布：關於這個筆記、Get started、Hello world、命令列應用程式、Code organization。 |
