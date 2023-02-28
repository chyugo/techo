<template>

	<div class="loginContainer" v-loading="loading" element-loading-background="rgba(0, 0, 0, 0.4)">

		<form action="" class="loginForm">
			<div style="align-self: center;">

				<transition name="el-fade-in-linear">
					<div v-show="show">
						<el-input id="username" v-model="username" placeholder="用户名" ></el-input>
						<el-input id="niao" placeholder="密码" v-model="password" show-password></el-input>
						<el-button plain class="loginbtn" @click="dologin">登录</el-button>
						<div style="margin-top:20px">
							<el-button type="text" @click="dialogFormVisible=true" style="color: white;">
								没有账号，去注册<span class="iconfont"
									style="font-weight: 700;font-size: 14px;margin-left: 2px;">&#xe674;</span>
							</el-button>
						</div>
						<el-button type="text" @click="useShare" style="color: white;">
							使用体验账户
						</el-button>

					</div>
				</transition>



			</div>
		</form>
		<el-dialog title="注册" :visible.sync="dialogFormVisible" width="40%" :close-on-click-modal=false>
			<el-form>
				<el-form-item label="用户名">
					<el-input v-model="registName" autocomplete="off" class="registinput"></el-input>
				</el-form-item>
				<el-form-item label="密码">
					注意：不要设置成常用的社交密码！
					<el-input v-model="password1" autocomplete="off" class="registinput"></el-input>
				</el-form-item>
				<el-form-item label="再输入一次密码">
					<el-input v-model="password2" autocomplete="off" class="registinput"></el-input>
				</el-form-item>
			</el-form>
			<div slot="footer" class="dialog-footer">
				<el-button @click="dialogFormVisible = false">取 消</el-button>
				<el-button type="primary" @click="doRegist">确 定</el-button>
			</div>
		</el-dialog>

	</div>
</template>

<script>
	import http from '../utils/request.js'

	export default {
		name: 'Login',
		data() {
			return {
				radio1: '我要兼职',
				msg: 'Welcome to Your Vue.js App',
				username: "",
				password: "",
				show: false,
				loading: false,
				dialogFormVisible: false,
				registName: '',
				password1: '',
				password2: '',
				input1:'',
				input2:'',
			}
		},
		methods: {
			dologin() {
				this.loading = true
				http.post("login/dologin", {
					userName: this.username,
					passWord: this.password
				}, {
					timeout: 18000
				}).then((response) => {
					if (response.data.err_code != 0) {
						this.$notify.error({
							title: response.data.err_msg,
							position: 'bottom-left',
							duration: 1400,
							showClose: false,
						});
						this.loading = false
					} else {
						this.$notify({
							title: '登录成功',
							position: 'bottom-left',
							type: 'success',
							duration: 1400,
							showClose: false,
						});
						this.$cookies.set("cookie", response.data.data.cookie, response.data.data.outTime)
						this.$router.push({
							path: "/home",

						});
						this.loading = false
					}
				}).catch(reason => {
					this.loading = false
				})

			},
			doRegist() {
				http.post("login/doRegist", {
					registName: this.registName,
					password1: this.password1,
					password2: this.password2,
				}, {
					timeout: 18000
				}).then((response) => {
					if (response.data.err_code != 0) {
						this.$message.error(response.data.err_msg);
						this.loading = false
						if (response.data.err_code == -3) {
							this.dialogFormVisible = true
						} else {

						}

					} else {
						this.dialogFormVisible = false
						this.$notify({
							title: "登陆成功",
							position: 'bottom-left',
							type: 'success',
							duration: 1400,
							showClose: false,
						});
						this.$cookies.set("cookie", response.data.data.cookie, response.data.data.outTime)
						this.$router.push({
							path: "/home",

						});
						this.loading = false
					}
				}).catch(reason => {
					this.loading = false
				})
			},
			useShare(){
				
				this.$notify({
					title: "用户名和密码均是share",
					position: 'bottom-left',
					type: 'success',
					duration: 1400,
					showClose: false,
				});
			}
		},
		mounted() {
			this.show = true
			const cookie = this.$cookies.get("cookie")
			if (cookie != '') {
				this.$router.push({
					path: "/mall",
				});
			}
		},
		
	}
</script>

<style>
	* {
		/* user-select: none; */

	}

	@media screen and (max-width:1000px) {
		.el-dialog {
			width: 90% !important;
		}

	}


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

	.registinput .el-input__inner {
		background-color: white !important;
		color: black !important;
	}

	.loginContainer {
		height: 100vh;
		background: url("../assets/1.jpeg") no-repeat;
		background-size: cover;
		background-attachment: fixed;
	}

	.loginContainer .loginForm {
		position: absolute;
		top: 50%;
		left: 50%;
		margin-top: -250px;
		margin-left: -250px;
		width: 500px;
		height: 500px;
		border-radius: 30px;
		backdrop-filter: blur(25px);
		text-align: center;
		display: flex;

	}

	.loginContainer .loginForm .title {
		font-size: 24px;
		color: black;

	}

	.loginContainer .loginForm .el-input {
		width: 60%;
		font-size: 16px;
		font-weight: 600;
		margin-bottom: 10px;


	}



	.loginContainer .loginForm .loginbtn {
		width: 60%;
		font-size: 16px;
		font-weight: 600;
		margin-top: 5px;
		background-color: rgba(0, 0, 0, 0.2);
		color: white;
	}

	.loginContainer .loginForm .regist {

		font-size: 16px;
		font-weight: 600;
		margin-top: 5px;
		color: white;
		margin-top: 30px;

	}



	.loginContainer .loginForm .loginbtn:hover {
		background-color: rgba(0, 0, 0, 0.2);

		color: rgba(255, 255, 255, 0.8) !important;
	}


	.loginContainer .loginForm .registbtn:hover {
		background-color: rgba(0, 0, 0, 0.2);

		color: rgba(255, 255, 255, 0.8) !important;
	}

	.loginContainer .el-radio-button__inner {
		margin-bottom: 20px;
		font-size: 14px !important;
		padding: 12px 45px !important;
		background-color: rgba(0, 0, 0, 0.2);
		color: white;
		border-left: 1px solid #449231 !important;
		border: 1px solid #449231;
	}

	.loginContainer .el-button.is-plain:focus {
		background-color: rgba(0, 0, 0, 0.2) !important;
		color: white;
	}

	.loginContainer .el-button.is-plain:active {
		background-color: rgba(0, 0, 0, 0.2) !important;
		color: white;
	}

	.loginContainer .el-button.is-plain:hover {
		color: white;
	}

	.loginContainer .el-input__inner {
		background-color: rgba(0, 0, 0, 0.2);
		color: white;
	}

	.loginContainer .choice {
		margin-top: 60px;
	}

	.loginContainer .el-radio-button__inner:hover {
		color: white;
	}

	.loginContainer .regist a {
		cursor: pointer;
		color: white;
		text-decoration: none;
	}
</style>
