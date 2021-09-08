import Vue from 'vue'
import App from './App.vue'
import router from './router'
import './plugins/element.js'

import axios from 'axios'
// 配置请求的根路径
// 因为照成了跨域问题，所以开发者模式暂时不用这里了
// axios.defaults.baseURL = ''

Vue.prototype.$http = axios

Vue.config.productionTip = false

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
