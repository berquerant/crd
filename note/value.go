package note

import (
	"github.com/berquerant/crd/errorx"
	"github.com/berquerant/crd/logx"
	"github.com/berquerant/crd/util"
	"gopkg.in/yaml.v3"
)

type Value struct {
	util.Rat
}

func MustNewValue(num, denom uint) Value {
	v, err := NewValue(num, denom)
	logx.PanicOnError(err)
	return v
}

func NewValue(num, denom uint) (Value, error) {
	r := util.NewRat(num, denom)
	v := Value{r}
	return v, v.validate()
}

func (v Value) validate() error {
	if v.Denom < 1 {
		return errorx.Invalid("Value should have positive denominator")
	}
	if v.Num < 1 {
		return errorx.Invalid("Value should have positive value")
	}
	return nil
}

func (v Value) MarshalYAML() (any, error) {
	return v.String(), nil
}

func (v *Value) UnmarshalYAML(value *yaml.Node) error {
	var r util.Rat
	if err := value.Decode(&r); err != nil {
		return err
	}

	*v = Value{r}
	return v.validate()
}
