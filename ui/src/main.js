import Vue from 'vue'
import App from './App.vue'
import router from './router'
//导入主样式
import './main.scss'

Vue.config.productionTip = false;

new Vue({
  router,
  render: h => h(App)
}).$mount('#app');
