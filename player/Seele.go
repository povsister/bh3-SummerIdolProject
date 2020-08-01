package player

import "povsister.app/bh3/summer-idol/log"

type seeleType int

const (
	WhiteSeele = iota
	BlackSeele
)

type SeeleVollerei struct {
	idol
	current seeleType
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
	switch s.current {
	case WhiteSeele:
		if s.Rival.CanIUseSkill(round, "我换我自己!") {
			s.current = BlackSeele
			log.Print("%s 转换为黑色形态!", s.Name)
		}
		s.roundAttack(round, s.current)
	case BlackSeele:
		if s.Rival.CanIUseSkill(round, "我换我自己!") {
			s.current = WhiteSeele
			log.Print("%s 转换为白色形态!", s.Name)
			rndHeal := s.RandNum(15)
			s.Health += rndHeal
			log.Print("%s 的生命值上升了 %d 点", s.Name, rndHeal)
		}
		s.roundAttack(round, s.current)
	default:
		panic("unknown seeleType")
	}
}

func (s *SeeleVollerei) roundAttack(round uint16, sT seeleType) {
	switch sT {
	case WhiteSeele:
		log.Print("白色%s 普攻 造成 %d 点伤害", s.Name, s.Rival.Attributes().trueDamage(s.Attack))
		s.Rival.TakeDamage(round, s.Attack, 1, Normal)
	case BlackSeele:
		log.Print("黑色%s 普攻 造成 %d 点伤害", s.Name, s.Rival.Attributes().trueDamage(s.Attack+10))
		s.Rival.TakeDamage(round, s.Attack+10, 1, Normal)
	default:
		panic("unknown seeleType")
	}
}

func (s *SeeleVollerei) TakeDamage(round uint16, damage int16, times uint8, form AttackType) {
	var trueDamage int16
	switch s.current {
	case WhiteSeele:
		trueDamage = damage - (s.Defence)
	case BlackSeele:
		trueDamage = damage - (s.Defence - 5)
	default:
		panic("unknown seeleType")
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
