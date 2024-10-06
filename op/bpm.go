package op

import (
	"github.com/berquerant/crd/errorx"
	"github.com/berquerant/crd/util"
	"gopkg.in/yaml.v3"
)

type BPM uint

func NewBPM(v uint) (BPM, error) {
	x := BPM(v)
	return x, x.validate()
}

func (b BPM) MarshalYAML() (any, error) {
	return uint(b), nil
}

func (b *BPM) UnmarshalYAML(value *yaml.Node) error {
	u, err := util.ParseUint(value.Value)
	if err != nil {
		return err
	}
	*b = BPM(u)
	return nil
}

func (b BPM) validate() error {
	if b == 0 {
		return errorx.Invalid("BPM should be positive")
	}
	return nil
}
