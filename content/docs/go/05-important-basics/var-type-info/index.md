---
title: 5.1 變數與型別資訊
---

## 變數 {#variables}

宣告變數時，可使用關鍵字 `var`，並使用 `=` 運算子來賦值。

範例：

```go
var x int
var y int = 100
```

沒有設定初始值的變數，都會有一個預設值。對 `int` 型別而言，這個預設值是 0，故此範例的 x 初始值為 0。

一次宣告多個變數：

```go
var x, y, z int
```

### Short declaration syntax

另一種更簡潔的語法是用 `:=` 運算子來一次完成兩件事：宣告變數且賦值，而且不用寫 `var`。此寫法稱為 short declaration syntax。

範例：

```go
sum := 100       // sum 是一個整數。
str := "hello"   // str 是一個字串。
x, y, z := true, sum, str // 一次宣告且設定多個變數。
```

使用 short declaration syntax 時，`:=` 的左側如果有多個變數，只要其中一個變數是新的（未曾宣告過的）即合法，否則編譯器會報錯。

```go
x := 10
y := 20
x, y := 1, 2       // 編譯錯誤! 因為 x 和 y 都已經宣告過。必須改用 `=` 才合法。
x, y, z := 1, 2, 3 // OK! 因為 z 是新宣告的變數。
```

## Scope

程式裡面的識別字（identifiers），像是變數、函式、型別等等，依照它們宣告時的所在位置和寫法，分為三種可見範圍：

| 範圍 | 說明 |
|-----|------|
| block | 宣告在 `{...}` 區塊裡面的識別字只有該區塊的程式碼可存取。 |
| package | 同一個 package 內的 .go 程式檔案皆可存取。 |
| global | 任何程式碼皆可存取。 |

有兩種情況會是 global 範圍：

- Go 的內建函式，例如 `panic()`。
- 在 package 層級宣告的識別字，名稱以英文大寫字母開頭來命名，就會被編譯器視為 **exported**，亦即公開的。

參考以下範例：

```go
package config

var ConfigFileName string = "d:/work/config.yaml" // 任何套件皆可存取。
var encoding string = "UTF-8" // 僅相同 package 的程式碼可以存取。

func createConfig() { // 僅相同 package 的程式碼可以存取。
    num := 100 // 只在此函式內可見。
}
```

那麼，如果有兩個 .go 程式檔案放在同一個 package 裡面，有辦法讓其中一個 .go 程式檔案中的全域變數隱藏起來，不讓另一個 .go 程式檔案存取嗎？

答案是不行。（這裡的全域變數指的是沒有包在任何 `{...}` 區塊中的變數）

這是因為，同一套件裡面的任何東西只要沒有被 `{...}` 包起來，在同一個套件範圍內都是共享的。這是 Go 語言賦予 package 的特性和規則。

> [!note]
> 也許有人會覺得這是 Go 語言的一個限制或缺點，但從另一個角度來看，寫程式的時候不用老是費心去考慮某些變數或函式到底要隱藏到什麼程度，也能讓事情變得簡單一點。

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

## 型別轉換 {#type-casting}

Go 是靜態型別語言，編譯器會自動推測型別，也會判斷型別是否相容。指派變數值的時候，若來源型別和目的型別不相容，便需要手動轉型，否則編譯器會報錯。

範例：

```go
var num int = 100

num = int64(50)   // 編譯錯誤。
num = 3.1416      // 編譯錯誤。
num = int(3.1416) // OK! num 的數值為 3。
```

更多語法和範例，建議參考官方文件：[Conversions](https://go.dev/ref/spec#Conversions)。

## 取得型別資訊 {#get-type}

這裡示範三種方法來取得變數的型別資訊：

- 使用 fmt.Printf 的 %T 旗號。
- 使用 reflect 套件。
- 使用 type assertion。

### 使用 fmt.Printf 的 %T 旗號 {#printf-t-flag}

```go
var count int = 42
fmt.Printf("variable count=%v is of type %T \n", count, count)
```

### 使用 reflect 套件 {#reflect-package}

使用 `reflect.TypeOf()` 方法：

```go
fmt.Printf("%v", reflect.TypeOf(10))   // int
fmt.Printf("%v", reflect.TypeOf("Go")) // string
```

### 使用 type assertion {#type-assertion}

```go
var x interface{} = 7

switch x.(type) {
case int:
    fmt.Println("int")
}
```

參閱 A Tour of Go: [Type assertions](https://go.dev/tour/methods/15)
