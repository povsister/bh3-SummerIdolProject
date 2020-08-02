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
		idol: s.deepCopyIdol(), current: s.current,
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
			log.Print("%s 的攻击上升了 10 点!", s.Name)
			log.Print("%s 的防御下降了 5 点!", s.Name)
			s.Attack += 10
			s.Defence -= 5
			s.Defence = notLessZero(s.Defence)
		}
		s.roundAttack(round, s.current)
	case BlackSeele:
		if s.Rival.CanIUseSkill(round, "我换我自己!") {
			s.current = WhiteSeele
			log.Print("%s 转换为白色形态!", s.Name)
			rndHeal := s.RandNum(15)
			s.Health += rndHeal
			s.Health = notGreater100(s.Health)
			log.Print("%s 的生命值上升了 %d 点", s.Name, rndHeal)
			log.Print("%s 的攻击下降了 10 点!", s.Name)
			log.Print("%s 的防御上升了 5 点!", s.Name)
			s.Attack -= 10
			s.Defence += 5
			s.Attack = notLessZero(s.Attack)
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
		log.Print("黑色%s 普攻 造成 %d 点伤害", s.Name, s.Rival.Attributes().trueDamage(s.Attack))
		s.Rival.TakeDamage(round, s.Attack, 1, Normal)
	default:
		panic("unknown seeleType")
	}
}

func (s *SeeleVollerei) Reset() {
	s.idol.Reset()
	s.current = WhiteSeele
}
