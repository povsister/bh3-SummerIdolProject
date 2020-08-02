package player

import "povsister.app/bh3/summer-idol/log"

type TheRaven struct {
	idol
	increaseDam bool
}

func (r *TheRaven) Reset() {
	r.idol.Reset()
	r.increaseDam = false
}

func (r *TheRaven) DeepCopy() Player {
	return &TheRaven{
		idol: r.deepCopyIdol(), increaseDam: r.increaseDam,
	}
}

func (r *TheRaven) PreBattle(round uint16) {
	if r.Rival.Attributes().ID == Kiana {
		r.increaseDam = true
		log.Print("%s 针对%s发动技能 不是针对你! 本场比赛造成的伤害提升25%", r.Name, r.Rival.IdolName())
	} else if r.Rand(25) {
		r.increaseDam = true
		log.Print("%s 发动技能 不是针对你! 本场比赛造成的伤害提升25%", r.Name)
	}
}

func (r *TheRaven) RoundAttack(round uint16) {
	if r.tryRecover() {
		return
	}
	if round%3 == 0 && r.Rival.CanIUseSkill(round, "别墅小岛!") {
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
	if r.increaseDam {
		return roundDamage(float64(damage) / 100 * 125)
	}
	return damage
}
