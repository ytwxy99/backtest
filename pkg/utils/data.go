package utils

import (
	"sort"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

// GetNowTimeStamp get now timestamp
func GetNowTimeStamp() int64 {
	return time.Now().Unix()
}

// GetNowTime get now time.time
func GetNowTime() time.Time {
	nowTimstamp := time.Now().Unix()
	return time.Unix(nowTimstamp, 0)
}

// GetNowData get now data
func GetNowData() string {
	nowTimstamp := time.Now().Unix()
	tm := time.Unix(nowTimstamp, 0)

	return tm.Format("2006-01-02 15:04:05")
}

func GetData(timstamp int64) string {
	tm := time.Unix(timstamp, 0)

	return tm.Format("2006-01-02 15:04:05")
}

// GetOldTimeStamp get old timestap
func GetOldTimeStamp(years int, months int, days int) int64 {
	return time.Now().AddDate(years, months, days).Unix()
}

// GetOldData get old data
func GetOldData(years int, months int, days int) string {
	return time.Now().AddDate(years, months, days).String()
}

// StringToFloat32 string convert to float32
func StringToFloat32(s string) float32 {
	f, err := strconv.ParseFloat(s, 32)
	if err != nil {
		logrus.Error("string convert to float32 error: ", err)
		return 0
	}

	return float32(f)
}

// StringToFloat64 string convert to float64
func StringToFloat64(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		logrus.Error("string convert to float32 error: ", err)
		return 0
	}

	return float64(f)
}

// Float32ToString float32 convert to string
func Float32ToString(f float32) string {
	return strconv.FormatFloat(float64(f), 'g', 10, 64)
}

type KVs2f32 struct {
	Key   string
	Value float32
}

func (*KVs2f32) MapSortStringFloat32(input map[string]float32) []KVs2f32 {
	var kv []KVs2f32
	for key, value := range input {
		kv = append(kv, KVs2f32{
			Key:   key,
			Value: value,
		})
	}

	sort.Slice(kv, func(i, j int) bool {
		return kv[i].Value > kv[j].Value // 降序
	})

	return kv
}

type KVs2i64 struct {
	Key   string
	Value int64
}

func (*KVs2i64) MapSortStringInt64(input map[string]int64) []KVs2i64 {
	var kv []KVs2i64
	for key, value := range input {
		kv = append(kv, KVs2i64{
			Key:   key,
			Value: value,
		})
	}

	sort.Slice(kv, func(i, j int) bool {
		return kv[i].Value > kv[j].Value // 降序
	})

	return kv
}
