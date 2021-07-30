/**
 creatTime: 2021/7/13
 creator: JustKeepSilence
 github: https://github.com/JustKeepSilence
 nodeVersion: 14.16.1
 **/

// examples of handle data in gdb by restful in node

/*
                                                route permission of data handler
        url                                 visitor         common_user             super_user
        batchWrite<T>Data                     ×                  ×                       √
        batchWrite<T>HistoricalData           ×                  ×                       √
        getRealTimeData                       ×                  √                       √
        get<T>HistoricalData                  ×                  √                       √
        get<T>RawHistoricalData               ×                  ×                       √
        get<T>HistoricalDataWithStamp         ×                  √                       √
        get<T>HistoricalDataWithCondition     ×                  ×                       √
        delete<T>HistoricalData               ×                  ×                       √
        cleanItemData                         ×                  ×                       √
        reLoadDb                              ×                  ×                       √

        1. for routes in data handler, you can add it in gdbClient
        2. T in url replace Float or Int or String or Bool, eg, batchWrite<T>Data replace batchWriteFloatData,
        batchWriteIntData,batchWriteStringData,batchWriteBoolData,the following examples use Float as examples
        for other type, you need to chang T to corresponding type
        Float ==> float32
        Int ==> int32
        String ==> string
        Bool ==> bool
*/

'use strict'
const axios = require('axios')
const fs = require('fs')
const {ip, configs, mockFloat32Data, mockInt32Data, mockStringData, mockBoolData} = require('./utils')
const moment = require('moment')
const now = moment()
const startTime = now.add(1, 'hour').unix() + 8 * 3600   // startTime
const endTime = now.add(3, 'hour').unix() + 8 * 3600   // endTime

/**
 * write realTime Data to gdb==>batchWrite<T>Data
 *
 * @param(Array){groupNames}
 * @param(Array){itemNames} two-dimensional array, every items correspond to the group in groups
 * for itemNames whose dataType is T, you should use batchWrite<T>Data to write data, such as dataType of item1
 * is float32, you should use batchWriteFloatData to write realTimeData
 * @param(Array){itemValues} two-dimensional array, dataType should correspond to <T>
 *
 */
// write float32 realTimeData
axios.post(`${ip}/data/batchWriteFloatData`, {
    groupNames: ['3DCS', '5DCS'],
    itemNames: [['xFloat'], ['xFloat', 'yFloat']],
    itemValues: [[1.0], [5.2, 3.0]]
}, configs).then(({data: {data}}) => {
    console.log(data)
    // { effectedRows: 3, times: 0 }
}).catch(({response: {data: {message}}}) => {
    console.log('err: ' + message)
})

// write int32 realTimeData
axios.post(`${ip}/data/batchWriteIntData`, {
    groupNames: ['5DCS'],
    itemNames: [['xInt', 'yInt']],
    itemValues: [[2, 3]]
}, configs).then(({data: {data}}) => {
    console.log(data)
    // { effectedRows: 2, times: 0 }
}).catch(({response: {data: {message}}}) => {
    console.log('err: ' + message)
})

// write string realTimeData
axios.post(`${ip}/data/batchWriteStringData`, {
    groupNames: ['5DCS'],
    itemNames: [['xString', 'yString']],
    itemValues: [['2', '3']]
}, configs).then(({data: {data}}) => {
    console.log(data)
    // { effectedRows: 2, times: 0 }
})

// write bool realTimeData
axios.post(`${ip}/data/batchWriteBoolData`, {
    groupNames: ['5DCS'],
    itemNames: [['xBool', 'yBool']],
    itemValues: [[true, false]]
}, configs).then(({data: {data}}) => {
    console.log(data)
    // { effectedRows: 2, times: 0 }
}).catch(({response: {data: {message}}}) => {
    console.log('err: ' + message)
})

/**
 * get realTimeData from gdb
 *
 * 1. if realTime of item not exist, the return value will be null
 * 2. you MUST NOT get the realTime data of the same name item in different groups
 * such as {groupNames: ['3DCS', '5DCS', '5DCS'], itemNames: ['xInt', 'xInt', 'xString']}
 * otherWise the result of xInt will be unreliable
 *
 * @param(Array){groupNames}
 * @param(Array){itemNames}
 */
