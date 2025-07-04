---
title: 文件
linkTitle: 文件 (10 min)
weight: 11
---

保持文件組織化，使用「集合（collections）」功能，根據喜好儲存和分類內容。

**預估時間：** 10 分鐘

你可以撰寫句子，也可以撰寫段落。你能否將這些段落整理成一份合乎邏輯的文件？

---

## 定義你的文件範圍（Scope）

好的文件會在一開始就定義其範圍。例如：

> 本文件描述「Frambus 專案」的設計。

更進一步的版本會明確指出**非範圍**——即那些讀者可能會以為應該包含，但實際不在本文件涵蓋範圍內的內容。例如：

> 本文件不說明與其相關的技術「Froobus 專案」的設計內容。

範圍與非範圍聲明不僅對讀者有益，對作者你也很重要。當寫作內容偏離範圍聲明（或闖入非範圍聲明）時，需要重新聚焦或修改你的範圍聲明。在審閱初稿時，刪除任何無法協助達成範圍聲明目標的章節。

### 練習

請指出以下段落有什麼問題：

> 本文件說明如何使用 Frambus API 來建立、更新與發布 Fwidgets。本文件不說明如何使用 Frambus API 刪除 Fwidgets，也不涵蓋 Linux 作業系統的歷史。

答案：非範圍部分僅應包含使用者合理期待本文件會涵蓋的資訊。沒有合理的使用者會期待本文件會涉及 Linux 的歷史。

---

## 說明你的讀者（Audience）

好的文件會明確指出其讀者對象。例如：

> 本文件針對以下族群：
>
> * 軟體工程師
> * 專案經理（PM）

除了讀者的職位，好的讀者聲明還會指出任何先備知識或經驗。例如：

> 本文件假設你已理解矩陣乘法與反向傳播的基本原理。

在某些情況下，讀者聲明還應指出先前必讀的資源或課程。例如：

> 在閱讀本文件之前，你必須先閱讀《Froobus 專案：新希望》。

---

## 在開頭總結重點

工程師和科學家通常很忙，不一定會讀完一份長達 76 頁的設計文件。假設同事們可能只讀第一段，因此務必在文件開頭就回答讀者最重要的問題。

專業寫作者會將大量精力放在第一頁，以提高讀者願意繼續閱讀的可能性。但第一頁常是最難寫的。因此要準備多次修改第一頁。

---

## 比較與對比（Compare and contrast）

在你的職涯中，很少會撰寫含有真正革命性概念的文件；大部分都是漸進式改良，建立在現有技術與概念之上。因此，請將你的想法與讀者已熟悉的概念進行比較與對比。例如：

> 這個新應用程式與 Frambus 應用程式相似，但圖形效果更佳。

或：

> Froobus API 處理與 Frambus API 相同的使用案例，但 Froobus API 更易使用。

### 練習

指出以下介紹段落有什麼問題：

> Frambus 天氣應用程式 v2 推出十項功能，是 v1 所不具備的。最重要的是，v2 提供兩週天氣預報，而 v1 只有一週預報。潮汐資訊不會改變。

答案：最後一句（關於潮汐的部分）不夠重要，不應出現在開頭。因為第一句提到有十項新功能，讀者自然會期待聽到更多，然而最後一句卻談及不是新功能的內容。

---

## 為你的讀者撰寫內容（Write for your audience）

本課程反覆強調定義讀者的重要性。本節聚焦於如何運用讀者定義來組織文件。

### 定義讀者需求

回答下列問題能幫助你確定文件內容：

* 你的目標讀者是誰？
* 他們的目標是什麼？為何要閱讀這份文件？
* 在閱讀之前，他們已經具備什麼知識？
* 閱讀之後，他們應該知道或能做什麼？

例如，假設你發明一種新的排序演算法，它與快速排序（quicksort）相似。以下是可能的回答：

* 目標讀者：本組織的軟體工程師。
* 讀者的目標：想找到更高效的排序方法，並閱讀本文判斷這種方法是否值得實作。
* 讀者的先備知識：會寫程式，曾學習排序演算法，包括快速排序，但多年未實作或評估排序演算法。
* 閱讀後能做什麼：

  * 理解此演算法與快速排序的比較與差異。
  * 辨識本演算法在兩種類型資料集上的性能優勢。
  * 以任一編程語言實作此演算法。
  * 辨識兩種執行效率不佳的邊界情況。

### 根據讀者需求組織文件

在定義完讀者需求後，將文件組織得有助讀者取得所需資訊。例如，根據上述答案，文件大綱可如下：

1. 演算法概述

   * 與快速排序比較，包括 Big O 複雜度比較
   * 演算法適用的最佳資料集
2. 演算法實作

   * 偽代碼（pseudocode）實作
   * 實作建議與常見錯誤
3. 深入分析

   * 邊界情況
   * 已知的未知（known unknowns）

---

## 下一單元：「標點符號（Punctuation）」，為選修單元

---

**內容授權**

除非另有說明，該頁面內容採用 [Creative Commons Attribution 4.0](https://creativecommons.org/licenses/by/4.0) 授權，程式碼範例則採用 Apache 2.0 授權。Java 為 Oracle 及其關聯公司之註冊商標。

**最後更新日期：2025‑03‑28 UTC**

---

如果你想繼續翻譯更多章節，或需要進一步說明，隨時告訴我！
