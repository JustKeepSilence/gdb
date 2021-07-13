// getHDataWithTs(groupNames, itemNames []string, timeStamps [][]int): 获取指定item在指定ts时刻的历史数据

var now = getNowTime()
var ts = parseInt(getTimeStamp(now, -60))  // 获取前1分钟时的历史值
var r = JSON.parse(getHDataWithTs(["calc", "calc"], ["x", "y"], [[ts], [ts]]))  // 获取历史值，返回的历史值为:{"x": [[timeStamps], [values]], "y": [[timeStamps],[values]]}
testItemValue(r)
if (r["x"][1].indexOf(null) === -1 && r["y"][1].indexOf(null) === -1){
    // 注意getHDataWithTs返回的是数组中可能存在null，所以需要使用indexOf来判断
    var z = (r["x"][1][0] + r["y"][1][0]) / 2
    writeRtData([{"itemName": "z", "value": z, "groupName": "calc"}]) 
}else{
    writeRtData([{"itemName": "z", "value": 1.2, "groupName": "calc"}]) 
}