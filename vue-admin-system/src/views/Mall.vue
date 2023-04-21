<template>
	<div>



		<el-card class="player" style="margin-top: 20px;">
			<aplayer :repeat="repeatType" :music="audio[0]" :list="audio" v-if="musicShow" :showlrc="true"></aplayer>
			<div class="controlButton">
				<el-button class="" type="text" style="margin-right: 10px;" @click="oneMusic">单曲循环</el-button>
				<el-button class="" type="text" style="margin-right: 10px;" @click="allMusic">列表循环（默认）
				</el-button>
			</div>
		</el-card>

		<div class="chooseBanner">
			<div v-show="showSetPrint" style="margin-right: 10px;width:auto">
				<el-popconfirm title="功能正处于测试阶段，手机用户绘图可能显示异常，继续吗？" @confirm="setPrint">
					<el-button slot="reference" type="text"><span class="gradient_text">准备生成报告</span></el-button>
				</el-popconfirm>
			</div>
			<el-date-picker v-show="showChooseMonth" v-model="chooseMonth" type="month" placeholder="选择查找的月份"
				value-format="timestamp" @change="doChooseMonth">
			</el-date-picker>
			<div type="text" @click="doCopy" v-show="chooseMonthTitle">
				<span class="gradient_text">复制标题</span>
			</div>
			<div type="text" @click="downloadImg" v-show="showDownload">
				<span class="gradient_text">1.下载图片</span>
			</div>
			<div v-show="showDownload" @click="downloadDoc">
				<span class="gradient_text">2.下载文档模板</span>
			</div>


		</div>

		<div v-bind:style="bindgraph" class="graph">


			<!-- 月度拉屎 -->
			<el-card class="card" v-bind:style="bindcard">
				<p style="font-size: 18px;font-weight: 700;">拉屎月份表</p>
				<el-calendar :first-day-of-week=7 v-model="calendarShitDate">
					<!-- 这里使用的是 2.5 slot 语法，对于新项目请使用 2.6 slot 语法-->
					<template slot="dateCell" slot-scope="{data}">
						<p style="font-size: 13px;">
							{{ data.day.split('-').slice(1).join('-') }}
						</p>
						<div style="" v-for="(item, index) in scheduleShitData" :key="index">
							<div v-if="item.workingDay.indexOf(data.day) != -1">
								<ul v-for="(content, index) in item.content" :key="index"
									style="height: 100%;padding: 0;margin: 0;">
									<li :class="content.type"
										style="font-size: 16px;list-style:none;margin: 0 5px;text-align: center;">
										{{ content.notice }}
									</li>
								</ul>

							</div>
						</div>
					</template>
				</el-calendar>
			</el-card>

			<!-- 月度腹胀 -->
			<el-card class="card" v-bind:style="bindcard">
				<p style="font-size: 18px;font-weight: 700;">腹胀月份表</p>
				<el-calendar :first-day-of-week=7 v-model="calendarAcidDate">
					<template slot="dateCell" slot-scope="{data}">
						<p style="font-size: 13px;">
							{{ data.day.split('-').slice(1).join('-') }}
						</p>
						<div style="" v-for="(item, index) in scheduleAcidData" :key="index">
							<div v-if="item.workingDay.indexOf(data.day) != -1">
								<ul v-for="(content, index) in item.content" :key="index"
									style="height: 100%;padding: 0;margin: 0;">
									<li :class="content.type"
										style="font-size: 16px;list-style:none;margin: 0 5px;text-align: center;">
										{{ content.notice }}
									</li>
								</ul>

							</div>
						</div>
					</template>
				</el-calendar>
			</el-card>

			<!-- 10天反酸 -->
			<el-card class="card" v-bind:style="bindcard">
				<div ref="echarts18" style="height: 400px;width:100%"></div>
				<div class="buttonGroup" v-show="showButtonGroup">
					<el-button-group>
						<div class="block">
							<el-pagination layout="prev, next" :total="acidPageCount" prev-text="前进" next-text="后退"
								@current-change="acidPageChange">
							</el-pagination>
						</div>
					</el-button-group>
				</div>
			</el-card>

			<!-- 反酸汇总 -->
			<el-card class="card" v-bind:style="bindcard">
				<div ref="echarts19" style="height: 400px;width:100%"></div>
			</el-card>

			<!-- 10天喝水 -->
			<el-card class="card" v-bind:style="bindcard">
				<div ref="echarts16" style="height: 400px;width:100%"></div>
				<div class="buttonGroup" v-show="showButtonGroup">
					<el-button-group>
						<div class="block">
							<el-pagination layout="prev, next" :total="waterPageCount" prev-text="前进" next-text="后退"
								@current-change="waterPageChange">
							</el-pagination>
						</div>
					</el-button-group>
				</div>
			</el-card>

			<!-- 喝水汇总 -->
			<el-card class="card" v-bind:style="bindcard">
				<div ref="echarts17" style="height: 400px;width:100%"></div>
			</el-card>



			<!-- 10天睡眠 -->
			<transition name="el-fade-in-linear">
				<el-card class="card" v-bind:style="bindcard">
					<div ref="echarts4" style="height: 400px;width:100%"></div>
					<div class="buttonGroup">
						<el-button-group v-show="showButtonGroup">
							<div class="block">
								<el-pagination layout="prev, next" :total="sleepPageCount" prev-text="前进" next-text="后退"
									@current-change="sleepPageChange">
								</el-pagination>
							</div>
						</el-button-group>
					</div>
				</el-card>
			</transition>

			<!-- 睡眠汇总 -->
			<el-card class="card" v-bind:style="bindcard">
				<div ref="echarts5" style="height: 400px;width:100%"></div>
			</el-card>


			<!-- 10天睡起 -->
			<transition name="el-fade-in-linear">
				<el-card class="card" v-bind:style="bindcard">
					<div ref="echarts14" style="height: 480px;width:100%"></div>
					<div class="buttonGroup">
						<el-button-group v-show="showButtonGroup">
							<div class="block">
								<el-pagination layout="prev, next" :total="wakePageCount" prev-text="前进" next-text="后退"
									@current-change="wakePageChange">
								</el-pagination>
							</div>
						</el-button-group>
					</div>
				</el-card>
			</transition>

			<!-- 睡起汇总 -->
			<el-card class="card" v-bind:style="bindcard">
				<div ref="echarts15" style="height: 480px;width:100%"></div>
			</el-card>

			<!-- shit -->
			<el-card class="card" v-bind:style="bindcard">
				<div ref="echarts12" style="height: 480px;width:100%"></div>
				<div class="buttonGroup">
					<el-button-group v-show="showButtonGroup">
						<div class="block">
							<el-pagination layout="prev, next" :total="shitPageCount" prev-text="前进" next-text="后退"
								@current-change="shitPageChange">
							</el-pagination>
						</div>
					</el-button-group>
				</div>
			</el-card>

			<!-- shit汇总 -->
			<el-card class="card" v-bind:style="bindcard">
				<div ref="echarts13" style="height: 480px;width:100%"></div>
			</el-card>

			<!-- 10天体重 -->
			<el-card class="card" v-bind:style="bindcard">
				<div ref="echarts7" style="height: 500px;width:100%"></div>
				<div class="buttonGroup" v-show="showButtonGroup">
					<el-button-group>
						<div class="block">
							<el-pagination layout="prev, next" :total="weightPageCount" prev-text="前进" next-text="后退"
								@current-change="weightPageChange">
							</el-pagination>
						</div>
					</el-button-group>
				</div>
			</el-card>

			<!-- 体重汇总 -->
			<el-card class="card" v-bind:style="bindcard">
				<div ref="echarts8" style="height: 500px;width:100%"></div>
			</el-card>


			<!-- 10天beat -->
			<el-card class="card" v-bind:style="bindcard">
				<div ref="echarts9" style="height: 400px;width:100%"></div>
				<div class="buttonGroup" v-show="showButtonGroup">
					<el-button-group>
						<div class="block">
							<el-pagination layout="prev, next" :total="heartPageCount" prev-text="前进" next-text="后退"
								@current-change="heartPageChange">
							</el-pagination>
						</div>
					</el-button-group>
				</div>
			</el-card>

			<!-- beat汇总 -->
			<el-card class="card" v-bind:style="bindcard">
				<div ref="echarts10" style="height: 400px;width:100%"></div>
			</el-card>

			<!-- 10天手机 -->
			<el-card class="card" v-bind:style="bindcard">
				<div ref="echarts3" style="height: 400px;width:100%"></div>
				<div class="buttonGroup" v-show="showButtonGroup">
					<div class="block">
						<el-pagination layout="prev, next" :total="phonePageCount" prev-text="前进" next-text="后退"
							@current-change="phonePageChange">
						</el-pagination>
					</div>
				</div>
			</el-card>

			<!-- 手机汇总 -->
			<el-card class="card" v-bind:style="bindcard">
				<div ref="echarts6" style="height: 400px;width:100%"></div>
			</el-card>


			<el-card class="card" v-bind:style="bindcard">
				<div style="text-align: center;">
					已累计记录phone
				</div>
				<div class="card_dashBoard">
					<div>
						<ul>
							<li v-for="item in phoneDashBoardH">
								<img class="phoneDashImg" :src="require(`../assets/num`+item+`.png`)" alt="">
							</li>
							小时
						</ul>
					</div>
					<div style="">
						<ul>
							<li v-for="item in phoneDashBoardM">
								<img class="phoneDashImg" :src="require(`../assets/num`+item+`.png`)" alt="">
							</li>
							分钟
						</ul>
					</div>
				</div>
				<div style="text-align: center;">
					一周累计
				</div>
				<div class="card_dashBoard">
					<div>
						<ul>
							<li v-for="item in phoneDashBoardEveryWeakH">
								<img class="phoneDashImg" :src="require(`../assets/num`+item+`.png`)" alt="">
							</li>
							小时
						</ul>
					</div>
					<div style="">
						<ul>
							<li v-for="item in phoneDashBoardEveryWeakM">
								<img class="phoneDashImg" :src="require(`../assets/num`+item+`.png`)" alt="">
							</li>
							分钟
						</ul>
					</div>
				</div>
				<div style="text-align: center;">
					一年估计
				</div>
				<div class="card_dashBoard">
					<div>
						<ul>
							<li v-for="item in phoneDashBoardEveryYearH">
								<img class="phoneDashImg" :src="require(`../assets/num`+item+`.png`)" alt="">
							</li>
							小时
						</ul>
					</div>
					<div style="">
						<ul>
							<li v-for="item in phoneDashBoardEveryYearM">
								<img class="phoneDashImg" :src="require(`../assets/num`+item+`.png`)" alt="">
							</li>
							分钟
						</ul>
					</div>
				</div>
				<div style="text-align: center;">
					百年估计
				</div>
				<div class="card_dashBoard">
					<div>
						<ul>
							<li v-for="item in phoneDashBoardEvery100YearH">
								<img class="phoneDashImg" :src="require(`../assets/num`+item+`.png`)" alt="">
							</li>
							小时
						</ul>
					</div>
					<div style="">
						<ul>
							<li v-for="item in phoneDashBoardEvery100YearM">
								<img class="phoneDashImg" :src="require(`../assets/num`+item+`.png`)" alt="">
							</li>
							分钟
						</ul>
					</div>
				</div>
			</el-card>


		</div>
	</div>