axios.post(`${ip}/data/getRealTimeData`, {
    groupNames: ['3DCS', '5DCS', '5DCS'],
    itemNames: ['xInt', 'xFloat', 'xString']
}, configs).then(({data: {data}}) => {
    console.log(data)
    // { realTimeData: { xFloat: 5.2, xInt: null, xString: '2' }, times: 0 }
}).catch(({response: {data: {message}}}) => {
    console.log('err: ' + message)
})

/**
 * write historicalData to gdb ==> batchWrite<T>HistoricalData
 *
 * @param(Array){groupName}
 * @param(Array){itemNames} for itemNames whose dataType is T, you should use batchWrite<T>HistoricalData to write
 * history data, such as dataType of item1 is float32, you should use batchWriteFloatHistoricalData to write history Data
 * @param(Array){itemValues}
 * @param(Array){timeStamps} timeStamp MUST be unix timeStamp
 */

// write float32 history data
let [floatTimeStamps, xFloatValues] = mockFloat32Data(now, Math.E)
let [, yFloatValues] = mockFloat32Data(now, Math.PI)
axios.post(`${ip}/data/batchWriteFloatHistoricalData`, {
    groupNames: ['5DCS', '5DCS'],
    itemNames: ['xFloat', 'yFloat'],
    itemValues: [xFloatValues, yFloatValues],
    timeStamps: [floatTimeStamps, floatTimeStamps]
}, configs).then(({data: {data}}) => {
    console.log(data)
}).catch(({response: {data: {message}}}) => {
    console.log('err: ' + message)
})

// write int32 history data
let [intTimeStamps, xIntValues] = mockInt32Data(now, Math.floor(Math.random() * (5 - 1 + 1) + 1))
let [, yIntValues] = mockInt32Data(now, Math.floor(Math.random() * (5 - 1 + 1) + 1))
axios.post(`${ip}/data/batchWriteIntHistoricalData`, {
    groupNames: ['5DCS', '5DCS'],
    itemNames: ['xInt', 'yInt'],
    itemValues: [xIntValues, yIntValues],
    timeStamps: [intTimeStamps, intTimeStamps]
}, configs).then(({data: {data}}) => {
    console.log(data)
}).catch(({response: {data: {message}}}) => {
    console.log(message)
})

// write string history data
let [stringTimeStamps, xStringValues] = mockStringData(now, Math.floor(Math.random() * (5 - 1 + 1) + 1).toString())
let [, yStringValues] = mockStringData(now, Math.floor(Math.random() * (5 - 1 + 1) + 1).toString())
axios.post(`${ip}/data/batchWriteStringHistoricalData`, {
    groupNames: ['5DCS', '5DCS'],
    itemNames: ['xString', 'yString'],
    itemValues: [xStringValues, yStringValues],
    timeStamps: [stringTimeStamps, stringTimeStamps]
}, configs).then(({data: {data}}) => {
    console.log(data)
})

// write bool history data
let [boolTimeStamps, xBoolValues] = mockBoolData(now, true)
let [, yBoolValues] = mockBoolData(now, false)
axios.post(`${ip}/data/batchWriteBoolHistoricalData`, {
    groupNames: ['5DCS', '5DCS'],
    itemNames: ['xBool', 'yBool'],
    itemValues: [xBoolValues, yBoolValues],
    timeStamps: [boolTimeStamps, boolTimeStamps]
}, configs).then(({data: {data}}) => {
    console.log(data)
})

/**
 * get history data from gdb ==> get<T>HistoricalData
 *
 * @param(Array){groupNames}
 * @param(Array){itemNames} for itemNames whose dataType is T, you should use get<T>HistoricalData to get
 * history data, such as dataType of item1 is float32, you should use getFloatHistoricalData to get history Data
 * @param(Array){startTimes}  startTime of history, MUST be unix timeStamp
 * @param(Array){endTimes} endTimes of history, MUST be unix timeStamp
 * @param(Array){intervals} intervals of getting history data, the scale is second
 * 1.you MUST NOT get the history data of the same name item in different groups
 * such as {groupNames: ['3DCS', '5DCS', '5DCS'], itemNames: ['xInt', 'xInt', 'xString']}
 * otherWise the result of xInt will be unreliable
 */

