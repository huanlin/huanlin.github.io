---
title: .NET Async FAQ
linkTitle: FAQ
draft: true
weight: 90
---

## 非同步呼叫即表示函式將執行於新的執行緒？ {#async-thread}

不一定。

在 async 方法中使用 `await` 關鍵字等待的非同步呼叫不一定會在另一條執行緒上執行；這取決於被等待的非同步操作的具體實現。請參考接下來的三個範例。

### 範例一：閒置等待 {#ex1-idle}

單純的閒置等待不會動用新的執行緒：

```csharp
// 不會使用新執行緒。
async Task DelayAsync()
{
    await Task.Delay(1000);
}
```

### 範例二：I/O 操作 {#ex2-io-bound}

非同步的 I/O 操作通常不涉及執行緒。例如檔案讀寫操作，底層作業系統是透過所謂的 I/O completion port 來達成，其中的操作並不需要動用新的執行緒。

```csharp
// 不會使用新執行緒，而是使用 I/O completion port。
async Task ReadFileAsync()
{
    var s = await File.ReadAllTextAsync("file.txt");
}
```

### 範例三：CPU 密集操作 {#ex3-cpu-bound}

如果是 CPU 密集操作，則通常會由另一條執行緒來執行非同步方法：

```csharp
// 會在另一條執行緒執行
async Task DoHeavyWorkAsync()
{
    await Task.Run(() => {
        // CPU 密集運算
    });
}
```

重點整理：

- I/O 相關的非同步操作通常不需要建立新的執行緒。
- CPU 密集的操作才需要由新執行緒來執行。
- `await` 的主要作用是非阻斷（non-blocking），而不是創建新執行緒。

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

採用 `async/await` 不只可以避免上述陷阱，還有其他好處。比如說，萬一非同步呼叫的過程發生錯誤，exception 物件的 stack trace 資訊會更完整詳細，能夠顯示真正發生錯誤的程式碼位置；相較之下，直接回傳 `Task` 的寫法，其 exception 的 stack trace 會不完整，可能不會提供正確的出錯位置。有關 `async/await` 寫法的優點，在 [David Fowler 的原文](https://github.com/davidfowl/AspNetCoreDiagnosticScenarios/blob/master/AsyncGuidance.md#prefer-asyncawait-over-directly-returning-task)裡面都有提到。完整起見，這裡用截圖的方式標示出來：

![Prefer async/await](images/fowler-prefer-async.png#center)

---

### 結論 {#prefer-async-conclusion}

多數情況下，直接回傳 `Task` 的好處抵不過它帶來的問題，故建議在回傳非同步呼叫的結果時優先選擇 `async/await` 寫法。

> [!note]
> 有一種見解是，當函式呼叫層層套疊很多層的時候，便應該傾向直接回傳 `Task` 物件。但我想還是應該基於是否真的足以產生「有實質意義上的效能差異」來決定，而不是有比較快就好。而且，直接回傳 `Task` 物件還有一些潛在問題和缺點（如前面提提過的），可能增加日後維護程式的麻煩，最好也納入考量。

### See also

- [Async Guidance](https://github.com/davidfowl/AspNetCoreDiagnosticScenarios/blob/master/AsyncGuidance.md) by David Fowler
- [Async/Await - Best Practices in Asynchronous Programming](https://learn.microsoft.com/en-us/archive/msdn-magazine/2013/march/async-await-best-practices-in-asynchronous-programming) by Stephen Cleary
