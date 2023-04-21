// 图表4
import {
	fomartTime
} from './time.js'
export function drawsleep(userData4) {

	const option4 = {
		animation: false,
		tooltip: {
			// show: true,
			// formatter: function(params) {
			// 	return fomartTime(params.data)
			// },
			position: function(point, params, dom, rect, size) {
				// 固定在顶部
				return [point[0]-70, '0%'];
			},
			
			trigger: 'axis',
			axisPointer: {
				type: 'shadow',

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

		grid: {
			left: '2%',
			right: '2%',
			bottom: '0%',
			containLabel: true,
		},
		title: [{
			textStyle: {
				color: '#000',
			},
			left: 'left',
			text: '睡眠时长'
		}, ],
		xAxis: {
			type: 'category',
			data: userData4.map(item => item.RecordDate),
			axisLabel: {
				interval: 0,
				rotate: 45,
				color: 'black',
			},
			splitLine: {
				show: true,
			},
			// boundaryGap: false,
		},
		yAxis: {
			type: 'value',
			// interval :2,
			axisLabel: {
				formatter: function(value, index) {
					return fomartTime(value)
				},
				fontSize: 11,
				color: 'black',
			},
			interval: 30,
			nameLocation: 'end',
			min: 240,
			max: 600,
		},
		// color: ["#f3dd18"],
		visualMap: {
			top: 50,
			right: 0,
			// showLabel: false,
			show: false,
			pieces: [{
					gt: 480,
					lte: 600,
					color: '#93CE07',
					label: '绿-容忍区间',
				},
				{
					gt: 420,
					lte: 480,
					color: '#FBDB0F',
					label: '黄-告警',
				},
				{
					gt: 240,
					lte: 420,
					color: '#FD0100',
					label: '红-急需改善'
				},
			],
			outOfRange: {
				color: '#93CE07'
			}
		},
		series: [{
			data: userData4.map(item => item.Active ? item.Active : 'b'),
			type: 'line',
			symbolSize: 8,
			smooth: true,
			itemStyle: {
				normal: {
					label: {
						show: true, //开启显示
						position: 'top', //在上方显示
						textStyle: { //数值样式
							color: 'black',
							fontSize: 11,
							// backgroundColor: 'rgba(141, 141, 141, 0.2)',
							backgroundColor: '#ffffff',
						},
						// offset: [10, 0],
						formatter: function(params) {
							return fomartTime(params.data)
						}
					}
				}
			},
			markLine: {
				silent: true,

				lineStyle: {
					color: '#6c6d6e'
				},
				data: [{
						yAxis: 540,
						name: '9小时',
					}, {
						yAxis: 480,
						name: '8小时',
					}, {
						yAxis: 420,
						name: '7小时',
					}, {
						yAxis: 360,
						name: '6小时',
					},

				],
				label: {
					formatter: '{b}',
					position: 'middle',
					color: '#dadbdd',
					fontWeight: '600'
				}
			}
		}]
	};
	return option4
}
