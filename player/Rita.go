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
		r.Rival.AffectHealth(round, 4, Unique)
		return
	}
	if r.Rand(35) {
		r.Rival.DirectTakeDamage(round, r.Rival.Attributes().trueDamage(r.Attack)-3, 1, Normal)
		r.Rival.AffectAttack(round, -4, Unique)
	} else {
		r.Rival.TakeDamage(round, r.Attack, 1, Normal)
	}
}

func (r *RitaRossweisse) TakeDamage(round uint16, damage int16, times uint8, form AttackType) {
	if round > 4 {
		switch round % 4 {
		case 1, 2:
			if form == Unique {
				// skill damage do not take effect
				r.Health -= r.reduceDamage(round, r.Rival.Attributes().Attack-r.Defence)
				return
			}
		}
	}
	for k := 0; uint8(k) < times; k++ {
		r.Health -= r.reduceDamage(round, damage-r.Defence)
	}
}

func (r *RitaRossweisse) DirectTakeDamage(round uint16, damage int16, times uint8, form AttackType) {
	if round > 4 {
		switch round % 4 {
		case 1, 2:
			if form == Unique {
				// skill damage do not take effect
				r.Health -= r.reduceDamage(round, r.Rival.Attributes().Attack)
				return
			}
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

func (r *RitaRossweisse) AffectAccuracy(round uint16, num int16, form AttackType) {
	if round > 4 && (round%4 == 1 || round%4 == 2) && form == Unique {
		// no effect
		return
	}
	r.Accuracy += num
	if r.Accuracy < 0 {
		r.Accuracy = 0
	} else if r.Accuracy > 100 {
		r.Accuracy = 100
	}
}

func (r *RitaRossweisse) AffectAttack(round uint16, num int16, form AttackType) {
	if round > 4 && (round%4 == 1 || round%4 == 2) && form == Unique {
		// no effect
		return
	}
	r.Attack += num
	if r.Attack < 0 {
		r.Attack = 0
	}
}

func (r *RitaRossweisse) AffectDefence(round uint16, num int16, form AttackType) {
	if round > 4 && (round%4 == 1 || round%4 == 2) && form == Unique {
		// no effect
		return
	}
	r.Defence += num
	if r.Defence < 0 {
		r.Defence = 0
	}
}
