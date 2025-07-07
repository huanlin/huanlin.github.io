---
title: 簡短的句子
linkTitle: 簡短的句子 (20 min)
weight: 6
---

> **預計閱讀時間：** 20 分鐘

軟體工程師通常會基於以下原因，試圖最小化實作中的程式碼行數：

* 較短的程式碼通常更容易讓他人閱讀。
* 較短的程式碼通常比較長的程式碼更容易維護。
* 額外的程式碼行數會引入額外的故障點（point of failure）。

事實上，同樣的規則也適用於技術寫作：

* 較短的文件讀起來比長文件快。
* 較短的文件通常比較長的文件更容易維護。
* 額外的文件行數會引入額外的故障點。

找到最短的文件實作方式需要時間，但最終是值得的。短句比長句更有力地傳達訊息，而且短句通常比長句更容易理解。

## 讓每個句子專注於單一想法

讓每個句子專注於一個想法、思緒或概念。就像程式中的陳述式執行單一任務一樣，句子也應該執行單一想法。例如，以下這個非常長的句子包含多個思緒：

> 1950 年代末期是程式語言的一個關鍵時代，因為 IBM 在 1957 年推出了 Fortran，而 John McCarthy 在隔年推出了 Lisp，這給予了程式設計師解決問題的迭代方法和遞迴方法。

將這個長句分解成一連串單一想法的句子，會得到以下結果：

> 1950 年代末期是程式語言的一個關鍵時代。IBM 在 1957 年推出了 Fortran，John McCarthy 在隔年發明了 Lisp。因此，到了 1950 年代末期，程式設計師可以用迭代或遞迴的方式解決問題。

### 練習 {#ex1}

請將下列過長的句子轉換成一連串較短的句子。無須大幅改寫，只需從單一句轉換成數句。

> In bash, use the if, then, and fi statements to implement a simple conditional branching block in which the if statement evaluates an expression, the then statement introduces a block of statements to run when the if expression is true, and the fi statement marks the end of the conditional branching block.
>
>
> 在 bash 中，可以使用 `if`、`then` 和 `fi` 指令來實作簡單的條件分支。`if` 負責判斷條件式是否成立。`then` 用來引入當條件為真時要執行的程式碼區塊。`fi` 則標示條件分支區塊的結尾。

<br>

{{< bs/collapse heading="點我看答案" expand=false >}}
In bash, use an if, then, and fi statement to implement a simple conditional branching block. The if statement evaluates an expression. The then statement introduces a block of statements to run when the if expression is true. The fi statement marks the end of the conditional branching block. (The resulting paragraph remains unclear but is still much easier to read than the original sentence.)

在 bash 中，可以使用 `if`、`then` 和 `fi` 指令來實作一個簡單的條件分支區塊。`if` 會判斷一個條件式。`then` 則用來引入當條件成立時要執行的程式碼區塊。`fi` 表示這個條件分支區塊的結尾。（雖然整段文字仍不夠清楚，但相較於原本的長句，已經更容易閱讀了。）
{{< /bs/collapse >}}

## 將一些長句轉換為列表

許多技術長句的內部，都藏著一個渴望掙脫的列表。例如，請看以下句子：

> 要改變迴圈的常規流程，您可以使用 `break` 陳述式（它會讓您跳出目前的迴圈）或 `continue` 陳述式（它會跳過目前迴圈迭代的其餘部分）。

當您在長句中看到連接詞「或」（**or**）時，可以考慮將該句重構為項目符號列表。當您在長句中看到嵌入的項目或任務列表時，可以考慮將該句重構為項目符號或編號列表。例如，前面的例子包含連接詞「或」（**or**），所以讓我們將該長句轉換為以下項目符號列表：

> 要改變迴圈的常規流程，請呼叫以下陳述式之一：
>
> * `break`，它會讓您跳出目前的迴圈。
> * `continue`，它會跳過目前迴圈迭代的其餘部分。

### 練習 {#ex2}

請將下列句子重構為更簡短、更清楚的版本，並確保你的答案包含一份清單：

