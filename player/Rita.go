package player

type RitaRossweisse struct {
	idol
}

func (r *RitaRossweisse) DeepCopy() Player {
	return &RitaRossweisse{
		idol: r.deepCopyIdol(),
	}
}
