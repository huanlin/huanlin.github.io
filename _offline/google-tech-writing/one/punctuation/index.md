---
title: 標點符號（選修）
linkTitle: 標點符號（選修）(5 分鐘)
weight: 12
---

> **預估時間：** 5 分鐘

本單元為選修，目的是提供標點符號的快速複習。

## 逗號

程式語言對標點符號有明確的規定。相較之下，英語中關於逗號的規則則較為模糊。作為一個指引，只要讀者在句中自然停頓的地方，就插入逗號。對於有音樂背景的人來說，如果句號是全音符（semibreve）的休止符，那麼逗號大概是二分音符（minim）或四分音符（crotchet）的休止符。換句話說，逗號的停頓時間比句號短。例如，當你大聲朗讀以下句子時，可能會在 just 這個字前短暫停頓：

> C behaves as a mid-level language, just a couple of steps up in abstraction from assembly language.

有些情況下必須使用逗號。例如，使用逗號來分隔嵌入式清單中的項目，如下所示：

> Our company uses C++, Python, Java, and JavaScript.

你可能會好奇清單中的最後一個逗號，也就是插入在第 N-1 項和第 N 項之間的逗號。這個逗號——稱為**序列逗號**（serial comma）或**牛津逗號**（Oxford comma）——在使用上是有爭議的。我們建議加上這個最後的逗號，因為技術寫作需要選擇最不模糊的解決方案。話雖如此，我們其實更喜歡透過將內嵌清單轉換成項目符號清單來避開這個爭議。

在表達條件的句子中，請在條件和結果之間加上逗號。例如，以下兩個句子都在正確的位置加上了逗號：

> If the program runs slowly, try the `--perf` flag.
>
> If the program runs slowly, then try the `--perf` flag.

你也可以在一對逗號之間插入一個簡短的定義或題外話，如以下範例：

>　Python, an easy-to-use language, has gained significant momentum in recent years.

最後，避免使用逗號將兩個獨立的想法連接在一起。例如，以下句子中的逗號犯了一種稱為**逗號拼接**（comma splice）的標點錯誤：

<i class="fa-solid fa-thumbs-down fa-lg" style="color: red;"></i> **不建議**

> Samantha is a wonderful coder, she writes abundant tests.

使用句號而非逗號來分隔兩個獨立的想法。例如：

<i class="fa-solid fa-thumbs-up fa-lg" style="color: green;"></i> **推薦**

> Samantha is a wonderful coder. She writes abundant tests.

### 練習 {#ex1}

請在以下段落中適當添加逗號：

> Protocol Buffers sometimes known as protobufs are our team's main structured data format. Use Protocol Buffers to represent store and transfer structured data. Unlike XML Protocol Buffers are compiled. Consequently clients transmit Protocol Buffers efficiently which has led to rapid adoption.

提示：朗讀這段文字，並在聽到短暫停頓的地方加上逗號。

<br />

{{< bs/collapse heading="點我看答案" expand=false >}}
參考答案：

> Protocol Buffers, sometimes known as protobufs, are our team's main structured data format. Use Protocol Buffers to represent, store, and transfer structured data. Unlike XML, Protocol Buffers are compiled. Consequently, clients transmit Protocol Buffers efficiently, which has led to rapid adoption.
{{< /bs/collapse >}}

## 分號

句號用來分隔不同的想法；分號則用來連接高度相關的想法。例如，請注意下列句子中分號如何連接第一個和第二個想法：

<i class="fa-solid fa-thumbs-up fa-lg" style="color: green;"></i> **推薦**

> Rerun Frambus after updating your configuration file; don't rerun Frambus after updating existing source code.

在使用分號之前，請先問自己，如果將分號兩邊的想法對調，句子是否仍然通順。例如，將前面的例子反過來仍然是一個有效的句子：

> Don't rerun Frambus after updating existing source code; rerun Frambus after updating your configuration file.

