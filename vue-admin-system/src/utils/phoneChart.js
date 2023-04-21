import {
	fomartTime
} from './time.js'

export function drawphone(userData3) {
	const option3 = {
		animation: false,
		tooltip: {
			show: true,
			trigger: 'axis',
			axisPointer: {
				type: 'shadow'
			},
			position: function(point, params, dom, rect, size) {
				// 固定在顶部
				return [point[0]-70, '0%'];
			},
			formatter: function(params) {
				return fomartTime(params[0].data)
			}
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
			textStyle: {
				color: '#000',
			},
			left: 'left',
			text: 'phone时长',
		},
		xAxis: {
			type: 'category',
			data: userData3.map(item => item.RecordDate),
			axisLabel: {
				interval: 0,
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
			max: 901
		},
		grid: {
			left: '2%',
			right: '2%',
			bottom: '0%',
			containLabel: true,
		},
		color: ["#a5552b"],
		series: [{
			data: userData3.map(item => item.PhoneMinute),
			type: 'line',
			symbolSize: 10,
			smooth: true,
			itemStyle: {
				normal: {
					label: {
						show: true, //开启显示
						position: 'top', //在上方显示
						backgroundColor: '#ffffff',
						textStyle: { //数值样式
							color: 'black',
							fontSize: 12
						},
						formatter: function(params) {
							return fomartTime(params.data)
						}
					}
				}
			},
			// showBackground: true,
			markLine: {
				silent: true,
				lineStyle: {
					color: '#fc041f',
					width: 2,
					type: [5, 1],
				},
				data: [{
					yAxis: 180,
					name: '3小时,贪玩线',
				}],
				label: {
					formatter: '{b}',
					position: 'middle',
					color: '#fc041f',
					fontWeight: '600'
				}
			}
		}]
	};
	return option3
}
