---
title: Interfaces
tags: [Go]
draft: true
---

## 隱含的介面實作 {#implicit-interfaces}

Go 的介面跟其他程式語言的介面都是用來定義共同行為，但是在 Go 語言中，介面的實作是隱含的、由編譯器自動識別的，亦即無須明白宣告某個型別實作了那些介面。也因為這個緣故，撰寫 Go 程式的時候大多是先寫具象（concrete）型別的程式碼，然後從幾個具有相似行為的具象類別當中找出介面。

> 其他物件導向語言則通常是先定義介面，然後才撰寫具象類別來實作該介面。

### 先實作具象型別 {#concrete-first}

以下範例摘自《Go by Example》第 2 章：

```go
type slackNotifier struct {           
    apiKey  string
    channel string
    // ..other fields
}

func (s *slackNotifier) notify(msg string) {   
    fmt.Println("slack:", msg)
}

func (s *slackNotifier) disconnect() {
    fmt.Println("slack: disconnecting")
}

type smsNotifier struct {
    gatewayIP string
    // ..other fields
}

func (s *smsNotifier) notify(msg string) {       
    fmt.Println("sms:", msg)
}
```

`slackNotifier` 和 `smsNotifier` 都是具象型別，分別代表兩種通知方式。

- `slackNotifier` 結構有兩個方法：`notify` 和 `disconnect`。
- `smsNotifier` 結構只有一個方法：`notify`。

於是，我們發現這兩個結構有一個共同點：它們的 `notify` 方法長得一模一樣。

### 發現介面 {#discover-interface}

發現具象型別的共同行為之後，便可以為這個行為定義一個介面。在前面的範例中，兩個結構的共同點是 `notify` 方法，故我們可以為它定義一個叫做 `notifier` 的介面：

```go
type notifier interface {
    notify(message string)   
}
```

然後實作此介面的方法：

```go
func notify(s *server, n notifier) {   
    if !s.slow() {
        return
    }
    msg := fmt.Sprintf(
        "%s server is slow: %s",       
        s.url, s.responseTime,
    )
    n.notify(msg)               
}
```

### 使用介面 {#wiring-up}

```go
func main() {
    authServer := &server {
        url: "auth",
        responseTime: time.Minute
    }
    slack := &slackNotifier { /* Slack specific configuration */ }
    sms   := &smsNotifier   { /* SMS specific configuration   */ }

    notify(authServer, slack)
    notify(authServer, sms)
}
```

程式執行結果：

```text
$ go run .
slack: auth server is slow: 1m0s
sms: auth server is slow: 1m0s
```

**延伸閱讀：** [Go Data Structures: Interfaces](https://research.swtch.com/interfaces)

## 小心誤用介面 {#interface-pollution}

Go 語言的介面在設計上跟許多物件導向程式語言的介面不大相同，可能是這個原因，使得介面經常被誤用或濫用。

Rob Pike（Go 語言的其中一名主要開發者）曾說過：

> Don't design with interfaces, discover them.

意思是，只有當我們發現這裡或那裡需要一個介面會比較好的時候，才去定義介面。這可以避免我們太早傷腦筋去設計「想像中的」介面，也能避免過度設計。

- 太早或者過度使用介面，很容易會加入一堆用處不大的抽象層，讓程式更複雜、更難維護。
- 如果沒有強烈的理由、或者不大確定增加一個介面能夠帶來明顯好處，就應該再三斟酌。
- 不要太擔心直接呼叫實作會造成什麼嚴重後果。與其用更多抽象層來應對未來可能的狀況，不如先解決眼下的需求。

## 介面該放在哪一邊？ {#where-should-interface-live}

關於介面要定義在哪裡，基本上有兩種做法：

- **服務端**
- **用戶端**

熟悉 C# 或 Java 的人通常會把介面定義在生產端（服務端），然後把介面提供給用戶端去按照介面來實作。然而，這跟 Go 的設計理念大相逕庭。

Go 的設計者認為不應由服務端來定義介面然後強迫所有的用戶端都必須按圖施工，而應該讓各個用戶端決定它是否需要某種形式的抽象，並選擇它認為最合適的抽象層。

> 參考自《100 Go Mistakes and How to Avoid Them》第 2 章。原文如下：
>
> ...it’s not up to the producer to force a given abstraction for all the clients. Instead, it’s up to the client to decide whether it needs some form of abstraction and then determine the best abstraction level for its needs.

因此，撰寫 Go 程式時，介面通常應該由用戶端來定義。不過，這是指大多數的情況，仍有少數特例是把介面寫在服務端（例如 Go 標準函式庫的 `encoding` 套件即定義了介面讓它的子套件提供實作）。

舉例：如果服務端預先定義了一個 `CustomerStorage` 介面讓所有用戶端照著這份規格來實作：

```go
package store

type CustomerStorage interface {
    StoreCustomer(customer Customer) error
    GetCustomer(id string) (Customer, error)
    UpdateCustomer(customer Customer) error
    GetAllCustomers() ([]Customer, error)
    GetCustomersWithoutContract() ([]Customer, error)
    GetCustomersWithNegativeBalance() ([]Customer, error)
}
```

萬一某個用戶端根本不需要這樣的寬鬆耦合呢？又或者某個用戶端只需要其中的 `GetAllCustomers` 方法呢？在服務端硬性規定了介面（規格），用戶端等於毫無選擇，只能照單全收。

若採用 Go 的建議作法，介面由用戶端自行決定，那麼用戶端就可以根據自己的需要來訂出最合用的介面。例如：

```go
package client

type customersGetter interface {
    GetAllCustomers() ([]store.Customer, error)
}
```

然而，把實作放在服務端，而介面定義在用戶端，難道不會發生循環參考的情形嗎？在 Go 是不會的，因為 Go 的介面是以隱含的方式對應至實作，而不需要明確宣告某個型別實作哪個介面。

> 事實上，Go 根本沒有 Java 的 `implements` 關鍵字，也沒有類似 C# 的明確實作介面的語法。

## 避免回傳介面 {#avoid-returning-interfaces}

returning an interface is, in many cases, considered a bad practice in Go.

in general, returning an interface restricts flexibility because we force all the clients to use one particular type of abstraction.

建議的做法是：

- 回傳結構（structs），而不要回傳介面。
- 如果可能的話，讓介面成為函式的傳入參數。

> 標準函式庫的 `io.Reader` 之所以定義了介面，是因為設計者事先已經確知（而非預測或想像）這樣的抽象層是有助於重複使用的。

總之，多數情況下，我們的函式不應該回傳介面，而應該回傳實作型別。否則，我們的設計會因為套件之間的相依關係而太過複雜，也會因為用戶端都必須倚賴一套完全相同的抽象規格而失去彈性。另一方面，如果我們確知（而非預料）某個抽象層對多數用戶端是有幫助的，那就可以考慮回傳介面。

再重複一次：我們不應強迫用戶端接受一份抽象規格；這份規格應該由用戶端自行發現和決定。如果某個用戶端需要從某個實作提取一份抽象規格，它也一樣能在用戶端完成此事。
