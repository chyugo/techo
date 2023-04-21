<template>
	<div class="header-container">
		<div class="l-content hidden-sm-and-down">
			<el-button icon="el-icon-menu" @click="handleMenu" size="mini"></el-button>
			<span class="text">
				<el-breadcrumb separator-class="el-icon-arrow-right">
					<el-breadcrumb-item v-for="item in breadcrumbList" :key="item.path" :to="{ path: item.path }">
						{{item.meta.title}}
					</el-breadcrumb-item>
				</el-breadcrumb>
			</span>
		</div>
		
		<div class="l-content2 hidden-md-and-up">
			<el-button type="text" class="topbutton" @click="tohome" size="mini">首页</el-button>
			<el-button type="text" class="topbutton" @click="tomall" size="mini">报表</el-button>
			<el-button type="text" class="topbutton" @click="touser" size="mini">录入</el-button>
			<el-button type="text" class="topbutton" @click="tomanage" size="mini">管理</el-button>
		</div>

		<div class="r-content ">

			<el-dropdown @command="handleCommand">
				<span class="el-dropdown-link">
					<img class="user" src="../assets/user.png" alt="">
				</span>
				<el-dropdown-menu slot="dropdown">
					<el-dropdown-item command="userSettings"><span class="iconfont"
							style="font-weight: 300;">&#xe602;</span> 个人中心</el-dropdown-item>
					<el-dropdown-item command="export"><span class="iconfont" style="font-weight: 300;">&#x100c5;</span>
						导出所有数据为Excel</el-dropdown-item>
					<el-dropdown-item command="logOut"><span class="iconfont" style="font-weight: 300;">&#xe668;</span>
						退出</el-dropdown-item>
				</el-dropdown-menu>
			</el-dropdown>
		</div>
	</div>
</template>

<script>
	import 'element-ui/lib/theme-chalk/display.css';
	import http from '../utils/request.js'
	export default {
		data() {
			return {

			}
		},
		methods: {

			handleMenu() {
				this.$store.commit('collapaseMenu')
			},
			tomall() {
				if (this.$route.path !== "/mall") {
					this.$router.push("/mall").catch(err => {})
				}
			},
			touser() {
				if (this.$route.path !== "/user") {
					this.$router.push("/user").catch(err => {})
				}
			},
			tomanage() {
				if (this.$route.path !== "/manage") {
					this.$router.push("/manage").catch(err => {})
				}
			},
			tohome() {
				if (this.$route.path !== "/home") {
					this.$router.push("/home").catch(err => {})
				}
			},
			handleCommand(command) {
				if (command == "logOut") {
					this.$cookies.remove("cookie")
					this.$router.push("/login").catch(err => {})
				}
				if (command == "export") {
					http.post("/export", {}, {
						responseType: 'arraybuffer'
					}).then((res) => {
						// 假设 data 是返回来的二进制数据
						const data = res.data
						const url = window.URL.createObjectURL(new Blob([data], {
							type: "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
						}))
						const link = document.createElement('a')
						link.style.display = 'none'
						link.href = url
						link.setAttribute('download', 'excel.xlsx')
						document.body.appendChild(link)
						link.click()
						document.body.removeChild(link)
					}).catch(err => {
						console.log(err)
					})
				}
				if (command == "userSettings") {
					this.$router.push("/settings").catch(err => {})
				}
			}
		},
		created() {
			console.log(this.$route)
		},
		computed: {
			breadcrumbList() {
				return this.$route.matched
			}
		},
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
	

	
	.header-container {
		padding: 0 20px;
		// background-color: #333;
		background-color: rgb(84, 92, 100);
		background-image: linear-gradient(95deg, #dbfbca, #f2da94 35%, #b2bc7b 67%, #7aac9a);

		height: 60px;
		display: flex;
		// 主轴两端显示
		justify-content: space-between;
		// 纵轴垂直居中
		align-items: center;

		.text {
			color: black;
			font-size: 14px;
			margin-left: 10px;
		}

		.r-content {
			.user {
				width: 30px;
				height: 30px;
				border-radius: 50%;
			}
		}

		.l-content,.l-content2{
			display: flex;
			align-items: center;
			justify-content: space-evenly;

			.topbutton {
				color: black;
				font-size: 13px;
				font-weight: 700;
				margin-right: 10px;
			}
		}
		.l-content2{
			width: 100%;
			justify-content: center;
			button{
				position: relative;
			}
			button:not(:last-child)::after{
				content: "|";
				position: absolute;
				right: -12px;
				top: 6px;
			}
		}

		/deep/ .el-breadcrumb__inner {
			color: black;
		}
	}
</style>
