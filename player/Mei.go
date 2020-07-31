package player

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
	if m.Rand(30) {
		m.Rival.Attributes().paralyzed = true
	}
	if round%2 == 0 {
		m.Rival.DirectTakeDamage(round, 3, 5, Unique)
		return
	}
	m.Rival.TakeDamage(round, m.Attack, 1, Normal)
}
