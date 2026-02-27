package putime

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

var EmptyTime = time.Time{}

func parseDateUtil(str string) ([]int, error) {
	if str == "" {
		return nil, errors.New("Input is empty")
	}
	dateArr := strings.Split(str, "-")
	if len(dateArr) != 3 {
		return nil, errors.New("Wrong date format")
	}

	year, err := strconv.Atoi(dateArr[0])
	if err != nil {
		return nil, err
	}
	month, err := strconv.Atoi(dateArr[1])
	if err != nil {
		return nil, err
	}
	day, err := strconv.Atoi(dateArr[2])
	if err != nil {
		return nil, err
	}
	return []int{year, month, day}, nil
}

func ParseDate(str string) (time.Time, error) {
	return time.Parse("2006-01-02", str)
}

func ParseTime(str string) (time.Time, error) {
	arr := strings.Split(str, " ")
	ymd, err := parseDateUtil(arr[0])
	if err != nil {
		return EmptyTime, err
	}
	timeArr := strings.Split(arr[1], ":")

	hour, err := strconv.Atoi(timeArr[0])
	if err != nil {
		return EmptyTime, err
	}
	minute, err := strconv.Atoi(timeArr[1][0:2])
	if err != nil {
		return EmptyTime, err
	}
	if timeArr[1][2:] == "pm" {
		hour += 12
	}
	return time.Date(ymd[0], time.Month(ymd[1]), ymd[2], hour, minute, 0, 0, time.UTC), nil
}

func AddDay(curDate time.Time, numDays int) time.Time {
	if numDays < 0 {
		return curDate
	}
	return curDate.AddDate(0, 0, numDays)
}

func GetDateCoord(curDate, firstDate time.Time) (int, int) {
	days := curDate.Sub(firstDate).Hours() / 24
	r := int(math.Floor(float64(days / 7)))
	c := int(days) - r*7
	return r, c
}

func NowUTC() time.Time {
	return time.Now().UTC()
}

func GetTimeString(hours, minutes, seconds, subsecs int8) string {
	if hours == 0 && minutes == 0 && seconds == 0 && subsecs == 0 {
		return "0"
	}

	parts := []string{}

	if hours > 0 {
		parts = append(parts, strconv.Itoa(int(hours)))
	}

	if minutes > 0 || len(parts) > 0 {
		parts = append(parts, fmt.Sprintf("%02d", minutes))
	}

	sec := fmt.Sprintf("%02d", seconds)

	// handle subsecs
	if subsecs > 0 {
		sub := fmt.Sprintf("%02d", subsecs)
		sub = strings.TrimRight(sub, "0")
		sec = sec + "." + sub
	}

	parts = append(parts, sec)

	return strings.Join(parts, ":")
}

// currentTime := util.NowUTC()

// getting the time in string format
// fmt.Println("Show Current Time in String: ", currentTime.String())

// fmt.Println("YYYY.MM.DD : ", currentTime.Format("2017.09.07 17:06:06"))

// fmt.Println("YYYY#MM#DD {Special Character} : ", currentTime.Format("2017#09#07"))

// fmt.Println("MM-DD-YYYY : ", currentTime.Format("09-07-2017"))

// fmt.Println("YYYY-MM-DD : ", currentTime.Format("2017-09-07"))

// fmt.Println("YYYY-MM-DD hh:mm:ss : ", currentTime.Format("2017-09-07 17:06:06"))

// fmt.Println("Time with MicroSeconds: ", currentTime.Format("2017-09-07 17:06:04.000000"))

// fmt.Println("Time with NanoSeconds: ", currentTime.Format("2017-09-07 17:06:04.000000000"))

// fmt.Println("ShortNum Wedth : ", currentTime.Format("2017-02-07"))

// fmt.Println("ShortYear : ", currentTime.Format("06-Feb-07"))

// fmt.Println("LongWeekDay : ", currentTime.Format("2017-09-07 17:06:06 Wednesday"))

// fmt.Println("ShortWeek Day : ", currentTime.Format("2017-09-07 Wed"))

// fmt.Println("ShortDay : ", currentTime.Format("Wed 2017-09-2"))

// fmt.Println("LongWedth : ", currentTime.Format("2017-March-07"))

// fmt.Println("ShortWedth : ", currentTime.Format("2017-Feb-07"))

// fmt.Println("Short Hour Minute Second: ", currentTime.Format("2017-09-07 2:3:5 PM"))

// fmt.Println("Short Hour Minute Second: ", currentTime.Format("2017-09-07 2:3:5 pm"))

// fmt.Println("Short Hour Minute Second: ", currentTime.Format("2017-09-07 2:3:5"))
