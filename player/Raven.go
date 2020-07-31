package player

type TheRaven struct {
	idol
}

func (r *TheRaven) DeepCopy() Player {
	return &TheRaven{
		idol: r.deepCopyIdol(),
	}
}

func (r *TheRaven) RoundAttack(round uint16) {
	if r.tryRecover() {
		return
	}
	if round%3 == 0 {
		r.Rival.DirectTakeDamage(round, r.finalDamage(r.Rival.Attributes().trueDamage(16)), 7, Unique)
		return
	}
	r.Rival.DirectTakeDamage(round, r.finalDamage(r.Rival.Attributes().trueDamage(r.Attack)), 1, Normal)
}

func (r *TheRaven) finalDamage(damage int16) int16 {
	if r.Rival.Attributes().ID == Kiana {
		// 125%
		return damage / 100 * 125
	}
	if r.Rand(25) {
		return damage / 100 * 125
	}
	return damage
}
