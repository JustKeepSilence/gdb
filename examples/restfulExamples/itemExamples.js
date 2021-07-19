/**
 creatTime: 2021/7/13
 creator: JustKeepSilence
 github: https://github.com/JustKeepSilence
 nodeVersion: 14.16.1
 **/

// examples of handle items in gdb by restful in node

/*
                                route permission of item handler
        url                         visitor         common_user             super_user
        addItems                     ×                  √                       √
        deleteItems                  ×                  ×                       √
        getItemsWithCount            √                  √                       √
        updateItems                  ×                  ×                       √
        checkItems                   ×                  √                       √
        cleanGroupItems              ×                  ×                       √

        for item route permissions, you can't change in gdbClient
*/

'use strict'
const axios = require('axios')
const {configs} = require('./utils')
const {ip} = require('./utils')

/**
 * add items to group in gdb
 *
 * 1. value of dataType can only be float32 or int32 or string or bool
 * 2. for custom columns, if not specify values when adding item, gdb will use default value when adding columns
 * for system column itemName and dataType, you MUST specify value
 * 3. itemName in group can't be duplicate
 * 4. only super_user or common_user have this route permission
 * @param(String){groupName}
 * @param(Array){itemValues} every item of itemValues should be like {"column1": "v1", "column2": "v2"}
 */
axios.post(`${ip}/item/addItems`, {
    groupName: '5DCS',
    itemValues: [{itemName: 'xFloat', dataType: 'float32'}, {itemName: 'yFloat', dataType: 'float32'},
        {itemName: 'xInt', dataType: 'int32'}, {itemName: 'yInt', dataType: 'int32'}, {
            itemName: 'xString',
            dataType: 'string'
        },
        {itemName: 'yString', dataType: 'string'}, {itemName: 'xBool', dataType: 'bool'}, {
            itemName: 'yBool',
            dataType: 'bool'
        }]
}, configs).then(({data: {data}}) => {
    console.log(data)
    // { effectedRows: 8, times: 151 }
})

/**
 * delete item from group in gdb
 *
 * 1. only super_user has this route permission
 * 2. this operation will NOT delete history of item, but you can't write data of this item to gdb any more
 * if you want to delete history data of item, please use /data/deleteFloatHistoricalData api
 * @param(String){groupName}
 * @param(String){condition}  condition is where clause in SQL
 */
axios.post(`${ip}/item/deleteItems`, {
    groupName: '5DCS',
    condition: `itemName like '%Bool%'`
}, configs).then(({data: {data}}) => {
    console.log(data)
    //  { effectedRows: 2, times: 426 }
})

/**
 * get items and itemCount of group from gdb
 *
 * @param(String){groupName}
 * @param(String){columnNames} columnNames should split with ',', if you want to get infos of all columns, you can use '*'
 * @param(String){condition} where clause in SQL
 * @param(number){startRow} pagination query parameters
 * @param(number){rowCount} pagination query parameters, if -1 ==> get all items
 */
axios.post(`${ip}/item/getItemsWithCount`, {
    groupName: '5DCS',
    columnNames: 'itemName, dataType',
    condition: '1=1',
    startRow: 0,
    rowCount: 10
}, configs).then(({data: {data}}) => {
    console.log(data)
    // {
    //   itemCount: 6,
    //   itemValues: [
    //     { dataType: 'int32', itemName: 'yInt' },
    //     { dataType: 'float32', itemName: 'xFloat' },
    //     { dataType: 'float32', itemName: 'yFloat' },
    //     { dataType: 'int32', itemName: 'xInt' },
    //     { dataType: 'string', itemName: 'yString' },
    //     { dataType: 'string', itemName: 'xString' }
    //   ]
    // }
})

/**
 * update items in group of gdb
 *
 * 1. only super_user has this route permission
 * 2. you can't update itemName column and dataType column
 * @param(String){groupName}
 * @param(String){condition} where clause in SQL
 * @param(String){clause} update clause in SQL
 */
axios.post(`${ip}/item/updateItems`, {
    groupName: '3DCS',
    condition: `itemName='X'`,
    clause: `descriptions='X',units=''`
}, configs).then(({data: {data}}) => {
    console.log(data)
    // { effectedRows: 1, times: 65 }
})

/**
 * check whether given items exist in group in gdb
 *
 * 1. only super_user or common_user have this route permission
 * 2. if item not existed, the returned code is 500, and message field record detail infos
 * @param(String){groupName}
 * @param(Array){itemNames}
 */
axios.post(`${ip}/item/checkItems`, {
    groupName: '5DCS',
    itemNames: ['xFloat', 'zFloat']
}, configs).then(({data: data}) => {
    console.log(data)
    // {"code":500,"message":"itemName:zFloat not existed","data":""}
})

/**
 * clean all items in group
 *
 * 1. only super_user has this route permission
 * 2. history data of item will be deleted
 *
 */
axios.post(`${ip}/item/cleanGroupItems`, {groupNames: ['3DCS']}, configs).then(({data: {data}}) => {
    console.log(data)
    // { effectedRows: 1, times: 0 }
})
