# 多語系內容與檔案命名

本網站使用 Hugo 原生多語系功能。目前支援的語系是繁體中文（`zh-TW`）與英文（`en`），繁體中文仍是預設語系。

## 基本原則

新建或修改內容時，檔名應明確加上語系代碼：

| 內容類型 | 繁體中文 | 英文 |
| --- | --- | --- |
| Leaf bundle | `index.zh-tw.md` | `index.en.md` |
| Section 首頁 | `_index.zh-tw.md` | `_index.en.md` |
| 一般內容檔案 | `article.zh-tw.md` | `article.en.md` |

語系後綴統一使用小寫的 `zh-tw` 與 `en`。不要使用 `zh`，因為本站設定的中文語系識別碼是 `zh-TW`，不是泛用的 `zh`。

倉庫中若仍看到 `.zh.md`，那是尚未整理的舊命名，不應作為新內容的範例；修改該內容時應一併改成 `.zh-tw.md`。

同一頁的不同語言版本應放在相同資料夾。例如：

```text
content/
└─ about/
   ├─ index.zh-tw.md
   └─ index.en.md
```

Hugo 會依相同的路徑與基礎檔名，把這兩個檔案視為同一頁的翻譯版本。放在 page bundle 內的圖片等資源也可以由兩個語言版本共用。

## 網址規則

目前的主要設定如下：

```toml
defaultContentLanguage = "zh-TW"
defaultContentLanguageInSubdir = false
```

因此網址會是：

| 語言 | 來源檔案 | 網址 |
| --- | --- | --- |
| 繁體中文 | `content/about/index.zh-tw.md` | `/about/` |
| 英文 | `content/about/index.en.md` | `/en/about/` |

中文是預設語系，所以網址不含 `/zh-tw/`。英文則位於 `/en/` 之下。

## 尚未加語系後綴的舊檔案

現有許多中文內容仍命名為 `index.md`、`_index.md` 或 `article.md`。Hugo 會把沒有語系後綴的內容歸入目前的預設語系，因此這些檔案現在仍會被視為繁體中文，既有網址也能正常運作。

不需要一次改完所有舊檔案。日後修改某個舊檔案時，應順便改成明確的中文檔名：

```text
index.md → index.zh-tw.md
```

只要繁體中文仍是預設語系，這項來源檔案改名不會改變生成的網址。

## 新增翻譯的步驟

以 About 頁為例：

1. 確認中文版使用 `content/about/index.zh-tw.md`。
2. 在相同資料夾新增 `content/about/index.en.md`。
3. 分別翻譯正文及 front matter，例如 `title`、`description`、`tagline` 和選單名稱。
4. 執行 `hugo --renderToMemory` 或專案既有的建置指令。
5. 確認 `/about/` 與 `/en/about/` 可以互相切換。

本站的語言選單只會在目前頁面確實有翻譯版本時顯示。沒有英文版的中文頁面不會自動以中文內容充當英文版，也不會提供一個會導向 404 的英文切換連結。

## 新增單一語言的內容

新文章即使暫時只有一種語言，也應明確標示語系：

```text
index.zh-tw.md  # 僅有中文版
index.en.md     # 僅有英文版
```

兩個語言的內容不必一對一存在。英文網站可以只收錄適合國際讀者的精選內容，中文網站也可以保留僅供中文讀者閱讀的文章。

## 將來變更預設語系

若日後將 `defaultContentLanguage` 改成 `en`：

- 所有沒有語系後綴的舊檔案都會被 Hugo 視為英文，因此切換前必須先完成明確的語系命名。
- 英文網址會由 `/en/...` 移至 `/...`。
- 中文網址通常會移至 `/zh-tw/...`。
- 必須規劃舊網址的重新導向與 SEO 遷移。

內容檔案現在逐步採用明確的語系後綴，可以降低未來切換預設語系時的整理成本；但是否切換仍應依英文內容的規模與既有網址影響另行評估。
