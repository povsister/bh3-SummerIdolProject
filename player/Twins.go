package player

type TheTwins struct {
	idol
	revived bool
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
	// Health should only be 20 just after revived
	if t.revived && t.Health == 20 {
		if t.Rand(50) {
			t.Rival.TakeDamage(round, Unique, 233)
		} else {
			t.Rival.TakeDamage(round, Unique, 50)
		}
		return
	}
	// normal attack
	t.Rival.TakeDamage(round, Normal, t.Attack)
}

func (t *TheTwins) TakeDamage(round uint16, form AttackType, damage int16) {
	t.Health -= t.trueDamage(damage)
	t.tryRevive()
}

func (t *TheTwins) DirectTakeDamage(round uint16, damage int16, form AttackType) {
	t.Health -= damage
	t.tryRevive()
}

// check if been revived or not
// also check if need revive
func (t *TheTwins) tryRevive() {
	if !t.revived && t.Health <= 0 {
		t.revived = true
		t.Health = 20
	}
}

func (t *TheTwins) Reset() {
	t.idol.Reset()
	t.revived = false
}
