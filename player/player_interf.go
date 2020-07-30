package player

import (
	"math/rand"
	"time"
)

type Player interface {
	RoundAttack(Player, uint16)
	TakeDamage(int16)
	DirectTakeDamage(int16)
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
}

func (i *idol) RoundAttack(player Player, round uint16) {
	panic(`not implemented`)
}

func (i *idol) TakeDamage(damage int16) {
	trueDamage := damage - i.Defence
	if trueDamage >= 0 {
		i.Health -= trueDamage
	}
}

func (i *idol) DirectTakeDamage(damage int16) {
	i.Health -= damage
}

func (i *idol) Rand(thresh int) bool {
	rand.Seed(time.Now().UnixNano())
	return thresh-1 <= rand.Intn(100)
}

func (i *idol) IdolName() string {
	return i.Name
}

func (i *idol) DeepCopy() Player {
	panic(`not implemented`)
}

func (i *idol) deepCopyIdol() idol {
	return idol{
		i.ID, i.Name, i.Health, i.Attack, i.Defence, i.Speed,
	}
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
		idol{Kiana, `琪亚娜`, 100, 24, 11, 23},
	},
	Mei: &RaidenMei{
		idol{Mei, `芽衣`, 100, 22, 12, 30},
	},
	Bronya: &BronyaZaychik{
		idol{Bronya, `布洛妮娅`, 100, 21, 10, 20},
	},
	Himeko: &MurataHimeko{
		idol{Himeko, `姬子`, 100, 23, 9, 12},
	},
	Rita: &RitaRossweisse{
		idol{Rita, `丽塔`, 100, 26, 11, 17},
	},
	Sakura: &YaeSakura{
		idol{Sakura, `八重樱&卡莲`, 100, 20, 9, 18},
	},
	Raven: &TheRaven{
		idol{Raven, `渡鸦`, 100, 23, 14, 14},
	},
	Theresa: &TheresaApocalypse{
		idol{Theresa, `德丽莎`, 100, 19, 12, 22},
	},
	Twins: &TheTwins{
		idol{Twins, `罗莎莉亚&莉莉娅`, 100, 18, 10, 10},
	},
	Seele: &SeeleVollerei{
		idol{Seele, `希儿`, 100, 23, 13, 26},
	},
	Durandal: &BiankaAtaegina{
		idol{Durandal, `幽兰黛尔`, 100, 19, 10, 15},
	},
	Fuka: &FuHua{
		idol{Fuka, `符华`, 100, 17, 15, 16},
	},
}
