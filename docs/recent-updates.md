# 首頁「最近更新」功能維護說明

本文件說明首頁「最近更新」區塊的設定位置、文章篩選規則，以及日後調整與驗證方式。

## 檔案結構

網站設定已使用 Hugo 的 configuration directory 拆分：

```text
config/
└─ _default/
   ├─ hugo.toml    # Hugo 與主題的一般設定
   └─ params.toml  # 網站自訂功能設定
```

相關實作檔案如下：

- `config/_default/params.toml`：控制最近更新功能。
- `layouts/_partials/recent-updates.html`：查詢文章並產生區塊 HTML。
- `layouts/portal/list.html`：在首頁載入 partial。
- `assets/scss/_styles_project.scss`：`.portal-recent*` 相關版面與樣式。

由於檔名是 `params.toml`，其中的設定會合併至 Hugo 的 `site.Params`，因此不需要再包一層 `[params]`。

## 功能設定

設定集中在 `config/_default/params.toml`：

```toml
[features.recent_updates]
enabled = false
title = "最近更新"
limit = 6
date_format = "2006-01-02"
sections = [
  "dotnet",
  "docs",
  "go",
  "google-tech-writing",
]

[features.recent_updates.section_labels]
dotnet = ".NET"
docs = "Documentation"
go = "Go"
google-tech-writing = "Tech Writing"
```

各欄位用途：

| 欄位 | 說明 |
| --- | --- |
| `enabled` | 是否顯示整個最近更新區塊。 |
| `title` | 區塊標題。 |
| `limit` | 最多顯示幾篇文章。 |
| `date_format` | 畫面上的日期格式，使用 Go 的日期格式語法。 |
| `sections` | 允許出現在區塊中的 Hugo section。 |
| `section_labels` | 將 section 名稱轉換成顯示於畫面上的標籤。未設定標籤時，模板會直接顯示 section 名稱。 |

### 啟用或關閉

顯示區塊：

```toml
enabled = true
```

隱藏區塊：

```toml
enabled = false
```

關閉時，模板不會輸出該區塊的 HTML。

## 文章篩選與排序規則

`layouts/_partials/recent-updates.html` 依序執行以下處理：

1. 從 `site.RegularPages` 取得一般內容頁。
2. 排除草稿文章。
3. 排除 front matter 中設有 `hide_from_recent: true` 的文章。
4. 只保留 `sections` 設定所列的分類。
5. 依 `.Lastmod` 由新到舊排序。
6. 依 `limit` 截取指定數量的文章。

如果某篇文章不應出現在最近更新中，可在它的 front matter 加入：

```yaml
hide_from_recent: true
```

這項設定不會刪除或隱藏文章本身，只會把文章排除於首頁的最近更新區塊。

## `.Lastmod` 的來源

主設定 `config/_default/hugo.toml` 已啟用：

```toml
enableGitInfo = true
```

模板使用 Hugo 提供的 `.Lastmod` 進行排序與顯示。依 Hugo 的日期解析規則，文章 front matter 的 `lastmod` 可以提供明確日期；啟用 Git 資訊後，Git 的最後修改時間也可以成為 `.Lastmod` 的來源。

若文章需要固定、可預期的更新日期，建議在 front matter 明確設定：

```yaml
lastmod: 2026-07-21
```

## 新增或調整分類

若要加入新的 section，例如 `devops`，需同時調整允許清單與顯示名稱：

```toml
sections = [
  "dotnet",
  "docs",
  "go",
  "google-tech-writing",
  "devops",
]

[features.recent_updates.section_labels]
dotnet = ".NET"
docs = "Documentation"
go = "Go"
google-tech-writing = "Tech Writing"
devops = "DevOps"
```

## 版面行為

最近更新預設使用兩欄排列；在較窄的螢幕上會切換為單欄。每筆資料包含：

- 分類標籤
- 文章標題
- 最後更新日期

相關樣式集中在 `assets/scss/_styles_project.scss` 的 `.portal-recent*` selectors。

## 驗證修改

修改設定或模板後，在專案根目錄執行：

```sh
npm run build
```

接著確認：

1. 建置沒有模板或 TOML 語法錯誤。
2. `enabled = true` 時，首頁顯示最近更新區塊。
3. 顯示筆數不超過 `limit`。
4. `enabled = false` 時，首頁不輸出最近更新區塊。
5. 窄螢幕下清單由兩欄切換為單欄。
