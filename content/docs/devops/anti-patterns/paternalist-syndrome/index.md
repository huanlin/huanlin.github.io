---
title: Paternalist syndrome
---

> 本文整理自 Jeffery D. Smith. *Operations Anti-Patterns, DevOps Solutions*. Manning Publications, 2020. Chapter 2: "The paternalist syndrome."
>
> 翻譯方法：使用 AI 工具翻譯，再由人工修訂。請注意：有些內容可能略過，故不是完整翻譯。

**家長式綜合症**（paternalist syndrome）這個名稱源自某一群體對其他人採取類似父母般的支配關係。家長式綜合症依賴把關者（gatekeepers）來決定工作該如何以及何時完成。這種權力的集中一開始看起來像是審慎的決策，但很快就可能演變成妨礙生產力的負擔。

## 2.1 Creating barriers instead of safeguards

有時候，來自其他團隊或部門的審核流程確實能帶來價值。但更多時候，另一個團隊的介入其實是出於別的原因——彼此之間缺乏信任，以及系統本身缺乏安全性。想像一下，我叫你喝湯，卻沒給你任何餐具。我的理由是，萬一你拿牛排刀喝湯時割傷了嘴怎麼辦。這聽起來瘋狂，但這正是許多公司正在做的事情——而且這是一種怠惰的作法。
如果真的關心系統的安全性，更好的做法是給你一把湯匙，而且**只給湯匙**：這是完成這項特定任務最合適的工具。

防止使用者在不知情的情況下執行危險操作，應該是系統可用性（usability）設計的重要目標之一。由於許多系統缺乏這類防護機制，組織通常會以限制能執行這些操作的人數來補救，僅讓少數「被選中的人」有權限執行這些任務。但成為某項任務的「授權執行者」，並不代表不會犯錯。被授權的人也可能操作失誤、輸入錯誤指令，或者誤解某個動作的影響。

透過限制這類任務的存取權限，實際上並沒有真正降低風險，而只是把風險的壓力集中到少數幾個人身上。同時，此作法也把問題歸咎為「人員問題」，而不是「系統問題」：這些被授權的工程師有足夠的能力了解衝擊，而其他人則不行。當你的系統缺乏安全機制時，這種缺陷就會以各種形式表現出來——例如繁瑣的交接流程、層層審核，以及過度限制的存取控制。

以下是過度限制存取權限所帶來的一些問題：

- 團隊之間負責的領域常常有所重疊。在技術團隊中，這種職責界線的模糊會讓人難以判斷，什麼時候一位成員的責任範圍結束，另一位成員的責任又從哪裡開始。舉例來說，身為一名開發人員，如果我負責維護測試套件，那我是否有權安裝支援該測試套件所需的軟體？我是否需要事先通知運維團隊這次安裝？
- 如果這次安裝導致伺服器故障，或出現不相容的情況，那麼該由誰來負責解決問題？
- 誰擁有實際進行故障排除的權限？

當職責產生重疊，上述問題便會浮現。雖然這些流程往往是以「安全防護」的名義設立，出發點是良善的，但最終往往會演變成脫離現實、效率低落的操作。

採用審核流程時，流程會隨著每一次未被現有機制明確處理的事件而變得更複雜。每出現一次事件，流程就會新增一條附加規則，或者更糟：增加一位審核者。沒過多久，你就會發現這套流程越來越笨重繁瑣，完全無法實現它原本應有的價值。

如果你曾參加過審核會議，你對這種情況應該不陌生：一間會議室裡坐滿人，通常是管理階層，他們試圖評估某項提案的風險。但這類會議往往逐漸淪為形式化的「蓋章機器」。大多數的變更其實根本不會帶來任何負面影響，久而久之，審核的標準就會被拉低。在許多組織中，變更審核很快就不再被視為一種有價值的流程，而成了一道障礙。你親歷過這種場面——內心深處，你清楚情況就是那麼糟糕。

