package utils

import "time"

func GetTimeStamp() uint32 {
	return uint32(time.Now().Unix())
}

func GetTimestampOfXDaysAgo(days int) uint32 {
	t := time.Now()

	// 计算x天前的时间
	t = t.AddDate(0, 0, -days)

	// 返回时间戳
	return uint32(t.Unix())
}
