package player

type SeeleType int

const (
	WhiteSeele = iota
	BlackSeele
)

type SeeleVollerei struct {
	idol
	current SeeleType
}

func (s *SeeleVollerei) DeepCopy() Player {
	return &SeeleVollerei{
		idol: s.deepCopyIdol(), current: WhiteSeele,
	}
}

func (s *SeeleVollerei) RoundAttack(round uint16) {
	if s.tryRecover() {
		return
	}
	if s.current == WhiteSeele {
		s.current = BlackSeele
		s.Rival.TakeDamage(round, Normal, s.Attack+10)
	} else {
		s.Health += s.RandNum(15)
		s.current = WhiteSeele
		s.Rival.TakeDamage(round, Normal, s.Attack)
	}
}

func (s *SeeleVollerei) TakeDamage(round uint16, form AttackType, damage int16) {
	var trueDamage int16
	if s.current == WhiteSeele {
		trueDamage = damage - (s.Defence)
	} else {
		trueDamage = damage - (s.Defence - 5)
	}
	if trueDamage >= 0 {
		s.Health -= trueDamage
	}
}

func (s *SeeleVollerei) Reset() {
	s.idol.Reset()
	s.current = WhiteSeele
}
