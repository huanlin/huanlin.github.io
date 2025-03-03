---
title: API Design-First Approach 心得筆記
slug: my-notes-on-api-design-first-approach.html
date: 2025-03-01
tags: ["API", "Git"]
---

有機會近距離接觸 API design-first approach 的工作現場，這是我的一點觀察筆記和讀書心得。

![](images/banner.png#center)

## 專有名詞

- 這裡的 **API** 指的是 HTTP based 或 RESTful API。
- 這裡的 **API spec** 指的是採用 OpenAPI 語法所撰寫或生成的 API specification 文件。API spec 的原始檔案通常是 .yaml 或 .json 格式的檔案。API spec 檔案可透過一些現成工具直接生成 HTML 網頁，方便相關人員隨時查閱。
- **Code-first approach** 指的是由程式碼來生成 API spec 的作法，也就是實作和設計規格都寫在程式檔案裡。
- **Design-first approach** 指的是先由 API designer 編寫 API spec，然後程式開發人員再按照規格實作。
- **Stakeholder** 指的是跟專案有關的人員，包括 PM、SA/SD、Developers 等等。

## 簡介

如果 API spec 是直接寫在程式碼檔案裡（需要在程式碼當中額外寫一些 annotations），通常不會有實作和設計規格不一致的問題，因為 API spec 隨時可用工具自動生成。換言之，code-first approach 大致不會有 API spec 跟實作不匹配的問題，故這裡不討論。

比較會有問題是 design-first approach，因為程式碼的實作很容易跟 API spec 不一致，以至於經常有這類疑義：到底哪那個才是對的？誰才應該是 the single source of truth？

## Design-first approach 的實務作法

採用 design-first approach 的團隊成員需要先有此共識：API spec 即是 the single source of truth（以下簡稱 **SSoT**）。

欲落實 API spec 是 SSoT，需仰賴團隊成員彼此的溝通與反饋：

- 當程式按照 API spec 來實作的時候碰到問題，必須將此問題反饋給 API designer，以便根據實際需要來更新 spec，並與實作保持一致。
- 一旦 API spec 有變動，必須有一個流程能夠確保相關人員（stakeholders）知道 API 有哪些變動，以便採取相應的行動。例如：負責寫教學文件和使用手冊的人可得知哪些文件需要一併更新。
- Stakeholders 最好也能有某種固定管道或程序來提交對 API spec 的變動需求。

欲滿足上述需求，一個常見的做法是以 Git repository 來作為 API spec 的儲存空間，並利用 GitHub（或 GitLab、BitBucket 等等）的 Pull Request（又稱為 Merge Request）功能來提交變更。這通常需要一些配套措施：

- 開發團隊成員與 stakeholders 皆可提交 Pull Request，且每一個 PR 必須至少有另一位成員的 review/approve 才能將變更合併至主要分支。
  說明：依專案和團隊規模而定，approval 的人數有時會規定得更多，例如至少三個人 vote 同意才能合併。PR 與 approval 的用意在於盡量讓相關人員能夠得知 API spec 即將有變動。

- Stakeholders 主動發現 API spec 的版本變動差異。
  說明：有些 stakeholders 不見得會經常參與 Pull Request 的 review 和 approvals 程序（例如負責撰寫使用者手冊的 technical writers）。像這種情形，他們可以從 API spec repository 建立一個自己的 branch，然後在需要比對 API spec 版本差異時，透過 git pull 操作來獲取更新，並利用 git 工具來比對 API spec 的版本差異。如此便可得知從上一次檢視 API spec 之後多了哪些改動。（僅靠開發團隊主動通知 stakeholder 相關的 API spec 變動是不大實際的，很容易遺漏和忘記）

> **掉書袋：** 根據 Conway 定律，軟體系統的結構會反映組織的內部溝通結構。

## Workflow

Design-first approach 的工作流程包含這幾個工作項目：

- 查看最新的 API spec
- 建議變更
- Review 與接受變更
- 比較版本差異
-
以下分別說明。

### 查看最新的 API spec

所有 stakeholders 都要能夠方便查看最新的 API spec，無論是透過 Git repo 的主分支，還是查看由 API spec 生成的網頁。

### 建議變更

建議 API 變更的步驟：

1. 開立工作單（例如 Jira ticket），建立新的 branch，然後在此分支進行變更。
2. 建立 Pull Request 來將變更合併至主分支。
3. 把相關人員加入此 PR 的 reviewers。

### Review 與接受變更

步驟：

1. 收到 PR 通知的人（被指定為 reviewers 的人）查看變更的內容，並提供 feedback。
2. 若沒問題，approve 此 PR（投認同票）。
3. Reviewers 的認同票數足夠時，提出 PR 的人（或有權限執行合併的人）將此變更合併。

### 比較版本差異

跟自己的副本（working copy or local copy）比較差異：

1. 利用 Git 差異比對工具來查看主分支跟自己的分支有何差異。
2. 把那些需要採取行動的差異記下來，成為自己的待辦事項。
3. 把主分支的變更合併至自己的分支。

### 教學文件如何跟上 API spec 的演進？

無論是 code-first 還是 design-first，到了寫給使用者看的 API 教學文件時，都會面臨同樣的問題：教學文件如何跟上 API spec 的變化？

也就是說，technical writer 要如何跟上 API spec 的變動，以確保文件的內容與目前的 API 規格與實作一致？若只是依賴開發團隊告知，總是難保遺漏和遺忘，故仍需要 technical writer 去主動發現和跟進 API 的演進。實際上要怎麼做呢？以下僅提供一點個人心得：

- 開發團隊若有任何針對 end user 舉辦的新版本功能展示會議，都要盡量參與，因為這些會議的內容往往有很棒的文件素材。
- 如前面提過的，從 API spec 的 git repository 開一個 branch，透過 git 差異比對來得知 API spec 的變動。
- 建立一組自己的 API 測試。這些測試主要目的是為了文件的編寫工作，故不需要是完整的測試。除了可以確認自己對 API 的理解，也能透過重複執行測試來發現某些 API 規格有了變動（例如原先寫好的測試突然有一天不能正常執行了）。
- 諮詢開發團隊，經常與他們保持密切聯繫。不確定的地方，別不好意思提問和確認。（提問之前也要自行做好各種檢查，以免頻繁提出「小白」等級的問題。）
- 提供文件的反饋機制，以便任何發現文件有錯誤的人能夠向作者反映。到了這一步，已經算是亡羊補牢了——使用者已經發現文件有錯誤，作者只能趕緊修正和改善文件內容。

## 結語

跟 code-first approach 相比，design-first approach 的溝通成本明顯高出許多，不易落實和長期堅持下去，因為此作法至少牽涉三個階段的作業：API 設計、程式開發、和文件撰寫。換言之，至少有三種角色之間需要同步，以確保各方產出一致且正確，實非易事。

相較之下，code-first approach 基本上把 API spec 和程式碼放在一起寫，故設計與實作通常不會出現分歧。剩下的，就只是想辦法讓使用者文件的編寫與發佈能夠跟上 API 的變動，這部分僅涉及文件作者和開發團隊之間的溝通，實施起來相對容易些。

最後附上一張由 AI 工具生成（[Mapify](https://mapify.so/)）、再手工修改的 mind map：

![本文摘要心智圖](images/mindmap.png#center)

## 參考資料

- [Designing APIs with Swagger and OpenAPI](https://www.amazon.com/Designing-Swagger-OpenAPI-Joshua-Ponelat/dp/1617296287) by Josh Ponelat, Lukas Rosenstock. Manning Publishing, 2022.