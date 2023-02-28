import {
	fomartTime
} from './time.js'

export function drawPhoneSum(userData6){
	const option6 = {
		tooltip: {
			show: true,
			trigger: 'axis',
			axisPointer: {
				type: 'shadow'
			},
			position: function(point, params, dom, rect, size) {
				// 固定在顶部
				return [point[0], '10%'];
			},
			formatter: function(params) {
				return fomartTime(params[0].data)
			},
		},
		toolbox: {
			feature: {
				magicType: {
					type: ["line", "bar"],
					title: {
						line: '切换为折线图',
						bar: '切换为柱状图',
					}
				},
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
			left: 'left',
			text: '【统计】phone时长',
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
		}, ],
		xAxis: {
			type: 'category',
			data: userData6.map(item => item.RecordDate),
			axisLabel: {
				rotate: 45,
				color: 'black',
			},
			axisTick: {
				alignWithLabel: true
			},
			splitLine: {
				show: true,
			},
	
		},
		yAxis: {
			type: 'value',
			// interval :2,
			axisLabel: {
				color: 'black',
				formatter: function(value, index) {
					return fomartTime(value)
				}
			},
			interval: 60,
			max: 901,
			splitArea: {
				show: false
			}
		},
		// dataZoom: [{
		// 		type: 'inside',
		// 		show: true,
		// 		start: 0,
		// 		end: 100,
		// 		moveOnMouseMove: true,
		// 		zoomOnMouseWheel: false,
		// 		preventDefaultMouseMove: false,
		// 	},
		// 	{
		// 		start: 0,
		// 		end: 100,
		// 	}
		// ],
		grid: {
			left: '2%',
			right: '2%',
			bottom: '0%',
			containLabel: true,
		},
		color: ["#a5552b"],
		series: [{
			data: userData6.map(item => item.PhoneMinute),
			type: 'bar',
			large: true,
			showBackground: true,
			// 采样降低
			// sampling: 'lttb',
		}]
	};
	return option6
}