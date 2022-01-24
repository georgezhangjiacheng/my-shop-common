package timeTranslate

import (
	"fmt"
	"strconv"
	"time"
)

/*
 *@Author georgezhangjc@163.com
 *@Date 2021/12/31 下午2:54
 *@Describe 时间转换工具
 */

func TimeStampToDate(timeStamp int64) (dateStamp string) {
	// 将时间戳转为当天00:00:00的时间戳 2020-10-12 13:12:45 -> 2020-10-12 00:00:00
	tm := time.Unix(timeStamp, 0).Format("2006-01-02")
	// 日期转时间戳
	loc, _ := time.LoadLocation("Local") //获取时区
	formatTime, _ := time.ParseInLocation("2006-01-02", tm, loc)
	// 转为字符的时间戳
	dateStamp = strconv.FormatInt(formatTime.Unix(), 10)
	return dateStamp
}

func TimeStampToStrDate(timeStamp int64) (datetime string) {

	//时间戳转化为日期（UTC）
	datetime = time.Unix(timeStamp, 0).Format("2006-01-02 15:04:05")
	return datetime
}

func TimeStampToStrDateShangHai(timeStamp int64) (datetime string) {

	//时间戳转化为北京日期字符串
	//timeLocal, _ := time.LoadLocation("Asia/Chongqing") //获取时区

	timeLocal := time.FixedZone("CST", 3600*8) // 使用东八区时间
	tm := time.Unix(timeStamp, 0).Format("2006-01-02 15:04:05")

	formatTime, _ := time.ParseInLocation("2006-01-02 15:04:05", tm, timeLocal)
	datetime = formatTime.Format("2006-01-02 15:04:05")
	return datetime
}

func IntDateToTimestamp(date int64) (dateStamp int64) {

	StrDate := strconv.FormatInt(date, 10)
	Date := StrDate[0:4] + "-" + StrDate[4:6] + "-" + StrDate[6:8]
	loc, _ := time.LoadLocation("Local") //获取时区
	formatTime, _ := time.ParseInLocation("2006-01-02", Date, loc)

	// 转为字符的时间戳
	dateStamp = formatTime.Unix()
	return dateStamp
}

func LastMonth(n int) (dateList []time.Time) {
	// 获取最近几个月的时间信息
	nowTime := time.Now()
	dateList = append(dateList, nowTime)
	for i := 0; i < n-1; i++ {
		// 找到上个月的这天
		lastMonthTime := nowTime.AddDate(0, -1, -nowTime.Day()+1)
		dateList = append(dateList, lastMonthTime)
		nowTime = lastMonthTime
	}
	return
}

func LastQuarter(n int) (rsp []string) {
	now := time.Now()
	month := now.Month()
	quarter := getQuarter(int(month))
	rsp = append(rsp, fmt.Sprintf("%d.Q%d", now.Year(), quarter))
	for i := 0; i < n-1; i++ {
		now = now.AddDate(0, -3, -now.Day()+1)
		s := getQuarter(int(now.Month()))
		rsp = append(rsp, fmt.Sprintf("%d.Q%d", now.Year(), s))
	}
	return
}

func getQuarter(month int) int {
	switch month {
	case 1, 2, 3:
		return 1
	case 4, 5, 6:
		return 2
	case 7, 8, 9:
		return 3
	case 10, 11, 12:
		return 4
	}
	return -1
}
