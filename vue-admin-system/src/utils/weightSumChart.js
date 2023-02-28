import {
	fomartTime
} from './time.js'

export function drawWeightSum(userData8) {
	const option8 = {
		tooltip: {
			show: true,
			trigger: 'axis',
			position: function(point, params, dom, rect, size) {
				// 固定在顶部
				return [point[0], '10%'];
			},
		},
		toolbox: {
			feature: {

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

			text: '【统计】体重',
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
		}],
		legend: {
			data: ['上午', '中午', '晚上']
		},
		xAxis: {
			type: 'category',
			data: userData8.map(item => item.RecordDate),
			axisLabel: {
				rotate: 45,
				color: 'black',
			},
			splitLine: {
				show: true,
			},
			// boundaryGap: false
		},
		yAxis: {
			type: 'value',
			// interval :2,
			axisLabel: {
				fontSize: 12,
				color: 'black',
			},
			nameLocation: 'end',
			min: 'dataMin',
			max: 'dataMax'
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
		series: [


			{
				name: '上午',
				data: userData8.map(item => item.MorningWeight == 0 ? null : item.MorningWeight),
				type: 'scatter',
				symbol: "triangle",
				smooth: true,
				// 采样降低
				// sampling: 'lttb',

			},
			{
				name: '中午',
				data: userData8.map(item => item.NoonWeight == 0 ? null : item.NoonWeight),
				type: 'scatter',
				symbol: "diamond",
				smooth: true,
				// 采样降低
				// sampling: 'lttb',

			}, {
				name: '晚上',
				data: userData8.map(item => item.NightWeight == 0 ? null : item.NightWeight),
				type: 'scatter',
				symbol: "rectangle",
				smooth: true,
				// 采样降低
				// sampling: 'lttb',

			},

		]
	};
	return option8
}
