export function fomartTime(value) {
	let unit = ['分', '小时', '天'],
		day = 0,
		hour = 0,
		min = 0,
		second = 0,
		returnStr = 0,
		arrVal = value.toString().split(".");
	if (arrVal.length > 1) {
		second = parseFloat("0." + arrVal[1]);
		second *= 60;
		value = parseInt(arrVal[0]);
	}
	returnStr = value + unit[0];
	if (value >= 60) {
		hour = parseInt(value / 60);
		min = value % 60;
	}
	if (hour) {
		returnStr = hour + unit[1];
		if (min) {
			returnStr += min + unit[0]
		}
	}
	if (second) {
		returnStr += second.toFixed(0) + '秒'
	}
	return returnStr
}

