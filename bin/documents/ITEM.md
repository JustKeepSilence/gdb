# Item
```
Item是GDB存储数据的单元,是group的内容,一个group里面可以有
N个Item,每个Item的属性在添加group的时候已经通过Column定义
```
<h2>1.Api</h2>
```
(1) url: /item/addItems
    des: 向指定的group添加item
    par: {"groupName": string, "values":[{"column1": s1, "column1": s2..}..]}
    res: {"effectedRows": int}
    egs: 
(2) url: /item/deleteItems
    des: 删除指定的item
    par: {"groupName": string, "condition": string}
    res: {"effectedRows", int}
    egs:
(3) url: /item/getItems
    des: 获取指定的item,如果startRow为-1则意味着查询全部,此时不需要提供rowCount参数,否则需要同时提供
         startRow和rowCount两个参数来进行Limit分页查询,column表示要查询的列,以逗号分割,*表示所有
         可以参见SQLite的查询语句
    par: {"groupName": string, "column": string, "condition": string, "startRow": int, rowCount: int}
    res: [{column: ...},{}]
    egs:
(4) url: /item/updateItems
    des: 更新制定的item
    par: {"groupName": string, "clause": string, "condition": string}
    res: {"effectedRows": int}
```

