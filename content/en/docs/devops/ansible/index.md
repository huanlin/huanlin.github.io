---
title: Ansible 簡介
weight: 7
description: >
  簡單介紹 Ansible 的用途、主要構成元件，以及工作腳本（Playbook）的寫法。
---

## Ansible 是什麼？

Ansible 是用來自動化 IT 日常工作的工具。

在複雜的 IT 作業環境中，可能有數十台甚至上百台的伺服器，有些 IT 維護工作很繁瑣，如果由人工手動執行，不僅費時費力，還容易出錯。

舉例來說，假設我們的應用程式分散部署於 10 台伺服器，現在應用程式準備發布 2.0 版，所以要逐一將那 10 台主機上面的應用程式更新至 2.0 版。或者某天發現需要升級所有主機的 Docker 版本，於是同樣的工作也必須在每一台主機上面執行一遍。還有許多其他類似的重複性工作，像是備份、建立或刪除使用者帳號、修改使用者權限或隸屬的群組、重啟系統等等。

如果以人工方式來執行剛才說的那些重複工作，工程師可能會用 SSH 連接第一台主機，執行完所有必要的維護操作之後，接著 SSH 到第二台主機執行一遍相同的操作，直到 10 台主機全部跑過一遍。有些操作比較簡單，這種作業方式或許還能勉力維持。但是對於某些涉及複雜步驟的操作，工程師往往得寫一份筆記，把詳細的步驟記在紙上或某個文字檔案裡，以免搞錯順序或遺漏某些步驟。即便有準備筆記，工程師也可能在執行到其中某個步驟時，因為某些因素而暫時離開；等到他回來準備接著繼續剩下的步驟時，還得確保自己記得上次離開之前做到哪個步驟。凡此種種，都是人工作業可能碰到的麻煩。

針對上述情境，使用 Ansilbe 能夠提供以下好處：

1. 只要在一台機器上就能對多台主機執行相同操作，而不用逐一遠端登入每一台主機。
2. 相較於撰寫臨時的操作步驟，工程師能夠以一致且簡潔易懂的 YAML 語法來編寫任務腳本。
3. 寫好的 Ansible 腳本可以重複使用，也可以用於不同的環境，例如開發環境、預備環境、正式環境。
4. 更加可靠，亦可減少因為人工失誤所造成的損失。

{{% admonition type=note title="Note" %}}
同類型工具如 Puppet 和 Chef 是採用 Ruby 語言來撰寫腳本，而 Ansible 採用簡單易學的 YAML 便有相對優勢。
{{% /admonition %}}

此外，Ansible 支援所有常見的作業系統和許多雲端平台，故大部分與 IT 維運有關的工作都可以透過 Ansilbe 來執行。

跟一些同類工具（例如 Pupper 和 Chef）相比，Ansible 還有一個優點：只需要在一台主控電腦上面安裝 Ansilbe，便可自動操作多台機器。也就是說，其他被操控的機器上面並不需要預先安裝任何 Ansible 的元件或服務，將來自然也不會有逐台機器升級 Ansilbe 元件的成本。

{{% admonition type=note title="Note" %}}
剛才說的主控電腦（有安裝 Ansible），在官方文件中叫做控制節點（control node）；而被操控的機器（不用安裝 Ansible）則稱為受管理節點（managed nodes）或 hosts。
{{% /admonition %}}

現在我們已經知道 Ansible 是個協助 IT 人員的自動化工具，也了解它有哪些優點，接著來看看它是如何運作的。

## Ansible 的運作方式

Ansilbe 是倚賴**模組**（modules）來執行實際的工作。

基本上，模組是負責執行特定任務的一些小程式。Ansible 所在的主控機器會將那些模組傳送至目標機器，然後執行它們的工作，並於任務完成後刪除。

