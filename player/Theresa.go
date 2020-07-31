package player

type TheresaApocalypse struct {
	idol
}

func (t *TheresaApocalypse) DeepCopy() Player {
	return &TheresaApocalypse{
		idol: t.deepCopyIdol(),
	}
}

func (t *TheresaApocalypse) RoundAttack(round uint16) {
	if t.tryRecover() {
		return
	}
	if round%3 == 0 {
		for i := 1; i <= 5; i++ {
			t.Rival.TakeDamage(round, 16, Unique)
		}
		t.tryWeakenRival()
		return
	}
	t.Rival.TakeDamage(round, t.Attack, Normal)
	t.tryWeakenRival()
}

func (t *TheresaApocalypse) tryWeakenRival() {
	if t.Rand(30) {
		t.Rival.Attributes().Defence -= 5
		if t.Rival.Attributes().Defence < 0 {
			t.Rival.Attributes().Defence = 0
		}
	}
}
