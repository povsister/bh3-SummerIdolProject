package player

type TheRaven struct {
	idol
}

func (r *TheRaven) DeepCopy() Player {
	return &TheRaven{
		idol: r.deepCopyIdol(),
	}
}
