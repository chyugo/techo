<template>
	<div>

		<div class="dobutton">
			<el-button class="topbutton" plain :icon="sleepIcon" @click="acid">反酸表</el-button>
			<el-button class="topbutton" plain :icon="sleepIcon" @click="water">喝水表</el-button>
			<el-button class="topbutton" plain :icon="sleepIcon" @click="sleep">睡眠表</el-button>
			<el-button class="topbutton" plain :icon="sleepIcon" @click="weight">体重表</el-button>
			<el-button class="topbutton" plain :icon="sleepIcon" @click="shit">拉屎表</el-button>
			<el-button class="topbutton" plain :icon="sleepIcon" @click="heart">心脏表</el-button>
			<el-button class="topbutton" plain :icon="sleepIcon" @click="phone">手机表</el-button>

		</div>
		<el-tabs v-model="activeName" class="tabs" :before-leave="beforeLeave">
			<el-tab-pane label="音乐表" name="music">
				<transition name="el-fade-in-linear">
					<div style="width: 100%;margin:0 auto" v-show="musicTableShow">
						<el-table :data="musicTableData" style="width: 100%" :border="true" :stripe="true">
							<el-table-column label="歌曲标题">
								<template slot-scope="scope">
									<span>{{ scope.row.title }}</span>
								</template>
							</el-table-column>
							<el-table-column label="艺术家">
								<template slot-scope="scope">
									<span>{{ scope.row.artist }}</span>
								</template>
							</el-table-column>
							<el-table-column label="封面图片">
								<template slot-scope="scope">
									<span>{{ scope.row.pic }}</span>
								</template>
							</el-table-column>

							<el-table-column label="操作">
								<template slot-scope="scope">
									<el-button size="mini" type="danger"
										@click="musicHandleDelete(scope.$index, scope.row)">删除
									</el-button>
								</template>
							</el-table-column>
						</el-table>
						<div style="text-align: center;">
							<el-pagination background layout="prev, pager, next" :pager-count=5
								:current-page="musicPageNow" :total="musicPageCount"
								@current-change="musicCurrentChange" @prev-click="musicCurrentChange"
								@next-click="musicCurrentChange" style="margin: 20px 0px;">
							</el-pagination>
						</div>
					</div>

				</transition>

			</el-tab-pane>
			<el-tab-pane label="反酸表" name="acid">
				<transition name="el-fade-in-linear">
					<div style="width: 100%;margin:0 auto" v-show="acidTableShow">
						<el-table :data="acidTableData" style="width: 100%" :border="true" :stripe="true">
							<el-table-column label="日期">
								<template slot-scope="scope">
									<span>{{ scope.row.RecordDate }}</span>
								</template>
							</el-table-column>
							<el-table-column label="反酸次数">
								<template slot-scope="scope">
									<span>{{ scope.row.Acid }}</span>
								</template>
							</el-table-column>
							<el-table-column label="评级">
								<template slot-scope="scope">
									<span>{{ scope.row.Judge =="1"?"没看见有":scope.row.Judge =="2"?"疑似有":scope.row.Judge =="3"?"纸上有":"明显有", }}</span>
								</template>
							</el-table-column>
							 <el-table-column label="操作">
								<template slot-scope="scope">
									<el-button size="mini" type="danger"
										@click="acidHandleDelete(scope.$index, scope.row)">删除
									</el-button>
								</template>
							</el-table-column>
						</el-table>
						<div style="text-align: center;">
							<el-pagination background layout="prev, pager, next" :pager-count=5
								:current-page="acidPageNow" :total="acidPageCount" @current-change="acidCurrentChange"
								@prev-click="acidCurrentChange" @next-click="acidCurrentChange"
								style="margin: 20px 0px;">
							</el-pagination>
						</div>
					</div>
				</transition>
			</el-tab-pane>
			<el-tab-pane label="喝水表" name="water">
				<transition name="el-fade-in-linear">
					<div style="width: 100%;margin:0 auto" v-show="waterTableShow">
						<el-table :data="waterTableData" style="width: 100%" :border="true" :stripe="true">
							<el-table-column label="日期">
								<template slot-scope="scope">
									<span>{{ scope.row.RecordDate }}</span>
								</template>
							</el-table-column>
							<el-table-column label="水量">
								<template slot-scope="scope">
									<span>{{ scope.row.Water }}</span>
								</template>
							</el-table-column>
							s <el-table-column label="操作">
								<template slot-scope="scope">
									<el-button size="mini" type="danger"
										@click="waterHandleDelete(scope.$index, scope.row)">删除
									</el-button>
								</template>
							</el-table-column>
						</el-table>
						<div style="text-align: center;">
							<el-pagination background layout="prev, pager, next" :pager-count=5
								:current-page="waterPageNow" :total="waterPageCount"
								@current-change="waterCurrentChange" @prev-click="waterCurrentChange"
								@next-click="waterCurrentChange" style="margin: 20px 0px;">
							</el-pagination>
						</div>
					</div>
				</transition>
			</el-tab-pane>
			
			<el-tab-pane label="睡眠表" name="sleep" style="margin: 0 auto;">
				<transition name="el-fade-in-linear">
					<div style="width: 100%;margin:0 auto" v-show="sleepTableShow">
						<el-table :data="sleepTableData" style="width: 100%" :border="true" :stripe="true">
							<!-- <el-table-column label="id" min-width="180">
								<template slot-scope="scope">
									<span style="margin-left: 10px">{{ scope.row.id }}</span>
								</template>
							</el-table-column> -->
							<el-table-column label="日期">
								<template slot-scope="scope">
									<span>{{ scope.row.RecordDate }}</span>
								</template>
							</el-table-column>
							<el-table-column label="sleep_time">
								<template slot-scope="scope">
									<span>{{ scope.row.SleepTime }}</span>
								</template>
							</el-table-column>
							<el-table-column label="get_time">
								<template slot-scope="scope">
									<span>{{ scope.row.GetTime }}</span>
								</template>
							</el-table-column>
							<el-table-column label="active">
								<template slot-scope="scope">
									<span>{{ scope.row.Active }}</span>
								</template>
							</el-table-column>
							<el-table-column label="操作">
								<template slot-scope="scope">
									<el-button size="mini" type="danger"
										@click="sleepHandleDelete(scope.$index, scope.row)">删除
									</el-button>
								</template>
							</el-table-column>
						</el-table>
						<div style="text-align: center;">
							<el-pagination background layout="prev, pager, next" :pager-count=5
								:current-page="sleepPageNow" :total="sleepPageCount"
								@current-change="sleepCurrentChange" @prev-click="sleepCurrentChange"
								@next-click="sleepCurrentChange" style="margin: 20px 0px;">
							</el-pagination>
						</div>
					</div>

				</transition>

			</el-tab-pane>

			<el-tab-pane label="体重表" name="weight">
				<transition name="el-fade-in-linear">
					<div style="width: 100%;margin:0 auto" v-show="weightTableShow">
						<el-table :data="weightTableData" style="width: 100%" :border="true" :stripe="true">
							<!-- <el-table-column label="id" min-width="180">
								<template slot-scope="scope">
									<span style="margin-left: 10px">{{ scope.row.id }}</span>
								</template>
							</el-table-column> -->
							<el-table-column label="日期">
								<template slot-scope="scope">
									<span>{{ scope.row.RecordDate }}</span>
								</template>
							</el-table-column>
							<el-table-column label="morning_weight">
								<template slot-scope="scope">
									<span>{{ scope.row.MorningWeight }}</span>
								</template>
							</el-table-column>
							<el-table-column label="noon_weight">
								<template slot-scope="scope">
									<span>{{ scope.row.NoonWeight }}</span>
								</template>
							</el-table-column>
							<el-table-column label="night_weight">
								<template slot-scope="scope">
									<span>{{ scope.row.NightWeight }}</span>
								</template>
							</el-table-column>
							<el-table-column label="操作">
								<template slot-scope="scope">
									<el-button size="mini" type="danger"
										@click="weightHandleDelete(scope.$index, scope.row)">删除
									</el-button>
								</template>
							</el-table-column>
						</el-table>
						<div style="text-align: center;">
							<el-pagination background layout="prev, pager, next" :pager-count=5
								:current-page="weightPageNow" :total="weightPageCount"
								@current-change="weightCurrentChange" @prev-click="weightCurrentChange"
								@next-click="weightCurrentChange" style="margin: 20px 0px;">
							</el-pagination>
						</div>
					</div>

				</transition>

			</el-tab-pane>
			<el-tab-pane label="拉屎表" name="shit">
				<transition name="el-fade-in-linear">
					<div style="width: 100%;margin:0 auto" v-show="shitTableShow">
						<el-table :data="shitTableData" style="width: 100%" :border="true" :stripe="true">
							<!-- 	<el-table-column label="id" min-width="180">
								<template slot-scope="scope">
									<span style="margin-left: 10px">{{ scope.row.id }}</span>
								</template>
							</el-table-column> -->
							<el-table-column label="日期">
								<template slot-scope="scope">
									<span>{{ scope.row.RecordDate }}</span>
								</template>
							</el-table-column>
							<el-table-column label="shit_time">
								<template slot-scope="scope">
									<span>{{ scope.row.ShitTime }}</span>
								</template>
							</el-table-column>
							<el-table-column label="操作">
								<template slot-scope="scope">
									<el-button size="mini" type="danger"
										@click="shitHandleDelete(scope.$index, scope.row)">
										删除
									</el-button>
								</template>
							</el-table-column>
						</el-table>
						<div style="text-align: center;">
							<el-pagination background layout="prev, pager, next" :pager-count=5
								:current-page="shitPageNow" :total="shitPageCount" @current-change="shitCurrentChange"
								@prev-click="shitCurrentChange" @next-click="shitCurrentChange"
								style="margin: 20px 0px;">
							</el-pagination>
						</div>
					</div>

				</transition>
			</el-tab-pane>
			<el-tab-pane label="心脏表" name="heart">
				<transition name="el-fade-in-linear">
					<div style="width: 100%;margin:0 auto" v-show="heartTableShow">
						<el-table :data="heartTableData" style="width: 100%" :border="true" :stripe="true">
							<!-- <el-table-column label="id" min-width="180">
								<template slot-scope="scope">
									<span style="margin-left: 10px">{{ scope.row.id }}</span>
								</template>
							</el-table-column> -->
							<el-table-column label="日期">
								<template slot-scope="scope">
									<span>{{ scope.row.date }}</span>
								</template>
							</el-table-column>
							<el-table-column label="normal">
								<template slot-scope="scope">
									<span>{{ scope.row.normal }}</span>
								</template>
							</el-table-column>
							<el-table-column label="abnormal">
								<template slot-scope="scope">
									<span>{{ scope.row.abnormal }}</span>
								</template>
							</el-table-column>

							<el-table-column label="操作">
								<template slot-scope="scope">
									<el-button size="mini" type="danger"
										@click="heartHandleDelete(scope.$index, scope.row)">删除
									</el-button>
								</template>
							</el-table-column>
						</el-table>
						<div style="text-align: center;">
							<el-pagination background layout="prev, pager, next" :pager-count=5
								:current-page="heartPageNow" :total="heartPageCount"
								@current-change="heartCurrentChange" @prev-click="heartCurrentChange"
								@next-click="heartCurrentChange" style="margin: 20px 0px;">
							</el-pagination>
						</div>
					</div>

				</transition>

			</el-tab-pane>
			<el-tab-pane label="手机表" name="phone">
				<transition name="el-fade-in-linear">
					<div style="width: 100%;margin:0 auto" v-show="phoneTableShow">
						<el-table :data="phoneTableData" style="width: 100%" :border="true" :stripe="true">
							<!-- <el-table-column label="id" min-width="180">
								<template slot-scope="scope">
									<span style="margin-left: 10px">{{ scope.row.id }}</span>
								</template>
							</el-table-column> -->
							<el-table-column label="日期">
								<template slot-scope="scope">
									<span>{{ scope.row.RecordDate }}</span>
								</template>
							</el-table-column>
							<el-table-column label="phoneTime">
								<template slot-scope="scope">
									<span>{{ scope.row.PhoneMinute }}</span>
								</template>
							</el-table-column>
							s <el-table-column label="操作">
								<template slot-scope="scope">
									<el-button size="mini" type="danger"
										@click="phoneHandleDelete(scope.$index, scope.row)">删除
									</el-button>
								</template>
							</el-table-column>
						</el-table>
						<div style="text-align: center;">
							<el-pagination background layout="prev, pager, next" :pager-count=5
								:current-page="phonePageNow" :total="phonePageCount"
								@current-change="phoneCurrentChange" @prev-click="phoneCurrentChange"
								@next-click="phoneCurrentChange" style="margin: 20px 0px;">
							</el-pagination>
						</div>
					</div>

				</transition>

			</el-tab-pane>


			

		</el-tabs>

	</div>
