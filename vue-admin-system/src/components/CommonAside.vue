<template>
	<div>
		<el-menu default-active="1-4-1" class="el-menu-vertical-demo hidden-sm-and-down" @open="handleOpen"
			@close="handleClose" :collapse="isCollapse" background-color="#545c64" text-color="#fff"
			active-text-color="#ffd04b">
			<h3>{{isCollapse ? '后台': '后台管理系统' }}</h3>
			<el-menu-item  @click="clickMenu(item)" v-for="item in noChildren" :key="item.name" :index="item.name">
				<i :class="`el-icon-${item.icon}`"></i>
				<!-- 模板字符串语法 -->
				<span slot="title" >{{item.label}}</span>
			</el-menu-item>
			<el-submenu v-for="item in hasChildren" :key="item.label" :index="item.label" >
				<template slot="title">
					<i :class="`el-icon-${item.icon}`"></i>

				</template>
				<el-menu-item-group v-for="subItem in item.children" :key="item.path" :index="item.path" >

					<el-menu-item @click="clickMenu(subItem)" :index="subItem.path">{{subItem.label}}</el-menu-item>
				</el-menu-item-group>

			</el-submenu>


		</el-menu>


	</div>
</template>

<script>
	import 'element-ui/lib/theme-chalk/display.css';
	export default {
		data() {
			return {

				menuData: [{
						path: '/',
						name: 'home',
						label: '首页',
						icon: 's-home',
						url: 'Home/Home',
					},
					{
						path: '/mall',
						name: 'mall',
						label: '报表',
						icon: 's-order',
						url: 'MallManage/MallManage',
					},
					{
						path: '/user',
						name: 'user',
						label: '录入数据',
						icon: 'edit',
						url: 'UserManage/UserManage',
					},
					{
						path: '/manage',
						name: 'manage',
						label: '管理',
						icon: 'setting',
						url: 'Manage',
					},
					// {
					// 	label: '其他',
					// 	icon: 'location',
					// 	children: [{
					// 		path: '/page2',
					// 		name: 'page2',
					// 		label: '页面2',
					// 		icon: 'setting',
					// 		url: 'Other/pageTwo',
					// 	}]
					// }
				],
			};
		},
		methods: {
			handleOpen(key, keyPath) {
				console.log(key, keyPath);
			},
			handleClose(key, keyPath) {
				console.log(key, keyPath);
			},
			clickMenu(item) {
				console.log(item)

				if (this.$route.path !== item.path && !(this.$route.path === '/home' && item.path === '/')) {
					this.$router.push(item.path).catch(err => {})
				}

			}
		},
		computed: {
			// 无子菜单
			noChildren() {
				return this.menuData.filter(item => !item.children)
			},
			// 有子菜单	 
			hasChildren() {
				return this.menuData.filter(item => item.children)
			},
			isCollapse() {
				return this.$store.state.tab.isCollapse
			}
		}
	}
</script>

<style lang="less" scoped>
	@font-face {
		font-family: 'iconfont';
		src: url('../assets/iconfont/iconfont.ttf') format('truetype');
	}
	.iconfont {
		font-family: "iconfont" !important;
		font-size: 16px;
		font-style: normal;
		-webkit-font-smoothing: antialiased;
		-moz-osx-font-smoothing: grayscale;
	}
	
	.el-menu-vertical-demo:not(.el-menu--collapse) {
		width: 200px;
		min-height: 400px;
	}

	.el-menu {
		height: 100vh;
		border-right: 0;

		h3 {
			color: #fff;
			text-align: center;
			line-height: 48px;
			font-size: 16px;
			font-weight: 400;
		}
	}

</style>
