---
title: 05 雜七雜八但是重要
tags: [Go]
---

整理一些比較重要或 tricky 的 Go 語法或標準函式庫用法。

> Tip: [到 Go Playground](https://go.dev/play/) 寫點程式來測試和驗證自己的理解。

## 變數 {#variables}

宣告變數時，可使用關鍵字 `var`，並使用 `=` 運算子來賦值。

範例：

```go
var x int
var y int = 100
```

沒有設定初始值的變數，都會有一個預設值。對 `int` 型別而言，這個預設值是 0，故此範例的 x 初始值為 0。

另一種更簡潔的語法是用 `:=` 運算子來一次完成兩件事：宣告變數且賦值，而且不用寫 `var`。此寫法稱為 short declaration syntax。

範例：

```go
sum := 100       // sum 是一個整數。
str := "hello"   // str 是一個字串。
```

### 型別轉換 {#type-casting}

Go 是靜態型別語言，編譯器會自動推測型別，也會判斷型別是否相容。指派變數值的時候，若來源型別和目的型別不相容，便需要手動轉型，否則編譯器會報錯。

範例：

```go
var num int = 100

num = int64(50)   // 編譯錯誤。
num = 3.1416      // 編譯錯誤。
num = int(3.1416) // OK! num 的數值為 3。
```

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

## 指標 {#pointers}

Go 具備類似 C/C++ 的指標語法，但是更安全。Go 不允許指標運算，而且它有資源回收器在背後監視著每一個指標；當某一塊記憶體沒有任何指標指向它，Go 才會將那塊記憶體釋放。

> Go 不允許指標運算，除非透過 `unsafe` 套件。參閱官方文件：<https://pkg.go.dev/unsafe>。

宣告一個指標變數的語法是在型別前面加上星號 `*`：

```go
var p *int
```

這裡 `p` 是一個指向 `int` 的指標。由於沒有給初始值，故 `p` 的內容會是指標的預設值：`nil`。

### 位址運算子 {#address-operator}

宣告為指標的變數，其內容就是一個記憶體位址，該位址所在的地方才是變數值所在的記憶體區塊。在操作指標時，除了 `*`，還會使用 `&` 符號：

- 在變數名稱前面加上 `&` 符號會取得該變數所在的記憶體位址（address）。
- 在指標變數名稱前面加上 `*` 符號則代表該指標所指向之變數的內容（value）。

範例：

```go
num := 100   // 編譯器決定 num 是個 int。
ptr := &num  // 編譯器決定 ptr 是個指向變數 num 的指標。
*ptr = 200   // 把 ptr 指向的變數的內容改為 200。

fmt.Println(num)
fmt.Printf("Type of ptr: %T", ptr) // 印出 ptr 的型別名稱。
```

執行結果：

```text
200
Type of ptr: *int
```

### 傳值還是傳址？ {#pass-by-value}

<mark>Go 只有傳值（pass by value）。</mark>也就是說，當我們把一個變數傳入某函式的參數時，該參數會是傳入之變數的新副本；在函式中修改那個參數值並不會改動先前的變數。

如果要讓函式可以修改傳入參數的變數內容，就要使用指標：

```go
func main() {
    num := 100 // 編譯器決定 num 是個 int。

    increase(&num)
    fmt.Println(num) // 印出 101
}

func increase(n *int) {
    *n++
}
```

### 從函式回傳指標 {#return-pointer}

Go 函式可以回傳一個指向函式區域變數的指標：

```go
func newInt() *int {
    num := 42
    return &num
}

func main() {
    c := newInt()
    fmt.Println(*c)      // 印出 42
    fmt.Printf("%T", c)  // 印出 *int
}
```

`newInt()` 裡面的區域變數 `num` 所佔據的記憶體不會在函式返回之後立即消失，因為呼叫端 `main()` 函式有一個指標 `c` 仍指向 `num` 所在的記憶區塊。等到 `num` 所在的記憶區塊完全沒有人參考時，Go 的資源回收器便會將它所佔據的記憶體回收。

## String

留意 Unicode 的字串長度：

```go
unicodeCharStr := "地鼠"
fmt.Println(len(unicodeCharStr)) // output: 6
```

程式印出的結果是 6 而不是 2。這是因為 <mark>Go 的字串內部不是字元陣列，而是代表每個 UTF-8 字元的 byte 陣列。</mark>

因此，如果取出字串中的某個字元，不能以陣列索引的語法，否則結果不會是我們想要的：

```go
unicodeCharStr := "地鼠"
for i := 0; i < len(unicodeCharStr); i++ {
    fmt.Print(string(unicodeCharStr[i]) + " ")
}
fmt.Println() // 輸出:  å  ° é ¼
```

要取出字串中的字元，可以用 `range`：

```go
unicodeCharStr := "地鼠"
for i, r := range unicodeCharStr {
    fmt.Printf("%d:%s ", i, string(r))
}
fmt.Println() // 輸出: 0:地 3:鼠
```

## Blank identifier

呼叫函式時，如果某個回傳值無需處理，可以用一個 blank identifier 字元，也就是底線（ `_` ）來承接該回傳值。

範例：

```go
-, err = ReadFile("no/file)
if (err != nil) {
    fmt.Println("Error: err)
}
```

此範例所要表達的是：我不在乎 `ReadFile()` 執行成功時回傳的結果，而只看它是否返回錯誤。

## `init` 函式 {#init-func}

Go 有一個特殊用途的函式，名稱固定叫做 `init()`。此函式會在一個 package 載入時自動執行，故通常會把一些初始化的操作寫在此函式中。

特性與用法：

- `init` 函式不能有參數和回傳值。
- 每個 package 可以有多個 `init` 函式，而各 packages 的 `init` 函式的執行順序便是按照 packages 之間的依賴關係來決定。換言之，先載入哪個 package，就會先執行那個 package 的 `init` 函式。
- 每個 .go 檔案也可以有多個 `init` 函式，這些函式會按照它們在檔案中出現的順序執行。但一般來說，通常不會在一個 .go 檔案裡面寫多個 `init` 函式，以免讓程式碼更難理解和維護。
- 應避免在 `init` 函式中執行耗時工作。

範例：

```go
var greeting string

func init() {
    fmt.Println("Hello")
}

func main() {
    fmt.Println("world")
}
```

執行結果：

```text
Hello
world
```

### Using `init()` as side effects {#init-as-side-effects}

有時候，我們的程式不會呼叫某個套件裡面的任何函式或變數，但是卻需要應用程式載入時執行那個套件的 `init()` 函式，以便完成某些預設的配置或初始化操作。碰到這種情形，就必須在 import 該套件時使用 blank identifier (`_`)，像這樣：

```go
import (
    _ "github.com/lib/pq"
    _ "image/png"
    ...
)
```

這種在引用套件的名稱前面加一個底線字元的寫法稱為 importing side effects。換言之，那些套件的 `init()` 函式便是我們的應用程式所需要的「副作用」。

> [!note]
> 如果應用程式沒有直接使用某個套件裡面的任何東西，該套件名稱就不能出現在 `import` 區塊裡，除非前面加上 `_` 字元來告訴 Go 編譯器：我需要這個套件載入時自動執行的 `init()`。

## 基本流程控制

### `for` loop {#for-loop}

底下是幾種常見的寫法：

```go
i := 1
for i <= 3 {
    fmt.Println(i)
}

for j := 0; j < 3; j++ {
    fmt.Println(j)
}

for {  // 無限迴圈
    fmt.Println("loop")
}
```

迴圈裡面可以用 `continue` 來進行下一圈，以及用 `break` 來跳離迴圈。

#### For-each range loop

使用 `range` 關鍵字來指定索引值的範圍：

```go
for i := range 3 {  // i = 0, 1, 2
    fmt.Println("range", i)
}
```

常用來處理 arrays、slices、maps、channels 等結構：

```go
strings := []string{"hello", "world"}
for i, s := range strings {
    fmt.Println(i, s)
}
```

執行結果：

```text
0 hello
1 world
```

上例中，若不在乎陣列的索引值，可使用 blank identifier `_` 取代 `i`：

```go
strings := []string{"hello", "world"}
for _, s := range strings {
    fmt.Println(s)
}
```

執行結果：

```text
hello
world
```

### `if` statement

Go 的 `if` 陳述式不需要使用小括弧 `()`，但必須使用大括號 `{}`。雖然某些簡單的表達式可以使用小括弧，例如 `if (j > 10) {}` 這樣的寫法可以通過編譯，但與常見的寫法不一致，而且多此一舉。

> [!note]
> 事實上，在 VS Code 中編寫 Go 程式的時候，即使在 `if` 陳述式中使用了小括弧，那些小括弧會在存檔時被 [gofmt](https://pkg.go.dev/cmd/gofmt) 工具自動消除。這是因為 Go 工具鍊在 VS Code 中的預設配置為存檔時自動格式化。

`else` 敘述一定要跟在 `if` 區塊結尾的大括號後面，不能寫成單獨一行，像底下這樣的寫法將無法通過編譯：

```go
if j > 10 {
}
else { // 編譯錯誤! else 必須寫在 if 區塊結尾的同一行。
}
```

#### If with a short statement {#if-with-a-statement}

類似 `for` 迴圈，`if` 敘述也可以先有一個短敘述（short statement），然後才跟著判斷式。

範例：

```go
func pow(x, n, lim float64) float64 {
    if v := math.Pow(x, n); v < lim {
        return v
    }
    return lim
}
```

第 2 行的意思是先把 `math.Pow()` 的結果指派給變數 `v`，然後判斷 `v` 是否小於 `lim`。

注意：由 `if` 的短敘述所宣告的變數只活在那個 `if` 區塊內。

## Defer 陳述句 {#defer}

Go 的 `defer` 關鍵字可用來將一個函式呼叫的執行時機延後至包覆函式（surrounding function）結束之前才執行，常用於清理資源（例如確保關閉資料庫連線）。

範例：

```go
func main() {
    defer fmt.Println("World") // 離開 main 函式之前才執行此敘述。
    fmt.Println("Hello")
}
```

輸出結果：

```text
Hello
World
```

### 清理資源 {#defer-cleanup}

範例：

```go
func doSomething() error {
  f, err := os.Open("test.txt")
  if err != nil {
    return err
  }
  defer f.Close()

  // 繼續處理檔案內容
}
```

注意：一旦檔案開啟成功，**接著立刻加上** `defer f.Close()`，然後才處理後續的檔案操作，如此便可確保此函式離開之前會關閉檔案。

### 後進先出 {#defer-lifo}

如果在一個函式中使用了多次 `defer`，那些被延遲的函式呼叫將會以後進先出的順序執行。

範例：

```go
func main() {
    defer fmt.Println(1)
    defer fmt.Println(2)
    defer fmt.Println(3)
}
```

輸出結果：

```text
3
2
1
```

另外要注意的是，延後執行的時機除了函式正常返回，還有一種情況：goroutine 發生了執行時期的 panics。相關細節與注意事項可參閱官方文件：[Defer statements](https://go.dev/ref/spec#Defer_statements)。


[100-mistakes]: https://www.manning.com/books/100-go-mistakes-and-how-to-avoid-them
[go-in-action]: https://www.manning.com/books/go-in-action-second-edition
[go-in-practice]: https://www.manning.com/books/go-in-practice-second-edition
[go-by-example]: https://www.manning.com/books/go-by-example