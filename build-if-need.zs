#!/usr/bin/zhscript --。
（
Progrem New_Builder 项目名
Up
Location:
${workspace_loc:/项目名/build-if-need.zs}
Arguments:
主类所在的文件路径，不必后缀，可多个
）
赋予gopath以/opt2/go。

加载lib/doscmd4。
如果调用‘existdir’、‘gopath’等于0那么先
	执行“echo "‘gopath’ 不存在"”。
	结束250。
了。

加载lib/defi。
加载lib/file4。
赋予根以调用‘得目录名’、先调用‘最终文件名’、‘参数0’了。
赋予目录以‘根’src。
加载lib/stdc。
调用‘改变目录函数’、‘目录’。

赋予其他以“export GOPATH=‘gopath’:‘根’”换行。

加载lib/gjk4。
赋予搜以调用‘dir.begin’、.、\.(go|so)$、s。
赋予最大时以0。
循环先
	赋予名以调用‘dir.next’、‘搜’。
	如果不‘名’那么跳出。
	如果调用‘尾匹配方法’、‘名’、.so那么先
		赋予其他以‘其他’“export LD_LIBRARY_PATH=”‘目录’/先调用‘得目录名’、‘名’了:$LD_LIBRARY_PATH换行。
		再来。
	了。
	赋予时以调用‘fileinfo’、‘名’、%tm。
	（执行“echo "‘时’ ‘名’"”。）
	如果‘时’大于‘最大时’那么
		赋予最大时以‘时’。
了。
调用‘dir.end’、‘搜’。

（赋予其他以‘其他’env换行。）

加载lib/clpars4。
调用‘命令行加回调’、
	-b、、1、下代码
		赋予‘参数0’【上】以‘参数1’。
	上代码、
	、、1、下代码
		当先
			赋予主以‘参数1’。
			调用‘exist’、‘主’那么跳出。
			赋予主以‘参数1’.go。
			调用‘exist’、‘主’那么跳出。
			执行“echo "‘主’ 不存在"”。
			结束250。
		了。
		赋予出以../bin/如果存在-b【上】那么‘-b’否则调用‘得文件主名’、‘主’。
		赋予出时以调用‘fileinfo’、‘出’、%tm。
		执行“echo "‘出时’ ‘出’"”。
		如果‘出时’大于‘最大时’那么先
			执行“echo "嘿"”。
			再来。
		了。
		
		执行“echo "哼"”。
		赋予错以执行下原样
export PATH=/opt/local/bin/:/usr/local/go/bin/:$PATH
export PKG_CONFIG_PATH=/opt/local/bin
上原样‘其他’下原样
go build -gcflags "-N -l" -o 上原样“‘出’ ‘主’”下原样
上原样。
		如果‘错’那么先
			执行“echo "啊‘错’"”。
			结束‘错’。
		了。
		执行“echo "哈"”。
	上代码。
调用‘命令行解析’、‘参数栈’。