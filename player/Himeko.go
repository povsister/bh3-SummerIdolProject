package player

type MurataHimeko struct {
	idol
}

func (h *MurataHimeko) DeepCopy() Player {
	return &MurataHimeko{
		idol: h.deepCopyIdol(),
	}
}
