package player

type FuHua struct {
	idol
}

func (f *FuHua) DeepCopy() Player {
	return &FuHua{
		idol: f.deepCopyIdol(),
	}
}

func (f *FuHua) RoundAttack(round uint16) {
	if f.tryRecover() {
		return
	}
	if round%3 == 0 {
		f.Rival.DirectTakeDamage(round, 18, 1, Unique)
		return
	}
	f.Rival.DirectTakeDamage(round, f.Attack, 1, Normal)
}

func (f *FuHua) TakeDamage(round uint16, damage int16, times uint8, form AttackType) {
	for k := 0; uint8(k) < times; k++ {
		if f.isHit(round) {
			f.Health -= f.trueDamage(damage)
		}
	}
}

func (f *FuHua) DirectTakeDamage(round uint16, damage int16, times uint8, form AttackType) {
	for k := 0; uint8(k) < times; k++ {
		if f.isHit(round) {
			f.Health -= damage
		}
	}
}

func (f *FuHua) isHit(round uint16) bool {
	// enemy has the 75% of accuracy after round 3
	if round == 3 && f.Speed > f.Rival.Attributes().Speed {
		return f.Rand(75)
	}
	if round > 3 {
		return f.Rand(75)
	}
	return true
}
