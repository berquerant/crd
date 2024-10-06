package op

import (
	"github.com/berquerant/crd/errorx"
	"github.com/berquerant/crd/logx"
	"github.com/berquerant/crd/util"
	"gopkg.in/yaml.v3"
)

type Meter struct {
	util.Rat
}

func NewMeter(num, denom uint) (Meter, error) {
	m := Meter{
		util.NewRat(num, denom),
	}
	return m, m.validate()
}

func MustNewMeter(num, denom uint) Meter {
	m, err := NewMeter(num, denom)
	logx.PanicOnError(err)
	return m
}

func (m *Meter) UnmarshalYAML(value *yaml.Node) error {
	var r util.Rat
	if err := value.Decode(&r); err != nil {
		return err
	}
	*m = Meter{r}
	return m.validate()
}

func (m Meter) validate() error {
	if m.Denom < 1 {
		return errorx.Invalid("Meter should have positive denominator")
	}
	if m.Num < 1 {
		return errorx.Invalid("Meter should be positive")
	}
	return nil
}
