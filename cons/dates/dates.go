package dates

import (
	"strconv"
	"time"
)

//GetWeekByEn 星期英转中
func GetWeekByEn(en string) (week string){
	weekList := make(map[string]string)

	weekList["Monday"]    = "星期一"
	weekList["Tuesday"]   = "星期二"
	weekList["Wednesday"] = "星期三"
	weekList["Thursday"]  = "星期四"
	weekList["Friday"]    = "星期五"
	weekList["Saturday"]  = "星期六"
	weekList["Sunday"]    = "星期日"

	wn, ok := weekList[en]
	if (ok){
		return wn
	}

	return
}

//GetWeekStartDate 获取当前周开始时间
func GetWeekStartDate()(date string){
	now := time.Now()
	offset := int(time.Monday - now.Weekday())
	if offset > 0 { //周日为0，周一 - 周日 大于0
		offset = -6
	}

	weekStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	date = weekStart.Format("2006-01-02")

	return
}

//GetDateChange +N天的日期
func GetDateChange(date string, n int64)(newDate string){
	timeLayout := "2006-01-02"  //转化所需模板
	loc, _ := time.LoadLocation("Local")    //获取时区
	tmp, _ := time.ParseInLocation(timeLayout, date, loc)
	timestamp := tmp.Unix()    //转化为时间戳 类型是int64

	newTimestamp := timestamp + n*3600*24
	newDate = time.Unix(newTimestamp, 0).Format(timeLayout)    //时间戳转化为日期

	return
}

//GetDisparityDay 计算相差N天
func GetDisparityDay(start,end string)(dd int64){
	timeLayout := "2006-01-02"  //转化所需模板
	loc, _ := time.LoadLocation("Local")    //获取时区
	startTmp, _ := time.ParseInLocation(timeLayout, start, loc)
	endTmp,_ := time.ParseInLocation(timeLayout, end, loc)
	dd = (endTmp.Unix() - startTmp.Unix())/(24*3600)
	return
}

//GetDisparityMonth 计算相差N月
func GetDisparityMonth(start,end string)(dm int64){
	timeLayout := "2006-01-02"  //转化所需模板
	loc, _ := time.LoadLocation("Local")    //获取时区
	startTmp, _ := time.ParseInLocation(timeLayout, start, loc)
	endTmp,_ := time.ParseInLocation(timeLayout, end, loc)
	dm = (endTmp.Unix() - startTmp.Unix())/(24*3600*30)
	return
}

//GetDisparityYear 计算相差N年
func GetDisparityYear(start,end string)(dm int64){
	timeLayout := "2006-01-02"  //转化所需模板
	loc, _ := time.LoadLocation("Local")    //获取时区
	startTmp, _ := time.ParseInLocation(timeLayout, start, loc)
	endTmp,_ := time.ParseInLocation(timeLayout, end, loc)
	dm = (endTmp.Unix() - startTmp.Unix())/(24*3600*365)
	return
}

//GetYoyDay N年月日前后的日期
func GetYoyDay(date string, y,m,d int)(newDate string){
	timeLayout := "2006-01-02"  //转化所需模板
	loc, _ := time.LoadLocation("Local")    //获取时区
	tmp, _ := time.ParseInLocation(timeLayout, date, loc)
	newDate = tmp.AddDate(y,m,d).Format("2006-01-02")
	return
}

//GetYoyYear N年前后的年份
func GetYoyYear(date string, y int)(newYear string){
	timeLayout := "2006"  //转化所需模板
	loc, _ := time.LoadLocation("Local")    //获取时区
	tmp, _ := time.ParseInLocation(timeLayout, date, loc)
	newYear = tmp.AddDate(y,0,0).Format("2006")
	return
}

//GetYmdToUnix 字符串转时间戳
func GetYmdToUnix(date string) (tInt int64, tStr string) {
	//toBeCharge := "2015-01-01"                                //待转化为时间戳的字符串 注意 这里的小时和分钟还要秒必须写 因为是跟着模板走的 修改模板的话也可以不写
	timeLayout := "2006-01-02"                             		//转化所需模板
	loc, _ := time.LoadLocation("Local")                  //重要：获取时区
	theTime, _ := time.ParseInLocation(timeLayout, date, loc) 	//使用模板在对应时区转化为time.time类型
	tInt = theTime.Unix()                                       //转化为时间戳 类型是int64
	tStr = strconv.FormatInt(tInt,10)
	return
}

