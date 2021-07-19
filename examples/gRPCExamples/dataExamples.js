/**
 creatTime: 2021/7/19
 creator: JustKeepSilence
 github: https://github.com/JustKeepSilence
 nodeVersion: 14.16.1
 **/

// gRPC examples of data handler in gdb, about parameter info, you can see:
// https://github.com/JustKeepSilence/gdb/blob/master/examples/restfulExamples/dataExamples.js

'use strict'
const fs = require('fs')
const {
    grpcRequest,
    mockFloat32Data,
    mockInt32Data,
    mockStringData,
    mockBoolData,
    streamWriteHistoryData
} = require('./utils')
const moment = require('moment')
const now = moment()
const endNow = moment()
const startTime = now.add(8, 'hour').unix() + 8 * 3600   // startTime
const endTime = endNow.add(9, 'hour').unix() + 8 * 3600   // endTime

// write float32 realTimeData
grpcRequest('/data/batchWriteFloatData', {
    groupNames: ['5DCS'],
    itemNames: [['xFloat', 'yFloat']],
    itemValues: [[1.0, 2.0]]
}).then(({data}) => {
    console.log(data)
    // { effectedRows: 2, times: '0' }
}).catch(({message}) => {
    console.log(message)
})

// write int32 realTimeData
grpcRequest('/data/batchWriteIntData', {
    groupNames: ['5DCS'],
    itemNames: [['xInt', 'yInt']],
    itemValues: [[2, 3]]
}).then(({data}) => {
    console.log(data)
    // { effectedRows: 2, times: '0' }
}).catch(({message}) => {
    console.log(message)
})

// write string realTimeData
grpcRequest('/data/batchWriteStringData', {
    groupNames: ['5DCS'],
    itemNames: [['xString', 'yString']],
    itemValues: [['2', '3']]
}).then(({data}) => {
    console.log(data)
    // { effectedRows: 2, times: '0' }
}).catch(({message}) => {
    console.log(message)
})

// write bool realTimeData
grpcRequest('/data/batchWriteBoolData', {
    groupNames: ['5DCS'],
    itemNames: [['xBool', 'yBool']],
    itemValues: [[true, false]]
}).then(({data}) => {
    console.log(data)
    // { effectedRows: 2, times: '0' }
}).catch(({message}) => {
    console.log(message)
})

// get realTimeData from gdb
grpcRequest('/data/getRealTimeData', {
    groupNames: ['5DCS', '5DCS', '5DCS', '5DCS'],
    itemNames: ['xInt', 'xFloat', 'xString', 'xBool']
}).then(({data}) => {
    console.log(data)
    // {
    //   realTimeData: '{"xBool":true,"xFloat":1,"xInt":2,"xString":"2"}',
    //   times: '0'
    // }
}).catch(({message}) => {
    console.log(message)
})

// write float32 history data
const [floatTimeStamps, xFloatValues] = mockFloat32Data(now, Math.E)
const [, yFloatValues] = mockFloat32Data(now, Math.PI)
grpcRequest('/data/batchWriteFloatHistoricalData', {
    groupNames: ['5DCS', '5DCS'],
    itemNames: ['xFloat', 'yFloat'],
    itemValues: [xFloatValues, yFloatValues],
    timeStamps: [floatTimeStamps, floatTimeStamps]
}).then(({data}) => {
    console.log(data)
    // { effectedRows: 1728000, times: '728' }
}).catch(({message}) => {
    console.log(message)
})

// write float32 history data withStream
streamWriteHistoryData('float32', {
    groupNames: ['5DCS', '5DCS'],
    itemNames: ['xFloat', 'yFloat'],
    itemValues: [xFloatValues, yFloatValues],
    timeStamps: [floatTimeStamps, floatTimeStamps]
}).then(({data}) => {
    console.log(data)
}).catch(({message}) => {
    console.log(message)
})

// write int32 history data
const [intTimeStamps, xIntValues] = mockInt32Data(now, Math.floor(Math.random() * (5 - 1 + 1) + 1))
const [, yIntValues] = mockInt32Data(now, Math.floor(Math.random() * (5 - 1 + 1) + 1))
grpcRequest('/data/batchWriteIntHistoricalData', {
    groupNames: ['5DCS', '5DCS'],
    itemNames: ['xInt', 'yInt'],
    itemValues: [xIntValues, yIntValues],
    timeStamps: [intTimeStamps, intTimeStamps]
}).then(({data}) => {
    console.log(data)
    // { effectedRows: 1728000, times: '728' }
}).catch(({message}) => {
    console.log(message)
})

// write int32 history data with stream
streamWriteHistoryData('int32', {
    groupNames: ['5DCS', '5DCS'],
    itemNames: ['xInt', 'yInt'],
    itemValues: [xIntValues, yIntValues],
    timeStamps: [intTimeStamps, intTimeStamps]
}).then(({data}) => {
    console.log(data)
    // { effectedRows: 1728000, times: '728' }
}).catch(({message}) => {
    console.log(message)
})

