package player

type RaidenMei struct {
	idol
}

func (m *RaidenMei) DeepCopy() Player {
	return &RaidenMei{
		idol: m.deepCopyIdol(),
	}
}
