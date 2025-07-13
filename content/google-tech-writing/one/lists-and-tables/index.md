---
title: 清單與表格
linkTitle: 清單與表格 (15 分鐘)
weight: 7
---

> **預估時間：** 15 分鐘

好的清單可以將技術上的混亂轉化為井然有序的內容。技術讀者通常喜歡清單。因此，在撰寫時，請尋找機會將段落文字轉換為清單。

## 選擇正確的清單類型

以下類型的清單在技術寫作中佔主導地位：

* 項目符號清單（Bulleted lists）
* 編號清單（Numbered lists）
* 內嵌清單（Embedded lists）

對於無序的項目，請使用**項目符號清單**（bulleted list）；對於有序的項目，請使用**編號清單**（numbered list）。換句話說：

* 若重新排列項目符號清單中的項目，整體意思不會改變。
* 若重新排列編號清單中的項目，整體意思會改變。

例如，我們將以下內容設為項目符號清單，因為重新排列其項目不會改變清單的含義：

> Bash 提供以下字串操作機制：
> * 從字串開頭刪除子字串
> * 將整個檔案讀入一個字串變數

相反地，以下這段必須是編號清單，因為項目順序會影響操作流程：

> 請按照以下步驟重新設定伺服器：
> 1. 停止伺服器。
> 2. 編輯設定檔。
> 3. 重新啟動伺服器。

**內嵌清單**（有時也稱為「行內清單」［run-in list］）是指將所有項目塞在同一個句子中。例如，下面這句話包含一個有四個項目的內嵌清單：

> Llamacatcher API 讓呼叫者可以建立和查詢駱馬、分析羊駝、刪除小羊駝以及追蹤單峰駱駝。

一般來說，內嵌清單不是呈現技術資訊的理想方式。建議將內嵌清單改寫為項目符號清單或編號清單。例如，你應該把前述例句改寫成以下段落：

> Llamacatcher API 讓呼叫者可以執行以下操作：
> * 建立和查詢駱馬。
> * 分析羊駝。
> * 刪除小羊駝。
> * 追蹤單峰駱駝。

### 練習 {#ex1}

請將以下段落改寫成一個或多個清單：

> 今天上班時，我必須撰寫三個單元測試、寫一份設計文件，還要審閱 Janet 最新的文件。下班後，我必須在不使用任何水的情況下洗車，然後在不使用毛巾的情況下將車子擦乾。

<br>

{{< bs/collapse heading="點我看答案" expand=false >}}
以下提供兩個參考答案，此為其一：
> 我今天上班時必須做以下幾件事：
> * 撰寫三個單元測試程式碼。
> * 撰寫設計文件。
> * 審閱 Janet 最新的文件。
>
> 下班後，我必須做以下幾件事：
> 1. 不用水洗車。
> 2. 不用毛巾擦乾車子。

以下是另一個參考答案：

> 我今天必須完成以下任務：
> * 工作時：
>   * 撰寫三個單元測試程式碼。
>   * 撰寫設計文件。
>   * 審閱 Janet 最新的文件。
>
> * 下班後：
>   1. 不用水洗車。
>   2. 不用毛巾擦乾車子。
{{< /bs/collapse >}}

## 保持清單項目一致（平行）

怎樣的清單是有效的？有效的清單具備平行性（parallel）；無效的清單則會包含不平行的項目。所謂平行的清單，是指其中所有項目彼此「相互呼應」，在以下幾個方面保持一致，能讓讀者一眼看出它們是同一類型的資訊：

* 文法結構（Grammar）
* 邏輯分類（Logical category）
* 大小寫格式（Capitalization）
* 標點符號（Punctuation）

相對地，只要有任一項目在上述其中一個方面不一致，整份清單就屬於非平行清單。

舉例來說，下面這個清單是平行的，因為所有項目都是複數名詞（文法），都是可以食用的東西（邏輯分類），採用標題式大寫（大小寫格式），且沒有句號或逗號（標點符號）：

* Carrots
* Potatoes
* Cabbages

相較之下，下面這份清單在各方面都缺乏一致，讀起來很不協調：

* Carrots
* Potatoes
* The summer light obscures all memories of winter.

再看下面這個例子，它是平行的，因為所有項目都是完整的句子，並且具備完整句子的大小寫與標點符號：

* Carrots contain lots of Vitamin A.
* Potatoes taste delicious.
* Cabbages provide oodles of Vitamin K.

讀者在看到清單的第一個項目時，就會期待後續的項目遵循相同的格式與風格，因此第一個項目等於建立了一種模式。要寫出清楚有效的清單，就必須保持這種平行性。

### 練習 {#ex2}

以下這份清單是平行的還是非平行的？

* Broccoli inspires feelings of love or hate.
* Potatoes taste delicious.
* Cabbages.

<br />

{{< bs/collapse heading="點我看答案" expand=false >}}
這份清單是非平行的。前兩個項目是完整的句子，但第三個項目不是句子。（不要被第三項的大小寫與標點所誤導。）
{{< /bs/collapse >}}

