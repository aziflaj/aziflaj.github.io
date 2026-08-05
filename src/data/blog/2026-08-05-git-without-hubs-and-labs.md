---
title: "Git without the hubs and the labs"
pubDatetime: 2026-08-05
description: "How to share code with Git patches, without GitHub, GitLab, or any other hub, lab, or bucket."
slug: git-without-hubs-and-labs
tags: [git, version-control, software-development]
---

Git has always been advertised as a decentralized version control system. In simple terms, that means you don't need a server, whatever a server is. And yet everyone and their mother uses some server, like GitHub or GitLab, or Bitbucket even, where you host your code after every push and can share it with everyone else who might be interested in whatever slop you clearly wrote without an LLM.

What I'm about to show you is an old-timer, grandpa-style code-sharing technique that removes the need for Hubs, Labs, Buckets and whatever we collectively decided to call "Git servers".

First, a small realignment of concepts is due. When you work on a git-controlled codebase, you don't really care about code. You care about **changes to code**, what we call **patches**. A patch, or diff, describes which files changed, which lines were added or removed, and what text replaced what. It looks something like this, taken from my previous Rust Ownership Shenanigans post:

```rs
diff --git a/file.rs b/file.rs
index 84962c8..f7f646b 100644
--- a/file.rs
+++ b/file.rs
@@ -3,11 +3,19 @@
 #[derive(Debug)]
 struct Coord(f64, f64);

+fn relocate(c: &mut Coord) {
+  println!("Relocating from {:?} to Null Island", c);
+
+  *c = Coord(0f64, 0f64);
+}
+
 fn main() {
   let your_house = Coord(45.463504, 6.576340);
-  let your_mums_house = your_house;
+  let mut your_mums_house = your_house;
+
+  println!("Before: {:?}", your_mums_house);

-  // println!("{:?}", your_house); // -> this won't compile
+  relocate(&mut your_mums_house);

-  println!("{:?}", your_mums_house);
+  println!("After: {:?}", your_mums_house);
 }
```

When you create a pull/merge request, your hub/lab/bucket takes this diff, puts some fancy CSS around it, and adds a few features like _"comment on this code block"_ and _"click this button to merge branches"_. That comment thing is nice to have, and the button is just some git commands hidden behind a click.

For git, this diff is enough. It doesn't really need the rest. First, to generate this diff, you can use the following command:

```bash
git diff > changes.patch
```

Nothing fancy, just put the output of `git diff` in a file called `changes.patch`. All needed for a PR/MR is in this file. So now we can print this `.patch` file, fold it carefully so as not to ruin the text, attach it to an [IP over Avian Carriers (RFC 1149)](https://www.rfc-editor.org/info/rfc1149/) packet and send it to some other developer.

When the avian packet arrives, assuming the pigeon didn't get eaten by a cat or get promoted to a PM in the meantime, that other developer can apply the patch to their codebase like this:

```bash
git apply changes.patch
```

And that's it. No servers, no hubs, no labs. Just you, your code, and a pigeon. Or, if you're a bit more modern, a drone. Or you can send the patch via email; that's how Linux kernel development is done. All patches and emails.

But email uses servers, and I said "no servers".
