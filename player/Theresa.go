package player

import "povsister.app/bh3/summer-idol/log"

type TheresaApocalypse struct {
	idol
}

func (t *TheresaApocalypse) DeepCopy() Player {
	return &TheresaApocalypse{
		idol: t.deepCopyIdol(),
	}
}

func (t *TheresaApocalypse) RoundAttack(round uint16) {
	if t.tryRecover() {
		return
	}
	if round%3 == 0 {
		log.Print("%s 发动技能 在线踢人! 造成 5 x %d 点伤害", t.Name, t.Rival.Attributes().trueDamage(16))
		t.Rival.TakeDamage(round, 16, 5, Unique)
		t.tryWeakenRival(round, Unique)
		return
	}
	log.Print("%s 普攻 造成 %d 点伤害", t.Name, t.Rival.Attributes().trueDamage(t.Attack))
	t.Rival.TakeDamage(round, t.Attack, 1, Normal)
	t.tryWeakenRival(round, Normal)
}

func (t *TheresaApocalypse) tryWeakenRival(round uint16, form AttackType) {
	if t.Rand(30) {
		t.Rival.AffectDefence(round, -5, form)
		log.Print("%s 发动技能 血犹大第一可爱! 降低对方 5 点防御", t.Name)
	}
}
