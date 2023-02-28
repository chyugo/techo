<template>
	<div>
		<div class="info">
			<p class="access">用户中心</p>
			<el-button class="infoButton" @click="dialogShow">修改密码</el-button>
			<el-dialog :visible.sync="dialogVisible" width="40%" :show-close=false :close-on-click-modal=false>
				<el-form>
					<el-form-item label="请输入新密码">
						<el-input v-model="password1" autocomplete="off"></el-input>
					</el-form-item>
					<el-form-item label="再输入一次新密码">
						<el-input v-model="password2" autocomplete="off"></el-input>
					</el-form-item>
				</el-form>
				<el-button type="primary" @click="doResetPassword">提交</el-button>
				<el-button type="text" @click="closeAbout">关闭</el-button>
			</el-dialog>
		</div>
	</div>
</template>

<script>
	import http from '../utils/request.js'
	import {
		getData
	} from '../api'
	import * as echarts from 'echarts'

	export default {
		data() {
			return {
				password1: '',
				password2: '',
				dialogVisible: false,

			}
		},
		methods: {
			dialogShow() {
				this.dialogVisible = true
			},
			closeAbout() {
				this.dialogVisible = false
			},
			doResetPassword() {
				http.post("user/resetPassword", {
					passWord1: this.password1,
					passWord2: this.password2,
				}, {

				}).then((response) => {
					if (response.data.err_code != 0) {
						this.$message.error(response.data.err_msg);
						this.loading = false
					} else {
						this.$message.success("修改成功");
						this.loading = false
					}
				}).catch(reason => {
					this.loading = false
				})
			}
		},
		mounted() {


		}

	}
</script>

<style lang="less" scoped>
	@media screen and (max-width:1000px) {
		.info {
			padding: 0px 20px;
		}

		/deep/.el-dialog {
			width: 90% !important;
		}

	}

	/deep/.el-dialog {
		border-radius: 10px;
	}

	.strong {
		font-weight: 700;
	}

	.info {
		margin-top: 110px;
		margin-bottom: 90px;
		text-align: center;
		color: rgb(27, 28, 29);
		font-size: 16px;
		font-weight: 700;
		line-height: 25px;

		.infoButton {
			margin-top: 30px;
			margin-bottom: 20px;
			width: 300px;
			height: 50px;
			border-radius: 10px;
		}

		.about {
			color: black;
			font-weight: 400;
			font-size: 14px;
		}
	}



	.dev {
		.devTop {
			width: 100%;
			font-size: 20px;
			text-align: center;
			margin: 30px 0px;
		}

		display: flex;
		flex-wrap: wrap;
		justify-content: space-evenly;
		background-color: rgb(247, 247, 247);
		padding-bottom: 60px;

		.devCard {
			width: 300px;
			// height: 110px;
			padding: 10px 20px;
			margin: 10px;
			border-radius: 20px;

			.devCardTitle {
				margin-bottom: 5px;
			}
		}
	}

	.graph {
		// margin-top: 40px;
		padding: 50px 0px;
		display: flex;
		flex-wrap: wrap;
		justify-content: space-around;

		.graphContent {
			width: 45%;
			height: 300px;
			margin-top: 10px;
			margin-bottom: 10px;
		}

		.el-card {
			border-radius: 35px;
		}


	}
</style>
