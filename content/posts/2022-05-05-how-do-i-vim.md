---
layout: post
title: How do I Vim
date: '2022-05-05'
---

My first editor was Notepad. Not [the ++
version](https://notepad-plus-plus.org/), the plain old white one that lacks syntax
highlighting, smart indentation, and
basically everything you need to write code. All it did was help me
create files with different extensions. That's not a big deal when all you're
writing are Windows Batch scripts, but it does get harder when you go beyond the
script kiddie stage. Since then, I have switched through a handful of code
editors and used them for a while before deciding they were garbage and moving
on to the next one.

For C/C++ I have used [Turbo
C++](https://en.wikipedia.org/wiki/Turbo_C%2B%2B), Notepad++, [Qt](https://www.qt.io/), and finally 
[Sublime Text](https://www.sublimetext.com/) and [gcc](https://gcc.gnu.org/) instead of an IDE. For Java, I've used [NetBeans](https://netbeans.apache.org/) and [Eclipse](https://www.eclipse.org/ide/)
before settling on [IntelliJ IDEA](https://www.jetbrains.com/idea/). I tried [Eclipse ADT](https://www.eclipse.org/downloads/packages/release/neon/m6/eclipse-android-developers) before switching to
[Android Studio](https://developer.android.com/studio). For Ruby, which is what I've been doing mostly for these past 6 years, I
started with [Atom](https://atom.io/), moved to [RubyMine](https://www.jetbrains.com/ruby/), then gave [VS Code](https://code.visualstudio.com/) a shot before moving to
Vim. Now, the only editor I use is [Neovim](https://neovim.io/), only
because it has [Github Copilot](https://copilot.github.com/) support. Besides that, Vim and Neovim are basically the same to me.

People who never used Vim before think it's either too hard to use,
or they think it's just a Notepad that's too hard to get out of, so not worth
using. But there's a reason your code editor and IDE of choice comes with a Vim mode,
so without further ado, allow me to introduce you to the Vim goodness.

{{< image src="/images/20220502/vimvside.jpeg" alt="chadvim" position="center" style="border-radius: 8px;" >}}

## Modal editing

When you first start it, you get what's known as **Normal mode**. This is where
your keys don't actually do what they're expected to do, e.g.:
- pressing `h`, `j`, `k`, or `l` will move the cursor around rather than writing the character ([for historical reasons](https://catonmat.net/why-vim-uses-hjkl-as-arrow-keys))
- pressing `w` or `b` also moves the cursor around, but in a different special way
- pressing `x` deletes characters
- pressing `dd` ~~deletes~~ crops lines
- pressing `yy` appears to do nothing, but pressing `p` after that duplicates
  the line (kinda like ~~copying~~ yanking and pasting)
- pressing `i` or `a` makes the keys start working as they should...

The reason why Vim starts in a Normal mode that doesn't seem normal at all (when
you switch from _normal_ code editors) is that most of the time, you are reading
the file and moving around rather than making changes. By pressing `i`, you
switch modes from **Normal** to **Insert**. This means that
now, instead of changing (editing) the content, you're inserting new stuff in
the file. Pressing `ESC` gets you back to Normal mode.

{{< image src="/images/20220502/replace.gif" alt="replace" position="center" style="border-radius: 8px;" >}}

In the gif you can see how replacing is done in Vim. Writing `:%s/vim/Vim/g` while on Normal mode changes
all `vim` occurrences into `Vim`; those of you with
[sed](https://www.gnu.org/software/sed/) knowledge will notice the similarity.
You can also notice (on the bottom left, right above where the command is
written) that Vim changes from **Normal** to **Command** mode, everything that
starts with `:` is a Vim command.

One of the reasons Vim seems hard to use at first is that people think they need
to know everything about it in order to start using it. Long before making it my
default code editor, I was using Vim to edit files in production servers. All I
needed for that was to know how to move around with `h/j/k/l`, `w/b` for moving
between words, `dd` to delete lines and `i` to start writing what I needed to
write. That's more than enough to ditch nano or whatever terminal editor you're
currently using.

There's also Visual mode (including Visual Block and Visual Line), similar to
click-dragging your mouse over the text; and Replace mode, similar to pressing
the `Ins` key on [most keyboards](https://www.computerhope.com/jargon/i/insertke.htm), but I rarely use them, so ... not much to say about them

¯\\_(ツ)_/¯

## Vim is hackable like Atom

...well, it's a bit more hackable than that. The default Vim you get in your OS
is just the barebones; it does almost what Notepad can do. But with a few lines
of `init.vim` file and a couple of plugins, you can turn it into a real IDE that
supports:

- Syntax highlighting
- Line numbers and _relative_ line numbers
- Indentation guides, visible invisible characters (like spaces, tabs, line
  breaks, or [this invisible cunt](https://www.youtube.com/shorts/x1kyIUZgzqo))
- Running tests by doing `Space + t` (current **T**est file), `Space + a` (**A**ll
  tests), `Space + s` (the test **S**urrounding the cursor), or `Space + l` (rerun
  the **L**ast
  test)
- Jumping to the test file for a class/module by pressing `:A`
- Automatically running a linter whenever a file is saved
- `Ctrl + f` to globally search for a given string (like `Ctrl + Shift + f` on
  most IDEs/editors)
- `Ctrl + p` to open a file
- Code completion suggestions as in every other IDE or Code editor.

{{< image src="/images/20220502/completion.gif" alt="code completion" position="center" style="border-radius: 8px;" >}}

The hackable thing about Vim is that the above commands are not set in stone.
Here are a couple of lines from my `init.vim`:

```vim
" Testing made easy
map <Leader>t :TestFile<CR>
map <Leader>s :TestNearest<CR>
map <Leader>l :TestLast<CR>
map <Leader>a :TestSuite<CR>
```

These are defining what `Space + t/s/l/a` (mentioned above) do. Think of
`<Leader>` as the `Space` key, and whenever `Space + t` are pressed, Vim acts
the same as typing `:TestFile` followed by Enter (`<CR>` stands for [Carriage Return](https://en.wikipedia.org/wiki/Carriage_return#Computers)). Same thing goes for my `Ctrl + p` and `Ctrl + f`; both of these are
user-defined, so they can do whatever the user (in this case, myself) wants:

```vim
nnoremap <C-p> :Telescope find_files<CR>
nnoremap <C-f> :Telescope live_grep<CR>
```


## I don't have to leave my terminal

In my job, I get to use the terminal a lot. After so many years of struggling with
computers, I'm more comfortable doing
stuff from the terminal rather than the GUI (that's fancy talk for _"I'm too
lazy to reach for the mouse"_). So when the first thing you do on a Monday
morning is to fire up [iTerm](https://iterm2.com/index.html), using a
terminal-based editor doesn't seem that crazy. _"But wait"_, -- you say, -- _"even
my Fancy Schmancy IDE has a built-in terminal"_. It sure does, but mine is
floating in the middle of the screen.

{{< image src="/images/20220502/floaterm.png" alt="floaterm" position="center" style="border-radius: 8px;" >}}

Plus I can fire up a new terminal window by doing `⌘ + T` and move around with
`⌘ + number` or `⌘ + left` and `⌘ + right`; using the keyboard is faster than
reaching for the mouse.

## Moving through panes

Here how a normal day of work might look [like](https://github.com/aziflaj/gogot):

{{< image src="/images/20220502/workdir.png" alt="gogot" position="center" style="border-radius: 8px;" >}}

There might be a few files opened side-by-side, a terminal window open to run some
commands, and so on. Moving between opened files is as easy as `Ctrl + h/j/k/l`
(on my config), toggling the terminal is just a `\ + e` away and as Drew Neil
says, [_text is edited at the speed of thought_](https://www.amazon.com/Practical-Vim-Edit-Speed-Thought/dp/1680501275).

FYI, to get out of Vim you do `:h :q` and RTFM.