1. 若要開始使用 Frambus 應用程式，你必須先在合適的商店找到該應用程式，使用有效的信用卡或金融卡付款，下載它，在 `/etc/Frambus` 檔案中設定 Foo 變數的值，然後重複唸兩次啟動咒語來執行它。
2. KornShell 是由 David Korn 在 1983 年發明的，當時他是貝爾實驗室的電腦科學家，這個 shell 是對 Bourne Shell 的擴充集合，包含更多功能、加強和改進，而 Bourne Shell 則是由另一位貝爾實驗室的電腦科學家 Stephen Bourne 在 1977 年發明的。

<br>

{{< bs/collapse heading="點我看答案" expand=false >}}
**句子 1**

請依照下列步驟開始使用 Frambus 應用程式：
1. 在合適的商店找到該應用程式。
2. 使用有效的信用卡或金融卡付款購買該應用程式。
3. 下載該應用程式。
4. 編輯 /etc/Frambus 檔案，為 Foo 變數指定一個值來設定該應用程式。
5. 唸兩次啟動咒語來執行該應用程式。

**句子 2**

下列兩位貝爾實驗室的電腦科學家發明了廣受歡迎的 shell：
- Stephen Bourne 在 1977 年發明了 Bourne Shell。
- David Korn 在 1983 年發明了 KornShell。
KornShell 是 Bourne Shell 的向下相容超集合，包含了許多對舊版 shell 的改進。
{{< /bs/collapse >}}

## 消除或減少無關的詞語

許多句子都包含填充詞——這些是消耗空間卻無法為讀者提供養分的文字垃圾。例如，看看您是否能找出以下句子中不必要的詞語：

> An input value greater than 100 causes the triggering of logging.
> <br />
> (大於100的輸入值會**導致觸發**日誌記錄。)

將 **causes the triggering of** 替換為更短的動詞 **triggers**，可以得到一個更短的句子：

> An input value greater than 100 triggers logging.
> <br />
> (大於100的輸入值會**觸發**日誌記錄。)

透過練習，您將能發現多餘的詞語，並在刪除或減少它們時感到無比的快樂。例如，請看以下句子：

> This design document provides a detailed description of Project Frambus.
> <br />
> (本設計文件**提供了關於** Frambus 專案的**詳細描述**。)

片語 **provides a detailed description of** 可以簡化為動詞 **describes** (或動詞 **details**)，因此結果句子可以變成：

> This design document describes Project Frambus.
> (本設計文件**描述了** Frambus 專案。)

下表建議了一些常見臃腫片語的替代方案：

| 冗長 (Wordy) | 簡潔 (Concise) |
| :--- | :--- |
| at this point in time | now |
| determine the location of | find |
| is able to | can |

### 練習 {#ex3}

請在不改變意思的前提下，將下列句子縮短：

1. In spite of the fact that Arnold writes buggy code, he writes error-free documentation.
2. Changing the sentence from passive voice to active voice enhances the clarification of the key points.
3. Determine whether Rikona is able to write code in COBOL.
4. Frambus causes the production of bugs, which will be chronicled in logs by the LogGenerator method.

<br>

{{< bs/collapse heading="點我看答案" expand=false >}}
以下是一些可能的解答：

1. Although Arnold writes buggy code, he writes error-free documentation. <br />
   **另一種寫法：** Arnold writes buggy code. However, he writes error-free documentation.
3. Changing the sentence from passive voice to active voice clarifies the key points.
4. Determine whether Rikona can code in COBOL.
5. Frambus produces bugs, which the LogGenerator method logs.

{{< /bs/collapse >}}

## 減少從屬子句 (選修)

子句是句子中一個獨立的邏輯片段，包含一個執行者和一個動作。每個句子都包含：

* 一個主要子句
* 零個或多個從屬子句

從屬子句（subordinate clause）修飾主要子句中的想法。顧名思義，從屬子句不如主要子句重要。例如，請看以下句子：

> Python is an interpreted programming language, which was invented in 1991. <br/>
> (Python是一種直譯式程式語言，它是在1991年發明的。)
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

有些從屬子句前面會加逗號，有些則不會。以下這句話中標示的從屬子句就是例子，它是以 **because** 開頭，而且沒有加逗號：

