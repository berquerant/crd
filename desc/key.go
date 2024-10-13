package desc

import "github.com/berquerant/crd/op"

type Key struct{}

func NewKey() *Key {
	return &Key{}
}

type KeyInfo struct {
	Scale    *op.Scale       `yaml:"scale"`
	Diatonic KeyInfoDiatonic `yaml:"diatonic"`
}

type KeyInfoDiatonic struct {
	Triads   [7]op.DiatonicChord `yaml:"triads"`
	Sevenths [7]op.DiatonicChord `yaml:"sevenths"`
}

func (k Key) Describe(scale *op.Scale) KeyInfo {
	c := op.NewDiatonicChorder(scale)
	return KeyInfo{
		Scale: scale,
		Diatonic: KeyInfoDiatonic{
			Triads:   c.Triads(),
			Sevenths: c.Sevenths(),
		},
	}
}
