package player

import "povsister.app/bh3/summer-idol/log"

type MurataHimeko struct {
	idol
}

func (h *MurataHimeko) DeepCopy() Player {
	return &MurataHimeko{
		idol: h.deepCopyIdol(),
	}
}

func (h *MurataHimeko) RoundAttack(round uint16) {
	if h.tryRecover() {
		return
	}
	// lower 35% accuracy every 2 round
	if h.Accuracy <= 0 {
		log.Print("%s 的命中率为0! 攻击不可能命中", h.Name)
		return
	}
	if round%2 == 0 && h.Rival.CanIUseSkill(round, "干杯,朋友!") {
		h.Attack *= 2
		log.Print("%s 发动技能 干杯,朋友! 使自己攻击力翻倍", h.Name)
		h.Accuracy -= 35
		log.Print("%s 的命中率下降 35 点", h.Name)
		h.Accuracy = notLessZero(h.Accuracy)
		if h.Rand(h.Accuracy) {
			log.Print("%s 发动技能 干杯,朋友! 造成 %d 点伤害", h.Name, h.getRealDamage(h.Attack))
			h.Rival.DirectTakeDamage(round, h.getRealDamage(h.Attack), 1, Skill)
		} else {
			log.Print("%s 发动技能 干杯,朋友! 可惜未能命中", h.Name)
		}
		return
	}
	if h.Rand(h.Accuracy) {
		log.Print("%s 普攻 造成 %d 点伤害", h.Name, h.getRealDamage(h.Attack))
		h.Rival.DirectTakeDamage(round, h.getRealDamage(h.Attack), 1, Normal)
	} else {
		log.Print("%s 普攻 未能命中", h.Name)
	}
}

func (h *MurataHimeko) getRealDamage(attack int16) int16 {
	switch h.Rival.Attributes().ID {
	case Durandal, Sakura, Twins:
		return (attack - h.Rival.Attributes().Defence) * 2
	}
	return attack - h.Rival.Attributes().Defence
}
