---
title: 07 集合
---

介紹三種集合類型：array、slice、map。

Go 內建了三種集合類型：陣列、切片（slice）、map。

- 陣列是固定長度的集合，所有元素都是相同型別。
- 切片是動態長度的集合，可理解為動態陣列，長度可任意變大縮小。
- Maps 是 key-value pairs，可透過 keys 來存取對應的 values。

## 陣列 {#array}

陣列是固定長度，故宣告時必須提供長度，例如：

```go
rgb := [3]byte{41, 190, 176}
```

此陣列的型別是 `[3]byte`，即一個內含三個 bytes 的陣列。宣告一個陣列時，其元素個數和元素類型共同決定了該陣列的類型。

此外，`byte` 型別是 `uint8` 的別名（alias）。故如果以下列程式碼印出 `rgb` 的型別：

```go
fmt.Printf("%T", rgb)  // 輸出型別
```

結果會是 `[3]uint8`。

### 傳遞陣列至函式 {#passing-arrays-to-func}

假設有一個解析度為 8K 畫質的圖片要保存於一個陣列。8K 的解析度是 7680 x 4320 個像素（pixels），也就是總共有 33,177,600 個像素。每個像素都是以 RGB 三原色的數值來表示，故用來保存該圖片的陣列所需要的空間為 7680 * 4320 * 3，將近 100 MB。然後，我們要把這個陣列傳遞給一個函式，將每個像素的顏色反轉。

```go
const resolution8K = 7_680 * 4_320 * 3           

func main() {
    image := [resolution8K]byte{ /* 顏色資料（略） */ }   
    invertColors(image)
}

func invertColors(colors [resolution8K]byte) {       
    for i := range colors {               
        colors[i] = 255 - colors[i]           
    }
}
```

這種寫法會傳遞整個陣列到函式，亦即要複製出另一個將近 100 MB 的陣列，實在太沒效率了。這種情況可以改為傳遞陣列的指標。

### 傳遞陣列的指標給函式 {#passing-array-pointers}

把上一節的範例改成傳遞陣列的指標：

```go
const resolution8K = 7_680 * 4_320 * 3

func main() {
    image := [resolution8K]byte{ /* 顏色資料（略） */ }   
    invertColors(&image)                   
}

func invertColors(colors *[resolution8K]byte) {       
    for i := range colors {
        colors[i] = 255 - colors[i]           
    }
}
```

如此一來，呼叫 `invertColors()` 函式的時候就只需要複製 8 bytes（在 64 位元的機器上，指標都是占 8 bytes）。

## Slice

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

由於此時 s2 的長度變成 4，超過原本配置的容量（3），需要擴容，即重新配置新的記憶空間；按照 slice 函式內部的演算法，新配置的容量會是原有容量的 2 倍，也就是 6。用以下程式碼便可確認：

```go
    s2 = append(s2, 9)
    fmt.Printf("len(s2): %d , cap(s2): %d\n", len(s2), cap(s2))
```

執行結果：

```text
len(s2): 4 , cap(s2): 6
```

接著，把 s2 傳入 `test()` 函式時，該函式裡面只增加一個元素，使其長度變成 5，而 5 個元素並未超出當前 slice 的容量（6），故這裡不會建立新的 slice 複本。換言之，此時 `test()` 函式中的變數 `s` 內部的 array 所指向的陣列區塊其實就是 s2 內部的 array 所指向的同一個陣列區塊。也因為這個緣故，s2 的內容會被 `test()` 修改。

> 註：這裡可能有人會誤解為「傳參考」，但 Go 函式只有傳值，沒有傳參考。詳細原因跟 slice 的結構有關。之後會再補相關說明，亦可參考文後的參考資料。

至於 s1，由於在傳入 `test()` 之前的長度和容量都是 3，而 `test()` 函式加入一個元素時，便超出了既有容量，導致 slice 需要擴容。一旦 slice 發生擴容，便會配置新的記憶區塊來儲存內部的陣列元素，於是 `s` 的內部陣列便不再指向 `s1` 的內部陣列，而是一個新建立的陣列。既然是兩個不同的陣列，對 s 做任何的修改也就不會影響 s1 了。

值得一提的是，slice 擴容的演算法並非每次都擴充為原本容量的 2 倍，而是由一個平滑遞減的成長因素來決定。相關細節可參考這個 commit: [runtime: make slice growth formula a bit smoother](https://go.googlesource.com/go/+/2dda92ff6f9f07eeb110ecbf0fc2d7a0ddd27f9d)。底下僅摘錄一部分內容：

```text
starting cap    growth factor
256             2.0
512             1.63
1024            1.44
2048            1.35
4096            1.30
```

## Map

*(TODO)*

## 其他補充 {#misc}

### 令新手迷惑的寫法 {#confusing}

底下兩種寫法，差別只在有沒有寫 `...`：

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

### 第三方套件

Go 標準函式庫提供的容器類型有陣列、slice、map、channel、heap、list、ring 等等。如果需要處理其他類型的資料結構，例如樹狀結構，可以試試一個叫做
Go Data Structures (GoDS) 的開源專案，網址是：<https://github.com/emirpasic/gods>。

## References

- The Go Blog: [Go Slices: usage and internals](https://go.dev/blog/slices-intro)
- [Slices in Go: Grow Big or Go Home](https://victoriametrics.com/blog/go-slice/)
- commit: [runtime: make slice growth formula a bit smoother](https://go.googlesource.com/go/+/2dda92ff6f9f07eeb110ecbf0fc2d7a0ddd27f9d)
- "[Go by Example](https://www.manning.com/books/go-by-example)" by Inanc Gumus (Manning)