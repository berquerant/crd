package astconv

import (
	"fmt"

	"github.com/berquerant/crd/errorx"
	"github.com/berquerant/crd/input"
	"github.com/berquerant/crd/input/ast"
	"github.com/berquerant/crd/note"
	"github.com/berquerant/crd/op"
	"github.com/berquerant/crd/util"
)

type Converter interface {
	// Convert AST (from text) into Instance (yaml).
	Convert(v ast.ChordOrRest) (*input.Instance, error)
}

func NewSyllableASTConverter(scale *op.Scale) *ASTConverter {
	var (
		vc ValuesConverterImpl
		cc = NewSyllableChordConverter(scale)
		mc MetaConverterImpl
	)
	return &ASTConverter{
		valuesConverter: &vc,
		chordConverter:  cc,
		metaConverter:   &mc,
	}
}

func NewDegreeASTConverter() *ASTConverter {
	var (
		vc ValuesConverterImpl
		cc DegreeChordConverter
		mc MetaConverterImpl
	)
	return &ASTConverter{
		valuesConverter: &vc,
		chordConverter:  &cc,
		metaConverter:   &mc,
	}
}

var (
	_ Converter = &ASTConverter{}
)

type ASTConverter struct {
	valuesConverter ValuesConverter
	chordConverter  ChordConverter
	metaConverter   MetaConverter
}

func (c ASTConverter) Convert(v ast.ChordOrRest) (*input.Instance, error) {
	switch v := v.(type) {
	case *ast.Chord:
		values, err := c.valuesConverter.Convert(v.Values)
		if err != nil {
			return nil, fmt.Errorf("%w: ChordValues", err)
		}
		chod, err := c.chordConverter.Convert(v)
		if err != nil {
			return nil, fmt.Errorf("%w: Chord", err)
		}
		meta := c.metaConverter.Convert(v.Meta)
		return &input.Instance{
			Values: values,
			Chord:  chod,
			Meta:   meta,
		}, nil
	case *ast.Rest:
		values, err := c.valuesConverter.Convert(v.Values)
		if err != nil {
			return nil, fmt.Errorf("%w: Rest", err)
		}
		meta := c.metaConverter.Convert(v.Meta)
		return &input.Instance{
			Values: values,
			Meta:   meta,
		}, nil
	default:
		return nil, errorx.Unexpected("Neither Chord nor Rest")
	}
}

type MetaConverter interface {
	Convert(v *ast.ChordMeta) *op.Meta
}

var (
	_ MetaConverter = &MetaConverterImpl{}
)

type MetaConverterImpl struct{}

func (MetaConverterImpl) Convert(v *ast.ChordMeta) *op.Meta {
	if v == nil || len(v.Data) == 0 {
		return nil
	}
	d := map[string]string{}
	for _, x := range v.Data {
		key := x.Key.Value()
		value := x.Value.Value()
		d[key] = value
	}
	r := op.Meta(d)
	return &r
}

type ValuesConverter interface {
	Convert(v *ast.ChordValues) ([]note.Value, error)
}

var (
	_ ValuesConverter = &ValuesConverterImpl{}
)

type ValuesConverterImpl struct{}

func (c ValuesConverterImpl) Convert(v *ast.ChordValues) ([]note.Value, error) {
	result := make([]note.Value, len(v.Values))
	for i, x := range v.Values {
		y, err := c.convertValue(x)
		if err != nil {
			return nil, fmt.Errorf("%w: ChordValues[%d]", err, i)
		}
		result[i] = y
	}
	return result, nil
}

func (c ValuesConverterImpl) convertValue(v *ast.ChordValue) (note.Value, error) {
	var d note.Value

	num, err := util.ParseUint(v.Num.Value())
	if err != nil {
		return d, fmt.Errorf("%w: ChordValue num %v", err, v.Num)
	}
	var denom uint = 1
	if v.Denom != nil {
		if x := v.Denom.Value(); x != "" {
			y, err := util.ParseUint(x)
			if err != nil {
				return d, fmt.Errorf("%w: ChordValue denom %v", err, v.Denom)
			}
			denom = y
		}
	}

	result, err := note.NewValue(num, denom)
	if err != nil {
		return result, fmt.Errorf("%w: ChordValue %v, %v", err, v.Num, v.Denom)
	}
	return result, nil
}

