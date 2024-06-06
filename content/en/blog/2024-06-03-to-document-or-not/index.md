---
title: To document or not?
slug: "to-document-or-not"
date: 2024-06-03
tags: ["Technical Writing"]
---

> This is a translation by AI. **如欲閱讀中文版，請使用頁面上方的語言切換器。**

After working as a full-time technical writer for almost a year, it seems like an appropriate time to summarize my experiences and observations. On the other hand, I recently gave two presentations on document production to internal teams using online live streaming, and I collected some questions and feedback worth organizing.

So in this note, I will discuss my views on "writing documents" in the form of Q&A and share some insights. I have compiled five questions this time:

- Do we need to write documents? Are documents useful?
- When should we consider creating "formal documents"?
- Is the Docs-as-Code method suitable for me?
- What do you find most bothersome about writing documents?
- Do you use AI to write documents? Will AI replace human writers?

## Do we need to write documents? Are documents useful? {#to-write-or-not}

There is a saying: "No one likes to write documents."

The author of the book "Living Documentation" also admitted in the book: "Updating documents is one of the least valued tasks. It is not fun and seems to have little return."

Some people have asked me in a more polite way: "Do you know who reads the documents you write? It seems like it's faster to just ask someone directly."

Others are more direct: "No one probably reads the documents you write."

These doubts or opinions about the usefulness of documents seem normal to me. When I unpack consumer electronics, I rarely read the product manuals. However, some functions or usage of certain appliances may be a bit more complicated. In those cases, I am glad to find the product manual. For example, I recently bought a vacuum cleaner that can remove dust mites. After using it for a while, I realized that I needed to clean the filter, but I couldn't figure out how to remove it. It wasn't until I read the manual that I realized I needed to use one of the narrower suction heads to disassemble it.

So, are documents useful? Will anyone read them?

When you need them, they are useful; just because you don't need them doesn't mean others don't. 

