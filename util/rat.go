package util

import (
	"fmt"
	"strings"

	"github.com/berquerant/crd/errorx"
	"gopkg.in/yaml.v3"
)

type Rat struct {
	Num   uint
	Denom uint
}

func NewRat(num, denom uint) Rat {
	return Rat{
		Num:   num,
		Denom: denom,
	}
}

func (r Rat) Float() float64 {
	return float64(r.Num) / float64(r.Denom)
}

func (r Rat) String() string {
	if r.Denom == 1 {
		return fmt.Sprint(r.Num)
	}
	return fmt.Sprintf("%d/%d", r.Num, r.Denom)
}

func (r Rat) MarshalYAML() (any, error) {
	return r.String(), nil
}

func (r *Rat) UnmarshalYAML(value *yaml.Node) error {
	x, err := ParseRat(value.Value)
	if err != nil {
		return err
	}
	*r = x
	return nil
}

func ParseRat(s string) (Rat, error) {
	xs := strings.SplitN(s, "/", 2)

	var d Rat
	if len(xs) == 1 {
		num, err := ParseUint(xs[0])
		if err != nil {
			return d, err
		}
		return NewRat(num, 1), nil
	}

	if len(xs) != 2 {
		return d, errorx.Unmarshal("Rat %s", s)
	}

	num, err := ParseUint(xs[0])
	if err != nil {
		return d, err
	}
	denom, err := ParseUint(xs[1])
	if err != nil {
		return d, err
	}
	return NewRat(num, denom), nil
}
