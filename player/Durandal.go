package player

import "povsister.app/bh3/summer-idol/log"

type BiankaAtaegina struct {
	idol
}

func (d *BiankaAtaegina) DeepCopy() Player {
	return &BiankaAtaegina{
		idol: d.deepCopyIdol(),
	}
}

func (d *BiankaAtaegina) RoundAttack(round uint16) {
	if d.tryRecover() {
		return
	}
	d.AffectAttack(round, 3, Normal)
	log.Print("%s 普攻 造成 %d 点伤害", d.Name, d.Rival.Attributes().trueDamage(d.Attack))
	d.Rival.TakeDamage(round, d.Attack, 1, Normal)
}

func (d *BiankaAtaegina) TakeDamage(round uint16, damage int16, times uint8, form AttackType) {
	if form == Unique && d.Rand(16) {
		d.Rival.TakeDamage(round, 30, 1, Unique)
		log.Print("%s 触发弹反! 免疫伤害并返还 %d 点伤害", d.Name, d.Rival.Attributes().trueDamage(30))
		return
	}
	// take normal damage or skill not triggered
	d.takeDamage(damage, times)
	log.HPStatus(d.Name, d.Health)
}

func (d *BiankaAtaegina) DirectTakeDamage(round uint16, damage int16, times uint8, form AttackType) {
	if form == Unique && d.Rand(16) {
		d.Rival.TakeDamage(round, 30, 1, Unique)
		log.Print("%s 触发弹反! 免疫伤害并返还 %d 点伤害", d.Name, d.Rival.Attributes().trueDamage(30))
		return
	}
	d.directTakeDamage(damage, times)
	log.HPStatus(d.Name, d.Health)
}
