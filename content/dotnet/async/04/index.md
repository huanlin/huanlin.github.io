---
title: 非同步程式設計常用技巧
draft: true
weight: 20
---

> 本章範例程式的原始碼位置：
>
> <https://github.com/huanlin/async-book-support> 裡面的 Examples/ch04 資料夾。

---

## 非同步方法的各種樣貌

本節要回答下列幾個問題：

- 在介面（interface）中如何定義非同步方法？
- 有「非同步屬性」嗎？
- 建構函式（constructor）也可以宣告成 `async` 嗎？
- 如何撰寫非同步匿名方法（async lambda）？

### 非同步介面方法

定義介面或者抽象方法時，不需要、也不可以使用 `async` 關鍵字。除此之外，其餘特徵與實作非同步方法的時候相同：函式名稱以「Async」結尾，而且回傳型別是 `Task` 或 `Task<T>`。

範例：

```cs
public interface IMyAsyncOperations
{
    // 注意這裡沒有使用 async 關鍵字!
    Task<string> GetDataAsync(string url);
}
```

實作介面時，則使用 `async` 來修飾非同步方法，例如：

```cs
public class MyAsyncOperations : IMyAsyncOperations
{
    public async Task<string> GetDataAsync(string url)
    {
        var client = new HttpClient();
        string content = await client.GetStringAsync(url);
        return content;
    }
}
```

請注意 `await` 等待的是某個方法所回傳的 `Task` 物件，而不管那個方法是不是 `async` 方法。因此，在介面中定義方法時，只要它回傳的物件類型是 `Task` 或 `Task<T>`，該物件便可以被 `await`。

### 非同步屬性？

實作非同步方法時需要加上 `async` 關鍵字，那麼非同步屬性呢？嗯，沒有這種東西。不過，屬性可以傳回 `Task` 類型的物件，而所有 `Task` 物件都可以被 `await`，因此，便可能出現類似底下範例的寫法：

```cs
public class MyAsyncOperations : IMyAsyncOperations
{
    public Task<string> Data
    {
        get { return GetDataAsync(); } // 註：沒有 async get 這種語法。
    }

    public async Task<string> GetDataAsync(string url)
    {
        // 略
    }
}
```

這樣好像實現了非同步屬性（可以 `await` 的屬性），然而它骨子裡其實是去呼叫另一個非同步方法。使用這個屬性的人可能會以為它返回的值已經保存在某個地方，卻沒想到其實每次存取該屬性時，背後會去呼叫一個方法。既然如此，那又何必硬要透過 `Task` 來偽裝呢？直接寫成非同步方法會更好。

### 非同步建構函式？

建構函式不可以加上 `async` 關鍵字（那麼它自然也就不能使用 `await`）。如果你想要在建立物件時，以非同步的方式來初始化物件的狀態，一種常見的作法是撰寫非同步工廠方法。參考以下範例：

```cs
class Foo
{
    private Foo()   // 把建構函式宣告成 private，不讓外界使用。
    {
    }

    // 外界必須透過此方法才能建立物件。
    public static Task<Foo> CreateInstanceAsync()
    {
        var foo = new Foo();          // 建立物件。
        return foo.InitializeAsync(); // 初始化。
    }

    // 初始化的工作都寫在這個非同步方法裡面。
    private async Task<Foo> InitializeAsync()
    {
        await DomSomeTaskAsync();
        return this;
    }
}
```

範例程式中的註解應足以說明一切，就不另外解釋了。

### 非同步 Lambda

只要在 Lambda 表示式前面加上 `async` 關鍵字，它就成了非同步匿名函式。

```cs
static void Main(string[] args)
{
    // 此委派需傳入一個 string，且回傳一個 Task<string>。
    Func<string, Task<string>> downloadPageAsync =
        async (string url) =>       // 非同步 Lambda
        {
            using (var client = new HttpClient())
            {
                return await client.GetStringAsync(url);
            }
        };

    var task = downloadPageAsync("http://www.huan-lin.blogspot.com");

    Console.WriteLine("網頁長度: {0}", task.Result.Length);
}
```

