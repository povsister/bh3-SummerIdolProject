package player

type Player interface {
	RoundAttack(uint16)
	TakeDamage(round uint16, damage int16, times uint8, form AttackType)
	DirectTakeDamage(round uint16, damage int16, times uint8, form AttackType)
	DeepCopy() Player
	IdolName() string
	Attributes() *idol
	AffectHealth(round uint16, num int16, form AttackType)
	AffectAttack(round uint16, num int16, form AttackType)
	AffectDefence(round uint16, num int16, form AttackType)
	AffectAccuracy(round uint16, num int16, form AttackType)
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
	ID       Candidate
	Name     string
	Health   int16
	Attack   int16
	Defence  int16
	Speed    int16
	Accuracy int16 // 0 - 100  default 100
	Rival    Player
	idolStatus
}

func (i *idol) AffectHealth(round uint16, num int16, form AttackType) {
	i.Health += num
}

func (i *idol) AffectAttack(round uint16, num int16, form AttackType) {
	i.Attack += num
	if i.Attack < 0 {
		i.Attack = 0
	}
}

func (i *idol) AffectDefence(round uint16, num int16, form AttackType) {
	i.Defence += num
	if i.Defence < 0 {
		i.Defence = 0
	}
}

func (i *idol) AffectAccuracy(round uint16, num int16, form AttackType) {
	i.Accuracy += num
	if i.Accuracy < 0 {
		i.Accuracy = 0
	} else if i.Accuracy > 100 {
		i.Accuracy = 100
	}
}

type AttackType uint8

const (
	Normal AttackType = iota
	Unique
)

func (i *idol) RoundAttack(round uint16) {
	panic(`not implemented`)
}

func (i *idol) TakeDamage(round uint16, damage int16, times uint8, form AttackType) {
	i.takeDamage(damage, times)
	//fmt.Printf("%s当前 %d HP\n", i.Name, i.Health)
}

// be noted that this also calculate the impact of defence
func (i *idol) takeDamage(damage int16, times uint8) {
	for k := 0; uint8(k) < times; k++ {
		i.Health -= i.trueDamage(damage)
	}
}

func (i *idol) DirectTakeDamage(round uint16, damage int16, times uint8, form AttackType) {
	i.directTakeDamage(damage, times)
	//fmt.Printf("%s当前 %d HP\n", i.Name, i.Health)
}

func (i *idol) directTakeDamage(damage int16, times uint8) {
	for k := 0; uint8(k) < times; k++ {
		i.Health -= damage
	}
}

// return true if rand value <= thresh
// value of rand num [0:99)
func (i *idol) Rand(thresh int16) bool {
	return (<-Rand % 100) <= int(thresh-1)
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
		i.Speed, 100, i.Rival, idolStatus{false, false, false},
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
	i.Accuracy = Players[i.ID].Attributes().Accuracy
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
		idol{Kiana, `琪亚娜`, 100, 24, 11, 23, 100, nil, defaultIdolStatus},
	},
	Mei: &RaidenMei{
		idol{Mei, `芽衣`, 100, 22, 12, 30, 100, nil, defaultIdolStatus},
	},
	Bronya: &BronyaZaychik{
		idol{Bronya, `布洛妮娅`, 100, 21, 10, 20, 100, nil, defaultIdolStatus},
	},
	Himeko: &MurataHimeko{
		idol{Himeko, `姬子`, 100, 23, 9, 12, 100, nil, defaultIdolStatus},
	},
	Rita: &RitaRossweisse{
		idol{Rita, `丽塔`, 100, 26, 11, 17, 100, nil, defaultIdolStatus},
	},
	Sakura: &YaeSakura{
		idol{Sakura, `八重樱&卡莲`, 100, 20, 9, 18, 100, nil, defaultIdolStatus},
	},
	Raven: &TheRaven{
		idol{Raven, `渡鸦`, 100, 23, 14, 14, 100, nil, defaultIdolStatus},
	},
	Theresa: &TheresaApocalypse{
		idol{Theresa, `德丽莎`, 100, 19, 12, 22, 100, nil, defaultIdolStatus},
	},
	Twins: &TheTwins{
		idol{Twins, `罗莎莉亚&莉莉娅`, 100, 18, 10, 10, 100, nil, defaultIdolStatus}, false, false,
	},
	Seele: &SeeleVollerei{
		idol{Seele, `希儿`, 100, 23, 13, 26, 100, nil, defaultIdolStatus}, WhiteSeele,
	},
	Durandal: &BiankaAtaegina{
		idol{Durandal, `幽兰黛尔`, 100, 19, 10, 15, 100, nil, defaultIdolStatus},
	},
	Fuka: &FuHua{
		idol{Fuka, `符华`, 100, 17, 15, 16, 100, nil, defaultIdolStatus},
	},
}
