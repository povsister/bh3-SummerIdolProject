package player

import (
	"math"
	"povsister.app/bh3/summer-idol/log"
)

type Player interface {
	PreBattle(round uint16)
	RoundAttack(round uint16)
	TakeDamage(round uint16, damage int16, times uint8, form AttackType)
	DirectTakeDamage(round uint16, damage int16, times uint8, form AttackType)
	DeepCopy() Player
	IdolName() string
	Attributes() *idol
	AffectAttr(aT attrType, round uint16, num int16, form AttackType)
	CanIUseSkill(round uint16, skillName string) bool
	NormalDamageType() DamageType
	ResetRound()
	IsDead() bool
	Reset()
}

type DamageType uint8

const (
	Physical DamageType = iota
	Elemental
)

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
	hit      bool // default true, indicates if player get hit by rival
	idolStatus
}

func (i *idol) PreBattle(round uint16) {
	// default nothing
}

func roundDamage(dam float64) int16 {
	return int16(math.Round(dam))
}

func (i *idol) NormalDamageType() DamageType {
	return Physical
}

func (i *idol) ResetRound() {
	i.hit = true
}

// tells the Rival if they can use skill on current round
// must be called by rival
func (i *idol) CanIUseSkill(round uint16, skillName string) bool {
	return true
}

type attrType uint8

const (
	attrHealth attrType = iota
	attrAttack
	attrDefence
	attrAccuracy
)

func notLessZero(num int16) int16 {
	if num < 0 {
		return 0
	}
	return num
}

func notGreater100(num int16) int16 {
	if num > 100 {
		return 100
	}
	return num
}

// affect attribute. This method must called by the rival
func (i *idol) AffectAttr(aT attrType, round uint16, num int16, form AttackType) {
	switch aT {
	case attrHealth:
		if i.hit {
			i.affectHealth(round, num, form)
		}
	case attrAttack:
		if i.hit {
			i.affectAttack(round, num, form)
		}
	case attrDefence:
		if i.hit {
			i.affectDefence(round, num, form)
		}
	case attrAccuracy:
		if i.hit {
			i.affectAccuracy(round, num, form)
		}
	default:
		panic(`unknown attrType`)
	}
}

func (i *idol) affectHealth(round uint16, num int16, form AttackType) {
	i.Health += num
	i.Health = notGreater100(i.Health)
	log.AttributeStatus(i.Name, `生命值`, num)
}

func (i *idol) affectAttack(round uint16, num int16, form AttackType) {
	i.Attack += num
	i.Attack = notLessZero(i.Attack)
	log.AttributeStatus(i.Name, `攻击`, num)
}

func (i *idol) affectDefence(round uint16, num int16, form AttackType) {
	i.Defence += num
	//i.Defence = notLessZero(i.Defence)
	log.AttributeStatus(i.Name, `防御`, num)
}

func (i *idol) affectAccuracy(round uint16, num int16, form AttackType) {
	i.Accuracy += num
	i.Accuracy = notLessZero(i.Accuracy)
	i.Accuracy = notGreater100(i.Accuracy)
	log.AttributeStatus(i.Name, `命中率`, num)
}

type AttackType uint8

const (
	Normal AttackType = iota // normal attack
	Skill                    // typical skill
	Unique                   // unique skill
)

func (i *idol) RoundAttack(round uint16) {
	panic(`not implemented`)
}

func (i *idol) TakeDamage(round uint16, damage int16, times uint8, form AttackType) {
	i.takeDamage(damage, times)
	log.HPStatus(i.Name, i.Health)
}

// be noted that this also calculate the impact of defence
func (i *idol) takeDamage(damage int16, times uint8) {
	for k := 0; uint8(k) < times; k++ {
		i.Health -= i.trueDamage(damage)
	}
}

func (i *idol) DirectTakeDamage(round uint16, damage int16, times uint8, form AttackType) {
	i.directTakeDamage(damage, times)
	log.HPStatus(i.Name, i.Health)
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
		i.Speed, i.Accuracy, i.Rival, i.hit,
		idolStatus{i.stunned, i.paralyzed, i.frozen},
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
	_ Candidate = iota
	Kiana
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
		idol{Kiana, `琪亚娜`, 100, 24, 11, 23, 100, nil, true, defaultIdolStatus},
	},
	Mei: &RaidenMei{
		idol{Mei, `芽衣`, 100, 22, 12, 30, 100, nil, true, defaultIdolStatus},
	},
	Bronya: &BronyaZaychik{
		idol{Bronya, `布洛妮娅`, 100, 21, 10, 20, 100, nil, true, defaultIdolStatus},
	},
	Himeko: &MurataHimeko{
		idol{Himeko, `姬子`, 100, 23, 9, 12, 100, nil, true, defaultIdolStatus},
	},
	Rita: &RitaRossweisse{
		idol{Rita, `丽塔`, 100, 26, 11, 17, 100, nil, true, defaultIdolStatus}, false,
	},
	Sakura: &YaeSakura{
		idol{Sakura, `八重樱&卡莲`, 100, 20, 9, 18, 100, nil, true, defaultIdolStatus},
	},
	Raven: &TheRaven{
		idol{Raven, `渡鸦`, 100, 23, 14, 14, 100, nil, true, defaultIdolStatus}, false,
	},
	Theresa: &TheresaApocalypse{
		idol{Theresa, `德丽莎`, 100, 19, 12, 22, 100, nil, true, defaultIdolStatus},
	},
	Twins: &TheTwins{
		idol{Twins, `罗莎莉亚&莉莉娅`, 100, 18, 10, 10, 100, nil, true, defaultIdolStatus}, false, false,
	},
	Seele: &SeeleVollerei{
		idol{Seele, `希儿`, 100, 23, 13, 26, 100, nil, true, defaultIdolStatus}, WhiteSeele,
	},
	Durandal: &BiankaAtaegina{
		idol{Durandal, `幽兰黛尔`, 100, 19, 10, 15, 100, nil, true, defaultIdolStatus}, nil, false,
	},
	Fuka: &FuHua{
		idol{Fuka, `符华`, 100, 17, 15, 16, 100, nil, true, defaultIdolStatus},
	},
}
