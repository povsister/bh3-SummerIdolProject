package player

type Player interface {
	RoundAttack(uint16)
	TakeDamage(uint16, AttackType, int16)
	DirectTakeDamage(uint16, int16, AttackType)
	DeepCopy() Player
	IdolName() string
	Attributes() *idol
	IsDead() bool
	Reset()
}

type idolStatus struct {
	stunned   bool
	paralyzed bool
	frozen    bool
}

var defaultIdolStatus = idolStatus{
	stunned:   false,
	paralyzed: false,
	frozen:    false,
}

type idol struct {
	ID      Candidate
	Name    string
	Health  int16
	Attack  int16
	Defence int16
	Speed   int16
	Rival   Player
	idolStatus
}

type AttackType uint8

const (
	Normal AttackType = iota
	Unique
)

func (i *idol) RoundAttack(round uint16) {
	panic(`not implemented`)
}

func (i *idol) TakeDamage(round uint16, form AttackType, damage int16) {
	i.Health -= i.trueDamage(damage)
}

func (i *idol) DirectTakeDamage(round uint16, damage int16, form AttackType) {
	i.Health -= damage
}

// return true if rand value <= thresh
// value of rand num [0:99)
func (i *idol) Rand(thresh int) bool {
	return (<-Rand % 100) <= (thresh - 1)
}

// return a random number from [1:upper]
func (i *idol) RandNum(upper int) int16 {
	return int16(<-Rand%upper + 1)
}

func (i *idol) IdolName() string {
	return i.Name
}

func (i *idol) DeepCopy() Player {
	panic(`not implemented`)
}

func (i *idol) deepCopyIdol() idol {
	return idol{
		i.ID, i.Name, i.Health, i.Attack, i.Defence,
		i.Speed, i.Rival, idolStatus{false, false, false},
	}
}

func (i *idol) trueDamage(damage int16) (ret int16) {
	if ret = damage - i.Defence; ret >= 0 {
		return
	}
	return 0
}

// return true if recovered from frozen/stunned/paralyzed
func (i *idol) tryRecover() bool {
	if i.paralyzed || i.frozen || i.stunned {
		i.resetStatus()
		return true
	}
	return false
}

func (i *idol) Attributes() *idol {
	return i
}

func (i *idol) IsDead() bool {
	return i.Health <= 0
}

func (i *idol) resetStatus() {
	i.idolStatus.stunned = false
	i.idolStatus.frozen = false
	i.idolStatus.paralyzed = false
}

func (i *idol) Reset() {
	i.Health = Players[i.ID].Attributes().Health
	i.Attack = Players[i.ID].Attributes().Attack
	i.Defence = Players[i.ID].Attributes().Defence
	i.Speed = Players[i.ID].Attributes().Speed
	i.resetStatus()
}

type Candidate uint8

const (
	Kiana Candidate = iota
	Mei
	Bronya
	Himeko
	Rita
	Sakura
	Raven
	Theresa
	Twins
	Seele
	Durandal
	Fuka
)

var Players = map[Candidate]Player{
	Kiana: &KianaKaslana{
		idol{Kiana, `琪亚娜`, 100, 24, 11, 23, nil, defaultIdolStatus},
	},
	Mei: &RaidenMei{
		idol{Mei, `芽衣`, 100, 22, 12, 30, nil, defaultIdolStatus},
	},
	Bronya: &BronyaZaychik{
		idol{Bronya, `布洛妮娅`, 100, 21, 10, 20, nil, defaultIdolStatus},
	},
	Himeko: &MurataHimeko{
		idol{Himeko, `姬子`, 100, 23, 9, 12, nil, defaultIdolStatus},
	},
	Rita: &RitaRossweisse{
		idol{Rita, `丽塔`, 100, 26, 11, 17, nil, defaultIdolStatus},
	},
	Sakura: &YaeSakura{
		idol{Sakura, `八重樱&卡莲`, 100, 20, 9, 18, nil, defaultIdolStatus},
	},
	Raven: &TheRaven{
		idol{Raven, `渡鸦`, 100, 23, 14, 14, nil, defaultIdolStatus},
	},
	Theresa: &TheresaApocalypse{
		idol{Theresa, `德丽莎`, 100, 19, 12, 22, nil, defaultIdolStatus},
	},
	Twins: &TheTwins{
		idol{Twins, `罗莎莉亚&莉莉娅`, 100, 18, 10, 10, nil, defaultIdolStatus}, false,
	},
	Seele: &SeeleVollerei{
		idol{Seele, `希儿`, 100, 23, 13, 26, nil, defaultIdolStatus}, WhiteSeele,
	},
	Durandal: &BiankaAtaegina{
		idol{Durandal, `幽兰黛尔`, 100, 19, 10, 15, nil, defaultIdolStatus},
	},
	Fuka: &FuHua{
		idol{Fuka, `符华`, 100, 17, 15, 16, nil, defaultIdolStatus},
	},
}
