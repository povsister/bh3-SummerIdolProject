package log

import (
	"fmt"
	"math"
)

var enableLog = false

func EnableLog(log bool) {
	enableLog = log
}

func Print(template string, args ...interface{}) {
	if enableLog {
		fmt.Printf(template+"\n", args...)
	}
}

func HPStatus(playerName string, health int16) {
	if enableLog {
		fmt.Printf("%s 当前剩余 %d HP\n", playerName, health)
	}
}

func AttributeStatus(playerName string, attrName string, diff int16) {
	if enableLog {
		if diff >= 0 {
			fmt.Printf("%s 的%s上升了 %d 点\n", playerName, attrName, int(math.Abs(float64(diff))))
		} else {
			fmt.Printf("%s 的%s下降了 %d 点\n", playerName, attrName, int(math.Abs(float64(diff))))
		}
	}
}
