---
title: Git reset
---

`git reset` 命令可將目前工作副本的 `HEAD` 指向先前的某一次 commit，通常表示要放棄近期的某個或某些 commits，以便將本機的 repository 回復至先前的某個版本。

此命令有三種模式：soft、hard、和 mixed。這裡只介紹 hard 模式，相關細節請參考文件：[git-reset](https://git-scm.com/docs/git-reset)。

範例：

```shell
git reset --hard HEAD~        # 倒退一個 commit
git reset --hard HEAD~2       # 倒退兩個 commit
git reset --hard HEAD~3       # 倒退三個 commit
git reset --hard 6f3efa2d     # 倒退至指定的 commit ID
```

加上 `--hard` 參數表示那些跳過的 commits 全都不要了。

## Push to remote

使用 `git reset` 變更本機的工作副本之後，還必須把變更推送至遠端的來源 repository，否則下次 `git pull` 又會從遠端把剛才跳過的 commits 恢復，並將 `HEAD` 指向最新版本。

推送的命令必須加上 `--force` 參數如下：

```shell
git push -force
```

要提醒的是，如果此分支是一個受保護的分支（GitHub 和 GitLab 皆有此功能），那麼上述命令會失敗。以 GitLab 為例，錯誤訊息會是：

```console
remote: GitLab: You are not allowed to force push code to a protected branch on this project.
```

解決方法是暫時解除該分支的保護：至 GitLab 網站，進入專案的管理頁面，然後進入 Settings > Repository > Protected branches，找到分支後，將它 `Unprotect`，或者把該分支的 `Allowed to force push` 選項開啟。然後再執行一次 `git push -f` 命令。

成功推送至遠端後，記得恢復該分支的保護。

{{< admonition warning "警告" >}}
此操作應該只用於自己獨立工作的分支上，絕對不要任用使用於多人共同協作的分支，例如 main 和 develop。場合。因為你透過強制 push 操作來「刪除」遠端 repo 中的那些 commits 可能已經被其他人拉回本機繼續修改，而當他們改完要推送變更至遠端 rep 時會被拒絕，所機魚的分支已經不存在了。
{{< /admonition >}}

See also:

- [GitLab Docs > Protected branches](https://docs.gitlab.com/ee/user/project/protected_branches.html)
- [GitLab Docs > Protected branches](https://docs.github.com/en/repositories/configuring-branches-and-merges-in-your-repository/managing-protected-branches/about-protected-branches)
