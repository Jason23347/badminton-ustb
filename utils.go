package main

import (
	"time"
	"math/rand"
)

// 生成一个不重复的随机数切片，范围在 1 到 20 之间
func generateUniqueRandomNumbers(min, max, count int) []int {
	// 创建一个新的随机数生成器
	randSource := rand.NewSource(time.Now().UnixNano())
	r := rand.New(randSource)

	// 创建一个从 min 到 max 的切片
	numbers := make([]int, max-min+1)
	for i := range numbers {
		numbers[i] = min + i
	}

	// 打乱切片中的元素
	r.Shuffle(len(numbers), func(i, j int) {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	})

	// 返回前 count 个随机数
	return numbers[:count]
}

// 生成每天加 7 天的日期
func getNextDate(currentDate time.Time, daysToAdd int) string {
	return currentDate.AddDate(0, 0, daysToAdd).Format("2006-01-02")
}
