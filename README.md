# LeetCode
LeetCode practice

题目来自：[leetcode-cn](https://leetcode-cn.com/problemset/all/)

* 说明
    - 数据结构和算法的部分代码练习也放在该项目里
        + 学习笔记链接：[数据结构与算法.md](https://github.com/xiaodongQ/devNoteBackup/blob/master/%E5%90%84%E8%AF%AD%E8%A8%80%E8%AE%B0%E5%BD%95/%E6%95%B0%E6%8D%AE%E7%BB%93%E6%9E%84%E4%B8%8E%E7%AE%97%E6%B3%95.md)
    - 一开始没有用VSCode插件，用插件后做题记录统一放：`./lc/`目录
    - 目前都用Go做的题，偶尔需要调试时单元测试比较方便，后续考虑C++

## VSCode LeetCode 插件

* 插件安装后很方便，能自动生成文件，还有测试和提交的快捷按钮，还能看到根据难易、各类标签排序的题目列表
* 插件安装和相关问题，记录在笔记里(`LeetCode插件配置`章节)
    - 笔记链接：[vscode快捷键及插件.md](https://github.com/xiaodongQ/devNoteBackup/blob/master/%E5%B7%A5%E5%85%B7%E4%BD%BF%E7%94%A8/vscode%E5%BF%AB%E6%8D%B7%E9%94%AE%E5%8F%8A%E6%8F%92%E4%BB%B6.md)

---

* 截取笔记内容：LeetCode插件配置
    - 在UI配置界面里搜leetcode
        + 各项配置参考：[settings](https://swift.ctolib.com/jdneo-vscode-leetcode.html#settings)
    - `Endpoint`选择中文站或者国际站(我用的中文站)
    - `Show Description`设置题目描述展示的方式，我选`Both`，生成的文件和VSCode里都会显示题目信息
    - `Workspace Folder`设置生成文件的路径
    - `File Path` 设置生成文件的格式，由于有时自己像运行些调试测试，所以配置生成的是`_test.go`形式的文件，便于单元测试
        + 设置File Path选项可以参考官网说明：[Customize the Relative Folder and the File Name](https://github.com/LeetCode-OpenSource/vscode-leetcode/wiki/Customize-the-Relative-Folder-and-the-File-Name-of-the-Problem-File)
    - 设置换行为`\n`，插件(不确定是gopls还是leetcode插件)只支持这类文件的自动调整(不过我windows上自动生成的文件还是`\r\n`。。。)
    - 不同机器用同一个账号，题目没有自动同步，自己的方式是切换一下站点再切回来(会自动退出原有登录)，重新登录

```json
"leetcode.endpoint": "leetcode-cn",
"leetcode.hint.configWebviewMarkdown": false,
"leetcode.workspaceFolder": "/home/xxx/lc",
"leetcode.defaultLanguage": "golang",
"leetcode.filePath": {
  "default": {
    "folder": "",
    "filename": "${id}.${kebab-case-name}_test.${ext}"
  }
},
"leetcode.hint.commandShortcut": false,
"leetcode.editor.shortcuts": [
  "submit",
  "test",
  "solution"
],
"leetcode.showDescription": "Both",
"leetcode.hint.commentDescription": false,
```