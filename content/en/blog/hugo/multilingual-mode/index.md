
In Hugo configuration file:

```
  [[module.mounts]]
    source = 'content/en'
    target = 'content'
    lang = 'en'
  [[module.mounts]]
    source = 'content/en' # Use content/en as the fallback folder to find pages when the selected languae is 'zh'
    target = 'content/'
    lang = 'zh'
```    
