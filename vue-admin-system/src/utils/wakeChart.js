import moment from "moment";
	export function drawWake(userData12) {
		const BASE_TIME = '2020-01-20'
		const legend = []
		const getupTimeArray = []
		const sleepTimeArray = []
		const travel = []
		for (let item of userData12) {
			legend.push(moment(item.recordDate).format('MM-DD'));
			getupTimeArray.push(item.GetTime ?
				`${BASE_TIME} ${moment(item.GetTime).format('YYYY-MM-DD HH:mm').substr(11, 10)}` :
				null);
			sleepTimeArray.push(item.SleepTime ?
				`${BASE_TIME} ${moment(item.SleepTime).format('YYYY-MM-DD HH:mm').substr(11, 10)}` :
				null);
			travel.push(item.ccNum !== '0' ?
				`${BASE_TIME} ${moment('2020-02-02 12:00').format('YYYY-MM-DD HH:mm').substr(11, 10)}` :
				null);
		}
		console.log("userData12", userData12)
		console.log("getupTimeArray", getupTimeArray)
		console.log("sleepTimeArray", sleepTimeArray)
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
					
					console.log("value[0].data",(moment(value[0].data).format('YYYY-MM-DD HH:mm')).substr(11, 20))
					return "起床" + (moment(value[1].data).format('YYYY-MM-DD HH:mm')).substr(11, 20) +
						"<br \>入睡" + (moment(value[0].data).format('YYYY-MM-DD HH:mm')).substr(11, 20)
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
				textStyle: {
					color: '#000',
				},
				text: '睡起'
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
					}
				},
				splitLine: {
					show: true,
				},
			},
			legend: {
				data: ['睡眠', '起床']
			},
			series: [
				{
					type: 'line',
					symbol:"emptyCircle",
					symbolSize: 8,
					name: '睡眠',
					data: sleepTimeArray,
					smooth: true,
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
					markLine: {
						symbol: 'none',
						data: [{
							name: '睡眠',
							yAxis: `${BASE_TIME} 22:30`
						}, ],
						label: {
							formatter: '{b}',
						}
					},

				},
				{
					type: 'line',
					name: '起床',
					symbol:"roundRect",
					symbolSize: 10,
					smooth: true,
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
					data: getupTimeArray,

				},
				
			]
		}
		// console.log("sleepTimeArray")
		return option12
	}
