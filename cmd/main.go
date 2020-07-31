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
	//var playerSlice []player.Candidate
	//for p, _ := range player.Players {
	//	playerSlice = append(playerSlice, p)
	//}
	//for i, p := range playerSlice {
	//	for k := i+1; k <= len(playerSlice); k++ {
	//		matchPair(p, playerSlice[k])
	//	}
	//}
	matchPair(player.Seele, player.Bronya)
}

func matchPair(p1, p2 player.Candidate) {
	resultChan := make(chan *arena.MatchResult, matches*arenas)
	var wg sync.WaitGroup
	for i := 1; i <= arenas; i++ {
		match := arena.NewMatch(&wg, p1, p2)
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
	fmt.Println(`=============================`)
	for k, v := range wins {
		fmt.Printf("%s wins: %d\n", player.Players[k].IdolName(), v)
	}
	fmt.Println(`=============================`)
}
