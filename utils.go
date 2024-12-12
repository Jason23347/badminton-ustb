package main

import (
	"math/rand"
	"time"
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

// filter is a generic function that filters a slice based on a condition
func filter[T any](items []T, condition func(T) bool) []T {
	var result []T
	for _, item := range items {
		if condition(item) {
			result = append(result, item)
		}
	}
	return result
}

func mapping[T any, C any](items []T, mapFunc func(T) C) []C {
	var result []C
	for _, item := range items {
		result = append(result, mapFunc(item))
	}
	return result
}
