package player

import (
	"math/rand"
	"sync"
	"time"
)

var rd = rand.New(rand.NewSource(time.Now().UnixNano()))
var once sync.Once

// return rand num from [0:99)
var Rand = make(chan int, 1000)

func InitRandGenerator() {
	once.Do(initGenerator)
}

func initGenerator() {
	go func() {
		for {
			Rand <- rd.Intn(1000000)
		}
	}()
}
