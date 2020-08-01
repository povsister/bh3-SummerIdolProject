package main

import (
	"povsister.app/bh3/summer-idol/arena"
	"povsister.app/bh3/summer-idol/log"
	"povsister.app/bh3/summer-idol/player"
	"sync"
)

const (
	simulateTimes = 100
)

func main() {
	log.EnableLog(true)
	pair(player.Seele, player.Rita)
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
		_, ok := <-resultChan
		if !ok {
			break
		}
	}
}
