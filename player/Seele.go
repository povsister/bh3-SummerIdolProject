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
		s.Rival.TakeDamage(round, s.Attack+10, 1, Normal)
	} else {
		s.Health += s.RandNum(15)
		s.current = WhiteSeele
		s.Rival.TakeDamage(round, s.Attack, 1, Normal)
	}
}

func (s *SeeleVollerei) TakeDamage(round uint16, damage int16, times uint8, form AttackType) {
	var trueDamage int16
	if s.current == WhiteSeele {
		trueDamage = damage - (s.Defence)
	} else {
		trueDamage = damage - (s.Defence - 5)
	}
	if trueDamage >= 0 {
		for k := 0; uint8(k) < times; k++ {
			s.Health -= trueDamage
		}
	}
}

func (s *SeeleVollerei) Reset() {
	s.idol.Reset()
	s.current = WhiteSeele
}
