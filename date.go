package dateTool

import "time"

//FirstDayMonday 如果为true，表示每周的第一天为周一，否则为周日
var FirstDayMonday bool

//时间格式化字符创切片
var TimeFormats = []string{"1/2/2006", "1/2/2006 15:4:5", "2006-1-2 15:4:5", "2006-1-2 15:4", "2006-1-2", "1-2", "15:4:5", "15:4", "15", "15:4:5 Jan 2, 2006 MST", "2006-01-02 15:04:05.999999999 -0700 MST"}

//时间结构体
type Now struct {
	time.Time
}

//创建时间函数
func New(t time.Time) *Now {
	return &Now{t}
}

//当前时间 精确到分钟，其它更小单元值忽略 使用默认零值
func BeginningOfMinute() time.Time {
	return New(time.Now()).BeginningOfMinute()
}

//当前时间 精确到小时，其它更小单元值忽略 使用默认零值
func BeginningOfHour() time.Time {
	return New(time.Now()).BeginningOfHour()
}

//当前时间 精确到天，其它更小单元值忽略 使用默认零值
func BeginningOfDay() time.Time {
	return New(time.Now()).BeginningOfDay()
}

//当前时间精确到天，表示周开始的第一天时间。默认周日为一周的第一天，，其它更小单元值忽略 使用默认零值
func BeginningOfWeek() time.Time {
	return New(time.Now()).BeginningOfWeek()
}

//当前月开始的第一天。 时间精确到天，表示月开始的第一天时间，其它更小单元值忽略 使用默认零值
func BeginningOfMonth() time.Time {
	return New(time.Now()).BeginningOfMonth()
}

//当前时间 所处在的季节开始时间  精确到秒
func BeginningOfQuarter() time.Time {
	return New(time.Now()).BeginningOfQuarter()
}

// 当前时间 年开始时间，精确到秒， 其它时间默认为零值
func BeginningOfYear() time.Time {
	return New(time.Now()).BeginningOfYear()
}

//当前时间  当前分钟的最后时间 精确到秒
func EndOfMinute() time.Time {
	return New(time.Now()).EndOfMinute()
}

//当前时间 当前小时的最后时间 精确到秒
func EndOfHour() time.Time {
	return New(time.Now()).EndOfHour()
}

//当前时间 当前天的最后时间  精确到秒
func EndOfDay() time.Time {
	return New(time.Now()).EndOfDay()
}

//当前时间 当前周的最后时间  精确到秒
func EndOfWeek() time.Time {
	return New(time.Now()).EndOfWeek()
}

//当前时间 当前月的最后时间  精确到秒
func EndOfMonth() time.Time {
	return New(time.Now()).EndOfMonth()
}

//当前时间 所处在的季节结束时间  精确到秒
func EndOfQuarter() time.Time {
	return New(time.Now()).EndOfQuarter()
}

// 当前时间  年最后的时间  精确秒
func EndOfYear() time.Time {
	return New(time.Now()).EndOfYear()
}

//获取当前时间 周一开始时间
func Monday() time.Time {
	return New(time.Now()).Monday()
}

//获取当前时间 周日开始时间  精确秒
func Sunday() time.Time {
	return New(time.Now()).Sunday()
}

//获取当前时间 周日结束时间   精确秒
func EndOfSunday() time.Time {
	return New(time.Now()).EndOfSunday()
}

//当前时间 转化为自定义格式
func Parse(strs ...string) (time.Time, error) {
	return New(time.Now()).Parse(strs...)
}

func MustParse(strs ...string) time.Time {
	return New(time.Now()).MustParse(strs...)
}

//判断当前时间是否在time1与time2时间之间
func Between(time1, time2 string) bool {
	return New(time.Now()).Between(time1, time2)
}
