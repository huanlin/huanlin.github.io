---
title: 04 模組與套件
tags: [Go]
aliases: ["04-code-organization"]
---

Go 的 **模組（module）** 與 **套件（package）** 是兩個重要觀念，因為它們決定了 Go 應用程式的組成結構與依賴關係（dependencies）。本章將介紹模組與套件的用法與相關細節。

## Modules

模組（module）是 Go 程式的基本派發（distribution）單位。一個 Go 專案——或者說一個 Git repository——通常只包含一個 module，而一個 module 是由一個或多個相關的 packages 組成（稍後會說明 package）。

> [!quote]
> A repository contains one or more modules. A **module** is a collection of related Go packages that are released together. A Go repository typically contains only one module, located at the root of the repository.
>
> Go 官方部落格：[How to Write Go Code](https://go.dev/doc/code)

Go 模組通常位於 repository 的根目錄，由一個名為 `go.mod` 的檔案來定義一個模組。換言之，一個 `go.mod` 檔案定義一個模組。此檔案的內容包括：

- 模組路徑。
- Go 版本的最低要求。
- 依賴哪些外部模組。

> [!note] 建議作法
>
> - 一個 repository 應只包含一個 module。
> - 將不同功能拆分到不同的 repository。
> - 使用 packages 來組織相關程式碼。
>
> 這樣不僅可以簡化 dependencies 的管理、程式碼更好維護，也符合 Go 的設計哲學。

建立 `go.mod` 檔案的方式，是在專案根目錄底下執行 `go mod init` 命令。比如說，專案名稱是 `todoapp`，便可使用以下命令來建立 `go.mod` 檔案：

```shell
mkdir todoapp
cd todoapp
go mod init todoapp
```

上述命令所建立的 `go.mod` 檔案，其內容會像這樣：

```text
module todoapp

go 1.23.0
```

如果這個專案的原始碼是放在 GitHub 平台上的 repository，那麼剛才的 `go mod init` 指令可以這麼寫：

```text
go mod init github.com/michael/todoapp
```

其中的 `michael` 是 GitHub 使用者帳戶。

上述命令同樣只是在當前目錄建立一個 `go.mod` 檔案，內容會變成：

```text
module github.com/michael/todoapp

go 1.23.0
```

其中的 `github.com/michael/todoapp` 即此模組的路徑名稱。

模組的路徑名稱的一個主要用途是作為該模組的唯一識別名稱，並不表示該路徑一定真的存在某個地方。如果這個專案實際上真的有放在 github.com 的指定路徑，那麼當其他應用程式要使用此模組時，Go 的模組管理工具就能自動找到並下載那個模組。此外，本機的檔案系統也不需要真的存在模組路徑，例如它可能存放在 D:/work/todoapp/ 目錄下。

假設此專案有著類似底下的目錄結構：

```text
todoapp/            -> 專案的根目錄
  ├─ go.mod         -> 定義專案的名稱和 dependencies
  ├─ main.go        -> 主程式套件
  └─ auth/          -> auth 套件
      ├─ auth.go    -> auth 相關功能的實作
      ├─ auth_test  -> auth 相關功能的測試
```

當 `main.go` 要引用（import）`auth` 套件時，它的 `import` 陳述句會像這樣：

```go
import (
    "github.com/michael/todoapp/auth"
)
```

也就是說，在引用 local folder 的套件時，`import` 的完整路徑名稱會是 `<module path>/<package path>`。此例的 module path 是 `github.com/michael/todoapp/`，而 package path 是 `auth`。

**See also:** [Golang imports tutorial](https://www.youtube.com/watch?v=Nv8J_Ruc280)

### Module paths

<mark>模組路徑是模組的正式名稱（唯一識別名稱）</mark>，宣告於模組的 go.mod 檔案，而且模組路徑要能表達該模組的用途，以及可以從何處找到它。

模組路徑通常包含三個部分：

- repository root path。
- repository 中的目錄名。
- 主要的版本編號：只有在主要版本編號為 2 或更高的版本才需要。換言之，當主要版號（major version number）晉升時，這個版號就必須成為 module path 的一部分。

範例：

```go
module example.com/mymodule
```

如果此範例的模組的版本是 v0.x.x 或 v1.x.x，那麼它的 v2.0.0 版（以及之後版本）的模組路徑就要加上主版本號，例如：

```go
module example.com/mymodule/v2
```

> [!note]
> 模組名稱雖然可以不包含 URL，但是帶有 URL 的模組名稱有助於找到並下載該模組，而且可以確保名稱唯一，避免跟其他模組名稱衝突或混淆。因此，建議的做法是以 URL 的寫法來指定模組名稱。

最後整理幾個重點：

- 每一個 module 都是以 module path 來作為唯一識別，這個模組路徑是宣告在一個 go.mod 檔案中。
- Modules 可以直接從版本控制儲存庫下載，或者從 module proxy 伺服器下載。
- 使用 `import` 來引用模組中的套件時，只能引用該模組 export 的（公開的）套件。

建議閱讀以下官方文件以了解更多有關 modules 的細節：

- [Go Modules Reference](https://go.dev/ref/mod)
- [go.mod file reference](https://go.dev/doc/modules/gomod-ref)
- [Module release and versioning workflow](https://go.dev/doc/modules/release-workflow)
- [Module version numbering](https://go.dev/doc/modules/version-numbers)

### Download a module

如果你的專案要使用某個外部模組，就必須先使用 `go get` 命令來把它下載至本機電腦。

範例：

```text
go get github.com/huanlin/learning-go
```

上述命令會直接到遠端的 GitHub 主機下載我的 `learning-go` 專案。

請注意，只有當模組符合此條件時，`go get` 命令才能成功下載：<mark>模組有包含 `go.mod` 檔案，而且 `go.mod` 檔案中的模組路徑與它所在的遠端 Git repository 的路徑相同。</mark>

> 除了 GitHub 之外，只要是 Git-based repositories，也都可以利用此命令來下載位於遠端主機的模組。例如：BitBucket。

**See also:** [What is 'go get' command in Go](https://gosamples.dev/go-get/).

#### Troubleshooting

我曾經在使用 `go get` 命令時碰到以下錯誤訊息：

```text
go: github.com/huanlin/learning-go@upgrade (v0.0.0-20240904141749-ce362a80bcf3)
    requires github.com/huanlin/learning-go@v0.0.0-20240904141749-ce362a80bcf3:
    parsing go.mod:
        module declares its path as: learning-go
                but was required as: github.com/huanlin/learning-go
```

後來發現是因為這個專案的 `go.mod` 檔案中最初的模組路徑是寫成 `learning-go`，像這樣：

```text
module learning-go
```

將檔案推送至遠端 GitHub 主機之後，如果用 `go get github.com/huanlin/learning-go` 命令來下載該模組，雖然 GitHub 主機上的確有這個 repository，可是 `go get` 命令發現該模組的 `go.mod` 檔案裡面寫的模組路徑並不是 `github.com/huanlin/learning-go`，於是顯示錯誤訊息並拒絕下載。

此時若修改 `go.mod` 檔案的內容，把模組路徑改為相符的 `github.com/huanlin/learning-go`，再推送至 GitHub 主機，然後再嘗試執行一遍剛才的 `go get` 命令，結果還是會得到同樣的錯誤訊息。這是 Go 在本機電腦的模組快取機制所造成的現象。

解決方法是隨意修改專案中的某個檔案，然後推送變更至 GitHub 主機，以產生一個新的 commit。接著以 `go get` 命令下載模組，並且在命令結尾處附加 @*<commit-hash>* 即可。比如說，commit hash 是 aa1ff21，那麼下載模組的命令會像這樣：

```text
go get github.com/huanlin/learning-go@aa1ff21
```

或者，也可以在 GitHub 上面建立一個 Release Tag，例如：`v0.0.1-beta`，然後告訴 `go get` 命令要下載這個版本的模組：

```text
go get github.com/huanlin/learning-go@v0.0.1-beta
```

### `go mod tidy` 命令 {#go-mod-tidy}

如果在編譯 go 程式時發生模組相關的錯誤（例如找不到某個模組），此時可以試試 `go mod tody` 命令。此命令會找到專案的 `go.mod` 檔案，並針對模組依賴關係執行以下工作：

- 下載專案有用到的外部模組。
- 更新 `go.mod` 檔案：在 `require` 區塊中加入必要的依賴，並刪除沒用到的依賴。
- 更新 `go.sum` 檔案。

舉例來說，假設 `go.mod` 檔案的內容如下：

```text
module github.com/huanlin/learning-go

go 1.23.0

require github.com/shirou/gopsutil/v4 v4.24.8
```

執行 `go mod tidy` 命令時，它會自動下載需要的套件：

```console
go: downloading github.com/shirou/gopsutil/v4 v4.24.8
go: downloading github.com/shoenig/go-m1cpu v0.1.6
go: downloading github.com/tklauser/go-sysconf v0.3.12
go: downloading github.com/power-devops/perfstat v0.0.0-20210106213030-5aafc221ea8c

...(略)
```

並且更新 `go.mod` 檔案：

```text
module github.com/huanlin/learning-go

go 1.23.0

require github.com/shirou/gopsutil/v4 v4.24.8

require (
    github.com/go-ole/go-ole v1.2.6 // indirect
    github.com/lufia/plan9stats v0.0.0-20211012122336-39d0f177ccd0 // indirect
    github.com/power-devops/perfstat v0.0.0-20210106213030-5aafc221ea8c // indirect
    github.com/shoenig/go-m1cpu v0.1.6 // indirect
    ...(略)
)
```

可以看到 `go mod tidy` 命令自動替 `go.mod` 檔案加入了一堆間接依賴的模組。

最後是 `go.sum` 檔案：

```text
github.com/davecgh/go-spew v1.1.1 h1:vj9j/u1bqnvCEfJOwUhtlOARqs3+rkHYY13jYWTU97c=
github.com/davecgh/go-spew v1.1.1/go.mod h1:J7Y8YcW2NihsgmVo/mv3lAwl/skON4iLHjSsI+c5H38=
github.com/go-ole/go-ole v1.2.6 h1:/Fpf6oFPoeFik9ty7siob0G6Ke8QvQEuVcuChpwXzpY=
github.com/go-ole/go-ole v1.2.6/go.mod h1:pprOEPIfldk/42T2oK7lQ4v4JSDwmV0As9GaiUsvbm0=
github.com/google/go-cmp v0.5.6/go.mod h1:v8dTdLbMG2kIc/vJvl+f65V22dbkXbowE6jgT/gNBxE=
...(略)
```

延伸閱讀：

- 官方文件：[go mod tidy](https://go.dev/ref/mod#go-mod-tidy)
- Youtube 影片：[how to import Golang local package](https://youtu.be/Nv8J_Ruc280?si=g5-SBXY1VYh5q1ko) （這影片把模組路徑和 import 套件時的路徑寫法講解得很清楚）

## Packages

Go 應用程式是由多個 packages 組成，一個 package 在檔案系統中就是一個資料夾，該資料夾底下的 .go 程式檔案必然隸屬同一個 package（否則無法通過編譯）。

> [!quote]
> A **package** is a collection of source files in the same directory that are compiled together. Functions, types, variables, and constants defined in one source file are visible to all other source files within the same package.
>
> Go 官方部落格：[How to Write Go Code](https://go.dev/doc/code)

換言之，package 是一個邏輯組成單位，讓我們把相關功能的程式檔案放在一起，也方便讓其他應用程式匯入使用。

範例：

```text
.                           -> 專案的根目錄
├─ go.mod                   -> 定義專案的名稱和 dependencies
├─ main.go                  -> 主程式（程式的進入點）
└─ auth/                    -> auth 套件
    ├─ auth.go              -> auth 相關功能的實作
    ├─ auth_test            -> auth 相關功能的測試
    └─ token/               -> token 套件
        ├─ token.go         -> token 相關功能的實作
        └─ token_test.go    -> token 相關功能的測試
```

Package 的名稱通常會跟它所在的資料夾名稱相同，但也可以不同。例如檔案 `auth.go` 裡的第一行通常會宣告套件名稱為 `auth`，但也可以是別的名稱，例如 `authentication`：

```go
package authentication

....
```

那樣的話，`auth_test.go` 的套件名稱也必須是 `authentication`，因為同一個資料夾底下的 .go 檔案必須隸屬同一個套件（否則無法通過編譯）。

### 匯入套件 {#import-packages}

在一個 Go 程式檔案中欲使用其他套件的識別字時，包括變數、函式、型別等等，必須使用 `import` 陳述句。比如說，要使用剛才提到的 `authentication` 套件，會在程式中這樣寫：

```go
import "mycompany.com/myapp/auth"
```

匯入套件時，套件的完整路徑名稱包含兩個部分：

- **模組路徑：** 套件路徑的開頭部分是模組路徑（module path），通常是模組的根 URL 或根目錄。如剛才範例中的 `mycompany.com/myapp` 即是模組路徑。
- **相對路徑：** 模組路徑後面接的是套件在模組中的相對路徑（相對於模組路徑），例如範例中的 `/auth` 即是 `authentication` 套件所在的相對路徑名。

如果要匯入多個套件，可以使用多行 `import` 語法，例如：

```go
import (
    "fmt"
    "mycompany.com/myapp/auth"
    "mycompany.com/myapp/db"
)
```

匯入套件時還可以指定套件的別名，以避免名稱衝突或提高程式碼的可讀性。例如：

```go
import (
    auth "mycompany.com/myapp/auth"
    database "mycompany.com/myapp/db"
)
```

### Package 有兩種 {#two-package-types}

另外要知道的是，Go 的 package 有兩種：

- 可執行套件：套件名稱一定是 `main`，而且不能被其他套件引用。
- 函式庫套件：套件名稱不是 `main` 的都是函式庫套件，可供其他套件引用。

至於不同的 package 之間要如何開放或隱藏某些變數、函式、或型別、或或服務，稍後會再說明。

### Package 名稱 {#package-names}

套件的命名慣例：

- 應簡潔明白，通常是名詞。
- 全部使用小寫英文字母。
- 雖然可以使用底線（`_`），但應盡量避免。
- 不可使用減號（`-`）。

詳見〈[附錄一：Go 程式風格指南]({{< ref "../a1-style-guide/index.md#package-naming" >}})〉或官方部落格：[Package names](https://go.dev/blog/package-names)。

## 標準 Go 專案目錄結構 {#std-project-layout}

在安排 Go 專案的目錄結構時，有兩種常見做法：

- **扁平結構：** 目前看起來大多數的 Go 專案是採用扁平結構，即主要的 Go 程式檔案會直接放在專案的根目錄下。
- **套件組織：** 如果專案規模較大，通常會將程式碼組織成多個套件，每個套件放在專案根目錄下的單獨目錄中。

可以參考[標準 Go 專案目錄結構](https://github.com/golang-standards/project-layout/blob/master/README_zh-TW.md)。它不是 Go 開發團隊制定的官方標準，而是根據常見作法所整理出來的通用結構。故參考這個標準目錄結構時，應該以專案實際的規模和需要來決定要有哪些資料夾，而不是照單全收，也不是資料夾拆分得越多就越好。

範例：

```text
myproject/
    ├── cmd/
    │   └── myapp/
    │       └── main.go
    ├── pkg/
    │   └── mypackage/
    │       └── mypackage.go
    ├── internal/
    │   └── myinternalpackage/
    │       └── myinternalpackage.go
    ├── api/
    │   └── v1/
    │       └── api.go
    ├── configs/
    │   └── config.yaml
    ├── scripts/
    │   └── build.sh
    ├── web/
    │   ├── static/
    │   └── templates/
    ├── go.mod
    └── go.sum
```

說明：

| 資料夾 | 用途說明 |
| ----- | --------|
| **cmd** | 每個子目錄代表一個應用程式的入口點（例如 myapp），包含 main.go 文件。 |
| **pkg** | 對外公開的套件，可讓其他專案使用。 |
| **internal** | 僅限專案內部使用的套件（Go 編譯器會確保這點）。 |
| **api** | API 程式碼。 |
| **configs** | 組態檔，例如 YAML、JSON。 |
| **scripts** | 腳本，用於建構、部署、生成等等。 |
| **web** | Web 相關資源，包括靜態文件和模板。 |

## 案例研究 {#case-study}

如果對 module 和 package 仍有不清楚的地方，不妨看一下別人的 Go 專案是如何組成的，包括 `go.mod` 檔案的內容、套件的階層結構、套件的命名等等。

這裡拿開源專案 gopsutil 為例，此專案的 GitHub 網址是：

<https://github.com/shirou/gopsutil>

首先，專案名稱 "gopsutil" 蠻符合 Go 對於模組和套件的命名建議：全部小寫的英文，而且盡量簡潔。此名稱是由 "go" 加上 "ps"（process）再加上 "util"（utility）所組成。

專案的根目錄底下有 `go.mod` 和 `go.sum` 檔案。`go.mod` 檔案描述模組的路徑以及用到哪些外部模組，`go.sum` 檔案則包含所有外部模組的 checksum，其內容由工具產生，不用自己編寫。

接著查看[此專案的 `go.mod` 檔案](https://github.com/shirou/gopsutil/blob/master/go.mod)的內容：

```text
module github.com/shirou/gopsutil/v4

go 1.18

require (
    github.com/ebitengine/purego v0.8.1
    github.com/google/go-cmp v0.6.0
    github.com/lufia/plan9stats v0.0.0-20211012122336-39d0f177ccd0
    github.com/power-devops/perfstat v0.0.0-20210106213030-5aafc221ea8c
    github.com/stretchr/testify v1.9.0
    github.com/tklauser/go-sysconf v0.3.12
    github.com/yusufpapurcu/wmi v1.2.4
    golang.org/x/sys v0.26.0
)

require (
    github.com/davecgh/go-spew v1.1.1 // indirect
    github.com/go-ole/go-ole v1.2.6 // indirect
    github.com/pmezard/go-difflib v1.0.0 // indirect
    github.com/tklauser/numcpus v0.6.1 // indirect
    gopkg.in/yaml.v3 v3.0.1 // indirect
)
```

從第一行可得知此專案的模組路徑是 `github.com/shirou/gopsutil/v4`。然後，在 GitHub 網站上查看[這個專案的 tags](https://github.com/shirou/gopsutil/tags) 可得知最新發布的版本，例如 `v4.24.8`。也就是說，如果應用程式需要用到 gopsutil 套件，我們的 Go 專案的 `go.mod` 檔案裡面要有底下的 `require` 宣告：

```text
require github.com/shirou/gopsutil/v4 v4.24.8
```

檔案 `go.mod` 裡面有兩個 `require` 區塊，其作用都是在描述此專案會用到外部模組。第一個 `require` 區塊是有直接引用的外部模組，第二個 `require` 區塊則是間接（indirect）用到的模組，即某些直接引用的模組內部有用到其他模組。在專案目錄下執行 `go mod tidy` 命令即可透過工具自動下載專案所依賴的外部模組，並自動更新這兩個 `require` 區塊的內容，包括移除沒用到的依賴，以及加入必要的依賴。

> 官方文件：[go mod tidy](https://go.dev/ref/mod#go-mod-tidy)

接著來看 gopsutil 專案的套件組成結構，如下圖：

![](images/gopsutil-folders.png#center)

其中的 `common`、`cpu`、`disk`、`internal/common` 等資料夾都是 packages。就如前面提過的，每一個 package 在磁碟檔案系統上面就是一個資料夾，而該資料夾底下的所有 Go 程式檔案都必須隸屬於同一個 package（不包含子目錄，因為子目錄也會成為一個 package）。

以 `cpu` 資料夾為例，它底下有許多 .go 檔案，例如 `cpu.go`、`cpu_aix.go`、`cpu_linux.go` 等等，這些 .go 檔案裡面都有底下這行 `package` 宣告，表示它們所屬的套件名稱是 `cpu`：

```go
package cpu
```

接在 `package cpu` 之後的是 `import` 宣告：

```go
import (
    "context"
    "errors"
    "fmt"
    "path/filepath"
    "strconv"
    "strings"

    "github.com/tklauser/go-sysconf"

    "github.com/shirou/gopsutil/v4/internal/common"
)
```

這裡有兩個地方值得注意。首先是<mark>套件的順序，建議寫法是先寫標準函式庫的套件，然後是其他套件。</mark>

其次，<mark>Go 程式在引用（import）第三方套件時不能寫相對路徑，而必須寫出套件的完整路徑名稱。</mark>套件的完整路徑名稱是由 go.mod 中宣告的模組路徑名稱再加上套件所在相對路徑名稱。

以 `"github.com/shirou/gopsutil/v4/internal/common"` 為例，它表示要使用的套件是位於 `github.com/shirou/gopsutil/v4` 模組（專案）底下的 `/internal/common` 套件。

## 尋找第三方套件

一個尋找 Go 第三方套件的好地方：<https://pkg.go.dev>。

在此網站的搜尋框輸入關鍵字，它會列出相關的套件。例如輸入 "util"，搜尋結果如下圖：

![](images/search-pkg.png#center)

搜尋結果的套件名稱旁邊如果有標示 "standard library"，即代表那是 Go 標準函式庫提供的套件；無此標示者則為第三方套件。

## Summary

- 一個 module 通常就是一個應用程式專案，而這個 module 裡面會有多個 packages。
- 一個 package 在檔案系統中就是一個資料夾，該資料夾底下的 .go 程式檔案必然隸屬同一個 package（否則無法通過編譯）。
- 模組路徑是模組的正式名稱（唯一識別名稱），宣告於模組的根目錄下的 `go.mod` 檔案；模組路徑要能表達該模組的用途，以及可以從何處找到它。
- Go 程式在 import 套件時，套件名稱的順序是先寫標準函式庫的套件，然後是其他套件。
- Go 程式在 import 第三方套件時不能寫相對路徑，而必須寫出套件的完整路徑名稱。套件的完整路徑名稱是由 go.mod 中宣告的模組路徑名稱再加上套件所在相對路徑名稱。
- 隸屬同一個 package 的程式碼可以互相存取宣告於 package 層級的名稱，如變數、函式、型別等等。
- 不同 package 的程式碼只能使用對方 export 出來的東西。
  - Go 語言沒有 `public`、`private` 或 `protected` 等識別字，而是根據變數名稱的第一個字母大小寫來判斷能否被外部引用。
  - 大寫字母開頭的名稱會被 export，即可供外界使用。
  - 小寫字母開頭的名稱無法被外界存取。

## References

- [How to Write Go Code](https://go.dev/doc/code)
- [Go Modules Reference](https://go.dev/ref/mod)
- [go.mod file reference](https://go.dev/doc/modules/gomod-ref)
- [Module release and versioning workflow](https://go.dev/doc/modules/release-workflow)
- [Module version numbering](https://go.dev/doc/modules/version-numbers)
- [標準 Go 專案目錄結構 (Standard Go Project Layout)](https://github.com/golang-standards/project-layout/blob/master/README_zh-TW.md)
- <https://pkg.go.dev/>

[100-mistakes]: https://www.manning.com/books/100-go-mistakes-and-how-to-avoid-them