package player

import "povsister.app/bh3/summer-idol/log"

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
		log.Print("%s 转换为黑色形态!", s.Name)
		log.Print("黑色%s 普攻 造成 %d 点伤害", s.Name, s.Rival.Attributes().trueDamage(s.Attack+10))
		s.Rival.TakeDamage(round, s.Attack+10, 1, Normal)
	} else {
		s.current = WhiteSeele
		log.Print("%s 转换为白色形态!", s.Name)
		s.AffectHealth(round, s.RandNum(15), Normal)
		log.Print("白色%s 普攻 造成 %d 点伤害", s.Name, s.Rival.Attributes().trueDamage(s.Attack))
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
	log.HPStatus(s.Name, s.Health)
}

func (s *SeeleVollerei) Reset() {
	s.idol.Reset()
	s.current = WhiteSeele
}
