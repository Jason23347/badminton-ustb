package main

import (
	"badminton-ustb/order"
	"fmt"
	"os"
	"time"
)

// 获取加 7 天后的日期
func main() {
	var fieldList []order.Field
	var orderableFieldIndicies []int

	// Config from args
	config := order.ParseConfig()

	for i := 0; i < config.MaxLoopCount; i++ {
		if config.RequestForm {
			formResult, err := order.GetForm()
			if err != nil {
				fmt.Fprintln(os.Stderr, "获取场地信息失败！")
				if config.RequestMode == order.TEST_MODE {
					return
				} else if config.RequestMode == order.LOOP_MODE {
					continue
				}
			}

			fieldList = formResult.FieldList()
		} else {
			fieldList = order.GetDefaultFieldList()
		}

		for index, field := range fieldList {
			if field.Price > 0 {
				orderableFieldIndicies = append(orderableFieldIndicies, index)
			}
		}

		orderableCount := len(orderableFieldIndicies)
		if orderableCount <= 0 {
			if config.RequestMode == order.TEST_MODE {
				fmt.Fprintln(os.Stderr, "没有场地了！明天再来吧TAT")
				return
			} else if config.RequestMode == order.LOOP_MODE {
				time.Sleep(time.Duration(config.LoopIntervalSeconds) * time.Second)
				continue
			}
		}

		// TODO: 连号的场地
		randomNumbers := generateUniqueRandomNumbers(0, orderableCount, config.FieldCount)
		fields := mapping(randomNumbers, func(index int) order.Field {
			return fieldList[index]
		})

		orderResult, err := order.SendOrder(order.FIELD_BADMINTON, fields)
		if err != nil {
			fmt.Fprintln(os.Stderr, "订场失败！")
		}

		fmt.Println(orderResult)
	}
}
