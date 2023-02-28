import {
	fomartTime
} from './time.js'
export function drawHeartSum(userData10) {
	const option10 = {
		tooltip: {
			trigger: 'axis',
			axisPointer: {
				type: 'shadow'
			},
			position: function(point, params, dom, rect, size) {
				// 固定在顶部
				return [point[0], '10%'];
			},
			formatter: '{a0}:{c}% <br />'
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
		title: {
			left: 'left',
			text: '【统计】心脏',
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
		},
		// legend: {

		// },
		grid: {
			left: '2%',
			right: '2%',
			bottom: '0%',
			containLabel: true,
		},
		xAxis: [{
			type: 'category',
			data: userData10.map(item => item.date),
			axisLabel: {
				color: 'black',
				rotate: 45,
			},
			axisTick: {
				alignWithLabel: true
			},
			splitLine: {
				show: true,
			},
		}],
		yAxis: [{
			type: 'value',
			// interval: 10,
			axisLabel: {
				formatter: '异常{value}%',
				color: 'black',
			},
		}],
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
		series: [{
				name: '异常',
				type: 'bar',
				stack: 'a',
				color: '#7c8aff',
				emphasis: {
					focus: 'series'
				},
				// 采样降低
				// sampling: 'lttb',
				data: userData10.map(item => item.abnormal)
			},




		]
	};
	return option10
}
