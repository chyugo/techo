import moment from "moment";
export function drawWakeSum(userData) {
	const BASE_TIME = '2020-01-20'
	const legend2 = []
	const getupTimeArray2 = []
	const sleepTimeArray2 = []

	const travel2 = []
	// for (let item = 1; item < 100; item++) {
	// 	userData.push({
	// 		getuptime: 1672527600000,
	// 		sleeptime: 1672512180000,
	// 		ccNum: '0',
	// 		recordDate: "2023-01-01"
	// 	}, {
	// 		getuptime: 1672612980000,
	// 		sleeptime: 1672667980000,

	// 		ccNum: '0',
	// 		recordDate: "2023-01-02"
	// 	}, {
	// 		getuptime: 1672527600000,
	// 		sleeptime: 1672578000000,

	// 		ccNum: '0',
	// 		recordDate: "2023-01-01"
	// 	}, {
	// 		getuptime: 1672527600000,
	// 		sleeptime: 1672512180000,

	// 		ccNum: '0',
	// 		recordDate: "2023-01-01"
	// 	}, {
	// 		getuptime: 1672612980000,
	// 		sleeptime: 1672667980000,

	// 		ccNum: '0',
	// 		recordDate: "2023-01-02"
	// 	}, {
	// 		getuptime: 1672527600000,
	// 		sleeptime: 1672578000000,

	// 		ccNum: '0',
	// 		recordDate: "2023-01-01"
	// 	}, {
	// 		getuptime: 1672527600000,
	// 		sleeptime: 1672512180000,

	// 		ccNum: '0',
	// 		recordDate: "2023-01-01"
	// 	}, {
	// 		getuptime: 1672612980000,
	// 		sleeptime: 1672667980000,

	// 		ccNum: '0',
	// 		recordDate: "2023-01-02"
	// 	}, {
	// 		getuptime: 1672527600000,
	// 		sleeptime: 1672578000000,

	// 		ccNum: '0',
	// 		recordDate: "2023-01-01"
	// 	}, {
	// 		getuptime: 1672612980000,
	// 		sleeptime: 1672667980000,

	// 		ccNum: '0',
	// 		recordDate: "2023-01-02"
	// 	})
	// }
	for (let item of userData) {
		legend2.push(moment(item.RecordDate).format('MM-DD'));
		getupTimeArray2.push(item.GetTime ?
			`${BASE_TIME} ${moment(item.GetTime).format('YYYY-MM-DD HH:mm').substr(11, 10)}` : null);
		sleepTimeArray2.push(item.SleepTime ?
			`${BASE_TIME} ${moment(item.SleepTime).format('YYYY-MM-DD HH:mm').substr(11, 10)}` : null);
		travel2.push(item.ccNum !== '0' ?
			`${BASE_TIME} ${moment('2020-02-02 12:00').format('YYYY-MM-DD HH:mm').substr(11, 10)}` : null);
	}
	const option13 = {
		grid: {
			left: '5%',
			right: '2%',
			bottom: '0%',
			containLabel: true
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
				return "起床" + (moment(value[1].data).format('YYYY-MM-DD HH:mm')).substr(11, 20) +
					"<br \>入睡" + (moment(value[0].data).format('YYYY-MM-DD HH:mm')).substr(11, 20)


			}
		},
		title: [{
			text: '【统计】睡起',
			textStyle: {
				//文字颜色
				color: '#ac3ea2',
				//字体风格,'normal','italic','oblique'
				fontStyle: 'italic',
				//字体粗细 'normal','bold','bolder','lighter',100 | 200 | 300 | 400...
				fontWeight: 'normal',
				//字体系列
				fontFamily: 'sans-serif',
				//字体大小
				fontSize: 18
			}
		}],
		xAxis: {
			type: 'category',
			data: userData.map(item => item.RecordDate),
			axisLabel: {
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
			data: ['睡眠', '起床'],
		},
		// dataZoom: [{
		// 		type: 'inside',
		// 		show: true,
		// 		start: 0,
		// 		end: 100,
		// 		// moveOnMouseWheel: 'false',
		// 		// zoomOnMouseWheel: 'false',
		// 		moveOnMouseMove: true,
		// 		zoomOnMouseWheel: false,
		// 		preventDefaultMouseMove: false,
		// 	},
		// 	{
		// 		start: 0,
		// 		end: 100,
		// 	}
		// ],
		series: [{
				type: 'scatter',
				name: '睡眠',
				data: sleepTimeArray2,
				smooth: true,
				symbolSize: 5,
				// 采样降低
				// sampling: 'lttb',
			},
			{
				type: 'scatter',
				name: '起床',
				smooth: true,
				data: getupTimeArray2,
				symbolSize: 5,
				// 采样降低
				// sampling: 'lttb',
			},

		]
	}
	console.log("sleepTimeArray")
	return option13

}