It's a bit like insurance - when you don't need it, you may feel that paying insurance premiums every year is unnecessary expenses; when you need it, you will be glad that you had insurance (maybe even feel that you didn't have enough insurance).

Therefore, if you are someone who writes documents, don't worry about what others say about the usefulness of documents. Especially when you are a full-time technical writer, writing documents is your job, something you should do; it is important to think about how to write good documents that are useful to readers.

If you are not a technical writer? Well, if you feel like writing, then write. Don't care about what people say. Writing itself is the best reward, and I believe those who have the habit of writing can understand the meaning of this sentence.

So, just write if you want to. Instead of worrying about whether to write or not to write, or whether documents are useful, it is better to think about what content you should write and how to write it well.

If you still think, "What you said makes some sense, but can you provide more specific or more compelling evidence?"

The following documents may serve as reference examples:

- [AWS Documentation](https://docs.aws.amazon.com/) and its [source code](https://github.com/awsdocs).
- [Microsoft Learning](http://www.microsoft.com/learning) and its [source code](https://github.com/MicrosoftLearning).
- [Kubernetes Documentation](https://kubernetes.io/docs/) and its [source code](https://github.com/kubernetes/kubernetes).

The technical documents listed above, in my opinion, embody the spirit of **documentation as a product** (referred to as docs as a product), which means conceptualizing, developing, and maintaining documents as formal (or official) product documents. Their source code is written in markdown and can be considered as documents produced using the **Docs-as-Code** method.

I listed the above reference documents because I am concerned that readers might mistakenly think that the documents discussed here also include some scattered or casual development notes. So, I want to emphasize that what we are going to discuss next is focused on "formal documents," such as product user guides, tutorials, API reference manuals, and so on.

## When should we consider creating formal documents? {#do-i-need-docs-as-a-product}

I don't want to pretend that I know everything. The need to spend time creating formal product documents depends on whether the **target users** of the product require them, and this varies for each company, software project, and development team.

For example, for some applications, the target users may be a few internal developers, and everyone can understand by looking at the code and its comments. In such cases, I think there is usually no need to spend additional time and effort creating formal documents, and even discussing "whether to have documents" seems unnecessary.

So, what if it is an open-source project with a larger number of target users (perhaps tens or hundreds of people)? And all users can access the source code and program comments, is there still a need to create formal product documents? This question may come with an additional note: "My code is self-explained, easy to understand at a glance."

OK, here's how I see it:

- If the code is really simple and easy to understand, and you believe that **anyone** should be able to understand it just by looking at the code and can solve problems, then perhaps there is no need to go through the trouble of creating formal documents; writing a readme and FAQ should be sufficient. Of course, this also depends on the scale and complexity of the application itself.
- Ask yourself: "If there is already an English version, do we still need a Chinese translation?" Although not exactly the same level of question, the essence is similar: <mark>Can it save users a considerable amount of time in learning, reading, and searching?</mark> If you believe that a Chinese translation would help accelerate learning, how can you deny the help of product documents to users?
- The "ask twice" principle. If the same question is asked twice, it usually reflects the users' dilemma and should be considered as a sign that it needs to be written in the documents.
- Take a look at the [Kubernetes documentation](https://kubernetes.io/docs/) and imagine if the K8s development team claimed, "Our code is easy to read and understand, and we have timely comments. It is all publicly available on GitHub, so there is no need to write product manuals and API tutorials."

Now let's talk about the Docs-as-Code approach.

## Is the Docs-as-Code method suitable for me? {#docs-as-code}

This question is not easy to answer. Based on the opinions I have encountered so far, people who are not accepting of it are mostly accustomed to the "what you see is what you get" editing style (such as Word, Confluence), so it is difficult for them to accept writing documents in markdown.

I'm afraid I can't give good advice on this point. Because I like to use markdown to write anything, the Docs-as-Code approach, which involves writing in markdown and using a git workflow to automatically build and publish a documentation website, resonates with me.

So my opinion is definitely biased towards the Docs-as-Code method.

However, if certain existing document management platforms support markdown, I think it would be a viable option. For example, Microsoft Azure DevOps Wiki, Document 360, and so on. This is because considering the possibility of migrating documents between different platforms in the future, it would be troublesome if the documents are locked to a specific platform (vendor lock-in) and cannot be easily transferred.

## What do you find most bothersome about writing documents? {#most-bothersome}

Based on my current experience, the most bothersome thing is when I need to write documents about a technology or product that I don't understand well, and it is not easy to gather reference materials.

Especially when it comes to architecture, design decisions, and some historical background factors, if the development team did not leave any records and updated them as the product evolved, it becomes difficult to dig out these design decisions when suddenly realizing the need to write product documents.

## Do you use AI to write documents? Will AI replace human writers? {#ai}

I often use AI assistants when writing documents. Sometimes I use them for translation, sometimes for polishing words, correcting spelling and grammar, and sometimes I even ask them to generate a piece of text or API examples.

As long as AI can help me improve efficiency, I will use it more frequently. As for whether it will replace my work, I'm not worried.

When AI can write better than me, I will use it to write. If I can use AI to write more and better, or have more leisure time, why not?

> However, can AI really write documents for "new products"? What kind of document materials can AI provide for a completely new development? You may know more about AI than I do, so please leave a comment and let me know.

## Conclusion

- Whether to create formal product documents depends on whether the target users of the product need them and should not be generalized.
- Documents are needed by those who need them; those who don't need them don't. XD (just kidding)
- Preserving design decisions is important for writing documents.
- Make good use of AI to increase productivity without fear of being replaced.

Finally, I would like to quote the following sentence from the book "Living Documentation":

> "Updating documents is one of the least valued tasks. It is not fun and seems to have little return. **But if you take it seriously and decide to adopt appropriate mechanisms to ensure its correctness, you can write good documents.**"

Did you know? After writing useful documents for users, not only will you be happy, but the rewards will gradually appear - this is based on my personal experience.

Keep writing!
