# Data
```
定义了和数据相关的接口
```
<h2>1.Api</h2>
```
(1) url: /data/batchWrite
    des: 向数据库中批量写入实时数据,如果itemValues应该分别为itemNames对应的slice和itemValue对应的slice
         如果需要写入指定的时间戳,则第三个元素应为对应的时间戳,并且时间戳应该为year-month-day hh:mm:ss对
         应的unixTimeStamp,如果没有指定则写入的时间戳为写入时的时间戳.写入的item只要有一个在数据库中不存在
         写入就会失败
    par: {"groupName": string, "itemValues":[[]string,[]string,[]string]}  
    res: {"effectedRows": int}
    egs: print(requests.post("http://192.168.0.199:8082/data/batchWrite", data=json.dumps([{"groupNames": "1DCS", "itemValues": [["item1", "item2"], ["1", "2"]]}])).text) # without timeStamp
         >>> {"effectedRows": 2}
         t = int(datetime(2021, 1, 4, 20, 11, 29)).timestamp()  # js: (new Date("2020-1-4 20:11:29").getTime() / 1000), go: time.Date(2020, 1, 4, 20, 11, 29, 0, time.UTC).Unix() - 8 * 3600
         print(requests.post("http://192.168.0.199:8082/data/batchWrite", data=json.dumps([{"groupNames": "1DCS", "itemValues": [["item1", "item2"], ["1", "2"]， [t, t]]}])).text) # with timeStamp
         >>> {"effectedRows": 2}
(2) url: /data/getRealTimeData
    des: 获取指定item的实时值,如果item不存在,则返回的实时值为null
    par: {"itemNames": []string}
    res: {"itemName1", v1, "itemName2": v2}
    egs: print(requests.post("http://192.168.0.199:8082/data/getRealTimeData", data=json.dumps({"itemNames": ["item1", "item2", "item3"]})).text)
         >>> {"item1": "1", "item2": "2", "item3": null}
(3) url: /data/getHistoricalData
    des: 获取指定item在指定时间范围内指定取数间隔下的的历史值,startTime和endTime可以为多段时间戳,interval为对应的取数间隔,单位为s
    par: {"itemNames": []string, "startTime": []int, "endTime": []int, "interval": []int}
    res: {"item1": [[timeStamp...], [values...]]}
    egs: print(requests.post("http://192.168.0.199:8082/data/getHistoricalData", data=json.dumps({"itemNames":["item1"],"startTime":[1609761600],"endTime":[1609763580],"interval":[360]})).text)
         >>> {"item1":[["1609761600","1609761985","1609762369","1609762758","1609763146","1609763535"],["778","88","278","746","865","876"]]}
(4) url: /data/getHistoricalDataWithStamp
    des: 获取指定item指定时间戳的历史数据
    par: {"itemNames": []string, "timeStamps": []string}
    res: {"item1": [[timeStamp...], [values...]]}
    egs: print(requests.post("http://192.168.0.199:8082/data/getHistoricalDataWithStamp", data=json.dumps({"itemNames":["item1"],"timeStamps":["1609761600","1609761985","1609762369","1609762758","1609763146","1609763535"]})).text)
         >>> {"item1":[["1609761600","1609761985","1609762369","1609762758","1609763146","1609763535"],["778","88","278","746","865","876"]]}
(5) url: /data/getHistoricalDataWithCondition
    des: 在getHistoricalData的基础之上可以获取指定筛选条件的历史数据,filterCondition为筛选条件,可以是任何符合语法的js语句,需要注意的是如果和item相关
         则形式应该为item["item1"]>10之类,deadZones为item对应的死区个数,即当有多少个值是相同的时候,则认为这段数据是无效的.如果不提供filterCondition
         或者filterCondition为空字符串则认为该项为true,如果不提供deadZones或者deadZones为[],则认为该项为true,即不进行死区筛选
    par: {"filterCondition": string, deadZones: []int}
    egs: print(requests.post("http://192.168.0.199:8082/data/getHistoricalDataWithCondition", data=json.dumps({"itemNames":["item1"],"startTime":[1609761600],"endTime":[1609763580],"interval":[360], "filterCondition": 'item["item2"]<800 && item["item1"] > 100', deadZones: [5]})).text)  
         # 获取item1>100并且item2<800时item1的数据,并且item1中连续的数据段不能超过5个(超过5个认为item1不刷新,数据无效)
         >>> {"item1":[["1609761600", "1609762369","1609762758","1609763146","1609763535"],["778","278","746","865","876"]]}
```

