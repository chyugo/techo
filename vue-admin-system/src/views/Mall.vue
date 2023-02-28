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

		<div style="text-align: center;margin-top:40px">

			<el-button v-show="showSetPrint" class="gradient_text" type="text" style="margin-right: 10px;" @click="setPrint">准备生成报告
			</el-button>
			<el-date-picker v-show="showChooseMonth" v-model="chooseMonth" type="month" placeholder="选择查找的月份" value-format="timestamp"
				@change="doChooseMonth">
			</el-date-picker>
			<el-button v-show="showDownload" class="gradient_text" type="text" style="margin:0px 10px;" @click="downloadImg">1.下载图片
			</el-button>
			<el-button  v-show="showDownload" class="gradient_text" type="text" style="margin-right: 10px;" @click="downloadDoc">2.下载文档模板
			</el-button>

			<el-button class="gradient_text" type="text" @click="doCopy" v-show="chooseMonthTitle"
				style="margin: 10px;">复制标题</el-button>


		</div>

		<div v-bind:style="bindgraph" class="graph">
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

		</div>
	</div>
</template>

<script>
	import aplayer from "vue-aplayer"
	import * as echarts from 'echarts'
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
				showDownload:false,
				showChooseMonth:false,
				showSetPrint:true,




			}
		},
		methods: {
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
					link.setAttribute('download', title+".docx")
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
				elink.download = "3-睡眠"+title;
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
				elink.download = "8-手机"+title;
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
				elink.download = "5-拉屎"+title;
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
				elink.download = "4-睡起"+title;
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
				elink.download = "7-体重"+title;
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
				elink.download = "6-心脏"+title;
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
				elink.download = "2-喝水"+title;
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
				elink.download = "1-反酸"+title;
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
					} else {
					}
				})
				
				http.post("water/searchMonth", {
					month: this.chooseMonth,
				}).then((response) => {
					if (response.data.data.waterList != null) {
						// 图表9
						const echarts16 = echarts.init(this.$refs.echarts16)
						const option16 = drawWater(response.data.data.waterList)
						echarts16.setOption(option16)
					} else {
					}
				})
				
				http.post("acid/searchMonth", {
					month: this.chooseMonth,
				}).then((response) => {
					if (response.data.data.acidList != null) {
						// 图表9
						const echarts18 = echarts.init(this.$refs.echarts18)
						const option18 = drawAcid(response.data.data.acidList)
						echarts18.setOption(option18)
					} else {
					}
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
					const option16= drawWater(response.data.data.waterList)
					echarts16.setOption(option16)
			
				})
			},
			
			acidPageChange(now) {
				this.acidPageNow = now
				http.post("acid/list", {
					acidPageNow: this.acidPageNow,
				}).then((response) => {
					const echarts18 = echarts.init(this.$refs.echarts18)
					const option18= drawAcid(response.data.data.acidList)
					echarts18.setOption(option18)
				})
			},
		},
		components: {
			aplayer
		},


		mounted() {
			this.getMusic()
			this.printType = this.$route.query.printType
			if (this.printType != null && this.printType == "print") {
				this.showChooseMonth = true
				this.showSetPrint = false
				this.bindgraph = "display: block !important;width: 92%;margin: 0% 2%;"
				this.bindcard =
					"margin: 20px auto;width: 1080px !important;border-radius: 10px !important;padding: 2% !important;"
			}
			http.post("phone/list", {
				phonePageNow: this.phonePageNow,
			}).then((response) => {
				if (response.data.data.phoneList != null) {
					// 图表3
					console.log(response.data.data)
					const echarts3 = echarts.init(this.$refs.echarts3)
					const option3 = drawphone(response.data.data.phoneList)
					echarts3.setOption(option3)
					this.phonePageCount = response.data.data.phonePageCount
				} else {
					// this.$message.error("手机空数据")
				}
			})


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
				} else {
					// this.$message.error("睡眠空数据")
				}
			})

			// 图表5
			http.post("sleep/sum", {}).then((response) => {
				if (response.data.data.sleepList != null) {
					console.log(response.data.data)
					const echarts5 = echarts.init(this.$refs.echarts5)
					const option5 = drawSleepSum(response.data.data.sleepList)
					echarts5.setOption(option5)
				} else {
					// this.$message.error("睡眠空数据")
				}
			})


			// 图表6
			http.post("phone/sum", {}).then((response) => {

				if (response.data.data.phoneList != null) {
					console.log(response.data.data)
					const echarts6 = echarts.init(this.$refs.echarts6)
					const option6 = drawPhoneSum(response.data.data.phoneList)
					echarts6.setOption(option6)
				} else {
					// this.$message.error("手机空数据")
				}
			})






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
				} else {
					// this.$message.error("体重空数据")
				}
			})


			// 图表8
			http.post("weight/sum", {}).then((response) => {
				if (response.data.data.weightList != null) {
					console.log(response.data.data)
					const echarts8 = echarts.init(this.$refs.echarts8)
					const option8 = drawWeightSum(response.data.data.weightList)
					echarts8.setOption(option8)
				} else {
					// this.$message.error("体重空数据")
				}
			})




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
				} else {
					// this.$message.error("律动空数据")
				}
			})

			// 图表10
			http.post("heart/sum", {}).then((response) => {

				if (response.data.data.heartList != null) {
					console.log(response.data.data)
					const echarts10 = echarts.init(this.$refs.echarts10)
					const option10 = drawHeartSum(response.data.data.heartList)
					echarts10.setOption(option10)
				} else {
					// this.$message.error("律动空数据")
				}

			})


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
				} else {
					// this.$message.error("拉屎空数据")
				}
			})

			// 图表13
			http.post("shit/sum", {}).then((response) => {
				if (response.data.data.shitList != null) {
					console.log(response.data.data)
					const echarts13 = echarts.init(this.$refs.echarts13)
					const option13 = drawShitSum(response.data.data.shitList)
					echarts13.setOption(option13)
				} else {
					// this.$message.error("拉屎空数据")
				}
			})

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
				} else {
					// this.$message.error("睡起空数据")
				}
			})


			// 图表15
			http.post("wake/sum", {}).then((response) => {
				if (response.data.data.wakeList != null) {
					console.log(response.data.data)
					const echarts15 = echarts.init(this.$refs.echarts15)
					const option15 = drawWakeSum(response.data.data.wakeList)
					echarts15.setOption(option15)
				} else {
					// this.$message.error("睡起空数据")
				}

			})
			
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
				} else {
					this.$message.error("喝水空数据")
				}
			
			})
			
			// 图表16
			http.post("water/sum", {
			}).then((response) => {
				if (response.data.data.waterList != null) {
					console.log(response.data.data)
					const echarts17 = echarts.init(this.$refs.echarts17)
					const option17 = drawWaterSum(response.data.data.waterList)
					echarts17.setOption(option17)
				} else {
					this.$message.error("喝水空数据")
				}
			
			})
			
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
				} else {
					this.$message.error("acid空数据")
				}
			
			})
			
			// 图表16
			http.post("acid/sum", {
			}).then((response) => {
				if (response.data.data.acidList != null) {
					console.log(response.data.data)
					const echarts19 = echarts.init(this.$refs.echarts19)
					const option19 = drawAcidSum(response.data.data.acidList)
					echarts19.setOption(option19)
				} else {
					this.$message.error("acid空数据")
				}
			
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

				.el-card__body {
					padding: 0;
				}
			}
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

	.graph {
		display: flex;
		flex-wrap: wrap;
		justify-content: space-around;

		.el-card {
			border-radius: 35px;
			padding: 10px;
			margin-top: 20px;
		}

		.card {
			width: 46%;
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
