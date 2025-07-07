---
title: 清單與表格
linkTitle: 清單與表格 (15 min)
weight: 7
---

> **預計閱讀時間：** 15 分鐘

好的清單可以將技術上的混亂轉化為井然有序的內容。技術讀者通常喜歡清單。因此，在撰寫時，請尋找機會將段落文字轉換為清單。

## 選擇正確的清單類型

以下類型的清單在技術寫作中佔主導地位：

* 項目符號清單
* 編號清單
* 嵌入式清單

對於無序的項目，請使用**項目符號清單**（bulleted list）；對於有序的項目，請使用**編號清單**（numbered list）。換句話說：

* 如果您重新排列項目符號清單中的項目，清單的含義不會改變。
* 如果您重新排列編號清單中的項目，清單的含義會改變。

例如，我們將以下內容設為項目符號清單，因為重新排列其項目不會改變清單的含義：

> Bash 提供以下字串操作機制：
> * 從字串開頭刪除子字串
> * 將整個檔案讀入一個字串變數

相比之下，以下清單必須是編號清單，因為重新排列其項目會改變清單的含義：

> 請按照以下步驟重新設定伺服器：
> 1. 停止伺服器。
> 2. 編輯設定檔。
> 3. 重新啟動伺服器。

**嵌入式清單**（有時稱為行內清單［run-in list］）將項目塞入一個句子中。例如，以下句子包含一個有四個項目的嵌入式清單。

> Llamacatcher API 讓呼叫者可以建立和查詢駱馬、分析羊駝、刪除小羊駝以及追蹤單峰駱駝。

一般來說，嵌入式清單是呈現技術資訊的較差方式。請嘗試將嵌入式清單轉換為項目符號清單或編號清單。例如，您應該將包含嵌入式清單的句子轉換為以下段落：

> Llamacatcher API 讓呼叫者可以執行以下操作：
> * 建立和查詢駱馬。
> * 分析羊駝。
> * 刪除小羊駝。
> * 追蹤單峰駱駝。

### 練習 {#ex1}

請將以下段落轉換成一個或多個清單：

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

## 保持清單項目平行

是什麼因素區分有效的清單和不良的清單？有效的清單具備平行結構；而不良的清單通常缺乏平行性。所有有效清單中的項目看起來彼此「相互呼應」。也就是說，有效清單中的所有項目在以下幾個方面保持一致：

* 文法
* 邏輯分類
* 大寫規則
* 標點符號

相反地，不平行清單中至少有一個項目在上述至少一項一致性檢查中失敗。

例如，以下清單是平行的，因為所有項目都是複數名詞（文法）、可食用的（邏輯分類）、標題大小寫（大寫規則），並且沒有句點或逗號（標點符號）。

* Carrots
* Potatoes
* Cabbages

相較之下，下面這份清單在所有面向上都嚴重缺乏平行性：

* Carrots
* Potatoes
* The summer light obscures all memories of winter.

以下清單是平行的，因為所有項目都是完整的句子，具有完整的句子大小寫和標點符號：

* Carrots contain lots of Vitamin A.
* Potatoes taste delicious.
* Cabbages provide oodles of Vitamin K.

清單中的第一個項目建立了一種模式，讀者期望在後續項目中看到這種模式的重複。

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

考慮讓編號清單中的所有項目都以祈使動詞開頭。祈使動詞（imperative verb）是一個命令，例如 **open** 或 **start**。舉例來說，請注意以下平行編號清單中的所有項目都以祈使動詞開頭：

1. Download the Frambus app from Google Play or iTunes.
2. Configure the Frambus app's settings.
3. Start the Frambus app.

以下編號清單不平行，因為其中兩個句子以祈使動詞開頭，但第三個項目沒有：

1. Instantiate the Froobus class.
2. Invoke the Froobus.Salmonella() method.
3. The process stalls.

### 練習 {#ex4}

請讓下列清單保持平行結構。確保結果中每個項目都以祈使動詞開頭：

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

## 適當為項目加上標點符號

雖然不同的風格指南對清單項目的標點符號提供了相互矛盾的建議，但 [Google 開發者文件風格指南](https://developers.google.com/style/lists#capitalization-and-end-punctuation)建議（有一些例外）每個清單項目都以大寫字母開頭。例如：

* Loops
* Conditionals
* Variable declarations

如果清單項目是一個句子，請使用適當的句尾標點符號。例如：

1. Open the program.
2. Click the settings icon.

## 建立有用的表格

善於分析的人往往喜歡表格。如果一個頁面包含多個段落和一個表格，工程師的目光會迅速轉向表格。

建立表格時請考慮以下準則：

* 為每一欄加上有意義的標頭。不要讓讀者猜測每一欄的內容。
* 避免在表格儲存格中放入過多文字。如果一個表格儲存格包含超過兩個句子，請自問該資訊是否屬於其他格式。
* 雖然不同欄（columns）可以包含不同類型的資料，但應力求在單一欄內保持平行結構。例如，特定表格的某一欄的所有儲存格不應混合數字資料和著名的馬戲團表演者。

> **注意：** 有些表格在所有裝置上都無法良好呈現。例如，在筆記型電腦上看起來很棒的表格，在手機上可能會看起來很糟糕。

## 為每個清單和表格加上引言

我們建議為每個清單和表格加上一個引言句，告訴讀者該清單或表格代表什麼。換句話說，為清單或表格提供上下文。引言句以冒號而不是句點結尾。

雖然不是強制要求，但我們建議在引言句中加入「以下」一詞。例如，考慮以下引言句：

> 以下清單標識了關鍵效能參數：
>
> 請按照以下步驟安裝 Frambus 套件：
>
> 下表總結了我們產品與主要競爭對手產品的功能：

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

**下一個單元：** [段落]({{< relref "../paragraphs/index.md" >}})