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
	if s.current == WhiteSeele {
		s.current = BlackSeele
		s.Rival.TakeDamage(s.Attack+10, Normal)
	} else {
		s.Health += s.RandNum(15)
		s.current = WhiteSeele
		s.Rival.TakeDamage(s.Attack, Normal)
	}
}

func (s *SeeleVollerei) TakeDamage(damage int16, form AttackType) {
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
