---
title: 單詞
linkTitle: 單詞 (10 min)
weight: 3
---

**預計時間：** 10 分鐘

我們對文件進行了廣泛的研究，結果發現，世界上最好的句子主要由單詞組成。

## 定義新的或不熟悉的術語

在撰寫或編輯時，要學會辨識目標讀者中某些人或所有人可能不熟悉的術語。當你發現這樣的術語時，請採取以下兩種策略之一：

* 如果該術語已存在，請連結到一個好的現有解釋。（不要重新發明輪子。）
* 如果你的文件是首次介紹該術語，請定義它。如果你的文件介紹了許多術語，請將定義收集到一個詞彙表中。

## 一致地使用術語

如果你在方法的中間更改了變數的名稱，你的程式碼將無法編譯。同樣，如果你在文件的中間重新命名一個術語，你的想法也將無法在用戶的腦中「編譯」。

要點是：在整份文件中始終如一地使用相同的、明確的單詞或術語。一旦你將一個組件命名為 `thingy`，就不要將其重新命名為 `thingamabob`。例如，以下段落錯誤地將 **Protocol Buffers** 重新命名為 **protobufs**：

> Protocol Buffers (or protobufs for short) provide their own definition language. Blah, blah, blah. And that's why protobufs have won so many county fairs.
>
> Protocol Buffers 提供了他們自己的定義語言。（中間省略），這就是為什麼 protobufs 能在這麼多縣市博覽會上獲勝的原因。

Google 軟體工程師 George Fairbanks 提供了關於一致命名的絕佳註解：

> 當我遇到兩個似乎是同義詞的詞時，我會懷疑作者是否試圖暗示一個我需要追查和理解的細微差別。

是的，技術寫作是殘酷和限制性的，但至少它提供了一個絕佳的變通方法。也就是說，在介紹一個冗長的概念名稱或產品名稱時，你也可以指定該名稱的縮寫版本。然後，你可以在整份文件中使用該縮寫名稱。例如，以下段落是沒有問題的：

> **Protocol Buffers**（簡稱 **protobufs**）提供了自己的定義語言。（中間省略），這就是為什麼 protobufs 能在這麼多縣市博覽會上獲勝的原因。

## 正確使用縮寫

在文件或章節中首次使用不熟悉的縮寫時，請先寫出全稱，然後將縮寫放在括號中。將全稱和縮寫都以粗體顯示。例如：

> This document is for engineers who are new to the **Telekinetic Tactile Network (TTN)** or need to understand how to order TTN replacement parts through finger motions.
>
> 本文件適用於剛接觸 **Telekinetic Tactile Network (TTN)** 的工程師，或需要瞭解如何透過手指動作訂購 TTN 替換零件的人員。

之後你便可以使用該縮寫，如下例所示：

> If no cache entry exists, the Mixer calls the **OttoGroup Server (OGS)** to fetch Ottos for the request. The OGS is a repository that holds all servable Ottos. The OGS is organized in a logical tree structure, with a root node and two levels of leaf nodes. The OGS root forwards the request to the leaves and collects the responses.
>
> 如果快取中沒有條目，混合器會呼叫 **OttoGroup 伺服器 (OGS)** 來為請求獲取 Ottos。OGS 是一個存放所有可提供 Ottos 的儲存庫。OGS 以邏輯樹狀結構組織，有一個根節點和兩層葉節點。OGS 根節點將請求轉發到葉節點並收集回應。

不要在同一份文件中來回交替使用縮寫和全稱。

### 使用縮寫還是全稱？

當然，你可以正確地介紹和使用縮寫，但你應該使用縮寫嗎？嗯，縮寫確實可以縮短句子長度。例如，TTN 比 Telekinetic Tactile Network 少了兩個詞。然而，縮寫實際上只是一層抽象；讀者必須在腦中將最近學到的縮寫擴展為全稱。例如，讀者在腦中將 TTN 轉換為 Telekinetic Tactile Network，所以「較短」的縮寫實際上比全稱需要更長的時間來處理。

大量使用的縮寫會發展出自己的身份。在出現多次後，讀者通常會停止將縮寫擴展為全稱。例如，許多網頁開發人員已經忘記 HTML 的全稱是什麼。

以下是使用縮寫的指南：

* 不要為只會使用幾次的縮寫下定義。
* 為同時滿足以下兩個條件的縮寫下定義：
  * 縮寫明顯短於全稱。
  * 縮寫在文件中出現多次。

## 辨識模稜兩可的代名詞

許多代名詞指向先前介紹過的名詞。這樣的代名詞類似於程式設計中的指標。就像程式設計中的指標一樣，代名詞也容易引發錯誤。不當使用代名詞會在讀者的腦中造成相當於空指標錯誤的認知問題。在許多情況下，你應該直接避免使用代名詞，而只是重複使用名詞。然而，有時代名詞的實用性超過了其風險（如此句所示）。

請考慮以下代名詞指南：

* 只在介紹了名詞之後才使用代名詞；絕不要在介紹名詞之前使用代名詞。
* 將代名詞放置在盡可能靠近其所指的名詞的地方。一般來說，如果你的名詞和代名詞之間相隔超過五個詞，請考慮重複使用名詞而不是使用代名詞。
* 如果你在你的名詞和代名詞之間引入了第二個名詞，請重複使用你的名詞而不是使用代名詞。

### It 和 they

以下代名詞在技術文件中最容易引起混淆：

* It
* They, them, and their

例如，在下面的句子中，`It` 是指 Python 還是 C++？

> Python 是直譯式的，而 C++ 是編譯式的。它擁有一批近乎狂熱的追隨者。

再舉一個例子，下面句子中的 `their` 是指什麼？

> 當使用 Frambus 或 Carambola 搭配 HoobyScooby 或 BoiseFram 時要小心，因為它們核心中的一個錯誤可能會導致意外的大規模解除好友關係。

### This 和 that

再考慮兩個有問題的代名詞：

* This
* That

例如，在下面這個模稜兩可的句子中，`This` 可能指使用者 ID、執行該過程，或所有這些：

> 執行該過程會設定權限並產生一個使用者 ID。這讓使用者可以向應用程式進行身份驗證。

為了幫助讀者，請避免以不清楚其所指的方式使用 `this` 或 `that`。使用以下任一策略來闡明 `this` 和 `that` 的模糊用法：

* 將 `this` 或 `that` 替換為適當的名詞。
* 在 `this` 或 `that` 之後立即放置一個名詞。

根據需要替換或補充明確的詞語，如以下對範例第二句的改寫：

> This user ID lets users authenticate.
>
> The process of configuring permissions lets users authenticate.
>
> The combination of permissions and a user ID lets users authenticate.

### 練習

請找出下列段落中所有可能的模糊代名詞所指為何。

1. Aparna and Phil share responsibilities with Maysam and Karan and they are next on call.
2. You may import Carambola data via your configuration file or dynamically at run time. This may be a security risk.

<br>

{{< bs/collapse heading="點我看答案" expand=false >}}

(1) 代名詞 they 可能指：

  * Aparna 和 Phil
  * Maysam 和 Karan
  * Aparna、Phil、Maysam 和 Karan（全部）
  * 任一個人，作為單數中性的「they」 <br>

(2) 代名詞 this 可能指：

  * 透過設定檔進行匯入
  * 在執行階段動態匯入
  * 兩者皆是

{{< /bs/collapse >}}

<br>