### 練習 {#ex3}

以下這份清單是平行的還是非平行的？

* The red dots represent sick trees.
* Immature trees are represented by the blue dots.
* The green dots represent healthy trees.

<br />

{{< bs/collapse heading="點我看答案" expand=false >}}
這是一份非平行的清單。第一和第三項使用主動語態，而第二項則使用被動語態。
{{< /bs/collapse >}}

## 編號清單項目以祈使動詞開頭

建議在編號清單中，讓每個項目都以祈使動詞開頭。祈使動詞（imperative verb）就是用來發出命令的動詞，例如「**開啟**（open）」、「**啟動**（start）」等。請看以下這份平行的編號清單，所有項目都以祈使動詞開頭：

1. Download the Frambus app from Google Play or iTunes. <br/>下載 Frambus 應用程式（從 Google Play 或 iTunes）。
2. Configure the Frambus app's settings. <br/>設定 Frambus 應用程式的偏好選項。
3. Start the Frambus app. <br/>啟動 Frambus 應用程式。

反之，下列編號清單就不具平行性，因為前三項中只有前兩項是以祈使動詞開頭，第三項不是：

1. Instantiate the Froobus class.
2. Invoke the Froobus.Salmonella() method.
3. The process stalls.

### 練習 {#ex4}

請讓下列清單保持平行結構，並確保清單中的每個項目都以祈使動詞開頭：

1. Stop Frambus
2. The key configuration file is `/etc/frambus`. Open this file with an ASCII text editor.
3. In this file, you will see a parameter named Carambola, which is currently set to the default value (32). Change this value to 64.
4. When you are finished setting this parameter, save and close the configuration file
5. now, start Frambus again.

<br />

{{< bs/collapse heading="點我看答案" expand=false >}}
參考答案：

1. Stop Frambus.
2. Open the key configuration file, /etc/frambus, with an ASCII text editor.
3. Change the Carambola parameter from its default value (32) to 64.
4. Save and close the configuration file.
5. Restart Frambus.
{{< /bs/collapse >}}

## 適當為清單項目加上標點符號

雖然各種風格指南對清單項目的標點符號用法有不同的建議，[Google 開發者文件風格指南](https://developers.google.com/style/lists#capitalization-and-end-punctuation)的建議是每個清單項目都以大寫字母開頭（有一些例外）。例如：

* Loops
* Conditionals
* Variable declarations

如果清單項目是一個句子，請使用適當的句尾標點符號。例如：

1. Open the program.
2. Click the settings icon.

## 建立實用的表格

善於分析的人往往喜歡表格。如果一個頁面包含多個段落和一個表格，工程師的目光通常會先掃向那個表格。

建立表格時，請參考以下準則：

* 為每一欄加上有意義的標頭。不要讓讀者猜測每一欄的內容是什麼。
* 避免在單一儲存格中放入過多文字。如果某儲存格的內容超過兩句話，請思考是否應改用其他格式呈現該資訊。
* 雖然不同欄位（columns）可以包含不同類型的資料，但應力求在同一欄內保持一致的寫法。例如，同一欄中的儲存格不應混雜數值與馬戲團表演者的姓名。

> **注意：** 有些表格在所有裝置上都無法良好呈現。例如，在筆記型電腦上看起來很清楚的表格，在手機上可能會很難閱讀。

## 為每個清單和表格撰寫引言

我們建議為每個清單和表格加上一個引言句，告訴讀者該清單或表格代表什麼。換句話說，為清單或表格提供上下文。引言句以冒號而不是句點結尾。
我們建議每一個清單或表格前加上一句說明文字，讓讀者了解該清單或表格的用意。換句話說，為清單或表格提供背景說明。這句引言句的結尾應該用冒號而非句號。

雖然不是硬性規定，但我們建議在引言中使用「**以下**（following）」這個詞。請參考下列範例引言句：

> The following list identifies key performance parameters: <br />以下清單標識了關鍵效能參數：
>
> Take the following steps to install the Frambus package: <br />請按照以下步驟安裝 Frambus 套件：
>
> The following table summarizes our product's features against our key competitors' features: <br />下表總結了我們產品與主要競爭對手產品的功能：

### 練習 {#ex5}

為下表寫一個引言句：

| 語言 | 發明者 | 推出年份 | 主要特點 |
| :--- | :--- | :--- | :--- |
| Lisp | John McCarthy | 1958 | 遞迴 |
| C++ | Bjarne Stroustrup | 1979 | 物件導向程式設計 |
| Python | Guido van Rossum | 1994 | 簡潔性 |

<br />

{{< bs/collapse heading="點我看答案" expand=false >}}
這裡提供兩種參考答案：

> 下表包含一些關於流行程式語言的關鍵事實：
>
> 下表標識了三種流行程式語言的發明者、發明年份和主要特點：
{{< /bs/collapse >}}

<br /><br />

**下一單元：** [段落]({{< relref "../paragraphs/index.md" >}})