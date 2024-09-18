---
title: 07 陣列與 slice
---

## 陣列 {#array}

*(有空再寫)*

## slice

*(基礎語法和內部結構晚點寫，先整理一個隱藏陷阱)*

### 隱藏的陷阱 {#slice-gotchas}

範例：

```go
func main() {
    s1 := []int{1, 2, 3}
    s2 := s1
    s2 = append(s2, 9)
    fmt.Println(s1)    // [1 2 3]
    fmt.Println(s2)    // [1 2 3 9]

    test(s1)
    fmt.Println(s1)    // [1 2 3]

    test(s2)
    fmt.Println(s2)    // [2 3 4 10] (為什麼會這樣?!)
}

func test(s []int) {
    s = append(s, 0)
    for i := range s {
        s[i]++
    }
}
```

執行結果：

```text
[1 2 3]
[1 2 3 9]
[1 2 3]
[2 3 4 10]
```

即使不清楚 slice 的運作機制，前兩次輸出的結果應該不難自行推導和理解：當時的 s1 和 s2 是兩個不同的 instances，故修改了 s2 並不會動到 s1。

然而，分別把 s1 和 s2 傳入 `test()` 函式之後，接著輸出的結果卻令人意外。這一次，s1 的內容依然不變，可是 s2 的內容卻被 `test()` 函式改變了。這是怎麼回事？

關鍵在於 slice 物件有沒有發生「容量擴充」的情形（以下簡稱「擴容」）：

**當 slice 需要擴容時，便會建立一個新的 slice 複本，並依特定演算法來替新複本配置新的容量。若無需擴容，則會使用同一塊陣列空間。**

由於 slice 的大小可以隨時變動，Go 使用了預先配置容量的方式來降低重新配置記憶區塊的頻率。比如說，原本的 slice 內容的長度（元素總數）為 3，接著又增加了一個元素，使其長度變成 4，而為了容納這個新進的元素，slice 的容量會變成原來的 2 倍，也就是 6，而不是 4。

以剛才的範例來說，可以先用以下程式碼來確認 s1 的長度和容量始終是 3。

```go
    fmt.Printf("len(s1): %d , cap(s1): %d\n", len(s1), cap(s1))
```

執行結果：

```text
len(s1): 3 , cap(s1): 3
```

剛開始，s2 的長度和容量也是 3，但是當程式執行完下面這行敘述：

```go
    s2 = append(s2, 9)
```

由於此時 s2 的長度變成 4，超過原本配置的容量（3），需要擴容，即重新配置新的記憶空間；而根據 slice 函式內部的演算法，新配置的容量會是原有容量的 2 倍，也就是 6。用以下程式碼便可確認：

```go
    s2 = append(s2, 9)
    fmt.Printf("len(s2): %d , cap(s2): %d\n", len(s2), cap(s2))
```

執行結果：

```text
len(s2): 4 , cap(s2): 6
```

接著，把 s2 傳入 `test()` 函式時，該函式裡面只增加一個元素，使其長度變成 5，而 5 個元素並未超出當前 slice 的容量（6），故這裡不會建立新的 slice 複本。換言之，此時 `test()` 函式中的變數 `s` 內部的 array 所指向的陣列區塊其實就是 s2 內部的 array 所指向的同一個陣列區塊。也因為這個緣故，s2 的內容會被 `test()` 修改。

> 註：這裡可能會有人誤解為「傳參考」，但 Go 函式只有傳值，沒有傳參考。詳細原因跟 slice 的結構有關。之後會再補相關說明。

至於 s1，由於在傳入 `test()` 之前的長度和容量都是 3，而 `test()` 函式加入一個元素時，便超出了既有容量，導致 slice 需要擴容。一旦 slice 發生擴容，便會配置新的記憶區塊來儲存內部的陣列元素，於是 `s` 的內部陣列便不再指向 `s1` 的內部陣列，而是一個新建立的陣列。既然是兩個不同的陣列，對 s 做任何的修改也就不會影響 s1 了。

值得一提的是，slice 擴容的演算法並非每次都擴充為原本容量的 2 倍，而是有一個漸進的成長因素來決定。相關細節可參考這個 commit: [runtime: make slice growth formula a bit smoother](https://go.googlesource.com/go/+/2dda92ff6f9f07eeb110ecbf0fc2d7a0ddd27f9d)。底下僅摘錄一部分內容：

```text
starting cap    growth factor
256             2.0
512             1.63
1024            1.44
2048            1.35
4096            1.30
```

### 令人迷惑的寫法 {#confusing}

有 `...` 和沒有的差別：

```go
a := [...]int{ 1: 10, 2: 20 } // a 是一個陣列
b := []int{ 1: 10, 2: 20 }    // b 是一個 slice

fmt.Printf("Type of a: %T\n", a)
fmt.Printf("Type of b: %T\n", b)
```

執行結果：

```text
Type of a: [3]int
Type of b: []int
```

## References

- The Go Blog: [Go Slices: usage and internals](https://go.dev/blog/slices-intro)
- [Slices in Go: Grow Big or Go Home](https://victoriametrics.com/blog/go-slice/)
- commit: [runtime: make slice growth formula a bit smoother](https://go.googlesource.com/go/+/2dda92ff6f9f07eeb110ecbf0fc2d7a0ddd27f9d)