請注意，`async` Lambda 表示式和一般具名的 `async` 方法一樣，剛開始是同步執行的（亦即執行於呼叫端所在的執行緒），直到碰到第一個 `await` 敘述才會分出不同的控制流。

## 非同步延遲

當我們想要在程式執行時暫停一段時間，經常會使用 `System.Threading.Thread.Sleep()` 方法。此方法會讓當前的執行緒（控制流）暫停，等到指定的等待時間過去，才會繼續往下執行。

若希望在暫停期間不要阻擋當前的執行緒，則可以使用非同步延遲的作法，也就是呼叫靜態方法 `Task.Delay()`。

`Task.Delay()` 可以傳入一個代表毫秒（milliseconds）的整數，也可以傳入代表時間長度的 TimeSpan 物件。

範例：

```cs
static async Task Main(string[] args)
{
    var task = MyTaskAsync();
    Console.WriteLine("已經返回 Main()");
    await task; // 等待工作完成
    Console.ReadLine();
}

static async Task MyTaskAsync()
{
    await Task.Delay(TimeSpan.FromSeconds(1)); // 非同步延遲 1 秒。
    Console.WriteLine("非同步工作結束"); // 延遲 1 秒後，回到這裡繼續執行。
}
```

> 範例程式：Ex01_AsyncDelay.csproj

執行結果：

```text
已經返回 Main()
非同步工作結束
```

### 重試機制

有時候，當某個非同步工作失敗了，我們可能不希望應用程式立刻拋出異常，而會想要等待一段時間再重試一次。此時 `Task.Delay()` 亦可派上用場。參考以下範例：

```cs
static async Task<string> MyDownloadPageAsync(string url)
{
    const int MaxRetryCount = 3;  // 最多重試 3 次

    using (var client = new HttpClient())
    {
        for (int i = 0; i < MaxRetryCount; i++)
        {
            try
            {
                return await client.GetStringAsync(url);
            }
            catch (Exception ex)
            {
                // 忽略錯誤。
                Console.WriteLine("第 {0} 次失敗: {1}", i+1, ex.Message);
            }
            await Task.Delay(TimeSpan.FromSeconds(i + 1));
        }

        // 最後一次失敗就讓它拋出異常。
        return await client.GetStringAsync(url);
    }
}
```

> 範例程式：Ch04/AsyncDelayAndRetry.csproj

此非同步方法會嘗試取得指定網址的網頁內容，而且每當發生錯誤，它會忽略錯誤並延遲一段時間，然後重試一次。在 `for` 迴圈中，每失敗一次，延遲的時間就多加一秒。當反覆失敗且重試次數達到三次，便離開這個反覆重試的迴圈。

## 傳回已完成的工作

`Task` 類別有一個靜態方法可用來傳回已完成的工作，這個方法是 `FromResult()`。底下是一個簡單範例：

```cs
Task<int> task1 = Task.FromResult(10);
```

什麼時候會用到它呢？

一個常見的應用場合是：在非同步方法中，希望以同步的方式傳回結果。舉例來說，假設要設計一個支援非同步操作的快取物件，你先定義了如下介面：

```cs
interface IMyCache
{
    Task<string> GetDataAsync();
}
```

這表示，你預期將來實作 `IMyCache` 介面時，`GetDataAsync` 會是個非同步方法。然而，在某些情況下，例如撰寫單元測試或模擬物件（mock），你可能不需要複雜的非同步操作，而只需要直接傳回一個現成的結果。那麼，`Task.FromResult` 便可能派上用場。如下所示：

```cs
class MyCache : IMyCache
{
    public async Task<string> GetDataAsync()
    {
        // 這裡的 async 和 await 是多餘的，稍後會說明。
        return await Task.FromResult("hello");
    }
}
```

以下程式碼則示範了如何在 Console 應用程式中使用 `MyCache` 類別的 `GetDataAsync` 方法：

