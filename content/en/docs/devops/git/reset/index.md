---
title: Git reset
---

## Step 1

```shell
git reset -- hard
```

## Step 2

Force push to the remote Git server:

```shell
git push -f
```

The push operation will fail if the branch is protected, and you may see an error message like below:

```console
remote: GitLab: You are not allowed to force push code to a protected branch on this project.
```

Open the project's page on GitLab, go to Settings > Repository > Protected branches, then find the branch that is protected. You can temporarily `Unprotect` the branch or enable the `Allowed to force push` option of the branch.

After doing a forced push successfuly, remember to protect the branch again.

