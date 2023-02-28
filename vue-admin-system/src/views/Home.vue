<template>
	<div>
		<div class="info">
			<p class="access">通过手机或电脑，记录入睡、起床、体重变化的规律</p>
			<p class="title">每日记录、每日都会生成报表</p>
			<el-button class="infoButton gradient_btn" @click="dialogShow">每日报表是什么？</el-button>

			<el-dialog :visible.sync="dialogVisible" width="40%" :show-close=false :close-on-click-modal=false>
				<p class="about">一个自制能力强大的人，靠的并不是什么玄而又玄的「意志」，</p>
				<p class="about">而是因为他们在守序状态下，人体各激素水平稳定。</p>
				<p class="about">因此能分泌足够多巴胺去支持他们做一件又一件事。</p>
				<br>
				<p class="about">而那些「自制」差的人，他的激素从头到尾就是紊乱的。</p>
				<p class="about">即使守序也只能缓解，不能从根本上解决紊乱。</p>
				<p class="about">因此做事没什么成就感，更不会守序。</p>
				<br>
				<p class="about">这里用于帮助大家找到自己的规律，调整生活的节奏。</p>
				<p class="about">可以使用手机或电脑录入数据，电脑会自动生成图表。</p>
				<p class="about">只是看图是没用的，还需要自己在生活中调整。</p>
				<p class="about">使未来比现在更好</p>
				<br>
				<br>
				<el-button type="text gradient_text" @click="closeAbout">关闭</el-button>
			</el-dialog>
			<div>
				<el-button class="gradient_text" type="text" @click="toLogin">立刻使用 <span class="iconfont"
						style="color:#58B758;font-size: 14px; ">&#xe674;</span></el-button>

			</div>
		</div>

		<div class="dev">
			<div class="devTop"><span class="iconfont" style="font-size: 25px;color: #7f7f7f;">&#xe602;</span> 开发记录
			</div>

			<el-card class="devCard" style="border: 2px solid #ff3f5f;">
				<div class="strong devCardTitle"><svg class="icon" aria-hidden="true" style="font-size: 30px;">
						<use xlink:href="#icon-chigua"></use>
					</svg> 使用方法</div>
				<!-- <span class="iconfont"
						style="font-size: 30px;font-weight: 300;">&#xe60a;</span> -->
				<p>1.每天在这个网址上录入自己的数据，电脑点击左侧<span class="strong">「录入数据」</span>，手机点击最上方<span class="strong">「录入」</span>。可录入入睡时间、起床时间、体重变化等</p>
				<p>2.建议每晚点击<span class="strong">「报表」</span>复盘总结</p>
				<p class="red-text bigger">3.为保险起见，建议每隔一段时间点击右上方小人，<span class="underline-text strong">下载所有原始数据（Excel格式）</span>。</p>
				<p class="red-text bigger">4.登录<span class="underline-text strong">一个月</span>后需要重新登录，建议把密码写在纸上以防忘记。</p>
			</el-card>
			<el-card class="devCard">
				<div class="login-info">
					<div class="strong devCardTitle"><span class="iconfont"
							style="font-size: 30px;font-weight: 300;color: #fc0107;">&#xe945;</span> 更新记录</div>
					<p>2023年2月25日 <span>预上线，add:1.功能条。2.时间提示。3.查询月份数据。4.喝水。fix:1.汇总条取消。2.tooltip变动。</span></p>
					<p>2023年1月22日 <span>预上线，add:1.音乐管理。2.提示消息变更。3.体重录入修改。fix:1.报表滑动</span></p>
					<p>2023年1月11日 <span>预上线，add:1.用户注册功能。 2.数据导出为Excel格式。3.修改密码。4.日志记录。 5.定时备份数据库。 6.用户个人中心页面。7.用户登出。
							update：1.主页。 2.手机版视图优化。 3.cookie的获取机制</span></p>
					<p>2023年1月9日 <span>预上线，fix:1.nginx刷新404。 2.数据库导入导出。 3.防火墙问题</span></p>
				</div>
			</el-card>

			<el-card class="devCard">
				<div class="strong devCardTitle"><span class="iconfont"
						style="font-size: 30px;font-weight: 300;color: #0f80ff;">&#xeb62;</span> 准备添加的功能</div>
				<p>1.防盗刷 2.通信加密 3.密码验证算法更新</p>
				<p>1.注册 2.logger</p>
			</el-card>
			<el-card class="devCard">
				<div class="user">
					<div class="strong devCardTitle"><span class="iconfont"
							style="font-size: 30px;color: #66ccff;">&#xebb3;</span>
						项目架构</div>
					前端 vue+ElementUI+echart
					后端 golang+mysql+gin
				</div>
			</el-card>
			<el-card class="devCard">
				<div class="strong devCardTitle"> 联系我</div>
				<p><span class="iconfont" style="font-size: 25px;font-weight: 300;">&#xe882;</span>1911447416</p>
			</el-card>

		</div>
		<div class="dev">
			<div class="devTop"><span class="iconfont" style="font-size: 30px;font-weight: 300;">&#x10049;</span>提建议
			</div>
			<el-card style="width:70%;border-radius: 30px;padding: 30px;">
				<div>
					<el-input type="textarea" :rows="5" placeholder="未登录用户也可提建议,希望能留下您的联系方式.😝(#^.^#)"
						v-model="suggestarea">
					</el-input>
				</div>
				<div style="display: flex;justify-content: flex-end;margin-top: 10px;">
					<el-button @click="sendSuggest"><span class="iconfont"
							style="font-weight: 300;color: #58B758;">&#xe893;</span>发送</el-button>
				</div>
			</el-card>
		</div>

		<div class="graph">
			<el-card class="graphContent">
				<!-- 折线图 -->
				<div ref="echarts1" style="height: 260px;"></div>
			</el-card>
			<el-card class="graphContent">
				<!-- 柱状图 -->
				<div ref="echarts2" style="height: 260px;"></div>

			</el-card>
			<el-card class="graphContent">
				<!-- 饼图 -->
				<div ref="echarts3" style="height: 240px;"></div>
			</el-card>
			<el-card class="graphContent">
				<!-- 饼图 -->
				<div ref="echarts4" style="height: 240px;"></div>
			</el-card>
		</div>
		</el-col>

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
				suggestarea: '',
				dialogVisible: false,
				orderData: {
					data: [{
						'Mon': 1493,
						'Tue': 1020,
						'Wed': 123,
						'Thu': 4213,
						'Fri': 5324,
						'Sat': 4213,
						'Sun': 1235,
					}],
				},
				userData: [{
					date: '周一',
					new: 5,
					active: 21,
				}, {
					date: '周二',
					new: 54,
					active: 41,
				}, {
					date: '周三',
					new: 56,
					active: 91,
				}, {
					date: '周四',
					new: 67,
					active: 100,
				}, {
					date: '周五',
					new: 90,
					active: 102,
				}, {
					date: '周六',
					new: 100,
					active: 10,
				}, {
					date: '周日',
					new: 201,
					active: 100,
				}],
				videoData: [{
					name: "a",
					value: 2999,
				}, {
					name: "b",
					value: 5999,
				}, {
					name: "c",
					value: 1500,
				}, {
					name: "d",
					value: 1999,
				}, {
					name: "e",
					value: 2222,
				}, {
					name: "f",
					value: 4500,
				}]
			}
		},
		methods: {
			dialogShow() {
				this.dialogVisible = true
			},
			closeAbout() {
				this.dialogVisible = false
			},
			toLogin() {
				this.$router.push({
					path: "/login",
				});
			},

			sendSuggest() {
				http.post("user/sentSuggest", {
					suggestContent: this.suggestarea,
				}, {}).then((response) => {
					if (response.data.err_code != 0) {
						this.$message.error(response.data.err_msg);
						this.loading = false

					} else {
						this.$message.success("发送成功");
						this.loading = false
					}
				}).catch(reason => {
					this.loading = false
				})

			}
		},
		mounted() {
			// 图表1
			const echarts1 = echarts.init(this.$refs.echarts1)
			var echarts1Option = {}
			// x轴数据
			echarts1Option.xAxis = {
				data: Object.keys(this.orderData.data[0])
			}
			// console.log(Object.keys(this.orderData.data[0]))

			echarts1Option.yAxis = {}
			echarts1Option.series = [{

				data: [1493, 1020, 123, 4213, 5324, 4213, 1235],
				type: 'line',
			}]

			echarts1Option.tooltip = {
					trigger: 'axis'
				},

				console.log(this.orderData.data)
			echarts1.setOption(echarts1Option)


			// 图表2
			const echarts2 = echarts.init(this.$refs.echarts2)
			const echarts2Option = {
				legend: {
					textStyle: {
						color: "#333",
					},
				},
				grid: {
					left: "20%",
				},
				tooltip: {
					trigger: "axis",
				},
				xAxis: {
					type: "category",
					data: this.userData.map(item => item.date),
					axisLine: {
						lineStyle: {
							color: "#17b3a3",

						}
					},
					axisLabel: {
						interval: 0,
						color: "#333",
					}
				},
				yAxis: [{
					type: "value",
					axisLine: {
						lineStyle: {
							color: "#17b3a3",
						}
					}
				}],
				color: ["#2ec7c9", "#b6a2de"],
				series: [{
					name: '新增',
					data: this.userData.map(item => item.new),
					type: 'bar',
				}, {
					name: '活跃',
					data: this.userData.map(item => item.active),
					type: 'bar',
				}],

			}
			echarts2.setOption(echarts2Option)


			// 饼图
			const echarts3 = echarts.init(this.$refs.echarts3)
			const echarts3Option = {
				tooltip: {
					trigger: 'item',
				},
				color: [
					"#0f78f4",
					"#dd53b6",
					"#9462e5",
					"#a6a6a6",
					"#e1bb22",
					"#39c362",
					"#3ed1cf"
				],
				series: [{
					data: this.videoData,
					type: "pie"

				}]
			}
			echarts3.setOption(echarts3Option)
			// 饼图
			const echarts4 = echarts.init(this.$refs.echarts4)
			echarts4.setOption(echarts3Option)

		}

	}
