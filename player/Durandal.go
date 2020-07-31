package player

type BiankaAtaegina struct {
	idol
}

func (d *BiankaAtaegina) DeepCopy() Player {
	return &BiankaAtaegina{
		idol: d.deepCopyIdol(),
	}
}

func (d *BiankaAtaegina) RoundAttack(round uint16) {
	if d.tryRecover() {
		return
	}
	d.Rival.TakeDamage(round, d.Attack+int16(3*round), 1, Normal)
}

func (d *BiankaAtaegina) TakeDamage(round uint16, damage int16, times uint8, form AttackType) {
	if form == Unique && d.Rand(16) {
		d.Rival.TakeDamage(round, 30, 1, Unique)
		return
	}
	// take normal damage or skill not triggered
	d.takeDamage(damage, times)
}

func (d *BiankaAtaegina) DirectTakeDamage(round uint16, damage int16, times uint8, form AttackType) {
	if form == Unique && d.Rand(16) {
		d.Rival.TakeDamage(round, 30, 1, Unique)
		return
	}
	d.directTakeDamage(damage, times)
}
