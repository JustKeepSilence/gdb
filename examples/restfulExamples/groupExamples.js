/**
creatTime: 2021/7/13
creator: JustKeepSilence
github: https://github.com/JustKeepSilence
nodeVersion: 14.16.1
**/

// examples of handle groups in gdb by restful in node

/*
                                route permission of group handler
        url                         visitor         common_user             super_user
        addGroups                     ×                  √                       √
        deleteGroups                  ×                  √                       √
        getGroups                     √                  √                       √
        getGroupProperty              √                  √                       √
        updateGroupNames              ×                  ×                       √
        updateGroupColumnNames        ×                  ×                       √
        deleteGroupColumns            ×                  ×                       √
        addGroupColumns               ×                  ×                       √

        for group route permissions, you can't change in gdbClient
*/

'use strict'
const axios = require('axios')
const {configs} = require('./utils')
const {ip} = require('./utils')

/**
 * addGroups to gdb
 *
 * 1. you can't add duplicate group to gdb and the groupName can't be gdbKeyWords
 * 2. you MUST NOT add column of id, itemName and dataType because these three columns will be automatically added.
 * 3. only super_user and common_user can add groups to gdb
 * @param(Array){groupInfos} the groups info to be added to gdb
 * @param(String){groupName} the groupName to be added to gdb
 * @param(Array)(columnNames) the group columns to be add to gdb
 */
axios.post(`${ip}/group/addGroups`, {groupInfos: [{groupName: '3DCS', columnNames: ['description', 'unit']},{groupName:'4DCS', columnNames: []}]}, configs).then(({data:{data}})=>{
    console.log(data)
    // { effectedRows: 2, times: 1278 }
})

/**
 * delete groups from gdb
 *
 * 1. you can't delete calc group
 * 2. only super_user can delete groups from gdb
 * 3. this operation will also delete history data of group
 * @param(Array){groupNames} the groupNames to be deleted from gdb
 */
axios.post(`${ip}/group/deleteGroups`, {groupNames: ['3DCS']}, configs).then(({data:{data}})=>{
    console.log(data)
    // { effectedRows: 1, times: 1278 }
})

/**
 * get all groups in gdb
 */
axios.post(`${ip}/group/getGroups`, {}, configs).then(({data: {data}})=>{
    console.log(data)
    // { groupNames: [ 'calc', '4DCS' ] }
})

/**
 * get columnNames and itemCount of the given group
 *
 * @param(String){groupName}
 * @param(String){condition} the condition of where clause in SQL
 */
axios.post(`${ip}/group/getGroupProperty`, {groupName: '4DCS', condition: '1=1'}, configs).then(({data: {data}})=>{
    console.log(data)
    // { itemCount: '0', itemColumnNames: [ 'itemName', 'dataType' ] }
})

/**
 * update groupNames in gdb
 *
 * 1. you can't update calc groupName
 * 2. only super_user has this route permission
 * 3. history data will migrate as well.
 * @param(Array){infos}  infos of updatedGroupNames
 * @param(String){oldGroupName}
 * @param(String){newGroupName}
 */
axios.post(`${ip}/group/updateGroupNames`, {infos:[{oldGroupName: '4DCS', newGroupName: '5DCS'}]}, configs).then(({data: {data}})=>{
    console.log(data)
    // { effectedRows: 1, times: 502 }
})

/**
 * update columnNames in group in gdb
 *
 * 1. you can't update id, itemName, dataType columnName
 * 2. only super_user has this route permission
 * @param(String){groupName}
 * @param(Array){oldColumnNames}
 * @param(Array){newColumnNames}
 */
axios.post(`${ip}/group/updateGroupColumnNames`, {groupName: '3DCS', oldColumnNames: ['description', 'unit'], newColumnNames: ['descriptions', 'units']}, configs).then(({data: {data}})=>{
    console.log(data)
    // { effectedCols: 2, times: 103 }
})

/**
 * delete columns from groups in gdb
 *
 * 1. you can't delete id, itemName, dataType columnName
 * 2. only super_user has this route permission
 * @param(String){groupName}
 * @param(Array){columnNames}
 */
axios.post(`${ip}/group/deleteGroupColumns`, {groupName: '3DCS', columnNames: ['units']}, configs).then(({data: {data}})=>{
    console.log(data)
    // { effectedCols: 1, times: 38 }
})

/**
 * add group columns to group in gdb
 *
 * 1. you can't add duplicate columns to group
 * 2. dataType of column in gdb is text
 * 3. you must specify default value for added columns
 * 4. only super_user has this route permission
 * @param(String){groupName}
 * @param(Array){columnNames}
 * @param(Array){defaultValues}
 */
axios.post(`${ip}/group/addGroupColumns`, {groupName: '3DCS', columnNames: ['units'], defaultValues:['m/s']}, configs).then(({data: {data}})=>{
    console.log(data)
    // { effectedCols: 1, times: 122 }
})