// write string history data
const [stringTimeStamps, xStringValues] = mockStringData(now, Math.floor(Math.random() * (5 - 1 + 1) + 1).toString())
const [, yStringValues] = mockStringData(now, Math.floor(Math.random() * (5 - 1 + 1) + 1).toString())
grpcRequest('/data/batchWriteStringHistoricalData', {
    groupNames: ['5DCS', '5DCS'],
    itemNames: ['xString', 'yString'],
    itemValues: [xStringValues, yStringValues],
    timeStamps: [stringTimeStamps, stringTimeStamps]
}).then(({data}) => {
    console.log(data)
    // { effectedRows: 1728000, times: '728' }
}).catch(({message}) => {
    console.log(message)
})

// write string history with stream
streamWriteHistoryData('string', {
    groupNames: ['5DCS', '5DCS'],
    itemNames: ['xString', 'yString'],
    itemValues: [xStringValues, yStringValues],
    timeStamps: [stringTimeStamps, stringTimeStamps]
}).then(({data}) => {
    console.log(data)
    // { effectedRows: 1728000, times: '728' }
}).catch(({message}) => {
    console.log(message)
})

// write bool history data
const [boolTimeStamps, xBoolValues] = mockBoolData(now, true)
const [, yBoolValues] = mockBoolData(now, false)
grpcRequest('/data/batchWriteBoolHistoricalData', {
    groupNames: ['5DCS', '5DCS'],
    itemNames: ['xBool', 'yBool'],
    itemValues: [xBoolValues, yBoolValues],
    timeStamps: [boolTimeStamps, boolTimeStamps]
}).then(({data}) => {
    console.log(data)
    // { effectedRows: 1728000, times: '728' }
}).catch(({message}) => {
    console.log(message)
})

// write bool history with stream
streamWriteHistoryData('bool', {
    groupNames: ['5DCS', '5DCS'],
    itemNames: ['xBool', 'yBool'],
    itemValues: [xBoolValues, yBoolValues],
    timeStamps: [boolTimeStamps, boolTimeStamps]
}).then(({data}) => {
    console.log(data)
    // { effectedRows: 1728000, times: '728' }
}).catch(({message}) => {
    console.log(message)
})

// get float32 history data, unlike restful,the returned data of gRPC is string
grpcRequest('/data/getFloatHistoricalData', {
    groupNames: ['5DCS', '5DCS'],
    itemNames: ['xFloat', 'yFloat'],
    startTimes: [startTime, startTime],
    endTimes: [endTime, endTime],
    intervals: [10, 10]
}).then(({data: {historicalData}}) => {
    fs.writeFile('./f.json', historicalData, err => {
        if (err) {
            console.log('err: ' + err.message)
        } else {
            console.log('ok')
        }
    })
})

// get int32 history data, unlike restful,the returned data of gRPC is string
grpcRequest('/data/getIntHistoricalData', {
    groupNames: ['5DCS', '5DCS'],
    itemNames: ['xInt', 'yInt'],
    startTimes: [startTime, startTime],
    endTimes: [endTime, endTime],
    intervals: [10, 10]
}).then(({data: {historicalData}}) => {
    fs.writeFile('./f.json', historicalData, err => {
        if (err) {
            console.log('err: ' + err.message)
        } else {
            console.log('ok')
        }
    })
})

// get string history data, unlike restful,the returned data of gRPC is string
grpcRequest('/data/getStringHistoricalData', {
    groupNames: ['5DCS', '5DCS'],
    itemNames: ['xString', 'yString'],
    startTimes: [startTime, startTime],
    endTimes: [endTime, endTime],
    intervals: [10, 10]
}).then(({data: {historicalData}}) => {
    fs.writeFile('./f.json', historicalData, err => {
        if (err) {
            console.log('err: ' + err.message)
        } else {
            console.log('ok')
        }
    })
})

// get bool history data, unlike restful,the returned data of gRPC is string
grpcRequest('/data/getBoolHistoricalData', {
    groupNames: ['5DCS', '5DCS'],
    itemNames: ['xBool', 'yBool'],
    startTimes: [startTime, startTime],
    endTimes: [endTime, endTime],
    intervals: [10, 10]
}).then(({data: {historicalData}}) => {
    fs.writeFile('./f.json', historicalData, err => {
        if (err) {
            console.log('err: ' + err.message)
        } else {
            console.log('ok')
        }
    })
})

// get all float32 history data, unlike restful,the returned data of gRPC is string
grpcRequest('/data/getFloatRawHistoricalData', {
    groupNames: ['5DCS'],
    itemNames: ['xFloat']
}).then(({data: {historicalData}}) => {
    fs.writeFile('./fr.json', historicalData, err => {
        if (err) {
            console.log('err: ' + err.message)
        } else {
            console.log('ok')
        }
    })
}).catch(({message}) => {
    console.log('err: ' + message)
})

// get all int32 history data, unlike restful,the returned data of gRPC is string
grpcRequest('/data/getIntRawHistoricalData', {
    groupNames: ['5DCS'],
    itemNames: ['xInt']
}).then(({data: {historicalData}}) => {
    fs.writeFile('./fr.json', historicalData, err => {
        if (err) {
            console.log('err: ' + err.message)
        } else {
            console.log('ok')
        }
    })
}).catch(({message}) => {
    console.log('err: ' + message)
})

