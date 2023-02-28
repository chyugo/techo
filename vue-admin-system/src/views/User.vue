<template>
	<div>

		<div class="dobutton">
			<el-button class="topbutton" plain :icon="sleepIcon" @click="acid">反酸</el-button>
			<el-button class="topbutton" plain :icon="sleepIcon" @click="water">喝水</el-button>
			<el-button class="topbutton" plain :icon="sleepIcon" @click="sleep">入睡起床</el-button>
			<el-button class="topbutton" plain :icon="sleepIcon" @click="weight">体重</el-button>
			<el-button class="topbutton" plain :icon="sleepIcon" @click="shit">拉屎</el-button>
			<el-button class="topbutton" plain :icon="sleepIcon" @click="heart">心脏</el-button>
			<el-button class="topbutton" plain :icon="sleepIcon" @click="phone">手机</el-button>
		</div>
		<el-tabs v-model="activeName" class="tabs">
			<el-tab-pane label="上传音乐" name="all">
				<el-upload ref="upload" class="upload-demo" action="" drag multiple :limit=1 :with-credentials="true"
					:headers="uploadHeader" :http-request="doUpload">
					<i class="el-icon-upload"></i>
					<div class="el-upload__text">歌曲：将文件拖到此处，或<em>点击上传</em></div>
				</el-upload>
				<div style="margin: 5px;">
					标题 <el-input v-model="musicTitle" style="width: 40%;" placeholder=""></el-input>
				</div>
				<div style="margin: 5px;">
					艺术家 <el-input v-model="musicArtist" style="width: 40%;" placeholder=""></el-input>
				</div>
				歌曲封面
				<el-upload ref="uploadPic" class="avatar-uploader" action="" drag multiple :limit=1
					:with-credentials="true" :headers="uploadHeader" :http-request="doUploadPic">
					<img v-if="imageUrl" :src="imageUrl" class="avatar">
					<i v-else class="el-icon-plus avatar-uploader-icon"></i>
				</el-upload>

				歌词
				<el-upload ref="uploadLrc" class="avatar-uploader" action="" drag multiple :limit=1
					:with-credentials="true" :headers="uploadHeader" :http-request="doUploadLrc">
					<i class="el-icon-plus avatar-uploader-icon"></i>
				</el-upload>
				<div style="margin: 5px;">
					<el-button type="primary" @click="SubmitMusic">提 交</el-button>
				</div>
			</el-tab-pane>
			<el-tab-pane label="反酸" name="acid">
				<div style="display: flex;justify-content: space-around;">
					<el-form ref="form">
						<el-form-item label="录入日期">
							<el-col>
								<el-date-picker v-model="postdate" align="right" type="date" placeholder="选择日期"
									:picker-options="pickerOptions" value-format="timestamp"
									style="width: 70%;margin-right: 5px;">
								</el-date-picker>
								<el-button plain @click="chosetoday">选择今天</el-button>
								<el-button plain @click="choseyesterday">选择昨天</el-button>
							</el-col>
			
						</el-form-item>
						<el-form-item label="今日反酸">
							<span style="margin-right: 5px;">
								<el-input v-model="acidValue" style="width: 30%;" placeholder=""></el-input> 次
							</span>
							<el-button class="gradient_btn" type="primary" @click="acidRecord">添加</el-button>
							<el-col>
								<div style="width: 80%;margin: 0 auto;">
									<!-- <el-slider :format-tooltip="formatTooltip1" v-model="morningslider" :step="0.1"
									:show-tooltip="false"></el-slider> -->
			
			
								</div>
							</el-col>
			
						</el-form-item>
			
					</el-form>
			
				</div>
			
			</el-tab-pane>
			
			<el-tab-pane label="喝水" name="water">
				<div style="display: flex;justify-content: space-around;">
					<el-form ref="form">
						<el-form-item label="录入日期">
							<el-col>
								<el-date-picker v-model="postdate" align="right" type="date" placeholder="选择日期"
									:picker-options="pickerOptions" value-format="timestamp"
									style="width: 70%;margin-right: 5px;">
								</el-date-picker>
								<el-button plain @click="chosetoday">选择今天</el-button>
								<el-button plain @click="choseyesterday">选择昨天</el-button>
							</el-col>
			
						</el-form-item>
						<el-form-item label="本次喝水">
							<span style="margin-right: 5px;">
								<el-input v-model="waterValue" style="width: 30%;" placeholder=""></el-input> 毫升
							</span>
							<el-button class="gradient_btn" type="primary" @click="waterRecord">添加</el-button>
							<el-col>
								<div style="width: 80%;margin: 0 auto;">
									<!-- <el-slider :format-tooltip="formatTooltip1" v-model="morningslider" :step="0.1"
									:show-tooltip="false"></el-slider> -->
			
			
								</div>
							</el-col>
			
						</el-form-item>
			
					</el-form>
			
				</div>
			
			</el-tab-pane>
			
			<el-tab-pane label="入睡起床" name="sleep" style="margin: 0 auto;">
				<div style="display: flex;justify-content: space-around;">
					<el-form ref="form">
						<!-- 现在是{{todayDate}} -->
						<el-form-item label="录入日期">
							<el-col>
								<el-date-picker v-model="postdate" align="right" type="date" placeholder="选择日期"
									:picker-options="pickerOptions" value-format="timestamp"
									style="width: 70%;margin-right: 5px;">
								</el-date-picker>
								<el-button plain @click="chosetoday">选择今天</el-button>
								<el-button plain @click="choseyesterday">选择昨天</el-button>
							</el-col>
						</el-form-item>
						<el-form-item label="入睡时间"><span class="iconfont" style="font-size: 25px;">&#xe692;</span>
							{{sleepTime}}
							<el-col>
								<el-time-picker placeholder="选择时间" v-model="sleepTime" style="" format='HH:mm'
									value-format="H:m" style="margin-right: 5px;width: 80%;">
								</el-time-picker>

								<div>
									<el-button class="gradient_btn" type="primary" @click="fallasleepRecord"><span
											class="iconfont" style="margin-right: 10px;">&#xe692;</span>录入</el-button>
								</div>
							</el-col>

						</el-form-item>
						<el-form-item label="起床时间"><span class="iconfont" style="font-size: 25px;">&#xe626;</span>
							{{getTime}}
							<el-col>
								<el-time-picker placeholder="选择时间" v-model="getTime" style="" format='HH:mm'
									value-format="H:m" style="margin-right: 5px;width: 80%;">
								</el-time-picker>
								<div>
									<el-button class="gradient_btn" type="primary" @click="getRecord"><span
											class="iconfont" style="margin-right: 10px;">&#xe626;</span>录入</el-button>
								</div>
							</el-col>
						</el-form-item>

						<el-form-item label="睡眠时长(时)">
							<span style="margin-right: 10px;">{{sleepHour}}小时</span>
							<el-col>
								<div style="width: 80%;margin: 0 auto;">
									<el-slider :format-tooltip="sleepHourFormatToolTip" v-model="sleepHourSlider"
										:show-tooltip="false" :max=12></el-slider>
								</div>
							</el-col>
						</el-form-item>

						<el-form-item label="睡眠时长(分)">
							<span style="margin-right: 10px;">{{sleepMinute}}分钟</span>

							<el-col>
								<div style="width: 80%;margin: 0 auto;">
									<el-slider :format-tooltip="sleepMinuteFormatToolTip" v-model="sleepMinuteSlider"
										:show-tooltip="false" :max="59"></el-slider>
								</div>
							</el-col>
						</el-form-item>
						<div style="display: flex;justify-content: space-around;margin-bottom: 50px;">
							<el-button class="gradient_btn" type="primary" @click="sleepDurationRecord"
								style="width: 80%;">录入</el-button>
						</div>

					</el-form>

				</div>
			</el-tab-pane>

			<el-tab-pane label="体重" name="weight">
				<div style="display: flex;justify-content: space-around;">
					<el-form ref="form">
						<!-- 现在是{{todayDate}} -->
						<el-form-item label="录入日期">
							<el-col>
								<el-date-picker v-model="postdate" align="right" type="date" placeholder="选择日期"
									:picker-options="pickerOptions" value-format="timestamp"
									style="width: 70%;margin-right: 5px;">
								</el-date-picker>
								<el-button plain @click="chosetoday">选择今天</el-button>
								<el-button plain @click="choseyesterday">选择昨天</el-button>
							</el-col>

						</el-form-item>
						<el-form-item label="早餐"><span class="iconfont"
								style="font-size: 25px;margin-right: 10px;">&#xe610;</span>
							<span style="margin-right: 5px;">
								<el-input v-model="morningData" style="width: 30%;" placeholder=""></el-input> 斤
							</span>
							<el-button class="gradient_btn" type="primary" @click="morningRecord">录入</el-button>
							<el-col>
								<div style="width: 80%;margin: 0 auto;">
									<!-- <el-slider :format-tooltip="formatTooltip1" v-model="morningslider" :step="0.1"
										:show-tooltip="false"></el-slider> -->


								</div>
							</el-col>

						</el-form-item>
						<el-form-item label="午餐"><span class="iconfont"
								style="font-size: 25px;margin-right: 10px;">&#xe610;</span>
							<span style="margin-right: 5px;">
								<el-input v-model="noonData" style="width: 30%;" placeholder=""></el-input> 斤
							</span>
							<el-button class="gradient_btn" type="primary" @click="noonRecord">录入</el-button>
							<el-col>
								<!-- <div style="width: 80%;margin: 0 auto;">
									<el-slider :format-tooltip="formatTooltip2" v-model="noonslider" :step="0.1"
										:show-tooltip="false"></el-slider>
								</div> -->
							</el-col>

						</el-form-item>

						<el-form-item label="晚餐"><span class="iconfont"
								style="font-size: 25px;margin-right: 10px;">&#xe610;</span>
							<span style="margin-right: 5px;">
								<el-input v-model="nightData" style="width: 30%;" placeholder=""></el-input> 斤
							</span>
							<el-button class="gradient_btn" type="primary" @click="nightRecord">录入</el-button>
							<el-col>
								<!-- <div style="width: 80%;margin: 0 auto;">
									<el-slider :format-tooltip="formatTooltip3" v-model="nightslider" :step="0.1"
										:show-tooltip="false"></el-slider>
								</div> -->



							</el-col>

						</el-form-item>


					</el-form>

				</div>

			</el-tab-pane>
			<el-tab-pane label="拉屎" name="shit">
				<div style="display: flex;justify-content: space-around;">
					<el-form ref="form">
						<!-- 现在是{{todayDate}} -->
						<el-form-item label="录入日期">
							<el-col>
								<el-date-picker v-model="postdate" align="right" type="date" placeholder="选择日期"
									:picker-options="pickerOptions" value-format="timestamp"
									style="width: 70%;margin-right: 5px;">
								</el-date-picker>
								<el-button plain @click="chosetoday">选择今天</el-button>
								<el-button @click="choseyesterday">选择昨天</el-button>
							</el-col>

						</el-form-item>
						<el-form-item label="拉屎时间"><span class="iconfont"
								style="font-size: 25px;color: #f9b000">&#xe604;</span>
							{{shitTime}}
							<el-col>
								<el-time-picker placeholder="选择时间" v-model="shitTime" style="" format='HH:mm'
									style="margin-right: 5px;width: 70%;" value-format="H:m">
								</el-time-picker>
								<el-button class="gradient_btn" type="primary" @click="shitRecord">录入</el-button>
							</el-col>

						</el-form-item>



					</el-form>

				</div>
			</el-tab-pane>
			<el-tab-pane label="心脏" name="heart">
				<div style="display: flex;justify-content: space-around;">
					<el-form ref="form">
						<!-- 现在是{{todayDate}} -->
						<el-form-item label="录入日期">
							<el-col>
								<el-date-picker v-model="postdate" align="right" type="date" placeholder="选择日期"
									:picker-options="pickerOptions" value-format="timestamp"
									style="width: 70%;margin-right: 5px;">
								</el-date-picker>
								<el-button plain @click="chosetoday">选择今天</el-button>
								<el-button plain @click="choseyesterday">选择昨天</el-button>
							</el-col>

						</el-form-item>

						<el-form-item label="检测的次数"><span class="iconfont"
								style="font-size: 25px;margin-right: 10px;color: red">&#xe9a0;</span>
							<span style="margin-right: 10px;">{{heartall}}次</span>

							<el-col>
								<div style="width: 80%;margin: 0 auto;">
									<el-slider :format-tooltip="heartFormatToolTip" v-model="heartSlider"
										:show-tooltip="false" :max="70"></el-slider>
								</div>
							</el-col>

						</el-form-item>

						<el-form-item label="检测到异常的次数"><span class="iconfont"
								style="font-size: 25px;margin-right: 10px;color: orange">&#xe6eb;</span>
							<span style="margin-right: 10px;">{{heartAbnomalAll}}次</span>
							<el-col>
								<div style="width: 80%;margin: 0 auto;">
									<el-slider :format-tooltip="heartAbnormalFormatToolTip"
										v-model="heartAbnormalSlider" :show-tooltip="false" :max="20"></el-slider>
								</div>
							</el-col>
						</el-form-item>
						<div style="display: flex;justify-content: space-around">
							<el-button class="gradient_btn" type="primary" @click="heartRecord" style="width: 80%;">录入
							</el-button>
						</div>

					</el-form>

				</div>
			</el-tab-pane>
			<el-tab-pane label="手机" name="phone">
				<div style="display: flex;justify-content: space-around;">
					<el-form ref="form">
						<!-- 现在是{{todayDate}} -->
						<div>
							<el-form-item label="录入日期">

								<el-col>
									<el-date-picker v-model="postdate" align="right" type="date" placeholder="选择日期"
										:picker-options="pickerOptions" value-format="timestamp"
										style="width: 70%;margin-right: 5px;">
									</el-date-picker>
									<el-button plain @click="chosetoday">选择今天</el-button>
									<el-button type="success" @click="choseyesterday">选择昨天</el-button>
								</el-col>

							</el-form-item>
						</div>

						<el-form-item label="时长(时)"> <span class="iconfont"
								style="font-size: 25px;margin-right: 10px;">&#xe616;</span>
							<span style="margin-right: 10px;">{{phoneHour}}小时</span>

							<el-col>
								<div style="width: 80%;margin: 0 auto;">
									<el-slider :format-tooltip="phoneHourFormatToolTip" v-model="phoneHourSlider"
										:show-tooltip="false" :max=24></el-slider>
								</div>
							</el-col>
						</el-form-item>

						<el-form-item label="时长(分)"> <span class="iconfont"
								style="font-size: 25px;margin-right: 10px;">&#xe616;</span>
							<span style="margin-right: 10px;">{{phoneMinute}}分钟</span>

							<el-col>
								<div style="width: 80%;margin: 0 auto;">
									<el-slider :format-tooltip="phoneMinuteFormatToolTip" v-model="phoneMinuteSlider"
										:show-tooltip="false" :max="59"></el-slider>
								</div>
							</el-col>
						</el-form-item>
						<div style="display: flex;justify-content: space-around">
							<el-button class="gradient_btn" type="primary" @click="phoneRecord" style="width: 80%;">录入
							</el-button>
						</div>

					</el-form>

				</div>
			</el-tab-pane>


		</el-tabs>
	</div>
