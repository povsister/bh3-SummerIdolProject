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

}
