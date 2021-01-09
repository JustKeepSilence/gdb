# GDB
```
GDB是基于Leveldb使用Go语言编写的高性能实时数据库,可以实现对大量数据实时值的存储和历史值的读取.并且提供了restful api以及
网页客户端来操作数据库,同时还提供了基于JS的二次计算开发功能，可以允许用户使用js语言来进行已有数据的二次数据开发工作.同时对于
Go语言使用者,还可以直接下载src中的源代码进行二次开发工作.
整个GDB的数据结构分为grou和item,使用之前必须先添加group,再向group中添加item,之后就可以向item中写入实时数据
```

# Quick Start
```
下载编译后的代码文件夹bin,其中config为配置文件夹,其中的config.json中为整个数据库的基本配置文件,ip指定了整个web服务的ip
地址,如果为空则会默认读取本机的ip，port指定了整个web服务的端口号,默认值为8082, path指定了整个实时数据库的存储位置.你可以
保持设置为默认值,然后启动db.exe,此时结果应该如下图所示:
```
![Image text](https://github.com/JustKeepSilence/gdb/blob/master/images/launch.png)



# Web Application Serve
1.login

```在支持es6的浏览器上输入对应的url: http:// + ip + port + "/index",初次运行没有cookie信息,会跳转到如下的登陆界面```
![Image text](https://github.com/JustKeepSilence/gdb/blob/master/images/login.png)

```其中远端服务器的地址即为整个服务运行的ip+port,这里就是192.168.0.114:8082,用户名为admin,密码为admin@123```

2.index

```首页的界面如下```
![Image text](https://github.com/JustKeepSilence/gdb/blob/master/images/index.png)

```分别显示了当前程序的内存使用情况，当前写入的item数量，最近一次写入的unix时间戳以及写入速率```


3.group

```group的界面如下```
![Image text](https://github.com/JustKeepSilence/gdb/blob/master/images/group.png)

```通过这个界面可以加组,加点,下载点表,查看点的实时值和历史值，下载指定时间的历史数据,编辑item的属性等一系列操作```

4.calc

```calc是和二次计算操作有关的界面```
![Image text](https://github.com/JustKeepSilence/gdb/blob/master/images/calc.png)

```通过这个界面可以添加基于二次js的二次计算项,也可以实时修改，监测计算项的运行情况.```

# Restful Api
[接口文档](https://justkeepsilence.gitbook.io/gdb/)

# gRPC
正在开发中...