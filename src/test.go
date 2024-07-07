type Portion int

const (
	Regular Portion = iota // 普通
	Small                  // 小盛
	Large                  // 大盛り
)

type Udon struct {
	men      Portion
	aburaage bool
	ebiten   unit
}

func NewUdon(p Portion, aburaage bool, ebiten unit) *Udon {
	return &Udon{
		men:      p,
		aburaage: aburaage,
		ebiten:   ebiten,
	}
}

var tempuraUdon = NewUdon(Large, false, 2)