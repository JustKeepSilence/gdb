// getHData(groupNames, itemNames []string, startTimeStamps []int, endTimeStamps []int, intervals []int):获取指定item的历史值，返回值为map
// getNowTime(): 获取当前时间.返回的格式为yyyy-MM-dd hh:mm:ss
// getTimeStamp(t, d): 获取指定时间t,时间间隔为ds时间的时间戳

var now = getNowTime()    // 获取当前时间
var st = parseInt(getTimeStamp(now, -24 * 3600))  // 起始时间为一天前
var et = parseInt(getTimeStamp(now, 0))  // 终止时间为现在
var interval = 1 // 取数间隔为60s
var itemNames = ["t1", "t2"]
var data = JSON.parse(getHData(["calc", "calc"],itemNames,[st], [et], [interval]))   // 获取历史数据,返回的为字符串:{"itemName": [[timeStamps],[values]]}
var sum = 0.0
for(var i = 0; i < itemNames.length;i++){
    var x = data[itemNames[i]][1]
    if (x !== null){
        // 历史数据存在
        var a = getAverage(x)   // 计算这段时间内历史数据的平均值
    }
    sum += 1.2
}
writeRtData([{"groupName": "calc", "itemName": "y", "value": sum}])  // 写入

function getAverage(a){
    var sum = 0;
    for(var i =0; i < a.length;i++){
        sum += a[i]
    }
    return sum / a.length
}
