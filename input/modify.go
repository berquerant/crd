package input

import (
	"fmt"
	"log/slog"

	"github.com/berquerant/crd/op"
)

type Modifier interface {
	Modify(v *Instance) error
}

var (
	_ Modifier = &ChordMetaTextMotifier{}
)

func NewChordMetaTextMofidier(symbolSep, slashSep string) *ChordMetaTextMotifier {
	return &ChordMetaTextMotifier{
		symbolSep: symbolSep,
		slashSep:  slashSep,
	}
}

type ChordMetaTextMotifier struct {
	symbolSep string
	slashSep  string
}

func (m ChordMetaTextMotifier) Modify(v *Instance) error {
	c := v.Chord
	if c == nil {
		return nil
	}
	if v.Meta == nil {
		v.Meta = op.NewMeta()
	}
	v.Meta.Set(MetaTextKey, m.generateText(c))
	return nil
}

func (m ChordMetaTextMotifier) generateText(c *Chord) string {
	slog.Debug("ChordMetaTextModifier", slog.Any("chord", c))
	s := fmt.Sprintf("%d%s%s%s",
		c.Degree.Value, c.Degree.Name.Coerce(),
		m.symbolSep,
		c.Chord)
	if x := c.Base; x != nil {
		return fmt.Sprintf("%s%s%s", s, m.slashSep, x)
	}
	return s
}
