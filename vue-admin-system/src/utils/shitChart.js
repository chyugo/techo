import moment from "moment";
	export function drawShit(userData12) {
		const BASE_TIME = '2020-01-20'
		const legend = []
		// const getupTimeArray = []
		// const sleepTimeArray = []
		const shitTimeArray = []
		const travel = []
		for (let item of userData12) {
			legend.push(moment(item.recordDate).format('MM-DD'));
			// getupTimeArray.push(item.getuptime ?
			// 	`${BASE_TIME} ${moment(item.getuptime).format('YYYY-MM-DD HH:mm').substr(11, 10)}` :
			// 	null);
			// sleepTimeArray.push(item.sleeptime ?
			// 	`${BASE_TIME} ${moment(item.sleeptime).format('YYYY-MM-DD HH:mm').substr(11, 10)}` :
			// 	null);
			shitTimeArray.push(item.ShitTime ?
				`${BASE_TIME} ${moment(item.ShitTime).format('YYYY-MM-DD HH:mm').substr(11, 10)}` : null
			);
			travel.push(item.ccNum !== '0' ?
				`${BASE_TIME} ${moment('2020-02-02 12:00').format('YYYY-MM-DD HH:mm').substr(11, 10)}` :
				null);
		}
		console.log("this.legend", legend)
		// console.log("this.onTime", getupTimeArray)
		// console.log("this.offTime", sleepTimeArray)
		console.log("shitTimeArray",shitTimeArray)
		console.log("this.travel", travel)
		const option12 = {
			grid: {
				left: '5%',
				right: '2%',
				bottom: '0%',
				containLabel: true
			},
			tooltip: {
				show: true,
				axisPointer: {
					type: 'shadow'
				},
				position: function(point, params, dom, rect, size) {
					// 固定在顶部
					return [point[0], '10%'];
				},
				trigger: 'axis',
				formatter: function(value) {
					return "拉屎" + (moment(value[0].data).format('YYYY-MM-DD HH:mm')).substr(11, 20)
					// console.log("value[0].data",(moment(value[0].data).format('YYYY-MM-DD HH:mm')).substr(11, 20))
					// return "起床" + (moment(value[1].data).format('YYYY-MM-DD HH:mm')).substr(11, 20) +
					// 	"<br \>入睡" + (moment(value[0].data).format('YYYY-MM-DD HH:mm')).substr(11, 20) +
					// 	"<br \>拉屎" + (moment(value[2].data).format('YYYY-MM-DD HH:mm')).substr(11, 20)
				}
			},
			toolbox: {
				feature: {
					restore: {
						title: '还原'
					},
					saveAsImage: {
						title: '保存图片'
					}
				},
				show: true
			},
			title: [{
				text: '拉屎',
				 textStyle :{
					 color: '#000',
				 }
			}],
			xAxis: {
				type: 'category',
				data: userData12.map(item => item.RecordDate),
				axisLabel: {
					interval: 0,
					rotate: 45,
					color: 'black',
				},
				splitLine: {
					show: true,
				},
			},
			yAxis: {
				type: 'time',
				maxInterval: 3600 * 1000,
				min: `${BASE_TIME} 0:00:00`,
				max: `${BASE_TIME} 23:59:59`,
				axisLabel: {
					color: 'black',
					formatter: function(value) {
						return (moment(value).format('YYYY-MM-DD HH:mm')).substr(11, 20)
						// return (moment(value).format('YYYY-MM-DD HH:mm')).substr(0, 20)
					}
				},
				splitLine: {
					show: true,
				},
			},
			legend: {
				// data: ['睡眠', '起床', "拉屎"]
				data: ["拉屎"]
			},
			series: [
				// {
				// 	type: 'line',
				// 	name: '睡眠',
				// 	data: sleepTimeArray,
				// 	smooth: true,
				// 	itemStyle: {
				// 		normal: {
				// 			label: {
				// 				show: true, //开启显示
				// 				backgroundColor: '#ffffff',
				// 				position: 'top', //在上方显示
				// 				textStyle: { //数值样式
				// 					color: 'black',
				// 					fontSize: 12
				// 				},
				// 				formatter: function(value) {
				// 					// console.log("value", value)
				// 					return value.seriesName.substr(0, 1) + value.data.substr(11, 20)
				// 				}
				// 			},

				// 		}
				// 	},
				// 	markLine: {
				// 		symbol: 'none',
				// 		data: [{
				// 			name: '睡眠',
				// 			yAxis: `${BASE_TIME} 22:30`
				// 		}, ],
				// 		label: {
				// 			formatter: '{b}',
				// 		}
				// 	},

				// },
				// {
				// 	type: 'line',
				// 	name: '起床',
				// 	smooth: true,
				// 	itemStyle: {
				// 		normal: {
				// 			label: {
				// 				show: true, //开启显示
				// 				backgroundColor: '#ffffff',
				// 				position: 'top', //在上方显示
				// 				textStyle: { //数值样式
				// 					color: 'black',
				// 					fontSize: 12
				// 				},
				// 				formatter: function(value) {
				// 					// console.log("value", value)
				// 					return value.seriesName.substr(0, 1) + value.data.substr(11, 20)
				// 				}
				// 			},

				// 		}
				// 	},
				// 	data: getupTimeArray,

				// },
				{
					type: 'line',
					symbolSize: 8,
					name: '拉屎',
					
					smooth: true,
					data: shitTimeArray,
					itemStyle: {
						normal: {
							label: {
								show: true, //开启显示
								backgroundColor: '#ffffff',
								position: 'top', //在上方显示
								textStyle: { //数值样式
									color: 'black',
									fontSize: 12
								},
								formatter: function(value) {
									// console.log("value", value)
									return value.seriesName.substr(0, 1) + value.data.substr(11, 20)
								}
							},

						}
					},
				},
			]
		}
		// console.log("sleepTimeArray")
		return option12
	}
