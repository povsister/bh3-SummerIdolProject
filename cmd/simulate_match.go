package main

import (
	"fmt"
	"povsister.app/bh3/summer-idol/arena"
	"povsister.app/bh3/summer-idol/log"
	"povsister.app/bh3/summer-idol/player"
	"sync"
)

const (
	simulateTimes = 1
)

func main() {
	log.EnableLog(true)
	pair(player.Himeko, player.Rita)
}

func pair(p1, p2 player.Candidate) {
	resultChan := make(chan *arena.MatchResult, simulateTimes)
	var wg sync.WaitGroup
	match := arena.NewMatch(&wg, p1, p2)
	match.SetMatchTimes(simulateTimes)
	wg.Add(1)
	go match.StartMatch(resultChan)
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for {
		res, ok := <-resultChan
		if !ok {
			break
		}
		fmt.Println(player.Players[res.Winner].IdolName(), "Wins !")
	}
}
