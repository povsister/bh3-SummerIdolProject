package player

type SeeleVollerei struct {
	idol
}

func (s *SeeleVollerei) DeepCopy() Player {
	return &SeeleVollerei{
		idol: s.deepCopyIdol(),
	}
}
