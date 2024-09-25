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

### Short declaration syntax

另一種更簡潔的語法是用 `:=` 運算子來一次完成兩件事：宣告變數且賦值，而且不用寫 `var`。此寫法稱為 short declaration syntax。

範例：

```go
sum := 100       // sum 是一個整數。
str := "hello"   // str 是一個字串。
v1, v2, v3 := 10, sum, str // 一次宣告且設定多個變數。
```

使用 short declaration syntax 時，`:=` 的左側如果有多個變數，只要其中一個變數是新的（未曾宣告過的）即合法，否則編譯器會報錯。

```go
x := 10
y := 20
x, y := 1, 2       // 編譯錯誤! 因為 x 和 y 都已經宣告過。必須改用 `=` 才合法。
x, y, z := 1, 2, 3 // OK! 因為 z 是新宣告的變數。
```

## 型別轉換 {#type-casting}

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
