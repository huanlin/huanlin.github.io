---
title: .NET Async FAQ
linkTitle: FAQ
draft: true
weight: 90
---

## To async or not? {#to-async-or-not}

這裡要討論的議題是：函式該直接回傳 `Task` 物件就好，還是一律使用 `async/await`？底下分別示範兩種寫法。

寫法一：直接回傳 `Task` 物件。

```csharp
public Task<int> DoSomethingAsync()
{
    return CallDependencyAsync();
}
```

寫法二：即使是簡單的非同步呼叫也一律使用 `async/await`。

```csharp
public async Task<int> DoSomethingAsync()
{
    return await CallDependencyAsync();
}
```

以上範例取自 David Fowler 撰寫的 Async Guidance 文件的其中一節：[Prefer async/await over directly returning Task](https://github.com/davidfowl/AspNetCoreDiagnosticScenarios/blob/master/AsyncGuidance.md#prefer-asyncawait-over-directly-returning-task)。Fowler 的建議是採用寫法二，也就是盡量使用 `async/await`，而不要直接回傳 `Task` 物件。他也在文中提到，直接回傳 `Task` 雖然能獲得稍微快一點的執行速度（因為它不用處理 async 狀態機的相關工作），但也失去了 async 狀態機帶來的一些好處，而且可能導致函式行為的改變。

兩種寫法的效能差異其實不大，通常不會是效能瓶頸之所在。故這裡推薦讀者採用 Fowler 的建議作法，也就是優先選擇採用 `async/await` 寫法來回傳非同步呼叫的結果。

如果你好奇直接回傳 `Task` 是否可能導致什麼比較嚴重的後果，以下範例展示了其中一種可能的狀況。

```csharp
public Task<string> GetWebPageTask()
{
    using var httpClient = new HttpClient();
    return httpClient.GetStringAsync("https://www.microsoft.com");
}
```

上面的程式碼經過編譯之後會有一個 `try/finally` 區塊，像這樣：

```csharp
public Task<string> GetWebPageTask()
{
    HttpClient httpClient = new HttpClient();
    try
    {
        return httpClient.GetStringAsync("https://www.microsoft.com");
    }
    finally
    {
        if (httpClient != null)
        {
            ((IDisposable)httpClient).Dispose();
        }
    }
}
```

由於 `httpClient.GetStringAsync()` 呼叫很可能尚未執行完畢，程式流程就進入了 `finally` 區塊而將 `httpClient` 物件回收，這將導致程式執行時發生 `TaskCanceledException`。這或許是 David Fowler 在其文章裡面說這種寫法將造成程式的「行為改變」的原因之一。

> Try it: <https://dotnetfiddle.net/NRXmfr>

採用 `async/await` 不只可以避免上述陷阱，還有其他好處，在 [David Fowler 的原文](https://github.com/davidfowl/AspNetCoreDiagnosticScenarios/blob/master/AsyncGuidance.md#prefer-asyncawait-over-directly-returning-task)裡面都有提到。完整起見，這裡用截圖的方式標示出來：

![Prefer async/await](images/fowler-prefer-async.png#center)

---

### 結論 {#prefer-async-conclusion}

多數情況下，直接回傳 `Task` 的好處抵不過它帶來的問題，故建議在回傳非同步呼叫的結果時優先選擇 `async/await` 寫法。

> [!note]
> 有一種見解是，當函式呼叫層層套疊很多層的時候，便應該傾向直接回傳 `Task` 物件。但我想還是應該基於是否真的足以產生「有實質意義上的效能差異」來決定，而不是有比較快就好。而且，直接回傳 `Task` 物件還可能造成一些潛在問題和陷阱（如前面舉的例子），而增加了日後維護程式的麻煩。

### See also

- [Async Guidance](https://github.com/davidfowl/AspNetCoreDiagnosticScenarios/blob/master/AsyncGuidance.md) by David Fowler
- [Async/Await - Best Practices in Asynchronous Programming](https://learn.microsoft.com/en-us/archive/msdn-magazine/2013/march/async-await-best-practices-in-asynchronous-programming) by Stephen Cleary
