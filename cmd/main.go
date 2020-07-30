package main

import (
	"fmt"
	"sync"

	"povsister.app/bh3/summer-idol/arena"
	"povsister.app/bh3/summer-idol/player"
)

const (
	matches = 5000
	arenas  = 200
)

func main() {
	resultChan := make(chan *arena.MatchResult, matches*arenas)
	var wg sync.WaitGroup
	for i := 1; i <= arenas; i++ {
		match := arena.NewMatch(&wg, player.Twins, player.Kiana)
		match.SetMatchTimes(matches)
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