```cs
static async Main(string[] args)
{
    Task.Run(async () =>
    {
        IMyCache myCache = new MyCache();
        var s = await myCache.GetDataAsync();
        System.Console.WriteLine(s);
    });

    Console.ReadLine();
}
```

> 範例程式：Ch04/ReturnCompletedTask.csproj

上述範例的寫法其實有個問題：在沒有非同步操作的地方多餘地使用了 `async` 和 `await` 關鍵字。接著要繼續討論這個問題。

### 不要 await Task.FromResult()

續上文，假設你就是負責撰寫 `MyCache` 類別的人，那麼，你自然知道自己寫的 `GetDataAsync` 方法其實並沒有執行任何非同步工作，因為它是呼叫 `Task.FromResult()` 來直接（立刻）傳回一個已經完成的工作。然而，剛才的 `GetDataAsync` 方法卻使用了 `async` 和 `await` 關鍵字，而這兩個關鍵字會使編譯器產生一些額外的程式碼，以便用來處理非同步呼叫的等待以及環境切換等處理。既然 `Task.FromResult` 會直接傳回現成的 `Task` 物件，那麼編譯器為非同步呼叫所產生的額外程式碼也就都是多餘的了。

簡單地說，我們應該遵循這個建議：不要 `await` 一個 `Task.FromResult()` 呼叫。

因此，先前的 `GetDataAsync` 方法應該把 `async` 和 `await` 去掉，變成這樣：

```cs
public Task<string> GetDataAsync()
{
    return Task.FromResult("hello");
}
```

你可以看到，在定義 `IMyCache` 介面時，雖然預期 `GetDataAsync` 是個非同步方法（注意方法名稱以 `Async` 結尾），但實作時仍有可能採取同步的方式——當你碰到這樣的寫法時，請特別留意是否無意間在非同步呼叫的流程中混雜了同步／阻斷式呼叫，例如底下這個錯誤示範：

```cs
public Task<string> GetDataAsync()
{
    string s = System.IO.File.ReadAllText(@"C:\temp\big.txt"); // 錯誤示範!
    return Task.FromResult(s);
}
```

> 第 3 章〈不要寫假的 `async` 方法〉一節中曾經討論過相關議題，如果已經忘得差不多，不妨回頭複習一下。

## 等待工作完成

當應用程式有多個非同步工作正在執行，你可能會需要等待其中特定的幾項工作完成後，才接續執行其他工作；或者，也有可能是另一種情況：只要任何一個非同步工作完成了，就接著執行後續工作。以下分別示範這兩個小技巧。

### 等待一組工作完成

欲等待一組非同步工作全部完成，可使用 `await` 關鍵字搭配 `Task.WhenAll` 方法。`Task.WhenAll` 接受不定個數的 `Task` 物件，並且會在這些工作完成後，傳回一個 `Task` 物件，讓我們可以用 `await` 來等待這個工作。範例如下：

```cs
Task task1 = Task.Delay(1000);
Task task2 = Task.Delay(2000);
Task task3 = Task.Delay(3000);

await Task.WhenAll(task1, task2, task3);

// 三個任務都完成後，控制流會回到這裡繼續執行。
```

如果所有工作都執行成功，而且它們的結果都是屬於同一種型別，那麼 `Task.WhenAll` 會傳回一個陣列，其中包含所有工作的結果。

範例：

```cs
Task<int> task1 = Task.FromResult(3);
Task<int> task2 = Task.FromResult(4);
Task<int> task3 = Task.FromResult(5);

int[] results = await Task.WhenAll<int>(task1, task2, task3);

// results 陣列包含三個整數：3, 4, 5。
```

>範例程式：Ch04/AsyncWaitAllTasks.csproj

### 等待任意一個工作完成

如果已經起始了多個非同步工作，而且只要其中任何一項工作完成時就繼續往下執行，此時可使用 `await` 關鍵字搭配 `Task.WhenAny` 方法。比如說，應用程式需要透過網路抓取的資料來源有三個，而且各自存放在不同的主機上，於是建立三個非同步工作來分頭抓取資料，而且只取用最快返回的結果，剩下的則予以忽略。

