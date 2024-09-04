---
title: 附錄一：Go 程式風格指南
---

{{% admonition type=note title="Note" %}}
這篇筆記還沒整理完，先別看喔。
{{% /admonition %}}

## Package name

好的套件名稱應簡潔明白，通常是單數名詞，而且全都是用英文小寫。注意不可以用底線（snake case）或大小寫混和（mixedCaps）。

範例：

- `list`
- `http`
- `strconv` （兩個單字的縮寫組合: string conversion）
- `syscall` （兩個單字的縮寫組合：system call）
- `fmt` （format 的縮寫）

應避免涵義廣泛的名稱，像是 `util`、`utility`、`helper`、`common` 等等。目的不夠明確的名稱不利於理解，而且更容易跟其他套件撞名。

> 參見 Go 官方部落格：[Package names](https://go.dev/blog/package-names)。

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

## 更多風格指南 {#more-guides}

- 官方文件：[Go Style Decisions](https://google.github.io/styleguide/go/decisions)
- 官方文件：[Go Style Best practices](https://google.github.io/styleguide/go/best-practices)
- Uber Go Style guide
  - [英文版](https://github.com/uber-go/guide/blob/master/style.md)
  - [繁體中文版](https://github.com/ianchen0119/uber_go_guide_tw)
