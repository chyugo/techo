import {
	fomartTime
} from './time.js'
export function drawWaterSum(userData10) {
	const option10 = {
		animation: false,
		tooltip: {
			show: true,
			axisPointer: {
				type: 'shadow'
			},
			position: function(point, params, dom, rect, size) {
				// 固定在顶部
				return [point[0]-70, '0%'];
			},
			trigger: 'axis'
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
			text: '【统计】喝水',
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

		grid: {
			left: '5%',
			right: '2%',
			bottom: '0%',
			containLabel: true,
		},
		xAxis: [{
			type: 'category',
			data: userData10.map(item => item.RecordDate),
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
		yAxis: {
			type: 'value',
			axisLabel: {
				color: 'black',
				fontSize: 12
			},
			interval: 300,
			nameLocation: 'end',
			min: 0,
			max: 3300
		},
		visualMap: {
			bottom: 0,
			left: 0,
			show: false,
			pieces: [{
					gt: 1200,
					lte: 2000,
					color: '#FBDB0F',
					label: '需要改善'
				},
				{
					gt: 2000,
					lte: 3000,
					color: '#93CE07',
					label: '达标'
				},
				{
					gt: 0,
					lte: 1200,
					color: '#FD0100',
					label: '急需改善'
				},

			],
			outOfRange: {
				color: '#93CE07'
			}
		},
		series: [{
				name: '当天',
				type: 'scatter',
				symbolSize: 8,
				color: '#82bafe',
				emphasis: {
					focus: 'series'
				},
				// 采样降低
				// sampling: 'lttb',
				data: userData10.map(item => item.Water),
				markLine: {
					silent: true,

					lineStyle: {
						color: '#6c6d6e'
					},
					data: [{
						yAxis: 2000,
						name: '达标',
					}, {
						yAxis: 1200,
						name: '至少',
					}, {
						yAxis: 0,
						name: '急需改善',
					}],
					label: {
						formatter: '{b}',
						position: 'middle',
						color: '#e8966b',
						fontWeight: '600'
					}
				}
			},

		]
	};
	return option10
}