範例：

```cs
var client = new HttpClient();

Task<string> task1 = client.GetStringAsync("https://microsoft.com");
Task<string> task2 = client.GetStringAsync("https://facebook.com");
Task<string> task3 = client.GetStringAsync("https://google.com");

Task<string> completedTask = await Task.WhenAny(task1, task2, task3);

string content = await completedTask;
Console.WriteLine("網頁長度: " + content.Length);
```

>範例程式：AsyncWaitAnyTask.csproj

當指定的一組工作當中有任何一個已經完成任務，而其他工作既沒有被 `await`，也沒有被取消（稍後會說明如何取消非同步工作），那麼它們就會被拋棄。被拋棄的工作仍會繼續執行，直到任務完成為止，但它們的結果都會被忽略，而且拋出的任何異常也都會被忽略。

## 取消工作

有些比較費時的工作，可能因為某些原因（例如使用者改變主意）而需要臨時取消，此時便可使用 TPL 提供的工作取消機制。此機制的主要型別有二：

- `CancellationToken` 結構：本身不具備任何取消工作的方法，而只是作為一種權杖（token），用來表示某工作的取消狀態，以便在各函式呼叫之間傳遞此狀態。
- `CancellationTokenSource` 類別：用來取消工作。它就像個管理員，可以控制多個 `CancellationToken` 物件。當你呼叫 `CancellationTokenSource` 的 `Cancel` 方法，就會令它所管理的所有 `CancellationToken` 物件進入取消狀態。

許多以 `Task` 為基礎的函式（例如稍後介紹的 `ContinueWith` 方法）都有提供參數讓呼叫端提供 `CancellationToken` 物件。要注意的是，一旦某項工作已經開始執行，TPL 是沒辦法讓它中途取消的；如果你在工作開始執行之前就將它取消，那就沒問題—— TPL 會把那項工作從它排定的工作佇列中移除。

如果想要取消某個已經開始執行的工作，你可以在那項工作的委派方法當中檢查 `CancellationToken` 物件的 `IsCancellationRequested` 屬性：若為 true，便中止目前的工作。

> 在應用程式中實現「取消工作」機制時，須了解有些操作一旦執行了，是根本無從取消的。例如透過網路發送 email 或即時訊息，除非在發送的過程當中把訊息暫時保存於某個佇列，並且提供一段「反悔時間」讓使用者有機會按下「取消」按鈕，那就還有可能來得及把訊息從佇列中移除，從而達到取消傳送訊息的效果。

## 工作的延續

這裡所謂「工作的延續」，指的是等到某個非同步工作執行完畢之後才要接著執行的工作；你要它稱為非同步工作的「延續工作」或「接續工作」也可以。有了 `async` 和 `await` 語法的幫忙，在處理延續工作這件事情上，程式寫起來變得輕鬆許多。

假設某些情況下，我們無法使用 `async` 和 `await` 語法，此時雖然能夠使用 `Task` 的 `Result` 屬性來等待工作的執行結果，或者呼叫 `Wait` 方法來等待工作完成，但它們都會阻擋當前的執行緒，如此便失去了非同步工作的意義。那麼，是否還有其他方法能夠通知我們某個非同步工作已經完成、以便繼續執行後續的工作？答案就是 `Task` 相關類別所提供的 `ContinueWith` 多載方法。

`ContinueWith` 方法可以讓我們輕易的銜接兩個非同步工作，例如：

```cs
webDownloadTask.ContinueWith(task =>
    {
        string content = task.Result;
        Console.WriteLine($"網頁內容長度：{content.Length}");
    });
```

`ContinueWith` 方法會建立新的 `Task` 物件來封裝接續的工作，而且這些工作在創建之時，是處於等待狀態（`TaskStatus.WaitingForActivation`），須等到前項工作（如上面範例的 `webDownloadTask`）執行完畢才會開始執行。此外，在預設情況下，透過 `ContinueWith` 方法所建立的工作都會動用到執行緒（由 `TaskScheduler` 從執行緒集區調用執行緒）；也就是說，它們都會在各自的執行緒上面執行（無論前項工作是否有動用執行緒）。

