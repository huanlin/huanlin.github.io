---
title: Doc as Code 與知識管理系統
slug: "doc-as-code-and-kms"
date: 2024-04-24
---

你任職的公司裡面有沒有使用現成的知識分享平台或知識管理系統？例如 [Confluence](https://www.atlassian.com/software/confluence)、[Document360](https://document360.com/) 等等。

目前我是用 markdown 編寫技術文件，搭配 Git workflow 來自動建置與發佈文件至網站。我寫的文件主要是軟體系統的 user guide 以及 API reference manual。像這類正規的產品文件，許多開放原始碼專案也都是採用類似的作法，例如 [Kubernetes Documentation](https://kubernetes.io/docs/home/)，以及[這篇筆記所在的網站](https://huanlin.cc)（使用 Hugo 建置）。

然而，使用 markdown 撰寫文件，並搭配靜態網站生成工具（如 Hugo）來建立技術文件網站，對企業內部需要知識管理的情況適用嗎？它能夠滿足企業對 KMS 的需求嗎？對此問題，我並沒有十分確定的答案，僅透過這篇筆記梳理目前的想法。

簡單起見，接下來的討論，會把 markdown + Git workflow 的文件製作方式稱為 Doc as Code，即 Documentation as Code 的簡寫。

## Doc as Code

優點：

- 採用 markdown 編寫，語法簡單、排版風格容易趨於一致。
- Markdown 是文字檔案，有許多工具可以對檔案內容進行處理。例如文字搜尋替換、整批翻譯、移轉至其他文件平台等等。
- 搭配 Git workflow 在團隊中實施 review 和 approval 程序，有助於傳遞知識、確保文件品質。
- DevOps team 本來就熟悉 Git 操作以及相關流程，故很容易一起加入文件協作的行列。如果有時間和興趣，團隊成員皆可自由貢獻至文件庫。
- 避免 vendor/platform lock-in。因為文件是以 markdown 檔案的形式保存，未來若要改用其他文件工具或平台，可輕易移轉既有內容，減少重工，從而確保先前製作文件的心血不會白費。
- 有很多開放原始碼的解決方案可供選擇。

缺點：

- Markdown 語法雖然簡單，但是對於沒寫過的人來說，還是有一點學習門檻。一旦拿來跟 WYSIWYG 介面的文件管理平台比較，親和力就輸了一截。（熟悉 markdown 的人應該會覺得 markdown 才更親切吧。）
- 搜尋能力僅限於「這個文件網站」，也就是同一個 Git repository 內的文件。如果企業內部想要建立一個統一的知識管理平台，Doc as Code 方法很難做到跨站搜尋和跨產品搜尋，除非把所有產品的文件全都放在同一個 Git repository（也就是同一個網站）。

## Full-blown KMS

現成的 KMS，我只有用過 Confluence。這裡僅以我個人有限的使用經驗來說說這類 full-blown KSM 的優缺點。

現成的 KMS 最大優點就是豐富的功能。通常有這些：

- WYSIWYG 文件編輯器，高親和力，隨時可寫。
- 所有文件全都集中在一個平台，方便管理和搜尋任何文件。
- 更好的 user feedback 與互動功能。
- 文件的權限管理功能。
- 文件的統計分析功能，例如某一篇文章的點閱次數。

就我所知，Document360 除了上述功能，還有：

- 支援 markdown 編輯器。（這個我喜歡！）
- 可製作美觀的 API reference manual。
- 可讓管理者自行定義文件的 review 與 approval 流程。

我認為支援 markdown 編輯器對 KMS 非常重要。因為如果只提供 WYSIWYG 編輯器，每個人寫出來的文件可能都有不同的排版風格，容易導致文件內容格式花俏、凌亂。

此外，文件的 review 和 approval 也很重要。因為 KMS 的操作介面通常很容易上手，且隨時瀏覽器打開頁面就能寫。這很容易產生一個現象：許多文件的品質就跟草稿差不多，例如一堆錯字、沒有為讀者設想閱讀順序，甚至各文章之間也沒有相互關聯。這些問題都和 KMS 工具本身沒有直接關聯，因為關鍵在於寫文件的人是否想要寫出易讀易懂的文件，以及作者本人是否有相關的寫作訓練或經驗。然而，KMS 工具隨時可以寫點東西，這樣的靈活性也多少是促成這種現象的間接原因。

缺點：

- Vendor/platform lock-in。一旦採用某個廠商的 KMS 解決方案，所有輸入至該平台的內容就會綁在上面。萬一日後發現這平台不好用，想換到另一個 KMS，可能要費不少工夫把既有的文件撈出來，再轉入另一個工具平台。
- 可能有額外費用，例如支付給 KMS 軟體廠商的年費。

## 該選哪個？ {#how-to-choose}

我認為現成的 KMS 平台很適合企業內部用來蒐集與管理零碎片段知識。這類片段知識通常不會太要求作者字斟句酌，只要有點到關鍵的技術 know-how，能讓讀者解決特定問題就好。雖然它也一定能製作出品質優秀的文件，但其主要關鍵還是在寫作的人。

採用 markdown 來編寫技術文件的 Doc as Code 方法，則很適合撰寫單一產品或一整套相關產品的正規文件。至於能夠在一個 Git repository 放進多少個產品的技術文件，這不好估算，我沒有一個比較客觀具體的數字。就我目前的經驗來說，把三到四個相關產品的 user guide 和 API reference manual 放進同一個 Git repository，跑 CI/CD pipeline 並不至於花太長時間（因為 Hugo 的執行速度超快），生成的網頁也沒有碰到搜尋效率不好的問題。

若在企業內部採用 Doc as Code 方法來建立各產品的文件，最終大致會是這個樣子：每一個產品文件就是一個網站，有自己的網址；各產品文件之間若有關聯，則會在文件中加入交叉參考連結。由於工具內建的 [local search 功能](https://gohugo.io/tools/search/)無法在企業內部做到跨站搜尋，也只能以交叉參考連結來把相關的產品文件串起來。

## 結語 {#conclusion}

無論採用 Doc as Code 方法還是現成的 KMS，知識管理的關鍵主角應該是內容。有了好的內容，再搭配方便強大的工具，才能真正起到傳遞知識和知識管理的作用。

怎樣寫出好的文件內容？關鍵還是在人。

Keep writing!
