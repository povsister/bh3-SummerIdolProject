package player

import "povsister.app/bh3/summer-idol/log"

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
		trueDam := r.finalDamage(r.Rival.Attributes().trueDamage(16))
		log.Print("%s 发动技能 别墅小岛! 造成 7 x %d 点伤害", r.Name, trueDam)
		r.Rival.DirectTakeDamage(round, trueDam, 7, Unique)
		return
	}
	trueDam := r.finalDamage(r.Rival.Attributes().trueDamage(r.Attack))
	log.Print("%s 普攻 造成 %d 点伤害", r.Name, trueDam)
	r.Rival.DirectTakeDamage(round, trueDam, 1, Normal)
}

func (r *TheRaven) finalDamage(damage int16) int16 {
	if r.Rival.Attributes().ID == Kiana {
		// 125%
		return roundDamage(float64(damage) / 100 * 125)
	}
	if r.Rand(25) {
		return roundDamage(float64(damage) / 100 * 125)
	}
	return damage
}
