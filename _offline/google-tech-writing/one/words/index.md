---
title: 字詞
linkTitle: 字詞 (10 分鐘)
weight: 3
---

> **預估時間：** 10 分鐘

我們對文件進行了廣泛的研究，結果發現，世界上最好的句子主要由字詞（words）組成。

## 定義新的或不熟悉的術語

在撰寫或編輯時，要學會辨識目標讀者中某些人或所有人可能不熟悉的術語。當你發現這樣的術語時，請採取以下兩種策略之一：

* 如果該術語已存在，請連結到一個好的現有解釋。（不要重新發明輪子。）
* 如果你的文件是首次介紹該術語，請定義它。如果你的文件介紹了許多術語，請將定義收集到一個詞彙表中。

## 一致地使用術語

如果你在方法的中間某處更改了變數的名稱，你的程式碼將無法編譯。同樣，如果你在文件當中重新命名了某個術語，你的想法也將無法（在使用者的腦袋裡）通過編譯。

重點：在整份文件中始終如一地使用相同的、明確的詞彙或術語。一旦你將一個元件命名為 **thingy**，就不要將其重新命名為 **thingamabob**。例如，以下段落錯誤地將 **Protocol Buffers** 重新命名為 **protobufs**：

> Protocol Buffers (or protobufs for short) provide their own definition language. Blah, blah, blah. And that's why protobufs have won so many county fairs.
>
> Protocol Buffers 提供了他們自己的定義語言。……（中間省略），這就是為什麼 protobufs 能在這麼多科技博覽會上勝出的原因。

Google 軟體工程師 George Fairbanks 提供了關於一致命名的絕佳註解：

> [!QUOTE]
> 當我遇到兩個似乎是同義詞的詞時，我會懷疑作者是否試圖暗示某種細微的差異，要我留意追查和理解。

是的，技術寫作確實很殘酷，限制也很多，但至少它提供了一個很好的變通方法。也就是說，在介紹一個冗長的概念名稱或產品名稱時，你可以指定該名稱的縮寫版本。然後，你可以在整個文件中使用該縮寫名稱。例如，以下段落是沒有問題的：

> **Protocol Buffers**（簡稱 **protobufs**）提供了自己的定義語言。……（中間省略），這就是為什麼 protobufs 能在這麼多科技博覽會上勝出的原因。

## 正確使用首字母縮寫

在文件或章節中首次使用不熟悉的首字母縮寫（acronym）時，請先寫出全稱，然後將縮寫放在括號中。全稱和縮寫都以粗體顯示。例如：

> This document is for engineers who are new to the **Telekinetic Tactile Network (TTN)** or need to understand how to order TTN replacement parts through finger motions.
>
> 本文件適用於剛接觸 **Telekinetic Tactile Network (TTN)** 的工程師，或需要瞭解如何透過手指動作訂購 TTN 替換零件的人員。

之後你便可以使用該縮寫，如下例所示：

> If no cache entry exists, the Mixer calls the **OttoGroup Server (OGS)** to fetch Ottos for the request. The OGS is a repository that holds all servable Ottos. The OGS is organized in a logical tree structure, with a root node and two levels of leaf nodes. The OGS root forwards the request to the leaves and collects the responses.
>
> 如果快取中沒有條目，混合器會呼叫 **OttoGroup 伺服器 (OGS)** 來為請求獲取 Ottos。OGS 是一個存放所有可提供 Ottos 的儲存庫。OGS 以邏輯樹狀結構組織，有一個根節點和兩層葉節點。OGS 根節點將請求轉發到葉節點並收集回應。

不要在同一份文件中來回交替使用縮寫和全稱。

### 使用縮寫還是全稱？

當然，你可以正確地介紹和使用首字母縮寫，但你應該使用縮寫嗎？嗯，縮寫確實可以縮短句子長度。例如，TTN 比 Telekinetic Tactile Network 少了兩個詞。然而，縮寫實際上只是一層抽象；讀者必須在腦中將最近學到的縮寫擴展為全名。例如，讀者在腦中將 TTN 轉換為 Telekinetic Tactile Network，所以「較短」的縮寫實際上比全名需要更長的時間來處理。

大量使用的首字母縮寫會發展出自己的身份。在出現多次後，讀者通常會停止將縮寫擴展為全名。例如，許多網頁開發人員已經忘記 HTML 的全名是什麼。

以下是首字母縮寫詞的使用準則：

* 不要定義只會使用幾次的縮寫詞。
* 若同時滿足以下兩個條件，便可以定義縮寫詞：
  * 縮寫明顯短於全稱。
  * 縮寫在文件中出現多次。

### 練習 {#ex1}

> Jeff Dean invented MapReduce in 1693, implementing the algorithm on a silicon-based computer fabricated from beach sand, wax-paper, a quill pen, and a toaster oven. This version of MR held several world performance records until 2014.

(Please note that the preceding passage is meant to be humorous, not factual.)

請修正以下段落。假設這是文件中第一次出現 **MapReduce**，且 **MR** 是最適合的縮寫：

