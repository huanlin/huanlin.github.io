---
title: 06 結構
tags: [Go]
---

Go 的設計者對物件導向程式設計（OOP; object-oriented programming）的看法跟一般認知的 OOP 不大相同。

Go 沒有類別和繼承機制，但是有結構（struct），而且：

- 我們可以將任何函式附加（attach）至**同一個 package 中**的任何具象型別。換言之，如果函式和型別隸屬不同 package，那就不行。比如說，我們無法將自己寫的函式附加至 Go 標準函式庫的 `time.Duration`。
- 型別能夠隱含地實作介面（無需明白宣告欲實作哪個介面）。

## 結構入門 {#struct-basic}

以下範例先定義了一個 `person` 結構，然後在程式中使用它。

```go
// 定義一個結構類型
type Person struct {
    Name string
    Age  int
}

func main() {
    p1 := Person{}               // 每個欄位都初始化為 0 或空值。
    p2 := Person{"Michael", 25}  // 按欄位宣告順序設定初值。
    p3 := Person{                // 依欄位名稱設定初值。
        Name: "Michael",
        Age:  25,          // 最後一律要加逗號!
    }

    fmt.Println(p1)        // { 0}
    fmt.Println(p2)        // {Michael 25}
    fmt.Println(p2 == p3)  // true
}
```

**Try it:** <https://go.dev/play/p/x-2GAFCIcG8>

此範例展示了以 **struct literal** 語法來建立和初始化一個結構，共有三種寫法：

- `p1`：每個欄位都初始化為 0 或空值。
- `p2`：按欄位宣告的順序逐一設定初值。此寫法不用寫欄位名稱，但必須給所有的欄位都提供初始值。
- `p3`：依欄位名稱來設定初值，格式為 `name: value`，可以只設定部分欄位。由於這裡是每一個欄位寫成單獨一行，故每一行最後一律要加逗號，否則編譯器會視為語法錯誤（這是 Go 設計者貼心的地方）。如果全寫在同一行，則最後一個欄位的結尾處不用加逗號。

其他值得留意的地方：

