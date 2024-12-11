package main

import (
	"badminton-ustb/order"
	"fmt"
	"os"
	"time"
)

// 获取加 7 天后的日期
func main() {
	nextDate := getNextDate(time.Now(), 7)
	randomNumbers := generateUniqueRandomNumbers(1, 20, 2)
	result, err := order.SendOrder(nextDate, "Y", randomNumbers)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Request failed.")
	}

	fmt.Println(result)
}