// get float32 history data
axios.post(`${ip}/data/getFloatHistoricalData`, {
    groupNames: ['5DCS', '5DCS'],
    itemNames: ['xFloat', 'yFloat'],
    startTimes: [startTime, startTime],
    endTimes: [endTime, endTime],
    intervals: [10, 10]
}, configs).then(({data: {data}}) => {
    console.log(data)
    // {
    //   historicalData: { xFloat: {timeStamps: [Array](timeStamp array), itemValues: [Array](value array) }, yFloat: { timeStamps: [Array], itemValues: [Array] } },
    //   times: 111  // operation times, scale is millionSecond
    // }
    //
    fs.writeFile('./f.json', JSON.stringify(data), err => {
        if (err) {
            console.log('fail to write json file:' + err)
        } else {
            console.log('write to json file successfully')
        }
    })
})

// get int32 history data
axios.post(`${ip}/data/getIntHistoricalData`, {
    groupNames: ['5DCS', '5DCS'],
    itemNames: ['xInt', 'yInt'],
    startTimes: [startTime, startTime],
    endTimes: [endTime, endTime],
    intervals: [10, 10]
}, configs).then(({data: {data}}) => {
    console.log(data)
    // {
    //   historicalData: { xInt: {timeStamps: [Array](timeStamp array), itemValues: [Array](value array) }, yInt: { timeStamps: [Array], itemValues: [Array] } },
    //   times: 111  // operation times, scale is millionSecond
    // }
    //
    fs.writeFile('./f.json', JSON.stringify(data), err => {
        if (err) {
            console.log('fail to write json file:' + err)
        } else {
            console.log('write to json file successfully')
        }
    })
})

// get string history data
axios.post(`${ip}/data/getStringHistoricalData`, {
    groupNames: ['5DCS', '5DCS'],
    itemNames: ['xString', 'yString'],
    startTimes: [startTime, startTime],
    endTimes: [endTime, endTime],
    intervals: [10, 10]
}, configs).then(({data: {data}}) => {
    console.log(data)
    // {
    //   historicalData: { xString: {timeStamps: [Array](timeStamp array), itemValues: [Array](value array) }, yString: { timeStamps: [Array], itemValues: [Array] } },
    //   times: 111  // operation times, scale is millionSecond
    // }
    //
    fs.writeFile('./f.json', JSON.stringify(data), err => {
        if (err) {
            console.log('fail to write json file:' + err)
        } else {
            console.log('write to json file successfully')
        }
    })
})

// get bool history data
axios.post(`${ip}/data/getBoolHistoricalData`, {
    groupNames: ['5DCS', '5DCS'],
    itemNames: ['xBool', 'yBool'],
    startTimes: [startTime, startTime],
    endTimes: [endTime, endTime],
    intervals: [10, 10]
}, configs).then(({data: {data}}) => {
    console.log(data)
    // {
    //   historicalData: { xBool: {timeStamps: [Array](timeStamp array), itemValues: [Array](value array) }, yBool: { timeStamps: [Array], itemValues: [Array] } },
    //   times: 111  // operation times, scale is millionSecond
    // }
    //
    fs.writeFile('./f.json', JSON.stringify(data), err => {
        if (err) {
            console.log('fail to write json file:' + err)
        } else {
            console.log('write to json file successfully')
        }
    })
})

/**
 * get all history of items from gdb ==> get<T>RawHistoricalData
 *
 * @param(Array){groupNames}
 * @param(Array){itemNames} for itemNames whose dataType is T, you should use get<T>RawHistoricalData to get
 * history data, such as dataType of item1 is float32, you should use getFloatRawHistoricalData to get history Data
 * 1.you MUST NOT get the history data of the same name item in different groups
 * such as {groupNames: ['3DCS', '5DCS', '5DCS'], itemNames: ['xInt', 'xInt', 'xString']}
 * otherWise the result of xInt will be unreliable
 */

// get all float32 history data
axios.post(`${ip}/data/getFloatRawHistoricalData`, {
    groupNames: ['5DCS', '5DCS'],
    itemNames: ['xFloat', 'yFloat']
}, configs).then(({data: {data}}) => {
    console.log(data)
    //{
    //   historicalData: { xFloat: {timeStamps: [Array](timeStamp array), itemValues: [Array](value array) }, yFloat: { timeStamps: [Array], itemValues: [Array] } },
    //   times: 381
    // }
    fs.writeFile('./fr.json', JSON.stringify(data), err => {
        if (err) {
            console.log('fail to write json file:' + err)
        } else {
            console.log('write to json file successfully')
        }
    })
})

