package player

type KianaKaslana struct {
	idol
}

func (k *KianaKaslana) DeepCopy() Player {
	return &KianaKaslana{
		idol: k.deepCopyIdol(),
	}
}

func (k *KianaKaslana) RoundAttack(round uint16) {
	if k.tryRecover() {
		return
	}
	if round%2 == 0 {
		k.Rival.TakeDamage(round, Unique, k.Attack+2*k.Rival.Attributes().Defence)
		if k.Rand(35) {
			k.stunned = true
		}
	} else {
		k.Rival.TakeDamage(round, Normal, k.Attack)
	}
}
