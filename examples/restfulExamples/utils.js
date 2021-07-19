const Base64 = require('js-base64')
const ip = 'http://192.168.0.129:8082'
const userName = 'admin'
const userToken = '7cb19c60ff6345c96aa723a406335389'  // you should login and then get the token
const token = 'Basic ' + Base64.encode(`${userName}:${userToken}`)
const configs = {headers: {Authorization: token}, maxBodyLength: Infinity}
const counts = 24 * 3600 * 10
// if run gdbService with authorization, you MUST use token

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

module.exports = {
    configs: configs,
    ip: ip,
    mockFloat32Data: mockFloat32Data,
    mockInt32Data: mockInt32Data,
    mockStringData:mockStringData,
    mockBoolData: mockBoolData,
}