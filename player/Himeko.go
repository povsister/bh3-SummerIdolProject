package player

type MurataHimeko struct {
	idol
}

func (h *MurataHimeko) DeepCopy() Player {
	return &MurataHimeko{
		idol: h.deepCopyIdol(),
	}
}

func (h *MurataHimeko) RoundAttack(round uint16) {
	if h.tryRecover() {
		return
	}
	// lower 35% accuracy every 2 round
	accuracy := 100 - int(round/2)*35
	if accuracy <= 0 {
		return
	}
	if round%2 == 0 {
		if h.Rand(accuracy) {
			h.Rival.DirectTakeDamage(round, h.getRealDamage(2*h.Attack), Unique)
		}
		return
	}
	if h.Rand(accuracy) {
		h.Rival.DirectTakeDamage(round, h.getRealDamage(h.Attack), Normal)
	}
}

func (h *MurataHimeko) getRealDamage(attack int16) int16 {
	switch h.Rival.Attributes().ID {
	case Durandal, Sakura, Twins:
		return (attack - h.Rival.Attributes().Defence) * 2
	}
	return attack - h.Rival.Attributes().Defence
}
