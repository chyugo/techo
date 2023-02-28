package utils

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

// 10位时间戳转换成2000-01-01
func Stamp10(t string) string {
	timeLayout := "2006-01-02 15:04:05"
	temp, err := strconv.ParseInt(t, 10, 64)
	temp2 := time.Unix(temp, 0).Format(timeLayout)
	if err != nil {
		fmt.Println("时间戳转换err", err)
		return ""
	}
	return temp2
}

// 10位时间戳转换成2000-01-01
func Stamp10_2(t string) string {
	timeLayout := "2006-01-02"
	temp, err := strconv.ParseInt(t, 10, 64)
	temp2 := time.Unix(temp, 0).Format(timeLayout)
	if err != nil {
		fmt.Println("时间戳转换err", err)
		return ""
	}
	return temp2
}

// 返回当日9点的时间戳 传入10位 传出10位
func NineStamp(i int64) int64 {
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>", i)
	datatime := time.Unix(i, 0)
	year, month, day := datatime.Date()
	fmt.Println("获取到时间,", year, month, day)
	return time.Date(year, month, day, 9, 0, 0, 0, time.Local).Unix()
}

// 传入10位时间戳，保留当天，指定时间，返回10位时间戳
func StampConv(i int64, hour int, min int) int64 {
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>", i)
	datatime := time.Unix(i, 0)
	year, month, day := datatime.Date()
	fmt.Println("获取到时间,", year, month, day)
	return time.Date(year, month, day, hour, min, 0, 0, time.Local).Unix()
}

func MinuteToHourMinute1(minute int) string {
	if minute == 0 {
		return "无数据"
	}
	h := minute / 60
	m := minute % 60
	fmt.Println("h", h, "m", m)
	return fmt.Sprintf("%d:%d", h, m)
}

func MinuteToHourMinute2(minute int) string {
	if minute == 0 {
		return "无数据"
	}
	h := minute / 60
	m := minute % 60
	return fmt.Sprintf("%d小时%d分钟", h, m)
}

func StamptoString(i int64) string {
	if i == 0 {
		return "无数据"
	}
	datatime := time.Unix(i, 0)
	year, month, day := datatime.Date()
	hour, minute, second := datatime.Clock()
	return time.Date(year, month, day, hour, minute, second, 0, time.Local).String()[0:19]
}

// 根据日期查找 返回当月和下月时间戳
func SearchMonthConv(month string) (int64, int64, error) {
	if !IsMatchNumber(month) {
		return -1, -1, errors.New("SearchMonthConv err.")
	}
	if month == "" || len(month) < 9 {
		return -1, -1, errors.New("SearchMonthConv err.")
	}

	month = month[:len(month)-3]
	tempMonth, err := strconv.ParseInt(month, 10, 64)
	if err != nil {
		fmt.Println("err", err)
		return -1, -1, errors.New("SearchMonthConv err.")
	}
	fmt.Println("tempMonth", tempMonth)
	thisMonth := time.Unix(tempMonth, 0).AddDate(0, 0, -5).Unix()
	nextMonth := time.Unix(tempMonth, 0).AddDate(0, 1, -1).Add(time.Hour * 23).Unix()
	fmt.Println("thisMonth", thisMonth, "nextMonth", nextMonth)
	return thisMonth, nextMonth, nil
}
