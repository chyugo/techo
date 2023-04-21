export function drawHeart(userData9) {
	const option9 = {
		animation: false,
		tooltip: {
			trigger: 'axis',
			axisPointer: {
				type: 'shadow'
			},
			position: function(point, params, dom, rect, size) {
				// 固定在顶部
				return [point[0]-70, '0%'];
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
			text: '心脏',
			textStyle: {
				color: '#000',
			}
		},
		// legend: {

		// },
		grid: {
			left: '5%',
			right: '2%',
			bottom: '0%',
			containLabel: true,
		},
		xAxis: [{
			type: 'category',
			data: userData9.map(item => item.date),
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
		}],
		yAxis: [{
			type: 'value',
			interval: 1,
			axisLabel: {
				formatter: '异常{value}%',
				color: 'black',
			},
	
		}],
		series: [{
				name: '异常',
				type: 'line',
				stack: 'a',
				color: '#d32532',
				symbolSize: 10,
				smooth: true,
				emphasis: {
					focus: 'series'
				},
				label: {
					show: true,
					position: 'top',
					fontWeight: '700',
					fontSize: 15,
					formatter: '{c}%',
					// backgroundColor: '#8d4103',
					backgroundColor: '#fff',
					color: '#000'
				},
				data: userData9.map(item => item.abnormal)
			},
			// {
			// 	name: '正常',
			// 	type: 'line',
			// 	symbol:"emptyCircle",
			// 	symbolSize: 8,
			// 	stack: 'a',
			// 	color: '#b8e5c4',
			// 	emphasis: {
			// 		focus: 'series'
			// 	},

			// 	data: userData9.map(item => item.normal)
			// },



		]
	};
	return option9

}
