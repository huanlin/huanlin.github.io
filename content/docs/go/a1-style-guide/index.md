---
title: 附錄一：Go 程式風格指南
---

{{% admonition type=note title="Note" %}}
這篇筆記還沒整理完，先別看喔。
{{% /admonition %}}

## Package naming {#package-naming}

套件的名稱應簡潔明白，通常是名詞，而且按照慣例全都用小寫英文字母。雖然可以使用底線字元 `_`，但最好盡量避免。減號字元 `-` 則不能用於套件名稱。

範例：

- `list`
- `http`
- `strconv` （兩個單字的縮寫組合: string conversion）
- `syscall` （兩個單字的縮寫組合：system call）
- `fmt` （format 的縮寫）

應避免涵義廣泛的名稱，像是 `util`、`utility`、`helper`、`common` 等等。目的不夠明確的名稱不利於理解，而且更容易跟其他套件撞名。

不過，Go 標準函式庫裡面也是有用 `util` 來命名的套件和檔案，例如：[types/util.go](https://github.com/golang/go/blob/master/src/go/types/util.go)、[httputil.go](https://github.com/golang/go/tree/master/src/net/http/httputil)、[ioutil.go](https://github.com/golang/go/blob/master/src/io/ioutil/ioutil.go) 等等。

> [!note] 關於減號字元
> 如果在某種情況下，套件的資料夾名稱必須用到減號 `-`，例如 `env-var` 可能比 `envvar` 更能清楚辨認為 environment variables 的縮寫，像這種情況，由於套件名稱不允許出現減號字元（編譯器會報錯），故可以考慮把套件名稱命名為 `env_var`，也就是用底線來取代原本的減號。這只是權宜作法，最好還是以全英文小寫來命名套件，且套件所在的資料夾名稱也和套件名稱一致。

**See also:** Go 官方部落格：[Package names](https://go.dev/blog/package-names)。

## File naming

檔案名稱應該以全部英文小寫搭配底線字元 (`_`) 來命名。一般而言，不同的單字會以底線字元隔開，但也有不少情況是兩個單字連在一起。

以下範例取自 Go 標準函式庫 [net/http 的原始碼](https://github.com/golang/go/blob/master/src/net/)：

- `responsecontroller.go`
- `roundtrip.go`
- `roundtrip_js.go`
- `routing_index.go`
- `routing_index_test.go`
- `transport.go`
- `transport_default_other.go`
- `transport_default_wasm.go`

注意事項：

- 檔案名稱如果是以 "." 或 "_" 開頭，Go tools 會忽略這些檔案。
- 測試程式的檔案必須以 `_test.go` 結尾，以便 Go 測試工具辨識。
- 檔案名稱如果以特定的作業系統或處理器架構的名稱結尾，將會影響實際的編譯結果。例如，`dirent_linux.go` 只會用於建置 Linux 環境的應用程式，`dir_windows` 只用於建置 Windows 環境的應用程式。

## Interface naming

### 使用 "er" 後綴 {#er-suffix}

根據[官方文件](https://go.dev/doc/effective_go#interface-names)，如果介面只包含一個方法，則該介面的慣例命名方式為"方法名稱+er"，例如： `Reader`、`Writer`、`Formatter` 等等。除此之外，便沒有進一步說明命名規則。也許我們可以延伸解讀為：只要能適當反映介面的行為，無論介面當中包含幾個方法，都可以用這種 "er" 後綴的命名方式。

### 使用 "I" 前綴 {#i-prefix}

許多物件導向語言建議使用大寫字母 `I` 前綴來命名介面（例如 `IWriter`），好處是能讓開發者更容易辨認哪些是介面，哪些是具象型別。Go 的官方風格則是強調根據介面的行為來命名，而不使用 `I` 前綴。不過，Go 官方文件並未禁止或反對使用 `I` 前綴來命名介面（至少我沒有看到），所以這種命名方式應該也可以納入選項。

> 還有一種比較累贅的命名方式：在名詞的後面加上 `Interface`，例如 `WriterInterface`。這種命名方式似乎比較少見。

## 程式自動排版 {#gofmt}

所有的 Go 程式都必須符合 `gofmt` 工具採用的排版格式。

## 變數 {#variables}

除非是 exported 變數，否則一律使用 camelCase 來命名變數。

```go
var myName = ""   // DO
var my_name = ""  // DON'T
```

偏好短一點的名稱。例如 `request` 可以用比較常見的縮寫 `req`。

在可見範圍比較小的區塊中的區域變數名稱可以用更精簡的名稱，例如索引值可以用一個字母 `i` 來命名。

```go
for i := 0; i < 100: i++ { }              // DO
for index := 0; index < 100: index++ { }  // DON'T
```

## 函式 {#functions}

不要寫 Getter 和 Setter 函式。

```go
func Person() person     // DO
func GetPerson() person  // DON't
```

不要在名稱裡重複套件名稱。

```go
encrypt.SHA()         // DO
encrypt.EncryptSHA()  // DON'T
```

## 推薦閱讀

- [Go Wiki: Go Code Review Comments](https://go.dev/wiki/CodeReviewComments)
- [Go Style Decisions](https://google.github.io/styleguide/go/decisions)
- [Go Style Best practices](https://google.github.io/styleguide/go/best-practices)

## 更多風格指南 {#more-guides}

- Uber Go Style guide
  - [英文版](https://github.com/uber-go/guide/blob/master/style.md)
  - [繁體中文版](https://github.com/ianchen0119/uber_go_guide_tw)
