# 介绍
```
GDB数据库是基于LevelDb的使用Go语言编写的实时数据库,可以实时添加实时数据,并获取数据的历史,适用于需要大量
存储数据的过程例如热工过程,实时网页应用等.
整个数据库的结构分为组Group和点Item,组就相当于SQL中的table,Item相当于每个table中的数据。使用之前需
要先添加对应的Group,之后再向Group中添加Item。有了Item之后就可以向Item写入实时数据,写入实时数据之后
就可以获取对应的历史数据。
GDB的特点:
1.基于LevelDB,批量写的速度很快,支持成千上万个连接同时并发写
2.使用bloomFilter,批量读的速度也很快
3.提供了高性能,可并发的webRestful接口,提供了网页程序和客户端程序,对于go语言使用者提供了
底层的函数,使用方便,可以和各个语言无缝衔接。
3.适合于需要大规模存储数据的场合,如电厂热工过程数据的存储,实时网页相关数据的存储
其中Group中的数据存储在SQLite中,实时数据和历史数据存储在Leveldb中.整个数据库的结构如下:
4.提供了网页应用程序,可以方便的对GDB数据库进行相应的操作
```
<img src="./images/db.png" style="width: 100%"/>
<i class="fa fa-github" style="margin-left: 250px;width: 30px"></i>https://github.com/JustKeepSilence