分號前後的想法必須各自是語法完整的句子。例如，以下的分號用法是不正確的，因為分號後的部分是[子句](https://developers.google.com/tech-writing/one/short-sentences#reduce_subordinate_clauses_optional)，而非完整句子：

<i class="fa-solid fa-thumbs-down fa-lg" style="color: red;"></i> **不建議**

> Rerun Frambus after updating your configuration file; not after updating existing source code.

<i class="fa-solid fa-thumbs-up fa-lg" style="color: green;"></i> **推薦**

> Rerun Frambus after updating your configuration file, not after updating existing source code.

您幾乎應該總是使用逗號，而非分號，來分隔內嵌清單中的項目。例如，以下使用分號的方式是不正確的：

<i class="fa-solid fa-thumbs-down fa-lg" style="color: red;"></i> **不建議**

> Style guides are bigger than the moon; more essential than oxygen; and completely inscrutable.

如本課程前面提到的，技術寫作通常偏好使用項目符號清單而非內嵌清單。然而，如果你真的偏好使用內嵌清單，請使用逗號而非分號來分隔項目，如以下範例所示：

<i class="fa-solid fa-thumbs-up fa-lg" style="color: green;"></i> **推薦**

> Style guides are bigger than the moon, more essential than oxygen, and completely inscrutable.

許多句子在分號後緊接著放置轉折詞或片語。在這種情況下，轉折詞後面應加逗號。請注意以下兩個例子中轉折詞後的逗號：

> Frambus provides no official open source package for string manipulation; however, subsets of string manipulation packages are available from other open source projects.

> Even seemingly trivial code changes can cause bugs; therefore, write abundant unit tests.

### 練習 {#ex2}

下列哪一個句號或逗號可以用分號取代？

1. Python is a popular programming language. The C language was developed long before Python.
2. Model learning for a low value of X appears in the top illustration. Model learning for a high value of X appears in the bottom illustration.
3. I'm thankful for my large monitor, powerful CPU, and blazing bandwidth.

<br />

{{< bs/collapse heading="點我看答案" expand=false >}}
1. 你不能將第 1 句的句號改成分號，因為這兩句話的關聯僅是模糊的。
2. 你可以將第 2 句的句號改成分號，因為這兩句話的關聯非常密切。
3. 你不能將第 3 句的逗號改成分號。這裡使用逗號來分隔各個項目是正確的。
{{< /bs/collapse >}}

## 破折號

破折號（Em dashes）是引人注目的標點符號，具有豐富的標點可能性。破折號表示比逗號更長的停頓——更大的中斷。對於音樂上熟悉的人來說，可以將逗號視為四分音符（四分休止符）的停頓，而破折號則是二分音符（二分休止符）的停頓。例如：

> C++ is a rich language—one requiring extensive experience to fully understand.

作者有時會使用一對長破折號來隔開插入語，如以下範例所示：

> Protocol Buffers—often nicknamed protobufs—encode structured data in an efficient yet extensible format.
>
> Protocol Buffers——常被暱稱為 protobufs——以高效且可擴充的格式編碼結構化資料。

在前面的範例中，我們能否用逗號代替長破折號？當然可以。我們為什麼選擇長破折號而不是逗號？感覺。藝術。經驗。

## En dash 和連字號

請參考下表所示的水平標點符號：

| 名稱 | 符號 | 相對寬度 |
| --- | --- | --- |
| 長破折號（em dash） | — | 最寬（通常為字母 m 的寬度） |
| 短破折號（en dash） | – | 中等（通常為字母 n 的寬度） |
| 連字號（hyphen） | \- | 最窄 |

> [!NOTE]
> 譯註：表格中的符號皆為英文半形符號。

有些風格指南建議在某些用途中使用 en dash。然而，Google 風格指南對 en dash 提出了以下直白的建議：

> 請勿使用。

連字號很棘手。在技術寫作中，連字號用於連接某些複合詞，例如：

* Self-attention
* N-gramm

令人困惑的是，三字複合詞通常在第一和第二個字之間會有連字號，但在第二和第三個字之間則沒有。例如：

* Decision-making system
* Floating-point feature

對於連字號有疑慮時，請查閱字典、詞彙表或寫作風格指南。

**注意：** 如果您查閱多本字典、詞彙表或風格指南關於連字號的用法，可能會遇到不一致的情況。

## 冒號

在技術寫作中，使用冒號提醒讀者接下來會有列表或表格。換句話說，介紹列表或表格的句子應以冒號結尾。以下範例中，請注意介紹列表的句子末尾的冒號：

> 請考慮以下重要的程式語言：
> * Python
> * Java
> * C++

技術寫作偏好使用項目符號清單或編號清單，而非內嵌清單。話雖如此，你仍可以使用冒號來引入內嵌清單，如以下範例所示：

> Consider the following important programming languages: Python, Java, and C++.

並非所有內嵌清單都需要冒號。例如：

> My three favorite programming languages are Python, Java, and anything other than C++.

## 括號

使用括號來補充次要觀點或插入離題內容。括號會讓讀者知道，括號內的文字不是重點。也因為括號內的內容不是重點，有些編輯認為，如果某段文字適合用括號包起來，那它可能根本不該出現在文件中。作為折衷方案，在技術寫作中應盡量減少使用括號。

關於句號與括號的規則並不總是明確，以下是一般的標準規則：

* 如果括號內包含的是完整句子，句號應放在括號內側。
* 如果括號只是出現在句子的結尾，但不包含整個句子，句號應放在括號外側。

例如：

> (Incidentally, Protocol Buffers make great birthday gifts.)
>
> Binary mode relies on the more compact native form (described later in this document).

<br /><br />

**下一單元：** [Markdown]({{< relref "markdown/" >}}) （選修）