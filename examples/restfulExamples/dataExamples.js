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
const {ip, configs, mockFloat32Data, mockInt32Data, mockStringData, mockBoolData} = require('./utils')
const moment = require('moment')
const now = moment()

/**
 * write realTime Data to gdb==>batchWrite<T>Data
 *
 * @param(Array){groupNames}
 * @param(Array){itemNames} two-dimensional array, every items correspond to the group in groups
 * @param(Array){itemValues} two-dimensional array, dataType should correspond to <T>
 *
 */
// axios.post(`${ip}/data/batchWriteFloatData`, {groupNames: ['3DCS', '5DCS'], itemNames: [['xFloat'], ['xFloat', 'yFloat']], itemValues: [[1.0], [5.2, 3.0]]}, configs).then(({data :{data}})=>{
//     console.log(data)
//     // { effectedRows: 3, times: 0 }
// })
// axios.post(`${ip}/data/batchWriteIntData`, {groupNames: ['5DCS'], itemNames: [['xInt', 'yInt']], itemValues: [[2, 3]]}, configs).then(({data :{data}})=>{
//     console.log(data)
        // { effectedRows: 2, times: 0 }
// })
// axios.post(`${ip}/data/batchWriteStringData`, {groupNames: ['5DCS'], itemNames: [['xString', 'yString']], itemValues: [['2', '3']]}, configs).then(({data :{data}})=>{
//     console.log(data)
        // { effectedRows: 2, times: 0 }
// })
// axios.post(`${ip}/data/batchWriteBoolData`, {groupNames: ['5DCS'], itemNames: [['xBool', 'yBool']], itemValues: [[true, false]]}, configs).then(({data :{data}})=>{
//     console.log(data)
       // { effectedRows: 2, times: 0 }
// })

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
// axios.post(`${ip}/data/getRealTimeData`, {groupNames: ['3DCS', '5DCS', '5DCS'], itemNames: ['xInt', 'xFloat', 'xString']}, configs).then(({data: {data}})=>{
//     console.log(data)
//     // { realTimeData: { xFloat: 5.2, xInt: null, xString: '2' }, times: 0 }
// })

/**
 * write historicalData to gdb ==> batchWrite<T>HistoricalData
 *
 * @param(Array){groupName}
 * @param(Array){itemNames}
 * @param(Array){itemValues}
 * @param(Array){timeStamps} timeStamp MUST be unix timeStamp
 *
 */
// let [timeStamps, xValues] = mockFloat32Data(now, Math.E)
// let [,yValues] = mockFloat32Data(now, Math.PI)
// axios.post(`${ip}/data/batchWriteFloatHistoricalData`, {groupNames: ['5DCS', '5DCS'], itemNames: ['xFloat', 'yFloat'], itemValues:[xValues, yValues], timeStamps: [timeStamps, timeStamps]}, configs).then(({data: {data}})=>{
//     console.log(data)
// })
// let [timeStamps, xValues] = mockInt32Data(now, Math.floor(Math.random() * (5 - 1 + 1) + 1))
// let [,yValues] = mockInt32Data(now, Math.floor(Math.random() * (5 - 1 + 1) + 1))
// axios.post(`${ip}/data/batchWriteIntHistoricalData`, {groupNames: ['5DCS', '5DCS'], itemNames: ['xInt', 'yInt'], itemValues:[xValues, yValues], timeStamps: [timeStamps, timeStamps]}, configs).then(({data: {data}})=>{
//     console.log(data)
// })
// let [timeStamps, xValues] = mockStringData(now, Math.floor(Math.random() * (5 - 1 + 1) + 1).toString())
// let [,yValues] = mockStringData(now, Math.floor(Math.random() * (5 - 1 + 1) + 1).toString())
// axios.post(`${ip}/data/batchWriteStringHistoricalData`, {groupNames: ['5DCS', '5DCS'], itemNames: ['xString', 'yString'], itemValues:[xValues, yValues], timeStamps: [timeStamps, timeStamps]}, configs).then(({data: {data}})=>{
//     console.log(data)
// })
let [timeStamps, xValues] = mockBoolData(now,true)
let [,yValues] = mockBoolData(now, false)
axios.post(`${ip}/data/batchWriteBoolHistoricalData`, {groupNames: ['5DCS', '5DCS'], itemNames: ['xBool', 'yBool'], itemValues:[xValues, yValues], timeStamps: [timeStamps, timeStamps]}, configs).then(({data: {data}})=>{
    console.log(data)
})