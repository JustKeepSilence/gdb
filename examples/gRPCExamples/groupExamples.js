/**
 creatTime: 2021/7/19
 creator: JustKeepSilence
 github: https://github.com/JustKeepSilence
 nodeVersion: 14.16.1
 **/

// gRPC examples of group handler in gdb, about parameter info, you can see:
// https://github.com/JustKeepSilence/gdb/blob/master/examples/restfulExamples/groupExamples.js


'use strict'
const {grpcRequest} = require('./utils')

// add groups to gdb
grpcRequest('/group/addGroups', {
    groupInfos: [{
        groupName: '3DCS',
        columnNames: ['description', 'unit']
    }, {groupName: '5DCS', columnNames: []}]
}).then(({data}) => {
    console.log(data)
    // { effectedRows: 2, times: '794' }
}).catch(({data: {message}}) => {
    console.log('err: ' + message)
})

// delete groups from gdb
grpcRequest('/group/deleteGroups', {groupNames: ['3DCS', '5DCS']}).then(({data}) => {
    console.log(data)
    //{ effectedRows: 2, times: '391' }
}).catch(({data: {message}}) => {
    console.log('err: ' + message)
})

// get groups in gdb
grpcRequest('/group/getGroups', {}).then(({data}) => {
    console.log(data)
}).catch(({data: {message}}) => {
    console.log(message)
})

// get columnNames and itemCount of the given group
grpcRequest('/group/getGroupProperty', {groupName: '3DCS', condition: '1=1'}).then(({data}) => {
    console.log(data)
    //{
    //   itemCount: 0,
    //   itemColumnNames: [ 'itemName', 'dataType', 'description', 'unit' ]
    // }
}).catch(({data: {message}}) => {
    console.log(message)
})

// update groupNames in gdb
grpcRequest('/group/updateGroupNames', {
    infos: [{
        oldGroupName: '3DCS',
        newGroupName: '4DCS'
    }]
}).then(({data}) => {
    console.log(data)
    // { effectedRows: 1, times: '475' }
}).catch(({data: {message}}) => {
    console.log('err: ' + message)
})

// update columnNames in group in gdb
grpcRequest('/group/updateGroupColumnNames', {
    groupName: '4DCS',
    oldColumnNames: ['description', 'unit'],
    newColumnNames: ['descriptions', 'units']
}).then(({data}) => {
    console.log(data)
    // { effectedCols: 2, times: '163' }
}).catch(({data: {message}}) => {
    console.log('err: ' + message)
})

// delete columns from groups in gdb
grpcRequest('/group/deleteGroupColumns', {
    groupName: '4DCS',
    columnNames: ['units']
}).then(({data}) => {
    console.log(data)
    // { effectedCols: 1, times: '46' }
}).catch(({data: {message}}) => {
    console.log('err: ' + message)
})

// add group columns to group in gdb
grpcRequest('/group/addGroupColumns', {
    groupName: '4DCS',
    columnNames: ['units'],
    defaultValues: ['m/s']
}).then(({data}) => {
    console.log(data)
    // { effectedCols: 1, times: '125' }
}).catch(({data: {message}}) => {
    console.log(message)
})
