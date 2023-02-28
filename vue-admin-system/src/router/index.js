import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import User from '../views/User.vue'
import Main from '../views/Main.vue'
import Mall from '../views/Mall.vue'
import Manage from '../views/Manage.vue'
import Settings from '../views/Settings.vue'
import PageTwo from '../views/PageTwo.vue'
import Login from '../views/Login.vue'
import http from '../utils/request.js'
import {
	Notification,
	MessageBox,
	Message,
	Loading
} from 'element-ui'
Vue.use(VueRouter)


// 0. 如果使用模块化机制编程，导入Vue和VueRouter，要调用 Vue.use(VueRouter)

// 1. 定义 (路由) 组件。
// 可以从其他文件 import 进来


// 2. 定义路由
// 每个路由应该映射一个组件。 其中"component" 可以是
// 通过 Vue.extend() 创建的组件构造器，
// 或者，只是一个组件配置对象。
// 我们晚点再讨论嵌套路由。
const routes = [{
		path: '/',
		component: Main,
		redirect: '/home', //访问/重定向
		meta: {
			title: '首页'
		},

		children: [
			// 子路由
			{
				path: '/home',
				component: Home,
				meta: {
					title: '欢迎'
				}
			},
			{
				path: '/user',
				component: User,
				meta: {
					title: '录入'
				}
			},
			{
				path: '/mall',
				component: Mall,
				meta: {
					title: '报表'
				}
			},
			{
				path: '/manage',
				component: Manage,
				meta: {
					title: '管理'
				}
			},
			{
				path: '/settings',
				component: Settings,
				meta: {
					title: '管理'
				}
			},
			{
				path: '/page2',
				component: PageTwo,
				meta: {
					title: '页面2'
				}
			},

		]
	},
	{
		path: '/login',
		component: Login,
		meta: {
			title: '欢迎'
		}
	},

]

// 3. 创建 router 实例，然后传 `routes` 配置
// 你还可以传别的配置参数, 不过先这么简单着吧。
const router = new VueRouter({
	mode: 'history',
	routes // (缩写) 相当于 routes: routes

})

// 鉴权
router.beforeEach((to, from, next) => {
	console.log("to", to)
	console.log("from", from)

	const cookie = $cookies.get("cookie")
	// http.post("login/dologin", {
	// 	userName: this.username,
	// 	passWord: this.password
	// }).then((response) => {
	// 	if (response.data.err_code != 0) {
	// 		this.$message.error(response.data.err_msg);
	// 		this.loading = false
	// 	} else {
	// 		this.$message.success("登录成功");
	// 		this.$cookies.set("cookie", response.data.data.cookie, response.data.data.outTime)
	// 		this.$router.push({
	// 			path: "/home",

	// 		});
	// 		this.loading = false
	// 	}
	// })

	console.log(cookie)


	if (to.path === '/login' || to.path === '/home') {
		next()
	} else {
		if (cookie) {
			http.post("/login/notOutTime", {
				cookie: cookie,
			}).then(response => {
				console.log(response.data.data.notOutTime)
				if (response.data.data.notOutTime == true) {
					// 没有过期
					next();
				} else {
					// // 过期
					// this.$notify({
					// 	title: '登录过期',
					// 	position: 'bottom-left',
					// 	type: 'success',
					// 	duration: 1400,
					// 	showClose: false,
					// });

					$cookies.remove("cookie")
					next('/login')
					
				}
			})
		} else {
			
			next('/login')
		}
	}



})

// 4. 创建和挂载根实例。
// 记得要通过 router 配置参数注入路由，
// 从而让整个应用都有路由功能
export default router
// 现在，应用已经启动了！
