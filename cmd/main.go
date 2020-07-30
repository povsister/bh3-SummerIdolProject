package main

import (
	"fmt"
	"sync"

	"povsister.app/bh3/summer-idol/arena"
	"povsister.app/bh3/summer-idol/player"
)

func main() {
	resultChan := make(chan *arena.MatchResult, 10000)
	var wg sync.WaitGroup

	for i := 1; i <= 30; i++ {
		match := arena.NewMatch(&wg, player.Bronya, player.Sakura)
		match.SetMatchTimes(10000)
		wg.Add(1)
		go match.StartMatch(resultChan)
	}
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	wins := make(map[player.Candidate]uint)
	for {
		res, ok := <-resultChan
		if !ok {
			break
		}
		wins[res.Winner] += 1
	}
	for k, v := range wins {
		fmt.Printf("%s wins: %d\n", player.Players[k].IdolName(), v)
	}
}
