package player

import "povsister.app/bh3/summer-idol/log"

type TheTwins struct {
	idol
	revived bool
	charged bool // is the skill charged ?
}

func (t *TheTwins) DeepCopy() Player {
	return &TheTwins{
		idol: t.deepCopyIdol(), revived: t.revived, charged: t.charged,
	}
}

func (t *TheTwins) RoundAttack(round uint16) {
	if t.tryRecover() {
		return
	}
	// is the skill charged after revived
	if t.revived && t.charged && t.Rival.CanIUseSkill(round, "变成星星吧!") {
		if t.Rand(50) {
			log.Print("%s 发动技能 变成星星吧! 造成 %d 点伤害", t.Name, t.Rival.Attributes().trueDamage(233))
			t.Rival.TakeDamage(round, 233, 1, Unique)
		} else {
			log.Print("%s 发动技能 变成星星吧! 造成 %d 点伤害", t.Name, t.Rival.Attributes().trueDamage(50))
			t.Rival.TakeDamage(round, 50, 1, Unique)
		}
		t.charged = false
		return
	}
	// normal attack
	log.Print("%s 普攻 造成 %d 点伤害", t.Name, t.Rival.Attributes().trueDamage(t.Attack))
	t.Rival.TakeDamage(round, t.Attack, 1, Normal)
}

func (t *TheTwins) TakeDamage(round uint16, damage int16, times uint8, form AttackType) {
	t.takeDamage(damage, times)
	log.HPStatus(t.Name, t.Health)
	t.tryRevive(round)
}

func (t *TheTwins) DirectTakeDamage(round uint16, damage int16, times uint8, form AttackType) {
	t.directTakeDamage(damage, times)
	log.HPStatus(t.Name, t.Health)
	t.tryRevive(round)
}

// check if been revived or not
// also check if need revive
func (t *TheTwins) tryRevive(round uint16) {
	if !t.revived && t.Health <= 0 && t.Rival.CanIUseSkill(round, "96度生命之水!") {
		t.revived = true
		t.charged = true
		t.Health = 20
		log.Print("%s 发动技能 96度生命之水! 原地复活并恢复至 20 HP", t.Name)
	}
}

func (t *TheTwins) Reset() {
	t.idol.Reset()
	t.revived = false
	t.charged = false
}
