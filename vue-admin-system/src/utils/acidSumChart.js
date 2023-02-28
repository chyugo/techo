import {
	fomartTime
} from './time.js'
export function drawAcidSum(userData10) {
	const option10 = {
		tooltip: {
			show: true,
			axisPointer: {
				type: 'shadow'
			},
			position: function(point, params, dom, rect, size) {
				// 固定在顶部
				return [point[0], '10%'];
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
			text: '【统计】反酸',
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
				fontSize: 12,
				color: 'black',
				formatter: '{value}次',
			},
			interval: 5,
			nameLocation: 'end',
			min: 0,
			max: 60
		},
		visualMap: {
			        bottom:0,
					left:0,
					show: false,
			pieces: [{
					gt: 0,
					lte: 20,
					color: '#93CE07',
					label: '绿-容忍区间',
				},
				{
					gt: 20,
					lte: 35,
					color: '#FBDB0F',
					label: '黄-告警',
				},
				{
					gt: 35,
					lte: 100,
					color: '#FD0100',
					label: '红-急需改善'
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
				data: userData10.map(item => item.Acid),
				markLine: {
					silent: true,
					lineStyle: {
						color: '#FD0100'
					},
					data: [{
						yAxis: 20,
						name: '容忍区间',
					}, {
						yAxis: 35,
						name: '告警',
					},{
						yAxis: 60,
						name: '急需改善',
					}],
					label: {
						formatter: '{b}',
						position: 'insideMiddleBottom',
						color: '#e8966b',
						fontWeight: '600'
					},
				}
			},

		]
	};
	return option10
}