</script>

<style lang="less" scoped>
	@media screen and (max-width:800px) {
		.info {
			padding: 0px 20px;

			.about {
				text-align: left;
			}

		}

		.graph {
			.graphContent {
				width: 85% !important;
			}
		}

		.el-dialog {
			width: 95% !important;
		}

	}

	@font-face {
		font-family: 'iconfont';
		src: url('../assets/iconfont/iconfont.ttf') format('truetype');
	}

	/deep/.el-dialog {
		border-radius: 10px;
	}

	.gradient_text {
		background-image: linear-gradient(95deg, #fd4536, #ec546e 35%, #c65f91 67%, #705fae);
		background-clip: text;
		color: transparent;
		transition: background-image 0.5s linear;
	}

	.gradient_text:hover {
		// background-image: linear-gradient(95deg, #fd4536, #ec546e 35%, #c65f91 67%, #705fae);
		opacity: .7;
	}

	.gradient_btn {
		background: linear-gradient(101deg, #fd4536, #ec546e 35%, #c65f91 67%, #705fae);
		color: #fff;
	}

	.gradient_btn:hover {
		opacity: .7;
	}

	.iconfont {
		font-family: "iconfont" !important;
		font-size: 16px;
		font-style: normal;
		-webkit-font-smoothing: antialiased;
		-moz-osx-font-smoothing: grayscale;
	}

	.icon {
		width: 1em;
		height: 1em;
		vertical-align: -0.15em;
		fill: currentColor;
		overflow: hidden;
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
	
	.bigger{
		font-size: 18px !important;
	}

	.red-text {
		color: #fd4536;
	}

	.underline-text {
		text-decoration: underline solid #ff3f5f;
		text-underline-position: under;
	}


	.dev {
		.devTop {
			width: 100%;
			font-size: 20px;
			text-align: center;
			margin: 30px 0px;

			svg {
				font-size: 20px;
			}
		}

		display: flex;
		flex-wrap: wrap;
		justify-content: space-evenly;
		background-color: rgb(247, 247, 247);
		padding-bottom: 60px;

		.devCard {
			width: 300px;
			// height: 110px;
			padding: 20px 15px;
			margin: 10px;
			border-radius: 20px;

			.devCardTitle {
				margin-bottom: 5px;
			}

			p {
				font-size: 15px;
				margin: 10px 0px;
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
