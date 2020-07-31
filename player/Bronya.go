package player

import "povsister.app/bh3/summer-idol/log"

type BronyaZaychik struct {
	idol
}

func (b *BronyaZaychik) DeepCopy() Player {
	return &BronyaZaychik{
		idol: b.deepCopyIdol(),
	}
}

func (b *BronyaZaychik) RoundAttack(round uint16) {
	if b.tryRecover() {
		return
	}
	if round%3 == 0 {
		num := b.RandNum(100)
		log.Print("%s 发动技能 摩托拜客哒! 造成 %d 点元素伤害", b.Name, num)
		b.Rival.DirectTakeDamage(round, num, 1, Unique)
		return
	}
	log.Print("%s 普攻 造成 %d 点伤害", b.Name, b.Rival.Attributes().trueDamage(b.Attack))
	b.Rival.TakeDamage(round, b.Attack, 1, Normal)
	if b.Rand(25) {
		log.Print("%s 钻头攻击 造成 4 x %d 点伤害", b.Name, b.Rival.Attributes().trueDamage(12))
		b.Rival.TakeDamage(round, 12, 4, Normal)
	}
}
