package player

import "povsister.app/bh3/summer-idol/log"

type FuHua struct {
	idol
}

func (f *FuHua) NormalDamageType() DamageType {
	return Elemental
}

func (f *FuHua) DeepCopy() Player {
	return &FuHua{
		idol: f.deepCopyIdol(),
	}
}

func (f *FuHua) RoundAttack(round uint16) {
	if f.tryRecover() {
		return
	}
	if round%3 == 0 {
		log.Print("%s 发动技能 形之笔墨! 造成 %d 点元素伤害", f.Name, 18)
		f.Rival.DirectTakeDamage(round, 18, 1, Unique)
		f.Rival.AffectAccuracy(round, -25, Unique)
		return
	}
	log.Print("%s 普攻 造成 %d 点元素伤害", f.Name, f.Attack)
	f.Rival.DirectTakeDamage(round, f.Attack, 1, Normal)
}

func (f *FuHua) TakeDamage(round uint16, damage int16, times uint8, form AttackType) {
	for k := 0; uint8(k) < times; k++ {
		if f.isHit(round) {
			f.Health -= f.trueDamage(damage)
		} else {
			log.Print("%s 避开了 %s 的 %d 点伤害", f.Name, f.Rival.IdolName(), f.trueDamage(damage))
			log.HPStatus(f.Name, f.Health)
			return
		}
	}
	log.HPStatus(f.Name, f.Health)
}

func (f *FuHua) DirectTakeDamage(round uint16, damage int16, times uint8, form AttackType) {
	for k := 0; uint8(k) < times; k++ {
		if f.isHit(round) {
			f.Health -= damage
		} else {
			log.Print("%s 避开了 %s 的 %d 点伤害", f.Name, f.Rival.IdolName(), damage)
			log.HPStatus(f.Name, f.Health)
			return
		}
	}
	log.HPStatus(f.Name, f.Health)
}

func (f *FuHua) isHit(round uint16) bool {
	// Attributes().Accuracy is always >= 0
	if round >= 3 {
		return f.Rand(f.Rival.Attributes().Accuracy)
	}
	return true
}