> Jeff Dean 於 1693 年發明了 **MapReduce**（簡稱 **MR**），並在一台以沙灘砂、蠟紙、羽毛筆和烤麵包機製成的電腦上實作了這個演算法。這個版本的 MR 曾保持多項世界效能紀錄，直到 2014 年為止。

（請注意，以上內容純屬幽默，並非事實。）

<br>

{{< bs/collapse heading="點我看答案" expand=false >}}
你可以採用幾種不同的方式來處理這段文字。一種做法是將縮寫 *MR* 與完整術語建立關聯，然後在後文使用該縮寫：

> Jeff Dean 於 1693 年發明了 **MapReduce**（**MR**）……這個版本的 MR 曾保持多項……

另一種做法是，考量在這麼短的段落中定義縮寫會增加讀者的負擔，故每次都使用完整術語 *MapReduce*：

> Jeff Dean 於 1693 年發明了 **MapReduce**……這個版本的 MapReduce 曾保持多項……

順帶一提，一位更講究的技術寫作員也會將「沙灘砂、蠟紙、羽毛筆和烤麵包機」轉換成項目符號清單。不過，這就是另一堂課要講的內容了。
{{< /bs/collapse >}}

---

## 辨識模稜兩可的代名詞

許多代名詞指向先前介紹過的名詞。這樣的代名詞類似於程式設計中的指標。就像程式設計中的指標一樣，代名詞也容易引發錯誤。不當使用代名詞會在讀者的腦中造成相當於空指標錯誤的認知問題。在許多情況下，你應該直接避免使用代名詞，而只是重複使用名詞。然而，有時代名詞的實用性超過了其風險（如此句所示）。

請考慮以下代名詞指南：

* 只在介紹了名詞之後才使用代名詞；絕不要在介紹名詞之前使用代名詞。
* 將代名詞放置在盡可能靠近其所指的名詞的地方。一般來說，如果你的名詞和代名詞之間相隔超過五個詞，請考慮重複使用名詞而不是使用代名詞。
* 如果你在你的名詞和代名詞之間引入了第二個名詞，請重複使用你的名詞而不是使用代名詞。

### It 和 they

以下代名詞在技術文件中最容易引起混淆：

* It
* They、them、和 their

例如，在下面的句子中，**It** 是指 Python 還是 C++？

> Python is interpreted, while C++ is compiled. **It** has an almost cult-like following.
>
> Python 是直譯式的，而 C++ 是編譯式的。**它**擁有一批近乎狂熱的追隨者。

再舉一個例子，下面句子中的 **their** 是指什麼？

> Be careful when using Frambus or Carambola with HoobyScooby or BoiseFram because a bug in **their** core may cause accidental mass unfriending.
>
> 當使用 Frambus 或 Carambola 搭配 HoobyScooby 或 BoiseFram 時要小心，因為**它們的**核心中的一個錯誤可能會導致意外的大規模解除好友關係。

### This 和 that

再考慮兩個有問題的代名詞：

* This
* That

例如，在下面這個模稜兩可的句子中，**This** 可能指使用者 ID、執行該程序，或兩者皆有：

> Running the process configures permissions and generates a user ID. **This** lets users authenticate to the app.
>
> 執行該程序會設定權限並產生一個使用者 ID。**這**讓使用者可以向應用程式進行身份驗證。

為了幫助讀者理解，請避免以不清楚所指為何的寫法來使用 **this** 或 **that**。你可以採用以下任一策略來釐清模糊的 **this** 和 **that** 用法：

* 將 **this** 或 **that** 替換為適當的名詞。
* 在 **this** 或 **that** 之後立即接上一個名詞。

請視需要來替換或補充明確的詞語，如以下對範例第二句的改寫：

> This user ID lets users authenticate.
>
> The process of configuring permissions lets users authenticate.
>
> The combination of permissions and a user ID lets users authenticate.

### 練習 {#ex2}

請找出下列段落中所有可能的模糊代名詞所指為何。

1. Aparna and Phil share responsibilities with Maysam and Karan and they are next on call.<br/>
  （Aparna 和 Phil 與 Maysam 及 Karan 共同分擔責任，而他們是下一個待命的人。）
2. You may import Carambola data via your configuration file or dynamically at run time. This may be a security risk. <br/>
  （您可以透過設定檔匯入 Carambola 資料，或在執行時動態匯入。這可能會帶來安全風險。）

<br>

{{< bs/collapse heading="點我看答案" expand=false >}}

(1) 代名詞 **they** 可能指：

* Aparna 和 Phil
* Maysam 和 Karan
* Aparna、Phil、Maysam 和 Karan（全部）
* 任一個人，作為單數中性的「they」 <br>

(2) 代名詞 **this** 可能指：

* 透過設定檔進行匯入
* 在執行階段動態匯入
* 兩者皆是

{{< /bs/collapse >}}

<br/><br/>

**下一單元：** [主動語態 vs. 被動語態]({{< relref "active-voice.md" >}})