// get all int32 history data
axios.post(`${ip}/data/getIntRawHistoricalData`, {
    groupNames: ['5DCS', '5DCS'],
    itemNames: ['xInt', 'yInt']
}, configs).then(({data: {data}}) => {
    console.log(data)
    //{
    //   historicalData: { xInt: {timeStamps: [Array](timeStamp array), itemValues: [Array](value array) }, yInt: { timeStamps: [Array], itemValues: [Array] } },
    //   times: 381
    // }
    fs.writeFile('./fr.json', JSON.stringify(data), err => {
        if (err) {
            console.log('fail to write json file:' + err)
        } else {
            console.log('write to json file successfully')
        }
    })
})

// get all string history data
axios.post(`${ip}/data/getStringRawHistoricalData`, {
    groupNames: ['5DCS', '5DCS'],
    itemNames: ['xString', 'yString']
}, configs).then(({data: {data}}) => {
    console.log(data)
    //{
    //   historicalData: { xString: {timeStamps: [Array](timeStamp array), itemValues: [Array](value array) }, yString: { timeStamps: [Array], itemValues: [Array] } },
    //   times: 381
    // }
    fs.writeFile('./fr.json', JSON.stringify(data), err => {
        if (err) {
            console.log('fail to write json file:' + err)
        } else {
            console.log('write to json file successfully')
        }
    })
})

// get all bool history data
axios.post(`${ip}/data/getBoolRawHistoricalData`, {
    groupNames: ['5DCS', '5DCS'],
    itemNames: ['xBool', 'yBool']
}, configs).then(({data: {data}}) => {
    console.log(data)
    //{
    //   historicalData: { xBool: {timeStamps: [Array](timeStamp array), itemValues: [Array](value array) }, yBool: { timeStamps: [Array], itemValues: [Array] } },
    //   times: 381
    // }
    fs.writeFile('./fr.json', JSON.stringify(data), err => {
        if (err) {
            console.log('fail to write json file:' + err)
        } else {
            console.log('write to json file successfully')
        }
    })
})

/**
 * get history data of items with specified timeStamps ==> get<T>HistoricalDataWithStamp
 *
 * @param(Array){groupNames}
 * @param(Array){itemNames} for itemNames whose dataType is T, you should use get<T>HistoricalDataWithStamp to get
 * history data, such as dataType of item1 is float32, you should use getFloatHistoricalDataWithStamp to get history Data
 * @param(Array){timeStamps}
 * 1.you MUST NOT get the history data of the same name item in different groups
 * such as {groupNames: ['3DCS', '5DCS', '5DCS'], itemNames: ['xInt', 'xInt', 'xString']}
 * otherWise the result of xInt will be unreliable
 */

// get float32 history data with ts
axios.post(`${ip}/data/getFloatHistoricalDataWithStamp`, {
    groupNames: ['5DCS', '5DCS'],
    itemNames: ['xFloat', 'yFloat'], timeStamps: [[startTime, endTime], [startTime, endTime]]
}, configs).then(({data: {data}}) => {
    console.log(data)
    // {
    //   historicalData: { xFloat: {timeStamps: [Array](timeStamp array), itemValues: [Array](value array) }, yFloat: { timeStamps: [Array], itemValues: [Array] } },
    //   times: 0
    // }
    console.log(JSON.stringify(data))
    // {historicalData:{xFloat:{timeStamps: [1627157580,1628021580], itemValues: [2.4523265,1.0813193]},yFloat:{timeStamps: [1627157580,1628021580],itemValues: [1.7954874,0.5654516]}},times:0}
})

// get int32 history data with ts
axios.post(`${ip}/data/getIntHistoricalDataWithStamp`, {
    groupNames: ['5DCS', '5DCS'],
    itemNames: ['xInt', 'yInt'], timeStamps: [[startTime, endTime], [startTime, endTime]]
}, configs).then(({data: {data}}) => {
    console.log(data)
    // {
    //   historicalData: { xInt: {timeStamps: [Array](timeStamp array), itemValues: [Array](value array) }, yInt: { timeStamps: [Array], itemValues: [Array] } },
    //   times: 0
    // }
    console.log(JSON.stringify(data))
    // {historicalData:{xInt:{timeStamps: [1627157734],itemValues: [105]},yInt:{timeStamps: [1627157734],itemValues: [95]}},times:0}
})

