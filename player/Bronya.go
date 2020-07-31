package player

type BronyaZaychik struct {
	idol
}

func (b *BronyaZaychik) DeepCopy() Player {
	return &BronyaZaychik{
		idol: b.deepCopyIdol(),
	}
}

func (b *BronyaZaychik) RoundAttack(round uint16) {
	if b.tryRecover() {
		return
	}
	if round%3 == 0 {
		b.Rival.DirectTakeDamage(round, b.RandNum(100), 1, Unique)
		return
	}
	b.Rival.TakeDamage(round, b.Attack, 1, Normal)
	if b.Rand(25) {
		b.Rival.TakeDamage(round, 12, 4, Normal)
	}
}
