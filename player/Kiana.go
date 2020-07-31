package player

type KianaKaslana struct {
	idol
	stunned bool
}

func (k *KianaKaslana) DeepCopy() Player {
	return &KianaKaslana{
		idol: k.deepCopyIdol(),
	}
}

func (k *KianaKaslana) RoundAttack(round uint16) {
	if k.stunned {
		k.stunned = false
		return
	}
	if round%2 == 0 {
		k.Rival.TakeDamage(k.Attack+2*k.Rival.Attributes().Defence, Unique)
		if k.Rand(35) {
			k.stunned = true
		}
	} else {
		k.Rival.TakeDamage(k.Attack, Normal)
	}
}

func (k *KianaKaslana) Reset() {
	k.idol.Reset()
	k.stunned = false
}
