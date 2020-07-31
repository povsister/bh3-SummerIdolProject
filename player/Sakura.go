package player

import "povsister.app/bh3/summer-idol/log"

type YaeSakura struct {
	idol
}

func (s *YaeSakura) DeepCopy() Player {
	return &YaeSakura{
		idol: s.deepCopyIdol(),
	}
}

func (s *YaeSakura) RoundAttack(round uint16) {
	if s.tryRecover() {
		return
	}
	if s.Rand(30) && s.Rival.CanUseSkill(round, "八重樱的饭团!") {
		log.Print("%s 发动技能 八重樱的饭团! 回复自身 25 HP", s.Name)
		s.Health += 25
	}
	if round%2 == 0 {
		log.Print("%s 发动技能 卡莲的饭团! 造成 25 点元素伤害", s.Name)
		s.Rival.DirectTakeDamage(round, 25, 1, Unique)
	} else {
		log.Print("%s 普攻 造成 %d 点伤害", s.Name, s.Rival.Attributes().trueDamage(s.Attack))
		s.Rival.TakeDamage(round, s.Attack, 1, Normal)
	}
}
