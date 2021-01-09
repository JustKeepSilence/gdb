# GDB
GDB是基于Leveldb使用Go语言编写的高性能实时数据库,可以实现对大量数据实时值的存储和历史值的读取.并且提供了restful api以及
网页客户端来操作数据库,同时还提供了基于JS的二次计算开发功能，可以允许用户使用js语言来进行已有数据的二次数据开发工作.同时对于
Go语言使用者,还可以直接下载src中的源代码进行二次开发工作.

# Quick Start
下载编译后的代码文件夹bin,其中config为配置文件夹,其中的config.json中为整个数据库的基本配置文件,ip指定了整个web服务的ip
地址,如果为空则会默认读取本机的ip，port指定了整个web服务的端口号,默认值为8082, path指定了整个实时数据库的存储位置.你可以
保持设置为默认值,然后启动db.exe,此时结果应该如下图所示:
![Image text](https://github.com/JustKeepSilence/gdb/blob/master/images/launch.png)

# Web Application Serve
在支持es6的浏览器上输入对应的url: http:// + ip + port + "/index",初次运行没有cookie信息,会跳转到如下的登陆界面
