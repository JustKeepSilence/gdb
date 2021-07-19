/**
 creatTime: 2021/7/19
 creator: JustKeepSilence
 github: https://github.com/JustKeepSilence
 nodeVersion: 14.16.1
 **/

// gRPC examples of item handler in gdb, about parameter info, you can see:
// https://github.com/JustKeepSilence/gdb/blob/master/examples/restfulExamples/itemExamples.js

'use strict'
const {grpcRequest} = require('./utils')

// add items to group in gdb, NOTE: map in proto file can't be repeated, so itemValues MUST be string
grpcRequest('/item/addItems', {
    groupName: '5DCS',
    itemValues: JSON.stringify([{itemName: 'xFloat', dataType: 'float32'}, {itemName: 'yFloat', dataType: 'float32'},
        {itemName: 'xInt', dataType: 'int32'}, {itemName: 'yInt', dataType: 'int32'}, {
            itemName: 'xString',
            dataType: 'string'
        },
        {itemName: 'yString', dataType: 'string'}, {itemName: 'xBool', dataType: 'bool'}, {
            itemName: 'yBool', dataType: 'bool'
        }])
}).then(({data}) => {
    console.log(data)
}).catch(({message}) => {
    console.log(message)
})

// delete item from group in gdb
grpcRequest('/item/deleteItems', {
    groupName: '5DCS',
    condition: `itemName like '%Bool%'`
}).then(({data}) => {
    console.log(data)
    // { effectedRows: 2, times: '135' }
}).catch(({message}) => {
    console.log(message)
})

// get items and itemCount of group from gdb
grpcRequest('/item/getItemsWithCount', {
    groupName: '5DCS',
    columnNames: 'itemName, dataType',
    condition: '1=1',
    startRow: 0,
    rowCount: 10
}).then(({data}) => {
    console.log(JSON.stringify(data))
}).catch(({message}) => {
    console.log(message)
})

// update items in group of gdb
grpcRequest('/item/updateItems', {
    groupName: '4DCS',
    condition: `itemName='X'`,
    clause: `descriptions='X',units=''`
}).then(({data}) => {
    console.log(data)
    // { effectedRows: 1, times: '125' }
}).catch(({message}) => {
    console.log(message)
})

// check whether given items exist in group in gdb
grpcRequest('/item/checkItems', {groupName: '5DCS', itemNames: ['xFloat', 'zFloat']}).then(({data}) => {
    console.log(data)
}).catch(({message}) => {
    console.log(message)
    // itemName: zFloat not existed
})

// clean all items in group
grpcRequest('/item/cleanGroupItems', {groupNames: ['4DCS', '5DCS']}).then(({data}) => {
    console.log(data)
    // { effectedRows: 2, times: '961' }
}).catch(({message}) => {
    console.log(message)
})
