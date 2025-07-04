---
title: 清晰的句子
linkTitle: 清晰的句子 (10 min)
weight: 5
---

喜劇作家追求最有趣的結果，恐怖作家力求最嚇人的情節，而技術寫作人員則以最清晰為目標。在技術寫作中，清晰度優先於所有其他規則。本單元將建議幾種讓您的句子變得優美清晰的方法。

## 選擇強而有力的動詞

許多技術寫作人員認為動詞是句子中最重要的部分。選對了動詞，句子的其餘部分就會水到渠成。不幸的是，有些作家只重複使用一小部分溫和的動詞，這就像每天用不新鮮的餅乾和濕軟的生菜招待客人一樣。挑選正確的動詞需要多花一點時間，但會產生更令人滿意的結果。

為了吸引和教育讀者，請選擇精確、有力、具體的動詞。減少使用不精確、軟弱或通用的動詞，例如：

* **be 的各種形式**：is、are、am、was、were 等。
* **occur** (發生)
* **happen** (發生)

例如，思考一下如何透過強化以下句子中的弱動詞來點燃一個更引人入勝的句子：

| 弱動詞 | 強動詞 |
| :--- | :--- |
| The exception **occurs** when dividing by zero. (當除以零時，會**發生**例外。) | Dividing by zero **raises** the exception. (除以零會**引發**例外。) |
| This error message **happens** when... (當...時，這個錯誤訊息會**發生**。) | The system **generates** this error message when... (當...時，系統會**產生**這個錯誤訊息。) |
| We **are** very careful to ensure... (我們**是**非常小心以確保...) | We carefully **ensure**... (我們小心地**確保**...) |

許多作家依賴 be 的各種形式，好像它們是調味架上唯一的香料。撒上不同的動詞，看著您的文章變得更開胃。話雖如此，be 的形式有時是動詞的最佳選擇，所以不要覺得您必須從寫作中刪除 be 的每一種形式。

請注意，通用動詞通常也預示著其他問題，例如：

* 句子中不精確或遺漏了行為者
* 被動語態的句子

## 減少使用 there is / there are

以 `There is` 或 `There are` 開頭的句子，是將一個通用的名詞嫁給一個通用的動詞。通用的婚禮會讓讀者感到厭煩。透過提供一個真實的主詞和一個真實的動詞，來向您的讀者展現真愛。

在最好的情況下，您可以直接刪除 `There is` 或 `There are`（以及句子後面可能的一兩個詞）。例如，思考以下句子：

> **There is** a variable called `met_trick` that stores the current accuracy.
> (**有一個**名為 `met_trick` 的變數，它儲存了當前的準確度。)

移除 `There is` 會用一個更好的主詞取代通用的主詞。例如，以下任一句子都比原句更清晰：

> A variable named `met_trick` **stores** the current accuracy.
> (一個名為 `met_trick` 的變數**儲存**了當前的準確度。)

> The `met_trick` variable **stores** the current accuracy.
> (`met_trick` 變數**儲存**了當前的準確度。)

您有時可以透過將真正的動詞和主詞從句尾移到句首來修復 `There is` 或 `There are` 的句子。例如，請注意在以下句子中，代名詞 `you` 出現在句尾附近：

> **There are** two disturbing facts about Perl you should know.
> (關於 Perl，**有**兩個令人不安的事實你應該知道。)

用 `You` 取代 `There are` 可以強化句子：

> **You should know** two disturbing facts about Perl.
> (**你應該知道**關於 Perl 的兩個令人不安的事實。)

在其他情況下，作者以 `There is` 或 `There are` 開頭，以避免創造真正主詞或動詞的麻煩。如果不存在主詞，請考慮創造一個。例如，以下的 `There is` 句子沒有指明接收實體：

> **There is** no guarantee that the updates will be received in sequential order.
> (**無法**保證更新將按順序接收。)

用一個有意義的主詞（例如 `clients`）取代 `There is`，可以為讀者創造更清晰的體驗：

> **Clients might not receive** the updates in sequential order。
> (**客戶可能不會**按順序接收到更新。)

## (可選) 盡量減少使用某些形容詞和副詞

形容詞和副詞在小說和詩歌中表現得非常出色。多虧了形容詞，普通的草地變成了「揮霍的」和「翠綠的」，而毫無生氣的頭髮則變成了「有光澤的」和「茂盛的」。副詞推動馬兒「瘋狂地」和「自由地」奔跑，狗兒「大聲地」和「兇猛地」吠叫。不幸的是，形容詞和副詞有時會讓技術讀者大聲而兇猛地吠叫。這是因為形容詞和副詞對於技術讀者來說，往往定義得太鬆散和主觀。更糟的是，形容詞和副詞會讓技術文件聽起來像危險的行銷材料。例如，思考以下技術文件中的段落：

> Setting this flag makes the application run screamingly fast.
> 
> 設定此旗標可讓應用程式執行得**飛快**。

誠然，「飛快」會引起讀者的注意，但不一定是好的方面。給您的技術讀者提供事實數據，而不是行銷術語。將無定形的副詞和形容詞重構為客觀的數字資訊。例如：

> Setting this flag makes the application run 225-250% faster.
> 
> 設定此旗標可讓應用程式執行速度**加快 225-250%**。

前面的改變是否剝奪了句子的一些魅力？是的，有一點，但修改後的句子獲得了準確性和可信度。

**注意**：不要把傳授知識（技術寫作）與宣傳或銷售產品（行銷寫作）混淆。當讀者期待的是學習，就應該提供知識性的內容；不要在教學內容中夾帶宣傳或銷售的資訊