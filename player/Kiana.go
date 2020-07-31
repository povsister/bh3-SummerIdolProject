package player

import "povsister.app/bh3/summer-idol/log"

type KianaKaslana struct {
	idol
}

func (k *KianaKaslana) DeepCopy() Player {
	return &KianaKaslana{
		idol: k.deepCopyIdol(),
	}
}

func (k *KianaKaslana) RoundAttack(round uint16) {
	if k.tryRecover() {
		return
	}
	if round%2 == 0 {
		log.Print("%s 发动技能 吃我一矛! 造成 %d 点伤害", k.Name, k.Rival.Attributes().trueDamage(k.Attack+2*k.Rival.Attributes().Defence))
		k.Rival.TakeDamage(round, k.Attack+2*k.Rival.Attributes().Defence, 1, Unique)
		if k.Rand(35) {
			log.Print("%s 因为 音浪~太强~ 眩晕一回合", k.Name)
			k.stunned = true
		}
	} else {
		log.Print("%s 普攻 造成 %d 点伤害", k.Name, k.Rival.Attributes().trueDamage(k.Attack))
		k.Rival.TakeDamage(round, k.Attack, 1, Normal)
	}
}
