package player

import (
	"math/rand"
	"time"
)

type BronyaZaychik struct {
	idol
}

func (b *BronyaZaychik) DeepCopy() Player {
	return &BronyaZaychik{
		idol: b.deepCopyIdol(),
	}
}

func (b *BronyaZaychik) RoundAttack(defender Player, round uint16) {
	if round%3 == 0 {
		rand.Seed(time.Now().UnixNano())
		defender.DirectTakeDamage(int16(rand.Intn(100) + 1))
		//defender.DirectTakeDamage(50)
		return
	}
	defender.TakeDamage(b.Attack)
	if b.Rand(25) {
		defender.TakeDamage(4 * 12)
	}
}