</template>

<script>
	import {
		Loading
	} from 'element-ui';
	import http from '../utils/request.js'
	import '../utils/config.js'
	import basePicUrl from '../utils/config.js';
	export default {
		data() {
			return {
				activeName: 'acid',
				sleepIcon: "el-icon-warning-outline",

				postdate: parseInt(new Date().getTime()),
				shitTime: '',
				getTime: '',
				sleepTime: '',
				sleepHour: '',
				sleepMinute: '',
				imageUrl: '',

				pickerOptions: {
					shortcuts: [{
						text: '今天',
						onClick(picker) {
							picker.$emit('pick', new Date());
						}
					}, {
						text: '昨天',
						onClick(picker) {
							const date = new Date();
							date.setTime(date.getTime() - 3600 * 1000 * 24);
							picker.$emit('pick', date);
						}
					}, {
						text: '一周前',
						onClick(picker) {
							const date = new Date();
							date.setTime(date.getTime() - 3600 * 1000 * 24 * 7);
							picker.$emit('pick', date);
						}
					}]
				},
				morningData: null,
				// morningslider: 0,
				noonData: null,
				// noonslider: 0,
				nightData: null,
				// nightslider: 0,

				heartall: 0,
				heartSlider: 0,
				heartAbnomalAll: 0,
				heartAbnormalSlider: 0,
				musicTitle: '',
				musicArtist: '',
				filePath: '',
				oldFileName: '',
				picPath: '',
				oldPicName: '',
				lrcPath: '',
				oldLrcName: '',

				todayDate: "",

				uploadHeader: {
					"Authorization": $cookies.get("cookie"),
					"Access-Control-Allow-Origin": "*",
				},
				phoneHour: 0,
				phoneHourSlider: 0,
				phoneMinute: 0,
				phoneMinuteSlider: 0,
				sleepHourSlider: 0,
				sleepMinute: 0,
				sleepMinuteSlider: 0,
				waterValue: 0,
				acidValue: 0,
			}
		},
		methods: {
			doUpload(item) {
				const loadingOptions = {
					fullscreen: true,
					lock: true,
					text: 'Loading',
					background: 'rgba(255, 255, 255, 0.9)'
				}
				Loading.service(loadingOptions);

				console.log("item", item)
				const file = item.file
				var form = new FormData()
				form.append("file", file)
				console.log("form", form)
				http.post("music/uploadMusic", form, {
					headers: {
						'Content-Type': "multipart/form-data"
					}
				}).then((response) => {
					Loading.service(loadingOptions).close()
					if (response.data.err_code != 0) {
						this.$message.error(response.data.err_msg);
					} else {
						this.$notify({
							title: '填写歌曲名和歌手名，上传歌曲封面，点击提交按钮确认上传',
							position: 'bottom-left',
							type: 'success',
							duration: 2000,
							showClose: false,
						});
						this.musicTitle = response.data.data.baseName
						this.filePath = response.data.data.filePath
						this.oldFileName = response.data.data.oldFileName

						console.log(response)
					}
				}).catch(reason => {})
			},
			doUploadPic(item) {
				const loadingOptions = {
					fullscreen: true,
					lock: true,
					text: 'Loading',
					background: 'rgba(255, 255, 255, 0.9)'
				}
				Loading.service(loadingOptions);
				console.log("item", item)
				const file = item.file
				var form = new FormData()
				form.append("file", file)
				console.log("form", form)
				http.post("music/uploadPic", form, {
					headers: {
						'Content-Type': "multipart/form-data"
					}
				}).then((response) => {
					Loading.service(loadingOptions).close()
					if (response.data.err_code != 0) {
						this.$message.error(response.data.err_msg);
					} else {
						this.$message({
							message: '若有歌词，继续上传。没有可提交。',
							type: 'success',
						})
						this.picPath = response.data.data.picPath
						this.imageUrl = basePicUrl + "/more_static/" + response.data.data.picPath
						this.oldPicName = response.data.data.oldPicName
						console.log(response)
					}
				}).catch(reason => {})
			},
			doUploadLrc(item) {
				const loadingOptions = {
					fullscreen: true,
					lock: true,
					text: 'Loading',
					background: 'rgba(255, 255, 255, 0.9)'
				}
				Loading.service(loadingOptions);
				console.log("item", item)
				const file = item.file
				var form = new FormData()
				form.append("file", file)
				console.log("form", form)
				http.post("music/uploadLrc", form, {
					headers: {
						'Content-Type': "multipart/form-data"
					}
				}).then((response) => {
					Loading.service(loadingOptions).close()
					if (response.data.err_code != 0) {
						this.$message.error(response.data.err_msg);
					} else {
						this.$message({
							message: '现在已可上传',
							type: 'success',
						})
						this.lrcPath = response.data.data.lrcPath
						this.oldLrcName = response.data.data.oldLrcName
						console.log(response)
					}
				}).catch(reason => {})
			},
			SubmitMusic() {
				const loadingOptions = {
					fullscreen: true,
					lock: true,
					text: 'Loading',
					background: 'rgba(255, 255, 255, 0.9)'
				}
				Loading.service(loadingOptions);
				http.post("music/submitMusic", {
					oldFileName: this.oldFileName,
					filePath: this.filePath,
					title: this.musicTitle,
					artist: this.musicArtist,
					picPath: this.picPath,
					oldPicName: this.oldPicName,
					lrcPath: this.lrcPath,
					oldLrcName: this.oldLrcName,
				}, {}).then((response) => {
					Loading.service(loadingOptions).close()
					if (response.data.err_code != 0) {
						this.$message.error(response.data.err_msg);
					} else {
						this.$notify({
							title: '上传成功',
							position: 'bottom-left',
							type: 'success',
							duration: 1400,
							showClose: false,
						});
						this.uploadresumeForm = false
						this.oldFileName = ''
						this.filePath = ''
						this.musicTitle = ''
						this.musicArtist = ''
						this.picPath = ''
						this.oldPicName = ''
						this.lrcPath = ''
						this.imageUrl = ''
						this.oldLrcName = ''
						console.log(response)
						this.$refs.upload.clearFiles();
						this.$refs.uploadPic.clearFiles();
						this.$refs.uploadLrc.clearFiles();
					}
				}).catch(reason => {})
				this.uploadresumeForm = false
			},
			sleep() {
				this.activeName = "sleep"
			},
			weight() {
				this.activeName = "weight"
			},
			shit() {
				this.activeName = "shit"
			},
			heart() {
				this.activeName = "heart"
			},
			phone() {
				this.activeName = "phone"
			},
			water() {
				this.activeName = "water"
			},
			acid(){
				this.activeName = "acid"
			},
			chosetoday() {
				this.postdate = parseFloat((new Date(Date.now()).getTime()) / 1000).toFixed(0) + '000'
				console.log("吖", this.postdate)
				console.log("初始", parseInt(new Date().getTime() / 1000) + '000')
			},
			choseyesterday() {
				this.postdate = parseFloat((new Date(Date.now() - 24 * 60 * 60 * 1000).getTime()) / 1000).toFixed(0) +
					'000'
				console.log("吖", this.postdate)
				console.log("初始", parseInt(new Date().getTime() / 1000) + '000')
			},
			onSubmit() {

			},
			formatTooltip1(val) {
				this.morningData = val + 100
				return (val + 100);
			},
			formatTooltip2(val) {
				this.noonData = val + 100
				return (val + 100);
			},
			formatTooltip3(val) {
				console.log("val", val)
				this.nightData = val + 100
				return (val + 100);
			},
			heartFormatToolTip(val) {
				console.log("val", val)
				this.heartall = val
				return (val);
			},
			heartAbnormalFormatToolTip(val) {
				console.log("val", val)
				this.heartAbnomalAll = val
				return (val);
			},
			phoneHourFormatToolTip(val) {
				console.log("val", val)
				this.phoneHour = val
				return (val);
			},
			phoneMinuteFormatToolTip(val) {
				console.log("val", val)
				this.phoneMinute = val
				return (val);
			},
			sleepHourFormatToolTip(val) {
				console.log("val", val)
				this.sleepHour = val
				return (val);
			},
			sleepMinuteFormatToolTip(val) {
				console.log("val", val)
				this.sleepMinute = val
				return (val);
			},
			morningRecord() {
				http.post("weight/record", {
					morningWeight: this.morningData * 10,
					recordDate: this.postdate,
				}).then((response) => {
					if (response.data.err_code != 0) {
						this.$notify.error({
							title: response.data.err_msg,
							position: 'bottom-left',
						});
					} else {
						this.$notify({
							title: '添加成功',
							position: 'bottom-left',
							type: 'success',
							duration: 1400,
							showClose: false,
						});
					}
				})
			},
			noonRecord() {
				http.post("weight/record", {
					noonWeight: this.noonData * 10,
					recordDate: this.postdate,
				}).then((response) => {
					if (response.data.err_code != 0) {
						this.$notify.error({
							title: response.data.err_msg,
							position: 'bottom-left',
						});
					} else {
						this.$notify({
							title: '添加成功',
							position: 'bottom-left',
							type: 'success',
							duration: 1400,
							showClose: false,
						});
					}
				})
			},
			nightRecord() {
				http.post("weight/record", {
					nightWeight: this.nightData * 10,
					recordDate: this.postdate,
				}).then((response) => {
					if (response.data.err_code != 0) {
						this.$notify.error({
							title: response.data.err_msg,
							position: 'bottom-left',
						});
					} else {
						this.$notify({
							title: '添加成功',
							position: 'bottom-left',
							type: 'success',
							duration: 1400,
							showClose: false,
						});
					}
				})
			},
			shitRecord() {
				http.post("shit/record", {
					shitTime: this.shitTime,
					recordDate: this.postdate,
				}).then((response) => {
					if (response.data.err_code != 0) {
						this.$notify.error({
							title: response.data.err_msg,
							position: 'bottom-left',
						});
					} else {
						this.$notify({
							title: '添加成功',
							position: 'bottom-left',
							type: 'success',
							duration: 1400,
							showClose: false,
						});
					}
				})
			},
			heartRecord() {
				http.post("heart/record", {
					heartall: this.heartall,
					heartAbnomalAll: this.heartAbnomalAll,
					recordDate: this.postdate,
				}).then((response) => {
					if (response.data.err_code != 0) {
						this.$notify.error({
							title: response.data.err_msg,
							position: 'bottom-left',
						});
					} else {
						this.$notify({
							title: '添加成功',
							position: 'bottom-left',
							type: 'success',
							duration: 1400,
							showClose: false,
						});
					}
				})
			},

			phoneRecord() {
				http.post("phone/record", {
					phoneHour: this.phoneHour,
					phoneMinute: this.phoneMinute,
					recordDate: this.postdate,
				}).then((response) => {
					if (response.data.err_code != 0) {
						this.$notify.error({
							title: response.data.err_msg,
							position: 'bottom-left',
						});
					} else {
						this.$notify({
							title: '添加成功',
							position: 'bottom-left',
							type: 'success',
							duration: 1400,
							showClose: false,
						});
					}
				})
			},
			fallasleepRecord() {
				http.post("sleep/record", {
					sleepTime: this.sleepTime,
					recordDate: this.postdate,
				}).then((response) => {
					if (response.data.err_code != 0) {
						this.$notify.error({
							title: response.data.err_msg,
							position: 'bottom-left',
						});
					} else {
						this.$notify({
							title: '添加成功',
							position: 'bottom-left',
							type: 'success',
							duration: 1400,
							showClose: false,
						});
					}
				})
			},
			getRecord() {
				http.post("sleep/record", {
					getTime: this.getTime,
					recordDate: this.postdate,
				}).then((response) => {
					if (response.data.err_code != 0) {
						this.$notify.error({
							title: response.data.err_msg,
							position: 'bottom-left',
						});
					} else {
						this.$notify({
							title: '添加成功',
							position: 'bottom-left',
							type: 'success',
							duration: 1400,
							showClose: false,
						});
					}
				})
			},
			sleepDurationRecord() {
				http.post("sleep/record", {
					sleepHour: this.sleepHour,
					sleepMinute: this.sleepMinute,
					recordDate: this.postdate,
				}).then((response) => {
					if (response.data.err_code != 0) {
						this.$notify.error({
							title: response.data.err_msg,
							position: 'bottom-left',
						});
					} else {
						this.$notify({
							title: '添加成功',
							position: 'bottom-left',
							type: 'success',
							duration: 1400,
							showClose: false,
						});
					}
				})
			},
			waterRecord() {
				http.post("water/record", {
					water: this.waterValue,
					recordDate: this.postdate,
				}).then((response) => {
					if (response.data.err_code != 0) {
						this.$notify.error({
							title: response.data.err_msg,
							position: 'bottom-left',
						});
					} else {
						this.$notify({
							title: '添加成功',
							position: 'bottom-left',
							type: 'success',
							duration: 1400,
							showClose: false,
						});
					}
				})
			},
			acidRecord() {
				http.post("acid/record", {
					acid: this.acidValue,
					recordDate: this.postdate,
				}).then((response) => {
					if (response.data.err_code != 0) {
						this.$notify.error({
							title: response.data.err_msg,
							position: 'bottom-left',
						});
					} else {
						this.$notify({
							title: '添加成功',
							position: 'bottom-left',
							type: 'success',
							duration: 1400,
							showClose: false,
						});
					}
				})
			},

		},
		mounted() {
			let d = new Date();
			this.todayDate = (d.getMonth() + 1) + "月" + d.getDate() + "日" + d.getHours() + "时" + d.getMinutes() + "分";

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
		background-image: linear-gradient(95deg, #fd4536, #ec546e 35%, #c65f91 67%, #705fae);
		color: #fff;
		border-color: white;
	}

	.gradient_btn:hover {
		opacity: .7;
	}

	.tabs {
		margin: 5px 30px;

	}

	.dobutton {
		margin: 5px 30px;
	}

	.dobutton {
		.topbutton {
			margin: 3px 3px;
		}

	}


	.el-col {
		button {
			margin-top: 5px;
		}

	}


	/deep/.el-upload-dragger {
		width: 200px !important;
		height: 200px !important;
	}

	.avatar-uploader .el-upload {
		border: 1px dashed #d9d9d9;
		border-radius: 6px;
		cursor: pointer;
		position: relative;
		overflow: hidden;
	}

	.avatar-uploader .el-upload:hover {
		border-color: #409EFF;
	}

	.avatar-uploader-icon {
		font-size: 28px;
		color: #8c939d;
		width: 200px;
		height: 200px;
		line-height: 200px;
		text-align: center;
	}

	.avatar {
		width: 200px;
		height: 200px;
		display: block;
	}
</style>
