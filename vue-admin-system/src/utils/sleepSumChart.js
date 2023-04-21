import {
	fomartTime
} from './time.js'


export function drawSleepSum(userData5) {
	const option5 = {
		animation: false,
		tooltip: {
			show: true,
			trigger: 'axis',
			position: function(point, params, dom, rect, size) {
				// 固定在顶部
				return [point[0]-70, '0%'];
			},
			formatter: function(params) {
				return fomartTime(params[0].data)
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
		title: [{
			left: 'left',
			text: '【统计】睡眠时长',
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
		}, ],
		xAxis: {
			type: 'category',
			boundaryGap: false,
			data: userData5.map(item => item.RecordDate),
			axisLabel: {
				rotate: 45,
				color: 'black',
			},
			splitLine: {
				show: true,

			},
		},
		yAxis: {
			type: 'value',
			interval: 30,
			axisLabel: {
				color: 'black',
				formatter: function(value, index) {
					return fomartTime(value)
				}
			},
			min: 240,
			max: 600,
		},
		grid: {
			left: '2%',
			right: '2%',
			bottom: '0%',
			containLabel: true,
		},

		// dataZoom: [{
		// 		type: 'slider',
		// 		show: true,
		// 		start: 0,
		// 		end: 100,
		// 		moveOnMouseMove: true,
		// 		zoomOnMouseWheel: false,
		// 		preventDefaultMouseMove: false,
		// 	},
		// ],
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
			data: userData5.map(item => item.Active ? item.Active : 'b'),
			// type: 'line',
			// 散点图
			type: 'scatter',
			// 采样降低
			// sampling: 'lttb',
			symbolSize: 8,
			
			markLine: {
				silent: true,
				lineStyle: {
					color: '#6c6c6c'
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
					color: '#b8b8b8',
					fontWeight: '600'
				}
			}
		}]
	};
	return option5;
}
