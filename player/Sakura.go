package player

type YaeSakura struct {
	idol
}

func (s *YaeSakura) DeepCopy() Player {
	return &YaeSakura{
		idol: s.deepCopyIdol(),
	}
}

func (s *YaeSakura) RoundAttack(round uint16) {
	if s.Rand(30) {
		s.Health += 25
	}
	if round%2 == 0 {
		s.Rival.DirectTakeDamage(25, Unique)
	} else {
		s.Rival.TakeDamage(s.Attack, Normal)
	}
}
