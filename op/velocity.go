package op

import (
	"errors"

	"github.com/berquerant/crd/util"
	"gopkg.in/yaml.v3"
)

type DynamicSign int

const (
	UnknownDynamicSign DynamicSign = iota
	Pianissimo
	Piano
	MezzoPiano
	MezzoForte
	Forte
	Fortissimo
)

const (
	dynamicSignPianissimo = "pp"
	dynamicSignPiano      = "p"
	dynamicSignMezzoPiano = "mp"
	dynamicSignMezzoForte = "mf"
	dynamicSignForte      = "f"
	dynamicSignFortissimo = "ff"
)

var (
	stringDynamicSignMap = map[string]DynamicSign{
		dynamicSignPianissimo: Pianissimo,
		dynamicSignPiano:      Piano,
		dynamicSignMezzoPiano: MezzoPiano,
		dynamicSignMezzoForte: MezzoForte,
		dynamicSignForte:      Forte,
		dynamicSignFortissimo: Fortissimo,
	}
	dynamicSignStringMap   = util.MustInverseMap(stringDynamicSignMap)
	dynamicSignVelocityMap = map[DynamicSign]Velocity{
		Pianissimo: 22,
		Piano:      43,
		MezzoPiano: 64,
		MezzoForte: 85,
		Forte:      106,
		Fortissimo: 127,
	}
)

func GetDynamicSignStrings() []string {
	ss := []string{}
	for k := range stringDynamicSignMap {
		ss = append(ss, k)
	}
	return ss
}

func NewDynamicSign(s string) DynamicSign {
	if x, ok := stringDynamicSignMap[s]; ok {
		return x
	}
	return UnknownDynamicSign
}

func (d DynamicSign) Velocity() Velocity {
	return dynamicSignVelocityMap[d]
}

func (d DynamicSign) String() string {
	return dynamicSignStringMap[d]
}

func (d DynamicSign) MarshalYAML() (any, error) {
	return d.String(), nil
}

func (d *DynamicSign) UnmarshalYAML(value *yaml.Node) error {
	x := NewDynamicSign(value.Value)
	if x == UnknownDynamicSign {
		return ErrUnknownDynamicSign
	}
	*d = x
	return nil
}

var (
	ErrUnknownDynamicSign = errors.New("UnknownDynamicSign")
)

type Velocity uint8
