import axios from 'axios'
import VueCookies from 'vue-cookies';
import {
	Notification,
	MessageBox,
	Message,
	Loading
} from 'element-ui'
import router from '@/router';


const http = axios.create({
	// baseURL: 'http://127.0.0.1:8092',
	baseURL: 'http://43.139.13.115:8092',
	timeout: 100000,
	headers:{"Content-Type":"application/x-www-form-urlencoded"}
})
http.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded';
// http.defaults.headers.common["Authorization"] = $cookies.get("cookie")
// 添加请求拦截器
http.interceptors.request.use(function(config) {
	// 在发送请求之前做些什么
	config.headers["Authorization"] = $cookies.get("cookie")
	return config;
}, function(error) {
	// 对请求错误做些什么
	return Promise.reject(error);
});

// 添加响应拦截器
http.interceptors.response.use(function(response) {
	// 2xx 范围内的状态码都会触发该函数。
	// 对响应数据做点什么
	console.log(response)
	if (response.data.err_code == -2) {
		this.$notify.error({
			title: 'cookie过期',
			position: 'bottom-left',
			// type: 'success',
			duration: 1400,
			showClose: false,
		});
		router.push("/login")
	}
	return response;
}, function(error) {
	// 超出 2xx 范围的状态码都会触发该函数。
	// 对响应错误做点什么
	console.log(error)
	if (error.code == "ERR_NETWORK") {
		Message.error("网络错误 code:" + error.code + " message:" + error.message)
	}
	return Promise.reject(error);
});

export default http
