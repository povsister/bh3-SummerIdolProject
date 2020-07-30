package player

type FuHua struct {
	idol
}

func (f *FuHua) DeepCopy() Player {
	return &FuHua{
		idol: f.deepCopyIdol(),
	}
}