type ChordConverter interface {
	Convert(v *ast.Chord) (*input.Chord, error)
}

var (
	_ ChordConverter = &DegreeChordConverter{}
	_ ChordConverter = &SyllableChordConverter{}
)

// SyllableChordConverter converts AST contains only syllables.
type SyllableChordConverter struct {
	scale *op.Scale
}

func NewSyllableChordConverter(scale *op.Scale) *SyllableChordConverter {
	return &SyllableChordConverter{
		scale: scale,
	}
}

func (c SyllableChordConverter) Convert(v *ast.Chord) (*input.Chord, error) {
	var result input.Chord

	if err := c.convertChordDegree(&result, v.Degree, v.Base); err != nil {
		return nil, err
	}

	if x := v.Symbol; x != nil {
		result.Chord = x.Symbol.Value()
	}

	return &result, nil
}

func (c SyllableChordConverter) convertChordDegree(result *input.Chord, v *ast.ChordDegree, base *ast.ChordBase) error {
	rootScaleNote, err := c.newScaleNote(v)
	if err != nil {
		return err
	}
	tendency, err := c.getTendency(rootScaleNote)
	if err != nil {
		return err
	}
	degree, err := c.scale.Tonic().GetDegree(rootScaleNote, tendency == op.Sharp)
	if err != nil {
		return err
	}
	result.Degree = degree
	if base == nil {
		return nil
	}

	baseScaleNote, err := c.newScaleNote(base.Degree)
	if err != nil {
		return err
	}
	baseTendency, err := c.getTendency(baseScaleNote)
	if err != nil {
		return err
	}
	baseDegree, err := rootScaleNote.GetDegree(baseScaleNote, baseTendency == op.Sharp)
	if err != nil {
		return err
	}
	result.Base = &baseDegree

	return nil
}

func (c SyllableChordConverter) getTendency(v *op.ScaleNote) (op.Accidental, error) {
	index, err := c.scale.GetNoteIndexByName(v.Name)
	if err != nil {
		return op.UnknownAccidental, errorx.Invalid("%w: get tendency %v on %v", err, v.Name, c.scale.Key)
	}
	t := c.scale.Notes[index].Accidental.Tendency(v.Accidental)
	if t == op.UnknownAccidental {
		return op.UnknownAccidental, errorx.Invalid("Get tendency from notes[%d] %#v and %#v",
			index, c.scale.Notes[index], v)
	}
	return t, nil
}

func (SyllableChordConverter) newScaleNote(v *ast.ChordDegree) (*op.ScaleNote, error) {
	name := note.NewName(v.Degree.Value())
	if name == note.UnknownName {
		return nil, errorx.Invalid("UnknownName %v", v.Degree)
	}
	accidental := op.Natural
	if x := v.Accidental; x != nil {
		accidental = op.NewAccidental(x.Value())
	}
	return &op.ScaleNote{
		Name:       name,
		Accidental: accidental,
	}, nil
}

// DegreeChordConverter converts AST contains only degrees.
type DegreeChordConverter struct{}

func (c DegreeChordConverter) Convert(v *ast.Chord) (*input.Chord, error) {
	var result input.Chord

	degree, err := c.convertDegree(v.Degree)
	if err != nil {
		return nil, fmt.Errorf("%w: ChordDegree", err)
	}
	result.Degree = degree

	if x := v.Symbol; x != nil {
		result.Chord = x.Symbol.Value()
	}

	if x := v.Base; x != nil {
		d, err := c.convertDegree(x.Degree)
		if err != nil {
			return nil, fmt.Errorf("%w: ChordBase", err)
		}
		result.Base = &d
	}

	return &result, nil
}

func (DegreeChordConverter) convertDegree(v *ast.ChordDegree) (note.Degree, error) {
	s := v.Degree.Value()
	if x := v.Accidental; x != nil {
		s += x.Value()
	}
	d, err := note.ParseDegree(s)
	if err != nil {
		return d, fmt.Errorf("%w: Degree is not uint %v", err, v.Degree)
	}
	return d, nil
}
