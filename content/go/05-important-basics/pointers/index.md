---
title: 5.3 指標
---

Go 具備類似 C/C++ 的指標語法，但是更安全。Go 不允許指標運算，而且它有資源回收器在背後監視著每一個指標；當某一塊記憶體沒有任何指標指向它，Go 才會將那塊記憶體釋放。

> Go 不允許指標運算，除非透過 `unsafe` 套件。參閱官方文件：<https://pkg.go.dev/unsafe>。

宣告一個指標變數的語法是在型別前面加上星號 `*`：

```go
var p *int
```

這裡 `p` 是一個指向 `int` 的指標。由於沒有給初始值，故 `p` 的內容會是指標的預設值：`nil`。

## 位址運算子 {#address-operator}

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

## 傳值還是傳址？ {#pass-by-value}

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

## 從函式回傳指標 {#return-pointer}

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

