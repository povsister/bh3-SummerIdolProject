package player

type BiankaAtaegina struct {
	idol
}

func (d *BiankaAtaegina) DeepCopy() Player {
	return &BiankaAtaegina{
		idol: d.deepCopyIdol(),
	}
}

func (d *BiankaAtaegina) RoundAttack(defender Player, round uint16) {
	defender.TakeDamage(d.Attack+int16(3*round), Normal)
}

func (d *BiankaAtaegina) TakeDamage(damage int16, from AttackType) {
	if from == Unique && d.Rand(16) {
		d.Rival.TakeDamage(30, Unique)
		return
	}
	trueDamage := damage - d.Defence
	if trueDamage >= 0 {
		d.Health -= trueDamage
	}
}

func (d *BiankaAtaegina) DirectTakeDamage(damage int16, from AttackType) {
	if from == Unique && d.Rand(16) {
		d.Rival.TakeDamage(30, Unique)
		return
	}
	d.Health -= damage
}