由於 `ContinueWith` 方法會傳回它所建立的 `Task` 物件，而該物件也能夠拿來做為其他工作的前項工作（antecedent），故可層層串接，不斷延續。例如：`task.ContinueWith(...).ContinueWith(...)`。

### 串接與組合多項工作

實務上，應用程式執行時可能會起始多個非同步工作，而這些工作之間亦可能有先後順序的關係，例如工作 A 執行完畢之後才能執行 B 和 C——亦即非同步工作 B 和 C 是工作 A 的延續。透過 `ContinueWith` 方法，我們便能夠串接任意數量的非同步工作，並確保這些非同步工作的執行順序，或者把多個小工作組合起來完成一項複雜工作。參考以下範例：

```cs
static void Main(string[] args)
{
    var taskA = Task.Run(() => Console.WriteLine("起始工作...."));

    Task taskB = taskA.ContinueWith( antecedentTask =>
        {
            Console.WriteLine("從 Task A 接續 Task B.");
            System.Threading.Thread.Sleep(4000); // 等待幾秒
        });

    Task taskC = taskA.ContinueWith( antecedentTask =>
        {
            Console.WriteLine("從 Task A 接續 Task C.");
        });

    Task taskD = taskA.ContinueWith( antecedentTask =>
        {
            Console.WriteLine("從 Task A 接續開始 Task D.");
        });

    Task.WaitAll(taskB, taskC);

    Console.WriteLine("程式結束");
}
```

> 此範例程式的專案：Ch04/Ex01_TaskContinuation.csproj

在此範例中：

- `taskA` 起始之後，第 8～13 行有一個延續工作，而等到 `taskA` 完成之後，會接著起始兩個延續工作：`taskB` 和 `taskC`。對 `B` 和 `C` 而言，`A` 是他們的前項工作。
- 前項工作執行完畢後，以 `ContinueWith` 方法所指定的延續工作將自動以非同步的方式開始執行。此例的 `taskA` 執行完畢之後，接著會分別以非同步的方式起始 `taskB` 和 `taskC`；至於哪一個延續工作會先執行則不一定（無法在編譯時期確定）。
- 以 `ContinueWith` 方法所指定的延續工作可以透過參數 `antecedentTask` 來獲取前項工作，以便得知前項工作當前的狀態。

此範例程式的執行結果如下（在你的機器上執行時，接續工作 B 和 C 的順序可能不同）：

```text
起始工作....
從 Task A 接續開始 Task D.
從 Task A 接續 Task B.
從 Task A 接續 Task C.
程式結束
```

### `TaskContinuationOptions`

`Task.ContinueWith` 方法有數個多載版本，其中有些可以傳入列舉型別 `TaskContinuationOptions` 來改變預設的延續行為。比如說，你可以使用 `OnlyOnRanToCompletion` 旗號來確保唯有在前項工作順利完成（沒有發生錯誤）的情況下才能執行延續工作，或者用 `OnlyOnFaulted` 和 `OnlyOnCanceled`來指定唯有在前項工作「發生錯誤」或者「被取消」的情況。又如 `NotOnRanToCompletion`，則是限定前項工作「沒有順利完成」的情況才能執行延續工作。