//GetYmdHisToUnix 字符串转时间戳
func GetYmdHisToUnix(date string) (tInt int64, tStr string) {
	//toBeCharge := "2015-01-01 00:00:00"                       //待转化为时间戳的字符串 注意 这里的小时和分钟还要秒必须写 因为是跟着模板走的 修改模板的话也可以不写
	timeLayout := "2006-01-02 15:04:05"                         //转化所需模板
	loc, _ := time.LoadLocation("Local")                  //重要：获取时区
	theTime, _ := time.ParseInLocation(timeLayout, date, loc) 	//使用模板在对应时区转化为time.time类型
	tInt = theTime.Unix()                                       //转化为时间戳 类型是int64
	tStr = strconv.FormatInt(tInt,10)
	return
}

//GetCurDayToUnix 当天开始结束时间
func GetCurDayToUnix()(startInt, endInt int64, startStr, endStr string){
	startDate := time.Now().Format("2006-01-02") + " 00:00:00"
	endDate := time.Now().Format("2006-01-02") + " 23:59:59"
	timeLayout := "2006-01-02 15:04:05"                         //转化所需模板
	loc, _ := time.LoadLocation("Local")                  //重要：获取时区

	startTime, _ := time.ParseInLocation(timeLayout, startDate, loc) 	//使用模板在对应时区转化为time.time类型
	startInt = startTime.Unix()                                         //转化为时间戳 类型是int64
	startStr = strconv.FormatInt(startInt,10)

	endTime, _ := time.ParseInLocation(timeLayout, endDate, loc) 	//使用模板在对应时区转化为time.time类型
	endInt = endTime.Unix()                                         //转化为时间戳 类型是int64
	endStr = strconv.FormatInt(endInt,10)
	return
}

//GetCurWeekToUnix 当周开始结束时间
func GetCurWeekToUnix()(startInt, endInt int64, startStr, endStr string){
	now := time.Now()
	offset := int(time.Monday - now.Weekday())
	if offset > 0 { //周日为0，周一 - 周日 大于0
		offset = -6
	}

	weekStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	startDate := weekStart.Format("2006-01-02") + " 00:00:00"

	weekEnd := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset+6)
	endDate := weekEnd.Format("2006-01-02") + " 23:59:59"

	timeLayout := "2006-01-02 15:04:05"                         //转化所需模板
	loc, _ := time.LoadLocation("Local")                  //重要：获取时区

	startTime, _ := time.ParseInLocation(timeLayout, startDate, loc) 	//使用模板在对应时区转化为time.time类型
	startInt = startTime.Unix()                                         //转化为时间戳 类型是int64
	startStr = strconv.FormatInt(startInt,10)

	endTime, _ := time.ParseInLocation(timeLayout, endDate, loc) 	//使用模板在对应时区转化为time.time类型
	endInt = endTime.Unix()                                         //转化为时间戳 类型是int64
	endStr = strconv.FormatInt(endInt,10)
	return
}

//GetCurMonthToUnix 当月开始结束时间
func GetCurMonthToUnix()(startInt, endInt int64, startStr, endStr string){
	startDate := time.Now().AddDate(0,0,- time.Now().Day() + 1).Format("2006-01-02") + " 00:00:00"
	endDate := time.Now().AddDate(0,1,- time.Now().Day()).Format("2006-01-02") + " 23:59:59"

	timeLayout := "2006-01-02 15:04:05"                         //转化所需模板
	loc, _ := time.LoadLocation("Local")                  //重要：获取时区

	startTime, _ := time.ParseInLocation(timeLayout, startDate, loc) 	//使用模板在对应时区转化为time.time类型
	startInt = startTime.Unix()                                         //转化为时间戳 类型是int64
	startStr = strconv.FormatInt(startInt,10)

	endTime, _ := time.ParseInLocation(timeLayout, endDate, loc) 	//使用模板在对应时区转化为time.time类型
	endInt = endTime.Unix()                                         //转化为时间戳 类型是int64
	endStr = strconv.FormatInt(endInt,10)
	return
}