</template>

<script>
	import aplayer from "vue-aplayer"
	import * as echarts from 'echarts'
	import { Loading } from 'element-ui';

	import {
		drawsleep
	} from '../utils/sleepChart.js'
	import {
		drawSleepSum
	} from '../utils/sleepSumChart.js'
	import {
		drawphone
	} from '../utils/phoneChart.js'

	import {
		drawPhoneSum
	} from '../utils/phoneSumChart.js'

	import {
		drawWeight
	} from '../utils/weightChart.js'
	import {
		drawWeightSum
	} from '../utils/weightSumChart.js'


	import {
		drawHeart
	} from '../utils/heartChart.js'

	import {
		drawHeartSum
	} from '../utils/heartSumChart.js'

	import {
		fomartTime
	} from '../utils/time.js'

	import {
		drawShit
	} from '../utils/shitChart.js'

	import {
		drawShitSum
	} from '../utils/shitSumChart.js'

	import {
		drawShitBSum
	} from '../utils/shitBChart.js'

	import {
		drawWake
	} from '../utils/wakeChart.js'

	import {
		drawWakeSum
	} from '../utils/wakeSumChart.js'

	import {
		drawWater
	} from '../utils/waterChart.js'

	import {
		drawWaterSum
	} from '../utils/waterSumChart.js'



	import {
		drawAcid
	} from '../utils/acidChart.js'

	import {
		drawAcidSum
	} from '../utils/acidSumChart.js'

	import moment from "moment";
	import http from '../utils/request.js'
	import {
		watch
	} from "vue"

	export default {
		data() {
			return {
				// 音频列表
				// audio: [{
				// 	title: 'さくら、もゆ。 -title arrange-',
				// 	artist: '忍',
				// 	src: 'http://43.139.13.115:8092/more_static/piano.mp3',
				// 	pic: 'http://43.139.13.115:8092/more_static/piano.jpeg', 
				// 	lrc: '',
				// }, {
				// 	title: '曇り空の向こうは晴れている(Instrumental) - (阴云彼端 晴空依然)',
				// 	artist: '22/7',
				// 	src: 'http://43.139.13.115:8092/more_static/piano2.m4a',
				// 	pic: 'http://43.139.13.115:8092/more_static/piano2.jpeg',
				// 	lrc: '',
				// }, ],

				// audio: [{
				// 	title: "さくら、もゆ。 -title arrange-",
				// 	artist: "忍",
				// 	src: "http://43.139.13.115:8092/more_static/piano.mp3",
				// 	pic: "http://43.139.13.115:8092/more_static/piano.jpeg", 
				// 	lrc: '',
				// } ],

				audio: [],
				musicShow: false,
				repeatType: "repeat-all",
				printType: "",
				chooseMonth: "",
				bindgraph: "",
				bindcard: "",
				phonePageNow: 1,
				phonePageCount: 0,
				sleepPageNow: 1,
				sleepPageCount: 0,
				shitPageNow: 1,
				shitPageCount: 0,
				wakePageNow: 1,
				wakePageCount: 0,
				weightPageCount: 0,
				weightPageNow: 1,
				heartPageNow: 1,
				heartPageCount: 0,
				waterPageNow: 1,
				waterPageCount: 0,
				acidPageNow: 1,
				acidPageCount: 0,

				showButtonGroup: true,
				thisMonth: "",
				nextMonth: "",
				chooseMonthTitle: false,
				showDownload: false,
				showChooseMonth: false,
				showSetPrint: true,
				phoneDashBoardH: [],
				phoneDashBoardM: [],
				phoneDashBoardEveryWeakH: [],
				phoneDashBoardEveryWeakM: [],
				phoneDashBoardEveryYearH: [],
				phoneDashBoardEveryYearM: [],
				phoneDashBoardEvery100YearH: [],
				phoneDashBoardEvery100YearM: [],

				echarts3: null,
				echarts4: null,
				echarts12: null,
				echarts14: null,
				echarts7: null,
				echarts9: null,
				echarts16: null,
				echarts18: null,
				tag: true,
				calendarShitDate: new Date(),
				calendarAcidDate: new Date(),
				scheduleShitData: [],
				scheduleAcidData: [],
				// scheduleShitData: [{
				// 	workingDay: "2023-03-02",
				// 	content: [{
				// 			notice: "疑似",
				// 			type: "important",
				// 		},

				// 	],
				// }, {
				// 	workingDay: "2023-03-01",
				// 	content: [{
				// 			notice: "擦有",
				// 			type: "important",
				// 		},

				// 	],
				// }, {
				// 	workingDay: "2023-03-04",
				// 	content: [{
				// 			notice: "明显",
				// 			type: "important",
				// 		},

				// 	],
				// }, {
				// 	workingDay: "2023-03-05",
				// 	content: [{
				// 			notice: "未现",
				// 			type: "safe",
				// 		},

				// 	],
				// }],

				// scheduleAcidData: [{
				// 	workingDay: "2023-03-02",
				// 	content: [{
				// 			notice: "小胀",
				// 			type: "important",
				// 		},

				// 	],
				// }, {
				// 	workingDay: "2023-03-01",
				// 	content: [{
				// 			notice: "不胀",
				// 			type: "safe",
				// 		},

				// 	],
				// }, {
				// 	workingDay: "2023-03-04",
				// 	content: [{
				// 			notice: "中胀",
				// 			type: "important",
				// 		},

				// 	],
				// }, {
				// 	workingDay: "2023-03-05",
				// 	content: [{
				// 			notice: "大胀",
				// 			type: "important",
				// 		},

				// 	],
				// }]


			}
		},
		watch: {
			"calendarShitDate"(newVal, oldVal) {
				console.log(newVal.getTime())
				http.post("shit/shitDate", {
					month: newVal.getTime()
				}, {}).then((res) => {
					this.scheduleShitData = res.data.data
					console.log(this.scheduleShitData)
				})
			},
			"calendarAcidDate"(newVal, oldVal) {
				console.log(newVal.getTime())
				http.post("acid/acidDate", {
					month: newVal.getTime()
				}, {}).then((res) => {
					this.scheduleAcidData = res.data.data
				})
			}
		},
		methods: {

			chooseCalendar(time) {
				console.log(time)
			},
			downloadDoc() {
				const title = "月度报告 （" + this.thisMonth + "~" + this.nextMonth + "）"
				http.post("doc/download", {
					title: title,
				}, {
					responseType: 'arraybuffer'
				}).then((res) => {
					// 假设 data 是返回来的二进制数据
					const data = res.data
					const url = window.URL.createObjectURL(new Blob([data], {
						type: "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
					}))
					const link = document.createElement('a')
					link.style.display = 'none'
					link.href = url
					link.setAttribute('download', title + ".docx")
					document.body.appendChild(link)
					link.click()
					document.body.removeChild(link)

				})
			},
			downloadImg() {
				const title = "（" + this.thisMonth + "~" + this.nextMonth + "）"
				const echarts4 = echarts.init(this.$refs.echarts4)
				var picInfo = echarts4.getDataURL({
					type: 'png',
					pixelRatio: 3, //放大两倍下载，之后压缩到同等大小展示。解决生成图片在移动端模糊问题
					backgroundColor: '#fff'
				}); //获取到的是一串base64信息
				const elink = document.createElement('a');
				elink.download = "3-睡眠" + title;
				elink.style.display = 'none';
				elink.href = picInfo;
				document.body.appendChild(elink);
				elink.click();
				URL.revokeObjectURL(elink.href); // 释放URL 对象
				document.body.removeChild(elink)

				const echarts3 = echarts.init(this.$refs.echarts3)
				var picInfo = echarts3.getDataURL({
					type: 'png',
					pixelRatio: 3, //放大两倍下载，之后压缩到同等大小展示。解决生成图片在移动端模糊问题
					backgroundColor: '#fff'
				}); //获取到的是一串base64信息
				elink.download = "8-手机" + title;
				elink.href = picInfo;
				document.body.appendChild(elink);
				elink.click();
				URL.revokeObjectURL(elink.href); // 释放URL 对象
				document.body.removeChild(elink)

				const echarts12 = echarts.init(this.$refs.echarts12)
				var picInfo = echarts12.getDataURL({
					type: 'png',
					pixelRatio: 3, //放大两倍下载，之后压缩到同等大小展示。解决生成图片在移动端模糊问题
					backgroundColor: '#fff'
				}); //获取到的是一串base64信息
				elink.download = "5-拉屎" + title;
				elink.href = picInfo;
				document.body.appendChild(elink);
				elink.click();
				URL.revokeObjectURL(elink.href); // 释放URL 对象
				document.body.removeChild(elink)

				const echarts14 = echarts.init(this.$refs.echarts14)
				var picInfo = echarts14.getDataURL({
					type: 'png',
					pixelRatio: 3, //放大两倍下载，之后压缩到同等大小展示。解决生成图片在移动端模糊问题
					backgroundColor: '#fff'
				}); //获取到的是一串base64信息
				elink.download = "4-睡起" + title;
				elink.href = picInfo;
				document.body.appendChild(elink);
				elink.click();
				URL.revokeObjectURL(elink.href); // 释放URL 对象
				document.body.removeChild(elink)

				const echarts7 = echarts.init(this.$refs.echarts7)
				var picInfo = echarts7.getDataURL({
					type: 'png',
					pixelRatio: 3, //放大两倍下载，之后压缩到同等大小展示。解决生成图片在移动端模糊问题
					backgroundColor: '#fff'
				}); //获取到的是一串base64信息
				elink.download = "7-体重" + title;
				elink.href = picInfo;
				document.body.appendChild(elink);
				elink.click();
				URL.revokeObjectURL(elink.href); // 释放URL 对象
				document.body.removeChild(elink)

				const echarts9 = echarts.init(this.$refs.echarts9)
				var picInfo = echarts9.getDataURL({
					type: 'png',
					pixelRatio: 3, //放大两倍下载，之后压缩到同等大小展示。解决生成图片在移动端模糊问题
					backgroundColor: '#fff'
				}); //获取到的是一串base64信息
				elink.download = "6-心脏" + title;
				elink.href = picInfo;
				document.body.appendChild(elink);
				elink.click();
				URL.revokeObjectURL(elink.href); // 释放URL 对象
				document.body.removeChild(elink)

				const echarts16 = echarts.init(this.$refs.echarts16)
				var picInfo = echarts16.getDataURL({
					type: 'png',
					pixelRatio: 3, //放大两倍下载，之后压缩到同等大小展示。解决生成图片在移动端模糊问题
					backgroundColor: '#fff'
				}); //获取到的是一串base64信息
				elink.download = "2-喝水" + title;
				elink.href = picInfo;
				document.body.appendChild(elink);
				elink.click();
				URL.revokeObjectURL(elink.href); // 释放URL 对象
				document.body.removeChild(elink)

				const echarts18 = echarts.init(this.$refs.echarts18)
				var picInfo = echarts18.getDataURL({
					type: 'png',
					pixelRatio: 3, //放大两倍下载，之后压缩到同等大小展示。解决生成图片在移动端模糊问题
					backgroundColor: '#fff'
				}); //获取到的是一串base64信息
				elink.download = "1-反酸" + title;
				elink.href = picInfo;
				document.body.appendChild(elink);
				elink.click();
				URL.revokeObjectURL(elink.href); // 释放URL 对象
				document.body.removeChild(elink)
			},
			doChooseMonth() {
				console.log(this.chooseMonth)
				this.showButtonGroup = false
				this.chooseMonthTitle = true
				this.showDownload = true
				http.post("sleep/searchMonth", {
					month: this.chooseMonth,
				}).then((response) => {
					if (response.data.data.sleepList != null) {
						// 图表3
						const echarts4 = echarts.init(this.$refs.echarts4)
						const option4 = drawsleep(response.data.data.sleepList)
						this.thisMonth = response.data.data.thisMonth
						this.nextMonth = response.data.data.nextMonth
						echarts4.setOption(option4)
						this.doCopy()

					} else {
						alert("sleepList")
					}
				})

				http.post("phone/searchMonth", {
					month: this.chooseMonth,
				}).then((response) => {
					// 图表3
					const echarts3 = echarts.init(this.$refs.echarts3)
					const option3 = drawphone(response.data.data.phoneList)
					echarts3.setOption(option3)
				})
				http.post("shit/searchMonth", {
					month: this.chooseMonth,
				}).then((response) => {
					// 图表12
					const echarts12 = echarts.init(this.$refs.echarts12)
					const option12 = drawShit(response.data.data.shitList)
					echarts12.setOption(option12)
				})
				http.post("wake/searchMonth", {
					month: this.chooseMonth,
				}).then((response) => {
					console.log(response.data.data)
					const echarts14 = echarts.init(this.$refs.echarts14)
					const option14 = drawWake(response.data.data.wakeList)
					echarts14.setOption(option14)
				})
				http.post("weight/searchMonth", {
					month: this.chooseMonth,
				}).then((response) => {
					// 图表8
					const echarts7 = echarts.init(this.$refs.echarts7)
					const option7 = drawWeight(response.data.data.weightList)
					echarts7.setOption(option7)
				})
				http.post("heart/searchMonth", {
					month: this.chooseMonth,
				}).then((response) => {
					if (response.data.data.heartList != null) {
						// 图表9
						const echarts9 = echarts.init(this.$refs.echarts9)
						const option9 = drawHeart(response.data.data.heartList)
						echarts9.setOption(option9)
					} else {}
				})

				http.post("water/searchMonth", {
					month: this.chooseMonth,
				}).then((response) => {
					if (response.data.data.waterList != null) {
						// 图表9
						const echarts16 = echarts.init(this.$refs.echarts16)
						const option16 = drawWater(response.data.data.waterList)
						echarts16.setOption(option16)
					} else {}
				})

				http.post("acid/searchMonth", {
					month: this.chooseMonth,
				}).then((response) => {
					if (response.data.data.acidList != null) {
						// 图表9
						const echarts18 = echarts.init(this.$refs.echarts18)
						const option18 = drawAcid(response.data.data.acidList)
						echarts18.setOption(option18)
					} else {}
				})

			},
			oneMusic() {
				this.repeatType = 'repeat-one'
				this.$notify({
					title: "已设置单曲播放",
					position: 'bottom-left',
					type: 'success',
					duration: 1400,
					showClose: false,
				});
			},
			allMusic() {
				this.repeatType = 'repeat-all'
				this.$notify({
					title: "已设置列表循环",
					position: 'bottom-left',
					type: 'success',
					duration: 1400,
					showClose: false,
				});
			},

			setPrint() {

				window.open("/mall?printType=print")

			},
			doCopy() {
				const title = " 月度报告 （" + this.thisMonth + "~" + this.nextMonth + "）"
				this.$copyText(title).then(
					e => {
						this.$notify({
							title: "复制成功",
							position: 'bottom-left',
							type: 'success',
							duration: 1400,
							showClose: false,
						});

					},
				)
			},
			getMusic() {
				http.post("music/getMusic", {}).then((response) => {
					console.log("music", response)
					if (response.data.err_code != 0) {

					} else {
						let List = response.data.data
						List.forEach(element => {

							let obj = {
								title: element.title,
								pic: element.pic,
								src: element.src,
								artist: element.artist,
								lrc: element.lrc
							}
							console.log("lrc", element.lrc)

							this.audio.push(obj);

						});
						this.audio = response.data.data
						console.log("audio", this.audio)
						this.musicShow = true
					}
				})
				// this.audio = [{
				// 	title: "さくら、もゆ。 -title arrange-",
				// 	artist: "忍",
				// 	src: "http://43.139.13.115:8092/more_static/piano.mp3",
				// 	pic: "http://43.139.13.115:8092/more_static/piano.jpeg",
				// 	lrc: '',
				// }]
				// let obj = {

				// 		title: 'さくら、もゆ。 -title arrange-',
				// 		artist: '忍',
				// 		src: 'http://43.139.13.115:8092/more_static/piano.mp3',
				// 		pic: 'http://43.139.13.115:8092/more_static/piano.jpeg', 
				// 		lrc: '',
				// }
				// this.audio.push(obj)
			},

			phonePageChange(now) {
				this.phonePageNow = now
				http.post("phone/list", {
					phonePageNow: this.phonePageNow,
				}).then((response) => {
					// 图表3
					const echarts3 = echarts.init(this.$refs.echarts3)
					const option3 = drawphone(response.data.data.phoneList)
					echarts3.setOption(option3)
				})
			},

			sleepPageChange(now) {
				this.sleepPageNow = now
				http.post("sleep/list", {
					sleepPageNow: this.sleepPageNow,
				}).then((response) => {
					// 图表3
					const echarts4 = echarts.init(this.$refs.echarts4)
					const option4 = drawsleep(response.data.data.sleepList)
					echarts4.setOption(option4)
				})
			},
			shitPageChange(now) {
				this.shitPageNow = now
				http.post("shit/list", {
					shitPageNow: this.shitPageNow,
				}).then((response) => {
					// 图表12
					const echarts12 = echarts.init(this.$refs.echarts12)
					const option12 = drawShit(response.data.data.shitList)
					echarts12.setOption(option12)
				})
			},

			wakePageChange(now) {
				this.wakePageNow = now
				http.post("wake/list", {
					wakePageNow: this.wakePageNow,
				}).then((response) => {
					console.log(response.data.data)
					const echarts14 = echarts.init(this.$refs.echarts14)
					const option14 = drawWake(response.data.data.wakeList)
					echarts14.setOption(option14)
				})
			},

			weightPageChange(now) {
				this.weightPageNow = now
				http.post("weight/list", {
					weightPageNow: this.weightPageNow,
				}).then((response) => {
					// 图表8
					const echarts7 = echarts.init(this.$refs.echarts7)
					const option7 = drawWeight(response.data.data.weightList)
					echarts7.setOption(option7)
				})
			},
			heartPageChange(now) {
				this.heartPageNow = now
				http.post("heart/list", {
					heartPageNow: this.heartPageNow,
				}).then((response) => {
					// 图表9
					const echarts9 = echarts.init(this.$refs.echarts9)
					const option9 = drawHeart(response.data.data.heartList)
					echarts9.setOption(option9)

				})
			},

			waterPageChange(now) {
				this.waterPageNow = now
				http.post("water/list", {
					waterPageNow: this.waterPageNow,
				}).then((response) => {
					const echarts16 = echarts.init(this.$refs.echarts16)
					const option16 = drawWater(response.data.data.waterList)
					echarts16.setOption(option16)

				})
			},

			acidPageChange(now) {
				this.acidPageNow = now
				http.post("acid/list", {
					acidPageNow: this.acidPageNow,
				}).then((response) => {
					const echarts18 = echarts.init(this.$refs.echarts18)
					const option18 = drawAcid(response.data.data.acidList)
					echarts18.setOption(option18)
				})
			},
		},
		components: {
			aplayer
		},

		created() {
			// let load = Loading.service({ fullscreen: true,background:"white",spinner:"el-icon-s-promotion",text:"生成中"})
			this.printType = this.$route.query.printType
			if (this.printType != null && this.printType == "print") {
				this.showChooseMonth = true
				this.showSetPrint = false
				this.bindgraph = "display: block !important;width: 92%;margin: 0% 2%;"
				this.bindcard =
					"margin: 20px auto;width: 1080px !important;border-radius: 10px !important;padding: 2% !important;"
			}

			Promise.all([
				new Promise((resolve) => {
					http.post("shit/shitDate", {
						month: Date.now()
					}, {}).then((res) => {
						this.scheduleShitData = res.data.data
						console.log(this.scheduleShitData)
						resolve("data1")
					})
				}),
				new Promise((resolve) => {
					http.post("phone/list", {
						phonePageNow: this.phonePageNow,
					}).then((response) => {
						if (response.data.data.phoneList != null) {
							// 图表3
							console.log(response.data.data)
							this.echarts3 = echarts.init(this.$refs.echarts3)
							const option3 = drawphone(response.data.data.phoneList)
							this.echarts3.setOption(option3)
							this.phonePageCount = response.data.data.phonePageCount
							resolve("data2")
						} else {
							// this.$message.error("手机空数据")
							resolve("data2")
						}
					})
				}),
				new Promise((resolve) => {
					http.post("phone/dashBoard", {}).then((response) => {
						if (response.data.data.phoneDashBoardH != null) {
							this.phoneDashBoardH = response.data.data.phoneDashBoardH
						} else {}
						if (response.data.data.phoneDashBoardM != null) {
							this.phoneDashBoardM = response.data.data.phoneDashBoardM
						} else {}
						if (response.data.data.everyWeakPhoneDashBoardH != null) {
							this.phoneDashBoardEveryWeakH = response.data.data.everyWeakPhoneDashBoardH
						} else {}
						if (response.data.data.everyWeakPhoneDashBoardM != null) {
							this.phoneDashBoardEveryWeakM = response.data.data.everyWeakPhoneDashBoardM
						} else {}
						if (response.data.data.everyYearPhoneDashBoardH != null) {
							this.phoneDashBoardEveryYearH = response.data.data.everyYearPhoneDashBoardH
						} else {}
						if (response.data.data.everyYearPhoneDashBoardM != null) {
							this.phoneDashBoardEveryYearM = response.data.data.everyYearPhoneDashBoardM
						} else {}
						if (response.data.data.every100YearPhoneDashBoardH != null) {
							this.phoneDashBoardEvery100YearH = response.data.data.every100YearPhoneDashBoardH
						} else {}
						if (response.data.data.every100YearPhoneDashBoardM != null) {
							this.phoneDashBoardEvery100YearM = response.data.data.every100YearPhoneDashBoardM
						} else {}
						resolve("data3")
					})
				}),
				new Promise((resolve) => {
					// 图表4
					http.post("sleep/list", {
						sleepPageNow: this.sleepPageNow,
					}).then((response) => {
						if (response.data.data.sleepList != null) {
							console.log(response.data.data)
							const echarts4 = echarts.init(this.$refs.echarts4)
							const option4 = drawsleep(response.data.data.sleepList)
							echarts4.setOption(option4)
							this.sleepPageCount = response.data.data.sleepPageCount
							resolve("data4")
						} else {
							// this.$message.error("睡眠空数据")
							resolve("data4")
						}
					})
				}),
				new Promise((resolve) => {
					// 图表5
					http.post("sleep/sum", {}).then((response) => {
						if (response.data.data.sleepList != null) {
							console.log(response.data.data)
							const echarts5 = echarts.init(this.$refs.echarts5)
							const option5 = drawSleepSum(response.data.data.sleepList)
							echarts5.setOption(option5)
							resolve("data5")
						} else {
							// this.$message.error("睡眠空数据")
							resolve("data5")
						}
					})
				}),
				new Promise((resolve) => {
					// 图表6
					http.post("phone/sum", {}).then((response) => {
					
						if (response.data.data.phoneList != null) {
							console.log(response.data.data)
							const echarts6 = echarts.init(this.$refs.echarts6)
							const option6 = drawPhoneSum(response.data.data.phoneList)
							echarts6.setOption(option6)
							resolve("data6")
						} else {
							// this.$message.error("手机空数据")
							resolve("data6")
						}
					})
				}),
				new Promise((resolve) => {
					// 图表7
					http.post("weight/list", {
						weightPageNow: this.weightPageNow,
					}).then((response) => {
						if (response.data.data.weightList != null) {
							console.log(response.data.data)
							const echarts7 = echarts.init(this.$refs.echarts7)
							const option7 = drawWeight(response.data.data.weightList)
							echarts7.setOption(option7)
							this.weightPageCount = response.data.data.weightPageCount
							resolve("data7")
						} else {
							// this.$message.error("体重空数据")
							resolve("data7")
						}
					})
				}),
				new Promise((resolve) => {
					// 图表8
					http.post("weight/sum", {}).then((response) => {
						if (response.data.data.weightList != null) {
							console.log(response.data.data)
							const echarts8 = echarts.init(this.$refs.echarts8)
							const option8 = drawWeightSum(response.data.data.weightList)
							echarts8.setOption(option8)
							resolve("data8")
						} else {
							// this.$message.error("体重空数据")
							resolve("data8")
						}
					})
				}),
				new Promise((resolve) => {
					// 图表9
					http.post("heart/list", {
						heartPageNow: this.heartPageNow,
					}).then((response) => {
						if (response.data.data.heartList != null) {
							console.log(response.data.data)
							const echarts9 = echarts.init(this.$refs.echarts9)
							const option9 = drawHeart(response.data.data.heartList)
							echarts9.setOption(option9)
							this.heartPageCount = response.data.data.heartPageCount
							resolve("data9")
						} else {
							// this.$message.error("律动空数据")
							resolve("data9")
						}
					})
				}),
				new Promise((resolve) => {
					// 图表10
					http.post("heart/sum", {}).then((response) => {
					
						if (response.data.data.heartList != null) {
							console.log(response.data.data)
							const echarts10 = echarts.init(this.$refs.echarts10)
							const option10 = drawHeartSum(response.data.data.heartList)
							echarts10.setOption(option10)
							resolve("data10")
						} else {
							// this.$message.error("律动空数据")
							resolve("data10")
						}
					})
				}),
				new Promise((resolve) => {
					// 图表12
					http.post("shit/list", {
						shitPageNow: this.shitPageNow,
					}).then((response) => {
						if (response.data.data.shitList != null) {
							console.log(response.data.data)
							const echarts12 = echarts.init(this.$refs.echarts12)
							const option12 = drawShit(response.data.data.shitList)
							echarts12.setOption(option12)
							this.shitPageCount = response.data.data.shitPageCount
							resolve("data11")
						} else {
							// this.$message.error("拉屎空数据")
							resolve("data11")
						}
					})
				}),
				new Promise((resolve) => {
					// 图表13
					http.post("shit/sum", {}).then((response) => {
						if (response.data.data.shitList != null) {
							console.log(response.data.data)
							const echarts13 = echarts.init(this.$refs.echarts13)
							const option13 = drawShitSum(response.data.data.shitList)
							echarts13.setOption(option13)
							resolve("data12")
						} else {
							// this.$message.error("拉屎空数据")
							resolve("data12")
						}
					})
				}),
				new Promise((resolve) => {
					// 图表14
					http.post("wake/list", {
						wakePageNow: this.wakePageNow,
					}).then((response) => {
						if (response.data.data.wakeList != null) {
							console.log(response.data.data)
							const echarts14 = echarts.init(this.$refs.echarts14)
							const option14 = drawWake(response.data.data.wakeList)
							echarts14.setOption(option14)
							this.wakePageCount = response.data.data.wakePageCount
							resolve("data13")
						} else {
							// this.$message.error("睡起空数据")
							resolve("data13")
						}
					})
				}),
				new Promise((resolve) => {
					// 图表15
					http.post("wake/sum", {}).then((response) => {
						if (response.data.data.wakeList != null) {
							console.log(response.data.data)
							const echarts15 = echarts.init(this.$refs.echarts15)
							const option15 = drawWakeSum(response.data.data.wakeList)
							echarts15.setOption(option15)
							resolve("data14")
						} else {
							// this.$message.error("睡起空数据")
							resolve("data14")
						}
					
					})
				}),
				new Promise((resolve) => {
					// 图表16
					http.post("water/list", {
						waterPageNow: this.waterPageNow,
					}).then((response) => {
						if (response.data.data.waterList != null) {
							console.log(response.data.data)
							const echarts16 = echarts.init(this.$refs.echarts16)
							const option16 = drawWater(response.data.data.waterList)
							echarts16.setOption(option16)
							this.waterPageCount = response.data.data.waterPageCount
							resolve("data15")
						} else {
							this.$message.error("喝水空数据")
							resolve("data15")
						}
					
					})
				}),
				new Promise((resolve) => {
					// 图表16
					http.post("water/sum", {}).then((response) => {
						if (response.data.data.waterList != null) {
							console.log(response.data.data)
							const echarts17 = echarts.init(this.$refs.echarts17)
							const option17 = drawWaterSum(response.data.data.waterList)
							echarts17.setOption(option17)
							resolve("data16")
						} else {
							this.$message.error("喝水空数据")
							resolve("data16")
						}
					
					})
				}),
				new Promise((resolve) => {
					// 图表16
					http.post("acid/list", {
						acidPageNow: this.acidPageNow,
					}).then((response) => {
						if (response.data.data.acidList != null) {
							console.log(response.data.data)
							const echarts18 = echarts.init(this.$refs.echarts18)
							const option18 = drawAcid(response.data.data.acidList)
							echarts18.setOption(option18)
							this.acidPageCount = response.data.data.acidPageCount
							resolve("data17")
						} else {
							this.$message.error("acid空数据")
							resolve("data17")
						}
					
					})
				}),
				new Promise((resolve) => {
					// 图表16
					http.post("acid/sum", {}).then((response) => {
						if (response.data.data.acidList != null) {
							console.log(response.data.data)
							const echarts19 = echarts.init(this.$refs.echarts19)
							const option19 = drawAcidSum(response.data.data.acidList)
							echarts19.setOption(option19)
							resolve("data18")
						} else {
							this.$message.error("acid空数据")
							resolve("data18")
						}
					})
				}),
			]).then(res => {
				console.log(res)
				//["data1","data2","data3","data4","data5","data6","data7","data8"]
				console.log("success.!!!")
				// load.close()
				/*
				
				渲染页面的代码
				
				*/
			})
		},
		mounted() {
			this.getMusic()
			
			http.post("acid/acidDate", {
				month: Date.now()
			}, {}).then((res) => {
				this.scheduleAcidData = res.data.data
			})

			setTimeout(() => {
				
				



				

				


				

				
				
				

				

				

				

				


				

				

				

				

				


			})

		}
	}
