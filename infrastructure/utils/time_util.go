package utils

import "time"

func GetTimeStamp() int64 {
	return time.Now().Unix()
}

func GetTimestampOfXDaysAgo(days int) int64 {
	t := time.Now()

	// 计算x天前的时间
	t = t.AddDate(0, 0, -days)

	// 返回时间戳
	return t.Unix()
}
