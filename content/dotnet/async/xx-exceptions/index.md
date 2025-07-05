---
title: 非同步方法與異常處理
draft: true
weight: 15
---

本章將討論異常處理，包括非同步操作出現異常的時候所面臨的挑戰，以及 .NET 如何解決這些問題。最後會介紹一些注意事項。

## 異常與非同步呼叫 {#exceptions-and-async-code}

關於非同步呼叫的異常處理，在微軟文件中有這麼一段：

> [!quote] 引述
> 非同步方法應該只在回應用法錯誤（usage error）時才拋出異常。使用錯誤不應該出現在正式環境（production）的程式碼中。舉例來說，如果在呼叫非同步方法時傳入 `null` 給某個參數而導致錯誤（通常以 `ArgumentNullException` 異常表示），你可以修改呼叫端的程式碼以確保永遠不會傳入 `null`。對於所有其他錯誤，非同步方法執行時發生的異常應該被指派給回傳的工作（task），即使非同步方法在工作回傳之前就已同步完成也一樣。通常一個工作僅包含一個異常。然而，如果該工作涉及多項操作（例如使用了 `WhenAll` 方法），則該工作可能會關聯多個異常。
>
> 來源：[Task-based asynchronous pattern (TAP) in .NET: Introduction and overview](https://learn.microsoft.com/en-us/dotnet/standard/asynchronous-programming-patterns/task-based-asynchronous-pattern-tap)

最後一句話：「通常一個工作僅包含一個異常。然而，如果該工作涉及多項操作（例如使用了 `WhenAll` 方法），則該工作可能會關聯多個異常。」這是什意思呢？

在進一步解釋之前，先提一個重點：**一個非同步方法可以拋出一般的異常，也能透過 `Task` 物件來回報錯誤。** 以下小節將詳細說明這句話的涵義。

先來看一個範例：

```cs
public static void Main()
{
    Console.WriteLine("Main thread ID: " + Thread.CurrentThread.ManagedThreadId);

    MyMethod();

    Thread.Sleep(500); // 確保來得及輸出 Worker thread ID。
}

static void MyMethod()
{
    try
    {
        var task = Task.Run(() =>
        {
            Console.WriteLine("Worker thread ID: " + Thread.CurrentThread.ManagedThreadId);
            throw new NotImplementedException();
        });

        // 將以下程式碼取消註解，再觀察執行結果。
        // task.Wait();
    }
    catch (Exception ex)
    {
        Console.WriteLine("捕捉到異常! " + ex.GetType());
    }
}
```

執行結果：

```text
Main thread ID: 1
Worker thread ID: 4
```

執行結果並沒有出現「捕捉到異常!」。但如果把範例程式中的 `// task.Wait();` 取消註解，再執行一次程式，則輸出結果會變成：

```text
Main thread ID: 1
Worker thread ID: 4
捕捉到異常! System.AggregateException
```

> [Try it on .NET Fiddle](https://dotnetfiddle.net/qgn2VI)

**解釋：**

- 雖然 `throw` 語句位於 `try` 區塊內，但此程式碼並非執行於主執行緒，而是執行於另一條執行緒。這是因為傳遞給 `Task.Run` 方法的 lambda 表達式會被編譯器拆成一個完全不同的方法，並且由另一條執行緒來執行。因此，這裡拋出的異常並不會被我們的 `catch` 區塊捕捉到。
- 一旦使用了 `Task.Wait` 方法（無論是靜態方法還是 instance 方法），異常就會被傳遞至呼叫端的執行緒，於是能夠被我們的 `try/catch` 區塊捕捉到。一個 `Task` 物件可能涉及多個非同步工作（例如把多個非同步工作傳入 `Task.WhereAll` 方法），亦即可能拋出多個異常，故 .NET 用一個 `AggregateException` 物件來保存相關的非同步工作所拋出的異常（即使只有一個異常也是如此）。

> 參閱微軟文件：[AggregateException 類別](https://learn.microsoft.com/zh-tw/dotnet/api/system.aggregateexception)

### AggregateException

`Task.Exception` 屬性的定義如下：

```cs
public AggregateException? Exception { get; }
```

`AggregateException` 類別就如其名稱所揭示的，是一個能夠保存多個 exceptions 的 exception 類別。如果某個 `Task` 順利執行完畢，過程中沒有拋出任何異常，那麼它的 `Exception` 屬性便會傳回 `null`。

然而，這個屬性實際上很少用到，以至於如果你用 `await` 來等待一個非同步工作，而那個 `Task` 失敗了，`await` 將會拋出 `AggregateException` 內的第一個異常，而不是整個 `AggregateException` 物件。也就是說，如果 `AggregateException` 內有多個異常，`await` 仍然只會拋出第一個異常並忽略其餘的異常。除了第一個異常之外的所有異常及其內部儲存的任何資訊都將遺失。以下程式碼展示了 `await` 如何拋出存儲的異常。

// TODO


以下程式碼展示了如何從 `AggregateException` 物件中取出所有內部的異常：

```cs
    try
    {
        var task = Task.Run(() =>  { ....(略) });
        task.Wait();
    }
    catch (AggregateException ae)
    {
        foreach (var ex in ae.InnerExceptions)
        {
            // 處理自訂異常。
            if (ex is CustomException)
            {
                Console.WriteLine(ex.Message);
            }
            // 若是其他異常類型，便再次拋出。
            else
            {
                throw ex;
            }
        }
    }
```

如果不想呼叫 `Task.Wait` 方法來將非同步工作的異常傳遞至呼叫端執行緒，另一種做法是透過 `Task` 物件的 `Exception` 屬性來取得 `AggregateException` 及其相關資訊。範例：

```cs
    var task = Task.Run(() =>  { ....(略) });

    while (!task.IsCompleted) { } // 僅作為示範，不建議這麼寫！

    if (task.Status == TaskStatus.Faulted)
    {
        foreach (var ex in task.Exception?.InnerExceptions ?? new(Array.Empty<Exception>()))
        {
            // 處理自訂異常。
            if (ex is CustomException)
            {
                Console.WriteLine(ex.Message);
            }
            // 若是其他異常類型，便再次拋出。
            else
            {
                throw ex;
            }
        }
    }
```

如果沒有存取 `Task` 物件的狀態（例如 `Exception` 屬性、`Result` 屬性），也不使用 `await` 或 `Task.Wait` 來等待非同步工作——換言之，呼叫端程式從未捕捉這些非同步工作的異常，那麼這些異常就會靜靜地藏在 `Task` 物件內部，就像什麼事都沒發生過（應用程式也不會異常終止）。

> 參閱微軟文件：[Exception handling (Task Parallel Library)](https://learn.microsoft.com/en-us/dotnet/standard/parallel-programming/exception-handling-task-parallel-library)

現在我們知道 .NET 會把非同步工作拋出的異常（一個或多個）集中保存於 `Task.Exception` 屬性，其型別為 `AggregateException`。那麼底下這段程式碼是否也是如此呢？

```cs
public async Task<int> MyMethod()
{
    throw new NotImplementedException();
}
```

此範例雖然會產生編譯警告：「This async method lacks 'await' operators and will run synchronously.」但的確是可以通過編譯且可以執行的。然而，程式執行時就只是拋出一般的異常，背後不會有「把異常從工作執行緒傳遞至呼叫端執行緒」的動作，因為編譯器不會對這種寫法產生非同步處理的程式碼。

請記住：**把一個方法加上 `async` 關鍵字不代表它一定會以非同步的方式執行；它只是一個旗號，告訴編譯器必須為 `await` 敘述產生必要的程式碼。** 既然這裡的 `async` 方法並沒有用到 `await` 關鍵字，編譯器就只會單純把方法的回傳值包在 `Task` 物件，而這裡的 `throw` 敘述經過編譯之後也沒有產生額外的程式碼，就只是跟一般（同步的）方法一樣拋出一個普通的異常。

那麼，如果有加上 `await` 呢？如以範例。

```cs
public async Task<int> MyMethod()
{
    await File.ReadAllBytesAsync("file.txt");
    throw new NotImplementedException();
}
```

就如稍早提過的， 只要有使用 `await` 或者 `Task.Wait` 方法（無論是靜態方法還是 instance 方法），非同步工作引發的異常就會被傳遞至呼叫端的執行緒，於是能夠被呼叫端的 `try/catch` 區塊捕捉到。

換言之，只要使用了正確的 `async/await` 語法，在多數比較單純的場景都可以用我們熟悉的 `try/catch` 語法來捕捉和處理非同步工作所引發的異常，如以下範例所示。

```cs
static async Task Main()
{
    try
    {
        await MyMethod();
    }
    catch (Exception ex)
    {
        Console.WriteLine("捕捉到異常! " + ex.GetType());
    }
}

static async Task<int> MyMethod()
{
    await File.ReadAllBytesAsync("file.txt");
    throw new NotImplementedException();
}
```

此方法會以非同步方式呼叫 `File.ReadAllBytesAsync` 並等待其執行結果。待非同步工作完成後，便緊接著拋出異常。若檔案 "file.txt" 不存在，程式執行時的輸出結果會是：

```text
捕捉到異常! System.IO.FileNotFoundException
```

> [Try it on .NET Fiddle](https://dotnetfiddle.net/F9fy4k)

---

> [!note] 重點整理
> 只要在呼叫非同步方法時使用 `await`，呼叫端便可以使用一般的 `try/catch` 語法來應付大多數的錯誤處理。如果沒有使用 `await`，例如使用 `Task.WhenAny` 或 `Task.WhenAll` 來蒐集多個非同步工作的結果，那就必須了解稍早介紹的 `AggregateException` 的用法。

到目前為止，可以說詳細解釋了本章開頭提到的一個重點：「一個非同步方法可以拋出一般的異常，也能透過 `Task` 物件來回報錯誤。」

## await 與 AggregateException

使用 `await` 來重新拋出異常和使用 `Task.Exception` 屬性之間只有一個區別，那就是它們如何使用 `AggregateException`。



---

stubs 待整理

> [!note]
> 另外要提醒的是，「延續」（continuation）操作通常在方法返回之後才開始執行，故如果其中的操作會拋出異常，從非同步方法（即此例的 `MyMethod`）回傳的 `Task` 物件的狀態通常會是 `Created`、`WaitingForActivation` 或 `Running`，並且在稍後才會變成 `Faulted` 狀態。
