/**
 creatTime: 2021/7/19
 creator: JustKeepSilence
 github: https://github.com/JustKeepSilence
 nodeVersion: 14.16.1
 **/

// gRCP examples for gdb of https mode, default sslTargetName is "gdb.com"

'use strict'
const grpc = require('grpc')
const Base64 = require('js-base64')
const fs = require('fs')
const async = require('async')

const ip = '192.168.0.103:8086'
const userName = 'admin'
const userToken = '7cb19c60ff6345c96aa723a406335389'  // you should login and then get the token
const token = 'Basic ' + Base64.encode(`${userName}:${userToken}`)
const cert = fs.readFileSync('../../ssl/gdbServer.crt')
const proto = grpc.load('../../nodeModel/gdb.proto')
const services = proto.model
const counts = 24 * 3600 * 10

const grpcRequest = async (path, data)=>{
    let client = null
    const option = { 'grpc.ssl_target_name_override': 'gdb.com', 'grpc.default_authority': 'gdb.com',  'grpc.max_receive_message_length': 1024 * 1024 * 1024}
    const credentials = grpc.credentials.createSsl(cert)
    const p = path.split('/')
    switch (p[1]) {
        case 'page':
            client = new services.Page(ip, credentials, option)
            break;
        case 'item':
            client = new services.Item(ip, credentials, option)
            break;
        case 'data':
            client = new services.Data(ip, credentials, option)
            break;
        case 'calculation':
            client = new services.Calc(ip, credentials, option)
            break;
        default:
            client = new services.Group(ip, credentials, option)
            break;
    }
    const metaData = new grpc.Metadata()
    metaData.add('Authorization', token)
    return new Promise((resolve, reject)=>{
        client[p[2]](data, metaData, (err, response) => {
            return err === null ? resolve({ data: response }) : reject({ message: err.details })
        })
    })
}

const streamWriteHistoryData = async (dataType, {groupNames, itemNames, itemValues, timeStamps})=>{
    return new Promise((resolve, reject)=>{
        const option = { 'grpc.ssl_target_name_override': 'gdb.com', 'grpc.default_authority': 'gdb.com',  'grpc.max_receive_message_length': 1024 * 1024 * 1024}
        const credentials = grpc.credentials.createSsl(cert)
        const client = new services.Data(ip, credentials, option)
        const metaData = new grpc.Metadata()
        metaData.add('Authorization', token)
        let call = null
        switch (dataType) {
            case "float32":
                call = client.batchWriteFloatHistoricalDataWithStream(metaData, (err, response)=>{
                    return err === null ? resolve({data:{response}}) : reject({message: err.details})
                })
                break
            case "int32":
                call = client.batchWriteIntHistoricalDataWithStream(metaData, (err, response)=>{
                    return err === null ? resolve({data:{response}}) : reject({message: err.details})
                })
                break
            case "string":
                call = client.batchWriteStringHistoricalDataWithStream(metaData, (err, response)=>{
                    return err === null ? resolve({data:{response}}) : reject({message: err.details})
                })
                break
            default:
                call = client.batchWriteBoolHistoricalDataWithStream(metaData, (err, response)=>{
                    return err === null ? resolve({data:{response}}) : reject({message: err.details})
                })
                break
        }
        function dataSender(groupNames, itemNames, itemValues, timeStamps){
            return function (callback){
                call.write({
                    groupNames,
                    itemNames,
                    itemValues,
                    timeStamps
                })
                callback(null, '')
            }
        }
        let fsr = []
        const step = 3600
        const c = itemValues[0].length / step // one hour
        for (let i = 0; i < c; i++) {
            let ts = []
            let values = []
            for (let j = 0; j < groupNames.length; j++) {
                values.push(sliceArray(i * step, (i + 1) * step, itemValues[j]))
                ts.push(sliceArray(i * step, (i + 1) * step, timeStamps[j]))
            }
            fsr[i] = dataSender(groupNames, itemNames, values, ts)
        }
        async.series(fsr, ()=>{
            call.end()
        })
    })
}

/**
 * mock float32 data ten days later
 */
function mockFloat32Data(now,coefficient){
    let timeStamps = []   // timeStamps
    let values = []   // values
    for (let i = 0; i < counts; i++) {
        const t = now.add(1, 'second')
        const ts = t.unix() + 8 * 3600   // timeStamp
        timeStamps.push(ts)
        if (t.second() >= 20 && t.second() <= 30){
            values.push(coefficient)   // write constant value
        }else{
            values.push(Math.random() * coefficient)
        }
    }
    return [timeStamps, values]
}

/**
 * mock int32 data ten days later
 */
function mockInt32Data(now,coefficient){
    let timeStamps = []   // timeStamps
    let values = []   // values
    for (let i = 0; i < counts; i++) {
        const t = now.add(1, 'second')
        const ts = t.unix() + 8 * 3600   // timeStamp
        timeStamps.push(ts)
        if (t.second() >= 20 && t.second() <= 30){
            values.push(coefficient)   // write constant value
        }else{
            values.push(Math.floor(Math.random() * (50 - 1 + 1) + 1) * coefficient)
        }
    }
    return [timeStamps, values]
}

/**
 * mock string data ten days later
 */
function mockStringData(now,coefficient){
    let timeStamps = []   // timeStamps
    let values = []   // values
    for (let i = 0; i < counts; i++) {
        const t = now.add(1, 'second')
        const ts = t.unix() + 8 * 3600   // timeStamp
        timeStamps.push(ts)
        if (t.second() >= 20 && t.second() <= 30){
            values.push(coefficient)   // write constant value
        }else{
            values.push(now.format('YYYY-MM-DD HH:mm:ss') + coefficient)
        }
    }
    return [timeStamps, values]
}

/**
 * mock bool data ten days later
 */
function mockBoolData(now, coefficient){
    let timeStamps = []   // timeStamps
    let values = []   // values
    for (let i = 0; i < counts; i++) {
        const t = now.add(1, 'second')
        const ts = t.unix() + 8 * 3600   // timeStamp
        timeStamps.push(ts)
        if (t.second() >= 20 && t.second() <= 30){
            values.push(coefficient)   // write constant value
        }else{
            values.push(i % 2 === 0 && coefficient)
        }
    }
    return [timeStamps, values]
}

function sliceArray(start, end, data) {
    if (start > end || end > data.length) {
        throw Error('invalid length')
    }
    let r = []
    for (let i = start; i < end; i++) {
        r.push(data[i])
    }
    return r
}

module.exports = {
    grpcRequest:grpcRequest,
    mockFloat32Data: mockFloat32Data,
    mockInt32Data: mockInt32Data,
    mockStringData:mockStringData,
    mockBoolData: mockBoolData,
    streamWriteHistoryData:streamWriteHistoryData
}
