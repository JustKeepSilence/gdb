# RESTFUL API
```
对于一般的开发者,我们提供了基于Restful的接口。要想使用GDB,首先需要创建组,然后向组中添加Item,之后再写入实时值
所以对应的接口可以分为三部分,分别是Group(和组相关),Item(和Item相关),Data(和数据相关)。
1.请求的方式:POST
2.返回的数据格式
{"Code": , "Message": , "data": },分别代表返回的状态码,返回的信息,返回的数据。请求成功Code为200,Message
为空,data为对应的数据。请求失败Code为500，Message为错误信息,data为null
```
<h2>1.Config</h2>
```
在使用整个数据库之前需要先对数据库进行配置,其中配置文件的名称为config.json,配置的格式是json文件
1.path: 用于指定整个数据库的路径,默认为当前路径下的db文件夹中
2.port: Resuful接口和网页程序的IP端口号,默认为9000
3.ip: 整个服务的ip地址
```
<h2>2.Notes</h2>
```
默认的web服务有两个端口号,一个是由用户配置的,另一个为8087,用于加载用户文档
```
<h2>3.Start</h2>
```
配置完成之后启动应用程序db.exe,开启你的GDB之旅吧~~~
```
<i class="fa fa-github" style="margin-left: 250px;width: 30px"></i>https://github.com/JustKeepSilence
