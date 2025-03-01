---
title: GitLab pages for different branches
date: "2024-07-08"
---

## Attempt 1

Stack Overflow: [Deploying GitLab pages for different branches](https://stackoverflow.com/questions/55596789/deploying-gitlab-pages-for-different-branches)

It doesn't work because each page's URL must be ended with "index.html," which will not work with all links in the website except the home page.

## Attempt 2

[GitLab Pages per Branch: The No-Compromise Hack to Serve Preview Pages](https://dev.to/zenika/gitlab-pages-preview-the-no-compromise-hack-to-serve-per-branch-pages-5599)

I couldn't get it work. A comment of [a Stackoverflow post](https://stackoverflow.com/questions/75853041/gitlab-pipeline-deploy-subfolders-to-gitlab-pages) explained:

## Something really useful

After spending several hours trying, I found the following posts on GitLab that are useful:

- [Allow customizing the artifacts path ("public") in GitLab Pages](https://gitlab.com/groups/gitlab-org/-/epics/10126)
- [Automatically use the "publish" property in the pages job as an artifact](https://gitlab.com/gitlab-org/gitlab/-/issues/398145)

## Conclusion

At the time of writing, GitLab Pages doesn't support deploying multiple branches.
