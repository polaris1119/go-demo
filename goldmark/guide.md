这是一篇讲解如何正确使用 **Markdown** 的排版示例，学会这个很有必要，能让你的帖子有更加清晰的排版。

> 引用文本：Markdown is a text formatting syntax inspired

## 语法指导

### 普通内容

这段内容展示了在内容里面一些小的格式，比如：

- **加粗** - `**加粗**`
- *倾斜* - `*倾斜*`
- ~~删除线~~ - `~~删除线~~`
- `Code 标记` - ``Code 标记``
- [超级链接](http://github.com) - `[超级链接](http://github.com)`
- [username@gmail.com](mailto:username@gmail.com) - `[username@gmail.com](mailto:username@gmail.com)`
- 自动链接 - https://studygolang.com

### 提及用户

@polaris ... 通过 @ 可以在发帖和回帖里面提及用户，信息提交以后，被提及的用户将会收到系统通知。以便让他来关注这个帖子或回帖。

### 表情符号 Emoji

支持表情符号，你可以用系统默认的 Emoji 符号（无法支持 Windows 用户）。
也可以用图片的表情，输入 `:` 将会出现智能提示。

#### 一些表情例子

:smile: :laughing: :dizzy_face: :sob: :cold_sweat: :sweat_smile:  :cry: :triumph: :heart_eyes: :relaxed: :sunglasses: :weary:

:+1: :-1: :100: :clap: :bell: :gift: :question: :bomb: :heart: :coffee: :cyclone: :bow: :kiss: :pray: :sweat_drops: :hankey: :exclamation: :anger:

### 大标题 - Heading 3

你可以选择使用 H2 至 H6，使用 ##(N) 打头，H1 不能使用，会自动转换成 H2。

> NOTE: 别忘了 # 后面需要有空格！

#### Heading 4

##### Heading 5

###### Heading 6

### 图片

```
![alt 文本](http://image-path.png)
![alt 文本](http://image-path.png "图片 Title 值")
```

### 代码块

#### 普通

```
*emphasize*    **strong**
_emphasize_    __strong__
@a = 1
```

#### 语法高亮支持

如果在 ``` 后面跟随语言名称，可以有语法高亮的效果哦，比如:

##### 演示 Go 代码高亮

```go
package main

import (
    "fmt"
)

func main() {
    fmt.Println("Hello World!")
}
```

##### 演示 JSON 代码高亮

```json
{"name":"Go语言中文网","url":"https://studygolang.com"}
``` 

> Tip: 语言名称支持大部分常用的语言

### 有序、无序列表

#### 无序列表

- Go
  - Gofmt
  - Revel
  - Gin
  - Echo
- PHP
  - Laravel
  - ThinkPHP

#### 有序列表

1. Go
    1. Gofmt
    2. Revel
    3. Gin
    4. Echo
2. PHP
    1. Laravel
    2. ThinkPHP
3. Java

### 表格

如果需要展示数据什么的，可以选择使用表格哦

| header 1 | header 3 |
| -------- | -------- |
| cell 1   | cell 2   |
| cell 3   | cell 4   |
| cell 5   | cell 6   |

### 段落

留空白的换行，将会被自动转换成一个段落，会有一定的段落间距，便于阅读。

请注意后面 Markdown 源代码的换行留空情况。

### 任务列表

- [x] 完成 Markdown 教程
- [] 完成简书项目
