## Introduction
Just another translate tool, for fit my self or others, I hope you like.

The name von comes from **John von Neumann** **:)**

## Install:
Make sure your golang version **>= 1.23.0**, setup the golang path and run:
```bash
go install github.com/qedhtp/von@latest
```

## Usages
#### Non-interactive mode

```bash
$ von apple

Phonetic symbol:
    UK: / ˈæp(ə)l /    US: / ˈæp(ə)l /    

Interpretation:
    n.苹果

Phrase:
    1.apple inc  苹果公司 ; 美国苹果公司 ; 苹果
    2.BIG APPLE  大苹果 ; 纽约 ; 大苹果城
    3.Apple Computer  苹果电脑公司 ; 苹果计算机 ; 苹果计算机公司

Examples:
    1.She crunched her apple noisily.
      她吃苹果发出嘎嚓嘎嚓的声音。
    2.He took another bite of apple.
      他又咬了一口苹果。
    3.Someone threw an apple core.
      有人扔了一个苹果核。
```
```
$ von "This is an example"
这是一个例子
```
#### Interactive mode

```bash
$ von -i

[von]>>> apple

Phonetic symbol:
    UK: / ˈæp(ə)l /    US: / ˈæp(ə)l /    

Interpretation:
    n.苹果

Phrase:
    1.apple inc  苹果公司 ; 美国苹果公司 ; 苹果
    2.BIG APPLE  大苹果 ; 纽约 ; 大苹果城
    3.Apple Computer  苹果电脑公司 ; 苹果计算机 ; 苹果计算机公司

Examples:
    1.She crunched her apple noisily.
      她吃苹果发出嘎嚓嘎嚓的声音。
    2.He took another bite of apple.
      他又咬了一口苹果。
    3.Someone threw an apple core.
      有人扔了一个苹果核。
```
```
[von]>>> this is an example
    这是一个例子
```
```
:clear   # clear screen
:exit    # exit von shell
```


