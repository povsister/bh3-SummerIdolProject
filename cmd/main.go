package main

import (
	"fmt"
	st "github.com/alexeyco/simpletable"
	"sync"

	"povsister.app/bh3/summer-idol/arena"
	"povsister.app/bh3/summer-idol/player"
)

const (
	matches = 10000
	arenas  = 10
)

func main() {
	var playerSlice []player.Candidate
	for p, _ := range player.Players {
		playerSlice = append(playerSlice, p)
	}
	table := st.New()
	var mu sync.Mutex // mutex for writing body cells
	var wg sync.WaitGroup
	for i := 0; i < len(playerSlice); i++ {
		for k := i + 1; k < len(playerSlice); k++ {
			wg.Add(1)
			go matchPair(playerSlice[i], playerSlice[k], &mu, &table.Body.Cells, &wg)
		}
	}
	table.Header = &st.Header{
		Cells: []*st.Cell{
			{Text: `对战双方`, Align: st.AlignLeft},
			{Text: `结果1(获胜)`, Align: st.AlignLeft},
			{Text: `结果2(获胜)`, Align: st.AlignLeft},
		},
	}
	table.SetStyle(st.StyleMarkdown)
	wg.Wait()
	fmt.Println(table.String())
}

func matchPair(p1, p2 player.Candidate, mu *sync.Mutex, bodyCells *[][]*st.Cell, pairWg *sync.WaitGroup) {
	defer pairWg.Done()
	resultChan := make(chan *arena.MatchResult, matches*arenas)
	var wg sync.WaitGroup
	var cells []*st.Cell
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
	cells = append(cells, &st.Cell{Text: fmt.Sprintf(`%s vs %s`,
		player.Players[p1].IdolName(), player.Players[p2].IdolName())})

	for k, v := range wins {
		cells = append(cells, &st.Cell{Text: fmt.Sprintf(`%s: %d`,
			player.Players[k].IdolName(), v), Align: st.AlignLeft})
	}
	if len(cells) <= 2 {
		if _, ok := wins[p1]; ok {
			cells = append(cells, &st.Cell{Text: fmt.Sprintf(`%s: %d`,
				player.Players[p2].IdolName(), 0), Align: st.AlignLeft})
		} else {
			cells = append(cells, &st.Cell{Text: fmt.Sprintf(`%s: %d`,
				player.Players[p1].IdolName(), 0), Align: st.AlignLeft})
		}
	}
	mu.Lock()
	*bodyCells = append(*bodyCells, cells)
	mu.Unlock()
}
