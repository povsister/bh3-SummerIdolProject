package arena

import (
	"povsister.app/bh3/summer-idol/player"
	"sync"
)

type arena struct {
	Round         uint16
	Rivals        [2]player.Player
	simulateTimes int    // simulate times of a single match
	firstAttack   *uint8 // log the player who attack first
	wg            *sync.WaitGroup
}

type MatchResult struct {
	Rivals [2]player.Candidate
	Winner player.Candidate
}

func NewMatch(wg *sync.WaitGroup, rivals ...player.Candidate) *arena {
	if len(rivals) != 2 {
		panic(`rivals must be a pair`)
	}
	player.InitRandGenerator()
	return &arena{
		1, [2]player.Player{
			player.Players[rivals[0]].DeepCopy(),
			player.Players[rivals[1]].DeepCopy(),
		}, 1000, nil, wg,
	}
}

func (a *arena) SetMatchTimes(times int) {
	a.simulateTimes = times
}

func (a *arena) StartMatch(ch chan *MatchResult) {
	defer a.wg.Done()
	var attacker, defender player.Player
	a.Rivals[0].Attributes().Rival = a.Rivals[1]
	a.Rivals[1].Attributes().Rival = a.Rivals[0]
	for i := 1; i <= a.simulateTimes; i++ {
		for {
			// must call attacker first
			attacker = a.attacker()
			defender = a.defender()
			// do the attack
			attacker.RoundAttack(defender, a.Round)
			if a.CheckResult(ch, attacker, defender) {
				break
			}
			// defender alive. Swap the attacker and defender
			attacker, defender = defender, attacker
			// do the attack
			attacker.RoundAttack(defender, a.Round)
			if a.CheckResult(ch, attacker, defender) {
				break
			}
			// no one died, continue to the next round
			a.NextRound()
		}
	}
}

func (a *arena) NextRound() {
	a.Round += 1
}

func (a *arena) CheckResult(ch chan *MatchResult, ps ...player.Player) bool {
	for i, p := range ps {
		if p.IsDead() {
			result := &MatchResult{
				[2]player.Candidate{
					a.Rivals[0].Attributes().ID,
					a.Rivals[1].Attributes().ID,
				},
				ps[1-i].Attributes().ID,
			}
			// reset arena and rivals status
			a.Reset()
			a.Rivals[0].Reset()
			a.Rivals[1].Reset()
			ch <- result
			return true
		}
	}
	return false
}

func (a *arena) Reset() {
	a.Round = 1
}

// return a player who should attack first
// with a pointer cache trick
func (a *arena) attacker() player.Player {
	if a.firstAttack != nil {
		return a.Rivals[*a.firstAttack]
	}
	var tmp uint8
	if a.Rivals[0].Attributes().Speed >= a.Rivals[1].Attributes().Speed {
		tmp = 0
		a.firstAttack = &tmp
		return a.Rivals[0]
	}
	tmp = 1
	a.firstAttack = &tmp
	return a.Rivals[1]
}

// return a player who should defend
func (a *arena) defender() player.Player {
	return a.Rivals[1-*a.firstAttack]
}
