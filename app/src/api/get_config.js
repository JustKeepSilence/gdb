/**
 * 获取json配置文件
 */

 import axios from 'axios'

 const request = axios.create({
    baseURL: '',
    timeout: 5000,
    headers: {'Content-type': 'multipart/form-data'}
 })

 const getConfig = request.get('E:\\go\\gdb\\app\\package.json').then(res => {
    console.log("获取本地的json文件" + JSON.stringify(res));
})

export default request