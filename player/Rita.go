package player

type RitaRossweisse struct {
	idol
}

func (r *RitaRossweisse) DeepCopy() Player {
	return &RitaRossweisse{
		idol: r.deepCopyIdol(),
	}
}

func (r *RitaRossweisse) RoundAttack(round uint16) {
	if r.tryRecover() {
		return
	}
	if round%4 == 0 {
		r.Rival.Attributes().Health += 4
		return
	}
	if r.Rand(35) {
		r.Rival.DirectTakeDamage(round, r.Rival.Attributes().trueDamage(r.Attack)-3, 1, Normal)
		r.Rival.Attributes().Attack -= 4
		if r.Rival.Attributes().Attack < 0 {
			r.Rival.Attributes().Attack = 0
		}
	} else {
		r.Rival.TakeDamage(round, r.Attack, 1, Normal)
	}
}

func (r *RitaRossweisse) TakeDamage(round uint16, damage int16, times uint8, form AttackType) {
	switch round % 4 {
	case 1, 2:
		if form == Unique {
			// skill damage do not take effect
			r.Health -= r.reduceDamage(round, r.Rival.Attributes().Attack-r.Defence)
			return
		}
	}
	for k := 0; uint8(k) < times; k++ {
		r.Health -= r.reduceDamage(round, damage-r.Defence)
	}
}

func (r *RitaRossweisse) DirectTakeDamage(round uint16, damage int16, times uint8, form AttackType) {
	switch round % 4 {
	case 1, 2:
		if form == Unique {
			// skill damage do not take effect
			r.Health -= r.reduceDamage(round, r.Rival.Attributes().Attack)
			return
		}
	}
	for k := 0; uint8(k) < times; k++ {
		r.Health -= r.reduceDamage(round, damage)
	}
}

func (r *RitaRossweisse) reduceDamage(round uint16, damage int16) int16 {
	if round/4 >= 1 {
		// damage * 40%
		return int16(float64(damage) / 10 * 4)
	}
	return damage
}