之所以說模組是小程式，是因為每一個模組都只負責執行一種任務。比如說，用來建立或複製檔案的模組、負責安裝 Nginx 的模組、用來啟動 Docker 容器的模組、或者建立雲端伺服器……等等。你可以從 Ansible 網站上找到一份[模組清單](https://docs.ansible.com/ansible/2.9/modules/list_of_all_modules.html)，以及它們的使用範例。

以下是 [Jenkins 模組](https://docs.ansible.com/ansible/latest/collections/community/general/jenkins_job_module.html)的使用範例：

```yaml
- name: Create a jenkins job using basic authentication
  community.general.jenkins_job:
    config: "{{ lookup('file', 'templates/test.xml') }}"
    name: test
    password: admin
    url: http://localhost:8080
    user: admin

- name: Delete a jenkins job using the token
  community.general.jenkins_job:
    name: test
    token: asdfasfasfasdfasdfadfasfasdfasdfc
    state: absent
    url: http://localhost:8080
    user: admin    
```

此範例只定義了兩個工作，一個是建立 Jenkins job，另一個是刪除 Jenkins job。

再看一個 [Docker 模組](https://docs.ansible.com/ansible/2.9/modules/docker_container_module.html)的使用範例：

```yaml
- name: Create a data container
  docker_container:
    name: mydata
    image: busybox
    volumes:
      - /data

- name: Restart a container
  docker_container:
    name: myapplication
    image: someuser/appimage
    state: started
    restart: yes
    links:
      - "myredis:aliasedredis"
    devices:
      - "/dev/sda:/dev/xvda:rwm"
    ports:
      - "8080:9000"
      - "127.0.0.1:8081:9001/udp"
    env:
      SECRET_KEY: "ssssh"
      BOOLEAN_KEY: "yes"
```

其中定義了兩個工作：建立容器和重啟容器。

從上面兩個範例也能看得出來，就如前面提過的：一個模組只負責某種特定的任務。然而，實務上我們可能會需要調派許多不同的模組，而且必須按照特定順序來執行，才能完成一項比較複雜的需求。那要如何編寫這些模組的任務腳本呢？這就必須了解 Ansible 的 Playbook。

## Playbook

在 Ansible 中，每一個**工作**（task）代表一個需要執行的動作，而多項工作集合起來的任務腳本，就叫做 **Playbook**。一個 Playbook 就是一個 YAML 檔案。

先看一個簡單範例：

```yaml
tasks:
  - name: Rename table foo to bar
    postgresql_table:
      table: foo
      rename: bar
  - name: Set owner to someuser
    postgresql_table:
      name: foo
      owner: john
  - name: Truncate table foo
    postgresql_table:
      name: foo
      truncate: yes
```

在此範例中，`tasks` 是一份工作清單，它包含了三項工作，每一項工作都以 `name` 屬性來簡單描述其工作內容。這三個工作都是使用 `postgresql_table` 模組，而且各自設定了必要的參數，例如 `table`、`rename`、`owner` 等等。

{{% admonition type=note title="Note" %}}
Ansible 2.1 版之後，官方建議盡量寫模組的全名，以免模組名稱重複而產生問題。此範例的 `postgresql_table` 模組的全名是 `community.postgresql.postgresql_table`。
{{% /admonition %}}

可是，要在哪些機器上執行這些工作、以及用什麼身分來執行那些工作呢？顯然上面的範例還缺少了一些東西。

底下是完整範例：

```yaml linenums="1" hl_lines="1-3"
- name: Update web servers
  hosts: databases
  remote_user: root

  tasks:
    - name: Rename table foo to bar
      postgresql_table:
        table: foo
        rename: bar
    - name: Set owner to someuser
      postgresql_table:
        name: foo
        owner: john
    - name: Truncate table foo
      postgresql_table:
        name: foo
        truncate: yes
```

增加的是前面三行：

- `name`：簡單扼要描述了任務內容，即下方 `task` 區塊中定義的三項工作合起來所完成的任務。
- `hosts`：代表那些工作要執行在哪些機器上。這裡的 `databases` 只是一個代稱，而它背後所代表的主機究竟定義在何處，稍後會進一步說明。
- `remote_user`：以什麼使用者身分執行工作。

此範例的整個區塊，Ansible 稱之為一個 **Play**。一個 Playbook 檔案裡面可以定義多個 Plays，以便實現更複雜的需求。

剛才的範例還有一個疑問尚未解決：`hosts: databases` 所代表的資料庫伺服器究竟定義在何處？

## Inventory

稍早提過，被操控的機器稱為受管理節點（managed nodes）或 hosts，而前述範例的 `hosts: databases` 表明欲操控的 hosts 是資料庫伺服器，其中的 `databases` 就是定義在一個 host 檔案裡——Ansible 稱之為 **Inventory**。

Inventory（或 host）檔案的內容看起來像這樣：

```txt
10.24.0.100

[databases]
10.24.0.1
10.24.0.2

[webservers]
10.24.0.10
10.24.0.11
```

基本上，就是把一群同類型設備的 IP 位址或主機名稱集合起來，並且給它們取個代表的名稱。於是，Playbook 檔案中只需要寫 host 檔案中的代表名稱，Ansible 就知道是哪些主機。

## 變數

Playbook 檔案中也可以使用 `vars` 來定義變數，以免重覆撰寫相同的內容。接著改寫前面的範例，把出現多次的資料表名稱 `foo` 定義成變數：

```yaml linenums="1" hl_lines="4 5 10 14 18"
- name: Update web servers
  hosts: databases
  remote_user: root
  vars:
    tablename: foo

  tasks:
    - name: Rename table {{ tablename }} to bar
      postgresql_table:
        table: {{ tablename}}
        rename: bar
    - name: Set owner to someuser
      postgresql_table:
        name: {{ tablename}}
        owner: john
    - name: Truncate table foo
      postgresql_table:
        name: {{ tablename}}
        truncate: yes
```

如此一來，萬一日後需要修改資料表名稱，就只要改變數區塊中的定義就行了。

{{% admonition type=info title="Info" %}}
有關 Ansible vs. Terraform 的議題，可參考 [Terraform 簡介](../terraform/terraform-overview.md)。
{{% /admonition %}}

## 參考資料

- 官方文件：[Ansible concepts](https://docs.ansible.com/ansible/2.9/user_guide/basic_concepts.html)