> I prefer to code in C++ <mark>because I like strong data typing.</mark>

在編輯時，請仔細檢查從屬子句。請記住「**一個句子 = 一個想法**」的[單一職責原則](https://en.wikipedia.org/wiki/Single-responsibility_principle)。句子中的從屬子句是擴展了單一想法，還是「分支」成一個獨立的想法？如果是後者，請考慮將有問題的從屬子句分割成獨立的句子。

### 練習 {#ex4}

判斷下列哪些句子中的從屬子句應該拆分成獨立的句子。（不要改寫句子，只要指出哪些句子應該重寫即可。）

1. Python is an interpreted language, which means that the language can execute source code directly. <br />(Python 是一種直譯式語言，這表示它可以直接執行原始碼。)
2. Bash is a modern shell scripting language that takes many of its features from KornShell 88, which was developed at Bell Labs. <br />(Bash 是一種現代的 shell 指令稿語言，它從 KornShell 88 繼承了許多功能，而 KornShell 88 是在貝爾實驗室開發的。)
3. Lisp is a programming language that relies on Polish prefix notation, which is one of the systems invented by the Polish logician Jan Łukasiewicz. <br />(Lisp 是一種程式語言，採用波蘭前置符號（Polish prefix notation），這種符號系統是由波蘭邏輯學家 Jan Łukasiewicz 發明的。)
4. I don't want to say that Fortran is old, but only radiocarbon dating can determine its true age. <br />(我不是說 Fortran 很老，但它的真實年齡可能需要用放射性碳定年法來測定。)

<br>

{{< bs/collapse heading="點我看答案" expand=false >}}
以下答案已將從屬子句以<mark>馬克筆</mark>或**粗體字**標示出來。

1. Python is an interpreted language, <mark>which means that the language can execute source code directly.</mark> <br />備註：這句中的從屬子句是用來延伸主句意思，句子本身沒問題。
2. Bash is a modern shell scripting language <mark>that takes many of its features from KornShell 88,</mark> **which was developed at Bell Labs.** <br/> 備註：這句的第一個從屬子句是用來延伸主句，但第二個從屬子句偏離主題太遠。建議拆成兩句。
3. Lisp is a programming language <mark>that relies on Polish prefix notation,</mark> **which is one of the systems invented by the Polish logician Jan Łukasiewicz.** <br />備註：這句的第一個從屬子句對句子非常重要，但第二個從屬子句讓讀者偏離主句太遠。建議拆成兩句。
4. I don't want to say that Fortran is old, <mark>but only radiocarbon dating can determine its true age.</mark> <br />備註：這句的從屬子句對句子很關鍵，因此句子本身沒問題。
{{< /bs/collapse >}}

## 區分 that 和 which

**That** 和 **which** 都可以用來引導從屬子句。它們之間有什麼區別？在某些國家，這兩個字幾乎可以互換使用。不過，美國的細心讀者通常會嚴正指出，你又搞混這兩個字了。

在美國，**which** 用於非必要（非限制性）的從屬子句，也就是那種刪掉句子仍然通順的子句；**that** 則用於必要（限制性）的從屬子句，是句子無法省略的部分。舉例來說，下面這句的重點是「Python 是一種直譯語言」，句子即使刪掉「Guido van Rossum invented」仍然合理：

> Python is an interpreted language, which Guido van Rossum invented.
> <br />
> Python 是一種直譯語言，這是 Guido van Rossum 發明的。

相較之下，下面這句需要用到「don't involve linear algebra」這個限制性子句：

> Fortran is perfect for mathematical calculations that don't involve linear algebra.
> <br/>
> Fortran 非常適合用來做不涉及線性代數的數學運算。

如果你朗讀一句話時，在從屬子句前聽到停頓，就用 **which**；如果沒有停頓，就用 **that**。你可以回頭讀前面兩句例句，第一句在從屬子句前有停頓嗎？

在 **which** 前面要加逗號，**that** 前面則不加逗號。

<br /><br />

**下一個單元：** [清單與表格]({{< relref "../lists-and-tables/index.md" >}})
