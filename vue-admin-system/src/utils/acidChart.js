import {
	fomartTime
} from './time.js'

export function drawAcid(userData7) {
	const option7 = {
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
		grid: {
			left: '5%',
			right: '2%',
			bottom: '0%',
			containLabel: true,
		},
		title: [{
			text: '反酸',
			textStyle: {
				color: '#000',
			}
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
			top: 50,
			right: 0,
			// showLabel: false,

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
			data: userData7.map(item => item.Acid == 0 ? null : item.Acid),
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
							fontSize: 15
						},

					}
				}
			},
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
				}, {
					yAxis: 60,
					name: '急需改善',
				}],
				label: {
					formatter: '{b}',
					position: 'insideMiddleBottom',
					color: '#e8966b',
					fontWeight: '600',
					
				},
			}
		}, ]
	};
	return option7
}
