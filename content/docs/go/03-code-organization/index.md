---
title: 03 Code organization
tags: [Go]
---

## Packages

一個 package 是一個或多個 .go 程式檔案所組成；這些程式檔案會放在同一個資料夾底下，而這個資料夾的名稱通常會跟 package 名稱一樣。

換言之，package 一個邏輯切割單位，讓不同用途的程式碼之間得以適度隔離。

另外要知道的是，Go 的 package 有兩種：

- 可執行套件：套件名稱一定是 `main`，而且不能被其他套件引用。
- 函式庫套件：套件名稱不是 `main` 的都是函式庫套件，可供其他套件引用。

至於不同的 package 之間要如何開放或隱藏某些資源或服務，請看下一節的說明。

## Scope

程式裡面有許多變數、函式、型別等識別字，依照它們宣告時的所在位置和寫法，分為三種可見範圍：

| 範圍 | 說明 |
|-----|------|
| block | 宣告在 `{...}` 區塊裡面的變數只有該區塊的程式碼可存取。 |
| package | 同一個 package 內的 .go 程式檔案可互相存取彼此的任何東西，包括變數、函式、型別等等，無論它們是否為 exported（公開成員）。 |
| global | 只要是 exported 變數（名稱以大寫英文開頭來命名的都是），就能夠被任何程式碼存取。 |

如前面提過的，package 是 Go 程式的邏輯切割單位，如果要讓 package 當中的某個東西公開讓其他套件也能使用，就必須以大寫英文開頭來命名。參考以下範例：

```go
package config

var ConfigFileName string = "d:/work/config.yaml" // 任何套件皆可存取。
var encoding string = "UTF-8" // 僅相同 package 的程式碼可以存取。

func createConfig() { // 僅相同 package 的程式碼可以存取。
    // ...
}
```

那麼，如果有兩個 .go 程式檔案放在同一個 package 裡面，有辦法讓其中一個 .go 程式檔案中的全域變數隱藏起來，不讓另一個 .go 程式檔案存取嗎？

答案是：不行。（這裡的全域變數指的是沒有包在任何 `{...}` 區塊中的變數）

這是因為，同一套件裡面的任何東西都是彼此共享的。這是 Go 語言賦予 package 的特性和規則。

> 也許有人會覺得這是 Go 語言的一個限制或缺點，但從另一個角度來看，寫程式的時候不用老是費心去考慮某些變數或函式到底要隱藏到什麼程度，也能讓事情變得簡單一點。

## Modules

每一個 project 都應該建立一個 `go.mod` 檔案來設定專案的基本資訊（名稱、版本）以及管理它所依賴的外部模組。

也就是說，一個 module 即代表一個應用程式專案。每個 module 是由一個或多個 packages 所組成，而且 module 名稱通常是以該專案的 Git repository 名稱來命名。

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

如果這個專案的原始碼是放在 GitHub 平台上的一個名為 "learning-go" 的 repository，那麼剛才的 `go mod init` 指令會這麼寫：

```text
mkdir todoapp
cd todoapp
go mod init github.com/todoapp
```

上述命令同樣只是在當前目錄建立一個 `go.mod` 檔案，內容會變成：

```text
module github.com/todoapp

go 1.23.0
```

{{% admonition type=note title="Note" %}}
模組名稱雖然可以不包含 URL，但是帶有 URL 的模組名稱有助於找到並下載該模組，而且可以確保名稱唯一，避免跟其他模組名稱衝突或混淆。因此，建議的做法是以 URL 的寫法來指定模組名稱。
{{% /admonition %}}

## More on variables

既然前面提到變數的可見範圍，這裡再介紹與變數有關的的兩個議題：variable shadowing 和 blank identifier。

### Variable shadowing

以下範例程式可以編譯和執行，但寫法容易令人 confuse：

```go
var case1 bool = true
var sum int = 100

func main() {
    if case1 {
        sum := add(5, 5) // 區域變數
        fmt.Println(sum)
    } else {
        m := add(10, 10) // 區域變數
        fmt.Println(sum)
    }

    fmt.Println(sum) // 使用全域變數
}

func add(x, y int) int {
    return x + y
}
```

程式中有幾處 `sum` 變數，有的是全域變數，有的是區域變數。雖然能通過編譯，但人眼容易誤讀，因為 `:=` 運算子可以同時宣告變數且賦值，使其左側的變數成為區域變數。如果使用 `=` 運算子，則會使用先前宣告過的變數，在此範例便是全域的 `sum`。

參見：[100 Go Mistakes and How to Avoid Them][100-mistakes] 的第 1 條：Unintended variable shadowing。

### Blank identifier

呼叫函式時，如果某個回傳值無需處理，可以用一個 blank identifier 字元（ `_` ）來承接該回傳值。

範例：

```go
-, err = ReadFile("no/file)
if (err != nil) {
    fmt.Println("Error: err)
}
```

此範例所要表達的是：我不在乎 `ReadFile()` 執行成功時回傳的結果，而只看它是否返回錯誤。


[100-mistakes]: https://www.manning.com/books/100-go-mistakes-and-how-to-avoid-them