</template>

<script>
	import http from '../utils/request.js'
	export default {
		data() {
			return {
				sleepTableShow: false,
				weightTableShow: false,
				shitTableShow: false,
				heartTableShow: false,
				phoneTableShow: false,
				waterTableShow: false,
				musicTableShow: false,
				acidTableShow: true,
				activeName: 'acid',
				sleepIcon: "el-icon-warning-outline",
				sleepTableData: [],
				sleepPageNow: 1,
				sleepPageCount: 0,

				weightTableData: [],
				weightPageNow: 1,
				weightPageCount: 0,

				shitTableData: [],
				shitPageNow: 1,
				shitPageCount: 0,

				heartTableData: [],
				heartPageNow: 1,
				heartPageCount: 0,

				phoneTableData: [],
				phonePageNow: 1,
				phonePageCount: 0,


				musicTableData: [],
				musicPageNow: 1,
				musicPageCount: 0,

				waterTableData: [],
				waterPageNow: 1,
				waterPageCount: 0,


				acidTableData: [],
				acidPageNow: 1,
				acidPageCount: 0,
			}
		},
		methods: {
			beforeLeave(activeName, oldActiveName) {
				this.sleepTableShow = false
				this.weightTableShow = false
				this.shitTableShow = false
				this.heartTableShow = false
				this.phoneTableShow = false
				this.musicTableShow = false
				this.waterTableShow = false
				this.acidTableShow = false
				console.log(activeName, oldActiveName)
				if (activeName == "sleep") {
					this.sleepTableShow = true
				}
				if (activeName == "weight") {
					this.weightTableShow = true
				}
				if (activeName == "shit") {
					this.shitTableShow = true
				}
				if (activeName == "heart") {
					this.heartTableShow = true
				}
				if (activeName == "phone") {
					this.phoneTableShow = true
				}
				if (activeName == "music") {
					this.musicTableShow = true
				}
				if (activeName == "water") {
					this.waterTableShow = true
				}
				if (activeName == "acid") {
					this.acidTableShow = true
				}

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
			acid() {
				this.activeName = "acid"
			},
			flushSleep() {
				http.post("sleep/find", {
					sleepPageNow: this.sleepPageNow,
				}).then((response) => {
					console.log("response.data.data", response.data.data)
					if (response.data.data.sleepList != null) {
						this.sleepTableData = response.data.data.sleepList
						this.sleepPageCount = response.data.data.sleepPageCount
						this.sleepTable = true
					} else {
						this.$message.error("睡眠空数据")
					}

				})
			},
			flushWeight() {
				http.post("weight/find", {
					weightPageNow: this.weightPageNow,
				}).then((response) => {
					console.log("response.data.data", response.data.data)
					if (response.data.data.weightList != null) {
						this.weightTableData = response.data.data.weightList
						this.weightPageCount = response.data.data.weightPageCount
					} else {
						this.$message.error("体重空数据")
					}
				})
			},
			flushShit() {
				http.post("shit/find", {
					shitPageNow: this.shitPageNow,
				}).then((response) => {
					console.log("response.data.data", response.data.data)
					if (response.data.data.shitList != null) {
						this.shitTableData = response.data.data.shitList
						this.shitPageCount = response.data.data.shitPageCount
					} else {
						this.$message.error("拉屎空数据")
					}
				})
			},
			flushHeart() {
				http.post("heart/find", {
					heartPageNow: this.heartPageNow,
				}).then((response) => {
					console.log("response.data.data", response.data.data)
					if (response.data.data.heartList != null) {
						this.heartTableData = response.data.data.heartList
						this.heartPageCount = response.data.data.heartPageCount
					} else {
						this.$message.error("睡眠空数据")
					}
				})
			},
			flushPhone() {
				http.post("phone/find", {
					phonePageNow: this.phonePageNow,
				}).then((response) => {
					console.log("response.data.data", response.data.data)
					if (response.data.data.phoneList != null) {
						this.phoneTableData = response.data.data.phoneList
						this.phonePageCount = response.data.data.phonePageCount
					} else {
						this.$message.error("手机空数据")
					}
				})
			},
			flushMusic() {
				http.post("music/find", {
					musicPageNow: this.musicPageNow,
				}).then((response) => {
					console.log("response.data.data", response.data.data)
					if (response.data.data.musicList != null) {
						this.musicTableData = response.data.data.musicList
						this.musicPageCount = response.data.data.musicPageCount
					} else {

					}
				})
			},
			flushWater() {
				http.post("water/find", {
					waterPageNow: this.waterPageNow,
				}).then((response) => {
					console.log("response.data.data", response.data.data)
					if (response.data.data.waterList != null) {
						this.waterTableData = response.data.data.waterList
						this.waterPageCount = response.data.data.waterPageCount
					} else {

					}
				})
			},
			flushAcid() {
				http.post("acid/find", {
					acidPageNow: this.acidPageNow,
				}).then((response) => {
					console.log("response.data.data", response.data.data)
					if (response.data.data.acidList != null) {
						this.acidTableData = response.data.data.acidList
						this.acidPageCount = response.data.data.acidPageCount
					} else {

					}
				})
			},
			sleepCurrentChange(pagenow) {
				this.sleepPageNow = pagenow
				this.flushSleep()
			},
			sleepHandleDelete(index, row) {
				console.log(row.id)
				http.post("sleep/delete", {
					sleepId: row.id,
				}).then((response) => {
					console.log("response.data.data", response.data.data)
					if (response.data.err_code != 0) {
						this.$message.error("删除失败")
					} else {
						this.$message.success("删除成功");
						// 刷新表格
						this.flushSleep()
					}
				})
			},
			weightCurrentChange(pagenow) {
				this.weightPageNow = pagenow
				this.flushWeight()
			},
			weightHandleDelete(index, row) {
				console.log(row.id)
				http.post("weight/delete", {
					weightId: row.id,
				}).then((response) => {
					console.log("response.data.data", response.data.data)
					if (response.data.err_code != 0) {
						this.$message.error("删除失败")
					} else {
						this.$message.success("删除成功");
						// 刷新表格
						this.flushWeight()
					}
				})
			},
			shitCurrentChange(pagenow) {
				this.shitPageNow = pagenow
				this.flushShit()
			},
			shitHandleDelete(index, row) {
				console.log(row.id)
				http.post("shit/delete", {
					shitId: row.id,
				}).then((response) => {
					console.log("response.data.data", response.data.data)
					if (response.data.err_code != 0) {
						this.$message.error("删除失败")
					} else {
						this.$message.success("删除成功");
						// 刷新表格
						this.flushShit()
					}
				})
			},
			heartCurrentChange(pagenow) {
				this.heartPageNow = pagenow
				this.flushHeart()
			},
			heartHandleDelete(index, row) {
				console.log(row.id)
				http.post("heart/delete", {
					heartId: row.id,
				}).then((response) => {
					console.log("response.data.data", response.data.data)
					if (response.data.err_code != 0) {
						this.$message.error("删除失败")
					} else {
						this.$message.success("删除成功");
						// 刷新表格
						this.flushHeart()
					}
				})
			},
			phoneCurrentChange(pagenow) {
				this.phonePageNow = pagenow
				this.flushPhone()
			},
			phoneHandleDelete(index, row) {
				console.log(row.id)
				http.post("phone/delete", {
					phoneId: row.id,
				}).then((response) => {
					console.log("response.data.data", response.data.data)
					if (response.data.err_code != 0) {
						this.$message.error("删除失败")
					} else {
						this.$message.success("删除成功");
						// 刷新表格
						this.flushPhone()
					}
				})
			},
			waterCurrentChange(pagenow) {
				this.waterPageNow = pagenow
				this.flushWater()
			},
			waterHandleDelete(index, row) {
				console.log(row.id)
				http.post("water/delete", {
					waterId: row.id,
				}).then((response) => {
					console.log("response.data.data", response.data.data)
					if (response.data.err_code != 0) {
						this.$message.error("删除失败")
					} else {
						this.$message.success("删除成功");
						// 刷新表格
						this.flushWater()
					}
				})
			},
			acidCurrentChange(pagenow) {
				this.acidPageNow = pagenow
				this.flushAcid()
			},
			acidHandleDelete(index, row) {
				console.log(row.id)
				http.post("acid/delete", {
					acidId: row.id,
				}).then((response) => {
					console.log("response.data.data", response.data.data)
					if (response.data.err_code != 0) {
						this.$message.error("删除失败")
					} else {
						this.$message.success("删除成功");
						// 刷新表格
						this.flushAcid()
					}
				})
			},

			musicCurrentChange(pagenow) {
				this.musicPageNow = pagenow
				this.flushMusic()
			},
			musicHandleDelete(index, row) {
				console.log(row.id)
				http.post("music/delete", {
					musicId: row.id,
				}).then((response) => {
					console.log("response.data.data", response.data.data)
					if (response.data.err_code != 0) {
						this.$message.error("删除失败")
					} else {
						this.$message.success("删除成功");
						// 刷新表格
						this.flushMusic()
					}
				})
			},

			onSubmit() {

			},


		},
		mounted() {
			console.log("this.sleepTableData", this.sleepTableData)
			http.post("sleep/find", {
				sleepPageNow: this.sleepPageNow,
			}).then((response) => {
				console.log("response.data.data", response.data.data)
				if (response.data.data.sleepList != null) {
					this.sleepTableData = response.data.data.sleepList
					this.sleepPageCount = response.data.data.sleepPageCount
				} else {
					this.$message.error("拉屎空数据")
				}
			})

			console.log("this.weightTableData", this.weightTableData)
			http.post("weight/find", {
				weightPageNow: this.weightPageNow,
			}).then((response) => {
				console.log("response.data.data", response.data.data)
				if (response.data.data.weightList != null) {
					this.weightTableData = response.data.data.weightList
					this.weightPageCount = response.data.data.weightPageCount
				} else {
					this.$message.error("体重空数据")
				}
			})

			console.log("this.shitTableData", this.shitTableData)
			http.post("shit/find", {
				shitPageNow: this.shitPageNow,
			}).then((response) => {
				console.log("response.data.data", response.data.data)
				if (response.data.data.shitList != null) {
					this.shitTableData = response.data.data.shitList
					this.shitPageCount = response.data.data.shitPageCount
				} else {
					this.$message.error("拉屎空数据")
				}
			})

			console.log("this.heartTableData", this.heartTableData)
			http.post("heart/find", {
				heartPageNow: this.heartPageNow,
			}).then((response) => {
				console.log("response.data.data", response.data.data)
				if (response.data.data.heartList != null) {
					this.heartTableData = response.data.data.heartList
					this.heartPageCount = response.data.data.heartPageCount
				} else {
					this.$message.error("heart空数据")
				}
			})

			console.log("this.phoneTableData", this.phoneTableData)
			http.post("phone/find", {
				phonePageNow: this.phonePageNow,
			}).then((response) => {
				console.log("response.data.data", response.data.data)
				if (response.data.data.phoneList != null) {
					this.phoneTableData = response.data.data.phoneList
					this.phonePageCount = response.data.data.phonePageCount
				} else {
					this.$message.error("手机空数据")
				}
			})

			http.post("music/find", {
				musicPageNow: this.musicPageNow,
			}).then((response) => {
				console.log("response.data.data musc>>>>>", response.data.data)
				if (response.data.data.musicList != null) {
					this.musicTableData = response.data.data.musicList
					this.musicPageCount = response.data.data.musicPageCount
				} else {}
			})


			http.post("water/find", {
				waterPageNow: this.waterPageNow,
			}).then((response) => {
				console.log("response.data.data musc>>>>>", response.data.data)
				if (response.data.data.waterList != null) {
					this.waterTableData = response.data.data.waterList
					this.waterPageCount = response.data.data.waterPageCount
				} else {}
			})

			http.post("acid/find", {
				acidPageNow: this.acidPageNow,
			}).then((response) => {
				console.log("response.data.data musc>>>>>", response.data.data)
				if (response.data.data.acidList != null) {
					this.acidTableData = response.data.data.acidList
					this.acidPageCount = response.data.data.acidPageCount
				} else {}
			})
		}
	}
</script>

<style lang="less" scoped>
	.tabs {
		margin: 5px 30px;

	}

	.dobutton {
		margin: 10px 30px;

		.topbutton {
			margin: 3px 3px;
		}
	}
</style>
