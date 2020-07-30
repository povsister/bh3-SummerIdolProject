package player

type TheTwins struct {
	idol
}

func (t *TheTwins) DeepCopy() Player {
	return &TheTwins{
		idol: t.deepCopyIdol(),
	}
}
