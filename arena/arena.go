package arena

import (
	"povsister.app/bh3/summer-idol/player"
)

type arena struct {
	Round         uint16
	Rivals        [2]player.Player
	simulateTimes int   // simulate times of a single match
	lastAttack    uint8 // log the player who attacked last round
}

type MatchResult struct {
	Rivals [2]player.Candidate
	Winner player.Candidate
}

func NewMatch(rivals ...player.Player) *arena {
	if len(rivals) != 2 {
		panic(`rivals must be a pair`)
	}
	return &arena{
		1, [2]player.Player{rivals[0], rivals[1]}, 1000, 0,
	}
}

func (a *arena) SetMatchTimes(times int) {
	a.simulateTimes = times
}

func (a *arena) StartMatch(ch chan *MatchResult) {
	var attacker, defender player.Player
	var result *MatchResult
	for i := 1; i <= a.simulateTimes; i++ {
		for {
			// must call attacker first
			attacker = a.attacker()
			defender = a.defender()
			// do the attack
			attacker.RoundAttack(defender, a.Round)
			if defender.IsDead() {
				result = &MatchResult{
					[2]player.Candidate{
						a.Rivals[0].Attributes().ID,
						a.Rivals[1].Attributes().ID,
					},
					attacker.Attributes().ID,
				}
				// reset arena and rivals status
				a.Reset()
				a.Rivals[0].Reset()
				a.Rivals[1].Reset()
				break
			}
			// defender alive. Continue to the next round
			a.NextRound()
		}
		ch <- result
	}
}

func (a *arena) NextRound() {
	a.Round += 1
}

func (a *arena) Reset() {
	a.Round = 1
}

// return a player who should attack at current round
func (a *arena) attacker() player.Player {
	if a.Round == 1 {
		if a.Rivals[0].Attributes().Speed >= a.Rivals[1].Attributes().Speed {
			a.lastAttack = 0
			return a.Rivals[0]
		}
		a.lastAttack = 1
		return a.Rivals[1]
	}

	if a.lastAttack == 0 {
		a.lastAttack = 1
		return a.Rivals[1]
	}
	a.lastAttack = 0
	return a.Rivals[0]
}

// return a player who should defend at current round
func (a *arena) defender() player.Player {
	if a.lastAttack == 0 {
		return a.Rivals[1]
	}
	return a.Rivals[0]
}
