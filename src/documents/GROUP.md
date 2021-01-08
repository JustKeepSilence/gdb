# Group
```
组Group是整个GDB数据库的基础,和其相关的接口主要是用于向GDB数据库中添加组,这里的response只给出了正确
响应时data字段的内容,egs给出了对应的python代码示例
```
<h2>1.Notes</h2>
```
组的操作不支持并发写操作,因为SQLite是RwMutex,所以支持并发的读,并发写的时候会造成数据库被锁住
```
<h2>2.Api</h2>
```
(1) url: /group/addGroups
    des: 向GDB中添加组
    par: [{"groupNames": "1DCS", columnNames:["Unit"]}...]
         1) 添加组的操作具有原子性,即如果一次性添加多个组,但是有组添加失败,则此次添加会全部失败.
         2) 添加已存在的组会失败
         3) 添加的groupName不能是预定义的关键字(不区分大小写),否则会失败
         4) 添加的时候会去除groupName两端的空格
         5) 所有group的item都有默认的id(primary key)和itemName两列,所以此两列不需要在ColumnNames
            中定义,否则会添加失败.同时也不能是""
    res: {"effectedRows": int}
    egs: 1) 直接添加组(没有定义ColumnNames的group的item默认只有itemName和id两列)
         print(requests.post("http://192.168.0.199:8082/group/addGroups", data=json.dumps([{"groupNames": "1DCS"}])).text)
         >>> {"code":200,"message":"","data":{"rowsEffected":2}}   
         2) 添加已存在组
         print(requests.post("http://192.168.0.199:8082/group/addGroups", data=json.dumps([{"groupNames": "1DCS"}])).text)
         >>> {"code":500,"message":"sqliteExecutionError: UNIQUE constraint failed: group_cfg.groupName\n","data":null}
         3) 自定义ColumnNames
         data = [{"groupNames": "2DCS"}, {"groupNames": "3DCS", "columnNames": ["description", "units", "upper", "lower"]}]
         print(requests.post("http://192.168.0.199:8082/group/addGroups", data=json.dumps(data)).text)
         >>> {"code":200,"message":"","data":{"rowsEffected":2}}   
         4) 原子操作
         data = [{"groupNames": "4DCS"}, {"groupNames": "Go", "columnNames": ["description", "units", "upper", "lower"]}]
         print(requests.post("http://192.168.0.199:8082/group/addGroups", data=json.dumps(data)).text)
         >>> {"code":500,"message":"groupNameError:Go","data":null}  # illegal groupName
         data = [{"groupNames": "4DCS"}, {"groupNames": "3DCS", "columnNames": ["description", "units", "upper", "lower"]}]
         print(requests.post("http://192.168.0.199:8082/group/addGroups", data=json.dumps(data)).text)  # add repeatedly
         >>> {"code":500,"message":"sqliteExecutionError: UNIQUE constraint failed: group_cfg.groupName\n","data":null}
         data = [{"groupNames": "4DCS"}, {"groupNames": "5DCS", "columnNames": ["iD", "units", "upper", "lower"]}]
         print(requests.post("http://192.168.0.199:8082/group/addGroups", data=json.dumps(data)).text)
         >>> {"code":500,"message":"ColumnNameError:iD","data":null}  # ColumnName error
```
```
(2) url: /group/deleteGroups
    des: 删除GDB中的组
    par: {"groupNames": [groupName1,groupName2...]}
         删除操作不具有原子性,如果slice为空,则不执行任何操作
    res: {"effectedRows": int}
    egs: 1) 操作不具有原子性
         data = {"groupNames": ["3DCS", "4DCS"]}
         print(requests.post("http://192.168.0.199:8082/group/deleteGroups", data=json.dumps(data)).text)
         >>> {"code":500,"message":"sqliteExecutionError: no such table: 4DCS","data":null}  # 3DCS删除成功
```
```
(3) url: /group/getGroups
    des: 获取GDB中所有的组
    par: null
    res: {"groupNames": []string}
    egs:
         print(requests.post("http://192.168.0.199:8082/group/getGroups").text)
         >>> "code":200,"message":"","data":{"groupNames":["1DCS","2DCS"]}}
```
```
(4) url: /group/getGroupProperty
    des: 获取指定groupName中的列名以及其中item的数目
    par: {"groupNames": [groupName1,groupName2...]}
    res: {"groupName1": {"itemCount": int, "itemColumnNames": []}}
    egs:
         data = {"groupNames": ["1DCS", "2DCS"]}
         print(requests.post("http://192.168.0.199:8082/group/getGroupProperty", data=json.dumps(data)).text)
         >>> {"code":200,"message":"","data":{"1DCS":{"itemColumnNames":["id","itemName"],"itemCount":"0"},"2DCS":{"itemColumnNames":["id","itemName"],"itemCount":"0"}}}  # go中的map是无序的
```
```
(5) url: /group/updateGroupNames
    des: 修改GDB中的组名
    par: [{"oldGroupName": string, "newGroupName": string}]
         此操作具有原子性
    res: {"effectedRows": int}
    egs:
         data = [{"oldGroupName": "1DCS", "newGroupName": "3DCS"}, {"oldGroupName": "2DCS", "newGroupName": "4DCS"}]
         print(requests.post("http://192.168.0.199:8082/group/updateGroups", data=json.dumps(data)).text)
         >>> {"code":200,"message":"","data":{"effectedRows":2}}
         2) 操作具有原子性
         data = [{"oldGroupName": "3DCS", "newGroupName": "1DCS"}, {"oldGroupName": "4DCS", "newGroupName": "1DCS"}]
         print(requests.post("http://192.168.0.199:8082/group/updateGroups", data=json.dumps(data)).text)
         >>> {"code":500,"message":"sqliteExecutionError: UNIQUE constraint failed: group_cfg.groupName","data":null}
```
```
(6) url: /group/updateGroupColumnNames
    des: 更新group的列名
    par: {"groupName": string, "oldColumnNames": []string, "newColumnNames": []string}
         此操作具有原子性
    res: {"effectedCols": int}
    egs: 
         data = {"groupName": "1DCS", "oldColumnNames": ["units", "types"], "newColumnNames": ["unit", "type"]}
         print(requests.post("http://192.168.0.199:8082/group/updateGroupColumnNames", data=json.dumps(data)).text)
         >>> {"code": 200, "message":"","data":{"effectedCols": 2}}
```
```
(7) url: /group/deleteGroupColumns
    des: 删除group中的列
    par: {"groupName": string, "columnNames": []string}
         此操作具有原子性
    res: {"effectedCols": int}
    egs: 
         data = {"groupName": "2DCS", "columnNames": ["description", "units", "groupName"]}
         print(requests.post("http://192.168.0.199:8082/group/deleteGroupColumns", data=json.dumps(data)).text)
         >>> {"code": 200, "message":"","data":{"effectedCols": 3}}
```
```
    url: /group/addGroupColumns 
    des: 增加group中的列
    par: {"groupName": string, "columnNames": []string}
         此操作具有原子性
    res: {"effectedCols": int}
    egs: 
         data = {"groupName": "2DCS", "columnNames": ["description", "groupName"]}
         print(requests.post("http://192.168.0.199:8082/group/addGroupColumns", data=json.dumps(data)).text)
         >>> {"code": 200, "message":"","data":{"effectedCols": 2}}
```
<h2>3.Errors</h2>
```
1.groupNameError: 组名错误,组名不能在预定义关键字中
2.columnNameError: 列名错误,列名不能在预定义关键字中
3.sqliteConnectionError: 连接SQLite错误
4.sqliteExecutionError: 执行SQLite语句错误
5.sqliteRowsError: 获取SQLite数据错误
6.sqliteTransactionError:SQLite事物错误
```
<h2>4.预定义的关键字</h2>
```
gdbKeyWords = 
break      default       func     interface       select
case       defer         go       map             struct
chan       else          goto     package         switch
const      fallthrough   if       range           type
continue   for           import   return          var
gdb        from          where    char            varchar
int        smallint      numeric  real            double 
precision  float         primary  key             foreign 
references not           null     create          table
insert     into          values   delete          update
set        where         drop     alter           add
truncate   distinct      all      and             or
join       as            *        order           by
desc       asc           between  union           except
is         avg           min      max             sum
count      group         having   realtimedata historicaldata 
```
<i class="fa fa-github" style="margin-left: 250px;width: 30px"></i>https://github.com/JustKeepSilence