> `TaskContinuationOptions` 的完整定義可參考 [MSDN 線上文件](http://bit.ly/1NkFkhG)。

你可以為一項工作指定多個延續工作，並分別指定某個延續工作負責處理前項工作順利成功的情況，而另一個延續工作則負責處理前項工作發生錯誤的情況。底下是一個簡單範例，示範如何利用 `TaskContinuationOptions` 來設定當前項工作的狀態符合特定條件時才去觸發（起始）某個延續工作——這等於是非同步的事件模型。

```cs
static void Main(string[] args)
{
    Task taskA = Task.Run(
        () => Div(10, 5) );  // 若把第二個參數改成 0，接續的工作會變成 taskOnFailed。

    Task taskOnFailed = taskA.ContinueWith(
        antecedentTask =>
        {
            Console.WriteLine("Task A 已失敗! IsFaulted={0}", antecedentTask.IsFaulted);
        },
        TaskContinuationOptions.OnlyOnFaulted); // 當前項工作失敗時才起始此工作

    Task taskOnCompleted = taskA.ContinueWith(
        antecedentTask =>
        {
            Console.WriteLine("Task A 已完成! IsCompleted={0}", antecedentTask.IsCompleted);
        },
        TaskContinuationOptions.OnlyOnRanToCompletion); // 當前項工作完成後才起始此工作

    taskOnCompleted.Wait();
}

static int Div(int dividend, int divisor)
{
    return dividend / divisor;
}
```

>此範例程式的專案：Ch03/Ex02_TaskContinuationOptions.csproj

在此範例中，唯有當 `taskA` 順利執行完畢，才會接著起始 `taskOnCompleted` 所代表的非同步工作。如果 `taskA` 執行失敗，則只有 `taskOnFailed` 工作會接著執行。

當你碰到需要「射後不理」（fire-and-forget）的非同步工作時，這種類似觸發事件的寫法便可派上用場。

> 你還可以使用 `ExecuteSynchronously` 旗號來告訴工作排程器：當前項工作執行完畢時，請立刻以當前的執行緒來執行此延續工作，而不要將它排入佇列，亦即不要將它執行於另一條執行緒。一般而言，只有在延續工作需要盡快執行的情況才需要使用這個旗號。

### 取消「中間的」延續工作

預設情況下，當你取消一項工作，則與該項工作關聯的「延續工作」都會立刻變成可執行的狀態（這句話可以簡單理解成：一旦前項工作取消，它的子工作便會「幾乎」立刻接著執行）。

考慮一個場景：假設 TaskB 是 TaskA 的延續工作，而 Task C 又是 TaskB 的延續工作，即三者的執行順序是 TaskA 然後 TaskB 然後 TaskC。理論上，延續工作應該會等到前項工作執行完畢之後才會執行，故 TaskB 以及 TaskC 都是在 TaskA 完成之後才會執行。

然而，當 TaskB 被中途取消、而且 TaskA 尚未完成，此時便很可能會出現一個特殊情形：TaskA 和 TaskC 都會繼續執行，可是 TaskC 會比它的「前項的前項工作」TaskA 更早執行完畢。請看以下範例：

```cs
static void Main(string[] args)
{
    Task taskA = Task.Run(DoSomething);
    var cancelManager = new CancellationTokenSource();
    Task taskB = taskA.ContinueWith(
        _ => Console.WriteLine("這裡不會執行"),
        cancelManager.Token);
    Task taskC = taskB.ContinueWith(
        _ => Console.WriteLine("TaskC 執行完畢。"));

    cancelManager.Cancel(); // 取消 taskB
    Console.ReadKey();
}

static void DoSomething()
{
    Thread.Sleep(1500);
    Console.WriteLine("TaskA 執行完畢。");
}
```

執行結果是：

```text
TaskC 執行完畢。
TaskA 執行完畢。
```

在此範例中，taskA 起始之後，接著立刻建立了延續工作 taskB，然後又建立 taskB 的延續工作 taskC，最後緊接著取消 taskB（第 11 行）。這裡用 ThreadSleep 方法來確保 taskA 要花一秒以上的時間才能執行完畢，故在取消 taskB 的時候，taskA 肯定還沒結束；在此同時，taskC 會因為前項工作 taskB 的取消而立刻開始執行，於是便會形成 taskC 比最初之前項工作 taskA 更早執行完畢的情形。

如果你要希望 taskA 不會因為延續工作 taskB 的取消而導致「孫代」工作 taskC 提前執行，可以在建立 taskB 的時候使用 `TaskContinuationOptions` 的 `LazyCancellation` 旗號來告訴工作排程器：當我（taskB）被取消時，請延後這個取消動作，直到前項工作結束之後才執行取消動作。故前面範例可以修改成：

```cs
Task taskB = taskA.ContinueWith(
    _ => Console.WriteLine("這裡不會執行"),
    cancelManager.Token,
    TaskContinuationOptions.LazyCancellation,
    TaskScheduler.Current);
```

結果就會變成 TaskA 先執行完畢，然後才是 TaskC：

```text
TaskA 執行完畢。
TaskC 執行完畢。
```

## 混合使用同步與非同步方法

雖然我們知道最好是從頭開始就一路維持非同步的寫法，然而在真實世界中，卻免不了會碰到需要混和使用的情形，包括：

- 在非同步方法中呼叫同步方法。
- 在同步方法中呼叫非同步方法，並取得非同步工作的結果。

以下分別以兩個小節來說明。

### 在非同步方法中呼叫同步方法

(TODO)

### 在同步方法中呼叫非同步方法

假設你有個非同步方法：

```cs
private async Task DoSomethingAsync(int milleSeconds)
{
    await Task.Delay(milleSeconds);
    // 省略其他程式碼（如果有的話）
}
```

當你需要在某個同步方法中呼叫上述方法時，**不要**這樣寫:

```
private void MyMethod()
{
    DoSomethingAsync(1000).Wait();  // 避免這樣寫!
}
```

若該方法有回傳值，則會是：

```cs
void MyMethod()
{
    var result = DoSomethingAsync(1000).Result; // 避免這樣寫!
}
```

若你在 UI 類型的應用程式（例如 Windows Form）中使用上述寫法，程式很可能會出現鎖死（deadlock）的情形。

稍微好一點的寫法是：

```cs
void MyMethod()
{
    Task.Run(
        async () =>
        {
            await DoSomethingAsync(1000);
        }
    ).Wait();
}
```

而目前已知的一個比較好的寫法，是使用微軟內部使用的工具類別：`AsyncHelper`。其[原始碼](https://github.com/IdentityServer/IdentityServer3.EntityFramework/blob/master/Source/Core.EntityFramework/Serialization/AsyncHelper.cs)如下：

```cs
internal static class AsyncHelper
{
    private static readonly TaskFactory _myTaskFactory =
        new TaskFactory(CancellationToken.None, TaskCreationOptions.None, TaskContinuationOptions.None, TaskScheduler.Default);

    public static void RunSync(Func<Task> func)
    {
        _myTaskFactory.StartNew(func).Unwrap().GetAwaiter().GetResult();
    }

    public static TResult RunSync<TResult>(Func<Task<TResult>> func)
    {
        return _myTaskFactory.StartNew(func).Unwrap().GetAwaiter().GetResult();
    }
}
```

你可以直接把上面的類別貼到自己的專案中使用。此工具類別提供兩個版本的 `RunSync` 方法，一個用於呼叫無回傳值的非同步方法，另一個用於有回傳值的場合。

當你需要在同步方法中呼叫非同步方法時，便可使用如下範例的寫法：

```cs
void MyMethod()
{
    // 若呼叫的非同步方法沒有回傳值：
    AsyncHelper.RunSync(DoSomethingAsync);

    // 若呼叫的非同步方法有回傳值：
    var result = AsyncHelper.RunSync(() => DoSomethingAsync(1000));
}
```

### 不要寫假的 async 方法

儘管我們知道撰寫非同步程式時，應從頭到尾都採用非同步呼叫，但有時候就是難以做到。比如說，假設你正要使用一個現成的函式庫，那個函式庫沒有提供非同步版本的 API，於是，為了可以從頭到尾都採用非同步呼叫，你可能會另外寫一個 `async` 方法來包裝那個同步呼叫的 API。像這樣：

```cs
public async Task MyDownloadPageFakedAsync(string url)
{
    var client = new WebClient();
    var task = Task.Run(() =>
    {
        string content = client.DownloadString(url);
        Console.WriteLine("網頁長度: " + content.Length);
    });
    await task;
}
```

> `WebClient` 有提供非同步的 `DownloadStringTaskAsync` 方法，這裡只是為了方便說明而刻意使用同步呼叫的版本。

使用此函式的人，有可能無法看到函式內部的實作，所以光從函式的宣告來看：有 `async` 關鍵字、返回 `Task` 物件，而且函式名稱以 "Async" 結尾，自然會認為那是個非同步方法。既然是非同步方法，那就不見得會動用執行緒。然而，實際上卻完全不是那麼回事。怎麼說呢？

請注意這裡使用了 `Task.Run()` 方法來建立一個非同步工作，以便在此函式中使用 `await`，以及在宣告時加上 `async`。換句話說，這裡使用了 `Task.Run()` 來把同步執行的工作偽裝成非同步方法。然而，正如第 2 章提過的，當你使用 `Task` 類別來建立非同步工作時，預設的工作排程器會向執行緒集區（thread pool）調動一個工作執行緒來執行任務。如此一來，使用此函式的人會以為它跟其他 `async` 函式一樣，卻不知道它背後其實使用了執行緒集區——如果大量用於 ASP.NET 應用程式中，可能導致效能或者延展性（scalability）不佳的問題（因為跟 ASP.NET 爭搶使用執行緒集區可能會導致集區耗盡）。

因此，如果沒辦法撰寫「真正的」非同步方法，最好別假裝它是。這樣的話，至少別人在使用你的函式庫時，一眼就能判斷那是個同步方法，不至於因為誤用而產生其他問題。

%%## 錯誤處理

(TODO)

## 進度回報

(TODO)

## 錯誤處理

當 `Task` 物件的 `Status` 屬性變成 `Faulted`，即所代表工作執行時「至少發生一個錯誤」。由於 TPL 可以串接組合多項工作，故在執行過程中，每個延續工作都有可能發生錯誤。我們可以透過 `Task` 物件的 `Exception` 屬性來取得一個根工作（root task）的所有錯誤。

`Task` 類別的 `Exception` 屬性的型別是 `AggregateException`，顧名思義，它能夠包含一個或多個 `Exception` 物件。

`AggregateException` 繼承自基礎類別 `Exception`，除了繼承下來的 `InnerException` 屬性，它還增加了複數形式的 `InnerExceptions`（注意後面有個 s），讓我們能夠取得某工作在執行過程當中發生的所有錯誤。

範例：

當你使用 `Task` 物件 `Result` 屬性來取得某工作的執行結果，或者呼叫 `Wait` 方法來等待工作執行完畢，一旦過程當中發生錯誤，就會拋出 `AggregateException`。你可以藉由 `try...catch` 語法來捕捉 `AggregateException`，這和透過 `Task` 物件的 `Exception` 屬性來獲取的執行時期例外是一樣的。

`Task` 物件會記住自己的錯誤是否有被外界獲取，如果沒有，則那些錯誤就會被視為「未觀察的錯誤」（unobserved exceptions）。TPL 會記錄這些沒有人注意到的錯誤，並觸發 `TaskScheduler.UnobservedTaskException` 事件（沒錯，它是個事件，而不是 exception），而此事件是你處理這些錯誤的最後機會。

## 為非同步方法撰寫單元測試

## 重點回顧

- 定義介面或者抽象方法時，不需要、也不可以使用 `async` 關鍵字。
- `await` 等待的是某個方法所回傳的 `Task` 物件，而不管那個方法是不是 `async` 方法。因此，在介面中定義方法時，只要它回傳的物件類型是 `Task` 或 `Task<T>`，該物件便可以被 `await`。
- 不要使用 `Thread.Sleep()`，而應盡量使用 `await Task.Delay()`。
- 不要 `await Task.FromResult()`。
- 欲等待一組非同步工作全部完成，可使用 `await Task.WhenAll()`。
- 欲等待任意一個非同步工作完成，可使用 `await Task.WhenAny()`。
