import {
	fomartTime
} from './time.js'

export function drawWater(userData7) {
	const option7 = {
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
		grid: {
			left: '5%',
			right: '2%',
			bottom: '0%',
			containLabel: true,
		},
		title: [{
			textStyle: {
				color: '#000',
			},
			text: '喝水'
		}],

		xAxis: {
			type: 'category',
			data: userData7.map(item => item.RecordDate),
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
			axisLabel: {
				color: 'black',
			},
			interval: 300,
			nameLocation: 'end',
			min: 0,
			max: 3300
		},
		visualMap: {
			top: 50,
			right: 0,
			// showLabel: false,
			show: false,
			pieces: [{
					gt: 1200,
					lte: 2000,
					color: '#FBDB0F',
					label: '需要改善',

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
			data: userData7.map(item => item.Water == 0 ? null : item.Water),
			type: 'line',
			symbolSize: 8,
			smooth: true,
			itemStyle: {
				normal: {
					label: {
						show: true, //开启显示
						backgroundColor: '#ffffff',
						fontWeight: '700',
						position: 'top', //在上方显示
						textStyle: { //数值样式
							color: 'black',
							fontSize: 12
						},

					}
				}
			},
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
		}, ]
	};
	return option7
}