在許多傳統組織中，「消除人為障礙、提升團隊協作」經常被掛在嘴邊。但一旦出事，這些理念立刻被丟在一邊。傳統組織習慣性地從「審核工具箱」中拿出新的限制來預防下一次事件，而在 DevOps 組織中，團隊會盡全力抵抗這種衝動。目標始終是：**消除那些不必要的人為障礙，同時保留真正有能創造價值的防護機制。**

人為設立的關卡會在團隊之間形成一種權力不對等的關係，導致申請方和審核方形成上對下的家長式管理。申請者感覺就像小孩向父母哀求借法拉利去約會，這種權力失衡自然會引發團隊間的摩擦。

讀到這裡，你可能已經發現自己公司也出現了類似問題。但你心裡或許有個疑問：「為什麼非得導入 DevOps？」解決組織問題的方法很多，但 DevOps 之所以蔚為風潮，確實有其原因。

首先，它直指問題核心——企業文化。有些工程師可能習慣用技術手段解決一切，但只要仔細觀察組織運作，就會發現真正的痛點往往在「人」（以及他們的協作方式），而非技術。沒有任何技術能自動改善你的優先順序決策流程，也沒有一種工具能神奇地讓各團隊目標一致、停止互相扯後腿。當年電子郵件號稱能解決溝通問題，接著是手機，現在又變成即時通訊軟體——這些工具只是讓我們「更快速地進行低效溝通」而已。當然，這不是說技術在 DevOps 轉型中不重要，只不過技術層面反而是最容易的部分。

DevOps 運動的另一個重要核心，在於它特別關注「人力潛能浪費」的成本。試想一下你的日常工作，肯定有些例行公事其實可以用程式或腳本自動處理，這樣你就能騰出時間做更有價值的事。DevOps 的精神就是讓你投入的工作能為組織創造最大、最持久的效益。

這意味著，你可能得把原本五分鐘就能手動完成的小任務，改寫成耗時一週的自動化流程。雖然當下會讓提出需求的人等更久，但這將是他們最後一次等待——因為從此你再也不用浪費那五分鐘了。消弭這些日常工作中浪費的零碎時間，長期下來會帶來驚人的生產力提升。

從技術團隊的角度來看，自動化也有助於留住人才。工程師可以從枯燥的例行作業中解放，轉而投入更具挑戰性、更有趣的任務，這對團隊士氣和技術成長都是正向循環。

DevOps 的目標包含以下幾點：

- 強化團隊協作
- 減少不必要的關卡與交接
- 提供工具與權限，讓開發團隊能真正掌握自己打造的系統
- 建立可重複、可預測的流程
- 共同分擔生產環境的責任

DevOps 透過遵循 CAMS 模型（文化、自動化、衡量指標與知識共享）協助企業達成這些目標。如前一章所述（參見 1.2 節），這四大要素能為 DevOps 的成功實踐創造必要條件。

首先，工作文化與思維模式的轉變，能讓團隊成員更有效率地完成任務，同時消除無謂的流程關卡。這種文化變革將自然引領團隊走向自動化需求——自動化是建立可重複流程的強大工具，當妥善實施時，能讓團隊任何成員都能一致性地執行作業。隨著文化轉變與責任共享程度的提升，團隊需要透過衡量指標來掌握系統狀態。比起單純確認「沒有錯誤」，能夠驗證「系統正常運作」顯然更有價值。最後，知識共享機制能確保 DevOps 理念在團隊中持續擴散，而非僅局限於少數人手中。

要實現責任共享，知識共享是必要前提。若未提供適當的學習途徑，就無法要求團隊成員承擔更多責任。成員需要對系統的各部組成有基本的了解，至少要有高層次的概念，這樣才有辦法因應職責範圍的增加。你會需要建立一套機制來促進資訊共享。隨著自動化程度提升、責任分工越來越廣，過去那種「囤積資訊」（information hoarding）的心態就會逐漸瓦解。

