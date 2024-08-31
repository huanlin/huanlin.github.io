---
title: 結構
tags: [Go]
draft: true
---

## 範例一：宣告一個結構型別 {#ex1}

以下範例先定義了一個 `person` 結構，然後在程式中使用它。

```go
type Person struct {
    name string
    age  int
}

func main() {
    fmt.Println("Hello, World!")

    james := Person{
        name: "James",
        age:  25,
    }
    fmt.Println(james.name, james.age)
}
```

**注意：**

- 給 `age` 成員賦值的時候，最後的逗號不可省略，否則編譯器會視為語法錯誤。這是 Go 設計者貼心的地方。
- 這裡的結構型別 `Person` 是以英文大寫開頭，表示可以公開給其他套件使用。

如果使用 `new` 來建立結構，會得到一個指向結構的指標；而使用 `&` 運算子也同樣會得到指向結構的指標。參考下範例所示：

```go
var p1 *Person = new(Person)
p2 := new(Person)
p3 := &Person{}

fmt.Printf("%T\n", p1)  // 輸出 p1 的型別名稱
fmt.Printf("%T\n", p2)  // 輸出 p2 的型別名稱
fmt.Printf("%T\n", p3)  // 輸出 p3 的型別名稱
```

這裡的 `p1`、`p2` 和 `p3` 都是指向一個新建立的 `Person` 結構的指標，所以三次輸出的型別名稱都一樣是 `*main.Person`。

## 範例二：使用匿名型別的結構 {#ex2}

以下範例展示了如何使用匿名型別的結構，並且直接初始化。

```go
func main() {
    james := struct {
        name string
        age  int
    }{
        name: "James",
        age:  25,
    }
    fmt.Println(james.name, james.age)
}
```

顯然，如果同樣的結構要使用很多次，應該使用範例一的寫法，也就是預先定義結構型別。

## 範例三：結構的欄位也可以是函式 {#ex3}

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

## 範例四：為結構附加方法 {#ex4}

範例三的寫法是把函式加入結構的成員，這裡要示範的寫法有點像是替既有結構額外附加（**擴充**）一個方法。

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

像 `func (a Animal) speak() string {` 這樣的寫法稱為 "**a method with a receiver**"。說它是個「方法」，以便和結構中的「函式」區別。

值得一提的是，每次呼叫 `a.speak()` 時，傳入 `speak()` 函式的 `a` 參數都是另一個新副本。如果想要讓 `speak()` 函式中修改原始傳入的 `a` 結構的內容，就要宣告成指標，像這樣：

```go
func (a *Animal) speak() string {
    ...
}
```

這裡只需要修改一行程式碼而已，其他地方不變。

## 範例五：結構成員可以匿名 {#nameless-field}

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

## 範例六：結構中的 tags {#tags}

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

## 範例七：將 tags 用於 JSON 序列化

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

## References

- Go in Practice, Second Edition