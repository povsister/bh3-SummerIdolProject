package player

import "povsister.app/bh3/summer-idol/log"

type BiankaAtaegina struct {
	idol
	fightBack            *bool
	fightBackAtLastRound bool
}

func (d *BiankaAtaegina) ResetRound() {
	d.idol.ResetRound()
	if d.fightBack != nil && *d.fightBack {
		d.fightBackAtLastRound = true
	} else {
		d.fightBackAtLastRound = false
	}
	d.fightBack = nil
}

func (d *BiankaAtaegina) isFightingBack(form AttackType) bool {
	if d.fightBack != nil {
		return *d.fightBack
	}
	var tmp bool
	if form == Unique && d.Rand(16) {
		tmp = true
		d.fightBack = &tmp
		return true
	}
	tmp = false
	d.fightBack = &tmp
	return false
}

func (d *BiankaAtaegina) DeepCopy() Player {
	return &BiankaAtaegina{
		idol: d.deepCopyIdol(), fightBack: nil, fightBackAtLastRound: false,
	}
}

func (d *BiankaAtaegina) RoundAttack(round uint16) {
	if d.paralyzed || d.stunned || d.frozen {
		if d.fightBack != nil && !*d.fightBack {
			// means the rival attacks first and no fightback
			d.resetStatus()
			return
		}
		if !d.fightBackAtLastRound && d.fightBack == nil {
			// means we attack first and no fightback
			d.resetStatus()
			return
		}
		if d.frozen {
			log.Print("%s 因弹反免疫了对方的冰冻效果", d.Name)
		} else if d.paralyzed {
			log.Print("%s 因弹反免疫了对方的麻痹效果", d.Name)
		} else if d.stunned {
			log.Print("%s 因弹反免疫了对方的眩晕效果", d.Name)
		} else {
			panic("unhandled abnormal status")
		}
		// reset the frozen/stunned/paralyzed status
		d.resetStatus()
	}
	d.Attack += 3
	log.Print("%s 的攻击上升了 3 点", d.Name)
	log.Print("%s 普攻 造成 %d 点伤害", d.Name, d.Rival.Attributes().trueDamage(d.Attack))
	d.Rival.TakeDamage(round, d.Attack, 1, Normal)
}

func (d *BiankaAtaegina) TakeDamage(round uint16, damage int16, times uint8, form AttackType) {
	if d.isFightingBack(form) {
		log.Print("%s 触发弹反! 免疫伤害并返还 %d 点伤害", d.Name, d.Rival.Attributes().trueDamage(30))
		d.Rival.TakeDamage(round, 30, 1, Normal)
		return
	}
	// take normal damage or skill not triggered
	d.takeDamage(damage, times)
	log.HPStatus(d.Name, d.Health)
}

func (d *BiankaAtaegina) DirectTakeDamage(round uint16, damage int16, times uint8, form AttackType) {
	if d.isFightingBack(form) {
		log.Print("%s 触发弹反! 免疫伤害并返还 %d 点伤害", d.Name, d.Rival.Attributes().trueDamage(30))
		d.Rival.TakeDamage(round, 30, 1, Normal)
		return
	}
	d.directTakeDamage(damage, times)
	log.HPStatus(d.Name, d.Health)
}

func (d *BiankaAtaegina) AffectAttr(aT attrType, round uint16, num int16, form AttackType) {
	switch aT {
	case attrHealth:
		if d.hit {
			d.affectHealth(round, num, form)
		}
	case attrAttack:
		if d.hit {
			d.affectAttack(round, num, form)
		}
	case attrDefence:
		if d.hit {
			d.affectDefence(round, num, form)
		}
	case attrAccuracy:
		if d.hit {
			d.affectAccuracy(round, num, form)
		}
	default:
		panic(`unknown attrType`)
	}
}

func (d *BiankaAtaegina) affectHealth(round uint16, num int16, form AttackType) {
	if d.isFightingBack(form) {
		// no effect
		log.Print("%s 触发弹反! 免疫对方对己方生命值的影响", d.Name)
		return
	}
	d.Health += num
	log.AttributeStatus(d.Name, "生命值", num)
}

func (d *BiankaAtaegina) affectAccuracy(round uint16, num int16, form AttackType) {
	if d.isFightingBack(form) {
		// no effect
		log.Print("%s 触发弹反! 免疫对方对己方命中率的影响", d.Name)
		return
	}
	d.Accuracy += num
	if d.Accuracy < 0 {
		d.Accuracy = 0
	} else if d.Accuracy > 100 {
		d.Accuracy = 100
	}
	log.AttributeStatus(d.Name, "命中率", num)
}

func (d *BiankaAtaegina) affectAttack(round uint16, num int16, form AttackType) {
	if d.isFightingBack(form) {
		// no effect
		log.Print("%s 触发弹反! 免疫对方对己方攻击的影响", d.Name)
		return
	}
	d.Attack += num
	if d.Attack < 0 {
		d.Attack = 0
	}
	log.AttributeStatus(d.Name, "攻击", num)
}

func (d *BiankaAtaegina) affectDefence(round uint16, num int16, form AttackType) {
	if d.isFightingBack(form) {
		// no effect
		log.Print("%s 触发弹反! 免疫对方对己方防御的影响", d.Name)
		return
	}
	d.Defence += num
	if d.Defence < 0 {
		d.Defence = 0
	}
	log.AttributeStatus(d.Name, "防御", num)
}
