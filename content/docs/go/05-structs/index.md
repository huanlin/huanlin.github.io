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