</script>


<style lang="less" scoped>
	@media screen and (max-width:1000px) {

		/deep/.el-card__body,
		.el-main {
			padding: 0px !important;
		}

		.graph {
			display: block !important;
			width: 92%;
			margin: 0% 2%;

			/deep/.card {
				margin: 0 auto;
				width: 100% !important;
				border-radius: 10px !important;
				padding: 2% !important;

				.phoneDashImg {
					width: 30px !important;
					height: 30px !important;
				}

				.el-card__body {
					padding: 0;
				}
			}
		}

		/deep/.el-calendar__body {
			padding: 0;
		}

	}

	@font-face {
		font-family: 'iconfont';
		src: url('../assets/iconfont/iconfont.ttf') format('truetype');
	}

	/deep/.el-calendar__header {
		justify-content: space-evenly
	}

	/deep/.el-calendar__body {
		.el-calendar-table__row {
			.prev {
				p {
					opacity: 0;
				}
			}

			.next {
				p {
					opacity: 0;
				}
			}

		}
	}

	/deep/.el-calendar-table {
		.el-calendar-day {
			padding: 8px 0px 8px 0px;
			height: 60px;
			text-align: center;
		}

		.el-calendar-day:hover {
			background-color: white;
		}

		td.is-selected {
			background-color: white;
		}

		thead {
			th {
				padding: 5px 0;
				font-size: 14px;
			}
		}
	}

	.iconfont {
		font-family: "iconfont" !important;
		font-size: 16px;
		font-style: normal;
		-webkit-font-smoothing: antialiased;
		-moz-osx-font-smoothing: grayscale;
	}



	.important {
		color: rgba(252, 255, 250, 1.0);
		background-color: rgba(210, 16, 5, 0.5);
	}

	.safe {
		color: rgba(0, 0, 0, 1.0);
		background-color: rgba(135, 235, 137, 0.5);
	}


	.secondarySty {
		color: tan;
		// background-color: rgba(235, 150, 22, 0.5);
	}

	.gradient_text {
		background: linear-gradient(95deg, #fd4536, #ec546e 35%, #c65f91 67%, #705fae);
		background-clip: text;
		color: transparent;
		transition: background-image 0.5s linear;
		cursor: pointer;
	}

	.gradient_text:hover {
		// background-image: linear-gradient(95deg, #fd4536, #ec546e 35%, #c65f91 67%, #705fae);
		opacity: .7;
	}

	.gradient_btn {
		background: linear-gradient(101deg, #fd4536, #ec546e 35%, #c65f91 67%, #705fae);
		color: #fff;
	}


	.bg-purple-dark {
		background: #99a9bf;
	}

	.bg-purple {
		background: #d3dce6;
	}

	.bg-purple-light {
		background: #e5e9f2;
	}

	.grid-content {
		border-radius: 4px;
		min-height: 36px;
	}

	.player {
		margin-top: 10px;
		border-radius: 35px;
		width: 97%;
		margin: 0 auto;

		.controlButton {
			margin: 0 auto;
			width: 80%;
			text-align: center;
		}

		.el-card__body {
			.aplayer {
				border-radius: 20px;
				box-shadow: 5px 5px 10px #d4dcdf;
			}
		}
	}

	.chooseBanner {


		text-align: center;
		margin-top: 40px;
		display: flex;
		flex-wrap: wrap;
		justify-content: center;

		div {
			display: flex;
			flex-direction: column;
			justify-content: center;
			margin: 10px;
		}
	}

	.graph {
		display: flex;
		flex-wrap: wrap;
		justify-content: space-around;
		padding-bottom: 50px;

		.el-card {
			border-radius: 35px;
			padding: 10px;
			margin-top: 20px;
		}

		.card {
			width: 46%;

			.card_dashBoard {
				width: 100%;
				display: flex;
				justify-content: center;

				ul {
					list-style: none;
					padding: 0;
					list-style: none;

					li {
						display: inline;
					}
				}
			}

			.phoneDashImg {
				width: 50px;
				height: 50px;
			}
		}

		.buttonGroup {
			display: flex;
			justify-content: center;

			/deep/.el-pagination {
				button {
					background-image: linear-gradient(95deg, #fd4536, #ec546e 35%, #c65f91 67%, #705fae);
					background-clip: text;
					background-size: auto;
					color: transparent;
					// transition: background-image 0.5s linear;
				}

				button:disabled {
					color: gray;
					background-image: none;
				}

			}

			.el-button--text {
				margin: 0px 10px;
			}

		}

	}
</style>
