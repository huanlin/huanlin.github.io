---
title: "[Hugo][Multilingual] - Fall back to the default language"
tags: ["Hugo", "Docsy"]
---

See discussion: [How to automatically fallback to the default language when a requested page doesn't have a translation?](https://github.com/google/docsy/discussions/1998)

Here is how I set it up for my website:

```toml
[module]
  [[module.mounts]]
    source = "content/en"
    target = "content"

  [[module.mounts]]
    source = "content/zh"
    target = "content"
    lang = "zh"

  [[module.mounts]]
    source = "content/en/docs" # zh missing-page fallback
    target = "content/docs"
    lang = "zh"

  [[module.mounts]]
    source = "content/en/blog" # zh missing-page fallback
    target = "content/blog"
    lang = "zh"
```

Hugo and Docsy versions:

- Hugo v0.125.6
- Docsy v0.10.0

## References

- [Do all page bundles need localized copies once you add a new language?](https://discourse.gohugo.io/t/do-all-page-bundles-need-localized-copies-once-you-add-a-new-language/37225)