// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'  // Vue
import App from './App'  // App
import router from './router'  // vue-router路由
import store from '@/store'
import '@/permission'  // 权限控制
// import '@/assets/fonts/iconfont.css'
import VueQuillEditor from 'vue-quill-editor'
import 'quill/dist/quill.core.css'
import 'quill/dist/quill.snow.css'
import 'quill/dist/quill.bubble.css'
import '@/icons' // icon
import '@/mobile'
  
Vue.use(VueQuillEditor)
const echarts =  require('echarts')


// Vue.use(VueSocketio, 'http://192.168.0.199:8082')  // 连接socketio，默认的namespace
// 就是/

Vue.prototype.$echarts = echarts;

// ElementUi
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
Vue.use(ElementUI)

Vue.config.productionTip = false

const requireAll = requireContext => requireContext.keys().map(requireContext)
const req = require.context('@/icons/svg', true, /\.svg$/)
requireAll(req)  // 自动导入svg中的所有图pain

/* eslint-disable no-new */
new Vue({
  el: '#app',
  store,
  router,  // 注入路由,这样在其他的组件中就可以使用this.$router来访问路由器,this.$route访问当前路由
  render: h => h(App)
})
