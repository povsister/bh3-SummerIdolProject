package player

type TheresaApocalypse struct {
	idol
}

func (t *TheresaApocalypse) DeepCopy() Player {
	return &TheresaApocalypse{
		idol: t.deepCopyIdol(),
	}
}
