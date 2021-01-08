/**
 * index js
 */

import request from '@/utils/request'
import { getCookie } from '@/utils/cookie'

const getDbInfo = (data)=>{
    return request({
        url: "http://" + getCookie('ip') +"/data/getDbInfo",
        method: "post",
        data
    })
}

const getSpeedHistory = (data)=>{
    return request({
        url: "http://" + getCookie('ip') +"/data/getDbSpeedHistory",
        method: "post",
        data
    })
}

export{
    getDbInfo,
    getSpeedHistory
}