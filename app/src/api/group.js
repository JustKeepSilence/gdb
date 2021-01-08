/**
 * 和group相关的请求
 */

import request from '@/utils/request'
import { getCookie } from '@/utils/cookie'

const getGroups = ()=>{
    return request({
        url: "http://" + getCookie('ip') +"/group/getGroups",
        method: "post",
    })
}

const addGroups = (data)=>{
    return request({
        url: "http://" + getCookie('ip') +'/group/addGroups',
        method: 'post',
        data
    })
}

const getGroupColumns = (data)=>{
    return request({
        url: "http://" + getCookie('ip') +'/group/getGroupProperty',
        method: 'post',
        data
    })
}

const addItem = (data)=>{
    return request({
        url: "http://" + getCookie('ip') +'/item/addItems',
        method: 'post',
        data
    })
}

const deleteGroup = (data)=>{
    return request({
        url: "http://" + getCookie('ip') +'/group/deleteGroups',
        method: 'post',
        data
    })
}

const getItems = (data)=>{
    return request({
        url : "http://" + getCookie('ip') +"/page/getItemsWithCount",
        method: 'post',
        data
    })
}

const getRealTimeData = (data)=>{
    return request({
        url: "http://" + getCookie('ip') +'/data/getRealTimeData',
        method: 'post',
        data
    })
}

const addItemsByExcel = (data)=>{
    return request({
        url: "http://" + getCookie('ip') +'/page/addItemsByExcel',
        method:'post',
        data
    })
}

const getHistoryData = (data)=>{
    return request({
        url :"http://" + getCookie('ip') +'/data/getHistoricalData',
        method: 'post',
        data
    })
}

const updateColumnNames = (data)=>{
    return request({
        url: "http://" + getCookie('ip') +'/group/updateGroupColumnNames',
        method: 'post',
        data
    })
}

const addColumns = (data)=>{
    return request({
        url: "http://" + getCookie('ip') +'/group/addGroupColumns',
        method: 'post',
        data
    })
}

const deleteColumns = (data)=>{
    return request({
        url: "http://" + getCookie('ip') +'/group/deleteGroupColumns',
        method: 'post',
        data
    })
}

const updateItems = (data)=>{
    return request({
        url: "http://" + getCookie('ip') +'/item/updateItems',
        method: 'post',
        data
    })
}

const deleteItems = (data)=>{
    return request({
        url: "http://" + getCookie('ip') +'/item/deleteItems',
        method: 'post',
        data
    })
}


export{
    getGroups,
    addGroups,
    getGroupColumns,
    addItem,
    deleteGroup,
    getItems,
    getRealTimeData,
    addItemsByExcel,
    getHistoryData,
    updateColumnNames,
    addColumns,
    deleteColumns,
    updateItems,
    deleteItems,
}
