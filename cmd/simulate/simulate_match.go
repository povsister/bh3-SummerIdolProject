package main

import (
	"fmt"
	"os"
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

	fmt.Println("选择对战两人的序号")
	index := 0
	for k, v := range player.Players {
		if index%4 == 0 {
			fmt.Println()
		}
		fmt.Printf("%d)%s   ", k, v.IdolName())
		index += 1
	}
	fmt.Println()
	fmt.Println()

	var p1, p2 uint8
	for {
		if p1 == 0 {
			fmt.Print("输入P1序号: ")
			_, err := fmt.Fscanln(os.Stdin, &p1)
			if err != nil {
				fmt.Println(err)
			}
			if p1 != 0 {
				if _, ok := player.Players[player.Candidate(p1)]; !ok {
					fmt.Println("无效P1序号", p1)
					p1 = 0
				}
			} else {
				fmt.Println("P1 is nil")
			}
		}

		if p2 == 0 {
			fmt.Print("输入P2序号: ")
			_, err := fmt.Fscanln(os.Stdin, &p2)
			if err != nil {
				fmt.Println(err)
			}
			if p1 != 0 {
				if _, ok := player.Players[player.Candidate(p2)]; !ok {
					fmt.Println("无效P2序号", p2)
					p2 = 0
				}
			} else {
				fmt.Println("P2 is nil")
			}
		}

		if p1 != 0 && p2 != 0 {
			break
		}
	}
	fmt.Println()
	pair(player.Candidate(p1), player.Candidate(p2))

	fmt.Println()
	fmt.Println("Press enter to continue")
	_, _ = fmt.Fscanln(os.Stdin)
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