// get all string history data, unlike restful,the returned data of gRPC is string
grpcRequest('/data/getStringRawHistoricalData', {
    groupNames: ['5DCS'],
    itemNames: ['xString']
}).then(({data: {historicalData}}) => {
    fs.writeFile('./fr.json', historicalData, err => {
        if (err) {
            console.log('err: ' + err.message)
        } else {
            console.log('ok')
        }
    })
}).catch(({message}) => {
    console.log('err: ' + message)
})

// get all bool history data, unlike restful,the returned data of gRPC is string
grpcRequest('/data/getBoolRawHistoricalData', {
    groupNames: ['5DCS'],
    itemNames: ['xBool']
}).then(({data: {historicalData}}) => {
    fs.writeFile('./fr.json', historicalData, err => {
        if (err) {
            console.log('err: ' + err.message)
        } else {
            console.log('ok')
        }
    })
}).catch(({message}) => {
    console.log('err: ' + message)
})

// get float32 history data with timeStamp, unlike restful,the returned data of gRPC is string
grpcRequest('/data/getFloatHistoricalDataWithStamp', {
    groupNames: ['5DCS', '5DCS'],
    itemNames: ['xFloat', 'yFloat'],
    timeStamps: [[startTime, endTime], [startTime, endTime]]
}).then(({data: {historicalData}}) => {
    console.log(historicalData)
    // {"xFloat":[[1626757086,1626767886],[1.2652606,1.0157638]],"yFloat":[[1626757086,1626767886],[0.03733571,2.721263]]}
})

// get int32 history data with timeStamp, unlike restful,the returned data of gRPC is string
grpcRequest('/data/getIntHistoricalDataWithStamp', {
    groupNames: ['5DCS', '5DCS'],
    itemNames: ['xInt', 'yInt'],
    timeStamps: [[startTime, endTime], [startTime, endTime]]
}).then(({data: {historicalData}}) => {
    console.log(historicalData)
    // {"xInt":[[1626757066,1626767866],[49,47]],"yInt":[[1626757066,1626767866],[5,18]]}
})

// get string history data with timeStamp, unlike restful,the returned data of gRPC is string
grpcRequest('/data/getStringHistoricalDataWithStamp', {
    groupNames: ['5DCS', '5DCS'],
    itemNames: ['xString', 'yString'],
    timeStamps: [[startTime, endTime], [startTime, endTime]]
}).then(({data: {historicalData}}) => {
    console.log(historicalData)
    // {"xString":[[1626757050,1626767850],["2","2"]],"yString":[[1626757050,1626767850],["1","1"]]}
})

// get bool history data with timeStamp, unlike restful,the returned data of gRPC is string
grpcRequest('/data/getBoolHistoricalDataWithStamp', {
    groupNames: ['5DCS', '5DCS'],
    itemNames: ['xBool', 'yBool'],
    timeStamps: [[startTime, endTime], [startTime, endTime]]
}).then(({data: {historicalData}}) => {
    console.log(historicalData)
    // {"xBool":[[1626757025,1626767825],[true,true]],"yBool":[[1626757025,1626767825],[false,false]]}
})

// get history with condition, without deadZones condition
grpcRequest('/data/getFloatHistoricalDataWithCondition', {
    groupName: '5DCS', itemNames: ['xFloat', 'yFloat'],
    startTime, endTime, interval: 1,
    filterCondition: `item["xFloat"]>= 1 && item["yFloat"]<=4 && item["xBool"]`, deadZones: []
}).then(({data: {historicalData}}) => {
    fs.writeFile('./fc.json', historicalData, err => {
        if (err) {
            console.log('err: ' + err.message)
        } else {
            console.log('ok')
        }
    })
})

// get history with filter condition and deadZones condition
grpcRequest('/data/getFloatHistoricalDataWithCondition', {
    groupName: '5DCS',
    itemNames: ['xFloat', 'yFloat'],
    startTime,
    endTime,
    interval: 1,
    filterCondition: `item["xFloat"]>= 1 && item["yFloat"]<=4 && item["xBool"]`,
    deadZones: [{itemName: "xFloat", deadZoneCount: 2}]
}).then(({data: {historicalData}}) => {
    fs.writeFile('./fcd.json', historicalData, err => {
        if (err) {
            console.log('err: ' + err.message)
        } else {
            console.log('ok')
        }
    })
})

// get history without filter condition with deadZones condition
grpcRequest('/data/getFloatHistoricalDataWithCondition', {
    groupName: '5DCS', itemNames: ['xFloat', 'yFloat'],
    startTime, endTime, interval: 1,
    filterCondition: `true`, deadZones: [{itemName: "xFloat", deadZoneCount: 5}]
}).then(({data: {historicalData}}) => {
    fs.writeFile('./fd.json', historicalData, err => {
        if (err) {
            console.log('err: ' + err.message)
        } else {
            console.log('ok')
        }
    })
})
