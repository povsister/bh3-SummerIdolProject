package player

import "math/rand"

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

func (s *SeeleVollerei) RoundAttack(defender Player, round uint16) {
	if s.current == WhiteSeele {
		s.current = BlackSeele
		defender.TakeDamage(s.Attack + 10)
	} else {
		s.Health += int16(rand.Intn(15) + 1)
		s.current = WhiteSeele
		defender.TakeDamage(s.Attack)
	}
}

func (s *SeeleVollerei) TakeDamage(damage int16) {
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
