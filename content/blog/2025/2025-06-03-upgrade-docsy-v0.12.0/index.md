---
title: "網站更新：Docsy 升級至 v0.12.0"
slug: "upgrade-website-to-docsy-v0.12.0"
date: 2025-06-03
tags: ["hugo", "docsy"]
---

更新日期：2025-06-03

## 主要更新

- Docsy 從 v0.11.0 升級至 v0.12.0。詳見 [Docsy changelog 0.12.0](https://www.docsy.dev/project/changelog/#0120)。
- Hugo 從 v0.138.0 升級至 v0.147.7。詳見 [Hugo News](https://gohugo.io/news/)，需特別注意的是 [v0.146.0](https://github.com/gohugoio/hugo/releases/tag/v0.146.0)，因為此版本有重大更新，而且是 breaking changes。

## 更新細節

底下是升級過程中，對個別檔案和資料夾進行的的修改。

### go.mod

網站根目錄下的 go.mod 檔案，在執行 `go mod tidy` 命令之後，自動更新了相關套件：

```go
require (
	github.com/google/docsy v0.12.0 // indirect
	github.com/hugomods/bootstrap v0.23.0 // indirect
	github.com/hugomods/icons/vendors/bootstrap v0.5.11 // indirect
	github.com/hugomods/icons/vendors/font-awesome v0.6.12 // indirect
)
```

### CI/CD Pipeline

檔案：`/.github/workflows/gh-pages.yaml`

```yaml
jobs:
  # Build job
  build:
    runs-on: ubuntu-latest
    env:
      HUGO_VERSION: 0.147.7
```

### Hugo 組態檔

```toml
[module.hugoVersion]
extended = true
min = "0.146.0"
```

在本機執行 `hugo server -D` 命令，結果出錯：

```text
hugo: downloading modules …
hugo: collected modules in 37848 msError: html/template:_partials/head.html:23:46: no such template "partials/page-description.html"
```

這是因為 Hugo v0.146.0 修改了一些重要的資料夾名稱，例如 `/layout` 資料夾底下：

- `shortcodes` 改名為 `_shortcodes`，而且 `/layout` 之下的任何子目錄都可以有這個資料夾。
- `partials` 改名為 `_partials`，而且 `/layout` 之下的任何子目錄都可以有這個資料夾。
- 刪除了 `_default` 資料夾。原本該資料夾底下的檔案和子目錄（例如 `_markup`）全都移至上一層。也就是說，`/layouts` 目錄下的 HTML 檔案就是 default templates。

> 我覺得這樣的變動是好的，因為資料夾的結構更合理，也更容易從名稱分辨其用途。

只要按照上述變動來修改資料夾名稱和相關的檔案，應該就能解決升級 Hugo 版本之後發生的錯誤。以剛才的錯誤訊息為例，由於我的網站有一個自訂的 head.html 放在 `/layouts/partials` 目錄下，而且裡面有使用另一個 template: `partials/page-description.html`，故修正方法為：

- 把 `partials` 目錄改名為 `_partials`。
- 接著再修改 `_partials/head.html` 檔案內容，把其中的 `partials/` 字串改為 `_partials/`。

至於其他詳細變動，請參考官方文件：[New template system in Hugo v0.146.0](https://gohugo.io/templates/new-templatesystem-overview/)。

## 附註：企業內網的更新程序

在有防火牆隔離的企業內網，可能會無法在建置網站的流程中自動從 Internet 下載相依套件。這種情況，我是預先以其他方式下載相依套件，然後把它們加入網站的 repository，也就是預先在 local 準備好所有需要的檔案，並且修改相關的組態檔案，把所有指向外部公開資源的 URLs 改為 local URLs。此過程有些繁瑣操作，容易忘記，所以在 GitHub 建立了一個展示用的專案，以便在未來碰到問題時可以參考。專案的網址如下：

<https://github.com/huanlin/docsy-example-standalone/>

預先加入的相依套件主要是：

- [/themes 資料夾](https://github.com/huanlin/docsy-example-standalone/tree/main/themes) 底下的所有子目錄，例如 docsy、twbs/bootstrap、FortAwesome/Font-Awesome。
- [/node_modules 資料夾](https://github.com/huanlin/docsy-example-standalone/tree/main/node_modules)底下全部的套件。

組態檔裡面有外部 URL 需要改為指向 local URL 的：

- [/hugo.toml](https://github.com/huanlin/docsy-example-standalone/blob/main/hugo.toml)（或 hugo.yaml）的 `module` 區段。
- [/themes 資料夾](https://github.com/huanlin/docsy-example-standalone/tree/main/themes) 底下所有套件的 hugo.yaml 或 hugo.yml 檔案中的 `module` 區段。可參考此檔案：[/themes/docsy/hugo.yaml](https://github.com/huanlin/docsy-example-standalone/blob/main/themes/docsy/hugo.yaml)。
