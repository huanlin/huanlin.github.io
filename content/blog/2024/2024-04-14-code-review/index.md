---
title: Code review
date: 2024-04-13
---

In this post:

- 為何做 code review？
- 走味的 code review。

## 為何做 code review？ {#why-code-review}

實施 code review 的兩個主要理由：提高程式品質，以及傳遞知識。

### 提高程式品質 {#better-code-quality}

Code review 可提升程式碼的品質，包括可讀性、可維護性等等，並減少應用程式的 bug。

程式碼的品質提高，洩漏至產品的 bug 數量減少，便能降低未來產品上線後的問題反饋、排查、debug、和部署新版本的成本。

### 促進了解，傳遞知識 {#sharing-knowledge}

- 新人可透過 code review 加速上手。
- 減少知識壁壘。code review 過程是公開透明的，所有團隊成員都能看到。
- 透過公開的溝通與討論，可了解其他人的想法，互相傳遞經驗與知識。

---

有了上述優點，實踐 code review 的團隊理論上應該會更快樂。然而實際上是否如此，恐怕要打個問號，因為在實施 code review 過程中還有許多「眉角」，一不小心就會走偏。接著就來看看有哪些需要注意的地方。

## 走味的 code review {#bad-smells}

這裡談一下走味的 code review，意思是實踐 code review 而導致的一些扭曲現象或者 anti-patterns。

### 馬馬虎虎 {#careless}

當團隊中的每一位成員忙於衝刺自己手上的工作，很可能沒有時間和心力去看別人寫的 code 並提出建議。也許今天比較有空，就看一下別人的 pull request；若沒空，就隨便瞄幾眼，或者跳過。（反正還有其他人會看，應該吧？）

> GitHub 的 pull request (PR)，在 GitLab 叫做 merge request (MR)。以下用 PR 來簡稱 pull request 或 merge request。

想像一下，某一個 PR 裡面包含上千行的程式碼，變動橫跨二、三十個檔案，這樣的 PR 要花多少心力和時間去仔細 review？

於是可能出現一種情形：少量變動的 PR，容易獲得其他人的 review 建議；<mark>若 PR 變動量大到一個程度，review 建議可能就只有四個字母：LGTM。</mark>也就是 looks good to me 的意思。（下圖取自 [Reddit 論壇](https://www.reddit.com/r/ProgrammerHumor/comments/w92k2i/lgtm/)）

![](images/lgtm.png#center)

然後，團隊成員可能會看到越來越多的 LGTM。至於有沒有真的看程式碼，就自由心證了。

### 合作無間 {#buddy-system}

在強制要求 code review 的團隊中，可能會在 CI/CD pipeline 裡面加入一個卡控規則：至少要有另一個人的 review 和 approve，CI/CD pipeline 才能往下執行。

這可能演變出一種行為：兩個好朋友之間彼此搭配合作，讓 PR 迅速過關。實際的操作可能是發私訊給夥伴，跟對方說：

「哈囉，我剛剛發了 PR，可以幫我 approve 一下嗎？」

然後，可能就會有某一位團隊成員感到納悶：「為什麼有些人的 PR 總是一個小時內就 approve 了，我的 PR 卻常常要等一兩天？」

### 重重關卡 {#complex-approval}

想像開發團隊分別為 testing、staging、和 production 三種環境訂下 code reivew 的關卡：

- 部署應用程式的新版本至 testing 環境時，需要 developer 和 project manager 的 review。
- QA 工程師在測試機完成測試之後，也需要 project manager 的核准才能將新版程式推送至 staging 環境。
- 最後，程式要部署至 production 環境時，則需要 project manager 和 CTO 都核准才能放行。

儘管這些關卡提供了更安全的卡控，但整個開發、測試和部署流程非常倚賴 project manager 這個角色，可能成為瓶頸，替原本已經有點笨重的流程平添更多阻礙。

### 目標持續移動 {#moving-target}

收到通知被指定為某個 PR 的 reviewer 之後，你花了時間認真且仔細地閱讀程式碼，然後提出了一些修改建議。就在你覺得差不多完成 review 工作時又收到通知，這個 PR 有新的 commits 推上來。因此，你又得重頭檢查一遍。

再次確認新的變動沒有其他問題之後，趕緊點下 "Approve" 按鈕，結果這個 PR 的狀態卻還是「需要 review」的狀態。怎麼回事？原來又有新的 commits 推上來。

像這種 review 過程仍有新的變動不斷推送至 repository 的情形，容易令 reviewer 感到挫折與煩躁，引發內心小劇場：「怎麼不事先確定程式碼都寫好了才發 PR 呢？」

### 粗魯的評論 {#mean-review}

透過 code review 所獲得的修改建議都一定是好的、對的嗎？不一定。我們甚至可能會收到到一些粗心的，甚至粗魯的 review 建議。

以下列舉幾種可能碰到的情形：

1. 僅挑枝微末節的毛病，卻忽略其中有更重要、需要修改的地方。
2. 凸顯存在感和個人貢獻，把焦點往自己身上拉。
3. 忽略客觀與其他面向，只要跟自己習慣不同就要求必須改。
4. 只要求改，不提供具體理由。
5. 陷入面子之爭。

對別人的工作產出提供建議，我覺得是有點嚴肅的事情，而且很容易冒犯到對方。故提出 review 建議的人，我想還是應該盡到說明的責任——除非錯誤非常明顯且毫無爭議（例如拼字錯誤）。提供連結指向 coding standard 或 best practice 也是 ok 的，但如果都不解釋原由，就不大好了。以下舉兩個比較誇張的例子，以凸顯什麼是令人感覺武斷、粗魯的 review 建議：

- This code looks like shit. (:thumbdown)
- Rename the `PowerOff` method to `TurnPowerOff`.

第一個例子很明顯，「looks like shit」是很粗魯的話，更別提如果只有簡單寫這句，毫無建設可言，那就是純粹的罵人，來找碴的。

第二個例子則有可能是好的，也可能是糟糕的建議。但如果只寫那樣一句，完全不說明任何理由，便可能讓接受該評論的原作者感到不被尊重，引起一些質疑：

- 「我的方法命名完全沒有違反團隊的 coding standard，為什麼一定要改成你習慣的名稱？」
- 「平日趕進度已經很忙，很多功能和品質方面的問題都等著處理，現在竟然要花時間去討論這種無關痛癢的小地方？」

Code review 是知識傳遞與彼此溝通想法的過程，其中必然有溝通的藝術成分，這點需要放在心上，以免因為忙碌而忘了基本的尊重。在提出 review 建議之前，不妨先站在對方的立場想一下他為什麼要這樣寫，而不那樣寫。如果想不到理由，或有任何不確定的地方，可以先詢問對方：

「請問這段程式碼這樣寫的原因是因為 XXX 嗎？或者是其他想法？我覺得如果改成這樣....」

當然啦，接受建議的那一方也不要過於敏感、太容易感到被人冒犯。如果是建設性的建議，保持開放的心胸與對方討論，虛心學習，會是比較好的。

## 結語 {#conclusion}

一般而言，code review 能提升程式碼與軟體產品的品質，降低日後的維護成本，也能促進團隊成員之間的溝通與知識傳遞。但實際上做起來仍有許多細節要留意，以免走偏了，甚至引發團隊成員之間的不愉快。其中的溝通藝術，是我認為最重要的部分。

總之，**若有爭議，時時想著「什麼才是對產品最好的」**，並以此前提來思考和回應。

---

這篇筆記就先寫到這裡吧。日後想到什麼再來補。

Keep coding.
