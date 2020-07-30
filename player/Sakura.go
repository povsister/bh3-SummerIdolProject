package player

type YaeSakura struct {
	idol
}

func (s *YaeSakura) DeepCopy() Player {
	return &YaeSakura{
		idol: s.deepCopyIdol(),
	}
}

func (s *YaeSakura) RoundAttack(defender Player, round uint16) {
	if s.Rand(35) {
		s.Health += 30
	}
	if round%2 == 0 {
		defender.DirectTakeDamage(25)
	} else {
		defender.TakeDamage(s.Attack)
	}
}
