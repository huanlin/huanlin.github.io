---
title: API design-first 方法的挑戰
slug: "api-design-first-challenges"
date: 2024-06-06
tags: ["API Design"]
---

底下框框內的文字是 Copilot 幫我產生的內容（已棄用）：

{{% admonition type=note title="AI 提供的文字" %}}
API design-first 是一種強大的方法，用於構建強大且可擴展的 API。然而，它也帶來了一系列的挑戰。在這篇博客文章中，我們將探討在採用 API design-first 方法時所面臨的一些常見挑戰，並討論克服這些挑戰的策略。

1. **缺乏文檔**: 設計 API 首先需要進行詳細的文檔記錄，以確保清晰和一致性。如果沒有適當的文檔，開發人員很難理解 API 的預期功能和使用方式。我們將探討創建全面易懂且易於維護的 API 文檔的技巧。

2. **版本控制和向後兼容性**: 隨著時間的推移，API 會不斷演進，保持向後兼容性變得至關重要。我們將討論 API 版本控制和處理變更而不破壞現有集成的策略。

3. (第 3 點和第 4 點省略)

通過解決這些挑戰，我們可以充分利用 API design-first 方法的優勢，構建強大且可擴展的 API。敬請期待我們即將推出的更多見解和最佳實踐。
{{% /admonition %}}

以下是正文，純手工。

---

## 簡介 {#introduction}

API **design-first**（設計優先）方法也稱為 contract-first（合約優先），指的是先設計 API，並以特定格式的語法寫成 API spec（規格），然後再由 developers 按照 API spec 來實作。

用來撰寫 API spec 的語法，目前看到比較受歡迎的是 OpenAPI（前身叫做 Swagger），而使用 OpenAPI 撰寫的規格，可以用一些現成工具（例如 SwaggerUI 或 Redoc）轉換成網頁形式的 API 參考手冊，相當方便。

相較於 design-first，另一種作法是 **code-first**，也就是先寫程式碼，不過通常還是會有一些不那麼正規嚴謹的需求分析與設計文件。而且，採用 code-first 方法的團隊也可以用工具產生 API 參考手冊，做法是在程式碼當中寫一些特定的標註（annotations）或註解，然後用工具讀取程式碼檔案來產生 OpenAPI spec  檔案，最後再轉換成 API 參考手冊。

採用 design-first 的團隊，其最終產生的 API 參考手冊的品質主要是由撰寫 API spec 的 designer 決定。另一方面，code-first 團隊所產生的 API 參考手冊則仰賴 developer 的辛勞。

接下來只談 design-first 方法可能遭遇的挑戰，其中有一些是來自我的工作觀察。

## Design-first 實施要點 {#key-points}

當有人對 API 的用法提出這類疑問：「A 和 B 哪個才是對的？」此時應該去查看 API spec。以 spec 為準。

如此一來，有疑義時，只要看 API spec 即可，不用非得去看程式碼或跑一些測試來確認。因此，採用 API design-first 方法時，團隊成員都要認同 API spec 是 single source of truth，這點至關重要。

幾個重點：

- 使用者手冊或教學文件有時會落後、寫錯、或遺漏一些細節，但採用 design-first 方法時，API spec 應該是最正確而且完整的參考依據。
- 當 developer 按照 API spec 來寫程式的時候，若碰到規格方面的問題，一定要將問題反映給 API designer 知道，以便更新 spec。如果只修改程式碼而沒有同步更新 API spec，將造成日後生成的 API 參考手冊與 API 實際行為不符，令使用者困惑。
- 一旦 API 與實作的程式碼有新的變動，必須要有管道能讓相關人員（stakeholders）也知道 API 有哪些變動。

## Design-first 的挑戰 {#challenges}

**挑戰一**：完成多少設計才開始實作？

若花太少時間在 API design，未來可能要付出更大心力來修改 API spec 和程式碼。若花太多時間 design，則會令實作較晚起步，可能耽誤開發時程。

故只能盡量試著平衡，並理解：design 不可能一次到位，但相較於實作完成後再回頭反覆修改，在設計階段解決問題的成本還是比較低。

**挑戰二**：當 API spec 變動時，如何把這些變動有效傳遞給相關人員知道？

前面提過，developer 實作 API 時若發現 spec 有問題，一定要反映給 designer 知道，以便設計與實作保持一致。除此之外，當 spec 有變動時，還有其他相關人員可能也需要知道，例如 PM 和 Technical Writer（沒錯，寫這篇筆記也是我替自己發聲 XD）。

比較好的做法，是制定一個標準的變動程序，包括如何提出變更、接受變更，以及發布變更。

## 結語 {#conclusion}

這篇筆記的重點在於組織內的訊息溝通，亦即如何避免有人漏接相關資訊，以確保 API 的設計、實作、和相關文件都能始終維持一致。以我看到的情形來說，這還真不容易做到。

Conway 定律：軟體系統的結構會反映組織的內部溝通結構。

確實如此。

## Reference

- [Designing APIs with Swagger and OpenAPI](https://www.amazon.com/Designing-Swagger-OpenAPI-Joshua-Ponelat/dp/1617296287) by by Joshua S. Ponelat and Lukas L. Rosenstock 