- 這裡的結構型別 `Person` 以及結構成員（欄位）都是以英文大寫開頭，表示它們都會公開給其他套件使用。如果要限定同一套件才能使用，則名稱必須以英文小寫開頭來命名，例如 `person`。結構名稱與其成員欄位的名稱不見得採用一致的大小寫命名方式，例如結構型別可能以英文小寫開頭來命名（例如 `person`），亦即 unexported（不給其他套件使用），而結構成員以大寫英文開頭來命名（例如 `Age`）。
- 範例程式的最後一行用 `==` 來比較 `p2` 和 `p3`，用來展示如何比較兩個結構的內容是否相同。不過，這種方法只適用於簡單的場合，對於比較複雜的情況，則需要改用其他方法甚至第三方套件。詳見稍後的小節：[比較兩個結構是否相同](#cmp-structs)。

### 底層型別 {#underlying-type}

宣告一個具名型別時，必須指定另一個型別作為其「底層型別」（underlying type）。例如前面範例在宣告 `Person` 結構時，其底層型別是匿名的 `struct` 型別：

```go
type Person struct {
    Name string
    Age  int
}
```

這個底層型別也可以是具名型別，例如 Go 標準函式庫的 `Durarion` 型別是這麼宣告的：

```go
type Duration int64
```

這表示 `Duration` 的底層型別是 `int64`。（參閱官方文件：[type Duration](https://pkg.go.dev/time#Duration)）

### 使用結構指標

上一節的範例是使用所謂的 struct literal 語法來建立一個 `Person` 結構的實體。還有兩種建立結構的寫法：

- 使用地址運算子 `&`。
- 使用關鍵字 `new`。

範例：

```go
p1 := &Person{Name: "Michael", Age: 25}

p2 := new(Person)
p2.Name = "Michael"
p2.Age = 25
```

以上兩種寫法都會得到一個指標，指向新建立的結構。也就是說，`p1` 和 `p2` 這兩個變數的型別都是「指向 Person 結構的指標」。

以下示範更多種寫法，並藉由執行結果來觀察變數的型別：

```go
p1 := Person{"Michael", 25}

// p2, p3, p4, p5 全都是指向 Person 結構的指標，只是寫法不同。
p2 := &Person{"Michael", 25}
p3 := &Person{Name: "Michael", Age: 25}
p4 := &Person{
    Name: "Michael",
    Age:  25,
}
p5 := new(Person)
p5.Name = "Michael"
p5.Age = 25

fmt.Printf("p1: %v, type: %T\n", p1, p1)
fmt.Printf("p2: %v, type: %T\n", p2, p2)
fmt.Printf("p3: %v, type: %T\n", p3, p3)
fmt.Printf("p4: %v, type: %T\n", p4, p4)
fmt.Println(p1 == *p2)   // true
fmt.Println(p2 == p5)    // false
```

執行結果：

```text
p1: {Michael 25}, type: main.Person
p2: &{Michael 25}, type: *main.Person
p3: &{Michael 25}, type: *main.Person
p4: &{Michael 25}, type: *main.Person
true
false
```

**Try it:** <https://go.dev/play/p/8srsvYtUflh>

說明：

- 範例中的 `p3` 和 `p4` 只是為了示範不同寫法，最後將它們的內容和型別印出來並沒有特別用意，單純是因為要通過編譯。（變數宣告了就必須使用，否則無法通過編譯。）
- 倒數第二行的 `p1 == *p2` 是比較兩個結構實體的內容。注意這裡的 `p2` 前面要加 `*`，因為它本身只是個指標，必須用 `*` 來表示它指向的實體。
- 最後一行程式碼輸出結果為 `false`，因為 `p2` 和 `p5` 都是指標，所以 `p2 == p5` 所比較的內容會是兩個變數的內容，亦即記憶體位址。`p2` 和 `p5` 分別指向不同的結構實體，故二者指向的記憶體位址當然不同，故二者不相等。

### 匿名結構 {#anonymous-struct}

匿名結構的使用時機：暫時性的場合，通常是在某個函式裡面只用到一次（故無需用費心替它命名）。

以下範例展示了如何使用匿名型別的結構，並且直接初始化。

```go
func main() {
    p1 := struct {
        name string
        age  int
    }{
        name: "Michael",
        age:  25,
    }
    fmt.Println(p1.name, p1.age)
}
```

## 結構的欄位也可以是函式 {#struct-func}

```go
func main() {
    animal := struct {
        name string
        speak func() string
    } {
        name: "cat",
        speak: func() string {
            return "meow"
        },
    }

    fmt.Println(fmt.Sprintf("動物名稱是 %s，牠說 %s", animal.name, animal.speak() ))
}
```

## 為結構附加方法 {#struct-method}

上一節的範例是把函式加入結構的成員，這裡要示範的寫法有點像是替既有結構額外附加（**擴充**）一個方法。

```go
type Animal struct {
    name string
}

func (a Animal) speak() string {
    switch a.name {
    case "cat":
        return "meow"
    case "dog":
        return "woof"
    default:
        return "nondescript animal noise?"
    }
}

func main() {
    a := Animal{
        name: "cat",
    }

    fmt.Println(a.speak())

    a.name = "dog"
    fmt.Println(a.speak())

    a.name = "llama"
    fmt.Println(a.speak())#
}
```

像 `func (a Animal) speak() string {` 這樣的寫法稱為 "**a method with a receiver**"。事實上，「**方法**」（method）這個名詞在 Go 語言中是有正式定義的：

**A method is a function with a receiver.**

> 參見 The Go Programming Language Specification: [Method declarations](https://go.dev/ref/spec#Method_declarations)。

剛才的範例中，每次呼叫 `a.speak()` 時傳入的參數 `a` 都是一個新副本。如果想要讓 `speak()` 方法中修改原始傳入的 `a` 結構的內容，就要宣告成指標，像這樣：

```go
func (a *Animal) speak() string {
    ...
}
```

這裡只需要修改一行程式碼而已，其他地方不變。

**重點整理：**

- 方法（methods）是帶有一個 *receiver* 的函式，而 receiver 是寫在函式名稱前面的一個特殊參數，該參數的型別則表明了這是哪個型別的方法。
- Receiver 有兩種：pointer receiver 和 value receiver。前者可以修改傳入物件的內容，後者不行。

> 熟悉物件導向程式語言的人可以把 receiver 參數理解為 `this` 或 `self`，即「當前的物件本身」。

## 結構成員可以匿名 {#nameless-field}

```go
type Animal struct {
    string
}
```

欲存取沒有名稱的欄位，必須使用欄位的型別：

```go
func main() {
    a := Animal{
        "cat",
    }

    func (a Animal) speak() {
        log.Println(a.string)
    }
    fmt.Println(a.speak())

    a.string = "dog"
    fmt.Println(a.speak())
}
```

由於匿名欄位只能以其型別來存取，故這種寫法有個限制：只能有一個匿名欄位。

> 如果有給欄位命名，那麼即使只有一個欄位，也必須以名稱來存取該欄位，而不能用型別。

## 結構中的 tags {#struct-tags}

結構的欄位可以附加額外的描述資訊（metadata），稱為「標籤」（tags）。

Tags 的寫法是用一對 backtick 字元 ( ` ) 包住一組或多組 key: "value" 字串。每一組 key-value pair 是以空白字元隔開。

範例：

```go
type Animal struct {
    name string `help: "動物的種類或名稱，只要是貓或狗就行。"`
}
```

這裡替 `name` 欄位加上了一個 tag。該 tag 的 key 是 `help`，而 value 是 `"動物的...."`。

以下示範如何讀取欄位的 tag 內容：

```go
func (a Animal) speak() string {
    switch a.name {
    case "cat":
        return "meow"
    case "dog":
        return "woof"
    default:
        if member, ok := reflect.TypeOf(a).FieldByName("name"); ok {
            return fmt.Sprintf("無效的動物名稱：%s", member.Tag.Get("help"))
        }
        return "nondescript animal noise?"
    }
}
```

這裡使用了 [Go 的 reflection 套件](https://pkg.go.dev/reflect)來取得結構的執行時期型別資訊，並以 `FieldByName` 來取得結構成員。取得結構成員之後，便可以透過它的 `Tag.Get("help")` 方法來取得 tag key 為 "help" 的內容。

### 範例：將 tags 用於 JSON 序列化 {#struct-tags-json}

```go
package main

import (
    "fmt"
    "encoding/json"
)

type Animal struct {
    Name string `json:"animal_name"`
    ScientificName string `json:"scientific_name"`
    Weight float32 `json:"animal_average_weight"`
}

func main() {
    a := Animal{
        Name: "cat",
        ScientificName: "Felis catus",
        Weight: 10.5,
    }

    output, err := json.Marshal(a)
    if err != nil {
        panic("couldn't encode json")
    }
    fmt.Println(string(output))
}
```

請注意這裡的 `Animal` 結構的所有欄位成員的名稱開頭第一個字元都是大寫英文字母，表示它們是公開給任何程式碼存取。如果欄位名稱以小寫英文字母開頭，將導致 `encoding/json` 套件的函式無法存取它們。

程式的執行結果如下：

```json
{"animal_name":"cat","scientific_name":"Felis catus","animal_average_weight":10.5}
```

## 比較兩個結構是否相同 {#cmp-structs}

欲比較兩個結構的內容（所有欄位）是否相等，Go 標準函式庫有提供 [reflect.DeepEqual() 函式](https://pkg.go.dev/reflect#DeepEqual)。不過，使用上可能不夠彈性，例如：

- 不允許浮點數的誤差。
- 結構中的未公開欄位（unexported fields）也會一併比較。

若碰到類似限制，可試試開源套件：[go-cmp](https://github.com/google/go-cmp)。

## References

- [Go 每日一庫之 go-cmp](https://darjun.github.io/2020/03/20/godailylib/go-cmp/)