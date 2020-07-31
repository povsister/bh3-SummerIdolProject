package player

type Player interface {
	RoundAttack(Player, uint16)
	TakeDamage(int16, AttackType)
	DirectTakeDamage(int16, AttackType)
	DeepCopy() Player
	IdolName() string
	Attributes() *idol
	IsDead() bool
	Reset()
}

type idol struct {
	ID      Candidate
	Name    string
	Health  int16
	Attack  int16
	Defence int16
	Speed   int16
	Rival   Player
}

type AttackType uint8

const (
	Normal AttackType = iota
	Unique
)

func (i *idol) RoundAttack(player Player, round uint16) {
	panic(`not implemented`)
}

func (i *idol) TakeDamage(damage int16, form AttackType) {
	i.Health -= i.trueDamage(damage)
}

func (i *idol) DirectTakeDamage(damage int16, form AttackType) {
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
		i.ID, i.Name, i.Health, i.Attack, i.Defence, i.Speed, i.Rival,
	}
}

func (i *idol) trueDamage(damage int16) (ret int16) {
	if ret = damage - i.Defence; ret >= 0 {
		return
	}
	return 0
}

func (i *idol) Attributes() *idol {
	return i
}

func (i *idol) IsDead() bool {
	return i.Health <= 0
}

func (i *idol) Reset() {
	i.Health = Players[i.ID].Attributes().Health
	i.Attack = Players[i.ID].Attributes().Attack
	i.Defence = Players[i.ID].Attributes().Defence
	i.Speed = Players[i.ID].Attributes().Speed
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
		idol{Kiana, `琪亚娜`, 100, 24, 11, 23, nil}, false,
	},
	Mei: &RaidenMei{
		idol{Mei, `芽衣`, 100, 22, 12, 30, nil},
	},
	Bronya: &BronyaZaychik{
		idol{Bronya, `布洛妮娅`, 100, 21, 10, 20, nil},
	},
	Himeko: &MurataHimeko{
		idol{Himeko, `姬子`, 100, 23, 9, 12, nil},
	},
	Rita: &RitaRossweisse{
		idol{Rita, `丽塔`, 100, 26, 11, 17, nil},
	},
	Sakura: &YaeSakura{
		idol{Sakura, `八重樱&卡莲`, 100, 20, 9, 18, nil},
	},
	Raven: &TheRaven{
		idol{Raven, `渡鸦`, 100, 23, 14, 14, nil},
	},
	Theresa: &TheresaApocalypse{
		idol{Theresa, `德丽莎`, 100, 19, 12, 22, nil},
	},
	Twins: &TheTwins{
		idol{Twins, `罗莎莉亚&莉莉娅`, 100, 18, 10, 10, nil}, false,
	},
	Seele: &SeeleVollerei{
		idol{Seele, `希儿`, 100, 23, 13, 26, nil}, WhiteSeele,
	},
	Durandal: &BiankaAtaegina{
		idol{Durandal, `幽兰黛尔`, 100, 19, 10, 15, nil},
	},
	Fuka: &FuHua{
		idol{Fuka, `符华`, 100, 17, 15, 16, nil},
	},
}
