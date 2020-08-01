package player

import "povsister.app/bh3/summer-idol/log"

type RaidenMei struct {
	idol
}

func (m *RaidenMei) DeepCopy() Player {
	return &RaidenMei{
		idol: m.deepCopyIdol(),
	}
}

func (m *RaidenMei) RoundAttack(round uint16) {
	if m.tryRecover() {
		return
	}
	if round%2 == 0 && m.Rival.CanIUseSkill(round, "雷电家的龙女仆!") {
		log.Print("%s 发动技能 雷电家的龙女仆! 造成 5 x 3 点元素伤害", m.Name)
		m.Rival.DirectTakeDamage(round, 3, 5, Unique)
		m.tryParalyze(round)
		return
	}
	log.Print("%s 普攻 造成 %d 点伤害", m.Name, m.Rival.Attributes().trueDamage(m.Attack))
	m.Rival.TakeDamage(round, m.Attack, 1, Normal)
	m.tryParalyze(round)
}

func (m *RaidenMei) tryParalyze(round uint16) {
	// only paralyze when rival got hit
	if !m.Rival.Attributes().hit {
		return
	}
	if m.Rand(30) && m.Rival.CanIUseSkill(round, "崩坏世界的歌姬!") {
		log.Print("%s 成功麻痹对方一回合", m.Name)
		m.Rival.Attributes().paralyzed = true
	}
}
