package player

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
		defender.DirectTakeDamage(b.RandNum(100), Unique)
		return
	}
	defender.TakeDamage(b.Attack, Normal)
	if b.Rand(25) {
		for i := 1; i <= 4; i++ {
			defender.TakeDamage(12, Normal)
		}
	}
}
