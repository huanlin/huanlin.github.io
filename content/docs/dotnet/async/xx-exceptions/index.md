---
title: 非同步方法與異常處理
draft: true
---

本章將討論異常處理，包括非同步操作出現異常的時候所面臨的挑戰，以及 .NET 如何解決這些問題。最後會介紹一些注意事項。

## 異常與非同步呼叫 {#exceptions-and-async-code}

關於非同步呼叫的異常處理，在微軟文件中有這麼一段：

> 非同步方法應該只在回應用法錯誤（usage error）時才拋出異常。使用錯誤不應該出現在生產環境的程式碼中。舉例來說，如果在呼叫非同步方法時傳入 `null` 給某個參數而導致錯誤（通常以 `ArgumentNullException` 異常表示），你可以修改呼叫端的程式碼以確保永遠不會傳入 `null`。對於所有其他錯誤，非同步方法執行時發生的異常應該被指派給回傳的工作（task），即使非同步方法在工作回傳之前就已同步完成也一樣。通常一個工作僅包含一個異常。然而，如果該工作涉及多項操作（例如 `WhenAll`），則該工作可能會關聯多個異常。
>
> 來源：[Task-based asynchronous pattern (TAP) in .NET: Introduction and overview](https://learn.microsoft.com/en-us/dotnet/standard/asynchronous-programming-patterns/task-based-asynchronous-pattern-tap)

最後一句話：「通常一個工作僅包含一個異常。然而，如果該工作涉及多項操作（例如 `WhenAll`），則該工作可能會關聯多個異常。」這是什意思呢？

在進一步解釋之前，先提一個重點：**一個非同步方法可以拋出一般的異常，也能透過 `Task` 物件來回報錯誤。** 接著請看以下範例和解說來嘗試理解這句話的意思。

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
        Console.WriteLine("捕捉到異常! " + ex.GetType().FullName);
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
- 一旦使用了 `Task.Wait` 方法（無論是靜態方法還是 instance 方法），異常就會被傳遞至呼叫端的執行緒，於是能夠被我們的 `try-catch` 區塊捕捉到。當程式中同時運行多個非同步工作，便可能拋出多個異常，故 .NET 將這些非同步工作產生的異常全都蒐集放到一個 `AggregateException` 物件中（即使只有一個異常也是這樣）。

### AggregateException

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

如果不想呼叫 `Task.Wait` 方法來傳遞非同步工作的異常，另一種做法是透過 `Task` 物件的 `Exception` 屬性來取得 `AggregateException` 及其相關資訊。範例：

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

如果你沒有存取 `Task` 物件的狀態（例如 `Exception` 屬性、`Result` 屬性），也不使用 `await` 或 `Task.Wait` 來等待非同步工作——換言之，呼叫端程式從未捕捉這些非同步工作的異常，那麼這些異常就會靜靜地藏在 `Task` 物件內部，就像什麼事都沒發生一樣（應用程式也不會異常終止）。

> 參閱微軟文件：[Exception handling (Task Parallel Library)](https://learn.microsoft.com/en-us/dotnet/standard/parallel-programming/exception-handling-task-parallel-library)


非同步（async）方法也有同樣的情形，因為 `await` 等同於呼叫 `ContinueWith`，所以如果我們寫一個簡單的 `async` 方法並拋出異常，像這樣：

```cs
static async Task MyMethod()
{
    try
    {
        await new HttpClient().GetStringAsync("https://microsoft.com");
        throw new NotImplementedException();
    }
    catch
    {
        Console.WriteLine("捕捉到異常!");
    }
}
```

執行結果跟上一個範例一樣不會出現「捕捉到異常!」。

> [Try it on .NET Fiddle](https://dotnetfiddle.net/F9fy4k)

此範例在呼叫 `File.ReadAllBytesAsync` 之後便立刻拋出異常。若以 `ContinueWith` 來改寫 `await`，程式碼會變成：

```cs
try
{
    File.ReadAllBytes("file.txt").ContinueWith(()=>
    {
        throw new NotImplementedException();
    });
}
catch
{
    Console.WriteLine("捕捉到異常!");
}
```

> 以上程式碼可通過編譯，但會有編譯警告。

這寫法跟稍早的 `Task.Run` 範例有同樣的問題：`throw` 語句寫在一個 lambda 函式中，而該匿名函式會由 `ContinueWith` 執行，而不是在 `try` 區塊的同一條執行緒上面執行，故這裡拋出的異常並不會由 `catch` 區塊捕捉到。於是，編譯器會替這段程式碼產生一段重複的 `try-catch` 敘述，類似這樣：

```cs
try
{
    File.ReadAllBytesAsync("file.txt").ContinueWith(()=>
    {
        try
        {
            throw new NotImplementedException();
        }
        catch
        {
            Console.WriteLine("被 catch 捕捉到");
        }
    });
}
catch
{
    Console.WriteLine("被 catch 捕捉到");
}
```

如此一來，當非同步方法拋出異常，便會跟原本程式碼的寫法有同樣的效果：捕捉到異常並輸出一段文字。

可是，並不是所有的 `try-catch` 寫法都能用剛才那樣複製程式碼的方式處理。比如說，`try-catch` 區塊寫在更上層的呼叫端。這樣的話，編譯器便無法決定上層的 `catch` 區塊應該長什麼樣（上層呼叫端的寫法可能很多種）。

再來看一個例子：

```cs
public async Task<int> MyMethod()
{
    DoSomething();
    throw new NotImplementedException();
}
```

對此寫法，編譯器可以說幾乎沒有對程式碼動任何手腳，而且會產生編譯警告："This async method lacks 'await' operators and will run synchronously."。請記住：**把一個方法加上 `async` 關鍵字不代表它一定會以非同步的方式執行；它只是一個旗號，告訴編譯器必須為 `await` 敘述產生必要的程式碼。** 如果 `async` 方法裡面沒有任何 `await` 敘述，編譯器就只是單純把回傳值包在一個 `Task` 物件，而這裡的 `throw` 敘述經過編譯之後也沒有額外處理，就只是跟一般（同步的）方法一樣拋出一個普通的異常。

那麼，如果有加上 `await` 呢？

```cs
public async Task<int> MyMethod()
{
    DoSomething();
    await File.ReadAllBytesAsync("file.txt");
    throw new NotImplementedException();
}
```

此方法會以非同步方式呼叫 `ReadAllBytesAsync` 並等待其執行結果。待非同步工作完成後，便緊接著拋出異常。對此情況，編譯器會替我們加入一個 `try-catch` 區塊，並將捕捉到的異常保存於回傳的 `Task` 物件中。類似以下程式碼：

```cs
public async Task<int> MyMethod()
{
    DoSomething();
    var result = new TaskCompletionSource<int>();
    File.ReadAllBytesAsync("file.txt").ContinueWith(t=>
    {
        try
        {
            throw new NotImplementedException();
        }
        catch(Exception ex)
        {
            result.TrySetException(new AggregateException(ex));
        }
    });
}
```

請注意編譯器額外產生的 `try` 區塊不是加在呼叫非同步方法 `ReadAllBytesAsync` 之前，而是放在傳入 `ContinueWith` 方法的 lambda 函式中。當程式執行時，考慮兩種情況：

- 情況一：在呼叫非同步方法 `ReadAllBytesAsync` 之前就發生錯誤（例如 `DoSomething` 拋出異常），則 `MyMethod` 方法就只是拋出一般的異常。
- 情況二：執行非同步方法 `ReadAllBytesAsync` 的時候出錯，則該異常會被包在 `Task` 物件中一並回傳給呼叫端。

如果在呼叫非同步方法時使用 `await` 來等待非同步工作執行完畢，錯誤處理的程式碼寫法都一樣，故無須在意上述兩種情形有何區別。但如果不是使用 `await`，例如使用 `Task.WhenAny` 或 `Task.WhenAll` 來蒐集多個工作的結果，那就必須了解上述兩種情形的細節差異。

> [!note]
> 另外要提醒的是，「延續」（continuation）操作通常在方法返回之後才開始執行，故如果其中的操作會拋出異常，從非同步方法（即此例的 `MyMethod`）回傳的 `Task` 物件的狀態通常會是 `Created`、`WaitingForActivation` 或 `Running`，並且在稍後才會變成 `Faulted` 狀態。


使用 `await` 來重新拋出異常和使用 `Task.Exception` 屬性之間只有一個區別，那就是它們如何使用 `AggregateException`。

## await 與 AggregateException