當團隊間的每次互動都被層層審核、申請程序和權力遊戲所阻礙，這些目標根本不可能實現。這類流程活動本質上是「把關任務」，若不謹慎處理，不僅會造成不必要的延遲、引發團隊摩擦，更可能促使人們不惜代價規避關卡，有時甚至會催生次優解決方案。

> **定義**
>
> 所謂的「把關」（gatekeeping），是指某個人或某個流程成為一種人為的障礙，用來限制或控管他人對某項資源的存取權限。

「家長式綜合症」的核心正是來自於「把關者」的存在。當組織因為缺乏信任而設置了某種把關機制時，這種家長式綜合症就會產生。這種信任的缺失，可能是過去某次事件造成的，也有可能從一開始就從未建立起來。家長式綜合症的運作基礎是這樣一種觀念：**只有某個特定的人或團體具備足夠的能力與可信度，能夠執行或核准某項行動。** 這樣的設計會導致團隊之間產生摩擦，因為這個把關流程成了其他團隊完成工作的阻礙。

當這道「門檻」本身沒有實質價值時，它就會被視為一種家長式干預：申請人需要解釋、說明甚至為自己的需求辯護，才能通過審核。

## 2.2 Introducing the gatekeepers

Stephanie 在一家地方性的醫療組織擔任 IT 營運部門的工作。有一天，她收到來自帳單團隊開發人員 Terrance 的請求，希望當天下午四點部署帳單系統應用程式。Terrance 想在週末的帳單處理作業開始前套用一個修補程式，並且保留足夠的時間以便必要時能夠回滾部署。

Stephanie 瞭解了所有細節後，認為下午四點是個合理的部署時間。她經常與帳單團隊合作，知道他們通常在中午前就不再使用該應用程式了。到了四點，她開始進行部署，整個過程順利完成，沒有任何問題。她注意到在她開始部署時，還有兩個人登入系統，但這種情況不罕見，因為很多人即使用完系統後也不會立即登出。Terrance 隨後確認應用程式運作正常，他們便認定此次部署成功。

然而隔天早上，Stephanie 被叫進主管辦公室，Terrance 已經坐在那裡，神情沮喪。Stephanie 的主管告訴她，帳單團隊有幾位成員因為應收帳款部門臨時提出的請求，當天比平常更晚還在使用系統。他們正在處理大量帳單的手動更新作業，這是一個需要三個步驟的流程。而部署正好發生在這個流程中段，導致他們損失了大量珍貴的資料輸入時間，甚至還需要重新輸入資料。帳單部門的主管非常憤怒，要求公司採取措施避免類似事件再次發生。

於是，這家公司就這樣誕生了他們的第一份「變更管理政策」。

> **定義**
>
> 「變更管理」（Change Management）是指組織用來引進應用程式或系統變更的一套標準化流程。這個流程通常包括一份說明即將執行工作的申請，並提交給負責審核的管理單位，以便在指定的時段內獲得批准。

在帳單部門與 IT 部門討論完這次事件後，雙方決定未來所有的部署都必須經過正式的審查流程。Stephanie 和其他營運人員將被要求在每次部署前，獲得帳單部門的事前同意。帳單部門會內部協調，確認可接受的部署時段，但他們也提出，希望至少提前一天收到通知，好讓他們有時間取得部門內所有人的簽核。

由於 Stephanie 的工作不只支援帳單系統，因此她也需要能提前規劃整體工作安排。於是她要求所有來自開發團隊的部署請求，**必須至少提前三天提交**，這樣她的團隊才能將請求排入行程，同時也能讓帳單部門有時間收集相關簽核。

此外，團隊也一致認同，在部署時若仍有使用者登入系統，部署就必須延後，以避免造成使用者被強制登出。

經過幾輪討論後，大家達成以下流程共識：

