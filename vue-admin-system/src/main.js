import Vue from 'vue'
import App from './App.vue'
import ElementUI from 'element-ui'
import VueClipBoard from 'vue-clipboard2'
import 'element-ui/lib/theme-chalk/index.css'
import '../src/assets/element-#58B758/index.css'
import router from './router'
import store from './store'
import VueCookies from 'vue-cookies';
import http from '../src/utils/request.js'
import './assets/iconfont/iconfont.js'

Vue.config.productionTip = false
Vue.use(ElementUI)
Vue.use(VueCookies)
Vue.use(VueClipBoard)
// 在这里挂载
new Vue({
	router,
	store,
	render: h => h(App),
}).$mount('#app')
