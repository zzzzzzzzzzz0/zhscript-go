实现一种不像计算机语言的脚本语言。按照中文语法习惯去分析执行语句。

中文脚本 zhscript golang 版，仅解释器。较之[ c++ 版](https://github.com/zzzzzzzzzzz0/zhscript)采用了 codecache 机制，使所有代码段只解析一次。

以下是相关项目：

[l-go](https://github.com/zzzzzzzzzzz0/l-go) 解释器命令行。（goclipse 工程）

[zsp-go](https://github.com/zzzzzzzzzzz0/zsp-go) 中文语法的动态网页。服务器端直接解释，不似 jsp、php 般再编译。（goclipse 工程）