// get string history data with ts
axios.post(`${ip}/data/getStringHistoricalDataWithStamp`, {
    groupNames: ['5DCS', '5DCS'],
    itemNames: ['xString', 'yString'], timeStamps: [[startTime, endTime], [startTime, endTime]]
}, configs).then(({data: {data}}) => {
    console.log(data)
    // {
    //   historicalData: { xString: {timeStamps: [Array](timeStamp array), itemValues: [Array](value array) }, yString: { timeStamps: [Array], itemValues: [Array] } },
    //   times: 0
    // }
    console.log(JSON.stringify(data))
    // {historicalData:{xString:{timeStamps: [1627157788],itemValues: ["2"]},yString:{timeStamps: [1627157788],itemValues: ["5"]}},times:0}
})

// get bool history data with ts
axios.post(`${ip}/data/getBoolHistoricalDataWithStamp`, {
    groupNames: ['5DCS', '5DCS'],
    itemNames: ['xBool', 'yBool'], timeStamps: [[startTime, endTime], [startTime, endTime]]
}, configs).then(({data: {data}}) => {
    console.log(data)
    // {
    //   historicalData: { xBool: {timeStamps: [Array](timeStamp array), itemValues: [Array](value array) }, yBool: { timeStamps: [Array], itemValues: [Array] } },
    //   times: 0
    // }
    console.log(JSON.stringify(data))
    // {historicalData:{xBool:{timeStamps: [1627157824], itemValues: [true]},yBool:{timeStamps: [1627157824],itemValues: [false]}},times:0}
})

/**
 * get history data item with given condition ==> get<T>HistoricalDataWithCondition
 *
 * @param(String){groupName}
 * @param(Array){itemNames} for itemNames whose dataType is T, you should use get<T>HistoricalDataWithCondition to get
 * history data, such as dataType of item1 is float32, you should use getFloatHistoricalDataWithCondition to get history Data
 * @param(Array){startTime} timeStamp MUST be unix timeStamp
 * @param(Array){endTime} timeStamp MUST be unix timeStamp
 * @param(Array){interval} interval of getting history data, the scale is second
 * @param(String){filterCondition}  filter condition must be correct js expression,itemName should be startedWith by item.
 * eg: item["itemName1"]>10 && item["itemName2"] > 30 ....
 * @param(Array){deadZones} every item is dict, key is itemName(String), value is deadZoneCount(int)
 * DeadZone is used to define the maximum number of continuous data allowed by itemName.eg,the deadZoneCount of item x
 * is 2, that is all data in x whose number of continuous > 2 will be filtered,eg history of item1 is [1,1,1,1], if deadZoneCount of item1 is 2
 * after filter, the result will be [1,1]
 * 1. itemNames and itemName in filterCondition and zones MUST be in the same group, itemNames in filterCondition may be different dataType
 * 2. if you don't want to use filterCondition, you should set it "true"
 * 3. if you don't want to use deadZone condition, you should set it []
 * 4. judgment priority is startTime, endTime, interval > filterCondition > deadZone condition
 */

// get history with condition, without deadZones condition
axios.post(`${ip}/data/getFloatHistoricalDataWithCondition`, {
    groupName: '5DCS', itemNames: ['xFloat', 'yFloat'],
    startTime, endTime, interval: 1,
    filterCondition: `item["xFloat"]>= 1 && item["yFloat"]<=4 && item["xBool"]`, deadZones: []
}, configs).then(({data: {data}}) => {
    fs.writeFile('./fc.json', JSON.stringify(data), err => {
        if (err) {
            console.log('fail to write json file: ' + err.message)
        } else {
            console.log('ok')
        }
    })
}).catch(({response: {data: {message}}}) => {
    console.log(message)
})

// get history with filter condition and deadZones condition
axios.post(`${ip}/data/getFloatHistoricalDataWithCondition`, {
    groupName: '5DCS',
    itemNames: ['xFloat', 'yFloat'],
    startTime,
    endTime,
    interval: 1,
    filterCondition: `item["xFloat"]>= 1 && item["yFloat"]<=4 && item["xBool"]`,
    deadZones: [{itemName: "xFloat", deadZoneCount: 2}]
}, configs).then(({data: {data}}) => {
    console.log(data)
    // {
    //   historicalData: { xFloat: {timeStamps: [Array](timeStamp array), itemValues: [Array](value array) }, yFloat: { timeStamps: [Array], itemValues: [Array] } },
    //   times: 35
    // }
    fs.writeFile('./fcd.json', JSON.stringify(data), err => {
        if (err) {
            console.log('fail to write json file: ' + err.message)
        } else {
            console.log('ok')
        }
    })
}).catch(({response: {data: {message}}}) => {
    console.log('err: ' + message)
})