1. 開發人員向營運部門提交變更申請（change ticket）。
2. 若申請未提前三天送出，將立即被退回，並要求開發人員重新填寫並調整日期。
3. 帳單部門會檢視他們的排程，若與其他作業衝突，則變更申請會被退回，並要求重新申請。
4. 若無衝突，變更將正式核准。
5. 執行部署作業。

這套流程的建立，表面上看起來確實緩解了各方的擔憂。但這其實是針對傳統問題的傳統解法。在 DevOps 組織裡，重點會放在**移除這類「把關式」的流程，轉而追求高效率、快速交付的能力**。這通常是透過自動化解決方案來實現，而不是新增更多流程瓶頸。

而且，這套流程還暗藏了一些潛在的副作用，會對團隊造成影響。下一段我們將深入探討這個流程的盲點與侷限。

## 2.3 Examining the gatekeepers

大家一致認為，這項新流程應該能確保同樣的問題不會再次發生，並且作為附帶效應，也能大幅改善各部門之間在變更事項上的溝通。但事實上，這項政策本身也引入了不少新的問題。由於團隊的注意力過度集中在「避免問題再發生」，反而忽略了這項流程對整體組織所帶來的額外負擔。

此外，把關者往往並不會從系統的整體觀點來思考，而是選擇「局部最佳化」，也就是只解決眼前的單一問題，但卻可能導致從整體系統角度來看出現新的問題。舉例來說，現在部署流程要求提前三天通知，這在實務上等同於帳單系統每週最多只能部署一次。這種限制讓某些緊急問題（例如 bug 修補）無法及時處理，反而可能逼得團隊必須繞過流程，直接進行部署。也就是說，**你為了解決「資料遺失風險」這個問題，設計了一個流程，卻因此犧牲了系統的靈活性與整體效率**。

再者，新流程也大幅增加了跨團隊之間的溝通需求。雖然「跟團隊成員好好溝通」本身是件好事，但當部署流程被過多的審核與回覆時間給拖慢時，整個進度可能會卡住，尤其當帳單部門回覆不即時時，問題就更明顯。這種 **「人為拖延」** 久了會讓人對整個流程感到厭煩，進而產生逃避心理。
我個人當然**從來沒**誇大過某個問題，為了能把它歸類成「緊急變更」好跳過流程，只為了週末能安心參加烤肉活動，不被 call 機打擾——但你應該能想像，有人是會這麼做的。

最後，還有一個意想不到的副作用是：**使用者的不作為也可能導致部署被取消。** 如果使用者忘記登出系統，這個無心之過就可能讓整場部署無法進行。這樣不僅會讓客戶錯過新功能或修復，對業務本身也完全沒有幫助。而且，要向他人解釋「為什麼沒部署成功」時，會變得超尷尬——「我們本來準備好了，但 Frank 沒回我們的 email，所以我們就整個取消了。」……聽起來是不是有點可笑？

如果使用者持續忘記登出，營運團隊在判斷到底哪些使用者是真正在使用系統、哪些只是單純掛著沒登出的情況時，會越來越困惑。

書中的表 2.1 會列出這項變更管理政策所帶來的新問題，幫助團隊在評估政策是否真正有效時，有更全面的思考基礎。

表 2.1 變更管理政策所衍生的新問題

| 變更項目 | 引入的問題 | 討論方向 |
|----------|-------------|-----------|
| 要求提前三天通知 | 限制帳單系統每週只能部署一次 | 這會如何影響緊急修補（bug fix）快速發布的能力？ |
| 團隊之間需要額外溝通 | 如果帳單團隊不在線上，審核流程會進一步延誤 | 帳單團隊要如何達成共識？這樣會對公司造成多少時間成本？ |
| 如果使用者還在線上，就中止部署 | 使用者可能在下班時忘記登出系統 | 營運團隊要如何判斷這位使用者是真的在用系統，還是單純忘了登出？ |


(TODO)