//GetChangeDayToUnix +N天开始结束时间
func GetChangeDayToUnix(d int)(startInt, endInt int64, startStr, endStr string){
	startDate := time.Now().AddDate(0,0,d).Format("2006-01-02") + " 00:00:00"
	endDate := time.Now().AddDate(0,0,d).Format("2006-01-02") + " 23:59:59"
	timeLayout := "2006-01-02 15:04:05"                         //转化所需模板
	loc, _ := time.LoadLocation("Local")                  //重要：获取时区

	startTime, _ := time.ParseInLocation(timeLayout, startDate, loc) 	//使用模板在对应时区转化为time.time类型
	startInt = startTime.Unix()                                         //转化为时间戳 类型是int64
	startStr = strconv.FormatInt(startInt,10)

	endTime, _ := time.ParseInLocation(timeLayout, endDate, loc) 	//使用模板在对应时区转化为time.time类型
	endInt = endTime.Unix()                                         //转化为时间戳 类型是int64
	endStr = strconv.FormatInt(endInt,10)
	return
}

//GetChangeWeekToUnix +N周开始结束时间
func GetChangeWeekToUnix(w int)(startInt, endInt int64, startStr, endStr string){
	now := time.Now()
	offset := int(time.Monday - now.Weekday())
	if offset > 0 { //周日为0，周一 - 周日 大于0
		offset = -6
	}

	weekStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset + 7*w)
	startDate := weekStart.Format("2006-01-02") + " 00:00:00"

	weekEnd := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset + 6 + 7*w)
	endDate := weekEnd.Format("2006-01-02") + " 23:59:59"

	timeLayout := "2006-01-02 15:04:05"                         //转化所需模板
	loc, _ := time.LoadLocation("Local")                  //重要：获取时区

	startTime, _ := time.ParseInLocation(timeLayout, startDate, loc) 	//使用模板在对应时区转化为time.time类型
	startInt = startTime.Unix()                                         //转化为时间戳 类型是int64
	startStr = strconv.FormatInt(startInt,10)

	endTime, _ := time.ParseInLocation(timeLayout, endDate, loc) 	//使用模板在对应时区转化为time.time类型
	endInt = endTime.Unix()                                         //转化为时间戳 类型是int64
	endStr = strconv.FormatInt(endInt,10)
	return
}

//GetChangeMonthToUnix +N月开始结束时间
func GetChangeMonthToUnix(m int)(startInt, endInt int64, startStr, endStr string){
	startDate := time.Now().AddDate(0,m,- time.Now().Day() + 1).Format("2006-01-02") + " 00:00:00"
	endDate := time.Now().AddDate(0,m+1,- time.Now().Day()).Format("2006-01-02") + " 23:59:59"

	timeLayout := "2006-01-02 15:04:05"                         //转化所需模板
	loc, _ := time.LoadLocation("Local")                  //重要：获取时区

	startTime, _ := time.ParseInLocation(timeLayout, startDate, loc) 	//使用模板在对应时区转化为time.time类型
	startInt = startTime.Unix()                                         //转化为时间戳 类型是int64
	startStr = strconv.FormatInt(startInt,10)

	endTime, _ := time.ParseInLocation(timeLayout, endDate, loc) 	//使用模板在对应时区转化为time.time类型
	endInt = endTime.Unix()                                         //转化为时间戳 类型是int64
	endStr = strconv.FormatInt(endInt,10)
	return
}

//GetWeekByDay 日期转星期
func GetWeekByDay(day string)(wInt int, week string, weekEn string){
	timeLayout := "2006-01-02"                         //转化所需模板
	loc, _ := time.LoadLocation("Local")                  //重要：获取时区

	dayTime, _ := time.ParseInLocation(timeLayout, day, loc) 	//使用模板在对应时区转化为time.time类型
	wInt = int(dayTime.Weekday())

	switch wInt {
	case 0:
		weekEn = "Sunday"
		week = "星期日"
		break
	case 1:
		weekEn = "Monday"
		week = "星期一"
		break
	case 2:
		weekEn = "Tuesday"
		week = "星期二"
		break
	case 3:
		weekEn = "Wednesday"
		week = "星期三"
		break
	case 4:
		weekEn = "Thursday"
		week = "星期四"
		break
	case 5:
		weekEn = "Friday"
		week = "星期五"
		break
	case 6:
		weekEn = "Saturday"
		week = "星期六"
		break
	default:
		break
	}
	return
}