// get history without filter condition with deadZones condition
axios.post(`${ip}/data/getFloatHistoricalDataWithCondition`, {
    groupName: '5DCS', itemNames: ['xFloat', 'yFloat'],
    startTime, endTime, interval: 1,
    filterCondition: `true`, deadZones: [{itemName: "xFloat", deadZoneCount: 5}]
}, configs).then(({data: {data}}) => {
    console.log(data)
    // {
    //   historicalData: { xFloat: {timeStamps: [Array](timeStamp array), itemValues: [Array](value array) }, yFloat: { timeStamps: [Array], itemValues: [Array] } },
    //   times: 28
    // }
    fs.writeFile('./fd.json', JSON.stringify(data), err => {
        if (err) {
            console.log('fail to write json file: ' + err.message)
        } else {
            console.log('ok')
        }
    })
}).catch(({response: {data: {message}}}) => {
    console.log(message)
})

/**
 * deleteHistoryData from gdb ==> delete<T>HistoricalData
 *
 * @param(Array){groupName}
 * @param(Array){itemNames}
 * @param(Array){startTimes}
 * @param(Array){endTimes}
 */

// delete float32 history
axios.post(`${ip}/data/deleteFloatHistoricalData`, {
    groupNames: ['5DCS', '5DCS'], itemNames: ['xFloat', 'yFloat'],
    startTimes: [startTime, startTime], endTimes: [endTime, endTime]
}, configs).then(({data: {data}}) => {
    console.log(data)
    // { effectedRows: 43200, times: 3 }
}).catch(({response: {data: {message}}}) => {
    console.log(message)
})

// delete int32 history
axios.post(`${ip}/data/deleteIntHistoricalData`, {
    groupNames: ['5DCS', '5DCS'], itemNames: ['xInt', 'yInt'],
    startTimes: [startTime, startTime], endTimes: [endTime, endTime]
}, configs).then(({data: {data}}) => {
    console.log(data)
    // { effectedRows: 43200, times: 3 }
}).catch(({response: {data: {message}}}) => {
    console.log(message)
})

// delete string history
axios.post(`${ip}/data/deleteStringHistoricalData`, {
    groupNames: ['5DCS', '5DCS'], itemNames: ['xString', 'yString'],
    startTimes: [startTime, startTime], endTimes: [endTime, endTime]
}, configs).then(({data: {data}}) => {
    console.log(data)
    // { effectedRows: 43200, times: 3 }
}).catch(({response: {data: {message}}}) => {
    console.log(message)
})

// delete bool history
axios.post(`${ip}/data/deleteBoolHistoricalData`, {
    groupNames: ['5DCS', '5DCS'], itemNames: ['xBool', 'yBool'],
    startTimes: [startTime, startTime], endTimes: [endTime, endTime]
}, configs).then(({data: {data}}) => {
    console.log(data)
    // { effectedRows: 43200, times: 3 }
}).catch(({response: {data: {message}}}) => {
    console.log(message)
})

/**
 * delete item and history data of the given item ==> cleanItemData
 *
 * @param(String){groupName}
 * @param(String){condition}
 * this operation may take longtime and may block writing operation, you should use it carefully
 */
axios.post(`${ip}/data/cleanItemData`, {groupName: '5DCS', condition: '1=1'}, configs).then(({data: {data}}) => {
    console.log(data)
    // { effectedRows: 8, times: 4395 }
}).catch(({response: {data: {message}}}) => {
    console.log('err: ' + message)
})

/**
 * reLoadDb: delete keys in db whose value is nil, then compact db again, during this time, all write operation will failed
 * so you should not write data to db during reload
 */
axios.post(`${ip}/data/reLoadDb`, {}, configs).then(({data: {data}}) => {
    console.log(data)
    // { effectedRows: 5760, times: 470 }
}).catch(({response: {data: {message}}}) => {
    console.log('err: ' + message)
})