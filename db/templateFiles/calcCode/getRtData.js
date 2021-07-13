// getRtData(itemNames: []string{}): 获取gdb中的实时值
// 输入参数：要获取的item
var r = JSON.parse(getRtData(["t1", "t2"], ["calc", "calc"]))  // 获取点的实时值,返回的数据类型均为string: {"t1": v1, "t2": v2}
if (r["t1"] !== null && r["t2"] !== null){
    // 一定要判断取到的数值是否为空来避免向数据库写入NaN
    var x = (r["t1"] + r["t2"]) / 2
    writeRtData([{"itemName": "x", "value": x, "groupName": "calc"}])  // 写入单点的实时值
}
