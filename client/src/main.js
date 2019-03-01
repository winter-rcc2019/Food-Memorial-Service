import Vue from 'vue'
import ElementUI from 'element-ui'
import locale from 'element-ui/lib/locale/lang/ja'
import 'element-ui/lib/theme-chalk/index.css'
import App from './App.vue'
import Top from './Top.vue'
import Insta from './Insta.vue'


Vue.use(ElementUI, {locale})
Vue.config.productionTip = false

new Vue({
  render: h => h(App),
}).$mount('#app')

new Vue({
  render: h => h(Top),
}).$mount('#toptop')
new Vue({
  render: h => h(Insta),
}).$mount('#inin')