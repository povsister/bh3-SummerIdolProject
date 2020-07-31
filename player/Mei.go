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
		for i := 1; i <= 5; i++ {
			m.Rival.DirectTakeDamage(round, 3, Unique)
		}
		return
	}
	m.Rival.TakeDamage(round, Normal, m.Attack)
}
