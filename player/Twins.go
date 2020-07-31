package player

type TheTwins struct {
	idol
	revived bool
	charged bool // is the skill charged ?
}

func (t *TheTwins) DeepCopy() Player {
	return &TheTwins{
		idol: t.deepCopyIdol(),
	}
}

func (t *TheTwins) RoundAttack(round uint16) {
	if t.tryRecover() {
		return
	}
	// is the skill charged after revived
	if t.revived && t.charged {
		if t.Rand(50) {
			t.Rival.TakeDamage(round, 233, 1, Unique)
		} else {
			t.Rival.TakeDamage(round, 50, 1, Unique)
		}
		t.charged = false
		return
	}
	// normal attack
	t.Rival.TakeDamage(round, t.Attack, 1, Normal)
}

func (t *TheTwins) TakeDamage(round uint16, damage int16, times uint8, form AttackType) {
	t.takeDamage(damage, times)
	t.tryRevive()
}

func (t *TheTwins) DirectTakeDamage(round uint16, damage int16, times uint8, form AttackType) {
	t.directTakeDamage(damage, times)
	t.tryRevive()
}

// check if been revived or not
// also check if need revive
func (t *TheTwins) tryRevive() {
	if !t.revived && t.Health <= 0 {
		t.revived = true
		t.charged = true
		t.Health = 20
	}
}

func (t *TheTwins) Reset() {
	t.idol.Reset()
	t.revived = false
	t.charged = false
}
