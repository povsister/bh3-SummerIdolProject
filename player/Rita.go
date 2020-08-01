package player

import (
	"povsister.app/bh3/summer-idol/log"
)

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
		log.Print("%s 发动技能 完美心意! 为对方回复 4 HP 并使对方下两个回合进入魅惑状态", r.Name)
		r.Rival.AffectHealth(round, 4, Unique)
		return
	}
	if r.Rand(35) {
		damage := r.Rival.Attributes().trueDamage(r.Attack) - 3
		if damage < 0 {
			damage = 0
		}
		log.Print("%s 普攻 造成 %d 点伤害", r.Name, damage)
		r.Rival.DirectTakeDamage(round, damage, 1, Skill)
		r.Rival.AffectAttack(round, -4, Skill)
	} else {
		log.Print("%s 普攻 造成 %d 点伤害", r.Name, r.Rival.Attributes().trueDamage(r.Attack))
		r.Rival.TakeDamage(round, r.Attack, 1, Normal)
	}
}

func (r *RitaRossweisse) TakeDamage(round uint16, damage int16, times uint8, form AttackType) {
	if round > 4 {
		switch round % 4 {
		case 1, 2:
			if form == Skill || form == Unique {
				// skill damage do not take effect
				r.Health -= r.reduceDamage(round, r.Rival.Attributes().Attack-r.Defence)
				log.Print("%s 的 魅惑 生效! 对方技能变为普通攻击造成 %d 点伤害", r.Name, r.reduceDamage(round, r.Rival.Attributes().Attack-r.Defence))
				return
			}
		}
	}
	for k := 0; uint8(k) < times; k++ {
		r.Health -= r.reduceDamage(round, damage-r.Defence)
	}
	log.HPStatus(r.Name, r.Health)
}

func (r *RitaRossweisse) DirectTakeDamage(round uint16, damage int16, times uint8, form AttackType) {
	if round > 4 {
		switch round % 4 {
		case 1, 2:
			if form == Skill || form == Unique {
				// skill damage do not take effect
				r.Health -= r.reduceDamage(round, r.Rival.Attributes().Attack)
				log.Print("%s 的 魅惑 生效! 对方技能变为普通攻击造成 %d 点伤害", r.Name, r.reduceDamage(round, r.Rival.Attributes().Attack))
				return
			}
		}
	}
	for k := 0; uint8(k) < times; k++ {
		r.Health -= r.reduceDamage(round, damage)
	}
	log.HPStatus(r.Name, r.Health)
}

func (r *RitaRossweisse) reduceDamage(round uint16, damage int16) int16 {
	if round/4 >= 1 {
		// damage * 40%
		log.Print("%s 的 完美心意 生效! %s 的攻击伤害永久降低百分之60", r.Name, r.Rival.IdolName())
		return int16(float64(damage) / 10 * 4)
	}
	return damage
}

func (r *RitaRossweisse) AffectAccuracy(round uint16, num int16, form AttackType) {
	if r.immunity(round, form) {
		// no effect
		log.Print("%s 的 魅惑 生效! 免疫对方对己方命中率的影响", r.Name)
		return
	}
	r.Accuracy += num
	if r.Accuracy < 0 {
		r.Accuracy = 0
	} else if r.Accuracy > 100 {
		r.Accuracy = 100
	}
	log.AttributeStatus(r.Name, "命中率", num)
}

func (r *RitaRossweisse) AffectAttack(round uint16, num int16, form AttackType) {
	if r.immunity(round, form) {
		// no effect
		log.Print("%s 的 魅惑 生效! 免疫对方对己方攻击的影响", r.Name)
		return
	}
	r.Attack += num
	if r.Attack < 0 {
		r.Attack = 0
	}
	log.AttributeStatus(r.Name, "攻击", num)
}

func (r *RitaRossweisse) AffectDefence(round uint16, num int16, form AttackType) {
	if r.immunity(round, form) {
		// no effect
		log.Print("%s 的 魅惑 生效! 免疫对方对己方防御的影响", r.Name)
		return
	}
	r.Defence += num
	if r.Defence < 0 {
		r.Defence = 0
	}
	log.AttributeStatus(r.Name, "防御", num)
}

func (r *RitaRossweisse) immunity(round uint16, form AttackType) bool {
	return round > 4 && (round%4 == 1 || round%4 == 2) && (form == Skill || form == Unique)
}

func (r *RitaRossweisse) CanIUseSkill(round uint16, skillName string) bool {
	if round > 4 && (round%4 == 1 || round%4 == 2) {
		// no effect
		log.Print("%s 的 魅惑 生效! %s 当前回合无法使用技能 %s", r.Name, r.Rival.IdolName(), skillName)
		return false
	}
	return true
}
