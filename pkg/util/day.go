package util

import (
	"github.com/jinzhu/now"
	"strconv"
	"time"
)

func WorkDayOfMonth(currentTime string) ([]interface{}, int) {
	currentDate, _ := time.ParseInLocation("2006-01-02 15:04:05", currentTime, time.Local)
	year := currentDate.Year()
	month, _ := strconv.Atoi(currentDate.Month().String())
	count := DaysOfMonth(year, month)
	bmonth := now.New(currentDate).BeginningOfMonth()
	var monthArr []string
	var tempArr []string
	var weekArr []interface{}
	for i := 0; i < count; i++ {
		dd := bmonth.AddDate(0, 0, i)
		if dd.Weekday().String() != "Saturday" && dd.Weekday().String() != "Sunday" {
			monthArr = append(monthArr, dd.Format("2006-01-02"))
			if len(tempArr) > 0 {
				aa, _ := time.ParseInLocation("2006-01-02", tempArr[len(tempArr)-1], time.Local)
				dd, _ = time.ParseInLocation("2006-01-02", dd.Format("2006-01-02"), time.Local)
				if dd.Sub(aa).Hours()/24 > 1 {
					weekArr = append(weekArr, tempArr)
					tempArr = nil
				}
			}
			tempArr = append(tempArr, dd.Format("2006-01-02"))
		}
	}
	return weekArr, count
}

func DaysOfMonth(year int, month int) (days int) {
	if month != 2 {
		if month == 4 || month == 6 || month == 9 || month == 11 {
			days = 30

		} else {
			days = 31
		}
	} else {
		if ((year%4) == 0 && (year%100) != 0) || (year%400) == 0 {
			days = 29
		} else {
			days = 28
		}
	}
	return days
}

func GetDetailDayByMonth(monthNo int) []string {
	year := strconv.Itoa(time.Now().Year())
	monthStr := strconv.Itoa(monthNo)
	if len(monthStr) == 1 {
		monthStr = "0" + monthStr
	}
	timeStr := year + "-" + monthStr + "-01"
	baseDay, err := time.ParseInLocation("2006-01-02", timeStr, time.Local)
	if err != nil {
		println(err.Error())
	}
	println(baseDay.Format("2006-01-02"))
	currentMonth := baseDay.Month()
	monthDayStrArr := []string{}
	monthDayValueArr := []time.Time{}
	supposeTime := baseDay
	for lastMonthDayStrArrMonth := currentMonth; true; {
		monthDayValueArr = append(monthDayValueArr, supposeTime)
		monthDayStrArr = append(monthDayStrArr, supposeTime.Format("2006-01-02"))
		supposeTime = supposeTime.Add(24 * time.Hour)
		lastMonthDayStrArrMonth = supposeTime.Month()
		if lastMonthDayStrArrMonth != currentMonth {
			break
		}
	}
	return monthDayStrArr
}
