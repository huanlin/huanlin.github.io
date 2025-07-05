---
title: 簡短的句子
linkTitle: 簡短的句子 (20 min)
weight: 6
---

軟體工程師通常會基於以下原因，試圖最小化實作中的程式碼行數：

* 較短的程式碼通常更容易讓他人閱讀。
* 較短的程式碼通常比較長的程式碼更容易維護。
* 額外的程式碼行數會引入額外的故障點。

事實上，同樣的規則也適用於技術寫作：

* 較短的文件讀起來比長文件快。
* 較短的文件通常比較長的文件更容易維護。
* 額外的文件行數會引入額外的故障點。

找到最短的文件實作方式需要時間，但最終是值得的。短句比長句更有力地傳達訊息，而且短句通常比長句更容易理解。

## 讓每個句子專注於單一想法

讓每個句子專注於一個想法、思緒或概念。就像程式中的陳述式執行單一任務一樣，句子也應該執行單一想法。例如，以下這個非常長的句子包含多個思緒：

> 1950年代末期是程式語言的一個關鍵時代，因為IBM在1957年推出了Fortran，而John McCarthy在隔年推出了Lisp，這給予了程式設計師解決問題的迭代方法和遞迴方法。

將這個長句分解成一連串單一想法的句子，會得到以下結果：

> 1950年代末期是程式語言的一個關鍵時代。IBM在1957年推出了Fortran。John McCarthy在隔年發明了Lisp。因此，到了1950年代末期，程式設計師可以用迭代或遞迴的方式解決問題。

## 將一些長句轉換為列表

許多技術長句的內部，都藏著一個渴望掙脫的列表。例如，請看以下句子：

> 要改變迴圈的常規流程，您可以使用 `break` 陳述式（它會讓您跳出目前的迴圈）或 `continue` 陳述式（它會跳過目前迴圈迭代的其餘部分）。

當您在長句中看到連接詞 `or` 時，可以考慮將該句重構為項目符號列表。當您在長句中看到嵌入的項目或任務列表時，可以考慮將該句重構為項目符號或編號列表。例如，前面的例子包含連接詞 `or`，所以讓我們將該長句轉換為以下項目符號列表：

> 要改變迴圈的常規流程，請呼叫以下陳述式之一：
>
> * `break`，它會讓您跳出目前的迴圈。
> * `continue`，它會跳過目前迴圈迭代的其餘部分。

## 消除或減少無關的詞語

許多句子都包含填充詞——這些是消耗空間卻無法為讀者提供養分的文字垃圾。例如，看看您是否能找出以下句子中不必要的詞語：

> An input value greater than 100 causes the triggering of logging.
> (大於100的輸入值會**導致觸發**日ログ記錄。)

將 `causes the triggering of` 替換為更短的動詞 `triggers`，可以得到一個更短的句子：

> An input value greater than 100 triggers logging.
> (大於100的輸入值會**觸發**日誌記錄。)

透過練習，您將能發現多餘的詞語，並在刪除或減少它們時感到無比的快樂。例如，請看以下句子：

> This design document provides a detailed description of Project Frambus.
> (本設計文件**提供了關於** Frambus 專案的**詳細描述**。)

片語 `provides a detailed description of` 可以簡化為動詞 `describes` (或動詞 `details`)，因此結果句子可以變成：

> This design document describes Project Frambus.
> (本設計文件**描述了** Frambus 專案。)

下表建議了一些常見臃腫片語的替代方案：

| 冗長 (Wordy) | 簡潔 (Concise) |
| :--- | :--- |
| at this point in time | now |
| determine the location of | find |
| is able to | can |

## 減少從屬子句 (可選)

子句是句子中一個獨立的邏輯片段，包含一個執行者和一個動作。每個句子都包含：

* 一個主要子句
* 零個或多個從屬子句

從屬子句修飾主要子句中的想法。顧名思義，從屬子句不如主要子句重要。例如，請看以下句子：

> Python is an interpreted programming language, which was invented in 1991.
> (Python是一種直譯式程式語言，**它是在1991年發明的**。)
>
> * 主要子句：Python is an interpreted programming language
> * 從屬子句：which was invented in 1991

您通常可以透過引導它們的詞語來識別從屬子句。以下列表（絕非完整）顯示了引導從屬子句的常見詞語：

* which
* that
* because
* whose
* until
* unless
* since

在編輯時，請仔細檢查從屬子句。請記住「一個句子 = 一個想法」的單一職責原則。句子中的從屬子句是擴展了單一想法，還是分支到了一個獨立的想法？如果是後者，請考慮將有問題的從屬子句分割成獨立的句子。

## 區分 `that` 和 `which`

`that` 和 `which` 都引導從屬子句。它們之間有什麼區別？在美國英語中，`which` 用於非必要的從屬子句，而 `that` 用於句子不可或缺的必要從屬子句。

例如，以下句子的關鍵訊息是「Python是一種直譯式語言」；即使沒有「Guido van Rossum發明了它」這部分，句子仍然可以成立：

> Python is an interpreted language, which Guido van Rossum invented.

相比之下，以下句子需要「don't involve linear algebra」這部分：

> Fortran is perfect for mathematical calculations that don't involve linear algebra.

如果您大聲朗讀一個句子，並在從屬子句前聽到停頓，那麼請使用 `which`。如果您沒有聽到停頓，請使用 `that`。

在 `which` 前面加上逗號；不要在 `that` 前面加上逗號。