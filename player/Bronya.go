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
		b.Rival.DirectTakeDamage(round, b.RandNum(100), Unique)
		return
	}
	b.Rival.TakeDamage(round, b.Attack, Normal)
	if b.Rand(25) {
		for i := 1; i <= 4; i++ {
			b.Rival.TakeDamage(round, 12, Normal)
		}
	}
}
