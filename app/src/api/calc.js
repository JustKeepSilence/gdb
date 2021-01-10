import request from '@/utils/request'
import { getCookie } from '@/utils/cookie'


const addCalulationItems = (data)=>{
    return request({
        url: "http://" + getCookie('ip') +'/calculation/addCalcItem',
        method: 'post',
        data
    })
}

const getCalculationItems = (data)=>{
    return request({
        url: "http://" + getCookie('ip') +'/calculation/getCalcItem',
        method:'post',
        data
    })
}

const updateCalculationItems = (data)=>{
    return request({
        url : "http://" + getCookie('ip') +'/calculation/updateCalcItem',
        method: 'post',
        data
    })
}

const getReCalc = (data)=>{
    return request({
        url: "http://" + getCookie('ip') +'/calculation/getReCalc',
        method: "post",
        data
    })
}

export{
    addCalulationItems,
    getCalculationItems,
    updateCalculationItems,
    getReCalc
}