import {
	fomartTime
} from './time.js'

export function drawWeight(userData7) {
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
			text: '体重变化',
			textStyle: {
				color: '#000',
			}
		}],
		legend: {
			data: ['上午', '中午', '晚上']
		},
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
			// interval :2,
			axisLabel: {
				fontSize: 12,
				color: 'black',
				hideOverlap:'true',
			},
			axisLine:{
				show:true,
			},
			interval: 0.3,
			nameLocation: 'end',
			min: 'dataMin',
			max: 'dataMax'
		},
		series: [{
				name: '上午',
				data: userData7.map(item => item.MorningWeight == 0? null :item.MorningWeight),
				type: 'line',
				symbol:"circle",
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
								fontSize: 10
							},

						}
					}
				},
			},
			{
				name: '中午',
				data:userData7.map(item => item.NoonWeight == 0? null :item.NoonWeight),
				type: 'line',
				symbol:"roundRect",
				symbolSize: 10,
				smooth: true,
				itemStyle: {
					normal: {
						label: {
							show: true, //开启显示
							backgroundColor: '#ffffff',
							offset: [10, 0],
							position: 'top', //在上方显示
							textStyle: { //数值样式
								color: 'black',
								fontSize: 10
							},

						}
					}
				},
			}, {
				name: '晚上',
				data: userData7.map(item => item.NightWeight == 0? null :item.NightWeight),
				type: 'line',
				symbol:"emptyCircle",
				symbolSize: 10,
				smooth: true,
				itemStyle: {
					normal: {
						label: {
							show: true, //开启显示
							backgroundColor: '#ffffff',
							offset: [10, 0],
							position: 'top', //在上方显示
							textStyle: { //数值样式
								color: 'black',
								fontSize: 10
							},

						}
					}
				},
			},

		]
	};
	return option7
}
